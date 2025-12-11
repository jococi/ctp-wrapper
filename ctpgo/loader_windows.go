//go:build windows

package ctpgo

import (
	"path/filepath"
	"syscall"
	"unsafe"
)

var (
	kernel32            = syscall.NewLazyDLL("kernel32.dll")
	procSetDllDirectory = kernel32.NewProc("SetDllDirectoryW")
)

// openLibrary 在 Windows 上加载动态库
// 需要先设置 DLL 搜索目录，确保能找到依赖库
func openLibrary(path string) (uintptr, error) {
	// 1. 转换为绝对路径
	absPath, err := filepath.Abs(path)
	if err != nil {
		return 0, err
	}

	// 2. 设置 DLL 搜索目录（让 Windows 能找到依赖库）
	dllDir := filepath.Dir(absPath)
	dllDirPtr, _ := syscall.UTF16PtrFromString(dllDir)
	procSetDllDirectory.Call(uintptr(unsafe.Pointer(dllDirPtr)))

	// 3. 加载库
	handle, err := syscall.LoadLibrary(absPath)
	if err != nil {
		return 0, err
	}

	return uintptr(handle), nil
}
