/**
 * CTP Trader API - 纯 C 接口封装
 * 
 * 自动生成，请勿手动修改
 * 特性：
 *   - 纯 C 接口，无 C++ 依赖
 *   - 不透明指针句柄
 *   - 回调携带 userData，支持多实例
 *   - 驼峰命名风格
 */

#ifndef CTP_TRADER_C_API_H
#define CTP_TRADER_C_API_H

#ifdef __cplusplus
extern "C" {
#endif

#include <stdbool.h>
#include <stdint.h>

// ========== 结构体定义 ==========
// CTP 的 Field 结构体是 C 兼容的 POD 类型，直接使用
// 平台相关的类型定义 (需要 THOST_TE_RESUME_TYPE)
#ifdef _WIN32
    #include "ctpapi/windows/ThostFtdcUserApiStruct.h"
#elif __APPLE__
    #include "ctpapi/macos/ThostFtdcUserApiStruct.h"
#elif __linux__
    #include "ctpapi/linux/ThostFtdcUserApiStruct.h"
#endif

// 导出宏定义
#ifdef _WIN32
    #ifdef CTP_EXPORTS
        #define CTP_API __declspec(dllexport)
    #else
        #define CTP_API __declspec(dllimport)
    #endif
#else
    #define CTP_API
#endif

// ========== 不透明句柄类型 ==========
typedef struct TraderApi_t* TraderApiHandle;
typedef struct TraderSpi_t* TraderSpiHandle;

// ========== 回调函数类型（带 userData） ==========
// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
typedef void (*TraderOnFrontConnectedCallback)(void* userData);
// 0x2003 收到错误报文
typedef void (*TraderOnFrontDisconnectedCallback)(void* userData, int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
typedef void (*TraderOnHeartBeatWarningCallback)(void* userData, int nTimeLapse);
// 客户端认证响应
typedef void (*TraderOnRspAuthenticateCallback)(void* userData, CThostFtdcRspAuthenticateField* pRspAuthenticateField, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 登录请求响应
typedef void (*TraderOnRspUserLoginCallback)(void* userData, CThostFtdcRspUserLoginField* pRspUserLogin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 登出请求响应
typedef void (*TraderOnRspUserLogoutCallback)(void* userData, CThostFtdcUserLogoutField* pUserLogout, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 用户口令更新请求响应
typedef void (*TraderOnRspUserPasswordUpdateCallback)(void* userData, CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 资金账户口令更新请求响应
typedef void (*TraderOnRspTradingAccountPasswordUpdateCallback)(void* userData, CThostFtdcTradingAccountPasswordUpdateField* pTradingAccountPasswordUpdate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 查询用户当前支持的认证模式的回复
typedef void (*TraderOnRspUserAuthMethodCallback)(void* userData, CThostFtdcRspUserAuthMethodField* pRspUserAuthMethod, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 获取图形验证码请求的回复
typedef void (*TraderOnRspGenUserCaptchaCallback)(void* userData, CThostFtdcRspGenUserCaptchaField* pRspGenUserCaptcha, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 获取短信验证码请求的回复
typedef void (*TraderOnRspGenUserTextCallback)(void* userData, CThostFtdcRspGenUserTextField* pRspGenUserText, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 报单录入请求响应
typedef void (*TraderOnRspOrderInsertCallback)(void* userData, CThostFtdcInputOrderField* pInputOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 预埋单录入请求响应
typedef void (*TraderOnRspParkedOrderInsertCallback)(void* userData, CThostFtdcParkedOrderField* pParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 预埋撤单录入请求响应
typedef void (*TraderOnRspParkedOrderActionCallback)(void* userData, CThostFtdcParkedOrderActionField* pParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 报单操作请求响应
typedef void (*TraderOnRspOrderActionCallback)(void* userData, CThostFtdcInputOrderActionField* pInputOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 查询最大报单数量响应
typedef void (*TraderOnRspQryMaxOrderVolumeCallback)(void* userData, CThostFtdcQryMaxOrderVolumeField* pQryMaxOrderVolume, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者结算结果确认响应
typedef void (*TraderOnRspSettlementInfoConfirmCallback)(void* userData, CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 删除预埋单响应
typedef void (*TraderOnRspRemoveParkedOrderCallback)(void* userData, CThostFtdcRemoveParkedOrderField* pRemoveParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 删除预埋撤单响应
typedef void (*TraderOnRspRemoveParkedOrderActionCallback)(void* userData, CThostFtdcRemoveParkedOrderActionField* pRemoveParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 执行宣告录入请求响应
typedef void (*TraderOnRspExecOrderInsertCallback)(void* userData, CThostFtdcInputExecOrderField* pInputExecOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 执行宣告操作请求响应
typedef void (*TraderOnRspExecOrderActionCallback)(void* userData, CThostFtdcInputExecOrderActionField* pInputExecOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 询价录入请求响应
typedef void (*TraderOnRspForQuoteInsertCallback)(void* userData, CThostFtdcInputForQuoteField* pInputForQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 报价录入请求响应
typedef void (*TraderOnRspQuoteInsertCallback)(void* userData, CThostFtdcInputQuoteField* pInputQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 报价操作请求响应
typedef void (*TraderOnRspQuoteActionCallback)(void* userData, CThostFtdcInputQuoteActionField* pInputQuoteAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 批量报单操作请求响应
typedef void (*TraderOnRspBatchOrderActionCallback)(void* userData, CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 期权自对冲录入请求响应
typedef void (*TraderOnRspOptionSelfCloseInsertCallback)(void* userData, CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 期权自对冲操作请求响应
typedef void (*TraderOnRspOptionSelfCloseActionCallback)(void* userData, CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 申请组合录入请求响应
typedef void (*TraderOnRspCombActionInsertCallback)(void* userData, CThostFtdcInputCombActionField* pInputCombAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询报单响应
typedef void (*TraderOnRspQryOrderCallback)(void* userData, CThostFtdcOrderField* pOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询成交响应
typedef void (*TraderOnRspQryTradeCallback)(void* userData, CThostFtdcTradeField* pTrade, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者持仓响应
typedef void (*TraderOnRspQryInvestorPositionCallback)(void* userData, CThostFtdcInvestorPositionField* pInvestorPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询资金账户响应
typedef void (*TraderOnRspQryTradingAccountCallback)(void* userData, CThostFtdcTradingAccountField* pTradingAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者响应
typedef void (*TraderOnRspQryInvestorCallback)(void* userData, CThostFtdcInvestorField* pInvestor, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易编码响应
typedef void (*TraderOnRspQryTradingCodeCallback)(void* userData, CThostFtdcTradingCodeField* pTradingCode, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询合约保证金率响应
typedef void (*TraderOnRspQryInstrumentMarginRateCallback)(void* userData, CThostFtdcInstrumentMarginRateField* pInstrumentMarginRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询合约手续费率响应
typedef void (*TraderOnRspQryInstrumentCommissionRateCallback)(void* userData, CThostFtdcInstrumentCommissionRateField* pInstrumentCommissionRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易所响应
typedef void (*TraderOnRspQryExchangeCallback)(void* userData, CThostFtdcExchangeField* pExchange, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询产品响应
typedef void (*TraderOnRspQryProductCallback)(void* userData, CThostFtdcProductField* pProduct, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询合约响应
typedef void (*TraderOnRspQryInstrumentCallback)(void* userData, CThostFtdcInstrumentField* pInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询行情响应
typedef void (*TraderOnRspQryDepthMarketDataCallback)(void* userData, CThostFtdcDepthMarketDataField* pDepthMarketData, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易员报盘机响应
typedef void (*TraderOnRspQryTraderOfferCallback)(void* userData, CThostFtdcTraderOfferField* pTraderOffer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者结算结果响应
typedef void (*TraderOnRspQrySettlementInfoCallback)(void* userData, CThostFtdcSettlementInfoField* pSettlementInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询转帐银行响应
typedef void (*TraderOnRspQryTransferBankCallback)(void* userData, CThostFtdcTransferBankField* pTransferBank, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者持仓明细响应
typedef void (*TraderOnRspQryInvestorPositionDetailCallback)(void* userData, CThostFtdcInvestorPositionDetailField* pInvestorPositionDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询客户通知响应
typedef void (*TraderOnRspQryNoticeCallback)(void* userData, CThostFtdcNoticeField* pNotice, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询结算信息确认响应
typedef void (*TraderOnRspQrySettlementInfoConfirmCallback)(void* userData, CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者持仓明细响应
typedef void (*TraderOnRspQryInvestorPositionCombineDetailCallback)(void* userData, CThostFtdcInvestorPositionCombineDetailField* pInvestorPositionCombineDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 查询保证金监管系统经纪公司资金账户密钥响应
typedef void (*TraderOnRspQryCFMMCTradingAccountKeyCallback)(void* userData, CThostFtdcCFMMCTradingAccountKeyField* pCFMMCTradingAccountKey, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询仓单折抵信息响应
typedef void (*TraderOnRspQryEWarrantOffsetCallback)(void* userData, CThostFtdcEWarrantOffsetField* pEWarrantOffset, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资者品种/跨品种保证金响应
typedef void (*TraderOnRspQryInvestorProductGroupMarginCallback)(void* userData, CThostFtdcInvestorProductGroupMarginField* pInvestorProductGroupMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易所保证金率响应
typedef void (*TraderOnRspQryExchangeMarginRateCallback)(void* userData, CThostFtdcExchangeMarginRateField* pExchangeMarginRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易所调整保证金率响应
typedef void (*TraderOnRspQryExchangeMarginRateAdjustCallback)(void* userData, CThostFtdcExchangeMarginRateAdjustField* pExchangeMarginRateAdjust, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询汇率响应
typedef void (*TraderOnRspQryExchangeRateCallback)(void* userData, CThostFtdcExchangeRateField* pExchangeRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询二级代理操作员银期权限响应
typedef void (*TraderOnRspQrySecAgentACIDMapCallback)(void* userData, CThostFtdcSecAgentACIDMapField* pSecAgentACIDMap, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询产品报价汇率
typedef void (*TraderOnRspQryProductExchRateCallback)(void* userData, CThostFtdcProductExchRateField* pProductExchRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询产品组
typedef void (*TraderOnRspQryProductGroupCallback)(void* userData, CThostFtdcProductGroupField* pProductGroup, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询做市商合约手续费率响应
typedef void (*TraderOnRspQryMMInstrumentCommissionRateCallback)(void* userData, CThostFtdcMMInstrumentCommissionRateField* pMMInstrumentCommissionRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询做市商期权合约手续费响应
typedef void (*TraderOnRspQryMMOptionInstrCommRateCallback)(void* userData, CThostFtdcMMOptionInstrCommRateField* pMMOptionInstrCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询报单手续费响应
typedef void (*TraderOnRspQryInstrumentOrderCommRateCallback)(void* userData, CThostFtdcInstrumentOrderCommRateField* pInstrumentOrderCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询资金账户响应
typedef void (*TraderOnRspQrySecAgentTradingAccountCallback)(void* userData, CThostFtdcTradingAccountField* pTradingAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询二级代理商资金校验模式响应
typedef void (*TraderOnRspQrySecAgentCheckModeCallback)(void* userData, CThostFtdcSecAgentCheckModeField* pSecAgentCheckMode, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询二级代理商信息响应
typedef void (*TraderOnRspQrySecAgentTradeInfoCallback)(void* userData, CThostFtdcSecAgentTradeInfoField* pSecAgentTradeInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询期权交易成本响应
typedef void (*TraderOnRspQryOptionInstrTradeCostCallback)(void* userData, CThostFtdcOptionInstrTradeCostField* pOptionInstrTradeCost, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询期权合约手续费响应
typedef void (*TraderOnRspQryOptionInstrCommRateCallback)(void* userData, CThostFtdcOptionInstrCommRateField* pOptionInstrCommRate, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询执行宣告响应
typedef void (*TraderOnRspQryExecOrderCallback)(void* userData, CThostFtdcExecOrderField* pExecOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询询价响应
typedef void (*TraderOnRspQryForQuoteCallback)(void* userData, CThostFtdcForQuoteField* pForQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询报价响应
typedef void (*TraderOnRspQryQuoteCallback)(void* userData, CThostFtdcQuoteField* pQuote, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询期权自对冲响应
typedef void (*TraderOnRspQryOptionSelfCloseCallback)(void* userData, CThostFtdcOptionSelfCloseField* pOptionSelfClose, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询投资单元响应
typedef void (*TraderOnRspQryInvestUnitCallback)(void* userData, CThostFtdcInvestUnitField* pInvestUnit, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询组合合约安全系数响应
typedef void (*TraderOnRspQryCombInstrumentGuardCallback)(void* userData, CThostFtdcCombInstrumentGuardField* pCombInstrumentGuard, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询申请组合响应
typedef void (*TraderOnRspQryCombActionCallback)(void* userData, CThostFtdcCombActionField* pCombAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询转帐流水响应
typedef void (*TraderOnRspQryTransferSerialCallback)(void* userData, CThostFtdcTransferSerialField* pTransferSerial, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询银期签约关系响应
typedef void (*TraderOnRspQryAccountregisterCallback)(void* userData, CThostFtdcAccountregisterField* pAccountregister, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 错误应答
typedef void (*TraderOnRspErrorCallback)(void* userData, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 报单通知
typedef void (*TraderOnRtnOrderCallback)(void* userData, CThostFtdcOrderField* pOrder);
// 成交通知
typedef void (*TraderOnRtnTradeCallback)(void* userData, CThostFtdcTradeField* pTrade);
// 报单录入错误回报
typedef void (*TraderOnErrRtnOrderInsertCallback)(void* userData, CThostFtdcInputOrderField* pInputOrder, CThostFtdcRspInfoField* pRspInfo);
// 报单操作错误回报
typedef void (*TraderOnErrRtnOrderActionCallback)(void* userData, CThostFtdcOrderActionField* pOrderAction, CThostFtdcRspInfoField* pRspInfo);
// 合约交易状态通知
typedef void (*TraderOnRtnInstrumentStatusCallback)(void* userData, CThostFtdcInstrumentStatusField* pInstrumentStatus);
// 交易所公告通知
typedef void (*TraderOnRtnBulletinCallback)(void* userData, CThostFtdcBulletinField* pBulletin);
// 交易通知
typedef void (*TraderOnRtnTradingNoticeCallback)(void* userData, CThostFtdcTradingNoticeInfoField* pTradingNoticeInfo);
// 提示条件单校验错误
typedef void (*TraderOnRtnErrorConditionalOrderCallback)(void* userData, CThostFtdcErrorConditionalOrderField* pErrorConditionalOrder);
// 执行宣告通知
typedef void (*TraderOnRtnExecOrderCallback)(void* userData, CThostFtdcExecOrderField* pExecOrder);
// 执行宣告录入错误回报
typedef void (*TraderOnErrRtnExecOrderInsertCallback)(void* userData, CThostFtdcInputExecOrderField* pInputExecOrder, CThostFtdcRspInfoField* pRspInfo);
// 执行宣告操作错误回报
typedef void (*TraderOnErrRtnExecOrderActionCallback)(void* userData, CThostFtdcExecOrderActionField* pExecOrderAction, CThostFtdcRspInfoField* pRspInfo);
// 询价录入错误回报
typedef void (*TraderOnErrRtnForQuoteInsertCallback)(void* userData, CThostFtdcInputForQuoteField* pInputForQuote, CThostFtdcRspInfoField* pRspInfo);
// 报价通知
typedef void (*TraderOnRtnQuoteCallback)(void* userData, CThostFtdcQuoteField* pQuote);
// 报价录入错误回报
typedef void (*TraderOnErrRtnQuoteInsertCallback)(void* userData, CThostFtdcInputQuoteField* pInputQuote, CThostFtdcRspInfoField* pRspInfo);
// 报价操作错误回报
typedef void (*TraderOnErrRtnQuoteActionCallback)(void* userData, CThostFtdcQuoteActionField* pQuoteAction, CThostFtdcRspInfoField* pRspInfo);
// 询价通知
typedef void (*TraderOnRtnForQuoteRspCallback)(void* userData, CThostFtdcForQuoteRspField* pForQuoteRsp);
// 保证金监控中心用户令牌
typedef void (*TraderOnRtnCFMMCTradingAccountTokenCallback)(void* userData, CThostFtdcCFMMCTradingAccountTokenField* pCFMMCTradingAccountToken);
// 批量报单操作错误回报
typedef void (*TraderOnErrRtnBatchOrderActionCallback)(void* userData, CThostFtdcBatchOrderActionField* pBatchOrderAction, CThostFtdcRspInfoField* pRspInfo);
// 期权自对冲通知
typedef void (*TraderOnRtnOptionSelfCloseCallback)(void* userData, CThostFtdcOptionSelfCloseField* pOptionSelfClose);
// 期权自对冲录入错误回报
typedef void (*TraderOnErrRtnOptionSelfCloseInsertCallback)(void* userData, CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, CThostFtdcRspInfoField* pRspInfo);
// 期权自对冲操作错误回报
typedef void (*TraderOnErrRtnOptionSelfCloseActionCallback)(void* userData, CThostFtdcOptionSelfCloseActionField* pOptionSelfCloseAction, CThostFtdcRspInfoField* pRspInfo);
// 申请组合通知
typedef void (*TraderOnRtnCombActionCallback)(void* userData, CThostFtdcCombActionField* pCombAction);
// 申请组合录入错误回报
typedef void (*TraderOnErrRtnCombActionInsertCallback)(void* userData, CThostFtdcInputCombActionField* pInputCombAction, CThostFtdcRspInfoField* pRspInfo);
// 请求查询签约银行响应
typedef void (*TraderOnRspQryContractBankCallback)(void* userData, CThostFtdcContractBankField* pContractBank, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询预埋单响应
typedef void (*TraderOnRspQryParkedOrderCallback)(void* userData, CThostFtdcParkedOrderField* pParkedOrder, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询预埋撤单响应
typedef void (*TraderOnRspQryParkedOrderActionCallback)(void* userData, CThostFtdcParkedOrderActionField* pParkedOrderAction, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询交易通知响应
typedef void (*TraderOnRspQryTradingNoticeCallback)(void* userData, CThostFtdcTradingNoticeField* pTradingNotice, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询经纪公司交易参数响应
typedef void (*TraderOnRspQryBrokerTradingParamsCallback)(void* userData, CThostFtdcBrokerTradingParamsField* pBrokerTradingParams, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询经纪公司交易算法响应
typedef void (*TraderOnRspQryBrokerTradingAlgosCallback)(void* userData, CThostFtdcBrokerTradingAlgosField* pBrokerTradingAlgos, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询监控中心用户令牌
typedef void (*TraderOnRspQueryCFMMCTradingAccountTokenCallback)(void* userData, CThostFtdcQueryCFMMCTradingAccountTokenField* pQueryCFMMCTradingAccountToken, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 银行发起银行资金转期货通知
typedef void (*TraderOnRtnFromBankToFutureByBankCallback)(void* userData, CThostFtdcRspTransferField* pRspTransfer);
// 银行发起期货资金转银行通知
typedef void (*TraderOnRtnFromFutureToBankByBankCallback)(void* userData, CThostFtdcRspTransferField* pRspTransfer);
// 银行发起冲正银行转期货通知
typedef void (*TraderOnRtnRepealFromBankToFutureByBankCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 银行发起冲正期货转银行通知
typedef void (*TraderOnRtnRepealFromFutureToBankByBankCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 期货发起银行资金转期货通知
typedef void (*TraderOnRtnFromBankToFutureByFutureCallback)(void* userData, CThostFtdcRspTransferField* pRspTransfer);
// 期货发起期货资金转银行通知
typedef void (*TraderOnRtnFromFutureToBankByFutureCallback)(void* userData, CThostFtdcRspTransferField* pRspTransfer);
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
typedef void (*TraderOnRtnRepealFromBankToFutureByFutureManualCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
typedef void (*TraderOnRtnRepealFromFutureToBankByFutureManualCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 期货发起查询银行余额通知
typedef void (*TraderOnRtnQueryBankBalanceByFutureCallback)(void* userData, CThostFtdcNotifyQueryAccountField* pNotifyQueryAccount);
// 期货发起银行资金转期货错误回报
typedef void (*TraderOnErrRtnBankToFutureByFutureCallback)(void* userData, CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo);
// 期货发起期货资金转银行错误回报
typedef void (*TraderOnErrRtnFutureToBankByFutureCallback)(void* userData, CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo);
// 系统运行时期货端手工发起冲正银行转期货错误回报
typedef void (*TraderOnErrRtnRepealBankToFutureByFutureManualCallback)(void* userData, CThostFtdcReqRepealField* pReqRepeal, CThostFtdcRspInfoField* pRspInfo);
// 系统运行时期货端手工发起冲正期货转银行错误回报
typedef void (*TraderOnErrRtnRepealFutureToBankByFutureManualCallback)(void* userData, CThostFtdcReqRepealField* pReqRepeal, CThostFtdcRspInfoField* pRspInfo);
// 期货发起查询银行余额错误回报
typedef void (*TraderOnErrRtnQueryBankBalanceByFutureCallback)(void* userData, CThostFtdcReqQueryAccountField* pReqQueryAccount, CThostFtdcRspInfoField* pRspInfo);
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
typedef void (*TraderOnRtnRepealFromBankToFutureByFutureCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
typedef void (*TraderOnRtnRepealFromFutureToBankByFutureCallback)(void* userData, CThostFtdcRspRepealField* pRspRepeal);
// 期货发起银行资金转期货应答
typedef void (*TraderOnRspFromBankToFutureByFutureCallback)(void* userData, CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 期货发起期货资金转银行应答
typedef void (*TraderOnRspFromFutureToBankByFutureCallback)(void* userData, CThostFtdcReqTransferField* pReqTransfer, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 期货发起查询银行余额应答
typedef void (*TraderOnRspQueryBankAccountMoneyByFutureCallback)(void* userData, CThostFtdcReqQueryAccountField* pReqQueryAccount, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 银行发起银期开户通知
typedef void (*TraderOnRtnOpenAccountByBankCallback)(void* userData, CThostFtdcOpenAccountField* pOpenAccount);
// 银行发起银期销户通知
typedef void (*TraderOnRtnCancelAccountByBankCallback)(void* userData, CThostFtdcCancelAccountField* pCancelAccount);
// 银行发起变更银行账号通知
typedef void (*TraderOnRtnChangeAccountByBankCallback)(void* userData, CThostFtdcChangeAccountField* pChangeAccount);
// 请求查询分类合约响应
typedef void (*TraderOnRspQryClassifiedInstrumentCallback)(void* userData, CThostFtdcInstrumentField* pInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求组合优惠比例响应
typedef void (*TraderOnRspQryCombPromotionParamCallback)(void* userData, CThostFtdcCombPromotionParamField* pCombPromotionParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者风险结算持仓查询响应
typedef void (*TraderOnRspQryRiskSettleInvstPositionCallback)(void* userData, CThostFtdcRiskSettleInvstPositionField* pRiskSettleInvstPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 风险结算产品查询响应
typedef void (*TraderOnRspQryRiskSettleProductStatusCallback)(void* userData, CThostFtdcRiskSettleProductStatusField* pRiskSettleProductStatus, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM期货合约参数查询响应
typedef void (*TraderOnRspQrySPBMFutureParameterCallback)(void* userData, CThostFtdcSPBMFutureParameterField* pSPBMFutureParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM期权合约参数查询响应
typedef void (*TraderOnRspQrySPBMOptionParameterCallback)(void* userData, CThostFtdcSPBMOptionParameterField* pSPBMOptionParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM品种内对锁仓折扣参数查询响应
typedef void (*TraderOnRspQrySPBMIntraParameterCallback)(void* userData, CThostFtdcSPBMIntraParameterField* pSPBMIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM跨品种抵扣参数查询响应
typedef void (*TraderOnRspQrySPBMInterParameterCallback)(void* userData, CThostFtdcSPBMInterParameterField* pSPBMInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM组合保证金套餐查询响应
typedef void (*TraderOnRspQrySPBMPortfDefinitionCallback)(void* userData, CThostFtdcSPBMPortfDefinitionField* pSPBMPortfDefinition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者SPBM套餐选择查询响应
typedef void (*TraderOnRspQrySPBMInvestorPortfDefCallback)(void* userData, CThostFtdcSPBMInvestorPortfDefField* pSPBMInvestorPortfDef, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者新型组合保证金系数查询响应
typedef void (*TraderOnRspQryInvestorPortfMarginRatioCallback)(void* userData, CThostFtdcInvestorPortfMarginRatioField* pInvestorPortfMarginRatio, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者产品SPBM明细查询响应
typedef void (*TraderOnRspQryInvestorProdSPBMDetailCallback)(void* userData, CThostFtdcInvestorProdSPBMDetailField* pInvestorProdSPBMDetail, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者商品组SPMM记录查询响应
typedef void (*TraderOnRspQryInvestorCommoditySPMMMarginCallback)(void* userData, CThostFtdcInvestorCommoditySPMMMarginField* pInvestorCommoditySPMMMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者商品群SPMM记录查询响应
typedef void (*TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback)(void* userData, CThostFtdcInvestorCommodityGroupSPMMMarginField* pInvestorCommodityGroupSPMMMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPMM合约参数查询响应
typedef void (*TraderOnRspQrySPMMInstParamCallback)(void* userData, CThostFtdcSPMMInstParamField* pSPMMInstParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPMM产品参数查询响应
typedef void (*TraderOnRspQrySPMMProductParamCallback)(void* userData, CThostFtdcSPMMProductParamField* pSPMMProductParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// SPBM附加跨品种抵扣参数查询响应
typedef void (*TraderOnRspQrySPBMAddOnInterParameterCallback)(void* userData, CThostFtdcSPBMAddOnInterParameterField* pSPBMAddOnInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS产品组合信息查询响应
typedef void (*TraderOnRspQryRCAMSCombProductInfoCallback)(void* userData, CThostFtdcRCAMSCombProductInfoField* pRCAMSCombProductInfo, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS同合约风险对冲参数查询响应
typedef void (*TraderOnRspQryRCAMSInstrParameterCallback)(void* userData, CThostFtdcRCAMSInstrParameterField* pRCAMSInstrParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS品种内风险对冲参数查询响应
typedef void (*TraderOnRspQryRCAMSIntraParameterCallback)(void* userData, CThostFtdcRCAMSIntraParameterField* pRCAMSIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS跨品种风险折抵参数查询响应
typedef void (*TraderOnRspQryRCAMSInterParameterCallback)(void* userData, CThostFtdcRCAMSInterParameterField* pRCAMSInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS空头期权风险调整参数查询响应
typedef void (*TraderOnRspQryRCAMSShortOptAdjustParamCallback)(void* userData, CThostFtdcRCAMSShortOptAdjustParamField* pRCAMSShortOptAdjustParam, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RCAMS策略组合持仓查询响应
typedef void (*TraderOnRspQryRCAMSInvestorCombPositionCallback)(void* userData, CThostFtdcRCAMSInvestorCombPositionField* pRCAMSInvestorCombPosition, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者品种RCAMS保证金查询响应
typedef void (*TraderOnRspQryInvestorProdRCAMSMarginCallback)(void* userData, CThostFtdcInvestorProdRCAMSMarginField* pInvestorProdRCAMSMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RULE合约保证金参数查询响应
typedef void (*TraderOnRspQryRULEInstrParameterCallback)(void* userData, CThostFtdcRULEInstrParameterField* pRULEInstrParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RULE品种内对锁仓折扣参数查询响应
typedef void (*TraderOnRspQryRULEIntraParameterCallback)(void* userData, CThostFtdcRULEIntraParameterField* pRULEIntraParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// RULE跨品种抵扣参数查询响应
typedef void (*TraderOnRspQryRULEInterParameterCallback)(void* userData, CThostFtdcRULEInterParameterField* pRULEInterParameter, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者产品RULE保证金查询响应
typedef void (*TraderOnRspQryInvestorProdRULEMarginCallback)(void* userData, CThostFtdcInvestorProdRULEMarginField* pInvestorProdRULEMargin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 投资者投资者新组保设置查询响应
typedef void (*TraderOnRspQryInvestorPortfSettingCallback)(void* userData, CThostFtdcInvestorPortfSettingField* pInvestorPortfSetting, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);

// ========== 回调表结构（便于批量设置） ==========
typedef struct {
    void* userData;
    TraderOnFrontConnectedCallback onFrontConnected;
    TraderOnFrontDisconnectedCallback onFrontDisconnected;
    TraderOnHeartBeatWarningCallback onHeartBeatWarning;
    TraderOnRspAuthenticateCallback onRspAuthenticate;
    TraderOnRspUserLoginCallback onRspUserLogin;
    TraderOnRspUserLogoutCallback onRspUserLogout;
    TraderOnRspUserPasswordUpdateCallback onRspUserPasswordUpdate;
    TraderOnRspTradingAccountPasswordUpdateCallback onRspTradingAccountPasswordUpdate;
    TraderOnRspUserAuthMethodCallback onRspUserAuthMethod;
    TraderOnRspGenUserCaptchaCallback onRspGenUserCaptcha;
    TraderOnRspGenUserTextCallback onRspGenUserText;
    TraderOnRspOrderInsertCallback onRspOrderInsert;
    TraderOnRspParkedOrderInsertCallback onRspParkedOrderInsert;
    TraderOnRspParkedOrderActionCallback onRspParkedOrderAction;
    TraderOnRspOrderActionCallback onRspOrderAction;
    TraderOnRspQryMaxOrderVolumeCallback onRspQryMaxOrderVolume;
    TraderOnRspSettlementInfoConfirmCallback onRspSettlementInfoConfirm;
    TraderOnRspRemoveParkedOrderCallback onRspRemoveParkedOrder;
    TraderOnRspRemoveParkedOrderActionCallback onRspRemoveParkedOrderAction;
    TraderOnRspExecOrderInsertCallback onRspExecOrderInsert;
    TraderOnRspExecOrderActionCallback onRspExecOrderAction;
    TraderOnRspForQuoteInsertCallback onRspForQuoteInsert;
    TraderOnRspQuoteInsertCallback onRspQuoteInsert;
    TraderOnRspQuoteActionCallback onRspQuoteAction;
    TraderOnRspBatchOrderActionCallback onRspBatchOrderAction;
    TraderOnRspOptionSelfCloseInsertCallback onRspOptionSelfCloseInsert;
    TraderOnRspOptionSelfCloseActionCallback onRspOptionSelfCloseAction;
    TraderOnRspCombActionInsertCallback onRspCombActionInsert;
    TraderOnRspQryOrderCallback onRspQryOrder;
    TraderOnRspQryTradeCallback onRspQryTrade;
    TraderOnRspQryInvestorPositionCallback onRspQryInvestorPosition;
    TraderOnRspQryTradingAccountCallback onRspQryTradingAccount;
    TraderOnRspQryInvestorCallback onRspQryInvestor;
    TraderOnRspQryTradingCodeCallback onRspQryTradingCode;
    TraderOnRspQryInstrumentMarginRateCallback onRspQryInstrumentMarginRate;
    TraderOnRspQryInstrumentCommissionRateCallback onRspQryInstrumentCommissionRate;
    TraderOnRspQryExchangeCallback onRspQryExchange;
    TraderOnRspQryProductCallback onRspQryProduct;
    TraderOnRspQryInstrumentCallback onRspQryInstrument;
    TraderOnRspQryDepthMarketDataCallback onRspQryDepthMarketData;
    TraderOnRspQryTraderOfferCallback onRspQryTraderOffer;
    TraderOnRspQrySettlementInfoCallback onRspQrySettlementInfo;
    TraderOnRspQryTransferBankCallback onRspQryTransferBank;
    TraderOnRspQryInvestorPositionDetailCallback onRspQryInvestorPositionDetail;
    TraderOnRspQryNoticeCallback onRspQryNotice;
    TraderOnRspQrySettlementInfoConfirmCallback onRspQrySettlementInfoConfirm;
    TraderOnRspQryInvestorPositionCombineDetailCallback onRspQryInvestorPositionCombineDetail;
    TraderOnRspQryCFMMCTradingAccountKeyCallback onRspQryCFMMCTradingAccountKey;
    TraderOnRspQryEWarrantOffsetCallback onRspQryEWarrantOffset;
    TraderOnRspQryInvestorProductGroupMarginCallback onRspQryInvestorProductGroupMargin;
    TraderOnRspQryExchangeMarginRateCallback onRspQryExchangeMarginRate;
    TraderOnRspQryExchangeMarginRateAdjustCallback onRspQryExchangeMarginRateAdjust;
    TraderOnRspQryExchangeRateCallback onRspQryExchangeRate;
    TraderOnRspQrySecAgentACIDMapCallback onRspQrySecAgentACIDMap;
    TraderOnRspQryProductExchRateCallback onRspQryProductExchRate;
    TraderOnRspQryProductGroupCallback onRspQryProductGroup;
    TraderOnRspQryMMInstrumentCommissionRateCallback onRspQryMMInstrumentCommissionRate;
    TraderOnRspQryMMOptionInstrCommRateCallback onRspQryMMOptionInstrCommRate;
    TraderOnRspQryInstrumentOrderCommRateCallback onRspQryInstrumentOrderCommRate;
    TraderOnRspQrySecAgentTradingAccountCallback onRspQrySecAgentTradingAccount;
    TraderOnRspQrySecAgentCheckModeCallback onRspQrySecAgentCheckMode;
    TraderOnRspQrySecAgentTradeInfoCallback onRspQrySecAgentTradeInfo;
    TraderOnRspQryOptionInstrTradeCostCallback onRspQryOptionInstrTradeCost;
    TraderOnRspQryOptionInstrCommRateCallback onRspQryOptionInstrCommRate;
    TraderOnRspQryExecOrderCallback onRspQryExecOrder;
    TraderOnRspQryForQuoteCallback onRspQryForQuote;
    TraderOnRspQryQuoteCallback onRspQryQuote;
    TraderOnRspQryOptionSelfCloseCallback onRspQryOptionSelfClose;
    TraderOnRspQryInvestUnitCallback onRspQryInvestUnit;
    TraderOnRspQryCombInstrumentGuardCallback onRspQryCombInstrumentGuard;
    TraderOnRspQryCombActionCallback onRspQryCombAction;
    TraderOnRspQryTransferSerialCallback onRspQryTransferSerial;
    TraderOnRspQryAccountregisterCallback onRspQryAccountregister;
    TraderOnRspErrorCallback onRspError;
    TraderOnRtnOrderCallback onRtnOrder;
    TraderOnRtnTradeCallback onRtnTrade;
    TraderOnErrRtnOrderInsertCallback onErrRtnOrderInsert;
    TraderOnErrRtnOrderActionCallback onErrRtnOrderAction;
    TraderOnRtnInstrumentStatusCallback onRtnInstrumentStatus;
    TraderOnRtnBulletinCallback onRtnBulletin;
    TraderOnRtnTradingNoticeCallback onRtnTradingNotice;
    TraderOnRtnErrorConditionalOrderCallback onRtnErrorConditionalOrder;
    TraderOnRtnExecOrderCallback onRtnExecOrder;
    TraderOnErrRtnExecOrderInsertCallback onErrRtnExecOrderInsert;
    TraderOnErrRtnExecOrderActionCallback onErrRtnExecOrderAction;
    TraderOnErrRtnForQuoteInsertCallback onErrRtnForQuoteInsert;
    TraderOnRtnQuoteCallback onRtnQuote;
    TraderOnErrRtnQuoteInsertCallback onErrRtnQuoteInsert;
    TraderOnErrRtnQuoteActionCallback onErrRtnQuoteAction;
    TraderOnRtnForQuoteRspCallback onRtnForQuoteRsp;
    TraderOnRtnCFMMCTradingAccountTokenCallback onRtnCFMMCTradingAccountToken;
    TraderOnErrRtnBatchOrderActionCallback onErrRtnBatchOrderAction;
    TraderOnRtnOptionSelfCloseCallback onRtnOptionSelfClose;
    TraderOnErrRtnOptionSelfCloseInsertCallback onErrRtnOptionSelfCloseInsert;
    TraderOnErrRtnOptionSelfCloseActionCallback onErrRtnOptionSelfCloseAction;
    TraderOnRtnCombActionCallback onRtnCombAction;
    TraderOnErrRtnCombActionInsertCallback onErrRtnCombActionInsert;
    TraderOnRspQryContractBankCallback onRspQryContractBank;
    TraderOnRspQryParkedOrderCallback onRspQryParkedOrder;
    TraderOnRspQryParkedOrderActionCallback onRspQryParkedOrderAction;
    TraderOnRspQryTradingNoticeCallback onRspQryTradingNotice;
    TraderOnRspQryBrokerTradingParamsCallback onRspQryBrokerTradingParams;
    TraderOnRspQryBrokerTradingAlgosCallback onRspQryBrokerTradingAlgos;
    TraderOnRspQueryCFMMCTradingAccountTokenCallback onRspQueryCFMMCTradingAccountToken;
    TraderOnRtnFromBankToFutureByBankCallback onRtnFromBankToFutureByBank;
    TraderOnRtnFromFutureToBankByBankCallback onRtnFromFutureToBankByBank;
    TraderOnRtnRepealFromBankToFutureByBankCallback onRtnRepealFromBankToFutureByBank;
    TraderOnRtnRepealFromFutureToBankByBankCallback onRtnRepealFromFutureToBankByBank;
    TraderOnRtnFromBankToFutureByFutureCallback onRtnFromBankToFutureByFuture;
    TraderOnRtnFromFutureToBankByFutureCallback onRtnFromFutureToBankByFuture;
    TraderOnRtnRepealFromBankToFutureByFutureManualCallback onRtnRepealFromBankToFutureByFutureManual;
    TraderOnRtnRepealFromFutureToBankByFutureManualCallback onRtnRepealFromFutureToBankByFutureManual;
    TraderOnRtnQueryBankBalanceByFutureCallback onRtnQueryBankBalanceByFuture;
    TraderOnErrRtnBankToFutureByFutureCallback onErrRtnBankToFutureByFuture;
    TraderOnErrRtnFutureToBankByFutureCallback onErrRtnFutureToBankByFuture;
    TraderOnErrRtnRepealBankToFutureByFutureManualCallback onErrRtnRepealBankToFutureByFutureManual;
    TraderOnErrRtnRepealFutureToBankByFutureManualCallback onErrRtnRepealFutureToBankByFutureManual;
    TraderOnErrRtnQueryBankBalanceByFutureCallback onErrRtnQueryBankBalanceByFuture;
    TraderOnRtnRepealFromBankToFutureByFutureCallback onRtnRepealFromBankToFutureByFuture;
    TraderOnRtnRepealFromFutureToBankByFutureCallback onRtnRepealFromFutureToBankByFuture;
    TraderOnRspFromBankToFutureByFutureCallback onRspFromBankToFutureByFuture;
    TraderOnRspFromFutureToBankByFutureCallback onRspFromFutureToBankByFuture;
    TraderOnRspQueryBankAccountMoneyByFutureCallback onRspQueryBankAccountMoneyByFuture;
    TraderOnRtnOpenAccountByBankCallback onRtnOpenAccountByBank;
    TraderOnRtnCancelAccountByBankCallback onRtnCancelAccountByBank;
    TraderOnRtnChangeAccountByBankCallback onRtnChangeAccountByBank;
    TraderOnRspQryClassifiedInstrumentCallback onRspQryClassifiedInstrument;
    TraderOnRspQryCombPromotionParamCallback onRspQryCombPromotionParam;
    TraderOnRspQryRiskSettleInvstPositionCallback onRspQryRiskSettleInvstPosition;
    TraderOnRspQryRiskSettleProductStatusCallback onRspQryRiskSettleProductStatus;
    TraderOnRspQrySPBMFutureParameterCallback onRspQrySPBMFutureParameter;
    TraderOnRspQrySPBMOptionParameterCallback onRspQrySPBMOptionParameter;
    TraderOnRspQrySPBMIntraParameterCallback onRspQrySPBMIntraParameter;
    TraderOnRspQrySPBMInterParameterCallback onRspQrySPBMInterParameter;
    TraderOnRspQrySPBMPortfDefinitionCallback onRspQrySPBMPortfDefinition;
    TraderOnRspQrySPBMInvestorPortfDefCallback onRspQrySPBMInvestorPortfDef;
    TraderOnRspQryInvestorPortfMarginRatioCallback onRspQryInvestorPortfMarginRatio;
    TraderOnRspQryInvestorProdSPBMDetailCallback onRspQryInvestorProdSPBMDetail;
    TraderOnRspQryInvestorCommoditySPMMMarginCallback onRspQryInvestorCommoditySPMMMargin;
    TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback onRspQryInvestorCommodityGroupSPMMMargin;
    TraderOnRspQrySPMMInstParamCallback onRspQrySPMMInstParam;
    TraderOnRspQrySPMMProductParamCallback onRspQrySPMMProductParam;
    TraderOnRspQrySPBMAddOnInterParameterCallback onRspQrySPBMAddOnInterParameter;
    TraderOnRspQryRCAMSCombProductInfoCallback onRspQryRCAMSCombProductInfo;
    TraderOnRspQryRCAMSInstrParameterCallback onRspQryRCAMSInstrParameter;
    TraderOnRspQryRCAMSIntraParameterCallback onRspQryRCAMSIntraParameter;
    TraderOnRspQryRCAMSInterParameterCallback onRspQryRCAMSInterParameter;
    TraderOnRspQryRCAMSShortOptAdjustParamCallback onRspQryRCAMSShortOptAdjustParam;
    TraderOnRspQryRCAMSInvestorCombPositionCallback onRspQryRCAMSInvestorCombPosition;
    TraderOnRspQryInvestorProdRCAMSMarginCallback onRspQryInvestorProdRCAMSMargin;
    TraderOnRspQryRULEInstrParameterCallback onRspQryRULEInstrParameter;
    TraderOnRspQryRULEIntraParameterCallback onRspQryRULEIntraParameter;
    TraderOnRspQryRULEInterParameterCallback onRspQryRULEInterParameter;
    TraderOnRspQryInvestorProdRULEMarginCallback onRspQryInvestorProdRULEMargin;
    TraderOnRspQryInvestorPortfSettingCallback onRspQryInvestorPortfSetting;
} TraderSpiCallbacks;

// ========== Trader API 函数 ==========

// 创建TraderApi
CTP_API TraderApiHandle TraderCreateFtdcTraderApi(const char* pszFlowPath);

// 获取API的版本信息
CTP_API const char * TraderGetApiVersion(void);


// 删除接口对象本身
CTP_API void TraderRelease(TraderApiHandle handle);

// 初始化
CTP_API void TraderInit(TraderApiHandle handle);

// 等待接口线程结束运行
CTP_API int TraderJoin(TraderApiHandle handle);

// 获取当前交易日
CTP_API const char * TraderGetTradingDay(TraderApiHandle handle);

// 获取已连接的前置的信息
CTP_API void TraderGetFrontInfo(TraderApiHandle handle, CThostFtdcFrontInfoField* pFrontInfo);

// 注册前置机网络地址
CTP_API void TraderRegisterFront(TraderApiHandle handle, char* pszFrontAddress);

// 注册名字服务器网络地址
CTP_API void TraderRegisterNameServer(TraderApiHandle handle, char* pszNsAddress);

// 注册名字服务器用户信息
CTP_API void TraderRegisterFensUserInfo(TraderApiHandle handle, CThostFtdcFensUserInfoField* pFensUserInfo);

// 订阅私有流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后私有流的内容
CTP_API void TraderSubscribePrivateTopic(TraderApiHandle handle, THOST_TE_RESUME_TYPE nResumeType);

// 订阅公共流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后公共流的内容 THOST_TERT_NONE:取消订阅公共流
CTP_API void TraderSubscribePublicTopic(TraderApiHandle handle, THOST_TE_RESUME_TYPE nResumeType);

// 客户端认证请求
CTP_API int TraderReqAuthenticate(TraderApiHandle handle, CThostFtdcReqAuthenticateField* pReqAuthenticateField, int nRequestID);

// 注册用户终端信息，用于中继服务器多连接模式 需要在终端认证成功后，用户登录前调用该接口
CTP_API int TraderRegisterUserSystemInfo(TraderApiHandle handle, CThostFtdcUserSystemInfoField* pUserSystemInfo);

// 上报用户终端信息，用于中继服务器操作员登录模式 操作员登录后，可以多次调用该接口上报客户信息
CTP_API int TraderSubmitUserSystemInfo(TraderApiHandle handle, CThostFtdcUserSystemInfoField* pUserSystemInfo);

// 用户登录请求
CTP_API int TraderReqUserLogin(TraderApiHandle handle, CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

// 登出请求
CTP_API int TraderReqUserLogout(TraderApiHandle handle, CThostFtdcUserLogoutField* pUserLogout, int nRequestID);

// 用户口令更新请求
CTP_API int TraderReqUserPasswordUpdate(TraderApiHandle handle, CThostFtdcUserPasswordUpdateField* pUserPasswordUpdate, int nRequestID);

// 资金账户口令更新请求
CTP_API int TraderReqTradingAccountPasswordUpdate(TraderApiHandle handle, CThostFtdcTradingAccountPasswordUpdateField* pTradingAccountPasswordUpdate, int nRequestID);

// 查询用户当前支持的认证模式
CTP_API int TraderReqUserAuthMethod(TraderApiHandle handle, CThostFtdcReqUserAuthMethodField* pReqUserAuthMethod, int nRequestID);

// 用户发出获取图形验证码请求
CTP_API int TraderReqGenUserCaptcha(TraderApiHandle handle, CThostFtdcReqGenUserCaptchaField* pReqGenUserCaptcha, int nRequestID);

// 用户发出获取短信验证码请求
CTP_API int TraderReqGenUserText(TraderApiHandle handle, CThostFtdcReqGenUserTextField* pReqGenUserText, int nRequestID);

// 用户发出带有图片验证码的登陆请求
CTP_API int TraderReqUserLoginWithCaptcha(TraderApiHandle handle, CThostFtdcReqUserLoginWithCaptchaField* pReqUserLoginWithCaptcha, int nRequestID);

// 用户发出带有短信验证码的登陆请求
CTP_API int TraderReqUserLoginWithText(TraderApiHandle handle, CThostFtdcReqUserLoginWithTextField* pReqUserLoginWithText, int nRequestID);

// 用户发出带有动态口令的登陆请求
CTP_API int TraderReqUserLoginWithOTP(TraderApiHandle handle, CThostFtdcReqUserLoginWithOTPField* pReqUserLoginWithOTP, int nRequestID);

// 报单录入请求
CTP_API int TraderReqOrderInsert(TraderApiHandle handle, CThostFtdcInputOrderField* pInputOrder, int nRequestID);

// 预埋单录入请求
CTP_API int TraderReqParkedOrderInsert(TraderApiHandle handle, CThostFtdcParkedOrderField* pParkedOrder, int nRequestID);

// 预埋撤单录入请求
CTP_API int TraderReqParkedOrderAction(TraderApiHandle handle, CThostFtdcParkedOrderActionField* pParkedOrderAction, int nRequestID);

// 报单操作请求
CTP_API int TraderReqOrderAction(TraderApiHandle handle, CThostFtdcInputOrderActionField* pInputOrderAction, int nRequestID);

// 查询最大报单数量请求
CTP_API int TraderReqQryMaxOrderVolume(TraderApiHandle handle, CThostFtdcQryMaxOrderVolumeField* pQryMaxOrderVolume, int nRequestID);

// 投资者结算结果确认
CTP_API int TraderReqSettlementInfoConfirm(TraderApiHandle handle, CThostFtdcSettlementInfoConfirmField* pSettlementInfoConfirm, int nRequestID);

// 请求删除预埋单
CTP_API int TraderReqRemoveParkedOrder(TraderApiHandle handle, CThostFtdcRemoveParkedOrderField* pRemoveParkedOrder, int nRequestID);

// 请求删除预埋撤单
CTP_API int TraderReqRemoveParkedOrderAction(TraderApiHandle handle, CThostFtdcRemoveParkedOrderActionField* pRemoveParkedOrderAction, int nRequestID);

// 执行宣告录入请求
CTP_API int TraderReqExecOrderInsert(TraderApiHandle handle, CThostFtdcInputExecOrderField* pInputExecOrder, int nRequestID);

// 执行宣告操作请求
CTP_API int TraderReqExecOrderAction(TraderApiHandle handle, CThostFtdcInputExecOrderActionField* pInputExecOrderAction, int nRequestID);

// 询价录入请求
CTP_API int TraderReqForQuoteInsert(TraderApiHandle handle, CThostFtdcInputForQuoteField* pInputForQuote, int nRequestID);

// 报价录入请求
CTP_API int TraderReqQuoteInsert(TraderApiHandle handle, CThostFtdcInputQuoteField* pInputQuote, int nRequestID);

// 报价操作请求
CTP_API int TraderReqQuoteAction(TraderApiHandle handle, CThostFtdcInputQuoteActionField* pInputQuoteAction, int nRequestID);

// 批量报单操作请求
CTP_API int TraderReqBatchOrderAction(TraderApiHandle handle, CThostFtdcInputBatchOrderActionField* pInputBatchOrderAction, int nRequestID);

// 期权自对冲录入请求
CTP_API int TraderReqOptionSelfCloseInsert(TraderApiHandle handle, CThostFtdcInputOptionSelfCloseField* pInputOptionSelfClose, int nRequestID);

// 期权自对冲操作请求
CTP_API int TraderReqOptionSelfCloseAction(TraderApiHandle handle, CThostFtdcInputOptionSelfCloseActionField* pInputOptionSelfCloseAction, int nRequestID);

// 申请组合录入请求
CTP_API int TraderReqCombActionInsert(TraderApiHandle handle, CThostFtdcInputCombActionField* pInputCombAction, int nRequestID);

// 请求查询报单
CTP_API int TraderReqQryOrder(TraderApiHandle handle, CThostFtdcQryOrderField* pQryOrder, int nRequestID);

// 请求查询成交
CTP_API int TraderReqQryTrade(TraderApiHandle handle, CThostFtdcQryTradeField* pQryTrade, int nRequestID);

// 请求查询投资者持仓
CTP_API int TraderReqQryInvestorPosition(TraderApiHandle handle, CThostFtdcQryInvestorPositionField* pQryInvestorPosition, int nRequestID);

// 请求查询资金账户
CTP_API int TraderReqQryTradingAccount(TraderApiHandle handle, CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID);

// 请求查询投资者
CTP_API int TraderReqQryInvestor(TraderApiHandle handle, CThostFtdcQryInvestorField* pQryInvestor, int nRequestID);

// 请求查询交易编码
CTP_API int TraderReqQryTradingCode(TraderApiHandle handle, CThostFtdcQryTradingCodeField* pQryTradingCode, int nRequestID);

// 请求查询合约保证金率
CTP_API int TraderReqQryInstrumentMarginRate(TraderApiHandle handle, CThostFtdcQryInstrumentMarginRateField* pQryInstrumentMarginRate, int nRequestID);

// 请求查询合约手续费率
CTP_API int TraderReqQryInstrumentCommissionRate(TraderApiHandle handle, CThostFtdcQryInstrumentCommissionRateField* pQryInstrumentCommissionRate, int nRequestID);

// 请求查询交易所
CTP_API int TraderReqQryExchange(TraderApiHandle handle, CThostFtdcQryExchangeField* pQryExchange, int nRequestID);

// 请求查询产品
CTP_API int TraderReqQryProduct(TraderApiHandle handle, CThostFtdcQryProductField* pQryProduct, int nRequestID);

// 请求查询合约
CTP_API int TraderReqQryInstrument(TraderApiHandle handle, CThostFtdcQryInstrumentField* pQryInstrument, int nRequestID);

// 请求查询行情
CTP_API int TraderReqQryDepthMarketData(TraderApiHandle handle, CThostFtdcQryDepthMarketDataField* pQryDepthMarketData, int nRequestID);

// 请求查询交易员报盘机
CTP_API int TraderReqQryTraderOffer(TraderApiHandle handle, CThostFtdcQryTraderOfferField* pQryTraderOffer, int nRequestID);

// 请求查询投资者结算结果
CTP_API int TraderReqQrySettlementInfo(TraderApiHandle handle, CThostFtdcQrySettlementInfoField* pQrySettlementInfo, int nRequestID);

// 请求查询转帐银行
CTP_API int TraderReqQryTransferBank(TraderApiHandle handle, CThostFtdcQryTransferBankField* pQryTransferBank, int nRequestID);

// 请求查询投资者持仓明细
CTP_API int TraderReqQryInvestorPositionDetail(TraderApiHandle handle, CThostFtdcQryInvestorPositionDetailField* pQryInvestorPositionDetail, int nRequestID);

// 请求查询客户通知
CTP_API int TraderReqQryNotice(TraderApiHandle handle, CThostFtdcQryNoticeField* pQryNotice, int nRequestID);

// 请求查询结算信息确认
CTP_API int TraderReqQrySettlementInfoConfirm(TraderApiHandle handle, CThostFtdcQrySettlementInfoConfirmField* pQrySettlementInfoConfirm, int nRequestID);

// 请求查询投资者持仓明细
CTP_API int TraderReqQryInvestorPositionCombineDetail(TraderApiHandle handle, CThostFtdcQryInvestorPositionCombineDetailField* pQryInvestorPositionCombineDetail, int nRequestID);

// 请求查询保证金监管系统经纪公司资金账户密钥
CTP_API int TraderReqQryCFMMCTradingAccountKey(TraderApiHandle handle, CThostFtdcQryCFMMCTradingAccountKeyField* pQryCFMMCTradingAccountKey, int nRequestID);

// 请求查询仓单折抵信息
CTP_API int TraderReqQryEWarrantOffset(TraderApiHandle handle, CThostFtdcQryEWarrantOffsetField* pQryEWarrantOffset, int nRequestID);

// 请求查询投资者品种/跨品种保证金
CTP_API int TraderReqQryInvestorProductGroupMargin(TraderApiHandle handle, CThostFtdcQryInvestorProductGroupMarginField* pQryInvestorProductGroupMargin, int nRequestID);

// 请求查询交易所保证金率
CTP_API int TraderReqQryExchangeMarginRate(TraderApiHandle handle, CThostFtdcQryExchangeMarginRateField* pQryExchangeMarginRate, int nRequestID);

// 请求查询交易所调整保证金率
CTP_API int TraderReqQryExchangeMarginRateAdjust(TraderApiHandle handle, CThostFtdcQryExchangeMarginRateAdjustField* pQryExchangeMarginRateAdjust, int nRequestID);

// 请求查询汇率
CTP_API int TraderReqQryExchangeRate(TraderApiHandle handle, CThostFtdcQryExchangeRateField* pQryExchangeRate, int nRequestID);

// 请求查询二级代理操作员银期权限
CTP_API int TraderReqQrySecAgentACIDMap(TraderApiHandle handle, CThostFtdcQrySecAgentACIDMapField* pQrySecAgentACIDMap, int nRequestID);

// 请求查询产品报价汇率
CTP_API int TraderReqQryProductExchRate(TraderApiHandle handle, CThostFtdcQryProductExchRateField* pQryProductExchRate, int nRequestID);

// 请求查询产品组
CTP_API int TraderReqQryProductGroup(TraderApiHandle handle, CThostFtdcQryProductGroupField* pQryProductGroup, int nRequestID);

// 请求查询做市商合约手续费率
CTP_API int TraderReqQryMMInstrumentCommissionRate(TraderApiHandle handle, CThostFtdcQryMMInstrumentCommissionRateField* pQryMMInstrumentCommissionRate, int nRequestID);

// 请求查询做市商期权合约手续费
CTP_API int TraderReqQryMMOptionInstrCommRate(TraderApiHandle handle, CThostFtdcQryMMOptionInstrCommRateField* pQryMMOptionInstrCommRate, int nRequestID);

// 请求查询报单手续费
CTP_API int TraderReqQryInstrumentOrderCommRate(TraderApiHandle handle, CThostFtdcQryInstrumentOrderCommRateField* pQryInstrumentOrderCommRate, int nRequestID);

// 请求查询资金账户
CTP_API int TraderReqQrySecAgentTradingAccount(TraderApiHandle handle, CThostFtdcQryTradingAccountField* pQryTradingAccount, int nRequestID);

// 请求查询二级代理商资金校验模式
CTP_API int TraderReqQrySecAgentCheckMode(TraderApiHandle handle, CThostFtdcQrySecAgentCheckModeField* pQrySecAgentCheckMode, int nRequestID);

// 请求查询二级代理商信息
CTP_API int TraderReqQrySecAgentTradeInfo(TraderApiHandle handle, CThostFtdcQrySecAgentTradeInfoField* pQrySecAgentTradeInfo, int nRequestID);

// 请求查询期权交易成本
CTP_API int TraderReqQryOptionInstrTradeCost(TraderApiHandle handle, CThostFtdcQryOptionInstrTradeCostField* pQryOptionInstrTradeCost, int nRequestID);

// 请求查询期权合约手续费
CTP_API int TraderReqQryOptionInstrCommRate(TraderApiHandle handle, CThostFtdcQryOptionInstrCommRateField* pQryOptionInstrCommRate, int nRequestID);

// 请求查询执行宣告
CTP_API int TraderReqQryExecOrder(TraderApiHandle handle, CThostFtdcQryExecOrderField* pQryExecOrder, int nRequestID);

// 请求查询询价
CTP_API int TraderReqQryForQuote(TraderApiHandle handle, CThostFtdcQryForQuoteField* pQryForQuote, int nRequestID);

// 请求查询报价
CTP_API int TraderReqQryQuote(TraderApiHandle handle, CThostFtdcQryQuoteField* pQryQuote, int nRequestID);

// 请求查询期权自对冲
CTP_API int TraderReqQryOptionSelfClose(TraderApiHandle handle, CThostFtdcQryOptionSelfCloseField* pQryOptionSelfClose, int nRequestID);

// 请求查询投资单元
CTP_API int TraderReqQryInvestUnit(TraderApiHandle handle, CThostFtdcQryInvestUnitField* pQryInvestUnit, int nRequestID);

// 请求查询组合合约安全系数
CTP_API int TraderReqQryCombInstrumentGuard(TraderApiHandle handle, CThostFtdcQryCombInstrumentGuardField* pQryCombInstrumentGuard, int nRequestID);

// 请求查询申请组合
CTP_API int TraderReqQryCombAction(TraderApiHandle handle, CThostFtdcQryCombActionField* pQryCombAction, int nRequestID);

// 请求查询转帐流水
CTP_API int TraderReqQryTransferSerial(TraderApiHandle handle, CThostFtdcQryTransferSerialField* pQryTransferSerial, int nRequestID);

// 请求查询银期签约关系
CTP_API int TraderReqQryAccountregister(TraderApiHandle handle, CThostFtdcQryAccountregisterField* pQryAccountregister, int nRequestID);

// 请求查询签约银行
CTP_API int TraderReqQryContractBank(TraderApiHandle handle, CThostFtdcQryContractBankField* pQryContractBank, int nRequestID);

// 请求查询预埋单
CTP_API int TraderReqQryParkedOrder(TraderApiHandle handle, CThostFtdcQryParkedOrderField* pQryParkedOrder, int nRequestID);

// 请求查询预埋撤单
CTP_API int TraderReqQryParkedOrderAction(TraderApiHandle handle, CThostFtdcQryParkedOrderActionField* pQryParkedOrderAction, int nRequestID);

// 请求查询交易通知
CTP_API int TraderReqQryTradingNotice(TraderApiHandle handle, CThostFtdcQryTradingNoticeField* pQryTradingNotice, int nRequestID);

// 请求查询经纪公司交易参数
CTP_API int TraderReqQryBrokerTradingParams(TraderApiHandle handle, CThostFtdcQryBrokerTradingParamsField* pQryBrokerTradingParams, int nRequestID);

// 请求查询经纪公司交易算法
CTP_API int TraderReqQryBrokerTradingAlgos(TraderApiHandle handle, CThostFtdcQryBrokerTradingAlgosField* pQryBrokerTradingAlgos, int nRequestID);

// 请求查询监控中心用户令牌
CTP_API int TraderReqQueryCFMMCTradingAccountToken(TraderApiHandle handle, CThostFtdcQueryCFMMCTradingAccountTokenField* pQueryCFMMCTradingAccountToken, int nRequestID);

// 期货发起银行资金转期货请求
CTP_API int TraderReqFromBankToFutureByFuture(TraderApiHandle handle, CThostFtdcReqTransferField* pReqTransfer, int nRequestID);

// 期货发起期货资金转银行请求
CTP_API int TraderReqFromFutureToBankByFuture(TraderApiHandle handle, CThostFtdcReqTransferField* pReqTransfer, int nRequestID);

// 期货发起查询银行余额请求
CTP_API int TraderReqQueryBankAccountMoneyByFuture(TraderApiHandle handle, CThostFtdcReqQueryAccountField* pReqQueryAccount, int nRequestID);

// 请求查询分类合约
CTP_API int TraderReqQryClassifiedInstrument(TraderApiHandle handle, CThostFtdcQryClassifiedInstrumentField* pQryClassifiedInstrument, int nRequestID);

// 请求组合优惠比例
CTP_API int TraderReqQryCombPromotionParam(TraderApiHandle handle, CThostFtdcQryCombPromotionParamField* pQryCombPromotionParam, int nRequestID);

// 投资者风险结算持仓查询
CTP_API int TraderReqQryRiskSettleInvstPosition(TraderApiHandle handle, CThostFtdcQryRiskSettleInvstPositionField* pQryRiskSettleInvstPosition, int nRequestID);

// 风险结算产品查询
CTP_API int TraderReqQryRiskSettleProductStatus(TraderApiHandle handle, CThostFtdcQryRiskSettleProductStatusField* pQryRiskSettleProductStatus, int nRequestID);

// SPBM期货合约参数查询
CTP_API int TraderReqQrySPBMFutureParameter(TraderApiHandle handle, CThostFtdcQrySPBMFutureParameterField* pQrySPBMFutureParameter, int nRequestID);

// SPBM期权合约参数查询
CTP_API int TraderReqQrySPBMOptionParameter(TraderApiHandle handle, CThostFtdcQrySPBMOptionParameterField* pQrySPBMOptionParameter, int nRequestID);

// SPBM品种内对锁仓折扣参数查询
CTP_API int TraderReqQrySPBMIntraParameter(TraderApiHandle handle, CThostFtdcQrySPBMIntraParameterField* pQrySPBMIntraParameter, int nRequestID);

// SPBM跨品种抵扣参数查询
CTP_API int TraderReqQrySPBMInterParameter(TraderApiHandle handle, CThostFtdcQrySPBMInterParameterField* pQrySPBMInterParameter, int nRequestID);

// SPBM组合保证金套餐查询
CTP_API int TraderReqQrySPBMPortfDefinition(TraderApiHandle handle, CThostFtdcQrySPBMPortfDefinitionField* pQrySPBMPortfDefinition, int nRequestID);

// 投资者SPBM套餐选择查询
CTP_API int TraderReqQrySPBMInvestorPortfDef(TraderApiHandle handle, CThostFtdcQrySPBMInvestorPortfDefField* pQrySPBMInvestorPortfDef, int nRequestID);

// 投资者新型组合保证金系数查询
CTP_API int TraderReqQryInvestorPortfMarginRatio(TraderApiHandle handle, CThostFtdcQryInvestorPortfMarginRatioField* pQryInvestorPortfMarginRatio, int nRequestID);

// 投资者产品SPBM明细查询
CTP_API int TraderReqQryInvestorProdSPBMDetail(TraderApiHandle handle, CThostFtdcQryInvestorProdSPBMDetailField* pQryInvestorProdSPBMDetail, int nRequestID);

// 投资者商品组SPMM记录查询
CTP_API int TraderReqQryInvestorCommoditySPMMMargin(TraderApiHandle handle, CThostFtdcQryInvestorCommoditySPMMMarginField* pQryInvestorCommoditySPMMMargin, int nRequestID);

// 投资者商品群SPMM记录查询
CTP_API int TraderReqQryInvestorCommodityGroupSPMMMargin(TraderApiHandle handle, CThostFtdcQryInvestorCommodityGroupSPMMMarginField* pQryInvestorCommodityGroupSPMMMargin, int nRequestID);

// SPMM合约参数查询
CTP_API int TraderReqQrySPMMInstParam(TraderApiHandle handle, CThostFtdcQrySPMMInstParamField* pQrySPMMInstParam, int nRequestID);

// SPMM产品参数查询
CTP_API int TraderReqQrySPMMProductParam(TraderApiHandle handle, CThostFtdcQrySPMMProductParamField* pQrySPMMProductParam, int nRequestID);

// SPBM附加跨品种抵扣参数查询
CTP_API int TraderReqQrySPBMAddOnInterParameter(TraderApiHandle handle, CThostFtdcQrySPBMAddOnInterParameterField* pQrySPBMAddOnInterParameter, int nRequestID);

// RCAMS产品组合信息查询
CTP_API int TraderReqQryRCAMSCombProductInfo(TraderApiHandle handle, CThostFtdcQryRCAMSCombProductInfoField* pQryRCAMSCombProductInfo, int nRequestID);

// RCAMS同合约风险对冲参数查询
CTP_API int TraderReqQryRCAMSInstrParameter(TraderApiHandle handle, CThostFtdcQryRCAMSInstrParameterField* pQryRCAMSInstrParameter, int nRequestID);

// RCAMS品种内风险对冲参数查询
CTP_API int TraderReqQryRCAMSIntraParameter(TraderApiHandle handle, CThostFtdcQryRCAMSIntraParameterField* pQryRCAMSIntraParameter, int nRequestID);

// RCAMS跨品种风险折抵参数查询
CTP_API int TraderReqQryRCAMSInterParameter(TraderApiHandle handle, CThostFtdcQryRCAMSInterParameterField* pQryRCAMSInterParameter, int nRequestID);

// RCAMS空头期权风险调整参数查询
CTP_API int TraderReqQryRCAMSShortOptAdjustParam(TraderApiHandle handle, CThostFtdcQryRCAMSShortOptAdjustParamField* pQryRCAMSShortOptAdjustParam, int nRequestID);

// RCAMS策略组合持仓查询
CTP_API int TraderReqQryRCAMSInvestorCombPosition(TraderApiHandle handle, CThostFtdcQryRCAMSInvestorCombPositionField* pQryRCAMSInvestorCombPosition, int nRequestID);

// 投资者品种RCAMS保证金查询
CTP_API int TraderReqQryInvestorProdRCAMSMargin(TraderApiHandle handle, CThostFtdcQryInvestorProdRCAMSMarginField* pQryInvestorProdRCAMSMargin, int nRequestID);

// RULE合约保证金参数查询
CTP_API int TraderReqQryRULEInstrParameter(TraderApiHandle handle, CThostFtdcQryRULEInstrParameterField* pQryRULEInstrParameter, int nRequestID);

// RULE品种内对锁仓折扣参数查询
CTP_API int TraderReqQryRULEIntraParameter(TraderApiHandle handle, CThostFtdcQryRULEIntraParameterField* pQryRULEIntraParameter, int nRequestID);

// RULE跨品种抵扣参数查询
CTP_API int TraderReqQryRULEInterParameter(TraderApiHandle handle, CThostFtdcQryRULEInterParameterField* pQryRULEInterParameter, int nRequestID);

// 投资者产品RULE保证金查询
CTP_API int TraderReqQryInvestorProdRULEMargin(TraderApiHandle handle, CThostFtdcQryInvestorProdRULEMarginField* pQryInvestorProdRULEMargin, int nRequestID);

// 投资者投资者新组保设置查询
CTP_API int TraderReqQryInvestorPortfSetting(TraderApiHandle handle, CThostFtdcQryInvestorPortfSettingField* pQryInvestorPortfSetting, int nRequestID);


// ========== Trader SPI 函数 ==========

// 创建 SPI 实例
CTP_API TraderSpiHandle TraderSpiCreate(void* userData);

// 销毁 SPI 实例
CTP_API void TraderSpiDestroy(TraderSpiHandle spi);

// 注册 SPI 到 API
CTP_API void TraderRegisterSpi(TraderApiHandle api, TraderSpiHandle spi);

// 批量设置回调
CTP_API void TraderSpiSetCallbacks(TraderSpiHandle spi, const TraderSpiCallbacks* callbacks);

// 单独设置回调
// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
CTP_API void TraderSpiSetOnFrontConnected(TraderSpiHandle spi, TraderOnFrontConnectedCallback callback);
// 0x2003 收到错误报文
CTP_API void TraderSpiSetOnFrontDisconnected(TraderSpiHandle spi, TraderOnFrontDisconnectedCallback callback);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
CTP_API void TraderSpiSetOnHeartBeatWarning(TraderSpiHandle spi, TraderOnHeartBeatWarningCallback callback);
// 客户端认证响应
CTP_API void TraderSpiSetOnRspAuthenticate(TraderSpiHandle spi, TraderOnRspAuthenticateCallback callback);
// 登录请求响应
CTP_API void TraderSpiSetOnRspUserLogin(TraderSpiHandle spi, TraderOnRspUserLoginCallback callback);
// 登出请求响应
CTP_API void TraderSpiSetOnRspUserLogout(TraderSpiHandle spi, TraderOnRspUserLogoutCallback callback);
// 用户口令更新请求响应
CTP_API void TraderSpiSetOnRspUserPasswordUpdate(TraderSpiHandle spi, TraderOnRspUserPasswordUpdateCallback callback);
// 资金账户口令更新请求响应
CTP_API void TraderSpiSetOnRspTradingAccountPasswordUpdate(TraderSpiHandle spi, TraderOnRspTradingAccountPasswordUpdateCallback callback);
// 查询用户当前支持的认证模式的回复
CTP_API void TraderSpiSetOnRspUserAuthMethod(TraderSpiHandle spi, TraderOnRspUserAuthMethodCallback callback);
// 获取图形验证码请求的回复
CTP_API void TraderSpiSetOnRspGenUserCaptcha(TraderSpiHandle spi, TraderOnRspGenUserCaptchaCallback callback);
// 获取短信验证码请求的回复
CTP_API void TraderSpiSetOnRspGenUserText(TraderSpiHandle spi, TraderOnRspGenUserTextCallback callback);
// 报单录入请求响应
CTP_API void TraderSpiSetOnRspOrderInsert(TraderSpiHandle spi, TraderOnRspOrderInsertCallback callback);
// 预埋单录入请求响应
CTP_API void TraderSpiSetOnRspParkedOrderInsert(TraderSpiHandle spi, TraderOnRspParkedOrderInsertCallback callback);
// 预埋撤单录入请求响应
CTP_API void TraderSpiSetOnRspParkedOrderAction(TraderSpiHandle spi, TraderOnRspParkedOrderActionCallback callback);
// 报单操作请求响应
CTP_API void TraderSpiSetOnRspOrderAction(TraderSpiHandle spi, TraderOnRspOrderActionCallback callback);
// 查询最大报单数量响应
CTP_API void TraderSpiSetOnRspQryMaxOrderVolume(TraderSpiHandle spi, TraderOnRspQryMaxOrderVolumeCallback callback);
// 投资者结算结果确认响应
CTP_API void TraderSpiSetOnRspSettlementInfoConfirm(TraderSpiHandle spi, TraderOnRspSettlementInfoConfirmCallback callback);
// 删除预埋单响应
CTP_API void TraderSpiSetOnRspRemoveParkedOrder(TraderSpiHandle spi, TraderOnRspRemoveParkedOrderCallback callback);
// 删除预埋撤单响应
CTP_API void TraderSpiSetOnRspRemoveParkedOrderAction(TraderSpiHandle spi, TraderOnRspRemoveParkedOrderActionCallback callback);
// 执行宣告录入请求响应
CTP_API void TraderSpiSetOnRspExecOrderInsert(TraderSpiHandle spi, TraderOnRspExecOrderInsertCallback callback);
// 执行宣告操作请求响应
CTP_API void TraderSpiSetOnRspExecOrderAction(TraderSpiHandle spi, TraderOnRspExecOrderActionCallback callback);
// 询价录入请求响应
CTP_API void TraderSpiSetOnRspForQuoteInsert(TraderSpiHandle spi, TraderOnRspForQuoteInsertCallback callback);
// 报价录入请求响应
CTP_API void TraderSpiSetOnRspQuoteInsert(TraderSpiHandle spi, TraderOnRspQuoteInsertCallback callback);
// 报价操作请求响应
CTP_API void TraderSpiSetOnRspQuoteAction(TraderSpiHandle spi, TraderOnRspQuoteActionCallback callback);
// 批量报单操作请求响应
CTP_API void TraderSpiSetOnRspBatchOrderAction(TraderSpiHandle spi, TraderOnRspBatchOrderActionCallback callback);
// 期权自对冲录入请求响应
CTP_API void TraderSpiSetOnRspOptionSelfCloseInsert(TraderSpiHandle spi, TraderOnRspOptionSelfCloseInsertCallback callback);
// 期权自对冲操作请求响应
CTP_API void TraderSpiSetOnRspOptionSelfCloseAction(TraderSpiHandle spi, TraderOnRspOptionSelfCloseActionCallback callback);
// 申请组合录入请求响应
CTP_API void TraderSpiSetOnRspCombActionInsert(TraderSpiHandle spi, TraderOnRspCombActionInsertCallback callback);
// 请求查询报单响应
CTP_API void TraderSpiSetOnRspQryOrder(TraderSpiHandle spi, TraderOnRspQryOrderCallback callback);
// 请求查询成交响应
CTP_API void TraderSpiSetOnRspQryTrade(TraderSpiHandle spi, TraderOnRspQryTradeCallback callback);
// 请求查询投资者持仓响应
CTP_API void TraderSpiSetOnRspQryInvestorPosition(TraderSpiHandle spi, TraderOnRspQryInvestorPositionCallback callback);
// 请求查询资金账户响应
CTP_API void TraderSpiSetOnRspQryTradingAccount(TraderSpiHandle spi, TraderOnRspQryTradingAccountCallback callback);
// 请求查询投资者响应
CTP_API void TraderSpiSetOnRspQryInvestor(TraderSpiHandle spi, TraderOnRspQryInvestorCallback callback);
// 请求查询交易编码响应
CTP_API void TraderSpiSetOnRspQryTradingCode(TraderSpiHandle spi, TraderOnRspQryTradingCodeCallback callback);
// 请求查询合约保证金率响应
CTP_API void TraderSpiSetOnRspQryInstrumentMarginRate(TraderSpiHandle spi, TraderOnRspQryInstrumentMarginRateCallback callback);
// 请求查询合约手续费率响应
CTP_API void TraderSpiSetOnRspQryInstrumentCommissionRate(TraderSpiHandle spi, TraderOnRspQryInstrumentCommissionRateCallback callback);
// 请求查询交易所响应
CTP_API void TraderSpiSetOnRspQryExchange(TraderSpiHandle spi, TraderOnRspQryExchangeCallback callback);
// 请求查询产品响应
CTP_API void TraderSpiSetOnRspQryProduct(TraderSpiHandle spi, TraderOnRspQryProductCallback callback);
// 请求查询合约响应
CTP_API void TraderSpiSetOnRspQryInstrument(TraderSpiHandle spi, TraderOnRspQryInstrumentCallback callback);
// 请求查询行情响应
CTP_API void TraderSpiSetOnRspQryDepthMarketData(TraderSpiHandle spi, TraderOnRspQryDepthMarketDataCallback callback);
// 请求查询交易员报盘机响应
CTP_API void TraderSpiSetOnRspQryTraderOffer(TraderSpiHandle spi, TraderOnRspQryTraderOfferCallback callback);
// 请求查询投资者结算结果响应
CTP_API void TraderSpiSetOnRspQrySettlementInfo(TraderSpiHandle spi, TraderOnRspQrySettlementInfoCallback callback);
// 请求查询转帐银行响应
CTP_API void TraderSpiSetOnRspQryTransferBank(TraderSpiHandle spi, TraderOnRspQryTransferBankCallback callback);
// 请求查询投资者持仓明细响应
CTP_API void TraderSpiSetOnRspQryInvestorPositionDetail(TraderSpiHandle spi, TraderOnRspQryInvestorPositionDetailCallback callback);
// 请求查询客户通知响应
CTP_API void TraderSpiSetOnRspQryNotice(TraderSpiHandle spi, TraderOnRspQryNoticeCallback callback);
// 请求查询结算信息确认响应
CTP_API void TraderSpiSetOnRspQrySettlementInfoConfirm(TraderSpiHandle spi, TraderOnRspQrySettlementInfoConfirmCallback callback);
// 请求查询投资者持仓明细响应
CTP_API void TraderSpiSetOnRspQryInvestorPositionCombineDetail(TraderSpiHandle spi, TraderOnRspQryInvestorPositionCombineDetailCallback callback);
// 查询保证金监管系统经纪公司资金账户密钥响应
CTP_API void TraderSpiSetOnRspQryCFMMCTradingAccountKey(TraderSpiHandle spi, TraderOnRspQryCFMMCTradingAccountKeyCallback callback);
// 请求查询仓单折抵信息响应
CTP_API void TraderSpiSetOnRspQryEWarrantOffset(TraderSpiHandle spi, TraderOnRspQryEWarrantOffsetCallback callback);
// 请求查询投资者品种/跨品种保证金响应
CTP_API void TraderSpiSetOnRspQryInvestorProductGroupMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProductGroupMarginCallback callback);
// 请求查询交易所保证金率响应
CTP_API void TraderSpiSetOnRspQryExchangeMarginRate(TraderSpiHandle spi, TraderOnRspQryExchangeMarginRateCallback callback);
// 请求查询交易所调整保证金率响应
CTP_API void TraderSpiSetOnRspQryExchangeMarginRateAdjust(TraderSpiHandle spi, TraderOnRspQryExchangeMarginRateAdjustCallback callback);
// 请求查询汇率响应
CTP_API void TraderSpiSetOnRspQryExchangeRate(TraderSpiHandle spi, TraderOnRspQryExchangeRateCallback callback);
// 请求查询二级代理操作员银期权限响应
CTP_API void TraderSpiSetOnRspQrySecAgentACIDMap(TraderSpiHandle spi, TraderOnRspQrySecAgentACIDMapCallback callback);
// 请求查询产品报价汇率
CTP_API void TraderSpiSetOnRspQryProductExchRate(TraderSpiHandle spi, TraderOnRspQryProductExchRateCallback callback);
// 请求查询产品组
CTP_API void TraderSpiSetOnRspQryProductGroup(TraderSpiHandle spi, TraderOnRspQryProductGroupCallback callback);
// 请求查询做市商合约手续费率响应
CTP_API void TraderSpiSetOnRspQryMMInstrumentCommissionRate(TraderSpiHandle spi, TraderOnRspQryMMInstrumentCommissionRateCallback callback);
// 请求查询做市商期权合约手续费响应
CTP_API void TraderSpiSetOnRspQryMMOptionInstrCommRate(TraderSpiHandle spi, TraderOnRspQryMMOptionInstrCommRateCallback callback);
// 请求查询报单手续费响应
CTP_API void TraderSpiSetOnRspQryInstrumentOrderCommRate(TraderSpiHandle spi, TraderOnRspQryInstrumentOrderCommRateCallback callback);
// 请求查询资金账户响应
CTP_API void TraderSpiSetOnRspQrySecAgentTradingAccount(TraderSpiHandle spi, TraderOnRspQrySecAgentTradingAccountCallback callback);
// 请求查询二级代理商资金校验模式响应
CTP_API void TraderSpiSetOnRspQrySecAgentCheckMode(TraderSpiHandle spi, TraderOnRspQrySecAgentCheckModeCallback callback);
// 请求查询二级代理商信息响应
CTP_API void TraderSpiSetOnRspQrySecAgentTradeInfo(TraderSpiHandle spi, TraderOnRspQrySecAgentTradeInfoCallback callback);
// 请求查询期权交易成本响应
CTP_API void TraderSpiSetOnRspQryOptionInstrTradeCost(TraderSpiHandle spi, TraderOnRspQryOptionInstrTradeCostCallback callback);
// 请求查询期权合约手续费响应
CTP_API void TraderSpiSetOnRspQryOptionInstrCommRate(TraderSpiHandle spi, TraderOnRspQryOptionInstrCommRateCallback callback);
// 请求查询执行宣告响应
CTP_API void TraderSpiSetOnRspQryExecOrder(TraderSpiHandle spi, TraderOnRspQryExecOrderCallback callback);
// 请求查询询价响应
CTP_API void TraderSpiSetOnRspQryForQuote(TraderSpiHandle spi, TraderOnRspQryForQuoteCallback callback);
// 请求查询报价响应
CTP_API void TraderSpiSetOnRspQryQuote(TraderSpiHandle spi, TraderOnRspQryQuoteCallback callback);
// 请求查询期权自对冲响应
CTP_API void TraderSpiSetOnRspQryOptionSelfClose(TraderSpiHandle spi, TraderOnRspQryOptionSelfCloseCallback callback);
// 请求查询投资单元响应
CTP_API void TraderSpiSetOnRspQryInvestUnit(TraderSpiHandle spi, TraderOnRspQryInvestUnitCallback callback);
// 请求查询组合合约安全系数响应
CTP_API void TraderSpiSetOnRspQryCombInstrumentGuard(TraderSpiHandle spi, TraderOnRspQryCombInstrumentGuardCallback callback);
// 请求查询申请组合响应
CTP_API void TraderSpiSetOnRspQryCombAction(TraderSpiHandle spi, TraderOnRspQryCombActionCallback callback);
// 请求查询转帐流水响应
CTP_API void TraderSpiSetOnRspQryTransferSerial(TraderSpiHandle spi, TraderOnRspQryTransferSerialCallback callback);
// 请求查询银期签约关系响应
CTP_API void TraderSpiSetOnRspQryAccountregister(TraderSpiHandle spi, TraderOnRspQryAccountregisterCallback callback);
// 错误应答
CTP_API void TraderSpiSetOnRspError(TraderSpiHandle spi, TraderOnRspErrorCallback callback);
// 报单通知
CTP_API void TraderSpiSetOnRtnOrder(TraderSpiHandle spi, TraderOnRtnOrderCallback callback);
// 成交通知
CTP_API void TraderSpiSetOnRtnTrade(TraderSpiHandle spi, TraderOnRtnTradeCallback callback);
// 报单录入错误回报
CTP_API void TraderSpiSetOnErrRtnOrderInsert(TraderSpiHandle spi, TraderOnErrRtnOrderInsertCallback callback);
// 报单操作错误回报
CTP_API void TraderSpiSetOnErrRtnOrderAction(TraderSpiHandle spi, TraderOnErrRtnOrderActionCallback callback);
// 合约交易状态通知
CTP_API void TraderSpiSetOnRtnInstrumentStatus(TraderSpiHandle spi, TraderOnRtnInstrumentStatusCallback callback);
// 交易所公告通知
CTP_API void TraderSpiSetOnRtnBulletin(TraderSpiHandle spi, TraderOnRtnBulletinCallback callback);
// 交易通知
CTP_API void TraderSpiSetOnRtnTradingNotice(TraderSpiHandle spi, TraderOnRtnTradingNoticeCallback callback);
// 提示条件单校验错误
CTP_API void TraderSpiSetOnRtnErrorConditionalOrder(TraderSpiHandle spi, TraderOnRtnErrorConditionalOrderCallback callback);
// 执行宣告通知
CTP_API void TraderSpiSetOnRtnExecOrder(TraderSpiHandle spi, TraderOnRtnExecOrderCallback callback);
// 执行宣告录入错误回报
CTP_API void TraderSpiSetOnErrRtnExecOrderInsert(TraderSpiHandle spi, TraderOnErrRtnExecOrderInsertCallback callback);
// 执行宣告操作错误回报
CTP_API void TraderSpiSetOnErrRtnExecOrderAction(TraderSpiHandle spi, TraderOnErrRtnExecOrderActionCallback callback);
// 询价录入错误回报
CTP_API void TraderSpiSetOnErrRtnForQuoteInsert(TraderSpiHandle spi, TraderOnErrRtnForQuoteInsertCallback callback);
// 报价通知
CTP_API void TraderSpiSetOnRtnQuote(TraderSpiHandle spi, TraderOnRtnQuoteCallback callback);
// 报价录入错误回报
CTP_API void TraderSpiSetOnErrRtnQuoteInsert(TraderSpiHandle spi, TraderOnErrRtnQuoteInsertCallback callback);
// 报价操作错误回报
CTP_API void TraderSpiSetOnErrRtnQuoteAction(TraderSpiHandle spi, TraderOnErrRtnQuoteActionCallback callback);
// 询价通知
CTP_API void TraderSpiSetOnRtnForQuoteRsp(TraderSpiHandle spi, TraderOnRtnForQuoteRspCallback callback);
// 保证金监控中心用户令牌
CTP_API void TraderSpiSetOnRtnCFMMCTradingAccountToken(TraderSpiHandle spi, TraderOnRtnCFMMCTradingAccountTokenCallback callback);
// 批量报单操作错误回报
CTP_API void TraderSpiSetOnErrRtnBatchOrderAction(TraderSpiHandle spi, TraderOnErrRtnBatchOrderActionCallback callback);
// 期权自对冲通知
CTP_API void TraderSpiSetOnRtnOptionSelfClose(TraderSpiHandle spi, TraderOnRtnOptionSelfCloseCallback callback);
// 期权自对冲录入错误回报
CTP_API void TraderSpiSetOnErrRtnOptionSelfCloseInsert(TraderSpiHandle spi, TraderOnErrRtnOptionSelfCloseInsertCallback callback);
// 期权自对冲操作错误回报
CTP_API void TraderSpiSetOnErrRtnOptionSelfCloseAction(TraderSpiHandle spi, TraderOnErrRtnOptionSelfCloseActionCallback callback);
// 申请组合通知
CTP_API void TraderSpiSetOnRtnCombAction(TraderSpiHandle spi, TraderOnRtnCombActionCallback callback);
// 申请组合录入错误回报
CTP_API void TraderSpiSetOnErrRtnCombActionInsert(TraderSpiHandle spi, TraderOnErrRtnCombActionInsertCallback callback);
// 请求查询签约银行响应
CTP_API void TraderSpiSetOnRspQryContractBank(TraderSpiHandle spi, TraderOnRspQryContractBankCallback callback);
// 请求查询预埋单响应
CTP_API void TraderSpiSetOnRspQryParkedOrder(TraderSpiHandle spi, TraderOnRspQryParkedOrderCallback callback);
// 请求查询预埋撤单响应
CTP_API void TraderSpiSetOnRspQryParkedOrderAction(TraderSpiHandle spi, TraderOnRspQryParkedOrderActionCallback callback);
// 请求查询交易通知响应
CTP_API void TraderSpiSetOnRspQryTradingNotice(TraderSpiHandle spi, TraderOnRspQryTradingNoticeCallback callback);
// 请求查询经纪公司交易参数响应
CTP_API void TraderSpiSetOnRspQryBrokerTradingParams(TraderSpiHandle spi, TraderOnRspQryBrokerTradingParamsCallback callback);
// 请求查询经纪公司交易算法响应
CTP_API void TraderSpiSetOnRspQryBrokerTradingAlgos(TraderSpiHandle spi, TraderOnRspQryBrokerTradingAlgosCallback callback);
// 请求查询监控中心用户令牌
CTP_API void TraderSpiSetOnRspQueryCFMMCTradingAccountToken(TraderSpiHandle spi, TraderOnRspQueryCFMMCTradingAccountTokenCallback callback);
// 银行发起银行资金转期货通知
CTP_API void TraderSpiSetOnRtnFromBankToFutureByBank(TraderSpiHandle spi, TraderOnRtnFromBankToFutureByBankCallback callback);
// 银行发起期货资金转银行通知
CTP_API void TraderSpiSetOnRtnFromFutureToBankByBank(TraderSpiHandle spi, TraderOnRtnFromFutureToBankByBankCallback callback);
// 银行发起冲正银行转期货通知
CTP_API void TraderSpiSetOnRtnRepealFromBankToFutureByBank(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByBankCallback callback);
// 银行发起冲正期货转银行通知
CTP_API void TraderSpiSetOnRtnRepealFromFutureToBankByBank(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByBankCallback callback);
// 期货发起银行资金转期货通知
CTP_API void TraderSpiSetOnRtnFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRtnFromBankToFutureByFutureCallback callback);
// 期货发起期货资金转银行通知
CTP_API void TraderSpiSetOnRtnFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRtnFromFutureToBankByFutureCallback callback);
// 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
CTP_API void TraderSpiSetOnRtnRepealFromBankToFutureByFutureManual(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByFutureManualCallback callback);
// 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
CTP_API void TraderSpiSetOnRtnRepealFromFutureToBankByFutureManual(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByFutureManualCallback callback);
// 期货发起查询银行余额通知
CTP_API void TraderSpiSetOnRtnQueryBankBalanceByFuture(TraderSpiHandle spi, TraderOnRtnQueryBankBalanceByFutureCallback callback);
// 期货发起银行资金转期货错误回报
CTP_API void TraderSpiSetOnErrRtnBankToFutureByFuture(TraderSpiHandle spi, TraderOnErrRtnBankToFutureByFutureCallback callback);
// 期货发起期货资金转银行错误回报
CTP_API void TraderSpiSetOnErrRtnFutureToBankByFuture(TraderSpiHandle spi, TraderOnErrRtnFutureToBankByFutureCallback callback);
// 系统运行时期货端手工发起冲正银行转期货错误回报
CTP_API void TraderSpiSetOnErrRtnRepealBankToFutureByFutureManual(TraderSpiHandle spi, TraderOnErrRtnRepealBankToFutureByFutureManualCallback callback);
// 系统运行时期货端手工发起冲正期货转银行错误回报
CTP_API void TraderSpiSetOnErrRtnRepealFutureToBankByFutureManual(TraderSpiHandle spi, TraderOnErrRtnRepealFutureToBankByFutureManualCallback callback);
// 期货发起查询银行余额错误回报
CTP_API void TraderSpiSetOnErrRtnQueryBankBalanceByFuture(TraderSpiHandle spi, TraderOnErrRtnQueryBankBalanceByFutureCallback callback);
// 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
CTP_API void TraderSpiSetOnRtnRepealFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRtnRepealFromBankToFutureByFutureCallback callback);
// 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
CTP_API void TraderSpiSetOnRtnRepealFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRtnRepealFromFutureToBankByFutureCallback callback);
// 期货发起银行资金转期货应答
CTP_API void TraderSpiSetOnRspFromBankToFutureByFuture(TraderSpiHandle spi, TraderOnRspFromBankToFutureByFutureCallback callback);
// 期货发起期货资金转银行应答
CTP_API void TraderSpiSetOnRspFromFutureToBankByFuture(TraderSpiHandle spi, TraderOnRspFromFutureToBankByFutureCallback callback);
// 期货发起查询银行余额应答
CTP_API void TraderSpiSetOnRspQueryBankAccountMoneyByFuture(TraderSpiHandle spi, TraderOnRspQueryBankAccountMoneyByFutureCallback callback);
// 银行发起银期开户通知
CTP_API void TraderSpiSetOnRtnOpenAccountByBank(TraderSpiHandle spi, TraderOnRtnOpenAccountByBankCallback callback);
// 银行发起银期销户通知
CTP_API void TraderSpiSetOnRtnCancelAccountByBank(TraderSpiHandle spi, TraderOnRtnCancelAccountByBankCallback callback);
// 银行发起变更银行账号通知
CTP_API void TraderSpiSetOnRtnChangeAccountByBank(TraderSpiHandle spi, TraderOnRtnChangeAccountByBankCallback callback);
// 请求查询分类合约响应
CTP_API void TraderSpiSetOnRspQryClassifiedInstrument(TraderSpiHandle spi, TraderOnRspQryClassifiedInstrumentCallback callback);
// 请求组合优惠比例响应
CTP_API void TraderSpiSetOnRspQryCombPromotionParam(TraderSpiHandle spi, TraderOnRspQryCombPromotionParamCallback callback);
// 投资者风险结算持仓查询响应
CTP_API void TraderSpiSetOnRspQryRiskSettleInvstPosition(TraderSpiHandle spi, TraderOnRspQryRiskSettleInvstPositionCallback callback);
// 风险结算产品查询响应
CTP_API void TraderSpiSetOnRspQryRiskSettleProductStatus(TraderSpiHandle spi, TraderOnRspQryRiskSettleProductStatusCallback callback);
// SPBM期货合约参数查询响应
CTP_API void TraderSpiSetOnRspQrySPBMFutureParameter(TraderSpiHandle spi, TraderOnRspQrySPBMFutureParameterCallback callback);
// SPBM期权合约参数查询响应
CTP_API void TraderSpiSetOnRspQrySPBMOptionParameter(TraderSpiHandle spi, TraderOnRspQrySPBMOptionParameterCallback callback);
// SPBM品种内对锁仓折扣参数查询响应
CTP_API void TraderSpiSetOnRspQrySPBMIntraParameter(TraderSpiHandle spi, TraderOnRspQrySPBMIntraParameterCallback callback);
// SPBM跨品种抵扣参数查询响应
CTP_API void TraderSpiSetOnRspQrySPBMInterParameter(TraderSpiHandle spi, TraderOnRspQrySPBMInterParameterCallback callback);
// SPBM组合保证金套餐查询响应
CTP_API void TraderSpiSetOnRspQrySPBMPortfDefinition(TraderSpiHandle spi, TraderOnRspQrySPBMPortfDefinitionCallback callback);
// 投资者SPBM套餐选择查询响应
CTP_API void TraderSpiSetOnRspQrySPBMInvestorPortfDef(TraderSpiHandle spi, TraderOnRspQrySPBMInvestorPortfDefCallback callback);
// 投资者新型组合保证金系数查询响应
CTP_API void TraderSpiSetOnRspQryInvestorPortfMarginRatio(TraderSpiHandle spi, TraderOnRspQryInvestorPortfMarginRatioCallback callback);
// 投资者产品SPBM明细查询响应
CTP_API void TraderSpiSetOnRspQryInvestorProdSPBMDetail(TraderSpiHandle spi, TraderOnRspQryInvestorProdSPBMDetailCallback callback);
// 投资者商品组SPMM记录查询响应
CTP_API void TraderSpiSetOnRspQryInvestorCommoditySPMMMargin(TraderSpiHandle spi, TraderOnRspQryInvestorCommoditySPMMMarginCallback callback);
// 投资者商品群SPMM记录查询响应
CTP_API void TraderSpiSetOnRspQryInvestorCommodityGroupSPMMMargin(TraderSpiHandle spi, TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback callback);
// SPMM合约参数查询响应
CTP_API void TraderSpiSetOnRspQrySPMMInstParam(TraderSpiHandle spi, TraderOnRspQrySPMMInstParamCallback callback);
// SPMM产品参数查询响应
CTP_API void TraderSpiSetOnRspQrySPMMProductParam(TraderSpiHandle spi, TraderOnRspQrySPMMProductParamCallback callback);
// SPBM附加跨品种抵扣参数查询响应
CTP_API void TraderSpiSetOnRspQrySPBMAddOnInterParameter(TraderSpiHandle spi, TraderOnRspQrySPBMAddOnInterParameterCallback callback);
// RCAMS产品组合信息查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSCombProductInfo(TraderSpiHandle spi, TraderOnRspQryRCAMSCombProductInfoCallback callback);
// RCAMS同合约风险对冲参数查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSInstrParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSInstrParameterCallback callback);
// RCAMS品种内风险对冲参数查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSIntraParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSIntraParameterCallback callback);
// RCAMS跨品种风险折抵参数查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSInterParameter(TraderSpiHandle spi, TraderOnRspQryRCAMSInterParameterCallback callback);
// RCAMS空头期权风险调整参数查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSShortOptAdjustParam(TraderSpiHandle spi, TraderOnRspQryRCAMSShortOptAdjustParamCallback callback);
// RCAMS策略组合持仓查询响应
CTP_API void TraderSpiSetOnRspQryRCAMSInvestorCombPosition(TraderSpiHandle spi, TraderOnRspQryRCAMSInvestorCombPositionCallback callback);
// 投资者品种RCAMS保证金查询响应
CTP_API void TraderSpiSetOnRspQryInvestorProdRCAMSMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProdRCAMSMarginCallback callback);
// RULE合约保证金参数查询响应
CTP_API void TraderSpiSetOnRspQryRULEInstrParameter(TraderSpiHandle spi, TraderOnRspQryRULEInstrParameterCallback callback);
// RULE品种内对锁仓折扣参数查询响应
CTP_API void TraderSpiSetOnRspQryRULEIntraParameter(TraderSpiHandle spi, TraderOnRspQryRULEIntraParameterCallback callback);
// RULE跨品种抵扣参数查询响应
CTP_API void TraderSpiSetOnRspQryRULEInterParameter(TraderSpiHandle spi, TraderOnRspQryRULEInterParameterCallback callback);
// 投资者产品RULE保证金查询响应
CTP_API void TraderSpiSetOnRspQryInvestorProdRULEMargin(TraderSpiHandle spi, TraderOnRspQryInvestorProdRULEMarginCallback callback);
// 投资者投资者新组保设置查询响应
CTP_API void TraderSpiSetOnRspQryInvestorPortfSetting(TraderSpiHandle spi, TraderOnRspQryInvestorPortfSettingCallback callback);

// ========== 跨平台统一登录接口 ==========
// 说明: macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数
// 此函数在 Linux/Windows 上忽略 systemInfo，在 macOS 上使用它

// 带系统信息的用户登录请求（跨平台统一接口）
// systemInfoLen: 系统信息长度，传 0 表示自动采集（仅 macOS 生效）
// systemInfo: 系统信息数据，传 NULL 表示自动采集（仅 macOS 生效）
CTP_API int TraderReqUserLoginWithSystemInfo(TraderApiHandle handle,
    CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID,
    int systemInfoLen, const char* systemInfo);

// ========== DataCollect 函数 ==========

// 获取终端信息（AES+RSA 加密）
// pSystemInfo: 输出缓冲区，至少 270 字节
// pLen: 输入缓冲区大小，输出实际数据长度
// 返回值: 0 成功，非 0 表示采集错误（按位判断）
CTP_API int DCGetSystemInfo(char* pSystemInfo, int* pLen);

// 获取终端信息（未 AES 加密）
CTP_API int DCGetSystemInfoUnAesEncode(char* pSystemInfo, int* pLen);

// 获取 DataCollect API 版本
CTP_API const char* DCGetDataCollectApiVersion(void);

#ifdef __cplusplus
}
#endif

#endif // CTP_TRADER_C_API_H
