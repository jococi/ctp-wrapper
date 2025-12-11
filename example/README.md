# CTP Go 包装使用示例

本目录包含使用 CTP Go 包装库的示例代码。

## 文件说明

- `main.go`: 完整的行情和交易 API 使用示例
- `.env.example`: 环境变量配置模板文件（复制为 `.env` 后填入真实信息）
- `README.md`: 本说明文件

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

**方式1：使用环境变量文件（推荐）**

复制 `.env.example` 为 `.env` 并填入真实的账户信息：

```bash
cp .env.example .env
# 编辑 .env 文件，填入你的账户信息
```

`.env` 文件内容示例：
```
CTP_USER_ID=your_user_id
CTP_PASSWORD=your_password
CTP_BROKER_ID=your_broker_id
CTP_MD_FRONT=tcp://182.254.243.31:40011
CTP_TRADER_FRONT=tcp://182.254.243.31:30001
```

**方式2：使用系统环境变量**

直接在系统中设置环境变量：
```bash
export CTP_USER_ID=your_user_id
export CTP_PASSWORD=your_password
export CTP_BROKER_ID=your_broker_id
```

### 3. 运行示例

```bash
cd example
go mod tidy
go run main.go
```

**注意：** 确保已正确配置环境变量，否则程序会提示缺少必要的配置信息。

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
2. **敏感信息保护**：**不要**在代码中硬编码账户信息，使用环境变量或 `.env` 文件
3. **`.env` 文件**：`.env` 文件已添加到 `.gitignore`，不会被提交到版本库，可以安全地存储敏感信息
4. **连接信息**：使用真实的服务器地址和登录凭证
5. **错误处理**：实际使用时应该添加更完善的错误处理逻辑
6. **线程安全**：回调函数可能在不同的 goroutine 中调用，需要注意线程安全
7. **穿透式认证**：生产环境可能需要穿透式认证，需要使用 `ReqAuthenticate` 接口
