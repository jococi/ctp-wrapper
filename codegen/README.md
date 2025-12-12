# CTP 代码生成器

本目录包含三个代码生成器，用于将 CTP 官方 C++ API 转换为跨语言的统一接口。

## 代码生成流程

```text
CTP C++ API (官方)
    ↓ [步骤 1: generate_c_api.py]
纯 C API (跨平台统一)
    ↓ [步骤 2: generate_go_api.py 或 generate_py_api.py]
Go/Python 包装代码
```

---

## 步骤 1: 生成 C API (`generate_c_api.py`)

将 CTP 官方 C++ API 转换为纯 C 接口，支持多实例，跨平台统一。

### 特性

- **纯 C 接口**：生成的头文件不依赖任何 C++ 语法
- **不透明指针**：使用句柄类型隐藏实现细节
- **userData 支持**：所有回调携带 `userData` 参数，支持多实例
- **驼峰命名**：函数和回调使用驼峰命名风格
- **跨平台统一**：支持 Windows/Linux/macOS，统一处理平台差异

### 三平台 API 差异分析

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

### 使用方法

**运行目录**：在 `codegen/` 目录下运行（即当前目录）

```bash
# 基本用法
python3 generate_c_api.py --input ../ctpapi/linux --output ../csrc

# 指定不同平台的头文件
python3 generate_c_api.py --input ../ctpapi/macos --output ../csrc
python3 generate_c_api.py --input ../ctpapi/windows --output ../csrc
```

**说明**：路径使用相对路径，相对于 `codegen/` 目录：

- `../ctpapi/` 指向 `ctp-wrapper/ctpapi/`
- `../csrc` 指向 `ctp-wrapper/csrc/`

### 生成的文件

```text
csrc/
├── ctp_md_c_api.h      # MdApi 的纯 C 头文件
├── ctp_md_c_api.cpp    # MdApi 的 C++ 实现
├── ctptrader_c_api.h    # TraderApi 的纯 C 头文件
└── ctptrader_c_api.cpp  # TraderApi 的 C++ 实现
```

### API 示例

#### 创建和初始化

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

#### 回调函数

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

#### 多实例支持

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

### 编译 C API

**注意**：生成的 C++ 实现文件包含条件编译，使用 `__APPLE__`、`_WIN32`、`__linux__` 宏自动适配平台。

推荐使用 `ctp-wrapper` 目录下的构建脚本进行编译，这些脚本会自动处理平台差异、依赖库复制和路径配置。

#### Linux / macOS

使用 Makefile 编译：

```bash
cd ../ctp-wrapper

# 编译所有库（行情API + 交易API）
make

# 仅编译行情API
make md

# 仅编译交易API
make trader

# 清理编译产物
make clean

# 查看帮助信息
make help
```

编译后的库文件会输出到 `libs/` 目录：

- Linux: `libs/libctpmd_c_api.so`, `libs/libctptrader_c_api.so`
- macOS: `libs/libctpmd_c_api.dylib`, `libs/libctptrader_c_api.dylib`

#### Windows

使用 `build.bat` 脚本编译：

```batch
cd ..\ctp-wrapper

REM 编译所有库（行情API + 交易API）
build.bat

REM 清理编译产物
build.bat clean

REM 查看帮助信息
build.bat help
```

编译后的库文件会输出到 `libs\` 目录：

- `libs\ctpmd_c_api.dll`
- `libs\ctptrader_c_api.dll`

**注意**：Windows 构建脚本会自动查找 Visual Studio 安装路径并设置 MSVC 环境。如果已从 Developer Command Prompt 运行，则直接使用已配置的环境。

---

## 步骤 2A: 生成 Go 包装 (`generate_go_api.py`)

将 C API 转换为 Go PureGo 包装代码，使用 `github.com/ebitengine/purego` 进行 FFI 调用。

### Go 包装特性

- **PureGo 实现**：使用 purego 库，无需 CGO
- **类型安全**：完整的 Go 类型定义和转换
- **多实例支持**：通过 userData 机制支持多实例
- **自动格式化**：生成后自动使用 `gofmt` 格式化代码

### Go 包装使用方法

**运行目录**：在 `codegen/` 目录下运行（即当前目录）

```bash
# 基本用法（需要先完成步骤 1）
python3 generate_go_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../ctpgo
```

**说明**：路径使用相对路径，相对于 `codegen/` 目录：

- `../csrc` 指向 `ctp-wrapper/csrc/`
- `../ctpapi/` 指向 `ctp-wrapper/ctpapi/`
- `../ctpgo` 指向 `ctp-wrapper/ctpgo/`

**参数说明**：

- `--input`: C API 头文件目录（包含 `ctp_md_c_api.h`, `ctptrader_c_api.h`）
- `--struct`: CTP 结构体头文件目录（包含 `ThostFtdcUserApiDataType.h`, `ThostFtdcUserApiStruct.h`）
- `--output`: 输出目录（Go 代码将生成到此目录）

### Go 包装生成的文件

```text
ctpgo/
├── go.mod                    # Go 模块定义
├── loader.go                 # 动态库加载（通用）
├── loader_unix.go            # Unix 平台动态库加载
├── loader_windows.go         # Windows 平台动态库加载
├── utils.go                  # 工具函数（字符串转换等）
├── datatype.go              # 枚举和类型别名定义
├── struct.go                # 结构体定义
├── md_api.go                # 行情 API 包装
├── trader_api.go            # 交易 API 包装
├── md_callbacks.go          # 行情回调接口定义
├── trader_callbacks.go      # 交易回调接口定义
├── md_default_spi.go        # 行情默认 SPI 实现
└── trader_default_spi.go    # 交易默认 SPI 实现
```

### Go 包装使用示例

```go
package main

