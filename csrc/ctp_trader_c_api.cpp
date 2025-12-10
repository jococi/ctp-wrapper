/**
 * CTP Trader API - C 接口实现
 * 
 * 自动生成，请勿手动修改
 */

#include "ctp_trader_c_api.h"

// 平台相关的原始 CTP 头文件
// 注意: 编译时需要正确设置头文件搜索路径
#ifdef _WIN32
    #include "ctpapi/windows/ThostFtdcTraderApi.h"
    #include "ctpapi/windows/DataCollect.h"
#elif __APPLE__
    #include "ctpapi/macos/ThostFtdcTraderApi.h"
    #include "ctpapi/macos/DataCollect.h"
#elif __linux__
    #include "ctpapi/linux/ThostFtdcTraderApi.h"
    #include "ctpapi/linux/DataCollect.h"
#endif

#include <cstring>

// ========== SPI 包装类 ==========
class TraderSpiWrapper : public CThostFtdcTraderSpi {
public:
    void* userData = nullptr;

    TraderOnFrontConnectedCallback onFrontConnected = nullptr;
    TraderOnFrontDisconnectedCallback onFrontDisconnected = nullptr;
    TraderOnHeartBeatWarningCallback onHeartBeatWarning = nullptr;
    TraderOnRspAuthenticateCallback onRspAuthenticate = nullptr;
    TraderOnRspUserLoginCallback onRspUserLogin = nullptr;
    TraderOnRspUserLogoutCallback onRspUserLogout = nullptr;
    TraderOnRspUserPasswordUpdateCallback onRspUserPasswordUpdate = nullptr;
    TraderOnRspTradingAccountPasswordUpdateCallback onRspTradingAccountPasswordUpdate = nullptr;
    TraderOnRspUserAuthMethodCallback onRspUserAuthMethod = nullptr;
    TraderOnRspGenUserCaptchaCallback onRspGenUserCaptcha = nullptr;
    TraderOnRspGenUserTextCallback onRspGenUserText = nullptr;
    TraderOnRspOrderInsertCallback onRspOrderInsert = nullptr;
    TraderOnRspParkedOrderInsertCallback onRspParkedOrderInsert = nullptr;
    TraderOnRspParkedOrderActionCallback onRspParkedOrderAction = nullptr;
    TraderOnRspOrderActionCallback onRspOrderAction = nullptr;
    TraderOnRspQryMaxOrderVolumeCallback onRspQryMaxOrderVolume = nullptr;
    TraderOnRspSettlementInfoConfirmCallback onRspSettlementInfoConfirm = nullptr;
    TraderOnRspRemoveParkedOrderCallback onRspRemoveParkedOrder = nullptr;
    TraderOnRspRemoveParkedOrderActionCallback onRspRemoveParkedOrderAction = nullptr;
    TraderOnRspExecOrderInsertCallback onRspExecOrderInsert = nullptr;
    TraderOnRspExecOrderActionCallback onRspExecOrderAction = nullptr;
    TraderOnRspForQuoteInsertCallback onRspForQuoteInsert = nullptr;
    TraderOnRspQuoteInsertCallback onRspQuoteInsert = nullptr;
    TraderOnRspQuoteActionCallback onRspQuoteAction = nullptr;
    TraderOnRspBatchOrderActionCallback onRspBatchOrderAction = nullptr;
    TraderOnRspOptionSelfCloseInsertCallback onRspOptionSelfCloseInsert = nullptr;
    TraderOnRspOptionSelfCloseActionCallback onRspOptionSelfCloseAction = nullptr;
    TraderOnRspCombActionInsertCallback onRspCombActionInsert = nullptr;
    TraderOnRspQryOrderCallback onRspQryOrder = nullptr;
    TraderOnRspQryTradeCallback onRspQryTrade = nullptr;
    TraderOnRspQryInvestorPositionCallback onRspQryInvestorPosition = nullptr;
    TraderOnRspQryTradingAccountCallback onRspQryTradingAccount = nullptr;
    TraderOnRspQryInvestorCallback onRspQryInvestor = nullptr;
    TraderOnRspQryTradingCodeCallback onRspQryTradingCode = nullptr;
    TraderOnRspQryInstrumentMarginRateCallback onRspQryInstrumentMarginRate = nullptr;
    TraderOnRspQryInstrumentCommissionRateCallback onRspQryInstrumentCommissionRate = nullptr;
    TraderOnRspQryExchangeCallback onRspQryExchange = nullptr;
    TraderOnRspQryProductCallback onRspQryProduct = nullptr;
    TraderOnRspQryInstrumentCallback onRspQryInstrument = nullptr;
    TraderOnRspQryDepthMarketDataCallback onRspQryDepthMarketData = nullptr;
    TraderOnRspQryTraderOfferCallback onRspQryTraderOffer = nullptr;
    TraderOnRspQrySettlementInfoCallback onRspQrySettlementInfo = nullptr;
    TraderOnRspQryTransferBankCallback onRspQryTransferBank = nullptr;
    TraderOnRspQryInvestorPositionDetailCallback onRspQryInvestorPositionDetail = nullptr;
    TraderOnRspQryNoticeCallback onRspQryNotice = nullptr;
    TraderOnRspQrySettlementInfoConfirmCallback onRspQrySettlementInfoConfirm = nullptr;
    TraderOnRspQryInvestorPositionCombineDetailCallback onRspQryInvestorPositionCombineDetail = nullptr;
    TraderOnRspQryCFMMCTradingAccountKeyCallback onRspQryCFMMCTradingAccountKey = nullptr;
    TraderOnRspQryEWarrantOffsetCallback onRspQryEWarrantOffset = nullptr;
    TraderOnRspQryInvestorProductGroupMarginCallback onRspQryInvestorProductGroupMargin = nullptr;
    TraderOnRspQryExchangeMarginRateCallback onRspQryExchangeMarginRate = nullptr;
    TraderOnRspQryExchangeMarginRateAdjustCallback onRspQryExchangeMarginRateAdjust = nullptr;
    TraderOnRspQryExchangeRateCallback onRspQryExchangeRate = nullptr;
    TraderOnRspQrySecAgentACIDMapCallback onRspQrySecAgentACIDMap = nullptr;
    TraderOnRspQryProductExchRateCallback onRspQryProductExchRate = nullptr;
    TraderOnRspQryProductGroupCallback onRspQryProductGroup = nullptr;
    TraderOnRspQryMMInstrumentCommissionRateCallback onRspQryMMInstrumentCommissionRate = nullptr;
    TraderOnRspQryMMOptionInstrCommRateCallback onRspQryMMOptionInstrCommRate = nullptr;
    TraderOnRspQryInstrumentOrderCommRateCallback onRspQryInstrumentOrderCommRate = nullptr;
    TraderOnRspQrySecAgentTradingAccountCallback onRspQrySecAgentTradingAccount = nullptr;
    TraderOnRspQrySecAgentCheckModeCallback onRspQrySecAgentCheckMode = nullptr;
    TraderOnRspQrySecAgentTradeInfoCallback onRspQrySecAgentTradeInfo = nullptr;
    TraderOnRspQryOptionInstrTradeCostCallback onRspQryOptionInstrTradeCost = nullptr;
    TraderOnRspQryOptionInstrCommRateCallback onRspQryOptionInstrCommRate = nullptr;
    TraderOnRspQryExecOrderCallback onRspQryExecOrder = nullptr;
    TraderOnRspQryForQuoteCallback onRspQryForQuote = nullptr;
    TraderOnRspQryQuoteCallback onRspQryQuote = nullptr;
    TraderOnRspQryOptionSelfCloseCallback onRspQryOptionSelfClose = nullptr;
    TraderOnRspQryInvestUnitCallback onRspQryInvestUnit = nullptr;
    TraderOnRspQryCombInstrumentGuardCallback onRspQryCombInstrumentGuard = nullptr;
    TraderOnRspQryCombActionCallback onRspQryCombAction = nullptr;
    TraderOnRspQryTransferSerialCallback onRspQryTransferSerial = nullptr;
    TraderOnRspQryAccountregisterCallback onRspQryAccountregister = nullptr;
    TraderOnRspErrorCallback onRspError = nullptr;
    TraderOnRtnOrderCallback onRtnOrder = nullptr;
    TraderOnRtnTradeCallback onRtnTrade = nullptr;
    TraderOnErrRtnOrderInsertCallback onErrRtnOrderInsert = nullptr;
    TraderOnErrRtnOrderActionCallback onErrRtnOrderAction = nullptr;
    TraderOnRtnInstrumentStatusCallback onRtnInstrumentStatus = nullptr;
    TraderOnRtnBulletinCallback onRtnBulletin = nullptr;
    TraderOnRtnTradingNoticeCallback onRtnTradingNotice = nullptr;
    TraderOnRtnErrorConditionalOrderCallback onRtnErrorConditionalOrder = nullptr;
    TraderOnRtnExecOrderCallback onRtnExecOrder = nullptr;
    TraderOnErrRtnExecOrderInsertCallback onErrRtnExecOrderInsert = nullptr;
    TraderOnErrRtnExecOrderActionCallback onErrRtnExecOrderAction = nullptr;
    TraderOnErrRtnForQuoteInsertCallback onErrRtnForQuoteInsert = nullptr;
    TraderOnRtnQuoteCallback onRtnQuote = nullptr;
    TraderOnErrRtnQuoteInsertCallback onErrRtnQuoteInsert = nullptr;
    TraderOnErrRtnQuoteActionCallback onErrRtnQuoteAction = nullptr;
    TraderOnRtnForQuoteRspCallback onRtnForQuoteRsp = nullptr;
    TraderOnRtnCFMMCTradingAccountTokenCallback onRtnCFMMCTradingAccountToken = nullptr;
    TraderOnErrRtnBatchOrderActionCallback onErrRtnBatchOrderAction = nullptr;
    TraderOnRtnOptionSelfCloseCallback onRtnOptionSelfClose = nullptr;
    TraderOnErrRtnOptionSelfCloseInsertCallback onErrRtnOptionSelfCloseInsert = nullptr;
    TraderOnErrRtnOptionSelfCloseActionCallback onErrRtnOptionSelfCloseAction = nullptr;
    TraderOnRtnCombActionCallback onRtnCombAction = nullptr;
    TraderOnErrRtnCombActionInsertCallback onErrRtnCombActionInsert = nullptr;
    TraderOnRspQryContractBankCallback onRspQryContractBank = nullptr;
    TraderOnRspQryParkedOrderCallback onRspQryParkedOrder = nullptr;
    TraderOnRspQryParkedOrderActionCallback onRspQryParkedOrderAction = nullptr;
    TraderOnRspQryTradingNoticeCallback onRspQryTradingNotice = nullptr;
    TraderOnRspQryBrokerTradingParamsCallback onRspQryBrokerTradingParams = nullptr;
    TraderOnRspQryBrokerTradingAlgosCallback onRspQryBrokerTradingAlgos = nullptr;
    TraderOnRspQueryCFMMCTradingAccountTokenCallback onRspQueryCFMMCTradingAccountToken = nullptr;
    TraderOnRtnFromBankToFutureByBankCallback onRtnFromBankToFutureByBank = nullptr;
    TraderOnRtnFromFutureToBankByBankCallback onRtnFromFutureToBankByBank = nullptr;
    TraderOnRtnRepealFromBankToFutureByBankCallback onRtnRepealFromBankToFutureByBank = nullptr;
    TraderOnRtnRepealFromFutureToBankByBankCallback onRtnRepealFromFutureToBankByBank = nullptr;
    TraderOnRtnFromBankToFutureByFutureCallback onRtnFromBankToFutureByFuture = nullptr;
    TraderOnRtnFromFutureToBankByFutureCallback onRtnFromFutureToBankByFuture = nullptr;
    TraderOnRtnRepealFromBankToFutureByFutureManualCallback onRtnRepealFromBankToFutureByFutureManual = nullptr;
    TraderOnRtnRepealFromFutureToBankByFutureManualCallback onRtnRepealFromFutureToBankByFutureManual = nullptr;
    TraderOnRtnQueryBankBalanceByFutureCallback onRtnQueryBankBalanceByFuture = nullptr;
    TraderOnErrRtnBankToFutureByFutureCallback onErrRtnBankToFutureByFuture = nullptr;
    TraderOnErrRtnFutureToBankByFutureCallback onErrRtnFutureToBankByFuture = nullptr;
    TraderOnErrRtnRepealBankToFutureByFutureManualCallback onErrRtnRepealBankToFutureByFutureManual = nullptr;
    TraderOnErrRtnRepealFutureToBankByFutureManualCallback onErrRtnRepealFutureToBankByFutureManual = nullptr;
    TraderOnErrRtnQueryBankBalanceByFutureCallback onErrRtnQueryBankBalanceByFuture = nullptr;
    TraderOnRtnRepealFromBankToFutureByFutureCallback onRtnRepealFromBankToFutureByFuture = nullptr;
    TraderOnRtnRepealFromFutureToBankByFutureCallback onRtnRepealFromFutureToBankByFuture = nullptr;
    TraderOnRspFromBankToFutureByFutureCallback onRspFromBankToFutureByFuture = nullptr;
    TraderOnRspFromFutureToBankByFutureCallback onRspFromFutureToBankByFuture = nullptr;
    TraderOnRspQueryBankAccountMoneyByFutureCallback onRspQueryBankAccountMoneyByFuture = nullptr;
    TraderOnRtnOpenAccountByBankCallback onRtnOpenAccountByBank = nullptr;
    TraderOnRtnCancelAccountByBankCallback onRtnCancelAccountByBank = nullptr;
    TraderOnRtnChangeAccountByBankCallback onRtnChangeAccountByBank = nullptr;
    TraderOnRspQryClassifiedInstrumentCallback onRspQryClassifiedInstrument = nullptr;
    TraderOnRspQryCombPromotionParamCallback onRspQryCombPromotionParam = nullptr;
    TraderOnRspQryRiskSettleInvstPositionCallback onRspQryRiskSettleInvstPosition = nullptr;
    TraderOnRspQryRiskSettleProductStatusCallback onRspQryRiskSettleProductStatus = nullptr;
    TraderOnRspQrySPBMFutureParameterCallback onRspQrySPBMFutureParameter = nullptr;
    TraderOnRspQrySPBMOptionParameterCallback onRspQrySPBMOptionParameter = nullptr;
    TraderOnRspQrySPBMIntraParameterCallback onRspQrySPBMIntraParameter = nullptr;
    TraderOnRspQrySPBMInterParameterCallback onRspQrySPBMInterParameter = nullptr;
    TraderOnRspQrySPBMPortfDefinitionCallback onRspQrySPBMPortfDefinition = nullptr;
    TraderOnRspQrySPBMInvestorPortfDefCallback onRspQrySPBMInvestorPortfDef = nullptr;
    TraderOnRspQryInvestorPortfMarginRatioCallback onRspQryInvestorPortfMarginRatio = nullptr;
    TraderOnRspQryInvestorProdSPBMDetailCallback onRspQryInvestorProdSPBMDetail = nullptr;
    TraderOnRspQryInvestorCommoditySPMMMarginCallback onRspQryInvestorCommoditySPMMMargin = nullptr;
    TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback onRspQryInvestorCommodityGroupSPMMMargin = nullptr;
    TraderOnRspQrySPMMInstParamCallback onRspQrySPMMInstParam = nullptr;
    TraderOnRspQrySPMMProductParamCallback onRspQrySPMMProductParam = nullptr;
    TraderOnRspQrySPBMAddOnInterParameterCallback onRspQrySPBMAddOnInterParameter = nullptr;
    TraderOnRspQryRCAMSCombProductInfoCallback onRspQryRCAMSCombProductInfo = nullptr;
    TraderOnRspQryRCAMSInstrParameterCallback onRspQryRCAMSInstrParameter = nullptr;
    TraderOnRspQryRCAMSIntraParameterCallback onRspQryRCAMSIntraParameter = nullptr;
    TraderOnRspQryRCAMSInterParameterCallback onRspQryRCAMSInterParameter = nullptr;
    TraderOnRspQryRCAMSShortOptAdjustParamCallback onRspQryRCAMSShortOptAdjustParam = nullptr;
    TraderOnRspQryRCAMSInvestorCombPositionCallback onRspQryRCAMSInvestorCombPosition = nullptr;
    TraderOnRspQryInvestorProdRCAMSMarginCallback onRspQryInvestorProdRCAMSMargin = nullptr;
    TraderOnRspQryRULEInstrParameterCallback onRspQryRULEInstrParameter = nullptr;
    TraderOnRspQryRULEIntraParameterCallback onRspQryRULEIntraParameter = nullptr;
    TraderOnRspQryRULEInterParameterCallback onRspQryRULEInterParameter = nullptr;
    TraderOnRspQryInvestorProdRULEMarginCallback onRspQryInvestorProdRULEMargin = nullptr;
    TraderOnRspQryInvestorPortfSettingCallback onRspQryInvestorPortfSetting = nullptr;

