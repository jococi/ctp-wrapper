package main

import (
	"bytes"
	"errors"
	"flag"
	"fmt"
	"go/format"
	"os"
	"path"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"
	"text/template"
	"unicode/utf8"

	"golang.org/x/text/encoding/simplifiedchinese"
)

var (
	// source path
	srcpath string
	// platform
	csys string
	// output path
	outpath string
	// language
	lang string
)

// Funciton Structure
type fieldStruct struct {
	FieldType string
	FieldName string
	Comment   string
	IsArray   bool // 标记是否为数组参数
}

// Template Structure
type tplStruct struct {
	FuncTypeName         string
	FuncRtn              string
	FuncName             string
	Comment              string
	FuncFields           []fieldStruct
	FuncFieldsWithoutApi []fieldStruct // Windows: 函数签名中不包含 api 字段
}

// Template
func tmpl(tplFileName string, content interface{}, funcMap template.FuncMap, outPath string) {
	_, curFile, _, _ := runtime.Caller(1)
	tplPath := path.Join(filepath.Dir(curFile), "wrap_tpl") // 模板文件在执行文件同级目录下的wrap_tpl文件夹下

	fm := make(template.FuncMap, 0)
	fm["trimStar"] = func(str string) string {
		return strings.TrimPrefix(str, "*")
	}
	// formatParamType: 将类型中的 * 移到变量名前，例如 "CThostFtdcUserLogoutField*" -> "CThostFtdcUserLogoutField"
	fm["formatParamType"] = func(fieldType string) string {
		if strings.HasSuffix(fieldType, "*") {
			return strings.TrimSuffix(fieldType, "*")
		}
		return fieldType
	}
	// formatParamName: 如果类型是指针，在变量名前加 *，例如 "pUserLogout" -> "*pUserLogout"
	fm["formatParamName"] = func(fieldType, fieldName string) string {
		if strings.HasSuffix(fieldType, "*") {
			return "*" + fieldName
		}
		return fieldName
	}

	for k, v := range funcMap {
		fm[k] = v
	}

	t := template.New(path.Base(tplFileName)).Delims("[[", "]]").Funcs(fm)
	t, err := t.ParseFiles(path.Join(tplPath, tplFileName))
	if err != nil {
		panic(err)
	}
	var buf = bytes.Buffer{}
	err = t.Execute(&buf, content) // ***
	if err != nil {
		panic(err)
	}

	// 写入 .h
	fname := strings.TrimSuffix(tplFileName, filepath.Ext(tplFileName))
	csys_flag := csys
	if csys_flag == "macos" {
		csys_flag = "darwin"
	}
	if strings.Contains(fname, "go") {
		if strings.Contains(fname, "win") || strings.Contains(fname, "nix") {
			fslice1 := strings.Split(fname, "_")
			fslice2 := strings.Split(fslice1[2], ".")
			fname = fslice1[0] + "_" + fslice1[1] + "_" + csys_flag + "." + fslice2[1]
		} else {
			fslice1 := strings.Split(fname, "_")
			fslice2 := strings.Split(fslice1[1], ".")
			fname = fslice1[0] + "_" + fslice2[0] + "_" + csys_flag + "." + fslice2[1]
		}
	}
	// 格式化生成的代码
	if strings.HasSuffix(fname, ".go") {
		formatted, err := format.Source(buf.Bytes())
		if err != nil {
			fmt.Fprintf(os.Stderr, "格式化错误: %v\n", err)
			panic(err)
		}
		err = os.WriteFile(path.Join(outPath, fname),
			formatted,
			os.ModePerm)
		if err != nil {
			panic(err)
		}
	} else {
		err = os.WriteFile(path.Join(outPath, fname),
			buf.Bytes(),
			os.ModePerm)
		if err != nil {
			panic(err)
		}
	}

}

