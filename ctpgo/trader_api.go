package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 交易 API 封装

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// ========== 回调类型定义 ==========

// TraderOnFrontConnectedCallback ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
type TraderOnFrontConnectedCallback func(userData uintptr)

// TraderOnFrontDisconnectedCallback 0x2003 收到错误报文
type TraderOnFrontDisconnectedCallback func(userData uintptr, nReason int32)

// TraderOnHeartBeatWarningCallback 心跳超时警告。当长时间未收到报文时，该方法被调用。
type TraderOnHeartBeatWarningCallback func(userData uintptr, nTimeLapse int32)

// TraderOnRspAuthenticateCallback 客户端认证响应
type TraderOnRspAuthenticateCallback func(userData uintptr, pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspUserLoginCallback 登录请求响应
type TraderOnRspUserLoginCallback func(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspUserLogoutCallback 登出请求响应
type TraderOnRspUserLogoutCallback func(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspUserPasswordUpdateCallback 用户口令更新请求响应
type TraderOnRspUserPasswordUpdateCallback func(userData uintptr, pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspTradingAccountPasswordUpdateCallback 资金账户口令更新请求响应
type TraderOnRspTradingAccountPasswordUpdateCallback func(userData uintptr, pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspUserAuthMethodCallback 查询用户当前支持的认证模式的回复
type TraderOnRspUserAuthMethodCallback func(userData uintptr, pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspGenUserCaptchaCallback 获取图形验证码请求的回复
type TraderOnRspGenUserCaptchaCallback func(userData uintptr, pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspGenUserTextCallback 获取短信验证码请求的回复
type TraderOnRspGenUserTextCallback func(userData uintptr, pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspOrderInsertCallback 报单录入请求响应
type TraderOnRspOrderInsertCallback func(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspParkedOrderInsertCallback 预埋单录入请求响应
type TraderOnRspParkedOrderInsertCallback func(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspParkedOrderActionCallback 预埋撤单录入请求响应
type TraderOnRspParkedOrderActionCallback func(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspOrderActionCallback 报单操作请求响应
type TraderOnRspOrderActionCallback func(userData uintptr, pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryMaxOrderVolumeCallback 查询最大报单数量响应
type TraderOnRspQryMaxOrderVolumeCallback func(userData uintptr, pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspSettlementInfoConfirmCallback 投资者结算结果确认响应
type TraderOnRspSettlementInfoConfirmCallback func(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspRemoveParkedOrderCallback 删除预埋单响应
type TraderOnRspRemoveParkedOrderCallback func(userData uintptr, pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspRemoveParkedOrderActionCallback 删除预埋撤单响应
type TraderOnRspRemoveParkedOrderActionCallback func(userData uintptr, pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspExecOrderInsertCallback 执行宣告录入请求响应
type TraderOnRspExecOrderInsertCallback func(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspExecOrderActionCallback 执行宣告操作请求响应
type TraderOnRspExecOrderActionCallback func(userData uintptr, pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspForQuoteInsertCallback 询价录入请求响应
type TraderOnRspForQuoteInsertCallback func(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQuoteInsertCallback 报价录入请求响应
type TraderOnRspQuoteInsertCallback func(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQuoteActionCallback 报价操作请求响应
type TraderOnRspQuoteActionCallback func(userData uintptr, pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspBatchOrderActionCallback 批量报单操作请求响应
type TraderOnRspBatchOrderActionCallback func(userData uintptr, pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspOptionSelfCloseInsertCallback 期权自对冲录入请求响应
type TraderOnRspOptionSelfCloseInsertCallback func(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspOptionSelfCloseActionCallback 期权自对冲操作请求响应
type TraderOnRspOptionSelfCloseActionCallback func(userData uintptr, pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspCombActionInsertCallback 申请组合录入请求响应
type TraderOnRspCombActionInsertCallback func(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryOrderCallback 请求查询报单响应
type TraderOnRspQryOrderCallback func(userData uintptr, pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTradeCallback 请求查询成交响应
type TraderOnRspQryTradeCallback func(userData uintptr, pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorPositionCallback 请求查询投资者持仓响应
type TraderOnRspQryInvestorPositionCallback func(userData uintptr, pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTradingAccountCallback 请求查询资金账户响应
type TraderOnRspQryTradingAccountCallback func(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorCallback 请求查询投资者响应
type TraderOnRspQryInvestorCallback func(userData uintptr, pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTradingCodeCallback 请求查询交易编码响应
type TraderOnRspQryTradingCodeCallback func(userData uintptr, pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInstrumentMarginRateCallback 请求查询合约保证金率响应
type TraderOnRspQryInstrumentMarginRateCallback func(userData uintptr, pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInstrumentCommissionRateCallback 请求查询合约手续费率响应
type TraderOnRspQryInstrumentCommissionRateCallback func(userData uintptr, pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryExchangeCallback 请求查询交易所响应
type TraderOnRspQryExchangeCallback func(userData uintptr, pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryProductCallback 请求查询产品响应
type TraderOnRspQryProductCallback func(userData uintptr, pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInstrumentCallback 请求查询合约响应
type TraderOnRspQryInstrumentCallback func(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryDepthMarketDataCallback 请求查询行情响应
type TraderOnRspQryDepthMarketDataCallback func(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTraderOfferCallback 请求查询交易员报盘机响应
type TraderOnRspQryTraderOfferCallback func(userData uintptr, pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySettlementInfoCallback 请求查询投资者结算结果响应
type TraderOnRspQrySettlementInfoCallback func(userData uintptr, pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTransferBankCallback 请求查询转帐银行响应
type TraderOnRspQryTransferBankCallback func(userData uintptr, pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorPositionDetailCallback 请求查询投资者持仓明细响应
type TraderOnRspQryInvestorPositionDetailCallback func(userData uintptr, pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryNoticeCallback 请求查询客户通知响应
type TraderOnRspQryNoticeCallback func(userData uintptr, pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySettlementInfoConfirmCallback 请求查询结算信息确认响应
type TraderOnRspQrySettlementInfoConfirmCallback func(userData uintptr, pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorPositionCombineDetailCallback 请求查询投资者持仓明细响应
type TraderOnRspQryInvestorPositionCombineDetailCallback func(userData uintptr, pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryCFMMCTradingAccountKeyCallback 查询保证金监管系统经纪公司资金账户密钥响应
type TraderOnRspQryCFMMCTradingAccountKeyCallback func(userData uintptr, pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryEWarrantOffsetCallback 请求查询仓单折抵信息响应
type TraderOnRspQryEWarrantOffsetCallback func(userData uintptr, pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorProductGroupMarginCallback 请求查询投资者品种/跨品种保证金响应
type TraderOnRspQryInvestorProductGroupMarginCallback func(userData uintptr, pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryExchangeMarginRateCallback 请求查询交易所保证金率响应
type TraderOnRspQryExchangeMarginRateCallback func(userData uintptr, pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryExchangeMarginRateAdjustCallback 请求查询交易所调整保证金率响应
type TraderOnRspQryExchangeMarginRateAdjustCallback func(userData uintptr, pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryExchangeRateCallback 请求查询汇率响应
type TraderOnRspQryExchangeRateCallback func(userData uintptr, pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySecAgentACIDMapCallback 请求查询二级代理操作员银期权限响应
type TraderOnRspQrySecAgentACIDMapCallback func(userData uintptr, pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryProductExchRateCallback 请求查询产品报价汇率
type TraderOnRspQryProductExchRateCallback func(userData uintptr, pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryProductGroupCallback 请求查询产品组
type TraderOnRspQryProductGroupCallback func(userData uintptr, pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryMMInstrumentCommissionRateCallback 请求查询做市商合约手续费率响应
type TraderOnRspQryMMInstrumentCommissionRateCallback func(userData uintptr, pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryMMOptionInstrCommRateCallback 请求查询做市商期权合约手续费响应
type TraderOnRspQryMMOptionInstrCommRateCallback func(userData uintptr, pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInstrumentOrderCommRateCallback 请求查询报单手续费响应
type TraderOnRspQryInstrumentOrderCommRateCallback func(userData uintptr, pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySecAgentTradingAccountCallback 请求查询资金账户响应
type TraderOnRspQrySecAgentTradingAccountCallback func(userData uintptr, pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySecAgentCheckModeCallback 请求查询二级代理商资金校验模式响应
type TraderOnRspQrySecAgentCheckModeCallback func(userData uintptr, pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySecAgentTradeInfoCallback 请求查询二级代理商信息响应
type TraderOnRspQrySecAgentTradeInfoCallback func(userData uintptr, pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryOptionInstrTradeCostCallback 请求查询期权交易成本响应
type TraderOnRspQryOptionInstrTradeCostCallback func(userData uintptr, pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryOptionInstrCommRateCallback 请求查询期权合约手续费响应
type TraderOnRspQryOptionInstrCommRateCallback func(userData uintptr, pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryExecOrderCallback 请求查询执行宣告响应
type TraderOnRspQryExecOrderCallback func(userData uintptr, pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryForQuoteCallback 请求查询询价响应
type TraderOnRspQryForQuoteCallback func(userData uintptr, pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryQuoteCallback 请求查询报价响应
type TraderOnRspQryQuoteCallback func(userData uintptr, pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryOptionSelfCloseCallback 请求查询期权自对冲响应
type TraderOnRspQryOptionSelfCloseCallback func(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestUnitCallback 请求查询投资单元响应
type TraderOnRspQryInvestUnitCallback func(userData uintptr, pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryCombInstrumentGuardCallback 请求查询组合合约安全系数响应
type TraderOnRspQryCombInstrumentGuardCallback func(userData uintptr, pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryCombActionCallback 请求查询申请组合响应
type TraderOnRspQryCombActionCallback func(userData uintptr, pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTransferSerialCallback 请求查询转帐流水响应
type TraderOnRspQryTransferSerialCallback func(userData uintptr, pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryAccountregisterCallback 请求查询银期签约关系响应
type TraderOnRspQryAccountregisterCallback func(userData uintptr, pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspErrorCallback 错误应答
type TraderOnRspErrorCallback func(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRtnOrderCallback 报单通知
type TraderOnRtnOrderCallback func(userData uintptr, pOrder *CThostFtdcOrderField)

// TraderOnRtnTradeCallback 成交通知
type TraderOnRtnTradeCallback func(userData uintptr, pTrade *CThostFtdcTradeField)

// TraderOnErrRtnOrderInsertCallback 报单录入错误回报
type TraderOnErrRtnOrderInsertCallback func(userData uintptr, pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnOrderActionCallback 报单操作错误回报
type TraderOnErrRtnOrderActionCallback func(userData uintptr, pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnInstrumentStatusCallback 合约交易状态通知
type TraderOnRtnInstrumentStatusCallback func(userData uintptr, pInstrumentStatus *CThostFtdcInstrumentStatusField)

// TraderOnRtnBulletinCallback 交易所公告通知
type TraderOnRtnBulletinCallback func(userData uintptr, pBulletin *CThostFtdcBulletinField)

// TraderOnRtnTradingNoticeCallback 交易通知
type TraderOnRtnTradingNoticeCallback func(userData uintptr, pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField)

// TraderOnRtnErrorConditionalOrderCallback 提示条件单校验错误
type TraderOnRtnErrorConditionalOrderCallback func(userData uintptr, pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField)

// TraderOnRtnExecOrderCallback 执行宣告通知
type TraderOnRtnExecOrderCallback func(userData uintptr, pExecOrder *CThostFtdcExecOrderField)

// TraderOnErrRtnExecOrderInsertCallback 执行宣告录入错误回报
type TraderOnErrRtnExecOrderInsertCallback func(userData uintptr, pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnExecOrderActionCallback 执行宣告操作错误回报
type TraderOnErrRtnExecOrderActionCallback func(userData uintptr, pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnForQuoteInsertCallback 询价录入错误回报
type TraderOnErrRtnForQuoteInsertCallback func(userData uintptr, pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnQuoteCallback 报价通知
type TraderOnRtnQuoteCallback func(userData uintptr, pQuote *CThostFtdcQuoteField)

// TraderOnErrRtnQuoteInsertCallback 报价录入错误回报
type TraderOnErrRtnQuoteInsertCallback func(userData uintptr, pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnQuoteActionCallback 报价操作错误回报
type TraderOnErrRtnQuoteActionCallback func(userData uintptr, pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnForQuoteRspCallback 询价通知
type TraderOnRtnForQuoteRspCallback func(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField)

// TraderOnRtnCFMMCTradingAccountTokenCallback 保证金监控中心用户令牌
type TraderOnRtnCFMMCTradingAccountTokenCallback func(userData uintptr, pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField)

// TraderOnErrRtnBatchOrderActionCallback 批量报单操作错误回报
type TraderOnErrRtnBatchOrderActionCallback func(userData uintptr, pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnOptionSelfCloseCallback 期权自对冲通知
type TraderOnRtnOptionSelfCloseCallback func(userData uintptr, pOptionSelfClose *CThostFtdcOptionSelfCloseField)

// TraderOnErrRtnOptionSelfCloseInsertCallback 期权自对冲录入错误回报
type TraderOnErrRtnOptionSelfCloseInsertCallback func(userData uintptr, pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnOptionSelfCloseActionCallback 期权自对冲操作错误回报
type TraderOnErrRtnOptionSelfCloseActionCallback func(userData uintptr, pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnCombActionCallback 申请组合通知
type TraderOnRtnCombActionCallback func(userData uintptr, pCombAction *CThostFtdcCombActionField)

// TraderOnErrRtnCombActionInsertCallback 申请组合录入错误回报
type TraderOnErrRtnCombActionInsertCallback func(userData uintptr, pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRspQryContractBankCallback 请求查询签约银行响应
type TraderOnRspQryContractBankCallback func(userData uintptr, pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryParkedOrderCallback 请求查询预埋单响应
type TraderOnRspQryParkedOrderCallback func(userData uintptr, pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryParkedOrderActionCallback 请求查询预埋撤单响应
type TraderOnRspQryParkedOrderActionCallback func(userData uintptr, pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryTradingNoticeCallback 请求查询交易通知响应
type TraderOnRspQryTradingNoticeCallback func(userData uintptr, pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryBrokerTradingParamsCallback 请求查询经纪公司交易参数响应
type TraderOnRspQryBrokerTradingParamsCallback func(userData uintptr, pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryBrokerTradingAlgosCallback 请求查询经纪公司交易算法响应
type TraderOnRspQryBrokerTradingAlgosCallback func(userData uintptr, pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQueryCFMMCTradingAccountTokenCallback 请求查询监控中心用户令牌
type TraderOnRspQueryCFMMCTradingAccountTokenCallback func(userData uintptr, pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRtnFromBankToFutureByBankCallback 银行发起银行资金转期货通知
type TraderOnRtnFromBankToFutureByBankCallback func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField)

// TraderOnRtnFromFutureToBankByBankCallback 银行发起期货资金转银行通知
type TraderOnRtnFromFutureToBankByBankCallback func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField)

// TraderOnRtnRepealFromBankToFutureByBankCallback 银行发起冲正银行转期货通知
type TraderOnRtnRepealFromBankToFutureByBankCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRtnRepealFromFutureToBankByBankCallback 银行发起冲正期货转银行通知
type TraderOnRtnRepealFromFutureToBankByBankCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRtnFromBankToFutureByFutureCallback 期货发起银行资金转期货通知
type TraderOnRtnFromBankToFutureByFutureCallback func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField)

// TraderOnRtnFromFutureToBankByFutureCallback 期货发起期货资金转银行通知
type TraderOnRtnFromFutureToBankByFutureCallback func(userData uintptr, pRspTransfer *CThostFtdcRspTransferField)

// TraderOnRtnRepealFromBankToFutureByFutureManualCallback 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
type TraderOnRtnRepealFromBankToFutureByFutureManualCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRtnRepealFromFutureToBankByFutureManualCallback 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
type TraderOnRtnRepealFromFutureToBankByFutureManualCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRtnQueryBankBalanceByFutureCallback 期货发起查询银行余额通知
type TraderOnRtnQueryBankBalanceByFutureCallback func(userData uintptr, pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField)

// TraderOnErrRtnBankToFutureByFutureCallback 期货发起银行资金转期货错误回报
type TraderOnErrRtnBankToFutureByFutureCallback func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnFutureToBankByFutureCallback 期货发起期货资金转银行错误回报
type TraderOnErrRtnFutureToBankByFutureCallback func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnRepealBankToFutureByFutureManualCallback 系统运行时期货端手工发起冲正银行转期货错误回报
type TraderOnErrRtnRepealBankToFutureByFutureManualCallback func(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnRepealFutureToBankByFutureManualCallback 系统运行时期货端手工发起冲正期货转银行错误回报
type TraderOnErrRtnRepealFutureToBankByFutureManualCallback func(userData uintptr, pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnErrRtnQueryBankBalanceByFutureCallback 期货发起查询银行余额错误回报
type TraderOnErrRtnQueryBankBalanceByFutureCallback func(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField)

// TraderOnRtnRepealFromBankToFutureByFutureCallback 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
type TraderOnRtnRepealFromBankToFutureByFutureCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRtnRepealFromFutureToBankByFutureCallback 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
type TraderOnRtnRepealFromFutureToBankByFutureCallback func(userData uintptr, pRspRepeal *CThostFtdcRspRepealField)

// TraderOnRspFromBankToFutureByFutureCallback 期货发起银行资金转期货应答
type TraderOnRspFromBankToFutureByFutureCallback func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspFromFutureToBankByFutureCallback 期货发起期货资金转银行应答
type TraderOnRspFromFutureToBankByFutureCallback func(userData uintptr, pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQueryBankAccountMoneyByFutureCallback 期货发起查询银行余额应答
type TraderOnRspQueryBankAccountMoneyByFutureCallback func(userData uintptr, pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRtnOpenAccountByBankCallback 银行发起银期开户通知
type TraderOnRtnOpenAccountByBankCallback func(userData uintptr, pOpenAccount *CThostFtdcOpenAccountField)

// TraderOnRtnCancelAccountByBankCallback 银行发起银期销户通知
type TraderOnRtnCancelAccountByBankCallback func(userData uintptr, pCancelAccount *CThostFtdcCancelAccountField)

// TraderOnRtnChangeAccountByBankCallback 银行发起变更银行账号通知
type TraderOnRtnChangeAccountByBankCallback func(userData uintptr, pChangeAccount *CThostFtdcChangeAccountField)

// TraderOnRspQryClassifiedInstrumentCallback 请求查询分类合约响应
type TraderOnRspQryClassifiedInstrumentCallback func(userData uintptr, pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryCombPromotionParamCallback 请求组合优惠比例响应
type TraderOnRspQryCombPromotionParamCallback func(userData uintptr, pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRiskSettleInvstPositionCallback 投资者风险结算持仓查询响应
type TraderOnRspQryRiskSettleInvstPositionCallback func(userData uintptr, pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRiskSettleProductStatusCallback 风险结算产品查询响应
type TraderOnRspQryRiskSettleProductStatusCallback func(userData uintptr, pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMFutureParameterCallback SPBM期货合约参数查询响应
type TraderOnRspQrySPBMFutureParameterCallback func(userData uintptr, pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMOptionParameterCallback SPBM期权合约参数查询响应
type TraderOnRspQrySPBMOptionParameterCallback func(userData uintptr, pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMIntraParameterCallback SPBM品种内对锁仓折扣参数查询响应
type TraderOnRspQrySPBMIntraParameterCallback func(userData uintptr, pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMInterParameterCallback SPBM跨品种抵扣参数查询响应
type TraderOnRspQrySPBMInterParameterCallback func(userData uintptr, pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMPortfDefinitionCallback SPBM组合保证金套餐查询响应
type TraderOnRspQrySPBMPortfDefinitionCallback func(userData uintptr, pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMInvestorPortfDefCallback 投资者SPBM套餐选择查询响应
type TraderOnRspQrySPBMInvestorPortfDefCallback func(userData uintptr, pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorPortfMarginRatioCallback 投资者新型组合保证金系数查询响应
type TraderOnRspQryInvestorPortfMarginRatioCallback func(userData uintptr, pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorProdSPBMDetailCallback 投资者产品SPBM明细查询响应
type TraderOnRspQryInvestorProdSPBMDetailCallback func(userData uintptr, pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorCommoditySPMMMarginCallback 投资者商品组SPMM记录查询响应
type TraderOnRspQryInvestorCommoditySPMMMarginCallback func(userData uintptr, pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback 投资者商品群SPMM记录查询响应
type TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback func(userData uintptr, pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPMMInstParamCallback SPMM合约参数查询响应
type TraderOnRspQrySPMMInstParamCallback func(userData uintptr, pSPMMInstParam *CThostFtdcSPMMInstParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPMMProductParamCallback SPMM产品参数查询响应
type TraderOnRspQrySPMMProductParamCallback func(userData uintptr, pSPMMProductParam *CThostFtdcSPMMProductParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQrySPBMAddOnInterParameterCallback SPBM附加跨品种抵扣参数查询响应
type TraderOnRspQrySPBMAddOnInterParameterCallback func(userData uintptr, pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSCombProductInfoCallback RCAMS产品组合信息查询响应
type TraderOnRspQryRCAMSCombProductInfoCallback func(userData uintptr, pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSInstrParameterCallback RCAMS同合约风险对冲参数查询响应
type TraderOnRspQryRCAMSInstrParameterCallback func(userData uintptr, pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSIntraParameterCallback RCAMS品种内风险对冲参数查询响应
type TraderOnRspQryRCAMSIntraParameterCallback func(userData uintptr, pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSInterParameterCallback RCAMS跨品种风险折抵参数查询响应
type TraderOnRspQryRCAMSInterParameterCallback func(userData uintptr, pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSShortOptAdjustParamCallback RCAMS空头期权风险调整参数查询响应
type TraderOnRspQryRCAMSShortOptAdjustParamCallback func(userData uintptr, pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRCAMSInvestorCombPositionCallback RCAMS策略组合持仓查询响应
type TraderOnRspQryRCAMSInvestorCombPositionCallback func(userData uintptr, pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorProdRCAMSMarginCallback 投资者品种RCAMS保证金查询响应
type TraderOnRspQryInvestorProdRCAMSMarginCallback func(userData uintptr, pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRULEInstrParameterCallback RULE合约保证金参数查询响应
type TraderOnRspQryRULEInstrParameterCallback func(userData uintptr, pRULEInstrParameter *CThostFtdcRULEInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRULEIntraParameterCallback RULE品种内对锁仓折扣参数查询响应
type TraderOnRspQryRULEIntraParameterCallback func(userData uintptr, pRULEIntraParameter *CThostFtdcRULEIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryRULEInterParameterCallback RULE跨品种抵扣参数查询响应
type TraderOnRspQryRULEInterParameterCallback func(userData uintptr, pRULEInterParameter *CThostFtdcRULEInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorProdRULEMarginCallback 投资者产品RULE保证金查询响应
type TraderOnRspQryInvestorProdRULEMarginCallback func(userData uintptr, pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderOnRspQryInvestorPortfSettingCallback 投资者投资者新组保设置查询响应
type TraderOnRspQryInvestorPortfSettingCallback func(userData uintptr, pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// TraderSpiCallbacks 回调结构体（用于批量设置）
type TraderSpiCallbacks struct {
	UserData                                  uintptr
	OnFrontConnected                          TraderOnFrontConnectedCallback
	OnFrontDisconnected                       TraderOnFrontDisconnectedCallback
	OnHeartBeatWarning                        TraderOnHeartBeatWarningCallback
	OnRspAuthenticate                         TraderOnRspAuthenticateCallback
	OnRspUserLogin                            TraderOnRspUserLoginCallback
	OnRspUserLogout                           TraderOnRspUserLogoutCallback
	OnRspUserPasswordUpdate                   TraderOnRspUserPasswordUpdateCallback
	OnRspTradingAccountPasswordUpdate         TraderOnRspTradingAccountPasswordUpdateCallback
	OnRspUserAuthMethod                       TraderOnRspUserAuthMethodCallback
	OnRspGenUserCaptcha                       TraderOnRspGenUserCaptchaCallback
	OnRspGenUserText                          TraderOnRspGenUserTextCallback
	OnRspOrderInsert                          TraderOnRspOrderInsertCallback
	OnRspParkedOrderInsert                    TraderOnRspParkedOrderInsertCallback
	OnRspParkedOrderAction                    TraderOnRspParkedOrderActionCallback
	OnRspOrderAction                          TraderOnRspOrderActionCallback
	OnRspQryMaxOrderVolume                    TraderOnRspQryMaxOrderVolumeCallback
	OnRspSettlementInfoConfirm                TraderOnRspSettlementInfoConfirmCallback
	OnRspRemoveParkedOrder                    TraderOnRspRemoveParkedOrderCallback
	OnRspRemoveParkedOrderAction              TraderOnRspRemoveParkedOrderActionCallback
	OnRspExecOrderInsert                      TraderOnRspExecOrderInsertCallback
	OnRspExecOrderAction                      TraderOnRspExecOrderActionCallback
	OnRspForQuoteInsert                       TraderOnRspForQuoteInsertCallback
	OnRspQuoteInsert                          TraderOnRspQuoteInsertCallback
	OnRspQuoteAction                          TraderOnRspQuoteActionCallback
	OnRspBatchOrderAction                     TraderOnRspBatchOrderActionCallback
	OnRspOptionSelfCloseInsert                TraderOnRspOptionSelfCloseInsertCallback
	OnRspOptionSelfCloseAction                TraderOnRspOptionSelfCloseActionCallback
	OnRspCombActionInsert                     TraderOnRspCombActionInsertCallback
	OnRspQryOrder                             TraderOnRspQryOrderCallback
	OnRspQryTrade                             TraderOnRspQryTradeCallback
	OnRspQryInvestorPosition                  TraderOnRspQryInvestorPositionCallback
	OnRspQryTradingAccount                    TraderOnRspQryTradingAccountCallback
	OnRspQryInvestor                          TraderOnRspQryInvestorCallback
	OnRspQryTradingCode                       TraderOnRspQryTradingCodeCallback
	OnRspQryInstrumentMarginRate              TraderOnRspQryInstrumentMarginRateCallback
	OnRspQryInstrumentCommissionRate          TraderOnRspQryInstrumentCommissionRateCallback
	OnRspQryExchange                          TraderOnRspQryExchangeCallback
	OnRspQryProduct                           TraderOnRspQryProductCallback
	OnRspQryInstrument                        TraderOnRspQryInstrumentCallback
	OnRspQryDepthMarketData                   TraderOnRspQryDepthMarketDataCallback
	OnRspQryOffer                             TraderOnRspQryTraderOfferCallback
	OnRspQrySettlementInfo                    TraderOnRspQrySettlementInfoCallback
	OnRspQryTransferBank                      TraderOnRspQryTransferBankCallback
	OnRspQryInvestorPositionDetail            TraderOnRspQryInvestorPositionDetailCallback
	OnRspQryNotice                            TraderOnRspQryNoticeCallback
	OnRspQrySettlementInfoConfirm             TraderOnRspQrySettlementInfoConfirmCallback
	OnRspQryInvestorPositionCombineDetail     TraderOnRspQryInvestorPositionCombineDetailCallback
	OnRspQryCFMMCTradingAccountKey            TraderOnRspQryCFMMCTradingAccountKeyCallback
	OnRspQryEWarrantOffset                    TraderOnRspQryEWarrantOffsetCallback
	OnRspQryInvestorProductGroupMargin        TraderOnRspQryInvestorProductGroupMarginCallback
	OnRspQryExchangeMarginRate                TraderOnRspQryExchangeMarginRateCallback
	OnRspQryExchangeMarginRateAdjust          TraderOnRspQryExchangeMarginRateAdjustCallback
	OnRspQryExchangeRate                      TraderOnRspQryExchangeRateCallback
	OnRspQrySecAgentACIDMap                   TraderOnRspQrySecAgentACIDMapCallback
	OnRspQryProductExchRate                   TraderOnRspQryProductExchRateCallback
	OnRspQryProductGroup                      TraderOnRspQryProductGroupCallback
	OnRspQryMMInstrumentCommissionRate        TraderOnRspQryMMInstrumentCommissionRateCallback
	OnRspQryMMOptionInstrCommRate             TraderOnRspQryMMOptionInstrCommRateCallback
	OnRspQryInstrumentOrderCommRate           TraderOnRspQryInstrumentOrderCommRateCallback
	OnRspQrySecAgentTradingAccount            TraderOnRspQrySecAgentTradingAccountCallback
	OnRspQrySecAgentCheckMode                 TraderOnRspQrySecAgentCheckModeCallback
	OnRspQrySecAgentTradeInfo                 TraderOnRspQrySecAgentTradeInfoCallback
	OnRspQryOptionInstrTradeCost              TraderOnRspQryOptionInstrTradeCostCallback
	OnRspQryOptionInstrCommRate               TraderOnRspQryOptionInstrCommRateCallback
	OnRspQryExecOrder                         TraderOnRspQryExecOrderCallback
	OnRspQryForQuote                          TraderOnRspQryForQuoteCallback
	OnRspQryQuote                             TraderOnRspQryQuoteCallback
	OnRspQryOptionSelfClose                   TraderOnRspQryOptionSelfCloseCallback
	OnRspQryInvestUnit                        TraderOnRspQryInvestUnitCallback
	OnRspQryCombInstrumentGuard               TraderOnRspQryCombInstrumentGuardCallback
	OnRspQryCombAction                        TraderOnRspQryCombActionCallback
	OnRspQryTransferSerial                    TraderOnRspQryTransferSerialCallback
	OnRspQryAccountregister                   TraderOnRspQryAccountregisterCallback
	OnRspError                                TraderOnRspErrorCallback
	OnRtnOrder                                TraderOnRtnOrderCallback
	OnRtnTrade                                TraderOnRtnTradeCallback
	OnErrRtnOrderInsert                       TraderOnErrRtnOrderInsertCallback
	OnErrRtnOrderAction                       TraderOnErrRtnOrderActionCallback
	OnRtnInstrumentStatus                     TraderOnRtnInstrumentStatusCallback
	OnRtnBulletin                             TraderOnRtnBulletinCallback
	OnRtnTradingNotice                        TraderOnRtnTradingNoticeCallback
	OnRtnErrorConditionalOrder                TraderOnRtnErrorConditionalOrderCallback
	OnRtnExecOrder                            TraderOnRtnExecOrderCallback
	OnErrRtnExecOrderInsert                   TraderOnErrRtnExecOrderInsertCallback
	OnErrRtnExecOrderAction                   TraderOnErrRtnExecOrderActionCallback
	OnErrRtnForQuoteInsert                    TraderOnErrRtnForQuoteInsertCallback
	OnRtnQuote                                TraderOnRtnQuoteCallback
	OnErrRtnQuoteInsert                       TraderOnErrRtnQuoteInsertCallback
	OnErrRtnQuoteAction                       TraderOnErrRtnQuoteActionCallback
	OnRtnForQuoteRsp                          TraderOnRtnForQuoteRspCallback
	OnRtnCFMMCTradingAccountToken             TraderOnRtnCFMMCTradingAccountTokenCallback
	OnErrRtnBatchOrderAction                  TraderOnErrRtnBatchOrderActionCallback
	OnRtnOptionSelfClose                      TraderOnRtnOptionSelfCloseCallback
	OnErrRtnOptionSelfCloseInsert             TraderOnErrRtnOptionSelfCloseInsertCallback
	OnErrRtnOptionSelfCloseAction             TraderOnErrRtnOptionSelfCloseActionCallback
	OnRtnCombAction                           TraderOnRtnCombActionCallback
	OnErrRtnCombActionInsert                  TraderOnErrRtnCombActionInsertCallback
	OnRspQryContractBank                      TraderOnRspQryContractBankCallback
	OnRspQryParkedOrder                       TraderOnRspQryParkedOrderCallback
	OnRspQryParkedOrderAction                 TraderOnRspQryParkedOrderActionCallback
	OnRspQryTradingNotice                     TraderOnRspQryTradingNoticeCallback
	OnRspQryBrokerTradingParams               TraderOnRspQryBrokerTradingParamsCallback
	OnRspQryBrokerTradingAlgos                TraderOnRspQryBrokerTradingAlgosCallback
	OnRspQueryCFMMCTradingAccountToken        TraderOnRspQueryCFMMCTradingAccountTokenCallback
	OnRtnFromBankToFutureByBank               TraderOnRtnFromBankToFutureByBankCallback
	OnRtnFromFutureToBankByBank               TraderOnRtnFromFutureToBankByBankCallback
	OnRtnRepealFromBankToFutureByBank         TraderOnRtnRepealFromBankToFutureByBankCallback
	OnRtnRepealFromFutureToBankByBank         TraderOnRtnRepealFromFutureToBankByBankCallback
	OnRtnFromBankToFutureByFuture             TraderOnRtnFromBankToFutureByFutureCallback
	OnRtnFromFutureToBankByFuture             TraderOnRtnFromFutureToBankByFutureCallback
	OnRtnRepealFromBankToFutureByFutureManual TraderOnRtnRepealFromBankToFutureByFutureManualCallback
	OnRtnRepealFromFutureToBankByFutureManual TraderOnRtnRepealFromFutureToBankByFutureManualCallback
	OnRtnQueryBankBalanceByFuture             TraderOnRtnQueryBankBalanceByFutureCallback
	OnErrRtnBankToFutureByFuture              TraderOnErrRtnBankToFutureByFutureCallback
	OnErrRtnFutureToBankByFuture              TraderOnErrRtnFutureToBankByFutureCallback
	OnErrRtnRepealBankToFutureByFutureManual  TraderOnErrRtnRepealBankToFutureByFutureManualCallback
	OnErrRtnRepealFutureToBankByFutureManual  TraderOnErrRtnRepealFutureToBankByFutureManualCallback
	OnErrRtnQueryBankBalanceByFuture          TraderOnErrRtnQueryBankBalanceByFutureCallback
	OnRtnRepealFromBankToFutureByFuture       TraderOnRtnRepealFromBankToFutureByFutureCallback
	OnRtnRepealFromFutureToBankByFuture       TraderOnRtnRepealFromFutureToBankByFutureCallback
	OnRspFromBankToFutureByFuture             TraderOnRspFromBankToFutureByFutureCallback
	OnRspFromFutureToBankByFuture             TraderOnRspFromFutureToBankByFutureCallback
	OnRspQueryBankAccountMoneyByFuture        TraderOnRspQueryBankAccountMoneyByFutureCallback
	OnRtnOpenAccountByBank                    TraderOnRtnOpenAccountByBankCallback
	OnRtnCancelAccountByBank                  TraderOnRtnCancelAccountByBankCallback
	OnRtnChangeAccountByBank                  TraderOnRtnChangeAccountByBankCallback
	OnRspQryClassifiedInstrument              TraderOnRspQryClassifiedInstrumentCallback
	OnRspQryCombPromotionParam                TraderOnRspQryCombPromotionParamCallback
	OnRspQryRiskSettleInvstPosition           TraderOnRspQryRiskSettleInvstPositionCallback
	OnRspQryRiskSettleProductStatus           TraderOnRspQryRiskSettleProductStatusCallback
	OnRspQrySPBMFutureParameter               TraderOnRspQrySPBMFutureParameterCallback
	OnRspQrySPBMOptionParameter               TraderOnRspQrySPBMOptionParameterCallback
	OnRspQrySPBMIntraParameter                TraderOnRspQrySPBMIntraParameterCallback
	OnRspQrySPBMInterParameter                TraderOnRspQrySPBMInterParameterCallback
	OnRspQrySPBMPortfDefinition               TraderOnRspQrySPBMPortfDefinitionCallback
	OnRspQrySPBMInvestorPortfDef              TraderOnRspQrySPBMInvestorPortfDefCallback
	OnRspQryInvestorPortfMarginRatio          TraderOnRspQryInvestorPortfMarginRatioCallback
	OnRspQryInvestorProdSPBMDetail            TraderOnRspQryInvestorProdSPBMDetailCallback
	OnRspQryInvestorCommoditySPMMMargin       TraderOnRspQryInvestorCommoditySPMMMarginCallback
	OnRspQryInvestorCommodityGroupSPMMMargin  TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback
	OnRspQrySPMMInstParam                     TraderOnRspQrySPMMInstParamCallback
	OnRspQrySPMMProductParam                  TraderOnRspQrySPMMProductParamCallback
	OnRspQrySPBMAddOnInterParameter           TraderOnRspQrySPBMAddOnInterParameterCallback
	OnRspQryRCAMSCombProductInfo              TraderOnRspQryRCAMSCombProductInfoCallback
	OnRspQryRCAMSInstrParameter               TraderOnRspQryRCAMSInstrParameterCallback
	OnRspQryRCAMSIntraParameter               TraderOnRspQryRCAMSIntraParameterCallback
	OnRspQryRCAMSInterParameter               TraderOnRspQryRCAMSInterParameterCallback
	OnRspQryRCAMSShortOptAdjustParam          TraderOnRspQryRCAMSShortOptAdjustParamCallback
	OnRspQryRCAMSInvestorCombPosition         TraderOnRspQryRCAMSInvestorCombPositionCallback
	OnRspQryInvestorProdRCAMSMargin           TraderOnRspQryInvestorProdRCAMSMarginCallback
	OnRspQryRULEInstrParameter                TraderOnRspQryRULEInstrParameterCallback
	OnRspQryRULEIntraParameter                TraderOnRspQryRULEIntraParameterCallback
	OnRspQryRULEInterParameter                TraderOnRspQryRULEInterParameterCallback
	OnRspQryInvestorProdRULEMargin            TraderOnRspQryInvestorProdRULEMarginCallback
	OnRspQryInvestorPortfSetting              TraderOnRspQryInvestorPortfSettingCallback
}

// ========== TraderSpi 接口 ==========

// TraderSpi 交易回调接口
type TraderSpi interface {
	OnFrontConnected()                                                                                                                                                                             // ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontDisconnected(nReason int32)                                                                                                                                                             // 0x2003 收到错误报文
	OnHeartBeatWarning(nTimeLapse int32)                                                                                                                                                           // 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnRspAuthenticate(pRspAuthenticateField *CThostFtdcRspAuthenticateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                     // 客户端认证响应
	OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                   // 登录请求响应
	OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                      // 登出请求响应
	OnRspUserPasswordUpdate(pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                              // 用户口令更新请求响应
	OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                // 资金账户口令更新请求响应
	OnRspUserAuthMethod(pRspUserAuthMethod *CThostFtdcRspUserAuthMethodField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 查询用户当前支持的认证模式的回复
	OnRspGenUserCaptcha(pRspGenUserCaptcha *CThostFtdcRspGenUserCaptchaField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 获取图形验证码请求的回复
	OnRspGenUserText(pRspGenUserText *CThostFtdcRspGenUserTextField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 获取短信验证码请求的回复
	OnRspOrderInsert(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                     // 报单录入请求响应
	OnRspParkedOrderInsert(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 预埋单录入请求响应
	OnRspParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                 // 预埋撤单录入请求响应
	OnRspOrderAction(pInputOrderAction *CThostFtdcInputOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                         // 报单操作请求响应
	OnRspQryMaxOrderVolume(pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                 // 查询最大报单数量响应
	OnRspSettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // 投资者结算结果确认响应
	OnRspRemoveParkedOrder(pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                 // 删除预埋单响应
	OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 删除预埋撤单响应
	OnRspExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                         // 执行宣告录入请求响应
	OnRspExecOrderAction(pInputExecOrderAction *CThostFtdcInputExecOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                             // 执行宣告操作请求响应
	OnRspForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                            // 询价录入请求响应
	OnRspQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                     // 报价录入请求响应
	OnRspQuoteAction(pInputQuoteAction *CThostFtdcInputQuoteActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                         // 报价操作请求响应
	OnRspBatchOrderAction(pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                          // 批量报单操作请求响应
	OnRspOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                       // 期权自对冲录入请求响应
	OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                           // 期权自对冲操作请求响应
	OnRspCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                      // 申请组合录入请求响应
	OnRspQryOrder(pOrder *CThostFtdcOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                                  // 请求查询报单响应
	OnRspQryTrade(pTrade *CThostFtdcTradeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                                  // 请求查询成交响应
	OnRspQryInvestorPosition(pInvestorPosition *CThostFtdcInvestorPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                 // 请求查询投资者持仓响应
	OnRspQryTradingAccount(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                       // 请求查询资金账户响应
	OnRspQryInvestor(pInvestor *CThostFtdcInvestorField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                         // 请求查询投资者响应
	OnRspQryTradingCode(pTradingCode *CThostFtdcTradingCodeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                // 请求查询交易编码响应
	OnRspQryInstrumentMarginRate(pInstrumentMarginRate *CThostFtdcInstrumentMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // 请求查询合约保证金率响应
	OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate *CThostFtdcInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                         // 请求查询合约手续费率响应
	OnRspQryExchange(pExchange *CThostFtdcExchangeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                         // 请求查询交易所响应
	OnRspQryProduct(pProduct *CThostFtdcProductField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                            // 请求查询产品响应
	OnRspQryInstrument(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                   // 请求查询合约响应
	OnRspQryDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 请求查询行情响应
	OnRspQryTraderOffer(pTraderOffer *CThostFtdcTraderOfferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                // 请求查询交易员报盘机响应
	OnRspQrySettlementInfo(pSettlementInfo *CThostFtdcSettlementInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                       // 请求查询投资者结算结果响应
	OnRspQryTransferBank(pTransferBank *CThostFtdcTransferBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 请求查询转帐银行响应
	OnRspQryInvestorPositionDetail(pInvestorPositionDetail *CThostFtdcInvestorPositionDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 请求查询投资者持仓明细响应
	OnRspQryNotice(pNotice *CThostFtdcNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                               // 请求查询客户通知响应
	OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                  // 请求查询结算信息确认响应
	OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail *CThostFtdcInvestorPositionCombineDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)          // 请求查询投资者持仓明细响应
	OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey *CThostFtdcCFMMCTradingAccountKeyField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 查询保证金监管系统经纪公司资金账户密钥响应
	OnRspQryEWarrantOffset(pEWarrantOffset *CThostFtdcEWarrantOffsetField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                       // 请求查询仓单折抵信息响应
	OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin *CThostFtdcInvestorProductGroupMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                   // 请求查询投资者品种/跨品种保证金响应
	OnRspQryExchangeMarginRate(pExchangeMarginRate *CThostFtdcExchangeMarginRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // 请求查询交易所保证金率响应
	OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust *CThostFtdcExchangeMarginRateAdjustField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                         // 请求查询交易所调整保证金率响应
	OnRspQryExchangeRate(pExchangeRate *CThostFtdcExchangeRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 请求查询汇率响应
	OnRspQrySecAgentACIDMap(pSecAgentACIDMap *CThostFtdcSecAgentACIDMapField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 请求查询二级代理操作员银期权限响应
	OnRspQryProductExchRate(pProductExchRate *CThostFtdcProductExchRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 请求查询产品报价汇率
	OnRspQryProductGroup(pProductGroup *CThostFtdcProductGroupField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 请求查询产品组
	OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate *CThostFtdcMMInstrumentCommissionRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                   // 请求查询做市商合约手续费率响应
	OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate *CThostFtdcMMOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                  // 请求查询做市商期权合约手续费响应
	OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate *CThostFtdcInstrumentOrderCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // 请求查询报单手续费响应
	OnRspQrySecAgentTradingAccount(pTradingAccount *CThostFtdcTradingAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                               // 请求查询资金账户响应
	OnRspQrySecAgentCheckMode(pSecAgentCheckMode *CThostFtdcSecAgentCheckModeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                              // 请求查询二级代理商资金校验模式响应
	OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo *CThostFtdcSecAgentTradeInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                              // 请求查询二级代理商信息响应
	OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost *CThostFtdcOptionInstrTradeCostField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // 请求查询期权交易成本响应
	OnRspQryOptionInstrCommRate(pOptionInstrCommRate *CThostFtdcOptionInstrCommRateField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // 请求查询期权合约手续费响应
	OnRspQryExecOrder(pExecOrder *CThostFtdcExecOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                      // 请求查询执行宣告响应
	OnRspQryForQuote(pForQuote *CThostFtdcForQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                         // 请求查询询价响应
	OnRspQryQuote(pQuote *CThostFtdcQuoteField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                                  // 请求查询报价响应
	OnRspQryOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 请求查询期权自对冲响应
	OnRspQryInvestUnit(pInvestUnit *CThostFtdcInvestUnitField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                   // 请求查询投资单元响应
	OnRspQryCombInstrumentGuard(pCombInstrumentGuard *CThostFtdcCombInstrumentGuardField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // 请求查询组合合约安全系数响应
	OnRspQryCombAction(pCombAction *CThostFtdcCombActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                   // 请求查询申请组合响应
	OnRspQryTransferSerial(pTransferSerial *CThostFtdcTransferSerialField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                       // 请求查询转帐流水响应
	OnRspQryAccountregister(pAccountregister *CThostFtdcAccountregisterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                    // 请求查询银期签约关系响应
	OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                                                                   // 错误应答
	OnRtnOrder(pOrder *CThostFtdcOrderField)                                                                                                                                                       // 报单通知
	OnRtnTrade(pTrade *CThostFtdcTradeField)                                                                                                                                                       // 成交通知
	OnErrRtnOrderInsert(pInputOrder *CThostFtdcInputOrderField, pRspInfo *CThostFtdcRspInfoField)                                                                                                  // 报单录入错误回报
	OnErrRtnOrderAction(pOrderAction *CThostFtdcOrderActionField, pRspInfo *CThostFtdcRspInfoField)                                                                                                // 报单操作错误回报
	OnRtnInstrumentStatus(pInstrumentStatus *CThostFtdcInstrumentStatusField)                                                                                                                      // 合约交易状态通知
	OnRtnBulletin(pBulletin *CThostFtdcBulletinField)                                                                                                                                              // 交易所公告通知
	OnRtnTradingNotice(pTradingNoticeInfo *CThostFtdcTradingNoticeInfoField)                                                                                                                       // 交易通知
	OnRtnErrorConditionalOrder(pErrorConditionalOrder *CThostFtdcErrorConditionalOrderField)                                                                                                       // 提示条件单校验错误
	OnRtnExecOrder(pExecOrder *CThostFtdcExecOrderField)                                                                                                                                           // 执行宣告通知
	OnErrRtnExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, pRspInfo *CThostFtdcRspInfoField)                                                                                      // 执行宣告录入错误回报
	OnErrRtnExecOrderAction(pExecOrderAction *CThostFtdcExecOrderActionField, pRspInfo *CThostFtdcRspInfoField)                                                                                    // 执行宣告操作错误回报
	OnErrRtnForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, pRspInfo *CThostFtdcRspInfoField)                                                                                         // 询价录入错误回报
	OnRtnQuote(pQuote *CThostFtdcQuoteField)                                                                                                                                                       // 报价通知
	OnErrRtnQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, pRspInfo *CThostFtdcRspInfoField)                                                                                                  // 报价录入错误回报
	OnErrRtnQuoteAction(pQuoteAction *CThostFtdcQuoteActionField, pRspInfo *CThostFtdcRspInfoField)                                                                                                // 报价操作错误回报
	OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField)                                                                                                                                     // 询价通知
	OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken *CThostFtdcCFMMCTradingAccountTokenField)                                                                                              // 保证金监控中心用户令牌
	OnErrRtnBatchOrderAction(pBatchOrderAction *CThostFtdcBatchOrderActionField, pRspInfo *CThostFtdcRspInfoField)                                                                                 // 批量报单操作错误回报
	OnRtnOptionSelfClose(pOptionSelfClose *CThostFtdcOptionSelfCloseField)                                                                                                                         // 期权自对冲通知
	OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, pRspInfo *CThostFtdcRspInfoField)                                                                    // 期权自对冲录入错误回报
	OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction *CThostFtdcOptionSelfCloseActionField, pRspInfo *CThostFtdcRspInfoField)                                                                  // 期权自对冲操作错误回报
	OnRtnCombAction(pCombAction *CThostFtdcCombActionField)                                                                                                                                        // 申请组合通知
	OnErrRtnCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, pRspInfo *CThostFtdcRspInfoField)                                                                                   // 申请组合录入错误回报
	OnRspQryContractBank(pContractBank *CThostFtdcContractBankField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                             // 请求查询签约银行响应
	OnRspQryParkedOrder(pParkedOrder *CThostFtdcParkedOrderField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                // 请求查询预埋单响应
	OnRspQryParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                              // 请求查询预埋撤单响应
	OnRspQryTradingNotice(pTradingNotice *CThostFtdcTradingNoticeField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                          // 请求查询交易通知响应
	OnRspQryBrokerTradingParams(pBrokerTradingParams *CThostFtdcBrokerTradingParamsField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // 请求查询经纪公司交易参数响应
	OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos *CThostFtdcBrokerTradingAlgosField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // 请求查询经纪公司交易算法响应
	OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)             // 请求查询监控中心用户令牌
	OnRtnFromBankToFutureByBank(pRspTransfer *CThostFtdcRspTransferField)                                                                                                                          // 银行发起银行资金转期货通知
	OnRtnFromFutureToBankByBank(pRspTransfer *CThostFtdcRspTransferField)                                                                                                                          // 银行发起期货资金转银行通知
	OnRtnRepealFromBankToFutureByBank(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                        // 银行发起冲正银行转期货通知
	OnRtnRepealFromFutureToBankByBank(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                        // 银行发起冲正期货转银行通知
	OnRtnFromBankToFutureByFuture(pRspTransfer *CThostFtdcRspTransferField)                                                                                                                        // 期货发起银行资金转期货通知
	OnRtnFromFutureToBankByFuture(pRspTransfer *CThostFtdcRspTransferField)                                                                                                                        // 期货发起期货资金转银行通知
	OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                // 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                // 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount *CThostFtdcNotifyQueryAccountField)                                                                                                          // 期货发起查询银行余额通知
	OnErrRtnBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)                                                                                       // 期货发起银行资金转期货错误回报
	OnErrRtnFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField)                                                                                       // 期货发起期货资金转银行错误回报
	OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)                                                                               // 系统运行时期货端手工发起冲正银行转期货错误回报
	OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal *CThostFtdcReqRepealField, pRspInfo *CThostFtdcRspInfoField)                                                                               // 系统运行时期货端手工发起冲正期货转银行错误回报
	OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField)                                                                           // 期货发起查询银行余额错误回报
	OnRtnRepealFromBankToFutureByFuture(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                      // 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
	OnRtnRepealFromFutureToBankByFuture(pRspRepeal *CThostFtdcRspRepealField)                                                                                                                      // 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
	OnRspFromBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                      // 期货发起银行资金转期货应答
	OnRspFromFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                      // 期货发起期货资金转银行应答
	OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                         // 期货发起查询银行余额应答
	OnRtnOpenAccountByBank(pOpenAccount *CThostFtdcOpenAccountField)                                                                                                                               // 银行发起银期开户通知
	OnRtnCancelAccountByBank(pCancelAccount *CThostFtdcCancelAccountField)                                                                                                                         // 银行发起银期销户通知
	OnRtnChangeAccountByBank(pChangeAccount *CThostFtdcChangeAccountField)                                                                                                                         // 银行发起变更银行账号通知
	OnRspQryClassifiedInstrument(pInstrument *CThostFtdcInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                         // 请求查询分类合约响应
	OnRspQryCombPromotionParam(pCombPromotionParam *CThostFtdcCombPromotionParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // 请求组合优惠比例响应
	OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition *CThostFtdcRiskSettleInvstPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // 投资者风险结算持仓查询响应
	OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus *CThostFtdcRiskSettleProductStatusField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // 风险结算产品查询响应
	OnRspQrySPBMFutureParameter(pSPBMFutureParameter *CThostFtdcSPBMFutureParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // SPBM期货合约参数查询响应
	OnRspQrySPBMOptionParameter(pSPBMOptionParameter *CThostFtdcSPBMOptionParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // SPBM期权合约参数查询响应
	OnRspQrySPBMIntraParameter(pSPBMIntraParameter *CThostFtdcSPBMIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // SPBM品种内对锁仓折扣参数查询响应
	OnRspQrySPBMInterParameter(pSPBMInterParameter *CThostFtdcSPBMInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // SPBM跨品种抵扣参数查询响应
	OnRspQrySPBMPortfDefinition(pSPBMPortfDefinition *CThostFtdcSPBMPortfDefinitionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // SPBM组合保证金套餐查询响应
	OnRspQrySPBMInvestorPortfDef(pSPBMInvestorPortfDef *CThostFtdcSPBMInvestorPortfDefField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // 投资者SPBM套餐选择查询响应
	OnRspQryInvestorPortfMarginRatio(pInvestorPortfMarginRatio *CThostFtdcInvestorPortfMarginRatioField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                         // 投资者新型组合保证金系数查询响应
	OnRspQryInvestorProdSPBMDetail(pInvestorProdSPBMDetail *CThostFtdcInvestorProdSPBMDetailField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 投资者产品SPBM明细查询响应
	OnRspQryInvestorCommoditySPMMMargin(pInvestorCommoditySPMMMargin *CThostFtdcInvestorCommoditySPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                // 投资者商品组SPMM记录查询响应
	OnRspQryInvestorCommodityGroupSPMMMargin(pInvestorCommodityGroupSPMMMargin *CThostFtdcInvestorCommodityGroupSPMMMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 投资者商品群SPMM记录查询响应
	OnRspQrySPMMInstParam(pSPMMInstParam *CThostFtdcSPMMInstParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                          // SPMM合约参数查询响应
	OnRspQrySPMMProductParam(pSPMMProductParam *CThostFtdcSPMMProductParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                 // SPMM产品参数查询响应
	OnRspQrySPBMAddOnInterParameter(pSPBMAddOnInterParameter *CThostFtdcSPBMAddOnInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // SPBM附加跨品种抵扣参数查询响应
	OnRspQryRCAMSCombProductInfo(pRCAMSCombProductInfo *CThostFtdcRCAMSCombProductInfoField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // RCAMS产品组合信息查询响应
	OnRspQryRCAMSInstrParameter(pRCAMSInstrParameter *CThostFtdcRCAMSInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // RCAMS同合约风险对冲参数查询响应
	OnRspQryRCAMSIntraParameter(pRCAMSIntraParameter *CThostFtdcRCAMSIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // RCAMS品种内风险对冲参数查询响应
	OnRspQryRCAMSInterParameter(pRCAMSInterParameter *CThostFtdcRCAMSInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                        // RCAMS跨品种风险折抵参数查询响应
	OnRspQryRCAMSShortOptAdjustParam(pRCAMSShortOptAdjustParam *CThostFtdcRCAMSShortOptAdjustParamField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                         // RCAMS空头期权风险调整参数查询响应
	OnRspQryRCAMSInvestorCombPosition(pRCAMSInvestorCombPosition *CThostFtdcRCAMSInvestorCombPositionField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                      // RCAMS策略组合持仓查询响应
	OnRspQryInvestorProdRCAMSMargin(pInvestorProdRCAMSMargin *CThostFtdcInvestorProdRCAMSMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // 投资者品种RCAMS保证金查询响应
	OnRspQryRULEInstrParameter(pRULEInstrParameter *CThostFtdcRULEInstrParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // RULE合约保证金参数查询响应
	OnRspQryRULEIntraParameter(pRULEIntraParameter *CThostFtdcRULEIntraParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // RULE品种内对锁仓折扣参数查询响应
	OnRspQryRULEInterParameter(pRULEInterParameter *CThostFtdcRULEInterParameterField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                           // RULE跨品种抵扣参数查询响应
	OnRspQryInvestorProdRULEMargin(pInvestorProdRULEMargin *CThostFtdcInvestorProdRULEMarginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 投资者产品RULE保证金查询响应
	OnRspQryInvestorPortfSetting(pInvestorPortfSetting *CThostFtdcInvestorPortfSettingField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                     // 投资者投资者新组保设置查询响应
}

// ========== TraderApi 结构体 ==========

// TraderApi 交易 API 封装
type TraderApi struct {
	handle    uintptr
	spi       TraderSpi
	spiHandle uintptr // C SPI 实例句柄
	userData  uintptr
	mu        sync.RWMutex
	flowPath  []byte // 保存 flowPath 的 C 字符串，防止被 GC 回收
}

// ========== C 函数声明 ==========

var (
	traderOnce sync.Once

	_TraderCreateFtdcTraderApi                             func(*byte) uintptr
	_TraderGetApiVersion                                   func() *byte
	_TraderRelease                                         func(uintptr)
	_TraderInit                                            func(uintptr)
	_TraderJoin                                            func(uintptr) int32
	_TraderGetTradingDay                                   func(uintptr) *byte
	_TraderGetFrontInfo                                    func(uintptr, *CThostFtdcFrontInfoField)
	_TraderRegisterFront                                   func(uintptr, *byte)
	_TraderRegisterNameServer                              func(uintptr, *byte)
	_TraderRegisterFensUserInfo                            func(uintptr, *CThostFtdcFensUserInfoField)
	_TraderSubscribePrivateTopic                           func(uintptr, THOST_TE_RESUME_TYPE)
	_TraderSubscribePublicTopic                            func(uintptr, THOST_TE_RESUME_TYPE)
	_TraderReqAuthenticate                                 func(uintptr, *CThostFtdcReqAuthenticateField, int32) int32
	_TraderRegisterUserSystemInfo                          func(uintptr, *CThostFtdcUserSystemInfoField) int32
	_TraderSubmitUserSystemInfo                            func(uintptr, *CThostFtdcUserSystemInfoField) int32
	_TraderReqUserLogin                                    func(uintptr, *CThostFtdcReqUserLoginField, int32) int32
	_TraderReqUserLogout                                   func(uintptr, *CThostFtdcUserLogoutField, int32) int32
	_TraderReqUserPasswordUpdate                           func(uintptr, *CThostFtdcUserPasswordUpdateField, int32) int32
	_TraderReqTradingAccountPasswordUpdate                 func(uintptr, *CThostFtdcTradingAccountPasswordUpdateField, int32) int32
	_TraderReqUserAuthMethod                               func(uintptr, *CThostFtdcReqUserAuthMethodField, int32) int32
	_TraderReqGenUserCaptcha                               func(uintptr, *CThostFtdcReqGenUserCaptchaField, int32) int32
	_TraderReqGenUserText                                  func(uintptr, *CThostFtdcReqGenUserTextField, int32) int32
	_TraderReqUserLoginWithCaptcha                         func(uintptr, *CThostFtdcReqUserLoginWithCaptchaField, int32) int32
	_TraderReqUserLoginWithText                            func(uintptr, *CThostFtdcReqUserLoginWithTextField, int32) int32
	_TraderReqUserLoginWithOTP                             func(uintptr, *CThostFtdcReqUserLoginWithOTPField, int32) int32
	_TraderReqOrderInsert                                  func(uintptr, *CThostFtdcInputOrderField, int32) int32
	_TraderReqParkedOrderInsert                            func(uintptr, *CThostFtdcParkedOrderField, int32) int32
	_TraderReqParkedOrderAction                            func(uintptr, *CThostFtdcParkedOrderActionField, int32) int32
	_TraderReqOrderAction                                  func(uintptr, *CThostFtdcInputOrderActionField, int32) int32
	_TraderReqQryMaxOrderVolume                            func(uintptr, *CThostFtdcQryMaxOrderVolumeField, int32) int32
	_TraderReqSettlementInfoConfirm                        func(uintptr, *CThostFtdcSettlementInfoConfirmField, int32) int32
	_TraderReqRemoveParkedOrder                            func(uintptr, *CThostFtdcRemoveParkedOrderField, int32) int32
	_TraderReqRemoveParkedOrderAction                      func(uintptr, *CThostFtdcRemoveParkedOrderActionField, int32) int32
	_TraderReqExecOrderInsert                              func(uintptr, *CThostFtdcInputExecOrderField, int32) int32
	_TraderReqExecOrderAction                              func(uintptr, *CThostFtdcInputExecOrderActionField, int32) int32
	_TraderReqForQuoteInsert                               func(uintptr, *CThostFtdcInputForQuoteField, int32) int32
	_TraderReqQuoteInsert                                  func(uintptr, *CThostFtdcInputQuoteField, int32) int32
	_TraderReqQuoteAction                                  func(uintptr, *CThostFtdcInputQuoteActionField, int32) int32
	_TraderReqBatchOrderAction                             func(uintptr, *CThostFtdcInputBatchOrderActionField, int32) int32
	_TraderReqOptionSelfCloseInsert                        func(uintptr, *CThostFtdcInputOptionSelfCloseField, int32) int32
	_TraderReqOptionSelfCloseAction                        func(uintptr, *CThostFtdcInputOptionSelfCloseActionField, int32) int32
	_TraderReqCombActionInsert                             func(uintptr, *CThostFtdcInputCombActionField, int32) int32
	_TraderReqQryOrder                                     func(uintptr, *CThostFtdcQryOrderField, int32) int32
	_TraderReqQryTrade                                     func(uintptr, *CThostFtdcQryTradeField, int32) int32
	_TraderReqQryInvestorPosition                          func(uintptr, *CThostFtdcQryInvestorPositionField, int32) int32
	_TraderReqQryTradingAccount                            func(uintptr, *CThostFtdcQryTradingAccountField, int32) int32
	_TraderReqQryInvestor                                  func(uintptr, *CThostFtdcQryInvestorField, int32) int32
	_TraderReqQryTradingCode                               func(uintptr, *CThostFtdcQryTradingCodeField, int32) int32
	_TraderReqQryInstrumentMarginRate                      func(uintptr, *CThostFtdcQryInstrumentMarginRateField, int32) int32
	_TraderReqQryInstrumentCommissionRate                  func(uintptr, *CThostFtdcQryInstrumentCommissionRateField, int32) int32
	_TraderReqQryExchange                                  func(uintptr, *CThostFtdcQryExchangeField, int32) int32
	_TraderReqQryProduct                                   func(uintptr, *CThostFtdcQryProductField, int32) int32
	_TraderReqQryInstrument                                func(uintptr, *CThostFtdcQryInstrumentField, int32) int32
	_TraderReqQryDepthMarketData                           func(uintptr, *CThostFtdcQryDepthMarketDataField, int32) int32
	_TraderReqQryTraderOffer                               func(uintptr, *CThostFtdcQryTraderOfferField, int32) int32
	_TraderReqQrySettlementInfo                            func(uintptr, *CThostFtdcQrySettlementInfoField, int32) int32
	_TraderReqQryTransferBank                              func(uintptr, *CThostFtdcQryTransferBankField, int32) int32
	_TraderReqQryInvestorPositionDetail                    func(uintptr, *CThostFtdcQryInvestorPositionDetailField, int32) int32
	_TraderReqQryNotice                                    func(uintptr, *CThostFtdcQryNoticeField, int32) int32
	_TraderReqQrySettlementInfoConfirm                     func(uintptr, *CThostFtdcQrySettlementInfoConfirmField, int32) int32
	_TraderReqQryInvestorPositionCombineDetail             func(uintptr, *CThostFtdcQryInvestorPositionCombineDetailField, int32) int32
	_TraderReqQryCFMMCTradingAccountKey                    func(uintptr, *CThostFtdcQryCFMMCTradingAccountKeyField, int32) int32
	_TraderReqQryEWarrantOffset                            func(uintptr, *CThostFtdcQryEWarrantOffsetField, int32) int32
	_TraderReqQryInvestorProductGroupMargin                func(uintptr, *CThostFtdcQryInvestorProductGroupMarginField, int32) int32
	_TraderReqQryExchangeMarginRate                        func(uintptr, *CThostFtdcQryExchangeMarginRateField, int32) int32
	_TraderReqQryExchangeMarginRateAdjust                  func(uintptr, *CThostFtdcQryExchangeMarginRateAdjustField, int32) int32
	_TraderReqQryExchangeRate                              func(uintptr, *CThostFtdcQryExchangeRateField, int32) int32
	_TraderReqQrySecAgentACIDMap                           func(uintptr, *CThostFtdcQrySecAgentACIDMapField, int32) int32
	_TraderReqQryProductExchRate                           func(uintptr, *CThostFtdcQryProductExchRateField, int32) int32
	_TraderReqQryProductGroup                              func(uintptr, *CThostFtdcQryProductGroupField, int32) int32
	_TraderReqQryMMInstrumentCommissionRate                func(uintptr, *CThostFtdcQryMMInstrumentCommissionRateField, int32) int32
	_TraderReqQryMMOptionInstrCommRate                     func(uintptr, *CThostFtdcQryMMOptionInstrCommRateField, int32) int32
	_TraderReqQryInstrumentOrderCommRate                   func(uintptr, *CThostFtdcQryInstrumentOrderCommRateField, int32) int32
	_TraderReqQrySecAgentTradingAccount                    func(uintptr, *CThostFtdcQryTradingAccountField, int32) int32
	_TraderReqQrySecAgentCheckMode                         func(uintptr, *CThostFtdcQrySecAgentCheckModeField, int32) int32
	_TraderReqQrySecAgentTradeInfo                         func(uintptr, *CThostFtdcQrySecAgentTradeInfoField, int32) int32
	_TraderReqQryOptionInstrTradeCost                      func(uintptr, *CThostFtdcQryOptionInstrTradeCostField, int32) int32
	_TraderReqQryOptionInstrCommRate                       func(uintptr, *CThostFtdcQryOptionInstrCommRateField, int32) int32
	_TraderReqQryExecOrder                                 func(uintptr, *CThostFtdcQryExecOrderField, int32) int32
	_TraderReqQryForQuote                                  func(uintptr, *CThostFtdcQryForQuoteField, int32) int32
	_TraderReqQryQuote                                     func(uintptr, *CThostFtdcQryQuoteField, int32) int32
	_TraderReqQryOptionSelfClose                           func(uintptr, *CThostFtdcQryOptionSelfCloseField, int32) int32
	_TraderReqQryInvestUnit                                func(uintptr, *CThostFtdcQryInvestUnitField, int32) int32
	_TraderReqQryCombInstrumentGuard                       func(uintptr, *CThostFtdcQryCombInstrumentGuardField, int32) int32
	_TraderReqQryCombAction                                func(uintptr, *CThostFtdcQryCombActionField, int32) int32
	_TraderReqQryTransferSerial                            func(uintptr, *CThostFtdcQryTransferSerialField, int32) int32
	_TraderReqQryAccountregister                           func(uintptr, *CThostFtdcQryAccountregisterField, int32) int32
	_TraderReqQryContractBank                              func(uintptr, *CThostFtdcQryContractBankField, int32) int32
	_TraderReqQryParkedOrder                               func(uintptr, *CThostFtdcQryParkedOrderField, int32) int32
	_TraderReqQryParkedOrderAction                         func(uintptr, *CThostFtdcQryParkedOrderActionField, int32) int32
	_TraderReqQryTradingNotice                             func(uintptr, *CThostFtdcQryTradingNoticeField, int32) int32
	_TraderReqQryBrokerTradingParams                       func(uintptr, *CThostFtdcQryBrokerTradingParamsField, int32) int32
	_TraderReqQryBrokerTradingAlgos                        func(uintptr, *CThostFtdcQryBrokerTradingAlgosField, int32) int32
	_TraderReqQueryCFMMCTradingAccountToken                func(uintptr, *CThostFtdcQueryCFMMCTradingAccountTokenField, int32) int32
	_TraderReqFromBankToFutureByFuture                     func(uintptr, *CThostFtdcReqTransferField, int32) int32
	_TraderReqFromFutureToBankByFuture                     func(uintptr, *CThostFtdcReqTransferField, int32) int32
	_TraderReqQueryBankAccountMoneyByFuture                func(uintptr, *CThostFtdcReqQueryAccountField, int32) int32
	_TraderReqQryClassifiedInstrument                      func(uintptr, *CThostFtdcQryClassifiedInstrumentField, int32) int32
	_TraderReqQryCombPromotionParam                        func(uintptr, *CThostFtdcQryCombPromotionParamField, int32) int32
	_TraderReqQryRiskSettleInvstPosition                   func(uintptr, *CThostFtdcQryRiskSettleInvstPositionField, int32) int32
	_TraderReqQryRiskSettleProductStatus                   func(uintptr, *CThostFtdcQryRiskSettleProductStatusField, int32) int32
	_TraderReqQrySPBMFutureParameter                       func(uintptr, *CThostFtdcQrySPBMFutureParameterField, int32) int32
	_TraderReqQrySPBMOptionParameter                       func(uintptr, *CThostFtdcQrySPBMOptionParameterField, int32) int32
	_TraderReqQrySPBMIntraParameter                        func(uintptr, *CThostFtdcQrySPBMIntraParameterField, int32) int32
	_TraderReqQrySPBMInterParameter                        func(uintptr, *CThostFtdcQrySPBMInterParameterField, int32) int32
	_TraderReqQrySPBMPortfDefinition                       func(uintptr, *CThostFtdcQrySPBMPortfDefinitionField, int32) int32
	_TraderReqQrySPBMInvestorPortfDef                      func(uintptr, *CThostFtdcQrySPBMInvestorPortfDefField, int32) int32
	_TraderReqQryInvestorPortfMarginRatio                  func(uintptr, *CThostFtdcQryInvestorPortfMarginRatioField, int32) int32
	_TraderReqQryInvestorProdSPBMDetail                    func(uintptr, *CThostFtdcQryInvestorProdSPBMDetailField, int32) int32
	_TraderReqQryInvestorCommoditySPMMMargin               func(uintptr, *CThostFtdcQryInvestorCommoditySPMMMarginField, int32) int32
	_TraderReqQryInvestorCommodityGroupSPMMMargin          func(uintptr, *CThostFtdcQryInvestorCommodityGroupSPMMMarginField, int32) int32
	_TraderReqQrySPMMInstParam                             func(uintptr, *CThostFtdcQrySPMMInstParamField, int32) int32
	_TraderReqQrySPMMProductParam                          func(uintptr, *CThostFtdcQrySPMMProductParamField, int32) int32
	_TraderReqQrySPBMAddOnInterParameter                   func(uintptr, *CThostFtdcQrySPBMAddOnInterParameterField, int32) int32
	_TraderReqQryRCAMSCombProductInfo                      func(uintptr, *CThostFtdcQryRCAMSCombProductInfoField, int32) int32
	_TraderReqQryRCAMSInstrParameter                       func(uintptr, *CThostFtdcQryRCAMSInstrParameterField, int32) int32
	_TraderReqQryRCAMSIntraParameter                       func(uintptr, *CThostFtdcQryRCAMSIntraParameterField, int32) int32
	_TraderReqQryRCAMSInterParameter                       func(uintptr, *CThostFtdcQryRCAMSInterParameterField, int32) int32
	_TraderReqQryRCAMSShortOptAdjustParam                  func(uintptr, *CThostFtdcQryRCAMSShortOptAdjustParamField, int32) int32
	_TraderReqQryRCAMSInvestorCombPosition                 func(uintptr, *CThostFtdcQryRCAMSInvestorCombPositionField, int32) int32
	_TraderReqQryInvestorProdRCAMSMargin                   func(uintptr, *CThostFtdcQryInvestorProdRCAMSMarginField, int32) int32
	_TraderReqQryRULEInstrParameter                        func(uintptr, *CThostFtdcQryRULEInstrParameterField, int32) int32
	_TraderReqQryRULEIntraParameter                        func(uintptr, *CThostFtdcQryRULEIntraParameterField, int32) int32
	_TraderReqQryRULEInterParameter                        func(uintptr, *CThostFtdcQryRULEInterParameterField, int32) int32
	_TraderReqQryInvestorProdRULEMargin                    func(uintptr, *CThostFtdcQryInvestorProdRULEMarginField, int32) int32
	_TraderReqQryInvestorPortfSetting                      func(uintptr, *CThostFtdcQryInvestorPortfSettingField, int32) int32
	_TraderSpiCreate                                       func(uintptr) uintptr
	_TraderSpiDestroy                                      func(uintptr)
	_TraderRegisterSpi                                     func(uintptr, uintptr)
	_TraderSpiSetCallbacks                                 func(uintptr, *TraderSpiCallbacks)
	_TraderSpiSetOnFrontConnected                          func(uintptr, uintptr)
	_TraderSpiSetOnFrontDisconnected                       func(uintptr, uintptr)
	_TraderSpiSetOnHeartBeatWarning                        func(uintptr, uintptr)
	_TraderSpiSetOnRspAuthenticate                         func(uintptr, uintptr)
	_TraderSpiSetOnRspUserLogin                            func(uintptr, uintptr)
	_TraderSpiSetOnRspUserLogout                           func(uintptr, uintptr)
	_TraderSpiSetOnRspUserPasswordUpdate                   func(uintptr, uintptr)
	_TraderSpiSetOnRspTradingAccountPasswordUpdate         func(uintptr, uintptr)
	_TraderSpiSetOnRspUserAuthMethod                       func(uintptr, uintptr)
	_TraderSpiSetOnRspGenUserCaptcha                       func(uintptr, uintptr)
	_TraderSpiSetOnRspGenUserText                          func(uintptr, uintptr)
	_TraderSpiSetOnRspOrderInsert                          func(uintptr, uintptr)
	_TraderSpiSetOnRspParkedOrderInsert                    func(uintptr, uintptr)
	_TraderSpiSetOnRspParkedOrderAction                    func(uintptr, uintptr)
	_TraderSpiSetOnRspOrderAction                          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryMaxOrderVolume                    func(uintptr, uintptr)
	_TraderSpiSetOnRspSettlementInfoConfirm                func(uintptr, uintptr)
	_TraderSpiSetOnRspRemoveParkedOrder                    func(uintptr, uintptr)
	_TraderSpiSetOnRspRemoveParkedOrderAction              func(uintptr, uintptr)
	_TraderSpiSetOnRspExecOrderInsert                      func(uintptr, uintptr)
	_TraderSpiSetOnRspExecOrderAction                      func(uintptr, uintptr)
	_TraderSpiSetOnRspForQuoteInsert                       func(uintptr, uintptr)
	_TraderSpiSetOnRspQuoteInsert                          func(uintptr, uintptr)
	_TraderSpiSetOnRspQuoteAction                          func(uintptr, uintptr)
	_TraderSpiSetOnRspBatchOrderAction                     func(uintptr, uintptr)
	_TraderSpiSetOnRspOptionSelfCloseInsert                func(uintptr, uintptr)
	_TraderSpiSetOnRspOptionSelfCloseAction                func(uintptr, uintptr)
	_TraderSpiSetOnRspCombActionInsert                     func(uintptr, uintptr)
	_TraderSpiSetOnRspQryOrder                             func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTrade                             func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorPosition                  func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTradingAccount                    func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestor                          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTradingCode                       func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInstrumentMarginRate              func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInstrumentCommissionRate          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryExchange                          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryProduct                           func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInstrument                        func(uintptr, uintptr)
	_TraderSpiSetOnRspQryDepthMarketData                   func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTraderOffer                       func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySettlementInfo                    func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTransferBank                      func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorPositionDetail            func(uintptr, uintptr)
	_TraderSpiSetOnRspQryNotice                            func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySettlementInfoConfirm             func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorPositionCombineDetail     func(uintptr, uintptr)
	_TraderSpiSetOnRspQryCFMMCTradingAccountKey            func(uintptr, uintptr)
	_TraderSpiSetOnRspQryEWarrantOffset                    func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorProductGroupMargin        func(uintptr, uintptr)
	_TraderSpiSetOnRspQryExchangeMarginRate                func(uintptr, uintptr)
	_TraderSpiSetOnRspQryExchangeMarginRateAdjust          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryExchangeRate                      func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySecAgentACIDMap                   func(uintptr, uintptr)
	_TraderSpiSetOnRspQryProductExchRate                   func(uintptr, uintptr)
	_TraderSpiSetOnRspQryProductGroup                      func(uintptr, uintptr)
	_TraderSpiSetOnRspQryMMInstrumentCommissionRate        func(uintptr, uintptr)
	_TraderSpiSetOnRspQryMMOptionInstrCommRate             func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInstrumentOrderCommRate           func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySecAgentTradingAccount            func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySecAgentCheckMode                 func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySecAgentTradeInfo                 func(uintptr, uintptr)
	_TraderSpiSetOnRspQryOptionInstrTradeCost              func(uintptr, uintptr)
	_TraderSpiSetOnRspQryOptionInstrCommRate               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryExecOrder                         func(uintptr, uintptr)
	_TraderSpiSetOnRspQryForQuote                          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryQuote                             func(uintptr, uintptr)
	_TraderSpiSetOnRspQryOptionSelfClose                   func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestUnit                        func(uintptr, uintptr)
	_TraderSpiSetOnRspQryCombInstrumentGuard               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryCombAction                        func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTransferSerial                    func(uintptr, uintptr)
	_TraderSpiSetOnRspQryAccountregister                   func(uintptr, uintptr)
	_TraderSpiSetOnRspError                                func(uintptr, uintptr)
	_TraderSpiSetOnRtnOrder                                func(uintptr, uintptr)
	_TraderSpiSetOnRtnTrade                                func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnOrderInsert                       func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnOrderAction                       func(uintptr, uintptr)
	_TraderSpiSetOnRtnInstrumentStatus                     func(uintptr, uintptr)
	_TraderSpiSetOnRtnBulletin                             func(uintptr, uintptr)
	_TraderSpiSetOnRtnTradingNotice                        func(uintptr, uintptr)
	_TraderSpiSetOnRtnErrorConditionalOrder                func(uintptr, uintptr)
	_TraderSpiSetOnRtnExecOrder                            func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnExecOrderInsert                   func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnExecOrderAction                   func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnForQuoteInsert                    func(uintptr, uintptr)
	_TraderSpiSetOnRtnQuote                                func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnQuoteInsert                       func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnQuoteAction                       func(uintptr, uintptr)
	_TraderSpiSetOnRtnForQuoteRsp                          func(uintptr, uintptr)
	_TraderSpiSetOnRtnCFMMCTradingAccountToken             func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnBatchOrderAction                  func(uintptr, uintptr)
	_TraderSpiSetOnRtnOptionSelfClose                      func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnOptionSelfCloseInsert             func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnOptionSelfCloseAction             func(uintptr, uintptr)
	_TraderSpiSetOnRtnCombAction                           func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnCombActionInsert                  func(uintptr, uintptr)
	_TraderSpiSetOnRspQryContractBank                      func(uintptr, uintptr)
	_TraderSpiSetOnRspQryParkedOrder                       func(uintptr, uintptr)
	_TraderSpiSetOnRspQryParkedOrderAction                 func(uintptr, uintptr)
	_TraderSpiSetOnRspQryTradingNotice                     func(uintptr, uintptr)
	_TraderSpiSetOnRspQryBrokerTradingParams               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryBrokerTradingAlgos                func(uintptr, uintptr)
	_TraderSpiSetOnRspQueryCFMMCTradingAccountToken        func(uintptr, uintptr)
	_TraderSpiSetOnRtnFromBankToFutureByBank               func(uintptr, uintptr)
	_TraderSpiSetOnRtnFromFutureToBankByBank               func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromBankToFutureByBank         func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromFutureToBankByBank         func(uintptr, uintptr)
	_TraderSpiSetOnRtnFromBankToFutureByFuture             func(uintptr, uintptr)
	_TraderSpiSetOnRtnFromFutureToBankByFuture             func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual func(uintptr, uintptr)
	_TraderSpiSetOnRtnQueryBankBalanceByFuture             func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnBankToFutureByFuture              func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnFutureToBankByFuture              func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual  func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual  func(uintptr, uintptr)
	_TraderSpiSetOnErrRtnQueryBankBalanceByFuture          func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromBankToFutureByFuture       func(uintptr, uintptr)
	_TraderSpiSetOnRtnRepealFromFutureToBankByFuture       func(uintptr, uintptr)
	_TraderSpiSetOnRspFromBankToFutureByFuture             func(uintptr, uintptr)
	_TraderSpiSetOnRspFromFutureToBankByFuture             func(uintptr, uintptr)
	_TraderSpiSetOnRspQueryBankAccountMoneyByFuture        func(uintptr, uintptr)
	_TraderSpiSetOnRtnOpenAccountByBank                    func(uintptr, uintptr)
	_TraderSpiSetOnRtnCancelAccountByBank                  func(uintptr, uintptr)
	_TraderSpiSetOnRtnChangeAccountByBank                  func(uintptr, uintptr)
	_TraderSpiSetOnRspQryClassifiedInstrument              func(uintptr, uintptr)
	_TraderSpiSetOnRspQryCombPromotionParam                func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRiskSettleInvstPosition           func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRiskSettleProductStatus           func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMFutureParameter               func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMOptionParameter               func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMIntraParameter                func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMInterParameter                func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMPortfDefinition               func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMInvestorPortfDef              func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorPortfMarginRatio          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorProdSPBMDetail            func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorCommoditySPMMMargin       func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin  func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPMMInstParam                     func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPMMProductParam                  func(uintptr, uintptr)
	_TraderSpiSetOnRspQrySPBMAddOnInterParameter           func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSCombProductInfo              func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSInstrParameter               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSIntraParameter               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSInterParameter               func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSShortOptAdjustParam          func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRCAMSInvestorCombPosition         func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorProdRCAMSMargin           func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRULEInstrParameter                func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRULEIntraParameter                func(uintptr, uintptr)
	_TraderSpiSetOnRspQryRULEInterParameter                func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorProdRULEMargin            func(uintptr, uintptr)
	_TraderSpiSetOnRspQryInvestorPortfSetting              func(uintptr, uintptr)
	_TraderReqUserLoginWithSystemInfo                      func(uintptr, *CThostFtdcReqUserLoginField, int32, int32, *byte) int32
	_DCGetSystemInfo                                       func(*byte, *int32) int32
	_DCGetSystemInfoUnAesEncode                            func(*byte, *int32) int32
	_DCGetDataCollectApiVersion                            func() *byte
)

// initTraderApi 初始化交易 API 函数
func initTraderApi(lib uintptr) {
	traderOnce.Do(func() {
		purego.RegisterLibFunc(&_TraderCreateFtdcTraderApi, lib, "TraderCreateFtdcTraderApi")
		purego.RegisterLibFunc(&_TraderGetApiVersion, lib, "TraderGetApiVersion")
		purego.RegisterLibFunc(&_TraderRelease, lib, "TraderRelease")
		purego.RegisterLibFunc(&_TraderInit, lib, "TraderInit")
		purego.RegisterLibFunc(&_TraderJoin, lib, "TraderJoin")
		purego.RegisterLibFunc(&_TraderGetTradingDay, lib, "TraderGetTradingDay")
		purego.RegisterLibFunc(&_TraderGetFrontInfo, lib, "TraderGetFrontInfo")
		purego.RegisterLibFunc(&_TraderRegisterFront, lib, "TraderRegisterFront")
		purego.RegisterLibFunc(&_TraderRegisterNameServer, lib, "TraderRegisterNameServer")
		purego.RegisterLibFunc(&_TraderRegisterFensUserInfo, lib, "TraderRegisterFensUserInfo")
		purego.RegisterLibFunc(&_TraderSubscribePrivateTopic, lib, "TraderSubscribePrivateTopic")
		purego.RegisterLibFunc(&_TraderSubscribePublicTopic, lib, "TraderSubscribePublicTopic")
		purego.RegisterLibFunc(&_TraderReqAuthenticate, lib, "TraderReqAuthenticate")
		purego.RegisterLibFunc(&_TraderRegisterUserSystemInfo, lib, "TraderRegisterUserSystemInfo")
		purego.RegisterLibFunc(&_TraderSubmitUserSystemInfo, lib, "TraderSubmitUserSystemInfo")
		purego.RegisterLibFunc(&_TraderReqUserLogin, lib, "TraderReqUserLogin")
		purego.RegisterLibFunc(&_TraderReqUserLogout, lib, "TraderReqUserLogout")
		purego.RegisterLibFunc(&_TraderReqUserPasswordUpdate, lib, "TraderReqUserPasswordUpdate")
		purego.RegisterLibFunc(&_TraderReqTradingAccountPasswordUpdate, lib, "TraderReqTradingAccountPasswordUpdate")
		purego.RegisterLibFunc(&_TraderReqUserAuthMethod, lib, "TraderReqUserAuthMethod")
		purego.RegisterLibFunc(&_TraderReqGenUserCaptcha, lib, "TraderReqGenUserCaptcha")
		purego.RegisterLibFunc(&_TraderReqGenUserText, lib, "TraderReqGenUserText")
		purego.RegisterLibFunc(&_TraderReqUserLoginWithCaptcha, lib, "TraderReqUserLoginWithCaptcha")
		purego.RegisterLibFunc(&_TraderReqUserLoginWithText, lib, "TraderReqUserLoginWithText")
		purego.RegisterLibFunc(&_TraderReqUserLoginWithOTP, lib, "TraderReqUserLoginWithOTP")
		purego.RegisterLibFunc(&_TraderReqOrderInsert, lib, "TraderReqOrderInsert")
		purego.RegisterLibFunc(&_TraderReqParkedOrderInsert, lib, "TraderReqParkedOrderInsert")
		purego.RegisterLibFunc(&_TraderReqParkedOrderAction, lib, "TraderReqParkedOrderAction")
		purego.RegisterLibFunc(&_TraderReqOrderAction, lib, "TraderReqOrderAction")
		purego.RegisterLibFunc(&_TraderReqQryMaxOrderVolume, lib, "TraderReqQryMaxOrderVolume")
		purego.RegisterLibFunc(&_TraderReqSettlementInfoConfirm, lib, "TraderReqSettlementInfoConfirm")
		purego.RegisterLibFunc(&_TraderReqRemoveParkedOrder, lib, "TraderReqRemoveParkedOrder")
		purego.RegisterLibFunc(&_TraderReqRemoveParkedOrderAction, lib, "TraderReqRemoveParkedOrderAction")
		purego.RegisterLibFunc(&_TraderReqExecOrderInsert, lib, "TraderReqExecOrderInsert")
		purego.RegisterLibFunc(&_TraderReqExecOrderAction, lib, "TraderReqExecOrderAction")
		purego.RegisterLibFunc(&_TraderReqForQuoteInsert, lib, "TraderReqForQuoteInsert")
		purego.RegisterLibFunc(&_TraderReqQuoteInsert, lib, "TraderReqQuoteInsert")
		purego.RegisterLibFunc(&_TraderReqQuoteAction, lib, "TraderReqQuoteAction")
		purego.RegisterLibFunc(&_TraderReqBatchOrderAction, lib, "TraderReqBatchOrderAction")
		purego.RegisterLibFunc(&_TraderReqOptionSelfCloseInsert, lib, "TraderReqOptionSelfCloseInsert")
		purego.RegisterLibFunc(&_TraderReqOptionSelfCloseAction, lib, "TraderReqOptionSelfCloseAction")
		purego.RegisterLibFunc(&_TraderReqCombActionInsert, lib, "TraderReqCombActionInsert")
		purego.RegisterLibFunc(&_TraderReqQryOrder, lib, "TraderReqQryOrder")
		purego.RegisterLibFunc(&_TraderReqQryTrade, lib, "TraderReqQryTrade")
		purego.RegisterLibFunc(&_TraderReqQryInvestorPosition, lib, "TraderReqQryInvestorPosition")
		purego.RegisterLibFunc(&_TraderReqQryTradingAccount, lib, "TraderReqQryTradingAccount")
		purego.RegisterLibFunc(&_TraderReqQryInvestor, lib, "TraderReqQryInvestor")
		purego.RegisterLibFunc(&_TraderReqQryTradingCode, lib, "TraderReqQryTradingCode")
		purego.RegisterLibFunc(&_TraderReqQryInstrumentMarginRate, lib, "TraderReqQryInstrumentMarginRate")
		purego.RegisterLibFunc(&_TraderReqQryInstrumentCommissionRate, lib, "TraderReqQryInstrumentCommissionRate")
		purego.RegisterLibFunc(&_TraderReqQryExchange, lib, "TraderReqQryExchange")
		purego.RegisterLibFunc(&_TraderReqQryProduct, lib, "TraderReqQryProduct")
		purego.RegisterLibFunc(&_TraderReqQryInstrument, lib, "TraderReqQryInstrument")
		purego.RegisterLibFunc(&_TraderReqQryDepthMarketData, lib, "TraderReqQryDepthMarketData")
		purego.RegisterLibFunc(&_TraderReqQryTraderOffer, lib, "TraderReqQryTraderOffer")
		purego.RegisterLibFunc(&_TraderReqQrySettlementInfo, lib, "TraderReqQrySettlementInfo")
		purego.RegisterLibFunc(&_TraderReqQryTransferBank, lib, "TraderReqQryTransferBank")
		purego.RegisterLibFunc(&_TraderReqQryInvestorPositionDetail, lib, "TraderReqQryInvestorPositionDetail")
		purego.RegisterLibFunc(&_TraderReqQryNotice, lib, "TraderReqQryNotice")
		purego.RegisterLibFunc(&_TraderReqQrySettlementInfoConfirm, lib, "TraderReqQrySettlementInfoConfirm")
		purego.RegisterLibFunc(&_TraderReqQryInvestorPositionCombineDetail, lib, "TraderReqQryInvestorPositionCombineDetail")
		purego.RegisterLibFunc(&_TraderReqQryCFMMCTradingAccountKey, lib, "TraderReqQryCFMMCTradingAccountKey")
		purego.RegisterLibFunc(&_TraderReqQryEWarrantOffset, lib, "TraderReqQryEWarrantOffset")
		purego.RegisterLibFunc(&_TraderReqQryInvestorProductGroupMargin, lib, "TraderReqQryInvestorProductGroupMargin")
		purego.RegisterLibFunc(&_TraderReqQryExchangeMarginRate, lib, "TraderReqQryExchangeMarginRate")
		purego.RegisterLibFunc(&_TraderReqQryExchangeMarginRateAdjust, lib, "TraderReqQryExchangeMarginRateAdjust")
		purego.RegisterLibFunc(&_TraderReqQryExchangeRate, lib, "TraderReqQryExchangeRate")
		purego.RegisterLibFunc(&_TraderReqQrySecAgentACIDMap, lib, "TraderReqQrySecAgentACIDMap")
		purego.RegisterLibFunc(&_TraderReqQryProductExchRate, lib, "TraderReqQryProductExchRate")
		purego.RegisterLibFunc(&_TraderReqQryProductGroup, lib, "TraderReqQryProductGroup")
		purego.RegisterLibFunc(&_TraderReqQryMMInstrumentCommissionRate, lib, "TraderReqQryMMInstrumentCommissionRate")
		purego.RegisterLibFunc(&_TraderReqQryMMOptionInstrCommRate, lib, "TraderReqQryMMOptionInstrCommRate")
		purego.RegisterLibFunc(&_TraderReqQryInstrumentOrderCommRate, lib, "TraderReqQryInstrumentOrderCommRate")
		purego.RegisterLibFunc(&_TraderReqQrySecAgentTradingAccount, lib, "TraderReqQrySecAgentTradingAccount")
		purego.RegisterLibFunc(&_TraderReqQrySecAgentCheckMode, lib, "TraderReqQrySecAgentCheckMode")
		purego.RegisterLibFunc(&_TraderReqQrySecAgentTradeInfo, lib, "TraderReqQrySecAgentTradeInfo")
		purego.RegisterLibFunc(&_TraderReqQryOptionInstrTradeCost, lib, "TraderReqQryOptionInstrTradeCost")
		purego.RegisterLibFunc(&_TraderReqQryOptionInstrCommRate, lib, "TraderReqQryOptionInstrCommRate")
		purego.RegisterLibFunc(&_TraderReqQryExecOrder, lib, "TraderReqQryExecOrder")
		purego.RegisterLibFunc(&_TraderReqQryForQuote, lib, "TraderReqQryForQuote")
		purego.RegisterLibFunc(&_TraderReqQryQuote, lib, "TraderReqQryQuote")
		purego.RegisterLibFunc(&_TraderReqQryOptionSelfClose, lib, "TraderReqQryOptionSelfClose")
		purego.RegisterLibFunc(&_TraderReqQryInvestUnit, lib, "TraderReqQryInvestUnit")
		purego.RegisterLibFunc(&_TraderReqQryCombInstrumentGuard, lib, "TraderReqQryCombInstrumentGuard")
		purego.RegisterLibFunc(&_TraderReqQryCombAction, lib, "TraderReqQryCombAction")
		purego.RegisterLibFunc(&_TraderReqQryTransferSerial, lib, "TraderReqQryTransferSerial")
		purego.RegisterLibFunc(&_TraderReqQryAccountregister, lib, "TraderReqQryAccountregister")
		purego.RegisterLibFunc(&_TraderReqQryContractBank, lib, "TraderReqQryContractBank")
		purego.RegisterLibFunc(&_TraderReqQryParkedOrder, lib, "TraderReqQryParkedOrder")
		purego.RegisterLibFunc(&_TraderReqQryParkedOrderAction, lib, "TraderReqQryParkedOrderAction")
		purego.RegisterLibFunc(&_TraderReqQryTradingNotice, lib, "TraderReqQryTradingNotice")
		purego.RegisterLibFunc(&_TraderReqQryBrokerTradingParams, lib, "TraderReqQryBrokerTradingParams")
		purego.RegisterLibFunc(&_TraderReqQryBrokerTradingAlgos, lib, "TraderReqQryBrokerTradingAlgos")
		purego.RegisterLibFunc(&_TraderReqQueryCFMMCTradingAccountToken, lib, "TraderReqQueryCFMMCTradingAccountToken")
		purego.RegisterLibFunc(&_TraderReqFromBankToFutureByFuture, lib, "TraderReqFromBankToFutureByFuture")
		purego.RegisterLibFunc(&_TraderReqFromFutureToBankByFuture, lib, "TraderReqFromFutureToBankByFuture")
		purego.RegisterLibFunc(&_TraderReqQueryBankAccountMoneyByFuture, lib, "TraderReqQueryBankAccountMoneyByFuture")
		purego.RegisterLibFunc(&_TraderReqQryClassifiedInstrument, lib, "TraderReqQryClassifiedInstrument")
		purego.RegisterLibFunc(&_TraderReqQryCombPromotionParam, lib, "TraderReqQryCombPromotionParam")
		purego.RegisterLibFunc(&_TraderReqQryRiskSettleInvstPosition, lib, "TraderReqQryRiskSettleInvstPosition")
		purego.RegisterLibFunc(&_TraderReqQryRiskSettleProductStatus, lib, "TraderReqQryRiskSettleProductStatus")
		purego.RegisterLibFunc(&_TraderReqQrySPBMFutureParameter, lib, "TraderReqQrySPBMFutureParameter")
		purego.RegisterLibFunc(&_TraderReqQrySPBMOptionParameter, lib, "TraderReqQrySPBMOptionParameter")
		purego.RegisterLibFunc(&_TraderReqQrySPBMIntraParameter, lib, "TraderReqQrySPBMIntraParameter")
		purego.RegisterLibFunc(&_TraderReqQrySPBMInterParameter, lib, "TraderReqQrySPBMInterParameter")
		purego.RegisterLibFunc(&_TraderReqQrySPBMPortfDefinition, lib, "TraderReqQrySPBMPortfDefinition")
		purego.RegisterLibFunc(&_TraderReqQrySPBMInvestorPortfDef, lib, "TraderReqQrySPBMInvestorPortfDef")
		purego.RegisterLibFunc(&_TraderReqQryInvestorPortfMarginRatio, lib, "TraderReqQryInvestorPortfMarginRatio")
		purego.RegisterLibFunc(&_TraderReqQryInvestorProdSPBMDetail, lib, "TraderReqQryInvestorProdSPBMDetail")
		purego.RegisterLibFunc(&_TraderReqQryInvestorCommoditySPMMMargin, lib, "TraderReqQryInvestorCommoditySPMMMargin")
		purego.RegisterLibFunc(&_TraderReqQryInvestorCommodityGroupSPMMMargin, lib, "TraderReqQryInvestorCommodityGroupSPMMMargin")
		purego.RegisterLibFunc(&_TraderReqQrySPMMInstParam, lib, "TraderReqQrySPMMInstParam")
		purego.RegisterLibFunc(&_TraderReqQrySPMMProductParam, lib, "TraderReqQrySPMMProductParam")
		purego.RegisterLibFunc(&_TraderReqQrySPBMAddOnInterParameter, lib, "TraderReqQrySPBMAddOnInterParameter")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSCombProductInfo, lib, "TraderReqQryRCAMSCombProductInfo")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSInstrParameter, lib, "TraderReqQryRCAMSInstrParameter")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSIntraParameter, lib, "TraderReqQryRCAMSIntraParameter")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSInterParameter, lib, "TraderReqQryRCAMSInterParameter")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSShortOptAdjustParam, lib, "TraderReqQryRCAMSShortOptAdjustParam")
		purego.RegisterLibFunc(&_TraderReqQryRCAMSInvestorCombPosition, lib, "TraderReqQryRCAMSInvestorCombPosition")
		purego.RegisterLibFunc(&_TraderReqQryInvestorProdRCAMSMargin, lib, "TraderReqQryInvestorProdRCAMSMargin")
		purego.RegisterLibFunc(&_TraderReqQryRULEInstrParameter, lib, "TraderReqQryRULEInstrParameter")
		purego.RegisterLibFunc(&_TraderReqQryRULEIntraParameter, lib, "TraderReqQryRULEIntraParameter")
		purego.RegisterLibFunc(&_TraderReqQryRULEInterParameter, lib, "TraderReqQryRULEInterParameter")
		purego.RegisterLibFunc(&_TraderReqQryInvestorProdRULEMargin, lib, "TraderReqQryInvestorProdRULEMargin")
		purego.RegisterLibFunc(&_TraderReqQryInvestorPortfSetting, lib, "TraderReqQryInvestorPortfSetting")
		purego.RegisterLibFunc(&_TraderSpiCreate, lib, "TraderSpiCreate")
		purego.RegisterLibFunc(&_TraderSpiDestroy, lib, "TraderSpiDestroy")
		purego.RegisterLibFunc(&_TraderRegisterSpi, lib, "TraderRegisterSpi")
		purego.RegisterLibFunc(&_TraderSpiSetCallbacks, lib, "TraderSpiSetCallbacks")
		purego.RegisterLibFunc(&_TraderSpiSetOnFrontConnected, lib, "TraderSpiSetOnFrontConnected")
		purego.RegisterLibFunc(&_TraderSpiSetOnFrontDisconnected, lib, "TraderSpiSetOnFrontDisconnected")
		purego.RegisterLibFunc(&_TraderSpiSetOnHeartBeatWarning, lib, "TraderSpiSetOnHeartBeatWarning")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspAuthenticate, lib, "TraderSpiSetOnRspAuthenticate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspUserLogin, lib, "TraderSpiSetOnRspUserLogin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspUserLogout, lib, "TraderSpiSetOnRspUserLogout")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspUserPasswordUpdate, lib, "TraderSpiSetOnRspUserPasswordUpdate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspTradingAccountPasswordUpdate, lib, "TraderSpiSetOnRspTradingAccountPasswordUpdate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspUserAuthMethod, lib, "TraderSpiSetOnRspUserAuthMethod")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspGenUserCaptcha, lib, "TraderSpiSetOnRspGenUserCaptcha")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspGenUserText, lib, "TraderSpiSetOnRspGenUserText")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspOrderInsert, lib, "TraderSpiSetOnRspOrderInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspParkedOrderInsert, lib, "TraderSpiSetOnRspParkedOrderInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspParkedOrderAction, lib, "TraderSpiSetOnRspParkedOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspOrderAction, lib, "TraderSpiSetOnRspOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryMaxOrderVolume, lib, "TraderSpiSetOnRspQryMaxOrderVolume")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspSettlementInfoConfirm, lib, "TraderSpiSetOnRspSettlementInfoConfirm")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspRemoveParkedOrder, lib, "TraderSpiSetOnRspRemoveParkedOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspRemoveParkedOrderAction, lib, "TraderSpiSetOnRspRemoveParkedOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspExecOrderInsert, lib, "TraderSpiSetOnRspExecOrderInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspExecOrderAction, lib, "TraderSpiSetOnRspExecOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspForQuoteInsert, lib, "TraderSpiSetOnRspForQuoteInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQuoteInsert, lib, "TraderSpiSetOnRspQuoteInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQuoteAction, lib, "TraderSpiSetOnRspQuoteAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspBatchOrderAction, lib, "TraderSpiSetOnRspBatchOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspOptionSelfCloseInsert, lib, "TraderSpiSetOnRspOptionSelfCloseInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspOptionSelfCloseAction, lib, "TraderSpiSetOnRspOptionSelfCloseAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspCombActionInsert, lib, "TraderSpiSetOnRspCombActionInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryOrder, lib, "TraderSpiSetOnRspQryOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTrade, lib, "TraderSpiSetOnRspQryTrade")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorPosition, lib, "TraderSpiSetOnRspQryInvestorPosition")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTradingAccount, lib, "TraderSpiSetOnRspQryTradingAccount")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestor, lib, "TraderSpiSetOnRspQryInvestor")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTradingCode, lib, "TraderSpiSetOnRspQryTradingCode")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInstrumentMarginRate, lib, "TraderSpiSetOnRspQryInstrumentMarginRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInstrumentCommissionRate, lib, "TraderSpiSetOnRspQryInstrumentCommissionRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryExchange, lib, "TraderSpiSetOnRspQryExchange")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryProduct, lib, "TraderSpiSetOnRspQryProduct")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInstrument, lib, "TraderSpiSetOnRspQryInstrument")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryDepthMarketData, lib, "TraderSpiSetOnRspQryDepthMarketData")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTraderOffer, lib, "TraderSpiSetOnRspQryTraderOffer")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySettlementInfo, lib, "TraderSpiSetOnRspQrySettlementInfo")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTransferBank, lib, "TraderSpiSetOnRspQryTransferBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorPositionDetail, lib, "TraderSpiSetOnRspQryInvestorPositionDetail")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryNotice, lib, "TraderSpiSetOnRspQryNotice")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySettlementInfoConfirm, lib, "TraderSpiSetOnRspQrySettlementInfoConfirm")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorPositionCombineDetail, lib, "TraderSpiSetOnRspQryInvestorPositionCombineDetail")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryCFMMCTradingAccountKey, lib, "TraderSpiSetOnRspQryCFMMCTradingAccountKey")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryEWarrantOffset, lib, "TraderSpiSetOnRspQryEWarrantOffset")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorProductGroupMargin, lib, "TraderSpiSetOnRspQryInvestorProductGroupMargin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryExchangeMarginRate, lib, "TraderSpiSetOnRspQryExchangeMarginRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryExchangeMarginRateAdjust, lib, "TraderSpiSetOnRspQryExchangeMarginRateAdjust")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryExchangeRate, lib, "TraderSpiSetOnRspQryExchangeRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySecAgentACIDMap, lib, "TraderSpiSetOnRspQrySecAgentACIDMap")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryProductExchRate, lib, "TraderSpiSetOnRspQryProductExchRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryProductGroup, lib, "TraderSpiSetOnRspQryProductGroup")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryMMInstrumentCommissionRate, lib, "TraderSpiSetOnRspQryMMInstrumentCommissionRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryMMOptionInstrCommRate, lib, "TraderSpiSetOnRspQryMMOptionInstrCommRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInstrumentOrderCommRate, lib, "TraderSpiSetOnRspQryInstrumentOrderCommRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySecAgentTradingAccount, lib, "TraderSpiSetOnRspQrySecAgentTradingAccount")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySecAgentCheckMode, lib, "TraderSpiSetOnRspQrySecAgentCheckMode")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySecAgentTradeInfo, lib, "TraderSpiSetOnRspQrySecAgentTradeInfo")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryOptionInstrTradeCost, lib, "TraderSpiSetOnRspQryOptionInstrTradeCost")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryOptionInstrCommRate, lib, "TraderSpiSetOnRspQryOptionInstrCommRate")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryExecOrder, lib, "TraderSpiSetOnRspQryExecOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryForQuote, lib, "TraderSpiSetOnRspQryForQuote")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryQuote, lib, "TraderSpiSetOnRspQryQuote")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryOptionSelfClose, lib, "TraderSpiSetOnRspQryOptionSelfClose")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestUnit, lib, "TraderSpiSetOnRspQryInvestUnit")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryCombInstrumentGuard, lib, "TraderSpiSetOnRspQryCombInstrumentGuard")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryCombAction, lib, "TraderSpiSetOnRspQryCombAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTransferSerial, lib, "TraderSpiSetOnRspQryTransferSerial")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryAccountregister, lib, "TraderSpiSetOnRspQryAccountregister")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspError, lib, "TraderSpiSetOnRspError")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnOrder, lib, "TraderSpiSetOnRtnOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnTrade, lib, "TraderSpiSetOnRtnTrade")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnOrderInsert, lib, "TraderSpiSetOnErrRtnOrderInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnOrderAction, lib, "TraderSpiSetOnErrRtnOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnInstrumentStatus, lib, "TraderSpiSetOnRtnInstrumentStatus")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnBulletin, lib, "TraderSpiSetOnRtnBulletin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnTradingNotice, lib, "TraderSpiSetOnRtnTradingNotice")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnErrorConditionalOrder, lib, "TraderSpiSetOnRtnErrorConditionalOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnExecOrder, lib, "TraderSpiSetOnRtnExecOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnExecOrderInsert, lib, "TraderSpiSetOnErrRtnExecOrderInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnExecOrderAction, lib, "TraderSpiSetOnErrRtnExecOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnForQuoteInsert, lib, "TraderSpiSetOnErrRtnForQuoteInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnQuote, lib, "TraderSpiSetOnRtnQuote")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnQuoteInsert, lib, "TraderSpiSetOnErrRtnQuoteInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnQuoteAction, lib, "TraderSpiSetOnErrRtnQuoteAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnForQuoteRsp, lib, "TraderSpiSetOnRtnForQuoteRsp")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnCFMMCTradingAccountToken, lib, "TraderSpiSetOnRtnCFMMCTradingAccountToken")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnBatchOrderAction, lib, "TraderSpiSetOnErrRtnBatchOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnOptionSelfClose, lib, "TraderSpiSetOnRtnOptionSelfClose")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnOptionSelfCloseInsert, lib, "TraderSpiSetOnErrRtnOptionSelfCloseInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnOptionSelfCloseAction, lib, "TraderSpiSetOnErrRtnOptionSelfCloseAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnCombAction, lib, "TraderSpiSetOnRtnCombAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnCombActionInsert, lib, "TraderSpiSetOnErrRtnCombActionInsert")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryContractBank, lib, "TraderSpiSetOnRspQryContractBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryParkedOrder, lib, "TraderSpiSetOnRspQryParkedOrder")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryParkedOrderAction, lib, "TraderSpiSetOnRspQryParkedOrderAction")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryTradingNotice, lib, "TraderSpiSetOnRspQryTradingNotice")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryBrokerTradingParams, lib, "TraderSpiSetOnRspQryBrokerTradingParams")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryBrokerTradingAlgos, lib, "TraderSpiSetOnRspQryBrokerTradingAlgos")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQueryCFMMCTradingAccountToken, lib, "TraderSpiSetOnRspQueryCFMMCTradingAccountToken")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnFromBankToFutureByBank, lib, "TraderSpiSetOnRtnFromBankToFutureByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnFromFutureToBankByBank, lib, "TraderSpiSetOnRtnFromFutureToBankByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromBankToFutureByBank, lib, "TraderSpiSetOnRtnRepealFromBankToFutureByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromFutureToBankByBank, lib, "TraderSpiSetOnRtnRepealFromFutureToBankByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnFromBankToFutureByFuture, lib, "TraderSpiSetOnRtnFromBankToFutureByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnFromFutureToBankByFuture, lib, "TraderSpiSetOnRtnFromFutureToBankByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual, lib, "TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual, lib, "TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnQueryBankBalanceByFuture, lib, "TraderSpiSetOnRtnQueryBankBalanceByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnBankToFutureByFuture, lib, "TraderSpiSetOnErrRtnBankToFutureByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnFutureToBankByFuture, lib, "TraderSpiSetOnErrRtnFutureToBankByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual, lib, "TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual, lib, "TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual")
		purego.RegisterLibFunc(&_TraderSpiSetOnErrRtnQueryBankBalanceByFuture, lib, "TraderSpiSetOnErrRtnQueryBankBalanceByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromBankToFutureByFuture, lib, "TraderSpiSetOnRtnRepealFromBankToFutureByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnRepealFromFutureToBankByFuture, lib, "TraderSpiSetOnRtnRepealFromFutureToBankByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspFromBankToFutureByFuture, lib, "TraderSpiSetOnRspFromBankToFutureByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspFromFutureToBankByFuture, lib, "TraderSpiSetOnRspFromFutureToBankByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQueryBankAccountMoneyByFuture, lib, "TraderSpiSetOnRspQueryBankAccountMoneyByFuture")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnOpenAccountByBank, lib, "TraderSpiSetOnRtnOpenAccountByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnCancelAccountByBank, lib, "TraderSpiSetOnRtnCancelAccountByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRtnChangeAccountByBank, lib, "TraderSpiSetOnRtnChangeAccountByBank")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryClassifiedInstrument, lib, "TraderSpiSetOnRspQryClassifiedInstrument")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryCombPromotionParam, lib, "TraderSpiSetOnRspQryCombPromotionParam")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRiskSettleInvstPosition, lib, "TraderSpiSetOnRspQryRiskSettleInvstPosition")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRiskSettleProductStatus, lib, "TraderSpiSetOnRspQryRiskSettleProductStatus")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMFutureParameter, lib, "TraderSpiSetOnRspQrySPBMFutureParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMOptionParameter, lib, "TraderSpiSetOnRspQrySPBMOptionParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMIntraParameter, lib, "TraderSpiSetOnRspQrySPBMIntraParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMInterParameter, lib, "TraderSpiSetOnRspQrySPBMInterParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMPortfDefinition, lib, "TraderSpiSetOnRspQrySPBMPortfDefinition")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMInvestorPortfDef, lib, "TraderSpiSetOnRspQrySPBMInvestorPortfDef")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorPortfMarginRatio, lib, "TraderSpiSetOnRspQryInvestorPortfMarginRatio")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorProdSPBMDetail, lib, "TraderSpiSetOnRspQryInvestorProdSPBMDetail")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorCommoditySPMMMargin, lib, "TraderSpiSetOnRspQryInvestorCommoditySPMMMargin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin, lib, "TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPMMInstParam, lib, "TraderSpiSetOnRspQrySPMMInstParam")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPMMProductParam, lib, "TraderSpiSetOnRspQrySPMMProductParam")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQrySPBMAddOnInterParameter, lib, "TraderSpiSetOnRspQrySPBMAddOnInterParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSCombProductInfo, lib, "TraderSpiSetOnRspQryRCAMSCombProductInfo")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSInstrParameter, lib, "TraderSpiSetOnRspQryRCAMSInstrParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSIntraParameter, lib, "TraderSpiSetOnRspQryRCAMSIntraParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSInterParameter, lib, "TraderSpiSetOnRspQryRCAMSInterParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSShortOptAdjustParam, lib, "TraderSpiSetOnRspQryRCAMSShortOptAdjustParam")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRCAMSInvestorCombPosition, lib, "TraderSpiSetOnRspQryRCAMSInvestorCombPosition")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorProdRCAMSMargin, lib, "TraderSpiSetOnRspQryInvestorProdRCAMSMargin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRULEInstrParameter, lib, "TraderSpiSetOnRspQryRULEInstrParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRULEIntraParameter, lib, "TraderSpiSetOnRspQryRULEIntraParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryRULEInterParameter, lib, "TraderSpiSetOnRspQryRULEInterParameter")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorProdRULEMargin, lib, "TraderSpiSetOnRspQryInvestorProdRULEMargin")
		purego.RegisterLibFunc(&_TraderSpiSetOnRspQryInvestorPortfSetting, lib, "TraderSpiSetOnRspQryInvestorPortfSetting")
		purego.RegisterLibFunc(&_TraderReqUserLoginWithSystemInfo, lib, "TraderReqUserLoginWithSystemInfo")
		purego.RegisterLibFunc(&_DCGetSystemInfo, lib, "DCGetSystemInfo")
		purego.RegisterLibFunc(&_DCGetSystemInfoUnAesEncode, lib, "DCGetSystemInfoUnAesEncode")
		purego.RegisterLibFunc(&_DCGetDataCollectApiVersion, lib, "DCGetDataCollectApiVersion")
	})
}

// ========== 实例管理 ==========

var (
	traderInstances   = make(map[uintptr]*TraderApi)
	traderInstancesMu sync.RWMutex
	traderNextID      uintptr = 1
)

func registerTraderInstance(api *TraderApi) uintptr {
	traderInstancesMu.Lock()
	defer traderInstancesMu.Unlock()
	id := traderNextID
	traderNextID++
	traderInstances[id] = api
	return id
}

func getTraderInstance(userData uintptr) *TraderApi {
	traderInstancesMu.RLock()
	defer traderInstancesMu.RUnlock()
	return traderInstances[userData]
}

func unregisterTraderInstance(userData uintptr) {
	traderInstancesMu.Lock()
	defer traderInstancesMu.Unlock()
	delete(traderInstances, userData)
}

// ========== 构造函数 ==========

// NewTraderApi 创建交易 API 实例
// 首次调用时会自动加载 CTP 库（如果尚未加载）
func NewTraderApi(flowPath string) *TraderApi {
	// 自动加载库（如果尚未加载）
	if err := autoLoadLibrary(); err != nil {
		// 如果自动加载失败，返回 nil（或者可以 panic，取决于设计）
		// 这里返回 nil，让调用者检查
		return nil
	}

	api := &TraderApi{}
	api.userData = registerTraderInstance(api)

	// 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录
	absFlowPath := flowPath
	if !filepath.IsAbs(flowPath) {
		var err error
		absFlowPath, err = filepath.Abs(flowPath)
		if err != nil {
			absFlowPath = flowPath
		}
	}

	// 确保路径以路径分隔符结尾
	if len(absFlowPath) > 0 && absFlowPath[len(absFlowPath)-1] != filepath.Separator {
		absFlowPath += string(filepath.Separator)
	}

	// 确保目录存在（CTP API 需要这个目录来创建 flow 文件）
	if err := os.MkdirAll(absFlowPath, 0755); err != nil {
		// 如果创建目录失败，记录错误但继续（CTP API 可能会自己创建）
		fmt.Printf("警告: 无法创建 flow 目录 %s: %v\n", absFlowPath, err)
		// 这里不返回错误，让 CTP API 自己处理
	}

	// 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收
	api.flowPath = make([]byte, len(absFlowPath)+1)
	copy(api.flowPath, absFlowPath)
	api.flowPath[len(absFlowPath)] = 0
	pathPtr := &api.flowPath[0]

	api.handle = _TraderCreateFtdcTraderApi(pathPtr)

	runtime.SetFinalizer(api, (*TraderApi).Release)
	return api
}

// ========== API 方法 ==========

// GetApiVersion 获取API的版本信息
func (api *TraderApi) GetApiVersion() string {
	ptr := _TraderGetApiVersion()
	if ptr == nil {
		return ""
	}
	return GoString(ptr)
}

// Release 删除接口对象本身
func (api *TraderApi) Release() {
	_TraderRelease(api.handle)
	unregisterTraderInstance(api.userData)
}

// Init 初始化
func (api *TraderApi) Init() {
	_TraderInit(api.handle)
}

// Join 等待接口线程结束运行
func (api *TraderApi) Join() int32 {
	return _TraderJoin(api.handle)
}

// GetTradingDay 获取当前交易日
func (api *TraderApi) GetTradingDay() string {
	ptr := _TraderGetTradingDay(api.handle)
	if ptr == nil {
		return ""
	}
	return GoString(ptr)
}

// GetFrontInfo 获取已连接的前置的信息
func (api *TraderApi) GetFrontInfo(pFrontInfo *CThostFtdcFrontInfoField) {
	_TraderGetFrontInfo(api.handle, pFrontInfo)
}

// RegisterFront 注册前置机网络地址
func (api *TraderApi) RegisterFront(pszFrontAddress string) {
	_TraderRegisterFront(api.handle, CString(pszFrontAddress))
}

// RegisterNameServer 注册名字服务器网络地址
func (api *TraderApi) RegisterNameServer(pszNsAddress string) {
	_TraderRegisterNameServer(api.handle, CString(pszNsAddress))
}

// RegisterFensUserInfo 注册名字服务器用户信息
func (api *TraderApi) RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField) {
	_TraderRegisterFensUserInfo(api.handle, pFensUserInfo)
}

// SubscribePrivateTopic 订阅私有流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后私有流的内容
func (api *TraderApi) SubscribePrivateTopic(nResumeType THOST_TE_RESUME_TYPE) {
	_TraderSubscribePrivateTopic(api.handle, nResumeType)
}

// SubscribePublicTopic 订阅公共流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后公共流的内容 THOST_TERT_NONE:取消订阅公共流
func (api *TraderApi) SubscribePublicTopic(nResumeType THOST_TE_RESUME_TYPE) {
	_TraderSubscribePublicTopic(api.handle, nResumeType)
}

// ReqAuthenticate 客户端认证请求
func (api *TraderApi) ReqAuthenticate(pReqAuthenticateField *CThostFtdcReqAuthenticateField, nRequestID int32) int32 {
	return _TraderReqAuthenticate(api.handle, pReqAuthenticateField, nRequestID)
}

// RegisterUserSystemInfo 注册用户终端信息，用于中继服务器多连接模式 需要在终端认证成功后，用户登录前调用该接口
func (api *TraderApi) RegisterUserSystemInfo(pUserSystemInfo *CThostFtdcUserSystemInfoField) int32 {
	return _TraderRegisterUserSystemInfo(api.handle, pUserSystemInfo)
}

// SubmitUserSystemInfo 上报用户终端信息，用于中继服务器操作员登录模式 操作员登录后，可以多次调用该接口上报客户信息
func (api *TraderApi) SubmitUserSystemInfo(pUserSystemInfo *CThostFtdcUserSystemInfoField) int32 {
	return _TraderSubmitUserSystemInfo(api.handle, pUserSystemInfo)
}

// ReqUserLogin 用户登录请求
func (api *TraderApi) ReqUserLogin(pReqUserLoginField *CThostFtdcReqUserLoginField, nRequestID int32) int32 {
	return _TraderReqUserLogin(api.handle, pReqUserLoginField, nRequestID)
}

// ReqUserLogout 登出请求
func (api *TraderApi) ReqUserLogout(pUserLogout *CThostFtdcUserLogoutField, nRequestID int32) int32 {
	return _TraderReqUserLogout(api.handle, pUserLogout, nRequestID)
}

// ReqUserPasswordUpdate 用户口令更新请求
func (api *TraderApi) ReqUserPasswordUpdate(pUserPasswordUpdate *CThostFtdcUserPasswordUpdateField, nRequestID int32) int32 {
	return _TraderReqUserPasswordUpdate(api.handle, pUserPasswordUpdate, nRequestID)
}

// ReqTradingAccountPasswordUpdate 资金账户口令更新请求
func (api *TraderApi) ReqTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate *CThostFtdcTradingAccountPasswordUpdateField, nRequestID int32) int32 {
	return _TraderReqTradingAccountPasswordUpdate(api.handle, pTradingAccountPasswordUpdate, nRequestID)
}

// ReqUserAuthMethod 查询用户当前支持的认证模式
func (api *TraderApi) ReqUserAuthMethod(pReqUserAuthMethod *CThostFtdcReqUserAuthMethodField, nRequestID int32) int32 {
	return _TraderReqUserAuthMethod(api.handle, pReqUserAuthMethod, nRequestID)
}

// ReqGenUserCaptcha 用户发出获取图形验证码请求
func (api *TraderApi) ReqGenUserCaptcha(pReqGenUserCaptcha *CThostFtdcReqGenUserCaptchaField, nRequestID int32) int32 {
	return _TraderReqGenUserCaptcha(api.handle, pReqGenUserCaptcha, nRequestID)
}

// ReqGenUserText 用户发出获取短信验证码请求
func (api *TraderApi) ReqGenUserText(pReqGenUserText *CThostFtdcReqGenUserTextField, nRequestID int32) int32 {
	return _TraderReqGenUserText(api.handle, pReqGenUserText, nRequestID)
}

// ReqUserLoginWithCaptcha 用户发出带有图片验证码的登陆请求
func (api *TraderApi) ReqUserLoginWithCaptcha(pReqUserLoginWithCaptcha *CThostFtdcReqUserLoginWithCaptchaField, nRequestID int32) int32 {
	return _TraderReqUserLoginWithCaptcha(api.handle, pReqUserLoginWithCaptcha, nRequestID)
}

// ReqUserLoginWithText 用户发出带有短信验证码的登陆请求
func (api *TraderApi) ReqUserLoginWithText(pReqUserLoginWithText *CThostFtdcReqUserLoginWithTextField, nRequestID int32) int32 {
	return _TraderReqUserLoginWithText(api.handle, pReqUserLoginWithText, nRequestID)
}

// ReqUserLoginWithOTP 用户发出带有动态口令的登陆请求
func (api *TraderApi) ReqUserLoginWithOTP(pReqUserLoginWithOTP *CThostFtdcReqUserLoginWithOTPField, nRequestID int32) int32 {
	return _TraderReqUserLoginWithOTP(api.handle, pReqUserLoginWithOTP, nRequestID)
}

// ReqOrderInsert 报单录入请求
func (api *TraderApi) ReqOrderInsert(pInputOrder *CThostFtdcInputOrderField, nRequestID int32) int32 {
	return _TraderReqOrderInsert(api.handle, pInputOrder, nRequestID)
}

// ReqParkedOrderInsert 预埋单录入请求
func (api *TraderApi) ReqParkedOrderInsert(pParkedOrder *CThostFtdcParkedOrderField, nRequestID int32) int32 {
	return _TraderReqParkedOrderInsert(api.handle, pParkedOrder, nRequestID)
}

// ReqParkedOrderAction 预埋撤单录入请求
func (api *TraderApi) ReqParkedOrderAction(pParkedOrderAction *CThostFtdcParkedOrderActionField, nRequestID int32) int32 {
	return _TraderReqParkedOrderAction(api.handle, pParkedOrderAction, nRequestID)
}

// ReqOrderAction 报单操作请求
func (api *TraderApi) ReqOrderAction(pInputOrderAction *CThostFtdcInputOrderActionField, nRequestID int32) int32 {
	return _TraderReqOrderAction(api.handle, pInputOrderAction, nRequestID)
}

// ReqQryMaxOrderVolume 查询最大报单数量请求
func (api *TraderApi) ReqQryMaxOrderVolume(pQryMaxOrderVolume *CThostFtdcQryMaxOrderVolumeField, nRequestID int32) int32 {
	return _TraderReqQryMaxOrderVolume(api.handle, pQryMaxOrderVolume, nRequestID)
}

// ReqSettlementInfoConfirm 投资者结算结果确认
func (api *TraderApi) ReqSettlementInfoConfirm(pSettlementInfoConfirm *CThostFtdcSettlementInfoConfirmField, nRequestID int32) int32 {
	return _TraderReqSettlementInfoConfirm(api.handle, pSettlementInfoConfirm, nRequestID)
}

// ReqRemoveParkedOrder 请求删除预埋单
func (api *TraderApi) ReqRemoveParkedOrder(pRemoveParkedOrder *CThostFtdcRemoveParkedOrderField, nRequestID int32) int32 {
	return _TraderReqRemoveParkedOrder(api.handle, pRemoveParkedOrder, nRequestID)
}

// ReqRemoveParkedOrderAction 请求删除预埋撤单
func (api *TraderApi) ReqRemoveParkedOrderAction(pRemoveParkedOrderAction *CThostFtdcRemoveParkedOrderActionField, nRequestID int32) int32 {
	return _TraderReqRemoveParkedOrderAction(api.handle, pRemoveParkedOrderAction, nRequestID)
}

// ReqExecOrderInsert 执行宣告录入请求
func (api *TraderApi) ReqExecOrderInsert(pInputExecOrder *CThostFtdcInputExecOrderField, nRequestID int32) int32 {
	return _TraderReqExecOrderInsert(api.handle, pInputExecOrder, nRequestID)
}

// ReqExecOrderAction 执行宣告操作请求
func (api *TraderApi) ReqExecOrderAction(pInputExecOrderAction *CThostFtdcInputExecOrderActionField, nRequestID int32) int32 {
	return _TraderReqExecOrderAction(api.handle, pInputExecOrderAction, nRequestID)
}

// ReqForQuoteInsert 询价录入请求
func (api *TraderApi) ReqForQuoteInsert(pInputForQuote *CThostFtdcInputForQuoteField, nRequestID int32) int32 {
	return _TraderReqForQuoteInsert(api.handle, pInputForQuote, nRequestID)
}

// ReqQuoteInsert 报价录入请求
func (api *TraderApi) ReqQuoteInsert(pInputQuote *CThostFtdcInputQuoteField, nRequestID int32) int32 {
	return _TraderReqQuoteInsert(api.handle, pInputQuote, nRequestID)
}

// ReqQuoteAction 报价操作请求
func (api *TraderApi) ReqQuoteAction(pInputQuoteAction *CThostFtdcInputQuoteActionField, nRequestID int32) int32 {
	return _TraderReqQuoteAction(api.handle, pInputQuoteAction, nRequestID)
}

// ReqBatchOrderAction 批量报单操作请求
func (api *TraderApi) ReqBatchOrderAction(pInputBatchOrderAction *CThostFtdcInputBatchOrderActionField, nRequestID int32) int32 {
	return _TraderReqBatchOrderAction(api.handle, pInputBatchOrderAction, nRequestID)
}

// ReqOptionSelfCloseInsert 期权自对冲录入请求
func (api *TraderApi) ReqOptionSelfCloseInsert(pInputOptionSelfClose *CThostFtdcInputOptionSelfCloseField, nRequestID int32) int32 {
	return _TraderReqOptionSelfCloseInsert(api.handle, pInputOptionSelfClose, nRequestID)
}

// ReqOptionSelfCloseAction 期权自对冲操作请求
func (api *TraderApi) ReqOptionSelfCloseAction(pInputOptionSelfCloseAction *CThostFtdcInputOptionSelfCloseActionField, nRequestID int32) int32 {
	return _TraderReqOptionSelfCloseAction(api.handle, pInputOptionSelfCloseAction, nRequestID)
}

// ReqCombActionInsert 申请组合录入请求
func (api *TraderApi) ReqCombActionInsert(pInputCombAction *CThostFtdcInputCombActionField, nRequestID int32) int32 {
	return _TraderReqCombActionInsert(api.handle, pInputCombAction, nRequestID)
}

// ReqQryOrder 请求查询报单
func (api *TraderApi) ReqQryOrder(pQryOrder *CThostFtdcQryOrderField, nRequestID int32) int32 {
	return _TraderReqQryOrder(api.handle, pQryOrder, nRequestID)
}

// ReqQryTrade 请求查询成交
func (api *TraderApi) ReqQryTrade(pQryTrade *CThostFtdcQryTradeField, nRequestID int32) int32 {
	return _TraderReqQryTrade(api.handle, pQryTrade, nRequestID)
}

// ReqQryInvestorPosition 请求查询投资者持仓
func (api *TraderApi) ReqQryInvestorPosition(pQryInvestorPosition *CThostFtdcQryInvestorPositionField, nRequestID int32) int32 {
	return _TraderReqQryInvestorPosition(api.handle, pQryInvestorPosition, nRequestID)
}

// ReqQryTradingAccount 请求查询资金账户
func (api *TraderApi) ReqQryTradingAccount(pQryTradingAccount *CThostFtdcQryTradingAccountField, nRequestID int32) int32 {
	return _TraderReqQryTradingAccount(api.handle, pQryTradingAccount, nRequestID)
}

// ReqQryInvestor 请求查询投资者
func (api *TraderApi) ReqQryInvestor(pQryInvestor *CThostFtdcQryInvestorField, nRequestID int32) int32 {
	return _TraderReqQryInvestor(api.handle, pQryInvestor, nRequestID)
}

// ReqQryTradingCode 请求查询交易编码
func (api *TraderApi) ReqQryTradingCode(pQryTradingCode *CThostFtdcQryTradingCodeField, nRequestID int32) int32 {
	return _TraderReqQryTradingCode(api.handle, pQryTradingCode, nRequestID)
}

// ReqQryInstrumentMarginRate 请求查询合约保证金率
func (api *TraderApi) ReqQryInstrumentMarginRate(pQryInstrumentMarginRate *CThostFtdcQryInstrumentMarginRateField, nRequestID int32) int32 {
	return _TraderReqQryInstrumentMarginRate(api.handle, pQryInstrumentMarginRate, nRequestID)
}

// ReqQryInstrumentCommissionRate 请求查询合约手续费率
func (api *TraderApi) ReqQryInstrumentCommissionRate(pQryInstrumentCommissionRate *CThostFtdcQryInstrumentCommissionRateField, nRequestID int32) int32 {
	return _TraderReqQryInstrumentCommissionRate(api.handle, pQryInstrumentCommissionRate, nRequestID)
}

// ReqQryExchange 请求查询交易所
func (api *TraderApi) ReqQryExchange(pQryExchange *CThostFtdcQryExchangeField, nRequestID int32) int32 {
	return _TraderReqQryExchange(api.handle, pQryExchange, nRequestID)
}

// ReqQryProduct 请求查询产品
func (api *TraderApi) ReqQryProduct(pQryProduct *CThostFtdcQryProductField, nRequestID int32) int32 {
	return _TraderReqQryProduct(api.handle, pQryProduct, nRequestID)
}

// ReqQryInstrument 请求查询合约
func (api *TraderApi) ReqQryInstrument(pQryInstrument *CThostFtdcQryInstrumentField, nRequestID int32) int32 {
	return _TraderReqQryInstrument(api.handle, pQryInstrument, nRequestID)
}

// ReqQryDepthMarketData 请求查询行情
func (api *TraderApi) ReqQryDepthMarketData(pQryDepthMarketData *CThostFtdcQryDepthMarketDataField, nRequestID int32) int32 {
	return _TraderReqQryDepthMarketData(api.handle, pQryDepthMarketData, nRequestID)
}

// ReqQryTraderOffer 请求查询交易员报盘机
func (api *TraderApi) ReqQryTraderOffer(pQryTraderOffer *CThostFtdcQryTraderOfferField, nRequestID int32) int32 {
	return _TraderReqQryTraderOffer(api.handle, pQryTraderOffer, nRequestID)
}

// ReqQrySettlementInfo 请求查询投资者结算结果
func (api *TraderApi) ReqQrySettlementInfo(pQrySettlementInfo *CThostFtdcQrySettlementInfoField, nRequestID int32) int32 {
	return _TraderReqQrySettlementInfo(api.handle, pQrySettlementInfo, nRequestID)
}

// ReqQryTransferBank 请求查询转帐银行
func (api *TraderApi) ReqQryTransferBank(pQryTransferBank *CThostFtdcQryTransferBankField, nRequestID int32) int32 {
	return _TraderReqQryTransferBank(api.handle, pQryTransferBank, nRequestID)
}

// ReqQryInvestorPositionDetail 请求查询投资者持仓明细
func (api *TraderApi) ReqQryInvestorPositionDetail(pQryInvestorPositionDetail *CThostFtdcQryInvestorPositionDetailField, nRequestID int32) int32 {
	return _TraderReqQryInvestorPositionDetail(api.handle, pQryInvestorPositionDetail, nRequestID)
}

// ReqQryNotice 请求查询客户通知
func (api *TraderApi) ReqQryNotice(pQryNotice *CThostFtdcQryNoticeField, nRequestID int32) int32 {
	return _TraderReqQryNotice(api.handle, pQryNotice, nRequestID)
}

// ReqQrySettlementInfoConfirm 请求查询结算信息确认
func (api *TraderApi) ReqQrySettlementInfoConfirm(pQrySettlementInfoConfirm *CThostFtdcQrySettlementInfoConfirmField, nRequestID int32) int32 {
	return _TraderReqQrySettlementInfoConfirm(api.handle, pQrySettlementInfoConfirm, nRequestID)
}

// ReqQryInvestorPositionCombineDetail 请求查询投资者持仓明细
func (api *TraderApi) ReqQryInvestorPositionCombineDetail(pQryInvestorPositionCombineDetail *CThostFtdcQryInvestorPositionCombineDetailField, nRequestID int32) int32 {
	return _TraderReqQryInvestorPositionCombineDetail(api.handle, pQryInvestorPositionCombineDetail, nRequestID)
}

// ReqQryCFMMCTradingAccountKey 请求查询保证金监管系统经纪公司资金账户密钥
func (api *TraderApi) ReqQryCFMMCTradingAccountKey(pQryCFMMCTradingAccountKey *CThostFtdcQryCFMMCTradingAccountKeyField, nRequestID int32) int32 {
	return _TraderReqQryCFMMCTradingAccountKey(api.handle, pQryCFMMCTradingAccountKey, nRequestID)
}

// ReqQryEWarrantOffset 请求查询仓单折抵信息
func (api *TraderApi) ReqQryEWarrantOffset(pQryEWarrantOffset *CThostFtdcQryEWarrantOffsetField, nRequestID int32) int32 {
	return _TraderReqQryEWarrantOffset(api.handle, pQryEWarrantOffset, nRequestID)
}

// ReqQryInvestorProductGroupMargin 请求查询投资者品种/跨品种保证金
func (api *TraderApi) ReqQryInvestorProductGroupMargin(pQryInvestorProductGroupMargin *CThostFtdcQryInvestorProductGroupMarginField, nRequestID int32) int32 {
	return _TraderReqQryInvestorProductGroupMargin(api.handle, pQryInvestorProductGroupMargin, nRequestID)
}

// ReqQryExchangeMarginRate 请求查询交易所保证金率
func (api *TraderApi) ReqQryExchangeMarginRate(pQryExchangeMarginRate *CThostFtdcQryExchangeMarginRateField, nRequestID int32) int32 {
	return _TraderReqQryExchangeMarginRate(api.handle, pQryExchangeMarginRate, nRequestID)
}

// ReqQryExchangeMarginRateAdjust 请求查询交易所调整保证金率
func (api *TraderApi) ReqQryExchangeMarginRateAdjust(pQryExchangeMarginRateAdjust *CThostFtdcQryExchangeMarginRateAdjustField, nRequestID int32) int32 {
	return _TraderReqQryExchangeMarginRateAdjust(api.handle, pQryExchangeMarginRateAdjust, nRequestID)
}

// ReqQryExchangeRate 请求查询汇率
func (api *TraderApi) ReqQryExchangeRate(pQryExchangeRate *CThostFtdcQryExchangeRateField, nRequestID int32) int32 {
	return _TraderReqQryExchangeRate(api.handle, pQryExchangeRate, nRequestID)
}

// ReqQrySecAgentACIDMap 请求查询二级代理操作员银期权限
func (api *TraderApi) ReqQrySecAgentACIDMap(pQrySecAgentACIDMap *CThostFtdcQrySecAgentACIDMapField, nRequestID int32) int32 {
	return _TraderReqQrySecAgentACIDMap(api.handle, pQrySecAgentACIDMap, nRequestID)
}

// ReqQryProductExchRate 请求查询产品报价汇率
func (api *TraderApi) ReqQryProductExchRate(pQryProductExchRate *CThostFtdcQryProductExchRateField, nRequestID int32) int32 {
	return _TraderReqQryProductExchRate(api.handle, pQryProductExchRate, nRequestID)
}

// ReqQryProductGroup 请求查询产品组
func (api *TraderApi) ReqQryProductGroup(pQryProductGroup *CThostFtdcQryProductGroupField, nRequestID int32) int32 {
	return _TraderReqQryProductGroup(api.handle, pQryProductGroup, nRequestID)
}

// ReqQryMMInstrumentCommissionRate 请求查询做市商合约手续费率
func (api *TraderApi) ReqQryMMInstrumentCommissionRate(pQryMMInstrumentCommissionRate *CThostFtdcQryMMInstrumentCommissionRateField, nRequestID int32) int32 {
	return _TraderReqQryMMInstrumentCommissionRate(api.handle, pQryMMInstrumentCommissionRate, nRequestID)
}

// ReqQryMMOptionInstrCommRate 请求查询做市商期权合约手续费
func (api *TraderApi) ReqQryMMOptionInstrCommRate(pQryMMOptionInstrCommRate *CThostFtdcQryMMOptionInstrCommRateField, nRequestID int32) int32 {
	return _TraderReqQryMMOptionInstrCommRate(api.handle, pQryMMOptionInstrCommRate, nRequestID)
}

// ReqQryInstrumentOrderCommRate 请求查询报单手续费
func (api *TraderApi) ReqQryInstrumentOrderCommRate(pQryInstrumentOrderCommRate *CThostFtdcQryInstrumentOrderCommRateField, nRequestID int32) int32 {
	return _TraderReqQryInstrumentOrderCommRate(api.handle, pQryInstrumentOrderCommRate, nRequestID)
}

// ReqQrySecAgentTradingAccount 请求查询资金账户
func (api *TraderApi) ReqQrySecAgentTradingAccount(pQryTradingAccount *CThostFtdcQryTradingAccountField, nRequestID int32) int32 {
	return _TraderReqQrySecAgentTradingAccount(api.handle, pQryTradingAccount, nRequestID)
}

// ReqQrySecAgentCheckMode 请求查询二级代理商资金校验模式
func (api *TraderApi) ReqQrySecAgentCheckMode(pQrySecAgentCheckMode *CThostFtdcQrySecAgentCheckModeField, nRequestID int32) int32 {
	return _TraderReqQrySecAgentCheckMode(api.handle, pQrySecAgentCheckMode, nRequestID)
}

// ReqQrySecAgentTradeInfo 请求查询二级代理商信息
func (api *TraderApi) ReqQrySecAgentTradeInfo(pQrySecAgentTradeInfo *CThostFtdcQrySecAgentTradeInfoField, nRequestID int32) int32 {
	return _TraderReqQrySecAgentTradeInfo(api.handle, pQrySecAgentTradeInfo, nRequestID)
}

// ReqQryOptionInstrTradeCost 请求查询期权交易成本
func (api *TraderApi) ReqQryOptionInstrTradeCost(pQryOptionInstrTradeCost *CThostFtdcQryOptionInstrTradeCostField, nRequestID int32) int32 {
	return _TraderReqQryOptionInstrTradeCost(api.handle, pQryOptionInstrTradeCost, nRequestID)
}

// ReqQryOptionInstrCommRate 请求查询期权合约手续费
func (api *TraderApi) ReqQryOptionInstrCommRate(pQryOptionInstrCommRate *CThostFtdcQryOptionInstrCommRateField, nRequestID int32) int32 {
	return _TraderReqQryOptionInstrCommRate(api.handle, pQryOptionInstrCommRate, nRequestID)
}

// ReqQryExecOrder 请求查询执行宣告
func (api *TraderApi) ReqQryExecOrder(pQryExecOrder *CThostFtdcQryExecOrderField, nRequestID int32) int32 {
	return _TraderReqQryExecOrder(api.handle, pQryExecOrder, nRequestID)
}

// ReqQryForQuote 请求查询询价
func (api *TraderApi) ReqQryForQuote(pQryForQuote *CThostFtdcQryForQuoteField, nRequestID int32) int32 {
	return _TraderReqQryForQuote(api.handle, pQryForQuote, nRequestID)
}

// ReqQryQuote 请求查询报价
func (api *TraderApi) ReqQryQuote(pQryQuote *CThostFtdcQryQuoteField, nRequestID int32) int32 {
	return _TraderReqQryQuote(api.handle, pQryQuote, nRequestID)
}

// ReqQryOptionSelfClose 请求查询期权自对冲
func (api *TraderApi) ReqQryOptionSelfClose(pQryOptionSelfClose *CThostFtdcQryOptionSelfCloseField, nRequestID int32) int32 {
	return _TraderReqQryOptionSelfClose(api.handle, pQryOptionSelfClose, nRequestID)
}

// ReqQryInvestUnit 请求查询投资单元
func (api *TraderApi) ReqQryInvestUnit(pQryInvestUnit *CThostFtdcQryInvestUnitField, nRequestID int32) int32 {
	return _TraderReqQryInvestUnit(api.handle, pQryInvestUnit, nRequestID)
}

// ReqQryCombInstrumentGuard 请求查询组合合约安全系数
func (api *TraderApi) ReqQryCombInstrumentGuard(pQryCombInstrumentGuard *CThostFtdcQryCombInstrumentGuardField, nRequestID int32) int32 {
	return _TraderReqQryCombInstrumentGuard(api.handle, pQryCombInstrumentGuard, nRequestID)
}

// ReqQryCombAction 请求查询申请组合
func (api *TraderApi) ReqQryCombAction(pQryCombAction *CThostFtdcQryCombActionField, nRequestID int32) int32 {
	return _TraderReqQryCombAction(api.handle, pQryCombAction, nRequestID)
}

// ReqQryTransferSerial 请求查询转帐流水
func (api *TraderApi) ReqQryTransferSerial(pQryTransferSerial *CThostFtdcQryTransferSerialField, nRequestID int32) int32 {
	return _TraderReqQryTransferSerial(api.handle, pQryTransferSerial, nRequestID)
}

// ReqQryAccountregister 请求查询银期签约关系
func (api *TraderApi) ReqQryAccountregister(pQryAccountregister *CThostFtdcQryAccountregisterField, nRequestID int32) int32 {
	return _TraderReqQryAccountregister(api.handle, pQryAccountregister, nRequestID)
}

// ReqQryContractBank 请求查询签约银行
func (api *TraderApi) ReqQryContractBank(pQryContractBank *CThostFtdcQryContractBankField, nRequestID int32) int32 {
	return _TraderReqQryContractBank(api.handle, pQryContractBank, nRequestID)
}

// ReqQryParkedOrder 请求查询预埋单
func (api *TraderApi) ReqQryParkedOrder(pQryParkedOrder *CThostFtdcQryParkedOrderField, nRequestID int32) int32 {
	return _TraderReqQryParkedOrder(api.handle, pQryParkedOrder, nRequestID)
}

// ReqQryParkedOrderAction 请求查询预埋撤单
func (api *TraderApi) ReqQryParkedOrderAction(pQryParkedOrderAction *CThostFtdcQryParkedOrderActionField, nRequestID int32) int32 {
	return _TraderReqQryParkedOrderAction(api.handle, pQryParkedOrderAction, nRequestID)
}

// ReqQryTradingNotice 请求查询交易通知
func (api *TraderApi) ReqQryTradingNotice(pQryTradingNotice *CThostFtdcQryTradingNoticeField, nRequestID int32) int32 {
	return _TraderReqQryTradingNotice(api.handle, pQryTradingNotice, nRequestID)
}

// ReqQryBrokerTradingParams 请求查询经纪公司交易参数
func (api *TraderApi) ReqQryBrokerTradingParams(pQryBrokerTradingParams *CThostFtdcQryBrokerTradingParamsField, nRequestID int32) int32 {
	return _TraderReqQryBrokerTradingParams(api.handle, pQryBrokerTradingParams, nRequestID)
}

// ReqQryBrokerTradingAlgos 请求查询经纪公司交易算法
func (api *TraderApi) ReqQryBrokerTradingAlgos(pQryBrokerTradingAlgos *CThostFtdcQryBrokerTradingAlgosField, nRequestID int32) int32 {
	return _TraderReqQryBrokerTradingAlgos(api.handle, pQryBrokerTradingAlgos, nRequestID)
}

// ReqQueryCFMMCTradingAccountToken 请求查询监控中心用户令牌
func (api *TraderApi) ReqQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken *CThostFtdcQueryCFMMCTradingAccountTokenField, nRequestID int32) int32 {
	return _TraderReqQueryCFMMCTradingAccountToken(api.handle, pQueryCFMMCTradingAccountToken, nRequestID)
}

// ReqFromBankToFutureByFuture 期货发起银行资金转期货请求
func (api *TraderApi) ReqFromBankToFutureByFuture(pReqTransfer *CThostFtdcReqTransferField, nRequestID int32) int32 {
	return _TraderReqFromBankToFutureByFuture(api.handle, pReqTransfer, nRequestID)
}

// ReqFromFutureToBankByFuture 期货发起期货资金转银行请求
func (api *TraderApi) ReqFromFutureToBankByFuture(pReqTransfer *CThostFtdcReqTransferField, nRequestID int32) int32 {
	return _TraderReqFromFutureToBankByFuture(api.handle, pReqTransfer, nRequestID)
}

// ReqQueryBankAccountMoneyByFuture 期货发起查询银行余额请求
func (api *TraderApi) ReqQueryBankAccountMoneyByFuture(pReqQueryAccount *CThostFtdcReqQueryAccountField, nRequestID int32) int32 {
	return _TraderReqQueryBankAccountMoneyByFuture(api.handle, pReqQueryAccount, nRequestID)
}

// ReqQryClassifiedInstrument 请求查询分类合约
func (api *TraderApi) ReqQryClassifiedInstrument(pQryClassifiedInstrument *CThostFtdcQryClassifiedInstrumentField, nRequestID int32) int32 {
	return _TraderReqQryClassifiedInstrument(api.handle, pQryClassifiedInstrument, nRequestID)
}

// ReqQryCombPromotionParam 请求组合优惠比例
func (api *TraderApi) ReqQryCombPromotionParam(pQryCombPromotionParam *CThostFtdcQryCombPromotionParamField, nRequestID int32) int32 {
	return _TraderReqQryCombPromotionParam(api.handle, pQryCombPromotionParam, nRequestID)
}

// ReqQryRiskSettleInvstPosition 投资者风险结算持仓查询
func (api *TraderApi) ReqQryRiskSettleInvstPosition(pQryRiskSettleInvstPosition *CThostFtdcQryRiskSettleInvstPositionField, nRequestID int32) int32 {
	return _TraderReqQryRiskSettleInvstPosition(api.handle, pQryRiskSettleInvstPosition, nRequestID)
}

// ReqQryRiskSettleProductStatus 风险结算产品查询
func (api *TraderApi) ReqQryRiskSettleProductStatus(pQryRiskSettleProductStatus *CThostFtdcQryRiskSettleProductStatusField, nRequestID int32) int32 {
	return _TraderReqQryRiskSettleProductStatus(api.handle, pQryRiskSettleProductStatus, nRequestID)
}

// ReqQrySPBMFutureParameter SPBM期货合约参数查询
func (api *TraderApi) ReqQrySPBMFutureParameter(pQrySPBMFutureParameter *CThostFtdcQrySPBMFutureParameterField, nRequestID int32) int32 {
	return _TraderReqQrySPBMFutureParameter(api.handle, pQrySPBMFutureParameter, nRequestID)
}

// ReqQrySPBMOptionParameter SPBM期权合约参数查询
func (api *TraderApi) ReqQrySPBMOptionParameter(pQrySPBMOptionParameter *CThostFtdcQrySPBMOptionParameterField, nRequestID int32) int32 {
	return _TraderReqQrySPBMOptionParameter(api.handle, pQrySPBMOptionParameter, nRequestID)
}

// ReqQrySPBMIntraParameter SPBM品种内对锁仓折扣参数查询
func (api *TraderApi) ReqQrySPBMIntraParameter(pQrySPBMIntraParameter *CThostFtdcQrySPBMIntraParameterField, nRequestID int32) int32 {
	return _TraderReqQrySPBMIntraParameter(api.handle, pQrySPBMIntraParameter, nRequestID)
}

// ReqQrySPBMInterParameter SPBM跨品种抵扣参数查询
func (api *TraderApi) ReqQrySPBMInterParameter(pQrySPBMInterParameter *CThostFtdcQrySPBMInterParameterField, nRequestID int32) int32 {
	return _TraderReqQrySPBMInterParameter(api.handle, pQrySPBMInterParameter, nRequestID)
}

// ReqQrySPBMPortfDefinition SPBM组合保证金套餐查询
func (api *TraderApi) ReqQrySPBMPortfDefinition(pQrySPBMPortfDefinition *CThostFtdcQrySPBMPortfDefinitionField, nRequestID int32) int32 {
	return _TraderReqQrySPBMPortfDefinition(api.handle, pQrySPBMPortfDefinition, nRequestID)
}

// ReqQrySPBMInvestorPortfDef 投资者SPBM套餐选择查询
func (api *TraderApi) ReqQrySPBMInvestorPortfDef(pQrySPBMInvestorPortfDef *CThostFtdcQrySPBMInvestorPortfDefField, nRequestID int32) int32 {
	return _TraderReqQrySPBMInvestorPortfDef(api.handle, pQrySPBMInvestorPortfDef, nRequestID)
}

// ReqQryInvestorPortfMarginRatio 投资者新型组合保证金系数查询
func (api *TraderApi) ReqQryInvestorPortfMarginRatio(pQryInvestorPortfMarginRatio *CThostFtdcQryInvestorPortfMarginRatioField, nRequestID int32) int32 {
	return _TraderReqQryInvestorPortfMarginRatio(api.handle, pQryInvestorPortfMarginRatio, nRequestID)
}

// ReqQryInvestorProdSPBMDetail 投资者产品SPBM明细查询
func (api *TraderApi) ReqQryInvestorProdSPBMDetail(pQryInvestorProdSPBMDetail *CThostFtdcQryInvestorProdSPBMDetailField, nRequestID int32) int32 {
	return _TraderReqQryInvestorProdSPBMDetail(api.handle, pQryInvestorProdSPBMDetail, nRequestID)
}

// ReqQryInvestorCommoditySPMMMargin 投资者商品组SPMM记录查询
func (api *TraderApi) ReqQryInvestorCommoditySPMMMargin(pQryInvestorCommoditySPMMMargin *CThostFtdcQryInvestorCommoditySPMMMarginField, nRequestID int32) int32 {
	return _TraderReqQryInvestorCommoditySPMMMargin(api.handle, pQryInvestorCommoditySPMMMargin, nRequestID)
}

// ReqQryInvestorCommodityGroupSPMMMargin 投资者商品群SPMM记录查询
func (api *TraderApi) ReqQryInvestorCommodityGroupSPMMMargin(pQryInvestorCommodityGroupSPMMMargin *CThostFtdcQryInvestorCommodityGroupSPMMMarginField, nRequestID int32) int32 {
	return _TraderReqQryInvestorCommodityGroupSPMMMargin(api.handle, pQryInvestorCommodityGroupSPMMMargin, nRequestID)
}

// ReqQrySPMMInstParam SPMM合约参数查询
func (api *TraderApi) ReqQrySPMMInstParam(pQrySPMMInstParam *CThostFtdcQrySPMMInstParamField, nRequestID int32) int32 {
	return _TraderReqQrySPMMInstParam(api.handle, pQrySPMMInstParam, nRequestID)
}

// ReqQrySPMMProductParam SPMM产品参数查询
func (api *TraderApi) ReqQrySPMMProductParam(pQrySPMMProductParam *CThostFtdcQrySPMMProductParamField, nRequestID int32) int32 {
	return _TraderReqQrySPMMProductParam(api.handle, pQrySPMMProductParam, nRequestID)
}

// ReqQrySPBMAddOnInterParameter SPBM附加跨品种抵扣参数查询
func (api *TraderApi) ReqQrySPBMAddOnInterParameter(pQrySPBMAddOnInterParameter *CThostFtdcQrySPBMAddOnInterParameterField, nRequestID int32) int32 {
	return _TraderReqQrySPBMAddOnInterParameter(api.handle, pQrySPBMAddOnInterParameter, nRequestID)
}

// ReqQryRCAMSCombProductInfo RCAMS产品组合信息查询
func (api *TraderApi) ReqQryRCAMSCombProductInfo(pQryRCAMSCombProductInfo *CThostFtdcQryRCAMSCombProductInfoField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSCombProductInfo(api.handle, pQryRCAMSCombProductInfo, nRequestID)
}

// ReqQryRCAMSInstrParameter RCAMS同合约风险对冲参数查询
func (api *TraderApi) ReqQryRCAMSInstrParameter(pQryRCAMSInstrParameter *CThostFtdcQryRCAMSInstrParameterField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSInstrParameter(api.handle, pQryRCAMSInstrParameter, nRequestID)
}

// ReqQryRCAMSIntraParameter RCAMS品种内风险对冲参数查询
func (api *TraderApi) ReqQryRCAMSIntraParameter(pQryRCAMSIntraParameter *CThostFtdcQryRCAMSIntraParameterField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSIntraParameter(api.handle, pQryRCAMSIntraParameter, nRequestID)
}

// ReqQryRCAMSInterParameter RCAMS跨品种风险折抵参数查询
func (api *TraderApi) ReqQryRCAMSInterParameter(pQryRCAMSInterParameter *CThostFtdcQryRCAMSInterParameterField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSInterParameter(api.handle, pQryRCAMSInterParameter, nRequestID)
}

// ReqQryRCAMSShortOptAdjustParam RCAMS空头期权风险调整参数查询
func (api *TraderApi) ReqQryRCAMSShortOptAdjustParam(pQryRCAMSShortOptAdjustParam *CThostFtdcQryRCAMSShortOptAdjustParamField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSShortOptAdjustParam(api.handle, pQryRCAMSShortOptAdjustParam, nRequestID)
}

// ReqQryRCAMSInvestorCombPosition RCAMS策略组合持仓查询
func (api *TraderApi) ReqQryRCAMSInvestorCombPosition(pQryRCAMSInvestorCombPosition *CThostFtdcQryRCAMSInvestorCombPositionField, nRequestID int32) int32 {
	return _TraderReqQryRCAMSInvestorCombPosition(api.handle, pQryRCAMSInvestorCombPosition, nRequestID)
}

// ReqQryInvestorProdRCAMSMargin 投资者品种RCAMS保证金查询
func (api *TraderApi) ReqQryInvestorProdRCAMSMargin(pQryInvestorProdRCAMSMargin *CThostFtdcQryInvestorProdRCAMSMarginField, nRequestID int32) int32 {
	return _TraderReqQryInvestorProdRCAMSMargin(api.handle, pQryInvestorProdRCAMSMargin, nRequestID)
}

// ReqQryRULEInstrParameter RULE合约保证金参数查询
func (api *TraderApi) ReqQryRULEInstrParameter(pQryRULEInstrParameter *CThostFtdcQryRULEInstrParameterField, nRequestID int32) int32 {
	return _TraderReqQryRULEInstrParameter(api.handle, pQryRULEInstrParameter, nRequestID)
}

// ReqQryRULEIntraParameter RULE品种内对锁仓折扣参数查询
func (api *TraderApi) ReqQryRULEIntraParameter(pQryRULEIntraParameter *CThostFtdcQryRULEIntraParameterField, nRequestID int32) int32 {
	return _TraderReqQryRULEIntraParameter(api.handle, pQryRULEIntraParameter, nRequestID)
}

// ReqQryRULEInterParameter RULE跨品种抵扣参数查询
func (api *TraderApi) ReqQryRULEInterParameter(pQryRULEInterParameter *CThostFtdcQryRULEInterParameterField, nRequestID int32) int32 {
	return _TraderReqQryRULEInterParameter(api.handle, pQryRULEInterParameter, nRequestID)
}

// ReqQryInvestorProdRULEMargin 投资者产品RULE保证金查询
func (api *TraderApi) ReqQryInvestorProdRULEMargin(pQryInvestorProdRULEMargin *CThostFtdcQryInvestorProdRULEMarginField, nRequestID int32) int32 {
	return _TraderReqQryInvestorProdRULEMargin(api.handle, pQryInvestorProdRULEMargin, nRequestID)
}

// ReqQryInvestorPortfSetting 投资者投资者新组保设置查询
func (api *TraderApi) ReqQryInvestorPortfSetting(pQryInvestorPortfSetting *CThostFtdcQryInvestorPortfSettingField, nRequestID int32) int32 {
	return _TraderReqQryInvestorPortfSetting(api.handle, pQryInvestorPortfSetting, nRequestID)
}

// SpiCreate ========== Trader SPI 函数 ========== 创建 SPI 实例
func (api *TraderApi) SpiCreate() uintptr {
	return _TraderSpiCreate(api.handle)
}

// SpiDestroy 销毁 SPI 实例
func (api *TraderApi) SpiDestroy() {
	_TraderSpiDestroy(api.handle)
}

// RegisterSpi 注册 SPI 到 API
func (api *TraderApi) RegisterSpi(spi uintptr) {
	_TraderRegisterSpi(api.handle, spi)
}

// SpiSetCallbacks 批量设置回调
func (api *TraderApi) SpiSetCallbacks(callbacks *TraderSpiCallbacks) {
	_TraderSpiSetCallbacks(api.handle, callbacks)
}

// ReqUserLoginWithSystemInfo ========== 跨平台统一登录接口 ========== 说明: macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数 此函数在 Linux/Windows 上忽略 systemInfo，在 macOS 上使用它 带系统信息的用户登录请求（跨平台统一接口） systemInfoLen: 系统信息长度，传 0 表示自动采集（仅 macOS 生效） systemInfo: 系统信息数据，传 NULL 表示自动采集（仅 macOS 生效）
func (api *TraderApi) ReqUserLoginWithSystemInfo(pReqUserLoginField *CThostFtdcReqUserLoginField, nRequestID int32, systemInfoLen int32, systemInfo string) int32 {
	return _TraderReqUserLoginWithSystemInfo(api.handle, pReqUserLoginField, nRequestID, systemInfoLen, CString(systemInfo))
}

// ========== SPI 回调设置方法 ==========

// SpiSetOnFrontConnected ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (api *TraderApi) SpiSetOnFrontConnected(callback TraderOnFrontConnectedCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnFrontConnected(api.spiHandle, ptr)
}

// SpiSetOnFrontDisconnected 0x2003 收到错误报文
func (api *TraderApi) SpiSetOnFrontDisconnected(callback TraderOnFrontDisconnectedCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnFrontDisconnected(api.spiHandle, ptr)
}

// SpiSetOnHeartBeatWarning 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (api *TraderApi) SpiSetOnHeartBeatWarning(callback TraderOnHeartBeatWarningCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnHeartBeatWarning(api.spiHandle, ptr)
}

// SpiSetOnRspAuthenticate 客户端认证响应
func (api *TraderApi) SpiSetOnRspAuthenticate(callback TraderOnRspAuthenticateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspAuthenticate(api.spiHandle, ptr)
}

// SpiSetOnRspUserLogin 登录请求响应
func (api *TraderApi) SpiSetOnRspUserLogin(callback TraderOnRspUserLoginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspUserLogin(api.spiHandle, ptr)
}

// SpiSetOnRspUserLogout 登出请求响应
func (api *TraderApi) SpiSetOnRspUserLogout(callback TraderOnRspUserLogoutCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspUserLogout(api.spiHandle, ptr)
}

// SpiSetOnRspUserPasswordUpdate 用户口令更新请求响应
func (api *TraderApi) SpiSetOnRspUserPasswordUpdate(callback TraderOnRspUserPasswordUpdateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspUserPasswordUpdate(api.spiHandle, ptr)
}

// SpiSetOnRspTradingAccountPasswordUpdate 资金账户口令更新请求响应
func (api *TraderApi) SpiSetOnRspTradingAccountPasswordUpdate(callback TraderOnRspTradingAccountPasswordUpdateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspTradingAccountPasswordUpdate(api.spiHandle, ptr)
}

// SpiSetOnRspUserAuthMethod 查询用户当前支持的认证模式的回复
func (api *TraderApi) SpiSetOnRspUserAuthMethod(callback TraderOnRspUserAuthMethodCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspUserAuthMethod(api.spiHandle, ptr)
}

// SpiSetOnRspGenUserCaptcha 获取图形验证码请求的回复
func (api *TraderApi) SpiSetOnRspGenUserCaptcha(callback TraderOnRspGenUserCaptchaCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspGenUserCaptcha(api.spiHandle, ptr)
}

// SpiSetOnRspGenUserText 获取短信验证码请求的回复
func (api *TraderApi) SpiSetOnRspGenUserText(callback TraderOnRspGenUserTextCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspGenUserText(api.spiHandle, ptr)
}

// SpiSetOnRspOrderInsert 报单录入请求响应
func (api *TraderApi) SpiSetOnRspOrderInsert(callback TraderOnRspOrderInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspOrderInsert(api.spiHandle, ptr)
}

// SpiSetOnRspParkedOrderInsert 预埋单录入请求响应
func (api *TraderApi) SpiSetOnRspParkedOrderInsert(callback TraderOnRspParkedOrderInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspParkedOrderInsert(api.spiHandle, ptr)
}

// SpiSetOnRspParkedOrderAction 预埋撤单录入请求响应
func (api *TraderApi) SpiSetOnRspParkedOrderAction(callback TraderOnRspParkedOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspParkedOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspOrderAction 报单操作请求响应
func (api *TraderApi) SpiSetOnRspOrderAction(callback TraderOnRspOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspQryMaxOrderVolume 查询最大报单数量响应
func (api *TraderApi) SpiSetOnRspQryMaxOrderVolume(callback TraderOnRspQryMaxOrderVolumeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryMaxOrderVolume(api.spiHandle, ptr)
}

// SpiSetOnRspSettlementInfoConfirm 投资者结算结果确认响应
func (api *TraderApi) SpiSetOnRspSettlementInfoConfirm(callback TraderOnRspSettlementInfoConfirmCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspSettlementInfoConfirm(api.spiHandle, ptr)
}

// SpiSetOnRspRemoveParkedOrder 删除预埋单响应
func (api *TraderApi) SpiSetOnRspRemoveParkedOrder(callback TraderOnRspRemoveParkedOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspRemoveParkedOrder(api.spiHandle, ptr)
}

// SpiSetOnRspRemoveParkedOrderAction 删除预埋撤单响应
func (api *TraderApi) SpiSetOnRspRemoveParkedOrderAction(callback TraderOnRspRemoveParkedOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspRemoveParkedOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspExecOrderInsert 执行宣告录入请求响应
func (api *TraderApi) SpiSetOnRspExecOrderInsert(callback TraderOnRspExecOrderInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspExecOrderInsert(api.spiHandle, ptr)
}

// SpiSetOnRspExecOrderAction 执行宣告操作请求响应
func (api *TraderApi) SpiSetOnRspExecOrderAction(callback TraderOnRspExecOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspExecOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspForQuoteInsert 询价录入请求响应
func (api *TraderApi) SpiSetOnRspForQuoteInsert(callback TraderOnRspForQuoteInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspForQuoteInsert(api.spiHandle, ptr)
}

// SpiSetOnRspQuoteInsert 报价录入请求响应
func (api *TraderApi) SpiSetOnRspQuoteInsert(callback TraderOnRspQuoteInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQuoteInsert(api.spiHandle, ptr)
}

// SpiSetOnRspQuoteAction 报价操作请求响应
func (api *TraderApi) SpiSetOnRspQuoteAction(callback TraderOnRspQuoteActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQuoteAction(api.spiHandle, ptr)
}

// SpiSetOnRspBatchOrderAction 批量报单操作请求响应
func (api *TraderApi) SpiSetOnRspBatchOrderAction(callback TraderOnRspBatchOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspBatchOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspOptionSelfCloseInsert 期权自对冲录入请求响应
func (api *TraderApi) SpiSetOnRspOptionSelfCloseInsert(callback TraderOnRspOptionSelfCloseInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspOptionSelfCloseInsert(api.spiHandle, ptr)
}

// SpiSetOnRspOptionSelfCloseAction 期权自对冲操作请求响应
func (api *TraderApi) SpiSetOnRspOptionSelfCloseAction(callback TraderOnRspOptionSelfCloseActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspOptionSelfCloseAction(api.spiHandle, ptr)
}

// SpiSetOnRspCombActionInsert 申请组合录入请求响应
func (api *TraderApi) SpiSetOnRspCombActionInsert(callback TraderOnRspCombActionInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspCombActionInsert(api.spiHandle, ptr)
}

// SpiSetOnRspQryOrder 请求查询报单响应
func (api *TraderApi) SpiSetOnRspQryOrder(callback TraderOnRspQryOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryOrder(api.spiHandle, ptr)
}

// SpiSetOnRspQryTrade 请求查询成交响应
func (api *TraderApi) SpiSetOnRspQryTrade(callback TraderOnRspQryTradeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTrade(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorPosition 请求查询投资者持仓响应
func (api *TraderApi) SpiSetOnRspQryInvestorPosition(callback TraderOnRspQryInvestorPositionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorPosition(api.spiHandle, ptr)
}

// SpiSetOnRspQryTradingAccount 请求查询资金账户响应
func (api *TraderApi) SpiSetOnRspQryTradingAccount(callback TraderOnRspQryTradingAccountCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTradingAccount(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestor 请求查询投资者响应
func (api *TraderApi) SpiSetOnRspQryInvestor(callback TraderOnRspQryInvestorCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestor(api.spiHandle, ptr)
}

// SpiSetOnRspQryTradingCode 请求查询交易编码响应
func (api *TraderApi) SpiSetOnRspQryTradingCode(callback TraderOnRspQryTradingCodeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTradingCode(api.spiHandle, ptr)
}

// SpiSetOnRspQryInstrumentMarginRate 请求查询合约保证金率响应
func (api *TraderApi) SpiSetOnRspQryInstrumentMarginRate(callback TraderOnRspQryInstrumentMarginRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInstrumentMarginRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryInstrumentCommissionRate 请求查询合约手续费率响应
func (api *TraderApi) SpiSetOnRspQryInstrumentCommissionRate(callback TraderOnRspQryInstrumentCommissionRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInstrumentCommissionRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryExchange 请求查询交易所响应
func (api *TraderApi) SpiSetOnRspQryExchange(callback TraderOnRspQryExchangeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryExchange(api.spiHandle, ptr)
}

// SpiSetOnRspQryProduct 请求查询产品响应
func (api *TraderApi) SpiSetOnRspQryProduct(callback TraderOnRspQryProductCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryProduct(api.spiHandle, ptr)
}

// SpiSetOnRspQryInstrument 请求查询合约响应
func (api *TraderApi) SpiSetOnRspQryInstrument(callback TraderOnRspQryInstrumentCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInstrument(api.spiHandle, ptr)
}

// SpiSetOnRspQryDepthMarketData 请求查询行情响应
func (api *TraderApi) SpiSetOnRspQryDepthMarketData(callback TraderOnRspQryDepthMarketDataCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryDepthMarketData(api.spiHandle, ptr)
}

// SpiSetOnRspQryTraderOffer 请求查询交易员报盘机响应
func (api *TraderApi) SpiSetOnRspQryTraderOffer(callback TraderOnRspQryTraderOfferCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTraderOffer(api.spiHandle, ptr)
}

// SpiSetOnRspQrySettlementInfo 请求查询投资者结算结果响应
func (api *TraderApi) SpiSetOnRspQrySettlementInfo(callback TraderOnRspQrySettlementInfoCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySettlementInfo(api.spiHandle, ptr)
}

// SpiSetOnRspQryTransferBank 请求查询转帐银行响应
func (api *TraderApi) SpiSetOnRspQryTransferBank(callback TraderOnRspQryTransferBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTransferBank(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorPositionDetail 请求查询投资者持仓明细响应
func (api *TraderApi) SpiSetOnRspQryInvestorPositionDetail(callback TraderOnRspQryInvestorPositionDetailCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorPositionDetail(api.spiHandle, ptr)
}

// SpiSetOnRspQryNotice 请求查询客户通知响应
func (api *TraderApi) SpiSetOnRspQryNotice(callback TraderOnRspQryNoticeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryNotice(api.spiHandle, ptr)
}

// SpiSetOnRspQrySettlementInfoConfirm 请求查询结算信息确认响应
func (api *TraderApi) SpiSetOnRspQrySettlementInfoConfirm(callback TraderOnRspQrySettlementInfoConfirmCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySettlementInfoConfirm(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorPositionCombineDetail 请求查询投资者持仓明细响应
func (api *TraderApi) SpiSetOnRspQryInvestorPositionCombineDetail(callback TraderOnRspQryInvestorPositionCombineDetailCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorPositionCombineDetail(api.spiHandle, ptr)
}

// SpiSetOnRspQryCFMMCTradingAccountKey 查询保证金监管系统经纪公司资金账户密钥响应
func (api *TraderApi) SpiSetOnRspQryCFMMCTradingAccountKey(callback TraderOnRspQryCFMMCTradingAccountKeyCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryCFMMCTradingAccountKey(api.spiHandle, ptr)
}

// SpiSetOnRspQryEWarrantOffset 请求查询仓单折抵信息响应
func (api *TraderApi) SpiSetOnRspQryEWarrantOffset(callback TraderOnRspQryEWarrantOffsetCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryEWarrantOffset(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorProductGroupMargin 请求查询投资者品种/跨品种保证金响应
func (api *TraderApi) SpiSetOnRspQryInvestorProductGroupMargin(callback TraderOnRspQryInvestorProductGroupMarginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorProductGroupMargin(api.spiHandle, ptr)
}

// SpiSetOnRspQryExchangeMarginRate 请求查询交易所保证金率响应
func (api *TraderApi) SpiSetOnRspQryExchangeMarginRate(callback TraderOnRspQryExchangeMarginRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryExchangeMarginRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryExchangeMarginRateAdjust 请求查询交易所调整保证金率响应
func (api *TraderApi) SpiSetOnRspQryExchangeMarginRateAdjust(callback TraderOnRspQryExchangeMarginRateAdjustCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryExchangeMarginRateAdjust(api.spiHandle, ptr)
}

// SpiSetOnRspQryExchangeRate 请求查询汇率响应
func (api *TraderApi) SpiSetOnRspQryExchangeRate(callback TraderOnRspQryExchangeRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryExchangeRate(api.spiHandle, ptr)
}

// SpiSetOnRspQrySecAgentACIDMap 请求查询二级代理操作员银期权限响应
func (api *TraderApi) SpiSetOnRspQrySecAgentACIDMap(callback TraderOnRspQrySecAgentACIDMapCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySecAgentACIDMap(api.spiHandle, ptr)
}

// SpiSetOnRspQryProductExchRate 请求查询产品报价汇率
func (api *TraderApi) SpiSetOnRspQryProductExchRate(callback TraderOnRspQryProductExchRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryProductExchRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryProductGroup 请求查询产品组
func (api *TraderApi) SpiSetOnRspQryProductGroup(callback TraderOnRspQryProductGroupCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryProductGroup(api.spiHandle, ptr)
}

// SpiSetOnRspQryMMInstrumentCommissionRate 请求查询做市商合约手续费率响应
func (api *TraderApi) SpiSetOnRspQryMMInstrumentCommissionRate(callback TraderOnRspQryMMInstrumentCommissionRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryMMInstrumentCommissionRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryMMOptionInstrCommRate 请求查询做市商期权合约手续费响应
func (api *TraderApi) SpiSetOnRspQryMMOptionInstrCommRate(callback TraderOnRspQryMMOptionInstrCommRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryMMOptionInstrCommRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryInstrumentOrderCommRate 请求查询报单手续费响应
func (api *TraderApi) SpiSetOnRspQryInstrumentOrderCommRate(callback TraderOnRspQryInstrumentOrderCommRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInstrumentOrderCommRate(api.spiHandle, ptr)
}

// SpiSetOnRspQrySecAgentTradingAccount 请求查询资金账户响应
func (api *TraderApi) SpiSetOnRspQrySecAgentTradingAccount(callback TraderOnRspQrySecAgentTradingAccountCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySecAgentTradingAccount(api.spiHandle, ptr)
}

// SpiSetOnRspQrySecAgentCheckMode 请求查询二级代理商资金校验模式响应
func (api *TraderApi) SpiSetOnRspQrySecAgentCheckMode(callback TraderOnRspQrySecAgentCheckModeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySecAgentCheckMode(api.spiHandle, ptr)
}

// SpiSetOnRspQrySecAgentTradeInfo 请求查询二级代理商信息响应
func (api *TraderApi) SpiSetOnRspQrySecAgentTradeInfo(callback TraderOnRspQrySecAgentTradeInfoCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySecAgentTradeInfo(api.spiHandle, ptr)
}

// SpiSetOnRspQryOptionInstrTradeCost 请求查询期权交易成本响应
func (api *TraderApi) SpiSetOnRspQryOptionInstrTradeCost(callback TraderOnRspQryOptionInstrTradeCostCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryOptionInstrTradeCost(api.spiHandle, ptr)
}

// SpiSetOnRspQryOptionInstrCommRate 请求查询期权合约手续费响应
func (api *TraderApi) SpiSetOnRspQryOptionInstrCommRate(callback TraderOnRspQryOptionInstrCommRateCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryOptionInstrCommRate(api.spiHandle, ptr)
}

// SpiSetOnRspQryExecOrder 请求查询执行宣告响应
func (api *TraderApi) SpiSetOnRspQryExecOrder(callback TraderOnRspQryExecOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryExecOrder(api.spiHandle, ptr)
}

// SpiSetOnRspQryForQuote 请求查询询价响应
func (api *TraderApi) SpiSetOnRspQryForQuote(callback TraderOnRspQryForQuoteCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryForQuote(api.spiHandle, ptr)
}

// SpiSetOnRspQryQuote 请求查询报价响应
func (api *TraderApi) SpiSetOnRspQryQuote(callback TraderOnRspQryQuoteCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryQuote(api.spiHandle, ptr)
}

// SpiSetOnRspQryOptionSelfClose 请求查询期权自对冲响应
func (api *TraderApi) SpiSetOnRspQryOptionSelfClose(callback TraderOnRspQryOptionSelfCloseCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryOptionSelfClose(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestUnit 请求查询投资单元响应
func (api *TraderApi) SpiSetOnRspQryInvestUnit(callback TraderOnRspQryInvestUnitCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestUnit(api.spiHandle, ptr)
}

// SpiSetOnRspQryCombInstrumentGuard 请求查询组合合约安全系数响应
func (api *TraderApi) SpiSetOnRspQryCombInstrumentGuard(callback TraderOnRspQryCombInstrumentGuardCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryCombInstrumentGuard(api.spiHandle, ptr)
}

// SpiSetOnRspQryCombAction 请求查询申请组合响应
func (api *TraderApi) SpiSetOnRspQryCombAction(callback TraderOnRspQryCombActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryCombAction(api.spiHandle, ptr)
}

// SpiSetOnRspQryTransferSerial 请求查询转帐流水响应
func (api *TraderApi) SpiSetOnRspQryTransferSerial(callback TraderOnRspQryTransferSerialCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTransferSerial(api.spiHandle, ptr)
}

// SpiSetOnRspQryAccountregister 请求查询银期签约关系响应
func (api *TraderApi) SpiSetOnRspQryAccountregister(callback TraderOnRspQryAccountregisterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryAccountregister(api.spiHandle, ptr)
}

// SpiSetOnRspError 错误应答
func (api *TraderApi) SpiSetOnRspError(callback TraderOnRspErrorCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspError(api.spiHandle, ptr)
}

// SpiSetOnRtnOrder 报单通知
func (api *TraderApi) SpiSetOnRtnOrder(callback TraderOnRtnOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnOrder(api.spiHandle, ptr)
}

// SpiSetOnRtnTrade 成交通知
func (api *TraderApi) SpiSetOnRtnTrade(callback TraderOnRtnTradeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnTrade(api.spiHandle, ptr)
}

// SpiSetOnErrRtnOrderInsert 报单录入错误回报
func (api *TraderApi) SpiSetOnErrRtnOrderInsert(callback TraderOnErrRtnOrderInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnOrderInsert(api.spiHandle, ptr)
}

// SpiSetOnErrRtnOrderAction 报单操作错误回报
func (api *TraderApi) SpiSetOnErrRtnOrderAction(callback TraderOnErrRtnOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRtnInstrumentStatus 合约交易状态通知
func (api *TraderApi) SpiSetOnRtnInstrumentStatus(callback TraderOnRtnInstrumentStatusCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnInstrumentStatus(api.spiHandle, ptr)
}

// SpiSetOnRtnBulletin 交易所公告通知
func (api *TraderApi) SpiSetOnRtnBulletin(callback TraderOnRtnBulletinCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnBulletin(api.spiHandle, ptr)
}

// SpiSetOnRtnTradingNotice 交易通知
func (api *TraderApi) SpiSetOnRtnTradingNotice(callback TraderOnRtnTradingNoticeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnTradingNotice(api.spiHandle, ptr)
}

// SpiSetOnRtnErrorConditionalOrder 提示条件单校验错误
func (api *TraderApi) SpiSetOnRtnErrorConditionalOrder(callback TraderOnRtnErrorConditionalOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnErrorConditionalOrder(api.spiHandle, ptr)
}

// SpiSetOnRtnExecOrder 执行宣告通知
func (api *TraderApi) SpiSetOnRtnExecOrder(callback TraderOnRtnExecOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnExecOrder(api.spiHandle, ptr)
}

// SpiSetOnErrRtnExecOrderInsert 执行宣告录入错误回报
func (api *TraderApi) SpiSetOnErrRtnExecOrderInsert(callback TraderOnErrRtnExecOrderInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnExecOrderInsert(api.spiHandle, ptr)
}

// SpiSetOnErrRtnExecOrderAction 执行宣告操作错误回报
func (api *TraderApi) SpiSetOnErrRtnExecOrderAction(callback TraderOnErrRtnExecOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnExecOrderAction(api.spiHandle, ptr)
}

// SpiSetOnErrRtnForQuoteInsert 询价录入错误回报
func (api *TraderApi) SpiSetOnErrRtnForQuoteInsert(callback TraderOnErrRtnForQuoteInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnForQuoteInsert(api.spiHandle, ptr)
}

// SpiSetOnRtnQuote 报价通知
func (api *TraderApi) SpiSetOnRtnQuote(callback TraderOnRtnQuoteCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnQuote(api.spiHandle, ptr)
}

// SpiSetOnErrRtnQuoteInsert 报价录入错误回报
func (api *TraderApi) SpiSetOnErrRtnQuoteInsert(callback TraderOnErrRtnQuoteInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnQuoteInsert(api.spiHandle, ptr)
}

// SpiSetOnErrRtnQuoteAction 报价操作错误回报
func (api *TraderApi) SpiSetOnErrRtnQuoteAction(callback TraderOnErrRtnQuoteActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnQuoteAction(api.spiHandle, ptr)
}

// SpiSetOnRtnForQuoteRsp 询价通知
func (api *TraderApi) SpiSetOnRtnForQuoteRsp(callback TraderOnRtnForQuoteRspCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnForQuoteRsp(api.spiHandle, ptr)
}

// SpiSetOnRtnCFMMCTradingAccountToken 保证金监控中心用户令牌
func (api *TraderApi) SpiSetOnRtnCFMMCTradingAccountToken(callback TraderOnRtnCFMMCTradingAccountTokenCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnCFMMCTradingAccountToken(api.spiHandle, ptr)
}

// SpiSetOnErrRtnBatchOrderAction 批量报单操作错误回报
func (api *TraderApi) SpiSetOnErrRtnBatchOrderAction(callback TraderOnErrRtnBatchOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnBatchOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRtnOptionSelfClose 期权自对冲通知
func (api *TraderApi) SpiSetOnRtnOptionSelfClose(callback TraderOnRtnOptionSelfCloseCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnOptionSelfClose(api.spiHandle, ptr)
}

// SpiSetOnErrRtnOptionSelfCloseInsert 期权自对冲录入错误回报
func (api *TraderApi) SpiSetOnErrRtnOptionSelfCloseInsert(callback TraderOnErrRtnOptionSelfCloseInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnOptionSelfCloseInsert(api.spiHandle, ptr)
}

// SpiSetOnErrRtnOptionSelfCloseAction 期权自对冲操作错误回报
func (api *TraderApi) SpiSetOnErrRtnOptionSelfCloseAction(callback TraderOnErrRtnOptionSelfCloseActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnOptionSelfCloseAction(api.spiHandle, ptr)
}

// SpiSetOnRtnCombAction 申请组合通知
func (api *TraderApi) SpiSetOnRtnCombAction(callback TraderOnRtnCombActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnCombAction(api.spiHandle, ptr)
}

// SpiSetOnErrRtnCombActionInsert 申请组合录入错误回报
func (api *TraderApi) SpiSetOnErrRtnCombActionInsert(callback TraderOnErrRtnCombActionInsertCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnCombActionInsert(api.spiHandle, ptr)
}

// SpiSetOnRspQryContractBank 请求查询签约银行响应
func (api *TraderApi) SpiSetOnRspQryContractBank(callback TraderOnRspQryContractBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryContractBank(api.spiHandle, ptr)
}

// SpiSetOnRspQryParkedOrder 请求查询预埋单响应
func (api *TraderApi) SpiSetOnRspQryParkedOrder(callback TraderOnRspQryParkedOrderCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryParkedOrder(api.spiHandle, ptr)
}

// SpiSetOnRspQryParkedOrderAction 请求查询预埋撤单响应
func (api *TraderApi) SpiSetOnRspQryParkedOrderAction(callback TraderOnRspQryParkedOrderActionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryParkedOrderAction(api.spiHandle, ptr)
}

// SpiSetOnRspQryTradingNotice 请求查询交易通知响应
func (api *TraderApi) SpiSetOnRspQryTradingNotice(callback TraderOnRspQryTradingNoticeCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryTradingNotice(api.spiHandle, ptr)
}

// SpiSetOnRspQryBrokerTradingParams 请求查询经纪公司交易参数响应
func (api *TraderApi) SpiSetOnRspQryBrokerTradingParams(callback TraderOnRspQryBrokerTradingParamsCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryBrokerTradingParams(api.spiHandle, ptr)
}

// SpiSetOnRspQryBrokerTradingAlgos 请求查询经纪公司交易算法响应
func (api *TraderApi) SpiSetOnRspQryBrokerTradingAlgos(callback TraderOnRspQryBrokerTradingAlgosCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryBrokerTradingAlgos(api.spiHandle, ptr)
}

// SpiSetOnRspQueryCFMMCTradingAccountToken 请求查询监控中心用户令牌
func (api *TraderApi) SpiSetOnRspQueryCFMMCTradingAccountToken(callback TraderOnRspQueryCFMMCTradingAccountTokenCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQueryCFMMCTradingAccountToken(api.spiHandle, ptr)
}

// SpiSetOnRtnFromBankToFutureByBank 银行发起银行资金转期货通知
func (api *TraderApi) SpiSetOnRtnFromBankToFutureByBank(callback TraderOnRtnFromBankToFutureByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnFromBankToFutureByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnFromFutureToBankByBank 银行发起期货资金转银行通知
func (api *TraderApi) SpiSetOnRtnFromFutureToBankByBank(callback TraderOnRtnFromFutureToBankByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnFromFutureToBankByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromBankToFutureByBank 银行发起冲正银行转期货通知
func (api *TraderApi) SpiSetOnRtnRepealFromBankToFutureByBank(callback TraderOnRtnRepealFromBankToFutureByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromBankToFutureByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromFutureToBankByBank 银行发起冲正期货转银行通知
func (api *TraderApi) SpiSetOnRtnRepealFromFutureToBankByBank(callback TraderOnRtnRepealFromFutureToBankByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromFutureToBankByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnFromBankToFutureByFuture 期货发起银行资金转期货通知
func (api *TraderApi) SpiSetOnRtnFromBankToFutureByFuture(callback TraderOnRtnFromBankToFutureByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnFromBankToFutureByFuture(api.spiHandle, ptr)
}

// SpiSetOnRtnFromFutureToBankByFuture 期货发起期货资金转银行通知
func (api *TraderApi) SpiSetOnRtnFromFutureToBankByFuture(callback TraderOnRtnFromFutureToBankByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnFromFutureToBankByFuture(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromBankToFutureByFutureManual 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (api *TraderApi) SpiSetOnRtnRepealFromBankToFutureByFutureManual(callback TraderOnRtnRepealFromBankToFutureByFutureManualCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromFutureToBankByFutureManual 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (api *TraderApi) SpiSetOnRtnRepealFromFutureToBankByFutureManual(callback TraderOnRtnRepealFromFutureToBankByFutureManualCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual(api.spiHandle, ptr)
}

// SpiSetOnRtnQueryBankBalanceByFuture 期货发起查询银行余额通知
func (api *TraderApi) SpiSetOnRtnQueryBankBalanceByFuture(callback TraderOnRtnQueryBankBalanceByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnQueryBankBalanceByFuture(api.spiHandle, ptr)
}

// SpiSetOnErrRtnBankToFutureByFuture 期货发起银行资金转期货错误回报
func (api *TraderApi) SpiSetOnErrRtnBankToFutureByFuture(callback TraderOnErrRtnBankToFutureByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnBankToFutureByFuture(api.spiHandle, ptr)
}

// SpiSetOnErrRtnFutureToBankByFuture 期货发起期货资金转银行错误回报
func (api *TraderApi) SpiSetOnErrRtnFutureToBankByFuture(callback TraderOnErrRtnFutureToBankByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnFutureToBankByFuture(api.spiHandle, ptr)
}

// SpiSetOnErrRtnRepealBankToFutureByFutureManual 系统运行时期货端手工发起冲正银行转期货错误回报
func (api *TraderApi) SpiSetOnErrRtnRepealBankToFutureByFutureManual(callback TraderOnErrRtnRepealBankToFutureByFutureManualCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual(api.spiHandle, ptr)
}

// SpiSetOnErrRtnRepealFutureToBankByFutureManual 系统运行时期货端手工发起冲正期货转银行错误回报
func (api *TraderApi) SpiSetOnErrRtnRepealFutureToBankByFutureManual(callback TraderOnErrRtnRepealFutureToBankByFutureManualCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual(api.spiHandle, ptr)
}

// SpiSetOnErrRtnQueryBankBalanceByFuture 期货发起查询银行余额错误回报
func (api *TraderApi) SpiSetOnErrRtnQueryBankBalanceByFuture(callback TraderOnErrRtnQueryBankBalanceByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnErrRtnQueryBankBalanceByFuture(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromBankToFutureByFuture 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
func (api *TraderApi) SpiSetOnRtnRepealFromBankToFutureByFuture(callback TraderOnRtnRepealFromBankToFutureByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromBankToFutureByFuture(api.spiHandle, ptr)
}

// SpiSetOnRtnRepealFromFutureToBankByFuture 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
func (api *TraderApi) SpiSetOnRtnRepealFromFutureToBankByFuture(callback TraderOnRtnRepealFromFutureToBankByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnRepealFromFutureToBankByFuture(api.spiHandle, ptr)
}

// SpiSetOnRspFromBankToFutureByFuture 期货发起银行资金转期货应答
func (api *TraderApi) SpiSetOnRspFromBankToFutureByFuture(callback TraderOnRspFromBankToFutureByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspFromBankToFutureByFuture(api.spiHandle, ptr)
}

// SpiSetOnRspFromFutureToBankByFuture 期货发起期货资金转银行应答
func (api *TraderApi) SpiSetOnRspFromFutureToBankByFuture(callback TraderOnRspFromFutureToBankByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspFromFutureToBankByFuture(api.spiHandle, ptr)
}

// SpiSetOnRspQueryBankAccountMoneyByFuture 期货发起查询银行余额应答
func (api *TraderApi) SpiSetOnRspQueryBankAccountMoneyByFuture(callback TraderOnRspQueryBankAccountMoneyByFutureCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQueryBankAccountMoneyByFuture(api.spiHandle, ptr)
}

// SpiSetOnRtnOpenAccountByBank 银行发起银期开户通知
func (api *TraderApi) SpiSetOnRtnOpenAccountByBank(callback TraderOnRtnOpenAccountByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnOpenAccountByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnCancelAccountByBank 银行发起银期销户通知
func (api *TraderApi) SpiSetOnRtnCancelAccountByBank(callback TraderOnRtnCancelAccountByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnCancelAccountByBank(api.spiHandle, ptr)
}

// SpiSetOnRtnChangeAccountByBank 银行发起变更银行账号通知
func (api *TraderApi) SpiSetOnRtnChangeAccountByBank(callback TraderOnRtnChangeAccountByBankCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRtnChangeAccountByBank(api.spiHandle, ptr)
}

// SpiSetOnRspQryClassifiedInstrument 请求查询分类合约响应
func (api *TraderApi) SpiSetOnRspQryClassifiedInstrument(callback TraderOnRspQryClassifiedInstrumentCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryClassifiedInstrument(api.spiHandle, ptr)
}

// SpiSetOnRspQryCombPromotionParam 请求组合优惠比例响应
func (api *TraderApi) SpiSetOnRspQryCombPromotionParam(callback TraderOnRspQryCombPromotionParamCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryCombPromotionParam(api.spiHandle, ptr)
}

// SpiSetOnRspQryRiskSettleInvstPosition 投资者风险结算持仓查询响应
func (api *TraderApi) SpiSetOnRspQryRiskSettleInvstPosition(callback TraderOnRspQryRiskSettleInvstPositionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRiskSettleInvstPosition(api.spiHandle, ptr)
}

// SpiSetOnRspQryRiskSettleProductStatus 风险结算产品查询响应
func (api *TraderApi) SpiSetOnRspQryRiskSettleProductStatus(callback TraderOnRspQryRiskSettleProductStatusCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRiskSettleProductStatus(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMFutureParameter SPBM期货合约参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMFutureParameter(callback TraderOnRspQrySPBMFutureParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMFutureParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMOptionParameter SPBM期权合约参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMOptionParameter(callback TraderOnRspQrySPBMOptionParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMOptionParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMIntraParameter SPBM品种内对锁仓折扣参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMIntraParameter(callback TraderOnRspQrySPBMIntraParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMIntraParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMInterParameter SPBM跨品种抵扣参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMInterParameter(callback TraderOnRspQrySPBMInterParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMInterParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMPortfDefinition SPBM组合保证金套餐查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMPortfDefinition(callback TraderOnRspQrySPBMPortfDefinitionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMPortfDefinition(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMInvestorPortfDef 投资者SPBM套餐选择查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMInvestorPortfDef(callback TraderOnRspQrySPBMInvestorPortfDefCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMInvestorPortfDef(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorPortfMarginRatio 投资者新型组合保证金系数查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorPortfMarginRatio(callback TraderOnRspQryInvestorPortfMarginRatioCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorPortfMarginRatio(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorProdSPBMDetail 投资者产品SPBM明细查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorProdSPBMDetail(callback TraderOnRspQryInvestorProdSPBMDetailCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorProdSPBMDetail(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorCommoditySPMMMargin 投资者商品组SPMM记录查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorCommoditySPMMMargin(callback TraderOnRspQryInvestorCommoditySPMMMarginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorCommoditySPMMMargin(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorCommodityGroupSPMMMargin 投资者商品群SPMM记录查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorCommodityGroupSPMMMargin(callback TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPMMInstParam SPMM合约参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPMMInstParam(callback TraderOnRspQrySPMMInstParamCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPMMInstParam(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPMMProductParam SPMM产品参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPMMProductParam(callback TraderOnRspQrySPMMProductParamCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPMMProductParam(api.spiHandle, ptr)
}

// SpiSetOnRspQrySPBMAddOnInterParameter SPBM附加跨品种抵扣参数查询响应
func (api *TraderApi) SpiSetOnRspQrySPBMAddOnInterParameter(callback TraderOnRspQrySPBMAddOnInterParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQrySPBMAddOnInterParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSCombProductInfo RCAMS产品组合信息查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSCombProductInfo(callback TraderOnRspQryRCAMSCombProductInfoCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSCombProductInfo(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSInstrParameter RCAMS同合约风险对冲参数查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSInstrParameter(callback TraderOnRspQryRCAMSInstrParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSInstrParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSIntraParameter RCAMS品种内风险对冲参数查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSIntraParameter(callback TraderOnRspQryRCAMSIntraParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSIntraParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSInterParameter RCAMS跨品种风险折抵参数查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSInterParameter(callback TraderOnRspQryRCAMSInterParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSInterParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSShortOptAdjustParam RCAMS空头期权风险调整参数查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSShortOptAdjustParam(callback TraderOnRspQryRCAMSShortOptAdjustParamCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSShortOptAdjustParam(api.spiHandle, ptr)
}

// SpiSetOnRspQryRCAMSInvestorCombPosition RCAMS策略组合持仓查询响应
func (api *TraderApi) SpiSetOnRspQryRCAMSInvestorCombPosition(callback TraderOnRspQryRCAMSInvestorCombPositionCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRCAMSInvestorCombPosition(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorProdRCAMSMargin 投资者品种RCAMS保证金查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorProdRCAMSMargin(callback TraderOnRspQryInvestorProdRCAMSMarginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorProdRCAMSMargin(api.spiHandle, ptr)
}

// SpiSetOnRspQryRULEInstrParameter RULE合约保证金参数查询响应
func (api *TraderApi) SpiSetOnRspQryRULEInstrParameter(callback TraderOnRspQryRULEInstrParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRULEInstrParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRULEIntraParameter RULE品种内对锁仓折扣参数查询响应
func (api *TraderApi) SpiSetOnRspQryRULEIntraParameter(callback TraderOnRspQryRULEIntraParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRULEIntraParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryRULEInterParameter RULE跨品种抵扣参数查询响应
func (api *TraderApi) SpiSetOnRspQryRULEInterParameter(callback TraderOnRspQryRULEInterParameterCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryRULEInterParameter(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorProdRULEMargin 投资者产品RULE保证金查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorProdRULEMargin(callback TraderOnRspQryInvestorProdRULEMarginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorProdRULEMargin(api.spiHandle, ptr)
}

// SpiSetOnRspQryInvestorPortfSetting 投资者投资者新组保设置查询响应
func (api *TraderApi) SpiSetOnRspQryInvestorPortfSetting(callback TraderOnRspQryInvestorPortfSettingCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_TraderSpiSetOnRspQryInvestorPortfSetting(api.spiHandle, ptr)
}

// SetSpi 设置回调接口
// 此方法会创建 C SPI 实例，注册 Go 回调函数，并将 SPI 注册到 API
func (api *TraderApi) SetSpi(spi TraderSpi) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.spi = spi

	// 如果已有 C SPI 实例，先销毁
	if api.spiHandle != 0 {
		_TraderSpiDestroy(api.spiHandle)
	}

	// 创建新的 C SPI 实例
	api.spiHandle = _TraderSpiCreate(api.userData)

	// 注册所有回调函数到 C SPI
	_TraderSpiSetOnFrontConnected(api.spiHandle, GetGoTraderOnFrontConnected())
	_TraderSpiSetOnFrontDisconnected(api.spiHandle, GetGoTraderOnFrontDisconnected())
	_TraderSpiSetOnHeartBeatWarning(api.spiHandle, GetGoTraderOnHeartBeatWarning())
	_TraderSpiSetOnRspAuthenticate(api.spiHandle, GetGoTraderOnRspAuthenticate())
	_TraderSpiSetOnRspUserLogin(api.spiHandle, GetGoTraderOnRspUserLogin())
	_TraderSpiSetOnRspUserLogout(api.spiHandle, GetGoTraderOnRspUserLogout())
	_TraderSpiSetOnRspUserPasswordUpdate(api.spiHandle, GetGoTraderOnRspUserPasswordUpdate())
	_TraderSpiSetOnRspTradingAccountPasswordUpdate(api.spiHandle, GetGoTraderOnRspTradingAccountPasswordUpdate())
	_TraderSpiSetOnRspUserAuthMethod(api.spiHandle, GetGoTraderOnRspUserAuthMethod())
	_TraderSpiSetOnRspGenUserCaptcha(api.spiHandle, GetGoTraderOnRspGenUserCaptcha())
	_TraderSpiSetOnRspGenUserText(api.spiHandle, GetGoTraderOnRspGenUserText())
	_TraderSpiSetOnRspOrderInsert(api.spiHandle, GetGoTraderOnRspOrderInsert())
	_TraderSpiSetOnRspParkedOrderInsert(api.spiHandle, GetGoTraderOnRspParkedOrderInsert())
	_TraderSpiSetOnRspParkedOrderAction(api.spiHandle, GetGoTraderOnRspParkedOrderAction())
	_TraderSpiSetOnRspOrderAction(api.spiHandle, GetGoTraderOnRspOrderAction())
	_TraderSpiSetOnRspQryMaxOrderVolume(api.spiHandle, GetGoTraderOnRspQryMaxOrderVolume())
	_TraderSpiSetOnRspSettlementInfoConfirm(api.spiHandle, GetGoTraderOnRspSettlementInfoConfirm())
	_TraderSpiSetOnRspRemoveParkedOrder(api.spiHandle, GetGoTraderOnRspRemoveParkedOrder())
	_TraderSpiSetOnRspRemoveParkedOrderAction(api.spiHandle, GetGoTraderOnRspRemoveParkedOrderAction())
	_TraderSpiSetOnRspExecOrderInsert(api.spiHandle, GetGoTraderOnRspExecOrderInsert())
	_TraderSpiSetOnRspExecOrderAction(api.spiHandle, GetGoTraderOnRspExecOrderAction())
	_TraderSpiSetOnRspForQuoteInsert(api.spiHandle, GetGoTraderOnRspForQuoteInsert())
	_TraderSpiSetOnRspQuoteInsert(api.spiHandle, GetGoTraderOnRspQuoteInsert())
	_TraderSpiSetOnRspQuoteAction(api.spiHandle, GetGoTraderOnRspQuoteAction())
	_TraderSpiSetOnRspBatchOrderAction(api.spiHandle, GetGoTraderOnRspBatchOrderAction())
	_TraderSpiSetOnRspOptionSelfCloseInsert(api.spiHandle, GetGoTraderOnRspOptionSelfCloseInsert())
	_TraderSpiSetOnRspOptionSelfCloseAction(api.spiHandle, GetGoTraderOnRspOptionSelfCloseAction())
	_TraderSpiSetOnRspCombActionInsert(api.spiHandle, GetGoTraderOnRspCombActionInsert())
	_TraderSpiSetOnRspQryOrder(api.spiHandle, GetGoTraderOnRspQryOrder())
	_TraderSpiSetOnRspQryTrade(api.spiHandle, GetGoTraderOnRspQryTrade())
	_TraderSpiSetOnRspQryInvestorPosition(api.spiHandle, GetGoTraderOnRspQryInvestorPosition())
	_TraderSpiSetOnRspQryTradingAccount(api.spiHandle, GetGoTraderOnRspQryTradingAccount())
	_TraderSpiSetOnRspQryInvestor(api.spiHandle, GetGoTraderOnRspQryInvestor())
	_TraderSpiSetOnRspQryTradingCode(api.spiHandle, GetGoTraderOnRspQryTradingCode())
	_TraderSpiSetOnRspQryInstrumentMarginRate(api.spiHandle, GetGoTraderOnRspQryInstrumentMarginRate())
	_TraderSpiSetOnRspQryInstrumentCommissionRate(api.spiHandle, GetGoTraderOnRspQryInstrumentCommissionRate())
	_TraderSpiSetOnRspQryExchange(api.spiHandle, GetGoTraderOnRspQryExchange())
	_TraderSpiSetOnRspQryProduct(api.spiHandle, GetGoTraderOnRspQryProduct())
	_TraderSpiSetOnRspQryInstrument(api.spiHandle, GetGoTraderOnRspQryInstrument())
	_TraderSpiSetOnRspQryDepthMarketData(api.spiHandle, GetGoTraderOnRspQryDepthMarketData())
	_TraderSpiSetOnRspQryTraderOffer(api.spiHandle, GetGoTraderOnRspQryTraderOffer())
	_TraderSpiSetOnRspQrySettlementInfo(api.spiHandle, GetGoTraderOnRspQrySettlementInfo())
	_TraderSpiSetOnRspQryTransferBank(api.spiHandle, GetGoTraderOnRspQryTransferBank())
	_TraderSpiSetOnRspQryInvestorPositionDetail(api.spiHandle, GetGoTraderOnRspQryInvestorPositionDetail())
	_TraderSpiSetOnRspQryNotice(api.spiHandle, GetGoTraderOnRspQryNotice())
	_TraderSpiSetOnRspQrySettlementInfoConfirm(api.spiHandle, GetGoTraderOnRspQrySettlementInfoConfirm())
	_TraderSpiSetOnRspQryInvestorPositionCombineDetail(api.spiHandle, GetGoTraderOnRspQryInvestorPositionCombineDetail())
	_TraderSpiSetOnRspQryCFMMCTradingAccountKey(api.spiHandle, GetGoTraderOnRspQryCFMMCTradingAccountKey())
	_TraderSpiSetOnRspQryEWarrantOffset(api.spiHandle, GetGoTraderOnRspQryEWarrantOffset())
	_TraderSpiSetOnRspQryInvestorProductGroupMargin(api.spiHandle, GetGoTraderOnRspQryInvestorProductGroupMargin())
	_TraderSpiSetOnRspQryExchangeMarginRate(api.spiHandle, GetGoTraderOnRspQryExchangeMarginRate())
	_TraderSpiSetOnRspQryExchangeMarginRateAdjust(api.spiHandle, GetGoTraderOnRspQryExchangeMarginRateAdjust())
	_TraderSpiSetOnRspQryExchangeRate(api.spiHandle, GetGoTraderOnRspQryExchangeRate())
	_TraderSpiSetOnRspQrySecAgentACIDMap(api.spiHandle, GetGoTraderOnRspQrySecAgentACIDMap())
	_TraderSpiSetOnRspQryProductExchRate(api.spiHandle, GetGoTraderOnRspQryProductExchRate())
	_TraderSpiSetOnRspQryProductGroup(api.spiHandle, GetGoTraderOnRspQryProductGroup())
	_TraderSpiSetOnRspQryMMInstrumentCommissionRate(api.spiHandle, GetGoTraderOnRspQryMMInstrumentCommissionRate())
	_TraderSpiSetOnRspQryMMOptionInstrCommRate(api.spiHandle, GetGoTraderOnRspQryMMOptionInstrCommRate())
	_TraderSpiSetOnRspQryInstrumentOrderCommRate(api.spiHandle, GetGoTraderOnRspQryInstrumentOrderCommRate())
	_TraderSpiSetOnRspQrySecAgentTradingAccount(api.spiHandle, GetGoTraderOnRspQrySecAgentTradingAccount())
	_TraderSpiSetOnRspQrySecAgentCheckMode(api.spiHandle, GetGoTraderOnRspQrySecAgentCheckMode())
	_TraderSpiSetOnRspQrySecAgentTradeInfo(api.spiHandle, GetGoTraderOnRspQrySecAgentTradeInfo())
	_TraderSpiSetOnRspQryOptionInstrTradeCost(api.spiHandle, GetGoTraderOnRspQryOptionInstrTradeCost())
	_TraderSpiSetOnRspQryOptionInstrCommRate(api.spiHandle, GetGoTraderOnRspQryOptionInstrCommRate())
	_TraderSpiSetOnRspQryExecOrder(api.spiHandle, GetGoTraderOnRspQryExecOrder())
	_TraderSpiSetOnRspQryForQuote(api.spiHandle, GetGoTraderOnRspQryForQuote())
	_TraderSpiSetOnRspQryQuote(api.spiHandle, GetGoTraderOnRspQryQuote())
	_TraderSpiSetOnRspQryOptionSelfClose(api.spiHandle, GetGoTraderOnRspQryOptionSelfClose())
	_TraderSpiSetOnRspQryInvestUnit(api.spiHandle, GetGoTraderOnRspQryInvestUnit())
	_TraderSpiSetOnRspQryCombInstrumentGuard(api.spiHandle, GetGoTraderOnRspQryCombInstrumentGuard())
	_TraderSpiSetOnRspQryCombAction(api.spiHandle, GetGoTraderOnRspQryCombAction())
	_TraderSpiSetOnRspQryTransferSerial(api.spiHandle, GetGoTraderOnRspQryTransferSerial())
	_TraderSpiSetOnRspQryAccountregister(api.spiHandle, GetGoTraderOnRspQryAccountregister())
	_TraderSpiSetOnRspError(api.spiHandle, GetGoTraderOnRspError())
	_TraderSpiSetOnRtnOrder(api.spiHandle, GetGoTraderOnRtnOrder())
	_TraderSpiSetOnRtnTrade(api.spiHandle, GetGoTraderOnRtnTrade())
	_TraderSpiSetOnErrRtnOrderInsert(api.spiHandle, GetGoTraderOnErrRtnOrderInsert())
	_TraderSpiSetOnErrRtnOrderAction(api.spiHandle, GetGoTraderOnErrRtnOrderAction())
	_TraderSpiSetOnRtnInstrumentStatus(api.spiHandle, GetGoTraderOnRtnInstrumentStatus())
	_TraderSpiSetOnRtnBulletin(api.spiHandle, GetGoTraderOnRtnBulletin())
	_TraderSpiSetOnRtnTradingNotice(api.spiHandle, GetGoTraderOnRtnTradingNotice())
	_TraderSpiSetOnRtnErrorConditionalOrder(api.spiHandle, GetGoTraderOnRtnErrorConditionalOrder())
	_TraderSpiSetOnRtnExecOrder(api.spiHandle, GetGoTraderOnRtnExecOrder())
	_TraderSpiSetOnErrRtnExecOrderInsert(api.spiHandle, GetGoTraderOnErrRtnExecOrderInsert())
	_TraderSpiSetOnErrRtnExecOrderAction(api.spiHandle, GetGoTraderOnErrRtnExecOrderAction())
	_TraderSpiSetOnErrRtnForQuoteInsert(api.spiHandle, GetGoTraderOnErrRtnForQuoteInsert())
	_TraderSpiSetOnRtnQuote(api.spiHandle, GetGoTraderOnRtnQuote())
	_TraderSpiSetOnErrRtnQuoteInsert(api.spiHandle, GetGoTraderOnErrRtnQuoteInsert())
	_TraderSpiSetOnErrRtnQuoteAction(api.spiHandle, GetGoTraderOnErrRtnQuoteAction())
	_TraderSpiSetOnRtnForQuoteRsp(api.spiHandle, GetGoTraderOnRtnForQuoteRsp())
	_TraderSpiSetOnRtnCFMMCTradingAccountToken(api.spiHandle, GetGoTraderOnRtnCFMMCTradingAccountToken())
	_TraderSpiSetOnErrRtnBatchOrderAction(api.spiHandle, GetGoTraderOnErrRtnBatchOrderAction())
	_TraderSpiSetOnRtnOptionSelfClose(api.spiHandle, GetGoTraderOnRtnOptionSelfClose())
	_TraderSpiSetOnErrRtnOptionSelfCloseInsert(api.spiHandle, GetGoTraderOnErrRtnOptionSelfCloseInsert())
	_TraderSpiSetOnErrRtnOptionSelfCloseAction(api.spiHandle, GetGoTraderOnErrRtnOptionSelfCloseAction())
	_TraderSpiSetOnRtnCombAction(api.spiHandle, GetGoTraderOnRtnCombAction())
	_TraderSpiSetOnErrRtnCombActionInsert(api.spiHandle, GetGoTraderOnErrRtnCombActionInsert())
	_TraderSpiSetOnRspQryContractBank(api.spiHandle, GetGoTraderOnRspQryContractBank())
	_TraderSpiSetOnRspQryParkedOrder(api.spiHandle, GetGoTraderOnRspQryParkedOrder())
	_TraderSpiSetOnRspQryParkedOrderAction(api.spiHandle, GetGoTraderOnRspQryParkedOrderAction())
	_TraderSpiSetOnRspQryTradingNotice(api.spiHandle, GetGoTraderOnRspQryTradingNotice())
	_TraderSpiSetOnRspQryBrokerTradingParams(api.spiHandle, GetGoTraderOnRspQryBrokerTradingParams())
	_TraderSpiSetOnRspQryBrokerTradingAlgos(api.spiHandle, GetGoTraderOnRspQryBrokerTradingAlgos())
	_TraderSpiSetOnRspQueryCFMMCTradingAccountToken(api.spiHandle, GetGoTraderOnRspQueryCFMMCTradingAccountToken())
	_TraderSpiSetOnRtnFromBankToFutureByBank(api.spiHandle, GetGoTraderOnRtnFromBankToFutureByBank())
	_TraderSpiSetOnRtnFromFutureToBankByBank(api.spiHandle, GetGoTraderOnRtnFromFutureToBankByBank())
	_TraderSpiSetOnRtnRepealFromBankToFutureByBank(api.spiHandle, GetGoTraderOnRtnRepealFromBankToFutureByBank())
	_TraderSpiSetOnRtnRepealFromFutureToBankByBank(api.spiHandle, GetGoTraderOnRtnRepealFromFutureToBankByBank())
	_TraderSpiSetOnRtnFromBankToFutureByFuture(api.spiHandle, GetGoTraderOnRtnFromBankToFutureByFuture())
	_TraderSpiSetOnRtnFromFutureToBankByFuture(api.spiHandle, GetGoTraderOnRtnFromFutureToBankByFuture())
	_TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual(api.spiHandle, GetGoTraderOnRtnRepealFromBankToFutureByFutureManual())
	_TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual(api.spiHandle, GetGoTraderOnRtnRepealFromFutureToBankByFutureManual())
	_TraderSpiSetOnRtnQueryBankBalanceByFuture(api.spiHandle, GetGoTraderOnRtnQueryBankBalanceByFuture())
	_TraderSpiSetOnErrRtnBankToFutureByFuture(api.spiHandle, GetGoTraderOnErrRtnBankToFutureByFuture())
	_TraderSpiSetOnErrRtnFutureToBankByFuture(api.spiHandle, GetGoTraderOnErrRtnFutureToBankByFuture())
	_TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual(api.spiHandle, GetGoTraderOnErrRtnRepealBankToFutureByFutureManual())
	_TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual(api.spiHandle, GetGoTraderOnErrRtnRepealFutureToBankByFutureManual())
	_TraderSpiSetOnErrRtnQueryBankBalanceByFuture(api.spiHandle, GetGoTraderOnErrRtnQueryBankBalanceByFuture())
	_TraderSpiSetOnRtnRepealFromBankToFutureByFuture(api.spiHandle, GetGoTraderOnRtnRepealFromBankToFutureByFuture())
	_TraderSpiSetOnRtnRepealFromFutureToBankByFuture(api.spiHandle, GetGoTraderOnRtnRepealFromFutureToBankByFuture())
	_TraderSpiSetOnRspFromBankToFutureByFuture(api.spiHandle, GetGoTraderOnRspFromBankToFutureByFuture())
	_TraderSpiSetOnRspFromFutureToBankByFuture(api.spiHandle, GetGoTraderOnRspFromFutureToBankByFuture())
	_TraderSpiSetOnRspQueryBankAccountMoneyByFuture(api.spiHandle, GetGoTraderOnRspQueryBankAccountMoneyByFuture())
	_TraderSpiSetOnRtnOpenAccountByBank(api.spiHandle, GetGoTraderOnRtnOpenAccountByBank())
	_TraderSpiSetOnRtnCancelAccountByBank(api.spiHandle, GetGoTraderOnRtnCancelAccountByBank())
	_TraderSpiSetOnRtnChangeAccountByBank(api.spiHandle, GetGoTraderOnRtnChangeAccountByBank())
	_TraderSpiSetOnRspQryClassifiedInstrument(api.spiHandle, GetGoTraderOnRspQryClassifiedInstrument())
	_TraderSpiSetOnRspQryCombPromotionParam(api.spiHandle, GetGoTraderOnRspQryCombPromotionParam())
	_TraderSpiSetOnRspQryRiskSettleInvstPosition(api.spiHandle, GetGoTraderOnRspQryRiskSettleInvstPosition())
	_TraderSpiSetOnRspQryRiskSettleProductStatus(api.spiHandle, GetGoTraderOnRspQryRiskSettleProductStatus())
	_TraderSpiSetOnRspQrySPBMFutureParameter(api.spiHandle, GetGoTraderOnRspQrySPBMFutureParameter())
	_TraderSpiSetOnRspQrySPBMOptionParameter(api.spiHandle, GetGoTraderOnRspQrySPBMOptionParameter())
	_TraderSpiSetOnRspQrySPBMIntraParameter(api.spiHandle, GetGoTraderOnRspQrySPBMIntraParameter())
	_TraderSpiSetOnRspQrySPBMInterParameter(api.spiHandle, GetGoTraderOnRspQrySPBMInterParameter())
	_TraderSpiSetOnRspQrySPBMPortfDefinition(api.spiHandle, GetGoTraderOnRspQrySPBMPortfDefinition())
	_TraderSpiSetOnRspQrySPBMInvestorPortfDef(api.spiHandle, GetGoTraderOnRspQrySPBMInvestorPortfDef())
	_TraderSpiSetOnRspQryInvestorPortfMarginRatio(api.spiHandle, GetGoTraderOnRspQryInvestorPortfMarginRatio())
	_TraderSpiSetOnRspQryInvestorProdSPBMDetail(api.spiHandle, GetGoTraderOnRspQryInvestorProdSPBMDetail())
	_TraderSpiSetOnRspQryInvestorCommoditySPMMMargin(api.spiHandle, GetGoTraderOnRspQryInvestorCommoditySPMMMargin())
	_TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin(api.spiHandle, GetGoTraderOnRspQryInvestorCommodityGroupSPMMMargin())
	_TraderSpiSetOnRspQrySPMMInstParam(api.spiHandle, GetGoTraderOnRspQrySPMMInstParam())
	_TraderSpiSetOnRspQrySPMMProductParam(api.spiHandle, GetGoTraderOnRspQrySPMMProductParam())
	_TraderSpiSetOnRspQrySPBMAddOnInterParameter(api.spiHandle, GetGoTraderOnRspQrySPBMAddOnInterParameter())
	_TraderSpiSetOnRspQryRCAMSCombProductInfo(api.spiHandle, GetGoTraderOnRspQryRCAMSCombProductInfo())
	_TraderSpiSetOnRspQryRCAMSInstrParameter(api.spiHandle, GetGoTraderOnRspQryRCAMSInstrParameter())
	_TraderSpiSetOnRspQryRCAMSIntraParameter(api.spiHandle, GetGoTraderOnRspQryRCAMSIntraParameter())
	_TraderSpiSetOnRspQryRCAMSInterParameter(api.spiHandle, GetGoTraderOnRspQryRCAMSInterParameter())
	_TraderSpiSetOnRspQryRCAMSShortOptAdjustParam(api.spiHandle, GetGoTraderOnRspQryRCAMSShortOptAdjustParam())
	_TraderSpiSetOnRspQryRCAMSInvestorCombPosition(api.spiHandle, GetGoTraderOnRspQryRCAMSInvestorCombPosition())
	_TraderSpiSetOnRspQryInvestorProdRCAMSMargin(api.spiHandle, GetGoTraderOnRspQryInvestorProdRCAMSMargin())
	_TraderSpiSetOnRspQryRULEInstrParameter(api.spiHandle, GetGoTraderOnRspQryRULEInstrParameter())
	_TraderSpiSetOnRspQryRULEIntraParameter(api.spiHandle, GetGoTraderOnRspQryRULEIntraParameter())
	_TraderSpiSetOnRspQryRULEInterParameter(api.spiHandle, GetGoTraderOnRspQryRULEInterParameter())
	_TraderSpiSetOnRspQryInvestorProdRULEMargin(api.spiHandle, GetGoTraderOnRspQryInvestorProdRULEMargin())
	_TraderSpiSetOnRspQryInvestorPortfSetting(api.spiHandle, GetGoTraderOnRspQryInvestorPortfSetting())

	// 将 C SPI 注册到 API
	_TraderRegisterSpi(api.handle, api.spiHandle)
}

// ========== DataCollect 函数 ==========

// GetSystemInfo ========== DataCollect 函数 ========== 获取终端信息（AES+RSA 加密） pSystemInfo: 输出缓冲区，至少 270 字节 pLen: 输入缓冲区大小，输出实际数据长度 返回值: 0 成功，非 0 表示采集错误（按位判断）
func GetSystemInfo() ([]byte, int32) {
	// 分配至少 270 字节的缓冲区
	buf := make([]byte, 512)
	len := int32(len(buf))
	ret := _DCGetSystemInfo(&buf[0], &len)
	if ret != 0 {
		return nil, ret
	}
	return buf[:len], 0
}

// GetSystemInfoUnAesEncode 获取终端信息（未 AES 加密）
func GetSystemInfoUnAesEncode() ([]byte, int32) {
	// 分配至少 270 字节的缓冲区
	buf := make([]byte, 512)
	len := int32(len(buf))
	ret := _DCGetSystemInfoUnAesEncode(&buf[0], &len)
	if ret != 0 {
		return nil, ret
	}
	return buf[:len], 0
}

// GetDataCollectApiVersion 获取 DataCollect API 版本
func GetDataCollectApiVersion() string {
	ptr := _DCGetDataCollectApiVersion()
	if ptr == nil {
		return ""
	}
	return GoString(ptr)
}
