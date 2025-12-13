"""
CTP 行情 API 封装

此文件由代码生成器自动生成，请勿手动修改
CTP 行情 API 封装
"""

import ctypes
import os
import threading
from abc import ABC, abstractmethod
from typing import Optional, List

from .loader import auto_load_library, get_md_lib_handle
from .struct import *
from .utils import *

# ========== 回调类型定义 ==========

# MdOnFrontConnectedCallback ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
MdOnFrontConnectedCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p)

# MdOnFrontDisconnectedCallback 0x2003 收到错误报文
MdOnFrontDisconnectedCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)

# MdOnHeartBeatWarningCallback 心跳超时警告。当长时间未收到报文时，该方法被调用。
MdOnHeartBeatWarningCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)

# MdOnRspUserLoginCallback 登录请求响应
MdOnRspUserLoginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserLoginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspUserLogoutCallback 登出请求响应
MdOnRspUserLogoutCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspQryMulticastInstrumentCallback 请求查询组播合约响应
MdOnRspQryMulticastInstrumentCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMulticastInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspErrorCallback 错误应答
MdOnRspErrorCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspSubMarketDataCallback 订阅行情应答
MdOnRspSubMarketDataCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspUnSubMarketDataCallback 取消订阅行情应答
MdOnRspUnSubMarketDataCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspSubForQuoteRspCallback 订阅询价应答
MdOnRspSubForQuoteRspCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRspUnSubForQuoteRspCallback 取消订阅询价应答
MdOnRspUnSubForQuoteRspCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# MdOnRtnDepthMarketDataCallback 深度行情通知
MdOnRtnDepthMarketDataCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcDepthMarketDataField))

# MdOnRtnForQuoteRspCallback 询价通知
MdOnRtnForQuoteRspCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteRspField))

# ========== 回调结构体定义 ==========

# MdSpiCallbacks 回调结构体（用于批量设置）
class MdSpiCallbacks(ctypes.Structure):
    """回调结构体（用于批量设置）"""
    _fields_ = [
        ("userData", ctypes.c_void_p),
        ("onFrontConnected", MdOnFrontConnectedCallback),
        ("onFrontDisconnected", MdOnFrontDisconnectedCallback),
        ("onHeartBeatWarning", MdOnHeartBeatWarningCallback),
        ("onRspUserLogin", MdOnRspUserLoginCallback),
        ("onRspUserLogout", MdOnRspUserLogoutCallback),
        ("onRspQryMulticastInstrument", MdOnRspQryMulticastInstrumentCallback),
        ("onRspError", MdOnRspErrorCallback),
        ("onRspSubMarketData", MdOnRspSubMarketDataCallback),
        ("onRspUnSubMarketData", MdOnRspUnSubMarketDataCallback),
        ("onRspSubForQuoteRsp", MdOnRspSubForQuoteRspCallback),
        ("onRspUnSubForQuoteRsp", MdOnRspUnSubForQuoteRspCallback),
        ("onRtnDepthMarketData", MdOnRtnDepthMarketDataCallback),
        ("onRtnForQuoteRsp", MdOnRtnForQuoteRspCallback),
    ]

# ========== MdSpi 接口 ==========

