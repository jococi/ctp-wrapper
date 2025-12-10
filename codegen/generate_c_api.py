#!/usr/bin/env python3
"""
CTP C++ API 转 C API 代码生成器

功能：
- 解析 CTP 官方 C++ 头文件（ThostFtdcMdApi.h, ThostFtdcTraderApi.h）
- 生成纯 C 头文件（带 user_data 的回调）
- 生成 C++ 实现文件
- 使用驼峰命名风格
- 跨平台统一接口（Linux/Windows/macOS）

跨平台差异处理：
- MdApi: 三平台完全一致
- TraderApi: macOS 的 ReqUserLogin 多两个参数 (length, systemInfo)
  解决方案: C API 统一提供带 systemInfo 的版本，Linux/Windows 内部忽略

用法：
    python generate.py --input ../ctpapi/linux --output ./output
"""

import re
import os
import argparse
from pathlib import Path
from dataclasses import dataclass
from typing import List, Optional


@dataclass
class Parameter:
    """函数参数"""
    type: str       # 参数类型
    name: str       # 参数名
    is_pointer: bool = False


@dataclass
class Method:
    """方法/回调定义"""
    name: str                   # 方法名
    return_type: str            # 返回类型
    params: List[Parameter]     # 参数列表
    comment: str = ""           # 注释
    is_static: bool = False     # 是否静态方法
    is_pure_virtual: bool = False  # 是否纯虚函数


def to_camel_case(name: str, lower_first: bool = True) -> str:
    """
    转换为驼峰命名
    OnFrontConnected -> onFrontConnected (lower_first=True)
    OnFrontConnected -> OnFrontConnected (lower_first=False)
    """
    if not name:
        return name
    if lower_first:
        return name[0].lower() + name[1:]
    return name


def parse_params(param_str: str) -> List[Parameter]:
    """解析参数列表字符串"""
    params = []
    if not param_str or param_str.strip() == "":
        return params
    
    # 分割参数
    param_list = []
    depth = 0
    current = ""
    for char in param_str:
        if char == '<':
            depth += 1
        elif char == '>':
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
        
        # 处理默认值
        if '=' in param:
            param = param.split('=')[0].strip()
        
        # 解析类型和名称
        # 例如: "const char *pszFlowPath" -> type="const char *", name="pszFlowPath"
        # 例如: "CThostFtdcRspInfoField *pRspInfo" -> type="CThostFtdcRspInfoField *", name="pRspInfo"
        
        is_pointer = '*' in param
        
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
            
            params.append(Parameter(type=param_type, name=name, is_pointer=is_pointer))
        elif len(parts) == 1:
            # 只有类型没有名字（如 void）
            pass
    
    return params


def parse_spi_class(content: str, class_name: str) -> List[Method]:
    """解析 Spi 类的虚函数"""
    methods = []
    
    # 找到类定义开始
    class_start = content.find(f'class {class_name}')
    if class_start == -1:
        print(f"警告: 未找到类 {class_name}")
        return methods
    
    # 找到类体开始的 {
    brace_start = content.find('{', class_start)
    if brace_start == -1:
        print(f"警告: 未找到类 {class_name} 的开始大括号")
        return methods
    
    # 找下一个 class 关键字或文件末尾作为类的结束边界
    next_class = content.find('\nclass ', brace_start)
    if next_class == -1:
        next_class = len(content)
    
    # 在这个范围内找 }; 作为类的结束
    class_end = content.rfind('\n};', brace_start, next_class)
    if class_end == -1:
        class_end = next_class
    
    class_body = content[brace_start:class_end]
    
    # 解析虚函数
    # CTP 的 Spi 虚函数格式:
    # ///注释内容
    # virtual void OnXxx(params){};
    # 或者
    # virtual void OnXxx(params) {};
    
    # 找所有 virtual void OnXxx 方法
    method_pattern = r'virtual\s+void\s+(On\w+)\s*\(([^)]*)\)\s*\{[^}]*\}'
    
    for m in re.finditer(method_pattern, class_body):
        name = m.group(1)
        param_str = m.group(2)
        
        # 尝试找前面的注释
        method_start = m.start()
        preceding = class_body[:method_start]
        comment = ""
        
        # 找最近的 /// 注释行（不带@）
        lines = preceding.split('\n')
        for line in reversed(lines):
            line = line.strip()
            if line.startswith('///') and not line.startswith('///@'):
                comment = line[3:].strip()
                break
            elif line and not line.startswith('//'):
                break
        
        params = parse_params(param_str)
        methods.append(Method(
            name=name,
            return_type="void",
            params=params,
            comment=comment,
            is_pure_virtual=True
        ))
    
    return methods


