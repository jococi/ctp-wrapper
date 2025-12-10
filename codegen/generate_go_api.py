#!/usr/bin/env python3
"""
CTP C API 转 Go CGO 包装代码生成器

功能：
- 解析 C API 头文件（ctp_trader_c_api.h, ctp_md_c_api.h）
- 生成使用 cgo 的 Go 包装代码，参考现有 ctpquote_api_darwin.go 风格
- 保持与现有 API 的兼容性，减少上层应用改动

用法：
    python generate_go_api.py --input ../csrc --output ../ctpgo
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
    type: str       # 参数类型（如 "CtpTraderApiHandle", "const char*"）
    name: str       # 参数名
    is_pointer: bool = False
    is_const: bool = False


@dataclass
class CFunction:
    """C 函数定义"""
    name: str                   # 函数名（如 "ctpTraderCreateFtdcTraderApi"）
    return_type: str            # 返回类型
    params: List[CParam]        # 参数列表
    comment: str = ""           # 注释


@dataclass
class CallbackType:
    """回调函数类型"""
    name: str                   # 类型名（如 "CtpTraderOnFrontConnectedCallback"）
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
    # typedef void (*CtpTraderOnFrontConnectedCallback)(void* userData);
    callback_pattern = r'typedef\s+void\s+\(\*(\w+)\)\s*\(([^)]*)\)\s*;'
    for m in re.finditer(callback_pattern, content):
        callback_name = m.group(1)
        param_str = m.group(2)
        
        # 提取注释（前面的行）
        callback_start = m.start()
        preceding = content[:callback_start]
        comment = extract_comment(preceding)
        
        params = parse_params(param_str)
        
        # 提取 Go 方法名：CtpTraderOnFrontConnectedCallback -> OnFrontConnected
        go_method_name = extract_go_method_name(callback_name)
        
        callbacks.append(CallbackType(
            name=callback_name,
            params=params,
            comment=comment,
            go_method_name=go_method_name
        ))
    
    # 提取函数声明
    # CTP_API CtpTraderApiHandle ctpTraderCreateFtdcTraderApi(const char* pszFlowPath);
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
    """从回调类型名提取 Go 方法名
    
    例如：
    CtpTraderOnFrontConnectedCallback -> OnFrontConnected
    CtpMdOnRtnDepthMarketDataCallback -> OnRtnDepthMarketData
    """
    # 移除前缀（CtpTrader 或 CtpMd）和后缀（Callback）
    name = callback_name
    if name.startswith("CtpTrader"):
        name = name[10:]  # 移除 "CtpTrader"
    elif name.startswith("CtpMd"):
        name = name[5:]   # 移除 "CtpMd"
    
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


def generate_cgo_wrapper(
    api_name: str,              # "Trader" 或 "Md"
    struct_name: str,            # "Trade" 或 "Quote"
    init_func_name: str,         # "InitTrade" 或 "InitQuote"
    functions: List[CFunction],
    callbacks: List[CallbackType],
    structs: List[str],
    header_file: str,            # C 头文件名
    lib_name: str                # 动态库名称（如 "libctptrader_c.dylib"）
) -> str:
    """生成 cgo 包装代码，参考 ctpquote_api_darwin.go 风格"""
    
    prefix = f"ctp{api_name.lower()}"
    prefix_upper = f"Ctp{api_name}"
    prefix_lower = prefix
    
    lines = []
    lines.append("package ctpgo")
    lines.append("")
    lines.append("/*")
    lines.append(f"#cgo CFLAGS: -I${{SRCDIR}}/../csrc")
    lines.append(f"#cgo LDFLAGS: -L${{SRCDIR}}/../libs/ -l{prefix}_c")
    lines.append(f'#include "{header_file}"')
    lines.append("*/")
    lines.append("import \"C\"")
    lines.append("import (")
    lines.append('\t"os"')
    lines.append('\t"unsafe"')
    lines.append(")")
    lines.append("")
    
    # 生成结构体定义
    lines.append(f"type {struct_name} struct {{")
    lines.append(f"\tapi\t\tunsafe.Pointer")
    lines.append(f"\tpSpi\t\tunsafe.Pointer")
    lines.append(f"\tversion\t\tstring")
    lines.append(f"\tpszFlowPath\tstring")
    
    # Md 特有的字段
    if api_name == "Md":
        lines.append("\tusingUdp\tbool")
        lines.append("\tusingMulticast\tbool")
    
    lines.append("")
    lines.append("\t// 回调函数字段")
    
    # 生成回调函数字段
    for cb in callbacks:
        if cb.comment:
            lines.append(f"\t// {cb.comment}")
        
        # 生成回调函数签名
        cb_params = []
        for param in cb.params:
            if param.name == "userData":
                continue  # 跳过 userData 参数
            go_type = c_type_to_go_type(param.type, param.is_pointer, param.is_const)
            cb_params.append(f"{go_type}")
        
        cb_sig = ", ".join(cb_params) if cb_params else ""
        lines.append(f"\t{cb.go_method_name}_ func({cb_sig})")
    
    lines.append("}")
    lines.append("")
    
    # 全局变量（用于单实例，保持兼容性）
    var_name = struct_name[0].lower()
    lines.append(f"var {var_name} *{struct_name}")
    lines.append("")
    
    # 生成 Init 函数
    if api_name == "Md":
        lines.append(f"func {init_func_name}(pszFlowPath string, usingUdp bool, usingMulticast bool) *{struct_name} {{")
    else:
        lines.append(f"func {init_func_name}(pszFlowPath string) *{struct_name} {{")
    
    lines.append(f"\t{var_name} = new({struct_name})")
    lines.append(f"\t{var_name}.pszFlowPath = pszFlowPath")
    if api_name == "Md":
        lines.append(f"\t{var_name}.usingUdp = usingUdp")
        lines.append(f"\t{var_name}.usingMulticast = usingMulticast")
    lines.append("\t// 执行目录下创建 log目录")
    lines.append(f"\t_, err := os.Stat({var_name}.pszFlowPath)")
    lines.append("\tif err != nil {")
    lines.append(f"\t\tos.Mkdir({var_name}.pszFlowPath, os.ModePerm)")
    lines.append("\t}")
    lines.append(f"\t{var_name}.api = {var_name}.CreateApi()")
    lines.append(f"\t{var_name}.pSpi = {var_name}.CreateSpi()")
    lines.append(f"\t{var_name}.version = {var_name}.GetApiVersion()")
    lines.append("")
    lines.append(f"\treturn {var_name}")
    lines.append("}")
    lines.append("")
    
    # 生成 API 方法
    api_methods = {
        "CreateApi": None,
        "CreateSpi": None,
        "GetApiVersion": None,
        "GetTradingDay": None,
        "Release": None,
        "Init": None,
        "Join": None,
        "RegisterFront": None,
        "RegisterNameServer": None,
        "RegisterFensUserInfo": None,
        "RegisterSpi": None,
    }
    
    # 查找对应的 C 函数
    for func in functions:
        func_name_lower = func.name.lower()
        
        # CreateApi
        if f"{prefix_lower}create" in func_name_lower and "api" in func_name_lower:
            api_methods["CreateApi"] = func
        # CreateSpi
        elif f"{prefix_lower}spicreate" in func_name_lower:
            api_methods["CreateSpi"] = func
        # GetApiVersion
        elif f"{prefix_lower}getapiversion" in func_name_lower:
            api_methods["GetApiVersion"] = func
        # GetTradingDay
        elif f"{prefix_lower}gettradingday" in func_name_lower:
            api_methods["GetTradingDay"] = func
        # Release
        elif f"{prefix_lower}release" in func_name_lower:
            api_methods["Release"] = func
        # Init
        elif f"{prefix_lower}init" in func_name_lower and len(func.params) == 1:
            api_methods["Init"] = func
        # Join
        elif f"{prefix_lower}join" in func_name_lower:
            api_methods["Join"] = func
        # RegisterFront
        elif f"{prefix_lower}registerfront" in func_name_lower:
            api_methods["RegisterFront"] = func
        # RegisterNameServer
        elif f"{prefix_lower}registernameserver" in func_name_lower:
            api_methods["RegisterNameServer"] = func
        # RegisterFensUserInfo
        elif f"{prefix_lower}registerfensuserinfo" in func_name_lower:
            api_methods["RegisterFensUserInfo"] = func
        # RegisterSpi
        elif f"{prefix_lower}registerspi" in func_name_lower:
            api_methods["RegisterSpi"] = func
    
    # 生成 CreateApi 方法
    if api_methods["CreateApi"]:
        func = api_methods["CreateApi"]
        lines.append(f"func ({var_name} *{struct_name}) CreateApi() unsafe.Pointer {{")
        if api_name == "Md":
            lines.append(f"\tapi := C.{func.name}(C.CString({var_name}.pszFlowPath), C._Bool({var_name}.usingUdp), C._Bool({var_name}.usingMulticast))")
        else:
            lines.append(f"\tapi := C.{func.name}(C.CString({var_name}.pszFlowPath))")
        lines.append(f"\treturn unsafe.Pointer(api)")
        lines.append("}")
        lines.append("")
    
    # 生成 CreateSpi 方法
    if api_methods["CreateSpi"]:
        func = api_methods["CreateSpi"]
        lines.append(f"func ({var_name} *{struct_name}) CreateSpi() unsafe.Pointer {{")
        lines.append(f"\tpSpi := C.{func.name}(unsafe.Pointer({var_name}))")
        lines.append(f"\treturn unsafe.Pointer(pSpi)")
        lines.append("}")
        lines.append("")
    
    # 生成 GetApiVersion 方法
    if api_methods["GetApiVersion"]:
        func = api_methods["GetApiVersion"]
        lines.append(f"func ({var_name} *{struct_name}) GetApiVersion() string {{")
        lines.append(f"\treturn C.GoString((*C.char)(C.{func.name}()))")
        lines.append("}")
        lines.append("")
    
    # 生成 GetTradingDay 方法
    if api_methods["GetTradingDay"]:
        func = api_methods["GetTradingDay"]
        lines.append(f"func ({var_name} *{struct_name}) GetTradingDay() string {{")
        lines.append(f"\treturn C.GoString((*C.char)(C.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api))))")
        lines.append("}")
        lines.append("")
    
    # 生成 Release 方法
    if api_methods["Release"]:
        func = api_methods["Release"]
        lines.append(f"func ({var_name} *{struct_name}) Release() {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api))")
        lines.append("}")
        lines.append("")
    
    # 生成 Init 方法
    if api_methods["Init"]:
        func = api_methods["Init"]
        lines.append(f"func ({var_name} *{struct_name}) Init() {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api))")
        lines.append("}")
        lines.append("")
    
    # 生成 Join 方法
    if api_methods["Join"]:
        func = api_methods["Join"]
        lines.append(f"func ({var_name} *{struct_name}) Join() int32 {{")
        lines.append(f"\tres := C.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api))")
        lines.append("\treturn int32(res)")
        lines.append("}")
        lines.append("")
    
    # 生成 RegisterFront 方法
    if api_methods["RegisterFront"]:
        func = api_methods["RegisterFront"]
        lines.append(f"func ({var_name} *{struct_name}) RegisterFront(pszFrontAddress []byte) {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api), (*C.char)(unsafe.Pointer(C.CBytes(pszFrontAddress))))")
        lines.append("}")
        lines.append("")
    
    # 生成 RegisterNameServer 方法
    if api_methods["RegisterNameServer"]:
        func = api_methods["RegisterNameServer"]
        lines.append(f"func ({var_name} *{struct_name}) RegisterNameServer(pszNsAddress []byte) {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api), (*C.char)(unsafe.Pointer(C.CBytes(pszNsAddress))))")
        lines.append("}")
        lines.append("")
    
    # 生成 RegisterFensUserInfo 方法
    if api_methods["RegisterFensUserInfo"]:
        func = api_methods["RegisterFensUserInfo"]
        lines.append(f"func ({var_name} *{struct_name}) RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField) {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api), (*C.struct_CThostFtdcFensUserInfoField)(unsafe.Pointer(pFensUserInfo)))")
        lines.append("}")
        lines.append("")
    
    # 生成 RegisterSpi 方法
    if api_methods["RegisterSpi"]:
        func = api_methods["RegisterSpi"]
        lines.append(f"func ({var_name} *{struct_name}) RegisterSpi() {{")
        lines.append(f"\tC.{func.name}((*C.{prefix_upper}ApiHandle)({var_name}.api), (*C.{prefix_upper}SpiHandle)({var_name}.pSpi))")
        lines.append("}")
        lines.append("")
    
    # 生成其他请求方法（Req*）
    for func in functions:
        if func.name.startswith(f"{prefix_lower}Req") or func.name.startswith(f"{prefix_lower}Subscribe") or func.name.startswith(f"{prefix_lower}UnSubscribe"):
            generate_request_method(lines, func, var_name, struct_name, prefix_upper, api_name)
    
    # 生成回调设置方法（On*）
    for cb in callbacks:
        generate_callback_setter(lines, cb, var_name, struct_name, prefix_upper)
    
    # 生成 //export 回调函数
    for cb in callbacks:
        generate_export_callback(lines, cb, var_name, struct_name, prefix_lower)
    
    return "\n".join(lines)


def generate_request_method(lines: List[str], func: CFunction, var_name: str, struct_name: str, prefix_upper: str, api_name: str):
    """生成请求方法"""
    # 提取方法名：ctpTraderReqUserLogin -> ReqUserLogin
    method_name = func.name.replace(f"ctp{api_name.lower()}", "")
    if method_name:
        method_name = method_name[0].upper() + method_name[1:]
    
    if func.comment:
        lines.append(f"// {func.comment}")
    
    # 检查是否是字符串数组参数（如 SubscribeMarketData）
    is_string_array = False
    for param in func.params:
        if param.type == "char" and param.is_pointer and "pp" in param.name.lower():
            is_string_array = True
            break
    
    # 构建参数列表
    go_params = []
    call_params = []
    
    for param in func.params:
        if param.type.endswith("Handle") and not param.is_pointer:
            # API/SPI 句柄，使用实例的字段
            if "Api" in param.type:
                call_params.append(f"(*C.{prefix_upper}ApiHandle)({var_name}.api)")
            elif "Spi" in param.type:
                call_params.append(f"(*C.{prefix_upper}SpiHandle)({var_name}.pSpi)")
            continue
        
        go_type = c_type_to_go_type(param.type, param.is_pointer, param.is_const)
        go_name = go_param_name(param.name)
        
        # 字符串数组特殊处理
        if is_string_array and param.type == "char" and param.is_pointer and "pp" in param.name.lower():
            go_params.append(f"{go_name} [][]byte")
            go_params.append("nCount int")
            # 在函数体内处理
            continue
        elif param.name == "nCount" and is_string_array:
            # nCount 参数已经在上面处理了
            continue
        
        if go_type == "string":
            go_params.append(f"{go_name} string")
            call_params.append(f"C.CString({go_name})")
        elif go_type.startswith("*CThostFtdc"):
            go_params.append(f"{go_name} *{go_type.replace('*CThostFtdc', 'CThostFtdc')}")
            call_params.append(f"(*C.struct_{go_type.replace('*CThostFtdc', 'CThostFtdc')})(unsafe.Pointer({go_name}))")
        elif go_type == "int32":
            go_params.append(f"{go_name} {go_type}")
            call_params.append(f"C.int({go_name})")
        elif go_type == "bool":
            go_params.append(f"{go_name} {go_type}")
            call_params.append(f"C._Bool({go_name})")
        else:
            go_params.append(f"{go_name} {go_type}")
            call_params.append(f"{go_name}")
    
    go_param_str = ", ".join(go_params)
    call_param_str = ", ".join(call_params)
    
    # 返回类型
    return_type_clean = func.return_type.replace("const", "").strip()
    is_return_pointer = "*" in func.return_type
    go_return_type = c_type_to_go_type(return_type_clean, is_return_pointer)
    
    if "const" in func.return_type and "char" in func.return_type and "*" in func.return_type:
        go_return_type = "string"
    
    if go_return_type == "int32" or go_return_type == "":
        go_return_type = "int32"
    
    lines.append(f"func ({var_name} *{struct_name}) {method_name}({go_param_str}) {go_return_type} {{")
    
    # 字符串数组特殊处理
    if is_string_array:
        lines.append("")
        lines.append("\ttmp_arr := make([]*C.char, nCount)")
        lines.append("\tfor i := 0; i < nCount; i++ {")
        lines.append("\t\ttmp_arr[i] = C.CString(string(ppInstrumentID[i]))")
        lines.append("\t}")
        lines.append("\tvar _ppPtr **C.char")
        lines.append("\tif nCount > 0 {")
        lines.append("\t\t_ppPtr = (**C.char)(unsafe.Pointer(&tmp_arr[0]))")
        lines.append("\t}")
        # 找到 API handle 参数位置
        api_param = None
        for param in func.params:
            if param.type.endswith("ApiHandle"):
                api_param = param
                break
        if api_param:
            call_param_str = f"(*C.{prefix_upper}ApiHandle)({var_name}.api), _ppPtr, C.int(nCount)"
        lines.append(f"\tres := C.{func.name}({call_param_str})")
        lines.append("\tfor i := 0; i < nCount; i++ {")
        lines.append("\t\tif tmp_arr[i] != nil {")
        lines.append("\t\t\tC.free(unsafe.Pointer(tmp_arr[i]))")
        lines.append("\t\t}")
        lines.append("\t}")
        lines.append("\treturn int32(res)")
    elif go_return_type == "string":
        lines.append(f"\tret := C.{func.name}({call_param_str})")
        lines.append("\treturn C.GoString((*C.char)(ret))")
    elif go_return_type == "int32":
        lines.append(f"\tres := C.{func.name}({call_param_str})")
        lines.append("\treturn int32(res)")
    else:
        lines.append(f"\tC.{func.name}({call_param_str})")
    
    lines.append("}")
    lines.append("")


def generate_callback_setter(lines: List[str], cb: CallbackType, var_name: str, struct_name: str, prefix_upper: str):
    """生成回调设置方法"""
    if cb.comment:
        lines.append(f"// {cb.comment}")
    
    # 构建回调函数签名
    cb_params = []
    for param in cb.params:
        if param.name == "userData":
            continue
        go_type = c_type_to_go_type(param.type, param.is_pointer, param.is_const)
        cb_params.append(f"{go_type}")
    
    cb_sig = ", ".join(cb_params) if cb_params else ""
    
    # 查找对应的 C 设置函数
    prefix_lower = prefix_upper.replace("Ctp", "").lower()
    setter_func_name = f"ctp{prefix_lower}SpiSet{cb.go_method_name}"
    
    # 生成 //export 函数名（用于 C 调用）
    export_func_name = f"{prefix_lower[0]}{cb.go_method_name}_"
    
    lines.append(f"func ({var_name} *{struct_name}) {cb.go_method_name}(fn func({cb_sig})) {{")
    lines.append(f"\t{var_name}.{cb.go_method_name}_ = fn")
    # 将 Go 的 //export 函数转换为 C 回调函数指针
    # 在 cgo 中，//export 函数会自动导出，我们可以直接引用并转换为 C 回调函数指针类型
    # 注意：回调函数签名必须匹配 C 回调类型（第一个参数是 void* userData）
    callback_type_name = cb.name.replace("Callback", "")
    lines.append(f"\tC.{setter_func_name}((*C.{prefix_upper}SpiHandle)({var_name}.pSpi), C.{callback_type_name}(C.{export_func_name}))")
    lines.append("}")
    lines.append("")


def generate_export_callback(lines: List[str], cb: CallbackType, var_name: str, struct_name: str, prefix_lower: str):
    """生成 //export 回调函数"""
    export_name = f"{prefix_lower[0]}{cb.go_method_name}_"
    
    # 构建 C 函数签名（匹配回调类型，第一个参数是 userData）
    go_params = []
    for param in cb.params:
        go_type = c_type_to_go_type(param.type, param.is_pointer, param.is_const)
        go_name = go_param_name(param.name)
        go_params.append(f"{go_name} {go_type}")
    
    go_sig = ", ".join(go_params)
    
    lines.append(f"//export {export_name}")
    lines.append(f"func {export_name}({go_sig}) {{")
    
    # 从 userData 恢复实例
    lines.append(f"\t{var_name} := (*{struct_name})(userData)")
    lines.append(f"\tif {var_name}.{cb.go_method_name}_ != nil {{")
    
    # 构建调用参数（跳过 userData）
    call_args = []
    for param in cb.params:
        if param.name == "userData":
            continue
        go_name = go_param_name(param.name)
        if param.type.startswith("struct "):
            struct_name_clean = param.type.replace("struct ", "").strip()
            call_args.append(f"(*CThostFtdc{struct_name_clean})(unsafe.Pointer({go_name}))")
        elif param.type == "int":
            call_args.append(f"int({go_name})")
        elif param.type == "bool":
            call_args.append(f"bool({go_name})")
        else:
            call_args.append(go_name)
    
    call_args_str = ", ".join(call_args) if call_args else ""
    lines.append(f"\t\t{var_name}.{cb.go_method_name}_({call_args_str})")
    lines.append("\t}")
    lines.append("}")
    lines.append("")


def main():
    parser = argparse.ArgumentParser(description='CTP C API 转 Go CGO 包装代码生成器')
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
        
        go_code = generate_cgo_wrapper(
            "Md", "Quote", "InitQuote",
            functions, callbacks, structs,
            "ctp_md_c_api.h", "libctpmd_c"
        )
        output_file = output_dir / "ctpquote_api_darwin.go"
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
        
        go_code = generate_cgo_wrapper(
            "Trader", "Trade", "InitTrade",
            functions, callbacks, structs,
            "ctp_trader_c_api.h", "libctptrader_c"
        )
        output_file = output_dir / "ctptrade_api_darwin.go"
        output_file.write_text(go_code, encoding='utf-8')
        print(f"  生成 {output_file}")
    else:
        print(f"警告: 未找到 {trader_header}")
    
    print()
    print("完成!")
    return 0


if __name__ == '__main__':
    exit(main())
