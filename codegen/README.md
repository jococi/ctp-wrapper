# CTP C API 代码生成器

将 CTP 官方 C++ API 转换为纯 C 接口，支持多实例，跨平台统一。

## 特性

- **纯 C 接口**：生成的头文件不依赖任何 C++ 语法
- **不透明指针**：使用句柄类型隐藏实现细节
- **userData 支持**：所有回调携带 `userData` 参数，支持多实例
- **驼峰命名**：函数和回调使用驼峰命名风格
- **跨平台统一**：支持 Windows/Linux/macOS，统一处理平台差异

## 三平台 API 差异分析

通过对比三个平台的 CTP 头文件，发现：

| 文件 | Linux/Windows | macOS | 差异 |
|------|---------------|-------|------|
| ThostFtdcMdApi.h | ✅ 完全一致 | ✅ 完全一致 | 仅编码不同 |
| ThostFtdcTraderApi.h | ✅ 一致 | ⚠️ 1处差异 | ReqUserLogin 签名不同 |
| ThostFtdcUserApiStruct.h | ✅ 完全一致 | ✅ 完全一致 | 仅编码不同 |
| ThostFtdcUserApiDataType.h | ✅ 完全一致 | ✅ 完全一致 | 仅编码不同 |

**唯一的 API 签名差异**：

```cpp
// Linux/Windows
virtual int ReqUserLogin(CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID) = 0;

// macOS (多两个参数用于系统信息采集)
virtual int ReqUserLogin(CThostFtdcReqUserLoginField *pReqUserLoginField, int nRequestID, 
                         TThostFtdcSystemInfoLenType length, TThostFtdcClientSystemInfoType systemInfo) = 0;
```

### 跨平台统一方案

本生成器通过条件编译自动处理差异：

```c
// 统一的登录接口 - 各平台行为一致
int ctpTraderReqUserLogin(CtpTraderApiHandle handle, 
    CThostFtdcReqUserLoginField* pField, int nRequestID);

// 带系统信息的登录接口 - 高级用户使用
int ctpTraderReqUserLoginWithSystemInfo(CtpTraderApiHandle handle,
    CThostFtdcReqUserLoginField* pField, int nRequestID,
    int systemInfoLen, const char* systemInfo);
```

- **Linux/Windows**：`ReqUserLogin` 直接调用底层 API，`systemInfo` 参数被忽略
- **macOS**：自动采集系统信息（或使用传入的 `systemInfo`），然后调用底层 API

## 使用方法

```bash
# 基本用法
python3 generate.py --input ../ctpapi/linux --output ./output

# 指定不同平台的头文件
python3 generate.py --input ../ctpapi/macos --output ./output_macos
```

## 生成的文件

```
output/
├── ctp_md_c_api.h      # MdApi 的纯 C 头文件
├── ctp_md_c_api.cpp    # MdApi 的 C++ 实现
├── ctp_trader_c_api.h  # TraderApi 的纯 C 头文件
└── ctp_trader_c_api.cpp # TraderApi 的 C++ 实现
```

## API 示例

### 创建和初始化

```c
#include "ctp_md_c_api.h"

// 创建 API 实例
CtpMdApiHandle api = ctpMdCreateFtdcMdApi("./log/", false, false);

// 创建 SPI 实例，传入用户数据指针
MyStrategy* strategy = create_strategy();
CtpMdSpiHandle spi = ctpMdSpiCreate(strategy);

// 设置回调
ctpMdSpiSetOnFrontConnected(spi, on_front_connected);
ctpMdSpiSetOnRtnDepthMarketData(spi, on_market_data);

// 注册 SPI 到 API
ctpMdRegisterSpi(api, spi);

// 注册前置地址
ctpMdRegisterFront(api, "tcp://180.168.146.187:10131");

// 初始化
ctpMdInit(api);
```

### 回调函数

```c
// 回调函数接收 userData 参数
void on_front_connected(void* userData) {
    MyStrategy* strategy = (MyStrategy*)userData;
    printf("Strategy %s connected!\n", strategy->name);
}

void on_market_data(void* userData, struct CThostFtdcDepthMarketDataField* pData) {
    MyStrategy* strategy = (MyStrategy*)userData;
    // 使用 userData 访问你的策略实例
    strategy->on_tick(pData);
}
```

