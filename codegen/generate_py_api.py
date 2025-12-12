#!/usr/bin/env python3
"""
CTP C API 转 Python ctypes 包装代码生成器

功能：
- 解析 C API 头文件（ctptrader_c_api.h, ctpmd_c_api.h）
- 解析 CTP 结构体和数据类型定义
- 生成使用 ctypes 的 Python 包装代码
- 支持多实例，使用 userData 机制

用法：
    python3 generate_py_api.py --input ../csrc --struct ../ctpapi/linux --output ../pyctp
"""

import re
import os
import argparse
import subprocess
import shutil
from pathlib import Path
from dataclasses import dataclass, field
from typing import List, Optional, Dict, Tuple, Set


# ========== 数据结构定义 ==========
# 复用 generate_go_api.py 中的数据结构定义

@dataclass
class CParam:
    """C 函数参数/结构体字段"""
    type: str           # 参数类型（如 "MdApiHandle", "const char*"）
    name: str           # 参数名
    is_pointer: bool = False
    is_const: bool = False
    is_array: bool = False
    array_size: int = 0
    comment: str = ""   # 字段注释


@dataclass
class CFunction:
    """C 函数定义"""
    name: str                   # 函数名（如 "MdCreateFtdcMdApi"）
    return_type: str            # 返回类型
    params: List[CParam]        # 参数列表
    comment: str = ""           # 注释


@dataclass
class CallbackType:
    """回调函数类型"""
    name: str                   # 类型名（如 "MdOnFrontConnectedCallback"）
    params: List[CParam]        # 参数列表（第一个通常是 void* userData）
    comment: str = ""
    py_method_name: str = ""    # Python 方法名（如 "OnFrontConnected"）


@dataclass
class CTypedef:
    """C 类型定义 (typedef)"""
    name: str           # 类型名 (如 TThostFtdcBrokerIDType)
    base_type: str      # 基础类型 (如 char)
    size: int = 0       # 数组大小（如果是数组类型，如 [11]）
    comment: str = ""   # 注释 (如 "经纪公司代码类型")


@dataclass
class CEnum:
    """C 枚举定义"""
    name: str                       # 枚举名 (如 THOST_TE_RESUME_TYPE)
    values: List[Tuple[str, int]]   # (值名, 值) 列表
    comment: str = ""               # 注释


@dataclass
class CDefine:
    """C #define 常量定义"""
    name: str           # 常量名 (如 THOST_FTDC_EXP_Normal)
    value: str          # 值 (如 '0')
    comment: str = ""   # 注释 (如 "正常")
    type_name: str = "" # 关联的类型名 (如 TThostFtdcExchangePropertyType)


@dataclass 
class CStruct:
    """C 结构体定义"""
    name: str                   # 结构体名 (如 CThostFtdcAccountPropertyField)
    fields: List[CParam]        # 字段列表
    comment: str = ""           # 结构体注释


# ========== C 类型到 Python ctypes 类型映射 ==========

# CTP 基础数据类型映射
CTP_TYPE_MAP = {
    # C 基础类型 -> ctypes 类型
    'char': 'c_char',
    'int': 'c_int32',
    'short': 'c_int16',
    'double': 'c_double',
    'float': 'c_float',
    'bool': 'c_bool',
    'void': '',
    
    # CTP 特定类型会在解析时动态添加
}


# ========== 解析函数 ==========
# 复用 generate_go_api.py 中的解析函数

def parse_datatype_header(header_path: Path) -> Tuple[Dict[str, CTypedef], Dict[str, CEnum], Dict[str, List[CDefine]]]:
    """
    解析 ThostFtdcUserApiDataType.h 获取类型定义、枚举定义和 #define 常量
    
    该头文件包含：
    1. 枚举定义 (如 enum THOST_TE_RESUME_TYPE { ... })
    2. 类型定义 (如 typedef char TThostFtdcBrokerIDType[11];)
    3. #define 常量 (如 #define THOST_FTDC_EXP_Normal '0')
    
    返回:
        typedefs: 类型名 -> CTypedef 的映射
        enums: 枚举名 -> CEnum 的映射
        defines: 类型名 -> [CDefine] 的映射（每个类型关联的常量列表）
    """
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    typedefs = {}
    enums = {}
    defines = {}  # 类型名 -> [CDefine]
    
    # ===== 解析枚举定义 =====
    # 格式: enum THOST_TE_RESUME_TYPE { THOST_TERT_RESTART = 0, THOST_TERT_RESUME, ... };
    enum_pattern = r'enum\s+(\w+)\s*\{([^}]+)\}'
    for m in re.finditer(enum_pattern, content, re.DOTALL):
        enum_name = m.group(1)
        enum_body = m.group(2)
        
        # 提取枚举注释（查找前面的 /// 注释）
        start = m.start()
        preceding = content[:start]
        comment = extract_comment(preceding)
        
        # 解析枚举值
        values = []
        current_value = 0
        for line in enum_body.split('\n'):
            line = line.strip()
            if not line or line.startswith('//'):
                continue
            
            # 移除末尾的逗号
            line = line.rstrip(',')
            if '=' in line:
                # 格式: THOST_TERT_RESTART = 0
                parts = line.split('=')
                name = parts[0].strip()
                value_str = parts[1].strip()
                try:
                    current_value = int(value_str)
                except ValueError:
                    pass
                values.append((name, current_value))
            elif line:
                # 无显式值，使用前一个值+1
                values.append((line, current_value))
            current_value += 1
        
        enums[enum_name] = CEnum(
            name=enum_name,
            values=values,
            comment=comment
        )
    
    # ===== 解析 #define 常量和关联的 typedef =====
    # 分块解析：每个 typedef 前面的 #define 都属于该 typedef
    # 格式: 
    #   ///注释
    #   #define THOST_FTDC_XXX_Yyy 'value'
    #   typedef char TThostFtdcXxxType;
    
    # 先找到所有 typedef 的位置
    typedef_positions = []
    typedef_pattern = r'typedef\s+(\w+)\s+(\w+)(?:\[(\d+)\])?\s*;'
    for m in re.finditer(typedef_pattern, content):
        typedef_positions.append((m.start(), m.end(), m))
    
    # 对于每个 typedef，查找它前面的 #define
    for i, (start, end, m) in enumerate(typedef_positions):
        base_type = m.group(1)
        type_name = m.group(2)
        array_size = int(m.group(3)) if m.group(3) else 0
        
        # 确定搜索范围的起点（上一个 typedef 的结束位置或文件开头）
        search_start = typedef_positions[i-1][1] if i > 0 else 0
        preceding_block = content[search_start:start]
        
        # 提取类型注释（从注释块中提取描述）
        comment = extract_comment(preceding_block)
        
        typedefs[type_name] = CTypedef(
            name=type_name,
            base_type=base_type,
            size=array_size,
            comment=comment
        )
        
        # 解析该块中的所有 #define
        # 格式: #define THOST_FTDC_XXX_Yyy 'value' 或 #define THOST_FTDC_XXX_Yyy "value"
        define_pattern = r"///([^\n]*)\n#define\s+(\w+)\s+['\"]?([^'\"\s]+)['\"]?"
        type_defines = []
        
        for dm in re.finditer(define_pattern, preceding_block):
            define_comment = dm.group(1).strip()
            define_name = dm.group(2)
            define_value = dm.group(3)
            
            type_defines.append(CDefine(
                name=define_name,
                value=define_value,
                comment=define_comment,
                type_name=type_name
            ))
        
        if type_defines:
            defines[type_name] = type_defines
    
    return typedefs, enums, defines


def parse_struct_header(header_path: Path, typedefs: Dict[str, CTypedef]) -> Dict[str, CStruct]:
    """
    解析 ThostFtdcUserApiStruct.h 获取结构体定义
    
    该头文件包含所有 CTP 业务结构体定义，如：
    struct CThostFtdcReqUserLoginField {
        TThostFtdcBrokerIDType  BrokerID;
        TThostFtdcUserIDType    UserID;
        ...
    };
    
    参数:
        header_path: 头文件路径
        typedefs: 已解析的类型定义（用于判断字段是否为数组）
        
    返回:
        结构体名 -> CStruct 的映射
    """
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    structs = {}
    
    # 匹配结构体定义: struct CThostFtdcXxxField { ... };
    struct_pattern = r'struct\s+(\w+)\s*\{([^}]+)\}'
    
    for m in re.finditer(struct_pattern, content, re.DOTALL):
        struct_name = m.group(1)
        body = m.group(2)
        
        # 提取结构体注释（查找前面的 /// 注释）
        start = m.start()
        preceding = content[:start]
        struct_comment = ""
        preceding_lines = preceding.split('\n')
        for line in reversed(preceding_lines):
            line = line.strip()
            if line.startswith('///'):
                struct_comment = line[3:].strip()
                break
            elif line and not line.startswith('//'):
                break
        
        fields = []
        
        # 解析字段: TThostFtdcTraderIDType TraderID;
        field_pattern = r'(\w+)\s+(\w+)\s*;'
        
        lines = body.split('\n')
        for i, line in enumerate(lines):
            line = line.strip()
            if not line or line.startswith('//'):
                continue
            
            fm = re.match(field_pattern, line)
            if fm:
                field_type = fm.group(1)
                field_name = fm.group(2)
                
                # 提取字段注释（查找前面的 /// 注释）
                field_comment = ""
                if i > 0:
                    prev_line = lines[i-1].strip()
                    if prev_line.startswith('///'):
                        field_comment = prev_line[3:].strip()
                
                # 判断是否是数组类型（根据 typedef 中的定义）
                is_array = False
                array_size = 0
                if field_type in typedefs:
                    td = typedefs[field_type]
                    if td.size > 0:
                        is_array = True
                        array_size = td.size
                
                fields.append(CParam(
                    type=field_type,
                    name=field_name,
                    is_pointer=False,
                    is_const=False,
                    is_array=is_array,
                    array_size=array_size,
                    comment=field_comment
                ))
        
        structs[struct_name] = CStruct(
            name=struct_name,
            fields=fields,
            comment=struct_comment
        )
    
    return structs


