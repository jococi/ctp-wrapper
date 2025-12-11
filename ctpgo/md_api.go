package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 行情 API 封装

import (
	"path/filepath"
	"runtime"
	"sync"
	"unsafe"

	"github.com/ebitengine/purego"
)

// ========== 回调类型定义 ==========

// MdOnFrontConnectedCallback ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
type MdOnFrontConnectedCallback func(userData uintptr)

// MdOnFrontDisconnectedCallback 0x2003 收到错误报文
type MdOnFrontDisconnectedCallback func(userData uintptr, nReason int32)

// MdOnHeartBeatWarningCallback 心跳超时警告。当长时间未收到报文时，该方法被调用。
type MdOnHeartBeatWarningCallback func(userData uintptr, nTimeLapse int32)

// MdOnRspUserLoginCallback 登录请求响应
type MdOnRspUserLoginCallback func(userData uintptr, pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspUserLogoutCallback 登出请求响应
type MdOnRspUserLogoutCallback func(userData uintptr, pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspQryMulticastInstrumentCallback 请求查询组播合约响应
type MdOnRspQryMulticastInstrumentCallback func(userData uintptr, pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspErrorCallback 错误应答
type MdOnRspErrorCallback func(userData uintptr, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspSubMarketDataCallback 订阅行情应答
type MdOnRspSubMarketDataCallback func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspUnSubMarketDataCallback 取消订阅行情应答
type MdOnRspUnSubMarketDataCallback func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspSubForQuoteRspCallback 订阅询价应答
type MdOnRspSubForQuoteRspCallback func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRspUnSubForQuoteRspCallback 取消订阅询价应答
type MdOnRspUnSubForQuoteRspCallback func(userData uintptr, pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)

// MdOnRtnDepthMarketDataCallback 深度行情通知
type MdOnRtnDepthMarketDataCallback func(userData uintptr, pDepthMarketData *CThostFtdcDepthMarketDataField)

// MdOnRtnForQuoteRspCallback 询价通知
type MdOnRtnForQuoteRspCallback func(userData uintptr, pForQuoteRsp *CThostFtdcForQuoteRspField)

// MdSpiCallbacks 回调结构体（用于批量设置）
type MdSpiCallbacks struct {
	UserData                    uintptr
	OnFrontConnected            MdOnFrontConnectedCallback
	OnFrontDisconnected         MdOnFrontDisconnectedCallback
	OnHeartBeatWarning          MdOnHeartBeatWarningCallback
	OnRspUserLogin              MdOnRspUserLoginCallback
	OnRspUserLogout             MdOnRspUserLogoutCallback
	OnRspQryMulticastInstrument MdOnRspQryMulticastInstrumentCallback
	OnRspError                  MdOnRspErrorCallback
	OnRspSubMarketData          MdOnRspSubMarketDataCallback
	OnRspUnSubMarketData        MdOnRspUnSubMarketDataCallback
	OnRspSubForQuoteRsp         MdOnRspSubForQuoteRspCallback
	OnRspUnSubForQuoteRsp       MdOnRspUnSubForQuoteRspCallback
	OnRtnDepthMarketData        MdOnRtnDepthMarketDataCallback
	OnRtnForQuoteRsp            MdOnRtnForQuoteRspCallback
}

// ========== MdSpi 接口 ==========

// MdSpi 行情回调接口
type MdSpi interface {
	OnFrontConnected()                                                                                                                                      // ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
	OnFrontDisconnected(nReason int32)                                                                                                                      // 0x2003 收到错误报文
	OnHeartBeatWarning(nTimeLapse int32)                                                                                                                    // 心跳超时警告。当长时间未收到报文时，该方法被调用。
	OnRspUserLogin(pRspUserLogin *CThostFtdcRspUserLoginField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                            // 登录请求响应
	OnRspUserLogout(pUserLogout *CThostFtdcUserLogoutField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                               // 登出请求响应
	OnRspQryMulticastInstrument(pMulticastInstrument *CThostFtdcMulticastInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) // 请求查询组播合约响应
	OnRspError(pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)                                                                            // 错误应答
	OnRspSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)            // 订阅行情应答
	OnRspUnSubMarketData(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)          // 取消订阅行情应答
	OnRspSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)           // 订阅询价应答
	OnRspUnSubForQuoteRsp(pSpecificInstrument *CThostFtdcSpecificInstrumentField, pRspInfo *CThostFtdcRspInfoField, nRequestID int32, bIsLast bool)         // 取消订阅询价应答
	OnRtnDepthMarketData(pDepthMarketData *CThostFtdcDepthMarketDataField)                                                                                  // 深度行情通知
	OnRtnForQuoteRsp(pForQuoteRsp *CThostFtdcForQuoteRspField)                                                                                              // 询价通知
}

// ========== MdApi 结构体 ==========

// MdApi 行情 API 封装
type MdApi struct {
	handle    uintptr
	spi       MdSpi
	spiHandle uintptr // C SPI 实例句柄
	userData  uintptr
	mu        sync.RWMutex
	flowPath  []byte // 保存 flowPath 的 C 字符串，防止被 GC 回收
}

// ========== C 函数声明 ==========

var (
	mdOnce sync.Once

	_MdCreateFtdcMdApi                   func(*byte, bool, bool) uintptr
	_MdGetApiVersion                     func() *byte
	_MdRelease                           func(uintptr)
	_MdInit                              func(uintptr)
	_MdJoin                              func(uintptr) int32
	_MdGetTradingDay                     func(uintptr) *byte
	_MdRegisterFront                     func(uintptr, *byte)
	_MdRegisterNameServer                func(uintptr, *byte)
	_MdRegisterFensUserInfo              func(uintptr, *CThostFtdcFensUserInfoField)
	_MdSubscribeMarketData               func(uintptr, **byte, int32) int32
	_MdUnSubscribeMarketData             func(uintptr, **byte, int32) int32
	_MdSubscribeForQuoteRsp              func(uintptr, **byte, int32) int32
	_MdUnSubscribeForQuoteRsp            func(uintptr, **byte, int32) int32
	_MdReqUserLogin                      func(uintptr, *CThostFtdcReqUserLoginField, int32) int32
	_MdReqUserLogout                     func(uintptr, *CThostFtdcUserLogoutField, int32) int32
	_MdReqQryMulticastInstrument         func(uintptr, *CThostFtdcQryMulticastInstrumentField, int32) int32
	_MdSpiCreate                         func(uintptr) uintptr
	_MdSpiDestroy                        func(uintptr)
	_MdRegisterSpi                       func(uintptr, uintptr)
	_MdSpiSetCallbacks                   func(uintptr, *MdSpiCallbacks)
	_MdSpiSetOnFrontConnected            func(uintptr, uintptr)
	_MdSpiSetOnFrontDisconnected         func(uintptr, uintptr)
	_MdSpiSetOnHeartBeatWarning          func(uintptr, uintptr)
	_MdSpiSetOnRspUserLogin              func(uintptr, uintptr)
	_MdSpiSetOnRspUserLogout             func(uintptr, uintptr)
	_MdSpiSetOnRspQryMulticastInstrument func(uintptr, uintptr)
	_MdSpiSetOnRspError                  func(uintptr, uintptr)
	_MdSpiSetOnRspSubMarketData          func(uintptr, uintptr)
	_MdSpiSetOnRspUnSubMarketData        func(uintptr, uintptr)
	_MdSpiSetOnRspSubForQuoteRsp         func(uintptr, uintptr)
	_MdSpiSetOnRspUnSubForQuoteRsp       func(uintptr, uintptr)
	_MdSpiSetOnRtnDepthMarketData        func(uintptr, uintptr)
	_MdSpiSetOnRtnForQuoteRsp            func(uintptr, uintptr)
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
// 首次调用时会自动加载 CTP 库（如果尚未加载）
func NewMdApi(flowPath string, usingUdp, multicast bool) *MdApi {
	// 自动加载库（如果尚未加载）
	if err := autoLoadLibrary(); err != nil {
		// 如果自动加载失败，返回 nil（或者可以 panic，取决于设计）
		// 这里返回 nil，让调用者检查
		return nil
	}

	api := &MdApi{}
	api.userData = registerMdInstance(api)

	// 将相对路径转换为绝对路径，确保 CTP API 能正确识别目录
	// CTP API 基于当前工作目录解析相对路径，但可能工作目录不是我们期望的
	// 所以转换为绝对路径更可靠
	absFlowPath := flowPath
	if !filepath.IsAbs(flowPath) {
		// 如果是相对路径，转换为基于当前工作目录的绝对路径
		var err error
		absFlowPath, err = filepath.Abs(flowPath)
		if err != nil {
			// 如果转换失败，使用原始路径
			absFlowPath = flowPath
		}
	}

	// 确保路径以路径分隔符结尾（CTP API 可能需要这样才能识别为目录）
	if len(absFlowPath) > 0 && absFlowPath[len(absFlowPath)-1] != filepath.Separator {
		absFlowPath += string(filepath.Separator)
	}

	// 将 flowPath 转换为 C 字符串并保存，防止被 GC 回收
	// CTP API 可能会在后续使用这个路径
	api.flowPath = make([]byte, len(absFlowPath)+1)
	copy(api.flowPath, absFlowPath)
	api.flowPath[len(absFlowPath)] = 0 // null terminator
	pathPtr := &api.flowPath[0]

	api.handle = _MdCreateFtdcMdApi(pathPtr, usingUdp, multicast)

	runtime.SetFinalizer(api, (*MdApi).Release)
	return api
}

// ========== API 方法 ==========

// GetApiVersion 获取API的版本信息
func (api *MdApi) GetApiVersion() string {
	ptr := _MdGetApiVersion()
	if ptr == nil {
		return ""
	}
	return GoString(ptr)
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
func (api *MdApi) GetTradingDay() string {
	ptr := _MdGetTradingDay(api.handle)
	if ptr == nil {
		return ""
	}
	return GoString(ptr)
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
func (api *MdApi) SubscribeMarketData(ppInstrumentID []string, nCount int32) int32 {
	if len(ppInstrumentID) == 0 {
		return 0
	}
	// 将字符串数组转换为 C 字符串数组
	ptrs, _ := CStringArray(ppInstrumentID)
	return _MdSubscribeMarketData(api.handle, ptrs, nCount)
}

// UnSubscribeMarketData 退订行情。
func (api *MdApi) UnSubscribeMarketData(ppInstrumentID []string, nCount int32) int32 {
	if len(ppInstrumentID) == 0 {
		return 0
	}
	// 将字符串数组转换为 C 字符串数组
	ptrs, _ := CStringArray(ppInstrumentID)
	return _MdUnSubscribeMarketData(api.handle, ptrs, nCount)
}

// SubscribeForQuoteRsp 订阅询价。
func (api *MdApi) SubscribeForQuoteRsp(ppInstrumentID []string, nCount int32) int32 {
	if len(ppInstrumentID) == 0 {
		return 0
	}
	// 将字符串数组转换为 C 字符串数组
	ptrs, _ := CStringArray(ppInstrumentID)
	return _MdSubscribeForQuoteRsp(api.handle, ptrs, nCount)
}

// UnSubscribeForQuoteRsp 退订询价。
func (api *MdApi) UnSubscribeForQuoteRsp(ppInstrumentID []string, nCount int32) int32 {
	if len(ppInstrumentID) == 0 {
		return 0
	}
	// 将字符串数组转换为 C 字符串数组
	ptrs, _ := CStringArray(ppInstrumentID)
	return _MdUnSubscribeForQuoteRsp(api.handle, ptrs, nCount)
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

// ========== SPI 回调设置方法 ==========

// SpiSetOnFrontConnected ========== 回调函数类型（带 userData） ========== 当客户端与交易后台建立起通信连接时（还未登录前），该方法被调用。
func (api *MdApi) SpiSetOnFrontConnected(callback MdOnFrontConnectedCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnFrontConnected(api.spiHandle, ptr)
}

// SpiSetOnFrontDisconnected 0x2003 收到错误报文
func (api *MdApi) SpiSetOnFrontDisconnected(callback MdOnFrontDisconnectedCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnFrontDisconnected(api.spiHandle, ptr)
}

// SpiSetOnHeartBeatWarning 心跳超时警告。当长时间未收到报文时，该方法被调用。
func (api *MdApi) SpiSetOnHeartBeatWarning(callback MdOnHeartBeatWarningCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnHeartBeatWarning(api.spiHandle, ptr)
}

// SpiSetOnRspUserLogin 登录请求响应
func (api *MdApi) SpiSetOnRspUserLogin(callback MdOnRspUserLoginCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspUserLogin(api.spiHandle, ptr)
}

// SpiSetOnRspUserLogout 登出请求响应
func (api *MdApi) SpiSetOnRspUserLogout(callback MdOnRspUserLogoutCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspUserLogout(api.spiHandle, ptr)
}

// SpiSetOnRspQryMulticastInstrument 请求查询组播合约响应
func (api *MdApi) SpiSetOnRspQryMulticastInstrument(callback MdOnRspQryMulticastInstrumentCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspQryMulticastInstrument(api.spiHandle, ptr)
}

// SpiSetOnRspError 错误应答
func (api *MdApi) SpiSetOnRspError(callback MdOnRspErrorCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspError(api.spiHandle, ptr)
}

// SpiSetOnRspSubMarketData 订阅行情应答
func (api *MdApi) SpiSetOnRspSubMarketData(callback MdOnRspSubMarketDataCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspSubMarketData(api.spiHandle, ptr)
}

// SpiSetOnRspUnSubMarketData 取消订阅行情应答
func (api *MdApi) SpiSetOnRspUnSubMarketData(callback MdOnRspUnSubMarketDataCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspUnSubMarketData(api.spiHandle, ptr)
}

// SpiSetOnRspSubForQuoteRsp 订阅询价应答
func (api *MdApi) SpiSetOnRspSubForQuoteRsp(callback MdOnRspSubForQuoteRspCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspSubForQuoteRsp(api.spiHandle, ptr)
}

// SpiSetOnRspUnSubForQuoteRsp 取消订阅询价应答
func (api *MdApi) SpiSetOnRspUnSubForQuoteRsp(callback MdOnRspUnSubForQuoteRspCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRspUnSubForQuoteRsp(api.spiHandle, ptr)
}

// SpiSetOnRtnDepthMarketData 深度行情通知
func (api *MdApi) SpiSetOnRtnDepthMarketData(callback MdOnRtnDepthMarketDataCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRtnDepthMarketData(api.spiHandle, ptr)
}

// SpiSetOnRtnForQuoteRsp 询价通知
func (api *MdApi) SpiSetOnRtnForQuoteRsp(callback MdOnRtnForQuoteRspCallback) {
	// 将函数类型转换为 uintptr
	ptr := *(*uintptr)(unsafe.Pointer(&callback))
	_MdSpiSetOnRtnForQuoteRsp(api.spiHandle, ptr)
}

// SetSpi 设置回调接口
// 此方法会创建 C SPI 实例，注册 Go 回调函数，并将 SPI 注册到 API
func (api *MdApi) SetSpi(spi MdSpi) {
	api.mu.Lock()
	defer api.mu.Unlock()
	api.spi = spi

	// 如果已有 C SPI 实例，先销毁
	if api.spiHandle != 0 {
		_MdSpiDestroy(api.spiHandle)
	}

	// 创建新的 C SPI 实例
	api.spiHandle = _MdSpiCreate(api.userData)

	// 注册所有回调函数到 C SPI
	// 使用回调文件中提供的辅助函数来获取函数指针（这些函数会包装 //export 函数以匹配正确的签名）
	_MdSpiSetOnFrontConnected(api.spiHandle, GetGoMdOnFrontConnected())
	_MdSpiSetOnFrontDisconnected(api.spiHandle, GetGoMdOnFrontDisconnected())
	_MdSpiSetOnHeartBeatWarning(api.spiHandle, GetGoMdOnHeartBeatWarning())
	_MdSpiSetOnRspUserLogin(api.spiHandle, GetGoMdOnRspUserLogin())
	_MdSpiSetOnRspUserLogout(api.spiHandle, GetGoMdOnRspUserLogout())
	_MdSpiSetOnRspQryMulticastInstrument(api.spiHandle, GetGoMdOnRspQryMulticastInstrument())
	_MdSpiSetOnRspError(api.spiHandle, GetGoMdOnRspError())
	_MdSpiSetOnRspSubMarketData(api.spiHandle, GetGoMdOnRspSubMarketData())
	_MdSpiSetOnRspUnSubMarketData(api.spiHandle, GetGoMdOnRspUnSubMarketData())
	_MdSpiSetOnRspSubForQuoteRsp(api.spiHandle, GetGoMdOnRspSubForQuoteRsp())
	_MdSpiSetOnRspUnSubForQuoteRsp(api.spiHandle, GetGoMdOnRspUnSubForQuoteRsp())
	_MdSpiSetOnRtnDepthMarketData(api.spiHandle, GetGoMdOnRtnDepthMarketData())
	_MdSpiSetOnRtnForQuoteRsp(api.spiHandle, GetGoMdOnRtnForQuoteRsp())

	// 将 C SPI 注册到 API
	_MdRegisterSpi(api.handle, api.spiHandle)
}
