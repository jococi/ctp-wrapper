# CTP Go 包装使用示例

本目录包含使用 CTP Go 包装库的示例代码。

## 文件说明

- `main.go`: 完整的行情和交易 API 使用示例

## 使用说明

### 1. 准备 CTP 库文件

将 CTP 的 C 包装库文件放在以下位置之一：

- `./libs/` 目录
- 环境变量 `CTP_LIB_PATH` 指定的路径
- 可执行文件所在目录的 `libs/` 子目录

需要的库文件：
- Linux: `libctp_md_c_api.so`, `libctp_trader_c_api.so`
- macOS: `libctp_md_c_api.dylib`, `libctp_trader_c_api.dylib`
- Windows: `ctp_md_c_api.dll`, `ctp_trader_c_api.dll`

### 2. 配置连接信息

在 `main.go` 中修改以下配置：

```go
// 前置服务器地址
mdApi.RegisterFront("tcp://your_md_server:port")
traderApi.RegisterFront("tcp://your_trader_server:port")

// 登录信息
copy(loginReq.UserID[:], "your_user_id")
copy(loginReq.Password[:], "your_password")
copy(loginReq.BrokerID[:], "your_broker_id")
```

### 3. 运行示例

```bash
cd example
go mod init example
go mod edit -replace github.com/wavy/agentic-trade/ctp-wrapper/ctpgo=../ctpgo
go mod tidy
go run main.go
```

## 自定义回调方法

### 使用默认回调

通过嵌入 `DefaultMdSpi` 或 `DefaultTraderSpi`，你只需要实现需要的方法：

```go
type MyMdSpi struct {
    ctpgo.DefaultMdSpi  // 嵌入默认实现
}

// 只实现需要的方法
func (s *MyMdSpi) OnRtnDepthMarketData(pDepthMarketData *ctpgo.CThostFtdcDepthMarketDataField) {
    // 你的实现
}
```

### 重要说明

- **默认回调方法不会影响自定义回调方法**：Go 的方法集（method set）机制确保如果你实现了某个方法，会优先使用你的实现
- 未实现的方法会使用默认的空实现（什么都不做）
- 这种方式可以让你只关注需要处理的回调，而不需要实现所有接口方法

### 完整实现接口

如果你不想使用默认实现，也可以直接实现所有接口方法：

```go
type MyMdSpi struct {
    // 不嵌入 DefaultMdSpi
}

func (s *MyMdSpi) OnFrontConnected() { ... }
func (s *MyMdSpi) OnFrontDisconnected(nReason int32) { ... }
// ... 实现所有接口方法
```

## 注意事项

1. **库文件路径**：确保 CTP 库文件在正确的位置，或设置 `CTP_LIB_PATH` 环境变量
2. **连接信息**：使用真实的服务器地址和登录凭证
3. **错误处理**：实际使用时应该添加更完善的错误处理逻辑
4. **线程安全**：回调函数可能在不同的 goroutine 中调用，需要注意线程安全
