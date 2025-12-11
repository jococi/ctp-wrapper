package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 交易回调实现
// 使用 purego.NewCallback 替代 CGO，支持 Windows 平台无需 C 编译器

import (
	"unsafe"

	"github.com/ebitengine/purego"
)

// ========== 回调函数 ==========

// goTraderOnFrontConnected 回调函数实现
func goTraderOnFrontConnected(userData uintptr) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontConnected()
}

// goTraderOnFrontDisconnected 回调函数实现
func goTraderOnFrontDisconnected(userData uintptr, nReason int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontDisconnected(nReason)
}

// goTraderOnHeartBeatWarning 回调函数实现
func goTraderOnHeartBeatWarning(userData uintptr, nTimeLapse int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnHeartBeatWarning(nTimeLapse)
}

// goTraderOnRspAuthenticate 回调函数实现（C 调用约定版本）
func goTraderOnRspAuthenticate(userData uintptr, pRspAuthenticateField unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspAuthenticate((*CThostFtdcRspAuthenticateField)(pRspAuthenticateField), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspUserLogin 回调函数实现（C 调用约定版本）
func goTraderOnRspUserLogin(userData uintptr, pRspUserLogin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogin((*CThostFtdcRspUserLoginField)(pRspUserLogin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspUserLogout 回调函数实现（C 调用约定版本）
func goTraderOnRspUserLogout(userData uintptr, pUserLogout unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogout((*CThostFtdcUserLogoutField)(pUserLogout), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspUserPasswordUpdate 回调函数实现（C 调用约定版本）
func goTraderOnRspUserPasswordUpdate(userData uintptr, pUserPasswordUpdate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserPasswordUpdate((*CThostFtdcUserPasswordUpdateField)(pUserPasswordUpdate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspTradingAccountPasswordUpdate 回调函数实现（C 调用约定版本）
func goTraderOnRspTradingAccountPasswordUpdate(userData uintptr, pTradingAccountPasswordUpdate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspTradingAccountPasswordUpdate((*CThostFtdcTradingAccountPasswordUpdateField)(pTradingAccountPasswordUpdate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspUserAuthMethod 回调函数实现（C 调用约定版本）
func goTraderOnRspUserAuthMethod(userData uintptr, pRspUserAuthMethod unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserAuthMethod((*CThostFtdcRspUserAuthMethodField)(pRspUserAuthMethod), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspGenUserCaptcha 回调函数实现（C 调用约定版本）
func goTraderOnRspGenUserCaptcha(userData uintptr, pRspGenUserCaptcha unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserCaptcha((*CThostFtdcRspGenUserCaptchaField)(pRspGenUserCaptcha), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspGenUserText 回调函数实现（C 调用约定版本）
func goTraderOnRspGenUserText(userData uintptr, pRspGenUserText unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserText((*CThostFtdcRspGenUserTextField)(pRspGenUserText), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspOrderInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspOrderInsert(userData uintptr, pInputOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderInsert((*CThostFtdcInputOrderField)(pInputOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspParkedOrderInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspParkedOrderInsert(userData uintptr, pParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderInsert((*CThostFtdcParkedOrderField)(pParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspParkedOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspParkedOrderAction(userData uintptr, pParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderAction((*CThostFtdcParkedOrderActionField)(pParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspOrderAction(userData uintptr, pInputOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderAction((*CThostFtdcInputOrderActionField)(pInputOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryMaxOrderVolume 回调函数实现（C 调用约定版本）
func goTraderOnRspQryMaxOrderVolume(userData uintptr, pQryMaxOrderVolume unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMaxOrderVolume((*CThostFtdcQryMaxOrderVolumeField)(pQryMaxOrderVolume), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspSettlementInfoConfirm 回调函数实现（C 调用约定版本）
func goTraderOnRspSettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(pSettlementInfoConfirm), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspRemoveParkedOrder 回调函数实现（C 调用约定版本）
func goTraderOnRspRemoveParkedOrder(userData uintptr, pRemoveParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrder((*CThostFtdcRemoveParkedOrderField)(pRemoveParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspRemoveParkedOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspRemoveParkedOrderAction(userData uintptr, pRemoveParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrderAction((*CThostFtdcRemoveParkedOrderActionField)(pRemoveParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspExecOrderInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspExecOrderInsert(userData uintptr, pInputExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderInsert((*CThostFtdcInputExecOrderField)(pInputExecOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspExecOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspExecOrderAction(userData uintptr, pInputExecOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderAction((*CThostFtdcInputExecOrderActionField)(pInputExecOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspForQuoteInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspForQuoteInsert(userData uintptr, pInputForQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspForQuoteInsert((*CThostFtdcInputForQuoteField)(pInputForQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQuoteInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspQuoteInsert(userData uintptr, pInputQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteInsert((*CThostFtdcInputQuoteField)(pInputQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQuoteAction 回调函数实现（C 调用约定版本）
func goTraderOnRspQuoteAction(userData uintptr, pInputQuoteAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteAction((*CThostFtdcInputQuoteActionField)(pInputQuoteAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspBatchOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspBatchOrderAction(userData uintptr, pInputBatchOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspBatchOrderAction((*CThostFtdcInputBatchOrderActionField)(pInputBatchOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspOptionSelfCloseInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(pInputOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspOptionSelfCloseAction 回调函数实现（C 调用约定版本）
func goTraderOnRspOptionSelfCloseAction(userData uintptr, pInputOptionSelfCloseAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseAction((*CThostFtdcInputOptionSelfCloseActionField)(pInputOptionSelfCloseAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspCombActionInsert 回调函数实现（C 调用约定版本）
func goTraderOnRspCombActionInsert(userData uintptr, pInputCombAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspCombActionInsert((*CThostFtdcInputCombActionField)(pInputCombAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryOrder 回调函数实现（C 调用约定版本）
func goTraderOnRspQryOrder(userData uintptr, pOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOrder((*CThostFtdcOrderField)(pOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTrade 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTrade(userData uintptr, pTrade unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTrade((*CThostFtdcTradeField)(pTrade), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorPosition 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorPosition(userData uintptr, pInvestorPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPosition((*CThostFtdcInvestorPositionField)(pInvestorPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTradingAccount 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTradingAccount(userData uintptr, pTradingAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingAccount((*CThostFtdcTradingAccountField)(pTradingAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestor 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestor(userData uintptr, pInvestor unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestor((*CThostFtdcInvestorField)(pInvestor), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTradingCode 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTradingCode(userData uintptr, pTradingCode unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingCode((*CThostFtdcTradingCodeField)(pTradingCode), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInstrumentMarginRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInstrumentMarginRate(userData uintptr, pInstrumentMarginRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentMarginRate((*CThostFtdcInstrumentMarginRateField)(pInstrumentMarginRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInstrumentCommissionRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInstrumentCommissionRate(userData uintptr, pInstrumentCommissionRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentCommissionRate((*CThostFtdcInstrumentCommissionRateField)(pInstrumentCommissionRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryExchange 回调函数实现（C 调用约定版本）
func goTraderOnRspQryExchange(userData uintptr, pExchange unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchange((*CThostFtdcExchangeField)(pExchange), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryProduct 回调函数实现（C 调用约定版本）
func goTraderOnRspQryProduct(userData uintptr, pProduct unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProduct((*CThostFtdcProductField)(pProduct), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInstrument 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInstrument(userData uintptr, pInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrument((*CThostFtdcInstrumentField)(pInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryDepthMarketData 回调函数实现（C 调用约定版本）
func goTraderOnRspQryDepthMarketData(userData uintptr, pDepthMarketData unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryDepthMarketData((*CThostFtdcDepthMarketDataField)(pDepthMarketData), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTraderOffer 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTraderOffer(userData uintptr, pTraderOffer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTraderOffer((*CThostFtdcTraderOfferField)(pTraderOffer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySettlementInfo 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySettlementInfo(userData uintptr, pSettlementInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfo((*CThostFtdcSettlementInfoField)(pSettlementInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTransferBank 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTransferBank(userData uintptr, pTransferBank unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferBank((*CThostFtdcTransferBankField)(pTransferBank), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorPositionDetail 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorPositionDetail(userData uintptr, pInvestorPositionDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionDetail((*CThostFtdcInvestorPositionDetailField)(pInvestorPositionDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryNotice 回调函数实现（C 调用约定版本）
func goTraderOnRspQryNotice(userData uintptr, pNotice unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryNotice((*CThostFtdcNoticeField)(pNotice), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySettlementInfoConfirm 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(pSettlementInfoConfirm), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorPositionCombineDetail 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorPositionCombineDetail(userData uintptr, pInvestorPositionCombineDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionCombineDetail((*CThostFtdcInvestorPositionCombineDetailField)(pInvestorPositionCombineDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryCFMMCTradingAccountKey 回调函数实现（C 调用约定版本）
func goTraderOnRspQryCFMMCTradingAccountKey(userData uintptr, pCFMMCTradingAccountKey unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCFMMCTradingAccountKey((*CThostFtdcCFMMCTradingAccountKeyField)(pCFMMCTradingAccountKey), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryEWarrantOffset 回调函数实现（C 调用约定版本）
func goTraderOnRspQryEWarrantOffset(userData uintptr, pEWarrantOffset unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryEWarrantOffset((*CThostFtdcEWarrantOffsetField)(pEWarrantOffset), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorProductGroupMargin 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorProductGroupMargin(userData uintptr, pInvestorProductGroupMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProductGroupMargin((*CThostFtdcInvestorProductGroupMarginField)(pInvestorProductGroupMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryExchangeMarginRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryExchangeMarginRate(userData uintptr, pExchangeMarginRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRate((*CThostFtdcExchangeMarginRateField)(pExchangeMarginRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryExchangeMarginRateAdjust 回调函数实现（C 调用约定版本）
func goTraderOnRspQryExchangeMarginRateAdjust(userData uintptr, pExchangeMarginRateAdjust unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRateAdjust((*CThostFtdcExchangeMarginRateAdjustField)(pExchangeMarginRateAdjust), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryExchangeRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryExchangeRate(userData uintptr, pExchangeRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeRate((*CThostFtdcExchangeRateField)(pExchangeRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySecAgentACIDMap 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySecAgentACIDMap(userData uintptr, pSecAgentACIDMap unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentACIDMap((*CThostFtdcSecAgentACIDMapField)(pSecAgentACIDMap), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryProductExchRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryProductExchRate(userData uintptr, pProductExchRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductExchRate((*CThostFtdcProductExchRateField)(pProductExchRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryProductGroup 回调函数实现（C 调用约定版本）
func goTraderOnRspQryProductGroup(userData uintptr, pProductGroup unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductGroup((*CThostFtdcProductGroupField)(pProductGroup), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryMMInstrumentCommissionRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryMMInstrumentCommissionRate(userData uintptr, pMMInstrumentCommissionRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMInstrumentCommissionRate((*CThostFtdcMMInstrumentCommissionRateField)(pMMInstrumentCommissionRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryMMOptionInstrCommRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryMMOptionInstrCommRate(userData uintptr, pMMOptionInstrCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMOptionInstrCommRate((*CThostFtdcMMOptionInstrCommRateField)(pMMOptionInstrCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInstrumentOrderCommRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInstrumentOrderCommRate(userData uintptr, pInstrumentOrderCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentOrderCommRate((*CThostFtdcInstrumentOrderCommRateField)(pInstrumentOrderCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySecAgentTradingAccount 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySecAgentTradingAccount(userData uintptr, pTradingAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradingAccount((*CThostFtdcTradingAccountField)(pTradingAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySecAgentCheckMode 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySecAgentCheckMode(userData uintptr, pSecAgentCheckMode unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentCheckMode((*CThostFtdcSecAgentCheckModeField)(pSecAgentCheckMode), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySecAgentTradeInfo 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySecAgentTradeInfo(userData uintptr, pSecAgentTradeInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradeInfo((*CThostFtdcSecAgentTradeInfoField)(pSecAgentTradeInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryOptionInstrTradeCost 回调函数实现（C 调用约定版本）
func goTraderOnRspQryOptionInstrTradeCost(userData uintptr, pOptionInstrTradeCost unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrTradeCost((*CThostFtdcOptionInstrTradeCostField)(pOptionInstrTradeCost), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryOptionInstrCommRate 回调函数实现（C 调用约定版本）
func goTraderOnRspQryOptionInstrCommRate(userData uintptr, pOptionInstrCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrCommRate((*CThostFtdcOptionInstrCommRateField)(pOptionInstrCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryExecOrder 回调函数实现（C 调用约定版本）
func goTraderOnRspQryExecOrder(userData uintptr, pExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExecOrder((*CThostFtdcExecOrderField)(pExecOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryForQuote 回调函数实现（C 调用约定版本）
func goTraderOnRspQryForQuote(userData uintptr, pForQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryForQuote((*CThostFtdcForQuoteField)(pForQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryQuote 回调函数实现（C 调用约定版本）
func goTraderOnRspQryQuote(userData uintptr, pQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryQuote((*CThostFtdcQuoteField)(pQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryOptionSelfClose 回调函数实现（C 调用约定版本）
func goTraderOnRspQryOptionSelfClose(userData uintptr, pOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionSelfClose((*CThostFtdcOptionSelfCloseField)(pOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestUnit 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestUnit(userData uintptr, pInvestUnit unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestUnit((*CThostFtdcInvestUnitField)(pInvestUnit), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryCombInstrumentGuard 回调函数实现（C 调用约定版本）
func goTraderOnRspQryCombInstrumentGuard(userData uintptr, pCombInstrumentGuard unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombInstrumentGuard((*CThostFtdcCombInstrumentGuardField)(pCombInstrumentGuard), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryCombAction 回调函数实现（C 调用约定版本）
func goTraderOnRspQryCombAction(userData uintptr, pCombAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombAction((*CThostFtdcCombActionField)(pCombAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTransferSerial 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTransferSerial(userData uintptr, pTransferSerial unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferSerial((*CThostFtdcTransferSerialField)(pTransferSerial), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryAccountregister 回调函数实现（C 调用约定版本）
func goTraderOnRspQryAccountregister(userData uintptr, pAccountregister unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryAccountregister((*CThostFtdcAccountregisterField)(pAccountregister), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspError 回调函数实现（C 调用约定版本）
func goTraderOnRspError(userData uintptr, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspError((*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRtnOrder 回调函数实现（C 调用约定版本）
func goTraderOnRtnOrder(userData uintptr, pOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOrder((*CThostFtdcOrderField)(pOrder))
}

// goTraderOnRtnTrade 回调函数实现（C 调用约定版本）
func goTraderOnRtnTrade(userData uintptr, pTrade unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTrade((*CThostFtdcTradeField)(pTrade))
}

// goTraderOnErrRtnOrderInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnOrderInsert(userData uintptr, pInputOrder unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderInsert((*CThostFtdcInputOrderField)(pInputOrder), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnOrderAction(userData uintptr, pOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderAction((*CThostFtdcOrderActionField)(pOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnInstrumentStatus 回调函数实现（C 调用约定版本）
func goTraderOnRtnInstrumentStatus(userData uintptr, pInstrumentStatus unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnInstrumentStatus((*CThostFtdcInstrumentStatusField)(pInstrumentStatus))
}

// goTraderOnRtnBulletin 回调函数实现（C 调用约定版本）
func goTraderOnRtnBulletin(userData uintptr, pBulletin unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnBulletin((*CThostFtdcBulletinField)(pBulletin))
}

// goTraderOnRtnTradingNotice 回调函数实现（C 调用约定版本）
func goTraderOnRtnTradingNotice(userData uintptr, pTradingNoticeInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTradingNotice((*CThostFtdcTradingNoticeInfoField)(pTradingNoticeInfo))
}

// goTraderOnRtnErrorConditionalOrder 回调函数实现（C 调用约定版本）
func goTraderOnRtnErrorConditionalOrder(userData uintptr, pErrorConditionalOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnErrorConditionalOrder((*CThostFtdcErrorConditionalOrderField)(pErrorConditionalOrder))
}

// goTraderOnRtnExecOrder 回调函数实现（C 调用约定版本）
func goTraderOnRtnExecOrder(userData uintptr, pExecOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnExecOrder((*CThostFtdcExecOrderField)(pExecOrder))
}

// goTraderOnErrRtnExecOrderInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnExecOrderInsert(userData uintptr, pInputExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderInsert((*CThostFtdcInputExecOrderField)(pInputExecOrder), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnExecOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnExecOrderAction(userData uintptr, pExecOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderAction((*CThostFtdcExecOrderActionField)(pExecOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnForQuoteInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnForQuoteInsert(userData uintptr, pInputForQuote unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnForQuoteInsert((*CThostFtdcInputForQuoteField)(pInputForQuote), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnQuote 回调函数实现（C 调用约定版本）
func goTraderOnRtnQuote(userData uintptr, pQuote unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQuote((*CThostFtdcQuoteField)(pQuote))
}

// goTraderOnErrRtnQuoteInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnQuoteInsert(userData uintptr, pInputQuote unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteInsert((*CThostFtdcInputQuoteField)(pInputQuote), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnQuoteAction 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnQuoteAction(userData uintptr, pQuoteAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteAction((*CThostFtdcQuoteActionField)(pQuoteAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnForQuoteRsp 回调函数实现（C 调用约定版本）
func goTraderOnRtnForQuoteRsp(userData uintptr, pForQuoteRsp unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnForQuoteRsp((*CThostFtdcForQuoteRspField)(pForQuoteRsp))
}

// goTraderOnRtnCFMMCTradingAccountToken 回调函数实现（C 调用约定版本）
func goTraderOnRtnCFMMCTradingAccountToken(userData uintptr, pCFMMCTradingAccountToken unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCFMMCTradingAccountToken((*CThostFtdcCFMMCTradingAccountTokenField)(pCFMMCTradingAccountToken))
}

// goTraderOnErrRtnBatchOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnBatchOrderAction(userData uintptr, pBatchOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBatchOrderAction((*CThostFtdcBatchOrderActionField)(pBatchOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnOptionSelfClose 回调函数实现（C 调用约定版本）
func goTraderOnRtnOptionSelfClose(userData uintptr, pOptionSelfClose unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOptionSelfClose((*CThostFtdcOptionSelfCloseField)(pOptionSelfClose))
}

// goTraderOnErrRtnOptionSelfCloseInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(pInputOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnOptionSelfCloseAction 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnOptionSelfCloseAction(userData uintptr, pOptionSelfCloseAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseAction((*CThostFtdcOptionSelfCloseActionField)(pOptionSelfCloseAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnCombAction 回调函数实现（C 调用约定版本）
func goTraderOnRtnCombAction(userData uintptr, pCombAction unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCombAction((*CThostFtdcCombActionField)(pCombAction))
}

// goTraderOnErrRtnCombActionInsert 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnCombActionInsert(userData uintptr, pInputCombAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnCombActionInsert((*CThostFtdcInputCombActionField)(pInputCombAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRspQryContractBank 回调函数实现（C 调用约定版本）
func goTraderOnRspQryContractBank(userData uintptr, pContractBank unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryContractBank((*CThostFtdcContractBankField)(pContractBank), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryParkedOrder 回调函数实现（C 调用约定版本）
func goTraderOnRspQryParkedOrder(userData uintptr, pParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrder((*CThostFtdcParkedOrderField)(pParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryParkedOrderAction 回调函数实现（C 调用约定版本）
func goTraderOnRspQryParkedOrderAction(userData uintptr, pParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrderAction((*CThostFtdcParkedOrderActionField)(pParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryTradingNotice 回调函数实现（C 调用约定版本）
func goTraderOnRspQryTradingNotice(userData uintptr, pTradingNotice unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingNotice((*CThostFtdcTradingNoticeField)(pTradingNotice), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryBrokerTradingParams 回调函数实现（C 调用约定版本）
func goTraderOnRspQryBrokerTradingParams(userData uintptr, pBrokerTradingParams unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingParams((*CThostFtdcBrokerTradingParamsField)(pBrokerTradingParams), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryBrokerTradingAlgos 回调函数实现（C 调用约定版本）
func goTraderOnRspQryBrokerTradingAlgos(userData uintptr, pBrokerTradingAlgos unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingAlgos((*CThostFtdcBrokerTradingAlgosField)(pBrokerTradingAlgos), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQueryCFMMCTradingAccountToken 回调函数实现（C 调用约定版本）
func goTraderOnRspQueryCFMMCTradingAccountToken(userData uintptr, pQueryCFMMCTradingAccountToken unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryCFMMCTradingAccountToken((*CThostFtdcQueryCFMMCTradingAccountTokenField)(pQueryCFMMCTradingAccountToken), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRtnFromBankToFutureByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnFromBankToFutureByBank(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByBank((*CThostFtdcRspTransferField)(pRspTransfer))
}

// goTraderOnRtnFromFutureToBankByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnFromFutureToBankByBank(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByBank((*CThostFtdcRspTransferField)(pRspTransfer))
}

// goTraderOnRtnRepealFromBankToFutureByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromBankToFutureByBank(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByBank((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRtnRepealFromFutureToBankByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromFutureToBankByBank(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByBank((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRtnFromBankToFutureByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRtnFromBankToFutureByFuture(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByFuture((*CThostFtdcRspTransferField)(pRspTransfer))
}

// goTraderOnRtnFromFutureToBankByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRtnFromFutureToBankByFuture(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByFuture((*CThostFtdcRspTransferField)(pRspTransfer))
}

// goTraderOnRtnRepealFromBankToFutureByFutureManual 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromBankToFutureByFutureManual(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFutureManual((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRtnRepealFromFutureToBankByFutureManual 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromFutureToBankByFutureManual(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFutureManual((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRtnQueryBankBalanceByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRtnQueryBankBalanceByFuture(userData uintptr, pNotifyQueryAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQueryBankBalanceByFuture((*CThostFtdcNotifyQueryAccountField)(pNotifyQueryAccount))
}

// goTraderOnErrRtnBankToFutureByFuture 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnBankToFutureByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBankToFutureByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnFutureToBankByFuture 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnFutureToBankByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnFutureToBankByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnRepealBankToFutureByFutureManual 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnRepealBankToFutureByFutureManual(userData uintptr, pReqRepeal unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealBankToFutureByFutureManual((*CThostFtdcReqRepealField)(pReqRepeal), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnRepealFutureToBankByFutureManual 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnRepealFutureToBankByFutureManual(userData uintptr, pReqRepeal unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealFutureToBankByFutureManual((*CThostFtdcReqRepealField)(pReqRepeal), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnErrRtnQueryBankBalanceByFuture 回调函数实现（C 调用约定版本）
func goTraderOnErrRtnQueryBankBalanceByFuture(userData uintptr, pReqQueryAccount unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQueryBankBalanceByFuture((*CThostFtdcReqQueryAccountField)(pReqQueryAccount), (*CThostFtdcRspInfoField)(pRspInfo))
}

// goTraderOnRtnRepealFromBankToFutureByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromBankToFutureByFuture(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFuture((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRtnRepealFromFutureToBankByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRtnRepealFromFutureToBankByFuture(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFuture((*CThostFtdcRspRepealField)(pRspRepeal))
}

// goTraderOnRspFromBankToFutureByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRspFromBankToFutureByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromBankToFutureByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspFromFutureToBankByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRspFromFutureToBankByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromFutureToBankByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQueryBankAccountMoneyByFuture 回调函数实现（C 调用约定版本）
func goTraderOnRspQueryBankAccountMoneyByFuture(userData uintptr, pReqQueryAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryBankAccountMoneyByFuture((*CThostFtdcReqQueryAccountField)(pReqQueryAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRtnOpenAccountByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnOpenAccountByBank(userData uintptr, pOpenAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOpenAccountByBank((*CThostFtdcOpenAccountField)(pOpenAccount))
}

// goTraderOnRtnCancelAccountByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnCancelAccountByBank(userData uintptr, pCancelAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCancelAccountByBank((*CThostFtdcCancelAccountField)(pCancelAccount))
}

// goTraderOnRtnChangeAccountByBank 回调函数实现（C 调用约定版本）
func goTraderOnRtnChangeAccountByBank(userData uintptr, pChangeAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnChangeAccountByBank((*CThostFtdcChangeAccountField)(pChangeAccount))
}

// goTraderOnRspQryClassifiedInstrument 回调函数实现（C 调用约定版本）
func goTraderOnRspQryClassifiedInstrument(userData uintptr, pInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryClassifiedInstrument((*CThostFtdcInstrumentField)(pInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryCombPromotionParam 回调函数实现（C 调用约定版本）
func goTraderOnRspQryCombPromotionParam(userData uintptr, pCombPromotionParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombPromotionParam((*CThostFtdcCombPromotionParamField)(pCombPromotionParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRiskSettleInvstPosition 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRiskSettleInvstPosition(userData uintptr, pRiskSettleInvstPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleInvstPosition((*CThostFtdcRiskSettleInvstPositionField)(pRiskSettleInvstPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRiskSettleProductStatus 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRiskSettleProductStatus(userData uintptr, pRiskSettleProductStatus unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleProductStatus((*CThostFtdcRiskSettleProductStatusField)(pRiskSettleProductStatus), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMFutureParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMFutureParameter(userData uintptr, pSPBMFutureParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMFutureParameter((*CThostFtdcSPBMFutureParameterField)(pSPBMFutureParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMOptionParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMOptionParameter(userData uintptr, pSPBMOptionParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMOptionParameter((*CThostFtdcSPBMOptionParameterField)(pSPBMOptionParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMIntraParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMIntraParameter(userData uintptr, pSPBMIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMIntraParameter((*CThostFtdcSPBMIntraParameterField)(pSPBMIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMInterParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMInterParameter(userData uintptr, pSPBMInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInterParameter((*CThostFtdcSPBMInterParameterField)(pSPBMInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMPortfDefinition 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMPortfDefinition(userData uintptr, pSPBMPortfDefinition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMPortfDefinition((*CThostFtdcSPBMPortfDefinitionField)(pSPBMPortfDefinition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMInvestorPortfDef 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMInvestorPortfDef(userData uintptr, pSPBMInvestorPortfDef unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInvestorPortfDef((*CThostFtdcSPBMInvestorPortfDefField)(pSPBMInvestorPortfDef), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorPortfMarginRatio 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorPortfMarginRatio(userData uintptr, pInvestorPortfMarginRatio unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfMarginRatio((*CThostFtdcInvestorPortfMarginRatioField)(pInvestorPortfMarginRatio), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorProdSPBMDetail 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorProdSPBMDetail(userData uintptr, pInvestorProdSPBMDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdSPBMDetail((*CThostFtdcInvestorProdSPBMDetailField)(pInvestorProdSPBMDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorCommoditySPMMMargin 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorCommoditySPMMMargin(userData uintptr, pInvestorCommoditySPMMMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommoditySPMMMargin((*CThostFtdcInvestorCommoditySPMMMarginField)(pInvestorCommoditySPMMMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorCommodityGroupSPMMMargin 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorCommodityGroupSPMMMargin(userData uintptr, pInvestorCommodityGroupSPMMMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommodityGroupSPMMMargin((*CThostFtdcInvestorCommodityGroupSPMMMarginField)(pInvestorCommodityGroupSPMMMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPMMInstParam 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPMMInstParam(userData uintptr, pSPMMInstParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMInstParam((*CThostFtdcSPMMInstParamField)(pSPMMInstParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPMMProductParam 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPMMProductParam(userData uintptr, pSPMMProductParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMProductParam((*CThostFtdcSPMMProductParamField)(pSPMMProductParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQrySPBMAddOnInterParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQrySPBMAddOnInterParameter(userData uintptr, pSPBMAddOnInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMAddOnInterParameter((*CThostFtdcSPBMAddOnInterParameterField)(pSPBMAddOnInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSCombProductInfo 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSCombProductInfo(userData uintptr, pRCAMSCombProductInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSCombProductInfo((*CThostFtdcRCAMSCombProductInfoField)(pRCAMSCombProductInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSInstrParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSInstrParameter(userData uintptr, pRCAMSInstrParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInstrParameter((*CThostFtdcRCAMSInstrParameterField)(pRCAMSInstrParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSIntraParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSIntraParameter(userData uintptr, pRCAMSIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSIntraParameter((*CThostFtdcRCAMSIntraParameterField)(pRCAMSIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSInterParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSInterParameter(userData uintptr, pRCAMSInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInterParameter((*CThostFtdcRCAMSInterParameterField)(pRCAMSInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSShortOptAdjustParam 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSShortOptAdjustParam(userData uintptr, pRCAMSShortOptAdjustParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSShortOptAdjustParam((*CThostFtdcRCAMSShortOptAdjustParamField)(pRCAMSShortOptAdjustParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRCAMSInvestorCombPosition 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRCAMSInvestorCombPosition(userData uintptr, pRCAMSInvestorCombPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInvestorCombPosition((*CThostFtdcRCAMSInvestorCombPositionField)(pRCAMSInvestorCombPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorProdRCAMSMargin 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorProdRCAMSMargin(userData uintptr, pInvestorProdRCAMSMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRCAMSMargin((*CThostFtdcInvestorProdRCAMSMarginField)(pInvestorProdRCAMSMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRULEInstrParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRULEInstrParameter(userData uintptr, pRULEInstrParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInstrParameter((*CThostFtdcRULEInstrParameterField)(pRULEInstrParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRULEIntraParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRULEIntraParameter(userData uintptr, pRULEIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEIntraParameter((*CThostFtdcRULEIntraParameterField)(pRULEIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryRULEInterParameter 回调函数实现（C 调用约定版本）
func goTraderOnRspQryRULEInterParameter(userData uintptr, pRULEInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInterParameter((*CThostFtdcRULEInterParameterField)(pRULEInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorProdRULEMargin 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorProdRULEMargin(userData uintptr, pInvestorProdRULEMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRULEMargin((*CThostFtdcInvestorProdRULEMarginField)(pInvestorProdRULEMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// goTraderOnRspQryInvestorPortfSetting 回调函数实现（C 调用约定版本）
func goTraderOnRspQryInvestorPortfSetting(userData uintptr, pInvestorPortfSetting unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfSetting((*CThostFtdcInvestorPortfSettingField)(pInvestorPortfSetting), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

// ========== 辅助函数：使用 purego.NewCallback 获取 C 函数指针 ==========
// 注意：purego.NewCallback 不支持 unsafe.Pointer 参数，需要用具体指针类型的 wrapper

// GetGoTraderOnFrontConnected 获取 goTraderOnFrontConnected 的 C 函数指针
func GetGoTraderOnFrontConnected() uintptr {
	return purego.NewCallback(goTraderOnFrontConnected)
}

// GetGoTraderOnFrontDisconnected 获取 goTraderOnFrontDisconnected 的 C 函数指针
func GetGoTraderOnFrontDisconnected() uintptr {
	return purego.NewCallback(goTraderOnFrontDisconnected)
}

// GetGoTraderOnHeartBeatWarning 获取 goTraderOnHeartBeatWarning 的 C 函数指针
func GetGoTraderOnHeartBeatWarning() uintptr {
	return purego.NewCallback(goTraderOnHeartBeatWarning)
}

// GetGoTraderOnRspAuthenticate 获取 goTraderOnRspAuthenticate 的 C 函数指针
func GetGoTraderOnRspAuthenticate() uintptr {
	wrapper := func(userData uintptr, pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspAuthenticate(userData, unsafe.Pointer(pRspAuthenticateField), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspUserLogin 获取 goTraderOnRspUserLogin 的 C 函数指针
func GetGoTraderOnRspUserLogin() uintptr {
	wrapper := func(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspUserLogin(userData, unsafe.Pointer(pRspUserLogin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspUserLogout 获取 goTraderOnRspUserLogout 的 C 函数指针
func GetGoTraderOnRspUserLogout() uintptr {
	wrapper := func(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspUserLogout(userData, unsafe.Pointer(pUserLogout), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspUserPasswordUpdate 获取 goTraderOnRspUserPasswordUpdate 的 C 函数指针
func GetGoTraderOnRspUserPasswordUpdate() uintptr {
	wrapper := func(userData uintptr, pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspUserPasswordUpdate(userData, unsafe.Pointer(pUserPasswordUpdate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspTradingAccountPasswordUpdate 获取 goTraderOnRspTradingAccountPasswordUpdate 的 C 函数指针
func GetGoTraderOnRspTradingAccountPasswordUpdate() uintptr {
	wrapper := func(userData uintptr, pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspTradingAccountPasswordUpdate(userData, unsafe.Pointer(pTradingAccountPasswordUpdate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspUserAuthMethod 获取 goTraderOnRspUserAuthMethod 的 C 函数指针
func GetGoTraderOnRspUserAuthMethod() uintptr {
	wrapper := func(userData uintptr, pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspUserAuthMethod(userData, unsafe.Pointer(pRspUserAuthMethod), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspGenUserCaptcha 获取 goTraderOnRspGenUserCaptcha 的 C 函数指针
func GetGoTraderOnRspGenUserCaptcha() uintptr {
	wrapper := func(userData uintptr, pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspGenUserCaptcha(userData, unsafe.Pointer(pRspGenUserCaptcha), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspGenUserText 获取 goTraderOnRspGenUserText 的 C 函数指针
func GetGoTraderOnRspGenUserText() uintptr {
	wrapper := func(userData uintptr, pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspGenUserText(userData, unsafe.Pointer(pRspGenUserText), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspOrderInsert 获取 goTraderOnRspOrderInsert 的 C 函数指针
func GetGoTraderOnRspOrderInsert() uintptr {
	wrapper := func(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspOrderInsert(userData, unsafe.Pointer(pInputOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspParkedOrderInsert 获取 goTraderOnRspParkedOrderInsert 的 C 函数指针
func GetGoTraderOnRspParkedOrderInsert() uintptr {
	wrapper := func(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspParkedOrderInsert(userData, unsafe.Pointer(pParkedOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspParkedOrderAction 获取 goTraderOnRspParkedOrderAction 的 C 函数指针
func GetGoTraderOnRspParkedOrderAction() uintptr {
	wrapper := func(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspParkedOrderAction(userData, unsafe.Pointer(pParkedOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspOrderAction 获取 goTraderOnRspOrderAction 的 C 函数指针
func GetGoTraderOnRspOrderAction() uintptr {
	wrapper := func(userData uintptr, pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspOrderAction(userData, unsafe.Pointer(pInputOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryMaxOrderVolume 获取 goTraderOnRspQryMaxOrderVolume 的 C 函数指针
func GetGoTraderOnRspQryMaxOrderVolume() uintptr {
	wrapper := func(userData uintptr, pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryMaxOrderVolume(userData, unsafe.Pointer(pQryMaxOrderVolume), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspSettlementInfoConfirm 获取 goTraderOnRspSettlementInfoConfirm 的 C 函数指针
func GetGoTraderOnRspSettlementInfoConfirm() uintptr {
	wrapper := func(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspSettlementInfoConfirm(userData, unsafe.Pointer(pSettlementInfoConfirm), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspRemoveParkedOrder 获取 goTraderOnRspRemoveParkedOrder 的 C 函数指针
func GetGoTraderOnRspRemoveParkedOrder() uintptr {
	wrapper := func(userData uintptr, pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspRemoveParkedOrder(userData, unsafe.Pointer(pRemoveParkedOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspRemoveParkedOrderAction 获取 goTraderOnRspRemoveParkedOrderAction 的 C 函数指针
func GetGoTraderOnRspRemoveParkedOrderAction() uintptr {
	wrapper := func(userData uintptr, pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspRemoveParkedOrderAction(userData, unsafe.Pointer(pRemoveParkedOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspExecOrderInsert 获取 goTraderOnRspExecOrderInsert 的 C 函数指针
func GetGoTraderOnRspExecOrderInsert() uintptr {
	wrapper := func(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspExecOrderInsert(userData, unsafe.Pointer(pInputExecOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspExecOrderAction 获取 goTraderOnRspExecOrderAction 的 C 函数指针
func GetGoTraderOnRspExecOrderAction() uintptr {
	wrapper := func(userData uintptr, pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspExecOrderAction(userData, unsafe.Pointer(pInputExecOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspForQuoteInsert 获取 goTraderOnRspForQuoteInsert 的 C 函数指针
func GetGoTraderOnRspForQuoteInsert() uintptr {
	wrapper := func(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspForQuoteInsert(userData, unsafe.Pointer(pInputForQuote), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQuoteInsert 获取 goTraderOnRspQuoteInsert 的 C 函数指针
func GetGoTraderOnRspQuoteInsert() uintptr {
	wrapper := func(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQuoteInsert(userData, unsafe.Pointer(pInputQuote), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQuoteAction 获取 goTraderOnRspQuoteAction 的 C 函数指针
func GetGoTraderOnRspQuoteAction() uintptr {
	wrapper := func(userData uintptr, pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQuoteAction(userData, unsafe.Pointer(pInputQuoteAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspBatchOrderAction 获取 goTraderOnRspBatchOrderAction 的 C 函数指针
func GetGoTraderOnRspBatchOrderAction() uintptr {
	wrapper := func(userData uintptr, pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspBatchOrderAction(userData, unsafe.Pointer(pInputBatchOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspOptionSelfCloseInsert 获取 goTraderOnRspOptionSelfCloseInsert 的 C 函数指针
func GetGoTraderOnRspOptionSelfCloseInsert() uintptr {
	wrapper := func(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspOptionSelfCloseInsert(userData, unsafe.Pointer(pInputOptionSelfClose), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspOptionSelfCloseAction 获取 goTraderOnRspOptionSelfCloseAction 的 C 函数指针
func GetGoTraderOnRspOptionSelfCloseAction() uintptr {
	wrapper := func(userData uintptr, pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspOptionSelfCloseAction(userData, unsafe.Pointer(pInputOptionSelfCloseAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspCombActionInsert 获取 goTraderOnRspCombActionInsert 的 C 函数指针
func GetGoTraderOnRspCombActionInsert() uintptr {
	wrapper := func(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspCombActionInsert(userData, unsafe.Pointer(pInputCombAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryOrder 获取 goTraderOnRspQryOrder 的 C 函数指针
func GetGoTraderOnRspQryOrder() uintptr {
	wrapper := func(userData uintptr, pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryOrder(userData, unsafe.Pointer(pOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTrade 获取 goTraderOnRspQryTrade 的 C 函数指针
func GetGoTraderOnRspQryTrade() uintptr {
	wrapper := func(userData uintptr, pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTrade(userData, unsafe.Pointer(pTrade), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorPosition 获取 goTraderOnRspQryInvestorPosition 的 C 函数指针
func GetGoTraderOnRspQryInvestorPosition() uintptr {
	wrapper := func(userData uintptr, pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorPosition(userData, unsafe.Pointer(pInvestorPosition), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTradingAccount 获取 goTraderOnRspQryTradingAccount 的 C 函数指针
func GetGoTraderOnRspQryTradingAccount() uintptr {
	wrapper := func(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTradingAccount(userData, unsafe.Pointer(pTradingAccount), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestor 获取 goTraderOnRspQryInvestor 的 C 函数指针
func GetGoTraderOnRspQryInvestor() uintptr {
	wrapper := func(userData uintptr, pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestor(userData, unsafe.Pointer(pInvestor), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTradingCode 获取 goTraderOnRspQryTradingCode 的 C 函数指针
func GetGoTraderOnRspQryTradingCode() uintptr {
	wrapper := func(userData uintptr, pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTradingCode(userData, unsafe.Pointer(pTradingCode), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInstrumentMarginRate 获取 goTraderOnRspQryInstrumentMarginRate 的 C 函数指针
func GetGoTraderOnRspQryInstrumentMarginRate() uintptr {
	wrapper := func(userData uintptr, pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInstrumentMarginRate(userData, unsafe.Pointer(pInstrumentMarginRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInstrumentCommissionRate 获取 goTraderOnRspQryInstrumentCommissionRate 的 C 函数指针
func GetGoTraderOnRspQryInstrumentCommissionRate() uintptr {
	wrapper := func(userData uintptr, pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInstrumentCommissionRate(userData, unsafe.Pointer(pInstrumentCommissionRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryExchange 获取 goTraderOnRspQryExchange 的 C 函数指针
func GetGoTraderOnRspQryExchange() uintptr {
	wrapper := func(userData uintptr, pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryExchange(userData, unsafe.Pointer(pExchange), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryProduct 获取 goTraderOnRspQryProduct 的 C 函数指针
func GetGoTraderOnRspQryProduct() uintptr {
	wrapper := func(userData uintptr, pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryProduct(userData, unsafe.Pointer(pProduct), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInstrument 获取 goTraderOnRspQryInstrument 的 C 函数指针
func GetGoTraderOnRspQryInstrument() uintptr {
	wrapper := func(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInstrument(userData, unsafe.Pointer(pInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryDepthMarketData 获取 goTraderOnRspQryDepthMarketData 的 C 函数指针
func GetGoTraderOnRspQryDepthMarketData() uintptr {
	wrapper := func(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryDepthMarketData(userData, unsafe.Pointer(pDepthMarketData), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTraderOffer 获取 goTraderOnRspQryTraderOffer 的 C 函数指针
func GetGoTraderOnRspQryTraderOffer() uintptr {
	wrapper := func(userData uintptr, pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTraderOffer(userData, unsafe.Pointer(pTraderOffer), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySettlementInfo 获取 goTraderOnRspQrySettlementInfo 的 C 函数指针
func GetGoTraderOnRspQrySettlementInfo() uintptr {
	wrapper := func(userData uintptr, pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySettlementInfo(userData, unsafe.Pointer(pSettlementInfo), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTransferBank 获取 goTraderOnRspQryTransferBank 的 C 函数指针
func GetGoTraderOnRspQryTransferBank() uintptr {
	wrapper := func(userData uintptr, pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTransferBank(userData, unsafe.Pointer(pTransferBank), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorPositionDetail 获取 goTraderOnRspQryInvestorPositionDetail 的 C 函数指针
func GetGoTraderOnRspQryInvestorPositionDetail() uintptr {
	wrapper := func(userData uintptr, pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorPositionDetail(userData, unsafe.Pointer(pInvestorPositionDetail), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryNotice 获取 goTraderOnRspQryNotice 的 C 函数指针
func GetGoTraderOnRspQryNotice() uintptr {
	wrapper := func(userData uintptr, pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryNotice(userData, unsafe.Pointer(pNotice), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySettlementInfoConfirm 获取 goTraderOnRspQrySettlementInfoConfirm 的 C 函数指针
func GetGoTraderOnRspQrySettlementInfoConfirm() uintptr {
	wrapper := func(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySettlementInfoConfirm(userData, unsafe.Pointer(pSettlementInfoConfirm), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorPositionCombineDetail 获取 goTraderOnRspQryInvestorPositionCombineDetail 的 C 函数指针
func GetGoTraderOnRspQryInvestorPositionCombineDetail() uintptr {
	wrapper := func(userData uintptr, pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorPositionCombineDetail(userData, unsafe.Pointer(pInvestorPositionCombineDetail), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryCFMMCTradingAccountKey 获取 goTraderOnRspQryCFMMCTradingAccountKey 的 C 函数指针
func GetGoTraderOnRspQryCFMMCTradingAccountKey() uintptr {
	wrapper := func(userData uintptr, pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryCFMMCTradingAccountKey(userData, unsafe.Pointer(pCFMMCTradingAccountKey), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryEWarrantOffset 获取 goTraderOnRspQryEWarrantOffset 的 C 函数指针
func GetGoTraderOnRspQryEWarrantOffset() uintptr {
	wrapper := func(userData uintptr, pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryEWarrantOffset(userData, unsafe.Pointer(pEWarrantOffset), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorProductGroupMargin 获取 goTraderOnRspQryInvestorProductGroupMargin 的 C 函数指针
func GetGoTraderOnRspQryInvestorProductGroupMargin() uintptr {
	wrapper := func(userData uintptr, pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorProductGroupMargin(userData, unsafe.Pointer(pInvestorProductGroupMargin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryExchangeMarginRate 获取 goTraderOnRspQryExchangeMarginRate 的 C 函数指针
func GetGoTraderOnRspQryExchangeMarginRate() uintptr {
	wrapper := func(userData uintptr, pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryExchangeMarginRate(userData, unsafe.Pointer(pExchangeMarginRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryExchangeMarginRateAdjust 获取 goTraderOnRspQryExchangeMarginRateAdjust 的 C 函数指针
func GetGoTraderOnRspQryExchangeMarginRateAdjust() uintptr {
	wrapper := func(userData uintptr, pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryExchangeMarginRateAdjust(userData, unsafe.Pointer(pExchangeMarginRateAdjust), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryExchangeRate 获取 goTraderOnRspQryExchangeRate 的 C 函数指针
func GetGoTraderOnRspQryExchangeRate() uintptr {
	wrapper := func(userData uintptr, pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryExchangeRate(userData, unsafe.Pointer(pExchangeRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySecAgentACIDMap 获取 goTraderOnRspQrySecAgentACIDMap 的 C 函数指针
func GetGoTraderOnRspQrySecAgentACIDMap() uintptr {
	wrapper := func(userData uintptr, pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySecAgentACIDMap(userData, unsafe.Pointer(pSecAgentACIDMap), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryProductExchRate 获取 goTraderOnRspQryProductExchRate 的 C 函数指针
func GetGoTraderOnRspQryProductExchRate() uintptr {
	wrapper := func(userData uintptr, pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryProductExchRate(userData, unsafe.Pointer(pProductExchRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryProductGroup 获取 goTraderOnRspQryProductGroup 的 C 函数指针
func GetGoTraderOnRspQryProductGroup() uintptr {
	wrapper := func(userData uintptr, pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryProductGroup(userData, unsafe.Pointer(pProductGroup), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryMMInstrumentCommissionRate 获取 goTraderOnRspQryMMInstrumentCommissionRate 的 C 函数指针
func GetGoTraderOnRspQryMMInstrumentCommissionRate() uintptr {
	wrapper := func(userData uintptr, pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryMMInstrumentCommissionRate(userData, unsafe.Pointer(pMMInstrumentCommissionRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryMMOptionInstrCommRate 获取 goTraderOnRspQryMMOptionInstrCommRate 的 C 函数指针
func GetGoTraderOnRspQryMMOptionInstrCommRate() uintptr {
	wrapper := func(userData uintptr, pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryMMOptionInstrCommRate(userData, unsafe.Pointer(pMMOptionInstrCommRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInstrumentOrderCommRate 获取 goTraderOnRspQryInstrumentOrderCommRate 的 C 函数指针
func GetGoTraderOnRspQryInstrumentOrderCommRate() uintptr {
	wrapper := func(userData uintptr, pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInstrumentOrderCommRate(userData, unsafe.Pointer(pInstrumentOrderCommRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySecAgentTradingAccount 获取 goTraderOnRspQrySecAgentTradingAccount 的 C 函数指针
func GetGoTraderOnRspQrySecAgentTradingAccount() uintptr {
	wrapper := func(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySecAgentTradingAccount(userData, unsafe.Pointer(pTradingAccount), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySecAgentCheckMode 获取 goTraderOnRspQrySecAgentCheckMode 的 C 函数指针
func GetGoTraderOnRspQrySecAgentCheckMode() uintptr {
	wrapper := func(userData uintptr, pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySecAgentCheckMode(userData, unsafe.Pointer(pSecAgentCheckMode), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySecAgentTradeInfo 获取 goTraderOnRspQrySecAgentTradeInfo 的 C 函数指针
func GetGoTraderOnRspQrySecAgentTradeInfo() uintptr {
	wrapper := func(userData uintptr, pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySecAgentTradeInfo(userData, unsafe.Pointer(pSecAgentTradeInfo), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryOptionInstrTradeCost 获取 goTraderOnRspQryOptionInstrTradeCost 的 C 函数指针
func GetGoTraderOnRspQryOptionInstrTradeCost() uintptr {
	wrapper := func(userData uintptr, pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryOptionInstrTradeCost(userData, unsafe.Pointer(pOptionInstrTradeCost), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryOptionInstrCommRate 获取 goTraderOnRspQryOptionInstrCommRate 的 C 函数指针
func GetGoTraderOnRspQryOptionInstrCommRate() uintptr {
	wrapper := func(userData uintptr, pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryOptionInstrCommRate(userData, unsafe.Pointer(pOptionInstrCommRate), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryExecOrder 获取 goTraderOnRspQryExecOrder 的 C 函数指针
func GetGoTraderOnRspQryExecOrder() uintptr {
	wrapper := func(userData uintptr, pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryExecOrder(userData, unsafe.Pointer(pExecOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryForQuote 获取 goTraderOnRspQryForQuote 的 C 函数指针
func GetGoTraderOnRspQryForQuote() uintptr {
	wrapper := func(userData uintptr, pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryForQuote(userData, unsafe.Pointer(pForQuote), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryQuote 获取 goTraderOnRspQryQuote 的 C 函数指针
func GetGoTraderOnRspQryQuote() uintptr {
	wrapper := func(userData uintptr, pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryQuote(userData, unsafe.Pointer(pQuote), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryOptionSelfClose 获取 goTraderOnRspQryOptionSelfClose 的 C 函数指针
func GetGoTraderOnRspQryOptionSelfClose() uintptr {
	wrapper := func(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryOptionSelfClose(userData, unsafe.Pointer(pOptionSelfClose), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestUnit 获取 goTraderOnRspQryInvestUnit 的 C 函数指针
func GetGoTraderOnRspQryInvestUnit() uintptr {
	wrapper := func(userData uintptr, pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestUnit(userData, unsafe.Pointer(pInvestUnit), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryCombInstrumentGuard 获取 goTraderOnRspQryCombInstrumentGuard 的 C 函数指针
func GetGoTraderOnRspQryCombInstrumentGuard() uintptr {
	wrapper := func(userData uintptr, pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryCombInstrumentGuard(userData, unsafe.Pointer(pCombInstrumentGuard), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryCombAction 获取 goTraderOnRspQryCombAction 的 C 函数指针
func GetGoTraderOnRspQryCombAction() uintptr {
	wrapper := func(userData uintptr, pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryCombAction(userData, unsafe.Pointer(pCombAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTransferSerial 获取 goTraderOnRspQryTransferSerial 的 C 函数指针
func GetGoTraderOnRspQryTransferSerial() uintptr {
	wrapper := func(userData uintptr, pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTransferSerial(userData, unsafe.Pointer(pTransferSerial), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryAccountregister 获取 goTraderOnRspQryAccountregister 的 C 函数指针
func GetGoTraderOnRspQryAccountregister() uintptr {
	wrapper := func(userData uintptr, pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryAccountregister(userData, unsafe.Pointer(pAccountregister), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspError 获取 goTraderOnRspError 的 C 函数指针
func GetGoTraderOnRspError() uintptr {
	wrapper := func(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspError(userData, unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnOrder 获取 goTraderOnRtnOrder 的 C 函数指针
func GetGoTraderOnRtnOrder() uintptr {
	wrapper := func(userData uintptr, pOrder *CThostFtdcOrderField) {
		goTraderOnRtnOrder(userData, unsafe.Pointer(pOrder))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnTrade 获取 goTraderOnRtnTrade 的 C 函数指针
func GetGoTraderOnRtnTrade() uintptr {
	wrapper := func(userData uintptr, pTrade *CThostFtdcTradeField) {
		goTraderOnRtnTrade(userData, unsafe.Pointer(pTrade))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnOrderInsert 获取 goTraderOnErrRtnOrderInsert 的 C 函数指针
func GetGoTraderOnErrRtnOrderInsert() uintptr {
	wrapper := func(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnOrderInsert(userData, unsafe.Pointer(pInputOrder), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnOrderAction 获取 goTraderOnErrRtnOrderAction 的 C 函数指针
func GetGoTraderOnErrRtnOrderAction() uintptr {
	wrapper := func(userData uintptr, pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnOrderAction(userData, unsafe.Pointer(pOrderAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnInstrumentStatus 获取 goTraderOnRtnInstrumentStatus 的 C 函数指针
func GetGoTraderOnRtnInstrumentStatus() uintptr {
	wrapper := func(userData uintptr, pInstrumentStatus *CThostFtdcInstrumentStatusField) {
		goTraderOnRtnInstrumentStatus(userData, unsafe.Pointer(pInstrumentStatus))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnBulletin 获取 goTraderOnRtnBulletin 的 C 函数指针
func GetGoTraderOnRtnBulletin() uintptr {
	wrapper := func(userData uintptr, pBulletin *CThostFtdcBulletinField) {
		goTraderOnRtnBulletin(userData, unsafe.Pointer(pBulletin))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnTradingNotice 获取 goTraderOnRtnTradingNotice 的 C 函数指针
func GetGoTraderOnRtnTradingNotice() uintptr {
	wrapper := func(userData uintptr, pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField) {
		goTraderOnRtnTradingNotice(userData, unsafe.Pointer(pTradingNoticeInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnErrorConditionalOrder 获取 goTraderOnRtnErrorConditionalOrder 的 C 函数指针
func GetGoTraderOnRtnErrorConditionalOrder() uintptr {
	wrapper := func(userData uintptr, pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField) {
		goTraderOnRtnErrorConditionalOrder(userData, unsafe.Pointer(pErrorConditionalOrder))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnExecOrder 获取 goTraderOnRtnExecOrder 的 C 函数指针
func GetGoTraderOnRtnExecOrder() uintptr {
	wrapper := func(userData uintptr, pExecOrder *CThostFtdcExecOrderField) {
		goTraderOnRtnExecOrder(userData, unsafe.Pointer(pExecOrder))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnExecOrderInsert 获取 goTraderOnErrRtnExecOrderInsert 的 C 函数指针
func GetGoTraderOnErrRtnExecOrderInsert() uintptr {
	wrapper := func(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnExecOrderInsert(userData, unsafe.Pointer(pInputExecOrder), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnExecOrderAction 获取 goTraderOnErrRtnExecOrderAction 的 C 函数指针
func GetGoTraderOnErrRtnExecOrderAction() uintptr {
	wrapper := func(userData uintptr, pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnExecOrderAction(userData, unsafe.Pointer(pExecOrderAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnForQuoteInsert 获取 goTraderOnErrRtnForQuoteInsert 的 C 函数指针
func GetGoTraderOnErrRtnForQuoteInsert() uintptr {
	wrapper := func(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnForQuoteInsert(userData, unsafe.Pointer(pInputForQuote), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnQuote 获取 goTraderOnRtnQuote 的 C 函数指针
func GetGoTraderOnRtnQuote() uintptr {
	wrapper := func(userData uintptr, pQuote *CThostFtdcQuoteField) {
		goTraderOnRtnQuote(userData, unsafe.Pointer(pQuote))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnQuoteInsert 获取 goTraderOnErrRtnQuoteInsert 的 C 函数指针
func GetGoTraderOnErrRtnQuoteInsert() uintptr {
	wrapper := func(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnQuoteInsert(userData, unsafe.Pointer(pInputQuote), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnQuoteAction 获取 goTraderOnErrRtnQuoteAction 的 C 函数指针
func GetGoTraderOnErrRtnQuoteAction() uintptr {
	wrapper := func(userData uintptr, pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnQuoteAction(userData, unsafe.Pointer(pQuoteAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnForQuoteRsp 获取 goTraderOnRtnForQuoteRsp 的 C 函数指针
func GetGoTraderOnRtnForQuoteRsp() uintptr {
	wrapper := func(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField) {
		goTraderOnRtnForQuoteRsp(userData, unsafe.Pointer(pForQuoteRsp))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnCFMMCTradingAccountToken 获取 goTraderOnRtnCFMMCTradingAccountToken 的 C 函数指针
func GetGoTraderOnRtnCFMMCTradingAccountToken() uintptr {
	wrapper := func(userData uintptr, pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField) {
		goTraderOnRtnCFMMCTradingAccountToken(userData, unsafe.Pointer(pCFMMCTradingAccountToken))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnBatchOrderAction 获取 goTraderOnErrRtnBatchOrderAction 的 C 函数指针
func GetGoTraderOnErrRtnBatchOrderAction() uintptr {
	wrapper := func(userData uintptr, pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnBatchOrderAction(userData, unsafe.Pointer(pBatchOrderAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnOptionSelfClose 获取 goTraderOnRtnOptionSelfClose 的 C 函数指针
func GetGoTraderOnRtnOptionSelfClose() uintptr {
	wrapper := func(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField) {
		goTraderOnRtnOptionSelfClose(userData, unsafe.Pointer(pOptionSelfClose))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnOptionSelfCloseInsert 获取 goTraderOnErrRtnOptionSelfCloseInsert 的 C 函数指针
func GetGoTraderOnErrRtnOptionSelfCloseInsert() uintptr {
	wrapper := func(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnOptionSelfCloseInsert(userData, unsafe.Pointer(pInputOptionSelfClose), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnOptionSelfCloseAction 获取 goTraderOnErrRtnOptionSelfCloseAction 的 C 函数指针
func GetGoTraderOnErrRtnOptionSelfCloseAction() uintptr {
	wrapper := func(userData uintptr, pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnOptionSelfCloseAction(userData, unsafe.Pointer(pOptionSelfCloseAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnCombAction 获取 goTraderOnRtnCombAction 的 C 函数指针
func GetGoTraderOnRtnCombAction() uintptr {
	wrapper := func(userData uintptr, pCombAction *CThostFtdcCombActionField) {
		goTraderOnRtnCombAction(userData, unsafe.Pointer(pCombAction))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnCombActionInsert 获取 goTraderOnErrRtnCombActionInsert 的 C 函数指针
func GetGoTraderOnErrRtnCombActionInsert() uintptr {
	wrapper := func(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnCombActionInsert(userData, unsafe.Pointer(pInputCombAction), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryContractBank 获取 goTraderOnRspQryContractBank 的 C 函数指针
func GetGoTraderOnRspQryContractBank() uintptr {
	wrapper := func(userData uintptr, pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryContractBank(userData, unsafe.Pointer(pContractBank), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryParkedOrder 获取 goTraderOnRspQryParkedOrder 的 C 函数指针
func GetGoTraderOnRspQryParkedOrder() uintptr {
	wrapper := func(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryParkedOrder(userData, unsafe.Pointer(pParkedOrder), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryParkedOrderAction 获取 goTraderOnRspQryParkedOrderAction 的 C 函数指针
func GetGoTraderOnRspQryParkedOrderAction() uintptr {
	wrapper := func(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryParkedOrderAction(userData, unsafe.Pointer(pParkedOrderAction), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryTradingNotice 获取 goTraderOnRspQryTradingNotice 的 C 函数指针
func GetGoTraderOnRspQryTradingNotice() uintptr {
	wrapper := func(userData uintptr, pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryTradingNotice(userData, unsafe.Pointer(pTradingNotice), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryBrokerTradingParams 获取 goTraderOnRspQryBrokerTradingParams 的 C 函数指针
func GetGoTraderOnRspQryBrokerTradingParams() uintptr {
	wrapper := func(userData uintptr, pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryBrokerTradingParams(userData, unsafe.Pointer(pBrokerTradingParams), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryBrokerTradingAlgos 获取 goTraderOnRspQryBrokerTradingAlgos 的 C 函数指针
func GetGoTraderOnRspQryBrokerTradingAlgos() uintptr {
	wrapper := func(userData uintptr, pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryBrokerTradingAlgos(userData, unsafe.Pointer(pBrokerTradingAlgos), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQueryCFMMCTradingAccountToken 获取 goTraderOnRspQueryCFMMCTradingAccountToken 的 C 函数指针
func GetGoTraderOnRspQueryCFMMCTradingAccountToken() uintptr {
	wrapper := func(userData uintptr, pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQueryCFMMCTradingAccountToken(userData, unsafe.Pointer(pQueryCFMMCTradingAccountToken), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnFromBankToFutureByBank 获取 goTraderOnRtnFromBankToFutureByBank 的 C 函数指针
func GetGoTraderOnRtnFromBankToFutureByBank() uintptr {
	wrapper := func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
		goTraderOnRtnFromBankToFutureByBank(userData, unsafe.Pointer(pRspTransfer))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnFromFutureToBankByBank 获取 goTraderOnRtnFromFutureToBankByBank 的 C 函数指针
func GetGoTraderOnRtnFromFutureToBankByBank() uintptr {
	wrapper := func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
		goTraderOnRtnFromFutureToBankByBank(userData, unsafe.Pointer(pRspTransfer))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromBankToFutureByBank 获取 goTraderOnRtnRepealFromBankToFutureByBank 的 C 函数指针
func GetGoTraderOnRtnRepealFromBankToFutureByBank() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromBankToFutureByBank(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromFutureToBankByBank 获取 goTraderOnRtnRepealFromFutureToBankByBank 的 C 函数指针
func GetGoTraderOnRtnRepealFromFutureToBankByBank() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromFutureToBankByBank(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnFromBankToFutureByFuture 获取 goTraderOnRtnFromBankToFutureByFuture 的 C 函数指针
func GetGoTraderOnRtnFromBankToFutureByFuture() uintptr {
	wrapper := func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
		goTraderOnRtnFromBankToFutureByFuture(userData, unsafe.Pointer(pRspTransfer))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnFromFutureToBankByFuture 获取 goTraderOnRtnFromFutureToBankByFuture 的 C 函数指针
func GetGoTraderOnRtnFromFutureToBankByFuture() uintptr {
	wrapper := func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
		goTraderOnRtnFromFutureToBankByFuture(userData, unsafe.Pointer(pRspTransfer))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromBankToFutureByFutureManual 获取 goTraderOnRtnRepealFromBankToFutureByFutureManual 的 C 函数指针
func GetGoTraderOnRtnRepealFromBankToFutureByFutureManual() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromBankToFutureByFutureManual(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromFutureToBankByFutureManual 获取 goTraderOnRtnRepealFromFutureToBankByFutureManual 的 C 函数指针
func GetGoTraderOnRtnRepealFromFutureToBankByFutureManual() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromFutureToBankByFutureManual(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnQueryBankBalanceByFuture 获取 goTraderOnRtnQueryBankBalanceByFuture 的 C 函数指针
func GetGoTraderOnRtnQueryBankBalanceByFuture() uintptr {
	wrapper := func(userData uintptr, pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField) {
		goTraderOnRtnQueryBankBalanceByFuture(userData, unsafe.Pointer(pNotifyQueryAccount))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnBankToFutureByFuture 获取 goTraderOnErrRtnBankToFutureByFuture 的 C 函数指针
func GetGoTraderOnErrRtnBankToFutureByFuture() uintptr {
	wrapper := func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnBankToFutureByFuture(userData, unsafe.Pointer(pReqTransfer), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnFutureToBankByFuture 获取 goTraderOnErrRtnFutureToBankByFuture 的 C 函数指针
func GetGoTraderOnErrRtnFutureToBankByFuture() uintptr {
	wrapper := func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnFutureToBankByFuture(userData, unsafe.Pointer(pReqTransfer), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnRepealBankToFutureByFutureManual 获取 goTraderOnErrRtnRepealBankToFutureByFutureManual 的 C 函数指针
func GetGoTraderOnErrRtnRepealBankToFutureByFutureManual() uintptr {
	wrapper := func(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnRepealBankToFutureByFutureManual(userData, unsafe.Pointer(pReqRepeal), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnRepealFutureToBankByFutureManual 获取 goTraderOnErrRtnRepealFutureToBankByFutureManual 的 C 函数指针
func GetGoTraderOnErrRtnRepealFutureToBankByFutureManual() uintptr {
	wrapper := func(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnRepealFutureToBankByFutureManual(userData, unsafe.Pointer(pReqRepeal), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnErrRtnQueryBankBalanceByFuture 获取 goTraderOnErrRtnQueryBankBalanceByFuture 的 C 函数指针
func GetGoTraderOnErrRtnQueryBankBalanceByFuture() uintptr {
	wrapper := func(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField) {
		goTraderOnErrRtnQueryBankBalanceByFuture(userData, unsafe.Pointer(pReqQueryAccount), unsafe.Pointer(pRspInfo))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromBankToFutureByFuture 获取 goTraderOnRtnRepealFromBankToFutureByFuture 的 C 函数指针
func GetGoTraderOnRtnRepealFromBankToFutureByFuture() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromBankToFutureByFuture(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnRepealFromFutureToBankByFuture 获取 goTraderOnRtnRepealFromFutureToBankByFuture 的 C 函数指针
func GetGoTraderOnRtnRepealFromFutureToBankByFuture() uintptr {
	wrapper := func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
		goTraderOnRtnRepealFromFutureToBankByFuture(userData, unsafe.Pointer(pRspRepeal))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspFromBankToFutureByFuture 获取 goTraderOnRspFromBankToFutureByFuture 的 C 函数指针
func GetGoTraderOnRspFromBankToFutureByFuture() uintptr {
	wrapper := func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspFromBankToFutureByFuture(userData, unsafe.Pointer(pReqTransfer), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspFromFutureToBankByFuture 获取 goTraderOnRspFromFutureToBankByFuture 的 C 函数指针
func GetGoTraderOnRspFromFutureToBankByFuture() uintptr {
	wrapper := func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspFromFutureToBankByFuture(userData, unsafe.Pointer(pReqTransfer), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQueryBankAccountMoneyByFuture 获取 goTraderOnRspQueryBankAccountMoneyByFuture 的 C 函数指针
func GetGoTraderOnRspQueryBankAccountMoneyByFuture() uintptr {
	wrapper := func(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQueryBankAccountMoneyByFuture(userData, unsafe.Pointer(pReqQueryAccount), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnOpenAccountByBank 获取 goTraderOnRtnOpenAccountByBank 的 C 函数指针
func GetGoTraderOnRtnOpenAccountByBank() uintptr {
	wrapper := func(userData uintptr, pOpenAccount *CThostFtdcOpenAccountField) {
		goTraderOnRtnOpenAccountByBank(userData, unsafe.Pointer(pOpenAccount))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnCancelAccountByBank 获取 goTraderOnRtnCancelAccountByBank 的 C 函数指针
func GetGoTraderOnRtnCancelAccountByBank() uintptr {
	wrapper := func(userData uintptr, pCancelAccount *CThostFtdcCancelAccountField) {
		goTraderOnRtnCancelAccountByBank(userData, unsafe.Pointer(pCancelAccount))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRtnChangeAccountByBank 获取 goTraderOnRtnChangeAccountByBank 的 C 函数指针
func GetGoTraderOnRtnChangeAccountByBank() uintptr {
	wrapper := func(userData uintptr, pChangeAccount *CThostFtdcChangeAccountField) {
		goTraderOnRtnChangeAccountByBank(userData, unsafe.Pointer(pChangeAccount))
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryClassifiedInstrument 获取 goTraderOnRspQryClassifiedInstrument 的 C 函数指针
func GetGoTraderOnRspQryClassifiedInstrument() uintptr {
	wrapper := func(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryClassifiedInstrument(userData, unsafe.Pointer(pInstrument), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryCombPromotionParam 获取 goTraderOnRspQryCombPromotionParam 的 C 函数指针
func GetGoTraderOnRspQryCombPromotionParam() uintptr {
	wrapper := func(userData uintptr, pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryCombPromotionParam(userData, unsafe.Pointer(pCombPromotionParam), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRiskSettleInvstPosition 获取 goTraderOnRspQryRiskSettleInvstPosition 的 C 函数指针
func GetGoTraderOnRspQryRiskSettleInvstPosition() uintptr {
	wrapper := func(userData uintptr, pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRiskSettleInvstPosition(userData, unsafe.Pointer(pRiskSettleInvstPosition), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRiskSettleProductStatus 获取 goTraderOnRspQryRiskSettleProductStatus 的 C 函数指针
func GetGoTraderOnRspQryRiskSettleProductStatus() uintptr {
	wrapper := func(userData uintptr, pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRiskSettleProductStatus(userData, unsafe.Pointer(pRiskSettleProductStatus), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMFutureParameter 获取 goTraderOnRspQrySPBMFutureParameter 的 C 函数指针
func GetGoTraderOnRspQrySPBMFutureParameter() uintptr {
	wrapper := func(userData uintptr, pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMFutureParameter(userData, unsafe.Pointer(pSPBMFutureParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMOptionParameter 获取 goTraderOnRspQrySPBMOptionParameter 的 C 函数指针
func GetGoTraderOnRspQrySPBMOptionParameter() uintptr {
	wrapper := func(userData uintptr, pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMOptionParameter(userData, unsafe.Pointer(pSPBMOptionParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMIntraParameter 获取 goTraderOnRspQrySPBMIntraParameter 的 C 函数指针
func GetGoTraderOnRspQrySPBMIntraParameter() uintptr {
	wrapper := func(userData uintptr, pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMIntraParameter(userData, unsafe.Pointer(pSPBMIntraParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMInterParameter 获取 goTraderOnRspQrySPBMInterParameter 的 C 函数指针
func GetGoTraderOnRspQrySPBMInterParameter() uintptr {
	wrapper := func(userData uintptr, pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMInterParameter(userData, unsafe.Pointer(pSPBMInterParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMPortfDefinition 获取 goTraderOnRspQrySPBMPortfDefinition 的 C 函数指针
func GetGoTraderOnRspQrySPBMPortfDefinition() uintptr {
	wrapper := func(userData uintptr, pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMPortfDefinition(userData, unsafe.Pointer(pSPBMPortfDefinition), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMInvestorPortfDef 获取 goTraderOnRspQrySPBMInvestorPortfDef 的 C 函数指针
func GetGoTraderOnRspQrySPBMInvestorPortfDef() uintptr {
	wrapper := func(userData uintptr, pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMInvestorPortfDef(userData, unsafe.Pointer(pSPBMInvestorPortfDef), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorPortfMarginRatio 获取 goTraderOnRspQryInvestorPortfMarginRatio 的 C 函数指针
func GetGoTraderOnRspQryInvestorPortfMarginRatio() uintptr {
	wrapper := func(userData uintptr, pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorPortfMarginRatio(userData, unsafe.Pointer(pInvestorPortfMarginRatio), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorProdSPBMDetail 获取 goTraderOnRspQryInvestorProdSPBMDetail 的 C 函数指针
func GetGoTraderOnRspQryInvestorProdSPBMDetail() uintptr {
	wrapper := func(userData uintptr, pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorProdSPBMDetail(userData, unsafe.Pointer(pInvestorProdSPBMDetail), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorCommoditySPMMMargin 获取 goTraderOnRspQryInvestorCommoditySPMMMargin 的 C 函数指针
func GetGoTraderOnRspQryInvestorCommoditySPMMMargin() uintptr {
	wrapper := func(userData uintptr, pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorCommoditySPMMMargin(userData, unsafe.Pointer(pInvestorCommoditySPMMMargin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorCommodityGroupSPMMMargin 获取 goTraderOnRspQryInvestorCommodityGroupSPMMMargin 的 C 函数指针
func GetGoTraderOnRspQryInvestorCommodityGroupSPMMMargin() uintptr {
	wrapper := func(userData uintptr, pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorCommodityGroupSPMMMargin(userData, unsafe.Pointer(pInvestorCommodityGroupSPMMMargin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPMMInstParam 获取 goTraderOnRspQrySPMMInstParam 的 C 函数指针
func GetGoTraderOnRspQrySPMMInstParam() uintptr {
	wrapper := func(userData uintptr, pSPMMInstParam *CThostFtdcSPMMInstParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPMMInstParam(userData, unsafe.Pointer(pSPMMInstParam), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPMMProductParam 获取 goTraderOnRspQrySPMMProductParam 的 C 函数指针
func GetGoTraderOnRspQrySPMMProductParam() uintptr {
	wrapper := func(userData uintptr, pSPMMProductParam *CThostFtdcSPMMProductParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPMMProductParam(userData, unsafe.Pointer(pSPMMProductParam), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQrySPBMAddOnInterParameter 获取 goTraderOnRspQrySPBMAddOnInterParameter 的 C 函数指针
func GetGoTraderOnRspQrySPBMAddOnInterParameter() uintptr {
	wrapper := func(userData uintptr, pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQrySPBMAddOnInterParameter(userData, unsafe.Pointer(pSPBMAddOnInterParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSCombProductInfo 获取 goTraderOnRspQryRCAMSCombProductInfo 的 C 函数指针
func GetGoTraderOnRspQryRCAMSCombProductInfo() uintptr {
	wrapper := func(userData uintptr, pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSCombProductInfo(userData, unsafe.Pointer(pRCAMSCombProductInfo), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSInstrParameter 获取 goTraderOnRspQryRCAMSInstrParameter 的 C 函数指针
func GetGoTraderOnRspQryRCAMSInstrParameter() uintptr {
	wrapper := func(userData uintptr, pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSInstrParameter(userData, unsafe.Pointer(pRCAMSInstrParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSIntraParameter 获取 goTraderOnRspQryRCAMSIntraParameter 的 C 函数指针
func GetGoTraderOnRspQryRCAMSIntraParameter() uintptr {
	wrapper := func(userData uintptr, pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSIntraParameter(userData, unsafe.Pointer(pRCAMSIntraParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSInterParameter 获取 goTraderOnRspQryRCAMSInterParameter 的 C 函数指针
func GetGoTraderOnRspQryRCAMSInterParameter() uintptr {
	wrapper := func(userData uintptr, pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSInterParameter(userData, unsafe.Pointer(pRCAMSInterParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSShortOptAdjustParam 获取 goTraderOnRspQryRCAMSShortOptAdjustParam 的 C 函数指针
func GetGoTraderOnRspQryRCAMSShortOptAdjustParam() uintptr {
	wrapper := func(userData uintptr, pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSShortOptAdjustParam(userData, unsafe.Pointer(pRCAMSShortOptAdjustParam), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRCAMSInvestorCombPosition 获取 goTraderOnRspQryRCAMSInvestorCombPosition 的 C 函数指针
func GetGoTraderOnRspQryRCAMSInvestorCombPosition() uintptr {
	wrapper := func(userData uintptr, pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRCAMSInvestorCombPosition(userData, unsafe.Pointer(pRCAMSInvestorCombPosition), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorProdRCAMSMargin 获取 goTraderOnRspQryInvestorProdRCAMSMargin 的 C 函数指针
func GetGoTraderOnRspQryInvestorProdRCAMSMargin() uintptr {
	wrapper := func(userData uintptr, pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorProdRCAMSMargin(userData, unsafe.Pointer(pInvestorProdRCAMSMargin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRULEInstrParameter 获取 goTraderOnRspQryRULEInstrParameter 的 C 函数指针
func GetGoTraderOnRspQryRULEInstrParameter() uintptr {
	wrapper := func(userData uintptr, pRULEInstrParameter *CThostFtdcRULEInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRULEInstrParameter(userData, unsafe.Pointer(pRULEInstrParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRULEIntraParameter 获取 goTraderOnRspQryRULEIntraParameter 的 C 函数指针
func GetGoTraderOnRspQryRULEIntraParameter() uintptr {
	wrapper := func(userData uintptr, pRULEIntraParameter *CThostFtdcRULEIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRULEIntraParameter(userData, unsafe.Pointer(pRULEIntraParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryRULEInterParameter 获取 goTraderOnRspQryRULEInterParameter 的 C 函数指针
func GetGoTraderOnRspQryRULEInterParameter() uintptr {
	wrapper := func(userData uintptr, pRULEInterParameter *CThostFtdcRULEInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryRULEInterParameter(userData, unsafe.Pointer(pRULEInterParameter), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorProdRULEMargin 获取 goTraderOnRspQryInvestorProdRULEMargin 的 C 函数指针
func GetGoTraderOnRspQryInvestorProdRULEMargin() uintptr {
	wrapper := func(userData uintptr, pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorProdRULEMargin(userData, unsafe.Pointer(pInvestorProdRULEMargin), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}

// GetGoTraderOnRspQryInvestorPortfSetting 获取 goTraderOnRspQryInvestorPortfSetting 的 C 函数指针
func GetGoTraderOnRspQryInvestorPortfSetting() uintptr {
	wrapper := func(userData uintptr, pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
		goTraderOnRspQryInvestorPortfSetting(userData, unsafe.Pointer(pInvestorPortfSetting), unsafe.Pointer(pRspInfo), nRequestID, bIsLast)
	}
	return purego.NewCallback(wrapper)
}
