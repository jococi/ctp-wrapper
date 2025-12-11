package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// 默认 SPI 空实现，可用于嵌入

// DefaultTraderSpi 默认交易回调实现（空实现）
// 使用方式：嵌入到自定义结构体中，只需实现需要的方法
// 例如：type MySpi struct { DefaultTraderSpi }
//
//	func (s *MySpi) OnRtnOrder(...) { ... }
type DefaultTraderSpi struct{}

func (s *DefaultTraderSpi) OnFrontConnected() {
	// 空实现
}

func (s *DefaultTraderSpi) OnFrontDisconnected(nReason int32) {
	// 空实现
}

func (s *DefaultTraderSpi) OnHeartBeatWarning(nTimeLapse int32) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspAuthenticate(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspUserPasswordUpdate(pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspUserAuthMethod(pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspGenUserCaptcha(pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspGenUserText(pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspOrderInsert(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspParkedOrderInsert(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspOrderAction(pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryMaxOrderVolume(pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspSettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspRemoveParkedOrder(pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspExecOrderAction(pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQuoteAction(pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspBatchOrderAction(pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryOrder(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTrade(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorPosition(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTradingAccount(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestor(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTradingCode(pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInstrumentMarginRate(pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryExchange(pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryProduct(pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInstrument(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTraderOffer(pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySettlementInfo(pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTransferBank(pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorPositionDetail(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryNotice(pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryEWarrantOffset(pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryExchangeMarginRate(pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryExchangeRate(pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySecAgentACIDMap(pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryProductExchRate(pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryProductGroup(pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySecAgentTradingAccount(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySecAgentCheckMode(pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryOptionInstrCommRate(pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryExecOrder(pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryForQuote(pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryQuote(pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestUnit(pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryCombInstrumentGuard(pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryCombAction(pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTransferSerial(pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryAccountregister(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnOrder(pOrder *CThostFtdcOrderField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnTrade(pTrade *CThostFtdcTradeField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnOrderInsert(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnOrderAction(pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnInstrumentStatus(pInstrumentStatus *CThostFtdcInstrumentStatusField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnBulletin(pBulletin *CThostFtdcBulletinField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnTradingNotice(pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnErrorConditionalOrder(pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnExecOrder(pExecOrder *CThostFtdcExecOrderField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnExecOrderAction(pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnQuote(pQuote *CThostFtdcQuoteField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnQuoteAction(pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnBatchOrderAction(pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnCombAction(pCombAction *CThostFtdcCombActionField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryContractBank(pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryParkedOrder(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryTradingNotice(pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryBrokerTradingParams(pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnFromBankToFutureByBank(pRspTransfer *CThostFtdcRspTransferField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnFromFutureToBankByBank(pRspTransfer *CThostFtdcRspTransferField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromBankToFutureByBank(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromFutureToBankByBank(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnFromBankToFutureByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnFromFutureToBankByFuture(pRspTransfer *CThostFtdcRspTransferField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromBankToFutureByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnRepealFromFutureToBankByFuture(pRspRepeal *CThostFtdcRspRepealField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspFromBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspFromFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnOpenAccountByBank(pOpenAccount *CThostFtdcOpenAccountField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnCancelAccountByBank(pCancelAccount *CThostFtdcCancelAccountField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRtnChangeAccountByBank(pChangeAccount *CThostFtdcChangeAccountField) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryClassifiedInstrument(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryCombPromotionParam(pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMFutureParameter(pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMOptionParameter(pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMIntraParameter(pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMInterParameter(pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMPortfDefinition(pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMInvestorPortfDef(pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorPortfMarginRatio(pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorProdSPBMDetail(pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorCommoditySPMMMargin(pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorCommodityGroupSPMMMargin(pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPMMInstParam(pSPMMInstParam *CThostFtdcSPMMInstParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPMMProductParam(pSPMMProductParam *CThostFtdcSPMMProductParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQrySPBMAddOnInterParameter(pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSCombProductInfo(pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSInstrParameter(pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSIntraParameter(pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSInterParameter(pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSShortOptAdjustParam(pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRCAMSInvestorCombPosition(pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorProdRCAMSMargin(pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRULEInstrParameter(pRULEInstrParameter *CThostFtdcRULEInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRULEIntraParameter(pRULEIntraParameter *CThostFtdcRULEIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryRULEInterParameter(pRULEInterParameter *CThostFtdcRULEInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorProdRULEMargin(pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}

func (s *DefaultTraderSpi) OnRspQryInvestorPortfSetting(pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	// 空实现
}
