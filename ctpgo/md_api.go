package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 行情 API 封装

import (
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// ========== MdSpi 接口 ==========

// MdSpi 行情回调接口
type MdSpi interface {
	OnFrontConnected() // ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontDisconnected(nReason int32) // 0x2003 收到错误报文
	OnHeartBeatWarning(nTimeLapse int32) // 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 登录请求响应
	OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 登出请求响应
	OnRspQryMulticastInstrument(pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 请求查询组播合约响应
	OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 错误应答
	OnRspSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 订阅行情应答
	OnRspUnSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 取消订阅行情应答
	OnRspSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 订阅询价应答
	OnRspUnSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 取消订阅询价应答
	OnRtnDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField) // 深度行情通知
	OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField) // 询价通知
}

// ========== MdApi 结构体 ==========

// MdApi 行情 API 封装
type MdApi struct {
	handle   uintptr
	spi      MdSpi
	userData uintptr
	mu       sync.RWMutex
}

// ========== C 函数声明 ==========

var (
	mdOnce sync.Once

	_MdCreateFtdcMdApi func(
		*byte, bool, bool
	) uintptr
	_MdGetApiVersion func(
		
	) *char *
	_MdRelease func(
		uintptr
	)
	_MdInit func(
		uintptr
	)
	_MdJoin func(
		uintptr
	) int32
	_MdGetTradingDay func(
		uintptr
	) *char *
	_MdRegisterFront func(
		uintptr, *byte
	)
	_MdRegisterNameServer func(
		uintptr, *byte
	)
	_MdRegisterFensUserInfo func(
		uintptr, *CThostFtdcFensUserInfoField
	)
	_MdSubscribeMarketData func(
		uintptr, *byte, int32
	) int32
	_MdUnSubscribeMarketData func(
		uintptr, *byte, int32
	) int32
	_MdSubscribeForQuoteRsp func(
		uintptr, *byte, int32
	) int32
	_MdUnSubscribeForQuoteRsp func(
		uintptr, *byte, int32
	) int32
	_MdReqUserLogin func(
		uintptr, *CThostFtdcReqUserLoginField, int32
	) int32
	_MdReqUserLogout func(
		uintptr, *CThostFtdcUserLogoutField, int32
	) int32
	_MdReqQryMulticastInstrument func(
		uintptr, *CThostFtdcQryMulticastInstrumentField, int32
	) int32
	_MdSpiCreate func(
		uintptr
	) uintptr
	_MdSpiDestroy func(
		uintptr
	)
	_MdRegisterSpi func(
		uintptr, uintptr
	)
	_MdSpiSetCallbacks func(
		uintptr, *MdSpiCallbacks
	)
	_MdSpiSetOnFrontConnected func(
		uintptr, MdOnFrontConnectedCallback
	)
	_MdSpiSetOnFrontDisconnected func(
		uintptr, MdOnFrontDisconnectedCallback
	)
	_MdSpiSetOnHeartBeatWarning func(
		uintptr, MdOnHeartBeatWarningCallback
	)
	_MdSpiSetOnRspUserLogin func(
		uintptr, MdOnRspUserLoginCallback
	)
	_MdSpiSetOnRspUserLogout func(
		uintptr, MdOnRspUserLogoutCallback
	)
	_MdSpiSetOnRspQryMulticastInstrument func(
		uintptr, MdOnRspQryMulticastInstrumentCallback
	)
	_MdSpiSetOnRspError func(
		uintptr, MdOnRspErrorCallback
	)
	_MdSpiSetOnRspSubMarketData func(
		uintptr, MdOnRspSubMarketDataCallback
	)
	_MdSpiSetOnRspUnSubMarketData func(
		uintptr, MdOnRspUnSubMarketDataCallback
	)
	_MdSpiSetOnRspSubForQuoteRsp func(
		uintptr, MdOnRspSubForQuoteRspCallback
	)
	_MdSpiSetOnRspUnSubForQuoteRsp func(
		uintptr, MdOnRspUnSubForQuoteRspCallback
	)
	_MdSpiSetOnRtnDepthMarketData func(
		uintptr, MdOnRtnDepthMarketDataCallback
	)
	_MdSpiSetOnRtnForQuoteRsp func(
		uintptr, MdOnRtnForQuoteRspCallback
	)
)

