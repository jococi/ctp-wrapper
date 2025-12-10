#!/usr/bin/env python3
"""
CTP C API 转 Go PureGo 包装代码生成器

功能：
- 解析 C API 头文件（ctp_trader_c_api.h, ctp_md_c_api.h）
- 解析 CTP 结构体和数据类型定义
- 生成使用 purego 的 Go 包装代码
- 支持多实例，使用 userData 机制

用法：
    python3 generate_go_api.py --input ../csrc --struct ../ctpapi/linux --output ../ctpgo
"""

import re
import os
import argparse
from pathlib import Path
from dataclasses import dataclass, field
from typing import List, Optional, Dict, Tuple, Set


# ========== 数据结构定义 ==========

@dataclass
class CParam:
    """C 函数参数"""
    type: str           # 参数类型（如 "MdApiHandle", "const char*"）
    name: str           # 参数名
    is_pointer: bool = False
    is_const: bool = False
    is_array: bool = False
    array_size: int = 0


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
    """C 类型定义"""
    name: str           # 类型名
    base_type: str      # 基础类型
    size: int = 0       # 数组大小（如果是数组类型）
    comment: str = ""


@dataclass 
class CStruct:
    """C 结构体定义"""
    name: str                   # 结构体名
    fields: List[CParam]        # 字段列表
    comment: str = ""


# ========== C 类型到 Go 类型映射 ==========

# CTP 数据类型映射
CTP_TYPE_MAP = {
    # 基础类型
    'char': 'byte',
    'int': 'int32',
    'short': 'int16',
    'double': 'float64',
    'float': 'float32',
    'bool': 'bool',
    'void': '',
    
    # CTP 特定类型 - 将在解析时动态添加
}


# ========== 解析函数 ==========

def parse_datatype_header(header_path: Path) -> Dict[str, CTypedef]:
    """解析 ThostFtdcUserApiDataType.h 获取类型定义"""
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    typedefs = {}
    
    # 匹配 typedef char TXxxType[N]; 或 typedef int TXxxType;
    # typedef char TThostFtdcTraderIDType[21];
    typedef_pattern = r'typedef\s+(\w+)\s+(\w+)(?:\[(\d+)\])?\s*;'
    
    for m in re.finditer(typedef_pattern, content):
        base_type = m.group(1)
        type_name = m.group(2)
        array_size = int(m.group(3)) if m.group(3) else 0
        
        # 提取注释
        start = m.start()
        preceding = content[:start]
        comment = extract_comment(preceding)
        
        typedefs[type_name] = CTypedef(
            name=type_name,
            base_type=base_type,
            size=array_size,
            comment=comment
        )
    
    return typedefs


def parse_struct_header(header_path: Path, typedefs: Dict[str, CTypedef]) -> Dict[str, CStruct]:
    """解析 ThostFtdcUserApiStruct.h 获取结构体定义"""
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    structs = {}
    
    # 匹配结构体定义
    # struct CThostFtdcXxxField { ... };
    struct_pattern = r'struct\s+(\w+)\s*\{([^}]+)\}'
    
    for m in re.finditer(struct_pattern, content, re.DOTALL):
        struct_name = m.group(1)
        body = m.group(2)
        
        # 提取结构体注释
        start = m.start()
        preceding = content[:start]
        struct_comment = extract_comment(preceding)
        
        fields = []
        
        # 解析字段
        # TThostFtdcTraderIDType  TraderID;
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
                
                # 提取字段注释
                field_comment = ""
                if i > 0:
                    prev_line = lines[i-1].strip()
                    if prev_line.startswith('///'):
                        field_comment = prev_line[3:].strip()
                
                # 判断是否是数组类型
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
                    array_size=array_size
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
        
        # 处理双指针 (char**)
        if '**' in param_clean:
            parts = param_clean.replace('**', ' ** ').split()
            parts = [p for p in parts if p]
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
    """从前置文本中提取最近的注释"""
    lines = preceding.split('\n')
    comment_lines = []
    
    for line in reversed(lines):
        line = line.strip()
        if line.startswith('//'):
            comment_text = line[2:].strip()
            comment_lines.insert(0, comment_text)
        elif line and not line.startswith('//'):
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

