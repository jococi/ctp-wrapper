# CTP Python 示例程序

这是使用 `pyctp` Python 包装库的示例程序，演示如何使用 CTP API 进行行情订阅和交易操作。

## 环境要求

- Python 3.7+
- CTP 动态库文件（`libctpmd_c_api.so` / `libctptrader_c_api.so` 或对应的 Windows/macOS 版本）

## 安装依赖

```bash
pip install python-dotenv
```

## 配置

1. 复制 `.env.example` 为 `.env`：
```bash
cp .env.example .env
```

2. 编辑 `.env` 文件，设置你的 CTP 账户信息：
```env
CTP_USER_ID=your_user_id
CTP_PASSWORD=your_password
CTP_BROKER_ID=your_broker_id
CTP_MD_FRONT=tcp://182.254.243.31:30011
CTP_TRADER_FRONT=tcp://182.254.243.31:30001
CTP_APP_ID=your_app_id          # 可选，用于穿透式认证
CTP_AUTH_CODE=your_auth_code    # 可选，用于穿透式认证
```

## 运行

```bash
python3 main.py
```

## 程序说明

程序分为两个部分：

### 1. 行情 API 示例

- 创建行情 API 实例
- 设置回调接口（`MyMdSpi`）
- 连接行情服务器
- 登录
- 订阅行情数据（`rb2605`, `lc2605`）

### 2. 交易 API 示例

- 创建交易 API 实例
- 设置回调接口（`MyTraderSpi`）
- 连接交易服务器
- 穿透式认证（如果配置了 AppID 和 AuthCode）
- 登录

## 自定义回调

程序中的 `MyMdSpi` 和 `MyTraderSpi` 类继承自 `DefaultMdSpi` 和 `DefaultTraderSpi`，只需实现需要的方法：

```python
class MyMdSpi(DefaultMdSpi):
    def OnFrontConnected(self):
        print("行情服务器连接成功")
    
    def OnRtnDepthMarketData(self, pDepthMarketData):
        # 处理行情数据
        pass
```

## 注意事项

1. **库路径**：程序会自动尝试加载 CTP 库，搜索路径包括：
   - `./libs`
   - `../libs`
   - `../../libs`
   - 系统库路径
   - 环境变量 `CTP_LIB_PATH` 指定的路径

2. **Flow 目录**：程序会在当前目录创建 `./flow` 和 `./flow_trader` 目录，用于存储订阅信息文件。

3. **字符串编码**：CTP API 使用 GB18030 编码，程序会自动处理编码转换。

4. **异步回调**：回调函数在 C 线程中执行，需要注意线程安全。

## 与 Go 版本的对比

Python 版本与 Go 版本功能相同，主要区别：

- **类型系统**：Python 使用 `ctypes`，Go 使用 `purego`
- **字符串处理**：Python 需要手动转换字节数组，Go 有自动转换
- **回调处理**：Python 使用 `CFUNCTYPE`，Go 使用 `purego.NewCallback`
- **错误处理**：Python 使用异常，Go 使用错误返回值

## 故障排除

如果遇到库加载失败：

1. 检查库文件是否存在
2. 设置 `CTP_LIB_PATH` 环境变量指向库文件所在目录
3. 检查库文件权限
4. 检查依赖库是否已安装
