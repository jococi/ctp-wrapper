# CTP接口封装项目

## 项目简介

本项目是一个对上海期货交易所CTP(Comprehensive Transaction Platform)交易接口的跨语言封装工具。它可以将CTP的C++接口转换为多种编程语言的接口，目前支持：

- C/C++动态链接库
- Golang
- Python

通过这个项目，开发者可以在不同的编程语言环境中使用CTP接口进行期货交易和行情数据获取，无需深入了解底层C++实现细节。

## 项目目录结构

```
.
├── ctpapi/              # 存放从CTP官网下载的SDK
│   ├── linux/           # Linux平台SDK
│   ├── macos/           # macOS平台SDK
│   └── windows/         # Windows平台SDK
├── wrapctp_gen/         # 接口封装代码生成工具
│   ├── wrap_tpl/        # 模板文件
│   └── wrapper_gen.go   # 代码生成器主程序
├── csrc/                # 生成的C/C++封装代码
│   ├── linux/           # Linux平台C/C++代码
│   ├── macos/           # macOS平台C/C++代码
│   └── windows/         # Windows平台C/C++代码
├── pyctp/               # 生成的Python语言封装
├── ctpgo/               # 生成的Golang语言封装
├── libs/                # 编译后的动态库存放目录
├── example/             # 使用示例
│   ├── test_py.py       # Python使用示例
│   └── test_go.go       # Golang使用示例
├── macos/               # macOS平台编译工程文件
└── windows/             # Windows平台编译工程文件
```

## 编译步骤

### 1. 准备CTP SDK

