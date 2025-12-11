package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 交易回调实现

/*
#include <stdint.h>
*/
import "C"

import "unsafe"

// ========== 回调函数 ==========

//export goTraderOnFrontConnected
func goTraderOnFrontConnected(userData uintptr) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontConnected()
}

//export goTraderOnFrontDisconnected
func goTraderOnFrontDisconnected(userData uintptr, nReason int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontDisconnected(nReason)
}

//export goTraderOnHeartBeatWarning
func goTraderOnHeartBeatWarning(userData uintptr, nTimeLapse int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnHeartBeatWarning(nTimeLapse)
}

//export goTraderOnRspAuthenticate
func goTraderOnRspAuthenticate(userData uintptr, pRspAuthenticateField unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspAuthenticate((*CThostFtdcRspAuthenticateField)(pRspAuthenticateField), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspUserLogin
func goTraderOnRspUserLogin(userData uintptr, pRspUserLogin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogin((*CThostFtdcRspUserLoginField)(pRspUserLogin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspUserLogout
func goTraderOnRspUserLogout(userData uintptr, pUserLogout unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogout((*CThostFtdcUserLogoutField)(pUserLogout), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspUserPasswordUpdate
func goTraderOnRspUserPasswordUpdate(userData uintptr, pUserPasswordUpdate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserPasswordUpdate((*CThostFtdcUserPasswordUpdateField)(pUserPasswordUpdate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspTradingAccountPasswordUpdate
func goTraderOnRspTradingAccountPasswordUpdate(userData uintptr, pTradingAccountPasswordUpdate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspTradingAccountPasswordUpdate((*CThostFtdcTradingAccountPasswordUpdateField)(pTradingAccountPasswordUpdate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspUserAuthMethod
func goTraderOnRspUserAuthMethod(userData uintptr, pRspUserAuthMethod unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserAuthMethod((*CThostFtdcRspUserAuthMethodField)(pRspUserAuthMethod), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspGenUserCaptcha
func goTraderOnRspGenUserCaptcha(userData uintptr, pRspGenUserCaptcha unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserCaptcha((*CThostFtdcRspGenUserCaptchaField)(pRspGenUserCaptcha), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspGenUserText
func goTraderOnRspGenUserText(userData uintptr, pRspGenUserText unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserText((*CThostFtdcRspGenUserTextField)(pRspGenUserText), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspOrderInsert
func goTraderOnRspOrderInsert(userData uintptr, pInputOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderInsert((*CThostFtdcInputOrderField)(pInputOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspParkedOrderInsert
func goTraderOnRspParkedOrderInsert(userData uintptr, pParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderInsert((*CThostFtdcParkedOrderField)(pParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspParkedOrderAction
func goTraderOnRspParkedOrderAction(userData uintptr, pParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderAction((*CThostFtdcParkedOrderActionField)(pParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspOrderAction
func goTraderOnRspOrderAction(userData uintptr, pInputOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderAction((*CThostFtdcInputOrderActionField)(pInputOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryMaxOrderVolume
func goTraderOnRspQryMaxOrderVolume(userData uintptr, pQryMaxOrderVolume unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMaxOrderVolume((*CThostFtdcQryMaxOrderVolumeField)(pQryMaxOrderVolume), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspSettlementInfoConfirm
func goTraderOnRspSettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(pSettlementInfoConfirm), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspRemoveParkedOrder
func goTraderOnRspRemoveParkedOrder(userData uintptr, pRemoveParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrder((*CThostFtdcRemoveParkedOrderField)(pRemoveParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspRemoveParkedOrderAction
func goTraderOnRspRemoveParkedOrderAction(userData uintptr, pRemoveParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrderAction((*CThostFtdcRemoveParkedOrderActionField)(pRemoveParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspExecOrderInsert
func goTraderOnRspExecOrderInsert(userData uintptr, pInputExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderInsert((*CThostFtdcInputExecOrderField)(pInputExecOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspExecOrderAction
func goTraderOnRspExecOrderAction(userData uintptr, pInputExecOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderAction((*CThostFtdcInputExecOrderActionField)(pInputExecOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspForQuoteInsert
func goTraderOnRspForQuoteInsert(userData uintptr, pInputForQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspForQuoteInsert((*CThostFtdcInputForQuoteField)(pInputForQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQuoteInsert
func goTraderOnRspQuoteInsert(userData uintptr, pInputQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteInsert((*CThostFtdcInputQuoteField)(pInputQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQuoteAction
func goTraderOnRspQuoteAction(userData uintptr, pInputQuoteAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteAction((*CThostFtdcInputQuoteActionField)(pInputQuoteAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspBatchOrderAction
func goTraderOnRspBatchOrderAction(userData uintptr, pInputBatchOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspBatchOrderAction((*CThostFtdcInputBatchOrderActionField)(pInputBatchOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspOptionSelfCloseInsert
func goTraderOnRspOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(pInputOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspOptionSelfCloseAction
func goTraderOnRspOptionSelfCloseAction(userData uintptr, pInputOptionSelfCloseAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseAction((*CThostFtdcInputOptionSelfCloseActionField)(pInputOptionSelfCloseAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspCombActionInsert
func goTraderOnRspCombActionInsert(userData uintptr, pInputCombAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspCombActionInsert((*CThostFtdcInputCombActionField)(pInputCombAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryOrder
func goTraderOnRspQryOrder(userData uintptr, pOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOrder((*CThostFtdcOrderField)(pOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTrade
func goTraderOnRspQryTrade(userData uintptr, pTrade unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTrade((*CThostFtdcTradeField)(pTrade), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorPosition
func goTraderOnRspQryInvestorPosition(userData uintptr, pInvestorPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPosition((*CThostFtdcInvestorPositionField)(pInvestorPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTradingAccount
func goTraderOnRspQryTradingAccount(userData uintptr, pTradingAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingAccount((*CThostFtdcTradingAccountField)(pTradingAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestor
func goTraderOnRspQryInvestor(userData uintptr, pInvestor unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestor((*CThostFtdcInvestorField)(pInvestor), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTradingCode
func goTraderOnRspQryTradingCode(userData uintptr, pTradingCode unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingCode((*CThostFtdcTradingCodeField)(pTradingCode), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInstrumentMarginRate
func goTraderOnRspQryInstrumentMarginRate(userData uintptr, pInstrumentMarginRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentMarginRate((*CThostFtdcInstrumentMarginRateField)(pInstrumentMarginRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInstrumentCommissionRate
func goTraderOnRspQryInstrumentCommissionRate(userData uintptr, pInstrumentCommissionRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentCommissionRate((*CThostFtdcInstrumentCommissionRateField)(pInstrumentCommissionRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryExchange
func goTraderOnRspQryExchange(userData uintptr, pExchange unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchange((*CThostFtdcExchangeField)(pExchange), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryProduct
func goTraderOnRspQryProduct(userData uintptr, pProduct unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProduct((*CThostFtdcProductField)(pProduct), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInstrument
func goTraderOnRspQryInstrument(userData uintptr, pInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrument((*CThostFtdcInstrumentField)(pInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryDepthMarketData
func goTraderOnRspQryDepthMarketData(userData uintptr, pDepthMarketData unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryDepthMarketData((*CThostFtdcDepthMarketDataField)(pDepthMarketData), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTraderOffer
func goTraderOnRspQryTraderOffer(userData uintptr, pTraderOffer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTraderOffer((*CThostFtdcTraderOfferField)(pTraderOffer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySettlementInfo
func goTraderOnRspQrySettlementInfo(userData uintptr, pSettlementInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfo((*CThostFtdcSettlementInfoField)(pSettlementInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTransferBank
func goTraderOnRspQryTransferBank(userData uintptr, pTransferBank unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferBank((*CThostFtdcTransferBankField)(pTransferBank), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorPositionDetail
func goTraderOnRspQryInvestorPositionDetail(userData uintptr, pInvestorPositionDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionDetail((*CThostFtdcInvestorPositionDetailField)(pInvestorPositionDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryNotice
func goTraderOnRspQryNotice(userData uintptr, pNotice unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryNotice((*CThostFtdcNoticeField)(pNotice), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySettlementInfoConfirm
func goTraderOnRspQrySettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfoConfirm((*CThostFtdcSettlementInfoConfirmField)(pSettlementInfoConfirm), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorPositionCombineDetail
func goTraderOnRspQryInvestorPositionCombineDetail(userData uintptr, pInvestorPositionCombineDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionCombineDetail((*CThostFtdcInvestorPositionCombineDetailField)(pInvestorPositionCombineDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryCFMMCTradingAccountKey
func goTraderOnRspQryCFMMCTradingAccountKey(userData uintptr, pCFMMCTradingAccountKey unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCFMMCTradingAccountKey((*CThostFtdcCFMMCTradingAccountKeyField)(pCFMMCTradingAccountKey), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryEWarrantOffset
func goTraderOnRspQryEWarrantOffset(userData uintptr, pEWarrantOffset unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryEWarrantOffset((*CThostFtdcEWarrantOffsetField)(pEWarrantOffset), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorProductGroupMargin
func goTraderOnRspQryInvestorProductGroupMargin(userData uintptr, pInvestorProductGroupMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProductGroupMargin((*CThostFtdcInvestorProductGroupMarginField)(pInvestorProductGroupMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryExchangeMarginRate
func goTraderOnRspQryExchangeMarginRate(userData uintptr, pExchangeMarginRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRate((*CThostFtdcExchangeMarginRateField)(pExchangeMarginRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryExchangeMarginRateAdjust
func goTraderOnRspQryExchangeMarginRateAdjust(userData uintptr, pExchangeMarginRateAdjust unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRateAdjust((*CThostFtdcExchangeMarginRateAdjustField)(pExchangeMarginRateAdjust), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryExchangeRate
func goTraderOnRspQryExchangeRate(userData uintptr, pExchangeRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeRate((*CThostFtdcExchangeRateField)(pExchangeRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySecAgentACIDMap
func goTraderOnRspQrySecAgentACIDMap(userData uintptr, pSecAgentACIDMap unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentACIDMap((*CThostFtdcSecAgentACIDMapField)(pSecAgentACIDMap), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryProductExchRate
func goTraderOnRspQryProductExchRate(userData uintptr, pProductExchRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductExchRate((*CThostFtdcProductExchRateField)(pProductExchRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryProductGroup
func goTraderOnRspQryProductGroup(userData uintptr, pProductGroup unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductGroup((*CThostFtdcProductGroupField)(pProductGroup), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryMMInstrumentCommissionRate
func goTraderOnRspQryMMInstrumentCommissionRate(userData uintptr, pMMInstrumentCommissionRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMInstrumentCommissionRate((*CThostFtdcMMInstrumentCommissionRateField)(pMMInstrumentCommissionRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryMMOptionInstrCommRate
func goTraderOnRspQryMMOptionInstrCommRate(userData uintptr, pMMOptionInstrCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMOptionInstrCommRate((*CThostFtdcMMOptionInstrCommRateField)(pMMOptionInstrCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInstrumentOrderCommRate
func goTraderOnRspQryInstrumentOrderCommRate(userData uintptr, pInstrumentOrderCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentOrderCommRate((*CThostFtdcInstrumentOrderCommRateField)(pInstrumentOrderCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySecAgentTradingAccount
func goTraderOnRspQrySecAgentTradingAccount(userData uintptr, pTradingAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradingAccount((*CThostFtdcTradingAccountField)(pTradingAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySecAgentCheckMode
func goTraderOnRspQrySecAgentCheckMode(userData uintptr, pSecAgentCheckMode unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentCheckMode((*CThostFtdcSecAgentCheckModeField)(pSecAgentCheckMode), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySecAgentTradeInfo
func goTraderOnRspQrySecAgentTradeInfo(userData uintptr, pSecAgentTradeInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradeInfo((*CThostFtdcSecAgentTradeInfoField)(pSecAgentTradeInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryOptionInstrTradeCost
func goTraderOnRspQryOptionInstrTradeCost(userData uintptr, pOptionInstrTradeCost unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrTradeCost((*CThostFtdcOptionInstrTradeCostField)(pOptionInstrTradeCost), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryOptionInstrCommRate
func goTraderOnRspQryOptionInstrCommRate(userData uintptr, pOptionInstrCommRate unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrCommRate((*CThostFtdcOptionInstrCommRateField)(pOptionInstrCommRate), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryExecOrder
func goTraderOnRspQryExecOrder(userData uintptr, pExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExecOrder((*CThostFtdcExecOrderField)(pExecOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryForQuote
func goTraderOnRspQryForQuote(userData uintptr, pForQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryForQuote((*CThostFtdcForQuoteField)(pForQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryQuote
func goTraderOnRspQryQuote(userData uintptr, pQuote unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryQuote((*CThostFtdcQuoteField)(pQuote), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryOptionSelfClose
func goTraderOnRspQryOptionSelfClose(userData uintptr, pOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionSelfClose((*CThostFtdcOptionSelfCloseField)(pOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestUnit
func goTraderOnRspQryInvestUnit(userData uintptr, pInvestUnit unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestUnit((*CThostFtdcInvestUnitField)(pInvestUnit), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryCombInstrumentGuard
func goTraderOnRspQryCombInstrumentGuard(userData uintptr, pCombInstrumentGuard unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombInstrumentGuard((*CThostFtdcCombInstrumentGuardField)(pCombInstrumentGuard), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryCombAction
func goTraderOnRspQryCombAction(userData uintptr, pCombAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombAction((*CThostFtdcCombActionField)(pCombAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTransferSerial
func goTraderOnRspQryTransferSerial(userData uintptr, pTransferSerial unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferSerial((*CThostFtdcTransferSerialField)(pTransferSerial), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryAccountregister
func goTraderOnRspQryAccountregister(userData uintptr, pAccountregister unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryAccountregister((*CThostFtdcAccountregisterField)(pAccountregister), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspError
func goTraderOnRspError(userData uintptr, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspError((*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRtnOrder
func goTraderOnRtnOrder(userData uintptr, pOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOrder((*CThostFtdcOrderField)(pOrder))
}

//export goTraderOnRtnTrade
func goTraderOnRtnTrade(userData uintptr, pTrade unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTrade((*CThostFtdcTradeField)(pTrade))
}

//export goTraderOnErrRtnOrderInsert
func goTraderOnErrRtnOrderInsert(userData uintptr, pInputOrder unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderInsert((*CThostFtdcInputOrderField)(pInputOrder), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnOrderAction
func goTraderOnErrRtnOrderAction(userData uintptr, pOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderAction((*CThostFtdcOrderActionField)(pOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnInstrumentStatus
func goTraderOnRtnInstrumentStatus(userData uintptr, pInstrumentStatus unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnInstrumentStatus((*CThostFtdcInstrumentStatusField)(pInstrumentStatus))
}

//export goTraderOnRtnBulletin
func goTraderOnRtnBulletin(userData uintptr, pBulletin unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnBulletin((*CThostFtdcBulletinField)(pBulletin))
}

//export goTraderOnRtnTradingNotice
func goTraderOnRtnTradingNotice(userData uintptr, pTradingNoticeInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTradingNotice((*CThostFtdcTradingNoticeInfoField)(pTradingNoticeInfo))
}

//export goTraderOnRtnErrorConditionalOrder
func goTraderOnRtnErrorConditionalOrder(userData uintptr, pErrorConditionalOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnErrorConditionalOrder((*CThostFtdcErrorConditionalOrderField)(pErrorConditionalOrder))
}

//export goTraderOnRtnExecOrder
func goTraderOnRtnExecOrder(userData uintptr, pExecOrder unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnExecOrder((*CThostFtdcExecOrderField)(pExecOrder))
}

//export goTraderOnErrRtnExecOrderInsert
func goTraderOnErrRtnExecOrderInsert(userData uintptr, pInputExecOrder unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderInsert((*CThostFtdcInputExecOrderField)(pInputExecOrder), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnExecOrderAction
func goTraderOnErrRtnExecOrderAction(userData uintptr, pExecOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderAction((*CThostFtdcExecOrderActionField)(pExecOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnForQuoteInsert
func goTraderOnErrRtnForQuoteInsert(userData uintptr, pInputForQuote unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnForQuoteInsert((*CThostFtdcInputForQuoteField)(pInputForQuote), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnQuote
func goTraderOnRtnQuote(userData uintptr, pQuote unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQuote((*CThostFtdcQuoteField)(pQuote))
}

//export goTraderOnErrRtnQuoteInsert
func goTraderOnErrRtnQuoteInsert(userData uintptr, pInputQuote unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteInsert((*CThostFtdcInputQuoteField)(pInputQuote), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnQuoteAction
func goTraderOnErrRtnQuoteAction(userData uintptr, pQuoteAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteAction((*CThostFtdcQuoteActionField)(pQuoteAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnForQuoteRsp
func goTraderOnRtnForQuoteRsp(userData uintptr, pForQuoteRsp unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnForQuoteRsp((*CThostFtdcForQuoteRspField)(pForQuoteRsp))
}

//export goTraderOnRtnCFMMCTradingAccountToken
func goTraderOnRtnCFMMCTradingAccountToken(userData uintptr, pCFMMCTradingAccountToken unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCFMMCTradingAccountToken((*CThostFtdcCFMMCTradingAccountTokenField)(pCFMMCTradingAccountToken))
}

//export goTraderOnErrRtnBatchOrderAction
func goTraderOnErrRtnBatchOrderAction(userData uintptr, pBatchOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBatchOrderAction((*CThostFtdcBatchOrderActionField)(pBatchOrderAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnOptionSelfClose
func goTraderOnRtnOptionSelfClose(userData uintptr, pOptionSelfClose unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOptionSelfClose((*CThostFtdcOptionSelfCloseField)(pOptionSelfClose))
}

//export goTraderOnErrRtnOptionSelfCloseInsert
func goTraderOnErrRtnOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseInsert((*CThostFtdcInputOptionSelfCloseField)(pInputOptionSelfClose), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnOptionSelfCloseAction
func goTraderOnErrRtnOptionSelfCloseAction(userData uintptr, pOptionSelfCloseAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseAction((*CThostFtdcOptionSelfCloseActionField)(pOptionSelfCloseAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnCombAction
func goTraderOnRtnCombAction(userData uintptr, pCombAction unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCombAction((*CThostFtdcCombActionField)(pCombAction))
}

//export goTraderOnErrRtnCombActionInsert
func goTraderOnErrRtnCombActionInsert(userData uintptr, pInputCombAction unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnCombActionInsert((*CThostFtdcInputCombActionField)(pInputCombAction), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRspQryContractBank
func goTraderOnRspQryContractBank(userData uintptr, pContractBank unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryContractBank((*CThostFtdcContractBankField)(pContractBank), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryParkedOrder
func goTraderOnRspQryParkedOrder(userData uintptr, pParkedOrder unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrder((*CThostFtdcParkedOrderField)(pParkedOrder), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryParkedOrderAction
func goTraderOnRspQryParkedOrderAction(userData uintptr, pParkedOrderAction unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrderAction((*CThostFtdcParkedOrderActionField)(pParkedOrderAction), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryTradingNotice
func goTraderOnRspQryTradingNotice(userData uintptr, pTradingNotice unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingNotice((*CThostFtdcTradingNoticeField)(pTradingNotice), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryBrokerTradingParams
func goTraderOnRspQryBrokerTradingParams(userData uintptr, pBrokerTradingParams unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingParams((*CThostFtdcBrokerTradingParamsField)(pBrokerTradingParams), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryBrokerTradingAlgos
func goTraderOnRspQryBrokerTradingAlgos(userData uintptr, pBrokerTradingAlgos unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingAlgos((*CThostFtdcBrokerTradingAlgosField)(pBrokerTradingAlgos), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQueryCFMMCTradingAccountToken
func goTraderOnRspQueryCFMMCTradingAccountToken(userData uintptr, pQueryCFMMCTradingAccountToken unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryCFMMCTradingAccountToken((*CThostFtdcQueryCFMMCTradingAccountTokenField)(pQueryCFMMCTradingAccountToken), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRtnFromBankToFutureByBank
func goTraderOnRtnFromBankToFutureByBank(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByBank((*CThostFtdcRspTransferField)(pRspTransfer))
}

//export goTraderOnRtnFromFutureToBankByBank
func goTraderOnRtnFromFutureToBankByBank(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByBank((*CThostFtdcRspTransferField)(pRspTransfer))
}

//export goTraderOnRtnRepealFromBankToFutureByBank
func goTraderOnRtnRepealFromBankToFutureByBank(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByBank((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRtnRepealFromFutureToBankByBank
func goTraderOnRtnRepealFromFutureToBankByBank(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByBank((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRtnFromBankToFutureByFuture
func goTraderOnRtnFromBankToFutureByFuture(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByFuture((*CThostFtdcRspTransferField)(pRspTransfer))
}

//export goTraderOnRtnFromFutureToBankByFuture
func goTraderOnRtnFromFutureToBankByFuture(userData uintptr, pRspTransfer unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByFuture((*CThostFtdcRspTransferField)(pRspTransfer))
}

//export goTraderOnRtnRepealFromBankToFutureByFutureManual
func goTraderOnRtnRepealFromBankToFutureByFutureManual(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFutureManual((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRtnRepealFromFutureToBankByFutureManual
func goTraderOnRtnRepealFromFutureToBankByFutureManual(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFutureManual((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRtnQueryBankBalanceByFuture
func goTraderOnRtnQueryBankBalanceByFuture(userData uintptr, pNotifyQueryAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQueryBankBalanceByFuture((*CThostFtdcNotifyQueryAccountField)(pNotifyQueryAccount))
}

//export goTraderOnErrRtnBankToFutureByFuture
func goTraderOnErrRtnBankToFutureByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBankToFutureByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnFutureToBankByFuture
func goTraderOnErrRtnFutureToBankByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnFutureToBankByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnRepealBankToFutureByFutureManual
func goTraderOnErrRtnRepealBankToFutureByFutureManual(userData uintptr, pReqRepeal unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealBankToFutureByFutureManual((*CThostFtdcReqRepealField)(pReqRepeal), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnRepealFutureToBankByFutureManual
func goTraderOnErrRtnRepealFutureToBankByFutureManual(userData uintptr, pReqRepeal unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealFutureToBankByFutureManual((*CThostFtdcReqRepealField)(pReqRepeal), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnErrRtnQueryBankBalanceByFuture
func goTraderOnErrRtnQueryBankBalanceByFuture(userData uintptr, pReqQueryAccount unsafe.Pointer, pRspInfo unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQueryBankBalanceByFuture((*CThostFtdcReqQueryAccountField)(pReqQueryAccount), (*CThostFtdcRspInfoField)(pRspInfo))
}

//export goTraderOnRtnRepealFromBankToFutureByFuture
func goTraderOnRtnRepealFromBankToFutureByFuture(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFuture((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRtnRepealFromFutureToBankByFuture
func goTraderOnRtnRepealFromFutureToBankByFuture(userData uintptr, pRspRepeal unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFuture((*CThostFtdcRspRepealField)(pRspRepeal))
}

//export goTraderOnRspFromBankToFutureByFuture
func goTraderOnRspFromBankToFutureByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromBankToFutureByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspFromFutureToBankByFuture
func goTraderOnRspFromFutureToBankByFuture(userData uintptr, pReqTransfer unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromFutureToBankByFuture((*CThostFtdcReqTransferField)(pReqTransfer), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQueryBankAccountMoneyByFuture
func goTraderOnRspQueryBankAccountMoneyByFuture(userData uintptr, pReqQueryAccount unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryBankAccountMoneyByFuture((*CThostFtdcReqQueryAccountField)(pReqQueryAccount), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRtnOpenAccountByBank
func goTraderOnRtnOpenAccountByBank(userData uintptr, pOpenAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOpenAccountByBank((*CThostFtdcOpenAccountField)(pOpenAccount))
}

//export goTraderOnRtnCancelAccountByBank
func goTraderOnRtnCancelAccountByBank(userData uintptr, pCancelAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCancelAccountByBank((*CThostFtdcCancelAccountField)(pCancelAccount))
}

//export goTraderOnRtnChangeAccountByBank
func goTraderOnRtnChangeAccountByBank(userData uintptr, pChangeAccount unsafe.Pointer) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnChangeAccountByBank((*CThostFtdcChangeAccountField)(pChangeAccount))
}

//export goTraderOnRspQryClassifiedInstrument
func goTraderOnRspQryClassifiedInstrument(userData uintptr, pInstrument unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryClassifiedInstrument((*CThostFtdcInstrumentField)(pInstrument), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryCombPromotionParam
func goTraderOnRspQryCombPromotionParam(userData uintptr, pCombPromotionParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombPromotionParam((*CThostFtdcCombPromotionParamField)(pCombPromotionParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRiskSettleInvstPosition
func goTraderOnRspQryRiskSettleInvstPosition(userData uintptr, pRiskSettleInvstPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleInvstPosition((*CThostFtdcRiskSettleInvstPositionField)(pRiskSettleInvstPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRiskSettleProductStatus
func goTraderOnRspQryRiskSettleProductStatus(userData uintptr, pRiskSettleProductStatus unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleProductStatus((*CThostFtdcRiskSettleProductStatusField)(pRiskSettleProductStatus), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMFutureParameter
func goTraderOnRspQrySPBMFutureParameter(userData uintptr, pSPBMFutureParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMFutureParameter((*CThostFtdcSPBMFutureParameterField)(pSPBMFutureParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMOptionParameter
func goTraderOnRspQrySPBMOptionParameter(userData uintptr, pSPBMOptionParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMOptionParameter((*CThostFtdcSPBMOptionParameterField)(pSPBMOptionParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMIntraParameter
func goTraderOnRspQrySPBMIntraParameter(userData uintptr, pSPBMIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMIntraParameter((*CThostFtdcSPBMIntraParameterField)(pSPBMIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMInterParameter
func goTraderOnRspQrySPBMInterParameter(userData uintptr, pSPBMInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInterParameter((*CThostFtdcSPBMInterParameterField)(pSPBMInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMPortfDefinition
func goTraderOnRspQrySPBMPortfDefinition(userData uintptr, pSPBMPortfDefinition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMPortfDefinition((*CThostFtdcSPBMPortfDefinitionField)(pSPBMPortfDefinition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMInvestorPortfDef
func goTraderOnRspQrySPBMInvestorPortfDef(userData uintptr, pSPBMInvestorPortfDef unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInvestorPortfDef((*CThostFtdcSPBMInvestorPortfDefField)(pSPBMInvestorPortfDef), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorPortfMarginRatio
func goTraderOnRspQryInvestorPortfMarginRatio(userData uintptr, pInvestorPortfMarginRatio unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfMarginRatio((*CThostFtdcInvestorPortfMarginRatioField)(pInvestorPortfMarginRatio), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorProdSPBMDetail
func goTraderOnRspQryInvestorProdSPBMDetail(userData uintptr, pInvestorProdSPBMDetail unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdSPBMDetail((*CThostFtdcInvestorProdSPBMDetailField)(pInvestorProdSPBMDetail), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorCommoditySPMMMargin
func goTraderOnRspQryInvestorCommoditySPMMMargin(userData uintptr, pInvestorCommoditySPMMMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommoditySPMMMargin((*CThostFtdcInvestorCommoditySPMMMarginField)(pInvestorCommoditySPMMMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorCommodityGroupSPMMMargin
func goTraderOnRspQryInvestorCommodityGroupSPMMMargin(userData uintptr, pInvestorCommodityGroupSPMMMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommodityGroupSPMMMargin((*CThostFtdcInvestorCommodityGroupSPMMMarginField)(pInvestorCommodityGroupSPMMMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPMMInstParam
func goTraderOnRspQrySPMMInstParam(userData uintptr, pSPMMInstParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMInstParam((*CThostFtdcSPMMInstParamField)(pSPMMInstParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPMMProductParam
func goTraderOnRspQrySPMMProductParam(userData uintptr, pSPMMProductParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMProductParam((*CThostFtdcSPMMProductParamField)(pSPMMProductParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQrySPBMAddOnInterParameter
func goTraderOnRspQrySPBMAddOnInterParameter(userData uintptr, pSPBMAddOnInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMAddOnInterParameter((*CThostFtdcSPBMAddOnInterParameterField)(pSPBMAddOnInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSCombProductInfo
func goTraderOnRspQryRCAMSCombProductInfo(userData uintptr, pRCAMSCombProductInfo unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSCombProductInfo((*CThostFtdcRCAMSCombProductInfoField)(pRCAMSCombProductInfo), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSInstrParameter
func goTraderOnRspQryRCAMSInstrParameter(userData uintptr, pRCAMSInstrParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInstrParameter((*CThostFtdcRCAMSInstrParameterField)(pRCAMSInstrParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSIntraParameter
func goTraderOnRspQryRCAMSIntraParameter(userData uintptr, pRCAMSIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSIntraParameter((*CThostFtdcRCAMSIntraParameterField)(pRCAMSIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSInterParameter
func goTraderOnRspQryRCAMSInterParameter(userData uintptr, pRCAMSInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInterParameter((*CThostFtdcRCAMSInterParameterField)(pRCAMSInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSShortOptAdjustParam
func goTraderOnRspQryRCAMSShortOptAdjustParam(userData uintptr, pRCAMSShortOptAdjustParam unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSShortOptAdjustParam((*CThostFtdcRCAMSShortOptAdjustParamField)(pRCAMSShortOptAdjustParam), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRCAMSInvestorCombPosition
func goTraderOnRspQryRCAMSInvestorCombPosition(userData uintptr, pRCAMSInvestorCombPosition unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInvestorCombPosition((*CThostFtdcRCAMSInvestorCombPositionField)(pRCAMSInvestorCombPosition), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorProdRCAMSMargin
func goTraderOnRspQryInvestorProdRCAMSMargin(userData uintptr, pInvestorProdRCAMSMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRCAMSMargin((*CThostFtdcInvestorProdRCAMSMarginField)(pInvestorProdRCAMSMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRULEInstrParameter
func goTraderOnRspQryRULEInstrParameter(userData uintptr, pRULEInstrParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInstrParameter((*CThostFtdcRULEInstrParameterField)(pRULEInstrParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRULEIntraParameter
func goTraderOnRspQryRULEIntraParameter(userData uintptr, pRULEIntraParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEIntraParameter((*CThostFtdcRULEIntraParameterField)(pRULEIntraParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryRULEInterParameter
func goTraderOnRspQryRULEInterParameter(userData uintptr, pRULEInterParameter unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInterParameter((*CThostFtdcRULEInterParameterField)(pRULEInterParameter), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorProdRULEMargin
func goTraderOnRspQryInvestorProdRULEMargin(userData uintptr, pInvestorProdRULEMargin unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRULEMargin((*CThostFtdcInvestorProdRULEMarginField)(pInvestorProdRULEMargin), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}

//export goTraderOnRspQryInvestorPortfSetting
func goTraderOnRspQryInvestorPortfSetting(userData uintptr, pInvestorPortfSetting unsafe.Pointer, pRspInfo unsafe.Pointer, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfSetting((*CThostFtdcInvestorPortfSettingField)(pInvestorPortfSetting), (*CThostFtdcRspInfoField)(pRspInfo), nRequestID, bIsLast)
}