def parse_c_header(header_path: Path) -> Tuple[List[CFunction], List[CallbackType]]:
    """解析 C API 头文件，返回函数列表和回调类型列表"""
    
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    functions = []
    callbacks = []
    
    # 提取回调函数类型定义
    # typedef void (*MdOnFrontConnectedCallback)(void* userData);
    callback_pattern = r'typedef\s+void\s+\(\*(\w+)\)\s*\(([^)]*)\)\s*;'
    for m in re.finditer(callback_pattern, content):
        callback_name = m.group(1)
        param_str = m.group(2)
        
        # 提取注释（前面的行）
        callback_start = m.start()
        preceding = content[:callback_start]
        comment = extract_comment(preceding)
        
        params = parse_params(param_str)
        
        # 提取 Python 方法名
        py_method_name = extract_py_method_name(callback_name)
        
        callbacks.append(CallbackType(
            name=callback_name,
            params=params,
            comment=comment,
            py_method_name=py_method_name
        ))
    
    # 提取函数声明
    # CTP_API MdApiHandle MdCreateFtdcMdApi(const char* pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast);
    function_pattern = r'CTP_API\s+([\w\s]+\s*\*?)\s+(\w+)\s*\(([^)]*)\)\s*;'
    for m in re.finditer(function_pattern, content, re.MULTILINE):
        return_type = m.group(1).strip()
        func_name = m.group(2)
        param_str = m.group(3)
        
        # 提取注释
        func_start = m.start()
        preceding = content[:func_start]
        comment = extract_comment(preceding)
        
        params = parse_params(param_str)
        functions.append(CFunction(
            name=func_name,
            return_type=return_type,
            params=params,
            comment=comment
        ))
    
    return functions, callbacks


def parse_params(param_str: str) -> List[CParam]:
    """解析参数列表字符串"""
    params = []
    if not param_str or param_str.strip() == "" or param_str.strip() == "void":
        return params
    
    # 分割参数
    param_list = []
    depth = 0
    current = ""
    for char in param_str:
        if char == '(':
            depth += 1
        elif char == ')':
            depth -= 1
        elif char == ',' and depth == 0:
            param_list.append(current.strip())
            current = ""
            continue
        current += char
    if current.strip():
        param_list.append(current.strip())
    
    for param in param_list:
        param = param.strip()
        if not param:
            continue
        
        # 解析类型和名称
        is_const = 'const' in param
        is_pointer = '*' in param
        is_array = '[]' in param or '**' in param
        
        # 移除 const
        param_clean = param.replace('const', '').strip()
        
        # 处理数组类型 (char* ppInstrumentID[] 或 char** ppInstrumentID)
        if '[]' in param_clean:
            # 对于 char* ppInstrumentID[] 格式，参数名在 [] 之前
            # 先找到 [ 的位置
            bracket_start = param_clean.index('[')
            # 提取 [ 之前的部分
            name_part = param_clean[:bracket_start].strip()
            
            # 从 name_part 中分离类型和参数名
            # char* ppInstrumentID 格式
            parts = name_part.split()
            if len(parts) >= 2:
                # 最后一个单词是参数名
                name = parts[-1]
                type_part = ' '.join(parts[:-1])
            elif len(parts) == 1:
                # 只有类型，没有参数名
                type_part = parts[0]
                name = ''
            else:
                type_part = ''
                name = ''
            
            # 提取类型
            if '*' in type_part:
                # char* [] 格式
                param_type = 'char'
                params.append(CParam(
                    type=param_type,
                    name=name,
                    is_pointer=True,
                    is_const=is_const,
                    is_array=True
                ))
            else:
                # 其他数组类型
                param_type = type_part if type_part else 'char'
                params.append(CParam(
                    type=param_type,
                    name=name,
                    is_pointer=False,
                    is_const=is_const,
                    is_array=True
                ))
            continue
        elif '**' in param_clean:
            # 处理双指针 (char** ppInstrumentID)
            parts = param_clean.replace('**', ' ** ').split()
            parts = [p for p in parts if p and p != '*']
            if len(parts) >= 2:
                name = parts[-1]
                param_type = parts[0]
                params.append(CParam(
                    type=param_type,
                    name=name,
                    is_pointer=True,
                    is_const=is_const,
                    is_array=True
                ))
            continue
        
        # 处理单指针
        if '*' in param_clean:
            parts = param_clean.replace('*', ' * ').split()
            parts = [p for p in parts if p and p != '*']
            if len(parts) >= 2:
                name = parts[-1]
                param_type = ' '.join(parts[:-1]).replace(' * ', '*').replace('* ', '*').strip()
                if '*' not in param_type:
                    param_type = parts[0]
                params.append(CParam(
                    type=param_type.replace('*', '').strip(),
                    name=name,
                    is_pointer=True,
                    is_const=is_const
                ))
            elif len(parts) == 1:
                # 只有类型没有名字
                params.append(CParam(
                    type=parts[0].replace('*', '').strip(),
                    name='',
                    is_pointer=True,
                    is_const=is_const
                ))
            continue
        
        # 非指针类型
        parts = param_clean.split()
        if len(parts) >= 2:
            name = parts[-1]
            param_type = ' '.join(parts[:-1])
            params.append(CParam(
                type=param_type,
                name=name,
                is_pointer=False,
                is_const=is_const
            ))
        elif len(parts) == 1:
            params.append(CParam(
                type=parts[0],
                name='',
                is_pointer=False,
                is_const=is_const
            ))
    
    return params


def extract_comment(preceding: str) -> str:
    """
    从前置文本中提取最近的注释
    
    CTP 头文件的注释格式通常为：
    /////////////////////////////////////////////////////////////////////////
    ///TFtdcBrokerIDType是一个经纪公司代码类型
    /////////////////////////////////////////////////////////////////////////
    typedef char TThostFtdcBrokerIDType[11];
    
    提取后返回完整的注释内容
    """
    lines = preceding.split('\n')
    comment_lines = []
    
    for line in reversed(lines):
        line = line.strip()
        if line.startswith('///'):
            # 提取 /// 后面的内容
            comment_text = line[3:].strip()
            comment_lines.insert(0, comment_text)
        elif line.startswith('//'):
            comment_text = line[2:].strip()
            comment_lines.insert(0, comment_text)
        elif line and not line.startswith('//'):
            # 遇到非注释行，停止
            break
    
    return ' '.join(comment_lines) if comment_lines else ""


def extract_py_method_name(callback_name: str) -> str:
    """从回调类型名提取 Python 方法名"""
    name = callback_name
    # 移除前缀 (Trader/Md)
    if name.startswith("TraderOn"):
        name = name[6:]  # 移除 "Trader"，保留 On
    elif name.startswith("MdOn"):
        name = name[2:]   # 移除 "Md"，保留 On
    
    if name.endswith("Callback"):
        name = name[:-8]  # 移除 "Callback"
    
    return name


# ========== Python ctypes 类型转换 ==========

def c_type_to_py_ctypes_type(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef], is_array: bool = False) -> str:
    """
    将 C 类型转换为 Python ctypes 类型
    
    类型映射规则：
    1. 句柄类型 (XxxHandle) -> c_void_p
    2. CTP Field 结构体指针 (*CThostFtdcXxxField) -> POINTER(CThostFtdcXxxField)
    3. CTP typedef 类型：
       - 数组类型 (typedef char TThostFtdcBrokerIDType[11]) -> c_char * 11
       - 简单类型 (typedef int TThostFtdcVolumeType) -> c_int32
       - 单字符类型 (typedef char TThostFtdcXxxType) -> c_char
    4. 基础 C 类型 -> 对应 ctypes 类型
    
    参数:
        c_type: C 类型名
        is_pointer: 是否是指针类型
        typedefs: 已解析的 typedef 映射
        is_array: 是否是数组（用于特殊处理 char*[]）
        
    返回:
        对应的 Python ctypes 类型字符串
    """
    c_type = c_type.strip().replace('const', '').replace('*', '').strip()
    
    # 句柄类型: MdApiHandle, TraderApiHandle, etc.
    if c_type.endswith('Handle'):
        return 'ctypes.c_void_p'
    
    # 枚举类型: THOST_TE_RESUME_TYPE 等
    if c_type.startswith('THOST_'):
        if is_pointer:
            return f'ctypes.POINTER({c_type})'
        return c_type
    
    # CTP Field 结构体类型
    if c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
        if is_pointer:
            return f'ctypes.POINTER({c_type})'
        return c_type
    
    # SpiCallbacks 结构体类型（MdSpiCallbacks, TraderSpiCallbacks）
    if c_type.endswith('SpiCallbacks'):
        if is_pointer:
            return f'ctypes.POINTER({c_type})'
        return c_type
    
    # CTP typedef 类型（如 TThostFtdcBrokerIDType）
    if c_type in typedefs:
        td = typedefs[c_type]
        if td.size > 0:
            # 数组类型: typedef char TThostFtdcBrokerIDType[11] -> ctypes.c_char * 11
            base_ctypes = CTP_TYPE_MAP.get(td.base_type, f'c_{td.base_type}')
            return f'ctypes.{base_ctypes} * {td.size}'
        else:
            # 非数组类型: typedef int TThostFtdcVolumeType -> ctypes.c_int32
            ctypes_type = CTP_TYPE_MAP.get(td.base_type, f'c_{td.base_type}')
            return f'ctypes.{ctypes_type}'
    
    # 基础 C 类型
    if is_array and is_pointer and c_type == 'char':
        # char* [] 格式（字符串数组），返回 ctypes.POINTER(ctypes.c_char_p)
        return 'ctypes.POINTER(ctypes.c_char_p)'
    elif is_pointer:
        if c_type == 'void':
            return 'ctypes.c_void_p'
        elif c_type == 'char':
            return 'ctypes.c_char_p'    # char* -> ctypes.c_char_p (C 字符串)
        else:
            # 检查是否是枚举类型
            if c_type.startswith('THOST_'):
                return f'ctypes.POINTER({c_type})'
            ctypes_type = CTP_TYPE_MAP.get(c_type, f'c_{c_type}')
            return f'ctypes.POINTER(ctypes.{ctypes_type})'
    else:
        if c_type == 'void':
            return ''
        # 检查是否是枚举类型
        if c_type.startswith('THOST_'):
            return c_type
        ctypes_type = CTP_TYPE_MAP.get(c_type, f'c_{c_type}')
        return f'ctypes.{ctypes_type}'


def c_type_to_py_callback_param(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef]) -> str:
    """将 C 类型转换为 Python 回调参数类型（用于接口定义）"""
    c_type = c_type.strip().replace('const', '').strip()
    
    # 指针类型
    if is_pointer:
        if c_type == 'void':
            return 'ctypes.c_void_p'
        elif c_type == 'char':
            return 'ctypes.c_char_p'
        elif c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
            return f'ctypes.POINTER({c_type})'
        else:
            ctypes_type = CTP_TYPE_MAP.get(c_type, f'c_{c_type}')
            return f'ctypes.POINTER(ctypes.{ctypes_type})'
    else:
        if c_type == 'void':
            return ''
        elif c_type == 'int':
            return 'ctypes.c_int32'
        elif c_type == 'bool':
            return 'ctypes.c_bool'
        else:
            ctypes_type = CTP_TYPE_MAP.get(c_type, f'c_{c_type}')
            return f'ctypes.{ctypes_type}'


# ========== 代码生成 ==========