class MdSpi(ABC):
    """行情回调接口"""

    @abstractmethod
    def OnFrontConnected(self, ):
        """========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。"""
        pass

    @abstractmethod
    def OnFrontDisconnected(self, nReason: ctypes.c_int32):
        """0x2003 收到错误报文"""
        pass

    @abstractmethod
    def OnHeartBeatWarning(self, nTimeLapse: ctypes.c_int32):
        """心跳超时警告。当长时间未收到报文时，该方法被调用。"""
        pass

    @abstractmethod
    def OnRspUserLogin(self, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登录请求响应"""
        pass

    @abstractmethod
    def OnRspUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登出请求响应"""
        pass

    @abstractmethod
    def OnRspQryMulticastInstrument(self, pMulticastInstrument: ctypes.POINTER(CThostFtdcMulticastInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询组播合约响应"""
        pass

    @abstractmethod
    def OnRspError(self, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """错误应答"""
        pass

    @abstractmethod
    def OnRspSubMarketData(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """订阅行情应答"""
        pass

    @abstractmethod
    def OnRspUnSubMarketData(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """取消订阅行情应答"""
        pass

    @abstractmethod
    def OnRspSubForQuoteRsp(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """订阅询价应答"""
        pass

    @abstractmethod
    def OnRspUnSubForQuoteRsp(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """取消订阅询价应答"""
        pass

    @abstractmethod
    def OnRtnDepthMarketData(self, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField)):
        """深度行情通知"""
        pass

    @abstractmethod
    def OnRtnForQuoteRsp(self, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
        """询价通知"""
        pass

# ========== MdApi 类 ==========

class MdApi:
    """行情 API 封装"""

    def __init__(self, flow_path: str, using_udp: bool = False, multicast: bool = False):
        """创建行情 API 实例"""
        # 自动加载库（如果尚未加载）
        auto_load_library()

        self._handle: Optional[ctypes.c_void_p] = None
        self._spi: Optional[MdSpi] = None
        self._spi_handle: Optional[ctypes.c_void_p] = None
        self._user_data: int = _register_md_instance(self)
        self._lock = threading.RLock()

        # 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录
        # CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的
        # 所以转换为绝对路径更可靠
        abs_flow_path = flow_path
        if not os.path.isabs(flow_path):
            # 如果是相对路径，转换为基于当前工作目录的绝对路径
            abs_flow_path = os.path.abspath(flow_path)

        # 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）
        if abs_flow_path and not abs_flow_path.endswith(os.sep):
            abs_flow_path += os.sep

        # 确保目录存在（CTP API 需要这个目录来创建 flow 文件）
        try:
            os.makedirs(abs_flow_path, exist_ok=True)
        except OSError as e:
            # 如果创建目录失败，记录错误但继续（CTP API 可能会自己创建）
            print(f"警告: 无法创建 flow 目录 {abs_flow_path}: {e}")
            # 这里不抛出异常，让 CTP API 自己处理

        # 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收
        # CTP API 可能会在后续使用这个路径
        self._flow_path = abs_flow_path.encode('utf-8') + b'\0'

        # 调用 C 函数创建 API
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")

        # 获取函数指针
        func = lib.MdCreateFtdcMdApi
        func.argtypes = [ctypes.c_char_p, ctypes.c_bool, ctypes.c_bool]
        func.restype = ctypes.c_void_p

        self._handle = func(self._flow_path, using_udp, multicast)
        if self._handle is None:
            raise RuntimeError("Failed to create MdApi")

    # ========== API 方法 ==========

    # GetApiVersion 获取API的版本信息
    def GetApiVersion(self) -> str:
        """获取 API 版本"""
        lib = get_md_lib_handle()
        if lib is None:
            return ""
        func = lib.MdGetApiVersion
        func.argtypes = []
        func.restype = ctypes.c_char_p
        ptr = func()
        return go_string(ptr) if ptr else ""

    # Release 删除接口对象本身
    def Release(self):
        """释放 API 实例"""
        with self._lock:
            if self._handle:
                lib = get_md_lib_handle()
                if lib:
                    func = lib.MdRelease
                    func.argtypes = [ctypes.c_void_p]
                    func.restype = None
                    func(self._handle)
                self._handle = None
            _unregister_md_instance(self._user_data)

    # Init 初始化
    def Init(self, ):
        """初始化"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdInit
        func.argtypes = [ctypes.c_void_p]
        func.restype = None
        func(self._handle)

    # Join 等待接口线程结束运行
    def Join(self, ) -> ctypes.c_int32:
        """等待接口线程结束运行"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdJoin
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_int32
        return func(self._handle)

    # GetTradingDay 获取当前交易日
    def GetTradingDay(self) -> str:
        """获取交易日"""
        lib = get_md_lib_handle()
        if lib is None:
            return ""
        func = lib.MdGetTradingDay
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_char_p
        ptr = func(self._handle)
        return go_string(ptr) if ptr else ""

    # RegisterFront 注册前置机网络地址
    def RegisterFront(self, pszFrontAddress: str):
        """注册前置机网络地址"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdRegisterFront
        func.argtypes = [ctypes.c_void_p, ctypes.c_char_p]
        func.restype = None
        func(self._handle, c_string(pszFrontAddress))

    # RegisterNameServer 注册名字服务器网络地址
    def RegisterNameServer(self, pszNsAddress: str):
        """注册名字服务器网络地址"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdRegisterNameServer
        func.argtypes = [ctypes.c_void_p, ctypes.c_char_p]
        func.restype = None
        func(self._handle, c_string(pszNsAddress))

    # RegisterFensUserInfo 注册名字服务器用户信息
    def RegisterFensUserInfo(self, pFensUserInfo: ctypes.POINTER(CThostFtdcFensUserInfoField)):
        """注册名字服务器用户信息"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdRegisterFensUserInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcFensUserInfoField)]
        func.restype = None
        func(self._handle, pFensUserInfo)

    # SubscribeMarketData 订阅行情。
    def SubscribeMarketData(self, ppInstrumentID: List[str], nCount: ctypes.c_int32) -> int:
        """订阅行情。"""
        if len(ppInstrumentID) == 0:
            return 0
        # 将字符串数组转换为 C 字符串数组
        ptrs, _ = c_string_array(ppInstrumentID)
        lib = get_md_lib_handle()
        if lib is None:
            return -1
        func = lib.MdSubscribeMarketData
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]
        func.restype = ctypes.c_int32
        return func(self._handle, ptrs, nCount)

    # UnSubscribeMarketData 退订行情。
    def UnSubscribeMarketData(self, ppInstrumentID: List[str], nCount: ctypes.c_int32) -> int:
        """退订行情。"""
        if len(ppInstrumentID) == 0:
            return 0
        # 将字符串数组转换为 C 字符串数组
        ptrs, _ = c_string_array(ppInstrumentID)
        lib = get_md_lib_handle()
        if lib is None:
            return -1
        func = lib.MdUnSubscribeMarketData
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]
        func.restype = ctypes.c_int32
        return func(self._handle, ptrs, nCount)

    # SubscribeForQuoteRsp 订阅询价。
    def SubscribeForQuoteRsp(self, ppInstrumentID: List[str], nCount: ctypes.c_int32) -> int:
        """订阅询价。"""
        if len(ppInstrumentID) == 0:
            return 0
        # 将字符串数组转换为 C 字符串数组
        ptrs, _ = c_string_array(ppInstrumentID)
        lib = get_md_lib_handle()
        if lib is None:
            return -1
        func = lib.MdSubscribeForQuoteRsp
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]
        func.restype = ctypes.c_int32
        return func(self._handle, ptrs, nCount)

    # UnSubscribeForQuoteRsp 退订询价。
    def UnSubscribeForQuoteRsp(self, ppInstrumentID: List[str], nCount: ctypes.c_int32) -> int:
        """退订询价。"""
        if len(ppInstrumentID) == 0:
            return 0
        # 将字符串数组转换为 C 字符串数组
        ptrs, _ = c_string_array(ppInstrumentID)
        lib = get_md_lib_handle()
        if lib is None:
            return -1
        func = lib.MdUnSubscribeForQuoteRsp
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(ctypes.c_char_p)]
        func.restype = ctypes.c_int32
        return func(self._handle, ptrs, nCount)

    # ReqUserLogin 用户登录请求
    def ReqUserLogin(self, pReqUserLoginField: ctypes.POINTER(CThostFtdcReqUserLoginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户登录请求"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdReqUserLogin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginField, nRequestID)

    # ReqUserLogout 登出请求
    def ReqUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """登出请求"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdReqUserLogout
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pUserLogout, nRequestID)

    # ReqQryMulticastInstrument 请求查询组播合约
    def ReqQryMulticastInstrument(self, pQryMulticastInstrument: ctypes.POINTER(CThostFtdcQryMulticastInstrumentField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询组播合约"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdReqQryMulticastInstrument
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMulticastInstrumentField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryMulticastInstrument, nRequestID)

    # SpiCreate ========== Md SPI 函数 ========== 创建 SPI 实例
    def SpiCreate(self, ) -> ctypes.c_void_p:
        """========== Md SPI 函数 ========== 创建 SPI 实例"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdSpiCreate
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_void_p
        return func(self._handle)

    # SpiDestroy 销毁 SPI 实例
    def SpiDestroy(self, ):
        """销毁 SPI 实例"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdSpiDestroy
        func.argtypes = [ctypes.c_void_p]
        func.restype = None
        func(self._handle)

    # RegisterSpi 注册 SPI 到 API
    def RegisterSpi(self, spi: ctypes.c_void_p):
        """注册 SPI 到 API"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdRegisterSpi
        func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]
        func.restype = None
        func(self._handle, spi)

    # SpiSetCallbacks 批量设置回调
    def SpiSetCallbacks(self, callbacks: ctypes.POINTER(MdSpiCallbacks)):
        """批量设置回调"""
        lib = get_md_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.MdSpiSetCallbacks
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(MdSpiCallbacks)]
        func.restype = None
        func(self._handle, callbacks)

    def set_spi(self, spi: MdSpi):
        """设置回调接口"""
        with self._lock:
            self._spi = spi

            # 如果已有 C SPI 实例，先销毁
            if self._spi_handle:
                lib = get_md_lib_handle()
                if lib:
                    func = lib.MdSpiDestroy
                    func.argtypes = [ctypes.c_void_p]
                    func.restype = None
                    func(self._spi_handle)
                self._spi_handle = None

            # 创建新的 C SPI 实例
            lib = get_md_lib_handle()
            if lib is None:
                raise RuntimeError("CTP library not loaded")

            func = lib.MdSpiCreate
            func.argtypes = [ctypes.c_void_p]
            func.restype = ctypes.c_void_p
            self._spi_handle = func(ctypes.c_void_p(self._user_data))

            # 注册所有回调函数到 C SPI
            _register_md_callback(self._spi_handle, lib, "FrontConnected", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "FrontDisconnected", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "HeartBeatWarning", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspUserLogin", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspUserLogout", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspQryMulticastInstrument", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspError", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspSubMarketData", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspUnSubMarketData", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspSubForQuoteRsp", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RspUnSubForQuoteRsp", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RtnDepthMarketData", self._spi, self._user_data)
            _register_md_callback(self._spi_handle, lib, "RtnForQuoteRsp", self._spi, self._user_data)

            # 将 C SPI 注册到 API
            func = lib.MdRegisterSpi
            func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]
            func.restype = None
            func(self._handle, self._spi_handle)

# ========== 实例管理 ==========

_md_instances: dict = {}
_md_instances_lock = threading.RLock()
_md_next_id = 1

def _register_md_instance(api: MdApi) -> int:
    """注册行情 API 实例"""
    global _md_next_id
    with _md_instances_lock:
        instance_id = _md_next_id
        _md_next_id += 1
        _md_instances[instance_id] = api
        return instance_id

def _get_md_instance(user_data: int) -> Optional[MdApi]:
    """获取行情 API 实例"""
    with _md_instances_lock:
        return _md_instances.get(user_data)

def _unregister_md_instance(user_data: int):
    """注销行情 API 实例"""
    with _md_instances_lock:
        _md_instances.pop(user_data, None)

def _register_md_callback(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi: MdSpi, user_data: int):
    """注册回调函数到 C SPI"""
    # 实际实现在 md_callbacks.py 中
    from .md_callbacks import _register_md_callback_impl
    _register_md_callback_impl(spi_handle, lib, callback_name, spi, user_data)
