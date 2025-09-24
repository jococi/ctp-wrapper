package ctpgo

import (
	"os"
	"path"
	"path/filepath"
	"runtime"
	"syscall"
	"unsafe"
)

type Quote struct {
	h              *syscall.DLL
	api            uintptr
	pSpi           uintptr
	version        string
	pszFlowPath    string
	usingUdp       bool
	usingMulticast bool

    [[ range .On]]// [[ .Comment ]]
    [[ .FuncName ]]_ func([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]])
    [[ end]]
}

func InitQuote(pszFlowPath string, usingUdp bool, usingMulticast bool) *Quote {
	q := new(Quote)
	// Load DLL
	workPath, _ := os.Getwd()
	_, curFile, _, _ := runtime.Caller(1)
	dllPath := filepath.Dir(curFile)
	dllPath = path.Join(dllPath, "../libs/")
	_ = os.Chdir(dllPath)
	q.h = syscall.MustLoadDLL("ctpquote_api.dll")
	os.Chdir(workPath)

	q.pszFlowPath = pszFlowPath
	q.usingUdp = usingUdp
	q.usingMulticast = usingMulticast
	// 执行目录下创建 log目录
	_, err := os.Stat(q.pszFlowPath)
	if err != nil {
		os.Mkdir(q.pszFlowPath, os.ModePerm)
	}
	q.api = q.CreateApi()
	q.pSpi = q.CreateSpi()
	q.version = q.GetApiVersion()

	return q
}

func (q *Quote) CreateApi() uintptr {
	bs, _ := syscall.BytePtrFromString(q.pszFlowPath)
	cUdp := 0
	if q.usingUdp {
		cUdp = 1
	}
	cMulticast := 0
	if q.usingMulticast {
		cMulticast = 1
	}
	api, _, _ := q.h.MustFindProc("qCreateApi").Call(uintptr(unsafe.Pointer(bs)), uintptr(cUdp), uintptr(cMulticast))
	return api
}

func (q *Quote) CreateSpi() uintptr {
	pSpi, _, _ := q.h.MustFindProc("qCreateSpi").Call()
	return pSpi
}

func (q *Quote) GetApiVersion() string {
	ver, _, _ := q.h.MustFindProc("qGetApiVersion").Call()
	p := (*byte)(unsafe.Pointer(ver))
	data := make([]byte, 0)
	for *p != 0 {
		data = append(data, *p)
		ver += unsafe.Sizeof(byte(0))
		p = (*byte)(unsafe.Pointer(ver))
	}

	return string(data)
}

func (q *Quote) GetTradingDay() string {
	ver, _, _ := q.h.MustFindProc("qGetTradingDay").Call(uintptr(q.api))
	p := (*byte)(unsafe.Pointer(ver))
	data := make([]byte, 0)
	for *p != 0 {
		data = append(data, *p)
		ver += unsafe.Sizeof(byte(0))
		p = (*byte)(unsafe.Pointer(ver))
	}

	return string(data)
}
[[ range .Fn]][[ if eq .FuncRtn "void"]]
[[if eq .FuncName "RegisterSpi"]]
func (q *Quote) [[ .FuncName ]]() { [[else]]
func (q *Quote) [[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]][[ if gt $i 1 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]][[end]]) { [[end]]
	q.h.MustFindProc("q[[ .FuncName ]]").Call([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ fldType .FieldType .FieldName ]][[ end ]])
}[[ else ]]
func (q *Quote) [[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]][[ if gt $i 1 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]][[end]]) [[ .FuncRtn ]]32 { [[ range $i,$v := .FuncFields ]][[ supType .FieldType .FieldName ]][[ end ]]
    res, _, _ := q.h.MustFindProc("q[[ .FuncName ]]").Call([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ fldType .FieldType .FieldName ]][[ end ]])
    return [[ .FuncRtn ]]32(res)
}[[ end ]][[end]]
[[ range .On ]]
// [[ .Comment ]]
func (q *Quote) [[ .FuncName ]](fn func([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type]][[ end ]])) {
    q.[[ .FuncName ]]_ = fn
	q.h.MustFindProc("q[[ .FuncName ]]").Call(q.pSpi, syscall.NewCallback(q.[[.FuncName]]__))
}
[[ end ]]
[[ range .On ]]
// [[ .Comment ]]
func (q *Quote) [[ .FuncName ]]__([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type]][[ end ]]) uintptr {
    if q.[[ .FuncName ]]_ != nil {    
	    q.[[ .FuncName ]]_([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ .FieldName|trimStar ]][[ end ]])
    }
	return 0
}
[[ end ]]