def parse_api_class(content: str, class_name: str) -> List[Method]:
    """解析 Api 类的方法"""
    methods = []
    
    # 找到类定义开始
    # 格式: class MD_API_EXPORT CThostFtdcMdApi 或 class CThostFtdcTraderApi
    class_start = content.find(class_name)
    if class_start == -1:
        print(f"警告: 未找到类 {class_name}")
        return methods
    
    # 找到类体开始的 {
    brace_start = content.find('{', class_start)
    if brace_start == -1:
        return methods
    
    # 找下一个 class 关键字或 protected/private 结尾
    # Api 类通常以 protected: ~CThostFtdcXxxApi(){}; 结尾
    class_end = content.find('protected:', brace_start)
    if class_end == -1:
        class_end = content.find('\n};', brace_start)
    if class_end == -1:
        class_end = len(content)
    
    class_body = content[brace_start:class_end]
    
    # 解析静态方法
    # static CThostFtdcMdApi *CreateFtdcMdApi(const char *pszFlowPath = "", ...);
    # static const char *GetApiVersion();
    static_pattern = r'static\s+([\w\s]+\s*\*?)\s*(\w+)\s*\(([^)]*)\)\s*;'
    for m in re.finditer(static_pattern, class_body):
        return_type = m.group(1).strip()
        name = m.group(2)
        param_str = m.group(3)
        
        # 找前面的注释
        method_start = m.start()
        preceding = class_body[:method_start]
        comment = extract_comment(preceding)
        
        params = parse_params(param_str)
        methods.append(Method(
            name=name,
            return_type=return_type,
            params=params,
            comment=comment,
            is_static=True
        ))
    
    # 解析虚方法（成员函数）
    # virtual void Release() = 0;
    # virtual int ReqUserLogin(...) = 0;
    # virtual const char *GetTradingDay() = 0;
    # 返回类型可能是: void, int, const char *, const char*
    virtual_pattern = r'virtual\s+((?:const\s+)?(?:char|int|void)(?:\s*\*)?)\s*(\w+)\s*\(([^)]*)\)\s*=\s*0\s*;'
    for m in re.finditer(virtual_pattern, class_body):
        return_type = m.group(1).strip()
        name = m.group(2)
        param_str = m.group(3)
        
        # 找前面的注释
        method_start = m.start()
        preceding = class_body[:method_start]
        comment = extract_comment(preceding)
        
        params = parse_params(param_str)
        methods.append(Method(
            name=name,
            return_type=return_type,
            params=params,
            comment=comment,
            is_pure_virtual=True
        ))
    
    return methods


def extract_comment(preceding: str) -> str:
    """从前置文本中提取最近的注释"""
    lines = preceding.split('\n')
    comment_lines = []
    
    for line in reversed(lines):
        line = line.strip()
        if line.startswith('///'):
            # 去掉 /// 和可能的 @xxx
            comment_text = line[3:].strip()
            if comment_text.startswith('@'):
                # 跳过 @param, @return 等
                continue
            comment_lines.insert(0, comment_text)
        elif line and not line.startswith('//'):
            break
    
    return ' '.join(comment_lines) if comment_lines else ""