func gen_cwrap(tplExeFunc func(title string, on []*tplStruct, fn []*tplStruct)) {
	srcp1 := path.Join(srcpath, "ThostFtdcMdApi.h")
	srcp2 := path.Join(srcpath, "ThostFtdcTraderApi.h")
	for _, hFileName := range []string{srcp1, srcp2} {
		var title string
		if strings.Contains(hFileName, "ThostFtdcMdApi") {
			title = "ctpquote_api"
		} else {
			title = "ctptrade_api"
		}
		bsFile, err := os.ReadFile(hFileName)
		if err != nil {
			panic(err)
		}
		// 汉字处理：检测编码，如果是 UTF-8 则直接使用，否则按 GB18030 解码
		var strFile string
		if utf8.Valid(bsFile) {
			// 文件已经是 UTF-8 编码，直接使用
			strFile = string(bsFile)
		} else {
			// 文件是 GB18030 编码，需要解码
			decoded, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
			strFile = string(decoded)
		}
		strFile = strings.ReplaceAll(strFile, "\r\n", "\n") // 换行符用 \n 避免 win和 lnx执行时不一致
		/*
			///登录请求响应
			virtual void OnRspUserLogin(CThostFtdcRspUserLoginField *pRspUserLogin, CThostFtdcRspInfoField *pRspInfo, int nRequestID, bool bIsLast) {};
		*/
		re := regexp.MustCompile(`\t///(.*)\n[^v]*virtual\s+(\w+)\s+(\w+)\(([^)]*)\)`) // 分解函数定义:注释,返回类型,函数名,参数字段四部分
		funs := re.FindAllStringSubmatch(strFile, -1)
		tplsOn := make([]*tplStruct, 0)
		tplsFn := make([]*tplStruct, 0)
		for _, fun := range funs {
			funComment, funRtn, funName, funParams := fun[1], fun[2], fun[3], fun[4]
			// 参数分解:类型,名称
			// 支持格式: int nRequestID, char* pszFlowPath, CThostFtdcFrontInfoField* pFrontInfo, CThostFtdcRspUserLoginField *pRspUserLogin
			// 类型名可能包含星号（紧跟在类型名后或中间有空格），参数名前面可能有星号但提取时不包含
			// 先按逗号分割参数，然后逐个解析
			paramList := strings.Split(funParams, ",")
			funFields := make([]fieldStruct, 0)
			for _, param := range paramList {
				param = strings.TrimSpace(param)
				if param == "" {
					continue
				}
				// 从后往前找参数名（最后一个标识符），前面的是类型
				// 先找到参数名（最后一个单词，可能后面跟[]）
				reName := regexp.MustCompile(`([A-Za-z_][A-Za-z0-9_]*)\s*(\[\])?\s*$`)
				nameMatch := reName.FindStringSubmatch(param)
				if len(nameMatch) < 2 {
					continue
				}
				fieldName := nameMatch[1]
				hasArray := nameMatch[2] == "[]"
				// 去掉参数名和[]，剩下的就是类型（可能包含*）
				typePart := strings.TrimSpace(param[:len(param)-len(fieldName)])
				if hasArray {
					typePart = strings.TrimSpace(typePart[:len(typePart)-2])
				}
				// 检查类型部分末尾是否有 *（可能有空格）
				fieldType := strings.TrimSpace(typePart)
				// 如果类型名和*之间有空格（如 "CThostFtdcRspUserLoginField *"），需要合并
				if strings.HasSuffix(fieldType, " *") {
					fieldType = strings.TrimSuffix(fieldType, " *") + "*"
				} else if strings.HasSuffix(fieldType, "* ") {
					fieldType = strings.TrimSuffix(fieldType, "* ") + "*"
				} else if !strings.HasSuffix(fieldType, "*") && (strings.Contains(param, " * "+fieldName) || strings.Contains(param, "* "+fieldName)) {
					// 类型名本身没有*，但参数前有*（有空格）
					fieldType = fieldType + "*"
				}
				// 注意：数组标记不添加到类型中，而是单独标记，在模板中会放在参数名后面
				funFields = append(funFields, fieldStruct{FieldType: fieldType, FieldName: fieldName, IsArray: hasArray})
			}
			if strings.HasPrefix(funName, "On") { // On 响应函数
				tplsOn = append(tplsOn, &tplStruct{
					Comment:      funComment,
					FuncTypeName: "FP_" + funName,
					FuncRtn:      funRtn,
					FuncName:     funName,
					FuncFields:   funFields,
				})
			} else {
				tplsFn = append(tplsFn, &tplStruct{
					Comment:      funComment,
					FuncTypeName: "FP_" + funName,
					FuncRtn:      funRtn,
					FuncName:     funName,
					FuncFields:   funFields,
				})
			}
		}
		tplExeFunc(title, tplsOn, tplsFn)
	}
}

