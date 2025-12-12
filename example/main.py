#!/usr/bin/env python3
"""
CTP Python 包装示例程序

演示如何使用 pyctp 进行行情订阅和交易操作
"""

import os
import sys
import time
import queue
from pathlib import Path

# 添加 pyctp 到路径（如果不在同一目录）
sys.path.insert(0, str(Path(__file__).parent.parent))

try:
    from dotenv import load_dotenv
    load_dotenv()
except ImportError:
    print("警告: python-dotenv 未安装，将使用系统环境变量")

import pyctp
from pyctp import (
    MdApi, MdSpi, DefaultMdSpi,
    TraderApi, TraderSpi, DefaultTraderSpi,
    CThostFtdcReqUserLoginField,
    CThostFtdcRspUserLoginField,
    CThostFtdcRspInfoField,
    CThostFtdcDepthMarketDataField,
    CThostFtdcReqAuthenticateField,
    CThostFtdcRspAuthenticateField,
    CThostFtdcOrderField,
    string_to_bytes, bytes_to_string, gb18030,
    ctypes
)
from pyctp.datatype import (
    TThostFtdcUserIDType, TThostFtdcPasswordType, TThostFtdcBrokerIDType,
    TThostFtdcAppIDType, TThostFtdcAuthCodeType
)