def generate_c_header(
    api_name: str,          # "Md" 或 "Trade"
    spi_methods: List[Method],
    api_methods: List[Method],
    cpp_spi_class: str,     # "CThostFtdcMdSpi"
    cpp_api_class: str      # "CThostFtdcMdApi"
) -> str:
    """生成纯 C 头文件"""
    
    prefix = f"{api_name}"
    prefix_lower = f"{api_name}"
    guard = f"CTP_{api_name.upper()}_C_API_H"
    
    lines = []
    lines.append(f"/**")
    lines.append(f" * CTP {api_name} API - 纯 C 接口封装")
    lines.append(f" * ")
    lines.append(f" * 自动生成，请勿手动修改")
    lines.append(f" * 特性：")
    lines.append(f" *   - 纯 C 接口，无 C++ 依赖")
    lines.append(f" *   - 不透明指针句柄")
    lines.append(f" *   - 回调携带 userData，支持多实例")
    lines.append(f" *   - 驼峰命名风格")
    lines.append(f" */")
    lines.append(f"")
    lines.append(f"#ifndef {guard}")
    lines.append(f"#define {guard}")
    lines.append(f"")
    lines.append(f"#ifdef __cplusplus")
    lines.append(f'extern "C" {{')
    lines.append(f"#endif")
    lines.append(f"")
    lines.append(f"#include <stdbool.h>")
    lines.append(f"#include <stdint.h>")
    lines.append(f"")
    
    # 平台相关的类型定义 (需要 THOST_TE_RESUME_TYPE)
    # 包含结构体定义头文件（CTP的结构体本身是C兼容的POD）
    lines.append(f"// ========== 结构体定义 ==========")
    lines.append(f"// CTP 的 Field 结构体是 C 兼容的 POD 类型，直接使用")
    lines.append(f"// 平台相关的类型定义 (需要 THOST_TE_RESUME_TYPE)")
    lines.append(f"#ifdef _WIN32")
    lines.append(f"    #include \"ctpapi/windows/ThostFtdcUserApiStruct.h\"")
    lines.append(f"#elif __APPLE__")
    lines.append(f"    #include \"ctpapi/macos/ThostFtdcUserApiStruct.h\"")
    lines.append(f"#elif __linux__")
    lines.append(f"    #include \"ctpapi/linux/ThostFtdcUserApiStruct.h\"")
    lines.append(f"#endif")
    lines.append(f"")
    
    # 平台相关的导出宏
    lines.append(f"// 导出宏定义")
    lines.append(f"#ifdef _WIN32")
    lines.append(f"    #ifdef CTP_EXPORTS")
    lines.append(f"        #define CTP_API __declspec(dllexport)")
    lines.append(f"    #else")
    lines.append(f"        #define CTP_API __declspec(dllimport)")
    lines.append(f"    #endif")
    lines.append(f"#else")
    lines.append(f"    #define CTP_API")
    lines.append(f"#endif")
    lines.append(f"")
    
    # 不透明句柄类型
    lines.append(f"// ========== 不透明句柄类型 ==========")
    lines.append(f"typedef struct {prefix}Api_t* {prefix}ApiHandle;")
    lines.append(f"typedef struct {prefix}Spi_t* {prefix}SpiHandle;")
    lines.append(f"")
    
    # 回调函数类型定义（带 userData）
    lines.append(f"// ========== 回调函数类型（带 userData） ==========")
    for method in spi_methods:
        callback_name = f"{prefix}On{method.name[2:]}"  # 去掉 "On" 前缀再加回来
        
        # 构建参数列表
        param_strs = ["void* userData"]
        for param in method.params:
            param_type = param.type
            if "CThostFtdc" in param_type:
                # 现在有了完整定义，直接使用类型名（不需要 struct 关键字）
                param_strs.append(f"{param_type} {param.name}")
            else:
                param_strs.append(f"{param_type} {param.name}")
        
        param_list = ", ".join(param_strs)
        
        if method.comment:
            lines.append(f"// {method.comment}")
        lines.append(f"typedef void (*{callback_name}Callback)({param_list});")
    lines.append(f"")
    
    # 回调表结构
    lines.append(f"// ========== 回调表结构（便于批量设置） ==========")
    lines.append(f"typedef struct {{")
    lines.append(f"    void* userData;")
    for method in spi_methods:
        callback_name = f"{prefix}On{method.name[2:]}"
        field_name = to_camel_case(f"on{method.name[2:]}")
        lines.append(f"    {callback_name}Callback {field_name};")
    lines.append(f"}} {prefix}SpiCallbacks;")
    lines.append(f"")
    
    # API 函数声明
    lines.append(f"// ========== {api_name} API 函数 ==========")
    lines.append(f"")
    
    # 静态方法（创建、版本等）
    for method in api_methods:
        if method.is_static:
            func_name = f"{prefix_lower}{method.name}"
            
            # 构建参数列表
            param_strs = []
            for param in method.params:
                param_type = param.type
                # 现在有了完整定义，直接使用类型名（不需要 struct 关键字）
                param_strs.append(f"{param_type} {param.name}")
            
            param_list = ", ".join(param_strs) if param_strs else "void"
            
            # 返回类型转换
            ret_type = method.return_type
            if "CThostFtdc" in ret_type and "Api" in ret_type:
                ret_type = f"{prefix}ApiHandle"
            elif ret_type == "const char":
                ret_type = "const char*"
            
            if method.comment:
                # 清理注释，只保留简短描述
                clean_comment = method.comment.split('\n')[0].strip()
                if clean_comment:
                    lines.append(f"// {clean_comment}")
            lines.append(f"CTP_API {ret_type} {func_name}({param_list});")
            lines.append(f"")
    lines.append(f"")
    
    # 实例方法
    for method in api_methods:
        if not method.is_static and method.name not in ["RegisterSpi"]:
            func_name = f"{prefix_lower}{method.name}"
            
            # 构建参数列表，第一个参数是 handle
            param_strs = [f"{prefix}ApiHandle handle"]
            
            # 特殊处理: TraderApi 的 ReqUserLogin 在 macOS 上签名不同
            # 统一接口：只保留标准参数，忽略 macOS 特有的 length 和 systemInfo 参数
            if api_name == "Trader" and method.name == "ReqUserLogin":
                # 只保留前两个参数（pReqUserLoginField 和 nRequestID）
                for param in method.params[:2]:
                    param_type = param.type
                    # 现在有了完整定义，直接使用类型名（不需要 struct 关键字）
                    param_strs.append(f"{param_type} {param.name}")
            else:
                # 其他方法正常处理所有参数
                for param in method.params:
                    param_type = param.type
                    # 现在有了完整定义，直接使用类型名（不需要 struct 关键字）
                    param_strs.append(f"{param_type} {param.name}")
            
            param_list = ", ".join(param_strs)
            
            # 返回类型
            ret_type = method.return_type
            if ret_type == "const char":
                ret_type = "const char*"
            
            if method.comment:
                # 清理注释，只保留简短描述
                clean_comment = method.comment.split('\n')[0].strip()
                if clean_comment:
                    lines.append(f"// {clean_comment}")
            lines.append(f"CTP_API {ret_type} {func_name}({param_list});")
            lines.append(f"")
    lines.append(f"")
    
    # SPI 函数声明
    lines.append(f"// ========== {api_name} SPI 函数 ==========")
    lines.append(f"")
    lines.append(f"// 创建 SPI 实例")
    lines.append(f"CTP_API {prefix}SpiHandle {prefix_lower}SpiCreate(void* userData);")
    lines.append(f"")
    lines.append(f"// 销毁 SPI 实例")
    lines.append(f"CTP_API void {prefix_lower}SpiDestroy({prefix}SpiHandle spi);")
    lines.append(f"")
    lines.append(f"// 注册 SPI 到 API")
    lines.append(f"CTP_API void {prefix_lower}RegisterSpi({prefix}ApiHandle api, {prefix}SpiHandle spi);")
    lines.append(f"")
    lines.append(f"// 批量设置回调")
    lines.append(f"CTP_API void {prefix_lower}SpiSetCallbacks({prefix}SpiHandle spi, const {prefix}SpiCallbacks* callbacks);")
    lines.append(f"")
    
    # 单独设置回调的函数
    lines.append(f"// 单独设置回调")
    for method in spi_methods:
        callback_name = f"{prefix}On{method.name[2:]}"
        func_name = f"{prefix_lower}SpiSetOn{method.name[2:]}"
        if method.comment:
            lines.append(f"// {method.comment}")
        lines.append(f"CTP_API void {func_name}({prefix}SpiHandle spi, {callback_name}Callback callback);")
    lines.append(f"")
    
    # 为 Trader API 添加跨平台登录函数和 DataCollect 函数
    if api_name == "Trader":
        lines.append(f"// ========== 跨平台统一登录接口 ==========")
        lines.append(f"// 说明: macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数")
        lines.append(f"// 此函数在 Linux/Windows 上忽略 systemInfo，在 macOS 上使用它")
        lines.append(f"")
        lines.append(f"// 带系统信息的用户登录请求（跨平台统一接口）")
        lines.append(f"// systemInfoLen: 系统信息长度，传 0 表示自动采集（仅 macOS 生效）")
        lines.append(f"// systemInfo: 系统信息数据，传 NULL 表示自动采集（仅 macOS 生效）")
        lines.append(f"CTP_API int TraderReqUserLoginWithSystemInfo(TraderApiHandle handle,")
        lines.append(f"    CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID,")
        lines.append(f"    int systemInfoLen, const char* systemInfo);")
        lines.append(f"")
        lines.append(f"// ========== DataCollect 函数 ==========")
        lines.append(f"")
        lines.append(f"// 获取终端信息（AES+RSA 加密）")
        lines.append(f"// pSystemInfo: 输出缓冲区，至少 270 字节")
        lines.append(f"// pLen: 输入缓冲区大小，输出实际数据长度")
        lines.append(f"// 返回值: 0 成功，非 0 表示采集错误（按位判断）")
        lines.append(f"CTP_API int DCGetSystemInfo(char* pSystemInfo, int* pLen);")
        lines.append(f"")
        lines.append(f"// 获取终端信息（未 AES 加密）")
        lines.append(f"CTP_API int DCGetSystemInfoUnAesEncode(char* pSystemInfo, int* pLen);")
        lines.append(f"")
        lines.append(f"// 获取 DataCollect API 版本")
        lines.append(f"CTP_API const char* DCGetDataCollectApiVersion(void);")
        lines.append(f"")
    
    lines.append(f"#ifdef __cplusplus")
    lines.append(f"}}")
    lines.append(f"#endif")
    lines.append(f"")
    lines.append(f"#endif // {guard}")
    lines.append(f"")
    
    return "\n".join(lines)


