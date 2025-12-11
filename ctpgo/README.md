# ctpgo - CTP Go 语言封装

CTP (Comprehensive Transaction Platform) 的 Go 语言封装库。

## 安装

```bash
go get github.com/your-org/ctp-wrapper/ctpgo
```

## 依赖说明

本包依赖 CTP 的动态库文件（`.so`、`.dylib`、`.dll`），这些库文件需要单独获取。

### 获取库文件

库文件需要从以下位置获取：

1. **从项目源码编译**（推荐）：
   ```bash
   cd ctp-wrapper
   make  # 或 make md trader
   # 库文件会生成在 libs/ 目录
   ```

2. **从 GitHub Releases 下载**：
   - 访问项目 Releases 页面
   - 下载对应平台的预编译库文件
   - 解压到项目的 `libs/` 目录

3. **从 CTP 官网获取**：
   - 访问 [上期所仿真交易平台](https://www.simnow.com.cn/static/apiDownload.action)
   - 下载对应平台的 SDK
   - 按照项目结构放置库文件

## 使用方法

### 方式1：自动查找库文件（推荐）

```go
package main

import "github.com/your-org/ctp-wrapper/ctpgo"

func main() {
    // 自动在多个位置查找库文件
    if err := ctpgo.LoadCTPLibraryAuto(); err != nil {
        panic(err)
    }
    
    // 使用 API...
}
```

自动查找顺序：
1. 当前工作目录的 `libs/` 子目录
2. 可执行文件所在目录的 `libs/` 子目录
3. 环境变量 `CTP_LIB_PATH` 指定的路径
4. 系统常见路径（`/usr/local/lib`、`/opt/ctp/lib` 等）

### 方式2：手动指定库文件路径

```go
package main

import "github.com/your-org/ctp-wrapper/ctpgo"

func main() {
    // 手动指定库文件所在目录
    if err := ctpgo.LoadCTPLibrary("./libs"); err != nil {
        panic(err)
    }
    
    // 使用 API...
}
```

### 方式3：使用环境变量

```bash
export CTP_LIB_PATH=/path/to/libs
go run main.go
```

## 目录结构建议

发布应用时，建议的目录结构：

```
your-app/
├── your-app              # 可执行文件
└── libs/                 # 库文件目录
    ├── libctp_md_c_api.so
    ├── libctp_trader_c_api.so
    ├── thostmduserapi_se.so
    └── thosttraderapi_se.so
```

## 发布说明

### Go 模块发布

本包的 Go 源代码可以独立发布到：
- GitHub/GitLab
- pkg.go.dev（通过 `go get`）

### 库文件发布

动态库文件需要单独发布，建议方式：
1. **GitHub Releases**：为每个平台创建 Release，包含预编译库文件
2. **CDN/对象存储**：将库文件上传到 CDN，提供下载链接
3. **文档说明**：在 README 中说明如何获取和放置库文件

### 发布检查清单

- [ ] Go 代码已通过 `go mod tidy` 整理
- [ ] 已添加版本标签（如 `v1.0.0`）
- [ ] 已创建 GitHub Release，包含各平台的库文件
- [ ] README 已更新，说明库文件获取方式
- [ ] 已测试 `go get` 安装流程

## 示例

完整示例请参考 `example/` 目录。

## 许可证

[你的许可证]

## 相关链接

- [CTP 官网](https://www.simnow.com.cn/)
- [项目主页](https://github.com/your-org/ctp-wrapper)