def generate_utils_py() -> str:
    """生成 utils.py"""
    return '''"""
CTP Python 包装工具函数

此文件由代码生成器自动生成，请勿手动修改
"""

import ctypes
from typing import List, Optional


def c_string(s: str) -> Optional[ctypes.c_char_p]:
    """将 Python 字符串转换为 C 字符串（c_char_p）"""
    if s is None or s == "":
        return None
    return ctypes.c_char_p(s.encode('utf-8'))


def c_string_array(ss: List[str]) -> tuple:
    """
    将 Python 字符串列表转换为 C 字符串数组
    
    返回:
        (POINTER(c_char_p), List[bytes]): 字符串指针数组和底层数据（需要保持引用防止 GC）
    """
    if not ss:
        return None, None
    
    # 创建字节数组保存字符串数据
    data = [s.encode('utf-8') for s in ss]
    # 创建 c_char_p 数组
    arr = (ctypes.c_char_p * len(ss))(*data)
    return ctypes.cast(arr, ctypes.POINTER(ctypes.c_char_p)), data


def go_string(ptr: ctypes.c_char_p) -> str:
    """将 C 字符串（c_char_p）转换为 Python 字符串"""
    if ptr is None:
        return ""
    return ptr.decode('utf-8')


def bytes_to_string(b: bytes) -> str:
    """将固定长度字节数组转换为字符串（去除尾部的 null）"""
    if b is None:
        return ""
    try:
        null_pos = b.find(b'\\0')
        if null_pos >= 0:
            return b[:null_pos].decode('utf-8', errors='ignore')
        return b.decode('utf-8', errors='ignore')
    except:
        return ""


def gb18030(b: bytes) -> str:
    """将 GB18030 编码的字节切片转换为 UTF-8 字符串"""
    if b is None:
        return ""
    try:
        null_pos = b.find(b'\\0')
        if null_pos >= 0:
            b = b[:null_pos]
        return b.decode('gb18030', errors='ignore')
    except:
        return ""


def string_to_bytes(s: str, size: int) -> bytes:
    """将字符串复制到固定长度字节数组"""
    if s is None:
        s = ""
    b = s.encode('utf-8')[:size]
    return b + b'\\0' * (size - len(b))


def bool_to_int(b: bool) -> int:
    """将 bool 转换为 int（C 风格）"""
    return 1 if b else 0


def int_to_bool(i: int) -> bool:
    """将 int 转换为 bool（C 风格）"""
    return i != 0
'''


def generate_datatype_py(typedefs: Dict[str, CTypedef], enums: Dict[str, CEnum], defines: Dict[str, List[CDefine]] = None) -> str:
    """
    生成 datatype.py - 完整的 CTP 数据类型定义
    
    将 ThostFtdcUserApiDataType.h 翻译成 Python ctypes 类型：
    
    1. enum -> class XXX(IntEnum) 或常量
    2. typedef char XXX[N] -> XXX = c_char * N
    3. typedef char XXX (单字符) + #define -> XXX = c_char + 常量
    4. typedef int XXX -> XXX = c_int32
    5. typedef short XXX -> XXX = c_int16
    6. typedef double XXX -> XXX = c_double
    """
    if defines is None:
        defines = {}
    
    lines = []
    lines.append('"""')
    lines.append('CTP 数据类型定义')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 数据类型定义 - 来自 ThostFtdcUserApiDataType.h')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from enum import IntEnum')
    lines.append('')
    
    # ========== 1. 生成 enum 类型 ==========
    if enums:
        lines.append('# ========== 枚举类型 ==========')
        lines.append('')
        
        for enum_name, enum in sorted(enums.items()):
            # 提取注释
            description = _extract_description(enum.comment)
            
            # 特殊处理 THOST_TE_RESUME_TYPE
            if enum_name == 'THOST_TE_RESUME_TYPE':
                description = '订阅类型'
            
            if description:
                lines.append(f'# {enum_name} {description}')
            else:
                lines.append(f'# {enum_name}')
            
            lines.append(f'class {enum_name}(IntEnum):')
            if description:
                lines.append(f'    """{description}"""')
            
            if enum.values:
                for value_name, value in enum.values:
                    value_comment = ""
                    if enum_name == 'THOST_TE_RESUME_TYPE':
                        value_comments = {
                            'THOST_TERT_RESTART': '从本交易日开始重传',
                            'THOST_TERT_RESUME': '从上次收到的续传',
                            'THOST_TERT_QUICK': '只传送登录后的流内容',
                            'THOST_TERT_NONE': '不传送'
                        }
                        value_comment = value_comments.get(value_name, "")
                    
                    if value_comment:
                        lines.append(f'    {value_name} = {value}  # {value_comment}')
                    else:
                        lines.append(f'    {value_name} = {value}')
            else:
                lines.append('    pass')
            lines.append('')
    
    # ========== 2. 生成 typedef 类型 ==========
    if typedefs:
        lines.append('# ========== 类型定义 ==========')
        lines.append('')
        
        # 按类型分组：字符数组、单字符（有常量）、单字符（无常量）、整数、短整数、浮点
        char_array_types = []      # typedef char XXX[N]
        char_enum_types = []       # typedef char XXX (有 #define)
        char_simple_types = []     # typedef char XXX (无 #define)
        int_types = []             # typedef int XXX
        short_types = []           # typedef short XXX
        double_types = []          # typedef double XXX
        
        for type_name, td in sorted(typedefs.items()):
            if td.base_type == 'char':
                if td.size > 0:
                    char_array_types.append((type_name, td))
                elif type_name in defines:
                    char_enum_types.append((type_name, td))
                else:
                    char_simple_types.append((type_name, td))
            elif td.base_type == 'int':
                int_types.append((type_name, td))
            elif td.base_type == 'short':
                short_types.append((type_name, td))
            elif td.base_type == 'double':
                double_types.append((type_name, td))
        
        # 2.1 字符数组类型
        if char_array_types:
            lines.append('# ----- 字符串类型 -----')
            lines.append('')
            for type_name, td in char_array_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'# {type_name} {description}')
                lines.append(f'{type_name} = ctypes.c_char * {td.size}')
                lines.append('')
        
        # 2.2 整数类型
        if int_types:
            lines.append('# ----- 整数类型 -----')
            lines.append('')
            for type_name, td in int_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'# {type_name} {description}')
                lines.append(f'{type_name} = ctypes.c_int32')
                lines.append('')
        
        # 2.3 短整数类型
        if short_types:
            lines.append('# ----- 短整数类型 -----')
            lines.append('')
            for type_name, td in short_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'# {type_name} {description}')
                lines.append(f'{type_name} = ctypes.c_int16')
                lines.append('')
        
        # 2.4 浮点类型
        if double_types:
            lines.append('# ----- 浮点类型 -----')
            lines.append('')
            for type_name, td in double_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'# {type_name} {description}')
                lines.append(f'{type_name} = ctypes.c_double')
                lines.append('')
        
        # 2.5 单字符枚举类型（有 #define 常量）
        if char_enum_types:
            lines.append('# ----- 字符枚举类型 -----')
            lines.append('')
            for type_name, td in char_enum_types:
                description = _extract_description(td.comment)
                type_defines = defines.get(type_name, [])
                
                # 检查常量值类型：
                # - 单字符（如 '0', 'A'）-> c_char 类型
                # - 多字符（如 '102001'）-> str 类型（C 多字符字面量，实际是字符串标识符）
                has_multi_char = any(len(d.value) > 1 for d in type_defines)
                
                if description:
                    lines.append(f'# {type_name} {description}')
                
                if has_multi_char:
                    # 多字符常量（如交易代码 '102001'），使用 str 类型
                    lines.append(f'{type_name} = str')
                else:
                    # 单字符常量，使用 bytes 类型（单字节）
                    lines.append(f'{type_name} = ctypes.c_char')
                lines.append('')
                
                # 生成关联的常量
                if type_defines:
                    for d in type_defines:
                        value = d.value
                        if has_multi_char:
                            # 字符串类型，使用双引号
                            value_str = f'"{value}"'
                        else:
                            # 单字符，使用字节字面量
                            value_str = f"b'{value}'"
                        
                        if d.comment:
                            lines.append(f'{d.name} = {value_str}  # {d.comment}')
                        else:
                            lines.append(f'{d.name} = {value_str}')
                    lines.append('')
        
        # 2.6 单字符简单类型（无 #define 常量）
        if char_simple_types:
            lines.append('# ----- 单字符类型 -----')
            lines.append('')
            for type_name, td in char_simple_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'# {type_name} {description}')
                lines.append(f'{type_name} = ctypes.c_char')
                lines.append('')
    
    return '\n'.join(lines)


def _extract_description(comment: str) -> str:
    """从注释中提取描述文本"""
    if not comment:
        return ""
    
    # 尝试从 "TFtdcXxxType是一个XXX类型" 格式提取
    match = re.search(r'是一个(.+?)类型', comment)
    if match:
        return match.group(1) + '类型'
    
    # 清理注释，移除多余的斜杠
    description = comment.replace('///', '').replace('//', '').strip()
    if description.startswith('/'):
        description = description[1:].strip()
    
    return description


def generate_struct_py(structs: Dict[str, CStruct], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 struct.py - CTP 业务结构体定义
    
    生成格式示例（参考 ThostFtdcUserApiStruct.h）：
    # CThostFtdcReqUserLoginField 用户登录请求
    class CThostFtdcReqUserLoginField(ctypes.Structure):
        _fields_ = [
            ("TradingDay", TThostFtdcDateType),     # 交易日
            ("BrokerID", TThostFtdcBrokerIDType),   # 经纪公司代码
            ("UserID", TThostFtdcUserIDType),       # 用户代码
            ...
        ]
    
    字段类型直接使用 CTP 原始类型名称，因为类型已在 datatype.py 中定义
    """
    lines = []
    lines.append('"""')
    lines.append('CTP 结构体定义')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 结构体定义')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from .datatype import *')
    lines.append('')
    
    # 生成结构体
    lines.append('# ========== CTP 结构体 ==========')
    lines.append('')
    
    for struct_name, struct in sorted(structs.items()):
        # 提取并清理结构体注释
        struct_comment = struct.comment.strip()
        if struct_comment:
            # 清理注释，移除多余的斜杠和空格
            struct_comment = struct_comment.replace('///', '').strip()
            if struct_comment.startswith('/'):
                struct_comment = struct_comment[1:].strip()
            # 生成格式: # CThostFtdcXxxField 结构体描述
            if struct_comment:
                lines.append(f'# {struct_name} {struct_comment}')
        
        lines.append(f'class {struct_name}(ctypes.Structure):')
        if struct_comment:
            lines.append(f'    """{struct_comment}"""')
        lines.append('    _fields_ = [')
        
        for field in struct.fields:
            # 直接使用 CTP 原始类型名称（类型已在 datatype.py 中定义）
            field_type = field.type
            
            # 获取字段注释
            field_comment = ""
            if field.comment:
                field_comment = f'  # {field.comment}'
            
            lines.append(f'        ("{field.name}", {field_type}),{field_comment}')
        
        lines.append('    ]')
        lines.append('')
    
    return '\n'.join(lines)