// Generate datatype
func gen_datatype(srcpath string, fn func([]*tplStruct)) {
	srcpath = path.Join(srcpath, "ThostFtdcUserApiDataType.h")
	bsFile, err := os.ReadFile(srcpath)
	if err != nil {
		panic(err)
	}
	// 检测编码，如果是 UTF-8 则直接使用，否则按 GB18030 解码
	var strFile string
	if utf8.Valid(bsFile) {
		strFile = string(bsFile)
	} else {
		decoded, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
		strFile = string(decoded)
	}
	strFile = strings.ReplaceAll(strings.ReplaceAll(strFile, "\r\n", "\n"), "\n\t", "\n")

	/*
		/////////////////////////////////////////////////////////////////////////
		///TFtdcExchangePropertyType是一个交易所属性类型
		/////////////////////////////////////////////////////////////////////////
		///正常
		#define THOST_FTDC_EXP_Normal '0'
		///根据成交生成报单
		#define THOST_FTDC_EXP_GenOrderByTrade '1'

		typedef char TThostFtdcExchangePropertyType;

		=>
		// 交易所属性类型
		type TThostFtdcExchangePropertyType byte
		const THOST_FTDC_EXP_Normal = '0' // 正常
		const THOST_FTDC_EXP_GenOrderByTrade = '1' // 根据成交生成报单
	*/
	re := regexp.MustCompile(`/+.+是一个(.*)\n/*\n([^;]+)typedef\s+(\w+)\s+(\w+)\s*;`) // 注释,defines,类型,名称
	types := re.FindAllStringSubmatch(strFile, -1)

	tss := make([]*tplStruct, 0)
	for _, v := range types {
		ts := &tplStruct{
			Comment:      v[1], // 注释
			FuncTypeName: v[3], // 基础类型
			FuncName:     v[4],
		}
		reSub := regexp.MustCompile(`/+(.*)\n#define\s+(\w+)\s+'(.+)'`) // 注释,名称,值 \w改为.因为有'#'的情况
		defines := reSub.FindAllStringSubmatch(v[2], -1)
		for _, v := range defines {
			ts.FuncFields = append(ts.FuncFields, fieldStruct{
				Comment:   v[1],
				FieldType: v[2],
				FieldName: v[3],
			})
		}
		tss = append(tss, ts)
	}

	/*	/////////////////////////////////////////////////////////////////////////
		///TFtdcTraderIDType是一个交易所交易员代码类型
		/////////////////////////////////////////////////////////////////////////
		typedef char TThostFtdcTraderIDType[21];
	*/
	re = regexp.MustCompile(`/+.+是一个(.*)\n/*\ntypedef\s+(\w+)\s+(.+)\s*;`)
	types = re.FindAllStringSubmatch(strFile, -1)
	for _, v := range types {
		ts := &tplStruct{
			Comment:      v[1],
			FuncTypeName: v[2], // 基础类型
			FuncName:     v[3],
		}
		// typedef char TThostFtdcTraderIDType[21]; -> type TThostFtdcTraderIDType [21]byte
		if strings.Contains(ts.FuncName, "[") {
			ts.FuncTypeName = "[" + strings.Split(ts.FuncName, "[")[1] + "byte"
			ts.FuncName = strings.Split(ts.FuncName, "[")[0]
		}
		tss = append(tss, ts)
	}
	fn(tss)
}

// Generate struct
func gen_struct(srcpath string, fn func([]*tplStruct)) {
	srcpath = path.Join(srcpath, "ThostFtdcUserApiStruct.h")
	bsFile, err := os.ReadFile(srcpath)
	if err != nil {
		panic(err)
	}
	// 检测编码，如果是 UTF-8 则直接使用，否则按 GB18030 解码
	var strFile string
	if utf8.Valid(bsFile) {
		strFile = string(bsFile)
	} else {
		decoded, _ := simplifiedchinese.GB18030.NewDecoder().Bytes(bsFile)
		strFile = string(decoded)
	}

	re := regexp.MustCompile(`///(\S*)\s*struct\s*(\w*)\s*{([^}]*)}`) // 分成struct的注释,名称,字段两部分
	structs := re.FindAllStringSubmatch(strFile, -1)
	tss := make([]*tplStruct, 0)
	for _, strc := range structs {
		ts := &tplStruct{
			Comment:      strc[1],
			FuncTypeName: strc[2],
		}
		re = regexp.MustCompile(`///([^\r\n]*)\s*(\w+)\s+([^;]+);`) // 所有字段再分解成各个单独字段: 注释(可能含空格),类型,名称
		fields := re.FindAllStringSubmatch(strc[3], -1)
		for _, v := range fields {
			ts.FuncFields = append(ts.FuncFields, fieldStruct{
				Comment:   v[1],
				FieldType: v[2],
				FieldName: v[3],
			})
		}
		tss = append(tss, ts)
	}
	fn(tss)
}

func cfm(title string, tplsOn, tplsFn []*tplStruct) {
	mpCpp := make(map[string]interface{})
	mpCpp["On"] = tplsOn
	mpCpp["Fn"] = tplsFn
	mpCpp["Pf"] = csys
	tmpl(title+".h.tpl", mpCpp, nil, outpath)
	tmpl(title+".cpp.tpl", mpCpp, nil, outpath)
}

