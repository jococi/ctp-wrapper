#!/usr/bin/env python3
"""
CTP C API 转 Go PureGo 包装代码生成器

功能：
- 解析 C API 头文件（ctptrader_c_api.h, ctpmd_c_api.h）
- 解析 CTP 结构体和数据类型定义
- 生成使用 purego 的 Go 包装代码
- 支持多实例，使用 userData 机制

用法：
    python3 generate_go_api.py --input ../csrc --struct ../ctpapi/linux --output ../ctpgo
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
    go_method_name: str = ""    # Go 方法名（如 "OnFrontConnected"）


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


# ========== C 类型到 Go 类型映射 ==========

# CTP 基础数据类型映射
CTP_TYPE_MAP = {
    # C 基础类型 -> Go 类型
    'char': 'byte',
    'int': 'int32',
    'short': 'int16',
    'double': 'float64',
    'float': 'float32',
    'bool': 'bool',
    'void': '',
    
    # CTP 特定类型会在解析时动态添加
}


# ========== 解析函数 ==========

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
        
        # 提取 Go 方法名
        go_method_name = extract_go_method_name(callback_name)
        
        callbacks.append(CallbackType(
            name=callback_name,
            params=params,
            comment=comment,
            go_method_name=go_method_name
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


def extract_go_method_name(callback_name: str) -> str:
    """从回调类型名提取 Go 方法名"""
    name = callback_name
    # 移除前缀 (Trader/Md)
    if name.startswith("TraderOn"):
        name = name[6:]  # 移除 "Trader"，保留 On
    elif name.startswith("MdOn"):
        name = name[2:]   # 移除 "Md"，保留 On
    
    if name.endswith("Callback"):
        name = name[:-8]  # 移除 "Callback"
    
    return name


# ========== Go 类型转换 ==========

def c_type_to_go_type(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef], is_array: bool = False) -> str:
    """
    将 C 类型转换为 Go 类型
    
    类型映射规则：
    1. 句柄类型 (XxxHandle) -> uintptr
    2. CTP Field 结构体指针 (*CThostFtdcXxxField) -> *CThostFtdcXxxField
    3. CTP typedef 类型：
       - 数组类型 (typedef char TThostFtdcBrokerIDType[11]) -> [11]byte
       - 简单类型 (typedef int TThostFtdcVolumeType) -> int32
       - 单字符类型 (typedef char TThostFtdcXxxType) -> byte
    4. 基础 C 类型 -> 对应 Go 类型
    
    参数:
        c_type: C 类型名
        is_pointer: 是否是指针类型
        typedefs: 已解析的 typedef 映射
        is_array: 是否是数组（用于特殊处理 char*[]）
        
    返回:
        对应的 Go 类型字符串
    """
    c_type = c_type.strip().replace('const', '').replace('*', '').strip()
    
    # 句柄类型: MdApiHandle, TraderApiHandle, etc.
    if c_type.endswith('Handle'):
        return 'uintptr'
    
    # CTP Field 结构体类型
    if c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
        if is_pointer:
            return f'*{c_type}'
        return c_type
    
    # CTP typedef 类型（如 TThostFtdcBrokerIDType）
    if c_type in typedefs:
        td = typedefs[c_type]
        if td.size > 0:
            # 数组类型: typedef char TThostFtdcBrokerIDType[11] -> [11]byte
            base_go = CTP_TYPE_MAP.get(td.base_type, td.base_type)
            return f'[{td.size}]{base_go}'
        else:
            # 非数组类型: typedef int TThostFtdcVolumeType -> int32
            return CTP_TYPE_MAP.get(td.base_type, td.base_type)
    
    # 基础 C 类型
    if is_array and is_pointer and c_type == 'char':
        # char* [] 格式（字符串数组），返回 **byte
        return '**byte'
    elif is_pointer:
        if c_type == 'void':
            return 'uintptr'  # void* -> uintptr
        elif c_type == 'char':
            return '*byte'    # char* -> *byte (C 字符串)
        else:
            go_type = CTP_TYPE_MAP.get(c_type, c_type)
            return f'*{go_type}'
    else:
        if c_type == 'void':
            return ''
        return CTP_TYPE_MAP.get(c_type, c_type)


def c_type_to_go_callback_param(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef]) -> str:
    """将 C 类型转换为 Go 回调参数类型（用于接口定义）"""
    c_type = c_type.strip().replace('const', '').strip()
    
    # 指针类型
    if is_pointer:
        if c_type == 'void':
            return 'uintptr'
        elif c_type == 'char':
            return '*byte'
        elif c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
            return f'*{c_type}'
        else:
            go_type = CTP_TYPE_MAP.get(c_type, c_type)
            return f'*{go_type}'
    else:
        if c_type == 'void':
            return ''
        elif c_type == 'int':
            return 'int32'
        elif c_type == 'bool':
            return 'bool'
        return CTP_TYPE_MAP.get(c_type, c_type)


def c_type_to_go_export_param(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef]) -> str:
    """将 C 类型转换为 Go //export 函数参数类型（使用 C 类型，cgo 会自动转换）"""
    c_type = c_type.strip().replace('const', '').strip()
    
    # 指针类型
    if is_pointer:
        if c_type == 'void':
            return 'uintptr'
        elif c_type == 'char':
            return '*C.char'
        elif c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
            # 对于 CTP Field 类型，使用 C 类型（cgo 会自动转换）
            return f'*C.struct_{c_type}'
        else:
            # 其他类型，尝试使用 C 类型
            return f'*C.{c_type}'
    else:
        if c_type == 'void':
            return ''
        elif c_type == 'int':
            return 'C.int'
        elif c_type == 'bool':
            return 'C.bool'
        elif c_type in CTP_TYPE_MAP:
            go_type = CTP_TYPE_MAP[c_type]
            if go_type == 'int32':
                return 'C.int'
            elif go_type == 'int16':
                return 'C.short'
            elif go_type == 'float64':
                return 'C.double'
            elif go_type == 'float32':
                return 'C.float'
            return f'C.{c_type}'
        return f'C.{c_type}'


# ========== 代码生成 ==========

def generate_utils_go() -> str:
    """生成 utils.go"""
    return '''package ctpgo

import (
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// CString 将 Go 字符串转换为 C 字符串（以 null 结尾的字节切片）
// 注意：返回的指针指向 Go 管理的内存，在传递给 C 后需要确保其生命周期
func CString(s string) *byte {
	if s == "" {
		return nil
	}
	bs := make([]byte, len(s)+1)
	copy(bs, s)
	bs[len(s)] = 0 // null terminator
	return &bs[0]
}

// CStringArray 将 Go 字符串切片转换为 C 字符串数组
// 返回指向字符串指针数组的指针和底层数据（需要保持引用防止 GC）
func CStringArray(ss []string) (**byte, [][]byte) {
	if len(ss) == 0 {
		return nil, nil
	}
	
	// 创建字节切片数组保存字符串数据
	data := make([][]byte, len(ss))
	ptrs := make([]*byte, len(ss))
	
	for i, s := range ss {
		data[i] = make([]byte, len(s)+1)
		copy(data[i], s)
		data[i][len(s)] = 0
		ptrs[i] = &data[i][0]
	}
	
	return &ptrs[0], data
}

// GoString 将 C 字符串（*byte）转换为 Go 字符串
func GoString(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	
	var length int
	for p := ptr; *p != 0; p = (*byte)(unsafe.Add(unsafe.Pointer(p), 1)) {
		length++
	}
	
	return string(unsafe.Slice(ptr, length))
}

// BytesToString 将固定长度字节数组转换为字符串（去除尾部的 null）
func BytesToString(b []byte) string {
	for i, v := range b {
		if v == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

// GB18030 将 GB18030 编码的字节切片转换为 UTF-8 字符串
func GB18030(b []byte) string {
	// 找到 null 终止符
	var end int
	for end = 0; end < len(b); end++ {
		if b[end] == 0 {
			break
		}
	}
	if end == 0 {
		return ""
	}
	
	decoder := simplifiedchinese.GB18030.NewDecoder()
	result, err := decoder.Bytes(b[:end])
	if err != nil {
		return string(b[:end])
	}
	return string(result)
}

// StringToBytes 将字符串复制到固定长度字节数组
func StringToBytes(s string, size int) []byte {
	b := make([]byte, size)
	copy(b, s)
	return b
}

// CopyStringToArray 将字符串复制到字节数组（用于填充 CTP 结构体字段）
func CopyStringToArray(dst []byte, src string) {
	copy(dst, src)
}

// BoolToInt 将 bool 转换为 int（C 风格）
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool 将 int 转换为 bool（C 风格）
func IntToBool(i int) bool {
	return i != 0
}
'''