def generate_cpp_impl(
    api_name: str,          # "Md" 或 "Trade"
    spi_methods: List[Method],
    api_methods: List[Method],
    cpp_spi_class: str,     # "CThostFtdcMdSpi"
    cpp_api_class: str      # "CThostFtdcMdApi"
) -> str:
    """生成 C++ 实现文件"""
    
    prefix = f"{api_name}"
    prefix_lower = f"{api_name}"
    header_name = f"ctp_{api_name.lower()}_c_api.h"
    
    lines = []
    lines.append(f"/**")
    lines.append(f" * CTP {api_name} API - C 接口实现")
    lines.append(f" * ")
    lines.append(f" * 自动生成，请勿手动修改")
    lines.append(f" */")
    lines.append(f"")
    lines.append(f"#include \"{header_name}\"")
    lines.append(f"")
    lines.append(f"// 平台相关的原始 CTP 头文件")
    lines.append(f"// 注意: 编译时需要正确设置头文件搜索路径")
    lines.append(f"#ifdef _WIN32")
    lines.append(f"    #include \"ctpapi/windows/ThostFtdc{api_name}Api.h\"")
    lines.append(f"    #include \"ctpapi/windows/DataCollect.h\"")
    lines.append(f"#elif __APPLE__")
    lines.append(f"    #include \"ctpapi/macos/ThostFtdc{api_name}Api.h\"")
    lines.append(f"    #include \"ctpapi/macos/DataCollect.h\"")
    lines.append(f"#elif __linux__")
    lines.append(f"    #include \"ctpapi/linux/ThostFtdc{api_name}Api.h\"")
    lines.append(f"    #include \"ctpapi/linux/DataCollect.h\"")
    lines.append(f"#endif")
    lines.append(f"")
    lines.append(f"#include <cstring>")
    lines.append(f"")
    
    # SPI 包装类
    lines.append(f"// ========== SPI 包装类 ==========")
    lines.append(f"class {prefix}SpiWrapper : public {cpp_spi_class} {{")
    lines.append(f"public:")
    lines.append(f"    virtual ~{prefix}SpiWrapper() = default;")
    lines.append(f"")
    lines.append(f"    void* userData = nullptr;")
    lines.append(f"")
    
    # 回调函数指针成员
    for method in spi_methods:
        callback_name = f"{prefix}On{method.name[2:]}"
        field_name = to_camel_case(f"on{method.name[2:]}")
        lines.append(f"    {callback_name}Callback {field_name} = nullptr;")
    lines.append(f"")
    
    # 重写虚函数
    for method in spi_methods:
        field_name = to_camel_case(f"on{method.name[2:]}")
        
        # 参数列表
        param_decls = []
        param_calls = ["userData"]
        for param in method.params:
            param_decls.append(f"{param.type} {param.name}")
            param_calls.append(param.name)
        
        param_decl_str = ", ".join(param_decls) if param_decls else ""
        param_call_str = ", ".join(param_calls)
        
        lines.append(f"    void {method.name}({param_decl_str}) override {{")
        lines.append(f"        if ({field_name}) {{")
        lines.append(f"            {field_name}({param_call_str});")
        lines.append(f"        }}")
        lines.append(f"    }}")
        lines.append(f"")
    
    lines.append(f"}};")
    lines.append(f"")
    
    # C 接口实现
    lines.append(f"// ========== C 接口实现 ==========")
    lines.append(f"")
    lines.append(f'extern "C" {{')
    lines.append(f"")
    
    # 静态方法实现
    for method in api_methods:
        if method.is_static:
            func_name = f"{prefix_lower}{method.name}"
            
            # 参数列表
            param_decls = []
            param_calls = []
            for param in method.params:
                param_type = param.type
                if "CThostFtdc" in param_type:
                    param_decls.append(f"struct {param_type} {param.name}")
                else:
                    param_decls.append(f"{param_type} {param.name}")
                param_calls.append(param.name)
            
            param_decl_str = ", ".join(param_decls) if param_decls else "void"
            param_call_str = ", ".join(param_calls)
            
            # 返回类型
            ret_type = method.return_type
            if "CThostFtdc" in ret_type and "Api" in ret_type:
                ret_type = f"{prefix}ApiHandle"
                lines.append(f"{ret_type} {func_name}({param_decl_str}) {{")
                lines.append(f"    return reinterpret_cast<{prefix}ApiHandle>(")
                lines.append(f"        {cpp_api_class}::{method.name}({param_call_str})")
                lines.append(f"    );")
                lines.append(f"}}")
            elif ret_type == "const char":
                ret_type = "const char*"
                lines.append(f"{ret_type} {func_name}({param_decl_str}) {{")
                lines.append(f"    return {cpp_api_class}::{method.name}({param_call_str});")
                lines.append(f"}}")
            else:
                lines.append(f"{ret_type} {func_name}({param_decl_str}) {{")
                lines.append(f"    return {cpp_api_class}::{method.name}({param_call_str});")
                lines.append(f"}}")
            lines.append(f"")
    
    # 为 Trader API 先生成跨平台登录函数（因为其他函数会依赖它）
    if api_name == "Trader":
        lines.append(f"// 跨平台统一登录接口实现（参考 ctpgo 实现）")
        lines.append(f"// macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数")
        lines.append(f"// 此函数内部自动采集系统信息，调用方无需关心平台差异")
        lines.append(f"int {prefix_lower}ReqUserLoginWithSystemInfo({prefix}ApiHandle handle,")
        lines.append(f"    CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID,")
        lines.append(f"    int systemInfoLen, const char* systemInfo) {{")
        lines.append(f"    auto* api = reinterpret_cast<{cpp_api_class}*>(handle);")
        lines.append(f"#ifdef __APPLE__")
        lines.append(f"    // macOS 版本：内部自动采集系统信息（参考 ctpgo 实现）")
        lines.append(f"    if (systemInfo == nullptr || systemInfoLen == 0) {{")
        lines.append(f"        // 使用 CTP 自带的类型和未 AES 加密的系统信息采集函数")
        lines.append(f"        TThostFtdcClientSystemInfoType sysInfo = {{0}};")
        lines.append(f"        int len = sizeof(sysInfo);  // CTP_GetSystemInfoUnAesEncode 需要 int&")
        lines.append(f"        CTP_GetSystemInfoUnAesEncode(sysInfo, len);")
        lines.append(f"        // ReqUserLogin 需要 TThostFtdcSystemInfoLenType，进行类型转换")
        lines.append(f"        return api->ReqUserLogin(pReqUserLoginField, nRequestID, static_cast<TThostFtdcSystemInfoLenType>(len), sysInfo);")
        lines.append(f"    }}")
        lines.append(f"    return api->ReqUserLogin(pReqUserLoginField, nRequestID, systemInfoLen, ")
        lines.append(f"        const_cast<char*>(systemInfo));")
        lines.append(f"#else")
        lines.append(f"    // Linux/Windows 版本忽略 systemInfo 参数")
        lines.append(f"    (void)systemInfoLen;")
        lines.append(f"    (void)systemInfo;")
        lines.append(f"    return api->ReqUserLogin(pReqUserLoginField, nRequestID);")
        lines.append(f"#endif")
        lines.append(f"}}")
        lines.append(f"")
    
    # 实例方法实现
    for method in api_methods:
        if not method.is_static and method.name not in ["RegisterSpi"]:
            func_name = f"{prefix_lower}{method.name}"
            
            # 返回类型（先定义，后面会用到）
            ret_type = method.return_type
            if ret_type == "const char":
                ret_type = "const char*"
            
            # 参数列表
            param_decls = [f"{prefix}ApiHandle handle"]
            param_calls = []
            
            # 特殊处理: TraderApi 的 ReqUserLogin 在 macOS 上签名不同
            # 统一接口：只保留标准参数，忽略 macOS 特有的 length 和 systemInfo 参数
            if api_name == "Trader" and method.name == "ReqUserLogin":
                # 只保留前两个参数（pReqUserLoginField 和 nRequestID）
                for param in method.params[:2]:
                    param_type = param.type
                    if "CThostFtdc" in param_type:
                        param_decls.append(f"struct {param_type} {param.name}")
                    else:
                        param_decls.append(f"{param_type} {param.name}")
                    param_calls.append(param.name)
                
                param_decl_str = ", ".join(param_decls)
                lines.append(f"{ret_type} {func_name}({param_decl_str}) {{")
                lines.append(f"    // 跨平台兼容: macOS 的 ReqUserLogin 需要额外参数")
                lines.append(f"    return {prefix_lower}ReqUserLoginWithSystemInfo(handle, pReqUserLoginField, nRequestID, 0, nullptr);")
                lines.append(f"}}")
                lines.append(f"")
                continue
            else:
                # 其他方法正常处理所有参数
                for param in method.params:
                    param_type = param.type
                    if "CThostFtdc" in param_type:
                        param_decls.append(f"struct {param_type} {param.name}")
                    else:
                        param_decls.append(f"{param_type} {param.name}")
                    param_calls.append(param.name)
            
            param_decl_str = ", ".join(param_decls)
            param_call_str = ", ".join(param_calls)
            
            lines.append(f"{ret_type} {func_name}({param_decl_str}) {{")
            lines.append(f"    auto* api = reinterpret_cast<{cpp_api_class}*>(handle);")
            
            # 修正参数调用：去掉数组括号
            param_call_str_fixed = param_call_str.replace('[]', '')
            
            if ret_type == "void":
                if param_call_str_fixed:
                    lines.append(f"    api->{method.name}({param_call_str_fixed});")
                else:
                    lines.append(f"    api->{method.name}();")
            else:
                if param_call_str_fixed:
                    lines.append(f"    return api->{method.name}({param_call_str_fixed});")
                else:
                    lines.append(f"    return api->{method.name}();")
            lines.append(f"}}")
            lines.append(f"")
    
    # SPI 函数实现
    lines.append(f"// SPI 创建与销毁")
    lines.append(f"{prefix}SpiHandle {prefix_lower}SpiCreate(void* userData) {{")
    lines.append(f"    auto* spi = new {prefix}SpiWrapper();")
    lines.append(f"    spi->userData = userData;")
    lines.append(f"    return reinterpret_cast<{prefix}SpiHandle>(spi);")
    lines.append(f"}}")
    lines.append(f"")
    
    lines.append(f"void {prefix_lower}SpiDestroy({prefix}SpiHandle spi) {{")
    lines.append(f"    delete reinterpret_cast<{prefix}SpiWrapper*>(spi);")
    lines.append(f"}}")
    lines.append(f"")
    
    lines.append(f"void {prefix_lower}RegisterSpi({prefix}ApiHandle api, {prefix}SpiHandle spi) {{")
    lines.append(f"    auto* apiPtr = reinterpret_cast<{cpp_api_class}*>(api);")
    lines.append(f"    auto* spiPtr = reinterpret_cast<{prefix}SpiWrapper*>(spi);")
    lines.append(f"    apiPtr->RegisterSpi(spiPtr);")
    lines.append(f"}}")
    lines.append(f"")
    
    # 批量设置回调
    lines.append(f"void {prefix_lower}SpiSetCallbacks({prefix}SpiHandle spi, const {prefix}SpiCallbacks* callbacks) {{")
    lines.append(f"    auto* spiPtr = reinterpret_cast<{prefix}SpiWrapper*>(spi);")
    lines.append(f"    spiPtr->userData = callbacks->userData;")
    for method in spi_methods:
        field_name = to_camel_case(f"on{method.name[2:]}")
        lines.append(f"    spiPtr->{field_name} = callbacks->{field_name};")
    lines.append(f"}}")
    lines.append(f"")
    
    # 单独设置回调
    for method in spi_methods:
        callback_name = f"{prefix}On{method.name[2:]}"
        func_name = f"{prefix_lower}SpiSetOn{method.name[2:]}"
        field_name = to_camel_case(f"on{method.name[2:]}")
        
        lines.append(f"void {func_name}({prefix}SpiHandle spi, {callback_name}Callback callback) {{")
        lines.append(f"    reinterpret_cast<{prefix}SpiWrapper*>(spi)->{field_name} = callback;")
        lines.append(f"}}")
        lines.append(f"")
    
    # 为 Trader API 添加 DataCollect 函数实现
    if api_name == "Trader":
        lines.append(f"// DataCollect 函数实现")
        lines.append(f"int DCGetSystemInfo(char* pSystemInfo, int* pLen) {{")
        lines.append(f"    return CTP_GetSystemInfo(pSystemInfo, *pLen);")
        lines.append(f"}}")
        lines.append(f"")
        lines.append(f"int DCGetSystemInfoUnAesEncode(char* pSystemInfo, int* pLen) {{")
        lines.append(f"#ifdef __APPLE__")
        lines.append(f"    // macOS 版本有专门的未 AES 加密函数")
        lines.append(f"    return CTP_GetSystemInfoUnAesEncode(pSystemInfo, *pLen);")
        lines.append(f"#else")
        lines.append(f"    // Linux/Windows 版本没有此函数，回退到 CTP_GetSystemInfo")
        lines.append(f"    // 注意：这会返回 AES 加密的数据，但对于内部登录流程不影响")
        lines.append(f"    // 因为 Linux/Windows 的 ReqUserLogin 不需要 systemInfo 参数")
        lines.append(f"    return CTP_GetSystemInfo(pSystemInfo, *pLen);")
        lines.append(f"#endif")
        lines.append(f"}}")
        lines.append(f"")
        lines.append(f"const char* DCGetDataCollectApiVersion(void) {{")
        lines.append(f"    return CTP_GetDataCollectApiVersion();")
        lines.append(f"}}")
        lines.append(f"")
    
    lines.append(f"}} // extern \"C\"")
    lines.append(f"")
    
    return "\n".join(lines)