func ccfm(title string, tplsOn, tplsFn []*tplStruct) {
	mpCpp := make(map[string]interface{})
	mpCpp["On"] = tplsOn
	mpCpp["Fn"] = tplsFn
	mpCpp["Pf"] = csys
	fm := make(template.FuncMap)
	fm["struct_Type"] = func(structType string) string {
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if baseType == "CThostFtdcMdSpi" {
			return "void"
		}
		if baseType == "CThostFtdcTraderSpi" {
			return "void"
		}
		if strings.HasSuffix(baseType, "Field") || strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return "struct " + baseType // struct CThostFtdcRspUserLoginField
		}
		if baseType == "bool" {
			return "bool" // 保持为 bool，不要改成 _Bool
		}
		if baseType == "THOST_TE_RESUME_TYPE" {
			return "int"
		}
		return baseType
	}
	// 添加字符串处理函数
	fm["hasPrefix"] = strings.HasPrefix
	fm["hasSuffix"] = strings.HasSuffix
	fm["trimPrefix"] = strings.TrimPrefix
	// 添加逻辑函数
	fm["or"] = func(a, b bool) bool { return a || b }
	tmpl("c"+title+".h.go.tpl", mpCpp, fm, outpath)
}

func pycfm(title string, on, fn []*tplStruct) {
	funcs := make(map[string]interface{})
	funcs["On"] = on
	funcs["Fn"] = fn
	funcs["Pf"] = csys
	fm := make(template.FuncMap)
	fm["baseType"] = func(preType string) string {
		if preType == "CThostFtdcMdSpi" {
			return "c_void_p"
		}
		if preType == "CThostFtdcTraderSpi" {
			return "c_void_p"
		}
		if preType == "int" { // SubscribeMarketData(char *ppInstrumentID[], int nCount)
			return "c_int32"
		}
		return "c_void_p" // char*  CThost结构体
	}
	fm["fnBaseType"] = func(fieldTypeName string) string {
		// 移除类型中的 * 后缀（Python 类型注解不需要 *）
		baseType := strings.TrimSuffix(fieldTypeName, "*")
		baseType = strings.TrimSpace(baseType)
		if baseType == "char" {
			// 对于地址参数（pszFrontAddress, pszNsAddress），类型注解使用 str
			return "str"
		}
		if baseType == "CThostFtdcMdSpi" {
			return "c_void_p"
		}
		if baseType == "CThostFtdcTraderSpi" {
			return "c_void_p"
		}
		return baseType
	}
	fm["evBaseType"] = func(fieldTypeName string) string {
		// 移除类型中的 * 后缀
		baseType := strings.TrimSuffix(fieldTypeName, "*")
		baseType = strings.TrimSpace(baseType)
		if strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return fmt.Sprintf("POINTER(%s)", baseType)
		}
		if baseType == "int" {
			return "c_int32"
		}
		if baseType == "bool" {
			return "c_bool"
		}
		return baseType
	}
	fm["param"] = func(fieldType, fieldName string) string {
		if fieldName == "ppInstrumentID" { // 类型为 Array[c_char_p]
			return fieldName
		}
		if fieldName == "pSpi" {
			return fmt.Sprintf("self.%s", fieldName)
		}
		// 移除类型中的 * 后缀
		baseType := strings.TrimSuffix(fieldType, "*")
		baseType = strings.TrimSpace(baseType)
		if baseType == "char" {
			// 使用 bytes(fieldName, encoding='ascii') 将 str 转换为 bytes
			return fmt.Sprintf("bytes(%s, encoding='ascii')", fieldName)
		}
		if baseType == "CThostFtdcMdSpi" {
			return fieldName
		}
		if baseType == "CThostFtdcTraderSpi" {
			return fieldName
		}
		if strings.HasPrefix(baseType, "CThostFtdc") {
			return fmt.Sprintf("byref(%s)", fieldName)
		}
		return fieldName
	}
	fm["onParam"] = func(fieldType, fieldName string) string {
		// 移除类型中的 * 后缀
		baseType := strings.TrimSuffix(fieldType, "*")
		baseType = strings.TrimSpace(baseType)
		if strings.HasPrefix(baseType, "CThostFtdc") {
			// self.OnRspQryInvestorPosition(copy.deepcopy(POINTER(CThostFtdcInvestorPositionField).from_param(pInvestorPosition).contents) if pInvestorPosition else CThostFtdcInvestorPositionField(), copy.deepcopy(POINTER(CThostFtdcRspInfoField).from_param(pRspInfo).contents), nRequestID, bIsLast)
			return fmt.Sprintf("copy.deepcopy(POINTER(%s).from_param(%s).contents) if %s else %s()", baseType, fieldName, fieldName, baseType)
		}
		return fieldName
	}
	tmpl(title+".py.tpl", funcs, fm, outpath)
}