class MyMdSpi(DefaultMdSpi):
    """自定义行情回调实现"""
    
    def OnFrontConnected(self):
        """连接成功回调"""
        print("行情服务器连接成功")
    
    def OnFrontDisconnected(self, nReason: ctypes.c_int32):
        """连接断开回调"""
        print(f"行情服务器断开连接，原因: {nReason.value}")
    
    def OnRspUserLogin(self, pRspUserLogin, pRspInfo, nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登录响应"""
        if pRspInfo and pRspInfo.ErrorID != 0:
            error_msg = gb18030(bytes(pRspInfo.ErrorMsg))
            print(f"登录失败: {error_msg}")
            return
        
        if pRspUserLogin:
            trading_day = bytes_to_string(bytes(pRspUserLogin.TradingDay))
            print(f"登录成功，交易日: {trading_day}")
    
    def OnRtnDepthMarketData(self, pDepthMarketData):
        """行情数据回调"""
        if not pDepthMarketData:
            return
        
        instrument_id = bytes_to_string(bytes(pDepthMarketData.InstrumentID))
        last_price = pDepthMarketData.LastPrice
        print(f"行情更新: {instrument_id}, 最新价: {last_price:.2f}")


class MyTraderSpi(DefaultTraderSpi):
    """自定义交易回调实现"""
    
    def __init__(self):
        super().__init__()
        self.auth_queue = queue.Queue()
    
    def OnFrontConnected(self):
        """连接成功回调"""
        print("交易服务器连接成功")
    
    def OnRspAuthenticate(self, pRspAuthenticateField, pRspInfo, nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """认证响应"""
        if pRspInfo and pRspInfo.ErrorID != 0:
            error_msg = gb18030(bytes(pRspInfo.ErrorMsg))
            print(f"认证失败: {error_msg}")
            self.auth_queue.put(False)
            return
        
        print("认证成功")
        self.auth_queue.put(True)
    
    def OnRspUserLogin(self, pRspUserLogin, pRspInfo, nRequestID: ctypes.c_int32, bIsLast: ctypes.c_bool):
        """登录响应"""
        if pRspInfo and pRspInfo.ErrorID != 0:
            error_msg = gb18030(bytes(pRspInfo.ErrorMsg))
            print(f"登录失败: {error_msg}")
            return
        
        if pRspUserLogin:
            trading_day = bytes_to_string(bytes(pRspUserLogin.TradingDay))
            print(f"登录成功，交易日: {trading_day}")
    
    def OnRtnOrder(self, pOrder):
        """报单通知"""
        if not pOrder:
            return
        
        instrument_id = bytes_to_string(bytes(pOrder.InstrumentID))
        order_status = chr(pOrder.OrderStatus) if pOrder.OrderStatus else '?'
        print(f"报单状态更新: {instrument_id}, 状态: {order_status}")


def main():
    # 从环境变量读取配置
    user_id = os.getenv("CTP_USER_ID", "")
    password = os.getenv("CTP_PASSWORD", "")
    broker_id = os.getenv("CTP_BROKER_ID", "")
    md_front = os.getenv("CTP_MD_FRONT", "tcp://182.254.243.31:30011")  # 默认 simnow 行情地址
    trader_front = os.getenv("CTP_TRADER_FRONT", "tcp://182.254.243.31:30001")  # 默认 simnow 交易地址
    app_id = os.getenv("CTP_APP_ID", "")
    auth_code = os.getenv("CTP_AUTH_CODE", "")
    
    # 检查必要的环境变量
    if not user_id or not password or not broker_id:
        print("错误: 请设置环境变量: CTP_USER_ID, CTP_PASSWORD, CTP_BROKER_ID")
        print("或者创建 .env 文件并设置这些变量")
        sys.exit(1)
    
    # ========== 行情 API 示例 ==========
    print("=== 行情 API 示例 ===")
    
    # 注意：库会在首次创建 API 时自动加载
    # 可以通过设置 CTP_LIB_PATH 环境变量指定库路径，否则使用默认路径
    
    # 1. 创建行情 API 实例（会自动加载库）
    flow_path = "./flow"
    # 确保目录存在
    Path(flow_path).mkdir(parents=True, exist_ok=True)
    
    try:
        md_api = MdApi(flow_path, False, False)
    except Exception as e:
        print(f"创建行情 API 失败: {e}")
        print("请确保库文件在 ../libs 或 ./libs 目录下，或设置 CTP_LIB_PATH 环境变量")
        sys.exit(1)
    
    # 2. 创建并设置回调
    md_spi = MyMdSpi()
    md_api.set_spi(md_spi)
    print(f"mdApi.GetApiVersion(): {md_api.GetApiVersion()}")
    
    # 3. 注册前置地址
    md_api.RegisterFront(md_front)
    
    # 4. 初始化
    md_api.Init()
    
    # 5. 等待连接
    time.sleep(1)
    
    # 6. 登录
    login_req = CThostFtdcReqUserLoginField()
    # 使用 string_to_bytes 填充固定长度字节数组字段
    # 注意：ctypes.c_char * N 字段可以直接接受 bytes 对象
    login_req.UserID = string_to_bytes(user_id, TThostFtdcUserIDType._length_)
    login_req.Password = string_to_bytes(password, TThostFtdcPasswordType._length_)
    login_req.BrokerID = string_to_bytes(broker_id, TThostFtdcBrokerIDType._length_)
    
    request_id = ctypes.c_int32(1)
    ret = md_api.ReqUserLogin(login_req, request_id)
    if ret != 0:
        print(f"登录请求失败，错误码: {ret}")
    
    # 7. 等待登录响应
    time.sleep(1)
    
    # 8. 订阅行情（需要替换为真实的合约代码）
    instruments = ["rb2605", "lc2605"]
    ret = md_api.SubscribeMarketData(instruments, ctypes.c_int32(len(instruments)))
    if ret != 0:
        print(f"订阅行情失败，错误码: {ret}")
    
    print("等待行情数据...")
    time.sleep(1)
    
    # 注意：不在这里释放行情API，让它继续运行以接收行情数据
    # md_api.Release() 将在程序结束时调用
    
    # ========== 交易 API 示例 ==========
    print("\n=== 交易 API 示例 ===")
    
    # 1. 创建交易 API 实例
    trader_flow_path = "./flow_trader"
    # 确保目录存在
    Path(trader_flow_path).mkdir(parents=True, exist_ok=True)
    
    trader_api = None
    try:
        trader_api = TraderApi(trader_flow_path)
    except Exception as e:
        print(f"创建交易 API 失败: {e}")
        print("请确保库文件在 ../libs 或 ./libs 目录下，或设置 CTP_LIB_PATH 环境变量")
        # 即使交易API创建失败，也要释放行情API
        md_api.Release()
        sys.exit(1)
    
    try:
        # 2. 创建并设置回调
        trader_spi = MyTraderSpi()
        trader_api.set_spi(trader_spi)
        print(f"traderApi.GetApiVersion(): {trader_api.GetApiVersion()}")
        print(f"pyctp.GetDataCollectApiVersion(): {pyctp.GetDataCollectApiVersion()}")
        
        # 3. 注册前置地址
        trader_api.RegisterFront(trader_front)
        
        # 4. 初始化
        trader_api.Init()
        
        # 5. 等待连接
        time.sleep(2)
        
        # 6. 穿透式认证（如果配置了 AppID 和 AuthCode）
        if app_id and auth_code:
            print("开始穿透式认证...")
            auth_req = CThostFtdcReqAuthenticateField()
            auth_req.BrokerID = string_to_bytes(broker_id, TThostFtdcBrokerIDType._length_)
            auth_req.UserID = string_to_bytes(user_id, TThostFtdcUserIDType._length_)
            auth_req.AppID = string_to_bytes(app_id, TThostFtdcAppIDType._length_)
            auth_req.AuthCode = string_to_bytes(auth_code, TThostFtdcAuthCodeType._length_)
            # UserProductInfo 可以留空或设置产品信息
            # auth_req.UserProductInfo = (ctypes.c_char * 33)(*string_to_bytes("your_product_info", 33))
            
            auth_request_id = ctypes.c_int32(0)
            ret = trader_api.ReqAuthenticate(auth_req, auth_request_id)
            if ret != 0:
                print(f"认证请求失败，错误码: {ret}")
            else:
                # 等待认证响应（设置超时）
                try:
                    auth_success = trader_spi.auth_queue.get(timeout=10)
                    if not auth_success:
                        print("认证失败，无法继续登录")
                        # 清理资源后退出
                        trader_api.Release()
                        md_api.Release()
                        return
                    print("认证成功，准备登录...")
                except queue.Empty:
                    print("认证超时")
                    # 清理资源后退出
                    trader_api.Release()
                    md_api.Release()
                    return
        else:
            print("未配置 AppID 和 AuthCode，跳过穿透式认证（仅适用于仿真环境）")
        
        # 7. 登录
        trader_login_req = CThostFtdcReqUserLoginField()
        trader_login_req.UserID = string_to_bytes(user_id, TThostFtdcUserIDType._length_)
        trader_login_req.Password = string_to_bytes(password, TThostFtdcPasswordType._length_)
        trader_login_req.BrokerID = string_to_bytes(broker_id, TThostFtdcBrokerIDType._length_)
        
        ret = trader_api.ReqUserLogin(trader_login_req, request_id)
        if ret != 0:
            print(f"登录请求失败，错误码: {ret}")
        
        print("等待交易响应...")
        time.sleep(10)
        
    finally:
        # 清理交易API
        trader_api.Release()
    
    # 清理行情API（在程序最后释放，确保两个API可以同时运行）
    md_api.Release()
    
    print("\n示例程序结束")


if __name__ == "__main__":
    main()