def generate_loader_py() -> str:
    """
    生成 loader.py - 动态库加载
    
    包含功能：
    1. get_system_lib_paths() - 获取系统库路径列表
    2. default_lib_paths - 默认库搜索路径
    3. load_ctp_library() - 加载 CTP C 包装库
    4. get_md_lib_handle() / get_trader_lib_handle() - 获取库句柄
    5. auto_load_library() - 自动加载库（只加载一次）
    """
    return '''"""
CTP Python 动态库加载

此文件由代码生成器自动生成，请勿手动修改
CTP 动态库加载
"""

import ctypes
import os
import sys
from pathlib import Path
from typing import Optional


_md_lib: Optional[ctypes.CDLL] = None
_trader_lib: Optional[ctypes.CDLL] = None
_load_error: Optional[Exception] = None


def get_system_lib_paths() -> list:
    """根据平台返回系统库路径列表"""
    paths = []
    
    # 从系统库路径环境变量中读取
    if sys.platform == "linux":
        # LD_LIBRARY_PATH 是 Linux 的标准库路径环境变量
        ld_path = os.getenv("LD_LIBRARY_PATH")
        if ld_path:
            paths.extend([p for p in ld_path.split(":") if p])
    elif sys.platform == "darwin":
        # DYLD_LIBRARY_PATH 是 macOS 的库路径环境变量
        dyld_path = os.getenv("DYLD_LIBRARY_PATH")
        if dyld_path:
            paths.extend([p for p in dyld_path.split(":") if p])
    elif sys.platform == "win32":
        # Windows 使用 PATH 环境变量，但通常系统会自动搜索
        pass
    
    # 添加标准系统路径
    if sys.platform == "linux":
        paths.extend([
            "/usr/local/lib",            # 用户安装的库
            "/usr/lib",                  # 系统库
            "/usr/lib/x86_64-linux-gnu", # Debian/Ubuntu 64位
            "/usr/lib64",                # 某些发行版的 64 位库路径
            "/opt/ctp/lib",              # CTP 专用安装路径
            "/opt/lib",                  # 通用 opt 路径
        ])
    elif sys.platform == "darwin":
        paths.extend([
            "/usr/local/lib",    # Homebrew (Intel)
            "/opt/homebrew/lib", # Homebrew (Apple Silicon)
            "/opt/local/lib",    # MacPorts
            "/usr/lib",          # 系统库
            "/opt/ctp/lib",      # CTP 专用安装路径
        ])
    elif sys.platform == "win32":
        # Windows 通常通过 PATH 环境变量查找，但也可以添加一些常见路径
        program_files = os.getenv("ProgramFiles")
        program_files_x86 = os.getenv("ProgramFiles(x86)")
        if program_files:
            paths.append(str(Path(program_files) / "CTP" / "lib"))
        if program_files_x86:
            paths.append(str(Path(program_files_x86) / "CTP" / "lib"))
        paths.extend([
            r"C:\\Windows\\System32",  # 系统目录
            r"C:\\CTP\\lib",           # 常见安装路径
        ])
    
    return paths


# 默认库路径列表，按优先级顺序尝试
# 可以通过环境变量 CTP_LIB_PATH 覆盖，环境变量优先级最高
_default_lib_paths = [
    "./libs",              # 当前目录下的 libs
    "../libs",             # 上一层级下的 libs
    "../../libs",          # 上两级目录下的 libs
    "./ctp-wrapper/libs",  # 项目根目录下的 ctp-wrapper/libs
    "../ctp-wrapper/libs", # 上一层级下的 ctp-wrapper/libs
] + get_system_lib_paths()


def load_ctp_library(lib_path: str) -> None:
    """
    从 C 包装库加载（包含回调支持）
    
    Args:
        lib_path: ctp_md_c_api 和 ctp_trader_c_api 库文件所在目录
        
    Raises:
        FileNotFoundError: 库文件不存在
        OSError: 加载库失败
    """
    global _md_lib, _trader_lib
    
    # 根据平台确定库文件名
    if sys.platform == "win32":
        md_lib_name = "ctpmd_c_api.dll"
        trader_lib_name = "ctptrader_c_api.dll"
    elif sys.platform == "linux":
        md_lib_name = "libctpmd_c_api.so"
        trader_lib_name = "libctptrader_c_api.so"
    elif sys.platform == "darwin":
        md_lib_name = "libctpmd_c_api.dylib"
        trader_lib_name = "libctptrader_c_api.dylib"
    else:
        raise OSError(f"Unsupported platform: {sys.platform}")
    
    md_path = Path(lib_path) / md_lib_name
    trader_path = Path(lib_path) / trader_lib_name
    
    # 检查文件是否存在
    if not md_path.exists():
        raise FileNotFoundError(f"md C wrapper library not found: {md_path}")
    if not trader_path.exists():
        raise FileNotFoundError(f"trader C wrapper library not found: {trader_path}")
    
    # 加载行情 C 包装库
    try:
        _md_lib = ctypes.CDLL(str(md_path))
    except OSError as e:
        raise OSError(f"Failed to load md C wrapper library: {e}") from e
    
    # 加载交易 C 包装库
    try:
        _trader_lib = ctypes.CDLL(str(trader_path))
    except OSError as e:
        raise OSError(f"Failed to load trader C wrapper library: {e}") from e


def get_md_lib_handle() -> Optional[ctypes.CDLL]:
    """获取行情库句柄"""
    return _md_lib


def get_trader_lib_handle() -> Optional[ctypes.CDLL]:
    """获取交易库句柄"""
    return _trader_lib


def auto_load_library() -> None:
    """
    自动加载库（只加载一次）
    优先使用环境变量 CTP_LIB_PATH，否则按顺序尝试默认路径列表
    
    Raises:
        FileNotFoundError: 所有路径都找不到库文件
        OSError: 加载库失败
    """
    global _load_error
    
    if _md_lib is not None and _trader_lib is not None:
        return  # 已经加载
    
    if _load_error is not None:
        raise _load_error  # 之前加载失败，直接抛出错误
    
    # 优先使用环境变量
    lib_path = os.getenv("CTP_LIB_PATH")
    if lib_path:
        try:
            load_ctp_library(lib_path)
            return
        except Exception as e:
            _load_error = e
            raise
    
    # 按顺序尝试默认路径列表
    for path in _default_lib_paths:
        try:
            load_ctp_library(path)
            return
        except (FileNotFoundError, OSError):
            continue  # 继续尝试下一个路径
    
    # 所有路径都失败
    _load_error = FileNotFoundError(
        f"Could not find CTP libraries in any of the following paths: {_default_lib_paths}"
    )
    raise _load_error
'''


