package ctpgo

/*
#cgo CFLAGS: -I${SRCDIR}/../
#cgo LDFLAGS: -L${SRCDIR}/../libs/ -lctpquote_api
#include "cctpquote_api_[[if eq .Pf "macos"]]darwin[[else]][[.Pf]][[end]].h"
*/
import "C"
import (
	"os"
	"unsafe"
)

type Quote struct {
	api            unsafe.Pointer
	pSpi           unsafe.Pointer
	version        string
	pszFlowPath    string
	usingUdp       bool
	usingMulticast bool

    [[ range .On]]// [[ .Comment ]]
    [[ .FuncName ]]_ func([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]])
    [[ end]]
}

var q *Quote

func InitQuote(pszFlowPath string, usingUdp bool, usingMulticast bool) *Quote {
	q = new(Quote)
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

func (q *Quote) CreateApi() unsafe.Pointer {
	api := C.qCreateApi(C.CString(q.pszFlowPath), C._Bool(q.usingUdp), C._Bool(q.usingMulticast))
	return api
}

func (q *Quote) CreateSpi() unsafe.Pointer {
	pSpi := C.qCreateSpi()
	return pSpi
}

func (q *Quote) GetApiVersion() string {
    return C.GoString((*C.char)(C.qGetApiVersion()))
}

func (q *Quote) GetTradingDay() string {
    return C.GoString((*C.char)(C.qGetTradingDay(q.api)))
}
[[ range .Fn]][[ if eq .FuncRtn "void"]] [[if eq .FuncName "RegisterSpi"]]
func (q *Quote) [[ .FuncName ]]() { [[else]]
func (q *Quote) [[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]][[ if gt $i 1 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]][[end]]) { [[end]]
    [[ range .FuncFields ]] [[ supType .FieldType .FieldName ]] [[ end ]]
    C.q[[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ fldType .FieldType .FieldName ]][[ end ]]) [[ postSup .FuncFields ]]
}[[ else ]]
func (q *Quote) [[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]][[ if gt $i 1 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]][[end]]) [[ .FuncRtn ]]32 {
    [[ range .FuncFields ]] [[ supType .FieldType .FieldName ]] [[ end ]]
    res := C.q[[ .FuncName ]]([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ fldType .FieldType .FieldName ]][[ end ]])
    [[ postSup .FuncFields ]] return [[ .FuncRtn ]]32(res)
}[[ end ]]
[[ end ]]

[[ range .On]]// [[ .Comment ]]
func (q *Quote) [[ .FuncName ]] (fn func([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[if eq .FieldName "*ppInstrumentID"]][[ .FieldName|trimStar ]] [][]byte [[else]][[ .FieldName|trimStar ]] [[ .FieldType|ctp_type ]][[end]][[ end ]])) {
	q.[[ .FuncName ]]_ = fn
	C.q[[ .FuncName ]](q.pSpi, C.q[[ .FuncName ]]_)
}
[[ end]]

[[ range .On ]]
//export q[[ .FuncName ]]_
func q[[ .FuncName ]]_([[ range $i, $v := .FuncFields ]][[ if gt $i 0]], [[ end ]][[ .FieldName|trimStar ]] [[ .FieldType|C_struct]][[ end ]]) C.int {
    if q.[[ .FuncName ]]_ != nil{
        q.[[ .FuncName ]]_([[ range $i,$v := .FuncFields ]][[ if gt $i 0 ]], [[ end ]][[ ctp_param .FieldType .FieldName ]][[ end ]])
    }
	return 0
}
[[ end ]]