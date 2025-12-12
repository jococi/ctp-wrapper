"""
CTP 行情回调实现

此文件由代码生成器自动生成，请勿手动修改
CTP 行情回调实现
"""

import ctypes
from .md_api import _get_md_instance
from .struct import *

# ========== 回调包装函数 ==========

def _go_md_OnFrontConnected(userData: ctypes.c_void_p):
    """回调函数实现: ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnFrontConnected()
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnFrontDisconnected(userData: ctypes.c_void_p, nReason: ctypes.c_int32):
    """回调函数实现: 0x2003 收到错误报文"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnFrontDisconnected(nReason)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnHeartBeatWarning(userData: ctypes.c_void_p, nTimeLapse: ctypes.c_int32):
    """回调函数实现: 心跳超时警告。当长时间未收到报文时，该方法被调用。"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnHeartBeatWarning(nTimeLapse)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspUserLogin(userData: ctypes.c_void_p, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 登录请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserLogin(pRspUserLogin.contents if pRspUserLogin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspUserLogout(userData: ctypes.c_void_p, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 登出请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserLogout(pUserLogout.contents if pUserLogout else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspQryMulticastInstrument(userData: ctypes.c_void_p, pMulticastInstrument: ctypes.POINTER(CThostFtdcMulticastInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询组播合约响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryMulticastInstrument(pMulticastInstrument.contents if pMulticastInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspError(userData: ctypes.c_void_p, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 错误应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspError(pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspSubMarketData(userData: ctypes.c_void_p, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 订阅行情应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspSubMarketData(pSpecificInstrument.contents if pSpecificInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspUnSubMarketData(userData: ctypes.c_void_p, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 取消订阅行情应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUnSubMarketData(pSpecificInstrument.contents if pSpecificInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspSubForQuoteRsp(userData: ctypes.c_void_p, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 订阅询价应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspSubForQuoteRsp(pSpecificInstrument.contents if pSpecificInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRspUnSubForQuoteRsp(userData: ctypes.c_void_p, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 取消订阅询价应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUnSubForQuoteRsp(pSpecificInstrument.contents if pSpecificInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRtnDepthMarketData(userData: ctypes.c_void_p, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField)):
    """回调函数实现: 深度行情通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnDepthMarketData(pDepthMarketData.contents if pDepthMarketData else None)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

def _go_md_OnRtnForQuoteRsp(userData: ctypes.c_void_p, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
    """回调函数实现: 询价通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_md_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnForQuoteRsp(pForQuoteRsp.contents if pForQuoteRsp else None)
    except Exception as e:
        # 回调异常不应该影响 C 层
        import traceback
        traceback.print_exc()

# ========== 回调注册函数 ==========

def _register_md_callback_impl(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi, user_data: int):
    """注册回调函数到 C SPI（内部实现）"""
    # 回调函数映射表
    callback_map = {
        "FrontConnected": (_go_md_OnFrontConnected, ctypes.CFUNCTYPE(None, ctypes.c_void_p)),
        "FrontDisconnected": (_go_md_OnFrontDisconnected, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)),
        "HeartBeatWarning": (_go_md_OnHeartBeatWarning, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)),
        "RspUserLogin": (_go_md_OnRspUserLogin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserLoginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUserLogout": (_go_md_OnRspUserLogout, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryMulticastInstrument": (_go_md_OnRspQryMulticastInstrument, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMulticastInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspError": (_go_md_OnRspError, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspSubMarketData": (_go_md_OnRspSubMarketData, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUnSubMarketData": (_go_md_OnRspUnSubMarketData, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspSubForQuoteRsp": (_go_md_OnRspSubForQuoteRsp, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUnSubForQuoteRsp": (_go_md_OnRspUnSubForQuoteRsp, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSpecificInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RtnDepthMarketData": (_go_md_OnRtnDepthMarketData, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcDepthMarketDataField))),
        "RtnForQuoteRsp": (_go_md_OnRtnForQuoteRsp, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteRspField))),
    }

    if callback_name not in callback_map:
        return

    callback_func, callback_type = callback_map[callback_name]

    # 创建 CFUNCTYPE 回调实例
    c_callback = callback_type(callback_func)

    # 注册到 C SPI
    func_name = f"MdSpiSetOn{callback_name}"
    if hasattr(lib, func_name):
        func = getattr(lib, func_name)
        func.argtypes = [ctypes.c_void_p, callback_type]
        func.restype = None
        func(spi_handle, c_callback)

    # 保存回调引用，防止被 GC 回收
    if not hasattr(spi, "_callbacks"):
        spi._callbacks = []
    spi._callbacks.append(c_callback)
