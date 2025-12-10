#!/usr/bin/env python3
"""
CTP C API 转 Go PureGo 包装代码生成器

功能：
- 解析 C API 头文件（ctp_trader_c_api.h, ctp_md_c_api.h）
- 生成使用 purego 的 Go 包装代码
- 支持多实例，使用 userData 机制

用法：
    python3 generate_go_api.py --input ../csrc --output ../ctpgo
"""

import re
import os
import argparse
from pathlib import Path
from dataclasses import dataclass
from typing import List, Optional, Dict, Tuple


@dataclass
class CParam:
    """C 函数参数"""
    type: str       # 参数类型（如 "MdApiHandle", "const char*"）
    name: str       # 参数名
    is_pointer: bool = False
    is_const: bool = False


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


def parse_c_header(header_path: Path) -> Tuple[List[CFunction], List[CallbackType], List[str]]:
    """解析 C 头文件，返回函数列表、回调类型列表和结构体前向声明列表"""
    
    try:
        content = header_path.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_path.read_text(encoding='gbk')
    
    functions = []
    callbacks = []
    structs = []
    
    # 提取结构体前向声明
    struct_pattern = r'struct\s+(\w+);'
    for m in re.finditer(struct_pattern, content):
        structs.append(m.group(1))
    
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
    
    return functions, callbacks, structs


