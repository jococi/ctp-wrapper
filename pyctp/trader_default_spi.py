"""
默认交易 SPI 空实现

此文件由代码生成器自动生成，请勿手动修改
默认 SPI 空实现，可用于嵌入
"""

import ctypes
from .trader_api import TraderSpi
from .struct import *

class DefaultTraderSpi(TraderSpi):
    """默认交易回调实现（空实现）"""
    
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

    def OnRspAuthenticate(self, pRspAuthenticateField: ctypes.POINTER(CThostFtdcRspAuthenticateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUserLogin(self, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUserPasswordUpdate(self, pUserPasswordUpdate: ctypes.POINTER(CThostFtdcUserPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspTradingAccountPasswordUpdate(self, pTradingAccountPasswordUpdate: ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspUserAuthMethod(self, pRspUserAuthMethod: ctypes.POINTER(CThostFtdcRspUserAuthMethodField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspGenUserCaptcha(self, pRspGenUserCaptcha: ctypes.POINTER(CThostFtdcRspGenUserCaptchaField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspGenUserText(self, pRspGenUserText: ctypes.POINTER(CThostFtdcRspGenUserTextField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspOrderInsert(self, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspParkedOrderInsert(self, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspParkedOrderAction(self, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspOrderAction(self, pInputOrderAction: ctypes.POINTER(CThostFtdcInputOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryMaxOrderVolume(self, pQryMaxOrderVolume: ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspSettlementInfoConfirm(self, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspRemoveParkedOrder(self, pRemoveParkedOrder: ctypes.POINTER(CThostFtdcRemoveParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspRemoveParkedOrderAction(self, pRemoveParkedOrderAction: ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspExecOrderInsert(self, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspExecOrderAction(self, pInputExecOrderAction: ctypes.POINTER(CThostFtdcInputExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspForQuoteInsert(self, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQuoteInsert(self, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQuoteAction(self, pInputQuoteAction: ctypes.POINTER(CThostFtdcInputQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspBatchOrderAction(self, pInputBatchOrderAction: ctypes.POINTER(CThostFtdcInputBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspOptionSelfCloseInsert(self, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspOptionSelfCloseAction(self, pInputOptionSelfCloseAction: ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspCombActionInsert(self, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryOrder(self, pOrder: ctypes.POINTER(CThostFtdcOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTrade(self, pTrade: ctypes.POINTER(CThostFtdcTradeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorPosition(self, pInvestorPosition: ctypes.POINTER(CThostFtdcInvestorPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTradingAccount(self, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestor(self, pInvestor: ctypes.POINTER(CThostFtdcInvestorField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTradingCode(self, pTradingCode: ctypes.POINTER(CThostFtdcTradingCodeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInstrumentMarginRate(self, pInstrumentMarginRate: ctypes.POINTER(CThostFtdcInstrumentMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInstrumentCommissionRate(self, pInstrumentCommissionRate: ctypes.POINTER(CThostFtdcInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryExchange(self, pExchange: ctypes.POINTER(CThostFtdcExchangeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryProduct(self, pProduct: ctypes.POINTER(CThostFtdcProductField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInstrument(self, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryDepthMarketData(self, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTraderOffer(self, pTraderOffer: ctypes.POINTER(CThostFtdcTraderOfferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySettlementInfo(self, pSettlementInfo: ctypes.POINTER(CThostFtdcSettlementInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTransferBank(self, pTransferBank: ctypes.POINTER(CThostFtdcTransferBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorPositionDetail(self, pInvestorPositionDetail: ctypes.POINTER(CThostFtdcInvestorPositionDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryNotice(self, pNotice: ctypes.POINTER(CThostFtdcNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySettlementInfoConfirm(self, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorPositionCombineDetail(self, pInvestorPositionCombineDetail: ctypes.POINTER(CThostFtdcInvestorPositionCombineDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryCFMMCTradingAccountKey(self, pCFMMCTradingAccountKey: ctypes.POINTER(CThostFtdcCFMMCTradingAccountKeyField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryEWarrantOffset(self, pEWarrantOffset: ctypes.POINTER(CThostFtdcEWarrantOffsetField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorProductGroupMargin(self, pInvestorProductGroupMargin: ctypes.POINTER(CThostFtdcInvestorProductGroupMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryExchangeMarginRate(self, pExchangeMarginRate: ctypes.POINTER(CThostFtdcExchangeMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryExchangeMarginRateAdjust(self, pExchangeMarginRateAdjust: ctypes.POINTER(CThostFtdcExchangeMarginRateAdjustField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryExchangeRate(self, pExchangeRate: ctypes.POINTER(CThostFtdcExchangeRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySecAgentACIDMap(self, pSecAgentACIDMap: ctypes.POINTER(CThostFtdcSecAgentACIDMapField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryProductExchRate(self, pProductExchRate: ctypes.POINTER(CThostFtdcProductExchRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryProductGroup(self, pProductGroup: ctypes.POINTER(CThostFtdcProductGroupField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryMMInstrumentCommissionRate(self, pMMInstrumentCommissionRate: ctypes.POINTER(CThostFtdcMMInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryMMOptionInstrCommRate(self, pMMOptionInstrCommRate: ctypes.POINTER(CThostFtdcMMOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInstrumentOrderCommRate(self, pInstrumentOrderCommRate: ctypes.POINTER(CThostFtdcInstrumentOrderCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySecAgentTradingAccount(self, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySecAgentCheckMode(self, pSecAgentCheckMode: ctypes.POINTER(CThostFtdcSecAgentCheckModeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySecAgentTradeInfo(self, pSecAgentTradeInfo: ctypes.POINTER(CThostFtdcSecAgentTradeInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryOptionInstrTradeCost(self, pOptionInstrTradeCost: ctypes.POINTER(CThostFtdcOptionInstrTradeCostField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryOptionInstrCommRate(self, pOptionInstrCommRate: ctypes.POINTER(CThostFtdcOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryExecOrder(self, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryForQuote(self, pForQuote: ctypes.POINTER(CThostFtdcForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryQuote(self, pQuote: ctypes.POINTER(CThostFtdcQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryOptionSelfClose(self, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestUnit(self, pInvestUnit: ctypes.POINTER(CThostFtdcInvestUnitField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryCombInstrumentGuard(self, pCombInstrumentGuard: ctypes.POINTER(CThostFtdcCombInstrumentGuardField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryCombAction(self, pCombAction: ctypes.POINTER(CThostFtdcCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTransferSerial(self, pTransferSerial: ctypes.POINTER(CThostFtdcTransferSerialField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryAccountregister(self, pAccountregister: ctypes.POINTER(CThostFtdcAccountregisterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspError(self, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRtnOrder(self, pOrder: ctypes.POINTER(CThostFtdcOrderField)):
        """空实现"""
        pass

    def OnRtnTrade(self, pTrade: ctypes.POINTER(CThostFtdcTradeField)):
        """空实现"""
        pass

    def OnErrRtnOrderInsert(self, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnOrderAction(self, pOrderAction: ctypes.POINTER(CThostFtdcOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnInstrumentStatus(self, pInstrumentStatus: ctypes.POINTER(CThostFtdcInstrumentStatusField)):
        """空实现"""
        pass

    def OnRtnBulletin(self, pBulletin: ctypes.POINTER(CThostFtdcBulletinField)):
        """空实现"""
        pass

    def OnRtnTradingNotice(self, pTradingNoticeInfo: ctypes.POINTER(CThostFtdcTradingNoticeInfoField)):
        """空实现"""
        pass

    def OnRtnErrorConditionalOrder(self, pErrorConditionalOrder: ctypes.POINTER(CThostFtdcErrorConditionalOrderField)):
        """空实现"""
        pass

    def OnRtnExecOrder(self, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField)):
        """空实现"""
        pass

    def OnErrRtnExecOrderInsert(self, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnExecOrderAction(self, pExecOrderAction: ctypes.POINTER(CThostFtdcExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnForQuoteInsert(self, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnQuote(self, pQuote: ctypes.POINTER(CThostFtdcQuoteField)):
        """空实现"""
        pass

    def OnErrRtnQuoteInsert(self, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnQuoteAction(self, pQuoteAction: ctypes.POINTER(CThostFtdcQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnForQuoteRsp(self, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
        """空实现"""
        pass

    def OnRtnCFMMCTradingAccountToken(self, pCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcCFMMCTradingAccountTokenField)):
        """空实现"""
        pass

    def OnErrRtnBatchOrderAction(self, pBatchOrderAction: ctypes.POINTER(CThostFtdcBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnOptionSelfClose(self, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField)):
        """空实现"""
        pass

    def OnErrRtnOptionSelfCloseInsert(self, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnOptionSelfCloseAction(self, pOptionSelfCloseAction: ctypes.POINTER(CThostFtdcOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnCombAction(self, pCombAction: ctypes.POINTER(CThostFtdcCombActionField)):
        """空实现"""
        pass

    def OnErrRtnCombActionInsert(self, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRspQryContractBank(self, pContractBank: ctypes.POINTER(CThostFtdcContractBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryParkedOrder(self, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryParkedOrderAction(self, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryTradingNotice(self, pTradingNotice: ctypes.POINTER(CThostFtdcTradingNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryBrokerTradingParams(self, pBrokerTradingParams: ctypes.POINTER(CThostFtdcBrokerTradingParamsField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryBrokerTradingAlgos(self, pBrokerTradingAlgos: ctypes.POINTER(CThostFtdcBrokerTradingAlgosField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQueryCFMMCTradingAccountToken(self, pQueryCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRtnFromBankToFutureByBank(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """空实现"""
        pass

    def OnRtnFromFutureToBankByBank(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """空实现"""
        pass

    def OnRtnRepealFromBankToFutureByBank(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRtnRepealFromFutureToBankByBank(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRtnFromBankToFutureByFuture(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """空实现"""
        pass

    def OnRtnFromFutureToBankByFuture(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """空实现"""
        pass

    def OnRtnRepealFromBankToFutureByFutureManual(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRtnRepealFromFutureToBankByFutureManual(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRtnQueryBankBalanceByFuture(self, pNotifyQueryAccount: ctypes.POINTER(CThostFtdcNotifyQueryAccountField)):
        """空实现"""
        pass

    def OnErrRtnBankToFutureByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnFutureToBankByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnRepealBankToFutureByFutureManual(self, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnRepealFutureToBankByFutureManual(self, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnErrRtnQueryBankBalanceByFuture(self, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """空实现"""
        pass

    def OnRtnRepealFromBankToFutureByFuture(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRtnRepealFromFutureToBankByFuture(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """空实现"""
        pass

    def OnRspFromBankToFutureByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspFromFutureToBankByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQueryBankAccountMoneyByFuture(self, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRtnOpenAccountByBank(self, pOpenAccount: ctypes.POINTER(CThostFtdcOpenAccountField)):
        """空实现"""
        pass

    def OnRtnCancelAccountByBank(self, pCancelAccount: ctypes.POINTER(CThostFtdcCancelAccountField)):
        """空实现"""
        pass

    def OnRtnChangeAccountByBank(self, pChangeAccount: ctypes.POINTER(CThostFtdcChangeAccountField)):
        """空实现"""
        pass

    def OnRspQryClassifiedInstrument(self, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryCombPromotionParam(self, pCombPromotionParam: ctypes.POINTER(CThostFtdcCombPromotionParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRiskSettleInvstPosition(self, pRiskSettleInvstPosition: ctypes.POINTER(CThostFtdcRiskSettleInvstPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRiskSettleProductStatus(self, pRiskSettleProductStatus: ctypes.POINTER(CThostFtdcRiskSettleProductStatusField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMFutureParameter(self, pSPBMFutureParameter: ctypes.POINTER(CThostFtdcSPBMFutureParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMOptionParameter(self, pSPBMOptionParameter: ctypes.POINTER(CThostFtdcSPBMOptionParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMIntraParameter(self, pSPBMIntraParameter: ctypes.POINTER(CThostFtdcSPBMIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMInterParameter(self, pSPBMInterParameter: ctypes.POINTER(CThostFtdcSPBMInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMPortfDefinition(self, pSPBMPortfDefinition: ctypes.POINTER(CThostFtdcSPBMPortfDefinitionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMInvestorPortfDef(self, pSPBMInvestorPortfDef: ctypes.POINTER(CThostFtdcSPBMInvestorPortfDefField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorPortfMarginRatio(self, pInvestorPortfMarginRatio: ctypes.POINTER(CThostFtdcInvestorPortfMarginRatioField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorProdSPBMDetail(self, pInvestorProdSPBMDetail: ctypes.POINTER(CThostFtdcInvestorProdSPBMDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorCommoditySPMMMargin(self, pInvestorCommoditySPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommoditySPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorCommodityGroupSPMMMargin(self, pInvestorCommodityGroupSPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommodityGroupSPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPMMInstParam(self, pSPMMInstParam: ctypes.POINTER(CThostFtdcSPMMInstParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPMMProductParam(self, pSPMMProductParam: ctypes.POINTER(CThostFtdcSPMMProductParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQrySPBMAddOnInterParameter(self, pSPBMAddOnInterParameter: ctypes.POINTER(CThostFtdcSPBMAddOnInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSCombProductInfo(self, pRCAMSCombProductInfo: ctypes.POINTER(CThostFtdcRCAMSCombProductInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSInstrParameter(self, pRCAMSInstrParameter: ctypes.POINTER(CThostFtdcRCAMSInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSIntraParameter(self, pRCAMSIntraParameter: ctypes.POINTER(CThostFtdcRCAMSIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSInterParameter(self, pRCAMSInterParameter: ctypes.POINTER(CThostFtdcRCAMSInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSShortOptAdjustParam(self, pRCAMSShortOptAdjustParam: ctypes.POINTER(CThostFtdcRCAMSShortOptAdjustParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRCAMSInvestorCombPosition(self, pRCAMSInvestorCombPosition: ctypes.POINTER(CThostFtdcRCAMSInvestorCombPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorProdRCAMSMargin(self, pInvestorProdRCAMSMargin: ctypes.POINTER(CThostFtdcInvestorProdRCAMSMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRULEInstrParameter(self, pRULEInstrParameter: ctypes.POINTER(CThostFtdcRULEInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRULEIntraParameter(self, pRULEIntraParameter: ctypes.POINTER(CThostFtdcRULEIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryRULEInterParameter(self, pRULEInterParameter: ctypes.POINTER(CThostFtdcRULEInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorProdRULEMargin(self, pInvestorProdRULEMargin: ctypes.POINTER(CThostFtdcInvestorProdRULEMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass

    def OnRspQryInvestorPortfSetting(self, pInvestorPortfSetting: ctypes.POINTER(CThostFtdcInvestorPortfSettingField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """空实现"""
        pass
