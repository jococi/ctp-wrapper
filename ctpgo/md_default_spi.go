package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// 默认 SPI 空实现，可用于嵌入

// DefaultMdSpi 默认行情回调实现（空实现）
// 使用方式：嵌入到自定义结构体中，只需实现需要的方法
// 例如：type MySpi struct { DefaultMdSpi }
//
//	func (s *MySpi) OnRtnDepthMarketData(...) { ... }
type DefaultMdSpi struct{}

func (s *DefaultMdSpi) OnFrontConnected() {
	// 空实现
}

func (s *DefaultMdSpi) OnFrontDisconnected(nReason int32) {
	// 空实现
}

func (s *DefaultMdSpi) OnHeartBeatWarning(nTimeLapse int32) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspQryMulticastInstrument(pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspUnSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRspUnSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultMdSpi) OnRtnDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField) {
	// 空实现
}

func (s *DefaultMdSpi) OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField) {
	// 空实现
}