def generate_datatype_go(typedefs: Dict[str, CTypedef], enums: Dict[str, CEnum], defines: Dict[str, List[CDefine]] = None) -> str:
    """
    生成 datatype.go - 完整的 CTP 数据类型定义
    
    将 ThostFtdcUserApiDataType.h 翻译成 Go 类型：
    
    1. enum -> type XXX int32 + const (...)
       enum THOST_TE_RESUME_TYPE { THOST_TERT_RESTART = 0, ... }
       -> type THOST_TE_RESUME_TYPE int32
       -> const ( THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0 ... )
    
    2. typedef char XXX[N] -> type XXX [N]byte
       typedef char TThostFtdcBrokerIDType[11];
       -> type TThostFtdcBrokerIDType [11]byte
    
    3. typedef char XXX (单字符) + #define -> type XXX byte + const (...)
       #define THOST_FTDC_EXP_Normal '0'
       typedef char TThostFtdcExchangePropertyType;
       -> type TThostFtdcExchangePropertyType byte
       -> const ( THOST_FTDC_EXP_Normal TThostFtdcExchangePropertyType = '0' ... )
    
    4. typedef int XXX -> type XXX int32
       typedef int TThostFtdcIPPortType;
       -> type TThostFtdcIPPortType int32
    
    5. typedef short XXX -> type XXX int16
       typedef short TThostFtdcCommPhaseNoType;
       -> type TThostFtdcCommPhaseNoType int16
    
    6. typedef double XXX -> type XXX float64
       typedef double TThostFtdcHedgeRateType;
       -> type TThostFtdcHedgeRateType float64
    """
    if defines is None:
        defines = {}
    
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 数据类型定义 - 来自 ThostFtdcUserApiDataType.h')
    lines.append('')
    
    # ========== 1. 生成 enum 类型 ==========
    if enums:
        lines.append('// ========== 枚举类型 ==========')
        lines.append('')
        
        for enum_name, enum in sorted(enums.items()):
            # 提取注释
            description = _extract_description(enum.comment)
            
            # 特殊处理 THOST_TE_RESUME_TYPE
            if enum_name == 'THOST_TE_RESUME_TYPE':
                description = '订阅类型'
            
            if description:
                lines.append(f'// {enum_name} {description}')
            else:
                lines.append(f'// {enum_name}')
            
            lines.append(f'type {enum_name} int32')
            lines.append('')
            
            if enum.values:
                lines.append('const (')
                
                # THOST_TE_RESUME_TYPE 的特殊注释
                enum_value_comments = {}
                if enum_name == 'THOST_TE_RESUME_TYPE':
                    enum_value_comments = {
                        'THOST_TERT_RESTART': '从本交易日开始重传',
                        'THOST_TERT_RESUME': '从上次收到的续传',
                        'THOST_TERT_QUICK': '只传送登录后的流内容',
                        'THOST_TERT_NONE': '不传送'
                    }
                
                for value_name, value in enum.values:
                    value_comment = enum_value_comments.get(value_name, "")
                    if value_comment:
                        lines.append(f'\t{value_name} {enum_name} = {value} // {value_comment}')
                    else:
                        lines.append(f'\t{value_name} {enum_name} = {value}')
                
                lines.append(')')
                lines.append('')
    
    # ========== 2. 生成 typedef 类型 ==========
    if typedefs:
        lines.append('// ========== 类型定义 ==========')
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
            lines.append('// ----- 字符串类型 -----')
            lines.append('')
            for type_name, td in char_array_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'// {type_name} {description}')
                lines.append(f'type {type_name} = [{td.size}]byte')
                lines.append('')
        
        # 2.2 整数类型
        if int_types:
            lines.append('// ----- 整数类型 -----')
            lines.append('')
            for type_name, td in int_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'// {type_name} {description}')
                lines.append(f'type {type_name} = int32')
                lines.append('')
        
        # 2.3 短整数类型
        if short_types:
            lines.append('// ----- 短整数类型 -----')
            lines.append('')
            for type_name, td in short_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'// {type_name} {description}')
                lines.append(f'type {type_name} = int16')
                lines.append('')
        
        # 2.4 浮点类型
        if double_types:
            lines.append('// ----- 浮点类型 -----')
            lines.append('')
            for type_name, td in double_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'// {type_name} {description}')
                lines.append(f'type {type_name} = float64')
                lines.append('')
        
        # 2.5 单字符枚举类型（有 #define 常量）
        if char_enum_types:
            lines.append('// ----- 字符枚举类型 -----')
            lines.append('')
            for type_name, td in char_enum_types:
                description = _extract_description(td.comment)
                type_defines = defines.get(type_name, [])
                
                # 检查常量值类型：
                # - 单字符（如 '0', 'A'）-> byte 类型
                # - 多字符（如 '102001'）-> string 类型（C 多字符字面量，实际是字符串标识符）
                has_multi_char = any(len(d.value) > 1 for d in type_defines)
                
                if description:
                    lines.append(f'// {type_name} {description}')
                
                if has_multi_char:
                    # 多字符常量（如交易代码 '102001'），使用 string 类型
                    lines.append(f'type {type_name} = string')
                else:
                    # 单字符常量，使用 byte 类型
                    lines.append(f'type {type_name} = byte')
                lines.append('')
                
                # 生成关联的常量（不指定类型）
                if type_defines:
                    lines.append('const (')
                    for d in type_defines:
                        value = d.value
                        if has_multi_char:
                            # 字符串类型，使用双引号
                            value_str = f'"{value}"'
                        else:
                            # 单字符，使用字符字面量
                            value_str = f"'{value}'"
                        
                        if d.comment:
                            lines.append(f'\t{d.name} = {value_str} // {d.comment}')
                        else:
                            lines.append(f'\t{d.name} = {value_str}')
                    lines.append(')')
                    lines.append('')
        
        # 2.6 单字符简单类型（无 #define 常量）
        if char_simple_types:
            lines.append('// ----- 单字符类型 -----')
            lines.append('')
            for type_name, td in char_simple_types:
                description = _extract_description(td.comment)
                if description:
                    lines.append(f'// {type_name} {description}')
                lines.append(f'type {type_name} = byte')
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


