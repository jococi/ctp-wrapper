"""
默认行情 SPI 空实现

此文件由代码生成器自动生成，请勿手动修改
默认 SPI 空实现，可用于嵌入
"""

import ctypes
from .md_api import MdSpi
from .struct import *

class DefaultMdSpi(MdSpi):
    """默认行情回调实现（空实现）"""
    
    # 使用方式：继承此类，只需实现需要的方法
    
    def OnFrontConnected(self, ):
        """空实现"""
        pass

    def OnFrontDisconnected(self, nReason: ctypes.c_int32):
        """空实现"""
        pass

    def OnHeartBeatWarning(self, nTimeLapse: ctypes.c_int32):
        """空实现"""
        pass

    def OnRspUserLogin(self, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryMulticastInstrument(self, pMulticastInstrument: ctypes.POINTER(CThostFtdcMulticastInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspError(self, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspSubMarketData(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUnSubMarketData(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspSubForQuoteRsp(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUnSubForQuoteRsp(self, pSpecificInstrument: ctypes.POINTER(CThostFtdcSpecificInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRtnDepthMarketData(self, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField)):
        """空实现"""
        pass

    def OnRtnForQuoteRsp(self, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
        """空实现"""
        pass