func wgocfm(title string, tplsOn, tplsFn []*tplStruct) {
	// for _, v := range [][]*tplStruct{tplsOn, tplsFn} {
	for _, fn := range tplsFn { // 主调函数
		// 增加 void* api 首个参数
		tmp := []fieldStruct{{FieldType: "void*", FieldName: "api"}}
		tmp = append(tmp, fn.FuncFields...)
		fn.FuncFields = tmp
		// Windows 版本：从函数签名中移除 api 参数（但 C 函数调用时仍然需要）
		// macOS 上的 ReqUserLogin 需要特殊处理：从函数签名中移除 length 和 systemInfo 参数
		// 但这些参数仍然会传递给 C 函数（在模板中处理）
		if csys == "macos" && fn.FuncName == "ReqUserLogin" {
			filtered := make([]fieldStruct, 0)
			for _, f := range fn.FuncFields {
				// 检查字段名，可能是 "length" 或 "*length" 等格式
				fieldName := strings.TrimPrefix(f.FieldName, "*")
				// 跳过 length 和 systemInfo 参数（这些会在函数内部自动生成）
				if fieldName != "length" && fieldName != "systemInfo" {
					filtered = append(filtered, f)
				}
			}
			fn.FuncFields = filtered
		}
		// 设置 FuncFieldsWithoutApi：Windows 版本不包含 api 字段
		if csys == "windows" {
			// 对于 Windows，创建一个不包含 api 的字段列表用于函数签名
			fn.FuncFieldsWithoutApi = make([]fieldStruct, 0)
			for _, f := range fn.FuncFields {
				fieldName := strings.TrimPrefix(f.FieldName, "*")
				if fieldName != "api" {
					fn.FuncFieldsWithoutApi = append(fn.FuncFieldsWithoutApi, f)
				}
			}
		} else {
			fn.FuncFieldsWithoutApi = fn.FuncFields
		}
	}
	// }
	mpCpp := make(map[string]interface{})
	mpCpp["On"] = tplsOn
	mpCpp["Fn"] = tplsFn
	mpCpp["Pf"] = csys
	fm := make(template.FuncMap)
	fm["ctp_type"] = func(structType string) string {
		// 先去掉类型后面的 *（如果有）
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if strings.HasSuffix(baseType, "Field") { // struct
			return fmt.Sprintf("*%s", baseType) // *CThostFtdcUserLogoutField
		}
		if baseType == "void" {
			return "uintptr"
		}
		if baseType == "char" {
			return "[]byte"
		}
		if baseType == "CThostFtdcMdSpi" {
			return "uintptr"
		}
		if baseType == "CThostFtdcTraderSpi" {
			return "uintptr"
		}
		if baseType == "THOST_TE_RESUME_TYPE" {
			return "THOST_TE_RESUME_TYPE"
		}
		return baseType
	}
	fm["fldType"] = func(structType string, str string) string {
		if strings.Contains(str, "*") {
			if str == "*pSpi" || str == "pSpi" {
				if title == "ctpquote_api" {
					fieldName := strings.TrimPrefix(str, "*")
					// 对于 Windows，pSpi 已经是 uintptr 类型，直接使用 uintptr 转换
					return fmt.Sprintf("uintptr(q.%s)", fieldName)
				} else {
					fieldName := strings.TrimPrefix(str, "*")
					// 对于 Windows，pSpi 已经是 uintptr 类型，直接使用 uintptr 转换
					return fmt.Sprintf("uintptr(t.%s)", fieldName)
				}
			} else {
				// Windows: 对于 char *ppInstrumentID[]，改用预先准备的 _ppPtr
				// 检查字段名是否为 ppInstrumentID（可能带 * 前缀）
				fieldName := strings.TrimPrefix(str, "*")
				if fieldName == "ppInstrumentID" {
					return "_ppPtr"
				}
				// 去掉类型中的 * 和 [] 后缀
				baseType := strings.TrimSuffix(structType, "*")
				baseType = strings.TrimSuffix(baseType, "[]")
				baseType = strings.TrimSpace(baseType)
				if baseType == "char" {
					// 对于 char* 类型的参数（[]byte），使用第一个元素的地址
					return fmt.Sprintf("uintptr(unsafe.Pointer(&%s[0]))", strings.TrimPrefix(str, "*"))
				} else {
					return fmt.Sprintf("uintptr(unsafe.Pointer(%s))", strings.TrimPrefix(str, "*"))
				}
			}
		} else {
			if str == "api" {
				if title == "ctpquote_api" {
					return fmt.Sprintf("uintptr(q.%s)", str)
				} else {
					return fmt.Sprintf("uintptr(t.%s)", str)
				}
			} else if str == "pSpi" {
				// 处理 pSpi 参数（不带 * 的情况）
				// 对于 Windows，pSpi 已经是 uintptr 类型，直接使用 uintptr 转换
				if title == "ctpquote_api" {
					return fmt.Sprintf("uintptr(q.%s)", str)
				} else {
					return fmt.Sprintf("uintptr(t.%s)", str)
				}
			} else {
				// Windows: 对于 ppInstrumentID（不带 *），使用预先准备的 _ppPtr
				// 检查字段名，去除可能的 * 前缀后比较
				fieldName := strings.TrimPrefix(str, "*")
				if fieldName == "ppInstrumentID" {
					return "_ppPtr"
				}
				// 检查是否是结构体指针类型（Field 结尾或 CThostFtdc 开头）
				baseType := strings.TrimSuffix(structType, "*")
				// 去掉 [] 后缀（数组类型）
				baseType = strings.TrimSuffix(baseType, "[]")
				baseType = strings.TrimSpace(baseType)
				// 检查是否是 char 类型的数组参数（[]byte），但要排除 ppInstrumentID
				if baseType == "char" && fieldName != "ppInstrumentID" {
					// 对于 char 类型的数组参数（[]byte），使用第一个元素的地址
					return fmt.Sprintf("uintptr(unsafe.Pointer(&%s[0]))", str)
				}
				if strings.HasSuffix(baseType, "Field") || strings.HasPrefix(baseType, "CThostFtdc") {
					// 结构体指针类型，使用 unsafe.Pointer
					return fmt.Sprintf("uintptr(unsafe.Pointer(%s))", str)
				}
				return fmt.Sprintf("uintptr(%s)", str)
			}
		}
	}
	// 为 Windows 版生成 ppInstrumentID 的指针数组与 KeepAlive 代码
	fm["supType"] = func(structType string, field string) string {
		if field == "*ppInstrumentID" || field == "ppInstrumentID" {
			return fmt.Sprintf(`
	var _ppPtr uintptr
	if nCount > 0 {
		ptrs := make([]*byte, nCount)
		for i := 0; i < nCount; i++ {
			if len(ppInstrumentID[i]) > 0 {
				ptrs[i] = &ppInstrumentID[i][0]
			} else {
				ptrs[i] = nil
			}
		}
		_ppPtr = uintptr(unsafe.Pointer(&ptrs[0]))
		runtime.KeepAlive(ppInstrumentID)
		runtime.KeepAlive(ptrs)
	}
	`)
		}
		return fmt.Sprintf("")
	}
	tmpl(title+"_win.go.tpl", mpCpp, fm, outpath)
}

