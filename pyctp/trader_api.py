"""
CTP 交易 API 封装

此文件由代码生成器自动生成，请勿手动修改
CTP 交易 API 封装
"""

import ctypes
import os
import threading
from abc import ABC, abstractmethod
from typing import Optional, List

from .loader import auto_load_library, get_trader_lib_handle
from .struct import *
from .utils import *

# ========== 回调类型定义 ==========

# TraderOnFrontConnectedCallback ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
TraderOnFrontConnectedCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p)

# TraderOnFrontDisconnectedCallback 0x2003 收到错误报文
TraderOnFrontDisconnectedCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)

# TraderOnHeartBeatWarningCallback 心跳超时警告。当长时间未收到报文时，该方法被调用。
TraderOnHeartBeatWarningCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)

# TraderOnRspAuthenticateCallback 客户端认证响应
TraderOnRspAuthenticateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspAuthenticateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspUserLoginCallback 登录请求响应
TraderOnRspUserLoginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserLoginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspUserLogoutCallback 登出请求响应
TraderOnRspUserLogoutCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspUserPasswordUpdateCallback 用户口令更新请求响应
TraderOnRspUserPasswordUpdateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserPasswordUpdateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspTradingAccountPasswordUpdateCallback 资金账户口令更新请求响应
TraderOnRspTradingAccountPasswordUpdateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspUserAuthMethodCallback 查询用户当前支持的认证模式的回复
TraderOnRspUserAuthMethodCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserAuthMethodField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspGenUserCaptchaCallback 获取图形验证码请求的回复
TraderOnRspGenUserCaptchaCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspGenUserCaptchaField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspGenUserTextCallback 获取短信验证码请求的回复
TraderOnRspGenUserTextCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspGenUserTextField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspOrderInsertCallback 报单录入请求响应
TraderOnRspOrderInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspParkedOrderInsertCallback 预埋单录入请求响应
TraderOnRspParkedOrderInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspParkedOrderActionCallback 预埋撤单录入请求响应
TraderOnRspParkedOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspOrderActionCallback 报单操作请求响应
TraderOnRspOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryMaxOrderVolumeCallback 查询最大报单数量响应
TraderOnRspQryMaxOrderVolumeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspSettlementInfoConfirmCallback 投资者结算结果确认响应
TraderOnRspSettlementInfoConfirmCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspRemoveParkedOrderCallback 删除预埋单响应
TraderOnRspRemoveParkedOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspRemoveParkedOrderActionCallback 删除预埋撤单响应
TraderOnRspRemoveParkedOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspExecOrderInsertCallback 执行宣告录入请求响应
TraderOnRspExecOrderInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspExecOrderActionCallback 执行宣告操作请求响应
TraderOnRspExecOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspForQuoteInsertCallback 询价录入请求响应
TraderOnRspForQuoteInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQuoteInsertCallback 报价录入请求响应
TraderOnRspQuoteInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQuoteActionCallback 报价操作请求响应
TraderOnRspQuoteActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspBatchOrderActionCallback 批量报单操作请求响应
TraderOnRspBatchOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputBatchOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspOptionSelfCloseInsertCallback 期权自对冲录入请求响应
TraderOnRspOptionSelfCloseInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspOptionSelfCloseActionCallback 期权自对冲操作请求响应
TraderOnRspOptionSelfCloseActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspCombActionInsertCallback 申请组合录入请求响应
TraderOnRspCombActionInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputCombActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryOrderCallback 请求查询报单响应
TraderOnRspQryOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTradeCallback 请求查询成交响应
TraderOnRspQryTradeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorPositionCallback 请求查询投资者持仓响应
TraderOnRspQryInvestorPositionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTradingAccountCallback 请求查询资金账户响应
TraderOnRspQryTradingAccountCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorCallback 请求查询投资者响应
TraderOnRspQryInvestorCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTradingCodeCallback 请求查询交易编码响应
TraderOnRspQryTradingCodeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingCodeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInstrumentMarginRateCallback 请求查询合约保证金率响应
TraderOnRspQryInstrumentMarginRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentMarginRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInstrumentCommissionRateCallback 请求查询合约手续费率响应
TraderOnRspQryInstrumentCommissionRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentCommissionRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryExchangeCallback 请求查询交易所响应
TraderOnRspQryExchangeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryProductCallback 请求查询产品响应
TraderOnRspQryProductCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInstrumentCallback 请求查询合约响应
TraderOnRspQryInstrumentCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryDepthMarketDataCallback 请求查询行情响应
TraderOnRspQryDepthMarketDataCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcDepthMarketDataField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTraderOfferCallback 请求查询交易员报盘机响应
TraderOnRspQryTraderOfferCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTraderOfferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySettlementInfoCallback 请求查询投资者结算结果响应
TraderOnRspQrySettlementInfoCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTransferBankCallback 请求查询转帐银行响应
TraderOnRspQryTransferBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTransferBankField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorPositionDetailCallback 请求查询投资者持仓明细响应
TraderOnRspQryInvestorPositionDetailCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryNoticeCallback 请求查询客户通知响应
TraderOnRspQryNoticeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcNoticeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySettlementInfoConfirmCallback 请求查询结算信息确认响应
TraderOnRspQrySettlementInfoConfirmCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorPositionCombineDetailCallback 请求查询投资者持仓明细响应
TraderOnRspQryInvestorPositionCombineDetailCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionCombineDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryCFMMCTradingAccountKeyCallback 查询保证金监管系统经纪公司资金账户密钥响应
TraderOnRspQryCFMMCTradingAccountKeyCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCFMMCTradingAccountKeyField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryEWarrantOffsetCallback 请求查询仓单折抵信息响应
TraderOnRspQryEWarrantOffsetCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcEWarrantOffsetField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorProductGroupMarginCallback 请求查询投资者品种/跨品种保证金响应
TraderOnRspQryInvestorProductGroupMarginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProductGroupMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryExchangeMarginRateCallback 请求查询交易所保证金率响应
TraderOnRspQryExchangeMarginRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeMarginRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryExchangeMarginRateAdjustCallback 请求查询交易所调整保证金率响应
TraderOnRspQryExchangeMarginRateAdjustCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeMarginRateAdjustField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryExchangeRateCallback 请求查询汇率响应
TraderOnRspQryExchangeRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySecAgentACIDMapCallback 请求查询二级代理操作员银期权限响应
TraderOnRspQrySecAgentACIDMapCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentACIDMapField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryProductExchRateCallback 请求查询产品报价汇率
TraderOnRspQryProductExchRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductExchRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryProductGroupCallback 请求查询产品组
TraderOnRspQryProductGroupCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductGroupField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryMMInstrumentCommissionRateCallback 请求查询做市商合约手续费率响应
TraderOnRspQryMMInstrumentCommissionRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMMInstrumentCommissionRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryMMOptionInstrCommRateCallback 请求查询做市商期权合约手续费响应
TraderOnRspQryMMOptionInstrCommRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMMOptionInstrCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInstrumentOrderCommRateCallback 请求查询报单手续费响应
TraderOnRspQryInstrumentOrderCommRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentOrderCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySecAgentTradingAccountCallback 请求查询资金账户响应
TraderOnRspQrySecAgentTradingAccountCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySecAgentCheckModeCallback 请求查询二级代理商资金校验模式响应
TraderOnRspQrySecAgentCheckModeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentCheckModeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySecAgentTradeInfoCallback 请求查询二级代理商信息响应
TraderOnRspQrySecAgentTradeInfoCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentTradeInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryOptionInstrTradeCostCallback 请求查询期权交易成本响应
TraderOnRspQryOptionInstrTradeCostCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionInstrTradeCostField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryOptionInstrCommRateCallback 请求查询期权合约手续费响应
TraderOnRspQryOptionInstrCommRateCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionInstrCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryExecOrderCallback 请求查询执行宣告响应
TraderOnRspQryExecOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryForQuoteCallback 请求查询询价响应
TraderOnRspQryForQuoteCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryQuoteCallback 请求查询报价响应
TraderOnRspQryQuoteCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryOptionSelfCloseCallback 请求查询期权自对冲响应
TraderOnRspQryOptionSelfCloseCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestUnitCallback 请求查询投资单元响应
TraderOnRspQryInvestUnitCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestUnitField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryCombInstrumentGuardCallback 请求查询组合合约安全系数响应
TraderOnRspQryCombInstrumentGuardCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombInstrumentGuardField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryCombActionCallback 请求查询申请组合响应
TraderOnRspQryCombActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTransferSerialCallback 请求查询转帐流水响应
TraderOnRspQryTransferSerialCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTransferSerialField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryAccountregisterCallback 请求查询银期签约关系响应
TraderOnRspQryAccountregisterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcAccountregisterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspErrorCallback 错误应答
TraderOnRspErrorCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRtnOrderCallback 报单通知
TraderOnRtnOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderField))

# TraderOnRtnTradeCallback 成交通知
TraderOnRtnTradeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradeField))

# TraderOnErrRtnOrderInsertCallback 报单录入错误回报
TraderOnErrRtnOrderInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnOrderActionCallback 报单操作错误回报
TraderOnErrRtnOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnInstrumentStatusCallback 合约交易状态通知
TraderOnRtnInstrumentStatusCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentStatusField))

# TraderOnRtnBulletinCallback 交易所公告通知
TraderOnRtnBulletinCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBulletinField))

# TraderOnRtnTradingNoticeCallback 交易通知
TraderOnRtnTradingNoticeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingNoticeInfoField))

# TraderOnRtnErrorConditionalOrderCallback 提示条件单校验错误
TraderOnRtnErrorConditionalOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcErrorConditionalOrderField))

# TraderOnRtnExecOrderCallback 执行宣告通知
TraderOnRtnExecOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderField))

# TraderOnErrRtnExecOrderInsertCallback 执行宣告录入错误回报
TraderOnErrRtnExecOrderInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnExecOrderActionCallback 执行宣告操作错误回报
TraderOnErrRtnExecOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnForQuoteInsertCallback 询价录入错误回报
TraderOnErrRtnForQuoteInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnQuoteCallback 报价通知
TraderOnRtnQuoteCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteField))

# TraderOnErrRtnQuoteInsertCallback 报价录入错误回报
TraderOnErrRtnQuoteInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnQuoteActionCallback 报价操作错误回报
TraderOnErrRtnQuoteActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnForQuoteRspCallback 询价通知
TraderOnRtnForQuoteRspCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteRspField))

# TraderOnRtnCFMMCTradingAccountTokenCallback 保证金监控中心用户令牌
TraderOnRtnCFMMCTradingAccountTokenCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCFMMCTradingAccountTokenField))

# TraderOnErrRtnBatchOrderActionCallback 批量报单操作错误回报
TraderOnErrRtnBatchOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBatchOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnOptionSelfCloseCallback 期权自对冲通知
TraderOnRtnOptionSelfCloseCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseField))

# TraderOnErrRtnOptionSelfCloseInsertCallback 期权自对冲录入错误回报
TraderOnErrRtnOptionSelfCloseInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnOptionSelfCloseActionCallback 期权自对冲操作错误回报
TraderOnErrRtnOptionSelfCloseActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnCombActionCallback 申请组合通知
TraderOnRtnCombActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombActionField))

# TraderOnErrRtnCombActionInsertCallback 申请组合录入错误回报
TraderOnErrRtnCombActionInsertCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputCombActionField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRspQryContractBankCallback 请求查询签约银行响应
TraderOnRspQryContractBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcContractBankField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryParkedOrderCallback 请求查询预埋单响应
TraderOnRspQryParkedOrderCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryParkedOrderActionCallback 请求查询预埋撤单响应
TraderOnRspQryParkedOrderActionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryTradingNoticeCallback 请求查询交易通知响应
TraderOnRspQryTradingNoticeCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingNoticeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryBrokerTradingParamsCallback 请求查询经纪公司交易参数响应
TraderOnRspQryBrokerTradingParamsCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBrokerTradingParamsField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryBrokerTradingAlgosCallback 请求查询经纪公司交易算法响应
TraderOnRspQryBrokerTradingAlgosCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBrokerTradingAlgosField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQueryCFMMCTradingAccountTokenCallback 请求查询监控中心用户令牌
TraderOnRspQueryCFMMCTradingAccountTokenCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRtnFromBankToFutureByBankCallback 银行发起银行资金转期货通知
TraderOnRtnFromBankToFutureByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))

# TraderOnRtnFromFutureToBankByBankCallback 银行发起期货资金转银行通知
TraderOnRtnFromFutureToBankByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))

# TraderOnRtnRepealFromBankToFutureByBankCallback 银行发起冲正银行转期货通知
TraderOnRtnRepealFromBankToFutureByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRtnRepealFromFutureToBankByBankCallback 银行发起冲正期货转银行通知
TraderOnRtnRepealFromFutureToBankByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRtnFromBankToFutureByFutureCallback 期货发起银行资金转期货通知
TraderOnRtnFromBankToFutureByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))

# TraderOnRtnFromFutureToBankByFutureCallback 期货发起期货资金转银行通知
TraderOnRtnFromFutureToBankByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))

# TraderOnRtnRepealFromBankToFutureByFutureManualCallback 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
TraderOnRtnRepealFromBankToFutureByFutureManualCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRtnRepealFromFutureToBankByFutureManualCallback 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
TraderOnRtnRepealFromFutureToBankByFutureManualCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRtnQueryBankBalanceByFutureCallback 期货发起查询银行余额通知
TraderOnRtnQueryBankBalanceByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcNotifyQueryAccountField))

# TraderOnErrRtnBankToFutureByFutureCallback 期货发起银行资金转期货错误回报
TraderOnErrRtnBankToFutureByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnFutureToBankByFutureCallback 期货发起期货资金转银行错误回报
TraderOnErrRtnFutureToBankByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnRepealBankToFutureByFutureManualCallback 系统运行时期货端手工发起冲正银行转期货错误回报
TraderOnErrRtnRepealBankToFutureByFutureManualCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqRepealField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnRepealFutureToBankByFutureManualCallback 系统运行时期货端手工发起冲正期货转银行错误回报
TraderOnErrRtnRepealFutureToBankByFutureManualCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqRepealField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnErrRtnQueryBankBalanceByFutureCallback 期货发起查询银行余额错误回报
TraderOnErrRtnQueryBankBalanceByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqQueryAccountField), ctypes.POINTER(CThostFtdcRspInfoField))