### 多实例支持

```c
// 实例 A
MyStrategy* strategyA = create_strategy("策略A");
CtpMdSpiHandle spiA = ctpMdSpiCreate(strategyA);
CtpMdApiHandle apiA = ctpMdCreateFtdcMdApi("./logA/", false, false);
ctpMdRegisterSpi(apiA, spiA);

// 实例 B - 回调不会串台！
MyStrategy* strategyB = create_strategy("策略B");
CtpMdSpiHandle spiB = ctpMdSpiCreate(strategyB);
CtpMdApiHandle apiB = ctpMdCreateFtdcMdApi("./logB/", false, false);
ctpMdRegisterSpi(apiB, spiB);

// 每个实例的回调会正确路由到对应的 userData
```

## 编译

**注意**：生成的 C++ 实现文件包含条件编译，使用 `__APPLE__`、`_WIN32`、`__linux__` 宏自动适配平台。

### Linux

```bash
# MdApi
g++ -shared -fPIC -o libctpmd_c.so ctp_md_c_api.cpp \
    -I../ctpapi/linux \
    -L../ctpapi/linux -lthostmduserapi_se

# TraderApi（包含 DataCollect）
g++ -shared -fPIC -o libctptrader_c.so ctp_trader_c_api.cpp \
    -I../ctpapi/linux \
    -L../ctpapi/linux -lthosttraderapi_se -lLinuxDataCollect
```

### macOS

```bash
# MdApi
clang++ -shared -fPIC -o libctpmd_c.dylib ctp_md_c_api.cpp \
    -I../ctpapi/macos/thostmduserapi_se.framework/Headers \
    -F../ctpapi/macos -framework thostmduserapi_se

# TraderApi（包含 DataCollect）
clang++ -shared -fPIC -o libctptrader_c.dylib ctp_trader_c_api.cpp \
    -I../ctpapi/macos/thosttraderapi_se.framework/Headers \
    -F../ctpapi/macos -framework thosttraderapi_se -framework MacDataCollect
```

### Windows

使用 Visual Studio 或 MSVC 编译：

```bash
# MdApi
cl /LD ctp_md_c_api.cpp /I..\ctpapi\windows /link thostmduserapi_se.lib

# TraderApi（包含 DataCollect）
cl /LD ctp_trader_c_api.cpp /I..\ctpapi\windows /link thosttraderapi_se.lib WinDataCollect.lib
```

## 与 Go 集成

生成的 C API 可以直接用 CGO 调用：

```go
package ctpmd

/*
#cgo CFLAGS: -I${SRCDIR}/output
#cgo LDFLAGS: -L${SRCDIR}/libs -lctpmd_c
#include "ctp_md_c_api.h"
*/
import "C"
import "unsafe"

type MdApi struct {
    handle C.CtpMdApiHandle
    spi    C.CtpMdSpiHandle
}

func NewMdApi(flowPath string, userData unsafe.Pointer) *MdApi {
    api := &MdApi{}
    cPath := C.CString(flowPath)
    defer C.free(unsafe.Pointer(cPath))
    
    api.handle = C.ctpMdCreateFtdcMdApi(cPath, false, false)
    api.spi = C.ctpMdSpiCreate(userData)
    C.ctpMdRegisterSpi(api.handle, api.spi)
    
    return api
}
```

## 注意事项

1. **结构体兼容**：CTP 的 Field 结构体是 C 兼容的 POD 类型，可直接使用
2. **内存管理**：API 和 SPI 实例需要手动释放
3. **线程安全**：CTP 回调在内部线程触发，注意线程同步

## 与旧版封装对比

| 特性 | 旧版（全局变量） | 新版（userData） |
|-----|----------------|-----------------|
| 多实例 | ❌ 不支持 | ✅ 支持 |
| 回调串台 | ⚠️ 可能发生 | ✅ 不会 |
| FFI 友好 | ⚠️ 需要 hack | ✅ 标准做法 |
| 性能 | 相同 | 相同 |