def c_type_to_go_type(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef]) -> str:
    """将 C 类型转换为 Go 类型"""
    c_type = c_type.strip().replace('const', '').strip()
    
    # 句柄类型
    if c_type.endswith('Handle'):
        return 'uintptr'
    
    # 检查是否是 CTP Field 类型
    if c_type.startswith('CThostFtdc') and c_type.endswith('Field'):
        if is_pointer:
            return f'*{c_type}'
        return c_type
    
    # 检查是否是 CTP typedef 类型
    if c_type in typedefs:
        td = typedefs[c_type]
        if td.size > 0:
            # 数组类型
            base_go = CTP_TYPE_MAP.get(td.base_type, td.base_type)
            return f'[{td.size}]{base_go}'
        else:
            # 非数组类型
            return CTP_TYPE_MAP.get(td.base_type, td.base_type)
    
    # 基础类型
    if is_pointer:
        if c_type == 'void':
            return 'uintptr'
        elif c_type == 'char':
            return '*byte'  # C 字符串
        else:
            go_type = CTP_TYPE_MAP.get(c_type, c_type)
            return f'*{go_type}'
    else:
        if c_type == 'void':
            return ''
        return CTP_TYPE_MAP.get(c_type, c_type)


def c_type_to_go_callback_param(c_type: str, is_pointer: bool, typedefs: Dict[str, CTypedef]) -> str:
    """将 C 类型转换为 Go 回调参数类型"""
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


