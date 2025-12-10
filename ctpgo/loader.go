package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 动态库加载

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/ebitengine/purego"
)

var (
	mdLib     uintptr
	traderLib uintptr
)

// LoadLibrary 加载 CTP 动态库
// libPath 为库文件所在目录路径
func LoadLibrary(libPath string) error {
	var mdLibName, traderLibName string
	
	switch runtime.GOOS {
	case "windows":
		mdLibName = "thostmduserapi_se.dll"
		traderLibName = "thosttraderapi_se.dll"
	case "linux":
		mdLibName = "thostmduserapi_se.so"
		traderLibName = "thosttraderapi_se.so"
	case "darwin":
		mdLibName = "thostmduserapi_se.framework/thostmduserapi_se"
		traderLibName = "thosttraderapi_se.framework/thosttraderapi_se"
	default:
		return fmt.Errorf("unsupported platform: %s", runtime.GOOS)
	}
	
	mdPath := filepath.Join(libPath, mdLibName)
	traderPath := filepath.Join(libPath, traderLibName)
	
	// 检查文件是否存在
	if _, err := os.Stat(mdPath); err != nil {
		return fmt.Errorf("md library not found: %s", mdPath)
	}
	if _, err := os.Stat(traderPath); err != nil {
		return fmt.Errorf("trader library not found: %s", traderPath)
	}
	
	// 加载行情库
	var err error
	mdLib, err = purego.Dlopen(mdPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load md library: %w", err)
	}
	
	// 加载交易库
	traderLib, err = purego.Dlopen(traderPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load trader library: %w", err)
	}
	
	// 初始化 API 函数
	initMdApi(mdLib)
	initTraderApi(traderLib)
	
	return nil
}

// LoadCTPLibrary 从 C 包装库加载（包含回调支持）
// libPath 为 ctp_md_c_api 和 ctp_trader_c_api 库文件所在目录
func LoadCTPLibrary(libPath string) error {
	var mdLibName, traderLibName string
	
	switch runtime.GOOS {
	case "windows":
		mdLibName = "ctp_md_c_api.dll"
		traderLibName = "ctp_trader_c_api.dll"
	case "linux":
		mdLibName = "libctp_md_c_api.so"
		traderLibName = "libctp_trader_c_api.so"
	case "darwin":
		mdLibName = "libctp_md_c_api.dylib"
		traderLibName = "libctp_trader_c_api.dylib"
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
	mdLib, err = purego.Dlopen(mdPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
	if err != nil {
		return fmt.Errorf("failed to load md C wrapper library: %w", err)
	}
	
	// 加载交易 C 包装库
	traderLib, err = purego.Dlopen(traderPath, purego.RTLD_NOW|purego.RTLD_GLOBAL)
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
