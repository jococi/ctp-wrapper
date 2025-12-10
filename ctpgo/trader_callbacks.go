package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 交易回调实现

import "unsafe"

// #include <stdint.h>
import "C"

// ========== 回调函数 ==========

//export goOnFrontConnected
func goOnFrontConnected(userData uintptr) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontConnected()
}

//export goOnFrontDisconnected
func goOnFrontDisconnected(userData uintptr, nReason int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnFrontDisconnected(nReason)
}

//export goOnHeartBeatWarning
func goOnHeartBeatWarning(userData uintptr, nTimeLapse int32) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnHeartBeatWarning(nTimeLapse)
}

//export goOnRspAuthenticate
func goOnRspAuthenticate(userData uintptr, pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspAuthenticate(pRspAuthenticateField, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUserLogin
func goOnRspUserLogin(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogin(pRspUserLogin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUserLogout
func goOnRspUserLogout(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserLogout(pUserLogout, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUserPasswordUpdate
func goOnRspUserPasswordUpdate(userData uintptr, pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserPasswordUpdate(pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspTradingAccountPasswordUpdate
func goOnRspTradingAccountPasswordUpdate(userData uintptr, pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspUserAuthMethod
func goOnRspUserAuthMethod(userData uintptr, pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspUserAuthMethod(pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspGenUserCaptcha
func goOnRspGenUserCaptcha(userData uintptr, pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserCaptcha(pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspGenUserText
func goOnRspGenUserText(userData uintptr, pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspGenUserText(pRspGenUserText, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspOrderInsert
func goOnRspOrderInsert(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderInsert(pInputOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspParkedOrderInsert
func goOnRspParkedOrderInsert(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderInsert(pParkedOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspParkedOrderAction
func goOnRspParkedOrderAction(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspParkedOrderAction(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspOrderAction
func goOnRspOrderAction(userData uintptr, pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOrderAction(pInputOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryMaxOrderVolume
func goOnRspQryMaxOrderVolume(userData uintptr, pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMaxOrderVolume(pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspSettlementInfoConfirm
func goOnRspSettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspSettlementInfoConfirm(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspRemoveParkedOrder
func goOnRspRemoveParkedOrder(userData uintptr, pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrder(pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspRemoveParkedOrderAction
func goOnRspRemoveParkedOrderAction(userData uintptr, pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspExecOrderInsert
func goOnRspExecOrderInsert(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderInsert(pInputExecOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspExecOrderAction
func goOnRspExecOrderAction(userData uintptr, pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspExecOrderAction(pInputExecOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspForQuoteInsert
func goOnRspForQuoteInsert(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspForQuoteInsert(pInputForQuote, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQuoteInsert
func goOnRspQuoteInsert(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteInsert(pInputQuote, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQuoteAction
func goOnRspQuoteAction(userData uintptr, pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQuoteAction(pInputQuoteAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspBatchOrderAction
func goOnRspBatchOrderAction(userData uintptr, pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspBatchOrderAction(pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspOptionSelfCloseInsert
func goOnRspOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseInsert(pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspOptionSelfCloseAction
func goOnRspOptionSelfCloseAction(userData uintptr, pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspCombActionInsert
func goOnRspCombActionInsert(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspCombActionInsert(pInputCombAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryOrder
func goOnRspQryOrder(userData uintptr, pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOrder(pOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTrade
func goOnRspQryTrade(userData uintptr, pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTrade(pTrade, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorPosition
func goOnRspQryInvestorPosition(userData uintptr, pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPosition(pInvestorPosition, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTradingAccount
func goOnRspQryTradingAccount(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingAccount(pTradingAccount, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestor
func goOnRspQryInvestor(userData uintptr, pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestor(pInvestor, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTradingCode
func goOnRspQryTradingCode(userData uintptr, pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingCode(pTradingCode, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInstrumentMarginRate
func goOnRspQryInstrumentMarginRate(userData uintptr, pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentMarginRate(pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInstrumentCommissionRate
func goOnRspQryInstrumentCommissionRate(userData uintptr, pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryExchange
func goOnRspQryExchange(userData uintptr, pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchange(pExchange, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryProduct
func goOnRspQryProduct(userData uintptr, pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProduct(pProduct, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInstrument
func goOnRspQryInstrument(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrument(pInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryDepthMarketData
func goOnRspQryDepthMarketData(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryDepthMarketData(pDepthMarketData, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTraderOffer
func goOnRspQryTraderOffer(userData uintptr, pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTraderOffer(pTraderOffer, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySettlementInfo
func goOnRspQrySettlementInfo(userData uintptr, pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfo(pSettlementInfo, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTransferBank
func goOnRspQryTransferBank(userData uintptr, pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferBank(pTransferBank, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorPositionDetail
func goOnRspQryInvestorPositionDetail(userData uintptr, pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionDetail(pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryNotice
func goOnRspQryNotice(userData uintptr, pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryNotice(pNotice, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySettlementInfoConfirm
func goOnRspQrySettlementInfoConfirm(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorPositionCombineDetail
func goOnRspQryInvestorPositionCombineDetail(userData uintptr, pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryCFMMCTradingAccountKey
func goOnRspQryCFMMCTradingAccountKey(userData uintptr, pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryEWarrantOffset
func goOnRspQryEWarrantOffset(userData uintptr, pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryEWarrantOffset(pEWarrantOffset, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorProductGroupMargin
func goOnRspQryInvestorProductGroupMargin(userData uintptr, pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryExchangeMarginRate
func goOnRspQryExchangeMarginRate(userData uintptr, pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRate(pExchangeMarginRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryExchangeMarginRateAdjust
func goOnRspQryExchangeMarginRateAdjust(userData uintptr, pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryExchangeRate
func goOnRspQryExchangeRate(userData uintptr, pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExchangeRate(pExchangeRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySecAgentACIDMap
func goOnRspQrySecAgentACIDMap(userData uintptr, pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentACIDMap(pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryProductExchRate
func goOnRspQryProductExchRate(userData uintptr, pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductExchRate(pProductExchRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryProductGroup
func goOnRspQryProductGroup(userData uintptr, pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryProductGroup(pProductGroup, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryMMInstrumentCommissionRate
func goOnRspQryMMInstrumentCommissionRate(userData uintptr, pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryMMOptionInstrCommRate
func goOnRspQryMMOptionInstrCommRate(userData uintptr, pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInstrumentOrderCommRate
func goOnRspQryInstrumentOrderCommRate(userData uintptr, pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySecAgentTradingAccount
func goOnRspQrySecAgentTradingAccount(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradingAccount(pTradingAccount, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySecAgentCheckMode
func goOnRspQrySecAgentCheckMode(userData uintptr, pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentCheckMode(pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySecAgentTradeInfo
func goOnRspQrySecAgentTradeInfo(userData uintptr, pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryOptionInstrTradeCost
func goOnRspQryOptionInstrTradeCost(userData uintptr, pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryOptionInstrCommRate
func goOnRspQryOptionInstrCommRate(userData uintptr, pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionInstrCommRate(pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryExecOrder
func goOnRspQryExecOrder(userData uintptr, pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryExecOrder(pExecOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryForQuote
func goOnRspQryForQuote(userData uintptr, pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryForQuote(pForQuote, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryQuote
func goOnRspQryQuote(userData uintptr, pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryQuote(pQuote, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryOptionSelfClose
func goOnRspQryOptionSelfClose(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryOptionSelfClose(pOptionSelfClose, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestUnit
func goOnRspQryInvestUnit(userData uintptr, pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestUnit(pInvestUnit, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryCombInstrumentGuard
func goOnRspQryCombInstrumentGuard(userData uintptr, pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombInstrumentGuard(pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryCombAction
func goOnRspQryCombAction(userData uintptr, pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombAction(pCombAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTransferSerial
func goOnRspQryTransferSerial(userData uintptr, pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTransferSerial(pTransferSerial, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryAccountregister
func goOnRspQryAccountregister(userData uintptr, pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryAccountregister(pAccountregister, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspError
func goOnRspError(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspError(pRspInfo, nRequestID, bIsLast)
}

//export goOnRtnOrder
func goOnRtnOrder(userData uintptr, pOrder *CThostFtdcOrderField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOrder(pOrder)
}

//export goOnRtnTrade
func goOnRtnTrade(userData uintptr, pTrade *CThostFtdcTradeField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTrade(pTrade)
}

//export goOnErrRtnOrderInsert
func goOnErrRtnOrderInsert(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderInsert(pInputOrder, pRspInfo)
}

//export goOnErrRtnOrderAction
func goOnErrRtnOrderAction(userData uintptr, pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOrderAction(pOrderAction, pRspInfo)
}

//export goOnRtnInstrumentStatus
func goOnRtnInstrumentStatus(userData uintptr, pInstrumentStatus *CThostFtdcInstrumentStatusField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnInstrumentStatus(pInstrumentStatus)
}

//export goOnRtnBulletin
func goOnRtnBulletin(userData uintptr, pBulletin *CThostFtdcBulletinField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnBulletin(pBulletin)
}

//export goOnRtnTradingNotice
func goOnRtnTradingNotice(userData uintptr, pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnTradingNotice(pTradingNoticeInfo)
}

//export goOnRtnErrorConditionalOrder
func goOnRtnErrorConditionalOrder(userData uintptr, pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnErrorConditionalOrder(pErrorConditionalOrder)
}

//export goOnRtnExecOrder
func goOnRtnExecOrder(userData uintptr, pExecOrder *CThostFtdcExecOrderField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnExecOrder(pExecOrder)
}

//export goOnErrRtnExecOrderInsert
func goOnErrRtnExecOrderInsert(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderInsert(pInputExecOrder, pRspInfo)
}

//export goOnErrRtnExecOrderAction
func goOnErrRtnExecOrderAction(userData uintptr, pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnExecOrderAction(pExecOrderAction, pRspInfo)
}

//export goOnErrRtnForQuoteInsert
func goOnErrRtnForQuoteInsert(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnForQuoteInsert(pInputForQuote, pRspInfo)
}

//export goOnRtnQuote
func goOnRtnQuote(userData uintptr, pQuote *CThostFtdcQuoteField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQuote(pQuote)
}

//export goOnErrRtnQuoteInsert
func goOnErrRtnQuoteInsert(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteInsert(pInputQuote, pRspInfo)
}

//export goOnErrRtnQuoteAction
func goOnErrRtnQuoteAction(userData uintptr, pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQuoteAction(pQuoteAction, pRspInfo)
}

//export goOnRtnForQuoteRsp
func goOnRtnForQuoteRsp(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnForQuoteRsp(pForQuoteRsp)
}

//export goOnRtnCFMMCTradingAccountToken
func goOnRtnCFMMCTradingAccountToken(userData uintptr, pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken)
}

//export goOnErrRtnBatchOrderAction
func goOnErrRtnBatchOrderAction(userData uintptr, pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBatchOrderAction(pBatchOrderAction, pRspInfo)
}

//export goOnRtnOptionSelfClose
func goOnRtnOptionSelfClose(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOptionSelfClose(pOptionSelfClose)
}

//export goOnErrRtnOptionSelfCloseInsert
func goOnErrRtnOptionSelfCloseInsert(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose, pRspInfo)
}

//export goOnErrRtnOptionSelfCloseAction
func goOnErrRtnOptionSelfCloseAction(userData uintptr, pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction, pRspInfo)
}

//export goOnRtnCombAction
func goOnRtnCombAction(userData uintptr, pCombAction *CThostFtdcCombActionField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCombAction(pCombAction)
}

//export goOnErrRtnCombActionInsert
func goOnErrRtnCombActionInsert(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnCombActionInsert(pInputCombAction, pRspInfo)
}

//export goOnRspQryContractBank
func goOnRspQryContractBank(userData uintptr, pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryContractBank(pContractBank, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryParkedOrder
func goOnRspQryParkedOrder(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrder(pParkedOrder, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryParkedOrderAction
func goOnRspQryParkedOrderAction(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryParkedOrderAction(pParkedOrderAction, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryTradingNotice
func goOnRspQryTradingNotice(userData uintptr, pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryTradingNotice(pTradingNotice, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryBrokerTradingParams
func goOnRspQryBrokerTradingParams(userData uintptr, pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingParams(pBrokerTradingParams, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryBrokerTradingAlgos
func goOnRspQryBrokerTradingAlgos(userData uintptr, pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQueryCFMMCTradingAccountToken
func goOnRspQueryCFMMCTradingAccountToken(userData uintptr, pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast)
}

//export goOnRtnFromBankToFutureByBank
func goOnRtnFromBankToFutureByBank(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByBank(pRspTransfer)
}

//export goOnRtnFromFutureToBankByBank
func goOnRtnFromFutureToBankByBank(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByBank(pRspTransfer)
}

//export goOnRtnRepealFromBankToFutureByBank
func goOnRtnRepealFromBankToFutureByBank(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByBank(pRspRepeal)
}

//export goOnRtnRepealFromFutureToBankByBank
func goOnRtnRepealFromFutureToBankByBank(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByBank(pRspRepeal)
}

//export goOnRtnFromBankToFutureByFuture
func goOnRtnFromBankToFutureByFuture(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromBankToFutureByFuture(pRspTransfer)
}

//export goOnRtnFromFutureToBankByFuture
func goOnRtnFromFutureToBankByFuture(userData uintptr, pRspTransfer *CThostFtdcRspTransferField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnFromFutureToBankByFuture(pRspTransfer)
}

//export goOnRtnRepealFromBankToFutureByFutureManual
func goOnRtnRepealFromBankToFutureByFutureManual(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal)
}

//export goOnRtnRepealFromFutureToBankByFutureManual
func goOnRtnRepealFromFutureToBankByFutureManual(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal)
}

//export goOnRtnQueryBankBalanceByFuture
func goOnRtnQueryBankBalanceByFuture(userData uintptr, pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount)
}

//export goOnErrRtnBankToFutureByFuture
func goOnErrRtnBankToFutureByFuture(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnBankToFutureByFuture(pReqTransfer, pRspInfo)
}

//export goOnErrRtnFutureToBankByFuture
func goOnErrRtnFutureToBankByFuture(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnFutureToBankByFuture(pReqTransfer, pRspInfo)
}

//export goOnErrRtnRepealBankToFutureByFutureManual
func goOnErrRtnRepealBankToFutureByFutureManual(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal, pRspInfo)
}

//export goOnErrRtnRepealFutureToBankByFutureManual
func goOnErrRtnRepealFutureToBankByFutureManual(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal, pRspInfo)
}

//export goOnErrRtnQueryBankBalanceByFuture
func goOnErrRtnQueryBankBalanceByFuture(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount, pRspInfo)
}

//export goOnRtnRepealFromBankToFutureByFuture
func goOnRtnRepealFromBankToFutureByFuture(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromBankToFutureByFuture(pRspRepeal)
}

//export goOnRtnRepealFromFutureToBankByFuture
func goOnRtnRepealFromFutureToBankByFuture(userData uintptr, pRspRepeal *CThostFtdcRspRepealField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnRepealFromFutureToBankByFuture(pRspRepeal)
}

//export goOnRspFromBankToFutureByFuture
func goOnRspFromBankToFutureByFuture(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromBankToFutureByFuture(pReqTransfer, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspFromFutureToBankByFuture
func goOnRspFromFutureToBankByFuture(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspFromFutureToBankByFuture(pReqTransfer, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQueryBankAccountMoneyByFuture
func goOnRspQueryBankAccountMoneyByFuture(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount, pRspInfo, nRequestID, bIsLast)
}

//export goOnRtnOpenAccountByBank
func goOnRtnOpenAccountByBank(userData uintptr, pOpenAccount *CThostFtdcOpenAccountField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnOpenAccountByBank(pOpenAccount)
}

//export goOnRtnCancelAccountByBank
func goOnRtnCancelAccountByBank(userData uintptr, pCancelAccount *CThostFtdcCancelAccountField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnCancelAccountByBank(pCancelAccount)
}

//export goOnRtnChangeAccountByBank
func goOnRtnChangeAccountByBank(userData uintptr, pChangeAccount *CThostFtdcChangeAccountField) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRtnChangeAccountByBank(pChangeAccount)
}

//export goOnRspQryClassifiedInstrument
func goOnRspQryClassifiedInstrument(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryClassifiedInstrument(pInstrument, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryCombPromotionParam
func goOnRspQryCombPromotionParam(userData uintptr, pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryCombPromotionParam(pCombPromotionParam, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRiskSettleInvstPosition
func goOnRspQryRiskSettleInvstPosition(userData uintptr, pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRiskSettleProductStatus
func goOnRspQryRiskSettleProductStatus(userData uintptr, pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMFutureParameter
func goOnRspQrySPBMFutureParameter(userData uintptr, pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMFutureParameter(pSPBMFutureParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMOptionParameter
func goOnRspQrySPBMOptionParameter(userData uintptr, pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMOptionParameter(pSPBMOptionParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMIntraParameter
func goOnRspQrySPBMIntraParameter(userData uintptr, pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMIntraParameter(pSPBMIntraParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMInterParameter
func goOnRspQrySPBMInterParameter(userData uintptr, pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInterParameter(pSPBMInterParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMPortfDefinition
func goOnRspQrySPBMPortfDefinition(userData uintptr, pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMPortfDefinition(pSPBMPortfDefinition, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMInvestorPortfDef
func goOnRspQrySPBMInvestorPortfDef(userData uintptr, pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMInvestorPortfDef(pSPBMInvestorPortfDef, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorPortfMarginRatio
func goOnRspQryInvestorPortfMarginRatio(userData uintptr, pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfMarginRatio(pInvestorPortfMarginRatio, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorProdSPBMDetail
func goOnRspQryInvestorProdSPBMDetail(userData uintptr, pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdSPBMDetail(pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorCommoditySPMMMargin
func goOnRspQryInvestorCommoditySPMMMargin(userData uintptr, pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommoditySPMMMargin(pInvestorCommoditySPMMMargin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorCommodityGroupSPMMMargin
func goOnRspQryInvestorCommodityGroupSPMMMargin(userData uintptr, pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorCommodityGroupSPMMMargin(pInvestorCommodityGroupSPMMMargin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPMMInstParam
func goOnRspQrySPMMInstParam(userData uintptr, pSPMMInstParam *CThostFtdcSPMMInstParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMInstParam(pSPMMInstParam, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPMMProductParam
func goOnRspQrySPMMProductParam(userData uintptr, pSPMMProductParam *CThostFtdcSPMMProductParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPMMProductParam(pSPMMProductParam, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQrySPBMAddOnInterParameter
func goOnRspQrySPBMAddOnInterParameter(userData uintptr, pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQrySPBMAddOnInterParameter(pSPBMAddOnInterParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSCombProductInfo
func goOnRspQryRCAMSCombProductInfo(userData uintptr, pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSCombProductInfo(pRCAMSCombProductInfo, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSInstrParameter
func goOnRspQryRCAMSInstrParameter(userData uintptr, pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInstrParameter(pRCAMSInstrParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSIntraParameter
func goOnRspQryRCAMSIntraParameter(userData uintptr, pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSIntraParameter(pRCAMSIntraParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSInterParameter
func goOnRspQryRCAMSInterParameter(userData uintptr, pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInterParameter(pRCAMSInterParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSShortOptAdjustParam
func goOnRspQryRCAMSShortOptAdjustParam(userData uintptr, pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSShortOptAdjustParam(pRCAMSShortOptAdjustParam, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRCAMSInvestorCombPosition
func goOnRspQryRCAMSInvestorCombPosition(userData uintptr, pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRCAMSInvestorCombPosition(pRCAMSInvestorCombPosition, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorProdRCAMSMargin
func goOnRspQryInvestorProdRCAMSMargin(userData uintptr, pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRCAMSMargin(pInvestorProdRCAMSMargin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRULEInstrParameter
func goOnRspQryRULEInstrParameter(userData uintptr, pRULEInstrParameter *CThostFtdcRULEInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInstrParameter(pRULEInstrParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRULEIntraParameter
func goOnRspQryRULEIntraParameter(userData uintptr, pRULEIntraParameter *CThostFtdcRULEIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEIntraParameter(pRULEIntraParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryRULEInterParameter
func goOnRspQryRULEInterParameter(userData uintptr, pRULEInterParameter *CThostFtdcRULEInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryRULEInterParameter(pRULEInterParameter, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorProdRULEMargin
func goOnRspQryInvestorProdRULEMargin(userData uintptr, pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorProdRULEMargin(pInvestorProdRULEMargin, pRspInfo, nRequestID, bIsLast)
}

//export goOnRspQryInvestorPortfSetting
func goOnRspQryInvestorPortfSetting(userData uintptr, pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	api := getTraderInstance(userData)
	if api == nil || api.spi == nil {
		return
	}
	api.spi.OnRspQryInvestorPortfSetting(pInvestorPortfSetting, pRspInfo, nRequestID, bIsLast)
}