// initMdApi 初始化行情 API 函数
func initMdApi(lib uintptr) {
	mdOnce.Do(func() {
		purego.RegisterLibFunc(&_MdCreateFtdcMdApi, lib, "MdCreateFtdcMdApi")
		purego.RegisterLibFunc(&_MdGetApiVersion, lib, "MdGetApiVersion")
		purego.RegisterLibFunc(&_MdRelease, lib, "MdRelease")
		purego.RegisterLibFunc(&_MdInit, lib, "MdInit")
		purego.RegisterLibFunc(&_MdJoin, lib, "MdJoin")
		purego.RegisterLibFunc(&_MdGetTradingDay, lib, "MdGetTradingDay")
		purego.RegisterLibFunc(&_MdRegisterFront, lib, "MdRegisterFront")
		purego.RegisterLibFunc(&_MdRegisterNameServer, lib, "MdRegisterNameServer")
		purego.RegisterLibFunc(&_MdRegisterFensUserInfo, lib, "MdRegisterFensUserInfo")
		purego.RegisterLibFunc(&_MdSubscribeMarketData, lib, "MdSubscribeMarketData")
		purego.RegisterLibFunc(&_MdUnSubscribeMarketData, lib, "MdUnSubscribeMarketData")
		purego.RegisterLibFunc(&_MdSubscribeForQuoteRsp, lib, "MdSubscribeForQuoteRsp")
		purego.RegisterLibFunc(&_MdUnSubscribeForQuoteRsp, lib, "MdUnSubscribeForQuoteRsp")
		purego.RegisterLibFunc(&_MdReqUserLogin, lib, "MdReqUserLogin")
		purego.RegisterLibFunc(&_MdReqUserLogout, lib, "MdReqUserLogout")
		purego.RegisterLibFunc(&_MdReqQryMulticastInstrument, lib, "MdReqQryMulticastInstrument")
		purego.RegisterLibFunc(&_MdSpiCreate, lib, "MdSpiCreate")
		purego.RegisterLibFunc(&_MdSpiDestroy, lib, "MdSpiDestroy")
		purego.RegisterLibFunc(&_MdRegisterSpi, lib, "MdRegisterSpi")
		purego.RegisterLibFunc(&_MdSpiSetCallbacks, lib, "MdSpiSetCallbacks")
		purego.RegisterLibFunc(&_MdSpiSetOnFrontConnected, lib, "MdSpiSetOnFrontConnected")
		purego.RegisterLibFunc(&_MdSpiSetOnFrontDisconnected, lib, "MdSpiSetOnFrontDisconnected")
		purego.RegisterLibFunc(&_MdSpiSetOnHeartBeatWarning, lib, "MdSpiSetOnHeartBeatWarning")
		purego.RegisterLibFunc(&_MdSpiSetOnRspUserLogin, lib, "MdSpiSetOnRspUserLogin")
		purego.RegisterLibFunc(&_MdSpiSetOnRspUserLogout, lib, "MdSpiSetOnRspUserLogout")
		purego.RegisterLibFunc(&_MdSpiSetOnRspQryMulticastInstrument, lib, "MdSpiSetOnRspQryMulticastInstrument")
		purego.RegisterLibFunc(&_MdSpiSetOnRspError, lib, "MdSpiSetOnRspError")
		purego.RegisterLibFunc(&_MdSpiSetOnRspSubMarketData, lib, "MdSpiSetOnRspSubMarketData")
		purego.RegisterLibFunc(&_MdSpiSetOnRspUnSubMarketData, lib, "MdSpiSetOnRspUnSubMarketData")
		purego.RegisterLibFunc(&_MdSpiSetOnRspSubForQuoteRsp, lib, "MdSpiSetOnRspSubForQuoteRsp")
		purego.RegisterLibFunc(&_MdSpiSetOnRspUnSubForQuoteRsp, lib, "MdSpiSetOnRspUnSubForQuoteRsp")
		purego.RegisterLibFunc(&_MdSpiSetOnRtnDepthMarketData, lib, "MdSpiSetOnRtnDepthMarketData")
		purego.RegisterLibFunc(&_MdSpiSetOnRtnForQuoteRsp, lib, "MdSpiSetOnRtnForQuoteRsp")
	})
}

// ========== 实例管理 ==========

var (
	mdInstances   = make(map[uintptr]*MdApi)
	mdInstancesMu sync.RWMutex
	mdNextID      uintptr = 1
)