import (
    "fmt"
    "unsafe"
    "github.com/your-org/ctpgo"
)

type MyStrategy struct {
    name string
}

func (s *MyStrategy) OnFrontConnected() {
    fmt.Printf("Strategy %s connected!\n", s.name)
}

func main() {
    // 自动加载动态库
    err := ctpgo.AutoLoadLibrary()
    if err != nil {
        panic(err)
    }
    
    // 创建 API 实例
    api := ctpgo.CreateFtdcMdApi("./log/", false, false)
    defer api.Release()
    
    // 创建策略实例
    strategy := &MyStrategy{name: "测试策略"}
    
    // 创建 SPI 并设置回调
    spi := ctpgo.NewMdSpi(unsafe.Pointer(strategy))
    spi.SetOnFrontConnected(func(userData unsafe.Pointer) {
        s := (*MyStrategy)(userData)
        s.OnFrontConnected()
    })
    
    // 注册 SPI
    api.RegisterSpi(spi)
    
    // 注册前置地址
    api.RegisterFront("tcp://180.168.146.187:10131")
    
    // 初始化
    api.Init()
    
    // 等待连接...
}
```

---

## 步骤 2B: 生成 Python 包装 (`generate_py_api.py`)

将 C API 转换为 Python ctypes 包装代码，提供 Pythonic 的接口。

### Python 包装特性

- **ctypes 实现**：使用 Python 标准库 ctypes，无需额外依赖
- **类型安全**：完整的 Python 类型定义和转换
- **多实例支持**：通过 userData 机制支持多实例
- **Pythonic API**：符合 Python 命名规范和习惯用法

### Python 包装使用方法

**运行目录**：在 `codegen/` 目录下运行（即当前目录）

```bash
# 基本用法（需要先完成步骤 1）
python3 generate_py_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../pyctp
```

**说明**：路径使用相对路径，相对于 `codegen/` 目录：

- `../csrc` 指向 `ctp-wrapper/csrc/`
- `../ctpapi/` 指向 `ctp-wrapper/ctpapi/`
- `../pyctp` 指向 `ctp-wrapper/pyctp/`

**参数说明**：

- `--input`: C API 头文件目录（包含 `ctp_md_c_api.h`, `ctptrader_c_api.h`）
- `--struct`: CTP 结构体头文件目录（包含 `ThostFtdcUserApiDataType.h`, `ThostFtdcUserApiStruct.h`）
- `--output`: 输出目录（Python 代码将生成到此目录）

### Python 包装生成的文件

```text
pyctp/
├── __init__.py              # 包初始化（导出主要接口）
├── loader.py                # 动态库加载
├── utils.py                 # 工具函数（字符串转换等）
├── datatype.py             # 枚举和类型别名定义
├── struct.py               # 结构体定义
├── md_api.py               # 行情 API 包装
├── trader_api.py           # 交易 API 包装
├── md_callbacks.py         # 行情回调接口定义
├── trader_callbacks.py     # 交易回调接口定义
├── md_default_spi.py       # 行情默认 SPI 实现
└── trader_default_spi.py   # 交易默认 SPI 实现
```

### Python 包装使用示例

```python
from pyctp import MdApi, MdSpi, auto_load_library

class MyStrategy:
    def __init__(self, name):
        self.name = name
    
    def on_front_connected(self):
        print(f"Strategy {self.name} connected!")

def main():
    # 自动加载动态库
    auto_load_library()
    
    # 创建 API 实例
    api = MdApi.CreateFtdcMdApi("./log/", False, False)
    
    # 创建策略实例
    strategy = MyStrategy("测试策略")
    
    # 创建 SPI 并设置回调
    spi = MdSpi(strategy)
    spi.set_on_front_connected(lambda user_data: 
        user_data.on_front_connected())
    
    # 注册 SPI
    api.RegisterSpi(spi)
    
    # 注册前置地址
    api.RegisterFront("tcp://180.168.146.187:10131")
    
    # 初始化
    api.Init()
    
    # 等待连接...

