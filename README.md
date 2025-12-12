# CTP 接口封装项目

## 项目简介

本项目是一个对上海期货交易所 CTP (Comprehensive Transaction Platform) 交易接口的跨语言封装工具。它可以将 CTP 的 C++ 接口转换为多种编程语言的接口，目前支持：

- **C/C++ 动态链接库**：纯 C API，支持多实例，跨平台统一
- **Golang**：使用 PureGo 实现，无需 CGO
- **Python**：使用 ctypes 实现，Pythonic 接口

通过这个项目，开发者可以在不同的编程语言环境中使用 CTP 接口进行期货交易和行情数据获取，无需深入了解底层 C++ 实现细节。

## 项目目录结构

```text
ctp-wrapper/
├── ctpapi/              # CTP 官方 API（需要从官网下载）
│   ├── linux/           # Linux 平台 SDK
│   ├── macos/           # macOS 平台 SDK
│   └── windows/         # Windows 平台 SDK
│   └── README.md        # CTP API 下载和安装说明
├── codegen/             # 代码生成器
│   ├── generate_c_api.py    # 生成 C API
│   ├── generate_go_api.py   # 生成 Go 包装
│   ├── generate_py_api.py   # 生成 Python 包装
│   └── README.md        # 代码生成器使用说明
├── csrc/                # 生成的 C API 代码
│   ├── ctpmd_c_api.h/cpp      # 行情 API
│   └── ctptrader_c_api.h/cpp  # 交易 API
├── ctpgo/               # 生成的 Golang 语言封装
├── pyctp/               # 生成的 Python 语言封装
├── libs/                # 编译后的动态库存放目录
├── example/             # 使用示例
│   ├── main.go          # Go 使用示例
│   ├── main.py          # Python 使用示例
│   └── README.md        # 示例说明
├── Makefile             # Linux/macOS 编译脚本
├── build.bat            # Windows 编译脚本
└── README.md            # 本文件
```

## 快速开始

### 1. 准备 CTP SDK

**详细说明请参考**：[`ctpapi/README.md`](ctpapi/README.md)

简要步骤：

1. 从[上期所仿真交易平台](https://www.simnow.com.cn/static/apiDownload.action)下载对应平台的行情和交易 API 文件
2. 将下载的文件解压后，按照目录结构放置到 `ctpapi/linux/`、`ctpapi/macos/` 或 `ctpapi/windows/` 目录
3. **macOS 平台重要**：进入 `ctpapi/macos/` 目录，运行 `sh setup_frameworks.sh` 脚本

### 2. 生成封装代码

**详细说明请参考**：[`codegen/README.md`](codegen/README.md)

代码生成分为两个步骤：

#### 步骤 1: 生成 C API

在 `codegen/` 目录下运行：

```bash
cd codegen

# Linux 平台
python3 generate_c_api.py --input ../ctpapi/linux --output ../csrc

# macOS 平台
python3 generate_c_api.py --input ../ctpapi/macos --output ../csrc

# Windows 平台
python3 generate_c_api.py --input ../ctpapi/windows --output ../csrc
```

#### 步骤 2: 编译 C API 动态库

在项目根目录 `ctp-wrapper/` 下运行：

**Linux / macOS**：

```bash
# 编译所有库（行情API + 交易API）
make

# 仅编译行情API
make md

# 仅编译交易API
make trader

# 清理编译产物
make clean
```

**Windows**：

```batch
REM 编译所有库
build.bat

REM 清理编译产物
build.bat clean
```

编译后的库文件会输出到 `libs/` 目录。

#### 步骤 3: 生成 Go/Python 包装代码

在 `codegen/` 目录下运行：

```bash
cd codegen

# 生成 Go 包装代码
python3 generate_go_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../ctpgo

# 生成 Python 包装代码
python3 generate_py_api.py \
    --input ../csrc \
    --struct ../ctpapi/linux \
    --output ../pyctp
```

**注意**：`--struct` 参数需要根据你的平台选择对应的目录（`linux`、`macos` 或 `windows`）。

### 3. 使用生成的代码

#### Python 使用示例

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

#### Go 使用示例

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

更多详细示例请参考 [`example/`](example/) 目录下的示例程序。

## 完整工作流程

```bash
# 1. 准备 CTP SDK（参考 ctpapi/README.md）
# 下载并解压到 ctpapi/ 对应平台目录
# macOS 需要运行: cd ctpapi/macos && sh setup_frameworks.sh

# 2. 生成 C API
cd codegen
python3 generate_c_api.py --input ../ctpapi/linux --output ../csrc

# 3. 编译 C API 动态库
cd ..
make  # 或 Windows: build.bat

# 4. 生成 Go/Python 包装代码
cd codegen
python3 generate_go_api.py --input ../csrc --struct ../ctpapi/linux --output ../ctpgo
python3 generate_py_api.py --input ../csrc --struct ../ctpapi/linux --output ../pyctp

# 5. 使用生成的代码（参考 example/ 目录）
```

## 特性

- **跨平台支持**：支持 Linux、macOS、Windows 三个平台
- **多实例支持**：通过 userData 机制支持多实例，回调不会串台
- **类型安全**：完整的类型定义和转换
- **易于使用**：提供 Pythonic/Go 风格的接口
- **自动生成**：代码自动生成，易于维护和更新

## 相关文档

- **CTP API 下载和安装**：[`ctpapi/README.md`](ctpapi/README.md)
- **代码生成器使用**：[`codegen/README.md`](codegen/README.md)
- **使用示例**：[`example/README.md`](example/README.md)
- **上期所仿真交易平台**：[https://www.simnow.com.cn/](https://www.simnow.com.cn/)
- **API 下载页面**：[https://www.simnow.com.cn/static/apiDownload.action](https://www.simnow.com.cn/static/apiDownload.action)

## 许可证

本项目代码遵循相应的开源许可证。CTP 官方 API 文件需要遵守上期所的相关规定和许可协议。
