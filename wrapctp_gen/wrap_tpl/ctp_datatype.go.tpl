/*
 * @Author: Wavy Wei 
 * @Date: 2022-12-12 14:21:36 
 * @Last Modified by: Wavy Wei
 * @Last Modified time: 2025-04-06 18:42:33
 */

package ctpgo
import (
	"bytes"

    "golang.org/x/text/encoding/simplifiedchinese"
)

func BytesToGBK(bs []byte) string {
	msg, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bytes.Split(bs, []byte("\x00"))[0])
	return string(msg)
}

func BytesToString(bs []byte) string {
	return string(bytes.Split(bs, []byte("\x00"))[0])
}

type THOST_TE_RESUME_TYPE = int32

const (
	THOST_TERT_RESTART THOST_TE_RESUME_TYPE = 0
	THOST_TERT_RESUME THOST_TE_RESUME_TYPE = 1
	THOST_TERT_QUICK THOST_TE_RESUME_TYPE = 2
	THOST_TERT_NONE THOST_TE_RESUME_TYPE = 3
)

[[ range .St]]// [[ .Comment ]]
type [[ .FuncName ]]  [[ if or (eq .FuncName "TThostFtdcFBTTradeCodeEnumType") (eq .FuncName "TThostFtdcVirementTradeCodeType") ]]string[[ else ]][[ .FuncTypeName|baseType ]] [[ end ]]
[[ range .FuncFields ]]const [[ .FieldType ]] [[ if eq (len .FieldName) 1 ]][[ .FuncName ]][[ end ]] = [[ if eq (len .FieldName) 1 ]]'[[ .FieldName ]]'[[ else ]]"[[ .FieldName ]]"[[ end ]] // [[ .Comment ]]
[[ end ]]
[[ if eq (len .FuncFields) 0  ]] [[ if eq (.FuncTypeName|isString) true ]]
func (s [[ .FuncName ]]) String() string {	return BytesToString(s[:])}
func (s [[ .FuncName ]]) GBString() string { return BytesToGBK(s[:])}
[[ end ]][[ else ]]var mp[[ .FuncName ]] = map[ [[ .FuncName ]] ]string{
[[ range .FuncFields ]][[ .FieldType ]]: "[[ .Comment ]]",
[[ end ]]
}
func (e [[ .FuncName ]]) String() string {
	if s, ok := mp[[ .FuncName ]][e]; ok {
		return s
	}
	return string(e) + "值未定义"
}
[[ end ]]
[[ end ]]