if __name__ == "__main__":
    main()
```

---

## 完整生成流程示例

### 一次性生成所有代码

**运行目录说明**：

- 步骤 1、3、4：在 `codegen/` 目录下运行
- 步骤 2：在 `ctp-wrapper/` 目录（项目根目录）下运行

```bash
# 1. 生成 C API（从 CTP C++ API）
# 在 codegen/ 目录下运行
cd codegen
python3 generate_c_api.py --input ../ctpapi/linux --output ../csrc

# 2. 编译 C API 动态库（使用 Makefile）
# 在 ctp-wrapper/ 目录（项目根目录）下运行
cd ..
make

# 或者 Windows 平台使用：
# build.bat

# 3. 生成 Go 包装代码
# 在 codegen/ 目录下运行
cd codegen
python3 generate_go_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../ctpgo

# 4. 生成 Python 包装代码
# 在 codegen/ 目录下运行
python3 generate_py_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../pyctp
```

### 目录结构

```text
ctp-wrapper/
├── codegen/              # 代码生成器（本目录）
│   ├── generate_c_api.py
│   ├── generate_go_api.py
│   ├── generate_py_api.py
│   └── README.md
├── csrc/                 # 生成的 C API（步骤 1 输出）
│   ├── ctp_md_c_api.h
│   ├── ctp_md_c_api.cpp
│   ├── ctptrader_c_api.h
│   └── ctptrader_c_api.cpp
├── ctpgo/                # 生成的 Go 包装（步骤 2A 输出）
│   ├── md_api.go
│   ├── trader_api.go
│   └── ...
├── pyctp/                # 生成的 Python 包装（步骤 2B 输出）
│   ├── md_api.py
│   ├── trader_api.py
│   └── ...
└── ctpapi/               # CTP 官方 API（输入）
    ├── linux/
    ├── macos/
    └── windows/
```

---

## 注意事项

### C API 生成器注意事项

1. **结构体兼容**：CTP 的 Field 结构体是 C 兼容的 POD 类型，可直接使用
2. **内存管理**：API 和 SPI 实例需要手动释放
3. **线程安全**：CTP 回调在内部线程触发，注意线程同步

### Go 包装注意事项

1. **动态库路径**：确保编译好的 C API 动态库在系统库路径中，或使用 `loader.go` 指定路径
2. **依赖管理**：生成的代码依赖 `github.com/ebitengine/purego` 和 `golang.org/x/text`
3. **类型转换**：字符串类型需要手动转换（使用 `utils.go` 中的辅助函数）

### Python 包装注意事项

1. **动态库路径**：确保编译好的 C API 动态库在系统库路径中，或使用 `loader.py` 指定路径
2. **Python 版本**：建议使用 Python 3.7+
3. **类型转换**：字符串类型需要手动转换（使用 `utils.py` 中的辅助函数）

---

## 与旧版封装对比

| 特性 | 旧版（全局变量） | 新版（userData） |
|-----|----------------|-----------------|
| 多实例 | ❌ 不支持 | ✅ 支持 |
| 回调串台 | ⚠️ 可能发生 | ✅ 不会 |
| FFI 友好 | ⚠️ 需要 hack | ✅ 标准做法 |
| 性能 | 相同 | 相同 |

---

## 故障排查

### 常见问题

1. **找不到 CTP 头文件**
   - 确保 `--input` 或 `--struct` 参数指向正确的 CTP API 目录
   - 检查目录中是否包含必要的头文件

2. **动态库加载失败**
   - 检查动态库文件是否存在
   - 检查动态库路径是否正确
   - 检查动态库依赖的其他库是否可用

3. **代码生成失败**
   - 检查 Python 版本（建议 3.7+）
   - 检查输入文件格式是否正确
   - 查看错误信息中的具体位置

4. **Go 代码格式化失败**
   - 确保已安装 Go 并添加到 PATH
   - 检查 `gofmt` 命令是否可用

---

## 贡献指南

如果需要修改代码生成逻辑：

1. **修改 C API 生成器**：编辑 `generate_c_api.py`
2. **修改 Go 包装生成器**：编辑 `generate_go_api.py`
3. **修改 Python 包装生成器**：编辑 `generate_py_api.py`

修改后重新运行生成器即可更新生成的代码。

**注意**：生成的代码文件（`csrc/`, `ctpgo/`, `pyctp/` 目录下的文件）是自动生成的，请勿手动修改。所有修改应在生成器脚本中进行。
