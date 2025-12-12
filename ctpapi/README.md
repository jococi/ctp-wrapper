# CTP 官方 API

本目录包含 CTP（综合交易平台）官方提供的行情和交易 API 文件。

## 文件来源

这些文件需要从 **上期所仿真交易平台** 下载：

🔗 **下载地址**：[https://www.simnow.com.cn/static/apiDownload.action](https://www.simnow.com.cn/static/apiDownload.action)

## 目录结构

```text
ctpapi/
├── linux/          # Linux 平台 API
├── macos/          # macOS 平台 API
└── windows/        # Windows 平台 API
```

## 各平台文件说明

### Linux 平台 (`linux/`)

**必需文件**：

- **头文件**：
  - `ThostFtdcMdApi.h` - 行情 API 头文件
  - `ThostFtdcTraderApi.h` - 交易 API 头文件
  - `ThostFtdcUserApiDataType.h` - 数据类型定义
  - `ThostFtdcUserApiStruct.h` - 结构体定义
  - `DataCollect.h` - 数据采集头文件

- **动态库**：
  - `thostmduserapi_se.so` - 行情 API 动态库
  - `thosttraderapi_se.so` - 交易 API 动态库
  - `LinuxDataCollect.so` - Linux 数据采集库

- **其他文件**：
  - `error.dtd` / `error.xml` - 错误码定义文件

### macOS 平台 (`macos/`)

**必需文件**：

- **头文件**：
  - `ThostFtdcMdApi.h` - 行情 API 头文件
  - `ThostFtdcTraderApi.h` - 交易 API 头文件
  - `ThostFtdcUserApiDataType.h` - 数据类型定义
  - `ThostFtdcUserApiStruct.h` - 结构体定义
  - `DataCollect.h` - 数据采集头文件

- **Framework**：
  - `thostmduserapi_se.framework/` - 行情 API Framework
  - `thosttraderapi_se.framework/` - 交易 API Framework
  - `MacDataCollect.framework/` - macOS 数据采集 Framework

**注意**：macOS 平台使用 Framework 格式，需要确保 Framework 结构完整。

### Windows 平台 (`windows/`)

**必需文件**：

- **头文件**：
  - `ThostFtdcMdApi.h` - 行情 API 头文件
  - `ThostFtdcTraderApi.h` - 交易 API 头文件
  - `ThostFtdcUserApiDataType.h` - 数据类型定义
  - `ThostFtdcUserApiStruct.h` - 结构体定义
  - `DataCollect.h` - 数据采集头文件

- **动态库和导入库**：
  - `thostmduserapi_se.dll` / `thostmduserapi_se.lib` - 行情 API
  - `thosttraderapi_se.dll` / `thosttraderapi_se.lib` - 交易 API
  - `WinDataCollect.dll` / `WinDataCollect.lib` - Windows 数据采集库

- **其他文件**：
  - `error.dtd` / `error.xml` - 错误码定义文件

## 下载和安装步骤

### 1. 访问下载页面

访问上期所仿真交易平台 API 下载页面：
[https://www.simnow.com.cn/static/apiDownload.action](https://www.simnow.com.cn/static/apiDownload.action)

### 2. 下载对应平台的 API 包

根据你的开发平台，下载对应的 API 包：

- **Linux**：下载 Linux 版本的行情 API 和交易 API
- **macOS**：下载 macOS 版本的行情 API 和交易 API
- **Windows**：下载 Windows 版本的行情 API 和交易 API

### 3. 解压并放置文件

将下载的文件解压后，按照上述目录结构放置到对应的平台目录中。

**Linux**：

- 解压行情 API 和交易 API 的压缩包
- 将所有文件复制到 `ctpapi/linux/` 目录

**macOS**：

- 解压行情 API 和交易 API 的压缩包
- 将 Framework 目录（`thostmduserapi_se.framework`、`thosttraderapi_se.framework`、`MacDataCollect.framework`）复制到 `ctpapi/macos/` 目录
- 将所有头文件（`.h` 文件）复制到 `ctpapi/macos/` 目录
- **重要**：进入 `ctpapi/macos/` 目录，运行 `sh setup_frameworks.sh` 脚本，该脚本会：
  - 修复 Framework 的符号链接（`Versions/Current` 等）
  - 创建头文件的软链接到 `macos/` 目录下，方便编译时引用

**Windows**：

- 解压行情 API 和交易 API 的压缩包
- 将所有文件复制到 `ctpapi/windows/` 目录

## 验证安装

安装完成后，检查以下文件是否存在：

### Linux

检查以下文件是否存在：

- `ctpapi/linux/thostmduserapi_se.so` - 行情 API 动态库
- `ctpapi/linux/thosttraderapi_se.so` - 交易 API 动态库
- `ctpapi/linux/ThostFtdcMdApi.h` - 行情 API 头文件
- `ctpapi/linux/ThostFtdcTraderApi.h` - 交易 API 头文件
- `ctpapi/linux/LinuxDataCollect.so` - 数据采集库

### macOS

检查以下文件是否存在：

- `ctpapi/macos/thostmduserapi_se.framework/` - 行情 API Framework
- `ctpapi/macos/thosttraderapi_se.framework/` - 交易 API Framework
- `ctpapi/macos/MacDataCollect.framework/` - 数据采集 Framework
- `ctpapi/macos/ThostFtdcMdApi.h` - 行情 API 头文件
- `ctpapi/macos/ThostFtdcTraderApi.h` - 交易 API 头文件

### Windows

检查以下文件是否存在：

- `ctpapi/windows/thostmduserapi_se.dll` 和 `thostmduserapi_se.lib` - 行情 API
- `ctpapi/windows/thosttraderapi_se.dll` 和 `thosttraderapi_se.lib` - 交易 API
- `ctpapi/windows/WinDataCollect.dll` 和 `WinDataCollect.lib` - 数据采集库
- `ctpapi/windows/ThostFtdcMdApi.h` - 行情 API 头文件
- `ctpapi/windows/ThostFtdcTraderApi.h` - 交易 API 头文件

## 注意事项

1. **版本兼容性**：确保下载的 API 版本与你的开发环境兼容
2. **文件完整性**：确保所有必需的文件都已正确放置
3. **权限问题**（Linux/macOS）：确保动态库文件具有执行权限
4. **macOS Framework 设置**：**重要** - 在 macOS 平台，文件复制完成后，必须进入 `ctpapi/macos/` 目录运行 `sh setup_frameworks.sh` 脚本。该脚本会修复 Framework 的符号链接并创建头文件软链接，这是 macOS 平台正常使用 CTP API 的必要步骤
5. **依赖库**：某些平台可能需要额外的系统库，请参考 CTP 官方文档

## 相关文档

- **CTP 官方文档**：下载包中通常包含开发文档
- **上期所仿真交易平台**：[https://www.simnow.com.cn/](https://www.simnow.com.cn/)
- **API 下载页面**：[https://www.simnow.com.cn/static/apiDownload.action](https://www.simnow.com.cn/static/apiDownload.action)

## 许可证

这些文件是 CTP 官方提供的 API，使用时请遵守上期所的相关规定和许可协议。