def generate_md_api_py(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_api.py - 行情 API 封装
    
    包含：
    1. 回调类型定义 (MdOnXxxCallback)
    2. MdSpi 接口（抽象基类）
    3. MdApi 类
    4. C 函数声明和初始化
    5. 实例管理
    6. API 方法和 SPI 方法
    7. SetSpi 方法
    """
    lines = []
    lines.append('"""')
    lines.append('CTP 行情 API 封装')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 行情 API 封装')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('import os')
    lines.append('import threading')
    lines.append('from abc import ABC, abstractmethod')
    lines.append('from typing import Optional, List')
    lines.append('')
    lines.append('from .loader import auto_load_library, get_md_lib_handle')
    lines.append('from .struct import *')
    lines.append('from .utils import *')
    lines.append('')
    
    # 生成回调类型定义
    lines.append('# ========== 回调类型定义 ==========')
    lines.append('')
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成回调函数类型（CFUNCTYPE）
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            callback_params.append(py_type)
        
        param_str = ', '.join(callback_params)
        comment = f'# {cb.name} {cb.comment}' if cb.comment else f'# {cb.name}'
        lines.append(comment)
        lines.append(f'{cb.name} = ctypes.CFUNCTYPE(None, {param_str})')
        lines.append('')
    
    # 生成 MdSpiCallbacks 结构体定义
    lines.append('# ========== 回调结构体定义 ==========')
    lines.append('')
    lines.append('# MdSpiCallbacks 回调结构体（用于批量设置）')
    lines.append('class MdSpiCallbacks(ctypes.Structure):')
    lines.append('    """回调结构体（用于批量设置）"""')
    lines.append('    _fields_ = [')
    lines.append('        ("userData", ctypes.c_void_p),')
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        callback_suffix = cb.py_method_name[2:]  # 移除 "On" 前缀
        lines.append(f'        ("on{callback_suffix}", {cb.name}),')
    lines.append('    ]')
    lines.append('')
    
    # 生成 SPI 接口
    lines.append('# ========== MdSpi 接口 ==========')
    lines.append('')
    lines.append('class MdSpi(ABC):')
    lines.append('    """行情回调接口"""')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成方法签名
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name}: {py_type}')
            else:
                method_params.append(py_type)
        
        param_str = ', '.join(method_params)
        comment = f'    # {cb.comment}' if cb.comment else ''
        lines.append(f'    @abstractmethod')
        lines.append(f'    def {cb.py_method_name}(self, {param_str}):')
        if cb.comment:
            lines.append(f'        """{cb.comment}"""')
        lines.append('        pass')
        lines.append('')
    
    # 生成 MdApi 类
    lines.append('# ========== MdApi 类 ==========')
    lines.append('')
    lines.append('class MdApi:')
    lines.append('    """行情 API 封装"""')
    lines.append('')
    lines.append('    def __init__(self, flow_path: str, using_udp: bool = False, multicast: bool = False):')
    lines.append('        """创建行情 API 实例"""')
    lines.append('        # 自动加载库（如果尚未加载）')
    lines.append('        auto_load_library()')
    lines.append('')
    lines.append('        self._handle: Optional[ctypes.c_void_p] = None')
    lines.append('        self._spi: Optional[MdSpi] = None')
    lines.append('        self._spi_handle: Optional[ctypes.c_void_p] = None')
    lines.append('        self._user_data: int = _register_md_instance(self)')
    lines.append('        self._lock = threading.RLock()')
    lines.append('')
    lines.append('        # 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录')
    lines.append('        # CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的')
    lines.append('        # 所以转换为绝对路径更可靠')
    lines.append('        abs_flow_path = flow_path')
    lines.append('        if not os.path.isabs(flow_path):')
    lines.append('            # 如果是相对路径，转换为基于当前工作目录的绝对路径')
    lines.append('            abs_flow_path = os.path.abspath(flow_path)')
    lines.append('')
    lines.append('        # 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）')
    lines.append('        if abs_flow_path and not abs_flow_path.endswith(os.sep):')
    lines.append('            abs_flow_path += os.sep')
    lines.append('')
    lines.append('        # 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收')
    lines.append('        # CTP API 可能会在后续使用这个路径')
    lines.append('        self._flow_path = abs_flow_path.encode(\'utf-8\') + b\'\\0\'')
    lines.append('')
    lines.append('        # 调用 C 函数创建 API')
    lines.append('        lib = get_md_lib_handle()')
    lines.append('        if lib is None:')
    lines.append('            raise RuntimeError("CTP library not loaded")')
    lines.append('')
    lines.append('        # 获取函数指针')
    lines.append('        func = lib.MdCreateFtdcMdApi')
    lines.append('        func.argtypes = [ctypes.c_char_p, ctypes.c_bool, ctypes.c_bool]')
    lines.append('        func.restype = ctypes.c_void_p')
    lines.append('')
    lines.append('        self._handle = func(self._flow_path, using_udp, multicast)')
    lines.append('        if self._handle is None:')
    lines.append('            raise RuntimeError("Failed to create MdApi")')
    lines.append('')
    
    # 生成 API 方法
    lines.append('    # ========== API 方法 ==========')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Md'):
            continue
        if func.name == 'MdCreateFtdcMdApi':
            continue  # 已经在 __init__ 中处理
        if func.name.startswith('MdSpiSetOn'):
            continue  # 在单独的 SPI 回调设置方法部分生成
        
        # 方法名（移除 Md 前缀）
        method_name = func.name[2:]
        
        # 生成参数
        params = []
        call_args = ['self._handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
            
            # 清理参数名（移除可能的 [] 后缀）
            param_name = p.name.replace('[]', '').strip() if p.name else ''
            
            # 处理字符串参数
            if p.is_array and p.type == 'char' and p.is_pointer:
                # char* [] 或 char** 类型（字符串数组）
                params.append(f'{param_name}: List[str]')
                call_args.append('_PLACEHOLDER_STRING_ARRAY_')
            elif p.type == 'char' and p.is_pointer and not p.is_array:
                params.append(f'{param_name}: str')
                call_args.append(f'c_string({param_name})')
            else:
                if param_name:
                    params.append(f'{param_name}: {py_type}')
                    call_args.append(param_name)
        
        param_str = ', '.join(params)
        
        # 返回类型
        ret_type = c_type_to_py_ctypes_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法
        comment = f'    # {method_name} {func.comment}' if func.comment else f'    # {method_name}'
        lines.append(comment)
        
        # 特殊处理方法
        if method_name == 'Release':
            lines.append(f'    def {method_name}(self):')
            lines.append(f'        """释放 API 实例"""')
            lines.append(f'        with self._lock:')
            lines.append(f'            if self._handle:')
            lines.append(f'                lib = get_md_lib_handle()')
            lines.append(f'                if lib:')
            lines.append(f'                    func = lib.{func.name}')
            lines.append(f'                    func.argtypes = [ctypes.c_void_p]')
            lines.append(f'                    func.restype = None')
            lines.append(f'                    func(self._handle)')
            lines.append(f'                self._handle = None')
            lines.append(f'            _unregister_md_instance(self._user_data)')
        elif method_name == 'GetApiVersion':
            lines.append(f'    def {method_name}(self) -> str:')
            lines.append(f'        """获取 API 版本"""')
            lines.append(f'        lib = get_md_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return ""')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = []')
            lines.append(f'        func.restype = ctypes.c_char_p')
            lines.append(f'        ptr = func()')
            lines.append(f'        return go_string(ptr) if ptr else ""')
        elif method_name == 'GetTradingDay':
            lines.append(f'    def {method_name}(self) -> str:')
            lines.append(f'        """获取交易日"""')
            lines.append(f'        lib = get_md_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return ""')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = [ctypes.c_void_p]')
            lines.append(f'        func.restype = ctypes.c_char_p')
            lines.append(f'        ptr = func(self._handle)')
            lines.append(f'        return go_string(ptr) if ptr else ""')
        elif '_PLACEHOLDER_STRING_ARRAY_' in call_args:
            # 处理字符串数组参数
            lines.append(f'    def {method_name}(self, {param_str}) -> int:')
            lines.append(f'        """{func.comment or method_name}"""')
            # 找到字符串数组参数的位置
            array_param_idx = call_args.index('_PLACEHOLDER_STRING_ARRAY_')
            array_param_name = params[array_param_idx - 1].split(':')[0].strip()  # 获取参数名
            lines.append(f'        if len({array_param_name}) == 0:')
            lines.append(f'            return 0')
            lines.append(f'        # 将字符串数组转换为 C 字符串数组')
            lines.append(f'        ptrs, _ = c_string_array({array_param_name})')
            # 替换占位符
            call_args[array_param_idx] = 'ptrs'
            lines.append(f'        lib = get_md_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return -1')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]')
            lines.append(f'        func.restype = ctypes.c_int32')
            call_str = ', '.join(call_args)
            lines.append(f'        return func({call_str})')
        elif ret_type:
            lines.append(f'    def {method_name}(self, {param_str}) -> {ret_type}:')
            lines.append(f'        """{func.comment or method_name}"""')
            lines.append(f'        lib = get_md_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            raise RuntimeError("CTP library not loaded")')
            lines.append(f'        func = lib.{func.name}')
            # 生成 argtypes
            argtypes = ['ctypes.c_void_p']
            for p in func.params[1:]:
                if p.is_array and p.type == 'char' and p.is_pointer:
                    argtypes.append('ctypes.POINTER(ctypes.c_char_p)')
                else:
                    py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
                    argtypes.append(py_type)
            lines.append(f'        func.argtypes = [{", ".join(argtypes)}]')
            lines.append(f'        func.restype = {ret_type}')
            call_str = ', '.join(call_args)
            lines.append(f'        return func({call_str})')
        else:
            lines.append(f'    def {method_name}(self, {param_str}):')
            lines.append(f'        """{func.comment or method_name}"""')
            lines.append(f'        lib = get_md_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            raise RuntimeError("CTP library not loaded")')
            lines.append(f'        func = lib.{func.name}')
            # 生成 argtypes
            argtypes = ['ctypes.c_void_p']
            for p in func.params[1:]:
                if p.is_array and p.type == 'char' and p.is_pointer:
                    argtypes.append('ctypes.POINTER(ctypes.c_char_p)')
                else:
                    py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
                    argtypes.append(py_type)
            lines.append(f'        func.argtypes = [{", ".join(argtypes)}]')
            lines.append(f'        func.restype = None')
            call_str = ', '.join(call_args)
            lines.append(f'        func({call_str})')
        
        lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('    def set_spi(self, spi: MdSpi):')
    lines.append('        """设置回调接口"""')
    lines.append('        with self._lock:')
    lines.append('            self._spi = spi')
    lines.append('')
    lines.append('            # 如果已有 C SPI 实例，先销毁')
    lines.append('            if self._spi_handle:')
    lines.append('                lib = get_md_lib_handle()')
    lines.append('                if lib:')
    lines.append('                    func = lib.MdSpiDestroy')
    lines.append('                    func.argtypes = [ctypes.c_void_p]')
    lines.append('                    func.restype = None')
    lines.append('                    func(self._spi_handle)')
    lines.append('                self._spi_handle = None')
    lines.append('')
    lines.append('            # 创建新的 C SPI 实例')
    lines.append('            lib = get_md_lib_handle()')
    lines.append('            if lib is None:')
    lines.append('                raise RuntimeError("CTP library not loaded")')
    lines.append('')
    lines.append('            func = lib.MdSpiCreate')
    lines.append('            func.argtypes = [ctypes.c_void_p]')
    lines.append('            func.restype = ctypes.c_void_p')
    lines.append('            self._spi_handle = func(ctypes.c_void_p(self._user_data))')
    lines.append('')
    lines.append('            # 注册所有回调函数到 C SPI')
    
    # 生成每个回调的设置调用
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        callback_suffix = cb.py_method_name[2:]  # 移除 "On" 前缀
        lines.append(f'            _register_md_callback(self._spi_handle, lib, "{callback_suffix}", self._spi, self._user_data)')
    
    lines.append('')
    lines.append('            # 将 C SPI 注册到 API')
    lines.append('            func = lib.MdRegisterSpi')
    lines.append('            func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]')
    lines.append('            func.restype = None')
    lines.append('            func(self._handle, self._spi_handle)')
    lines.append('')
    
    # 生成实例管理
    lines.append('# ========== 实例管理 ==========')
    lines.append('')
    lines.append('_md_instances: dict = {}')
    lines.append('_md_instances_lock = threading.RLock()')
    lines.append('_md_next_id = 1')
    lines.append('')
    lines.append('def _register_md_instance(api: MdApi) -> int:')
    lines.append('    """注册行情 API 实例"""')
    lines.append('    global _md_next_id')
    lines.append('    with _md_instances_lock:')
    lines.append('        instance_id = _md_next_id')
    lines.append('        _md_next_id += 1')
    lines.append('        _md_instances[instance_id] = api')
    lines.append('        return instance_id')
    lines.append('')
    lines.append('def _get_md_instance(user_data: int) -> Optional[MdApi]:')
    lines.append('    """获取行情 API 实例"""')
    lines.append('    with _md_instances_lock:')
    lines.append('        return _md_instances.get(user_data)')
    lines.append('')
    lines.append('def _unregister_md_instance(user_data: int):')
    lines.append('    """注销行情 API 实例"""')
    lines.append('    with _md_instances_lock:')
    lines.append('        _md_instances.pop(user_data, None)')
    lines.append('')
    lines.append('def _register_md_callback(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi: MdSpi, user_data: int):')
    lines.append('    """注册回调函数到 C SPI"""')
    lines.append('    # 实际实现在 md_callbacks.py 中')
    lines.append('    from .md_callbacks import _register_md_callback_impl')
    lines.append('    _register_md_callback_impl(spi_handle, lib, callback_name, spi, user_data)')
    lines.append('')
    
    return '\n'.join(lines)