func xgocfm(title string, tplsOn, tplsFn []*tplStruct) {
	// 注意：不需要在这里添加 void* api 参数
	// 因为模板中已经硬编码了 void *api
	// 如果添加会导致重复：void *api, void* api

	// macOS 上的 ReqUserLogin 需要特殊处理：从函数签名中移除 length 和 systemInfo 参数
	// 但这些参数仍然会传递给 C 函数（在模板中处理）
	if csys == "macos" {
		for _, fn := range tplsFn {
			if fn.FuncName == "ReqUserLogin" {
				filtered := make([]fieldStruct, 0)
				for _, f := range fn.FuncFields {
					// 检查字段名，可能是 "length" 或 "*length" 等格式
					fieldName := strings.TrimPrefix(f.FieldName, "*")
					// 跳过 length 和 systemInfo 参数（这些会在函数内部自动生成）
					if fieldName != "length" && fieldName != "systemInfo" {
						filtered = append(filtered, f)
					}
				}
				fn.FuncFields = filtered
			}
		}
	}

	funcs := make(map[string]interface{})
	funcs["On"] = tplsOn
	funcs["Fn"] = tplsFn
	funcs["Pf"] = csys
	fm := make(template.FuncMap)
	// 添加字符串处理函数到模板
	fm["hasPrefix"] = strings.HasPrefix
	fm["hasSuffix"] = strings.HasSuffix
	fm["trimPrefix"] = strings.TrimPrefix
	// 添加逻辑函数
	fm["or"] = func(a, b bool) bool { return a || b }
	fm["and"] = func(a, b bool) bool { return a && b }
	fm["ne"] = func(a, b interface{}) bool { return a != b }
	fm["eq"] = func(a, b interface{}) bool { return a == b }
	fm["struct_Type"] = func(structType string) string {
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if baseType == "CThostFtdcMdSpi" {
			return "void"
		}
		if baseType == "CThostFtdcTraderSpi" {
			return "void"
		}
		if strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return "struct " + baseType // struct CThostFtdcRspUserLoginField
		}
		if baseType == "bool" {
			return "bool" // 在 C 头文件中使用 bool 而不是 _Bool
		}
		if baseType == "THOST_TE_RESUME_TYPE" {
			return "int"
		}
		return baseType
	}
	fm["C_struct"] = func(structType string) string {
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return "*C.struct_" + baseType // field *C.struct_CThostFtdcRspUserLoginField
		}
		if baseType == "int" {
			return "C.int"
		}
		if baseType == "bool" {
			return "C._Bool"
		}
		return baseType
	}
	fm["ctp_type"] = func(structType string) string {
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return fmt.Sprintf("*%s", baseType) // *CThostFtdcUserLogoutField
		}
		if baseType == "char" {
			return "[]byte"
		}
		return baseType
	}
	fm["ctp_param"] = func(structType, field string) string {
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if strings.HasPrefix(baseType, "CThostFtdc") { // struct
			return fmt.Sprintf("(*%s)(unsafe.Pointer(%s))", baseType, strings.TrimPrefix(field, "*")) // (*CThostFtdcRspUserLoginField)(unsafe.Pointer(field))
		}
		if baseType == "int" {
			return "int(" + field + ")"
		}
		if baseType == "bool" {
			return "bool(" + field + ")"
		}
		return field
	}
	fm["fldType"] = func(structType string, field string) string {
		if field == "*ppInstrumentID" || field == "ppInstrumentID" {
			return "_ppPtr"
		}
		// 先检查 pSpi 字段，因为它已经是 unsafe.Pointer 类型
		if field == "*pSpi" || field == "pSpi" {
			if title == "ctpquote_api" {
				return fmt.Sprintf("q.%s", strings.TrimPrefix(field, "*"))
			} else {
				return fmt.Sprintf("t.%s", strings.TrimPrefix(field, "*"))
			}
		}
		// 先去掉类型末尾的 *，因为 * 应该在变量名前
		baseType := strings.TrimSuffix(structType, "*")
		baseType = strings.TrimSpace(baseType)

		if strings.HasSuffix(baseType, "Field") || strings.HasPrefix(baseType, "CThostFtdc") {
			return fmt.Sprintf("(*C.struct_%s)(unsafe.Pointer(%s))", baseType, strings.TrimPrefix(field, "*"))
		}
		if baseType == "char" {
			return fmt.Sprintf("(*C.char)(unsafe.Pointer(C.CBytes(%s)))", strings.TrimPrefix(field, "*"))
		}
		if baseType == "int" {
			return fmt.Sprintf("C.int(%s)", field)
		}
		if baseType == "bool" {
			return fmt.Sprintf("C._Bool(%s)", field)
		}
		if baseType == "THOST_TE_RESUME_TYPE" {
			return fmt.Sprintf("C.int(%s)", field)
		}
		if field == "api" {
			if title == "ctpquote_api" {
				return "q.api"
			} else {
				return "t.api"
			}
		}
		if baseType == "TThostFtdcSystemInfoLenType" {
			return fmt.Sprintf("C.int(%s)", field)
		}
		if baseType == "TThostFtdcClientSystemInfoType" {
			return fmt.Sprintf("(*C.char)(unsafe.Pointer(&%s[0]))", field)
		}
		// 处理 void* 类型
		if baseType == "void" {
			return fmt.Sprintf("unsafe.Pointer(%s)", strings.TrimPrefix(field, "*"))
		}
		// 对于其他未知类型，直接返回字段名（去掉可能的 * 前缀）
		return strings.TrimPrefix(field, "*")
	}

	fm["supType"] = func(structType string, field string) string {
		if field == "*ppInstrumentID" || field == "ppInstrumentID" {
			return fmt.Sprintf(`
    tmp_arr := make([]*C.char, nCount)
    for i := 0; i < nCount; i++ {
        tmp_arr[i] = C.CString(string(ppInstrumentID[i]))
    }
    var _ppPtr **C.char
    if nCount > 0 {
        _ppPtr = (**C.char)(unsafe.Pointer(&tmp_arr[0]))
    }
    `)
		}
		return fmt.Sprintf("")
	}
	fm["postSup"] = func(fields []fieldStruct) string {
		for _, f := range fields {
			if f.FieldName == "*ppInstrumentID" || f.FieldName == "ppInstrumentID" {
				return "\tfor i := 0; i < nCount; i++ {\n\t\tif tmp_arr[i] != nil { C.free(unsafe.Pointer(tmp_arr[i])) }\n\t}\n"
			}
		}
		return ""
	}

	// xgocfm 只生成 Go 文件（ctpquote_api_darwin.go 或 ctptrade_api_darwin.go）
	// C 头文件由 ccfm 函数生成
	tmpl(title+"_nix.go.tpl", funcs, fm, outpath)
}