# TraderOnRtnRepealFromBankToFutureByFutureCallback 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知
TraderOnRtnRepealFromBankToFutureByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRtnRepealFromFutureToBankByFutureCallback 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知
TraderOnRtnRepealFromFutureToBankByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))

# TraderOnRspFromBankToFutureByFutureCallback 期货发起银行资金转期货应答
TraderOnRspFromBankToFutureByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspFromFutureToBankByFutureCallback 期货发起期货资金转银行应答
TraderOnRspFromFutureToBankByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQueryBankAccountMoneyByFutureCallback 期货发起查询银行余额应答
TraderOnRspQueryBankAccountMoneyByFutureCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqQueryAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRtnOpenAccountByBankCallback 银行发起银期开户通知
TraderOnRtnOpenAccountByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOpenAccountField))

# TraderOnRtnCancelAccountByBankCallback 银行发起银期销户通知
TraderOnRtnCancelAccountByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCancelAccountField))

# TraderOnRtnChangeAccountByBankCallback 银行发起变更银行账号通知
TraderOnRtnChangeAccountByBankCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcChangeAccountField))

# TraderOnRspQryClassifiedInstrumentCallback 请求查询分类合约响应
TraderOnRspQryClassifiedInstrumentCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryCombPromotionParamCallback 请求组合优惠比例响应
TraderOnRspQryCombPromotionParamCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombPromotionParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRiskSettleInvstPositionCallback 投资者风险结算持仓查询响应
TraderOnRspQryRiskSettleInvstPositionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRiskSettleInvstPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRiskSettleProductStatusCallback 风险结算产品查询响应
TraderOnRspQryRiskSettleProductStatusCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRiskSettleProductStatusField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMFutureParameterCallback SPBM期货合约参数查询响应
TraderOnRspQrySPBMFutureParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMFutureParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMOptionParameterCallback SPBM期权合约参数查询响应
TraderOnRspQrySPBMOptionParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMOptionParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMIntraParameterCallback SPBM品种内对锁仓折扣参数查询响应
TraderOnRspQrySPBMIntraParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMInterParameterCallback SPBM跨品种抵扣参数查询响应
TraderOnRspQrySPBMInterParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMPortfDefinitionCallback SPBM组合保证金套餐查询响应
TraderOnRspQrySPBMPortfDefinitionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMPortfDefinitionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMInvestorPortfDefCallback 投资者SPBM套餐选择查询响应
TraderOnRspQrySPBMInvestorPortfDefCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMInvestorPortfDefField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorPortfMarginRatioCallback 投资者新型组合保证金系数查询响应
TraderOnRspQryInvestorPortfMarginRatioCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPortfMarginRatioField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorProdSPBMDetailCallback 投资者产品SPBM明细查询响应
TraderOnRspQryInvestorProdSPBMDetailCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdSPBMDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorCommoditySPMMMarginCallback 投资者商品组SPMM记录查询响应
TraderOnRspQryInvestorCommoditySPMMMarginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorCommoditySPMMMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback 投资者商品群SPMM记录查询响应
TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorCommodityGroupSPMMMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPMMInstParamCallback SPMM合约参数查询响应
TraderOnRspQrySPMMInstParamCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPMMInstParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPMMProductParamCallback SPMM产品参数查询响应
TraderOnRspQrySPMMProductParamCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPMMProductParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQrySPBMAddOnInterParameterCallback SPBM附加跨品种抵扣参数查询响应
TraderOnRspQrySPBMAddOnInterParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMAddOnInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSCombProductInfoCallback RCAMS产品组合信息查询响应
TraderOnRspQryRCAMSCombProductInfoCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSCombProductInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSInstrParameterCallback RCAMS同合约风险对冲参数查询响应
TraderOnRspQryRCAMSInstrParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInstrParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSIntraParameterCallback RCAMS品种内风险对冲参数查询响应
TraderOnRspQryRCAMSIntraParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSInterParameterCallback RCAMS跨品种风险折抵参数查询响应
TraderOnRspQryRCAMSInterParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSShortOptAdjustParamCallback RCAMS空头期权风险调整参数查询响应
TraderOnRspQryRCAMSShortOptAdjustParamCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSShortOptAdjustParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRCAMSInvestorCombPositionCallback RCAMS策略组合持仓查询响应
TraderOnRspQryRCAMSInvestorCombPositionCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInvestorCombPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorProdRCAMSMarginCallback 投资者品种RCAMS保证金查询响应
TraderOnRspQryInvestorProdRCAMSMarginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdRCAMSMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRULEInstrParameterCallback RULE合约保证金参数查询响应
TraderOnRspQryRULEInstrParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEInstrParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRULEIntraParameterCallback RULE品种内对锁仓折扣参数查询响应
TraderOnRspQryRULEIntraParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryRULEInterParameterCallback RULE跨品种抵扣参数查询响应
TraderOnRspQryRULEInterParameterCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorProdRULEMarginCallback 投资者产品RULE保证金查询响应
TraderOnRspQryInvestorProdRULEMarginCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdRULEMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# TraderOnRspQryInvestorPortfSettingCallback 投资者投资者新组保设置查询响应
TraderOnRspQryInvestorPortfSettingCallback = ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPortfSettingField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)

# ========== 回调结构体定义 ==========

# TraderSpiCallbacks 回调结构体（用于批量设置）
class TraderSpiCallbacks(ctypes.Structure):
    """回调结构体（用于批量设置）"""
    _fields_ = [
        ("userData", ctypes.c_void_p),
        ("onFrontConnected", TraderOnFrontConnectedCallback),
        ("onFrontDisconnected", TraderOnFrontDisconnectedCallback),
        ("onHeartBeatWarning", TraderOnHeartBeatWarningCallback),
        ("onRspAuthenticate", TraderOnRspAuthenticateCallback),
        ("onRspUserLogin", TraderOnRspUserLoginCallback),
        ("onRspUserLogout", TraderOnRspUserLogoutCallback),
        ("onRspUserPasswordUpdate", TraderOnRspUserPasswordUpdateCallback),
        ("onRspTradingAccountPasswordUpdate", TraderOnRspTradingAccountPasswordUpdateCallback),
        ("onRspUserAuthMethod", TraderOnRspUserAuthMethodCallback),
        ("onRspGenUserCaptcha", TraderOnRspGenUserCaptchaCallback),
        ("onRspGenUserText", TraderOnRspGenUserTextCallback),
        ("onRspOrderInsert", TraderOnRspOrderInsertCallback),
        ("onRspParkedOrderInsert", TraderOnRspParkedOrderInsertCallback),
        ("onRspParkedOrderAction", TraderOnRspParkedOrderActionCallback),
        ("onRspOrderAction", TraderOnRspOrderActionCallback),
        ("onRspQryMaxOrderVolume", TraderOnRspQryMaxOrderVolumeCallback),
        ("onRspSettlementInfoConfirm", TraderOnRspSettlementInfoConfirmCallback),
        ("onRspRemoveParkedOrder", TraderOnRspRemoveParkedOrderCallback),
        ("onRspRemoveParkedOrderAction", TraderOnRspRemoveParkedOrderActionCallback),
        ("onRspExecOrderInsert", TraderOnRspExecOrderInsertCallback),
        ("onRspExecOrderAction", TraderOnRspExecOrderActionCallback),
        ("onRspForQuoteInsert", TraderOnRspForQuoteInsertCallback),
        ("onRspQuoteInsert", TraderOnRspQuoteInsertCallback),
        ("onRspQuoteAction", TraderOnRspQuoteActionCallback),
        ("onRspBatchOrderAction", TraderOnRspBatchOrderActionCallback),
        ("onRspOptionSelfCloseInsert", TraderOnRspOptionSelfCloseInsertCallback),
        ("onRspOptionSelfCloseAction", TraderOnRspOptionSelfCloseActionCallback),
        ("onRspCombActionInsert", TraderOnRspCombActionInsertCallback),
        ("onRspQryOrder", TraderOnRspQryOrderCallback),
        ("onRspQryTrade", TraderOnRspQryTradeCallback),
        ("onRspQryInvestorPosition", TraderOnRspQryInvestorPositionCallback),
        ("onRspQryTradingAccount", TraderOnRspQryTradingAccountCallback),
        ("onRspQryInvestor", TraderOnRspQryInvestorCallback),
        ("onRspQryTradingCode", TraderOnRspQryTradingCodeCallback),
        ("onRspQryInstrumentMarginRate", TraderOnRspQryInstrumentMarginRateCallback),
        ("onRspQryInstrumentCommissionRate", TraderOnRspQryInstrumentCommissionRateCallback),
        ("onRspQryExchange", TraderOnRspQryExchangeCallback),
        ("onRspQryProduct", TraderOnRspQryProductCallback),
        ("onRspQryInstrument", TraderOnRspQryInstrumentCallback),
        ("onRspQryDepthMarketData", TraderOnRspQryDepthMarketDataCallback),
        ("onRspQryTraderOffer", TraderOnRspQryTraderOfferCallback),
        ("onRspQrySettlementInfo", TraderOnRspQrySettlementInfoCallback),
        ("onRspQryTransferBank", TraderOnRspQryTransferBankCallback),
        ("onRspQryInvestorPositionDetail", TraderOnRspQryInvestorPositionDetailCallback),
        ("onRspQryNotice", TraderOnRspQryNoticeCallback),
        ("onRspQrySettlementInfoConfirm", TraderOnRspQrySettlementInfoConfirmCallback),
        ("onRspQryInvestorPositionCombineDetail", TraderOnRspQryInvestorPositionCombineDetailCallback),
        ("onRspQryCFMMCTradingAccountKey", TraderOnRspQryCFMMCTradingAccountKeyCallback),
        ("onRspQryEWarrantOffset", TraderOnRspQryEWarrantOffsetCallback),
        ("onRspQryInvestorProductGroupMargin", TraderOnRspQryInvestorProductGroupMarginCallback),
        ("onRspQryExchangeMarginRate", TraderOnRspQryExchangeMarginRateCallback),
        ("onRspQryExchangeMarginRateAdjust", TraderOnRspQryExchangeMarginRateAdjustCallback),
        ("onRspQryExchangeRate", TraderOnRspQryExchangeRateCallback),
        ("onRspQrySecAgentACIDMap", TraderOnRspQrySecAgentACIDMapCallback),
        ("onRspQryProductExchRate", TraderOnRspQryProductExchRateCallback),
        ("onRspQryProductGroup", TraderOnRspQryProductGroupCallback),
        ("onRspQryMMInstrumentCommissionRate", TraderOnRspQryMMInstrumentCommissionRateCallback),
        ("onRspQryMMOptionInstrCommRate", TraderOnRspQryMMOptionInstrCommRateCallback),
        ("onRspQryInstrumentOrderCommRate", TraderOnRspQryInstrumentOrderCommRateCallback),
        ("onRspQrySecAgentTradingAccount", TraderOnRspQrySecAgentTradingAccountCallback),
        ("onRspQrySecAgentCheckMode", TraderOnRspQrySecAgentCheckModeCallback),
        ("onRspQrySecAgentTradeInfo", TraderOnRspQrySecAgentTradeInfoCallback),
        ("onRspQryOptionInstrTradeCost", TraderOnRspQryOptionInstrTradeCostCallback),
        ("onRspQryOptionInstrCommRate", TraderOnRspQryOptionInstrCommRateCallback),
        ("onRspQryExecOrder", TraderOnRspQryExecOrderCallback),
        ("onRspQryForQuote", TraderOnRspQryForQuoteCallback),
        ("onRspQryQuote", TraderOnRspQryQuoteCallback),
        ("onRspQryOptionSelfClose", TraderOnRspQryOptionSelfCloseCallback),
        ("onRspQryInvestUnit", TraderOnRspQryInvestUnitCallback),
        ("onRspQryCombInstrumentGuard", TraderOnRspQryCombInstrumentGuardCallback),
        ("onRspQryCombAction", TraderOnRspQryCombActionCallback),
        ("onRspQryTransferSerial", TraderOnRspQryTransferSerialCallback),
        ("onRspQryAccountregister", TraderOnRspQryAccountregisterCallback),
        ("onRspError", TraderOnRspErrorCallback),
        ("onRtnOrder", TraderOnRtnOrderCallback),
        ("onRtnTrade", TraderOnRtnTradeCallback),
        ("onErrRtnOrderInsert", TraderOnErrRtnOrderInsertCallback),
        ("onErrRtnOrderAction", TraderOnErrRtnOrderActionCallback),
        ("onRtnInstrumentStatus", TraderOnRtnInstrumentStatusCallback),
        ("onRtnBulletin", TraderOnRtnBulletinCallback),
        ("onRtnTradingNotice", TraderOnRtnTradingNoticeCallback),
        ("onRtnErrorConditionalOrder", TraderOnRtnErrorConditionalOrderCallback),
        ("onRtnExecOrder", TraderOnRtnExecOrderCallback),
        ("onErrRtnExecOrderInsert", TraderOnErrRtnExecOrderInsertCallback),
        ("onErrRtnExecOrderAction", TraderOnErrRtnExecOrderActionCallback),
        ("onErrRtnForQuoteInsert", TraderOnErrRtnForQuoteInsertCallback),
        ("onRtnQuote", TraderOnRtnQuoteCallback),
        ("onErrRtnQuoteInsert", TraderOnErrRtnQuoteInsertCallback),
        ("onErrRtnQuoteAction", TraderOnErrRtnQuoteActionCallback),
        ("onRtnForQuoteRsp", TraderOnRtnForQuoteRspCallback),
        ("onRtnCFMMCTradingAccountToken", TraderOnRtnCFMMCTradingAccountTokenCallback),
        ("onErrRtnBatchOrderAction", TraderOnErrRtnBatchOrderActionCallback),
        ("onRtnOptionSelfClose", TraderOnRtnOptionSelfCloseCallback),
        ("onErrRtnOptionSelfCloseInsert", TraderOnErrRtnOptionSelfCloseInsertCallback),
        ("onErrRtnOptionSelfCloseAction", TraderOnErrRtnOptionSelfCloseActionCallback),
        ("onRtnCombAction", TraderOnRtnCombActionCallback),
        ("onErrRtnCombActionInsert", TraderOnErrRtnCombActionInsertCallback),
        ("onRspQryContractBank", TraderOnRspQryContractBankCallback),
        ("onRspQryParkedOrder", TraderOnRspQryParkedOrderCallback),
        ("onRspQryParkedOrderAction", TraderOnRspQryParkedOrderActionCallback),
        ("onRspQryTradingNotice", TraderOnRspQryTradingNoticeCallback),
        ("onRspQryBrokerTradingParams", TraderOnRspQryBrokerTradingParamsCallback),
        ("onRspQryBrokerTradingAlgos", TraderOnRspQryBrokerTradingAlgosCallback),
        ("onRspQueryCFMMCTradingAccountToken", TraderOnRspQueryCFMMCTradingAccountTokenCallback),
        ("onRtnFromBankToFutureByBank", TraderOnRtnFromBankToFutureByBankCallback),
        ("onRtnFromFutureToBankByBank", TraderOnRtnFromFutureToBankByBankCallback),
        ("onRtnRepealFromBankToFutureByBank", TraderOnRtnRepealFromBankToFutureByBankCallback),
        ("onRtnRepealFromFutureToBankByBank", TraderOnRtnRepealFromFutureToBankByBankCallback),
        ("onRtnFromBankToFutureByFuture", TraderOnRtnFromBankToFutureByFutureCallback),
        ("onRtnFromFutureToBankByFuture", TraderOnRtnFromFutureToBankByFutureCallback),
        ("onRtnRepealFromBankToFutureByFutureManual", TraderOnRtnRepealFromBankToFutureByFutureManualCallback),
        ("onRtnRepealFromFutureToBankByFutureManual", TraderOnRtnRepealFromFutureToBankByFutureManualCallback),
        ("onRtnQueryBankBalanceByFuture", TraderOnRtnQueryBankBalanceByFutureCallback),
        ("onErrRtnBankToFutureByFuture", TraderOnErrRtnBankToFutureByFutureCallback),
        ("onErrRtnFutureToBankByFuture", TraderOnErrRtnFutureToBankByFutureCallback),
        ("onErrRtnRepealBankToFutureByFutureManual", TraderOnErrRtnRepealBankToFutureByFutureManualCallback),
        ("onErrRtnRepealFutureToBankByFutureManual", TraderOnErrRtnRepealFutureToBankByFutureManualCallback),
        ("onErrRtnQueryBankBalanceByFuture", TraderOnErrRtnQueryBankBalanceByFutureCallback),
        ("onRtnRepealFromBankToFutureByFuture", TraderOnRtnRepealFromBankToFutureByFutureCallback),
        ("onRtnRepealFromFutureToBankByFuture", TraderOnRtnRepealFromFutureToBankByFutureCallback),
        ("onRspFromBankToFutureByFuture", TraderOnRspFromBankToFutureByFutureCallback),
        ("onRspFromFutureToBankByFuture", TraderOnRspFromFutureToBankByFutureCallback),
        ("onRspQueryBankAccountMoneyByFuture", TraderOnRspQueryBankAccountMoneyByFutureCallback),
        ("onRtnOpenAccountByBank", TraderOnRtnOpenAccountByBankCallback),
        ("onRtnCancelAccountByBank", TraderOnRtnCancelAccountByBankCallback),
        ("onRtnChangeAccountByBank", TraderOnRtnChangeAccountByBankCallback),
        ("onRspQryClassifiedInstrument", TraderOnRspQryClassifiedInstrumentCallback),
        ("onRspQryCombPromotionParam", TraderOnRspQryCombPromotionParamCallback),
        ("onRspQryRiskSettleInvstPosition", TraderOnRspQryRiskSettleInvstPositionCallback),
        ("onRspQryRiskSettleProductStatus", TraderOnRspQryRiskSettleProductStatusCallback),
        ("onRspQrySPBMFutureParameter", TraderOnRspQrySPBMFutureParameterCallback),
        ("onRspQrySPBMOptionParameter", TraderOnRspQrySPBMOptionParameterCallback),
        ("onRspQrySPBMIntraParameter", TraderOnRspQrySPBMIntraParameterCallback),
        ("onRspQrySPBMInterParameter", TraderOnRspQrySPBMInterParameterCallback),
        ("onRspQrySPBMPortfDefinition", TraderOnRspQrySPBMPortfDefinitionCallback),
        ("onRspQrySPBMInvestorPortfDef", TraderOnRspQrySPBMInvestorPortfDefCallback),
        ("onRspQryInvestorPortfMarginRatio", TraderOnRspQryInvestorPortfMarginRatioCallback),
        ("onRspQryInvestorProdSPBMDetail", TraderOnRspQryInvestorProdSPBMDetailCallback),
        ("onRspQryInvestorCommoditySPMMMargin", TraderOnRspQryInvestorCommoditySPMMMarginCallback),
        ("onRspQryInvestorCommodityGroupSPMMMargin", TraderOnRspQryInvestorCommodityGroupSPMMMarginCallback),
        ("onRspQrySPMMInstParam", TraderOnRspQrySPMMInstParamCallback),
        ("onRspQrySPMMProductParam", TraderOnRspQrySPMMProductParamCallback),
        ("onRspQrySPBMAddOnInterParameter", TraderOnRspQrySPBMAddOnInterParameterCallback),
        ("onRspQryRCAMSCombProductInfo", TraderOnRspQryRCAMSCombProductInfoCallback),
        ("onRspQryRCAMSInstrParameter", TraderOnRspQryRCAMSInstrParameterCallback),
        ("onRspQryRCAMSIntraParameter", TraderOnRspQryRCAMSIntraParameterCallback),
        ("onRspQryRCAMSInterParameter", TraderOnRspQryRCAMSInterParameterCallback),
        ("onRspQryRCAMSShortOptAdjustParam", TraderOnRspQryRCAMSShortOptAdjustParamCallback),
        ("onRspQryRCAMSInvestorCombPosition", TraderOnRspQryRCAMSInvestorCombPositionCallback),
        ("onRspQryInvestorProdRCAMSMargin", TraderOnRspQryInvestorProdRCAMSMarginCallback),
        ("onRspQryRULEInstrParameter", TraderOnRspQryRULEInstrParameterCallback),
        ("onRspQryRULEIntraParameter", TraderOnRspQryRULEIntraParameterCallback),
        ("onRspQryRULEInterParameter", TraderOnRspQryRULEInterParameterCallback),
        ("onRspQryInvestorProdRULEMargin", TraderOnRspQryInvestorProdRULEMarginCallback),
        ("onRspQryInvestorPortfSetting", TraderOnRspQryInvestorPortfSettingCallback),
    ]

