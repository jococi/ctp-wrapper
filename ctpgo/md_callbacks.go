package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 行情回调实现

import "unsafe"

// #include <stdint.h>
import "C"

// ========== 回调函数 ==========

//export goOnFrontConnected
func goOnFrontConnected(userData uintptr) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontConnected()
}

//export goOnFrontDisconnected
func goOnFrontDisconnected(userData uintptr, nReason int32) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontDisconnected(nReason)
}

//export goOnHeartBeatWarning
func goOnHeartBeatWarning(userData uintptr, nTimeLapse int32) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnHeartBeatWarning(nTimeLapse)
}

//export goOnRspUserLogin
func goOnRspUserLogin(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogin(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUserLogout
func goOnRspUserLogout(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogout(pUserLogout, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryMulticastInstrument
func goOnRspQryMulticastInstrument(userData uintptr, pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMulticastInstrument(pMulticastInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspError
func goOnRspError(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspError(pRspInfo, nRequestID, bIsLast)
}

//export goOnRspSubMarketData
func goOnRspSubMarketData(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSubMarketData(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUnSubMarketData
func goOnRspUnSubMarketData(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUnSubMarketData(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspSubForQuoteRsp
func goOnRspSubForQuoteRsp(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSubForQuoteRsp(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUnSubForQuoteRsp
func goOnRspUnSubForQuoteRsp(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUnSubForQuoteRsp(pSpecificInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRtnDepthMarketData
func goOnRtnDepthMarketData(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnDepthMarketData(pDepthMarketData)
}

//export goOnRtnForQuoteRsp
func goOnRtnForQuoteRsp(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnForQuoteRsp(pForQuoteRsp)
}