1. 从[http://www.sfit.com.cn/5_2_DocumentDown_2.htm](http://www.sfit.com.cn/5_2_DocumentDown_2.htm)下载对应的行情和交易API文件
2. 将下载的不同系统的接口文件，按照项目根目录下的`csource/ctpapi`的`linux`、`macos`、`windows`文件夹里的命名格式进行命名，并存放/替换到对应文件夹。

```
ctpapi/
├── linux/
│   ├── ThostFtdcMdApi.h
│   ├── ThostFtdcTraderApi.h
│   ├── ThostFtdcUserApiDataType.h
│   ├── ThostFtdcUserApiStruct.h
│   ├── libthostmduserapi_se.so
│   └── libthosttraderapi_se.so
├── macos/
│   ├── ThostFtdcMdApi.h
│   ├── ThostFtdcTraderApi.h
│   ├── ThostFtdcUserApiDataType.h
│   ├── ThostFtdcUserApiStruct.h
│   ├── libthostmduserapi_se.dylib
│   └── libthosttraderapi_se.dylib
└── windows/
    ├── ThostFtdcMdApi.h
    ├── ThostFtdcTraderApi.h
    ├── ThostFtdcUserApiDataType.h
    ├── ThostFtdcUserApiStruct.h
    ├── thostmduserapi_se.dll
    └── thosttraderapi_se.dll
```

### 2. 生成封装代码

进入wrapctp_gen目录，使用以下命令生成各平台的接口代码：

```shell
# 生成C/C++接口代码
go run wrapper_gen.go -csys macos -lang c -srcpath ../ctpapi/ -outpath ../csrc/macos
go run wrapper_gen.go -csys windows -lang c -srcpath ../ctpapi/ -outpath ../csrc/windows
go run wrapper_gen.go -csys linux -lang c -srcpath ../ctpapi/ -outpath ../csrc/linux

# 生成Python接口代码
go run wrapper_gen.go -csys macos -lang python -srcpath ../ctpapi/ -outpath ../pyctp/macos
go run wrapper_gen.go -csys windows -lang python -srcpath ../ctpapi/ -outpath ../pyctp/windows
go run wrapper_gen.go -csys linux -lang python -srcpath ../ctpapi/ -outpath ../pyctp/linux

# 生成Golang接口代码
go run wrapper_gen.go -csys macos -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/
go run wrapper_gen.go -csys windows -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/
go run wrapper_gen.go -csys linux -lang golang -srcpath ../ctpapi/ -outpath ../ctpgo/
```

### 3. 编译C/C++动态链接库

#### macOS平台 (使用Clang)

```shell
clang++ -shared -fPIC -std=c++11 \
  -o libs/libctpquote_api.dylib \
  -I. -O3 \
  -L./ctpapi/macos \
  -lthostmduserapi_se -lthosttraderapi_se -lMacDataCollect \
  -install_name @rpath/libctpquote_api.dylib \
  csrc/macos/ctpquote_api.cpp

clang++ -shared -fPIC -std=c++11 \
  -o libs/libctptrade_api.dylib \
  -I. -O3 \
  -L./ctpapi/macos \
  -lthostmduserapi_se -lthosttraderapi_se -lMacDataCollect \
  -framework IOKit -framework Foundation -framework CoreFoundation \
  -install_name @rpath/libctptrade_api.dylib \
  csrc/macos/ctptrade_api.cpp
```

#### macOS平台 (Xcode)

打开项目根目录下的macos文件夹下的AlgoTrade工作空间，对其中项目进行编译。

__如果使用Xcode进行编译，需要在Xcode的 Build Settings -> Apple Clang - Code Generation -> Symbols Hidden by Default设置为`No`，或者在需要导出的函数前添加 `__attribute__((visibility("default")))`__

1. 将生成动态链接库文件，默认在项目根目录`macos/DerivedData/Build/Products/Release`下面的`.dylib`文件拷贝到项目根目录`libs`下。
2. 同时，也可以将该路径添加到系统环境变量。

#### Linux平台 (使用G++)

```shell
g++ -shared -fPIC -std=c++11 \
  -o libs/libctpquote_api.so \
  -I. -O3 \
  -L./ctpapi/linux \
  -Wl,-rpath,'$ORIGIN' \
  csrc/linux/ctpquote_api.cpp \
  -lthostmduserapi_se -lthosttraderapi_se -lLinuxDataCollect

g++ -shared -fPIC -std=c++11 \
  -o libs/libctptrade_api.so \
  -I. -O3 \
  -L./ctpapi/linux \
  -Wl,-rpath,'$ORIGIN' \
  csrc/linux/ctptrade_api.cpp \
  -lthostmduserapi_se -lthosttraderapi_se -lLinuxDataCollect
```

2. 将 _SFIT接口下载_ 的`.so`文件拷贝到项目根目录`libs/`下，并且将动态链接库所在路径即项目根目录`libs/`添加到系统路径里面，命令语句：`export LD_LIBRARY_PATH=$LD_LIBRARY_PATH:<root of project>/libs/` （添加完成后需要`source <configuration file>`，立即生效）


3. 设置Linux信息采集库权限`chmod u+s libLinuxDataCollect.so`

#### Windows平台 (使用Visual Studio)

1. 使用Visual Studio打开项目根目录的`windows`文件夹下的AlgoTrade解决方案，对其中项目进行编译。
2. 将生成动态链接库文件，默认在项目根目录`windows\x64\Release`下面的`.dll`文件和 _SFIT接口下载_ 的`.dll`文件拷贝到项目根目录`libs`下。
3. 为了方便全局使用，将该路径添加到系统环境变量。（添加完成后，需要重启terminal/cmd）。

### 4. 安装和使用

#### Python使用方法

开发模式安装:
```shell
pip3 install -e .
```

VSCode配置（使用代码提示的配置）:
```json
{
    "python.analysis.typeCheckingMode": "basic",
    "python.analysis.autoImportCompletions": true,
    "python.analysis.indexing": true,
    "python.analysis.packageIndexDepths": [
        {
            "name": "pyctp",
            "depth": 10
        }
    ],
    "python.languageServer": "Pylance",
    "python.analysis.extraPaths": [
        "./ctp-wrapper"
    ]
}
```

#### Golang使用方法

在go.mod中添加replace指令:
```
replace ctpgo => /path/to/ctp-wrapper
require ctpgo v0.0.1
```

## 使用示例

### Python示例

```python
import pyctp

# 创建行情接口实例
cq = pyctp.Quote()
print(cq.GetApiVersion())

# 创建交易接口实例
ctpgo = pyctp.Trade()
print(ctpgo.GetApiVersion())
```

### Golang示例

```go
package main

import (
	"ctpgo/ctpgo"
	"fmt"
)

func main() {
	// 创建行情接口实例
	pq := ctpgo.InitQuote()
	fmt.Println(pq.GetApiVersion())

	// 创建交易接口实例
	pt := ctpgo.InitTrade()
	fmt.Println(pt.GetApiVersion())
}
```

更多详细示例请参考`example`目录下的示例程序。