def process_api(input_dir: Path, output_dir: Path, api_name: str, 
                spi_class: str, api_class: str):
    """处理单个 API（Md 或 Trade）"""
    
    header_file = input_dir / f"ThostFtdc{api_name}Api.h"
    if not header_file.exists():
        print(f"错误: 找不到文件 {header_file}")
        return
    
    print(f"解析 {header_file}...")
    # CTP 头文件可能是 GBK 编码
    try:
        content = header_file.read_text(encoding='utf-8')
    except UnicodeDecodeError:
        content = header_file.read_text(encoding='gbk')
    
    # 解析 Spi 类
    spi_methods = parse_spi_class(content, spi_class)
    print(f"  找到 {len(spi_methods)} 个 Spi 回调方法")
    
    # 解析 Api 类
    api_methods = parse_api_class(content, api_class)
    print(f"  找到 {len(api_methods)} 个 Api 方法")
    
    # 生成 C 头文件
    c_header = generate_c_header(api_name, spi_methods, api_methods, spi_class, api_class)
    header_output = output_dir / f"ctp_{api_name.lower()}_c_api.h"
    header_output.write_text(c_header, encoding='utf-8')
    print(f"  生成 {header_output}")
    
    # 生成 C++ 实现文件
    cpp_impl = generate_cpp_impl(api_name, spi_methods, api_methods, spi_class, api_class)
    impl_output = output_dir / f"ctp_{api_name.lower()}_c_api.cpp"
    impl_output.write_text(cpp_impl, encoding='utf-8')
    print(f"  生成 {impl_output}")


def main():
    parser = argparse.ArgumentParser(description='CTP C++ API 转 C API 代码生成器')
    parser.add_argument('--input', '-i', default='../ctpapi/linux',
                       help='CTP 头文件目录 (默认: ../ctpapi/linux)')
    parser.add_argument('--output', '-o', default='./output',
                       help='输出目录 (默认: ./output)')
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
    process_api(input_dir, output_dir, "Md", 
                "CThostFtdcMdSpi", "CThostFtdcMdApi")
    print()
    
    # 处理 TraderApi
    process_api(input_dir, output_dir, "Trader",
                "CThostFtdcTraderSpi", "CThostFtdcTraderApi")
    
    print()
    print("完成!")
    return 0


if __name__ == '__main__':
    exit(main())