def generate_struct_go(structs: Dict[str, CStruct], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 struct.go - CTP 业务结构体定义
    
    生成格式示例（参考 ThostFtdcUserApiStruct.h）：
    // CThostFtdcReqUserLoginField 用户登录请求
    type CThostFtdcReqUserLoginField struct {
        TradingDay TThostFtdcDateType     // 交易日
        BrokerID   TThostFtdcBrokerIDType // 经纪公司代码
        UserID     TThostFtdcUserIDType   // 用户代码
        ...
    }
    
    字段类型直接使用 CTP 原始类型名称，因为类型已在 datatype.go 中定义
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 结构体定义')
    lines.append('')
    
    # 生成结构体
    lines.append('// ========== CTP 结构体 ==========')
    lines.append('')
    
    for struct_name, struct in sorted(structs.items()):
        # 提取并清理结构体注释
        struct_comment = struct.comment.strip()
        if struct_comment:
            # 清理注释，移除多余的斜杠和空格
            struct_comment = struct_comment.replace('///', '').strip()
            if struct_comment.startswith('/'):
                struct_comment = struct_comment[1:].strip()
            # 生成格式: // CThostFtdcXxxField 结构体描述
            if struct_comment:
                lines.append(f'// {struct_name} {struct_comment}')
        
        lines.append(f'type {struct_name} struct {{')
        
        for field in struct.fields:
            # 直接使用 CTP 原始类型名称（类型已在 datatype.go 中定义）
            go_type = field.type
            
            # 获取字段注释（使用字段自身的注释，参考 ThostFtdcUserApiStruct.h 的格式）
            field_comment = ""
            if field.comment:
                # 使用字段自己的注释（如：交易日、经纪公司代码 等）
                field_comment = f' // {field.comment}'
            
            lines.append(f'\t{field.name} {go_type}{field_comment}')
        
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_md_api_go(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_api.go - 行情 API 封装
    
    包含：
    1. 回调类型定义 (MdOnXxxCallback)
    2. MdSpiCallbacks 结构体
    3. MdSpi 接口
    4. MdApi 结构体
    5. C 函数声明和初始化
    6. 实例管理
    7. NewMdApi 构造函数
    8. API 方法和 SPI 方法
    9. SetSpi 方法
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 行情 API 封装')
    lines.append('')
    lines.append('import (')
    lines.append('\t"fmt"')
    lines.append('\t"os"')
    lines.append('\t"path/filepath"')
    lines.append('\t"runtime"')
    lines.append('\t"sync"')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
    lines.append('')
    
    # 生成回调类型定义
    lines.append('// ========== 回调类型定义 ==========')
    lines.append('')
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成回调函数类型
        callback_params = []
        for p in cb.params:
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                callback_params.append(f'{p.name} {go_type}')
            else:
                callback_params.append(go_type)
        
        param_str = ', '.join(callback_params)
        lines.append(f'// {cb.name} {cb.comment}' if cb.comment else f'// {cb.name}')
        lines.append(f'type {cb.name} func({param_str})')
        lines.append('')
    
    # 生成回调结构体
    lines.append('// MdSpiCallbacks 回调结构体（用于批量设置）')
    lines.append('type MdSpiCallbacks struct {')
    lines.append('\tUserData uintptr')
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        field_name = cb.name.replace('Md', '').replace('Callback', '')
        lines.append(f'\t{field_name} {cb.name}')
    lines.append('}')
    lines.append('')
    
    # 生成 SPI 接口
    lines.append('// ========== MdSpi 接口 ==========')
    lines.append('')
    lines.append('// MdSpi 行情回调接口')
    lines.append('type MdSpi interface {')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成方法签名
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name} {go_type}')
            else:
                method_params.append(go_type)
        
        param_str = ', '.join(method_params)
        comment = f' // {cb.comment}' if cb.comment else ''
        lines.append(f'\t{cb.go_method_name}({param_str}){comment}')
    
    lines.append('}')
    lines.append('')
    
    # 生成 MdApi 结构体
    lines.append('// ========== MdApi 结构体 ==========')
    lines.append('')
    lines.append('// MdApi 行情 API 封装')
    lines.append('type MdApi struct {')
    lines.append('\thandle    uintptr')
    lines.append('\tspi       MdSpi')
    lines.append('\tspiHandle uintptr // C SPI 实例句柄')
    lines.append('\tuserData  uintptr')
    lines.append('\tmu        sync.RWMutex')
    lines.append('\tflowPath  []byte // 保存 flowPath 的 C 字符串，防止被 GC 回收')
    lines.append('}')
    lines.append('')
    
    # 生成函数变量声明
    lines.append('// ========== C 函数声明 ==========')
    lines.append('')
    lines.append('var (')
    lines.append('\tmdOnce sync.Once')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Md'):
            continue
        var_name = f'_{func.name}'
        
        # 生成参数类型
        param_types = []
        for i, p in enumerate(func.params):
            # 特殊处理 MdSpiSetOnXxx 函数：第二个参数（回调函数指针）应该是 uintptr
            if func.name.startswith('MdSpiSetOn') and i == 1:
                param_types.append('uintptr')
            else:
                go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs, p.is_array)
                param_types.append(go_type)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成函数签名（统一使用单行格式）
        param_str = ', '.join(param_types) if param_types else ''
        if param_str:
            if ret_type:
                lines.append(f'\t{var_name} func({param_str}) {ret_type}')
            else:
                lines.append(f'\t{var_name} func({param_str})')
        else:
            if ret_type:
                lines.append(f'\t{var_name} func() {ret_type}')
            else:
                lines.append(f'\t{var_name} func()')
    
    lines.append(')')
    lines.append('')
    
    # 生成初始化函数
    lines.append('// initMdApi 初始化行情 API 函数')
    lines.append('func initMdApi(lib uintptr) {')
    lines.append('\tmdOnce.Do(func() {')
    
    for func in functions:
        if not func.name.startswith('Md'):
            continue
        var_name = f'_{func.name}'
        lines.append(f'\t\tpurego.RegisterLibFunc(&{var_name}, lib, "{func.name}")')
    
    lines.append('\t})')
    lines.append('}')
    lines.append('')
    
    # 生成实例管理
    lines.append('// ========== 实例管理 ==========')
    lines.append('')
    lines.append('var (')
    lines.append('\tmdInstances   = make(map[uintptr]*MdApi)')
    lines.append('\tmdInstancesMu sync.RWMutex')
    lines.append('\tmdNextID      uintptr = 1')
    lines.append(')')
    lines.append('')
    
    lines.append('func registerMdInstance(api *MdApi) uintptr {')
    lines.append('\tmdInstancesMu.Lock()')
    lines.append('\tdefer mdInstancesMu.Unlock()')
    lines.append('\tid := mdNextID')
    lines.append('\tmdNextID++')
    lines.append('\tmdInstances[id] = api')
    lines.append('\treturn id')
    lines.append('}')
    lines.append('')
    
    lines.append('func getMdInstance(userData uintptr) *MdApi {')
    lines.append('\tmdInstancesMu.RLock()')
    lines.append('\tdefer mdInstancesMu.RUnlock()')
    lines.append('\treturn mdInstances[userData]')
    lines.append('}')
    lines.append('')
    
    lines.append('func unregisterMdInstance(userData uintptr) {')
    lines.append('\tmdInstancesMu.Lock()')
    lines.append('\tdefer mdInstancesMu.Unlock()')
    lines.append('\tdelete(mdInstances, userData)')
    lines.append('}')
    lines.append('')
    
    # 生成构造函数
    lines.append('// ========== 构造函数 ==========')
    lines.append('')
    lines.append('// NewMdApi 创建行情 API 实例')
    lines.append('// 首次调用时会自动加载 CTP 库（如果尚未加载）')
    lines.append('func NewMdApi(flowPath string, usingUdp, multicast bool) *MdApi {')
    lines.append('\t// 自动加载库（如果尚未加载）')
    lines.append('\tif err := autoLoadLibrary(); err != nil {')
    lines.append('\t\t// 如果自动加载失败，返回 nil（或者可以 panic，取决于设计）')
    lines.append('\t\t// 这里返回 nil，让调用者检查')
    lines.append('\t\treturn nil')
    lines.append('\t}')
    lines.append('')
    lines.append('\tapi := &MdApi{}')
    lines.append('\tapi.userData = registerMdInstance(api)')
    lines.append('')
    lines.append('\t// 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录')
    lines.append('\t// CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的')
    lines.append('\t// 所以转换为绝对路径更可靠')
    lines.append('\tabsFlowPath := flowPath')
    lines.append('\tif !filepath.IsAbs(flowPath) {')
    lines.append('\t\t// 如果是相对路径，转换为基于当前工作目录的绝对路径')
    lines.append('\t\tvar err error')
    lines.append('\t\tabsFlowPath, err = filepath.Abs(flowPath)')
    lines.append('\t\tif err != nil {')
    lines.append('\t\t\t// 如果转换失败，使用原始路径')
    lines.append('\t\t\tabsFlowPath = flowPath')
    lines.append('\t\t}')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）')
    lines.append('\tif len(absFlowPath) > 0 && absFlowPath[len(absFlowPath)-1] != filepath.Separator {')
    lines.append('\t\tabsFlowPath += string(filepath.Separator)')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 确保目录存在（CTP API 需要这个目录来创建 flow 文件）')
    lines.append('\tif err := os.MkdirAll(absFlowPath, 0755); err != nil {')
    lines.append('\t\t// 如果创建目录失败，记录错误但继续（CTP API 可能会自己创建）')
    lines.append('\t\tfmt.Printf("警告: 无法创建 flow 目录 %s: %v\\n", absFlowPath, err)')
    lines.append('\t\t// 这里不返回错误，让 CTP API 自己处理')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收')
    lines.append('\t// CTP API 可能会在后续使用这个路径')
    lines.append('\tapi.flowPath = make([]byte, len(absFlowPath)+1)')
    lines.append('\tcopy(api.flowPath, absFlowPath)')
    lines.append('\tapi.flowPath[len(absFlowPath)] = 0 // null terminator')
    lines.append('\tpathPtr := &api.flowPath[0]')
    lines.append('')
    lines.append('\tapi.handle = _MdCreateFtdcMdApi(pathPtr, usingUdp, multicast)')
    lines.append('')
    lines.append('\truntime.SetFinalizer(api, (*MdApi).Release)')
    lines.append('\treturn api')
    lines.append('}')
    lines.append('')
    
    # 生成 API 方法
    lines.append('// ========== API 方法 ==========')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Md'):
            continue
        if func.name == 'MdCreateFtdcMdApi':
            continue  # 已经在 NewMdApi 中处理
        if func.name.startswith('MdSpiSetOn'):
            continue  # 在单独的 SPI 回调设置方法部分生成
        
        # 方法名（移除 Md 前缀）
        method_name = func.name[2:]
        
        # 生成参数
        params = []
        call_args = ['api.handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            
            # 清理参数名（移除可能的 [] 后缀）
            param_name = p.name.replace('[]', '').strip() if p.name else ''
            
            # 处理字符串参数
            if p.is_array and p.type == 'char' and p.is_pointer:
                # char* [] 或 char** 类型（字符串数组）
                params.append(f'{param_name} []string')
                # 将在方法体中特殊处理
                call_args.append('_PLACEHOLDER_STRING_ARRAY_')
            elif p.type == 'char' and p.is_pointer and not p.is_array:
                params.append(f'{param_name} string')
                call_args.append(f'CString({param_name})')
            else:
                if param_name:
                    params.append(f'{param_name} {go_type}')
                    call_args.append(param_name)
        
        param_str = ', '.join(params)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法
        comment = f'// {method_name} {func.comment}' if func.comment else f'// {method_name}'
        lines.append(comment)
        
        # 特殊处理方法
        if method_name == 'Release':
            # Release 方法需要先调用 C 函数，然后注销实例
            lines.append(f'func (api *MdApi) {method_name}() {{')
            lines.append(f'\t_{func.name}(api.handle)')
            lines.append('\tunregisterMdInstance(api.userData)')
        elif method_name == 'GetApiVersion':
            lines.append(f'func (api *MdApi) {method_name}() string {{')
            lines.append(f'\tptr := _{func.name}()')
            lines.append('\tif ptr == nil {')
            lines.append('\t\treturn ""')
            lines.append('\t}')
            lines.append('\treturn GoString(ptr)')
        elif method_name == 'GetTradingDay':
            lines.append(f'func (api *MdApi) {method_name}() string {{')
            lines.append(f'\tptr := _{func.name}(api.handle)')
            lines.append('\tif ptr == nil {')
            lines.append('\t\treturn ""')
            lines.append('\t}')
            lines.append('\treturn GoString(ptr)')
        elif '_PLACEHOLDER_STRING_ARRAY_' in call_args:
            # 处理字符串数组参数
            lines.append(f'func (api *MdApi) {method_name}({param_str}) int32 {{')
            # 找到字符串数组参数的位置
            array_param_idx = call_args.index('_PLACEHOLDER_STRING_ARRAY_')
            array_param_name = params[array_param_idx - 1].split()[0]  # 获取参数名
            lines.append(f'\tif len({array_param_name}) == 0 {{')
            lines.append('\t\treturn 0')
            lines.append('\t}')
            lines.append(f'\t// 将字符串数组转换为 C 字符串数组')
            lines.append(f'\tptrs, _ := CStringArray({array_param_name})')
            # 替换占位符
            call_args[array_param_idx] = 'ptrs'
            call_str = ', '.join(call_args)
            lines.append(f'\treturn _{func.name}({call_str})')
        elif ret_type:
            call_str = ', '.join(call_args)
            lines.append(f'func (api *MdApi) {method_name}({param_str}) {ret_type} {{')
            lines.append(f'\treturn _{func.name}({call_str})')
        else:
            call_str = ', '.join(call_args)
            lines.append(f'func (api *MdApi) {method_name}({param_str}) {{')
            lines.append(f'\t_{func.name}({call_str})')
        
        lines.append('}')
        lines.append('')
    
    # Release 方法已经在 API 方法生成中处理，不需要重复生成
    
    # 生成 SpiSetOnXxx 方法（用于单独设置某个回调）
    lines.append('// ========== SPI 回调设置方法 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 方法名：SpiSetOnXxx
        callback_suffix = cb.go_method_name[2:]  # 移除 "On" 前缀，如 OnFrontConnected -> FrontConnected
        method_name = f'SpiSetOn{callback_suffix}'
        
        # 注释
        comment = cb.comment if cb.comment else ''
        lines.append(f'// {method_name} {comment}')
        lines.append(f'func (api *MdApi) {method_name}(callback {cb.name}) {{')
        lines.append('\t// 将函数类型转换为 uintptr')
        lines.append('\tptr := *(*uintptr)(unsafe.Pointer(&callback))')
        lines.append(f'\t_MdSpiSetOn{callback_suffix}(api.spiHandle, ptr)')
        lines.append('}')
        lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('// SetSpi 设置回调接口')
    lines.append('// 此方法会创建 C SPI 实例，注册 Go 回调函数，并将 SPI 注册到 API')
    lines.append('func (api *MdApi) SetSpi(spi MdSpi) {')
    lines.append('\tapi.mu.Lock()')
    lines.append('\tdefer api.mu.Unlock()')
    lines.append('\tapi.spi = spi')
    lines.append('')
    lines.append('\t// 如果已有 C SPI 实例，先销毁')
    lines.append('\tif api.spiHandle != 0 {')
    lines.append('\t\t_MdSpiDestroy(api.spiHandle)')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 创建新的 C SPI 实例')
    lines.append('\tapi.spiHandle = _MdSpiCreate(api.userData)')
    lines.append('')
    lines.append('\t// 注册所有回调函数到 C SPI')
    lines.append('\t// 使用回调文件中提供的辅助函数来获取函数指针（这些函数会包装 //export 函数以匹配正确的签名）')
    
    # 生成每个回调的设置调用
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        callback_name = cb.go_method_name
        lines.append(f'\t_MdSpiSetOn{callback_name[2:]}(api.spiHandle, GetGoMdOn{callback_name[2:]}())')
    
    lines.append('')
    lines.append('\t// 将 C SPI 注册到 API')
    lines.append('\t_MdRegisterSpi(api.handle, api.spiHandle)')
    lines.append('}')
    lines.append('')
    
    return '\n'.join(lines)


def generate_trader_api_go(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_api.go - 交易 API 封装
    
    结构与 md_api.go 类似，包含回调定义、接口、结构体、API方法等
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 交易 API 封装')
    lines.append('')
    lines.append('import (')
    lines.append('\t"fmt"')
    lines.append('\t"os"')
    lines.append('\t"path/filepath"')
    lines.append('\t"runtime"')
    lines.append('\t"sync"')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
    lines.append('')
    
    # 生成回调类型定义
    lines.append('// ========== 回调类型定义 ==========')
    lines.append('')
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成回调函数类型
        callback_params = []
        for p in cb.params:
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                callback_params.append(f'{p.name} {go_type}')
            else:
                callback_params.append(go_type)
        
        param_str = ', '.join(callback_params)
        lines.append(f'// {cb.name} {cb.comment}' if cb.comment else f'// {cb.name}')
        lines.append(f'type {cb.name} func({param_str})')
        lines.append('')
    
    # 生成回调结构体
    lines.append('// TraderSpiCallbacks 回调结构体（用于批量设置）')
    lines.append('type TraderSpiCallbacks struct {')
    lines.append('\tUserData uintptr')
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        field_name = cb.name.replace('Trader', '').replace('Callback', '')
        lines.append(f'\t{field_name} {cb.name}')
    lines.append('}')
    lines.append('')
    
    # 生成 SPI 接口
    lines.append('// ========== TraderSpi 接口 ==========')
    lines.append('')
    lines.append('// TraderSpi 交易回调接口')
    lines.append('type TraderSpi interface {')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成方法签名
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name} {go_type}')
            else:
                method_params.append(go_type)
        
        param_str = ', '.join(method_params)
        comment = f' // {cb.comment}' if cb.comment else ''
        lines.append(f'\t{cb.go_method_name}({param_str}){comment}')
    
    lines.append('}')
    lines.append('')
    
    # 生成 TraderApi 结构体
    lines.append('// ========== TraderApi 结构体 ==========')
    lines.append('')
    lines.append('// TraderApi 交易 API 封装')
    lines.append('type TraderApi struct {')
    lines.append('\thandle    uintptr')
    lines.append('\tspi       TraderSpi')
    lines.append('\tspiHandle uintptr // C SPI 实例句柄')
    lines.append('\tuserData  uintptr')
    lines.append('\tmu        sync.RWMutex')
    lines.append('\tflowPath  []byte // 保存 flowPath 的 C 字符串，防止被 GC 回收')
    lines.append('}')
    lines.append('')
    
    # 生成函数变量声明
    lines.append('// ========== C 函数声明 ==========')
    lines.append('')
    lines.append('var (')
    lines.append('\ttraderOnce sync.Once')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Trader') and not func.name.startswith('DC'):
            continue
        var_name = f'_{func.name}'
        
        # 生成参数类型
        param_types = []
        for i, p in enumerate(func.params):
            # 特殊处理 TraderSpiSetOnXxx 函数：第二个参数（回调函数指针）应该是 uintptr
            if func.name.startswith('TraderSpiSetOn') and i == 1:
                param_types.append('uintptr')
            else:
                go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs, p.is_array)
                param_types.append(go_type)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成函数签名（统一使用单行格式）
        param_str = ', '.join(param_types) if param_types else ''
        if param_str:
            if ret_type:
                lines.append(f'\t{var_name} func({param_str}) {ret_type}')
            else:
                lines.append(f'\t{var_name} func({param_str})')
        else:
            if ret_type:
                lines.append(f'\t{var_name} func() {ret_type}')
            else:
                lines.append(f'\t{var_name} func()')
    
    lines.append(')')
    lines.append('')
    
    # 生成初始化函数
    lines.append('// initTraderApi 初始化交易 API 函数')
    lines.append('func initTraderApi(lib uintptr) {')
    lines.append('\ttraderOnce.Do(func() {')
    
    for func in functions:
        if not func.name.startswith('Trader') and not func.name.startswith('DC'):
            continue
        var_name = f'_{func.name}'
        lines.append(f'\t\tpurego.RegisterLibFunc(&{var_name}, lib, "{func.name}")')
    
    lines.append('\t})')
    lines.append('}')
    lines.append('')
    
    # 生成实例管理
    lines.append('// ========== 实例管理 ==========')
    lines.append('')
    lines.append('var (')
    lines.append('\ttraderInstances   = make(map[uintptr]*TraderApi)')
    lines.append('\ttraderInstancesMu sync.RWMutex')
    lines.append('\ttraderNextID      uintptr = 1')
    lines.append(')')
    lines.append('')
    
    lines.append('func registerTraderInstance(api *TraderApi) uintptr {')
    lines.append('\ttraderInstancesMu.Lock()')
    lines.append('\tdefer traderInstancesMu.Unlock()')
    lines.append('\tid := traderNextID')
    lines.append('\ttraderNextID++')
    lines.append('\ttraderInstances[id] = api')
    lines.append('\treturn id')
    lines.append('}')
    lines.append('')
    
    lines.append('func getTraderInstance(userData uintptr) *TraderApi {')
    lines.append('\ttraderInstancesMu.RLock()')
    lines.append('\tdefer traderInstancesMu.RUnlock()')
    lines.append('\treturn traderInstances[userData]')
    lines.append('}')
    lines.append('')
    
    lines.append('func unregisterTraderInstance(userData uintptr) {')
    lines.append('\ttraderInstancesMu.Lock()')
    lines.append('\tdefer traderInstancesMu.Unlock()')
    lines.append('\tdelete(traderInstances, userData)')
    lines.append('}')
    lines.append('')
    
    # 生成构造函数
    lines.append('// ========== 构造函数 ==========')
    lines.append('')
    lines.append('// NewTraderApi 创建交易 API 实例')
    lines.append('// 首次调用时会自动加载 CTP 库（如果尚未加载）')
    lines.append('func NewTraderApi(flowPath string) *TraderApi {')
    lines.append('\t// 自动加载库（如果尚未加载）')
    lines.append('\tif err := autoLoadLibrary(); err != nil {')
    lines.append('\t\t// 如果自动加载失败，返回 nil（或者可以 panic，取决于设计）')
    lines.append('\t\t// 这里返回 nil，让调用者检查')
    lines.append('\t\treturn nil')
    lines.append('\t}')
    lines.append('')
    lines.append('\tapi := &TraderApi{}')
    lines.append('\tapi.userData = registerTraderInstance(api)')
    lines.append('')
    lines.append('\t// 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录')
    lines.append('\tabsFlowPath := flowPath')
    lines.append('\tif !filepath.IsAbs(flowPath) {')
    lines.append('\t\tvar err error')
    lines.append('\t\tabsFlowPath, err = filepath.Abs(flowPath)')
    lines.append('\t\tif err != nil {')
    lines.append('\t\t\tabsFlowPath = flowPath')
    lines.append('\t\t}')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 确保路径以路径分隔符结尾')
    lines.append('\tif len(absFlowPath) > 0 && absFlowPath[len(absFlowPath)-1] != filepath.Separator {')
    lines.append('\t\tabsFlowPath += string(filepath.Separator)')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 确保目录存在（CTP API 需要这个目录来创建 flow 文件）')
    lines.append('\tif err := os.MkdirAll(absFlowPath, 0755); err != nil {')
    lines.append('\t\t// 如果创建目录失败，记录错误但继续（CTP API 可能会自己创建）')
    lines.append('\t\tfmt.Printf("警告: 无法创建 flow 目录 %s: %v\\n", absFlowPath, err)')
    lines.append('\t\t// 这里不返回错误，让 CTP API 自己处理')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收')
    lines.append('\tapi.flowPath = make([]byte, len(absFlowPath)+1)')
    lines.append('\tcopy(api.flowPath, absFlowPath)')
    lines.append('\tapi.flowPath[len(absFlowPath)] = 0')
    lines.append('\tpathPtr := &api.flowPath[0]')
    lines.append('')
    lines.append('\tapi.handle = _TraderCreateFtdcTraderApi(pathPtr)')
    lines.append('')
    lines.append('\truntime.SetFinalizer(api, (*TraderApi).Release)')
    lines.append('\treturn api')
    lines.append('}')
    lines.append('')
    
    # 生成 API 方法
    lines.append('// ========== API 方法 ==========')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Trader'):
            continue
        if func.name == 'TraderCreateFtdcTraderApi':
            continue  # 已经在 NewTraderApi 中处理
        if func.name.startswith('TraderSpiSetOn'):
            continue  # 在单独的 SPI 回调设置方法部分生成

        # 方法名（移除 Trader 前缀）
        method_name = func.name[6:]
        
        # 生成参数
        params = []
        call_args = ['api.handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            
            # 清理参数名（移除可能的 [] 后缀）
            param_name = p.name.replace('[]', '').strip() if p.name else ''
            
            # 处理字符串参数
            if p.is_array and p.type == 'char' and p.is_pointer:
                # char* [] 或 char** 类型（字符串数组）
                params.append(f'{param_name} []string')
                # 将在方法体中特殊处理
                call_args.append('_PLACEHOLDER_STRING_ARRAY_')
            elif p.type == 'char' and p.is_pointer and not p.is_array:
                params.append(f'{param_name} string')
                call_args.append(f'CString({param_name})')
            else:
                if param_name:
                    params.append(f'{param_name} {go_type}')
                    call_args.append(param_name)
        
        param_str = ', '.join(params)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法
        comment = f'// {method_name} {func.comment}' if func.comment else f'// {method_name}'
        lines.append(comment)
        
        # 特殊处理方法
        if method_name == 'Release':
            # Release 方法需要先调用 C 函数，然后注销实例
            lines.append(f'func (api *TraderApi) {method_name}() {{')
            lines.append(f'\t_{func.name}(api.handle)')
            lines.append('\tunregisterTraderInstance(api.userData)')
        elif method_name == 'GetApiVersion':
            lines.append(f'func (api *TraderApi) {method_name}() string {{')
            lines.append(f'\tptr := _{func.name}()')
            lines.append('\tif ptr == nil {')
            lines.append('\t\treturn ""')
            lines.append('\t}')
            lines.append('\treturn GoString(ptr)')
        elif method_name == 'GetTradingDay':
            lines.append(f'func (api *TraderApi) {method_name}() string {{')
            lines.append(f'\tptr := _{func.name}(api.handle)')
            lines.append('\tif ptr == nil {')
            lines.append('\t\treturn ""')
            lines.append('\t}')
            lines.append('\treturn GoString(ptr)')
        elif '_PLACEHOLDER_STRING_ARRAY_' in call_args:
            # 处理字符串数组参数
            lines.append(f'func (api *TraderApi) {method_name}({param_str}) int32 {{')
            # 找到字符串数组参数的位置
            array_param_idx = call_args.index('_PLACEHOLDER_STRING_ARRAY_')
            array_param_name = params[array_param_idx - 1].split()[0]  # 获取参数名
            lines.append(f'\tif len({array_param_name}) == 0 {{')
            lines.append('\t\treturn 0')
            lines.append('\t}')
            lines.append(f'\t// 将字符串数组转换为 C 字符串数组')
            lines.append(f'\tptrs, _ := CStringArray({array_param_name})')
            # 替换占位符
            call_args[array_param_idx] = 'ptrs'
            call_str = ', '.join(call_args)
            lines.append(f'\treturn _{func.name}({call_str})')
        elif ret_type:
            call_str = ', '.join(call_args)
            lines.append(f'func (api *TraderApi) {method_name}({param_str}) {ret_type} {{')
            lines.append(f'\treturn _{func.name}({call_str})')
        else:
            call_str = ', '.join(call_args)
            lines.append(f'func (api *TraderApi) {method_name}({param_str}) {{')
            lines.append(f'\t_{func.name}({call_str})')
        
        lines.append('}')
        lines.append('')
    
    # Release 方法已经在 API 方法生成中处理，不需要重复生成
    
    # 生成 SpiSetOnXxx 方法（用于单独设置某个回调）
    lines.append('// ========== SPI 回调设置方法 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 方法名：SpiSetOnXxx
        callback_suffix = cb.go_method_name[2:]  # 移除 "On" 前缀，如 OnFrontConnected -> FrontConnected
        method_name = f'SpiSetOn{callback_suffix}'
        
        # 注释
        comment = cb.comment if cb.comment else ''
        lines.append(f'// {method_name} {comment}')
        lines.append(f'func (api *TraderApi) {method_name}(callback {cb.name}) {{')
        lines.append('\t// 将函数类型转换为 uintptr')
        lines.append('\tptr := *(*uintptr)(unsafe.Pointer(&callback))')
        lines.append(f'\t_TraderSpiSetOn{callback_suffix}(api.spiHandle, ptr)')
        lines.append('}')
        lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('// SetSpi 设置回调接口')
    lines.append('// 此方法会创建 C SPI 实例，注册 Go 回调函数，并将 SPI 注册到 API')
    lines.append('func (api *TraderApi) SetSpi(spi TraderSpi) {')
    lines.append('\tapi.mu.Lock()')
    lines.append('\tdefer api.mu.Unlock()')
    lines.append('\tapi.spi = spi')
    lines.append('')
    lines.append('\t// 如果已有 C SPI 实例，先销毁')
    lines.append('\tif api.spiHandle != 0 {')
    lines.append('\t\t_TraderSpiDestroy(api.spiHandle)')
    lines.append('\t}')
    lines.append('')
    lines.append('\t// 创建新的 C SPI 实例')
    lines.append('\tapi.spiHandle = _TraderSpiCreate(api.userData)')
    lines.append('')
    lines.append('\t// 注册所有回调函数到 C SPI')
    
    # 生成每个回调的设置调用
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        callback_name = cb.go_method_name
        lines.append(f'\t_TraderSpiSetOn{callback_name[2:]}(api.spiHandle, GetGoTraderOn{callback_name[2:]}())')
    
    lines.append('')
    lines.append('\t// 将 C SPI 注册到 API')
    lines.append('\t_TraderRegisterSpi(api.handle, api.spiHandle)')
    lines.append('}')
    lines.append('')
    
    # 生成 DataCollect 函数（DC 开头的独立函数）
    dc_functions = [f for f in functions if f.name.startswith('DC')]
    if dc_functions:
        lines.append('// ========== DataCollect 函数 ==========')
        lines.append('')
        
        for func in dc_functions:
            # 公开函数名去掉 DC 前缀
            public_name = func.name[2:] if func.name.startswith('DC') else func.name
            comment = f'// {public_name} {func.comment}' if func.comment else f'// {public_name}'
            lines.append(comment)
            
            if func.name == 'DCGetDataCollectApiVersion':
                # 返回字符串的函数
                lines.append(f'func {public_name}() string {{')
                lines.append(f'\tptr := _{func.name}()')
                lines.append('\tif ptr == nil {')
                lines.append('\t\treturn ""')
                lines.append('\t}')
                lines.append('\treturn GoString(ptr)')
                lines.append('}')
            elif func.name in ('DCGetSystemInfo', 'DCGetSystemInfoUnAesEncode'):
                # 获取系统信息的函数，返回 ([]byte, error) 更符合 Go 风格
                lines.append(f'func {public_name}() ([]byte, int32) {{')
                lines.append('\t// 分配至少 270 字节的缓冲区')
                lines.append('\tbuf := make([]byte, 512)')
                lines.append('\tlen := int32(len(buf))')
                lines.append(f'\tret := _{func.name}(&buf[0], &len)')
                lines.append('\tif ret != 0 {')
                lines.append('\t\treturn nil, ret')
                lines.append('\t}')
                lines.append('\treturn buf[:len], 0')
                lines.append('}')
            else:
                # 其他 DC 函数（未来可能添加的）
                params = []
                call_args = []
                for p in func.params:
                    go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
                    if p.name:
                        params.append(f'{p.name} {go_type}')
                        call_args.append(p.name)
                
                param_str = ', '.join(params)
                ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
                
                if ret_type:
                    lines.append(f'func {public_name}({param_str}) {ret_type} {{')
                    call_str = ', '.join(call_args)
                    lines.append(f'\treturn _{func.name}({call_str})')
                else:
                    lines.append(f'func {public_name}({param_str}) {{')
                    call_str = ', '.join(call_args)
                    lines.append(f'\t_{func.name}({call_str})')
                lines.append('')
            lines.append('')
    
    return '\n'.join(lines)


def generate_md_callbacks_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_callbacks.go - 行情回调实现
    
    使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器
    
    生成内容：
    1. goMdOnXxx 回调函数（使用 unsafe.Pointer 接收 CTP Field 指针）
    2. GetGoMdOnXxx 辅助函数（使用 purego.NewCallback 获取 C 函数指针）
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 行情回调实现')
    lines.append('// 使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器')
    lines.append('// 注意：Windows 的 syscall.NewCallback 要求回调函数必须返回 uintptr')
    lines.append('')
    lines.append('import (')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
    lines.append('')
    
    # 生成回调函数
    lines.append('// ========== 回调函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成 Go 回调函数
        # 对于 CTP Field 类型，使用 unsafe.Pointer 作为参数，然后在函数内部转换
        go_params = ['userData uintptr']
        call_args = []
        for p in cb.params[1:]:
            if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                # CTP Field 类型：使用 unsafe.Pointer 作为参数
                go_params.append(f'{p.name} unsafe.Pointer')
                # 在函数调用时转换为 Go 类型
                call_args.append(f'(*{p.type})({p.name})')
            else:
                go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
                if p.name:
                    go_params.append(f'{p.name} {go_type}')
                    call_args.append(p.name)
        
        param_str = ', '.join(go_params)
        func_name = f'goMd{cb.go_method_name}'
        
        comment = f'// {func_name} 回调函数实现'
        if any(p.type.startswith('CThostFtdc') for p in cb.params[1:]):
            comment += '（C 调用约定版本）'
        lines.append(comment)
        # Windows 的 syscall.NewCallback 要求回调函数必须返回 uintptr
        lines.append(f'func {func_name}({param_str}) uintptr {{')
        lines.append('\tapi := getMdInstance(userData)')
        lines.append('\tif api == nil || api.spi == nil {')
        lines.append('\t\treturn 0')
        lines.append('\t}')
        
        # 调用 SPI 方法
        call_str = ', '.join(call_args) if call_args else ''
        lines.append(f'\tapi.spi.{cb.go_method_name}({call_str})')
        lines.append('\treturn 0')
        lines.append('}')
        lines.append('')
    
    # 生成 GetGoMdOnXxx 辅助函数
    lines.append('// ========== 辅助函数：使用 purego.NewCallback 获取 C 函数指针 ==========')
    lines.append('// 这些函数使用 purego.NewCallback 将 Go 函数转换为 C 函数指针，无需 CGO')
    lines.append('// purego.NewCallback 返回 uintptr，需要转换为函数类型')
    lines.append('// 注意：purego.NewCallback 不支持 unsafe.Pointer 参数，需要用具体指针类型的 wrapper')
    lines.append('// 注意：Windows 要求 wrapper 函数也必须返回 uintptr')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        func_name = f'goMd{cb.go_method_name}'
        getter_name = f'GetGoMd{cb.go_method_name}'
        
        # 检查是否有 CTP Field 指针参数（这些参数在 goMdOnXxx 中用的是 unsafe.Pointer）
        has_ctp_field_ptr = any(
            p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer
            for p in cb.params[1:]
        )
        
        lines.append(f'// {getter_name} 获取 {func_name} 的 C 函数指针')
        lines.append(f'func {getter_name}() uintptr {{')
        
        if has_ctp_field_ptr:
            # 需要 wrapper：purego.NewCallback 不支持 unsafe.Pointer，需要用具体指针类型
            # Windows 要求 wrapper 函数也必须返回 uintptr
            wrapper_params = ['userData uintptr']
            call_args = ['userData']
            for p in cb.params[1:]:
                if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                    # wrapper 用具体指针类型，调用时转为 unsafe.Pointer
                    wrapper_params.append(f'{p.name} *{p.type}')
                    call_args.append(f'unsafe.Pointer({p.name})')
                else:
                    go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
                    if p.name:
                        wrapper_params.append(f'{p.name} {go_type}')
                        call_args.append(p.name)
            
            wrapper_param_str = ', '.join(wrapper_params)
            call_arg_str = ', '.join(call_args)
            lines.append(f'\twrapper := func({wrapper_param_str}) uintptr {{')
            lines.append(f'\t\treturn {func_name}({call_arg_str})')
            lines.append('\t}')
            lines.append('\treturn purego.NewCallback(wrapper)')
        else:
            # 无 CTP Field 指针参数，直接使用
            lines.append(f'\treturn purego.NewCallback({func_name})')
        
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_trader_callbacks_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_callbacks.go - 交易回调实现
    
    使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 交易回调实现')
    lines.append('// 使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器')
    lines.append('// 注意：Windows 的 syscall.NewCallback 要求回调函数必须返回 uintptr')
    lines.append('')
    lines.append('import (')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
    lines.append('')
    
    # 生成回调函数
    lines.append('// ========== 回调函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成 Go 回调函数
        go_params = ['userData uintptr']
        call_args = []
        for p in cb.params[1:]:
            if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                go_params.append(f'{p.name} unsafe.Pointer')
                call_args.append(f'(*{p.type})({p.name})')
            else:
                go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
                if p.name:
                    go_params.append(f'{p.name} {go_type}')
                    call_args.append(p.name)
        
        param_str = ', '.join(go_params)
        func_name = f'goTrader{cb.go_method_name}'
        
        comment = f'// {func_name} 回调函数实现'
        if any(p.type.startswith('CThostFtdc') for p in cb.params[1:]):
            comment += '（C 调用约定版本）'
        lines.append(comment)
        # Windows 的 syscall.NewCallback 要求回调函数必须返回 uintptr
        lines.append(f'func {func_name}({param_str}) uintptr {{')
        lines.append('\tapi := getTraderInstance(userData)')
        lines.append('\tif api == nil || api.spi == nil {')
        lines.append('\t\treturn 0')
        lines.append('\t}')
        
        call_str = ', '.join(call_args) if call_args else ''
        lines.append(f'\tapi.spi.{cb.go_method_name}({call_str})')
        lines.append('\treturn 0')
        lines.append('}')
        lines.append('')
    
    # 生成 GetGoTraderOnXxx 辅助函数
    lines.append('// ========== 辅助函数：使用 purego.NewCallback 获取 C 函数指针 ==========')
    lines.append('// 注意：purego.NewCallback 不支持 unsafe.Pointer 参数，需要用具体指针类型的 wrapper')
    lines.append('// 注意：Windows 要求 wrapper 函数也必须返回 uintptr')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        func_name = f'goTrader{cb.go_method_name}'
        getter_name = f'GetGoTrader{cb.go_method_name}'
        
        # 检查是否有 CTP Field 指针参数
        has_ctp_field_ptr = any(
            p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer
            for p in cb.params[1:]
        )
        
        lines.append(f'// {getter_name} 获取 {func_name} 的 C 函数指针')
        lines.append(f'func {getter_name}() uintptr {{')
        
        if has_ctp_field_ptr:
            # 需要 wrapper：purego.NewCallback 不支持 unsafe.Pointer，需要用具体指针类型
            # Windows 要求 wrapper 函数也必须返回 uintptr
            wrapper_params = ['userData uintptr']
            call_args = ['userData']
            for p in cb.params[1:]:
                if p.type.startswith('CThostFtdc') and p.type.endswith('Field') and p.is_pointer:
                    wrapper_params.append(f'{p.name} *{p.type}')
                    call_args.append(f'unsafe.Pointer({p.name})')
                else:
                    go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
                    if p.name:
                        wrapper_params.append(f'{p.name} {go_type}')
                        call_args.append(p.name)
            
            wrapper_param_str = ', '.join(wrapper_params)
            call_arg_str = ', '.join(call_args)
            lines.append(f'\twrapper := func({wrapper_param_str}) uintptr {{')
            lines.append(f'\t\treturn {func_name}({call_arg_str})')
            lines.append('\t}')
            lines.append('\treturn purego.NewCallback(wrapper)')
        else:
            # 无 CTP Field 指针参数，直接使用
            lines.append(f'\treturn purego.NewCallback({func_name})')
        
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_loader_windows_go() -> str:
    """
    生成 loader_windows.go - Windows 平台动态库加载
    
    使用 syscall.LoadLibrary 加载 DLL，并通过 SetDllDirectory 设置依赖库搜索路径
    """
    return '''//go:build windows

package ctpgo

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procSetDllDirectory = kernel32.NewProc("SetDllDirectoryW")
)

// openLibrary 在 Windows 上加载动态库
// 需要先设置 DLL 搜索目录，确保能找到依赖库
func openLibrary(path string) (uintptr, error) {
	// 1. 转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return 0, err
	}

	// 2. 设置 DLL 搜索目录（让 Windows 能找到依赖库）
	dllDir := filepath.Dir(absPath)
	dllDirPtr, _ := syscall.UTF16PtrFromString(dllDir)
	procSetDllDirectory.Call(uintptr(unsafe.Pointer(dllDirPtr)))

	// 3. 加载库
	handle, err := syscall.LoadLibrary(absPath)
	if err != nil {
		return 0, err
	}

	return uintptr(handle), nil
}
'''


def generate_loader_unix_go() -> str:
    """
    生成 loader_unix.go - Unix 平台动态库加载
    
    使用 purego.Dlopen 加载动态库（支持 Linux、macOS、FreeBSD、NetBSD）
    """
    return '''//go:build darwin || freebsd || linux || netbsd

package ctpgo

import "github.com/ebitengine/purego"

// openLibrary 在 Unix 上使用 purego.Dlopen 加载动态库
func openLibrary(path string) (uintptr, error) {
	return purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
'''


def generate_loader_go() -> str:
    """
    生成 loader.go - 动态库加载
    
    包含功能：
    1. getSystemLibPaths() - 获取系统库路径列表
    2. defaultLibPaths - 默认库搜索路径
    3. LoadCTPLibrary() - 加载 CTP C 包装库
    4. GetMdLibHandle() / GetTraderLibHandle() - 获取库句柄
    5. autoLoadLibrary() - 自动加载库（只加载一次）
    """
    return '''package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 动态库加载

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	mdLib     uintptr
	traderLib uintptr
	loadOnce  sync.Once
	loadErr   error
)

// getSystemLibPaths 根据平台返回系统库路径列表
// 包括标准系统路径和从环境变量读取的路径（类似 Python 的 sys.path）
func getSystemLibPaths() []string {
	var paths []string

	// 从系统库路径环境变量中读取（类似 Python 的 sys.path）
	switch runtime.GOOS {
	case "linux":
		// LD_LIBRARY_PATH 是 Linux 的标准库路径环境变量
		if ldPath := os.Getenv("LD_LIBRARY_PATH"); ldPath != "" {
			for _, p := range strings.Split(ldPath, ":") {
				if p != "" {
					paths = append(paths, p)
				}
			}
		}
	case "darwin":
		// DYLD_LIBRARY_PATH 是 macOS 的库路径环境变量
		if dyldPath := os.Getenv("DYLD_LIBRARY_PATH"); dyldPath != "" {
			for _, p := range strings.Split(dyldPath, ":") {
				if p != "" {
					paths = append(paths, p)
				}
			}
		}
	case "windows":
		// Windows 使用 PATH 环境变量，但通常系统会自动搜索
		// 这里可以添加一些特定路径
	}

	// 添加标准系统路径
	switch runtime.GOOS {
	case "linux":
		paths = append(paths,
			"/usr/local/lib",            // 用户安装的库
			"/usr/lib",                  // 系统库
			"/usr/lib/x86_64-linux-gnu", // Debian/Ubuntu 64位
			"/usr/lib64",                // 某些发行版的 64 位库路径
			"/opt/ctp/lib",              // CTP 专用安装路径
			"/opt/lib",                  // 通用 opt 路径
		)
	case "darwin":
		paths = append(paths,
			"/usr/local/lib",    // Homebrew (Intel)
			"/opt/homebrew/lib", // Homebrew (Apple Silicon)
			"/opt/local/lib",    // MacPorts
			"/usr/lib",          // 系统库
			"/opt/ctp/lib",      // CTP 专用安装路径
		)
	case "windows":
		// Windows 通常通过 PATH 环境变量查找，但也可以添加一些常见路径
		programFiles := os.Getenv("ProgramFiles")
		programFilesX86 := os.Getenv("ProgramFiles(x86)")
		if programFiles != "" {
			paths = append(paths, filepath.Join(programFiles, "CTP", "lib"))
		}
		if programFilesX86 != "" {
			paths = append(paths, filepath.Join(programFilesX86, "CTP", "lib"))
		}
		paths = append(paths,
			`C:\Windows\System32`, // 系统目录
			`C:\CTP\lib`,          // 常见安装路径
		)
	}

	return paths
}

var (
	// 默认库路径列表，按优先级顺序尝试
	// 可以通过环境变量 CTP_LIB_PATH 覆盖，环境变量优先级最高
	defaultLibPaths = func() []string {
		paths := []string{
			"./libs",              // 当前目录下的 libs
			"../libs",             // 上一层级下的 libs
			"../../libs",          // 上两级目录下的 libs
			"./ctp-wrapper/libs",  // 项目根目录下的 ctp-wrapper/libs
			"../ctp-wrapper/libs", // 上一层级下的 ctp-wrapper/libs
		}
		// 添加系统路径
		paths = append(paths, getSystemLibPaths()...)
		return paths
	}()
)

// LoadCTPLibrary 从 C 包装库加载（包含回调支持）
// libPath 为 ctp_md_c_api 和 ctp_trader_c_api 库文件所在目录
func LoadCTPLibrary(libPath string) error {
	var mdLibName, traderLibName string

	switch runtime.GOOS {
	case "windows":
		mdLibName = "ctpmd_c_api.dll"
		traderLibName = "ctptrader_c_api.dll"
	case "linux":
		mdLibName = "libctpmd_c_api.so"
		traderLibName = "libctptrader_c_api.so"
	case "darwin":
		mdLibName = "libctpmd_c_api.dylib"
		traderLibName = "libctptrader_c_api.dylib"
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	mdPath := filepath.Join(libPath, mdLibName)
	traderPath := filepath.Join(libPath, traderLibName)

	// 检查文件是否存在
	if _, err := os.Stat(mdPath); err != nil {
		return fmt.Errorf("md C wrapper library not found: %s", mdPath)
	}
	if _, err := os.Stat(traderPath); err != nil {
		return fmt.Errorf("trader C wrapper library not found: %s", traderPath)
	}

	// 加载行情 C 包装库
	var err error
	mdLib, err = openLibrary(mdPath)
	if err != nil {
		return fmt.Errorf("failed to load md C wrapper library: %w", err)
	}

	// 加载交易 C 包装库
	traderLib, err = openLibrary(traderPath)
	if err != nil {
		return fmt.Errorf("failed to load trader C wrapper library: %w", err)
	}

	// 初始化 API 函数
	initMdApi(mdLib)
	initTraderApi(traderLib)

	return nil
}

// GetMdLibHandle 获取行情库句柄
func GetMdLibHandle() uintptr {
	return mdLib
}

// GetTraderLibHandle 获取交易库句柄
func GetTraderLibHandle() uintptr {
	return traderLib
}

// getExecutableDir 获取可执行文件所在目录
func getExecutableDir() string {
	execPath, err := os.Executable()
	if err != nil {
		return "."
	}
	// 解析符号链接，获取真实路径
	realPath, err := filepath.EvalSymlinks(execPath)
	if err != nil {
		return filepath.Dir(execPath)
	}
	return filepath.Dir(realPath)
}

// autoLoadLibrary 自动加载库（只加载一次）
// 优先使用环境变量 CTP_LIB_PATH，否则按顺序尝试默认路径列表
func autoLoadLibrary() error {
	loadOnce.Do(func() {
		// 优先使用环境变量
		libPath := os.Getenv("CTP_LIB_PATH")
		if libPath != "" {
			// 环境变量指定的路径，只尝试一次
			loadErr = LoadCTPLibrary(libPath)
			return
		}

		// 获取可执行文件所在目录
		execDir := getExecutableDir()

		// 构建基于可执行文件位置的路径列表
		execBasedPaths := []string{
			filepath.Join(execDir, "libs"),                      // 可执行文件同目录下的 libs
			filepath.Join(execDir, "..", "libs"),                // 可执行文件父目录下的 libs
			filepath.Join(execDir, "..", "..", "libs"),          // 上两级目录下的 libs
			filepath.Join(execDir, "ctp-wrapper", "libs"),       // 可执行文件同目录下的 ctp-wrapper/libs
			filepath.Join(execDir, "..", "ctp-wrapper", "libs"), // 可执行文件父目录下的 ctp-wrapper/libs
		}

		// 先尝试基于可执行文件位置的路径
		for _, path := range execBasedPaths {
			err := LoadCTPLibrary(path)
			if err == nil {
				// 加载成功，直接返回
				loadErr = nil
				return
			}
		}

		// 再按顺序尝试基于当前工作目录的默认路径列表
		for _, path := range defaultLibPaths {
			err := LoadCTPLibrary(path)
			if err == nil {
				// 加载成功，直接返回
				loadErr = nil
				return
			}
			// 加载失败，继续尝试下一个路径
			loadErr = err
		}
		// 所有路径都失败，返回最后一个错误
	})
	return loadErr
}
'''


def generate_md_default_spi_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 md_default_spi.go - 行情 SPI 默认空实现
    
    用于嵌入到自定义结构体中，只需实现需要的方法
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// 默认 SPI 空实现，可用于嵌入')
    lines.append('')
    lines.append('// DefaultMdSpi 默认行情回调实现（空实现）')
    lines.append('// 使用方式：嵌入到自定义结构体中，只需实现需要的方法')
    lines.append('// 例如：type MySpi struct { DefaultMdSpi }')
    lines.append('//')
    lines.append('//\tfunc (s *MySpi) OnRtnDepthMarketData(...) { ... }')
    lines.append('type DefaultMdSpi struct{}')
    lines.append('')
    
    # 生成 MdSpi 接口的所有空实现
    md_callbacks = [cb for cb in callbacks if cb.name.startswith('Md')]
    for cb in md_callbacks:
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name} {go_type}')
            else:
                method_params.append(go_type)
        param_str = ', '.join(method_params)
        lines.append(f'func (s *DefaultMdSpi) {cb.go_method_name}({param_str}) {{')
        lines.append('\t// 空实现')
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_trader_default_spi_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """
    生成 trader_default_spi.go - 交易 SPI 默认空实现
    
    用于嵌入到自定义结构体中，只需实现需要的方法
    """
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// 默认 SPI 空实现，可用于嵌入')
    lines.append('')
    lines.append('// DefaultTraderSpi 默认交易回调实现（空实现）')
    lines.append('// 使用方式：嵌入到自定义结构体中，只需实现需要的方法')
    lines.append('// 例如：type MySpi struct { DefaultTraderSpi }')
    lines.append('//')
    lines.append('//\tfunc (s *MySpi) OnRtnOrder(...) { ... }')
    lines.append('type DefaultTraderSpi struct{}')
    lines.append('')
    
    # 生成 TraderSpi 接口的所有空实现
    trader_callbacks = [cb for cb in callbacks if cb.name.startswith('Trader')]
    for cb in trader_callbacks:
        method_params = []
        for p in cb.params[1:]:  # 跳过 userData
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                method_params.append(f'{p.name} {go_type}')
            else:
                method_params.append(go_type)
        param_str = ', '.join(method_params)
        lines.append(f'func (s *DefaultTraderSpi) {cb.go_method_name}({param_str}) {{')
        lines.append('\t// 空实现')
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


# ========== 主函数 ==========

def main():
    parser = argparse.ArgumentParser(description='CTP C API 转 Go PureGo 包装代码生成器')
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
    
    # utils.go
    utils_file = output_dir / 'utils.go'
    utils_file.write_text(generate_utils_go(), encoding='utf-8')
    print(f"  生成 {utils_file}")
    
    # datatype.go - 类型别名、枚举常量定义（不包含结构体）
    datatype_file = output_dir / 'datatype.go'
    datatype_content = generate_datatype_go(typedefs, enums, defines)
    datatype_file.write_text(datatype_content, encoding='utf-8')
    print(f"  生成 {datatype_file}")
    
    # struct.go - 结构体定义（与 datatype.go 中的结构体部分相同）
    struct_file = output_dir / 'struct.go'
    struct_content = generate_struct_go(structs, typedefs)
    struct_file.write_text(struct_content, encoding='utf-8')
    print(f"  生成 {struct_file}")
    
    # loader.go
    loader_file = output_dir / 'loader.go'
    loader_file.write_text(generate_loader_go(), encoding='utf-8')
    print(f"  生成 {loader_file}")
    
    # loader_windows.go
    loader_windows_file = output_dir / 'loader_windows.go'
    loader_windows_file.write_text(generate_loader_windows_go(), encoding='utf-8')
    print(f"  生成 {loader_windows_file}")
    
    # loader_unix.go
    loader_unix_file = output_dir / 'loader_unix.go'
    loader_unix_file.write_text(generate_loader_unix_go(), encoding='utf-8')
    print(f"  生成 {loader_unix_file}")
    
    # md_api.go
    md_api_file = output_dir / 'md_api.go'
    md_api_file.write_text(generate_md_api_go(md_functions, md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_api_file}")
    
    # trader_api.go
    trader_api_file = output_dir / 'trader_api.go'
    trader_api_file.write_text(generate_trader_api_go(trader_functions, trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_api_file}")
    
    # md_callbacks.go
    md_callbacks_file = output_dir / 'md_callbacks.go'
    md_callbacks_file.write_text(generate_md_callbacks_go(md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_callbacks_file}")
    
    # trader_callbacks.go
    trader_callbacks_file = output_dir / 'trader_callbacks.go'
    trader_callbacks_file.write_text(generate_trader_callbacks_go(trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_callbacks_file}")
    
    # md_default_spi.go
    md_default_spi_file = output_dir / 'md_default_spi.go'
    md_default_spi_file.write_text(generate_md_default_spi_go(md_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {md_default_spi_file}")
    
    # trader_default_spi.go
    trader_default_spi_file = output_dir / 'trader_default_spi.go'
    trader_default_spi_file.write_text(generate_trader_default_spi_go(trader_callbacks, typedefs), encoding='utf-8')
    print(f"  生成 {trader_default_spi_file}")
    
    # go.mod
    go_mod_file = output_dir / 'go.mod'
    if not go_mod_file.exists():
        go_mod_content = '''module ctpgo

go 1.23.4

require (
	github.com/ebitengine/purego v0.9.1
	golang.org/x/text v0.14.0
)
'''
        go_mod_file.write_text(go_mod_content, encoding='utf-8')
        print(f"  生成 {go_mod_file}")
    
    print("\n代码生成完成!")
    print(f"生成的文件列表:")
    print(f"  - utils.go              : 工具函数")
    print(f"  - datatype.go           : 枚举和结构体定义 ({len(structs)} 个)")
    print(f"  - struct.go             : 结构体定义 ({len(structs)} 个)")
    print(f"  - loader_windows.go     : Windows 动态库加载")
    print(f"  - loader_unix.go        : Unix 动态库加载")
    print(f"  - loader.go             : 动态库加载")
    print(f"  - md_api.go             : 行情 API ({len(md_functions)} 个方法)")
    print(f"  - trader_api.go         : 交易 API ({len(trader_functions)} 个方法)")
    print(f"  - md_callbacks.go       : 行情回调 ({len(md_callbacks)} 个)")
    print(f"  - trader_callbacks.go   : 交易回调 ({len(trader_callbacks)} 个)")
    print(f"  - md_default_spi.go     : 行情默认 SPI 实现")
    print(f"  - trader_default_spi.go : 交易默认 SPI 实现")
    
    # 自动格式化生成的 Go 代码
    format_go_files(output_dir)


def format_go_files(output_dir: Path):
    """
    使用 gofmt 格式化生成的 Go 文件
    
    gofmt 会自动对齐结构体字段、import 等，使代码符合 Go 规范
    """
    # 检查 gofmt 是否可用
    gofmt_path = shutil.which('gofmt')
    if not gofmt_path:
        print("\n警告: gofmt 未找到，跳过代码格式化")
        print("  请确保 Go 已安装并添加到 PATH 环境变量")
        return
    
    print(f"\n使用 gofmt 格式化代码...")
    
    # 获取所有 .go 文件
    go_files = list(output_dir.glob('*.go'))
    
    if not go_files:
        print("  没有找到 .go 文件")
        return
    
    # 格式化每个文件
    success_count = 0
    for go_file in go_files:
        try:
            result = subprocess.run(
                ['gofmt', '-w', str(go_file)],
                capture_output=True,
                text=True
            )
            if result.returncode == 0:
                success_count += 1
            else:
                print(f"  格式化失败 {go_file.name}: {result.stderr}")
        except Exception as e:
            print(f"  格式化出错 {go_file.name}: {e}")
    
    print(f"  成功格式化 {success_count}/{len(go_files)} 个文件")


if __name__ == '__main__':
    main()