# ========== TraderSpi 接口 ==========

class TraderSpi(ABC):
    """交易回调接口"""

    @abstractmethod
    def OnFrontConnected(self, ):
        """========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。"""
        pass

    @abstractmethod
    def OnFrontDisconnected(self, nReason: ctypes.c_int32):
        """0x2003 收到错误报文"""
        pass

    @abstractmethod
    def OnHeartBeatWarning(self, nTimeLapse: ctypes.c_int32):
        """心跳超时警告。当长时间未收到报文时，该方法被调用。"""
        pass

    @abstractmethod
    def OnRspAuthenticate(self, pRspAuthenticateField: ctypes.POINTER(CThostFtdcRspAuthenticateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """客户端认证响应"""
        pass

    @abstractmethod
    def OnRspUserLogin(self, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登录请求响应"""
        pass

    @abstractmethod
    def OnRspUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登出请求响应"""
        pass

    @abstractmethod
    def OnRspUserPasswordUpdate(self, pUserPasswordUpdate: ctypes.POINTER(CThostFtdcUserPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """用户口令更新请求响应"""
        pass

    @abstractmethod
    def OnRspTradingAccountPasswordUpdate(self, pTradingAccountPasswordUpdate: ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """资金账户口令更新请求响应"""
        pass

    @abstractmethod
    def OnRspUserAuthMethod(self, pRspUserAuthMethod: ctypes.POINTER(CThostFtdcRspUserAuthMethodField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """查询用户当前支持的认证模式的回复"""
        pass

    @abstractmethod
    def OnRspGenUserCaptcha(self, pRspGenUserCaptcha: ctypes.POINTER(CThostFtdcRspGenUserCaptchaField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """获取图形验证码请求的回复"""
        pass

    @abstractmethod
    def OnRspGenUserText(self, pRspGenUserText: ctypes.POINTER(CThostFtdcRspGenUserTextField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """获取短信验证码请求的回复"""
        pass

    @abstractmethod
    def OnRspOrderInsert(self, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """报单录入请求响应"""
        pass

    @abstractmethod
    def OnRspParkedOrderInsert(self, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """预埋单录入请求响应"""
        pass

    @abstractmethod
    def OnRspParkedOrderAction(self, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """预埋撤单录入请求响应"""
        pass

    @abstractmethod
    def OnRspOrderAction(self, pInputOrderAction: ctypes.POINTER(CThostFtdcInputOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """报单操作请求响应"""
        pass

    @abstractmethod
    def OnRspQryMaxOrderVolume(self, pQryMaxOrderVolume: ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """查询最大报单数量响应"""
        pass

    @abstractmethod
    def OnRspSettlementInfoConfirm(self, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者结算结果确认响应"""
        pass

    @abstractmethod
    def OnRspRemoveParkedOrder(self, pRemoveParkedOrder: ctypes.POINTER(CThostFtdcRemoveParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """删除预埋单响应"""
        pass

    @abstractmethod
    def OnRspRemoveParkedOrderAction(self, pRemoveParkedOrderAction: ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """删除预埋撤单响应"""
        pass

    @abstractmethod
    def OnRspExecOrderInsert(self, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """执行宣告录入请求响应"""
        pass

    @abstractmethod
    def OnRspExecOrderAction(self, pInputExecOrderAction: ctypes.POINTER(CThostFtdcInputExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """执行宣告操作请求响应"""
        pass

    @abstractmethod
    def OnRspForQuoteInsert(self, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """询价录入请求响应"""
        pass

    @abstractmethod
    def OnRspQuoteInsert(self, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """报价录入请求响应"""
        pass

    @abstractmethod
    def OnRspQuoteAction(self, pInputQuoteAction: ctypes.POINTER(CThostFtdcInputQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """报价操作请求响应"""
        pass

    @abstractmethod
    def OnRspBatchOrderAction(self, pInputBatchOrderAction: ctypes.POINTER(CThostFtdcInputBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """批量报单操作请求响应"""
        pass

    @abstractmethod
    def OnRspOptionSelfCloseInsert(self, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """期权自对冲录入请求响应"""
        pass

    @abstractmethod
    def OnRspOptionSelfCloseAction(self, pInputOptionSelfCloseAction: ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """期权自对冲操作请求响应"""
        pass

    @abstractmethod
    def OnRspCombActionInsert(self, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """申请组合录入请求响应"""
        pass

    @abstractmethod
    def OnRspQryOrder(self, pOrder: ctypes.POINTER(CThostFtdcOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询报单响应"""
        pass

    @abstractmethod
    def OnRspQryTrade(self, pTrade: ctypes.POINTER(CThostFtdcTradeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询成交响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorPosition(self, pInvestorPosition: ctypes.POINTER(CThostFtdcInvestorPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者持仓响应"""
        pass

    @abstractmethod
    def OnRspQryTradingAccount(self, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询资金账户响应"""
        pass

    @abstractmethod
    def OnRspQryInvestor(self, pInvestor: ctypes.POINTER(CThostFtdcInvestorField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者响应"""
        pass

    @abstractmethod
    def OnRspQryTradingCode(self, pTradingCode: ctypes.POINTER(CThostFtdcTradingCodeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易编码响应"""
        pass

    @abstractmethod
    def OnRspQryInstrumentMarginRate(self, pInstrumentMarginRate: ctypes.POINTER(CThostFtdcInstrumentMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询合约保证金率响应"""
        pass

    @abstractmethod
    def OnRspQryInstrumentCommissionRate(self, pInstrumentCommissionRate: ctypes.POINTER(CThostFtdcInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询合约手续费率响应"""
        pass

    @abstractmethod
    def OnRspQryExchange(self, pExchange: ctypes.POINTER(CThostFtdcExchangeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易所响应"""
        pass

    @abstractmethod
    def OnRspQryProduct(self, pProduct: ctypes.POINTER(CThostFtdcProductField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询产品响应"""
        pass

    @abstractmethod
    def OnRspQryInstrument(self, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询合约响应"""
        pass

    @abstractmethod
    def OnRspQryDepthMarketData(self, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询行情响应"""
        pass

    @abstractmethod
    def OnRspQryTraderOffer(self, pTraderOffer: ctypes.POINTER(CThostFtdcTraderOfferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易员报盘机响应"""
        pass

    @abstractmethod
    def OnRspQrySettlementInfo(self, pSettlementInfo: ctypes.POINTER(CThostFtdcSettlementInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者结算结果响应"""
        pass

    @abstractmethod
    def OnRspQryTransferBank(self, pTransferBank: ctypes.POINTER(CThostFtdcTransferBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询转帐银行响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorPositionDetail(self, pInvestorPositionDetail: ctypes.POINTER(CThostFtdcInvestorPositionDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者持仓明细响应"""
        pass

    @abstractmethod
    def OnRspQryNotice(self, pNotice: ctypes.POINTER(CThostFtdcNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询客户通知响应"""
        pass

    @abstractmethod
    def OnRspQrySettlementInfoConfirm(self, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询结算信息确认响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorPositionCombineDetail(self, pInvestorPositionCombineDetail: ctypes.POINTER(CThostFtdcInvestorPositionCombineDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者持仓明细响应"""
        pass

    @abstractmethod
    def OnRspQryCFMMCTradingAccountKey(self, pCFMMCTradingAccountKey: ctypes.POINTER(CThostFtdcCFMMCTradingAccountKeyField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """查询保证金监管系统经纪公司资金账户密钥响应"""
        pass

    @abstractmethod
    def OnRspQryEWarrantOffset(self, pEWarrantOffset: ctypes.POINTER(CThostFtdcEWarrantOffsetField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询仓单折抵信息响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorProductGroupMargin(self, pInvestorProductGroupMargin: ctypes.POINTER(CThostFtdcInvestorProductGroupMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资者品种/跨品种保证金响应"""
        pass

    @abstractmethod
    def OnRspQryExchangeMarginRate(self, pExchangeMarginRate: ctypes.POINTER(CThostFtdcExchangeMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易所保证金率响应"""
        pass

    @abstractmethod
    def OnRspQryExchangeMarginRateAdjust(self, pExchangeMarginRateAdjust: ctypes.POINTER(CThostFtdcExchangeMarginRateAdjustField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易所调整保证金率响应"""
        pass

    @abstractmethod
    def OnRspQryExchangeRate(self, pExchangeRate: ctypes.POINTER(CThostFtdcExchangeRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询汇率响应"""
        pass

    @abstractmethod
    def OnRspQrySecAgentACIDMap(self, pSecAgentACIDMap: ctypes.POINTER(CThostFtdcSecAgentACIDMapField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询二级代理操作员银期权限响应"""
        pass

    @abstractmethod
    def OnRspQryProductExchRate(self, pProductExchRate: ctypes.POINTER(CThostFtdcProductExchRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询产品报价汇率"""
        pass

    @abstractmethod
    def OnRspQryProductGroup(self, pProductGroup: ctypes.POINTER(CThostFtdcProductGroupField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询产品组"""
        pass

    @abstractmethod
    def OnRspQryMMInstrumentCommissionRate(self, pMMInstrumentCommissionRate: ctypes.POINTER(CThostFtdcMMInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询做市商合约手续费率响应"""
        pass

    @abstractmethod
    def OnRspQryMMOptionInstrCommRate(self, pMMOptionInstrCommRate: ctypes.POINTER(CThostFtdcMMOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询做市商期权合约手续费响应"""
        pass

    @abstractmethod
    def OnRspQryInstrumentOrderCommRate(self, pInstrumentOrderCommRate: ctypes.POINTER(CThostFtdcInstrumentOrderCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询报单手续费响应"""
        pass

    @abstractmethod
    def OnRspQrySecAgentTradingAccount(self, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询资金账户响应"""
        pass

    @abstractmethod
    def OnRspQrySecAgentCheckMode(self, pSecAgentCheckMode: ctypes.POINTER(CThostFtdcSecAgentCheckModeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询二级代理商资金校验模式响应"""
        pass

    @abstractmethod
    def OnRspQrySecAgentTradeInfo(self, pSecAgentTradeInfo: ctypes.POINTER(CThostFtdcSecAgentTradeInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询二级代理商信息响应"""
        pass

    @abstractmethod
    def OnRspQryOptionInstrTradeCost(self, pOptionInstrTradeCost: ctypes.POINTER(CThostFtdcOptionInstrTradeCostField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询期权交易成本响应"""
        pass

    @abstractmethod
    def OnRspQryOptionInstrCommRate(self, pOptionInstrCommRate: ctypes.POINTER(CThostFtdcOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询期权合约手续费响应"""
        pass

    @abstractmethod
    def OnRspQryExecOrder(self, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询执行宣告响应"""
        pass

    @abstractmethod
    def OnRspQryForQuote(self, pForQuote: ctypes.POINTER(CThostFtdcForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询询价响应"""
        pass

    @abstractmethod
    def OnRspQryQuote(self, pQuote: ctypes.POINTER(CThostFtdcQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询报价响应"""
        pass

    @abstractmethod
    def OnRspQryOptionSelfClose(self, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询期权自对冲响应"""
        pass

    @abstractmethod
    def OnRspQryInvestUnit(self, pInvestUnit: ctypes.POINTER(CThostFtdcInvestUnitField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询投资单元响应"""
        pass

    @abstractmethod
    def OnRspQryCombInstrumentGuard(self, pCombInstrumentGuard: ctypes.POINTER(CThostFtdcCombInstrumentGuardField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询组合合约安全系数响应"""
        pass

    @abstractmethod
    def OnRspQryCombAction(self, pCombAction: ctypes.POINTER(CThostFtdcCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询申请组合响应"""
        pass

    @abstractmethod
    def OnRspQryTransferSerial(self, pTransferSerial: ctypes.POINTER(CThostFtdcTransferSerialField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询转帐流水响应"""
        pass

    @abstractmethod
    def OnRspQryAccountregister(self, pAccountregister: ctypes.POINTER(CThostFtdcAccountregisterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询银期签约关系响应"""
        pass

    @abstractmethod
    def OnRspError(self, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """错误应答"""
        pass

    @abstractmethod
    def OnRtnOrder(self, pOrder: ctypes.POINTER(CThostFtdcOrderField)):
        """报单通知"""
        pass

    @abstractmethod
    def OnRtnTrade(self, pTrade: ctypes.POINTER(CThostFtdcTradeField)):
        """成交通知"""
        pass

    @abstractmethod
    def OnErrRtnOrderInsert(self, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """报单录入错误回报"""
        pass

    @abstractmethod
    def OnErrRtnOrderAction(self, pOrderAction: ctypes.POINTER(CThostFtdcOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """报单操作错误回报"""
        pass

    @abstractmethod
    def OnRtnInstrumentStatus(self, pInstrumentStatus: ctypes.POINTER(CThostFtdcInstrumentStatusField)):
        """合约交易状态通知"""
        pass

    @abstractmethod
    def OnRtnBulletin(self, pBulletin: ctypes.POINTER(CThostFtdcBulletinField)):
        """交易所公告通知"""
        pass

    @abstractmethod
    def OnRtnTradingNotice(self, pTradingNoticeInfo: ctypes.POINTER(CThostFtdcTradingNoticeInfoField)):
        """交易通知"""
        pass

    @abstractmethod
    def OnRtnErrorConditionalOrder(self, pErrorConditionalOrder: ctypes.POINTER(CThostFtdcErrorConditionalOrderField)):
        """提示条件单校验错误"""
        pass

    @abstractmethod
    def OnRtnExecOrder(self, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField)):
        """执行宣告通知"""
        pass

    @abstractmethod
    def OnErrRtnExecOrderInsert(self, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """执行宣告录入错误回报"""
        pass

    @abstractmethod
    def OnErrRtnExecOrderAction(self, pExecOrderAction: ctypes.POINTER(CThostFtdcExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """执行宣告操作错误回报"""
        pass

    @abstractmethod
    def OnErrRtnForQuoteInsert(self, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """询价录入错误回报"""
        pass

    @abstractmethod
    def OnRtnQuote(self, pQuote: ctypes.POINTER(CThostFtdcQuoteField)):
        """报价通知"""
        pass

    @abstractmethod
    def OnErrRtnQuoteInsert(self, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """报价录入错误回报"""
        pass

    @abstractmethod
    def OnErrRtnQuoteAction(self, pQuoteAction: ctypes.POINTER(CThostFtdcQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """报价操作错误回报"""
        pass

    @abstractmethod
    def OnRtnForQuoteRsp(self, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
        """询价通知"""
        pass

    @abstractmethod
    def OnRtnCFMMCTradingAccountToken(self, pCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcCFMMCTradingAccountTokenField)):
        """保证金监控中心用户令牌"""
        pass

    @abstractmethod
    def OnErrRtnBatchOrderAction(self, pBatchOrderAction: ctypes.POINTER(CThostFtdcBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """批量报单操作错误回报"""
        pass

    @abstractmethod
    def OnRtnOptionSelfClose(self, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField)):
        """期权自对冲通知"""
        pass

    @abstractmethod
    def OnErrRtnOptionSelfCloseInsert(self, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """期权自对冲录入错误回报"""
        pass

    @abstractmethod
    def OnErrRtnOptionSelfCloseAction(self, pOptionSelfCloseAction: ctypes.POINTER(CThostFtdcOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """期权自对冲操作错误回报"""
        pass

    @abstractmethod
    def OnRtnCombAction(self, pCombAction: ctypes.POINTER(CThostFtdcCombActionField)):
        """申请组合通知"""
        pass

    @abstractmethod
    def OnErrRtnCombActionInsert(self, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """申请组合录入错误回报"""
        pass

    @abstractmethod
    def OnRspQryContractBank(self, pContractBank: ctypes.POINTER(CThostFtdcContractBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询签约银行响应"""
        pass

    @abstractmethod
    def OnRspQryParkedOrder(self, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询预埋单响应"""
        pass

    @abstractmethod
    def OnRspQryParkedOrderAction(self, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询预埋撤单响应"""
        pass

    @abstractmethod
    def OnRspQryTradingNotice(self, pTradingNotice: ctypes.POINTER(CThostFtdcTradingNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询交易通知响应"""
        pass

    @abstractmethod
    def OnRspQryBrokerTradingParams(self, pBrokerTradingParams: ctypes.POINTER(CThostFtdcBrokerTradingParamsField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询经纪公司交易参数响应"""
        pass

    @abstractmethod
    def OnRspQryBrokerTradingAlgos(self, pBrokerTradingAlgos: ctypes.POINTER(CThostFtdcBrokerTradingAlgosField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询经纪公司交易算法响应"""
        pass

    @abstractmethod
    def OnRspQueryCFMMCTradingAccountToken(self, pQueryCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询监控中心用户令牌"""
        pass

    @abstractmethod
    def OnRtnFromBankToFutureByBank(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """银行发起银行资金转期货通知"""
        pass

    @abstractmethod
    def OnRtnFromFutureToBankByBank(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """银行发起期货资金转银行通知"""
        pass

    @abstractmethod
    def OnRtnRepealFromBankToFutureByBank(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """银行发起冲正银行转期货通知"""
        pass

    @abstractmethod
    def OnRtnRepealFromFutureToBankByBank(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """银行发起冲正期货转银行通知"""
        pass

    @abstractmethod
    def OnRtnFromBankToFutureByFuture(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """期货发起银行资金转期货通知"""
        pass

    @abstractmethod
    def OnRtnFromFutureToBankByFuture(self, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
        """期货发起期货资金转银行通知"""
        pass

    @abstractmethod
    def OnRtnRepealFromBankToFutureByFutureManual(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知"""
        pass

    @abstractmethod
    def OnRtnRepealFromFutureToBankByFutureManual(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知"""
        pass

    @abstractmethod
    def OnRtnQueryBankBalanceByFuture(self, pNotifyQueryAccount: ctypes.POINTER(CThostFtdcNotifyQueryAccountField)):
        """期货发起查询银行余额通知"""
        pass

    @abstractmethod
    def OnErrRtnBankToFutureByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """期货发起银行资金转期货错误回报"""
        pass

    @abstractmethod
    def OnErrRtnFutureToBankByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """期货发起期货资金转银行错误回报"""
        pass

    @abstractmethod
    def OnErrRtnRepealBankToFutureByFutureManual(self, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """系统运行时期货端手工发起冲正银行转期货错误回报"""
        pass

    @abstractmethod
    def OnErrRtnRepealFutureToBankByFutureManual(self, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """系统运行时期货端手工发起冲正期货转银行错误回报"""
        pass

    @abstractmethod
    def OnErrRtnQueryBankBalanceByFuture(self, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
        """期货发起查询银行余额错误回报"""
        pass

    @abstractmethod
    def OnRtnRepealFromBankToFutureByFuture(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知"""
        pass

    @abstractmethod
    def OnRtnRepealFromFutureToBankByFuture(self, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
        """期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知"""
        pass

    @abstractmethod
    def OnRspFromBankToFutureByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """期货发起银行资金转期货应答"""
        pass

    @abstractmethod
    def OnRspFromFutureToBankByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """期货发起期货资金转银行应答"""
        pass

    @abstractmethod
    def OnRspQueryBankAccountMoneyByFuture(self, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """期货发起查询银行余额应答"""
        pass

    @abstractmethod
    def OnRtnOpenAccountByBank(self, pOpenAccount: ctypes.POINTER(CThostFtdcOpenAccountField)):
        """银行发起银期开户通知"""
        pass

    @abstractmethod
    def OnRtnCancelAccountByBank(self, pCancelAccount: ctypes.POINTER(CThostFtdcCancelAccountField)):
        """银行发起银期销户通知"""
        pass

    @abstractmethod
    def OnRtnChangeAccountByBank(self, pChangeAccount: ctypes.POINTER(CThostFtdcChangeAccountField)):
        """银行发起变更银行账号通知"""
        pass

    @abstractmethod
    def OnRspQryClassifiedInstrument(self, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求查询分类合约响应"""
        pass

    @abstractmethod
    def OnRspQryCombPromotionParam(self, pCombPromotionParam: ctypes.POINTER(CThostFtdcCombPromotionParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """请求组合优惠比例响应"""
        pass

    @abstractmethod
    def OnRspQryRiskSettleInvstPosition(self, pRiskSettleInvstPosition: ctypes.POINTER(CThostFtdcRiskSettleInvstPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者风险结算持仓查询响应"""
        pass

    @abstractmethod
    def OnRspQryRiskSettleProductStatus(self, pRiskSettleProductStatus: ctypes.POINTER(CThostFtdcRiskSettleProductStatusField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """风险结算产品查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMFutureParameter(self, pSPBMFutureParameter: ctypes.POINTER(CThostFtdcSPBMFutureParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM期货合约参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMOptionParameter(self, pSPBMOptionParameter: ctypes.POINTER(CThostFtdcSPBMOptionParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM期权合约参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMIntraParameter(self, pSPBMIntraParameter: ctypes.POINTER(CThostFtdcSPBMIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM品种内对锁仓折扣参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMInterParameter(self, pSPBMInterParameter: ctypes.POINTER(CThostFtdcSPBMInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM跨品种抵扣参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMPortfDefinition(self, pSPBMPortfDefinition: ctypes.POINTER(CThostFtdcSPBMPortfDefinitionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM组合保证金套餐查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMInvestorPortfDef(self, pSPBMInvestorPortfDef: ctypes.POINTER(CThostFtdcSPBMInvestorPortfDefField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者SPBM套餐选择查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorPortfMarginRatio(self, pInvestorPortfMarginRatio: ctypes.POINTER(CThostFtdcInvestorPortfMarginRatioField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者新型组合保证金系数查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorProdSPBMDetail(self, pInvestorProdSPBMDetail: ctypes.POINTER(CThostFtdcInvestorProdSPBMDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者产品SPBM明细查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorCommoditySPMMMargin(self, pInvestorCommoditySPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommoditySPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者商品组SPMM记录查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorCommodityGroupSPMMMargin(self, pInvestorCommodityGroupSPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommodityGroupSPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者商品群SPMM记录查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPMMInstParam(self, pSPMMInstParam: ctypes.POINTER(CThostFtdcSPMMInstParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPMM合约参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPMMProductParam(self, pSPMMProductParam: ctypes.POINTER(CThostFtdcSPMMProductParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPMM产品参数查询响应"""
        pass

    @abstractmethod
    def OnRspQrySPBMAddOnInterParameter(self, pSPBMAddOnInterParameter: ctypes.POINTER(CThostFtdcSPBMAddOnInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """SPBM附加跨品种抵扣参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSCombProductInfo(self, pRCAMSCombProductInfo: ctypes.POINTER(CThostFtdcRCAMSCombProductInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS产品组合信息查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSInstrParameter(self, pRCAMSInstrParameter: ctypes.POINTER(CThostFtdcRCAMSInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS同合约风险对冲参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSIntraParameter(self, pRCAMSIntraParameter: ctypes.POINTER(CThostFtdcRCAMSIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS品种内风险对冲参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSInterParameter(self, pRCAMSInterParameter: ctypes.POINTER(CThostFtdcRCAMSInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS跨品种风险折抵参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSShortOptAdjustParam(self, pRCAMSShortOptAdjustParam: ctypes.POINTER(CThostFtdcRCAMSShortOptAdjustParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS空头期权风险调整参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRCAMSInvestorCombPosition(self, pRCAMSInvestorCombPosition: ctypes.POINTER(CThostFtdcRCAMSInvestorCombPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RCAMS策略组合持仓查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorProdRCAMSMargin(self, pInvestorProdRCAMSMargin: ctypes.POINTER(CThostFtdcInvestorProdRCAMSMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者品种RCAMS保证金查询响应"""
        pass

    @abstractmethod
    def OnRspQryRULEInstrParameter(self, pRULEInstrParameter: ctypes.POINTER(CThostFtdcRULEInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RULE合约保证金参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRULEIntraParameter(self, pRULEIntraParameter: ctypes.POINTER(CThostFtdcRULEIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RULE品种内对锁仓折扣参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryRULEInterParameter(self, pRULEInterParameter: ctypes.POINTER(CThostFtdcRULEInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """RULE跨品种抵扣参数查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorProdRULEMargin(self, pInvestorProdRULEMargin: ctypes.POINTER(CThostFtdcInvestorProdRULEMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者产品RULE保证金查询响应"""
        pass

    @abstractmethod
    def OnRspQryInvestorPortfSetting(self, pInvestorPortfSetting: ctypes.POINTER(CThostFtdcInvestorPortfSettingField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """投资者投资者新组保设置查询响应"""
        pass

# ========== TraderApi 类 ==========

class TraderApi:
    """交易 API 封装"""

    def __init__(self, flow_path: str):
        """创建交易 API 实例"""
        # 自动加载库（如果尚未加载）
        auto_load_library()

        self._handle: Optional[ctypes.c_void_p] = None
        self._spi: Optional[TraderSpi] = None
        self._spi_handle: Optional[ctypes.c_void_p] = None
        self._user_data: int = _register_trader_instance(self)
        self._lock = threading.RLock()

        # 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录
        # CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的
        # 所以转换为绝对路径更可靠
        abs_flow_path = flow_path
        if not os.path.isabs(flow_path):
            # 如果是相对路径，转换为基于当前工作目录的绝对路径
            abs_flow_path = os.path.abspath(flow_path)

        # 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）
        if abs_flow_path and not abs_flow_path.endswith(os.sep):
            abs_flow_path += os.sep

        # 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收
        # CTP API 可能会在后续使用这个路径
        self._flow_path = abs_flow_path.encode('utf-8') + b'\0'

        # 调用 C 函数创建 API
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")

        # 获取函数指针
        func = lib.TraderCreateFtdcTraderApi
        func.argtypes = [ctypes.c_char_p]
        func.restype = ctypes.c_void_p

        self._handle = func(self._flow_path)
        if self._handle is None:
            raise RuntimeError("Failed to create TraderApi")

    # ========== API 方法 ==========

    # GetApiVersion 获取API的版本信息
    def GetApiVersion(self) -> str:
        """获取 API 版本"""
        lib = get_trader_lib_handle()
        if lib is None:
            return ""
        func = lib.TraderGetApiVersion
        func.argtypes = []
        func.restype = ctypes.c_char_p
        ptr = func()
        return go_string(ptr) if ptr else ""

    # Release 删除接口对象本身
    def Release(self):
        """释放 API 实例"""
        with self._lock:
            if self._handle:
                lib = get_trader_lib_handle()
                if lib:
                    func = lib.TraderRelease
                    func.argtypes = [ctypes.c_void_p]
                    func.restype = None
                    func(self._handle)
                self._handle = None
            _unregister_trader_instance(self._user_data)

    # Init 初始化
    def Init(self, ):
        """初始化"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderInit
        func.argtypes = [ctypes.c_void_p]
        func.restype = None
        func(self._handle)

    # Join 等待接口线程结束运行
    def Join(self, ) -> ctypes.c_int32:
        """等待接口线程结束运行"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderJoin
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_int32
        return func(self._handle)

    # GetTradingDay 获取当前交易日
    def GetTradingDay(self) -> str:
        """获取交易日"""
        lib = get_trader_lib_handle()
        if lib is None:
            return ""
        func = lib.TraderGetTradingDay
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_char_p
        ptr = func(self._handle)
        return go_string(ptr) if ptr else ""

    # GetFrontInfo 获取已连接的前置的信息
    def GetFrontInfo(self, pFrontInfo: ctypes.POINTER(CThostFtdcFrontInfoField)):
        """获取已连接的前置的信息"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderGetFrontInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcFrontInfoField)]
        func.restype = None
        func(self._handle, pFrontInfo)

    # RegisterFront 注册前置机网络地址
    def RegisterFront(self, pszFrontAddress: str):
        """注册前置机网络地址"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderRegisterFront
        func.argtypes = [ctypes.c_void_p, ctypes.c_char_p]
        func.restype = None
        func(self._handle, c_string(pszFrontAddress))

    # RegisterNameServer 注册名字服务器网络地址
    def RegisterNameServer(self, pszNsAddress: str):
        """注册名字服务器网络地址"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderRegisterNameServer
        func.argtypes = [ctypes.c_void_p, ctypes.c_char_p]
        func.restype = None
        func(self._handle, c_string(pszNsAddress))

    # RegisterFensUserInfo 注册名字服务器用户信息
    def RegisterFensUserInfo(self, pFensUserInfo: ctypes.POINTER(CThostFtdcFensUserInfoField)):
        """注册名字服务器用户信息"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderRegisterFensUserInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcFensUserInfoField)]
        func.restype = None
        func(self._handle, pFensUserInfo)

    # SubscribePrivateTopic 订阅私有流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后私有流的内容
    def SubscribePrivateTopic(self, nResumeType: THOST_TE_RESUME_TYPE):
        """订阅私有流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后私有流的内容"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSubscribePrivateTopic
        func.argtypes = [ctypes.c_void_p, THOST_TE_RESUME_TYPE]
        func.restype = None
        func(self._handle, nResumeType)

    # SubscribePublicTopic 订阅公共流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后公共流的内容 THOST_TERT_NONE:取消订阅公共流
    def SubscribePublicTopic(self, nResumeType: THOST_TE_RESUME_TYPE):
        """订阅公共流。 THOST_TERT_RESTART:从本交易日开始重传 THOST_TERT_RESUME:从上次收到的续传 THOST_TERT_QUICK:只传送登录后公共流的内容 THOST_TERT_NONE:取消订阅公共流"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSubscribePublicTopic
        func.argtypes = [ctypes.c_void_p, THOST_TE_RESUME_TYPE]
        func.restype = None
        func(self._handle, nResumeType)

    # ReqAuthenticate 客户端认证请求
    def ReqAuthenticate(self, pReqAuthenticateField: ctypes.POINTER(CThostFtdcReqAuthenticateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """客户端认证请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqAuthenticate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqAuthenticateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqAuthenticateField, nRequestID)

    # RegisterUserSystemInfo 注册用户终端信息，用于中继服务器多连接模式 需要在终端认证成功后，用户登录前调用该接口
    def RegisterUserSystemInfo(self, pUserSystemInfo: ctypes.POINTER(CThostFtdcUserSystemInfoField)) -> ctypes.c_int32:
        """注册用户终端信息，用于中继服务器多连接模式 需要在终端认证成功后，用户登录前调用该接口"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderRegisterUserSystemInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserSystemInfoField)]
        func.restype = ctypes.c_int32
        return func(self._handle, pUserSystemInfo)

    # SubmitUserSystemInfo 上报用户终端信息，用于中继服务器操作员登录模式 操作员登录后，可以多次调用该接口上报客户信息
    def SubmitUserSystemInfo(self, pUserSystemInfo: ctypes.POINTER(CThostFtdcUserSystemInfoField)) -> ctypes.c_int32:
        """上报用户终端信息，用于中继服务器操作员登录模式 操作员登录后，可以多次调用该接口上报客户信息"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSubmitUserSystemInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserSystemInfoField)]
        func.restype = ctypes.c_int32
        return func(self._handle, pUserSystemInfo)

    # ReqUserLogin 用户登录请求
    def ReqUserLogin(self, pReqUserLoginField: ctypes.POINTER(CThostFtdcReqUserLoginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户登录请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLogin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginField, nRequestID)

    # ReqUserLogout 登出请求
    def ReqUserLogout(self, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """登出请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLogout
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pUserLogout, nRequestID)

    # ReqUserPasswordUpdate 用户口令更新请求
    def ReqUserPasswordUpdate(self, pUserPasswordUpdate: ctypes.POINTER(CThostFtdcUserPasswordUpdateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户口令更新请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserPasswordUpdate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserPasswordUpdateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pUserPasswordUpdate, nRequestID)

    # ReqTradingAccountPasswordUpdate 资金账户口令更新请求
    def ReqTradingAccountPasswordUpdate(self, pTradingAccountPasswordUpdate: ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """资金账户口令更新请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqTradingAccountPasswordUpdate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pTradingAccountPasswordUpdate, nRequestID)

    # ReqUserAuthMethod 查询用户当前支持的认证模式
    def ReqUserAuthMethod(self, pReqUserAuthMethod: ctypes.POINTER(CThostFtdcReqUserAuthMethodField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """查询用户当前支持的认证模式"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserAuthMethod
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserAuthMethodField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserAuthMethod, nRequestID)

    # ReqGenUserCaptcha 用户发出获取图形验证码请求
    def ReqGenUserCaptcha(self, pReqGenUserCaptcha: ctypes.POINTER(CThostFtdcReqGenUserCaptchaField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户发出获取图形验证码请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqGenUserCaptcha
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqGenUserCaptchaField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqGenUserCaptcha, nRequestID)

    # ReqGenUserText 用户发出获取短信验证码请求
    def ReqGenUserText(self, pReqGenUserText: ctypes.POINTER(CThostFtdcReqGenUserTextField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户发出获取短信验证码请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqGenUserText
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqGenUserTextField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqGenUserText, nRequestID)

    # ReqUserLoginWithCaptcha 用户发出带有图片验证码的登陆请求
    def ReqUserLoginWithCaptcha(self, pReqUserLoginWithCaptcha: ctypes.POINTER(CThostFtdcReqUserLoginWithCaptchaField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户发出带有图片验证码的登陆请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLoginWithCaptcha
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginWithCaptchaField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginWithCaptcha, nRequestID)

    # ReqUserLoginWithText 用户发出带有短信验证码的登陆请求
    def ReqUserLoginWithText(self, pReqUserLoginWithText: ctypes.POINTER(CThostFtdcReqUserLoginWithTextField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户发出带有短信验证码的登陆请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLoginWithText
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginWithTextField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginWithText, nRequestID)

    # ReqUserLoginWithOTP 用户发出带有动态口令的登陆请求
    def ReqUserLoginWithOTP(self, pReqUserLoginWithOTP: ctypes.POINTER(CThostFtdcReqUserLoginWithOTPField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """用户发出带有动态口令的登陆请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLoginWithOTP
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginWithOTPField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginWithOTP, nRequestID)

    # ReqOrderInsert 报单录入请求
    def ReqOrderInsert(self, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """报单录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqOrderInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputOrder, nRequestID)

    # ReqParkedOrderInsert 预埋单录入请求
    def ReqParkedOrderInsert(self, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """预埋单录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqParkedOrderInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pParkedOrder, nRequestID)

    # ReqParkedOrderAction 预埋撤单录入请求
    def ReqParkedOrderAction(self, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """预埋撤单录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqParkedOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pParkedOrderAction, nRequestID)

    # ReqOrderAction 报单操作请求
    def ReqOrderAction(self, pInputOrderAction: ctypes.POINTER(CThostFtdcInputOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """报单操作请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputOrderAction, nRequestID)

    # ReqQryMaxOrderVolume 查询最大报单数量请求
    def ReqQryMaxOrderVolume(self, pQryMaxOrderVolume: ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """查询最大报单数量请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryMaxOrderVolume
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryMaxOrderVolume, nRequestID)

    # ReqSettlementInfoConfirm 投资者结算结果确认
    def ReqSettlementInfoConfirm(self, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者结算结果确认"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqSettlementInfoConfirm
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pSettlementInfoConfirm, nRequestID)

    # ReqRemoveParkedOrder 请求删除预埋单
    def ReqRemoveParkedOrder(self, pRemoveParkedOrder: ctypes.POINTER(CThostFtdcRemoveParkedOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求删除预埋单"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqRemoveParkedOrder
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pRemoveParkedOrder, nRequestID)

    # ReqRemoveParkedOrderAction 请求删除预埋撤单
    def ReqRemoveParkedOrderAction(self, pRemoveParkedOrderAction: ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求删除预埋撤单"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqRemoveParkedOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pRemoveParkedOrderAction, nRequestID)

    # ReqExecOrderInsert 执行宣告录入请求
    def ReqExecOrderInsert(self, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """执行宣告录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqExecOrderInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputExecOrder, nRequestID)

    # ReqExecOrderAction 执行宣告操作请求
    def ReqExecOrderAction(self, pInputExecOrderAction: ctypes.POINTER(CThostFtdcInputExecOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """执行宣告操作请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqExecOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputExecOrderAction, nRequestID)

    # ReqForQuoteInsert 询价录入请求
    def ReqForQuoteInsert(self, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """询价录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqForQuoteInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputForQuoteField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputForQuote, nRequestID)

    # ReqQuoteInsert 报价录入请求
    def ReqQuoteInsert(self, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """报价录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQuoteInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputQuote, nRequestID)

    # ReqQuoteAction 报价操作请求
    def ReqQuoteAction(self, pInputQuoteAction: ctypes.POINTER(CThostFtdcInputQuoteActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """报价操作请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQuoteAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputQuoteAction, nRequestID)

    # ReqBatchOrderAction 批量报单操作请求
    def ReqBatchOrderAction(self, pInputBatchOrderAction: ctypes.POINTER(CThostFtdcInputBatchOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """批量报单操作请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqBatchOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputBatchOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputBatchOrderAction, nRequestID)

    # ReqOptionSelfCloseInsert 期权自对冲录入请求
    def ReqOptionSelfCloseInsert(self, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """期权自对冲录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqOptionSelfCloseInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputOptionSelfClose, nRequestID)

    # ReqOptionSelfCloseAction 期权自对冲操作请求
    def ReqOptionSelfCloseAction(self, pInputOptionSelfCloseAction: ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """期权自对冲操作请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqOptionSelfCloseAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputOptionSelfCloseAction, nRequestID)

    # ReqCombActionInsert 申请组合录入请求
    def ReqCombActionInsert(self, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """申请组合录入请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqCombActionInsert
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputCombActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pInputCombAction, nRequestID)

    # ReqQryOrder 请求查询报单
    def ReqQryOrder(self, pQryOrder: ctypes.POINTER(CThostFtdcQryOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询报单"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryOrder
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryOrder, nRequestID)

    # ReqQryTrade 请求查询成交
    def ReqQryTrade(self, pQryTrade: ctypes.POINTER(CThostFtdcQryTradeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询成交"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTrade
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTradeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTrade, nRequestID)

    # ReqQryInvestorPosition 请求查询投资者持仓
    def ReqQryInvestorPosition(self, pQryInvestorPosition: ctypes.POINTER(CThostFtdcQryInvestorPositionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者持仓"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorPosition
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorPositionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorPosition, nRequestID)

    # ReqQryTradingAccount 请求查询资金账户
    def ReqQryTradingAccount(self, pQryTradingAccount: ctypes.POINTER(CThostFtdcQryTradingAccountField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询资金账户"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTradingAccount
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTradingAccountField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTradingAccount, nRequestID)

    # ReqQryInvestor 请求查询投资者
    def ReqQryInvestor(self, pQryInvestor: ctypes.POINTER(CThostFtdcQryInvestorField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestor
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestor, nRequestID)

    # ReqQryTradingCode 请求查询交易编码
    def ReqQryTradingCode(self, pQryTradingCode: ctypes.POINTER(CThostFtdcQryTradingCodeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易编码"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTradingCode
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTradingCodeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTradingCode, nRequestID)

    # ReqQryInstrumentMarginRate 请求查询合约保证金率
    def ReqQryInstrumentMarginRate(self, pQryInstrumentMarginRate: ctypes.POINTER(CThostFtdcQryInstrumentMarginRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询合约保证金率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInstrumentMarginRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInstrumentMarginRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInstrumentMarginRate, nRequestID)

    # ReqQryInstrumentCommissionRate 请求查询合约手续费率
    def ReqQryInstrumentCommissionRate(self, pQryInstrumentCommissionRate: ctypes.POINTER(CThostFtdcQryInstrumentCommissionRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询合约手续费率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInstrumentCommissionRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInstrumentCommissionRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInstrumentCommissionRate, nRequestID)

    # ReqQryExchange 请求查询交易所
    def ReqQryExchange(self, pQryExchange: ctypes.POINTER(CThostFtdcQryExchangeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易所"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryExchange
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryExchangeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryExchange, nRequestID)

    # ReqQryProduct 请求查询产品
    def ReqQryProduct(self, pQryProduct: ctypes.POINTER(CThostFtdcQryProductField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询产品"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryProduct
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryProductField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryProduct, nRequestID)

    # ReqQryInstrument 请求查询合约
    def ReqQryInstrument(self, pQryInstrument: ctypes.POINTER(CThostFtdcQryInstrumentField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询合约"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInstrument
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInstrumentField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInstrument, nRequestID)

    # ReqQryDepthMarketData 请求查询行情
    def ReqQryDepthMarketData(self, pQryDepthMarketData: ctypes.POINTER(CThostFtdcQryDepthMarketDataField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询行情"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryDepthMarketData
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryDepthMarketDataField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryDepthMarketData, nRequestID)

    # ReqQryTraderOffer 请求查询交易员报盘机
    def ReqQryTraderOffer(self, pQryTraderOffer: ctypes.POINTER(CThostFtdcQryTraderOfferField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易员报盘机"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTraderOffer
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTraderOfferField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTraderOffer, nRequestID)

    # ReqQrySettlementInfo 请求查询投资者结算结果
    def ReqQrySettlementInfo(self, pQrySettlementInfo: ctypes.POINTER(CThostFtdcQrySettlementInfoField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者结算结果"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySettlementInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySettlementInfoField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySettlementInfo, nRequestID)

    # ReqQryTransferBank 请求查询转帐银行
    def ReqQryTransferBank(self, pQryTransferBank: ctypes.POINTER(CThostFtdcQryTransferBankField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询转帐银行"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTransferBank
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTransferBankField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTransferBank, nRequestID)

    # ReqQryInvestorPositionDetail 请求查询投资者持仓明细
    def ReqQryInvestorPositionDetail(self, pQryInvestorPositionDetail: ctypes.POINTER(CThostFtdcQryInvestorPositionDetailField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者持仓明细"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorPositionDetail
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorPositionDetailField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorPositionDetail, nRequestID)

    # ReqQryNotice 请求查询客户通知
    def ReqQryNotice(self, pQryNotice: ctypes.POINTER(CThostFtdcQryNoticeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询客户通知"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryNotice
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryNoticeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryNotice, nRequestID)

    # ReqQrySettlementInfoConfirm 请求查询结算信息确认
    def ReqQrySettlementInfoConfirm(self, pQrySettlementInfoConfirm: ctypes.POINTER(CThostFtdcQrySettlementInfoConfirmField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询结算信息确认"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySettlementInfoConfirm
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySettlementInfoConfirmField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySettlementInfoConfirm, nRequestID)

    # ReqQryInvestorPositionCombineDetail 请求查询投资者持仓明细
    def ReqQryInvestorPositionCombineDetail(self, pQryInvestorPositionCombineDetail: ctypes.POINTER(CThostFtdcQryInvestorPositionCombineDetailField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者持仓明细"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorPositionCombineDetail
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorPositionCombineDetailField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorPositionCombineDetail, nRequestID)

    # ReqQryCFMMCTradingAccountKey 请求查询保证金监管系统经纪公司资金账户密钥
    def ReqQryCFMMCTradingAccountKey(self, pQryCFMMCTradingAccountKey: ctypes.POINTER(CThostFtdcQryCFMMCTradingAccountKeyField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询保证金监管系统经纪公司资金账户密钥"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryCFMMCTradingAccountKey
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryCFMMCTradingAccountKeyField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryCFMMCTradingAccountKey, nRequestID)

    # ReqQryEWarrantOffset 请求查询仓单折抵信息
    def ReqQryEWarrantOffset(self, pQryEWarrantOffset: ctypes.POINTER(CThostFtdcQryEWarrantOffsetField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询仓单折抵信息"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryEWarrantOffset
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryEWarrantOffsetField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryEWarrantOffset, nRequestID)

    # ReqQryInvestorProductGroupMargin 请求查询投资者品种/跨品种保证金
    def ReqQryInvestorProductGroupMargin(self, pQryInvestorProductGroupMargin: ctypes.POINTER(CThostFtdcQryInvestorProductGroupMarginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资者品种/跨品种保证金"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorProductGroupMargin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorProductGroupMarginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorProductGroupMargin, nRequestID)

    # ReqQryExchangeMarginRate 请求查询交易所保证金率
    def ReqQryExchangeMarginRate(self, pQryExchangeMarginRate: ctypes.POINTER(CThostFtdcQryExchangeMarginRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易所保证金率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryExchangeMarginRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryExchangeMarginRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryExchangeMarginRate, nRequestID)

    # ReqQryExchangeMarginRateAdjust 请求查询交易所调整保证金率
    def ReqQryExchangeMarginRateAdjust(self, pQryExchangeMarginRateAdjust: ctypes.POINTER(CThostFtdcQryExchangeMarginRateAdjustField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易所调整保证金率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryExchangeMarginRateAdjust
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryExchangeMarginRateAdjustField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryExchangeMarginRateAdjust, nRequestID)

    # ReqQryExchangeRate 请求查询汇率
    def ReqQryExchangeRate(self, pQryExchangeRate: ctypes.POINTER(CThostFtdcQryExchangeRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询汇率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryExchangeRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryExchangeRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryExchangeRate, nRequestID)

    # ReqQrySecAgentACIDMap 请求查询二级代理操作员银期权限
    def ReqQrySecAgentACIDMap(self, pQrySecAgentACIDMap: ctypes.POINTER(CThostFtdcQrySecAgentACIDMapField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询二级代理操作员银期权限"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySecAgentACIDMap
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySecAgentACIDMapField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySecAgentACIDMap, nRequestID)

    # ReqQryProductExchRate 请求查询产品报价汇率
    def ReqQryProductExchRate(self, pQryProductExchRate: ctypes.POINTER(CThostFtdcQryProductExchRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询产品报价汇率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryProductExchRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryProductExchRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryProductExchRate, nRequestID)

    # ReqQryProductGroup 请求查询产品组
    def ReqQryProductGroup(self, pQryProductGroup: ctypes.POINTER(CThostFtdcQryProductGroupField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询产品组"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryProductGroup
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryProductGroupField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryProductGroup, nRequestID)

    # ReqQryMMInstrumentCommissionRate 请求查询做市商合约手续费率
    def ReqQryMMInstrumentCommissionRate(self, pQryMMInstrumentCommissionRate: ctypes.POINTER(CThostFtdcQryMMInstrumentCommissionRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询做市商合约手续费率"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryMMInstrumentCommissionRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMMInstrumentCommissionRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryMMInstrumentCommissionRate, nRequestID)

    # ReqQryMMOptionInstrCommRate 请求查询做市商期权合约手续费
    def ReqQryMMOptionInstrCommRate(self, pQryMMOptionInstrCommRate: ctypes.POINTER(CThostFtdcQryMMOptionInstrCommRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询做市商期权合约手续费"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryMMOptionInstrCommRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMMOptionInstrCommRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryMMOptionInstrCommRate, nRequestID)

    # ReqQryInstrumentOrderCommRate 请求查询报单手续费
    def ReqQryInstrumentOrderCommRate(self, pQryInstrumentOrderCommRate: ctypes.POINTER(CThostFtdcQryInstrumentOrderCommRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询报单手续费"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInstrumentOrderCommRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInstrumentOrderCommRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInstrumentOrderCommRate, nRequestID)

    # ReqQrySecAgentTradingAccount 请求查询资金账户
    def ReqQrySecAgentTradingAccount(self, pQryTradingAccount: ctypes.POINTER(CThostFtdcQryTradingAccountField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询资金账户"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySecAgentTradingAccount
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTradingAccountField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTradingAccount, nRequestID)

    # ReqQrySecAgentCheckMode 请求查询二级代理商资金校验模式
    def ReqQrySecAgentCheckMode(self, pQrySecAgentCheckMode: ctypes.POINTER(CThostFtdcQrySecAgentCheckModeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询二级代理商资金校验模式"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySecAgentCheckMode
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySecAgentCheckModeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySecAgentCheckMode, nRequestID)

    # ReqQrySecAgentTradeInfo 请求查询二级代理商信息
    def ReqQrySecAgentTradeInfo(self, pQrySecAgentTradeInfo: ctypes.POINTER(CThostFtdcQrySecAgentTradeInfoField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询二级代理商信息"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySecAgentTradeInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySecAgentTradeInfoField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySecAgentTradeInfo, nRequestID)

    # ReqQryOptionInstrTradeCost 请求查询期权交易成本
    def ReqQryOptionInstrTradeCost(self, pQryOptionInstrTradeCost: ctypes.POINTER(CThostFtdcQryOptionInstrTradeCostField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询期权交易成本"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryOptionInstrTradeCost
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryOptionInstrTradeCostField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryOptionInstrTradeCost, nRequestID)

    # ReqQryOptionInstrCommRate 请求查询期权合约手续费
    def ReqQryOptionInstrCommRate(self, pQryOptionInstrCommRate: ctypes.POINTER(CThostFtdcQryOptionInstrCommRateField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询期权合约手续费"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryOptionInstrCommRate
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryOptionInstrCommRateField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryOptionInstrCommRate, nRequestID)

    # ReqQryExecOrder 请求查询执行宣告
    def ReqQryExecOrder(self, pQryExecOrder: ctypes.POINTER(CThostFtdcQryExecOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询执行宣告"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryExecOrder
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryExecOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryExecOrder, nRequestID)

    # ReqQryForQuote 请求查询询价
    def ReqQryForQuote(self, pQryForQuote: ctypes.POINTER(CThostFtdcQryForQuoteField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询询价"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryForQuote
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryForQuoteField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryForQuote, nRequestID)

    # ReqQryQuote 请求查询报价
    def ReqQryQuote(self, pQryQuote: ctypes.POINTER(CThostFtdcQryQuoteField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询报价"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryQuote
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryQuoteField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryQuote, nRequestID)

    # ReqQryOptionSelfClose 请求查询期权自对冲
    def ReqQryOptionSelfClose(self, pQryOptionSelfClose: ctypes.POINTER(CThostFtdcQryOptionSelfCloseField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询期权自对冲"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryOptionSelfClose
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryOptionSelfCloseField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryOptionSelfClose, nRequestID)

    # ReqQryInvestUnit 请求查询投资单元
    def ReqQryInvestUnit(self, pQryInvestUnit: ctypes.POINTER(CThostFtdcQryInvestUnitField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询投资单元"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestUnit
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestUnitField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestUnit, nRequestID)

    # ReqQryCombInstrumentGuard 请求查询组合合约安全系数
    def ReqQryCombInstrumentGuard(self, pQryCombInstrumentGuard: ctypes.POINTER(CThostFtdcQryCombInstrumentGuardField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询组合合约安全系数"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryCombInstrumentGuard
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryCombInstrumentGuardField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryCombInstrumentGuard, nRequestID)

    # ReqQryCombAction 请求查询申请组合
    def ReqQryCombAction(self, pQryCombAction: ctypes.POINTER(CThostFtdcQryCombActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询申请组合"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryCombAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryCombActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryCombAction, nRequestID)

    # ReqQryTransferSerial 请求查询转帐流水
    def ReqQryTransferSerial(self, pQryTransferSerial: ctypes.POINTER(CThostFtdcQryTransferSerialField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询转帐流水"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTransferSerial
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTransferSerialField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTransferSerial, nRequestID)

    # ReqQryAccountregister 请求查询银期签约关系
    def ReqQryAccountregister(self, pQryAccountregister: ctypes.POINTER(CThostFtdcQryAccountregisterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询银期签约关系"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryAccountregister
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryAccountregisterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryAccountregister, nRequestID)

    # ReqQryContractBank 请求查询签约银行
    def ReqQryContractBank(self, pQryContractBank: ctypes.POINTER(CThostFtdcQryContractBankField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询签约银行"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryContractBank
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryContractBankField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryContractBank, nRequestID)

    # ReqQryParkedOrder 请求查询预埋单
    def ReqQryParkedOrder(self, pQryParkedOrder: ctypes.POINTER(CThostFtdcQryParkedOrderField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询预埋单"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryParkedOrder
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryParkedOrderField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryParkedOrder, nRequestID)

    # ReqQryParkedOrderAction 请求查询预埋撤单
    def ReqQryParkedOrderAction(self, pQryParkedOrderAction: ctypes.POINTER(CThostFtdcQryParkedOrderActionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询预埋撤单"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryParkedOrderAction
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryParkedOrderActionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryParkedOrderAction, nRequestID)

    # ReqQryTradingNotice 请求查询交易通知
    def ReqQryTradingNotice(self, pQryTradingNotice: ctypes.POINTER(CThostFtdcQryTradingNoticeField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询交易通知"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryTradingNotice
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryTradingNoticeField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryTradingNotice, nRequestID)

    # ReqQryBrokerTradingParams 请求查询经纪公司交易参数
    def ReqQryBrokerTradingParams(self, pQryBrokerTradingParams: ctypes.POINTER(CThostFtdcQryBrokerTradingParamsField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询经纪公司交易参数"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryBrokerTradingParams
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryBrokerTradingParamsField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryBrokerTradingParams, nRequestID)

    # ReqQryBrokerTradingAlgos 请求查询经纪公司交易算法
    def ReqQryBrokerTradingAlgos(self, pQryBrokerTradingAlgos: ctypes.POINTER(CThostFtdcQryBrokerTradingAlgosField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询经纪公司交易算法"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryBrokerTradingAlgos
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryBrokerTradingAlgosField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryBrokerTradingAlgos, nRequestID)

    # ReqQueryCFMMCTradingAccountToken 请求查询监控中心用户令牌
    def ReqQueryCFMMCTradingAccountToken(self, pQueryCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询监控中心用户令牌"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQueryCFMMCTradingAccountToken
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQueryCFMMCTradingAccountToken, nRequestID)

    # ReqFromBankToFutureByFuture 期货发起银行资金转期货请求
    def ReqFromBankToFutureByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """期货发起银行资金转期货请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqFromBankToFutureByFuture
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqTransfer, nRequestID)

    # ReqFromFutureToBankByFuture 期货发起期货资金转银行请求
    def ReqFromFutureToBankByFuture(self, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """期货发起期货资金转银行请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqFromFutureToBankByFuture
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqTransfer, nRequestID)

    # ReqQueryBankAccountMoneyByFuture 期货发起查询银行余额请求
    def ReqQueryBankAccountMoneyByFuture(self, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """期货发起查询银行余额请求"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQueryBankAccountMoneyByFuture
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqQueryAccountField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqQueryAccount, nRequestID)

    # ReqQryClassifiedInstrument 请求查询分类合约
    def ReqQryClassifiedInstrument(self, pQryClassifiedInstrument: ctypes.POINTER(CThostFtdcQryClassifiedInstrumentField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求查询分类合约"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryClassifiedInstrument
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryClassifiedInstrumentField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryClassifiedInstrument, nRequestID)

    # ReqQryCombPromotionParam 请求组合优惠比例
    def ReqQryCombPromotionParam(self, pQryCombPromotionParam: ctypes.POINTER(CThostFtdcQryCombPromotionParamField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """请求组合优惠比例"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryCombPromotionParam
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryCombPromotionParamField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryCombPromotionParam, nRequestID)

    # ReqQryRiskSettleInvstPosition 投资者风险结算持仓查询
    def ReqQryRiskSettleInvstPosition(self, pQryRiskSettleInvstPosition: ctypes.POINTER(CThostFtdcQryRiskSettleInvstPositionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者风险结算持仓查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRiskSettleInvstPosition
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRiskSettleInvstPositionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRiskSettleInvstPosition, nRequestID)

    # ReqQryRiskSettleProductStatus 风险结算产品查询
    def ReqQryRiskSettleProductStatus(self, pQryRiskSettleProductStatus: ctypes.POINTER(CThostFtdcQryRiskSettleProductStatusField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """风险结算产品查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRiskSettleProductStatus
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRiskSettleProductStatusField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRiskSettleProductStatus, nRequestID)

    # ReqQrySPBMFutureParameter SPBM期货合约参数查询
    def ReqQrySPBMFutureParameter(self, pQrySPBMFutureParameter: ctypes.POINTER(CThostFtdcQrySPBMFutureParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM期货合约参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMFutureParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMFutureParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMFutureParameter, nRequestID)

    # ReqQrySPBMOptionParameter SPBM期权合约参数查询
    def ReqQrySPBMOptionParameter(self, pQrySPBMOptionParameter: ctypes.POINTER(CThostFtdcQrySPBMOptionParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM期权合约参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMOptionParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMOptionParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMOptionParameter, nRequestID)

    # ReqQrySPBMIntraParameter SPBM品种内对锁仓折扣参数查询
    def ReqQrySPBMIntraParameter(self, pQrySPBMIntraParameter: ctypes.POINTER(CThostFtdcQrySPBMIntraParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM品种内对锁仓折扣参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMIntraParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMIntraParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMIntraParameter, nRequestID)

    # ReqQrySPBMInterParameter SPBM跨品种抵扣参数查询
    def ReqQrySPBMInterParameter(self, pQrySPBMInterParameter: ctypes.POINTER(CThostFtdcQrySPBMInterParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM跨品种抵扣参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMInterParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMInterParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMInterParameter, nRequestID)

    # ReqQrySPBMPortfDefinition SPBM组合保证金套餐查询
    def ReqQrySPBMPortfDefinition(self, pQrySPBMPortfDefinition: ctypes.POINTER(CThostFtdcQrySPBMPortfDefinitionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM组合保证金套餐查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMPortfDefinition
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMPortfDefinitionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMPortfDefinition, nRequestID)

    # ReqQrySPBMInvestorPortfDef 投资者SPBM套餐选择查询
    def ReqQrySPBMInvestorPortfDef(self, pQrySPBMInvestorPortfDef: ctypes.POINTER(CThostFtdcQrySPBMInvestorPortfDefField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者SPBM套餐选择查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMInvestorPortfDef
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMInvestorPortfDefField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMInvestorPortfDef, nRequestID)

    # ReqQryInvestorPortfMarginRatio 投资者新型组合保证金系数查询
    def ReqQryInvestorPortfMarginRatio(self, pQryInvestorPortfMarginRatio: ctypes.POINTER(CThostFtdcQryInvestorPortfMarginRatioField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者新型组合保证金系数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorPortfMarginRatio
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorPortfMarginRatioField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorPortfMarginRatio, nRequestID)

    # ReqQryInvestorProdSPBMDetail 投资者产品SPBM明细查询
    def ReqQryInvestorProdSPBMDetail(self, pQryInvestorProdSPBMDetail: ctypes.POINTER(CThostFtdcQryInvestorProdSPBMDetailField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者产品SPBM明细查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorProdSPBMDetail
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorProdSPBMDetailField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorProdSPBMDetail, nRequestID)

    # ReqQryInvestorCommoditySPMMMargin 投资者商品组SPMM记录查询
    def ReqQryInvestorCommoditySPMMMargin(self, pQryInvestorCommoditySPMMMargin: ctypes.POINTER(CThostFtdcQryInvestorCommoditySPMMMarginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者商品组SPMM记录查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorCommoditySPMMMargin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorCommoditySPMMMarginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorCommoditySPMMMargin, nRequestID)

    # ReqQryInvestorCommodityGroupSPMMMargin 投资者商品群SPMM记录查询
    def ReqQryInvestorCommodityGroupSPMMMargin(self, pQryInvestorCommodityGroupSPMMMargin: ctypes.POINTER(CThostFtdcQryInvestorCommodityGroupSPMMMarginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者商品群SPMM记录查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorCommodityGroupSPMMMargin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorCommodityGroupSPMMMarginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorCommodityGroupSPMMMargin, nRequestID)

    # ReqQrySPMMInstParam SPMM合约参数查询
    def ReqQrySPMMInstParam(self, pQrySPMMInstParam: ctypes.POINTER(CThostFtdcQrySPMMInstParamField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPMM合约参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPMMInstParam
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPMMInstParamField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPMMInstParam, nRequestID)

    # ReqQrySPMMProductParam SPMM产品参数查询
    def ReqQrySPMMProductParam(self, pQrySPMMProductParam: ctypes.POINTER(CThostFtdcQrySPMMProductParamField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPMM产品参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPMMProductParam
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPMMProductParamField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPMMProductParam, nRequestID)

    # ReqQrySPBMAddOnInterParameter SPBM附加跨品种抵扣参数查询
    def ReqQrySPBMAddOnInterParameter(self, pQrySPBMAddOnInterParameter: ctypes.POINTER(CThostFtdcQrySPBMAddOnInterParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """SPBM附加跨品种抵扣参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQrySPBMAddOnInterParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQrySPBMAddOnInterParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQrySPBMAddOnInterParameter, nRequestID)

    # ReqQryRCAMSCombProductInfo RCAMS产品组合信息查询
    def ReqQryRCAMSCombProductInfo(self, pQryRCAMSCombProductInfo: ctypes.POINTER(CThostFtdcQryRCAMSCombProductInfoField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS产品组合信息查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSCombProductInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSCombProductInfoField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSCombProductInfo, nRequestID)

    # ReqQryRCAMSInstrParameter RCAMS同合约风险对冲参数查询
    def ReqQryRCAMSInstrParameter(self, pQryRCAMSInstrParameter: ctypes.POINTER(CThostFtdcQryRCAMSInstrParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS同合约风险对冲参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSInstrParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSInstrParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSInstrParameter, nRequestID)

    # ReqQryRCAMSIntraParameter RCAMS品种内风险对冲参数查询
    def ReqQryRCAMSIntraParameter(self, pQryRCAMSIntraParameter: ctypes.POINTER(CThostFtdcQryRCAMSIntraParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS品种内风险对冲参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSIntraParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSIntraParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSIntraParameter, nRequestID)

    # ReqQryRCAMSInterParameter RCAMS跨品种风险折抵参数查询
    def ReqQryRCAMSInterParameter(self, pQryRCAMSInterParameter: ctypes.POINTER(CThostFtdcQryRCAMSInterParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS跨品种风险折抵参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSInterParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSInterParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSInterParameter, nRequestID)

    # ReqQryRCAMSShortOptAdjustParam RCAMS空头期权风险调整参数查询
    def ReqQryRCAMSShortOptAdjustParam(self, pQryRCAMSShortOptAdjustParam: ctypes.POINTER(CThostFtdcQryRCAMSShortOptAdjustParamField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS空头期权风险调整参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSShortOptAdjustParam
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSShortOptAdjustParamField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSShortOptAdjustParam, nRequestID)

    # ReqQryRCAMSInvestorCombPosition RCAMS策略组合持仓查询
    def ReqQryRCAMSInvestorCombPosition(self, pQryRCAMSInvestorCombPosition: ctypes.POINTER(CThostFtdcQryRCAMSInvestorCombPositionField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RCAMS策略组合持仓查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRCAMSInvestorCombPosition
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRCAMSInvestorCombPositionField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRCAMSInvestorCombPosition, nRequestID)

    # ReqQryInvestorProdRCAMSMargin 投资者品种RCAMS保证金查询
    def ReqQryInvestorProdRCAMSMargin(self, pQryInvestorProdRCAMSMargin: ctypes.POINTER(CThostFtdcQryInvestorProdRCAMSMarginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者品种RCAMS保证金查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorProdRCAMSMargin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorProdRCAMSMarginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorProdRCAMSMargin, nRequestID)

    # ReqQryRULEInstrParameter RULE合约保证金参数查询
    def ReqQryRULEInstrParameter(self, pQryRULEInstrParameter: ctypes.POINTER(CThostFtdcQryRULEInstrParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RULE合约保证金参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRULEInstrParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRULEInstrParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRULEInstrParameter, nRequestID)

    # ReqQryRULEIntraParameter RULE品种内对锁仓折扣参数查询
    def ReqQryRULEIntraParameter(self, pQryRULEIntraParameter: ctypes.POINTER(CThostFtdcQryRULEIntraParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RULE品种内对锁仓折扣参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRULEIntraParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRULEIntraParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRULEIntraParameter, nRequestID)

    # ReqQryRULEInterParameter RULE跨品种抵扣参数查询
    def ReqQryRULEInterParameter(self, pQryRULEInterParameter: ctypes.POINTER(CThostFtdcQryRULEInterParameterField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """RULE跨品种抵扣参数查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryRULEInterParameter
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryRULEInterParameterField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryRULEInterParameter, nRequestID)

    # ReqQryInvestorProdRULEMargin 投资者产品RULE保证金查询
    def ReqQryInvestorProdRULEMargin(self, pQryInvestorProdRULEMargin: ctypes.POINTER(CThostFtdcQryInvestorProdRULEMarginField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者产品RULE保证金查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorProdRULEMargin
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorProdRULEMarginField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorProdRULEMargin, nRequestID)

    # ReqQryInvestorPortfSetting 投资者投资者新组保设置查询
    def ReqQryInvestorPortfSetting(self, pQryInvestorPortfSetting: ctypes.POINTER(CThostFtdcQryInvestorPortfSettingField), nRequestID: ctypes.c_int32) -> ctypes.c_int32:
        """投资者投资者新组保设置查询"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqQryInvestorPortfSetting
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryInvestorPortfSettingField), ctypes.c_int32]
        func.restype = ctypes.c_int32
        return func(self._handle, pQryInvestorPortfSetting, nRequestID)

    # SpiCreate ========== Trader SPI 函数 ========== 创建 SPI 实例
    def SpiCreate(self, ) -> ctypes.c_void_p:
        """========== Trader SPI 函数 ========== 创建 SPI 实例"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSpiCreate
        func.argtypes = [ctypes.c_void_p]
        func.restype = ctypes.c_void_p
        return func(self._handle)

    # SpiDestroy 销毁 SPI 实例
    def SpiDestroy(self, ):
        """销毁 SPI 实例"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSpiDestroy
        func.argtypes = [ctypes.c_void_p]
        func.restype = None
        func(self._handle)

    # RegisterSpi 注册 SPI 到 API
    def RegisterSpi(self, spi: ctypes.c_void_p):
        """注册 SPI 到 API"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderRegisterSpi
        func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]
        func.restype = None
        func(self._handle, spi)

    # SpiSetCallbacks 批量设置回调
    def SpiSetCallbacks(self, callbacks: ctypes.POINTER(TraderSpiCallbacks)):
        """批量设置回调"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderSpiSetCallbacks
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(TraderSpiCallbacks)]
        func.restype = None
        func(self._handle, callbacks)

    # ReqUserLoginWithSystemInfo ========== 跨平台统一登录接口 ========== 说明: macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数 此函数在 Linux/Windows 上忽略 systemInfo，在 macOS 上使用它 带系统信息的用户登录请求（跨平台统一接口） systemInfoLen: 系统信息长度，传 0 表示自动采集（仅 macOS 生效） systemInfo: 系统信息数据，传 NULL 表示自动采集（仅 macOS 生效）
    def ReqUserLoginWithSystemInfo(self, pReqUserLoginField: ctypes.POINTER(CThostFtdcReqUserLoginField), nRequestID: ctypes.c_int32, systemInfoLen: ctypes.c_int32, systemInfo: str) -> ctypes.c_int32:
        """========== 跨平台统一登录接口 ========== 说明: macOS 版本的 ReqUserLogin 需要额外的 systemInfo 参数 此函数在 Linux/Windows 上忽略 systemInfo，在 macOS 上使用它 带系统信息的用户登录请求（跨平台统一接口） systemInfoLen: 系统信息长度，传 0 表示自动采集（仅 macOS 生效） systemInfo: 系统信息数据，传 NULL 表示自动采集（仅 macOS 生效）"""
        lib = get_trader_lib_handle()
        if lib is None:
            raise RuntimeError("CTP library not loaded")
        func = lib.TraderReqUserLoginWithSystemInfo
        func.argtypes = [ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqUserLoginField), ctypes.c_int32, ctypes.c_int32, ctypes.c_char_p]
        func.restype = ctypes.c_int32
        return func(self._handle, pReqUserLoginField, nRequestID, systemInfoLen, c_string(systemInfo))

    def set_spi(self, spi: TraderSpi):
        """设置回调接口"""
        with self._lock:
            self._spi = spi

            if self._spi_handle:
                lib = get_trader_lib_handle()
                if lib:
                    func = lib.TraderSpiDestroy
                    func.argtypes = [ctypes.c_void_p]
                    func.restype = None
                    func(self._spi_handle)
                self._spi_handle = None

            lib = get_trader_lib_handle()
            if lib is None:
                raise RuntimeError("CTP library not loaded")

            func = lib.TraderSpiCreate
            func.argtypes = [ctypes.c_void_p]
            func.restype = ctypes.c_void_p
            self._spi_handle = func(ctypes.c_void_p(self._user_data))

            # 注册所有回调函数到 C SPI
            _register_trader_callback(self._spi_handle, lib, "FrontConnected", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "FrontDisconnected", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "HeartBeatWarning", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspAuthenticate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspUserLogin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspUserLogout", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspUserPasswordUpdate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspTradingAccountPasswordUpdate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspUserAuthMethod", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspGenUserCaptcha", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspGenUserText", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspOrderInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspParkedOrderInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspParkedOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryMaxOrderVolume", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspSettlementInfoConfirm", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspRemoveParkedOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspRemoveParkedOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspExecOrderInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspExecOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspForQuoteInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQuoteInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQuoteAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspBatchOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspOptionSelfCloseInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspOptionSelfCloseAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspCombActionInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTrade", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorPosition", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTradingAccount", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestor", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTradingCode", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInstrumentMarginRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInstrumentCommissionRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryExchange", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryProduct", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInstrument", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryDepthMarketData", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTraderOffer", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySettlementInfo", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTransferBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorPositionDetail", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryNotice", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySettlementInfoConfirm", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorPositionCombineDetail", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryCFMMCTradingAccountKey", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryEWarrantOffset", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorProductGroupMargin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryExchangeMarginRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryExchangeMarginRateAdjust", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryExchangeRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySecAgentACIDMap", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryProductExchRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryProductGroup", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryMMInstrumentCommissionRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryMMOptionInstrCommRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInstrumentOrderCommRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySecAgentTradingAccount", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySecAgentCheckMode", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySecAgentTradeInfo", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryOptionInstrTradeCost", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryOptionInstrCommRate", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryExecOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryForQuote", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryQuote", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryOptionSelfClose", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestUnit", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryCombInstrumentGuard", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryCombAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTransferSerial", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryAccountregister", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspError", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnTrade", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnOrderInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnInstrumentStatus", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnBulletin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnTradingNotice", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnErrorConditionalOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnExecOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnExecOrderInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnExecOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnForQuoteInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnQuote", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnQuoteInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnQuoteAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnForQuoteRsp", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnCFMMCTradingAccountToken", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnBatchOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnOptionSelfClose", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnOptionSelfCloseInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnOptionSelfCloseAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnCombAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnCombActionInsert", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryContractBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryParkedOrder", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryParkedOrderAction", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryTradingNotice", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryBrokerTradingParams", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryBrokerTradingAlgos", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQueryCFMMCTradingAccountToken", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnFromBankToFutureByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnFromFutureToBankByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromBankToFutureByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromFutureToBankByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnFromBankToFutureByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnFromFutureToBankByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromBankToFutureByFutureManual", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromFutureToBankByFutureManual", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnQueryBankBalanceByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnBankToFutureByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnFutureToBankByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnRepealBankToFutureByFutureManual", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnRepealFutureToBankByFutureManual", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "ErrRtnQueryBankBalanceByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromBankToFutureByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnRepealFromFutureToBankByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspFromBankToFutureByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspFromFutureToBankByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQueryBankAccountMoneyByFuture", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnOpenAccountByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnCancelAccountByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RtnChangeAccountByBank", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryClassifiedInstrument", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryCombPromotionParam", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRiskSettleInvstPosition", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRiskSettleProductStatus", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMFutureParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMOptionParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMIntraParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMInterParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMPortfDefinition", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMInvestorPortfDef", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorPortfMarginRatio", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorProdSPBMDetail", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorCommoditySPMMMargin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorCommodityGroupSPMMMargin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPMMInstParam", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPMMProductParam", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQrySPBMAddOnInterParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSCombProductInfo", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSInstrParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSIntraParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSInterParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSShortOptAdjustParam", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRCAMSInvestorCombPosition", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorProdRCAMSMargin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRULEInstrParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRULEIntraParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryRULEInterParameter", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorProdRULEMargin", self._spi, self._user_data)
            _register_trader_callback(self._spi_handle, lib, "RspQryInvestorPortfSetting", self._spi, self._user_data)

            func = lib.TraderRegisterSpi
            func.argtypes = [ctypes.c_void_p, ctypes.c_void_p]
            func.restype = None
            func(self._handle, self._spi_handle)

# ========== 实例管理 ==========

_trader_instances: dict = {}
_trader_instances_lock = threading.RLock()
_trader_next_id = 1

def _register_trader_instance(api: TraderApi) -> int:
    """注册交易 API 实例"""
    global _trader_next_id
    with _trader_instances_lock:
        instance_id = _trader_next_id
        _trader_next_id += 1
        _trader_instances[instance_id] = api
        return instance_id

def _get_trader_instance(user_data: int) -> Optional[TraderApi]:
    """获取交易 API 实例"""
    with _trader_instances_lock:
        return _trader_instances.get(user_data)

def _unregister_trader_instance(user_data: int):
    """注销交易 API 实例"""
    with _trader_instances_lock:
        _trader_instances.pop(user_data, None)

def _register_trader_callback(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi: TraderSpi, user_data: int):
    """注册回调函数到 C SPI"""
    # 实际实现在 trader_callbacks.py 中
    from .trader_callbacks import _register_trader_callback_impl
    _register_trader_callback_impl(spi_handle, lib, callback_name, spi, user_data)

# ========== DataCollect 函数 ==========

# GetSystemInfo ========== DataCollect 函数 ========== 获取终端信息（AES+RSA 加密） pSystemInfo: 输出缓冲区，至少 270 字节 pLen: 输入缓冲区大小，输出实际数据长度 返回值: 0 成功，非 0 表示采集错误（按位判断）
def GetSystemInfo() -> tuple[bytes, int]:
    """
    获取终端信息（AES+RSA 加密）
    
    返回:
        tuple[bytes, int]: (系统信息字节数组, 错误码)
        错误码为 0 表示成功，非 0 表示采集错误（按位判断）
    """
    lib = get_trader_lib_handle()
    if lib is None:
        return None, -1
    
    func = lib.DCGetSystemInfo
    func.argtypes = [ctypes.POINTER(ctypes.c_byte), ctypes.POINTER(ctypes.c_int32)]
    func.restype = ctypes.c_int32
    
    # 分配至少 270 字节的缓冲区
    buf_size = 512
    buf = (ctypes.c_byte * buf_size)()
    buf_len = ctypes.c_int32(buf_size)
    
    ret = func(ctypes.cast(buf, ctypes.POINTER(ctypes.c_byte)), ctypes.byref(buf_len))
    
    if ret != 0:
        return None, ret
    
    # 返回实际长度的字节数组
    return bytes(buf[:buf_len.value]), 0

# GetSystemInfoUnAesEncode 获取终端信息（未 AES 加密）
def GetSystemInfoUnAesEncode() -> tuple[bytes, int]:
    """
    获取终端信息（未 AES 加密）
    
    返回:
        tuple[bytes, int]: (系统信息字节数组, 错误码)
        错误码为 0 表示成功，非 0 表示采集错误（按位判断）
    """
    lib = get_trader_lib_handle()
    if lib is None:
        return None, -1
    
    func = lib.DCGetSystemInfoUnAesEncode
    func.argtypes = [ctypes.POINTER(ctypes.c_byte), ctypes.POINTER(ctypes.c_int32)]
    func.restype = ctypes.c_int32
    
    # 分配至少 270 字节的缓冲区
    buf_size = 512
    buf = (ctypes.c_byte * buf_size)()
    buf_len = ctypes.c_int32(buf_size)
    
    ret = func(ctypes.cast(buf, ctypes.POINTER(ctypes.c_byte)), ctypes.byref(buf_len))
    
    if ret != 0:
        return None, ret
    
    # 返回实际长度的字节数组
    return bytes(buf[:buf_len.value]), 0

# GetDataCollectApiVersion 获取 DataCollect API 版本
def GetDataCollectApiVersion() -> str:
    """
    获取 DataCollect API 版本
    
    返回:
        str: API 版本字符串
    """
    lib = get_trader_lib_handle()
    if lib is None:
        return ""
    
    func = lib.DCGetDataCollectApiVersion
    func.argtypes = []
    func.restype = ctypes.c_char_p
    
    ptr = func()
    return go_string(ptr) if ptr else ""
