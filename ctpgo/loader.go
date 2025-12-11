package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 动态库加载

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"sync"
)

var (
	mdLib     uintptr
	traderLib uintptr
	loadOnce  sync.Once
	loadErr   error
)

// getSystemLibPaths 根据平台返回系统库路径列表
// 包括标准系统路径和从环境变量读取的路径（类似 Python 的 sys.path）
func getSystemLibPaths() []string {
	var paths []string

	// 从系统库路径环境变量中读取（类似 Python 的 sys.path）
	switch runtime.GOOS {
	case "linux":
		// LD_LIBRARY_PATH 是 Linux 的标准库路径环境变量
		if ldPath := os.Getenv("LD_LIBRARY_PATH"); ldPath != "" {
			for _, p := range strings.Split(ldPath, ":") {
				if p != "" {
					paths = append(paths, p)
				}
			}
		}
	case "darwin":
		// DYLD_LIBRARY_PATH 是 macOS 的库路径环境变量
		if dyldPath := os.Getenv("DYLD_LIBRARY_PATH"); dyldPath != "" {
			for _, p := range strings.Split(dyldPath, ":") {
				if p != "" {
					paths = append(paths, p)
				}
			}
		}
	case "windows":
		// Windows 使用 PATH 环境变量，但通常系统会自动搜索
		// 这里可以添加一些特定路径
	}

	// 添加标准系统路径
	switch runtime.GOOS {
	case "linux":
		paths = append(paths,
			"/usr/local/lib",            // 用户安装的库
			"/usr/lib",                  // 系统库
			"/usr/lib/x86_64-linux-gnu", // Debian/Ubuntu 64位
			"/usr/lib64",                // 某些发行版的 64 位库路径
			"/opt/ctp/lib",              // CTP 专用安装路径
			"/opt/lib",                  // 通用 opt 路径
		)
	case "darwin":
		paths = append(paths,
			"/usr/local/lib",    // Homebrew (Intel)
			"/opt/homebrew/lib", // Homebrew (Apple Silicon)
			"/opt/local/lib",    // MacPorts
			"/usr/lib",          // 系统库
			"/opt/ctp/lib",      // CTP 专用安装路径
		)
	case "windows":
		// Windows 通常通过 PATH 环境变量查找，但也可以添加一些常见路径
		programFiles := os.Getenv("ProgramFiles")
		programFilesX86 := os.Getenv("ProgramFiles(x86)")
		if programFiles != "" {
			paths = append(paths, filepath.Join(programFiles, "CTP", "lib"))
		}
		if programFilesX86 != "" {
			paths = append(paths, filepath.Join(programFilesX86, "CTP", "lib"))
		}
		paths = append(paths,
			`C:\Windows\System32`, // 系统目录
			`C:\CTP\lib`,          // 常见安装路径
		)
	}

	return paths
}

var (
	// 默认库路径列表，按优先级顺序尝试
	// 可以通过环境变量 CTP_LIB_PATH 覆盖，环境变量优先级最高
	defaultLibPaths = func() []string {
		paths := []string{
			"./libs",              // 当前目录下的 libs
			"../libs",             // 上一层级下的 libs
			"../../libs",          // 上两级目录下的 libs
			"./ctp-wrapper/libs",  // 项目根目录下的 ctp-wrapper/libs
			"../ctp-wrapper/libs", // 上一层级下的 ctp-wrapper/libs
		}
		// 添加系统路径
		paths = append(paths, getSystemLibPaths()...)
		return paths
	}()
)

// LoadCTPLibrary 从 C 包装库加载（包含回调支持）
// libPath 为 ctp_md_c_api 和 ctp_trader_c_api 库文件所在目录
func LoadCTPLibrary(libPath string) error {
	var mdLibName, traderLibName string

	switch runtime.GOOS {
	case "windows":
		mdLibName = "ctpmd_c_api.dll"
		traderLibName = "ctptrader_c_api.dll"
	case "linux":
		mdLibName = "libctpmd_c_api.so"
		traderLibName = "libctptrader_c_api.so"
	case "darwin":
		mdLibName = "libctpmd_c_api.dylib"
		traderLibName = "libctptrader_c_api.dylib"
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}

	mdPath := filepath.Join(libPath, mdLibName)
	traderPath := filepath.Join(libPath, traderLibName)

	// 检查文件是否存在
	if _, err := os.Stat(mdPath); err != nil {
		return fmt.Errorf("md C wrapper library not found: %s", mdPath)
	}
	if _, err := os.Stat(traderPath); err != nil {
		return fmt.Errorf("trader C wrapper library not found: %s", traderPath)
	}

	// 加载行情 C 包装库
	var err error
	mdLib, err = openLibrary(mdPath)
	if err != nil {
		return fmt.Errorf("failed to load md C wrapper library: %w", err)
	}

	// 加载交易 C 包装库
	traderLib, err = openLibrary(traderPath)
	if err != nil {
		return fmt.Errorf("failed to load trader C wrapper library: %w", err)
	}

	// 初始化 API 函数
	initMdApi(mdLib)
	initTraderApi(traderLib)

	return nil
}

// GetMdLibHandle 获取行情库句柄
func GetMdLibHandle() uintptr {
	return mdLib
}

// GetTraderLibHandle 获取交易库句柄
func GetTraderLibHandle() uintptr {
	return traderLib
}

// autoLoadLibrary 自动加载库（只加载一次）
// 优先使用环境变量 CTP_LIB_PATH，否则按顺序尝试默认路径列表
func autoLoadLibrary() error {
	loadOnce.Do(func() {
		// 优先使用环境变量
		libPath := os.Getenv("CTP_LIB_PATH")
		if libPath != "" {
			// 环境变量指定的路径，只尝试一次
			loadErr = LoadCTPLibrary(libPath)
			return
		}

		// 按顺序尝试默认路径列表
		for _, path := range defaultLibPaths {
			err := LoadCTPLibrary(path)
			if err == nil {
				// 加载成功，直接返回
				loadErr = nil
				return
			}
			// 加载失败，继续尝试下一个路径
			loadErr = err
		}
		// 所有路径都失败，返回最后一个错误
	})
	return loadErr
}