def parse_params(param_str: str) -> List[CParam]:
    """解析参数列表字符串"""
    params = []
    if not param_str or param_str.strip() == "":
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
        
        # 特殊处理：const char* 类型
        if 'const' in param and 'char' in param and '*' in param:
            # 提取参数名（最后一个标识符）
            parts = param.split()
            name = parts[-1].replace('*', '')
            params.append(CParam(
                type="char",
                name=name,
                is_pointer=True,
                is_const=True
            ))
            continue
        
        # 找最后一个标识符作为参数名
        parts = param.replace('*', ' * ').split()
        if len(parts) >= 2:
            name = parts[-1]
            type_parts = parts[:-1]
            # 重组类型
            param_type = ""
            for i, p in enumerate(type_parts):
                if p == '*':
                    param_type = param_type.rstrip() + '*'
                else:
                    param_type += p + " "
            param_type = param_type.strip()
            
            params.append(CParam(
                type=param_type,
                name=name,
                is_pointer=is_pointer,
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
    if name.startswith("TraderOn"):
        name = name[7:]  # 移除 "Trader"
    elif name.startswith("MdOn"):
        name = name[2:]   # 移除 "Md"
    
    if name.endswith("Callback"):
        name = name[:-8]  # 移除 "Callback"
    
    return name


def c_type_to_go_type(c_type: str, is_pointer: bool, is_const: bool = False) -> str:
    """将 C 类型转换为 Go 类型（用于 cgo）"""
    c_type = c_type.strip()
    
    # 处理 const 关键字
    if is_const:
        c_type = c_type.replace("const", "").strip()
    
    # 移除类型中的指针符号
    c_type = c_type.replace("*", "").strip()
    
    # 处理指针
    if is_pointer:
        if c_type == "void" or c_type == "":
            return "unsafe.Pointer"
        elif c_type == "char":
            return "string"  # const char* 在 Go 中是 string
        elif c_type.startswith("struct "):
            struct_name = c_type.replace("struct ", "").strip()
            return f"*CThostFtdc{struct_name}"
        elif c_type.endswith("Handle"):
            # 句柄类型，在 Go 中用 unsafe.Pointer 表示
            return "unsafe.Pointer"
        else:
            # 其他指针类型
            return f"*{c_type}"
    else:
        # 非指针类型
        if c_type == "void" or c_type == "":
            return ""
        elif c_type == "int":
            return "int32"
        elif c_type == "bool":
            return "bool"
        elif c_type == "char":
            return "byte"
        elif c_type.startswith("struct "):
            struct_name = c_type.replace("struct ", "").strip()
            return f"CThostFtdc{struct_name}"
        elif c_type.endswith("Handle"):
            return "unsafe.Pointer"
        else:
            return c_type


def go_param_name(c_name: str) -> str:
    """将 C 参数名转换为 Go 风格（驼峰命名，首字母小写）"""
    if not c_name:
        return c_name
    
    # 如果已经是驼峰命名，直接返回
    if c_name[0].islower():
        return c_name
    
    # 转换为驼峰命名
    parts = re.split(r'[_-]', c_name)
    result = parts[0].lower()
    for part in parts[1:]:
        if part:
            result += part.capitalize()
    
    return result


def c_type_to_purego_signature(c_type: str, is_pointer: bool, is_const: bool = False) -> str:
    """将 C 类型转换为 PureGo 函数签名中的类型"""
    c_type = c_type.strip()
    
    if is_const:
        c_type = c_type.replace("const", "").strip()
    
    c_type = c_type.replace("*", "").strip()
    
    if is_pointer:
        if c_type == "void" or c_type == "":
            return "unsafe.Pointer"
        elif c_type == "char":
            return "*byte"  # purego 中字符串用 *byte
        elif c_type.startswith("struct "):
            struct_name = c_type.replace("struct ", "").strip()
            return f"*CThostFtdc{struct_name}"
        elif c_type.endswith("Handle"):
            return "unsafe.Pointer"
        else:
            return f"*{c_type}"
    else:
        if c_type == "void" or c_type == "":
            return ""
        elif c_type == "int":
            return "int32"
        elif c_type == "bool":
            return "bool"
        elif c_type == "char":
            return "byte"
        elif c_type.startswith("struct "):
            struct_name = c_type.replace("struct ", "").strip()
            return f"CThostFtdc{struct_name}"
        elif c_type.endswith("Handle"):
            return "unsafe.Pointer"
        else:
            return c_type


def generate_purego_wrapper(
    api_name: str,
    struct_name: str,
    init_func_name: str,
    functions: List[CFunction],
    callbacks: List[CallbackType],
    structs: List[str],
    header_file: str,
    lib_name: str
) -> str:
    """生成 purego 包装代码"""
    
    prefix_upper = api_name
    prefix_lower = api_name.lower()
    var_name = struct_name[0].lower()
    
    lines = []
    lines.append("package ctpgo")
    lines.append("")
    lines.append("// 此文件由代码生成器自动生成，请勿手动修改")
    lines.append(f"// 基于 {header_file} 生成")
    lines.append("")
    lines.append("import (")
    lines.append('\t"errors"')
    lines.append('\t"os"')
    lines.append('\t"runtime"')
    lines.append('\t"sync"')
    lines.append('\t"unsafe"')
    lines.append("")
    lines.append('\t"github.com/ebitengine/purego"')
    lines.append(")")
    lines.append("")
    
    # 生成库加载相关代码
    lines.append("var (")
    lines.append(f"\t{prefix_lower}Lib unsafe.Pointer")
    lines.append(f"\t{prefix_lower}LibOnce sync.Once")
    lines.append(f"\t{prefix_lower}LibErr error")
    lines.append(")")
    lines.append("")
    
    # 生成库加载函数
    lines.append(f"func load{prefix_upper}Lib() error {{")
    lines.append(f"\t{prefix_lower}LibOnce.Do(func() {{")
    lines.append(f'\t\tlibPath := os.Getenv("CTP_{prefix_upper.upper()}_LIB")')
    lines.append(f'\t\tif libPath == "" {{')
    lines.append('\t\t\t// 根据平台选择默认库名')
    lines.append('\t\t\tswitch runtime.GOOS {')
    lines.append('\t\t\tcase "linux":')
    lines.append(f'\t\t\t\tlibPath = "{lib_name.replace(".so", "").replace(".dylib", "").replace(".dll", "")}.so"')
    lines.append('\t\t\tcase "darwin":')
    lines.append(f'\t\t\t\tlibPath = "{lib_name.replace(".so", "").replace(".dylib", "").replace(".dll", "")}.dylib"')
    lines.append('\t\t\tcase "windows":')
    lines.append(f'\t\t\t\tlibPath = "{lib_name.replace(".so", "").replace(".dylib", "").replace(".dll", "")}.dll"')
    lines.append('\t\t\tdefault:')
    lines.append(f'\t\t\t\tlibPath = "{lib_name}"')
    lines.append('\t\t\t}')
    lines.append('\t\t}')
    lines.append(f'\t\t{prefix_lower}Lib, {prefix_lower}LibErr = purego.Dlopen(libPath, purego.RTLD_NOW)')
    lines.append(f'\t\tif {prefix_lower}LibErr != nil {{')
    lines.append(f'\t\t\t{prefix_lower}LibErr = errors.New("failed to load {prefix_lower} library from " + libPath + ": " + {prefix_lower}LibErr.Error())')
    lines.append('\t\t}')
    lines.append('\t})')
    lines.append(f'\treturn {prefix_lower}LibErr')
    lines.append("}")
    lines.append("")
    
    # 生成结构体定义
    lines.append(f"type {struct_name} struct {{")
    lines.append(f"\tapi\t\tunsafe.Pointer")
    lines.append(f"\tpSpi\t\tunsafe.Pointer")
    lines.append(f"\tversion\t\tstring")
    lines.append(f"\tpszFlowPath\tstring")
    
    if api_name == "Md":
        lines.append("\tusingUdp\tbool")
        lines.append("\tusingMulticast\tbool")
    
    lines.append("")
    lines.append("\t// 回调函数字段")
    
    # 生成回调函数字段
    for cb in callbacks:
        if cb.comment:
            lines.append(f"\t// {cb.comment}")
        
        cb_params = []
        for param in cb.params:
            if param.name == "userData":
                continue
            go_type = c_type_to_go_type(param.type, param.is_pointer, param.is_const)
            cb_params.append(f"{go_type}")
        
        cb_sig = ", ".join(cb_params) if cb_params else ""
        lines.append(f"\t{cb.go_method_name}_ func({cb_sig})")
    
    lines.append("}")
    lines.append("")
    
    # 生成 Init 函数
    if api_name == "Md":
        lines.append(f"func {init_func_name}(pszFlowPath string, usingUdp bool, usingMulticast bool) (*{struct_name}, error) {{")
    else:
        lines.append(f"func {init_func_name}(pszFlowPath string) (*{struct_name}, error) {{")
    
    lines.append(f'\tif err := load{prefix_upper}Lib(); err != nil {{')
    lines.append('\t\treturn nil, err')
    lines.append('\t}')
    lines.append("")
    lines.append(f"\t{var_name} := &{struct_name}{{")
    lines.append(f"\t\tpszFlowPath: pszFlowPath,")
    if api_name == "Md":
        lines.append("\t\tusingUdp: usingUdp,")
        lines.append("\t\tusingMulticast: usingMulticast,")
    lines.append("\t}")
    lines.append("")
    lines.append("\t// 执行目录下创建 log 目录")
    lines.append(f"\t_, err := os.Stat({var_name}.pszFlowPath)")
    lines.append("\tif err != nil {")
    lines.append(f"\t\tos.Mkdir({var_name}.pszFlowPath, os.ModePerm)")
    lines.append("\t}")
    lines.append("")
    lines.append(f"\tapi, err := {var_name}.CreateApi()")
    lines.append("\tif err != nil {")
    lines.append("\t\treturn nil, err")
    lines.append("\t}")
    lines.append(f"\t{var_name}.api = api")
    lines.append("")
    lines.append(f"\tspi, err := {var_name}.CreateSpi()")
    lines.append("\tif err != nil {")
    lines.append("\t\treturn nil, err")
    lines.append("\t}")
    lines.append(f"\t{var_name}.pSpi = spi")
    lines.append("")
    lines.append(f"\t{var_name}.version = {var_name}.GetApiVersion()")
    lines.append("")
    lines.append(f"\treturn {var_name}, nil")
    lines.append("}")
    lines.append("")
    
    # 生成函数绑定辅助函数
    lines.append(f"func bind{prefix_upper}Func(name string, fn interface{{}}) error {{")
    lines.append(f'\tif err := load{prefix_upper}Lib(); err != nil {{')
    lines.append('\t\treturn err')
    lines.append('\t}')
    lines.append(f'\treturn purego.RegisterFunc({prefix_lower}Lib, name, fn)')
    lines.append("}")
    lines.append("")
    
    # 生成各个 API 方法
    generate_purego_methods(lines, functions, callbacks, var_name, struct_name, prefix_upper, prefix_lower, api_name)
    
    return "\n".join(lines)


def generate_purego_methods(
    lines: List[str],
    functions: List[CFunction],
    callbacks: List[CallbackType],
    var_name: str,
    struct_name: str,
    prefix_upper: str,
    prefix_lower: str,
    api_name: str
):
    """生成 purego 方法"""
    
    # 查找关键函数
    api_methods = {}
    for func in functions:
        func_name_lower = func.name.lower()
        if f"{prefix_lower}create" in func_name_lower and "api" in func_name_lower and "ftdc" in func_name_lower:
            api_methods["CreateApi"] = func
        elif f"{prefix_lower}spicreate" in func_name_lower:
            api_methods["CreateSpi"] = func
        elif f"{prefix_lower}getapiversion" in func_name_lower:
            api_methods["GetApiVersion"] = func
        elif f"{prefix_lower}gettradingday" in func_name_lower:
            api_methods["GetTradingDay"] = func
        elif f"{prefix_lower}release" in func_name_lower:
            api_methods["Release"] = func
        elif f"{prefix_lower}init" in func_name_lower and len(func.params) == 1:
            api_methods["Init"] = func
        elif f"{prefix_lower}join" in func_name_lower:
            api_methods["Join"] = func
        elif f"{prefix_lower}registerfront" in func_name_lower:
            api_methods["RegisterFront"] = func
        elif f"{prefix_lower}registernameserver" in func_name_lower:
            api_methods["RegisterNameServer"] = func
        elif f"{prefix_lower}registerfensuserinfo" in func_name_lower:
            api_methods["RegisterFensUserInfo"] = func
        elif f"{prefix_lower}registerspi" in func_name_lower:
            api_methods["RegisterSpi"] = func
    
    # 生成 CreateApi
    if "CreateApi" in api_methods:
        func = api_methods["CreateApi"]
        lines.append(f"func ({var_name} *{struct_name}) CreateApi() (unsafe.Pointer, error) {{")
        if api_name == "Md":
            lines.append(f'\tvar fn func(*byte, bool, bool) unsafe.Pointer')
            lines.append(f'\tif err := bind{prefix_upper}Func("{func.name}", &fn); err != nil {{')
            lines.append('\t\treturn nil, err')
            lines.append('\t}')
            lines.append(f'\tpath := CString({var_name}.pszFlowPath)')
            lines.append(f'\tdefer CFree(path)')
            lines.append(f'\treturn fn(path, {var_name}.usingUdp, {var_name}.usingMulticast), nil')
        else:
            lines.append(f'\tvar fn func(*byte) unsafe.Pointer')
            lines.append(f'\tif err := bind{prefix_upper}Func("{func.name}", &fn); err != nil {{')
            lines.append('\t\treturn nil, err')
            lines.append('\t}')
            lines.append(f'\tpath := CString({var_name}.pszFlowPath)')
            lines.append(f'\tdefer CFree(path)')
            lines.append(f'\treturn fn(path), nil')
        lines.append("}")
        lines.append("")
    
    # 生成 CreateSpi
    if "CreateSpi" in api_methods:
        func = api_methods["CreateSpi"]
        lines.append(f"func ({var_name} *{struct_name}) CreateSpi() (unsafe.Pointer, error) {{")
        lines.append(f'\tvar fn func(unsafe.Pointer) unsafe.Pointer')
        lines.append(f'\tif err := bind{prefix_upper}Func("{func.name}", &fn); err != nil {{')
        lines.append('\t\treturn nil, err')
        lines.append('\t}')
        lines.append(f'\treturn fn(unsafe.Pointer({var_name})), nil')
        lines.append("}")
        lines.append("")
    
    # 生成其他方法（简化版，实际需要完整实现）
    # 这里只生成几个关键方法作为示例
    
    # 生成 GetApiVersion
    if "GetApiVersion" in api_methods:
        func = api_methods["GetApiVersion"]
        lines.append(f"func ({var_name} *{struct_name}) GetApiVersion() string {{")
        lines.append(f'\tvar fn func() *byte')
        lines.append(f'\tif err := bind{prefix_upper}Func("{func.name}", &fn); err != nil {{')
        lines.append('\t\treturn ""')
        lines.append('\t}')
        lines.append('\tptr := fn()')
        lines.append('\tif ptr == nil {')
        lines.append('\t\treturn ""')
        lines.append('\t}')
        lines.append('\treturn CGoString(ptr)')
        lines.append("}")
        lines.append("")
    
    # 生成辅助函数
    lines.append("// CString 分配 C 字符串（使用 C malloc）")
    lines.append("func CString(s string) *byte {")
    lines.append('\tif s == "" {')
    lines.append('\t\treturn nil')
    lines.append('\t}')
    lines.append('\t// 注意：这里简化处理，实际应该使用 C.malloc')
    lines.append('\t// purego 中需要手动管理内存')
    lines.append('\tbs := []byte(s)')
    lines.append('\tbs = append(bs, 0) // null terminator')
    lines.append('\treturn &bs[0]')
    lines.append("}")
    lines.append("")
    
    lines.append("// CFree 释放 C 字符串")
    lines.append("func CFree(p *byte) {")
    lines.append('\t// 注意：实际应该使用 C.free')
    lines.append('\t_ = p')
    lines.append("}")
    lines.append("")
    
    lines.append("// CGoString 将 C 字符串转换为 Go string")
    lines.append("func CGoString(ptr *byte) string {")
    lines.append('\tif ptr == nil {')
    lines.append('\t\treturn ""')
    lines.append('\t}')
    lines.append('\t// 找到 null terminator')
    lines.append('\tvar bs []byte')
    lines.append('\tfor p := ptr; *p != 0; p = (*byte)(unsafe.Pointer(uintptr(unsafe.Pointer(p)) + 1)) {')
    lines.append('\t\tbs = append(bs, *p)')
    lines.append('\t}')
    lines.append('\treturn string(bs)')
    lines.append("}")
    lines.append("")


def main():
    parser = argparse.ArgumentParser(description='CTP C API 转 Go PureGo 包装代码生成器')
    parser.add_argument('--input', '-i', default='../csrc',
                       help='C API 头文件目录 (默认: ../csrc)')
    parser.add_argument('--output', '-o', default='../ctpgo',
                       help='输出目录 (默认: ../ctpgo)')
    args = parser.parse_args()
    
    input_dir = Path(args.input)
    output_dir = Path(args.output)
    
    if not input_dir.exists():
        print(f"错误: 输入目录不存在 {input_dir}")
        return 1
    
    output_dir.mkdir(parents=True, exist_ok=True)
    
    print(f"输入目录: {input_dir}")
    print(f"输出目录: {output_dir}")
    print()
    
    # 处理 MdApi
    md_header = input_dir / "ctp_md_c_api.h"
    if md_header.exists():
        print(f"解析 {md_header}...")
        functions, callbacks, structs = parse_c_header(md_header)
        print(f"  找到 {len(functions)} 个函数")
        print(f"  找到 {len(callbacks)} 个回调类型")
        
        # 更新回调的 go_method_name
        for cb in callbacks:
            cb.go_method_name = extract_go_method_name(cb.name)
        
        go_code = generate_purego_wrapper(
            "Md", "Quote", "InitQuote",
            functions, callbacks, structs,
            "ctp_md_c_api.h", "libctpmd_c.so"
        )
        output_file = output_dir / "ctpquote_api_purego.go"
        output_file.write_text(go_code, encoding='utf-8')
        print(f"  生成 {output_file}")
    else:
        print(f"警告: 未找到 {md_header}")
    
    print()
    
    # 处理 TraderApi
    trader_header = input_dir / "ctp_trader_c_api.h"
    if trader_header.exists():
        print(f"解析 {trader_header}...")
        functions, callbacks, structs = parse_c_header(trader_header)
        print(f"  找到 {len(functions)} 个函数")
        print(f"  找到 {len(callbacks)} 个回调类型")
        
        # 更新回调的 go_method_name
        for cb in callbacks:
            cb.go_method_name = extract_go_method_name(cb.name)
        
        go_code = generate_purego_wrapper(
            "Trader", "Trade", "InitTrade",
            functions, callbacks, structs,
            "ctp_trader_c_api.h", "libctptrader_c.so"
        )
        output_file = output_dir / "ctptrade_api_purego.go"
        output_file.write_text(go_code, encoding='utf-8')
        print(f"  生成 {output_file}")
    else:
        print(f"警告: 未找到 {trader_header}")
    
    print()
    print("完成!")
    print()
    print("注意：")
    print("1. 生成的代码需要进一步完善，特别是字符串内存管理和所有方法的实现")
    print("2. 回调函数需要使用 purego.NewCallback 进行注册")
    print("3. 建议参考现有的 CGO 版本完善实现")
    return 0


if __name__ == '__main__':
    exit(main())