func pydfm(ts []*tplStruct) {
	fm := make(template.FuncMap)
	fm["baseType"] = func(preType string) string {
		if preType == "int" { // typedef int TThostFtdcIPPortType; -> TThostFtdcIPPortType = c_int32
			return "c_int32"
		}
		if preType == "double" { // typedef double TThostFtdcPriceType; -> TThostFtdcPriceType = c_double
			return "c_double"
		}
		if preType == "short" { // typedef short TThostFtdcSequenceSeriesType; -> TThostFtdcSequenceSeriesType = c_short
			return "c_short"
		}
		if preType == "char" { // typedef char TThostFtdcNewsUrgencyType;
			return "c_char"
		}
		// [nn]byte
		if strings.Contains(preType, "[") { // typedef char TThostFtdcTraderIDType[21]; -> TThostFtdcTraderIDType = c_char*21
			return "c_char*" + strings.Split(strings.Split(preType, "[")[1], "]")[0]
		}
		return preType
	}
	fm["atoi"] = func(char string) int {
		return int(char[0])
	}
	tmpl("ctp_datatype.py.tpl", ts, fm, outpath)
}

func godfm(ts []*tplStruct) {
	mpCpp := make(map[string]interface{})
	mpCpp["St"] = ts
	mpCpp["Pf"] = csys
	fm := make(template.FuncMap)
	fm["baseType"] = func(preType string) string {
		if preType == "int" { // typedef int TThostFtdcIPPortType; -> TThostFtdcIPPortType = c_int32
			return "int32"
		}
		if preType == "double" { // typedef double TThostFtdcPriceType; -> TThostFtdcPriceType = c_double
			return "float64"
		}
		if preType == "short" { // typedef short TThostFtdcSequenceSeriesType; -> TThostFtdcSequenceSeriesType = c_short
			return "int16"
		}
		if preType == "char" { // typedef char TThostFtdcNewsUrgencyType;
			return "byte"
		}
		return preType
	}

	tmpl("ctp_datatype.go.tpl", mpCpp, fm, outpath)
}

