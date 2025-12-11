//go:build darwin || freebsd || linux || netbsd

package ctpgo

import "github.com/ebitengine/purego"

// openLibrary 在 Unix 上使用 purego.Dlopen 加载动态库
func openLibrary(path string) (uintptr, error) {
	return purego.Dlopen(path, purego.RTLD_NOW|purego.RTLD_GLOBAL)
}