func registerMdInstance(api *MdApi) uintptr {
	mdInstancesMu.Lock()
	defer mdInstancesMu.Unlock()
	id := mdNextID
	mdNextID++
	mdInstances[id] = api
	return id
}

func getMdInstance(userData uintptr) *MdApi {
	mdInstancesMu.RLock()
	defer mdInstancesMu.RUnlock()
	return mdInstances[userData]
}

func unregisterMdInstance(userData uintptr) {
	mdInstancesMu.Lock()
	defer mdInstancesMu.Unlock()
	delete(mdInstances, userData)
}

// ========== 构造函数 ==========

// NewMdApi 创建行情 API 实例
func NewMdApi(flowPath string, usingUdp, multicast bool) *MdApi {
	api := &MdApi{}
	api.userData = registerMdInstance(api)
	
	pathPtr := CString(flowPath)
	api.handle = _MdCreateFtdcMdApi(pathPtr, usingUdp, multicast)
	
	runtime.SetFinalizer(api, (*MdApi).Release)
	return api
}

// ========== API 方法 ==========

// GetApiVersion 获取API的版本信息
func (api *MdApi) GetApiVersion() *char * {
	return _MdGetApiVersion(api.handle)
}

// Release 删除接口对象本身
func (api *MdApi) Release() {
	_MdRelease(api.handle)
}

// Init 初始化
func (api *MdApi) Init() {
	_MdInit(api.handle)
}

// Join 等待接口线程结束运行
func (api *MdApi) Join() int32 {
	return _MdJoin(api.handle)
}

// GetTradingDay 获取当前交易日
func (api *MdApi) GetTradingDay() *char * {
	return _MdGetTradingDay(api.handle)
}

// RegisterFront 注册前置机网络地址
func (api *MdApi) RegisterFront(pszFrontAddress string) {
	_MdRegisterFront(api.handle, CString(pszFrontAddress))
}

// RegisterNameServer 注册名字服务器网络地址
func (api *MdApi) RegisterNameServer(pszNsAddress string) {
	_MdRegisterNameServer(api.handle, CString(pszNsAddress))
}

// RegisterFensUserInfo 注册名字服务器用户信息
func (api *MdApi) RegisterFensUserInfo(pFensUserInfo *CThostFtdcFensUserInfoField) {
	_MdRegisterFensUserInfo(api.handle, pFensUserInfo)
}

// SubscribeMarketData 订阅行情。
func (api *MdApi) SubscribeMarketData(ppInstrumentID[] string, nCount int32) int32 {
	return _MdSubscribeMarketData(api.handle, CString(ppInstrumentID[]), nCount)
}

// UnSubscribeMarketData 退订行情。
func (api *MdApi) UnSubscribeMarketData(ppInstrumentID[] string, nCount int32) int32 {
	return _MdUnSubscribeMarketData(api.handle, CString(ppInstrumentID[]), nCount)
}

// SubscribeForQuoteRsp 订阅询价。
func (api *MdApi) SubscribeForQuoteRsp(ppInstrumentID[] string, nCount int32) int32 {
	return _MdSubscribeForQuoteRsp(api.handle, CString(ppInstrumentID[]), nCount)
}

// UnSubscribeForQuoteRsp 退订询价。
func (api *MdApi) UnSubscribeForQuoteRsp(ppInstrumentID[] string, nCount int32) int32 {
	return _MdUnSubscribeForQuoteRsp(api.handle, CString(ppInstrumentID[]), nCount)
}

// ReqUserLogin 用户登录请求
func (api *MdApi) ReqUserLogin(pReqUserLoginField *CThostFtdcReqUserLoginField, nRequestID int32) int32 {
	return _MdReqUserLogin(api.handle, pReqUserLoginField, nRequestID)
}

// ReqUserLogout 登出请求
func (api *MdApi) ReqUserLogout(pUserLogout *CThostFtdcUserLogoutField, nRequestID int32) int32 {
	return _MdReqUserLogout(api.handle, pUserLogout, nRequestID)
}

// ReqQryMulticastInstrument 请求查询组播合约
func (api *MdApi) ReqQryMulticastInstrument(pQryMulticastInstrument *CThostFtdcQryMulticastInstrumentField, nRequestID int32) int32 {
	return _MdReqQryMulticastInstrument(api.handle, pQryMulticastInstrument, nRequestID)
}

// SpiCreate ========== Md SPI 函数 ========== 创建 SPI 实例
func (api *MdApi) SpiCreate() uintptr {
	return _MdSpiCreate(api.handle)
}

