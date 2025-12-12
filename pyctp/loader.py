"""
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
            r"C:\Windows\System32",  # 系统目录
            r"C:\CTP\lib",           # 常见安装路径
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
