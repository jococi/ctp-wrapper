"""
CTP 交易回调实现

此文件由代码生成器自动生成，请勿手动修改
CTP 交易回调实现
"""

import ctypes
from .trader_api import _get_trader_instance
from .struct import *

# ========== 回调包装函数 ==========

def _go_trader_OnFrontConnected(userData: ctypes.c_void_p):
    """回调函数实现: ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnFrontConnected()
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnFrontDisconnected(userData: ctypes.c_void_p, nReason: ctypes.c_int32):
    """回调函数实现: 0x2003 收到错误报文"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnFrontDisconnected(nReason)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnHeartBeatWarning(userData: ctypes.c_void_p, nTimeLapse: ctypes.c_int32):
    """回调函数实现: 心跳超时警告。当长时间未收到报文时，该方法被调用。"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnHeartBeatWarning(nTimeLapse)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspAuthenticate(userData: ctypes.c_void_p, pRspAuthenticateField: ctypes.POINTER(CThostFtdcRspAuthenticateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 客户端认证响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspAuthenticate(pRspAuthenticateField.contents if pRspAuthenticateField else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspUserLogin(userData: ctypes.c_void_p, pRspUserLogin: ctypes.POINTER(CThostFtdcRspUserLoginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 登录请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserLogin(pRspUserLogin.contents if pRspUserLogin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspUserLogout(userData: ctypes.c_void_p, pUserLogout: ctypes.POINTER(CThostFtdcUserLogoutField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 登出请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserLogout(pUserLogout.contents if pUserLogout else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspUserPasswordUpdate(userData: ctypes.c_void_p, pUserPasswordUpdate: ctypes.POINTER(CThostFtdcUserPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 用户口令更新请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserPasswordUpdate(pUserPasswordUpdate.contents if pUserPasswordUpdate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspTradingAccountPasswordUpdate(userData: ctypes.c_void_p, pTradingAccountPasswordUpdate: ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 资金账户口令更新请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspTradingAccountPasswordUpdate(pTradingAccountPasswordUpdate.contents if pTradingAccountPasswordUpdate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspUserAuthMethod(userData: ctypes.c_void_p, pRspUserAuthMethod: ctypes.POINTER(CThostFtdcRspUserAuthMethodField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 查询用户当前支持的认证模式的回复"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspUserAuthMethod(pRspUserAuthMethod.contents if pRspUserAuthMethod else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspGenUserCaptcha(userData: ctypes.c_void_p, pRspGenUserCaptcha: ctypes.POINTER(CThostFtdcRspGenUserCaptchaField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 获取图形验证码请求的回复"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspGenUserCaptcha(pRspGenUserCaptcha.contents if pRspGenUserCaptcha else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspGenUserText(userData: ctypes.c_void_p, pRspGenUserText: ctypes.POINTER(CThostFtdcRspGenUserTextField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 获取短信验证码请求的回复"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspGenUserText(pRspGenUserText.contents if pRspGenUserText else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspOrderInsert(userData: ctypes.c_void_p, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 报单录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspOrderInsert(pInputOrder.contents if pInputOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspParkedOrderInsert(userData: ctypes.c_void_p, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 预埋单录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspParkedOrderInsert(pParkedOrder.contents if pParkedOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspParkedOrderAction(userData: ctypes.c_void_p, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 预埋撤单录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspParkedOrderAction(pParkedOrderAction.contents if pParkedOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspOrderAction(userData: ctypes.c_void_p, pInputOrderAction: ctypes.POINTER(CThostFtdcInputOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 报单操作请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspOrderAction(pInputOrderAction.contents if pInputOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryMaxOrderVolume(userData: ctypes.c_void_p, pQryMaxOrderVolume: ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 查询最大报单数量响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryMaxOrderVolume(pQryMaxOrderVolume.contents if pQryMaxOrderVolume else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspSettlementInfoConfirm(userData: ctypes.c_void_p, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者结算结果确认响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspSettlementInfoConfirm(pSettlementInfoConfirm.contents if pSettlementInfoConfirm else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspRemoveParkedOrder(userData: ctypes.c_void_p, pRemoveParkedOrder: ctypes.POINTER(CThostFtdcRemoveParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 删除预埋单响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspRemoveParkedOrder(pRemoveParkedOrder.contents if pRemoveParkedOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspRemoveParkedOrderAction(userData: ctypes.c_void_p, pRemoveParkedOrderAction: ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 删除预埋撤单响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspRemoveParkedOrderAction(pRemoveParkedOrderAction.contents if pRemoveParkedOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspExecOrderInsert(userData: ctypes.c_void_p, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 执行宣告录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspExecOrderInsert(pInputExecOrder.contents if pInputExecOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspExecOrderAction(userData: ctypes.c_void_p, pInputExecOrderAction: ctypes.POINTER(CThostFtdcInputExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 执行宣告操作请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspExecOrderAction(pInputExecOrderAction.contents if pInputExecOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspForQuoteInsert(userData: ctypes.c_void_p, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 询价录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspForQuoteInsert(pInputForQuote.contents if pInputForQuote else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQuoteInsert(userData: ctypes.c_void_p, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 报价录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQuoteInsert(pInputQuote.contents if pInputQuote else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQuoteAction(userData: ctypes.c_void_p, pInputQuoteAction: ctypes.POINTER(CThostFtdcInputQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 报价操作请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQuoteAction(pInputQuoteAction.contents if pInputQuoteAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspBatchOrderAction(userData: ctypes.c_void_p, pInputBatchOrderAction: ctypes.POINTER(CThostFtdcInputBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 批量报单操作请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspBatchOrderAction(pInputBatchOrderAction.contents if pInputBatchOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspOptionSelfCloseInsert(userData: ctypes.c_void_p, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 期权自对冲录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspOptionSelfCloseInsert(pInputOptionSelfClose.contents if pInputOptionSelfClose else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspOptionSelfCloseAction(userData: ctypes.c_void_p, pInputOptionSelfCloseAction: ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 期权自对冲操作请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspOptionSelfCloseAction(pInputOptionSelfCloseAction.contents if pInputOptionSelfCloseAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspCombActionInsert(userData: ctypes.c_void_p, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 申请组合录入请求响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspCombActionInsert(pInputCombAction.contents if pInputCombAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryOrder(userData: ctypes.c_void_p, pOrder: ctypes.POINTER(CThostFtdcOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询报单响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryOrder(pOrder.contents if pOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTrade(userData: ctypes.c_void_p, pTrade: ctypes.POINTER(CThostFtdcTradeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询成交响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTrade(pTrade.contents if pTrade else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorPosition(userData: ctypes.c_void_p, pInvestorPosition: ctypes.POINTER(CThostFtdcInvestorPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者持仓响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorPosition(pInvestorPosition.contents if pInvestorPosition else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTradingAccount(userData: ctypes.c_void_p, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询资金账户响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTradingAccount(pTradingAccount.contents if pTradingAccount else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestor(userData: ctypes.c_void_p, pInvestor: ctypes.POINTER(CThostFtdcInvestorField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestor(pInvestor.contents if pInvestor else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTradingCode(userData: ctypes.c_void_p, pTradingCode: ctypes.POINTER(CThostFtdcTradingCodeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易编码响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTradingCode(pTradingCode.contents if pTradingCode else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInstrumentMarginRate(userData: ctypes.c_void_p, pInstrumentMarginRate: ctypes.POINTER(CThostFtdcInstrumentMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询合约保证金率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInstrumentMarginRate(pInstrumentMarginRate.contents if pInstrumentMarginRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInstrumentCommissionRate(userData: ctypes.c_void_p, pInstrumentCommissionRate: ctypes.POINTER(CThostFtdcInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询合约手续费率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInstrumentCommissionRate(pInstrumentCommissionRate.contents if pInstrumentCommissionRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryExchange(userData: ctypes.c_void_p, pExchange: ctypes.POINTER(CThostFtdcExchangeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易所响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryExchange(pExchange.contents if pExchange else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryProduct(userData: ctypes.c_void_p, pProduct: ctypes.POINTER(CThostFtdcProductField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询产品响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryProduct(pProduct.contents if pProduct else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInstrument(userData: ctypes.c_void_p, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询合约响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInstrument(pInstrument.contents if pInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryDepthMarketData(userData: ctypes.c_void_p, pDepthMarketData: ctypes.POINTER(CThostFtdcDepthMarketDataField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询行情响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryDepthMarketData(pDepthMarketData.contents if pDepthMarketData else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTraderOffer(userData: ctypes.c_void_p, pTraderOffer: ctypes.POINTER(CThostFtdcTraderOfferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易员报盘机响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTraderOffer(pTraderOffer.contents if pTraderOffer else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySettlementInfo(userData: ctypes.c_void_p, pSettlementInfo: ctypes.POINTER(CThostFtdcSettlementInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者结算结果响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySettlementInfo(pSettlementInfo.contents if pSettlementInfo else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTransferBank(userData: ctypes.c_void_p, pTransferBank: ctypes.POINTER(CThostFtdcTransferBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询转帐银行响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTransferBank(pTransferBank.contents if pTransferBank else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorPositionDetail(userData: ctypes.c_void_p, pInvestorPositionDetail: ctypes.POINTER(CThostFtdcInvestorPositionDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者持仓明细响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorPositionDetail(pInvestorPositionDetail.contents if pInvestorPositionDetail else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryNotice(userData: ctypes.c_void_p, pNotice: ctypes.POINTER(CThostFtdcNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询客户通知响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryNotice(pNotice.contents if pNotice else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySettlementInfoConfirm(userData: ctypes.c_void_p, pSettlementInfoConfirm: ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询结算信息确认响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySettlementInfoConfirm(pSettlementInfoConfirm.contents if pSettlementInfoConfirm else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorPositionCombineDetail(userData: ctypes.c_void_p, pInvestorPositionCombineDetail: ctypes.POINTER(CThostFtdcInvestorPositionCombineDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者持仓明细响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorPositionCombineDetail(pInvestorPositionCombineDetail.contents if pInvestorPositionCombineDetail else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryCFMMCTradingAccountKey(userData: ctypes.c_void_p, pCFMMCTradingAccountKey: ctypes.POINTER(CThostFtdcCFMMCTradingAccountKeyField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 查询保证金监管系统经纪公司资金账户密钥响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryCFMMCTradingAccountKey(pCFMMCTradingAccountKey.contents if pCFMMCTradingAccountKey else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryEWarrantOffset(userData: ctypes.c_void_p, pEWarrantOffset: ctypes.POINTER(CThostFtdcEWarrantOffsetField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询仓单折抵信息响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryEWarrantOffset(pEWarrantOffset.contents if pEWarrantOffset else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorProductGroupMargin(userData: ctypes.c_void_p, pInvestorProductGroupMargin: ctypes.POINTER(CThostFtdcInvestorProductGroupMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资者品种/跨品种保证金响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorProductGroupMargin(pInvestorProductGroupMargin.contents if pInvestorProductGroupMargin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryExchangeMarginRate(userData: ctypes.c_void_p, pExchangeMarginRate: ctypes.POINTER(CThostFtdcExchangeMarginRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易所保证金率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryExchangeMarginRate(pExchangeMarginRate.contents if pExchangeMarginRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryExchangeMarginRateAdjust(userData: ctypes.c_void_p, pExchangeMarginRateAdjust: ctypes.POINTER(CThostFtdcExchangeMarginRateAdjustField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易所调整保证金率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryExchangeMarginRateAdjust(pExchangeMarginRateAdjust.contents if pExchangeMarginRateAdjust else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryExchangeRate(userData: ctypes.c_void_p, pExchangeRate: ctypes.POINTER(CThostFtdcExchangeRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询汇率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryExchangeRate(pExchangeRate.contents if pExchangeRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySecAgentACIDMap(userData: ctypes.c_void_p, pSecAgentACIDMap: ctypes.POINTER(CThostFtdcSecAgentACIDMapField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询二级代理操作员银期权限响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySecAgentACIDMap(pSecAgentACIDMap.contents if pSecAgentACIDMap else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryProductExchRate(userData: ctypes.c_void_p, pProductExchRate: ctypes.POINTER(CThostFtdcProductExchRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询产品报价汇率"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryProductExchRate(pProductExchRate.contents if pProductExchRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryProductGroup(userData: ctypes.c_void_p, pProductGroup: ctypes.POINTER(CThostFtdcProductGroupField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询产品组"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryProductGroup(pProductGroup.contents if pProductGroup else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryMMInstrumentCommissionRate(userData: ctypes.c_void_p, pMMInstrumentCommissionRate: ctypes.POINTER(CThostFtdcMMInstrumentCommissionRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询做市商合约手续费率响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryMMInstrumentCommissionRate(pMMInstrumentCommissionRate.contents if pMMInstrumentCommissionRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryMMOptionInstrCommRate(userData: ctypes.c_void_p, pMMOptionInstrCommRate: ctypes.POINTER(CThostFtdcMMOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询做市商期权合约手续费响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryMMOptionInstrCommRate(pMMOptionInstrCommRate.contents if pMMOptionInstrCommRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInstrumentOrderCommRate(userData: ctypes.c_void_p, pInstrumentOrderCommRate: ctypes.POINTER(CThostFtdcInstrumentOrderCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询报单手续费响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInstrumentOrderCommRate(pInstrumentOrderCommRate.contents if pInstrumentOrderCommRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySecAgentTradingAccount(userData: ctypes.c_void_p, pTradingAccount: ctypes.POINTER(CThostFtdcTradingAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询资金账户响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySecAgentTradingAccount(pTradingAccount.contents if pTradingAccount else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySecAgentCheckMode(userData: ctypes.c_void_p, pSecAgentCheckMode: ctypes.POINTER(CThostFtdcSecAgentCheckModeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询二级代理商资金校验模式响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySecAgentCheckMode(pSecAgentCheckMode.contents if pSecAgentCheckMode else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySecAgentTradeInfo(userData: ctypes.c_void_p, pSecAgentTradeInfo: ctypes.POINTER(CThostFtdcSecAgentTradeInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询二级代理商信息响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySecAgentTradeInfo(pSecAgentTradeInfo.contents if pSecAgentTradeInfo else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryOptionInstrTradeCost(userData: ctypes.c_void_p, pOptionInstrTradeCost: ctypes.POINTER(CThostFtdcOptionInstrTradeCostField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询期权交易成本响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryOptionInstrTradeCost(pOptionInstrTradeCost.contents if pOptionInstrTradeCost else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryOptionInstrCommRate(userData: ctypes.c_void_p, pOptionInstrCommRate: ctypes.POINTER(CThostFtdcOptionInstrCommRateField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询期权合约手续费响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryOptionInstrCommRate(pOptionInstrCommRate.contents if pOptionInstrCommRate else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryExecOrder(userData: ctypes.c_void_p, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询执行宣告响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryExecOrder(pExecOrder.contents if pExecOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryForQuote(userData: ctypes.c_void_p, pForQuote: ctypes.POINTER(CThostFtdcForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询询价响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryForQuote(pForQuote.contents if pForQuote else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryQuote(userData: ctypes.c_void_p, pQuote: ctypes.POINTER(CThostFtdcQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询报价响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryQuote(pQuote.contents if pQuote else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryOptionSelfClose(userData: ctypes.c_void_p, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询期权自对冲响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryOptionSelfClose(pOptionSelfClose.contents if pOptionSelfClose else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestUnit(userData: ctypes.c_void_p, pInvestUnit: ctypes.POINTER(CThostFtdcInvestUnitField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询投资单元响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestUnit(pInvestUnit.contents if pInvestUnit else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryCombInstrumentGuard(userData: ctypes.c_void_p, pCombInstrumentGuard: ctypes.POINTER(CThostFtdcCombInstrumentGuardField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询组合合约安全系数响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryCombInstrumentGuard(pCombInstrumentGuard.contents if pCombInstrumentGuard else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryCombAction(userData: ctypes.c_void_p, pCombAction: ctypes.POINTER(CThostFtdcCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询申请组合响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryCombAction(pCombAction.contents if pCombAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTransferSerial(userData: ctypes.c_void_p, pTransferSerial: ctypes.POINTER(CThostFtdcTransferSerialField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询转帐流水响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTransferSerial(pTransferSerial.contents if pTransferSerial else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryAccountregister(userData: ctypes.c_void_p, pAccountregister: ctypes.POINTER(CThostFtdcAccountregisterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询银期签约关系响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryAccountregister(pAccountregister.contents if pAccountregister else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspError(userData: ctypes.c_void_p, pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 错误应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspError(pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnOrder(userData: ctypes.c_void_p, pOrder: ctypes.POINTER(CThostFtdcOrderField)):
    """回调函数实现: 报单通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnOrder(pOrder.contents if pOrder else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnTrade(userData: ctypes.c_void_p, pTrade: ctypes.POINTER(CThostFtdcTradeField)):
    """回调函数实现: 成交通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnTrade(pTrade.contents if pTrade else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnOrderInsert(userData: ctypes.c_void_p, pInputOrder: ctypes.POINTER(CThostFtdcInputOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 报单录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnOrderInsert(pInputOrder.contents if pInputOrder else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnOrderAction(userData: ctypes.c_void_p, pOrderAction: ctypes.POINTER(CThostFtdcOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 报单操作错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnOrderAction(pOrderAction.contents if pOrderAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnInstrumentStatus(userData: ctypes.c_void_p, pInstrumentStatus: ctypes.POINTER(CThostFtdcInstrumentStatusField)):
    """回调函数实现: 合约交易状态通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnInstrumentStatus(pInstrumentStatus.contents if pInstrumentStatus else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnBulletin(userData: ctypes.c_void_p, pBulletin: ctypes.POINTER(CThostFtdcBulletinField)):
    """回调函数实现: 交易所公告通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnBulletin(pBulletin.contents if pBulletin else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnTradingNotice(userData: ctypes.c_void_p, pTradingNoticeInfo: ctypes.POINTER(CThostFtdcTradingNoticeInfoField)):
    """回调函数实现: 交易通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnTradingNotice(pTradingNoticeInfo.contents if pTradingNoticeInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnErrorConditionalOrder(userData: ctypes.c_void_p, pErrorConditionalOrder: ctypes.POINTER(CThostFtdcErrorConditionalOrderField)):
    """回调函数实现: 提示条件单校验错误"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnErrorConditionalOrder(pErrorConditionalOrder.contents if pErrorConditionalOrder else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnExecOrder(userData: ctypes.c_void_p, pExecOrder: ctypes.POINTER(CThostFtdcExecOrderField)):
    """回调函数实现: 执行宣告通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnExecOrder(pExecOrder.contents if pExecOrder else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnExecOrderInsert(userData: ctypes.c_void_p, pInputExecOrder: ctypes.POINTER(CThostFtdcInputExecOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 执行宣告录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnExecOrderInsert(pInputExecOrder.contents if pInputExecOrder else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnExecOrderAction(userData: ctypes.c_void_p, pExecOrderAction: ctypes.POINTER(CThostFtdcExecOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 执行宣告操作错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnExecOrderAction(pExecOrderAction.contents if pExecOrderAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnForQuoteInsert(userData: ctypes.c_void_p, pInputForQuote: ctypes.POINTER(CThostFtdcInputForQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 询价录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnForQuoteInsert(pInputForQuote.contents if pInputForQuote else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnQuote(userData: ctypes.c_void_p, pQuote: ctypes.POINTER(CThostFtdcQuoteField)):
    """回调函数实现: 报价通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnQuote(pQuote.contents if pQuote else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnQuoteInsert(userData: ctypes.c_void_p, pInputQuote: ctypes.POINTER(CThostFtdcInputQuoteField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 报价录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnQuoteInsert(pInputQuote.contents if pInputQuote else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnQuoteAction(userData: ctypes.c_void_p, pQuoteAction: ctypes.POINTER(CThostFtdcQuoteActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 报价操作错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnQuoteAction(pQuoteAction.contents if pQuoteAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnForQuoteRsp(userData: ctypes.c_void_p, pForQuoteRsp: ctypes.POINTER(CThostFtdcForQuoteRspField)):
    """回调函数实现: 询价通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnForQuoteRsp(pForQuoteRsp.contents if pForQuoteRsp else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnCFMMCTradingAccountToken(userData: ctypes.c_void_p, pCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcCFMMCTradingAccountTokenField)):
    """回调函数实现: 保证金监控中心用户令牌"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnCFMMCTradingAccountToken(pCFMMCTradingAccountToken.contents if pCFMMCTradingAccountToken else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnBatchOrderAction(userData: ctypes.c_void_p, pBatchOrderAction: ctypes.POINTER(CThostFtdcBatchOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 批量报单操作错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnBatchOrderAction(pBatchOrderAction.contents if pBatchOrderAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnOptionSelfClose(userData: ctypes.c_void_p, pOptionSelfClose: ctypes.POINTER(CThostFtdcOptionSelfCloseField)):
    """回调函数实现: 期权自对冲通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnOptionSelfClose(pOptionSelfClose.contents if pOptionSelfClose else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnOptionSelfCloseInsert(userData: ctypes.c_void_p, pInputOptionSelfClose: ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 期权自对冲录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnOptionSelfCloseInsert(pInputOptionSelfClose.contents if pInputOptionSelfClose else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnOptionSelfCloseAction(userData: ctypes.c_void_p, pOptionSelfCloseAction: ctypes.POINTER(CThostFtdcOptionSelfCloseActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 期权自对冲操作错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnOptionSelfCloseAction(pOptionSelfCloseAction.contents if pOptionSelfCloseAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnCombAction(userData: ctypes.c_void_p, pCombAction: ctypes.POINTER(CThostFtdcCombActionField)):
    """回调函数实现: 申请组合通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnCombAction(pCombAction.contents if pCombAction else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnCombActionInsert(userData: ctypes.c_void_p, pInputCombAction: ctypes.POINTER(CThostFtdcInputCombActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 申请组合录入错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnCombActionInsert(pInputCombAction.contents if pInputCombAction else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryContractBank(userData: ctypes.c_void_p, pContractBank: ctypes.POINTER(CThostFtdcContractBankField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询签约银行响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryContractBank(pContractBank.contents if pContractBank else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryParkedOrder(userData: ctypes.c_void_p, pParkedOrder: ctypes.POINTER(CThostFtdcParkedOrderField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询预埋单响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryParkedOrder(pParkedOrder.contents if pParkedOrder else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryParkedOrderAction(userData: ctypes.c_void_p, pParkedOrderAction: ctypes.POINTER(CThostFtdcParkedOrderActionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询预埋撤单响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryParkedOrderAction(pParkedOrderAction.contents if pParkedOrderAction else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryTradingNotice(userData: ctypes.c_void_p, pTradingNotice: ctypes.POINTER(CThostFtdcTradingNoticeField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询交易通知响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryTradingNotice(pTradingNotice.contents if pTradingNotice else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryBrokerTradingParams(userData: ctypes.c_void_p, pBrokerTradingParams: ctypes.POINTER(CThostFtdcBrokerTradingParamsField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询经纪公司交易参数响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryBrokerTradingParams(pBrokerTradingParams.contents if pBrokerTradingParams else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryBrokerTradingAlgos(userData: ctypes.c_void_p, pBrokerTradingAlgos: ctypes.POINTER(CThostFtdcBrokerTradingAlgosField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询经纪公司交易算法响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryBrokerTradingAlgos(pBrokerTradingAlgos.contents if pBrokerTradingAlgos else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQueryCFMMCTradingAccountToken(userData: ctypes.c_void_p, pQueryCFMMCTradingAccountToken: ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询监控中心用户令牌"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQueryCFMMCTradingAccountToken(pQueryCFMMCTradingAccountToken.contents if pQueryCFMMCTradingAccountToken else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnFromBankToFutureByBank(userData: ctypes.c_void_p, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
    """回调函数实现: 银行发起银行资金转期货通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnFromBankToFutureByBank(pRspTransfer.contents if pRspTransfer else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnFromFutureToBankByBank(userData: ctypes.c_void_p, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
    """回调函数实现: 银行发起期货资金转银行通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnFromFutureToBankByBank(pRspTransfer.contents if pRspTransfer else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromBankToFutureByBank(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 银行发起冲正银行转期货通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromBankToFutureByBank(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromFutureToBankByBank(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 银行发起冲正期货转银行通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromFutureToBankByBank(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnFromBankToFutureByFuture(userData: ctypes.c_void_p, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
    """回调函数实现: 期货发起银行资金转期货通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnFromBankToFutureByFuture(pRspTransfer.contents if pRspTransfer else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnFromFutureToBankByFuture(userData: ctypes.c_void_p, pRspTransfer: ctypes.POINTER(CThostFtdcRspTransferField)):
    """回调函数实现: 期货发起期货资金转银行通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnFromFutureToBankByFuture(pRspTransfer.contents if pRspTransfer else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromBankToFutureByFutureManual(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 系统运行时期货端手工发起冲正银行转期货请求，银行处理完毕后报盘发回的通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromBankToFutureByFutureManual(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromFutureToBankByFutureManual(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 系统运行时期货端手工发起冲正期货转银行请求，银行处理完毕后报盘发回的通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromFutureToBankByFutureManual(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnQueryBankBalanceByFuture(userData: ctypes.c_void_p, pNotifyQueryAccount: ctypes.POINTER(CThostFtdcNotifyQueryAccountField)):
    """回调函数实现: 期货发起查询银行余额通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnQueryBankBalanceByFuture(pNotifyQueryAccount.contents if pNotifyQueryAccount else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnBankToFutureByFuture(userData: ctypes.c_void_p, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 期货发起银行资金转期货错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnBankToFutureByFuture(pReqTransfer.contents if pReqTransfer else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnFutureToBankByFuture(userData: ctypes.c_void_p, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 期货发起期货资金转银行错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnFutureToBankByFuture(pReqTransfer.contents if pReqTransfer else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnRepealBankToFutureByFutureManual(userData: ctypes.c_void_p, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 系统运行时期货端手工发起冲正银行转期货错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnRepealBankToFutureByFutureManual(pReqRepeal.contents if pReqRepeal else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnRepealFutureToBankByFutureManual(userData: ctypes.c_void_p, pReqRepeal: ctypes.POINTER(CThostFtdcReqRepealField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 系统运行时期货端手工发起冲正期货转银行错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnRepealFutureToBankByFutureManual(pReqRepeal.contents if pReqRepeal else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnErrRtnQueryBankBalanceByFuture(userData: ctypes.c_void_p, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField)):
    """回调函数实现: 期货发起查询银行余额错误回报"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnErrRtnQueryBankBalanceByFuture(pReqQueryAccount.contents if pReqQueryAccount else None, pRspInfo.contents if pRspInfo else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromBankToFutureByFuture(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 期货发起冲正银行转期货请求，银行处理完毕后报盘发回的通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromBankToFutureByFuture(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnRepealFromFutureToBankByFuture(userData: ctypes.c_void_p, pRspRepeal: ctypes.POINTER(CThostFtdcRspRepealField)):
    """回调函数实现: 期货发起冲正期货转银行请求，银行处理完毕后报盘发回的通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnRepealFromFutureToBankByFuture(pRspRepeal.contents if pRspRepeal else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspFromBankToFutureByFuture(userData: ctypes.c_void_p, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 期货发起银行资金转期货应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspFromBankToFutureByFuture(pReqTransfer.contents if pReqTransfer else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspFromFutureToBankByFuture(userData: ctypes.c_void_p, pReqTransfer: ctypes.POINTER(CThostFtdcReqTransferField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 期货发起期货资金转银行应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspFromFutureToBankByFuture(pReqTransfer.contents if pReqTransfer else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQueryBankAccountMoneyByFuture(userData: ctypes.c_void_p, pReqQueryAccount: ctypes.POINTER(CThostFtdcReqQueryAccountField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 期货发起查询银行余额应答"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQueryBankAccountMoneyByFuture(pReqQueryAccount.contents if pReqQueryAccount else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnOpenAccountByBank(userData: ctypes.c_void_p, pOpenAccount: ctypes.POINTER(CThostFtdcOpenAccountField)):
    """回调函数实现: 银行发起银期开户通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnOpenAccountByBank(pOpenAccount.contents if pOpenAccount else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnCancelAccountByBank(userData: ctypes.c_void_p, pCancelAccount: ctypes.POINTER(CThostFtdcCancelAccountField)):
    """回调函数实现: 银行发起银期销户通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnCancelAccountByBank(pCancelAccount.contents if pCancelAccount else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRtnChangeAccountByBank(userData: ctypes.c_void_p, pChangeAccount: ctypes.POINTER(CThostFtdcChangeAccountField)):
    """回调函数实现: 银行发起变更银行账号通知"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRtnChangeAccountByBank(pChangeAccount.contents if pChangeAccount else None)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryClassifiedInstrument(userData: ctypes.c_void_p, pInstrument: ctypes.POINTER(CThostFtdcInstrumentField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求查询分类合约响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryClassifiedInstrument(pInstrument.contents if pInstrument else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryCombPromotionParam(userData: ctypes.c_void_p, pCombPromotionParam: ctypes.POINTER(CThostFtdcCombPromotionParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 请求组合优惠比例响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryCombPromotionParam(pCombPromotionParam.contents if pCombPromotionParam else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRiskSettleInvstPosition(userData: ctypes.c_void_p, pRiskSettleInvstPosition: ctypes.POINTER(CThostFtdcRiskSettleInvstPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者风险结算持仓查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRiskSettleInvstPosition(pRiskSettleInvstPosition.contents if pRiskSettleInvstPosition else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRiskSettleProductStatus(userData: ctypes.c_void_p, pRiskSettleProductStatus: ctypes.POINTER(CThostFtdcRiskSettleProductStatusField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 风险结算产品查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRiskSettleProductStatus(pRiskSettleProductStatus.contents if pRiskSettleProductStatus else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMFutureParameter(userData: ctypes.c_void_p, pSPBMFutureParameter: ctypes.POINTER(CThostFtdcSPBMFutureParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM期货合约参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMFutureParameter(pSPBMFutureParameter.contents if pSPBMFutureParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMOptionParameter(userData: ctypes.c_void_p, pSPBMOptionParameter: ctypes.POINTER(CThostFtdcSPBMOptionParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM期权合约参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMOptionParameter(pSPBMOptionParameter.contents if pSPBMOptionParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMIntraParameter(userData: ctypes.c_void_p, pSPBMIntraParameter: ctypes.POINTER(CThostFtdcSPBMIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM品种内对锁仓折扣参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMIntraParameter(pSPBMIntraParameter.contents if pSPBMIntraParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMInterParameter(userData: ctypes.c_void_p, pSPBMInterParameter: ctypes.POINTER(CThostFtdcSPBMInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM跨品种抵扣参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMInterParameter(pSPBMInterParameter.contents if pSPBMInterParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMPortfDefinition(userData: ctypes.c_void_p, pSPBMPortfDefinition: ctypes.POINTER(CThostFtdcSPBMPortfDefinitionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM组合保证金套餐查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMPortfDefinition(pSPBMPortfDefinition.contents if pSPBMPortfDefinition else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMInvestorPortfDef(userData: ctypes.c_void_p, pSPBMInvestorPortfDef: ctypes.POINTER(CThostFtdcSPBMInvestorPortfDefField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者SPBM套餐选择查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMInvestorPortfDef(pSPBMInvestorPortfDef.contents if pSPBMInvestorPortfDef else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorPortfMarginRatio(userData: ctypes.c_void_p, pInvestorPortfMarginRatio: ctypes.POINTER(CThostFtdcInvestorPortfMarginRatioField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者新型组合保证金系数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorPortfMarginRatio(pInvestorPortfMarginRatio.contents if pInvestorPortfMarginRatio else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorProdSPBMDetail(userData: ctypes.c_void_p, pInvestorProdSPBMDetail: ctypes.POINTER(CThostFtdcInvestorProdSPBMDetailField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者产品SPBM明细查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorProdSPBMDetail(pInvestorProdSPBMDetail.contents if pInvestorProdSPBMDetail else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorCommoditySPMMMargin(userData: ctypes.c_void_p, pInvestorCommoditySPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommoditySPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者商品组SPMM记录查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorCommoditySPMMMargin(pInvestorCommoditySPMMMargin.contents if pInvestorCommoditySPMMMargin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorCommodityGroupSPMMMargin(userData: ctypes.c_void_p, pInvestorCommodityGroupSPMMMargin: ctypes.POINTER(CThostFtdcInvestorCommodityGroupSPMMMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者商品群SPMM记录查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorCommodityGroupSPMMMargin(pInvestorCommodityGroupSPMMMargin.contents if pInvestorCommodityGroupSPMMMargin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPMMInstParam(userData: ctypes.c_void_p, pSPMMInstParam: ctypes.POINTER(CThostFtdcSPMMInstParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPMM合约参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPMMInstParam(pSPMMInstParam.contents if pSPMMInstParam else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPMMProductParam(userData: ctypes.c_void_p, pSPMMProductParam: ctypes.POINTER(CThostFtdcSPMMProductParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPMM产品参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPMMProductParam(pSPMMProductParam.contents if pSPMMProductParam else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQrySPBMAddOnInterParameter(userData: ctypes.c_void_p, pSPBMAddOnInterParameter: ctypes.POINTER(CThostFtdcSPBMAddOnInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: SPBM附加跨品种抵扣参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQrySPBMAddOnInterParameter(pSPBMAddOnInterParameter.contents if pSPBMAddOnInterParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSCombProductInfo(userData: ctypes.c_void_p, pRCAMSCombProductInfo: ctypes.POINTER(CThostFtdcRCAMSCombProductInfoField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS产品组合信息查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSCombProductInfo(pRCAMSCombProductInfo.contents if pRCAMSCombProductInfo else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSInstrParameter(userData: ctypes.c_void_p, pRCAMSInstrParameter: ctypes.POINTER(CThostFtdcRCAMSInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS同合约风险对冲参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSInstrParameter(pRCAMSInstrParameter.contents if pRCAMSInstrParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSIntraParameter(userData: ctypes.c_void_p, pRCAMSIntraParameter: ctypes.POINTER(CThostFtdcRCAMSIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS品种内风险对冲参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSIntraParameter(pRCAMSIntraParameter.contents if pRCAMSIntraParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSInterParameter(userData: ctypes.c_void_p, pRCAMSInterParameter: ctypes.POINTER(CThostFtdcRCAMSInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS跨品种风险折抵参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSInterParameter(pRCAMSInterParameter.contents if pRCAMSInterParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSShortOptAdjustParam(userData: ctypes.c_void_p, pRCAMSShortOptAdjustParam: ctypes.POINTER(CThostFtdcRCAMSShortOptAdjustParamField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS空头期权风险调整参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSShortOptAdjustParam(pRCAMSShortOptAdjustParam.contents if pRCAMSShortOptAdjustParam else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRCAMSInvestorCombPosition(userData: ctypes.c_void_p, pRCAMSInvestorCombPosition: ctypes.POINTER(CThostFtdcRCAMSInvestorCombPositionField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RCAMS策略组合持仓查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRCAMSInvestorCombPosition(pRCAMSInvestorCombPosition.contents if pRCAMSInvestorCombPosition else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorProdRCAMSMargin(userData: ctypes.c_void_p, pInvestorProdRCAMSMargin: ctypes.POINTER(CThostFtdcInvestorProdRCAMSMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者品种RCAMS保证金查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorProdRCAMSMargin(pInvestorProdRCAMSMargin.contents if pInvestorProdRCAMSMargin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRULEInstrParameter(userData: ctypes.c_void_p, pRULEInstrParameter: ctypes.POINTER(CThostFtdcRULEInstrParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RULE合约保证金参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRULEInstrParameter(pRULEInstrParameter.contents if pRULEInstrParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRULEIntraParameter(userData: ctypes.c_void_p, pRULEIntraParameter: ctypes.POINTER(CThostFtdcRULEIntraParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RULE品种内对锁仓折扣参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRULEIntraParameter(pRULEIntraParameter.contents if pRULEIntraParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryRULEInterParameter(userData: ctypes.c_void_p, pRULEInterParameter: ctypes.POINTER(CThostFtdcRULEInterParameterField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: RULE跨品种抵扣参数查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryRULEInterParameter(pRULEInterParameter.contents if pRULEInterParameter else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorProdRULEMargin(userData: ctypes.c_void_p, pInvestorProdRULEMargin: ctypes.POINTER(CThostFtdcInvestorProdRULEMarginField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者产品RULE保证金查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorProdRULEMargin(pInvestorProdRULEMargin.contents if pInvestorProdRULEMargin else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

def _go_trader_OnRspQryInvestorPortfSetting(userData: ctypes.c_void_p, pInvestorPortfSetting: ctypes.POINTER(CThostFtdcInvestorPortfSettingField), pRspInfo: ctypes.POINTER(CThostFtdcRspInfoField), nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
    """回调函数实现: 投资者投资者新组保设置查询响应"""
    user_data = userData.value if hasattr(userData, 'value') else userData
    api = _get_trader_instance(user_data)
    if api is None or api._spi is None:
        return

    try:
        api._spi.OnRspQryInvestorPortfSetting(pInvestorPortfSetting.contents if pInvestorPortfSetting else None, pRspInfo.contents if pRspInfo else None, nRequestID, bIsLast)
    except Exception as e:
        import traceback
        traceback.print_exc()

# ========== 回调注册函数 ==========

def _register_trader_callback_impl(spi_handle: ctypes.c_void_p, lib: ctypes.CDLL, callback_name: str, spi, user_data: int):
    """注册回调函数到 C SPI（内部实现）"""
    # 回调函数映射表
    callback_map = {
        "FrontConnected": (_go_trader_OnFrontConnected, ctypes.CFUNCTYPE(None, ctypes.c_void_p)),
        "FrontDisconnected": (_go_trader_OnFrontDisconnected, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)),
        "HeartBeatWarning": (_go_trader_OnHeartBeatWarning, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.c_int32)),
        "RspAuthenticate": (_go_trader_OnRspAuthenticate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspAuthenticateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUserLogin": (_go_trader_OnRspUserLogin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserLoginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUserLogout": (_go_trader_OnRspUserLogout, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserLogoutField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUserPasswordUpdate": (_go_trader_OnRspUserPasswordUpdate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcUserPasswordUpdateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspTradingAccountPasswordUpdate": (_go_trader_OnRspTradingAccountPasswordUpdate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountPasswordUpdateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspUserAuthMethod": (_go_trader_OnRspUserAuthMethod, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspUserAuthMethodField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspGenUserCaptcha": (_go_trader_OnRspGenUserCaptcha, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspGenUserCaptchaField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspGenUserText": (_go_trader_OnRspGenUserText, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspGenUserTextField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspOrderInsert": (_go_trader_OnRspOrderInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspParkedOrderInsert": (_go_trader_OnRspParkedOrderInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspParkedOrderAction": (_go_trader_OnRspParkedOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspOrderAction": (_go_trader_OnRspOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryMaxOrderVolume": (_go_trader_OnRspQryMaxOrderVolume, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQryMaxOrderVolumeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspSettlementInfoConfirm": (_go_trader_OnRspSettlementInfoConfirm, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspRemoveParkedOrder": (_go_trader_OnRspRemoveParkedOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspRemoveParkedOrderAction": (_go_trader_OnRspRemoveParkedOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRemoveParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspExecOrderInsert": (_go_trader_OnRspExecOrderInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspExecOrderAction": (_go_trader_OnRspExecOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspForQuoteInsert": (_go_trader_OnRspForQuoteInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQuoteInsert": (_go_trader_OnRspQuoteInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQuoteAction": (_go_trader_OnRspQuoteAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspBatchOrderAction": (_go_trader_OnRspBatchOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputBatchOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspOptionSelfCloseInsert": (_go_trader_OnRspOptionSelfCloseInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspOptionSelfCloseAction": (_go_trader_OnRspOptionSelfCloseAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspCombActionInsert": (_go_trader_OnRspCombActionInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputCombActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryOrder": (_go_trader_OnRspQryOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTrade": (_go_trader_OnRspQryTrade, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorPosition": (_go_trader_OnRspQryInvestorPosition, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTradingAccount": (_go_trader_OnRspQryTradingAccount, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestor": (_go_trader_OnRspQryInvestor, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTradingCode": (_go_trader_OnRspQryTradingCode, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingCodeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInstrumentMarginRate": (_go_trader_OnRspQryInstrumentMarginRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentMarginRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInstrumentCommissionRate": (_go_trader_OnRspQryInstrumentCommissionRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentCommissionRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryExchange": (_go_trader_OnRspQryExchange, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryProduct": (_go_trader_OnRspQryProduct, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInstrument": (_go_trader_OnRspQryInstrument, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryDepthMarketData": (_go_trader_OnRspQryDepthMarketData, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcDepthMarketDataField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTraderOffer": (_go_trader_OnRspQryTraderOffer, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTraderOfferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySettlementInfo": (_go_trader_OnRspQrySettlementInfo, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTransferBank": (_go_trader_OnRspQryTransferBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTransferBankField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorPositionDetail": (_go_trader_OnRspQryInvestorPositionDetail, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryNotice": (_go_trader_OnRspQryNotice, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcNoticeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySettlementInfoConfirm": (_go_trader_OnRspQrySettlementInfoConfirm, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSettlementInfoConfirmField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorPositionCombineDetail": (_go_trader_OnRspQryInvestorPositionCombineDetail, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPositionCombineDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryCFMMCTradingAccountKey": (_go_trader_OnRspQryCFMMCTradingAccountKey, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCFMMCTradingAccountKeyField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryEWarrantOffset": (_go_trader_OnRspQryEWarrantOffset, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcEWarrantOffsetField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorProductGroupMargin": (_go_trader_OnRspQryInvestorProductGroupMargin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProductGroupMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryExchangeMarginRate": (_go_trader_OnRspQryExchangeMarginRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeMarginRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryExchangeMarginRateAdjust": (_go_trader_OnRspQryExchangeMarginRateAdjust, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeMarginRateAdjustField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryExchangeRate": (_go_trader_OnRspQryExchangeRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExchangeRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySecAgentACIDMap": (_go_trader_OnRspQrySecAgentACIDMap, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentACIDMapField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryProductExchRate": (_go_trader_OnRspQryProductExchRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductExchRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryProductGroup": (_go_trader_OnRspQryProductGroup, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcProductGroupField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryMMInstrumentCommissionRate": (_go_trader_OnRspQryMMInstrumentCommissionRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMMInstrumentCommissionRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryMMOptionInstrCommRate": (_go_trader_OnRspQryMMOptionInstrCommRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcMMOptionInstrCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInstrumentOrderCommRate": (_go_trader_OnRspQryInstrumentOrderCommRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentOrderCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySecAgentTradingAccount": (_go_trader_OnRspQrySecAgentTradingAccount, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySecAgentCheckMode": (_go_trader_OnRspQrySecAgentCheckMode, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentCheckModeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySecAgentTradeInfo": (_go_trader_OnRspQrySecAgentTradeInfo, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSecAgentTradeInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryOptionInstrTradeCost": (_go_trader_OnRspQryOptionInstrTradeCost, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionInstrTradeCostField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryOptionInstrCommRate": (_go_trader_OnRspQryOptionInstrCommRate, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionInstrCommRateField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryExecOrder": (_go_trader_OnRspQryExecOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryForQuote": (_go_trader_OnRspQryForQuote, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryQuote": (_go_trader_OnRspQryQuote, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryOptionSelfClose": (_go_trader_OnRspQryOptionSelfClose, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestUnit": (_go_trader_OnRspQryInvestUnit, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestUnitField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryCombInstrumentGuard": (_go_trader_OnRspQryCombInstrumentGuard, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombInstrumentGuardField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryCombAction": (_go_trader_OnRspQryCombAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTransferSerial": (_go_trader_OnRspQryTransferSerial, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTransferSerialField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryAccountregister": (_go_trader_OnRspQryAccountregister, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcAccountregisterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspError": (_go_trader_OnRspError, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RtnOrder": (_go_trader_OnRtnOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderField))),
        "RtnTrade": (_go_trader_OnRtnTrade, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradeField))),
        "ErrRtnOrderInsert": (_go_trader_OnErrRtnOrderInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOrderField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnOrderAction": (_go_trader_OnErrRtnOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnInstrumentStatus": (_go_trader_OnRtnInstrumentStatus, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentStatusField))),
        "RtnBulletin": (_go_trader_OnRtnBulletin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBulletinField))),
        "RtnTradingNotice": (_go_trader_OnRtnTradingNotice, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingNoticeInfoField))),
        "RtnErrorConditionalOrder": (_go_trader_OnRtnErrorConditionalOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcErrorConditionalOrderField))),
        "RtnExecOrder": (_go_trader_OnRtnExecOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderField))),
        "ErrRtnExecOrderInsert": (_go_trader_OnErrRtnExecOrderInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputExecOrderField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnExecOrderAction": (_go_trader_OnErrRtnExecOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcExecOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnForQuoteInsert": (_go_trader_OnErrRtnForQuoteInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputForQuoteField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnQuote": (_go_trader_OnRtnQuote, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteField))),
        "ErrRtnQuoteInsert": (_go_trader_OnErrRtnQuoteInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputQuoteField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnQuoteAction": (_go_trader_OnErrRtnQuoteAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQuoteActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnForQuoteRsp": (_go_trader_OnRtnForQuoteRsp, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcForQuoteRspField))),
        "RtnCFMMCTradingAccountToken": (_go_trader_OnRtnCFMMCTradingAccountToken, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCFMMCTradingAccountTokenField))),
        "ErrRtnBatchOrderAction": (_go_trader_OnErrRtnBatchOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBatchOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnOptionSelfClose": (_go_trader_OnRtnOptionSelfClose, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseField))),
        "ErrRtnOptionSelfCloseInsert": (_go_trader_OnErrRtnOptionSelfCloseInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputOptionSelfCloseField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnOptionSelfCloseAction": (_go_trader_OnErrRtnOptionSelfCloseAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOptionSelfCloseActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnCombAction": (_go_trader_OnRtnCombAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombActionField))),
        "ErrRtnCombActionInsert": (_go_trader_OnErrRtnCombActionInsert, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInputCombActionField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RspQryContractBank": (_go_trader_OnRspQryContractBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcContractBankField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryParkedOrder": (_go_trader_OnRspQryParkedOrder, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryParkedOrderAction": (_go_trader_OnRspQryParkedOrderAction, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcParkedOrderActionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryTradingNotice": (_go_trader_OnRspQryTradingNotice, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcTradingNoticeField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryBrokerTradingParams": (_go_trader_OnRspQryBrokerTradingParams, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBrokerTradingParamsField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryBrokerTradingAlgos": (_go_trader_OnRspQryBrokerTradingAlgos, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcBrokerTradingAlgosField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQueryCFMMCTradingAccountToken": (_go_trader_OnRspQueryCFMMCTradingAccountToken, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcQueryCFMMCTradingAccountTokenField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RtnFromBankToFutureByBank": (_go_trader_OnRtnFromBankToFutureByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))),
        "RtnFromFutureToBankByBank": (_go_trader_OnRtnFromFutureToBankByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))),
        "RtnRepealFromBankToFutureByBank": (_go_trader_OnRtnRepealFromBankToFutureByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RtnRepealFromFutureToBankByBank": (_go_trader_OnRtnRepealFromFutureToBankByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RtnFromBankToFutureByFuture": (_go_trader_OnRtnFromBankToFutureByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))),
        "RtnFromFutureToBankByFuture": (_go_trader_OnRtnFromFutureToBankByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspTransferField))),
        "RtnRepealFromBankToFutureByFutureManual": (_go_trader_OnRtnRepealFromBankToFutureByFutureManual, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RtnRepealFromFutureToBankByFutureManual": (_go_trader_OnRtnRepealFromFutureToBankByFutureManual, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RtnQueryBankBalanceByFuture": (_go_trader_OnRtnQueryBankBalanceByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcNotifyQueryAccountField))),
        "ErrRtnBankToFutureByFuture": (_go_trader_OnErrRtnBankToFutureByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnFutureToBankByFuture": (_go_trader_OnErrRtnFutureToBankByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnRepealBankToFutureByFutureManual": (_go_trader_OnErrRtnRepealBankToFutureByFutureManual, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqRepealField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnRepealFutureToBankByFutureManual": (_go_trader_OnErrRtnRepealFutureToBankByFutureManual, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqRepealField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "ErrRtnQueryBankBalanceByFuture": (_go_trader_OnErrRtnQueryBankBalanceByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqQueryAccountField), ctypes.POINTER(CThostFtdcRspInfoField))),
        "RtnRepealFromBankToFutureByFuture": (_go_trader_OnRtnRepealFromBankToFutureByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RtnRepealFromFutureToBankByFuture": (_go_trader_OnRtnRepealFromFutureToBankByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRspRepealField))),
        "RspFromBankToFutureByFuture": (_go_trader_OnRspFromBankToFutureByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspFromFutureToBankByFuture": (_go_trader_OnRspFromFutureToBankByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqTransferField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQueryBankAccountMoneyByFuture": (_go_trader_OnRspQueryBankAccountMoneyByFuture, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcReqQueryAccountField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RtnOpenAccountByBank": (_go_trader_OnRtnOpenAccountByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcOpenAccountField))),
        "RtnCancelAccountByBank": (_go_trader_OnRtnCancelAccountByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCancelAccountField))),
        "RtnChangeAccountByBank": (_go_trader_OnRtnChangeAccountByBank, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcChangeAccountField))),
        "RspQryClassifiedInstrument": (_go_trader_OnRspQryClassifiedInstrument, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInstrumentField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryCombPromotionParam": (_go_trader_OnRspQryCombPromotionParam, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcCombPromotionParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRiskSettleInvstPosition": (_go_trader_OnRspQryRiskSettleInvstPosition, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRiskSettleInvstPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRiskSettleProductStatus": (_go_trader_OnRspQryRiskSettleProductStatus, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRiskSettleProductStatusField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMFutureParameter": (_go_trader_OnRspQrySPBMFutureParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMFutureParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMOptionParameter": (_go_trader_OnRspQrySPBMOptionParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMOptionParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMIntraParameter": (_go_trader_OnRspQrySPBMIntraParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMInterParameter": (_go_trader_OnRspQrySPBMInterParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMPortfDefinition": (_go_trader_OnRspQrySPBMPortfDefinition, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMPortfDefinitionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMInvestorPortfDef": (_go_trader_OnRspQrySPBMInvestorPortfDef, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMInvestorPortfDefField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorPortfMarginRatio": (_go_trader_OnRspQryInvestorPortfMarginRatio, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPortfMarginRatioField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorProdSPBMDetail": (_go_trader_OnRspQryInvestorProdSPBMDetail, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdSPBMDetailField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorCommoditySPMMMargin": (_go_trader_OnRspQryInvestorCommoditySPMMMargin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorCommoditySPMMMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorCommodityGroupSPMMMargin": (_go_trader_OnRspQryInvestorCommodityGroupSPMMMargin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorCommodityGroupSPMMMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPMMInstParam": (_go_trader_OnRspQrySPMMInstParam, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPMMInstParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPMMProductParam": (_go_trader_OnRspQrySPMMProductParam, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPMMProductParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQrySPBMAddOnInterParameter": (_go_trader_OnRspQrySPBMAddOnInterParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcSPBMAddOnInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSCombProductInfo": (_go_trader_OnRspQryRCAMSCombProductInfo, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSCombProductInfoField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSInstrParameter": (_go_trader_OnRspQryRCAMSInstrParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInstrParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSIntraParameter": (_go_trader_OnRspQryRCAMSIntraParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSInterParameter": (_go_trader_OnRspQryRCAMSInterParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSShortOptAdjustParam": (_go_trader_OnRspQryRCAMSShortOptAdjustParam, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSShortOptAdjustParamField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRCAMSInvestorCombPosition": (_go_trader_OnRspQryRCAMSInvestorCombPosition, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRCAMSInvestorCombPositionField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorProdRCAMSMargin": (_go_trader_OnRspQryInvestorProdRCAMSMargin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdRCAMSMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRULEInstrParameter": (_go_trader_OnRspQryRULEInstrParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEInstrParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRULEIntraParameter": (_go_trader_OnRspQryRULEIntraParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEIntraParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryRULEInterParameter": (_go_trader_OnRspQryRULEInterParameter, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcRULEInterParameterField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorProdRULEMargin": (_go_trader_OnRspQryInvestorProdRULEMargin, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorProdRULEMarginField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
        "RspQryInvestorPortfSetting": (_go_trader_OnRspQryInvestorPortfSetting, ctypes.CFUNCTYPE(None, ctypes.c_void_p, ctypes.POINTER(CThostFtdcInvestorPortfSettingField), ctypes.POINTER(CThostFtdcRspInfoField), ctypes.c_int32, ctypes.c_bool)),
    }

    if callback_name not in callback_map:
        return

    callback_func, callback_type = callback_map[callback_name]

    # 创建 CFUNCTYPE 回调实例
    c_callback = callback_type(callback_func)

    # 注册到 C SPI
    func_name = f"TraderSpiSetOn{callback_name}"
    if hasattr(lib, func_name):
        func = getattr(lib, func_name)
        func.argtypes = [ctypes.c_void_p, callback_type]
        func.restype = None
        func(spi_handle, c_callback)

    # 保存回调引用，防止被 GC 回收
    if not hasattr(spi, "_callbacks"):
        spi._callbacks = []
    spi._callbacks.append(c_callback)
