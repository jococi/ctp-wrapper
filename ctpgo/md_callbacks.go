package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 行情回调实现
// 使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// ========== 回调函数 ==========

// goMdOnFrontConnected 回调函数实现
func goMdOnFrontConnected(userData uintptr) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontConnected()
}

// goMdOnFrontDisconnected 回调函数实现
func goMdOnFrontDisconnected(userData uintptr, nReason int32) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontDisconnected(nReason)
}

// goMdOnHeartBeatWarning 回调函数实现
func goMdOnHeartBeatWarning(userData uintptr, nTimeLapse int32) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnHeartBeatWarning(nTimeLapse)
}

// goMdOnRspUserLogin 回调函数实现（C 调用约定版本）
func goMdOnRspUserLogin(userData uintptr, pRspUserLogin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogin((*CThostFtdcRspUserLoginField)(pRspUserLogin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspUserLogout 回调函数实现（C 调用约定版本）
func goMdOnRspUserLogout(userData uintptr, pUserLogout unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogout((*CThostFtdcUserLogoutField)(pUserLogout), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspQryMulticastInstrument 回调函数实现（C 调用约定版本）
func goMdOnRspQryMulticastInstrument(userData uintptr, pMulticastInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMulticastInstrument((*CThostFtdcMulticastInstrumentField)(pMulticastInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspError 回调函数实现（C 调用约定版本）
func goMdOnRspError(userData uintptr, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspError((*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspSubMarketData 回调函数实现（C 调用约定版本）
func goMdOnRspSubMarketData(userData uintptr, pSpecificInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSubMarketData((*CThostFtdcSpecificInstrumentField)(pSpecificInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspUnSubMarketData 回调函数实现（C 调用约定版本）
func goMdOnRspUnSubMarketData(userData uintptr, pSpecificInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUnSubMarketData((*CThostFtdcSpecificInstrumentField)(pSpecificInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspSubForQuoteRsp 回调函数实现（C 调用约定版本）
func goMdOnRspSubForQuoteRsp(userData uintptr, pSpecificInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSubForQuoteRsp((*CThostFtdcSpecificInstrumentField)(pSpecificInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRspUnSubForQuoteRsp 回调函数实现（C 调用约定版本）
func goMdOnRspUnSubForQuoteRsp(userData uintptr, pSpecificInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUnSubForQuoteRsp((*CThostFtdcSpecificInstrumentField)(pSpecificInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goMdOnRtnDepthMarketData 回调函数实现（C 调用约定版本）
func goMdOnRtnDepthMarketData(userData uintptr, pDepthMarketData unsafe.Pointer) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnDepthMarketData((*CThostFtdcDepthMarketDataField)(pDepthMarketData))
}

// goMdOnRtnForQuoteRsp 回调函数实现（C 调用约定版本）
func goMdOnRtnForQuoteRsp(userData uintptr, pForQuoteRsp unsafe.Pointer) {
	api := getMdInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnForQuoteRsp((*CThostFtdcForQuoteRspField)(pForQuoteRsp))
}

// ========== 辅助函数：使用 purego.NewCallback 获取 C 函数指针 ==========
// 这些函数使用 purego.NewCallback 将 Go 函数转换为 C 函数指针，无需 CGO
// purego.NewCallback 返回 uintptr，需要转换为函数类型
// 注意：purego.NewCallback 不支持 unsafe.Pointer 参数，需要用具体指针类型的 wrapper

// GetGoMdOnFrontConnected 获取 goMdOnFrontConnected 的 C 函数指针
func GetGoMdOnFrontConnected() uintptr {
	return purego.NewCallback(goMdOnFrontConnected)
}

// GetGoMdOnFrontDisconnected 获取 goMdOnFrontDisconnected 的 C 函数指针
func GetGoMdOnFrontDisconnected() uintptr {
	return purego.NewCallback(goMdOnFrontDisconnected)
}

// GetGoMdOnHeartBeatWarning 获取 goMdOnHeartBeatWarning 的 C 函数指针
func GetGoMdOnHeartBeatWarning() uintptr {
	return purego.NewCallback(goMdOnHeartBeatWarning)
}

// GetGoMdOnRspUserLogin 获取 goMdOnRspUserLogin 的 C 函数指针
func GetGoMdOnRspUserLogin() uintptr {
	wrapper := func(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspUserLogin(userData, unsafe.Pointer(pRspUserLogin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspUserLogout 获取 goMdOnRspUserLogout 的 C 函数指针
func GetGoMdOnRspUserLogout() uintptr {
	wrapper := func(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspUserLogout(userData, unsafe.Pointer(pUserLogout), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspQryMulticastInstrument 获取 goMdOnRspQryMulticastInstrument 的 C 函数指针
func GetGoMdOnRspQryMulticastInstrument() uintptr {
	wrapper := func(userData uintptr, pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspQryMulticastInstrument(userData, unsafe.Pointer(pMulticastInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspError 获取 goMdOnRspError 的 C 函数指针
func GetGoMdOnRspError() uintptr {
	wrapper := func(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspError(userData, unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspSubMarketData 获取 goMdOnRspSubMarketData 的 C 函数指针
func GetGoMdOnRspSubMarketData() uintptr {
	wrapper := func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspSubMarketData(userData, unsafe.Pointer(pSpecificInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspUnSubMarketData 获取 goMdOnRspUnSubMarketData 的 C 函数指针
func GetGoMdOnRspUnSubMarketData() uintptr {
	wrapper := func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspUnSubMarketData(userData, unsafe.Pointer(pSpecificInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspSubForQuoteRsp 获取 goMdOnRspSubForQuoteRsp 的 C 函数指针
func GetGoMdOnRspSubForQuoteRsp() uintptr {
	wrapper := func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspSubForQuoteRsp(userData, unsafe.Pointer(pSpecificInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRspUnSubForQuoteRsp 获取 goMdOnRspUnSubForQuoteRsp 的 C 函数指针
func GetGoMdOnRspUnSubForQuoteRsp() uintptr {
	wrapper := func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goMdOnRspUnSubForQuoteRsp(userData, unsafe.Pointer(pSpecificInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRtnDepthMarketData 获取 goMdOnRtnDepthMarketData 的 C 函数指针
func GetGoMdOnRtnDepthMarketData() uintptr {
	wrapper := func(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField) {
		goMdOnRtnDepthMarketData(userData, unsafe.Pointer(pDepthMarketData))
	}
	return purego.NewCallback(wrapper)
}

// GetGoMdOnRtnForQuoteRsp 获取 goMdOnRtnForQuoteRsp 的 C 函数指针
func GetGoMdOnRtnForQuoteRsp() uintptr {
	wrapper := func(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField) {
		goMdOnRtnForQuoteRsp(userData, unsafe.Pointer(pForQuoteRsp))
	}
	return purego.NewCallback(wrapper)
}