def generate_types_go(structs: Dict[str, CStruct], typedefs: Dict[str, CTypedef]) -> str:
    """生成 ctp_types.go"""
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 结构体定义')
    lines.append('')
    
    # 生成常用类型别名
    lines.append('// ========== 类型别名 ==========')
    lines.append('')
    lines.append('// THOST_TE_RESUME_TYPE 订阅类型')
    lines.append('type THOST_TE_RESUME_TYPE int32')
    lines.append('')
    lines.append('const (')
    lines.append('\tTHOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0 // 从本交易日开始重传')
    lines.append('\tTHOST_TERT_RESUME  THOST_TE_RESUME_TYPE = 1 // 从上次收到的续传')
    lines.append('\tTHOST_TERT_QUICK   THOST_TE_RESUME_TYPE = 2 // 只传送登录后的流内容')
    lines.append('\tTHOST_TERT_NONE    THOST_TE_RESUME_TYPE = 3 // 不传送')
    lines.append(')')
    lines.append('')
    
    # 生成结构体
    lines.append('// ========== CTP 结构体 ==========')
    lines.append('')
    
    for struct_name, struct in sorted(structs.items()):
        if struct.comment:
            lines.append(f'// {struct_name} {struct.comment}')
        lines.append(f'type {struct_name} struct {{')
        
        for field in struct.fields:
            go_type = c_type_to_go_type(field.type, field.is_pointer, typedefs)
            field_comment = ""
            
            # 获取字段注释
            if field.type in typedefs and typedefs[field.type].comment:
                field_comment = f' // {typedefs[field.type].comment}'
            
            lines.append(f'\t{field.name} {go_type}{field_comment}')
        
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_md_api_go(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """生成 md_api.go"""
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 行情 API 封装')
    lines.append('')
    lines.append('import (')
    lines.append('\t"runtime"')
    lines.append('\t"sync"')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
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
    lines.append('\thandle   uintptr')
    lines.append('\tspi      MdSpi')
    lines.append('\tuserData uintptr')
    lines.append('\tmu       sync.RWMutex')
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
        lines.append(f'\t{var_name} func(')
        
        # 生成参数类型
        param_types = []
        for p in func.params:
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            param_types.append(go_type)
        
        lines.append(f'\t\t{", ".join(param_types)}')
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        if ret_type:
            lines.append(f'\t) {ret_type}')
        else:
            lines.append('\t)')
    
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
    lines.append('func NewMdApi(flowPath string, usingUdp, multicast bool) *MdApi {')
    lines.append('\tapi := &MdApi{}')
    lines.append('\tapi.userData = registerMdInstance(api)')
    lines.append('\t')
    lines.append('\tpathPtr := CString(flowPath)')
    lines.append('\tapi.handle = _MdCreateFtdcMdApi(pathPtr, usingUdp, multicast)')
    lines.append('\t')
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
        
        # 方法名（移除 Md 前缀）
        method_name = func.name[2:]
        
        # 生成参数
        params = []
        call_args = ['api.handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            
            # 处理字符串参数
            if p.type == 'char' and p.is_pointer:
                params.append(f'{p.name} string')
                call_args.append(f'CString({p.name})')
            elif p.is_array and p.type == 'char':
                # char** 类型（字符串数组）
                params.append(f'{p.name} []string')
                call_args.append(f'/* {p.name} */')  # 需要特殊处理
            else:
                if p.name:
                    params.append(f'{p.name} {go_type}')
                    call_args.append(p.name)
        
        param_str = ', '.join(params)
        call_str = ', '.join(call_args)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法
        comment = f'// {method_name} {func.comment}' if func.comment else f'// {method_name}'
        lines.append(comment)
        
        if ret_type:
            lines.append(f'func (api *MdApi) {method_name}({param_str}) {ret_type} {{')
            lines.append(f'\treturn _{func.name}({call_str})')
        else:
            lines.append(f'func (api *MdApi) {method_name}({param_str}) {{')
            lines.append(f'\t_{func.name}({call_str})')
        
        lines.append('}')
        lines.append('')
    
    # 生成 Release 方法
    lines.append('// Release 释放 API 实例')
    lines.append('func (api *MdApi) Release() {')
    lines.append('\tif api.handle != 0 {')
    lines.append('\t\t_MdRelease(api.handle)')
    lines.append('\t\tunregisterMdInstance(api.userData)')
    lines.append('\t\tapi.handle = 0')
    lines.append('\t}')
    lines.append('}')
    lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('// SetSpi 设置回调接口')
    lines.append('func (api *MdApi) SetSpi(spi MdSpi) {')
    lines.append('\tapi.mu.Lock()')
    lines.append('\tdefer api.mu.Unlock()')
    lines.append('\tapi.spi = spi')
    lines.append('}')
    lines.append('')
    
    return '\n'.join(lines)


def generate_trader_api_go(functions: List[CFunction], callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """生成 trader_api.go"""
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 交易 API 封装')
    lines.append('')
    lines.append('import (')
    lines.append('\t"runtime"')
    lines.append('\t"sync"')
    lines.append('\t"unsafe"')
    lines.append('')
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(')')
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
    lines.append('\thandle   uintptr')
    lines.append('\tspi      TraderSpi')
    lines.append('\tuserData uintptr')
    lines.append('\tmu       sync.RWMutex')
    lines.append('}')
    lines.append('')
    
    # 生成函数变量声明
    lines.append('// ========== C 函数声明 ==========')
    lines.append('')
    lines.append('var (')
    lines.append('\ttraderOnce sync.Once')
    lines.append('')
    
    for func in functions:
        if not func.name.startswith('Trader'):
            continue
        var_name = f'_{func.name}'
        lines.append(f'\t{var_name} func(')
        
        # 生成参数类型
        param_types = []
        for p in func.params:
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            param_types.append(go_type)
        
        lines.append(f'\t\t{", ".join(param_types)}')
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        if ret_type:
            lines.append(f'\t) {ret_type}')
        else:
            lines.append('\t)')
    
    lines.append(')')
    lines.append('')
    
    # 生成初始化函数
    lines.append('// initTraderApi 初始化交易 API 函数')
    lines.append('func initTraderApi(lib uintptr) {')
    lines.append('\ttraderOnce.Do(func() {')
    
    for func in functions:
        if not func.name.startswith('Trader'):
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
    lines.append('func NewTraderApi(flowPath string) *TraderApi {')
    lines.append('\tapi := &TraderApi{}')
    lines.append('\tapi.userData = registerTraderInstance(api)')
    lines.append('\t')
    lines.append('\tpathPtr := CString(flowPath)')
    lines.append('\tapi.handle = _TraderCreateFtdcTraderApi(pathPtr)')
    lines.append('\t')
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
        
        # 方法名（移除 Trader 前缀）
        method_name = func.name[6:]
        
        # 生成参数
        params = []
        call_args = ['api.handle']
        
        for p in func.params[1:]:  # 跳过第一个 handle 参数
            go_type = c_type_to_go_type(p.type, p.is_pointer, typedefs)
            
            # 处理字符串参数
            if p.type == 'char' and p.is_pointer:
                params.append(f'{p.name} string')
                call_args.append(f'CString({p.name})')
            elif p.is_array and p.type == 'char':
                # char** 类型（字符串数组）
                params.append(f'{p.name} []string')
                call_args.append(f'/* {p.name} */')  # 需要特殊处理
            else:
                if p.name:
                    params.append(f'{p.name} {go_type}')
                    call_args.append(p.name)
        
        param_str = ', '.join(params)
        call_str = ', '.join(call_args)
        
        # 返回类型
        ret_type = c_type_to_go_type(func.return_type, '*' in func.return_type, typedefs)
        
        # 生成方法
        comment = f'// {method_name} {func.comment}' if func.comment else f'// {method_name}'
        lines.append(comment)
        
        if ret_type:
            lines.append(f'func (api *TraderApi) {method_name}({param_str}) {ret_type} {{')
            lines.append(f'\treturn _{func.name}({call_str})')
        else:
            lines.append(f'func (api *TraderApi) {method_name}({param_str}) {{')
            lines.append(f'\t_{func.name}({call_str})')
        
        lines.append('}')
        lines.append('')
    
    # 生成 Release 方法
    lines.append('// Release 释放 API 实例')
    lines.append('func (api *TraderApi) Release() {')
    lines.append('\tif api.handle != 0 {')
    lines.append('\t\t_TraderRelease(api.handle)')
    lines.append('\t\tunregisterTraderInstance(api.userData)')
    lines.append('\t\tapi.handle = 0')
    lines.append('\t}')
    lines.append('}')
    lines.append('')
    
    # 生成 SetSpi 方法
    lines.append('// SetSpi 设置回调接口')
    lines.append('func (api *TraderApi) SetSpi(spi TraderSpi) {')
    lines.append('\tapi.mu.Lock()')
    lines.append('\tdefer api.mu.Unlock()')
    lines.append('\tapi.spi = spi')
    lines.append('}')
    lines.append('')
    
    return '\n'.join(lines)


def generate_md_callbacks_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """生成 md_callbacks.go - 行情回调实现"""
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 行情回调实现')
    lines.append('')
    lines.append('import "unsafe"')
    lines.append('')
    lines.append('// #include <stdint.h>')
    lines.append('import "C"')
    lines.append('')
    
    # 生成回调函数
    lines.append('// ========== 回调函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Md'):
            continue
        
        # 生成 Go 回调函数
        go_params = ['userData uintptr']
        for p in cb.params[1:]:
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                go_params.append(f'{p.name} {go_type}')
        
        param_str = ', '.join(go_params)
        func_name = f'go{cb.go_method_name}'
        
        lines.append(f'//export {func_name}')
        lines.append(f'func {func_name}({param_str}) {{')
        lines.append('\tapi := getMdInstance(userData)')
        lines.append('\tif api == nil || api.spi == nil {')
        lines.append('\t\treturn')
        lines.append('\t}')
        
        # 调用 SPI 方法
        call_args = []
        for p in cb.params[1:]:
            if p.name:
                call_args.append(p.name)
        
        call_str = ', '.join(call_args)
        lines.append(f'\tapi.spi.{cb.go_method_name}({call_str})')
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_trader_callbacks_go(callbacks: List[CallbackType], typedefs: Dict[str, CTypedef]) -> str:
    """生成 trader_callbacks.go - 交易回调实现"""
    lines = []
    lines.append('package ctpgo')
    lines.append('')
    lines.append('// 此文件由代码生成器自动生成，请勿手动修改')
    lines.append('// CTP 交易回调实现')
    lines.append('')
    lines.append('import "unsafe"')
    lines.append('')
    lines.append('// #include <stdint.h>')
    lines.append('import "C"')
    lines.append('')
    
    # 生成回调函数
    lines.append('// ========== 回调函数 ==========')
    lines.append('')
    
    for cb in callbacks:
        if not cb.name.startswith('Trader'):
            continue
        
        # 生成 Go 回调函数
        go_params = ['userData uintptr']
        for p in cb.params[1:]:
            go_type = c_type_to_go_callback_param(p.type, p.is_pointer, typedefs)
            if p.name:
                go_params.append(f'{p.name} {go_type}')
        
        param_str = ', '.join(go_params)
        func_name = f'go{cb.go_method_name}'
        
        lines.append(f'//export {func_name}')
        lines.append(f'func {func_name}({param_str}) {{')
        lines.append('\tapi := getTraderInstance(userData)')
        lines.append('\tif api == nil || api.spi == nil {')
        lines.append('\t\treturn')
        lines.append('\t}')
        
        # 调用 SPI 方法
        call_args = []
        for p in cb.params[1:]:
            if p.name:
                call_args.append(p.name)
        
        call_str = ', '.join(call_args)
        lines.append(f'\tapi.spi.{cb.go_method_name}({call_str})')
        lines.append('}')
        lines.append('')
    
    return '\n'.join(lines)


def generate_loader_go() -> str:
    """生成 loader.go - 动态库加载"""
    return '''package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 动态库加载

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	mdLib     uintptr
	traderLib uintptr
)

// LoadLibrary 加载 CTP 动态库
// libPath 为库文件所在目录路径
func LoadLibrary(libPath string) error {
	var mdLibName, traderLibName string
	
	switch runtime.GOOS {
	case "windows":
		mdLibName = "thostmduserapi_se.dll"
		traderLibName = "thosttraderapi_se.dll"
	case "linux":
		mdLibName = "thostmduserapi_se.so"
		traderLibName = "thosttraderapi_se.so"
	case "darwin":
		mdLibName = "thostmduserapi_se.framework/thostmduserapi_se"
		traderLibName = "thosttraderapi_se.framework/thosttraderapi_se"
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	
	mdPath := filepath.Join(libPath, mdLibName)
	traderPath := filepath.Join(libPath, traderLibName)
	
	// 检查文件是否存在
	if _, err := os.Stat(mdPath); err != nil {
		return fmt.Errorf("md library not found: %s", mdPath)
	}
	if _, err := os.Stat(traderPath); err != nil {
		return fmt.Errorf("trader library not found: %s", traderPath)
	}
	
	// 加载行情库
	var err error
	mdLib, err = purego.Dlopen(mdPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load md library: %w", err)
	}
	
	// 加载交易库
	traderLib, err = purego.Dlopen(traderPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load trader library: %w", err)
	}
	
	// 初始化 API 函数
	initMdApi(mdLib)
	initTraderApi(traderLib)
	
	return nil
}

// LoadCTPLibrary 从 C 包装库加载（包含回调支持）
// libPath 为 ctp_md_c_api 和 ctp_trader_c_api 库文件所在目录
func LoadCTPLibrary(libPath string) error {
	var mdLibName, traderLibName string
	
	switch runtime.GOOS {
	case "windows":
		mdLibName = "ctp_md_c_api.dll"
		traderLibName = "ctp_trader_c_api.dll"
	case "linux":
		mdLibName = "libctp_md_c_api.so"
		traderLibName = "libctp_trader_c_api.so"
	case "darwin":
		mdLibName = "libctp_md_c_api.dylib"
		traderLibName = "libctp_trader_c_api.dylib"
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
	mdLib, err = purego.Dlopen(mdPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load md C wrapper library: %w", err)
	}
	
	// 加载交易 C 包装库
	traderLib, err = purego.Dlopen(traderPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
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
'''


def generate_default_spi_go() -> str:
    """生成 default_spi.go - 默认 SPI 实现"""
    return '''package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// 默认 SPI 空实现，可用于嵌入

// DefaultMdSpi 默认行情回调实现（空实现）
// 使用方式：嵌入到自定义结构体中，只需实现需要的方法
type DefaultMdSpi struct{}

// DefaultTraderSpi 默认交易回调实现（空实现）
// 使用方式：嵌入到自定义结构体中，只需实现需要的方法
type DefaultTraderSpi struct{}
'''


# ========== 主函数 ==========

def main():
    parser = argparse.ArgumentParser(description='CTP C API 转 Go PureGo 包装代码生成器')
    parser.add_argument('--input', required=True, help='C API 头文件目录 (包含 ctp_md_c_api.h, ctp_trader_c_api.h)')
    parser.add_argument('--struct', required=True, help='CTP 结构体头文件目录 (包含 ThostFtdcUserApiDataType.h, ThostFtdcUserApiStruct.h)')
    parser.add_argument('--output', required=True, help='输出目录')
    args = parser.parse_args()
    
    input_dir = Path(args.input)
    struct_dir = Path(args.struct)
    output_dir = Path(args.output)
    
    # 确保输出目录存在
    output_dir.mkdir(parents=True, exist_ok=True)
    
    print(f"解析类型定义: {struct_dir / 'ThostFtdcUserApiDataType.h'}")
    typedefs = parse_datatype_header(struct_dir / 'ThostFtdcUserApiDataType.h')
    print(f"  找到 {len(typedefs)} 个类型定义")
    
    print(f"解析结构体定义: {struct_dir / 'ThostFtdcUserApiStruct.h'}")
    structs = parse_struct_header(struct_dir / 'ThostFtdcUserApiStruct.h', typedefs)
    print(f"  找到 {len(structs)} 个结构体")
    
    # 解析行情 API
    md_header = input_dir / 'ctp_md_c_api.h'
    print(f"解析行情 API: {md_header}")
    md_functions, md_callbacks = parse_c_header(md_header)
    print(f"  找到 {len(md_functions)} 个函数, {len(md_callbacks)} 个回调")
    
    # 解析交易 API
    trader_header = input_dir / 'ctp_trader_c_api.h'
    print(f"解析交易 API: {trader_header}")
    trader_functions, trader_callbacks = parse_c_header(trader_header)
    print(f"  找到 {len(trader_functions)} 个函数, {len(trader_callbacks)} 个回调")
    
    # 生成代码文件
    print(f"\n生成代码到: {output_dir}")
    
    # utils.go
    utils_file = output_dir / 'utils.go'
    utils_file.write_text(generate_utils_go(), encoding='utf-8')
    print(f"  生成 {utils_file}")
    
    # ctp_types.go
    types_file = output_dir / 'ctp_types.go'
    types_file.write_text(generate_types_go(structs, typedefs), encoding='utf-8')
    print(f"  生成 {types_file}")
    
    # loader.go
    loader_file = output_dir / 'loader.go'
    loader_file.write_text(generate_loader_go(), encoding='utf-8')
    print(f"  生成 {loader_file}")
    
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
    
    # default_spi.go
    default_spi_file = output_dir / 'default_spi.go'
    default_spi_file.write_text(generate_default_spi_go(), encoding='utf-8')
    print(f"  生成 {default_spi_file}")
    
    # go.mod
    go_mod_file = output_dir / 'go.mod'
    if not go_mod_file.exists():
        go_mod_content = '''module ctpgo

go 1.21

require (
	github.com/ebitengine/purego v0.7.1
	golang.org/x/text v0.14.0
)
'''
        go_mod_file.write_text(go_mod_content, encoding='utf-8')
        print(f"  生成 {go_mod_file}")
    
    print("\n代码生成完成!")
    print(f"生成的文件列表:")
    print(f"  - utils.go          : 工具函数")
    print(f"  - ctp_types.go      : CTP 结构体定义 ({len(structs)} 个)")
    print(f"  - loader.go         : 动态库加载")
    print(f"  - md_api.go         : 行情 API ({len(md_functions)} 个方法)")
    print(f"  - trader_api.go     : 交易 API ({len(trader_functions)} 个方法)")
    print(f"  - md_callbacks.go   : 行情回调 ({len(md_callbacks)} 个)")
    print(f"  - trader_callbacks.go: 交易回调 ({len(trader_callbacks)} 个)")
    print(f"  - default_spi.go    : 默认 SPI 实现")


if __name__ == '__main__':
    main()