func pysfm(ts []*tplStruct) {
	mpPlatform := make(map[string]interface{})
	mpPlatform["St"] = ts
	mpPlatform["Pf"] = csys // 添加平台参数
	fm := make(template.FuncMap)
	fm["baseType"] = func(fieldType string) string {
		return fieldType
	}
	tmpl("ctp_struct.py.tpl", mpPlatform, fm, outpath)
}

func gosfm(ts []*tplStruct) {
	mpCpp := make(map[string]interface{})
	mpCpp["St"] = ts
	mpCpp["Pf"] = csys
	tmpl("ctp_struct.go.tpl", mpCpp, nil, outpath)
}

// Judge Platform type
func iscontain(item string, item_arr []string) error {
	for _, cit := range item_arr {
		if item == cit {
			return nil
		}
	}
	return errors.New("must input: linux or windows or macos")
}

func main() {

	flag.StringVar(&srcpath, "srcpath", ".",
		"Base path of CTP API")

	flag.StringVar(&csys, "csys", "macos",
		"Generate Wrapper for System, support: linux,windows,macos")

	flag.StringVar(&outpath, "outpath", ".",
		"Output path")

	flag.StringVar(&lang, "lang", "python",
		"Language Wrapper, support: python,golang")

	flag.Parse()

	var sys_arr = []string{"linux", "windows", "macos"}
	err := iscontain(csys, sys_arr)
	if err != nil {
		panic(err)
	}
	srcpath = path.Join(srcpath, csys)
	// 检查输出路径是否存在，如果不存在则创建
	if _, err := os.Stat(outpath); os.IsNotExist(err) {
		err = os.MkdirAll(outpath, os.ModePerm)
		if err != nil {
			panic(fmt.Errorf("创建输出目录失败: %v", err))
		}
		fmt.Printf("已创建输出目录: %s\n", outpath)
	}

	if lang == "c" {
		gen_cwrap(cfm)
	} else if lang == "python" {
		gen_datatype(srcpath, pydfm)
		gen_struct(srcpath, pysfm)
		gen_cwrap(pycfm)
	} else if lang == "golang" {
		gen_datatype(srcpath, godfm)
		gen_struct(srcpath, gosfm)
		if csys == "windows" {
			gen_cwrap(wgocfm)
		} else {
			gen_cwrap(ccfm)
			gen_cwrap(xgocfm)
		}
	}

}
