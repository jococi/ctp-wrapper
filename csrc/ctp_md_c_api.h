/**
 * CTP Md API - 纯 C 接口封装
 * 
 * 自动生成，请勿手动修改
 * 特性：
 *   - 纯 C 接口，无 C++ 依赖
 *   - 不透明指针句柄
 *   - 回调携带 userData，支持多实例
 *   - 驼峰命名风格
 */

#ifndef CTP_MD_C_API_H
#define CTP_MD_C_API_H

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
typedef struct MdApi_t* MdApiHandle;
typedef struct MdSpi_t* MdSpiHandle;

// ========== 回调函数类型（带 userData） ==========
// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
typedef void (*MdOnFrontConnectedCallback)(void* userData);
// 0x2003 收到错误报文
typedef void (*MdOnFrontDisconnectedCallback)(void* userData, int nReason);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
typedef void (*MdOnHeartBeatWarningCallback)(void* userData, int nTimeLapse);
// 登录请求响应
typedef void (*MdOnRspUserLoginCallback)(void* userData, CThostFtdcRspUserLoginField* pRspUserLogin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 登出请求响应
typedef void (*MdOnRspUserLogoutCallback)(void* userData, CThostFtdcUserLogoutField* pUserLogout, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 请求查询组播合约响应
typedef void (*MdOnRspQryMulticastInstrumentCallback)(void* userData, CThostFtdcMulticastInstrumentField* pMulticastInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 错误应答
typedef void (*MdOnRspErrorCallback)(void* userData, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 订阅行情应答
typedef void (*MdOnRspSubMarketDataCallback)(void* userData, CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 取消订阅行情应答
typedef void (*MdOnRspUnSubMarketDataCallback)(void* userData, CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 订阅询价应答
typedef void (*MdOnRspSubForQuoteRspCallback)(void* userData, CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 取消订阅询价应答
typedef void (*MdOnRspUnSubForQuoteRspCallback)(void* userData, CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast);
// 深度行情通知
typedef void (*MdOnRtnDepthMarketDataCallback)(void* userData, CThostFtdcDepthMarketDataField* pDepthMarketData);
// 询价通知
typedef void (*MdOnRtnForQuoteRspCallback)(void* userData, CThostFtdcForQuoteRspField* pForQuoteRsp);

// ========== 回调表结构（便于批量设置） ==========
typedef struct {
    void* userData;
    MdOnFrontConnectedCallback onFrontConnected;
    MdOnFrontDisconnectedCallback onFrontDisconnected;
    MdOnHeartBeatWarningCallback onHeartBeatWarning;
    MdOnRspUserLoginCallback onRspUserLogin;
    MdOnRspUserLogoutCallback onRspUserLogout;
    MdOnRspQryMulticastInstrumentCallback onRspQryMulticastInstrument;
    MdOnRspErrorCallback onRspError;
    MdOnRspSubMarketDataCallback onRspSubMarketData;
    MdOnRspUnSubMarketDataCallback onRspUnSubMarketData;
    MdOnRspSubForQuoteRspCallback onRspSubForQuoteRsp;
    MdOnRspUnSubForQuoteRspCallback onRspUnSubForQuoteRsp;
    MdOnRtnDepthMarketDataCallback onRtnDepthMarketData;
    MdOnRtnForQuoteRspCallback onRtnForQuoteRsp;
} MdSpiCallbacks;

// ========== Md API 函数 ==========

// 创建MdApi modify for udp marketdata
CTP_API MdApiHandle MdCreateFtdcMdApi(const char* pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast);

// 获取API的版本信息
CTP_API const char * MdGetApiVersion(void);


// 删除接口对象本身
CTP_API void MdRelease(MdApiHandle handle);

// 初始化
CTP_API void MdInit(MdApiHandle handle);

// 等待接口线程结束运行
CTP_API int MdJoin(MdApiHandle handle);

// 获取当前交易日
CTP_API const char * MdGetTradingDay(MdApiHandle handle);

// 注册前置机网络地址
CTP_API void MdRegisterFront(MdApiHandle handle, char* pszFrontAddress);

// 注册名字服务器网络地址
CTP_API void MdRegisterNameServer(MdApiHandle handle, char* pszNsAddress);

// 注册名字服务器用户信息
CTP_API void MdRegisterFensUserInfo(MdApiHandle handle, CThostFtdcFensUserInfoField* pFensUserInfo);

// 订阅行情。
CTP_API int MdSubscribeMarketData(MdApiHandle handle, char* ppInstrumentID[], int nCount);

// 退订行情。
CTP_API int MdUnSubscribeMarketData(MdApiHandle handle, char* ppInstrumentID[], int nCount);

// 订阅询价。
CTP_API int MdSubscribeForQuoteRsp(MdApiHandle handle, char* ppInstrumentID[], int nCount);

// 退订询价。
CTP_API int MdUnSubscribeForQuoteRsp(MdApiHandle handle, char* ppInstrumentID[], int nCount);

// 用户登录请求
CTP_API int MdReqUserLogin(MdApiHandle handle, CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID);

// 登出请求
CTP_API int MdReqUserLogout(MdApiHandle handle, CThostFtdcUserLogoutField* pUserLogout, int nRequestID);

// 请求查询组播合约
CTP_API int MdReqQryMulticastInstrument(MdApiHandle handle, CThostFtdcQryMulticastInstrumentField* pQryMulticastInstrument, int nRequestID);


// ========== Md SPI 函数 ==========

// 创建 SPI 实例
CTP_API MdSpiHandle MdSpiCreate(void* userData);

// 销毁 SPI 实例
CTP_API void MdSpiDestroy(MdSpiHandle spi);

// 注册 SPI 到 API
CTP_API void MdRegisterSpi(MdApiHandle api, MdSpiHandle spi);

// 批量设置回调
CTP_API void MdSpiSetCallbacks(MdSpiHandle spi, const MdSpiCallbacks* callbacks);

// 单独设置回调
// 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
CTP_API void MdSpiSetOnFrontConnected(MdSpiHandle spi, MdOnFrontConnectedCallback callback);
// 0x2003 收到错误报文
CTP_API void MdSpiSetOnFrontDisconnected(MdSpiHandle spi, MdOnFrontDisconnectedCallback callback);
// 心跳超时警告。当长时间未收到报文时，该方法被调用。
CTP_API void MdSpiSetOnHeartBeatWarning(MdSpiHandle spi, MdOnHeartBeatWarningCallback callback);
// 登录请求响应
CTP_API void MdSpiSetOnRspUserLogin(MdSpiHandle spi, MdOnRspUserLoginCallback callback);
// 登出请求响应
CTP_API void MdSpiSetOnRspUserLogout(MdSpiHandle spi, MdOnRspUserLogoutCallback callback);
// 请求查询组播合约响应
CTP_API void MdSpiSetOnRspQryMulticastInstrument(MdSpiHandle spi, MdOnRspQryMulticastInstrumentCallback callback);
// 错误应答
CTP_API void MdSpiSetOnRspError(MdSpiHandle spi, MdOnRspErrorCallback callback);
// 订阅行情应答
CTP_API void MdSpiSetOnRspSubMarketData(MdSpiHandle spi, MdOnRspSubMarketDataCallback callback);
// 取消订阅行情应答
CTP_API void MdSpiSetOnRspUnSubMarketData(MdSpiHandle spi, MdOnRspUnSubMarketDataCallback callback);
// 订阅询价应答
CTP_API void MdSpiSetOnRspSubForQuoteRsp(MdSpiHandle spi, MdOnRspSubForQuoteRspCallback callback);
// 取消订阅询价应答
CTP_API void MdSpiSetOnRspUnSubForQuoteRsp(MdSpiHandle spi, MdOnRspUnSubForQuoteRspCallback callback);
// 深度行情通知
CTP_API void MdSpiSetOnRtnDepthMarketData(MdSpiHandle spi, MdOnRtnDepthMarketDataCallback callback);
// 询价通知
CTP_API void MdSpiSetOnRtnForQuoteRsp(MdSpiHandle spi, MdOnRtnForQuoteRspCallback callback);

#ifdef __cplusplus
}
#endif

#endif // CTP_MD_C_API_H
