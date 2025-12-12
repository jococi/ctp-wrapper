"""
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
        null_pos = b.find(b'\0')
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
        null_pos = b.find(b'\0')
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
    return b + b'\0' * (size - len(b))


def bool_to_int(b: bool) -> int:
    """将 bool 转换为 int（C 风格）"""
    return 1 if b else 0


def int_to_bool(i: int) -> bool:
    """将 int 转换为 bool（C 风格）"""
    return i != 0