    void OnFrontConnected() override {
        if (onFrontConnected) {
            onFrontConnected(userData);
        }
    }

    void OnFrontDisconnected(int nReason) override {
        if (onFrontDisconnected) {
            onFrontDisconnected(userData, nReason);
        }
    }

    void OnHeartBeatWarning(int nTimeLapse) override {
        if (onHeartBeatWarning) {
            onHeartBeatWarning(userData, nTimeLapse);
        }
    }

    void OnRspAuthenticate(CThostFtdcRspAuthenticateField* pRspAuthenticateField, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspAuthenticate) {
            onRspAuthenticate(userData, pRspAuthenticateField, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUserLogin(CThostFtdcRspUserLoginField* pRspUserLogin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserLogin) {
            onRspUserLogin(userData, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUserLogout(CThostFtdcUserLogoutField* pUserLogout, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserLogout) {
            onRspUserLogout(userData, pUserLogout, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUserPasswordUpdate(CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserPasswordUpdate) {
            onRspUserPasswordUpdate(userData, pUserPasswordUpdate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspTradingAccountPasswordUpdate(CThostFtdcTradingAccountPasswordUpdateField* pTradingAccountPasswordUpdate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspTradingAccountPasswordUpdate) {
            onRspTradingAccountPasswordUpdate(userData, pTradingAccountPasswordUpdate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUserAuthMethod(CThostFtdcRspUserAuthMethodField* pRspUserAuthMethod, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserAuthMethod) {
            onRspUserAuthMethod(userData, pRspUserAuthMethod, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspGenUserCaptcha(CThostFtdcRspGenUserCaptchaField* pRspGenUserCaptcha, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspGenUserCaptcha) {
            onRspGenUserCaptcha(userData, pRspGenUserCaptcha, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspGenUserText(CThostFtdcRspGenUserTextField* pRspGenUserText, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspGenUserText) {
            onRspGenUserText(userData, pRspGenUserText, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspOrderInsert(CThostFtdcInputOrderField* pInputOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspOrderInsert) {
            onRspOrderInsert(userData, pInputOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspParkedOrderInsert(CThostFtdcParkedOrderField* pParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspParkedOrderInsert) {
            onRspParkedOrderInsert(userData, pParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspParkedOrderAction(CThostFtdcParkedOrderActionField* pParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspParkedOrderAction) {
            onRspParkedOrderAction(userData, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspOrderAction(CThostFtdcInputOrderActionField* pInputOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspOrderAction) {
            onRspOrderAction(userData, pInputOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryMaxOrderVolume(CThostFtdcQryMaxOrderVolumeField* pQryMaxOrderVolume, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryMaxOrderVolume) {
            onRspQryMaxOrderVolume(userData, pQryMaxOrderVolume, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspSettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspSettlementInfoConfirm) {
            onRspSettlementInfoConfirm(userData, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspRemoveParkedOrder(CThostFtdcRemoveParkedOrderField* pRemoveParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspRemoveParkedOrder) {
            onRspRemoveParkedOrder(userData, pRemoveParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspRemoveParkedOrderAction(CThostFtdcRemoveParkedOrderActionField* pRemoveParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspRemoveParkedOrderAction) {
            onRspRemoveParkedOrderAction(userData, pRemoveParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspExecOrderInsert(CThostFtdcInputExecOrderField* pInputExecOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspExecOrderInsert) {
            onRspExecOrderInsert(userData, pInputExecOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspExecOrderAction(CThostFtdcInputExecOrderActionField* pInputExecOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspExecOrderAction) {
            onRspExecOrderAction(userData, pInputExecOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspForQuoteInsert(CThostFtdcInputForQuoteField* pInputForQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspForQuoteInsert) {
            onRspForQuoteInsert(userData, pInputForQuote, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQuoteInsert(CThostFtdcInputQuoteField* pInputQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQuoteInsert) {
            onRspQuoteInsert(userData, pInputQuote, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQuoteAction(CThostFtdcInputQuoteActionField* pInputQuoteAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQuoteAction) {
            onRspQuoteAction(userData, pInputQuoteAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspBatchOrderAction(CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspBatchOrderAction) {
            onRspBatchOrderAction(userData, pInputBatchOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspOptionSelfCloseInsert) {
            onRspOptionSelfCloseInsert(userData, pInputOptionSelfClose, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspOptionSelfCloseAction(CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspOptionSelfCloseAction) {
            onRspOptionSelfCloseAction(userData, pInputOptionSelfCloseAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspCombActionInsert(CThostFtdcInputCombActionField* pInputCombAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspCombActionInsert) {
            onRspCombActionInsert(userData, pInputCombAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryOrder(CThostFtdcOrderField* pOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryOrder) {
            onRspQryOrder(userData, pOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTrade(CThostFtdcTradeField* pTrade, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTrade) {
            onRspQryTrade(userData, pTrade, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorPosition(CThostFtdcInvestorPositionField* pInvestorPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorPosition) {
            onRspQryInvestorPosition(userData, pInvestorPosition, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTradingAccount(CThostFtdcTradingAccountField* pTradingAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTradingAccount) {
            onRspQryTradingAccount(userData, pTradingAccount, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestor(CThostFtdcInvestorField* pInvestor, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestor) {
            onRspQryInvestor(userData, pInvestor, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTradingCode(CThostFtdcTradingCodeField* pTradingCode, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTradingCode) {
            onRspQryTradingCode(userData, pTradingCode, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInstrumentMarginRate(CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInstrumentMarginRate) {
            onRspQryInstrumentMarginRate(userData, pInstrumentMarginRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInstrumentCommissionRate(CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInstrumentCommissionRate) {
            onRspQryInstrumentCommissionRate(userData, pInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryExchange(CThostFtdcExchangeField* pExchange, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryExchange) {
            onRspQryExchange(userData, pExchange, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryProduct(CThostFtdcProductField* pProduct, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryProduct) {
            onRspQryProduct(userData, pProduct, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInstrument(CThostFtdcInstrumentField* pInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInstrument) {
            onRspQryInstrument(userData, pInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryDepthMarketData(CThostFtdcDepthMarketDataField* pDepthMarketData, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryDepthMarketData) {
            onRspQryDepthMarketData(userData, pDepthMarketData, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTraderOffer(CThostFtdcTraderOfferField* pTraderOffer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTraderOffer) {
            onRspQryTraderOffer(userData, pTraderOffer, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySettlementInfo(CThostFtdcSettlementInfoField* pSettlementInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySettlementInfo) {
            onRspQrySettlementInfo(userData, pSettlementInfo, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTransferBank(CThostFtdcTransferBankField* pTransferBank, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTransferBank) {
            onRspQryTransferBank(userData, pTransferBank, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorPositionDetail(CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorPositionDetail) {
            onRspQryInvestorPositionDetail(userData, pInvestorPositionDetail, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryNotice(CThostFtdcNoticeField* pNotice, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryNotice) {
            onRspQryNotice(userData, pNotice, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySettlementInfoConfirm(CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySettlementInfoConfirm) {
            onRspQrySettlementInfoConfirm(userData, pSettlementInfoConfirm, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorPositionCombineDetail(CThostFtdcInvestorPositionCombineDetailField* pInvestorPositionCombineDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorPositionCombineDetail) {
            onRspQryInvestorPositionCombineDetail(userData, pInvestorPositionCombineDetail, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryCFMMCTradingAccountKey(CThostFtdcCFMMCTradingAccountKeyField* pCFMMCTradingAccountKey, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryCFMMCTradingAccountKey) {
            onRspQryCFMMCTradingAccountKey(userData, pCFMMCTradingAccountKey, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryEWarrantOffset(CThostFtdcEWarrantOffsetField* pEWarrantOffset, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryEWarrantOffset) {
            onRspQryEWarrantOffset(userData, pEWarrantOffset, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorProductGroupMargin(CThostFtdcInvestorProductGroupMarginField* pInvestorProductGroupMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorProductGroupMargin) {
            onRspQryInvestorProductGroupMargin(userData, pInvestorProductGroupMargin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryExchangeMarginRate(CThostFtdcExchangeMarginRateField* pExchangeMarginRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryExchangeMarginRate) {
            onRspQryExchangeMarginRate(userData, pExchangeMarginRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryExchangeMarginRateAdjust(CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryExchangeMarginRateAdjust) {
            onRspQryExchangeMarginRateAdjust(userData, pExchangeMarginRateAdjust, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryExchangeRate(CThostFtdcExchangeRateField* pExchangeRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryExchangeRate) {
            onRspQryExchangeRate(userData, pExchangeRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySecAgentACIDMap(CThostFtdcSecAgentACIDMapField* pSecAgentACIDMap, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySecAgentACIDMap) {
            onRspQrySecAgentACIDMap(userData, pSecAgentACIDMap, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryProductExchRate(CThostFtdcProductExchRateField* pProductExchRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryProductExchRate) {
            onRspQryProductExchRate(userData, pProductExchRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryProductGroup(CThostFtdcProductGroupField* pProductGroup, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryProductGroup) {
            onRspQryProductGroup(userData, pProductGroup, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryMMInstrumentCommissionRate(CThostFtdcMMInstrumentCommissionRateField* pMMInstrumentCommissionRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryMMInstrumentCommissionRate) {
            onRspQryMMInstrumentCommissionRate(userData, pMMInstrumentCommissionRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryMMOptionInstrCommRate(CThostFtdcMMOptionInstrCommRateField* pMMOptionInstrCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryMMOptionInstrCommRate) {
            onRspQryMMOptionInstrCommRate(userData, pMMOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInstrumentOrderCommRate(CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInstrumentOrderCommRate) {
            onRspQryInstrumentOrderCommRate(userData, pInstrumentOrderCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySecAgentTradingAccount(CThostFtdcTradingAccountField* pTradingAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySecAgentTradingAccount) {
            onRspQrySecAgentTradingAccount(userData, pTradingAccount, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySecAgentCheckMode(CThostFtdcSecAgentCheckModeField* pSecAgentCheckMode, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySecAgentCheckMode) {
            onRspQrySecAgentCheckMode(userData, pSecAgentCheckMode, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySecAgentTradeInfo(CThostFtdcSecAgentTradeInfoField* pSecAgentTradeInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySecAgentTradeInfo) {
            onRspQrySecAgentTradeInfo(userData, pSecAgentTradeInfo, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryOptionInstrTradeCost(CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryOptionInstrTradeCost) {
            onRspQryOptionInstrTradeCost(userData, pOptionInstrTradeCost, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryOptionInstrCommRate(CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryOptionInstrCommRate) {
            onRspQryOptionInstrCommRate(userData, pOptionInstrCommRate, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryExecOrder(CThostFtdcExecOrderField* pExecOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryExecOrder) {
            onRspQryExecOrder(userData, pExecOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryForQuote(CThostFtdcForQuoteField* pForQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryForQuote) {
            onRspQryForQuote(userData, pForQuote, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryQuote(CThostFtdcQuoteField* pQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryQuote) {
            onRspQryQuote(userData, pQuote, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryOptionSelfClose(CThostFtdcOptionSelfCloseField* pOptionSelfClose, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryOptionSelfClose) {
            onRspQryOptionSelfClose(userData, pOptionSelfClose, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestUnit(CThostFtdcInvestUnitField* pInvestUnit, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestUnit) {
            onRspQryInvestUnit(userData, pInvestUnit, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryCombInstrumentGuard(CThostFtdcCombInstrumentGuardField* pCombInstrumentGuard, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryCombInstrumentGuard) {
            onRspQryCombInstrumentGuard(userData, pCombInstrumentGuard, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryCombAction(CThostFtdcCombActionField* pCombAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryCombAction) {
            onRspQryCombAction(userData, pCombAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTransferSerial(CThostFtdcTransferSerialField* pTransferSerial, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTransferSerial) {
            onRspQryTransferSerial(userData, pTransferSerial, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryAccountregister(CThostFtdcAccountregisterField* pAccountregister, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryAccountregister) {
            onRspQryAccountregister(userData, pAccountregister, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspError(CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspError) {
            onRspError(userData, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRtnOrder(CThostFtdcOrderField* pOrder) override {
        if (onRtnOrder) {
            onRtnOrder(userData, pOrder);
        }
    }

    void OnRtnTrade(CThostFtdcTradeField* pTrade) override {
        if (onRtnTrade) {
            onRtnTrade(userData, pTrade);
        }
    }

    void OnErrRtnOrderInsert(CThostFtdcInputOrderField* pInputOrder, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnOrderInsert) {
            onErrRtnOrderInsert(userData, pInputOrder, pRspInfo);
        }
    }

    void OnErrRtnOrderAction(CThostFtdcOrderActionField* pOrderAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnOrderAction) {
            onErrRtnOrderAction(userData, pOrderAction, pRspInfo);
        }
    }

    void OnRtnInstrumentStatus(CThostFtdcInstrumentStatusField* pInstrumentStatus) override {
        if (onRtnInstrumentStatus) {
            onRtnInstrumentStatus(userData, pInstrumentStatus);
        }
    }

    void OnRtnBulletin(CThostFtdcBulletinField* pBulletin) override {
        if (onRtnBulletin) {
            onRtnBulletin(userData, pBulletin);
        }
    }

    void OnRtnTradingNotice(CThostFtdcTradingNoticeInfoField* pTradingNoticeInfo) override {
        if (onRtnTradingNotice) {
            onRtnTradingNotice(userData, pTradingNoticeInfo);
        }
    }

    void OnRtnErrorConditionalOrder(CThostFtdcErrorConditionalOrderField* pErrorConditionalOrder) override {
        if (onRtnErrorConditionalOrder) {
            onRtnErrorConditionalOrder(userData, pErrorConditionalOrder);
        }
    }

    void OnRtnExecOrder(CThostFtdcExecOrderField* pExecOrder) override {
        if (onRtnExecOrder) {
            onRtnExecOrder(userData, pExecOrder);
        }
    }

    void OnErrRtnExecOrderInsert(CThostFtdcInputExecOrderField* pInputExecOrder, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnExecOrderInsert) {
            onErrRtnExecOrderInsert(userData, pInputExecOrder, pRspInfo);
        }
    }

    void OnErrRtnExecOrderAction(CThostFtdcExecOrderActionField* pExecOrderAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnExecOrderAction) {
            onErrRtnExecOrderAction(userData, pExecOrderAction, pRspInfo);
        }
    }

    void OnErrRtnForQuoteInsert(CThostFtdcInputForQuoteField* pInputForQuote, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnForQuoteInsert) {
            onErrRtnForQuoteInsert(userData, pInputForQuote, pRspInfo);
        }
    }

    void OnRtnQuote(CThostFtdcQuoteField* pQuote) override {
        if (onRtnQuote) {
            onRtnQuote(userData, pQuote);
        }
    }

    void OnErrRtnQuoteInsert(CThostFtdcInputQuoteField* pInputQuote, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnQuoteInsert) {
            onErrRtnQuoteInsert(userData, pInputQuote, pRspInfo);
        }
    }

    void OnErrRtnQuoteAction(CThostFtdcQuoteActionField* pQuoteAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnQuoteAction) {
            onErrRtnQuoteAction(userData, pQuoteAction, pRspInfo);
        }
    }

    void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField* pForQuoteRsp) override {
        if (onRtnForQuoteRsp) {
            onRtnForQuoteRsp(userData, pForQuoteRsp);
        }
    }

    void OnRtnCFMMCTradingAccountToken(CThostFtdcCFMMCTradingAccountTokenField* pCFMMCTradingAccountToken) override {
        if (onRtnCFMMCTradingAccountToken) {
            onRtnCFMMCTradingAccountToken(userData, pCFMMCTradingAccountToken);
        }
    }

    void OnErrRtnBatchOrderAction(CThostFtdcBatchOrderActionField* pBatchOrderAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnBatchOrderAction) {
            onErrRtnBatchOrderAction(userData, pBatchOrderAction, pRspInfo);
        }
    }

    void OnRtnOptionSelfClose(CThostFtdcOptionSelfCloseField* pOptionSelfClose) override {
        if (onRtnOptionSelfClose) {
            onRtnOptionSelfClose(userData, pOptionSelfClose);
        }
    }

    void OnErrRtnOptionSelfCloseInsert(CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnOptionSelfCloseInsert) {
            onErrRtnOptionSelfCloseInsert(userData, pInputOptionSelfClose, pRspInfo);
        }
    }

    void OnErrRtnOptionSelfCloseAction(CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnOptionSelfCloseAction) {
            onErrRtnOptionSelfCloseAction(userData, pOptionSelfCloseAction, pRspInfo);
        }
    }

    void OnRtnCombAction(CThostFtdcCombActionField* pCombAction) override {
        if (onRtnCombAction) {
            onRtnCombAction(userData, pCombAction);
        }
    }

    void OnErrRtnCombActionInsert(CThostFtdcInputCombActionField* pInputCombAction, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnCombActionInsert) {
            onErrRtnCombActionInsert(userData, pInputCombAction, pRspInfo);
        }
    }

    void OnRspQryContractBank(CThostFtdcContractBankField* pContractBank, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryContractBank) {
            onRspQryContractBank(userData, pContractBank, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryParkedOrder(CThostFtdcParkedOrderField* pParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryParkedOrder) {
            onRspQryParkedOrder(userData, pParkedOrder, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryParkedOrderAction(CThostFtdcParkedOrderActionField* pParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryParkedOrderAction) {
            onRspQryParkedOrderAction(userData, pParkedOrderAction, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryTradingNotice(CThostFtdcTradingNoticeField* pTradingNotice, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryTradingNotice) {
            onRspQryTradingNotice(userData, pTradingNotice, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryBrokerTradingParams(CThostFtdcBrokerTradingParamsField* pBrokerTradingParams, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryBrokerTradingParams) {
            onRspQryBrokerTradingParams(userData, pBrokerTradingParams, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryBrokerTradingAlgos(CThostFtdcBrokerTradingAlgosField* pBrokerTradingAlgos, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryBrokerTradingAlgos) {
            onRspQryBrokerTradingAlgos(userData, pBrokerTradingAlgos, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQueryCFMMCTradingAccountToken(CThostFtdcQueryCFMMCTradingAccountTokenField* pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQueryCFMMCTradingAccountToken) {
            onRspQueryCFMMCTradingAccountToken(userData, pQueryCFMMCTradingAccountToken, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRtnFromBankToFutureByBank(CThostFtdcRspTransferField* pRspTransfer) override {
        if (onRtnFromBankToFutureByBank) {
            onRtnFromBankToFutureByBank(userData, pRspTransfer);
        }
    }

    void OnRtnFromFutureToBankByBank(CThostFtdcRspTransferField* pRspTransfer) override {
        if (onRtnFromFutureToBankByBank) {
            onRtnFromFutureToBankByBank(userData, pRspTransfer);
        }
    }

    void OnRtnRepealFromBankToFutureByBank(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromBankToFutureByBank) {
            onRtnRepealFromBankToFutureByBank(userData, pRspRepeal);
        }
    }

    void OnRtnRepealFromFutureToBankByBank(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromFutureToBankByBank) {
            onRtnRepealFromFutureToBankByBank(userData, pRspRepeal);
        }
    }

    void OnRtnFromBankToFutureByFuture(CThostFtdcRspTransferField* pRspTransfer) override {
        if (onRtnFromBankToFutureByFuture) {
            onRtnFromBankToFutureByFuture(userData, pRspTransfer);
        }
    }

    void OnRtnFromFutureToBankByFuture(CThostFtdcRspTransferField* pRspTransfer) override {
        if (onRtnFromFutureToBankByFuture) {
            onRtnFromFutureToBankByFuture(userData, pRspTransfer);
        }
    }

    void OnRtnRepealFromBankToFutureByFutureManual(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromBankToFutureByFutureManual) {
            onRtnRepealFromBankToFutureByFutureManual(userData, pRspRepeal);
        }
    }

    void OnRtnRepealFromFutureToBankByFutureManual(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromFutureToBankByFutureManual) {
            onRtnRepealFromFutureToBankByFutureManual(userData, pRspRepeal);
        }
    }

    void OnRtnQueryBankBalanceByFuture(CThostFtdcNotifyQueryAccountField* pNotifyQueryAccount) override {
        if (onRtnQueryBankBalanceByFuture) {
            onRtnQueryBankBalanceByFuture(userData, pNotifyQueryAccount);
        }
    }

    void OnErrRtnBankToFutureByFuture(CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnBankToFutureByFuture) {
            onErrRtnBankToFutureByFuture(userData, pReqTransfer, pRspInfo);
        }
    }

    void OnErrRtnFutureToBankByFuture(CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnFutureToBankByFuture) {
            onErrRtnFutureToBankByFuture(userData, pReqTransfer, pRspInfo);
        }
    }

    void OnErrRtnRepealBankToFutureByFutureManual(CThostFtdcReqRepealField* pReqRepeal, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnRepealBankToFutureByFutureManual) {
            onErrRtnRepealBankToFutureByFutureManual(userData, pReqRepeal, pRspInfo);
        }
    }

    void OnErrRtnRepealFutureToBankByFutureManual(CThostFtdcReqRepealField* pReqRepeal, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnRepealFutureToBankByFutureManual) {
            onErrRtnRepealFutureToBankByFutureManual(userData, pReqRepeal, pRspInfo);
        }
    }

    void OnErrRtnQueryBankBalanceByFuture(CThostFtdcReqQueryAccountField* pReqQueryAccount, CThostFtdcRspInfoField* pRspInfo) override {
        if (onErrRtnQueryBankBalanceByFuture) {
            onErrRtnQueryBankBalanceByFuture(userData, pReqQueryAccount, pRspInfo);
        }
    }

    void OnRtnRepealFromBankToFutureByFuture(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromBankToFutureByFuture) {
            onRtnRepealFromBankToFutureByFuture(userData, pRspRepeal);
        }
    }

    void OnRtnRepealFromFutureToBankByFuture(CThostFtdcRspRepealField* pRspRepeal) override {
        if (onRtnRepealFromFutureToBankByFuture) {
            onRtnRepealFromFutureToBankByFuture(userData, pRspRepeal);
        }
    }

    void OnRspFromBankToFutureByFuture(CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspFromBankToFutureByFuture) {
            onRspFromBankToFutureByFuture(userData, pReqTransfer, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspFromFutureToBankByFuture(CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspFromFutureToBankByFuture) {
            onRspFromFutureToBankByFuture(userData, pReqTransfer, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQueryBankAccountMoneyByFuture(CThostFtdcReqQueryAccountField* pReqQueryAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQueryBankAccountMoneyByFuture) {
            onRspQueryBankAccountMoneyByFuture(userData, pReqQueryAccount, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRtnOpenAccountByBank(CThostFtdcOpenAccountField* pOpenAccount) override {
        if (onRtnOpenAccountByBank) {
            onRtnOpenAccountByBank(userData, pOpenAccount);
        }
    }

    void OnRtnCancelAccountByBank(CThostFtdcCancelAccountField* pCancelAccount) override {
        if (onRtnCancelAccountByBank) {
            onRtnCancelAccountByBank(userData, pCancelAccount);
        }
    }

    void OnRtnChangeAccountByBank(CThostFtdcChangeAccountField* pChangeAccount) override {
        if (onRtnChangeAccountByBank) {
            onRtnChangeAccountByBank(userData, pChangeAccount);
        }
    }

    void OnRspQryClassifiedInstrument(CThostFtdcInstrumentField* pInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryClassifiedInstrument) {
            onRspQryClassifiedInstrument(userData, pInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryCombPromotionParam(CThostFtdcCombPromotionParamField* pCombPromotionParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryCombPromotionParam) {
            onRspQryCombPromotionParam(userData, pCombPromotionParam, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRiskSettleInvstPosition(CThostFtdcRiskSettleInvstPositionField* pRiskSettleInvstPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRiskSettleInvstPosition) {
            onRspQryRiskSettleInvstPosition(userData, pRiskSettleInvstPosition, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRiskSettleProductStatus(CThostFtdcRiskSettleProductStatusField* pRiskSettleProductStatus, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRiskSettleProductStatus) {
            onRspQryRiskSettleProductStatus(userData, pRiskSettleProductStatus, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMFutureParameter(CThostFtdcSPBMFutureParameterField* pSPBMFutureParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMFutureParameter) {
            onRspQrySPBMFutureParameter(userData, pSPBMFutureParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMOptionParameter(CThostFtdcSPBMOptionParameterField* pSPBMOptionParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMOptionParameter) {
            onRspQrySPBMOptionParameter(userData, pSPBMOptionParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMIntraParameter(CThostFtdcSPBMIntraParameterField* pSPBMIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMIntraParameter) {
            onRspQrySPBMIntraParameter(userData, pSPBMIntraParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMInterParameter(CThostFtdcSPBMInterParameterField* pSPBMInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMInterParameter) {
            onRspQrySPBMInterParameter(userData, pSPBMInterParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMPortfDefinition(CThostFtdcSPBMPortfDefinitionField* pSPBMPortfDefinition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMPortfDefinition) {
            onRspQrySPBMPortfDefinition(userData, pSPBMPortfDefinition, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMInvestorPortfDef(CThostFtdcSPBMInvestorPortfDefField* pSPBMInvestorPortfDef, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMInvestorPortfDef) {
            onRspQrySPBMInvestorPortfDef(userData, pSPBMInvestorPortfDef, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorPortfMarginRatio(CThostFtdcInvestorPortfMarginRatioField* pInvestorPortfMarginRatio, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorPortfMarginRatio) {
            onRspQryInvestorPortfMarginRatio(userData, pInvestorPortfMarginRatio, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorProdSPBMDetail(CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorProdSPBMDetail) {
            onRspQryInvestorProdSPBMDetail(userData, pInvestorProdSPBMDetail, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorCommoditySPMMMargin(CThostFtdcInvestorCommoditySPMMMarginField* pInvestorCommoditySPMMMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorCommoditySPMMMargin) {
            onRspQryInvestorCommoditySPMMMargin(userData, pInvestorCommoditySPMMMargin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorCommodityGroupSPMMMargin(CThostFtdcInvestorCommodityGroupSPMMMarginField* pInvestorCommodityGroupSPMMMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorCommodityGroupSPMMMargin) {
            onRspQryInvestorCommodityGroupSPMMMargin(userData, pInvestorCommodityGroupSPMMMargin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPMMInstParam(CThostFtdcSPMMInstParamField* pSPMMInstParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPMMInstParam) {
            onRspQrySPMMInstParam(userData, pSPMMInstParam, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPMMProductParam(CThostFtdcSPMMProductParamField* pSPMMProductParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPMMProductParam) {
            onRspQrySPMMProductParam(userData, pSPMMProductParam, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQrySPBMAddOnInterParameter(CThostFtdcSPBMAddOnInterParameterField* pSPBMAddOnInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQrySPBMAddOnInterParameter) {
            onRspQrySPBMAddOnInterParameter(userData, pSPBMAddOnInterParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSCombProductInfo(CThostFtdcRCAMSCombProductInfoField* pRCAMSCombProductInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSCombProductInfo) {
            onRspQryRCAMSCombProductInfo(userData, pRCAMSCombProductInfo, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSInstrParameter(CThostFtdcRCAMSInstrParameterField* pRCAMSInstrParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSInstrParameter) {
            onRspQryRCAMSInstrParameter(userData, pRCAMSInstrParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSIntraParameter(CThostFtdcRCAMSIntraParameterField* pRCAMSIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSIntraParameter) {
            onRspQryRCAMSIntraParameter(userData, pRCAMSIntraParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSInterParameter(CThostFtdcRCAMSInterParameterField* pRCAMSInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSInterParameter) {
            onRspQryRCAMSInterParameter(userData, pRCAMSInterParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSShortOptAdjustParam(CThostFtdcRCAMSShortOptAdjustParamField* pRCAMSShortOptAdjustParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSShortOptAdjustParam) {
            onRspQryRCAMSShortOptAdjustParam(userData, pRCAMSShortOptAdjustParam, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRCAMSInvestorCombPosition(CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRCAMSInvestorCombPosition) {
            onRspQryRCAMSInvestorCombPosition(userData, pRCAMSInvestorCombPosition, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorProdRCAMSMargin(CThostFtdcInvestorProdRCAMSMarginField* pInvestorProdRCAMSMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorProdRCAMSMargin) {
            onRspQryInvestorProdRCAMSMargin(userData, pInvestorProdRCAMSMargin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRULEInstrParameter(CThostFtdcRULEInstrParameterField* pRULEInstrParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRULEInstrParameter) {
            onRspQryRULEInstrParameter(userData, pRULEInstrParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRULEIntraParameter(CThostFtdcRULEIntraParameterField* pRULEIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRULEIntraParameter) {
            onRspQryRULEIntraParameter(userData, pRULEIntraParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryRULEInterParameter(CThostFtdcRULEInterParameterField* pRULEInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryRULEInterParameter) {
            onRspQryRULEInterParameter(userData, pRULEInterParameter, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorProdRULEMargin(CThostFtdcInvestorProdRULEMarginField* pInvestorProdRULEMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorProdRULEMargin) {
            onRspQryInvestorProdRULEMargin(userData, pInvestorProdRULEMargin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryInvestorPortfSetting(CThostFtdcInvestorPortfSettingField* pInvestorPortfSetting, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryInvestorPortfSetting) {
            onRspQryInvestorPortfSetting(userData, pInvestorPortfSetting, pRspInfo, nRequestID, bIsLast);
        }
    }

};

// ========== C 接口实现 ==========

extern "C" {

TraderApiHandle TraderCreateFtdcTraderApi(const char* pszFlowPath) {
    return reinterpret_cast<TraderApiHandle>(
        CThostFtdcTraderApi::CreateFtdcTraderApi(pszFlowPath)
    );
}

const char * TraderGetApiVersion(void) {
    return CThostFtdcTraderApi::GetApiVersion();
}

// 跨平台统一登录接口实现（参考 ctpgo 实现）
// macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数
// 此函数内部自动采集系统信息，调用方无需关心平台差异
int TraderReqUserLoginWithSystemInfo(TraderApiHandle handle,
    CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID,
    int systemInfoLen, const char* systemInfo) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
#ifdef __APPLE__
    // macOS 版本：内部自动采集系统信息（参考 ctpgo 实现）
    if (systemInfo == nullptr || systemInfoLen == 0) {
        // 使用 CTP 自带的类型和未 AES 加密的系统信息采集函数
        TThostFtdcClientSystemInfoType sysInfo = {0};
        int len = sizeof(sysInfo);  // CTP_GetSystemInfoUnAesEncode 需要 int&
        CTP_GetSystemInfoUnAesEncode(sysInfo, len);
        // ReqUserLogin 需要 TThostFtdcSystemInfoLenType，进行类型转换
        return api->ReqUserLogin(pReqUserLoginField, nRequestID, static_cast<TThostFtdcSystemInfoLenType>(len), sysInfo);
    }
    return api->ReqUserLogin(pReqUserLoginField, nRequestID, systemInfoLen, 
        const_cast<char*>(systemInfo));
#else
    // Linux/Windows 版本忽略 systemInfo 参数
    (void)systemInfoLen;
    (void)systemInfo;
    return api->ReqUserLogin(pReqUserLoginField, nRequestID);
#endif
}

void TraderRelease(TraderApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->Release();
}

void TraderInit(TraderApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->Init();
}

int TraderJoin(TraderApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->Join();
}

const char * TraderGetTradingDay(TraderApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->GetTradingDay();
}

void TraderGetFrontInfo(TraderApiHandle handle, struct CThostFtdcFrontInfoField* pFrontInfo) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->GetFrontInfo(pFrontInfo);
}

void TraderRegisterFront(TraderApiHandle handle, char* pszFrontAddress) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->RegisterFront(pszFrontAddress);
}

void TraderRegisterNameServer(TraderApiHandle handle, char* pszNsAddress) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->RegisterNameServer(pszNsAddress);
}

void TraderRegisterFensUserInfo(TraderApiHandle handle, struct CThostFtdcFensUserInfoField* pFensUserInfo) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->RegisterFensUserInfo(pFensUserInfo);
}

void TraderSubscribePrivateTopic(TraderApiHandle handle, THOST_TE_RESUME_TYPE nResumeType) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->SubscribePrivateTopic(nResumeType);
}

void TraderSubscribePublicTopic(TraderApiHandle handle, THOST_TE_RESUME_TYPE nResumeType) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    api->SubscribePublicTopic(nResumeType);
}

int TraderReqAuthenticate(TraderApiHandle handle, struct CThostFtdcReqAuthenticateField* pReqAuthenticateField, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqAuthenticate(pReqAuthenticateField, nRequestID);
}

int TraderRegisterUserSystemInfo(TraderApiHandle handle, struct CThostFtdcUserSystemInfoField* pUserSystemInfo) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->RegisterUserSystemInfo(pUserSystemInfo);
}

int TraderSubmitUserSystemInfo(TraderApiHandle handle, struct CThostFtdcUserSystemInfoField* pUserSystemInfo) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->SubmitUserSystemInfo(pUserSystemInfo);
}

int TraderReqUserLogin(TraderApiHandle handle, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID) {
    // 跨平台兼容: macOS 的 ReqUserLogin 需要额外参数
    return TraderReqUserLoginWithSystemInfo(handle, pReqUserLoginField, nRequestID, 0, nullptr);
}

int TraderReqUserLogout(TraderApiHandle handle, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserLogout(pUserLogout, nRequestID);
}

int TraderReqUserPasswordUpdate(TraderApiHandle handle, struct CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserPasswordUpdate(pUserPasswordUpdate, nRequestID);
}

int TraderReqTradingAccountPasswordUpdate(TraderApiHandle handle, struct CThostFtdcTradingAccountPasswordUpdateField* pTradingAccountPasswordUpdate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate, nRequestID);
}

int TraderReqUserAuthMethod(TraderApiHandle handle, struct CThostFtdcReqUserAuthMethodField* pReqUserAuthMethod, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserAuthMethod(pReqUserAuthMethod, nRequestID);
}

int TraderReqGenUserCaptcha(TraderApiHandle handle, struct CThostFtdcReqGenUserCaptchaField* pReqGenUserCaptcha, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqGenUserCaptcha(pReqGenUserCaptcha, nRequestID);
}

int TraderReqGenUserText(TraderApiHandle handle, struct CThostFtdcReqGenUserTextField* pReqGenUserText, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqGenUserText(pReqGenUserText, nRequestID);
}

int TraderReqUserLoginWithCaptcha(TraderApiHandle handle, struct CThostFtdcReqUserLoginWithCaptchaField* pReqUserLoginWithCaptcha, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha, nRequestID);
}

int TraderReqUserLoginWithText(TraderApiHandle handle, struct CThostFtdcReqUserLoginWithTextField* pReqUserLoginWithText, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserLoginWithText(pReqUserLoginWithText, nRequestID);
}

int TraderReqUserLoginWithOTP(TraderApiHandle handle, struct CThostFtdcReqUserLoginWithOTPField* pReqUserLoginWithOTP, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqUserLoginWithOTP(pReqUserLoginWithOTP, nRequestID);
}

int TraderReqOrderInsert(TraderApiHandle handle, struct CThostFtdcInputOrderField* pInputOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqOrderInsert(pInputOrder, nRequestID);
}

int TraderReqParkedOrderInsert(TraderApiHandle handle, struct CThostFtdcParkedOrderField* pParkedOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqParkedOrderInsert(pParkedOrder, nRequestID);
}

int TraderReqParkedOrderAction(TraderApiHandle handle, struct CThostFtdcParkedOrderActionField* pParkedOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqParkedOrderAction(pParkedOrderAction, nRequestID);
}

int TraderReqOrderAction(TraderApiHandle handle, struct CThostFtdcInputOrderActionField* pInputOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqOrderAction(pInputOrderAction, nRequestID);
}

int TraderReqQryMaxOrderVolume(TraderApiHandle handle, struct CThostFtdcQryMaxOrderVolumeField* pQryMaxOrderVolume, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryMaxOrderVolume(pQryMaxOrderVolume, nRequestID);
}

int TraderReqSettlementInfoConfirm(TraderApiHandle handle, struct CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqSettlementInfoConfirm(pSettlementInfoConfirm, nRequestID);
}

int TraderReqRemoveParkedOrder(TraderApiHandle handle, struct CThostFtdcRemoveParkedOrderField* pRemoveParkedOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqRemoveParkedOrder(pRemoveParkedOrder, nRequestID);
}

int TraderReqRemoveParkedOrderAction(TraderApiHandle handle, struct CThostFtdcRemoveParkedOrderActionField* pRemoveParkedOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqRemoveParkedOrderAction(pRemoveParkedOrderAction, nRequestID);
}

int TraderReqExecOrderInsert(TraderApiHandle handle, struct CThostFtdcInputExecOrderField* pInputExecOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqExecOrderInsert(pInputExecOrder, nRequestID);
}

int TraderReqExecOrderAction(TraderApiHandle handle, struct CThostFtdcInputExecOrderActionField* pInputExecOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqExecOrderAction(pInputExecOrderAction, nRequestID);
}

int TraderReqForQuoteInsert(TraderApiHandle handle, struct CThostFtdcInputForQuoteField* pInputForQuote, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqForQuoteInsert(pInputForQuote, nRequestID);
}

int TraderReqQuoteInsert(TraderApiHandle handle, struct CThostFtdcInputQuoteField* pInputQuote, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQuoteInsert(pInputQuote, nRequestID);
}

int TraderReqQuoteAction(TraderApiHandle handle, struct CThostFtdcInputQuoteActionField* pInputQuoteAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQuoteAction(pInputQuoteAction, nRequestID);
}

int TraderReqBatchOrderAction(TraderApiHandle handle, struct CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqBatchOrderAction(pInputBatchOrderAction, nRequestID);
}

int TraderReqOptionSelfCloseInsert(TraderApiHandle handle, struct CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqOptionSelfCloseInsert(pInputOptionSelfClose, nRequestID);
}

int TraderReqOptionSelfCloseAction(TraderApiHandle handle, struct CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqOptionSelfCloseAction(pInputOptionSelfCloseAction, nRequestID);
}

int TraderReqCombActionInsert(TraderApiHandle handle, struct CThostFtdcInputCombActionField* pInputCombAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqCombActionInsert(pInputCombAction, nRequestID);
}

int TraderReqQryOrder(TraderApiHandle handle, struct CThostFtdcQryOrderField* pQryOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryOrder(pQryOrder, nRequestID);
}

int TraderReqQryTrade(TraderApiHandle handle, struct CThostFtdcQryTradeField* pQryTrade, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTrade(pQryTrade, nRequestID);
}

int TraderReqQryInvestorPosition(TraderApiHandle handle, struct CThostFtdcQryInvestorPositionField* pQryInvestorPosition, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorPosition(pQryInvestorPosition, nRequestID);
}

int TraderReqQryTradingAccount(TraderApiHandle handle, struct CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTradingAccount(pQryTradingAccount, nRequestID);
}

int TraderReqQryInvestor(TraderApiHandle handle, struct CThostFtdcQryInvestorField* pQryInvestor, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestor(pQryInvestor, nRequestID);
}

int TraderReqQryTradingCode(TraderApiHandle handle, struct CThostFtdcQryTradingCodeField* pQryTradingCode, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTradingCode(pQryTradingCode, nRequestID);
}

int TraderReqQryInstrumentMarginRate(TraderApiHandle handle, struct CThostFtdcQryInstrumentMarginRateField* pQryInstrumentMarginRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInstrumentMarginRate(pQryInstrumentMarginRate, nRequestID);
}

int TraderReqQryInstrumentCommissionRate(TraderApiHandle handle, struct CThostFtdcQryInstrumentCommissionRateField* pQryInstrumentCommissionRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate, nRequestID);
}

int TraderReqQryExchange(TraderApiHandle handle, struct CThostFtdcQryExchangeField* pQryExchange, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryExchange(pQryExchange, nRequestID);
}

int TraderReqQryProduct(TraderApiHandle handle, struct CThostFtdcQryProductField* pQryProduct, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryProduct(pQryProduct, nRequestID);
}

int TraderReqQryInstrument(TraderApiHandle handle, struct CThostFtdcQryInstrumentField* pQryInstrument, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInstrument(pQryInstrument, nRequestID);
}

int TraderReqQryDepthMarketData(TraderApiHandle handle, struct CThostFtdcQryDepthMarketDataField* pQryDepthMarketData, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryDepthMarketData(pQryDepthMarketData, nRequestID);
}

int TraderReqQryTraderOffer(TraderApiHandle handle, struct CThostFtdcQryTraderOfferField* pQryTraderOffer, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTraderOffer(pQryTraderOffer, nRequestID);
}

int TraderReqQrySettlementInfo(TraderApiHandle handle, struct CThostFtdcQrySettlementInfoField* pQrySettlementInfo, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySettlementInfo(pQrySettlementInfo, nRequestID);
}

int TraderReqQryTransferBank(TraderApiHandle handle, struct CThostFtdcQryTransferBankField* pQryTransferBank, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTransferBank(pQryTransferBank, nRequestID);
}

int TraderReqQryInvestorPositionDetail(TraderApiHandle handle, struct CThostFtdcQryInvestorPositionDetailField* pQryInvestorPositionDetail, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorPositionDetail(pQryInvestorPositionDetail, nRequestID);
}

int TraderReqQryNotice(TraderApiHandle handle, struct CThostFtdcQryNoticeField* pQryNotice, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryNotice(pQryNotice, nRequestID);
}

int TraderReqQrySettlementInfoConfirm(TraderApiHandle handle, struct CThostFtdcQrySettlementInfoConfirmField* pQrySettlementInfoConfirm, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm, nRequestID);
}

int TraderReqQryInvestorPositionCombineDetail(TraderApiHandle handle, struct CThostFtdcQryInvestorPositionCombineDetailField* pQryInvestorPositionCombineDetail, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail, nRequestID);
}

int TraderReqQryCFMMCTradingAccountKey(TraderApiHandle handle, struct CThostFtdcQryCFMMCTradingAccountKeyField* pQryCFMMCTradingAccountKey, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey, nRequestID);
}

int TraderReqQryEWarrantOffset(TraderApiHandle handle, struct CThostFtdcQryEWarrantOffsetField* pQryEWarrantOffset, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryEWarrantOffset(pQryEWarrantOffset, nRequestID);
}

int TraderReqQryInvestorProductGroupMargin(TraderApiHandle handle, struct CThostFtdcQryInvestorProductGroupMarginField* pQryInvestorProductGroupMargin, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin, nRequestID);
}

int TraderReqQryExchangeMarginRate(TraderApiHandle handle, struct CThostFtdcQryExchangeMarginRateField* pQryExchangeMarginRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryExchangeMarginRate(pQryExchangeMarginRate, nRequestID);
}

int TraderReqQryExchangeMarginRateAdjust(TraderApiHandle handle, struct CThostFtdcQryExchangeMarginRateAdjustField* pQryExchangeMarginRateAdjust, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust, nRequestID);
}

int TraderReqQryExchangeRate(TraderApiHandle handle, struct CThostFtdcQryExchangeRateField* pQryExchangeRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryExchangeRate(pQryExchangeRate, nRequestID);
}

int TraderReqQrySecAgentACIDMap(TraderApiHandle handle, struct CThostFtdcQrySecAgentACIDMapField* pQrySecAgentACIDMap, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySecAgentACIDMap(pQrySecAgentACIDMap, nRequestID);
}

int TraderReqQryProductExchRate(TraderApiHandle handle, struct CThostFtdcQryProductExchRateField* pQryProductExchRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryProductExchRate(pQryProductExchRate, nRequestID);
}

int TraderReqQryProductGroup(TraderApiHandle handle, struct CThostFtdcQryProductGroupField* pQryProductGroup, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryProductGroup(pQryProductGroup, nRequestID);
}

int TraderReqQryMMInstrumentCommissionRate(TraderApiHandle handle, struct CThostFtdcQryMMInstrumentCommissionRateField* pQryMMInstrumentCommissionRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate, nRequestID);
}

int TraderReqQryMMOptionInstrCommRate(TraderApiHandle handle, struct CThostFtdcQryMMOptionInstrCommRateField* pQryMMOptionInstrCommRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate, nRequestID);
}

int TraderReqQryInstrumentOrderCommRate(TraderApiHandle handle, struct CThostFtdcQryInstrumentOrderCommRateField* pQryInstrumentOrderCommRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate, nRequestID);
}

int TraderReqQrySecAgentTradingAccount(TraderApiHandle handle, struct CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySecAgentTradingAccount(pQryTradingAccount, nRequestID);
}

int TraderReqQrySecAgentCheckMode(TraderApiHandle handle, struct CThostFtdcQrySecAgentCheckModeField* pQrySecAgentCheckMode, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySecAgentCheckMode(pQrySecAgentCheckMode, nRequestID);
}

int TraderReqQrySecAgentTradeInfo(TraderApiHandle handle, struct CThostFtdcQrySecAgentTradeInfoField* pQrySecAgentTradeInfo, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo, nRequestID);
}

int TraderReqQryOptionInstrTradeCost(TraderApiHandle handle, struct CThostFtdcQryOptionInstrTradeCostField* pQryOptionInstrTradeCost, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost, nRequestID);
}

int TraderReqQryOptionInstrCommRate(TraderApiHandle handle, struct CThostFtdcQryOptionInstrCommRateField* pQryOptionInstrCommRate, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryOptionInstrCommRate(pQryOptionInstrCommRate, nRequestID);
}

int TraderReqQryExecOrder(TraderApiHandle handle, struct CThostFtdcQryExecOrderField* pQryExecOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryExecOrder(pQryExecOrder, nRequestID);
}

int TraderReqQryForQuote(TraderApiHandle handle, struct CThostFtdcQryForQuoteField* pQryForQuote, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryForQuote(pQryForQuote, nRequestID);
}

int TraderReqQryQuote(TraderApiHandle handle, struct CThostFtdcQryQuoteField* pQryQuote, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryQuote(pQryQuote, nRequestID);
}

int TraderReqQryOptionSelfClose(TraderApiHandle handle, struct CThostFtdcQryOptionSelfCloseField* pQryOptionSelfClose, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryOptionSelfClose(pQryOptionSelfClose, nRequestID);
}

int TraderReqQryInvestUnit(TraderApiHandle handle, struct CThostFtdcQryInvestUnitField* pQryInvestUnit, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestUnit(pQryInvestUnit, nRequestID);
}

int TraderReqQryCombInstrumentGuard(TraderApiHandle handle, struct CThostFtdcQryCombInstrumentGuardField* pQryCombInstrumentGuard, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryCombInstrumentGuard(pQryCombInstrumentGuard, nRequestID);
}

int TraderReqQryCombAction(TraderApiHandle handle, struct CThostFtdcQryCombActionField* pQryCombAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryCombAction(pQryCombAction, nRequestID);
}

int TraderReqQryTransferSerial(TraderApiHandle handle, struct CThostFtdcQryTransferSerialField* pQryTransferSerial, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTransferSerial(pQryTransferSerial, nRequestID);
}

int TraderReqQryAccountregister(TraderApiHandle handle, struct CThostFtdcQryAccountregisterField* pQryAccountregister, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryAccountregister(pQryAccountregister, nRequestID);
}

int TraderReqQryContractBank(TraderApiHandle handle, struct CThostFtdcQryContractBankField* pQryContractBank, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryContractBank(pQryContractBank, nRequestID);
}

int TraderReqQryParkedOrder(TraderApiHandle handle, struct CThostFtdcQryParkedOrderField* pQryParkedOrder, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryParkedOrder(pQryParkedOrder, nRequestID);
}

int TraderReqQryParkedOrderAction(TraderApiHandle handle, struct CThostFtdcQryParkedOrderActionField* pQryParkedOrderAction, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryParkedOrderAction(pQryParkedOrderAction, nRequestID);
}

int TraderReqQryTradingNotice(TraderApiHandle handle, struct CThostFtdcQryTradingNoticeField* pQryTradingNotice, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryTradingNotice(pQryTradingNotice, nRequestID);
}

int TraderReqQryBrokerTradingParams(TraderApiHandle handle, struct CThostFtdcQryBrokerTradingParamsField* pQryBrokerTradingParams, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryBrokerTradingParams(pQryBrokerTradingParams, nRequestID);
}

int TraderReqQryBrokerTradingAlgos(TraderApiHandle handle, struct CThostFtdcQryBrokerTradingAlgosField* pQryBrokerTradingAlgos, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos, nRequestID);
}

int TraderReqQueryCFMMCTradingAccountToken(TraderApiHandle handle, struct CThostFtdcQueryCFMMCTradingAccountTokenField* pQueryCFMMCTradingAccountToken, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken, nRequestID);
}

int TraderReqFromBankToFutureByFuture(TraderApiHandle handle, struct CThostFtdcReqTransferField* pReqTransfer, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqFromBankToFutureByFuture(pReqTransfer, nRequestID);
}

int TraderReqFromFutureToBankByFuture(TraderApiHandle handle, struct CThostFtdcReqTransferField* pReqTransfer, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqFromFutureToBankByFuture(pReqTransfer, nRequestID);
}

int TraderReqQueryBankAccountMoneyByFuture(TraderApiHandle handle, struct CThostFtdcReqQueryAccountField* pReqQueryAccount, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQueryBankAccountMoneyByFuture(pReqQueryAccount, nRequestID);
}

int TraderReqQryClassifiedInstrument(TraderApiHandle handle, struct CThostFtdcQryClassifiedInstrumentField* pQryClassifiedInstrument, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryClassifiedInstrument(pQryClassifiedInstrument, nRequestID);
}

int TraderReqQryCombPromotionParam(TraderApiHandle handle, struct CThostFtdcQryCombPromotionParamField* pQryCombPromotionParam, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryCombPromotionParam(pQryCombPromotionParam, nRequestID);
}

int TraderReqQryRiskSettleInvstPosition(TraderApiHandle handle, struct CThostFtdcQryRiskSettleInvstPositionField* pQryRiskSettleInvstPosition, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition, nRequestID);
}

int TraderReqQryRiskSettleProductStatus(TraderApiHandle handle, struct CThostFtdcQryRiskSettleProductStatusField* pQryRiskSettleProductStatus, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus, nRequestID);
}

int TraderReqQrySPBMFutureParameter(TraderApiHandle handle, struct CThostFtdcQrySPBMFutureParameterField* pQrySPBMFutureParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMFutureParameter(pQrySPBMFutureParameter, nRequestID);
}

int TraderReqQrySPBMOptionParameter(TraderApiHandle handle, struct CThostFtdcQrySPBMOptionParameterField* pQrySPBMOptionParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMOptionParameter(pQrySPBMOptionParameter, nRequestID);
}

int TraderReqQrySPBMIntraParameter(TraderApiHandle handle, struct CThostFtdcQrySPBMIntraParameterField* pQrySPBMIntraParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMIntraParameter(pQrySPBMIntraParameter, nRequestID);
}

int TraderReqQrySPBMInterParameter(TraderApiHandle handle, struct CThostFtdcQrySPBMInterParameterField* pQrySPBMInterParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMInterParameter(pQrySPBMInterParameter, nRequestID);
}

int TraderReqQrySPBMPortfDefinition(TraderApiHandle handle, struct CThostFtdcQrySPBMPortfDefinitionField* pQrySPBMPortfDefinition, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMPortfDefinition(pQrySPBMPortfDefinition, nRequestID);
}

int TraderReqQrySPBMInvestorPortfDef(TraderApiHandle handle, struct CThostFtdcQrySPBMInvestorPortfDefField* pQrySPBMInvestorPortfDef, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMInvestorPortfDef(pQrySPBMInvestorPortfDef, nRequestID);
}

int TraderReqQryInvestorPortfMarginRatio(TraderApiHandle handle, struct CThostFtdcQryInvestorPortfMarginRatioField* pQryInvestorPortfMarginRatio, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorPortfMarginRatio(pQryInvestorPortfMarginRatio, nRequestID);
}

int TraderReqQryInvestorProdSPBMDetail(TraderApiHandle handle, struct CThostFtdcQryInvestorProdSPBMDetailField* pQryInvestorProdSPBMDetail, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorProdSPBMDetail(pQryInvestorProdSPBMDetail, nRequestID);
}

int TraderReqQryInvestorCommoditySPMMMargin(TraderApiHandle handle, struct CThostFtdcQryInvestorCommoditySPMMMarginField* pQryInvestorCommoditySPMMMargin, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorCommoditySPMMMargin(pQryInvestorCommoditySPMMMargin, nRequestID);
}

int TraderReqQryInvestorCommodityGroupSPMMMargin(TraderApiHandle handle, struct CThostFtdcQryInvestorCommodityGroupSPMMMarginField* pQryInvestorCommodityGroupSPMMMargin, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorCommodityGroupSPMMMargin(pQryInvestorCommodityGroupSPMMMargin, nRequestID);
}

int TraderReqQrySPMMInstParam(TraderApiHandle handle, struct CThostFtdcQrySPMMInstParamField* pQrySPMMInstParam, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPMMInstParam(pQrySPMMInstParam, nRequestID);
}

int TraderReqQrySPMMProductParam(TraderApiHandle handle, struct CThostFtdcQrySPMMProductParamField* pQrySPMMProductParam, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPMMProductParam(pQrySPMMProductParam, nRequestID);
}

int TraderReqQrySPBMAddOnInterParameter(TraderApiHandle handle, struct CThostFtdcQrySPBMAddOnInterParameterField* pQrySPBMAddOnInterParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQrySPBMAddOnInterParameter(pQrySPBMAddOnInterParameter, nRequestID);
}

int TraderReqQryRCAMSCombProductInfo(TraderApiHandle handle, struct CThostFtdcQryRCAMSCombProductInfoField* pQryRCAMSCombProductInfo, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSCombProductInfo(pQryRCAMSCombProductInfo, nRequestID);
}

int TraderReqQryRCAMSInstrParameter(TraderApiHandle handle, struct CThostFtdcQryRCAMSInstrParameterField* pQryRCAMSInstrParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSInstrParameter(pQryRCAMSInstrParameter, nRequestID);
}

int TraderReqQryRCAMSIntraParameter(TraderApiHandle handle, struct CThostFtdcQryRCAMSIntraParameterField* pQryRCAMSIntraParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSIntraParameter(pQryRCAMSIntraParameter, nRequestID);
}

int TraderReqQryRCAMSInterParameter(TraderApiHandle handle, struct CThostFtdcQryRCAMSInterParameterField* pQryRCAMSInterParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSInterParameter(pQryRCAMSInterParameter, nRequestID);
}

int TraderReqQryRCAMSShortOptAdjustParam(TraderApiHandle handle, struct CThostFtdcQryRCAMSShortOptAdjustParamField* pQryRCAMSShortOptAdjustParam, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSShortOptAdjustParam(pQryRCAMSShortOptAdjustParam, nRequestID);
}

int TraderReqQryRCAMSInvestorCombPosition(TraderApiHandle handle, struct CThostFtdcQryRCAMSInvestorCombPositionField* pQryRCAMSInvestorCombPosition, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRCAMSInvestorCombPosition(pQryRCAMSInvestorCombPosition, nRequestID);
}

int TraderReqQryInvestorProdRCAMSMargin(TraderApiHandle handle, struct CThostFtdcQryInvestorProdRCAMSMarginField* pQryInvestorProdRCAMSMargin, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorProdRCAMSMargin(pQryInvestorProdRCAMSMargin, nRequestID);
}

int TraderReqQryRULEInstrParameter(TraderApiHandle handle, struct CThostFtdcQryRULEInstrParameterField* pQryRULEInstrParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRULEInstrParameter(pQryRULEInstrParameter, nRequestID);
}

int TraderReqQryRULEIntraParameter(TraderApiHandle handle, struct CThostFtdcQryRULEIntraParameterField* pQryRULEIntraParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRULEIntraParameter(pQryRULEIntraParameter, nRequestID);
}

int TraderReqQryRULEInterParameter(TraderApiHandle handle, struct CThostFtdcQryRULEInterParameterField* pQryRULEInterParameter, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryRULEInterParameter(pQryRULEInterParameter, nRequestID);
}

int TraderReqQryInvestorProdRULEMargin(TraderApiHandle handle, struct CThostFtdcQryInvestorProdRULEMarginField* pQryInvestorProdRULEMargin, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorProdRULEMargin(pQryInvestorProdRULEMargin, nRequestID);
}

int TraderReqQryInvestorPortfSetting(TraderApiHandle handle, struct CThostFtdcQryInvestorPortfSettingField* pQryInvestorPortfSetting, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcTraderApi*>(handle);
    return api->ReqQryInvestorPortfSetting(pQryInvestorPortfSetting, nRequestID);
}

// SPI 创建与销毁
TraderSpiHandle TraderSpiCreate(void* userData) {
    auto* spi = new TraderSpiWrapper();
    spi->userData = userData;
    return reinterpret_cast<TraderSpiHandle>(spi);
}

void TraderSpiDestroy(TraderSpiHandle spi) {
    delete reinterpret_cast<TraderSpiWrapper*>(spi);
}

void TraderRegisterSpi(TraderApiHandle api, TraderSpiHandle spi) {
    auto* apiPtr = reinterpret_cast<CThostFtdcTraderApi*>(api);
    auto* spiPtr = reinterpret_cast<TraderSpiWrapper*>(spi);
    apiPtr->RegisterSpi(spiPtr);
}

void TraderSpiSetCallbacks(TraderSpiHandle spi, const TraderSpiCallbacks* callbacks) {
    auto* spiPtr = reinterpret_cast<TraderSpiWrapper*>(spi);
    spiPtr->userData = callbacks->userData;
    spiPtr->onFrontConnected = callbacks->onFrontConnected;
    spiPtr->onFrontDisconnected = callbacks->onFrontDisconnected;
    spiPtr->onHeartBeatWarning = callbacks->onHeartBeatWarning;
    spiPtr->onRspAuthenticate = callbacks->onRspAuthenticate;
    spiPtr->onRspUserLogin = callbacks->onRspUserLogin;
    spiPtr->onRspUserLogout = callbacks->onRspUserLogout;
    spiPtr->onRspUserPasswordUpdate = callbacks->onRspUserPasswordUpdate;
    spiPtr->onRspTradingAccountPasswordUpdate = callbacks->onRspTradingAccountPasswordUpdate;
    spiPtr->onRspUserAuthMethod = callbacks->onRspUserAuthMethod;
    spiPtr->onRspGenUserCaptcha = callbacks->onRspGenUserCaptcha;
    spiPtr->onRspGenUserText = callbacks->onRspGenUserText;
    spiPtr->onRspOrderInsert = callbacks->onRspOrderInsert;
    spiPtr->onRspParkedOrderInsert = callbacks->onRspParkedOrderInsert;
    spiPtr->onRspParkedOrderAction = callbacks->onRspParkedOrderAction;
    spiPtr->onRspOrderAction = callbacks->onRspOrderAction;
    spiPtr->onRspQryMaxOrderVolume = callbacks->onRspQryMaxOrderVolume;
    spiPtr->onRspSettlementInfoConfirm = callbacks->onRspSettlementInfoConfirm;
    spiPtr->onRspRemoveParkedOrder = callbacks->onRspRemoveParkedOrder;
    spiPtr->onRspRemoveParkedOrderAction = callbacks->onRspRemoveParkedOrderAction;
    spiPtr->onRspExecOrderInsert = callbacks->onRspExecOrderInsert;
    spiPtr->onRspExecOrderAction = callbacks->onRspExecOrderAction;
    spiPtr->onRspForQuoteInsert = callbacks->onRspForQuoteInsert;
    spiPtr->onRspQuoteInsert = callbacks->onRspQuoteInsert;
    spiPtr->onRspQuoteAction = callbacks->onRspQuoteAction;
    spiPtr->onRspBatchOrderAction = callbacks->onRspBatchOrderAction;
    spiPtr->onRspOptionSelfCloseInsert = callbacks->onRspOptionSelfCloseInsert;
    spiPtr->onRspOptionSelfCloseAction = callbacks->onRspOptionSelfCloseAction;
    spiPtr->onRspCombActionInsert = callbacks->onRspCombActionInsert;
    spiPtr->onRspQryOrder = callbacks->onRspQryOrder;
    spiPtr->onRspQryTrade = callbacks->onRspQryTrade;
    spiPtr->onRspQryInvestorPosition = callbacks->onRspQryInvestorPosition;
    spiPtr->onRspQryTradingAccount = callbacks->onRspQryTradingAccount;
    spiPtr->onRspQryInvestor = callbacks->onRspQryInvestor;
    spiPtr->onRspQryTradingCode = callbacks->onRspQryTradingCode;
    spiPtr->onRspQryInstrumentMarginRate = callbacks->onRspQryInstrumentMarginRate;
    spiPtr->onRspQryInstrumentCommissionRate = callbacks->onRspQryInstrumentCommissionRate;
    spiPtr->onRspQryExchange = callbacks->onRspQryExchange;
    spiPtr->onRspQryProduct = callbacks->onRspQryProduct;
    spiPtr->onRspQryInstrument = callbacks->onRspQryInstrument;
    spiPtr->onRspQryDepthMarketData = callbacks->onRspQryDepthMarketData;
    spiPtr->onRspQryTraderOffer = callbacks->onRspQryTraderOffer;
    spiPtr->onRspQrySettlementInfo = callbacks->onRspQrySettlementInfo;
    spiPtr->onRspQryTransferBank = callbacks->onRspQryTransferBank;
    spiPtr->onRspQryInvestorPositionDetail = callbacks->onRspQryInvestorPositionDetail;
    spiPtr->onRspQryNotice = callbacks->onRspQryNotice;
    spiPtr->onRspQrySettlementInfoConfirm = callbacks->onRspQrySettlementInfoConfirm;
    spiPtr->onRspQryInvestorPositionCombineDetail = callbacks->onRspQryInvestorPositionCombineDetail;
    spiPtr->onRspQryCFMMCTradingAccountKey = callbacks->onRspQryCFMMCTradingAccountKey;
    spiPtr->onRspQryEWarrantOffset = callbacks->onRspQryEWarrantOffset;
    spiPtr->onRspQryInvestorProductGroupMargin = callbacks->onRspQryInvestorProductGroupMargin;
    spiPtr->onRspQryExchangeMarginRate = callbacks->onRspQryExchangeMarginRate;
    spiPtr->onRspQryExchangeMarginRateAdjust = callbacks->onRspQryExchangeMarginRateAdjust;
    spiPtr->onRspQryExchangeRate = callbacks->onRspQryExchangeRate;
    spiPtr->onRspQrySecAgentACIDMap = callbacks->onRspQrySecAgentACIDMap;
    spiPtr->onRspQryProductExchRate = callbacks->onRspQryProductExchRate;
    spiPtr->onRspQryProductGroup = callbacks->onRspQryProductGroup;
    spiPtr->onRspQryMMInstrumentCommissionRate = callbacks->onRspQryMMInstrumentCommissionRate;
    spiPtr->onRspQryMMOptionInstrCommRate = callbacks->onRspQryMMOptionInstrCommRate;
    spiPtr->onRspQryInstrumentOrderCommRate = callbacks->onRspQryInstrumentOrderCommRate;
    spiPtr->onRspQrySecAgentTradingAccount = callbacks->onRspQrySecAgentTradingAccount;
    spiPtr->onRspQrySecAgentCheckMode = callbacks->onRspQrySecAgentCheckMode;
    spiPtr->onRspQrySecAgentTradeInfo = callbacks->onRspQrySecAgentTradeInfo;
    spiPtr->onRspQryOptionInstrTradeCost = callbacks->onRspQryOptionInstrTradeCost;
    spiPtr->onRspQryOptionInstrCommRate = callbacks->onRspQryOptionInstrCommRate;
    spiPtr->onRspQryExecOrder = callbacks->onRspQryExecOrder;
    spiPtr->onRspQryForQuote = callbacks->onRspQryForQuote;
    spiPtr->onRspQryQuote = callbacks->onRspQryQuote;
    spiPtr->onRspQryOptionSelfClose = callbacks->onRspQryOptionSelfClose;
    spiPtr->onRspQryInvestUnit = callbacks->onRspQryInvestUnit;
    spiPtr->onRspQryCombInstrumentGuard = callbacks->onRspQryCombInstrumentGuard;
    spiPtr->onRspQryCombAction = callbacks->onRspQryCombAction;
    spiPtr->onRspQryTransferSerial = callbacks->onRspQryTransferSerial;
    spiPtr->onRspQryAccountregister = callbacks->onRspQryAccountregister;
    spiPtr->onRspError = callbacks->onRspError;
    spiPtr->onRtnOrder = callbacks->onRtnOrder;
    spiPtr->onRtnTrade = callbacks->onRtnTrade;
    spiPtr->onErrRtnOrderInsert = callbacks->onErrRtnOrderInsert;
    spiPtr->onErrRtnOrderAction = callbacks->onErrRtnOrderAction;
    spiPtr->onRtnInstrumentStatus = callbacks->onRtnInstrumentStatus;
    spiPtr->onRtnBulletin = callbacks->onRtnBulletin;
    spiPtr->onRtnTradingNotice = callbacks->onRtnTradingNotice;
    spiPtr->onRtnErrorConditionalOrder = callbacks->onRtnErrorConditionalOrder;
    spiPtr->onRtnExecOrder = callbacks->onRtnExecOrder;
    spiPtr->onErrRtnExecOrderInsert = callbacks->onErrRtnExecOrderInsert;
    spiPtr->onErrRtnExecOrderAction = callbacks->onErrRtnExecOrderAction;
    spiPtr->onErrRtnForQuoteInsert = callbacks->onErrRtnForQuoteInsert;
    spiPtr->onRtnQuote = callbacks->onRtnQuote;
    spiPtr->onErrRtnQuoteInsert = callbacks->onErrRtnQuoteInsert;
    spiPtr->onErrRtnQuoteAction = callbacks->onErrRtnQuoteAction;
    spiPtr->onRtnForQuoteRsp = callbacks->onRtnForQuoteRsp;
    spiPtr->onRtnCFMMCTradingAccountToken = callbacks->onRtnCFMMCTradingAccountToken;
    spiPtr->onErrRtnBatchOrderAction = callbacks->onErrRtnBatchOrderAction;
    spiPtr->onRtnOptionSelfClose = callbacks->onRtnOptionSelfClose;
    spiPtr->onErrRtnOptionSelfCloseInsert = callbacks->onErrRtnOptionSelfCloseInsert;
    spiPtr->onErrRtnOptionSelfCloseAction = callbacks->onErrRtnOptionSelfCloseAction;
    spiPtr->onRtnCombAction = callbacks->onRtnCombAction;
    spiPtr->onErrRtnCombActionInsert = callbacks->onErrRtnCombActionInsert;
    spiPtr->onRspQryContractBank = callbacks->onRspQryContractBank;
    spiPtr->onRspQryParkedOrder = callbacks->onRspQryParkedOrder;
    spiPtr->onRspQryParkedOrderAction = callbacks->onRspQryParkedOrderAction;
    spiPtr->onRspQryTradingNotice = callbacks->onRspQryTradingNotice;
    spiPtr->onRspQryBrokerTradingParams = callbacks->onRspQryBrokerTradingParams;
    spiPtr->onRspQryBrokerTradingAlgos = callbacks->onRspQryBrokerTradingAlgos;
    spiPtr->onRspQueryCFMMCTradingAccountToken = callbacks->onRspQueryCFMMCTradingAccountToken;
    spiPtr->onRtnFromBankToFutureByBank = callbacks->onRtnFromBankToFutureByBank;
    spiPtr->onRtnFromFutureToBankByBank = callbacks->onRtnFromFutureToBankByBank;
    spiPtr->onRtnRepealFromBankToFutureByBank = callbacks->onRtnRepealFromBankToFutureByBank;
    spiPtr->onRtnRepealFromFutureToBankByBank = callbacks->onRtnRepealFromFutureToBankByBank;
    spiPtr->onRtnFromBankToFutureByFuture = callbacks->onRtnFromBankToFutureByFuture;
    spiPtr->onRtnFromFutureToBankByFuture = callbacks->onRtnFromFutureToBankByFuture;
    spiPtr->onRtnRepealFromBankToFutureByFutureManual = callbacks->onRtnRepealFromBankToFutureByFutureManual;
    spiPtr->onRtnRepealFromFutureToBankByFutureManual = callbacks->onRtnRepealFromFutureToBankByFutureManual;
    spiPtr->onRtnQueryBankBalanceByFuture = callbacks->onRtnQueryBankBalanceByFuture;
    spiPtr->onErrRtnBankToFutureByFuture = callbacks->onErrRtnBankToFutureByFuture;
    spiPtr->onErrRtnFutureToBankByFuture = callbacks->onErrRtnFutureToBankByFuture;
    spiPtr->onErrRtnRepealBankToFutureByFutureManual = callbacks->onErrRtnRepealBankToFutureByFutureManual;
    spiPtr->onErrRtnRepealFutureToBankByFutureManual = callbacks->onErrRtnRepealFutureToBankByFutureManual;
    spiPtr->onErrRtnQueryBankBalanceByFuture = callbacks->onErrRtnQueryBankBalanceByFuture;
    spiPtr->onRtnRepealFromBankToFutureByFuture = callbacks->onRtnRepealFromBankToFutureByFuture;
    spiPtr->onRtnRepealFromFutureToBankByFuture = callbacks->onRtnRepealFromFutureToBankByFuture;
    spiPtr->onRspFromBankToFutureByFuture = callbacks->onRspFromBankToFutureByFuture;
    spiPtr->onRspFromFutureToBankByFuture = callbacks->onRspFromFutureToBankByFuture;
    spiPtr->onRspQueryBankAccountMoneyByFuture = callbacks->onRspQueryBankAccountMoneyByFuture;
    spiPtr->onRtnOpenAccountByBank = callbacks->onRtnOpenAccountByBank;
    spiPtr->onRtnCancelAccountByBank = callbacks->onRtnCancelAccountByBank;
    spiPtr->onRtnChangeAccountByBank = callbacks->onRtnChangeAccountByBank;
    spiPtr->onRspQryClassifiedInstrument = callbacks->onRspQryClassifiedInstrument;
    spiPtr->onRspQryCombPromotionParam = callbacks->onRspQryCombPromotionParam;
    spiPtr->onRspQryRiskSettleInvstPosition = callbacks->onRspQryRiskSettleInvstPosition;
    spiPtr->onRspQryRiskSettleProductStatus = callbacks->onRspQryRiskSettleProductStatus;
    spiPtr->onRspQrySPBMFutureParameter = callbacks->onRspQrySPBMFutureParameter;
    spiPtr->onRspQrySPBMOptionParameter = callbacks->onRspQrySPBMOptionParameter;
    spiPtr->onRspQrySPBMIntraParameter = callbacks->onRspQrySPBMIntraParameter;
    spiPtr->onRspQrySPBMInterParameter = callbacks->onRspQrySPBMInterParameter;
    spiPtr->onRspQrySPBMPortfDefinition = callbacks->onRspQrySPBMPortfDefinition;
    spiPtr->onRspQrySPBMInvestorPortfDef = callbacks->onRspQrySPBMInvestorPortfDef;
    spiPtr->onRspQryInvestorPortfMarginRatio = callbacks->onRspQryInvestorPortfMarginRatio;
    spiPtr->onRspQryInvestorProdSPBMDetail = callbacks->onRspQryInvestorProdSPBMDetail;
    spiPtr->onRspQryInvestorCommoditySPMMMargin = callbacks->onRspQryInvestorCommoditySPMMMargin;
    spiPtr->onRspQryInvestorCommodityGroupSPMMMargin = callbacks->onRspQryInvestorCommodityGroupSPMMMargin;
    spiPtr->onRspQrySPMMInstParam = callbacks->onRspQrySPMMInstParam;
    spiPtr->onRspQrySPMMProductParam = callbacks->onRspQrySPMMProductParam;
    spiPtr->onRspQrySPBMAddOnInterParameter = callbacks->onRspQrySPBMAddOnInterParameter;
    spiPtr->onRspQryRCAMSCombProductInfo = callbacks->onRspQryRCAMSCombProductInfo;
    spiPtr->onRspQryRCAMSInstrParameter = callbacks->onRspQryRCAMSInstrParameter;
    spiPtr->onRspQryRCAMSIntraParameter = callbacks->onRspQryRCAMSIntraParameter;
    spiPtr->onRspQryRCAMSInterParameter = callbacks->onRspQryRCAMSInterParameter;
    spiPtr->onRspQryRCAMSShortOptAdjustParam = callbacks->onRspQryRCAMSShortOptAdjustParam;
    spiPtr->onRspQryRCAMSInvestorCombPosition = callbacks->onRspQryRCAMSInvestorCombPosition;
    spiPtr->onRspQryInvestorProdRCAMSMargin = callbacks->onRspQryInvestorProdRCAMSMargin;
    spiPtr->onRspQryRULEInstrParameter = callbacks->onRspQryRULEInstrParameter;
    spiPtr->onRspQryRULEIntraParameter = callbacks->onRspQryRULEIntraParameter;
    spiPtr->onRspQryRULEInterParameter = callbacks->onRspQryRULEInterParameter;
    spiPtr->onRspQryInvestorProdRULEMargin = callbacks->onRspQryInvestorProdRULEMargin;
    spiPtr->onRspQryInvestorPortfSetting = callbacks->onRspQryInvestorPortfSetting;
}

void TraderSpiSetOnFrontConnected(TraderSpiHandle spi, TraderOnFrontConnectedCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onFrontConnected = callback;
}

void TraderSpiSetOnFrontDisconnected(TraderSpiHandle spi, TraderOnFrontDisconnectedCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onFrontDisconnected = callback;
}

void TraderSpiSetOnHeartBeatWarning(TraderSpiHandle spi, TraderOnHeartBeatWarningCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onHeartBeatWarning = callback;
}

void TraderSpiSetOnRspAuthenticate(TraderSpiHandle spi, TraderOnRspAuthenticateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspAuthenticate = callback;
}

void TraderSpiSetOnRspUserLogin(TraderSpiHandle spi, TraderOnRspUserLoginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspUserLogin = callback;
}

void TraderSpiSetOnRspUserLogout(TraderSpiHandle spi, TraderOnRspUserLogoutCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspUserLogout = callback;
}

void TraderSpiSetOnRspUserPasswordUpdate(TraderSpiHandle spi, TraderOnRspUserPasswordUpdateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspUserPasswordUpdate = callback;
}

void TraderSpiSetOnRspTradingAccountPasswordUpdate(TraderSpiHandle spi, TraderOnRspTradingAccountPasswordUpdateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspTradingAccountPasswordUpdate = callback;
}

void TraderSpiSetOnRspUserAuthMethod(TraderSpiHandle spi, TraderOnRspUserAuthMethodCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspUserAuthMethod = callback;
}

void TraderSpiSetOnRspGenUserCaptcha(TraderSpiHandle spi, TraderOnRspGenUserCaptchaCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspGenUserCaptcha = callback;
}

void TraderSpiSetOnRspGenUserText(TraderSpiHandle spi, TraderOnRspGenUserTextCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspGenUserText = callback;
}

void TraderSpiSetOnRspOrderInsert(TraderSpiHandle spi, TraderOnRspOrderInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspOrderInsert = callback;
}

void TraderSpiSetOnRspParkedOrderInsert(TraderSpiHandle spi, TraderOnRspParkedOrderInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspParkedOrderInsert = callback;
}

void TraderSpiSetOnRspParkedOrderAction(TraderSpiHandle spi, TraderOnRspParkedOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspParkedOrderAction = callback;
}

void TraderSpiSetOnRspOrderAction(TraderSpiHandle spi, TraderOnRspOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspOrderAction = callback;
}

void TraderSpiSetOnRspQryMaxOrderVolume(TraderSpiHandle spi, TraderOnRspQryMaxOrderVolumeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryMaxOrderVolume = callback;
}

void TraderSpiSetOnRspSettlementInfoConfirm(TraderSpiHandle spi, TraderOnRspSettlementInfoConfirmCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspSettlementInfoConfirm = callback;
}

void TraderSpiSetOnRspRemoveParkedOrder(TraderSpiHandle spi, TraderOnRspRemoveParkedOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspRemoveParkedOrder = callback;
}

void TraderSpiSetOnRspRemoveParkedOrderAction(TraderSpiHandle spi, TraderOnRspRemoveParkedOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspRemoveParkedOrderAction = callback;
}

void TraderSpiSetOnRspExecOrderInsert(TraderSpiHandle spi, TraderOnRspExecOrderInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspExecOrderInsert = callback;
}

void TraderSpiSetOnRspExecOrderAction(TraderSpiHandle spi, TraderOnRspExecOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspExecOrderAction = callback;
}

void TraderSpiSetOnRspForQuoteInsert(TraderSpiHandle spi, TraderOnRspForQuoteInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspForQuoteInsert = callback;
}

void TraderSpiSetOnRspQuoteInsert(TraderSpiHandle spi, TraderOnRspQuoteInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQuoteInsert = callback;
}

void TraderSpiSetOnRspQuoteAction(TraderSpiHandle spi, TraderOnRspQuoteActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQuoteAction = callback;
}

void TraderSpiSetOnRspBatchOrderAction(TraderSpiHandle spi, TraderOnRspBatchOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspBatchOrderAction = callback;
}

void TraderSpiSetOnRspOptionSelfCloseInsert(TraderSpiHandle spi, TraderOnRspOptionSelfCloseInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspOptionSelfCloseInsert = callback;
}

void TraderSpiSetOnRspOptionSelfCloseAction(TraderSpiHandle spi, TraderOnRspOptionSelfCloseActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspOptionSelfCloseAction = callback;
}

void TraderSpiSetOnRspCombActionInsert(TraderSpiHandle spi, TraderOnRspCombActionInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspCombActionInsert = callback;
}

void TraderSpiSetOnRspQryOrder(TraderSpiHandle spi, TraderOnRspQryOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryOrder = callback;
}

void TraderSpiSetOnRspQryTrade(TraderSpiHandle spi, TraderOnRspQryTradeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTrade = callback;
}

void TraderSpiSetOnRspQryInvestorPosition(TraderSpiHandle spi, TraderOnRspQryInvestorPositionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorPosition = callback;
}

void TraderSpiSetOnRspQryTradingAccount(TraderSpiHandle spi, TraderOnRspQryTradingAccountCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTradingAccount = callback;
}

void TraderSpiSetOnRspQryInvestor(TraderSpiHandle spi, TraderOnRspQryInvestorCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestor = callback;
}

void TraderSpiSetOnRspQryTradingCode(TraderSpiHandle spi, TraderOnRspQryTradingCodeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTradingCode = callback;
}

void TraderSpiSetOnRspQryInstrumentMarginRate(TraderSpiHandle spi, TraderOnRspQryInstrumentMarginRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInstrumentMarginRate = callback;
}

void TraderSpiSetOnRspQryInstrumentCommissionRate(TraderSpiHandle spi, TraderOnRspQryInstrumentCommissionRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInstrumentCommissionRate = callback;
}

void TraderSpiSetOnRspQryExchange(TraderSpiHandle spi, TraderOnRspQryExchangeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryExchange = callback;
}

void TraderSpiSetOnRspQryProduct(TraderSpiHandle spi, TraderOnRspQryProductCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryProduct = callback;
}

void TraderSpiSetOnRspQryInstrument(TraderSpiHandle spi, TraderOnRspQryInstrumentCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInstrument = callback;
}

void TraderSpiSetOnRspQryDepthMarketData(TraderSpiHandle spi, TraderOnRspQryDepthMarketDataCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryDepthMarketData = callback;
}

void TraderSpiSetOnRspQryTraderOffer(TraderSpiHandle spi, TraderOnRspQryTraderOfferCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTraderOffer = callback;
}

void TraderSpiSetOnRspQrySettlementInfo(TraderSpiHandle spi, TraderOnRspQrySettlementInfoCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySettlementInfo = callback;
}

void TraderSpiSetOnRspQryTransferBank(TraderSpiHandle spi, TraderOnRspQryTransferBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTransferBank = callback;
}

void TraderSpiSetOnRspQryInvestorPositionDetail(TraderSpiHandle spi, TraderOnRspQryInvestorPositionDetailCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorPositionDetail = callback;
}

void TraderSpiSetOnRspQryNotice(TraderSpiHandle spi, TraderOnRspQryNoticeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryNotice = callback;
}

void TraderSpiSetOnRspQrySettlementInfoConfirm(TraderSpiHandle spi, TraderOnRspQrySettlementInfoConfirmCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySettlementInfoConfirm = callback;
}

void TraderSpiSetOnRspQryInvestorPositionCombineDetail(TraderSpiHandle spi, TraderOnRspQryInvestorPositionCombineDetailCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorPositionCombineDetail = callback;
}

void TraderSpiSetOnRspQryCFMMCTradingAccountKey(TraderSpiHandle spi, TraderOnRspQryCFMMCTradingAccountKeyCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryCFMMCTradingAccountKey = callback;
}

void TraderSpiSetOnRspQryEWarrantOffset(TraderSpiHandle spi, TraderOnRspQryEWarrantOffsetCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryEWarrantOffset = callback;
}

void TraderSpiSetOnRspQryInvestorProductGroupMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProductGroupMarginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorProductGroupMargin = callback;
}

void TraderSpiSetOnRspQryExchangeMarginRate(TraderSpiHandle spi, TraderOnRspQryExchangeMarginRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryExchangeMarginRate = callback;
}

void TraderSpiSetOnRspQryExchangeMarginRateAdjust(TraderSpiHandle spi, TraderOnRspQryExchangeMarginRateAdjustCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryExchangeMarginRateAdjust = callback;
}

void TraderSpiSetOnRspQryExchangeRate(TraderSpiHandle spi, TraderOnRspQryExchangeRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryExchangeRate = callback;
}

void TraderSpiSetOnRspQrySecAgentACIDMap(TraderSpiHandle spi, TraderOnRspQrySecAgentACIDMapCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySecAgentACIDMap = callback;
}

void TraderSpiSetOnRspQryProductExchRate(TraderSpiHandle spi, TraderOnRspQryProductExchRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryProductExchRate = callback;
}

void TraderSpiSetOnRspQryProductGroup(TraderSpiHandle spi, TraderOnRspQryProductGroupCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryProductGroup = callback;
}

void TraderSpiSetOnRspQryMMInstrumentCommissionRate(TraderSpiHandle spi, TraderOnRspQryMMInstrumentCommissionRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryMMInstrumentCommissionRate = callback;
}

void TraderSpiSetOnRspQryMMOptionInstrCommRate(TraderSpiHandle spi, TraderOnRspQryMMOptionInstrCommRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryMMOptionInstrCommRate = callback;
}

void TraderSpiSetOnRspQryInstrumentOrderCommRate(TraderSpiHandle spi, TraderOnRspQryInstrumentOrderCommRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInstrumentOrderCommRate = callback;
}

void TraderSpiSetOnRspQrySecAgentTradingAccount(TraderSpiHandle spi, TraderOnRspQrySecAgentTradingAccountCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySecAgentTradingAccount = callback;
}

void TraderSpiSetOnRspQrySecAgentCheckMode(TraderSpiHandle spi, TraderOnRspQrySecAgentCheckModeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySecAgentCheckMode = callback;
}

void TraderSpiSetOnRspQrySecAgentTradeInfo(TraderSpiHandle spi, TraderOnRspQrySecAgentTradeInfoCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySecAgentTradeInfo = callback;
}

void TraderSpiSetOnRspQryOptionInstrTradeCost(TraderSpiHandle spi, TraderOnRspQryOptionInstrTradeCostCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryOptionInstrTradeCost = callback;
}

void TraderSpiSetOnRspQryOptionInstrCommRate(TraderSpiHandle spi, TraderOnRspQryOptionInstrCommRateCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryOptionInstrCommRate = callback;
}

void TraderSpiSetOnRspQryExecOrder(TraderSpiHandle spi, TraderOnRspQryExecOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryExecOrder = callback;
}

void TraderSpiSetOnRspQryForQuote(TraderSpiHandle spi, TraderOnRspQryForQuoteCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryForQuote = callback;
}

void TraderSpiSetOnRspQryQuote(TraderSpiHandle spi, TraderOnRspQryQuoteCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryQuote = callback;
}

void TraderSpiSetOnRspQryOptionSelfClose(TraderSpiHandle spi, TraderOnRspQryOptionSelfCloseCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryOptionSelfClose = callback;
}

void TraderSpiSetOnRspQryInvestUnit(TraderSpiHandle spi, TraderOnRspQryInvestUnitCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestUnit = callback;
}

void TraderSpiSetOnRspQryCombInstrumentGuard(TraderSpiHandle spi, TraderOnRspQryCombInstrumentGuardCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryCombInstrumentGuard = callback;
}

void TraderSpiSetOnRspQryCombAction(TraderSpiHandle spi, TraderOnRspQryCombActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryCombAction = callback;
}

void TraderSpiSetOnRspQryTransferSerial(TraderSpiHandle spi, TraderOnRspQryTransferSerialCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTransferSerial = callback;
}

void TraderSpiSetOnRspQryAccountregister(TraderSpiHandle spi, TraderOnRspQryAccountregisterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryAccountregister = callback;
}

void TraderSpiSetOnRspError(TraderSpiHandle spi, TraderOnRspErrorCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspError = callback;
}

void TraderSpiSetOnRtnOrder(TraderSpiHandle spi, TraderOnRtnOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnOrder = callback;
}

void TraderSpiSetOnRtnTrade(TraderSpiHandle spi, TraderOnRtnTradeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnTrade = callback;
}

void TraderSpiSetOnErrRtnOrderInsert(TraderSpiHandle spi, TraderOnErrRtnOrderInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnOrderInsert = callback;
}

void TraderSpiSetOnErrRtnOrderAction(TraderSpiHandle spi, TraderOnErrRtnOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnOrderAction = callback;
}

void TraderSpiSetOnRtnInstrumentStatus(TraderSpiHandle spi, TraderOnRtnInstrumentStatusCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnInstrumentStatus = callback;
}

void TraderSpiSetOnRtnBulletin(TraderSpiHandle spi, TraderOnRtnBulletinCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnBulletin = callback;
}

void TraderSpiSetOnRtnTradingNotice(TraderSpiHandle spi, TraderOnRtnTradingNoticeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnTradingNotice = callback;
}

void TraderSpiSetOnRtnErrorConditionalOrder(TraderSpiHandle spi, TraderOnRtnErrorConditionalOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnErrorConditionalOrder = callback;
}

void TraderSpiSetOnRtnExecOrder(TraderSpiHandle spi, TraderOnRtnExecOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnExecOrder = callback;
}

void TraderSpiSetOnErrRtnExecOrderInsert(TraderSpiHandle spi, TraderOnErrRtnExecOrderInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnExecOrderInsert = callback;
}

void TraderSpiSetOnErrRtnExecOrderAction(TraderSpiHandle spi, TraderOnErrRtnExecOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnExecOrderAction = callback;
}

void TraderSpiSetOnErrRtnForQuoteInsert(TraderSpiHandle spi, TraderOnErrRtnForQuoteInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnForQuoteInsert = callback;
}

void TraderSpiSetOnRtnQuote(TraderSpiHandle spi, TraderOnRtnQuoteCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnQuote = callback;
}

void TraderSpiSetOnErrRtnQuoteInsert(TraderSpiHandle spi, TraderOnErrRtnQuoteInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnQuoteInsert = callback;
}

void TraderSpiSetOnErrRtnQuoteAction(TraderSpiHandle spi, TraderOnErrRtnQuoteActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnQuoteAction = callback;
}

void TraderSpiSetOnRtnForQuoteRsp(TraderSpiHandle spi, TraderOnRtnForQuoteRspCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnForQuoteRsp = callback;
}

void TraderSpiSetOnRtnCFMMCTradingAccountToken(TraderSpiHandle spi, TraderOnRtnCFMMCTradingAccountTokenCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnCFMMCTradingAccountToken = callback;
}

void TraderSpiSetOnErrRtnBatchOrderAction(TraderSpiHandle spi, TraderOnErrRtnBatchOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnBatchOrderAction = callback;
}

void TraderSpiSetOnRtnOptionSelfClose(TraderSpiHandle spi, TraderOnRtnOptionSelfCloseCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnOptionSelfClose = callback;
}

void TraderSpiSetOnErrRtnOptionSelfCloseInsert(TraderSpiHandle spi, TraderOnErrRtnOptionSelfCloseInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnOptionSelfCloseInsert = callback;
}

void TraderSpiSetOnErrRtnOptionSelfCloseAction(TraderSpiHandle spi, TraderOnErrRtnOptionSelfCloseActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnOptionSelfCloseAction = callback;
}

void TraderSpiSetOnRtnCombAction(TraderSpiHandle spi, TraderOnRtnCombActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnCombAction = callback;
}

void TraderSpiSetOnErrRtnCombActionInsert(TraderSpiHandle spi, TraderOnErrRtnCombActionInsertCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnCombActionInsert = callback;
}

void TraderSpiSetOnRspQryContractBank(TraderSpiHandle spi, TraderOnRspQryContractBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryContractBank = callback;
}

void TraderSpiSetOnRspQryParkedOrder(TraderSpiHandle spi, TraderOnRspQryParkedOrderCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryParkedOrder = callback;
}

void TraderSpiSetOnRspQryParkedOrderAction(TraderSpiHandle spi, TraderOnRspQryParkedOrderActionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryParkedOrderAction = callback;
}

void TraderSpiSetOnRspQryTradingNotice(TraderSpiHandle spi, TraderOnRspQryTradingNoticeCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryTradingNotice = callback;
}

void TraderSpiSetOnRspQryBrokerTradingParams(TraderSpiHandle spi, TraderOnRspQryBrokerTradingParamsCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryBrokerTradingParams = callback;
}

void TraderSpiSetOnRspQryBrokerTradingAlgos(TraderSpiHandle spi, TraderOnRspQryBrokerTradingAlgosCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryBrokerTradingAlgos = callback;
}

void TraderSpiSetOnRspQueryCFMMCTradingAccountToken(TraderSpiHandle spi, TraderOnRspQueryCFMMCTradingAccountTokenCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQueryCFMMCTradingAccountToken = callback;
}

void TraderSpiSetOnRtnFromBankToFutureByBank(TraderSpiHandle spi, TraderOnRtnFromBankToFutureByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnFromBankToFutureByBank = callback;
}

void TraderSpiSetOnRtnFromFutureToBankByBank(TraderSpiHandle spi, TraderOnRtnFromFutureToBankByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnFromFutureToBankByBank = callback;
}

void TraderSpiSetOnRtnRepealFromBankToFutureByBank(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromBankToFutureByBank = callback;
}

void TraderSpiSetOnRtnRepealFromFutureToBankByBank(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromFutureToBankByBank = callback;
}

void TraderSpiSetOnRtnFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRtnFromBankToFutureByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnFromBankToFutureByFuture = callback;
}

void TraderSpiSetOnRtnFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRtnFromFutureToBankByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnFromFutureToBankByFuture = callback;
}

void TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByFutureManualCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromBankToFutureByFutureManual = callback;
}

void TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByFutureManualCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromFutureToBankByFutureManual = callback;
}

void TraderSpiSetOnRtnQueryBankBalanceByFuture(TraderSpiHandle spi, TraderOnRtnQueryBankBalanceByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnQueryBankBalanceByFuture = callback;
}

void TraderSpiSetOnErrRtnBankToFutureByFuture(TraderSpiHandle spi, TraderOnErrRtnBankToFutureByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnBankToFutureByFuture = callback;
}

void TraderSpiSetOnErrRtnFutureToBankByFuture(TraderSpiHandle spi, TraderOnErrRtnFutureToBankByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnFutureToBankByFuture = callback;
}

void TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual(TraderSpiHandle spi, TraderOnErrRtnRepealBankToFutureByFutureManualCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnRepealBankToFutureByFutureManual = callback;
}

void TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual(TraderSpiHandle spi, TraderOnErrRtnRepealFutureToBankByFutureManualCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnRepealFutureToBankByFutureManual = callback;
}

void TraderSpiSetOnErrRtnQueryBankBalanceByFuture(TraderSpiHandle spi, TraderOnErrRtnQueryBankBalanceByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onErrRtnQueryBankBalanceByFuture = callback;
}

void TraderSpiSetOnRtnRepealFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromBankToFutureByFuture = callback;
}

void TraderSpiSetOnRtnRepealFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnRepealFromFutureToBankByFuture = callback;
}

void TraderSpiSetOnRspFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRspFromBankToFutureByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspFromBankToFutureByFuture = callback;
}

void TraderSpiSetOnRspFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRspFromFutureToBankByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspFromFutureToBankByFuture = callback;
}

void TraderSpiSetOnRspQueryBankAccountMoneyByFuture(TraderSpiHandle spi, TraderOnRspQueryBankAccountMoneyByFutureCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQueryBankAccountMoneyByFuture = callback;
}

void TraderSpiSetOnRtnOpenAccountByBank(TraderSpiHandle spi, TraderOnRtnOpenAccountByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnOpenAccountByBank = callback;
}

void TraderSpiSetOnRtnCancelAccountByBank(TraderSpiHandle spi, TraderOnRtnCancelAccountByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnCancelAccountByBank = callback;
}

void TraderSpiSetOnRtnChangeAccountByBank(TraderSpiHandle spi, TraderOnRtnChangeAccountByBankCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRtnChangeAccountByBank = callback;
}

void TraderSpiSetOnRspQryClassifiedInstrument(TraderSpiHandle spi, TraderOnRspQryClassifiedInstrumentCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryClassifiedInstrument = callback;
}

void TraderSpiSetOnRspQryCombPromotionParam(TraderSpiHandle spi, TraderOnRspQryCombPromotionParamCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryCombPromotionParam = callback;
}

void TraderSpiSetOnRspQryRiskSettleInvstPosition(TraderSpiHandle spi, TraderOnRspQryRiskSettleInvstPositionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRiskSettleInvstPosition = callback;
}

void TraderSpiSetOnRspQryRiskSettleProductStatus(TraderSpiHandle spi, TraderOnRspQryRiskSettleProductStatusCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRiskSettleProductStatus = callback;
}

void TraderSpiSetOnRspQrySPBMFutureParameter(TraderSpiHandle spi, TraderOnRspQrySPBMFutureParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMFutureParameter = callback;
}

void TraderSpiSetOnRspQrySPBMOptionParameter(TraderSpiHandle spi, TraderOnRspQrySPBMOptionParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMOptionParameter = callback;
}

void TraderSpiSetOnRspQrySPBMIntraParameter(TraderSpiHandle spi, TraderOnRspQrySPBMIntraParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMIntraParameter = callback;
}

void TraderSpiSetOnRspQrySPBMInterParameter(TraderSpiHandle spi, TraderOnRspQrySPBMInterParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMInterParameter = callback;
}

void TraderSpiSetOnRspQrySPBMPortfDefinition(TraderSpiHandle spi, TraderOnRspQrySPBMPortfDefinitionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMPortfDefinition = callback;
}

void TraderSpiSetOnRspQrySPBMInvestorPortfDef(TraderSpiHandle spi, TraderOnRspQrySPBMInvestorPortfDefCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMInvestorPortfDef = callback;
}

void TraderSpiSetOnRspQryInvestorPortfMarginRatio(TraderSpiHandle spi, TraderOnRspQryInvestorPortfMarginRatioCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorPortfMarginRatio = callback;
}

void TraderSpiSetOnRspQryInvestorProdSPBMDetail(TraderSpiHandle spi, TraderOnRspQryInvestorProdSPBMDetailCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorProdSPBMDetail = callback;
}

void TraderSpiSetOnRspQryInvestorCommoditySPMMMargin(TraderSpiHandle spi, TraderOnRspQryInvestorCommoditySPMMMarginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorCommoditySPMMMargin = callback;
}

void TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin(TraderSpiHandle spi, TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorCommodityGroupSPMMMargin = callback;
}

void TraderSpiSetOnRspQrySPMMInstParam(TraderSpiHandle spi, TraderOnRspQrySPMMInstParamCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPMMInstParam = callback;
}

void TraderSpiSetOnRspQrySPMMProductParam(TraderSpiHandle spi, TraderOnRspQrySPMMProductParamCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPMMProductParam = callback;
}

void TraderSpiSetOnRspQrySPBMAddOnInterParameter(TraderSpiHandle spi, TraderOnRspQrySPBMAddOnInterParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQrySPBMAddOnInterParameter = callback;
}

void TraderSpiSetOnRspQryRCAMSCombProductInfo(TraderSpiHandle spi, TraderOnRspQryRCAMSCombProductInfoCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSCombProductInfo = callback;
}

void TraderSpiSetOnRspQryRCAMSInstrParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSInstrParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSInstrParameter = callback;
}

void TraderSpiSetOnRspQryRCAMSIntraParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSIntraParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSIntraParameter = callback;
}

void TraderSpiSetOnRspQryRCAMSInterParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSInterParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSInterParameter = callback;
}

void TraderSpiSetOnRspQryRCAMSShortOptAdjustParam(TraderSpiHandle spi, TraderOnRspQryRCAMSShortOptAdjustParamCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSShortOptAdjustParam = callback;
}

void TraderSpiSetOnRspQryRCAMSInvestorCombPosition(TraderSpiHandle spi, TraderOnRspQryRCAMSInvestorCombPositionCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRCAMSInvestorCombPosition = callback;
}

void TraderSpiSetOnRspQryInvestorProdRCAMSMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProdRCAMSMarginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorProdRCAMSMargin = callback;
}

void TraderSpiSetOnRspQryRULEInstrParameter(TraderSpiHandle spi, TraderOnRspQryRULEInstrParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRULEInstrParameter = callback;
}

void TraderSpiSetOnRspQryRULEIntraParameter(TraderSpiHandle spi, TraderOnRspQryRULEIntraParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRULEIntraParameter = callback;
}

void TraderSpiSetOnRspQryRULEInterParameter(TraderSpiHandle spi, TraderOnRspQryRULEInterParameterCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryRULEInterParameter = callback;
}

void TraderSpiSetOnRspQryInvestorProdRULEMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProdRULEMarginCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorProdRULEMargin = callback;
}

void TraderSpiSetOnRspQryInvestorPortfSetting(TraderSpiHandle spi, TraderOnRspQryInvestorPortfSettingCallback callback) {
    reinterpret_cast<TraderSpiWrapper*>(spi)->onRspQryInvestorPortfSetting = callback;
}

// DataCollect 函数实现
int DCGetSystemInfo(char* pSystemInfo, int* pLen) {
    return CTP_GetSystemInfo(pSystemInfo, *pLen);
}

int DCGetSystemInfoUnAesEncode(char* pSystemInfo, int* pLen) {
#ifdef __APPLE__
    // macOS 版本有专门的未 AES 加密函数
    return CTP_GetSystemInfoUnAesEncode(pSystemInfo, *pLen);
#else
    // Linux/Windows 版本没有此函数，回退到 CTP_GetSystemInfo
    // 注意：这会返回 AES 加密的数据，但对于内部登录流程不影响
    // 因为 Linux/Windows 的 ReqUserLogin 不需要 systemInfo 参数
    return CTP_GetSystemInfo(pSystemInfo, *pLen);
#endif
}

const char* DCGetDataCollectApiVersion(void) {
    return CTP_GetDataCollectApiVersion();
}

} // extern "C"
