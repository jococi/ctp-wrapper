# PureGo 代码生成器

生成使用 purego（而非 cgo）的 Go 包装代码。

## 功能

- 解析 C API 头文件（`ctp_trader_c_api.h`, `ctp_md_c_api.h`）
- 生成使用 purego 的 Go 包装代码
- 不使用 cgo，纯 Go 实现
- 自动处理类型转换（字符串、指针、句柄等）

## 使用方法

```bash
# 基本用法
python3 generate_purego.py --input ../csrc --output ./output

# 指定动态库名称
python3 generate_purego.py \
    --input ../csrc \
    --output ./output \
    --lib-md libctpmd_c.dylib \
    --lib-trader libctptrader_c.dylib
```

## 生成的文件

- `ctp_md_purego.go` - MdApi 的 purego 包装
- `ctp_trader_purego.go` - TraderApi 的 purego 包装

## 使用示例

```go
package main

import (
    "fmt"
    "ctptrader_purego"
)

func main() {
    // 创建 API 实例
    api := ctptrader_purego.CreateFtdcTraderApi("./log/")
    
    // 获取版本信息
    version := ctptrader_purego.GetApiVersion()
    fmt.Printf("API Version: %s\n", version)
    
    // 初始化
    ctptrader_purego.Init(api)
    
    // ... 其他操作
}
```

## 注意事项

1. **依赖**: 需要安装 `github.com/ebitengine/purego` 包
2. **动态库路径**: 生成器会尝试从 `../libs` 目录加载动态库
3. **字符串处理**: 当前实现使用简化的字符串转换，生产环境可能需要使用 C 的 malloc/free
4. **回调函数**: 回调函数的处理需要额外实现（当前生成器主要处理 API 函数）

## 与 cgo 版本的区别

| 特性 | cgo 版本 | purego 版本 |
|------|---------|------------|
| 编译要求 | 需要 C 编译器 | 不需要 C 编译器 |
| 交叉编译 | 困难 | 容易 |
| 性能 | 取决于场景：<br>- 单次 API 调用：略好（官方优化，类型转换自动）<br>- 高频回调（如行情）：可能略差（cgo 上下文切换开销累积） | 取决于场景：<br>- 单次 API 调用：可能略差（手动类型转换）<br>- 高频回调：可能略好（避免 cgo 上下文切换） |
| 二进制大小 | 较大 | 较小 |

**性能说明**：对于 CTP 这种场景，API 调用本身耗时（网络 I/O、数据处理）远大于调用开销，两者性能差异通常可忽略。选择主要考虑编译和部署便利性。

## 改进建议

1. 使用 C 的 malloc/free 进行字符串内存管理
2. 实现回调函数的 purego 包装
3. 添加错误处理机制
4. 添加单元测试
