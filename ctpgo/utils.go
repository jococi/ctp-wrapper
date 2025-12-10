package ctpgo

import (
	"unsafe"

	"golang.org/x/text/encoding/simplifiedchinese"
)

// CString 将 Go 字符串转换为 C 字符串（以 null 结尾的字节切片）
// 注意：返回的指针指向 Go 管理的内存，在传递给 C 后需要确保其生命周期
func CString(s string) *byte {
	if s == "" {
		return nil
	}
	bs := make([]byte, len(s)+1)
	copy(bs, s)
	bs[len(s)] = 0 // null terminator
	return &bs[0]
}

// CStringArray 将 Go 字符串切片转换为 C 字符串数组
// 返回指向字符串指针数组的指针和底层数据（需要保持引用防止 GC）
func CStringArray(ss []string) (**byte, [][]byte) {
	if len(ss) == 0 {
		return nil, nil
	}
	
	// 创建字节切片数组保存字符串数据
	data := make([][]byte, len(ss))
	ptrs := make([]*byte, len(ss))
	
	for i, s := range ss {
		data[i] = make([]byte, len(s)+1)
		copy(data[i], s)
		data[i][len(s)] = 0
		ptrs[i] = &data[i][0]
	}
	
	return &ptrs[0], data
}

// GoString 将 C 字符串（*byte）转换为 Go 字符串
func GoString(ptr *byte) string {
	if ptr == nil {
		return ""
	}
	
	var length int
	for p := ptr; *p != 0; p = (*byte)(unsafe.Add(unsafe.Pointer(p), 1)) {
		length++
	}
	
	return string(unsafe.Slice(ptr, length))
}

// BytesToString 将固定长度字节数组转换为字符串（去除尾部的 null）
func BytesToString(b []byte) string {
	for i, v := range b {
		if v == 0 {
			return string(b[:i])
		}
	}
	return string(b)
}

// GB18030 将 GB18030 编码的字节切片转换为 UTF-8 字符串
func GB18030(b []byte) string {
	// 找到 null 终止符
	var end int
	for end = 0; end < len(b); end++ {
		if b[end] == 0 {
			break
		}
	}
	if end == 0 {
		return ""
	}
	
	decoder := simplifiedchinese.GB18030.NewDecoder()
	result, err := decoder.Bytes(b[:end])
	if err != nil {
		return string(b[:end])
	}
	return string(result)
}

// StringToBytes 将字符串复制到固定长度字节数组
func StringToBytes(s string, size int) []byte {
	b := make([]byte, size)
	copy(b, s)
	return b
}

// CopyStringToArray 将字符串复制到字节数组（用于填充 CTP 结构体字段）
func CopyStringToArray(dst []byte, src string) {
	copy(dst, src)
}

// BoolToInt 将 bool 转换为 int（C 风格）
func BoolToInt(b bool) int {
	if b {
		return 1
	}
	return 0
}

// IntToBool 将 int 转换为 bool（C 风格）
func IntToBool(i int) bool {
	return i != 0
}