// SpiDestroy 销毁 SPI 实例
func (api *MdApi) SpiDestroy() {
	_MdSpiDestroy(api.handle)
}

// RegisterSpi 注册 SPI 到 API
func (api *MdApi) RegisterSpi(spi uintptr) {
	_MdRegisterSpi(api.handle, spi)
}

// SpiSetCallbacks 批量设置回调
func (api *MdApi) SpiSetCallbacks(callbacks *MdSpiCallbacks) {
	_MdSpiSetCallbacks(api.handle, callbacks)
}

// SpiSetOnFrontConnected 单独设置回调 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (api *MdApi) SpiSetOnFrontConnected(callback MdOnFrontConnectedCallback) {
	_MdSpiSetOnFrontConnected(api.handle, callback)
}

// SpiSetOnFrontDisconnected 0x2003 收到错误报文
func (api *MdApi) SpiSetOnFrontDisconnected(callback MdOnFrontDisconnectedCallback) {
	_MdSpiSetOnFrontDisconnected(api.handle, callback)
}

// SpiSetOnHeartBeatWarning 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (api *MdApi) SpiSetOnHeartBeatWarning(callback MdOnHeartBeatWarningCallback) {
	_MdSpiSetOnHeartBeatWarning(api.handle, callback)
}

// SpiSetOnRspUserLogin 登录请求响应
func (api *MdApi) SpiSetOnRspUserLogin(callback MdOnRspUserLoginCallback) {
	_MdSpiSetOnRspUserLogin(api.handle, callback)
}

// SpiSetOnRspUserLogout 登出请求响应
func (api *MdApi) SpiSetOnRspUserLogout(callback MdOnRspUserLogoutCallback) {
	_MdSpiSetOnRspUserLogout(api.handle, callback)
}

// SpiSetOnRspQryMulticastInstrument 请求查询组播合约响应
func (api *MdApi) SpiSetOnRspQryMulticastInstrument(callback MdOnRspQryMulticastInstrumentCallback) {
	_MdSpiSetOnRspQryMulticastInstrument(api.handle, callback)
}

// SpiSetOnRspError 错误应答
func (api *MdApi) SpiSetOnRspError(callback MdOnRspErrorCallback) {
	_MdSpiSetOnRspError(api.handle, callback)
}

// SpiSetOnRspSubMarketData 订阅行情应答
func (api *MdApi) SpiSetOnRspSubMarketData(callback MdOnRspSubMarketDataCallback) {
	_MdSpiSetOnRspSubMarketData(api.handle, callback)
}

// SpiSetOnRspUnSubMarketData 取消订阅行情应答
func (api *MdApi) SpiSetOnRspUnSubMarketData(callback MdOnRspUnSubMarketDataCallback) {
	_MdSpiSetOnRspUnSubMarketData(api.handle, callback)
}

// SpiSetOnRspSubForQuoteRsp 订阅询价应答
func (api *MdApi) SpiSetOnRspSubForQuoteRsp(callback MdOnRspSubForQuoteRspCallback) {
	_MdSpiSetOnRspSubForQuoteRsp(api.handle, callback)
}

// SpiSetOnRspUnSubForQuoteRsp 取消订阅询价应答
func (api *MdApi) SpiSetOnRspUnSubForQuoteRsp(callback MdOnRspUnSubForQuoteRspCallback) {
	_MdSpiSetOnRspUnSubForQuoteRsp(api.handle, callback)
}

// SpiSetOnRtnDepthMarketData 深度行情通知
func (api *MdApi) SpiSetOnRtnDepthMarketData(callback MdOnRtnDepthMarketDataCallback) {
	_MdSpiSetOnRtnDepthMarketData(api.handle, callback)
}

// SpiSetOnRtnForQuoteRsp 询价通知
func (api *MdApi) SpiSetOnRtnForQuoteRsp(callback MdOnRtnForQuoteRspCallback) {
	_MdSpiSetOnRtnForQuoteRsp(api.handle, callback)
}

// Release 释放 API 实例
func (api *MdApi) Release() {
	if api.handle != 0 {
		_MdRelease(api.handle)
		unregisterMdInstance(api.userData)
		api.handle = 0
	}
}

// SetSpi 设置回调接口
func (api *MdApi) SetSpi(spi MdSpi) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.spi = spi
}