def generate_trader_api_py(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_api.py - 交易 API 封装
    
    结构与 md_api.py 类似
    """
    lines = []
    lines.append('"""')
    lines.append('CTP 交易 API 封装')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 交易 API 封装')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('import os')
    lines.append('import threading')
    lines.append('from abc import ABC, abstractmethod')
    lines.append('from typing import Optional, List')
    lines.append('')
    lines.append('from .loader import auto_load_library, get_trader_lib_handle')
    lines.append('from .struct import *')
    lines.append('from .utils import *')
    lines.append('')
    
    # 生成回调类型定义
    lines.append('# ========== 回调类型定义 ==========')
    lines.append('')
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成回调函数类型（CFUNCTYPE）
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            callback_params.append(py_type)
        
        param_str = ', '.join(callback_params)
        comment = f'# {cb.name} {cb.comment}' if cb.comment else f'# {cb.name}'
        lines.append(comment)
        lines.append(f'{cb.name} = ctypes.CFUNCTYPE(None, {param_str})')
        lines.append('')
    
    # 生成 TraderSpiCallbacks 结构体定义
    lines.append('# ========== 回调结构体定义 ==========')
    lines.append('')
    lines.append('# TraderSpiCallbacks 回调结构体（用于批量设置）')
    lines.append('class TraderSpiCallbacks(ctypes.Structure):')
    lines.append('    """回调结构体（用于批量设置）"""')
    lines.append('    _fields_ = [')
    lines.append('        ("userData", ctypes.c_void_p),')
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        callback_suffix = cb.py_method_name[2:]  # 移除 "On" 前缀
        lines.append(f'        ("on{callback_suffix}", {cb.name}),')
    lines.append('    ]')
    lines.append('')
    
    # 生成 SPI 接口
    lines.append('# ========== TraderSpi 接口 ==========')
    lines.append('')
    lines.append('class TraderSpi(ABC):')
    lines.append('    """交易回调接口"""')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成方法签名
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name}: {py_type}')
            else:
                method_params.append(py_type)
        
        param_str = ', '.join(method_params)
        comment = f'    # {cb.comment}' if cb.comment else ''
        lines.append(f'    @abstractmethod')
        lines.append(f'    def {cb.py_method_name}(self, {param_str}):')
        if cb.comment:
            lines.append(f'        """{cb.comment}"""')
        lines.append('        pass')
        lines.append('')
    
    # 生成 TraderApi 类
    lines.append('# ========== TraderApi 类 ==========')
    lines.append('')
    lines.append('class TraderApi:')
    lines.append('    """交易 API 封装"""')
    lines.append('')
    lines.append('    def __init__(self, flow_path: str):')
    lines.append('        """创建交易 API 实例"""')
    lines.append('        # 自动加载库（如果尚未加载）')
    lines.append('        auto_load_library()')
    lines.append('')
    lines.append('        self._handle: Optional[ctypes.c_void_p] = None')
    lines.append('        self._spi: Optional[TraderSpi] = None')
    lines.append('        self._spi_handle: Optional[ctypes.c_void_p] = None')
    lines.append('        self._user_data: int = _register_trader_instance(self)')
    lines.append('        self._lock = threading.RLock()')
    lines.append('')
    lines.append('        # 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录')
    lines.append('        # CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的')
    lines.append('        # 所以转换为绝对路径更可靠')
    lines.append('        abs_flow_path = flow_path')
    lines.append('        if not os.path.isabs(flow_path):')
    lines.append('            # 如果是相对路径，转换为基于当前工作目录的绝对路径')
    lines.append('            abs_flow_path = os.path.abspath(flow_path)')
    lines.append('')
    lines.append('        # 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）')
    lines.append('        if abs_flow_path and not abs_flow_path.endswith(os.sep):')
    lines.append('            abs_flow_path += os.sep')
    lines.append('')
    lines.append('        # 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收')
    lines.append('        # CTP API 可能会在后续使用这个路径')
    lines.append('        self._flow_path = abs_flow_path.encode(\'utf-8\') + b\'\\0\'')
    lines.append('')
    lines.append('        # 调用 C 函数创建 API')
    lines.append('        lib = get_trader_lib_handle()')
    lines.append('        if lib is None:')
    lines.append('            raise RuntimeError("CTP library not loaded")')
    lines.append('')
    lines.append('        # 获取函数指针')
    lines.append('        func = lib.TraderCreateFtdcTraderApi')
    lines.append('        func.argtypes = [ctypes.c_char_p]')
    lines.append('        func.restype = ctypes.c_void_p')
    lines.append('')
    lines.append('        self._handle = func(self._flow_path)')
    lines.append('        if self._handle is None:')
    lines.append('            raise RuntimeError("Failed to create TraderApi")')
    lines.append('')
    
    # 生成 API 方法（类似 md_api.py）
    lines.append('    # ========== API 方法 ==========')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Trader'):
            continue
        if func.name == 'TraderCreateFtdcTraderApi':
            continue  # 已经在 __init__ 中处理
        if func.name.startswith('TraderSpiSetOn'):
            continue  # 在单独的 SPI 回调设置方法部分生成
        
        # 方法名（移除 Trader 前缀）
        method_name = func.name[6:]
        
        # 生成参数
        params = []
        call_args = ['self._handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
            
            # 清理参数名
            param_name = p.name.replace('[]', '').strip() if p.name else ''
            
            # 处理字符串参数
            if p.is_array and p.type == 'char' and p.is_pointer:
                params.append(f'{param_name}: List[str]')
                call_args.append('_PLACEHOLDER_STRING_ARRAY_')
            elif p.type == 'char' and p.is_pointer and not p.is_array:
                params.append(f'{param_name}: str')
                call_args.append(f'c_string({param_name})')
            else:
                if param_name:
                    params.append(f'{param_name}: {py_type}')
                    call_args.append(param_name)
        
        param_str = ', '.join(params)
        
        # 返回类型
        ret_type = c_type_to_py_ctypes_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法（类似 md_api.py 的处理）
        comment = f'    # {method_name} {func.comment}' if func.comment else f'    # {method_name}'
        lines.append(comment)
        
        if method_name == 'Release':
            lines.append(f'    def {method_name}(self):')
            lines.append(f'        """释放 API 实例"""')
            lines.append(f'        with self._lock:')
            lines.append(f'            if self._handle:')
            lines.append(f'                lib = get_trader_lib_handle()')
            lines.append(f'                if lib:')
            lines.append(f'                    func = lib.{func.name}')
            lines.append(f'                    func.argtypes = [ctypes.c_void_p]')
            lines.append(f'                    func.restype = None')
            lines.append(f'                    func(self._handle)')
            lines.append(f'                self._handle = None')
            lines.append(f'            _unregister_trader_instance(self._user_data)')
        elif method_name == 'GetApiVersion':
            lines.append(f'    def {method_name}(self) -> str:')
            lines.append(f'        """获取 API 版本"""')
            lines.append(f'        lib = get_trader_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return ""')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = []')
            lines.append(f'        func.restype = ctypes.c_char_p')
            lines.append(f'        ptr = func()')
            lines.append(f'        return go_string(ptr) if ptr else ""')
        elif method_name == 'GetTradingDay':
            lines.append(f'    def {method_name}(self) -> str:')
            lines.append(f'        """获取交易日"""')
            lines.append(f'        lib = get_trader_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return ""')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = [ctypes.c_void_p]')
            lines.append(f'        func.restype = ctypes.c_char_p')
            lines.append(f'        ptr = func(self._handle)')
            lines.append(f'        return go_string(ptr) if ptr else ""')
        elif '_PLACEHOLDER_STRING_ARRAY_' in call_args:
            array_param_idx = call_args.index('_PLACEHOLDER_STRING_ARRAY_')
            array_param_name = params[array_param_idx - 1].split(':')[0].strip()
            lines.append(f'    def {method_name}(self, {param_str}) -> int:')
            lines.append(f'        """{func.comment or method_name}"""')
            lines.append(f'        if len({array_param_name}) == 0:')
            lines.append(f'            return 0')
            lines.append(f'        ptrs, _ = c_string_array({array_param_name})')
            call_args[array_param_idx] = 'ptrs'
            lines.append(f'        lib = get_trader_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            return -1')
            lines.append(f'        func = lib.{func.name}')
            lines.append(f'        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]')
            lines.append(f'        func.restype = ctypes.c_int32')
            call_str = ', '.join(call_args)
            lines.append(f'        return func({call_str})')
        elif ret_type:
            lines.append(f'    def {method_name}(self, {param_str}) -> {ret_type}:')
            lines.append(f'        """{func.comment or method_name}"""')
            lines.append(f'        lib = get_trader_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            raise RuntimeError("CTP library not loaded")')
            lines.append(f'        func = lib.{func.name}')
            argtypes = ['ctypes.c_void_p']
            for p in func.params[1:]:
                if p.is_array and p.type == 'char' and p.is_pointer:
                    argtypes.append('ctypes.POINTER(ctypes.c_char_p)')
                else:
                    py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
                    argtypes.append(py_type)
            lines.append(f'        func.argtypes = [{", ".join(argtypes)}]')
            lines.append(f'        func.restype = {ret_type}')
            call_str = ', '.join(call_args)
            lines.append(f'        return func({call_str})')
        else:
            lines.append(f'    def {method_name}(self, {param_str}):')
            lines.append(f'        """{func.comment or method_name}"""')
            lines.append(f'        lib = get_trader_lib_handle()')
            lines.append(f'        if lib is None:')
            lines.append(f'            raise RuntimeError("CTP library not loaded")')
            lines.append(f'        func = lib.{func.name}')
            argtypes = ['ctypes.c_void_p']
            for p in func.params[1:]:
                if p.is_array and p.type == 'char' and p.is_pointer:
                    argtypes.append('ctypes.POINTER(ctypes.c_char_p)')
                else:
                    py_type = c_type_to_py_ctypes_type(p.type, p.is_pointer, typedefs, p.is_array)
                    argtypes.append(py_type)
            lines.append(f'        func.argtypes = [{", ".join(argtypes)}]')
            lines.append(f'        func.restype = None')
            call_str = ', '.join(call_args)
            lines.append(f'        func({call_str})')
        
        lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('    def set_spi(self, spi: TraderSpi):')
    lines.append('        """设置回调接口"""')
    lines.append('        with self._lock:')
    lines.append('            self._spi = spi')
    lines.append('')
    lines.append('            if self._spi_handle:')
    lines.append('                lib = get_trader_lib_handle()')
    lines.append('                if lib:')
    lines.append('                    func = lib.TraderSpiDestroy')
    lines.append('                    func.argtypes = [ctypes.c_void_p]')
    lines.append('                    func.restype = None')
    lines.append('                    func(self._spi_handle)')
    lines.append('                self._spi_handle = None')
    lines.append('')
    lines.append('            lib = get_trader_lib_handle()')
    lines.append('            if lib is None:')
    lines.append('                raise RuntimeError("CTP library not loaded")')
    lines.append('')
    lines.append('            func = lib.TraderSpiCreate')
    lines.append('            func.argtypes = [ctypes.c_void_p]')
    lines.append('            func.restype = ctypes.c_void_p')
    lines.append('            self._spi_handle = func(ctypes.c_void_p(self._user_data))')
    lines.append('')
    lines.append('            # 注册所有回调函数到 C SPI')
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        callback_suffix = cb.py_method_name[2:]
        lines.append(f'            _register_trader_callback(self._spi_handle, lib, "{callback_suffix}", self._spi, self._user_data)')
    
    lines.append('')
    lines.append('            func = lib.TraderRegisterSpi')
    lines.append('            func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]')
    lines.append('            func.restype = None')
    lines.append('            func(self._handle, self._spi_handle)')
    lines.append('')
    
    # 生成实例管理
    lines.append('# ========== 实例管理 ==========')
    lines.append('')
    lines.append('_trader_instances: dict = {}')
    lines.append('_trader_instances_lock = threading.RLock()')
    lines.append('_trader_next_id = 1')
    lines.append('')
    lines.append('def _register_trader_instance(api: TraderApi) -> int:')
    lines.append('    """注册交易 API 实例"""')
    lines.append('    global _trader_next_id')
    lines.append('    with _trader_instances_lock:')
    lines.append('        instance_id = _trader_next_id')
    lines.append('        _trader_next_id += 1')
    lines.append('        _trader_instances[instance_id] = api')
    lines.append('        return instance_id')
    lines.append('')
    lines.append('def _get_trader_instance(user_data: int) -> Optional[TraderApi]:')
    lines.append('    """获取交易 API 实例"""')
    lines.append('    with _trader_instances_lock:')
    lines.append('        return _trader_instances.get(user_data)')
    lines.append('')
    lines.append('def _unregister_trader_instance(user_data: int):')
    lines.append('    """注销交易 API 实例"""')
    lines.append('    with _trader_instances_lock:')
    lines.append('        _trader_instances.pop(user_data, None)')
    lines.append('')
    lines.append('def _register_trader_callback(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi: TraderSpi, user_data: int):')
    lines.append('    """注册回调函数到 C SPI"""')
    lines.append('    # 实际实现在 trader_callbacks.py 中')
    lines.append('    from .trader_callbacks import _register_trader_callback_impl')
    lines.append('    _register_trader_callback_impl(spi_handle, lib, callback_name, spi, user_data)')
    lines.append('')
    
    # 生成 DataCollect 函数
    lines.append('# ========== DataCollect 函数 ==========')
    lines.append('')
    lines.append('def GetSystemInfo() -> tuple[bytes, int]:')
    lines.append('    """')
    lines.append('    获取终端信息（AES+RSA 加密）')
    lines.append('    ')
    lines.append('    返回:')
    lines.append('        tuple[bytes, int]: (系统信息字节数组, 错误码)')
    lines.append('        错误码为 0 表示成功，非 0 表示采集错误（按位判断）')
    lines.append('    """')
    lines.append('    lib = get_trader_lib_handle()')
    lines.append('    if lib is None:')
    lines.append('        return None, -1')
    lines.append('    ')
    lines.append('    func = lib.DCGetSystemInfo')
    lines.append('    func.argtypes = [ctypes.POINTER(ctypes.c_byte), ctypes.POINTER(ctypes.c_int32)]')
    lines.append('    func.restype = ctypes.c_int32')
    lines.append('    ')
    lines.append('    # 分配至少 270 字节的缓冲区')
    lines.append('    buf_size = 512')
    lines.append('    buf = (ctypes.c_byte * buf_size)()')
    lines.append('    buf_len = ctypes.c_int32(buf_size)')
    lines.append('    ')
    lines.append('    ret = func(ctypes.cast(buf, ctypes.POINTER(ctypes.c_byte)), ctypes.byref(buf_len))')
    lines.append('    ')
    lines.append('    if ret != 0:')
    lines.append('        return None, ret')
    lines.append('    ')
    lines.append('    # 返回实际长度的字节数组')
    lines.append('    return bytes(buf[:buf_len.value]), 0')
    lines.append('')
    lines.append('')
    lines.append('def GetSystemInfoUnAesEncode() -> tuple[bytes, int]:')
    lines.append('    """')
    lines.append('    获取终端信息（未 AES 加密）')
    lines.append('    ')
    lines.append('    返回:')
    lines.append('        tuple[bytes, int]: (系统信息字节数组, 错误码)')
    lines.append('        错误码为 0 表示成功，非 0 表示采集错误（按位判断）')
    lines.append('    """')
    lines.append('    lib = get_trader_lib_handle()')
    lines.append('    if lib is None:')
    lines.append('        return None, -1')
    lines.append('    ')
    lines.append('    func = lib.DCGetSystemInfoUnAesEncode')
    lines.append('    func.argtypes = [ctypes.POINTER(ctypes.c_byte), ctypes.POINTER(ctypes.c_int32)]')
    lines.append('    func.restype = ctypes.c_int32')
    lines.append('    ')
    lines.append('    # 分配至少 270 字节的缓冲区')
    lines.append('    buf_size = 512')
    lines.append('    buf = (ctypes.c_byte * buf_size)()')
    lines.append('    buf_len = ctypes.c_int32(buf_size)')
    lines.append('    ')
    lines.append('    ret = func(ctypes.cast(buf, ctypes.POINTER(ctypes.c_byte)), ctypes.byref(buf_len))')
    lines.append('    ')
    lines.append('    if ret != 0:')
    lines.append('        return None, ret')
    lines.append('    ')
    lines.append('    # 返回实际长度的字节数组')
    lines.append('    return bytes(buf[:buf_len.value]), 0')
    lines.append('')
    lines.append('')
    lines.append('def GetDataCollectApiVersion() -> str:')
    lines.append('    """')
    lines.append('    获取 DataCollect API 版本')
    lines.append('    ')
    lines.append('    返回:')
    lines.append('        str: API 版本字符串')
    lines.append('    """')
    lines.append('    lib = get_trader_lib_handle()')
    lines.append('    if lib is None:')
    lines.append('        return ""')
    lines.append('    ')
    lines.append('    func = lib.DCGetDataCollectApiVersion')
    lines.append('    func.argtypes = []')
    lines.append('    func.restype = ctypes.c_char_p')
    lines.append('    ')
    lines.append('    ptr = func()')
    lines.append('    return go_string(ptr) if ptr else ""')
    lines.append('')
    
    return '\n'.join(lines)


def generate_md_callbacks_py(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_callbacks.py - 行情回调实现
    
    使用 ctypes.CFUNCTYPE 创建回调函数，将 C 回调转发到 Python SPI
    """
    lines = []
    lines.append('"""')
    lines.append('CTP 行情回调实现')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 行情回调实现')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from .md_api import _get_md_instance')
    lines.append('from .struct import *')
    lines.append('')
    
    # 生成回调包装函数
    lines.append('# ========== 回调包装函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成 Python 回调包装函数
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            # 生成参数名，如果没有名称则生成默认名称
            if p.name:
                param_name = p.name
            else:
                # 为没有名称的参数生成默认名称
                if p.type == 'void' and p.is_pointer:
                    param_name = 'userData'
                else:
                    param_name = f'param_{len(callback_params)}'
            callback_params.append(f'{param_name}: {py_type}' if py_type else param_name)
        
        param_str = ', '.join(callback_params)
        func_name = f'_go_md_{cb.py_method_name}'
        
        lines.append(f'def {func_name}({param_str}):')
        lines.append(f'    """回调函数实现: {cb.comment or cb.name}"""')
        lines.append(f'    user_data = userData.value if hasattr(userData, \'value\') else userData')
        lines.append(f'    api = _get_md_instance(user_data)')
        lines.append(f'    if api is None or api._spi is None:')
        lines.append(f'        return')
        lines.append('')
        
        # 构建调用参数
        call_args = []
        for p in cb.params[1:]:  # 跳过 userData
            if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                # CTP Field 类型，需要解引用
                call_args.append(f'{p.name}.contents if {p.name} else None')
            else:
                call_args.append(p.name)
        
        call_str = ', '.join(call_args) if call_args else ''
        lines.append(f'    try:')
        lines.append(f'        api._spi.{cb.py_method_name}({call_str})')
        lines.append(f'    except Exception as e:')
        lines.append(f'        # 回调异常不应该影响 C 层')
        lines.append(f'        import traceback')
        lines.append(f'        traceback.print_exc()')
        lines.append('')
    
    # 生成回调注册函数
    lines.append('# ========== 回调注册函数 ==========')
    lines.append('')
    lines.append('def _register_md_callback_impl(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi, user_data: int):')
    lines.append('    """注册回调函数到 C SPI（内部实现）"""')
    lines.append('    # 回调函数映射表')
    lines.append('    callback_map = {')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        callback_suffix = cb.py_method_name[2:]  # 移除 "On" 前缀
        func_name = f'_go_md_{cb.py_method_name}'
        # 生成回调类型定义
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            callback_params.append(py_type)
        param_str = ', '.join(callback_params)
        callback_type = f'ctypes.CFUNCTYPE(None, {param_str})'
        
        lines.append(f'        "{callback_suffix}": ({func_name}, {callback_type}),')
    
    lines.append('    }')
    lines.append('')
    lines.append('    if callback_name not in callback_map:')
    lines.append('        return')
    lines.append('')
    lines.append('    callback_func, callback_type = callback_map[callback_name]')
    lines.append('')
    lines.append('    # 创建 CFUNCTYPE 回调实例')
    lines.append('    c_callback = callback_type(callback_func)')
    lines.append('')
    lines.append('    # 注册到 C SPI')
    lines.append('    func_name = f"MdSpiSetOn{callback_name}"')
    lines.append('    if hasattr(lib, func_name):')
    lines.append('        func = getattr(lib, func_name)')
    lines.append('        func.argtypes = [ctypes.c_void_p, callback_type]')
    lines.append('        func.restype = None')
    lines.append('        func(spi_handle, c_callback)')
    lines.append('')
    lines.append('    # 保存回调引用，防止被 GC 回收')
    lines.append('    if not hasattr(spi, "_callbacks"):')
    lines.append('        spi._callbacks = []')
    lines.append('    spi._callbacks.append(c_callback)')
    lines.append('')
    
    return '\n'.join(lines)


def generate_trader_callbacks_py(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_callbacks.py - 交易回调实现
    """
    lines = []
    lines.append('"""')
    lines.append('CTP 交易回调实现')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('CTP 交易回调实现')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from .trader_api import _get_trader_instance')
    lines.append('from .struct import *')
    lines.append('')
    
    # 生成回调包装函数（类似 md_callbacks.py）
    lines.append('# ========== 回调包装函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            # 生成参数名，如果没有名称则生成默认名称
            if p.name:
                param_name = p.name
            else:
                # 为没有名称的参数生成默认名称
                if p.type == 'void' and p.is_pointer:
                    param_name = 'userData'
                else:
                    param_name = f'param_{len(callback_params)}'
            callback_params.append(f'{param_name}: {py_type}' if py_type else param_name)
        
        param_str = ', '.join(callback_params)
        func_name = f'_go_trader_{cb.py_method_name}'
        
        lines.append(f'def {func_name}({param_str}):')
        lines.append(f'    """回调函数实现: {cb.comment or cb.name}"""')
        lines.append(f'    user_data = userData.value if hasattr(userData, \'value\') else userData')
        lines.append(f'    api = _get_trader_instance(user_data)')
        lines.append(f'    if api is None or api._spi is None:')
        lines.append(f'        return')
        lines.append('')
        
        call_args = []
        for p in cb.params[1:]:
            if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                call_args.append(f'{p.name}.contents if {p.name} else None')
            else:
                call_args.append(p.name)
        
        call_str = ', '.join(call_args) if call_args else ''
        lines.append(f'    try:')
        lines.append(f'        api._spi.{cb.py_method_name}({call_str})')
        lines.append(f'    except Exception as e:')
        lines.append(f'        import traceback')
        lines.append(f'        traceback.print_exc()')
        lines.append('')
    
    # 生成回调注册函数
    lines.append('# ========== 回调注册函数 ==========')
    lines.append('')
    lines.append('def _register_trader_callback_impl(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi, user_data: int):')
    lines.append('    """注册回调函数到 C SPI（内部实现）"""')
    lines.append('    # 回调函数映射表')
    lines.append('    callback_map = {')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        callback_suffix = cb.py_method_name[2:]
        func_name = f'_go_trader_{cb.py_method_name}'
        # 生成回调类型定义
        callback_params = []
        for p in cb.params:
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            callback_params.append(py_type)
        param_str = ', '.join(callback_params)
        callback_type = f'ctypes.CFUNCTYPE(None, {param_str})'
        
        lines.append(f'        "{callback_suffix}": ({func_name}, {callback_type}),')
    
    lines.append('    }')
    lines.append('')
    lines.append('    if callback_name not in callback_map:')
    lines.append('        return')
    lines.append('')
    lines.append('    callback_func, callback_type = callback_map[callback_name]')
    lines.append('')
    lines.append('    # 创建 CFUNCTYPE 回调实例')
    lines.append('    c_callback = callback_type(callback_func)')
    lines.append('')
    lines.append('    # 注册到 C SPI')
    lines.append('    func_name = f"TraderSpiSetOn{callback_name}"')
    lines.append('    if hasattr(lib, func_name):')
    lines.append('        func = getattr(lib, func_name)')
    lines.append('        func.argtypes = [ctypes.c_void_p, callback_type]')
    lines.append('        func.restype = None')
    lines.append('        func(spi_handle, c_callback)')
    lines.append('')
    lines.append('    # 保存回调引用，防止被 GC 回收')
    lines.append('    if not hasattr(spi, "_callbacks"):')
    lines.append('        spi._callbacks = []')
    lines.append('    spi._callbacks.append(c_callback)')
    lines.append('')
    
    return '\n'.join(lines)


def generate_md_default_spi_py(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_default_spi.py - 行情 SPI 默认空实现
    """
    lines = []
    lines.append('"""')
    lines.append('默认行情 SPI 空实现')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('默认 SPI 空实现，可用于嵌入')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from .md_api import MdSpi')
    lines.append('from .struct import *')
    lines.append('')
    lines.append('class DefaultMdSpi(MdSpi):')
    lines.append('    """默认行情回调实现（空实现）"""')
    lines.append('    ')
    lines.append('    # 使用方式：继承此类，只需实现需要的方法')
    lines.append('    ')
    
    md_callbacks = [cb for cb in callbacks if cb.name.startswith('Md')]
    for cb in md_callbacks:
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name}: {py_type}')
            else:
                method_params.append(py_type)
        param_str = ', '.join(method_params)
        lines.append(f'    def {cb.py_method_name}(self, {param_str}):')
        lines.append('        """空实现"""')
        lines.append('        pass')
        lines.append('')
    
    return '\n'.join(lines)


def generate_trader_default_spi_py(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_default_spi.py - 交易 SPI 默认空实现
    """
    lines = []
    lines.append('"""')
    lines.append('默认交易 SPI 空实现')
    lines.append('')
    lines.append('此文件由代码生成器自动生成，请勿手动修改')
    lines.append('默认 SPI 空实现，可用于嵌入')
    lines.append('"""')
    lines.append('')
    lines.append('import ctypes')
    lines.append('from .trader_api import TraderSpi')
    lines.append('from .struct import *')
    lines.append('')
    lines.append('class DefaultTraderSpi(TraderSpi):')
    lines.append('    """默认交易回调实现（空实现）"""')
    lines.append('    ')
    lines.append('    # 使用方式：继承此类，只需实现需要的方法')
    lines.append('    ')
    
    trader_callbacks = [cb for cb in callbacks if cb.name.startswith('Trader')]
    for cb in trader_callbacks:
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            py_type = c_type_to_py_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name}: {py_type}')
            else:
                method_params.append(py_type)
        param_str = ', '.join(method_params)
        lines.append(f'    def {cb.py_method_name}(self, {param_str}):')
        lines.append('        """空实现"""')
        lines.append('        pass')
        lines.append('')
    
    return '\n'.join(lines)


# ========== 主函数 ==========

def main():
    parser = argparse.ArgumentParser(description='CTP C API 转 Python ctypes 包装代码生成器')
    parser.add_argument('--input', required=True, help='C API 头文件目录 (包含 ctpmd_c_api.h, ctptrader_c_api.h)')
    parser.add_argument('--struct', required=True, help='CTP 结构体头文件目录 (包含 ThostFtdcUserApiDataType.h, ThostFtdcUserApiStruct.h)')
    parser.add_argument('--output', required=True, help='输出目录')
    args = parser.parse_args()
    
    input_dir = Path(args.input)
    struct_dir = Path(args.struct)
    output_dir = Path(args.output)
    
    # 确保输出目录存在
    output_dir.mkdir(parents=True, exist_ok=True)
    
    print(f"解析类型定义: {struct_dir / 'ThostFtdcUserApiDataType.h'}")
    typedefs, enums, defines = parse_datatype_header(struct_dir / 'ThostFtdcUserApiDataType.h')
    print(f"  找到 {len(typedefs)} 个类型定义, {len(enums)} 个枚举, {len(defines)} 个常量组")
    
    print(f"解析结构体定义: {struct_dir / 'ThostFtdcUserApiStruct.h'}")
    structs = parse_struct_header(struct_dir / 'ThostFtdcUserApiStruct.h', typedefs)
    print(f"  找到 {len(structs)} 个结构体")
    
    # 解析行情 API
    md_header = input_dir / 'ctpmd_c_api.h'
    print(f"解析行情 API: {md_header}")
    md_functions, md_callbacks = parse_c_header(md_header)
    print(f"  找到 {len(md_functions)} 个函数, {len(md_callbacks)} 个回调")
    
    # 解析交易 API
    trader_header = input_dir / 'ctptrader_c_api.h'
    print(f"解析交易 API: {trader_header}")
    trader_functions, trader_callbacks = parse_c_header(trader_header)
    print(f"  找到 {len(trader_functions)} 个函数, {len(trader_callbacks)} 个回调")
    
    # 生成代码文件
    print(f"\n生成代码到: {output_dir}")
    
    # utils.py
    utils_file = output_dir / 'utils.py'
    utils_file.write_text(generate_utils_py(), encoding='utf-8')
    print(f"  生成 {utils_file}")
    
    # datatype.py - 类型别名、枚举常量定义（不包含结构体）
    datatype_file = output_dir / 'datatype.py'
    datatype_content = generate_datatype_py(typedefs, enums, defines)
    datatype_file.write_text(datatype_content, encoding='utf-8')
    print(f"  生成 {datatype_file}")
    
    # struct.py - 结构体定义
    struct_file = output_dir / 'struct.py'
    struct_content = generate_struct_py(structs, typedefs)
    struct_file.write_text(struct_content, encoding='utf-8')
    print(f"  生成 {struct_file}")
    
    # loader.py
    loader_file = output_dir / 'loader.py'
    loader_file.write_text(generate_loader_py(), encoding='utf-8')
    print(f"  生成 {loader_file}")
    
    # md_api.py
    md_api_file = output_dir / 'md_api.py'
    md_api_file.write_text(generate_md_api_py(md_functions, md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_api_file}")
    
    # trader_api.py
    trader_api_file = output_dir / 'trader_api.py'
    trader_api_file.write_text(generate_trader_api_py(trader_functions, trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_api_file}")
    
    # md_callbacks.py
    md_callbacks_file = output_dir / 'md_callbacks.py'
    md_callbacks_file.write_text(generate_md_callbacks_py(md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_callbacks_file}")
    
    # trader_callbacks.py
    trader_callbacks_file = output_dir / 'trader_callbacks.py'
    trader_callbacks_file.write_text(generate_trader_callbacks_py(trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_callbacks_file}")
    
    # md_default_spi.py
    md_default_spi_file = output_dir / 'md_default_spi.py'
    md_default_spi_file.write_text(generate_md_default_spi_py(md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_default_spi_file}")
    
    # trader_default_spi.py
    trader_default_spi_file = output_dir / 'trader_default_spi.py'
    trader_default_spi_file.write_text(generate_trader_default_spi_py(trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_default_spi_file}")
    
    # __init__.py
    init_file = output_dir / '__init__.py'
    init_content = '''"""
CTP Python 包装库

此文件由代码生成器自动生成，请勿手动修改
"""

from .loader import auto_load_library, load_ctp_library, get_md_lib_handle, get_trader_lib_handle
from .datatype import *
from .struct import *
from .utils import *
from .md_api import MdApi, MdSpi
from .trader_api import TraderApi, TraderSpi, GetSystemInfo, GetSystemInfoUnAesEncode, GetDataCollectApiVersion
from .md_default_spi import DefaultMdSpi
from .trader_default_spi import DefaultTraderSpi

__all__ = [
    'auto_load_library',
    'load_ctp_library',
    'get_md_lib_handle',
    'get_trader_lib_handle',
    'MdApi',
    'MdSpi',
    'TraderApi',
    'TraderSpi',
    'DefaultMdSpi',
    'DefaultTraderSpi',
    'GetSystemInfo',
    'GetSystemInfoUnAesEncode',
    'GetDataCollectApiVersion',
]
'''
    init_file.write_text(init_content, encoding='utf-8')
    print(f"  生成 {init_file}")
    
    print("\n代码生成完成!")
    print(f"生成的文件列表:")
    print(f"  - utils.py              : 工具函数")
    print(f"  - datatype.py           : 枚举和类型定义 ({len(typedefs)} 个类型, {len(enums)} 个枚举)")
    print(f"  - struct.py             : 结构体定义 ({len(structs)} 个)")
    print(f"  - loader.py             : 动态库加载")
    print(f"  - md_api.py             : 行情 API ({len(md_functions)} 个方法)")
    print(f"  - trader_api.py         : 交易 API ({len(trader_functions)} 个方法)")
    print(f"  - md_callbacks.py       : 行情回调 ({len(md_callbacks)} 个)")
    print(f"  - trader_callbacks.py   : 交易回调 ({len(trader_callbacks)} 个)")
    print(f"  - md_default_spi.py     : 行情默认 SPI 实现")
    print(f"  - trader_default_spi.py : 交易默认 SPI 实现")
    print(f"  - __init__.py           : 包初始化")


if __name__ == '__main__':
    main()
