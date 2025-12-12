"""
CTP 数据类型定义

此文件由代码生成器自动生成，请勿手动修改
CTP 数据类型定义 - 来自 ThostFtdcUserApiDataType.h
"""

import ctypes
from enum import IntEnum

# ========== 枚举类型 ==========

# THOST_TE_RESUME_TYPE 订阅类型
class THOST_TE_RESUME_TYPE(IntEnum):
    """订阅类型"""
    THOST_TERT_RESTART = 0  # 从本交易日开始重传
    THOST_TERT_RESUME = 1  # 从上次收到的续传
    THOST_TERT_QUICK = 2  # 只传送登录后的流内容
    THOST_TERT_NONE = 3  # 不传送

# ========== 类型定义 ==========

# ----- 字符串类型 -----

# TThostFtdcAMLAccountTypeType 账户类型
TThostFtdcAMLAccountTypeType = ctypes.c_char * 5

# TThostFtdcAMLCapitalIOType 资金收付标识类型
TThostFtdcAMLCapitalIOType = ctypes.c_char * 3

# TThostFtdcAMLCapitalPurposeType 资金用途类型
TThostFtdcAMLCapitalPurposeType = ctypes.c_char * 129

# TThostFtdcAMLCustomerCardTypeType 客户身份证件/证明文件类型
TThostFtdcAMLCustomerCardTypeType = ctypes.c_char * 81

# TThostFtdcAMLDistrictIDType 金融机构网点所在地区行政区划代码类型
TThostFtdcAMLDistrictIDType = ctypes.c_char * 7

# TThostFtdcAMLFileNameType AML文件名类型
TThostFtdcAMLFileNameType = ctypes.c_char * 257

# TThostFtdcAMLIdCardTypeType 证件类型
TThostFtdcAMLIdCardTypeType = ctypes.c_char * 3

# TThostFtdcAMLInstitutionIDType 金融机构网点代码类型
TThostFtdcAMLInstitutionIDType = ctypes.c_char * 13

# TThostFtdcAMLInstitutionNameType 金融机构网点名称类型
TThostFtdcAMLInstitutionNameType = ctypes.c_char * 65

# TThostFtdcAMLInstitutionTypeType 金融机构网点代码类型
TThostFtdcAMLInstitutionTypeType = ctypes.c_char * 3

# TThostFtdcAMLInvestorTypeType 投资者类型
TThostFtdcAMLInvestorTypeType = ctypes.c_char * 3

# TThostFtdcAMLParamIDType 参数代码类型
TThostFtdcAMLParamIDType = ctypes.c_char * 21

# TThostFtdcAMLRelationShipType 金融机构网点与大额交易的关系类型
TThostFtdcAMLRelationShipType = ctypes.c_char * 3

# TThostFtdcAMLReportNameType 报文名称类型
TThostFtdcAMLReportNameType = ctypes.c_char * 81

# TThostFtdcAMLReportTypeType 报文类型
TThostFtdcAMLReportTypeType = ctypes.c_char * 2

# TThostFtdcAMLSeqCodeType 业务标识号类型
TThostFtdcAMLSeqCodeType = ctypes.c_char * 65

# TThostFtdcAMLSerialNoType 编号类型
TThostFtdcAMLSerialNoType = ctypes.c_char * 5

# TThostFtdcAMLSiteType 交易地点类型
TThostFtdcAMLSiteType = ctypes.c_char * 10

# TThostFtdcAMLStatusType 状态类型
TThostFtdcAMLStatusType = ctypes.c_char * 2

# TThostFtdcAMLTradeDirectType 资金进出方向类型
TThostFtdcAMLTradeDirectType = ctypes.c_char * 3

# TThostFtdcAMLTradeModelType 资金进出方式类型
TThostFtdcAMLTradeModelType = ctypes.c_char * 3

# TThostFtdcAMLTradingTypeType 交易方式类型
TThostFtdcAMLTradingTypeType = ctypes.c_char * 7

# TThostFtdcAMLTransactClassType 涉外收支交易分类与代码类型
TThostFtdcAMLTransactClassType = ctypes.c_char * 7

# TThostFtdcAbstractType 消息摘要类型
TThostFtdcAbstractType = ctypes.c_char * 81

# TThostFtdcAccountIDType 投资者帐号类型
TThostFtdcAccountIDType = ctypes.c_char * 13

# TThostFtdcAddInfoType 附加信息类型
TThostFtdcAddInfoType = ctypes.c_char * 129

# TThostFtdcAdditionalInfoType 系统外部信息类型
TThostFtdcAdditionalInfoType = ctypes.c_char * 261

# TThostFtdcAddressType 通讯地址类型
TThostFtdcAddressType = ctypes.c_char * 101

# TThostFtdcAdvanceMonthArrayType 月份提前数组类型
TThostFtdcAdvanceMonthArrayType = ctypes.c_char * 13

# TThostFtdcAgentBrokerIDType 代理经纪公司代码类型
TThostFtdcAgentBrokerIDType = ctypes.c_char * 13

# TThostFtdcAgentGroupIDType 经纪人组代码类型
TThostFtdcAgentGroupIDType = ctypes.c_char * 13

# TThostFtdcAgentGroupNameType 经纪人组名称类型
TThostFtdcAgentGroupNameType = ctypes.c_char * 41

# TThostFtdcAgentIDType 经纪人代码类型
TThostFtdcAgentIDType = ctypes.c_char * 13

# TThostFtdcAgentNameType 经纪人名称类型
TThostFtdcAgentNameType = ctypes.c_char * 41

# TThostFtdcAmAccountType 投资账户类型
TThostFtdcAmAccountType = ctypes.c_char * 23

# TThostFtdcAmlCheckFlowType 反洗钱数据抽取审核流程类型
TThostFtdcAmlCheckFlowType = ctypes.c_char * 2

# TThostFtdcAppIDType App代码类型
TThostFtdcAppIDType = ctypes.c_char * 33

# TThostFtdcAreaCodeType 区号类型
TThostFtdcAreaCodeType = ctypes.c_char * 11

# TThostFtdcAssetmgrApprovalNOType 资产管理业务批文号类型
TThostFtdcAssetmgrApprovalNOType = ctypes.c_char * 51

# TThostFtdcAssetmgrCFullNameType 代理资产管理业务的期货公司全称类型
TThostFtdcAssetmgrCFullNameType = ctypes.c_char * 101

# TThostFtdcAssetmgrMgrNameType 资产管理业务负责人姓名类型
TThostFtdcAssetmgrMgrNameType = ctypes.c_char * 401

# TThostFtdcAuthCodeType 客户端认证码类型
TThostFtdcAuthCodeType = ctypes.c_char * 17

# TThostFtdcAuthInfoType 客户端认证信息类型
TThostFtdcAuthInfoType = ctypes.c_char * 129

# TThostFtdcAuthKeyType 令牌密钥类型
TThostFtdcAuthKeyType = ctypes.c_char * 41

# TThostFtdcAuthenticDataType 认证数据类型
TThostFtdcAuthenticDataType = ctypes.c_char * 129

# TThostFtdcBankAccountNameType 银行帐户名称类型
TThostFtdcBankAccountNameType = ctypes.c_char * 71

# TThostFtdcBankAccountType 银行账户类型
TThostFtdcBankAccountType = ctypes.c_char * 41

# TThostFtdcBankAccountTypeType 账户类别类型
TThostFtdcBankAccountTypeType = ctypes.c_char * 2

# TThostFtdcBankBranchIDType 分中心代码类型
TThostFtdcBankBranchIDType = ctypes.c_char * 11

# TThostFtdcBankBrchIDType 银行分中心代码类型
TThostFtdcBankBrchIDType = ctypes.c_char * 5

# TThostFtdcBankCodingForFutureType 银行对期货公司的编码类型
TThostFtdcBankCodingForFutureType = ctypes.c_char * 33

# TThostFtdcBankCustNoType 银行客户号类型
TThostFtdcBankCustNoType = ctypes.c_char * 21

# TThostFtdcBankFlagType 银行统一标识类型
TThostFtdcBankFlagType = ctypes.c_char * 4

# TThostFtdcBankIDByBankType 银行自己的编码类型
TThostFtdcBankIDByBankType = ctypes.c_char * 21

# TThostFtdcBankIDType 银行代码类型
TThostFtdcBankIDType = ctypes.c_char * 4

# TThostFtdcBankMainKeyType 银行主密钥类型
TThostFtdcBankMainKeyType = ctypes.c_char * 129

# TThostFtdcBankNameType 银行名称类型
TThostFtdcBankNameType = ctypes.c_char * 101

# TThostFtdcBankOperNoType 银行操作员号类型
TThostFtdcBankOperNoType = ctypes.c_char * 4

# TThostFtdcBankReturnCodeType 银行对返回码的定义类型
TThostFtdcBankReturnCodeType = ctypes.c_char * 7

# TThostFtdcBankSerialType 银行流水号类型
TThostFtdcBankSerialType = ctypes.c_char * 13

# TThostFtdcBankServerDescriptionType 银行服务器描述信息类型
TThostFtdcBankServerDescriptionType = ctypes.c_char * 129

# TThostFtdcBankSubBranchIDType 银行分支机构编码类型
TThostFtdcBankSubBranchIDType = ctypes.c_char * 31

# TThostFtdcBankTransKeyType 银行传输密钥类型
TThostFtdcBankTransKeyType = ctypes.c_char * 129

# TThostFtdcBankWorkKeyType 银行工作密钥类型
TThostFtdcBankWorkKeyType = ctypes.c_char * 129

# TThostFtdcBase64AdditionalInfoType base64系统外部信息类型
TThostFtdcBase64AdditionalInfoType = ctypes.c_char * 349

# TThostFtdcBase64ClientSystemInfoType base64交易终端系统信息类型
TThostFtdcBase64ClientSystemInfoType = ctypes.c_char * 365

# TThostFtdcBatchSerialNoType 批次号类型
TThostFtdcBatchSerialNoType = ctypes.c_char * 21

# TThostFtdcBillNameType 票据名称类型
TThostFtdcBillNameType = ctypes.c_char * 33

# TThostFtdcBillNoType 票据号类型
TThostFtdcBillNoType = ctypes.c_char * 15

# TThostFtdcBranchIDType 营业部编号类型
TThostFtdcBranchIDType = ctypes.c_char * 9

# TThostFtdcBranchNetCodeType 机构网点号类型
TThostFtdcBranchNetCodeType = ctypes.c_char * 31

# TThostFtdcBranchNetNameType 机构网点名称类型
TThostFtdcBranchNetNameType = ctypes.c_char * 71

# TThostFtdcBrandCodeType 牌号类型
TThostFtdcBrandCodeType = ctypes.c_char * 257

# TThostFtdcBrokerAbbrType 经纪公司简称类型
TThostFtdcBrokerAbbrType = ctypes.c_char * 9

# TThostFtdcBrokerDNSType 域名类型
TThostFtdcBrokerDNSType = ctypes.c_char * 256

# TThostFtdcBrokerIDType 经纪公司代码类型
TThostFtdcBrokerIDType = ctypes.c_char * 11

# TThostFtdcBrokerNameType 经纪公司名称类型
TThostFtdcBrokerNameType = ctypes.c_char * 81

# TThostFtdcBusinessPeriodType 经营期限类型
TThostFtdcBusinessPeriodType = ctypes.c_char * 21

# TThostFtdcBusinessScopeType 经营范围类型
TThostFtdcBusinessScopeType = ctypes.c_char * 1001

# TThostFtdcBusinessUnitType 业务单元类型
TThostFtdcBusinessUnitType = ctypes.c_char * 21

# TThostFtdcCFMMCKeyType 密钥类型
TThostFtdcCFMMCKeyType = ctypes.c_char * 21

# TThostFtdcCFMMCTokenType 令牌类型
TThostFtdcCFMMCTokenType = ctypes.c_char * 21

# TThostFtdcCSRCAmTypeType 机构类型
TThostFtdcCSRCAmTypeType = ctypes.c_char * 5

# TThostFtdcCSRCBankAccountType 银行账户类型
TThostFtdcCSRCBankAccountType = ctypes.c_char * 23

# TThostFtdcCSRCBankFlagType 银行标识类型
TThostFtdcCSRCBankFlagType = ctypes.c_char * 3

# TThostFtdcCSRCCancelFlagType 新增或变更标志类型
TThostFtdcCSRCCancelFlagType = ctypes.c_char * 2

# TThostFtdcCSRCClientIDType 交易编码类型
TThostFtdcCSRCClientIDType = ctypes.c_char * 11

# TThostFtdcCSRCDateType 日期类型
TThostFtdcCSRCDateType = ctypes.c_char * 11

# TThostFtdcCSRCExchangeInstIDType 合约代码类型
TThostFtdcCSRCExchangeInstIDType = ctypes.c_char * 31

# TThostFtdcCSRCFreezeStatusType 休眠状态类型
TThostFtdcCSRCFreezeStatusType = ctypes.c_char * 2

# TThostFtdcCSRCIdentifiedCardNoType 证件号码类型
TThostFtdcCSRCIdentifiedCardNoType = ctypes.c_char * 51

# TThostFtdcCSRCInvestorIDType 客户代码类型
TThostFtdcCSRCInvestorIDType = ctypes.c_char * 13

# TThostFtdcCSRCInvestorNameType 客户名称类型
TThostFtdcCSRCInvestorNameType = ctypes.c_char * 201

# TThostFtdcCSRCMemo1Type 说明类型
TThostFtdcCSRCMemo1Type = ctypes.c_char * 41

# TThostFtdcCSRCMemoType 说明类型
TThostFtdcCSRCMemoType = ctypes.c_char * 101

# TThostFtdcCSRCMortgageNameType 质押品名称类型
TThostFtdcCSRCMortgageNameType = ctypes.c_char * 7

# TThostFtdcCSRCNationalType 国籍类型
TThostFtdcCSRCNationalType = ctypes.c_char * 4

# TThostFtdcCSRCOpenInvestorNameType 客户名称类型
TThostFtdcCSRCOpenInvestorNameType = ctypes.c_char * 101

# TThostFtdcCSRCOpenNameType 开户人类型
TThostFtdcCSRCOpenNameType = ctypes.c_char * 401

# TThostFtdcCSRCOptionsTypeType 期权类型
TThostFtdcCSRCOptionsTypeType = ctypes.c_char * 2

# TThostFtdcCSRCReasonType 事由类型
TThostFtdcCSRCReasonType = ctypes.c_char * 3

# TThostFtdcCSRCSecAgentIDType 二级代理ID类型
TThostFtdcCSRCSecAgentIDType = ctypes.c_char * 11

# TThostFtdcCSRCTargetInstrIDType 标的合约类型
TThostFtdcCSRCTargetInstrIDType = ctypes.c_char * 31

# TThostFtdcCSRCTargetProductIDType 标的品种类型
TThostFtdcCSRCTargetProductIDType = ctypes.c_char * 3

# TThostFtdcCSRCTimeType 时间类型
TThostFtdcCSRCTimeType = ctypes.c_char * 11

# TThostFtdcCSRCTradeIDType 成交流水号类型
TThostFtdcCSRCTradeIDType = ctypes.c_char * 21

# TThostFtdcCapitalCurrencyType 注册资本币种类型
TThostFtdcCapitalCurrencyType = ctypes.c_char * 4

# TThostFtdcCaptchaInfoType 图片验证信息类型
TThostFtdcCaptchaInfoType = ctypes.c_char * 2561

# TThostFtdcCertCodeType 证件号码类型
TThostFtdcCertCodeType = ctypes.c_char * 21

# TThostFtdcCffexDepartmentCodeType 营业部代码类型
TThostFtdcCffexDepartmentCodeType = ctypes.c_char * 9

# TThostFtdcCffexDepartmentNameType 开户营业部类型
TThostFtdcCffexDepartmentNameType = ctypes.c_char * 101

# TThostFtdcCffmcDateType 日期类型
TThostFtdcCffmcDateType = ctypes.c_char * 11

# TThostFtdcCffmcTimeType 时间类型
TThostFtdcCffmcTimeType = ctypes.c_char * 11

# TThostFtdcChannelType 渠道类型
TThostFtdcChannelType = ctypes.c_char * 51

# TThostFtdcCharacterIDType 交易特征代码类型
TThostFtdcCharacterIDType = ctypes.c_char * 5

# TThostFtdcCheckResultMemoType 核对结果说明类型
TThostFtdcCheckResultMemoType = ctypes.c_char * 1025

# TThostFtdcCityType 市类型
TThostFtdcCityType = ctypes.c_char * 51

# TThostFtdcClassifyType 类别类型
TThostFtdcClassifyType = ctypes.c_char * 41

# TThostFtdcClearAccountType 结算账户类型
TThostFtdcClearAccountType = ctypes.c_char * 33

# TThostFtdcClearBrchIDType 机构结算帐户联行号类型
TThostFtdcClearBrchIDType = ctypes.c_char * 6

# TThostFtdcClearDepIDType 机构结算帐户机构号类型
TThostFtdcClearDepIDType = ctypes.c_char * 6

# TThostFtdcClearNameType 机构结算帐户名称类型
TThostFtdcClearNameType = ctypes.c_char * 71

# TThostFtdcClearbarchIDType 结算账户联行号类型
TThostFtdcClearbarchIDType = ctypes.c_char * 6

# TThostFtdcClientClassifyType 客户分类码类型
TThostFtdcClientClassifyType = ctypes.c_char * 11

# TThostFtdcClientIDType 交易编码类型
TThostFtdcClientIDType = ctypes.c_char * 11

# TThostFtdcClientLoginRemarkType 客户登录备注2类型
TThostFtdcClientLoginRemarkType = ctypes.c_char * 151

# TThostFtdcClientModeType 开户模式类型
TThostFtdcClientModeType = ctypes.c_char * 3

# TThostFtdcClientSystemInfoType 交易终端系统信息类型
TThostFtdcClientSystemInfoType = ctypes.c_char * 273

# TThostFtdcCollectTimeType 信息采集时间类型
TThostFtdcCollectTimeType = ctypes.c_char * 21

# TThostFtdcCombHedgeFlagType 组合投机套保标志类型
TThostFtdcCombHedgeFlagType = ctypes.c_char * 5

# TThostFtdcCombOffsetFlagType 组合开平标志类型
TThostFtdcCombOffsetFlagType = ctypes.c_char * 5

# TThostFtdcCombinInstrIDType 套利合约代码类型
TThostFtdcCombinInstrIDType = ctypes.c_char * 61

# TThostFtdcCombinSettlePriceType 各腿结算价类型
TThostFtdcCombinSettlePriceType = ctypes.c_char * 61

# TThostFtdcCombineIDType 组合编号类型
TThostFtdcCombineIDType = ctypes.c_char * 25

# TThostFtdcCombineTypeType 组合类型
TThostFtdcCombineTypeType = ctypes.c_char * 25

# TThostFtdcComeFromType 消息来源类型
TThostFtdcComeFromType = ctypes.c_char * 21

# TThostFtdcCommModelMemoType 手续费率模板备注类型
TThostFtdcCommModelMemoType = ctypes.c_char * 1025

# TThostFtdcCommModelNameType 手续费率模板名称类型
TThostFtdcCommModelNameType = ctypes.c_char * 161

# TThostFtdcCommandTypeType DB命令类型
TThostFtdcCommandTypeType = ctypes.c_char * 65

# TThostFtdcCommentType 盈亏算法说明类型
TThostFtdcCommentType = ctypes.c_char * 31

# TThostFtdcCompanyCodeType 企业代码类型
TThostFtdcCompanyCodeType = ctypes.c_char * 51

# TThostFtdcCompanyTypeType 企业性质类型
TThostFtdcCompanyTypeType = ctypes.c_char * 16

# TThostFtdcContentType 消息正文类型
TThostFtdcContentType = ctypes.c_char * 501

# TThostFtdcContractCodeType 合同编号类型
TThostFtdcContractCodeType = ctypes.c_char * 41

# TThostFtdcCorporateIdentifiedCardNoType 法人代表证件号码类型
TThostFtdcCorporateIdentifiedCardNoType = ctypes.c_char * 101

# TThostFtdcCounterIDType 计数器代码类型
TThostFtdcCounterIDType = ctypes.c_char * 33

# TThostFtdcCountryCodeType 国家代码类型
TThostFtdcCountryCodeType = ctypes.c_char * 21

# TThostFtdcCountryType 国家类型
TThostFtdcCountryType = ctypes.c_char * 16

# TThostFtdcCryptoKeyVersionType api与front通信密钥版本号类型
TThostFtdcCryptoKeyVersionType = ctypes.c_char * 31

# TThostFtdcCurrExchCertNoType 凭证号类型
TThostFtdcCurrExchCertNoType = ctypes.c_char * 13

# TThostFtdcCurrencyCodeType 币种类型
TThostFtdcCurrencyCodeType = ctypes.c_char * 4

# TThostFtdcCurrencyIDType 币种代码类型
TThostFtdcCurrencyIDType = ctypes.c_char * 4

# TThostFtdcCurrencyNameType 币种名称类型
TThostFtdcCurrencyNameType = ctypes.c_char * 31

# TThostFtdcCurrencySignType 币种符号类型
TThostFtdcCurrencySignType = ctypes.c_char * 4

# TThostFtdcCurrencySwapMemoType 换汇需确认信息类型
TThostFtdcCurrencySwapMemoType = ctypes.c_char * 101

# TThostFtdcCustNumberType 客户编号类型
TThostFtdcCustNumberType = ctypes.c_char * 36

# TThostFtdcDBLinkIDType DBLink标识号类型
TThostFtdcDBLinkIDType = ctypes.c_char * 31

# TThostFtdcDRIdentityNameType 交易中心名称类型
TThostFtdcDRIdentityNameType = ctypes.c_char * 65

# TThostFtdcDataTypeType 数据类型
TThostFtdcDataTypeType = ctypes.c_char * 129

# TThostFtdcDateExprType 日期表达式类型
TThostFtdcDateExprType = ctypes.c_char * 1025

# TThostFtdcDateTimeType 日期时间类型
TThostFtdcDateTimeType = ctypes.c_char * 17

# TThostFtdcDateType 日期类型
TThostFtdcDateType = ctypes.c_char * 9

# TThostFtdcDepositSeqNoType 出入金流水号类型
TThostFtdcDepositSeqNoType = ctypes.c_char * 15

# TThostFtdcDescrInfoForReturnCodeType 返回码描述类型
TThostFtdcDescrInfoForReturnCodeType = ctypes.c_char * 129

# TThostFtdcDescriptionType 描述类型
TThostFtdcDescriptionType = ctypes.c_char * 401

# TThostFtdcDeviceIDType 渠道标志类型
TThostFtdcDeviceIDType = ctypes.c_char * 3

# TThostFtdcDigestType 摘要类型
TThostFtdcDigestType = ctypes.c_char * 36

# TThostFtdcEMailType 电子邮件类型
TThostFtdcEMailType = ctypes.c_char * 41

# TThostFtdcEnumValueIDType 枚举值代码类型
TThostFtdcEnumValueIDType = ctypes.c_char * 65

# TThostFtdcEnumValueLabelType 枚举值名称类型
TThostFtdcEnumValueLabelType = ctypes.c_char * 65

# TThostFtdcEnumValueResultType 枚举值结果类型
TThostFtdcEnumValueResultType = ctypes.c_char * 33

# TThostFtdcEnumValueTypeType 枚举值类型
TThostFtdcEnumValueTypeType = ctypes.c_char * 33

# TThostFtdcErrorMsgType 错误信息类型
TThostFtdcErrorMsgType = ctypes.c_char * 81

# TThostFtdcEventTypeType 业务操作类型
TThostFtdcEventTypeType = ctypes.c_char * 33

# TThostFtdcExchangeAbbrType 交易所简称类型
TThostFtdcExchangeAbbrType = ctypes.c_char * 9

# TThostFtdcExchangeFlagType 交易所标志类型
TThostFtdcExchangeFlagType = ctypes.c_char * 2

# TThostFtdcExchangeIDType 交易所代码类型
TThostFtdcExchangeIDType = ctypes.c_char * 9

# TThostFtdcExchangeInstIDType 合约在交易所的代码类型
TThostFtdcExchangeInstIDType = ctypes.c_char * 81

# TThostFtdcExchangeNameType 交易所名称类型
TThostFtdcExchangeNameType = ctypes.c_char * 61

# TThostFtdcExecOrderSysIDType 执行宣告系统编号类型
TThostFtdcExecOrderSysIDType = ctypes.c_char * 21

# TThostFtdcFBEBankAccountNameType 换汇银行账户名类型
TThostFtdcFBEBankAccountNameType = ctypes.c_char * 61

# TThostFtdcFBEBankAccountType 换汇银行账户类型
TThostFtdcFBEBankAccountType = ctypes.c_char * 33

# TThostFtdcFBEBankNoType 换汇银行行号类型
TThostFtdcFBEBankNoType = ctypes.c_char * 13

# TThostFtdcFBEBatchSerialType 换汇批次号类型
TThostFtdcFBEBatchSerialType = ctypes.c_char * 21

# TThostFtdcFBEBusinessSerialType 换汇记账流水号类型
TThostFtdcFBEBusinessSerialType = ctypes.c_char * 31

# TThostFtdcFBEBusinessTypeType 换汇业务类型
TThostFtdcFBEBusinessTypeType = ctypes.c_char * 3

# TThostFtdcFBECertNoType 换汇凭证号类型
TThostFtdcFBECertNoType = ctypes.c_char * 13

# TThostFtdcFBEExtendMsgType 换汇扩展信息类型
TThostFtdcFBEExtendMsgType = ctypes.c_char * 61

# TThostFtdcFBEFileNameType 换汇相关文件名类型
TThostFtdcFBEFileNameType = ctypes.c_char * 21

# TThostFtdcFBEOpenBankType 换汇账户开户行类型
TThostFtdcFBEOpenBankType = ctypes.c_char * 61

# TThostFtdcFBEPostScriptType 换汇附言类型
TThostFtdcFBEPostScriptType = ctypes.c_char * 61

# TThostFtdcFBERemarkType 换汇备注类型
TThostFtdcFBERemarkType = ctypes.c_char * 71

# TThostFtdcFBERtnMsgType 换汇返回信息类型
TThostFtdcFBERtnMsgType = ctypes.c_char * 61

# TThostFtdcFBESystemSerialType 换汇流水号类型
TThostFtdcFBESystemSerialType = ctypes.c_char * 21

# TThostFtdcFBETimeType 各种换汇时间类型
TThostFtdcFBETimeType = ctypes.c_char * 7

# TThostFtdcFaxType 传真类型
TThostFtdcFaxType = ctypes.c_char * 41

# TThostFtdcFetchAmtType 银行可取余额类型
TThostFtdcFetchAmtType = ctypes.c_char * 20

# TThostFtdcFieldContentType 字段内容类型
TThostFtdcFieldContentType = ctypes.c_char * 2049

# TThostFtdcFieldNameType 字段名类型
TThostFtdcFieldNameType = ctypes.c_char * 2049

# TThostFtdcFileNameType 文件名称类型
TThostFtdcFileNameType = ctypes.c_char * 257

# TThostFtdcForceCloseSceneIdType 强平场景编号类型
TThostFtdcForceCloseSceneIdType = ctypes.c_char * 24

# TThostFtdcFunctionIDType 功能代码类型
TThostFtdcFunctionIDType = ctypes.c_char * 25

# TThostFtdcFunctionNameType 功能名称类型
TThostFtdcFunctionNameType = ctypes.c_char * 65

# TThostFtdcFunctionUrlType 功能链接类型
TThostFtdcFunctionUrlType = ctypes.c_char * 1025

# TThostFtdcFunctionValueCodeType 功能编码类型
TThostFtdcFunctionValueCodeType = ctypes.c_char * 257

# TThostFtdcFundProjectIDType 资金项目编号类型
TThostFtdcFundProjectIDType = ctypes.c_char * 5

# TThostFtdcFutureAccPwdType 期货资金密码类型
TThostFtdcFutureAccPwdType = ctypes.c_char * 17

# TThostFtdcFutureAccountNameType 期货帐户名称类型
TThostFtdcFutureAccountNameType = ctypes.c_char * 129

# TThostFtdcFutureAccountType 期货资金账号类型
TThostFtdcFutureAccountType = ctypes.c_char * 22

# TThostFtdcFutureBranchIDType 期货分支机构编码类型
TThostFtdcFutureBranchIDType = ctypes.c_char * 31

# TThostFtdcFutureIDType 期货公司代码类型
TThostFtdcFutureIDType = ctypes.c_char * 11

# TThostFtdcFutureMainKeyType 期货公司主密钥类型
TThostFtdcFutureMainKeyType = ctypes.c_char * 129

# TThostFtdcFutureTransKeyType 期货公司传输密钥类型
TThostFtdcFutureTransKeyType = ctypes.c_char * 129

# TThostFtdcFutureWorkKeyType 期货公司工作密钥类型
TThostFtdcFutureWorkKeyType = ctypes.c_char * 129

# TThostFtdcFuturesIDType 监控中心为客户分配的代码类型
TThostFtdcFuturesIDType = ctypes.c_char * 21

# TThostFtdcGradeType 等级类型
TThostFtdcGradeType = ctypes.c_char * 41

# TThostFtdcHandshakeDataType 握手数据内容类型
TThostFtdcHandshakeDataType = ctypes.c_char * 301

# TThostFtdcIDBNameType 握手数据内容类型
TThostFtdcIDBNameType = ctypes.c_char * 100

# TThostFtdcIPAddressType IP地址类型
TThostFtdcIPAddressType = ctypes.c_char * 33

# TThostFtdcIdentifiedCardNoType 证件号码类型
TThostFtdcIdentifiedCardNoType = ctypes.c_char * 51

# TThostFtdcImportSequenceIDType 动态令牌导入批次编号类型
TThostFtdcImportSequenceIDType = ctypes.c_char * 17

# TThostFtdcInTheMoneyFlagType 平值期权标志类型
TThostFtdcInTheMoneyFlagType = ctypes.c_char * 2

# TThostFtdcIndividualNameType 个人姓名类型
TThostFtdcIndividualNameType = ctypes.c_char * 51

# TThostFtdcIndustryIDType 行业编码类型
TThostFtdcIndustryIDType = ctypes.c_char * 17

# TThostFtdcInstrumentCodeType 合约标识码类型
TThostFtdcInstrumentCodeType = ctypes.c_char * 31

# TThostFtdcInstrumentIDExprType 合约代码表达式类型
TThostFtdcInstrumentIDExprType = ctypes.c_char * 41

# TThostFtdcInstrumentIDType 合约代码类型
TThostFtdcInstrumentIDType = ctypes.c_char * 81

# TThostFtdcInstrumentIDsType 多个产品代码,用+分隔,如cu+zn类型
TThostFtdcInstrumentIDsType = ctypes.c_char * 101

# TThostFtdcInstrumentNameExprType 合约名称表达式类型
TThostFtdcInstrumentNameExprType = ctypes.c_char * 41

# TThostFtdcInstrumentNameType 合约名称类型
TThostFtdcInstrumentNameType = ctypes.c_char * 21

# TThostFtdcInvBrchIDType 机构投资人联行号类型
TThostFtdcInvBrchIDType = ctypes.c_char * 6

# TThostFtdcInvDepIDType 机构投资人账号机构号类型
TThostFtdcInvDepIDType = ctypes.c_char * 6

# TThostFtdcInvestUnitIDType 投资单元代码类型
TThostFtdcInvestUnitIDType = ctypes.c_char * 17

# TThostFtdcInvestVarietyType 投资品种类型
TThostFtdcInvestVarietyType = ctypes.c_char * 101

# TThostFtdcInvestorFullNameType 投资者全称类型
TThostFtdcInvestorFullNameType = ctypes.c_char * 101

# TThostFtdcInvestorGroupNameType 投资者分组名称类型
TThostFtdcInvestorGroupNameType = ctypes.c_char * 41

# TThostFtdcInvestorIDRuleExprType 号段规则表达式类型
TThostFtdcInvestorIDRuleExprType = ctypes.c_char * 513

# TThostFtdcInvestorIDRuleNameType 号段规则名称类型
TThostFtdcInvestorIDRuleNameType = ctypes.c_char * 61

# TThostFtdcInvestorIDType 投资者代码类型
TThostFtdcInvestorIDType = ctypes.c_char * 13

# TThostFtdcIsSettlementType 是否为非结算会员类型
TThostFtdcIsSettlementType = ctypes.c_char * 2

# TThostFtdcIsStockType 是否股民类型
TThostFtdcIsStockType = ctypes.c_char * 11

# TThostFtdcLedgerManageBankType 开户银行类型
TThostFtdcLedgerManageBankType = ctypes.c_char * 101

# TThostFtdcLedgerManageIDType 分户管理资产编码类型
TThostFtdcLedgerManageIDType = ctypes.c_char * 51

# TThostFtdcLicenseNOType 营业执照类型
TThostFtdcLicenseNOType = ctypes.c_char * 33

# TThostFtdcLicenseNoType 营业执照号类型
TThostFtdcLicenseNoType = ctypes.c_char * 51

# TThostFtdcLogLevelType 日志级别类型
TThostFtdcLogLevelType = ctypes.c_char * 33

# TThostFtdcLoginRemarkType 登录备注类型
TThostFtdcLoginRemarkType = ctypes.c_char * 36

# TThostFtdcLongFBEBankAccountNameType 长换汇银行账户名类型
TThostFtdcLongFBEBankAccountNameType = ctypes.c_char * 161

# TThostFtdcLongIndividualNameType 长个人姓名类型
TThostFtdcLongIndividualNameType = ctypes.c_char * 161

# TThostFtdcLongTimeType 长时间类型
TThostFtdcLongTimeType = ctypes.c_char * 13

# TThostFtdcMacAddressType Mac地址类型
TThostFtdcMacAddressType = ctypes.c_char * 21

# TThostFtdcMarketIDType 市场代码类型
TThostFtdcMarketIDType = ctypes.c_char * 31

# TThostFtdcMemoType 备注类型
TThostFtdcMemoType = ctypes.c_char * 161

# TThostFtdcMessageFormatVersionType 信息格式版本类型
TThostFtdcMessageFormatVersionType = ctypes.c_char * 36

# TThostFtdcMobilePhoneType 手机类型
TThostFtdcMobilePhoneType = ctypes.c_char * 21

# TThostFtdcMobileType 手机类型
TThostFtdcMobileType = ctypes.c_char * 41

# TThostFtdcNationalType 国籍类型
TThostFtdcNationalType = ctypes.c_char * 31

# TThostFtdcNewsTypeType 公告类型
TThostFtdcNewsTypeType = ctypes.c_char * 3

# TThostFtdcNocIDType 组织机构代码类型
TThostFtdcNocIDType = ctypes.c_char * 21

# TThostFtdcOTCTraderIDType OTC交易员代码类型
TThostFtdcOTCTraderIDType = ctypes.c_char * 31

# TThostFtdcOTPVendorsIDType 动态令牌提供商类型
TThostFtdcOTPVendorsIDType = ctypes.c_char * 2

# TThostFtdcOTPVendorsNameType 动态令牌提供商名称类型
TThostFtdcOTPVendorsNameType = ctypes.c_char * 61

# TThostFtdcOldCityType 城市类型
TThostFtdcOldCityType = ctypes.c_char * 41

# TThostFtdcOldExchangeInstIDType 合约在交易所的代码类型
TThostFtdcOldExchangeInstIDType = ctypes.c_char * 31

# TThostFtdcOldIPAddressType IP地址类型
TThostFtdcOldIPAddressType = ctypes.c_char * 16

# TThostFtdcOldInstrumentIDType 合约代码类型
TThostFtdcOldInstrumentIDType = ctypes.c_char * 31

# TThostFtdcOpenBankType 银行账户的开户行类型
TThostFtdcOpenBankType = ctypes.c_char * 101

# TThostFtdcOpenNameType 银行账户的开户人名称类型
TThostFtdcOpenNameType = ctypes.c_char * 61

# TThostFtdcOperNoType 交易柜员类型
TThostFtdcOperNoType = ctypes.c_char * 17

# TThostFtdcOperationMemoType 操作摘要类型
TThostFtdcOperationMemoType = ctypes.c_char * 1025

# TThostFtdcOperatorCodeType 操作员类型
TThostFtdcOperatorCodeType = ctypes.c_char * 17

# TThostFtdcOperatorIDType 操作员代码类型
TThostFtdcOperatorIDType = ctypes.c_char * 65

# TThostFtdcOptionContentType 选项说明类型
TThostFtdcOptionContentType = ctypes.c_char * 61

# TThostFtdcOptionIDType 选项编号类型
TThostFtdcOptionIDType = ctypes.c_char * 13

# TThostFtdcOrderLocalIDType 本地报单编号类型
TThostFtdcOrderLocalIDType = ctypes.c_char * 13

# TThostFtdcOrderMemoType 报单回显字段类型
TThostFtdcOrderMemoType = ctypes.c_char * 13

# TThostFtdcOrderRefType 报单引用类型
TThostFtdcOrderRefType = ctypes.c_char * 13

# TThostFtdcOrderSysIDType 报单编号类型
TThostFtdcOrderSysIDType = ctypes.c_char * 21

# TThostFtdcOrganCodeType 机构编码类型
TThostFtdcOrganCodeType = ctypes.c_char * 36

# TThostFtdcOrganFlagType 机构标识类型
TThostFtdcOrganFlagType = ctypes.c_char * 2

# TThostFtdcOrganNOType 结算账户类型
TThostFtdcOrganNOType = ctypes.c_char * 6

# TThostFtdcOrganNameType 机构名称类型
TThostFtdcOrganNameType = ctypes.c_char * 71

# TThostFtdcPKNameType FBT表操作主键名类型
TThostFtdcPKNameType = ctypes.c_char * 201

# TThostFtdcPKValueType FBT表操作主键值类型
TThostFtdcPKValueType = ctypes.c_char * 501

# TThostFtdcPageControlType 换汇页面控制类型
TThostFtdcPageControlType = ctypes.c_char * 2

# TThostFtdcParamNameType 参数名类型
TThostFtdcParamNameType = ctypes.c_char * 41

# TThostFtdcParamValueType 参数值类型
TThostFtdcParamValueType = ctypes.c_char * 41

# TThostFtdcParkedOrderActionIDType 预埋撤单编号类型
TThostFtdcParkedOrderActionIDType = ctypes.c_char * 13

# TThostFtdcParkedOrderIDType 预埋报单编号类型
TThostFtdcParkedOrderIDType = ctypes.c_char * 13

# TThostFtdcParticipantIDType 会员代码类型
TThostFtdcParticipantIDType = ctypes.c_char * 11

# TThostFtdcPartyNameType 参与人名称类型
TThostFtdcPartyNameType = ctypes.c_char * 81

# TThostFtdcPasswordKeyType 密钥类型
TThostFtdcPasswordKeyType = ctypes.c_char * 129

# TThostFtdcPasswordType 密码类型
TThostFtdcPasswordType = ctypes.c_char * 41

# TThostFtdcPhotoNameType 影像名称类型
TThostFtdcPhotoNameType = ctypes.c_char * 161

# TThostFtdcPhotoTypeIDType 影像类型
TThostFtdcPhotoTypeIDType = ctypes.c_char * 5

# TThostFtdcPhotoTypeNameType 影像类型
TThostFtdcPhotoTypeNameType = ctypes.c_char * 41

# TThostFtdcPlateReturnCodeType 银期转帐平台对返回码的定义类型
TThostFtdcPlateReturnCodeType = ctypes.c_char * 5

# TThostFtdcPositionType 货位类型
TThostFtdcPositionType = ctypes.c_char * 41

# TThostFtdcPriceDecimalType 价格小数位类型
TThostFtdcPriceDecimalType = ctypes.c_char * 2

# TThostFtdcProcessIDType 业务流水号类型
TThostFtdcProcessIDType = ctypes.c_char * 33

# TThostFtdcProcessNameType 存储过程名称类型
TThostFtdcProcessNameType = ctypes.c_char * 257

# TThostFtdcProcessTypeType 流程功能类型
TThostFtdcProcessTypeType = ctypes.c_char * 3

# TThostFtdcProductDateType 产期类型
TThostFtdcProductDateType = ctypes.c_char * 41

# TThostFtdcProductIDType 产品ID类型
TThostFtdcProductIDType = ctypes.c_char * 41

# TThostFtdcProductInfoType 产品信息类型
TThostFtdcProductInfoType = ctypes.c_char * 11

# TThostFtdcProductNameType 产品名称类型
TThostFtdcProductNameType = ctypes.c_char * 21

# TThostFtdcProfessionType 职业类型
TThostFtdcProfessionType = ctypes.c_char * 101

# TThostFtdcPropertyIDType 属性代码类型
TThostFtdcPropertyIDType = ctypes.c_char * 33

# TThostFtdcPropertyNameType 属性名称类型
TThostFtdcPropertyNameType = ctypes.c_char * 65

# TThostFtdcPropertyStringType 用于查询的投资属性字段类型
TThostFtdcPropertyStringType = ctypes.c_char * 2049

# TThostFtdcProtocolInfoType 协议信息类型
TThostFtdcProtocolInfoType = ctypes.c_char * 11

# TThostFtdcProvinceType 省类型
TThostFtdcProvinceType = ctypes.c_char * 51

# TThostFtdcPublishPathType 发布路径类型
TThostFtdcPublishPathType = ctypes.c_char * 257

# TThostFtdcQuestionContentType 特有信息说明类型
TThostFtdcQuestionContentType = ctypes.c_char * 41

# TThostFtdcQuestionIDType 特有信息编号类型
TThostFtdcQuestionIDType = ctypes.c_char * 5

# TThostFtdcRandomStringType 随机串类型
TThostFtdcRandomStringType = ctypes.c_char * 17

# TThostFtdcRangeIntFromType 限定值下限类型
TThostFtdcRangeIntFromType = ctypes.c_char * 33

# TThostFtdcRangeIntToType 限定值上限类型
TThostFtdcRangeIntToType = ctypes.c_char * 33

# TThostFtdcRangeIntTypeType 限定值类型
TThostFtdcRangeIntTypeType = ctypes.c_char * 33

# TThostFtdcRateTemplateIDType 模型代码类型
TThostFtdcRateTemplateIDType = ctypes.c_char * 9

# TThostFtdcRateTemplateNameType 模型名称类型
TThostFtdcRateTemplateNameType = ctypes.c_char * 61

# TThostFtdcRecordNumType 记录数类型
TThostFtdcRecordNumType = ctypes.c_char * 7

# TThostFtdcRegionType 区类型
TThostFtdcRegionType = ctypes.c_char * 16

# TThostFtdcReportTypeIDType 交易报告类型
TThostFtdcReportTypeIDType = ctypes.c_char * 3

# TThostFtdcRetCodeType 响应代码类型
TThostFtdcRetCodeType = ctypes.c_char * 5

# TThostFtdcRetInfoType 响应信息类型
TThostFtdcRetInfoType = ctypes.c_char * 129

# TThostFtdcReturnCodeType 返回代码类型
TThostFtdcReturnCodeType = ctypes.c_char * 7

# TThostFtdcRightTemplateIDType 模板代码类型
TThostFtdcRightTemplateIDType = ctypes.c_char * 9

# TThostFtdcRightTemplateNameType 模板名称类型
TThostFtdcRightTemplateNameType = ctypes.c_char * 61

# TThostFtdcRiskNofityInfoType 客户风险通知消息类型
TThostFtdcRiskNofityInfoType = ctypes.c_char * 257

# TThostFtdcRiskRateType 风险度类型
TThostFtdcRiskRateType = ctypes.c_char * 21

# TThostFtdcRoleIDType 角色编号类型
TThostFtdcRoleIDType = ctypes.c_char * 11

# TThostFtdcRoleNameType 角色名称类型
TThostFtdcRoleNameType = ctypes.c_char * 41

# TThostFtdcRuleIdType 策略id类型
TThostFtdcRuleIdType = ctypes.c_char * 51

# TThostFtdcSHFEInstLifePhaseType 上期所合约生命周期状态类型
TThostFtdcSHFEInstLifePhaseType = ctypes.c_char * 3

# TThostFtdcSHFEProductClassType 产品类型
TThostFtdcSHFEProductClassType = ctypes.c_char * 11

# TThostFtdcSPMMModelDescType SPMM模板描述类型
TThostFtdcSPMMModelDescType = ctypes.c_char * 129

# TThostFtdcSPMMModelIDType SPMM模板ID类型
TThostFtdcSPMMModelIDType = ctypes.c_char * 33

# TThostFtdcSPMMProductIDType SPMM商品群商品组ID类型
TThostFtdcSPMMProductIDType = ctypes.c_char * 41

# TThostFtdcSRiskRateType 风险度类型
TThostFtdcSRiskRateType = ctypes.c_char * 21

# TThostFtdcSentenceType 语句类型
TThostFtdcSentenceType = ctypes.c_char * 501

# TThostFtdcSequenceLabelType 序列编号类型
TThostFtdcSequenceLabelType = ctypes.c_char * 2

# TThostFtdcSerialNumberType 序列号类型
TThostFtdcSerialNumberType = ctypes.c_char * 17

# TThostFtdcServiceNameType 服务名类型
TThostFtdcServiceNameType = ctypes.c_char * 61

# TThostFtdcSettleManagerIDType 结算配置代码类型
TThostFtdcSettleManagerIDType = ctypes.c_char * 33

# TThostFtdcSettleManagerNameType 结算配置名称类型
TThostFtdcSettleManagerNameType = ctypes.c_char * 129

# TThostFtdcSettlementGroupIDType 结算组代码类型
TThostFtdcSettlementGroupIDType = ctypes.c_char * 9

# TThostFtdcSettlementParamValueType 参数代码值类型
TThostFtdcSettlementParamValueType = ctypes.c_char * 256

# TThostFtdcSoftwareProviderIDType 交易软件商ID类型
TThostFtdcSoftwareProviderIDType = ctypes.c_char * 22

# TThostFtdcStrikeTimeType 执行时间类型
TThostFtdcStrikeTimeType = ctypes.c_char * 13

# TThostFtdcSubBranchIDType 分支机构类型
TThostFtdcSubBranchIDType = ctypes.c_char * 31

# TThostFtdcSubBranchNameType 分支机构名称类型
TThostFtdcSubBranchNameType = ctypes.c_char * 71

# TThostFtdcSuperOrganCodeType 上级机构编码,即期货公司总部、银行总行类型
TThostFtdcSuperOrganCodeType = ctypes.c_char * 12

# TThostFtdcSwapBusinessTypeType 换汇业务种类类型
TThostFtdcSwapBusinessTypeType = ctypes.c_char * 3

# TThostFtdcSyncDescriptionType 追平描述类型
TThostFtdcSyncDescriptionType = ctypes.c_char * 257

# TThostFtdcSysVersionType 系统版本类型
TThostFtdcSysVersionType = ctypes.c_char * 41

# TThostFtdcSystemIDType 系统编号类型
TThostFtdcSystemIDType = ctypes.c_char * 21

# TThostFtdcSystemNameType 系统名称类型
TThostFtdcSystemNameType = ctypes.c_char * 41

# TThostFtdcTableNameType FBT表名类型
TThostFtdcTableNameType = ctypes.c_char * 61

# TThostFtdcTargetIDType 同步目标编号类型
TThostFtdcTargetIDType = ctypes.c_char * 4

# TThostFtdcTaxNoType 税务登记号类型
TThostFtdcTaxNoType = ctypes.c_char * 31

# TThostFtdcTelephoneType 联系电话类型
TThostFtdcTelephoneType = ctypes.c_char * 41

# TThostFtdcTimeSpanType 时间跨度类型
TThostFtdcTimeSpanType = ctypes.c_char * 9

# TThostFtdcTimeType 时间类型
TThostFtdcTimeType = ctypes.c_char * 9

# TThostFtdcToolIDType 工具代码类型
TThostFtdcToolIDType = ctypes.c_char * 9

# TThostFtdcToolNameType 工具名称类型
TThostFtdcToolNameType = ctypes.c_char * 81

# TThostFtdcTradeAmtType 银行总余额类型
TThostFtdcTradeAmtType = ctypes.c_char * 20

# TThostFtdcTradeCodeType 交易代码类型
TThostFtdcTradeCodeType = ctypes.c_char * 7

# TThostFtdcTradeDateType 交易日期类型
TThostFtdcTradeDateType = ctypes.c_char * 9

# TThostFtdcTradeIDType 成交编号类型
TThostFtdcTradeIDType = ctypes.c_char * 21

# TThostFtdcTradeSerialType 发起方流水号类型
TThostFtdcTradeSerialType = ctypes.c_char * 9

# TThostFtdcTradeTimeType 交易时间类型
TThostFtdcTradeTimeType = ctypes.c_char * 9

# TThostFtdcTraderIDType 交易所交易员代码类型
TThostFtdcTraderIDType = ctypes.c_char * 21

# TThostFtdcUOABrokerIDType 境外中介机构ID类型
TThostFtdcUOABrokerIDType = ctypes.c_char * 11

# TThostFtdcUOACountryCodeType 国家代码类型
TThostFtdcUOACountryCodeType = ctypes.c_char * 11

# TThostFtdcUOAEMailType 电子邮箱类型
TThostFtdcUOAEMailType = ctypes.c_char * 101

# TThostFtdcUOAIdCardTypeType 统一开户证件类型
TThostFtdcUOAIdCardTypeType = ctypes.c_char * 3

# TThostFtdcUOAOrganTypeType 单位性质类型
TThostFtdcUOAOrganTypeType = ctypes.c_char * 11

# TThostFtdcUOAProcessStatusType 流程状态类型
TThostFtdcUOAProcessStatusType = ctypes.c_char * 3

# TThostFtdcUOAZipCodeType 邮政编码类型
TThostFtdcUOAZipCodeType = ctypes.c_char * 11

# TThostFtdcUOMType 计量单位类型
TThostFtdcUOMType = ctypes.c_char * 11

# TThostFtdcURLLinkType WEB地址类型
TThostFtdcURLLinkType = ctypes.c_char * 201

# TThostFtdcUploadModeType 上传文件类型
TThostFtdcUploadModeType = ctypes.c_char * 21

# TThostFtdcUseAmtType 银行可用余额类型
TThostFtdcUseAmtType = ctypes.c_char * 20

# TThostFtdcUserEventInfoType 用户事件信息类型
TThostFtdcUserEventInfoType = ctypes.c_char * 1025

# TThostFtdcUserIDType 用户代码类型
TThostFtdcUserIDType = ctypes.c_char * 16

# TThostFtdcUserNameType 用户名称类型
TThostFtdcUserNameType = ctypes.c_char * 81

# TThostFtdcUserProductIDType 产品标识类型
TThostFtdcUserProductIDType = ctypes.c_char * 33

# TThostFtdcUserProductMemoType 产品说明类型
TThostFtdcUserProductMemoType = ctypes.c_char * 129

# TThostFtdcUserProductNameType 产品名称类型
TThostFtdcUserProductNameType = ctypes.c_char * 65

# TThostFtdcVersionType 版本号类型
TThostFtdcVersionType = ctypes.c_char * 4

# TThostFtdcWarehouseType 仓库类型
TThostFtdcWarehouseType = ctypes.c_char * 257

# TThostFtdcWebSiteType 网址类型
TThostFtdcWebSiteType = ctypes.c_char * 101

# TThostFtdcWebsiteType 网站地址类型
TThostFtdcWebsiteType = ctypes.c_char * 51

# TThostFtdcWeightType 公定重量类型
TThostFtdcWeightType = ctypes.c_char * 41

# TThostFtdcWithDrawParamValueType 可提控制参数内容类型
TThostFtdcWithDrawParamValueType = ctypes.c_char * 41

# TThostFtdcWorkPlaceType 工作单位类型
TThostFtdcWorkPlaceType = ctypes.c_char * 101

# TThostFtdcYieldlyType 产地类型
TThostFtdcYieldlyType = ctypes.c_char * 41

# TThostFtdcZipCodeType 邮政编码类型
TThostFtdcZipCodeType = ctypes.c_char * 7

# ----- 整数类型 -----

# TThostFtdcAMLFileAmountType 反洗钱资金类型
TThostFtdcAMLFileAmountType = ctypes.c_int32

# TThostFtdcAdditionalInfoLenType 补充信息长度类型
TThostFtdcAdditionalInfoLenType = ctypes.c_int32

# TThostFtdcApplicationIDType 应用标识类型
TThostFtdcApplicationIDType = ctypes.c_int32

# TThostFtdcBankProxyIDType 银行代理标识类型
TThostFtdcBankProxyIDType = ctypes.c_int32

# TThostFtdcBoolType 布尔型类型
TThostFtdcBoolType = ctypes.c_int32

# TThostFtdcBulletinIDType 公告编号类型
TThostFtdcBulletinIDType = ctypes.c_int32

# TThostFtdcCaptchaInfoLenType 图片验证信息长度类型
TThostFtdcCaptchaInfoLenType = ctypes.c_int32

# TThostFtdcCheckNoType 操作次数类型
TThostFtdcCheckNoType = ctypes.c_int32

# TThostFtdcComTypeType 组合成交类型
TThostFtdcComTypeType = ctypes.c_int32

# TThostFtdcCommApiPointerType 通讯API指针类型
TThostFtdcCommApiPointerType = ctypes.c_int32

# TThostFtdcCommandNoType DB命令序号类型
TThostFtdcCommandNoType = ctypes.c_int32

# TThostFtdcCommodityGroupIDType 商品群号类型
TThostFtdcCommodityGroupIDType = ctypes.c_int32

# TThostFtdcCommonIntType 通用int类型
TThostFtdcCommonIntType = ctypes.c_int32

# TThostFtdcCorrectSerialType 被冲正交易流水号类型
TThostFtdcCorrectSerialType = ctypes.c_int32

# TThostFtdcCurrentAuthMethodType 当前可用的认证模式，0代表无需认证模式 A从低位开始最后一位代表图片验证码，倒数第二位代表动态口令，倒数第三位代表短信验证码类型
TThostFtdcCurrentAuthMethodType = ctypes.c_int32

# TThostFtdcDBOPSeqNoType 递增的序列号类型
TThostFtdcDBOPSeqNoType = ctypes.c_int32

# TThostFtdcDCEPriorityType 优先级类型
TThostFtdcDCEPriorityType = ctypes.c_int32

# TThostFtdcDRIdentityIDType 交易中心代码类型
TThostFtdcDRIdentityIDType = ctypes.c_int32

# TThostFtdcDataCenterIDType 数据中心代码类型
TThostFtdcDataCenterIDType = ctypes.c_int32

# TThostFtdcErrorIDType 错误代码类型
TThostFtdcErrorIDType = ctypes.c_int32

# TThostFtdcExReturnCodeType 交易所返回码类型
TThostFtdcExReturnCodeType = ctypes.c_int32

# TThostFtdcFBETotalExCntType 换汇交易总笔数类型
TThostFtdcFBETotalExCntType = ctypes.c_int32

# TThostFtdcFBTCoreIDType 银期转帐核心系统标识类型
TThostFtdcFBTCoreIDType = ctypes.c_int32

# TThostFtdcFBTRequestIDType 请求ID类型
TThostFtdcFBTRequestIDType = ctypes.c_int32

# TThostFtdcFrontIDType 前置编号类型
TThostFtdcFrontIDType = ctypes.c_int32

# TThostFtdcFutureSerialType 期货公司流水号类型
TThostFtdcFutureSerialType = ctypes.c_int32

# TThostFtdcHandshakeDataLenType 握手数据内容长度类型
TThostFtdcHandshakeDataLenType = ctypes.c_int32

# TThostFtdcIPPortType IP端口类型
TThostFtdcIPPortType = ctypes.c_int32

# TThostFtdcImplyLevelType 派生层数类型
TThostFtdcImplyLevelType = ctypes.c_int32

# TThostFtdcInstallCountType 安装数量类型
TThostFtdcInstallCountType = ctypes.c_int32

# TThostFtdcInstallIDType 安装编号类型
TThostFtdcInstallIDType = ctypes.c_int32

# TThostFtdcIsCheckPrepaType 是否校验开户可用资金类型
TThostFtdcIsCheckPrepaType = ctypes.c_int32

# TThostFtdcLastDriftType 上次OTP漂移值类型
TThostFtdcLastDriftType = ctypes.c_int32

# TThostFtdcLastSuccessType 上次OTP成功值类型
TThostFtdcLastSuccessType = ctypes.c_int32

# TThostFtdcLegIDType 单腿编号类型
TThostFtdcLegIDType = ctypes.c_int32

# TThostFtdcLegMultipleType 单腿乘数类型
TThostFtdcLegMultipleType = ctypes.c_int32

# TThostFtdcMillisecType 时间（毫秒）类型
TThostFtdcMillisecType = ctypes.c_int32

# TThostFtdcMonthCountType 月份数量类型
TThostFtdcMonthCountType = ctypes.c_int32

# TThostFtdcMonthType 月份类型
TThostFtdcMonthType = ctypes.c_int32

# TThostFtdcOrderActionRefType 报单操作引用类型
TThostFtdcOrderActionRefType = ctypes.c_int32

# TThostFtdcParamIDType 参数代码类型
TThostFtdcParamIDType = ctypes.c_int32

# TThostFtdcPlateSerialType 平台流水号类型
TThostFtdcPlateSerialType = ctypes.c_int32

# TThostFtdcPortfolioDefIDType SPBM组合套餐ID类型
TThostFtdcPortfolioDefIDType = ctypes.c_int32

# TThostFtdcPriorityType 优先级类型
TThostFtdcPriorityType = ctypes.c_int32

# TThostFtdcQueryDepthType 查询深度类型
TThostFtdcQueryDepthType = ctypes.c_int32

# TThostFtdcQueryFreqType 查询频率类型
TThostFtdcQueryFreqType = ctypes.c_int32

# TThostFtdcRCAMSPriorityType 优先级类型
TThostFtdcRCAMSPriorityType = ctypes.c_int32

# TThostFtdcRecordCountType 记录数类型
TThostFtdcRecordCountType = ctypes.c_int32

# TThostFtdcRepealTimeIntervalType 冲正时间间隔类型
TThostFtdcRepealTimeIntervalType = ctypes.c_int32

# TThostFtdcRepealedTimesType 已经冲正次数类型
TThostFtdcRepealedTimesType = ctypes.c_int32

# TThostFtdcRequestIDType 请求编号类型
TThostFtdcRequestIDType = ctypes.c_int32

# TThostFtdcRsaKeyVersionType 公钥版本号类型
TThostFtdcRsaKeyVersionType = ctypes.c_int32

# TThostFtdcSecType 时间（秒）类型
TThostFtdcSecType = ctypes.c_int32

# TThostFtdcSeqNoType 流水号类型
TThostFtdcSeqNoType = ctypes.c_int32

# TThostFtdcSequenceNo12Type 序号类型
TThostFtdcSequenceNo12Type = ctypes.c_int32

# TThostFtdcSequenceNoType 序号类型
TThostFtdcSequenceNoType = ctypes.c_int32

# TThostFtdcSerialType 流水号类型
TThostFtdcSerialType = ctypes.c_int32

# TThostFtdcServerPortType 服务端口号类型
TThostFtdcServerPortType = ctypes.c_int32

# TThostFtdcServiceIDType 服务编号类型
TThostFtdcServiceIDType = ctypes.c_int32

# TThostFtdcServiceLineNoType 服务线路编号类型
TThostFtdcServiceLineNoType = ctypes.c_int32

# TThostFtdcSessionIDType 会话编号类型
TThostFtdcSessionIDType = ctypes.c_int32

# TThostFtdcSettlementIDType 结算编号类型
TThostFtdcSettlementIDType = ctypes.c_int32

# TThostFtdcSpreadIdType 抵扣组优先级类型
TThostFtdcSpreadIdType = ctypes.c_int32

# TThostFtdcStrikeSequenceType 执行序号类型
TThostFtdcStrikeSequenceType = ctypes.c_int32

# TThostFtdcSubEntryFundNoType 分项资金流水号类型
TThostFtdcSubEntryFundNoType = ctypes.c_int32

# TThostFtdcSystemInfoLenType 系统信息长度类型
TThostFtdcSystemInfoLenType = ctypes.c_int32

# TThostFtdcTIDType 交易ID类型
TThostFtdcTIDType = ctypes.c_int32

# TThostFtdcThostFunctionCodeType Thost终端功能代码类型
TThostFtdcThostFunctionCodeType = ctypes.c_int32

# TThostFtdcTimestampType 时间戳类型
TThostFtdcTimestampType = ctypes.c_int32

# TThostFtdcTopicIDType 主题代码类型
TThostFtdcTopicIDType = ctypes.c_int32

# TThostFtdcTotalTimesType 每日累计转帐次数类型
TThostFtdcTotalTimesType = ctypes.c_int32

# TThostFtdcTradeGroupIDType 成交组号类型
TThostFtdcTradeGroupIDType = ctypes.c_int32

# TThostFtdcTradeSerialNoType 发起方流水号类型
TThostFtdcTradeSerialNoType = ctypes.c_int32

# TThostFtdcTradingSegmentSNType 交易阶段编号类型
TThostFtdcTradingSegmentSNType = ctypes.c_int32

# TThostFtdcUserTextSeqType 用户短信验证码的编号类型
TThostFtdcUserTextSeqType = ctypes.c_int32

# TThostFtdcVolumeMultipleType 合约数量乘数类型
TThostFtdcVolumeMultipleType = ctypes.c_int32

# TThostFtdcVolumeType 数量类型
TThostFtdcVolumeType = ctypes.c_int32

# TThostFtdcYearType 年份类型
TThostFtdcYearType = ctypes.c_int32

# ----- 短整数类型 -----

# TThostFtdcCommPhaseNoType 通讯时段编号类型
TThostFtdcCommPhaseNoType = ctypes.c_int16

# TThostFtdcSequenceSeriesType 序列系列号类型
TThostFtdcSequenceSeriesType = ctypes.c_int16

# ----- 浮点类型 -----

# TThostFtdcAMLMoneyType 反洗钱资金类型
TThostFtdcAMLMoneyType = ctypes.c_double

# TThostFtdcAMLOpParamValueType 业务参数代码值类型
TThostFtdcAMLOpParamValueType = ctypes.c_double

# TThostFtdcAdjustValueType 空头期权风险调整标准类型
TThostFtdcAdjustValueType = ctypes.c_double

# TThostFtdcBigMoneyType 资金类型
TThostFtdcBigMoneyType = ctypes.c_double

# TThostFtdcCSRCMoneyType 资金类型
TThostFtdcCSRCMoneyType = ctypes.c_double

# TThostFtdcCSRCPriceType 价格类型
TThostFtdcCSRCPriceType = ctypes.c_double

# TThostFtdcCSRCStrikePriceType 执行价类型
TThostFtdcCSRCStrikePriceType = ctypes.c_double

# TThostFtdcCurrencyUnitType 币种单位数量类型
TThostFtdcCurrencyUnitType = ctypes.c_double

# TThostFtdcCustFeeType 应收客户费用（元）类型
TThostFtdcCustFeeType = ctypes.c_double

# TThostFtdcDeltaType Delta类型
TThostFtdcDeltaType = ctypes.c_double

# TThostFtdcDiscountRatioType 折扣率类型
TThostFtdcDiscountRatioType = ctypes.c_double

# TThostFtdcExRateType 换汇汇率类型
TThostFtdcExRateType = ctypes.c_double

# TThostFtdcExchangeRateType 汇率类型
TThostFtdcExchangeRateType = ctypes.c_double

# TThostFtdcFBEAmtType 各种换汇金额类型
TThostFtdcFBEAmtType = ctypes.c_double

# TThostFtdcFutureFeeType 应收期货公司费用（元）类型
TThostFtdcFutureFeeType = ctypes.c_double

# TThostFtdcHedgeRateType HedgeRate类型
TThostFtdcHedgeRateType = ctypes.c_double

# TThostFtdcLargeVolumeType 大额数量类型
TThostFtdcLargeVolumeType = ctypes.c_double

# TThostFtdcMoneyType 资金类型
TThostFtdcMoneyType = ctypes.c_double

# TThostFtdcPriceType 价格类型
TThostFtdcPriceType = ctypes.c_double

# TThostFtdcRatioType 比率类型
TThostFtdcRatioType = ctypes.c_double

# TThostFtdcRiskValueType 期货风险值类型
TThostFtdcRiskValueType = ctypes.c_double

# TThostFtdcSPMMDiscountRatioType SPMM折扣率类型
TThostFtdcSPMMDiscountRatioType = ctypes.c_double

# TThostFtdcSingleMaxAmtType 单笔最高限额类型
TThostFtdcSingleMaxAmtType = ctypes.c_double

# TThostFtdcSingleMinAmtType 单笔最低限额类型
TThostFtdcSingleMinAmtType = ctypes.c_double

# TThostFtdcStdPositionType 标准持仓类型
TThostFtdcStdPositionType = ctypes.c_double

# TThostFtdcTotalAmtType 每日累计转帐额度类型
TThostFtdcTotalAmtType = ctypes.c_double

# TThostFtdcTradeAmountType 交易金额（元）类型
TThostFtdcTradeAmountType = ctypes.c_double

# TThostFtdcUnderlyingMultipleType 基础商品乘数类型
TThostFtdcUnderlyingMultipleType = ctypes.c_double

# ----- 字符枚举类型 -----

TThostFtdcAMLCheckStatusType = ctypes.c_char

THOST_FTDC_AMLCHS_Init = b'0'  # 未复核
THOST_FTDC_AMLCHS_Checking = b'1'  # 复核中
THOST_FTDC_AMLCHS_Checked = b'2'  # 已复核
THOST_FTDC_AMLCHS_RefuseReport = b'3'  # 拒绝上报

TThostFtdcAMLGenStatusType = ctypes.c_char

THOST_FTDC_GEN_Program = b'0'  # 程序生成
THOST_FTDC_GEN_HandWork = b'1'  # 人工生成

TThostFtdcAPIProductClassType = ctypes.c_char

THOST_FTDC_APC_FutureSingle = b'1'  # 期货单一合约
THOST_FTDC_APC_OptionSingle = b'2'  # 期权单一合约
THOST_FTDC_APC_Futures = b'3'  # 可交易期货(含期货组合和期货单一合约)
THOST_FTDC_APC_Options = b'4'  # 可交易期权(含期权组合和期权单一合约)
THOST_FTDC_APC_TradingComb = b'5'  # 可下单套利组合
THOST_FTDC_APC_UnTradingComb = b'6'  # 可申请的组合（可以申请的组合合约 包含可以交易的合约）
THOST_FTDC_APC_AllTrading = b'7'  # 所有可以交易合约
THOST_FTDC_APC_All = b'8'  # 所有合约（包含不能交易合约 慎用）

TThostFtdcAccountSettlementParamIDType = ctypes.c_char

THOST_FTDC_ASPI_BaseMargin = b'1'  # 基础保证金
THOST_FTDC_ASPI_LowestInterest = b'2'  # 最低权益标准

TThostFtdcAccountSourceTypeType = ctypes.c_char

THOST_FTDC_AST_FBTransfer = b'0'  # 银期同步
THOST_FTDC_AST_ManualEntry = b'1'  # 手工录入

TThostFtdcActionDirectionType = ctypes.c_char

THOST_FTDC_ACD_Add = b'1'  # 增加
THOST_FTDC_ACD_Del = b'2'  # 删除
THOST_FTDC_ACD_Upd = b'3'  # 更新

TThostFtdcActionFlagType = ctypes.c_char

THOST_FTDC_AF_Delete = b'0'  # 删除
THOST_FTDC_AF_Modify = b'3'  # 修改

TThostFtdcActionTypeType = ctypes.c_char

THOST_FTDC_ACTP_Exec = b'1'  # 执行
THOST_FTDC_ACTP_Abandon = b'2'  # 放弃

TThostFtdcActiveTypeType = ctypes.c_char

THOST_FTDC_ACT_Intraday = b'1'  # 仅当日生效
THOST_FTDC_ACT_Long = b'2'  # 长期生效

TThostFtdcAlgoTypeType = ctypes.c_char

THOST_FTDC_AT_HandlePositionAlgo = b'1'  # 持仓处理算法
THOST_FTDC_AT_FindMarginRateAlgo = b'2'  # 寻找保证金率算法

TThostFtdcAlgorithmType = ctypes.c_char

THOST_FTDC_AG_All = b'1'  # 浮盈浮亏都计算
THOST_FTDC_AG_OnlyLost = b'2'  # 浮盈不计，浮亏计
THOST_FTDC_AG_OnlyGain = b'3'  # 浮盈计，浮亏不计
THOST_FTDC_AG_None = b'4'  # 浮盈浮亏都不计算

TThostFtdcAllWithoutTradeType = ctypes.c_char

THOST_FTDC_AWT_Enable = b'0'  # 无仓无成交不受可提比例限制
THOST_FTDC_AWT_Disable = b'2'  # 受可提比例限制
THOST_FTDC_AWT_NoHoldEnable = b'3'  # 无仓不受可提比例限制

TThostFtdcAmTypeType = ctypes.c_char

THOST_FTDC_AMT_Bank = b'1'  # 银行
THOST_FTDC_AMT_Securities = b'2'  # 证券公司
THOST_FTDC_AMT_Fund = b'3'  # 基金公司
THOST_FTDC_AMT_Insurance = b'4'  # 保险公司
THOST_FTDC_AMT_Trust = b'5'  # 信托公司
THOST_FTDC_AMT_Other = b'9'  # 其他

TThostFtdcAmlCheckLevelType = ctypes.c_char

THOST_FTDC_AMLCL_CheckLevel0 = b'0'  # 零级审核
THOST_FTDC_AMLCL_CheckLevel1 = b'1'  # 一级审核
THOST_FTDC_AMLCL_CheckLevel2 = b'2'  # 二级审核
THOST_FTDC_AMLCL_CheckLevel3 = b'3'  # 三级审核

TThostFtdcAmlDateTypeType = ctypes.c_char

THOST_FTDC_AMLDT_DrawDay = b'0'  # 检查日期
THOST_FTDC_AMLDT_TouchDay = b'1'  # 发生日期

TThostFtdcAppTypeType = ctypes.c_char

THOST_FTDC_APP_TYPE_Investor = b'1'  # 直连的投资者
THOST_FTDC_APP_TYPE_InvestorRelay = b'2'  # 为每个投资者都创建连接的中继
THOST_FTDC_APP_TYPE_OperatorRelay = b'3'  # 所有投资者共享一个操作员连接的中继
THOST_FTDC_APP_TYPE_UnKnown = b'4'  # 未知

TThostFtdcApplyOperateIDType = ctypes.c_char

THOST_FTDC_AOID_OpenInvestor = b'1'  # 开户
THOST_FTDC_AOID_ModifyIDCard = b'2'  # 修改身份信息
THOST_FTDC_AOID_ModifyNoIDCard = b'3'  # 修改一般信息
THOST_FTDC_AOID_ApplyTradingCode = b'4'  # 申请交易编码
THOST_FTDC_AOID_CancelTradingCode = b'5'  # 撤销交易编码
THOST_FTDC_AOID_CancelInvestor = b'6'  # 销户
THOST_FTDC_AOID_FreezeAccount = b'8'  # 账户休眠
THOST_FTDC_AOID_ActiveFreezeAccount = b'9'  # 激活休眠账户

TThostFtdcApplyStatusIDType = ctypes.c_char

THOST_FTDC_ASID_NoComplete = b'1'  # 未补全
THOST_FTDC_ASID_Submited = b'2'  # 已提交
THOST_FTDC_ASID_Checked = b'3'  # 已审核
THOST_FTDC_ASID_Refused = b'4'  # 已拒绝
THOST_FTDC_ASID_Deleted = b'5'  # 已删除

TThostFtdcApplyTypeType = ctypes.c_char

THOST_FTDC_APPT_NotStrikeNum = b'4'  # 不执行数量

TThostFtdcAssetmgrClientTypeType = ctypes.c_char

THOST_FTDC_AMCT_Person = b'1'  # 个人资管客户
THOST_FTDC_AMCT_Organ = b'2'  # 单位资管客户
THOST_FTDC_AMCT_SpecialOrgan = b'4'  # 特殊单位资管客户

TThostFtdcAssetmgrTypeType = ctypes.c_char

THOST_FTDC_ASST_Futures = b'3'  # 期货类
THOST_FTDC_ASST_SpecialOrgan = b'4'  # 综合类

TThostFtdcAuthTypeType = ctypes.c_char

THOST_FTDC_AU_WHITE = b'0'  # 白名单校验
THOST_FTDC_AU_BLACK = b'1'  # 黑名单校验

TThostFtdcAvailabilityFlagType = ctypes.c_char

THOST_FTDC_AVAF_Invalid = b'0'  # 未确认
THOST_FTDC_AVAF_Valid = b'1'  # 有效
THOST_FTDC_AVAF_Repeal = b'2'  # 冲正

TThostFtdcBackUpStatusType = ctypes.c_char

THOST_FTDC_BUS_UnBak = b'0'  # 未生成备份数据
THOST_FTDC_BUS_BakUp = b'1'  # 备份数据生成中
THOST_FTDC_BUS_BakUped = b'2'  # 已生成备份数据
THOST_FTDC_BUS_BakFail = b'3'  # 备份数据失败

TThostFtdcBalanceAlgorithmType = ctypes.c_char

THOST_FTDC_BLAG_Default = b'1'  # 不计算期权市值盈亏
THOST_FTDC_BLAG_IncludeOptValLost = b'2'  # 计算期权市值亏损

TThostFtdcBankAccStatusType = ctypes.c_char

THOST_FTDC_BAS_Normal = b'0'  # 正常
THOST_FTDC_BAS_Freeze = b'1'  # 冻结
THOST_FTDC_BAS_ReportLoss = b'2'  # 挂失

TThostFtdcBankAccTypeType = ctypes.c_char

THOST_FTDC_BAT_BankBook = b'1'  # 银行存折
THOST_FTDC_BAT_SavingCard = b'2'  # 储蓄卡
THOST_FTDC_BAT_CreditCard = b'3'  # 信用卡

TThostFtdcBankAcountOriginType = ctypes.c_char

THOST_FTDC_BAO_ByAccProperty = b'0'  # 手工录入
THOST_FTDC_BAO_ByFBTransfer = b'1'  # 银期转账

TThostFtdcBankRepealFlagType = ctypes.c_char

THOST_FTDC_BRF_BankNotNeedRepeal = b'0'  # 银行无需自动冲正
THOST_FTDC_BRF_BankWaitingRepeal = b'1'  # 银行待自动冲正
THOST_FTDC_BRF_BankBeenRepealed = b'2'  # 银行已自动冲正

TThostFtdcBanlanceTypeType = ctypes.c_char

THOST_FTDC_BLT_CurrentMoney = b'0'  # 当前余额
THOST_FTDC_BLT_UsableMoney = b'1'  # 可用余额
THOST_FTDC_BLT_FetchableMoney = b'2'  # 可取余额
THOST_FTDC_BLT_FreezeMoney = b'3'  # 冻结余额

TThostFtdcBasisPriceTypeType = ctypes.c_char

THOST_FTDC_IPT_LastSettlement = b'1'  # 上一合约结算价
THOST_FTDC_IPT_LaseClose = b'2'  # 上一合约收盘价

TThostFtdcBatchStatusType = ctypes.c_char

THOST_FTDC_BS_NoUpload = b'1'  # 未上传
THOST_FTDC_BS_Uploaded = b'2'  # 已上传
THOST_FTDC_BS_Failed = b'3'  # 审核失败

TThostFtdcBillGenStatusType = ctypes.c_char

THOST_FTDC_BGS_None = b'0'  # 未生成
THOST_FTDC_BGS_NoGenerated = b'1'  # 生成中
THOST_FTDC_BGS_Generated = b'2'  # 已生成

TThostFtdcBillHedgeFlagType = ctypes.c_char

THOST_FTDC_BHF_Speculation = b'1'  # 投机
THOST_FTDC_BHF_Arbitrage = b'2'  # 套利
THOST_FTDC_BHF_Hedge = b'3'  # 套保

TThostFtdcBizTypeType = ctypes.c_char

THOST_FTDC_BZTP_Future = b'1'  # 期货
THOST_FTDC_BZTP_Stock = b'2'  # 证券

TThostFtdcBrokerDataSyncStatusType = ctypes.c_char

THOST_FTDC_BDS_Synchronized = b'1'  # 已同步
THOST_FTDC_BDS_Synchronizing = b'2'  # 同步中

TThostFtdcBrokerFunctionCodeType = ctypes.c_char

THOST_FTDC_BFC_ForceUserLogout = b'1'  # 强制用户登出
THOST_FTDC_BFC_UserPasswordUpdate = b'2'  # 变更用户口令
THOST_FTDC_BFC_SyncBrokerData = b'3'  # 同步经纪公司数据
THOST_FTDC_BFC_BachSyncBrokerData = b'4'  # 批量同步经纪公司数据
THOST_FTDC_BFC_OrderInsert = b'5'  # 报单插入
THOST_FTDC_BFC_OrderAction = b'6'  # 报单操作
THOST_FTDC_BFC_AllQuery = b'7'  # 全部查询
THOST_FTDC_BFC_log = b'a'  # 系统功能：登入/登出/修改密码等
THOST_FTDC_BFC_BaseQry = b'b'  # 基本查询：查询基础数据，如合约，交易所等常量
THOST_FTDC_BFC_TradeQry = b'c'  # 交易查询：如查成交，委托
THOST_FTDC_BFC_Trade = b'd'  # 交易功能：报单，撤单
THOST_FTDC_BFC_Virement = b'e'  # 银期转账
THOST_FTDC_BFC_Risk = b'f'  # 风险监控
THOST_FTDC_BFC_Session = b'g'  # 查询/管理：查询会话，踢人等
THOST_FTDC_BFC_RiskNoticeCtl = b'h'  # 风控通知控制
THOST_FTDC_BFC_RiskNotice = b'i'  # 风控通知发送
THOST_FTDC_BFC_BrokerDeposit = b'j'  # 察看经纪公司资金权限
THOST_FTDC_BFC_QueryFund = b'k'  # 资金查询
THOST_FTDC_BFC_QueryOrder = b'l'  # 报单查询
THOST_FTDC_BFC_QueryTrade = b'm'  # 成交查询
THOST_FTDC_BFC_QueryPosition = b'n'  # 持仓查询
THOST_FTDC_BFC_QueryMarketData = b'o'  # 行情查询
THOST_FTDC_BFC_QueryUserEvent = b'p'  # 用户事件查询
THOST_FTDC_BFC_QueryRiskNotify = b'q'  # 风险通知查询
THOST_FTDC_BFC_QueryFundChange = b'r'  # 出入金查询
THOST_FTDC_BFC_QueryInvestor = b's'  # 投资者信息查询
THOST_FTDC_BFC_QueryTradingCode = b't'  # 交易编码查询
THOST_FTDC_BFC_ForceClose = b'u'  # 强平
THOST_FTDC_BFC_PressTest = b'v'  # 压力测试
THOST_FTDC_BFC_RemainCalc = b'w'  # 权益反算
THOST_FTDC_BFC_NetPositionInd = b'x'  # 净持仓保证金指标
THOST_FTDC_BFC_RiskPredict = b'y'  # 风险预算
THOST_FTDC_BFC_DataExport = b'z'  # 数据导出
THOST_FTDC_BFC_RiskTargetSetup = b'A'  # 风控指标设置
THOST_FTDC_BFC_MarketDataWarn = b'B'  # 行情预警
THOST_FTDC_BFC_QryBizNotice = b'C'  # 业务通知查询
THOST_FTDC_BFC_CfgBizNotice = b'D'  # 业务通知模板设置
THOST_FTDC_BFC_SyncOTP = b'E'  # 同步动态令牌
THOST_FTDC_BFC_SendBizNotice = b'F'  # 发送业务通知
THOST_FTDC_BFC_CfgRiskLevelStd = b'G'  # 风险级别标准设置
THOST_FTDC_BFC_TbCommand = b'H'  # 交易终端应急功能
THOST_FTDC_BFC_DeleteOrder = b'J'  # 删除未知单
THOST_FTDC_BFC_ParkedOrderInsert = b'K'  # 预埋报单插入
THOST_FTDC_BFC_ParkedOrderAction = b'L'  # 预埋报单操作
THOST_FTDC_BFC_ExecOrderNoCheck = b'M'  # 资金不够仍允许行权
THOST_FTDC_BFC_Designate = b'N'  # 指定
THOST_FTDC_BFC_StockDisposal = b'O'  # 证券处置
THOST_FTDC_BFC_BrokerDepositWarn = b'Q'  # 席位资金预警
THOST_FTDC_BFC_CoverWarn = b'S'  # 备兑不足预警
THOST_FTDC_BFC_PreExecOrder = b'T'  # 行权试算
THOST_FTDC_BFC_ExecOrderRisk = b'P'  # 行权交收风险
THOST_FTDC_BFC_PosiLimitWarn = b'U'  # 持仓限额预警
THOST_FTDC_BFC_QryPosiLimit = b'V'  # 持仓限额查询
THOST_FTDC_BFC_FBSign = b'W'  # 银期签到签退
THOST_FTDC_BFC_FBAccount = b'X'  # 银期签约解约

TThostFtdcBrokerRepealFlagType = ctypes.c_char

THOST_FTDC_BRORF_BrokerNotNeedRepeal = b'0'  # 期商无需自动冲正
THOST_FTDC_BRORF_BrokerWaitingRepeal = b'1'  # 期商待自动冲正
THOST_FTDC_BRORF_BrokerBeenRepealed = b'2'  # 期商已自动冲正

TThostFtdcBrokerTypeType = ctypes.c_char

THOST_FTDC_BT_Trade = b'0'  # 交易会员
THOST_FTDC_BT_TradeSettle = b'1'  # 交易结算会员

TThostFtdcBrokerUserTypeType = ctypes.c_char

THOST_FTDC_BUT_Investor = b'1'  # 投资者
THOST_FTDC_BUT_BrokerUser = b'2'  # 操作员

TThostFtdcBusinessClassType = ctypes.c_char

THOST_FTDC_BT_Profit = b'0'  # 盈利
THOST_FTDC_BT_Loss = b'1'  # 亏损
THOST_FTDC_BT_Other = b'Z'  # 其他

TThostFtdcBusinessTypeType = ctypes.c_char

THOST_FTDC_BT_Request = b'1'  # 请求
THOST_FTDC_BT_Response = b'2'  # 应答
THOST_FTDC_BT_Notice = b'3'  # 通知

TThostFtdcByGroupType = ctypes.c_char

THOST_FTDC_BG_Investor = b'2'  # 按投资者统计
THOST_FTDC_BG_Group = b'1'  # 按类统计

TThostFtdcByInvestorRangeType = ctypes.c_char

THOST_FTDC_BIR_Property = b'1'  # 属性统计
THOST_FTDC_BIR_All = b'2'  # 统计所有

TThostFtdcCCBFeeModeType = ctypes.c_char

THOST_FTDC_CCBFM_ByAmount = b'1'  # 按金额扣收
THOST_FTDC_CCBFM_ByMonth = b'2'  # 按月扣收

TThostFtdcCFFEXUploadFileNameType = ctypes.c_char

THOST_FTDC_CFUFN_SUFN_T = b'T'  # ^\d{4}_SG\d{1}_\d{8}_\d{1}_Trade
THOST_FTDC_CFUFN_SUFN_P = b'P'  # ^\d{4}_SG\d{1}_\d{8}_\d{1}_SettlementDetail
THOST_FTDC_CFUFN_SUFN_F = b'F'  # ^\d{4}_SG\d{1}_\d{8}_\d{1}_Capital
THOST_FTDC_CFUFN_SUFN_S = b'S'  # ^\d{4}_SG\d{1}_\d{8}_\d{1}_OptionExec

TThostFtdcCFMMCKeyKindType = ctypes.c_char

THOST_FTDC_CFMMCKK_REQUEST = b'R'  # 主动请求更新
THOST_FTDC_CFMMCKK_AUTO = b'A'  # CFMMC自动更新
THOST_FTDC_CFMMCKK_MANUAL = b'M'  # CFMMC手动更新

TThostFtdcCSRCDataQueyTypeType = ctypes.c_char

THOST_FTDC_CSRCQ_Current = b'0'  # 查询当前交易日报送的数据
THOST_FTDC_CSRCQ_History = b'1'  # 查询历史报送的代理经纪公司的数据

TThostFtdcCSRCFundIOTypeType = ctypes.c_char

THOST_FTDC_CFIOT_FundIO = b'0'  # 出入金
THOST_FTDC_CFIOT_SwapCurrency = b'1'  # 银期换汇

TThostFtdcCTPTypeType = ctypes.c_char

THOST_FTDC_CTPT_Unkown = b'0'  # 未知类型
THOST_FTDC_CTPT_MainCenter = b'1'  # 主中心
THOST_FTDC_CTPT_BackUp = b'2'  # 备中心

TThostFtdcCZCEUploadFileNameType = ctypes.c_char

THOST_FTDC_CUFN_CUFN_O = b'O'  # ^\d{8}_zz_\d{4}
THOST_FTDC_CUFN_CUFN_T = b'T'  # ^\d{8}成交表
THOST_FTDC_CUFN_CUFN_P = b'P'  # ^\d{8}单腿持仓表new
THOST_FTDC_CUFN_CUFN_N = b'N'  # ^\d{8}非平仓了结表
THOST_FTDC_CUFN_CUFN_L = b'L'  # ^\d{8}平仓表
THOST_FTDC_CUFN_CUFN_F = b'F'  # ^\d{8}资金表
THOST_FTDC_CUFN_CUFN_C = b'C'  # ^\d{8}组合持仓表
THOST_FTDC_CUFN_CUFN_M = b'M'  # ^\d{8}保证金参数表

TThostFtdcCashExchangeCodeType = ctypes.c_char

THOST_FTDC_CEC_Exchange = b'1'  # 汇
THOST_FTDC_CEC_Cash = b'2'  # 钞

TThostFtdcCertificationTypeType = ctypes.c_char

THOST_FTDC_CFT_IDCard = b'0'  # 身份证
THOST_FTDC_CFT_Passport = b'1'  # 护照
THOST_FTDC_CFT_OfficerIDCard = b'2'  # 军官证
THOST_FTDC_CFT_SoldierIDCard = b'3'  # 士兵证
THOST_FTDC_CFT_HomeComingCard = b'4'  # 回乡证
THOST_FTDC_CFT_HouseholdRegister = b'5'  # 户口簿
THOST_FTDC_CFT_LicenseNo = b'6'  # 营业执照号
THOST_FTDC_CFT_InstitutionCodeCard = b'7'  # 组织机构代码证
THOST_FTDC_CFT_TempLicenseNo = b'8'  # 临时营业执照号
THOST_FTDC_CFT_NoEnterpriseLicenseNo = b'9'  # 民办非企业登记证书
THOST_FTDC_CFT_OtherCard = b'x'  # 其他证件
THOST_FTDC_CFT_SuperDepAgree = b'a'  # 主管部门批文

TThostFtdcCfmmcReturnCodeType = ctypes.c_char

THOST_FTDC_CRC_Success = b'0'  # 成功
THOST_FTDC_CRC_Working = b'1'  # 该客户已经有流程在处理中
THOST_FTDC_CRC_InfoFail = b'2'  # 监控中客户资料检查失败
THOST_FTDC_CRC_IDCardFail = b'3'  # 监控中实名制检查失败
THOST_FTDC_CRC_OtherFail = b'4'  # 其他错误

TThostFtdcCheckInstrTypeType = ctypes.c_char

THOST_FTDC_CIT_HasExch = b'0'  # 合约交易所不存在
THOST_FTDC_CIT_HasATP = b'1'  # 合约本系统不存在
THOST_FTDC_CIT_HasDiff = b'2'  # 合约比较不一致

TThostFtdcCheckLevelType = ctypes.c_char

THOST_FTDC_CL_Zero = b'0'  # 零级复核
THOST_FTDC_CL_One = b'1'  # 一级复核
THOST_FTDC_CL_Two = b'2'  # 二级复核

TThostFtdcCheckStatusType = ctypes.c_char

THOST_FTDC_CHS_Init = b'0'  # 未复核
THOST_FTDC_CHS_Checking = b'1'  # 复核中
THOST_FTDC_CHS_Checked = b'2'  # 已复核
THOST_FTDC_CHS_Refuse = b'3'  # 拒绝
THOST_FTDC_CHS_Cancel = b'4'  # 作废

TThostFtdcClassTypeType = ctypes.c_char

THOST_FTDC_INS_ALL = b'0'  # 所有合约
THOST_FTDC_INS_FUTURE = b'1'  # 期货、即期、期转现、Tas、金属指数合约
THOST_FTDC_INS_OPTION = b'2'  # 期货、现货期权合约
THOST_FTDC_INS_COMB = b'3'  # 组合合约

TThostFtdcClientIDStatusType = ctypes.c_char

THOST_FTDC_UOACS_NoApply = b'1'  # 未申请
THOST_FTDC_UOACS_Submited = b'2'  # 已提交申请
THOST_FTDC_UOACS_Sended = b'3'  # 已发送申请
THOST_FTDC_UOACS_Success = b'4'  # 完成
THOST_FTDC_UOACS_Refuse = b'5'  # 拒绝
THOST_FTDC_UOACS_Cancel = b'6'  # 已撤销编码

TThostFtdcClientIDTypeType = ctypes.c_char

THOST_FTDC_CIDT_Speculation = b'1'  # 投机
THOST_FTDC_CIDT_Arbitrage = b'2'  # 套利
THOST_FTDC_CIDT_Hedge = b'3'  # 套保
THOST_FTDC_CIDT_MarketMaker = b'5'  # 做市商

TThostFtdcClientRegionType = ctypes.c_char

THOST_FTDC_CR_Domestic = b'1'  # 国内客户
THOST_FTDC_CR_GMT = b'2'  # 港澳台客户
THOST_FTDC_CR_Foreign = b'3'  # 国外客户

TThostFtdcClientTypeType = ctypes.c_char

THOST_FTDC_CfMMCCT_All = b'0'  # 所有
THOST_FTDC_CfMMCCT_Person = b'1'  # 个人
THOST_FTDC_CfMMCCT_Company = b'2'  # 单位
THOST_FTDC_CfMMCCT_Other = b'3'  # 其他
THOST_FTDC_CfMMCCT_SpecialOrgan = b'4'  # 特殊法人
THOST_FTDC_CfMMCCT_Asset = b'5'  # 资管户

TThostFtdcCloseDealTypeType = ctypes.c_char

THOST_FTDC_CDT_Normal = b'0'  # 正常
THOST_FTDC_CDT_SpecFirst = b'1'  # 投机平仓优先

TThostFtdcCloseStyleType = ctypes.c_char

THOST_FTDC_ICS_Close = b'0'  # 先开先平
THOST_FTDC_ICS_CloseToday = b'1'  # 先平今再平昨

TThostFtdcCodeSourceTypeType = ctypes.c_char

THOST_FTDC_CST_UnifyAccount = b'0'  # 统一开户(已规范)
THOST_FTDC_CST_ManualEntry = b'1'  # 手工录入(未规范)

TThostFtdcCombDirectionType = ctypes.c_char

THOST_FTDC_CMDR_Comb = b'0'  # 申请组合
THOST_FTDC_CMDR_UnComb = b'1'  # 申请拆分
THOST_FTDC_CMDR_DelComb = b'2'  # 操作员删组合单

TThostFtdcCombinationTypeType = ctypes.c_char

THOST_FTDC_COMBT_Future = b'0'  # 期货组合
THOST_FTDC_COMBT_BUL = b'1'  # 垂直价差BUL
THOST_FTDC_COMBT_BER = b'2'  # 垂直价差BER
THOST_FTDC_COMBT_STD = b'3'  # 跨式组合
THOST_FTDC_COMBT_STG = b'4'  # 宽跨式组合
THOST_FTDC_COMBT_PRT = b'5'  # 备兑组合
THOST_FTDC_COMBT_CAS = b'6'  # 时间价差组合
THOST_FTDC_COMBT_OPL = b'7'  # 期权对锁组合
THOST_FTDC_COMBT_BFO = b'8'  # 买备兑组合
THOST_FTDC_COMBT_BLS = b'9'  # 买入期权垂直价差组合
THOST_FTDC_COMBT_BES = b'a'  # 卖出期权垂直价差组合

TThostFtdcCommApiTypeType = ctypes.c_char

THOST_FTDC_CAPIT_Client = b'1'  # 客户端
THOST_FTDC_CAPIT_Server = b'2'  # 服务端
THOST_FTDC_CAPIT_UserApi = b'3'  # 交易系统的UserApi

TThostFtdcConditionalOrderSortTypeType = ctypes.c_char

THOST_FTDC_COST_LastPriceAsc = b'0'  # 使用最新价升序
THOST_FTDC_COST_LastPriceDesc = b'1'  # 使用最新价降序
THOST_FTDC_COST_AskPriceAsc = b'2'  # 使用卖价升序
THOST_FTDC_COST_AskPriceDesc = b'3'  # 使用卖价降序
THOST_FTDC_COST_BidPriceAsc = b'4'  # 使用买价升序
THOST_FTDC_COST_BidPriceDesc = b'5'  # 使用买价降序

TThostFtdcConnectModeType = ctypes.c_char

THOST_FTDC_CM_ShortConnect = b'0'  # 短连接
THOST_FTDC_CM_LongConnect = b'1'  # 长连接

TThostFtdcContingentConditionType = ctypes.c_char

THOST_FTDC_CC_Immediately = b'1'  # 立即
THOST_FTDC_CC_Touch = b'2'  # 止损
THOST_FTDC_CC_TouchProfit = b'3'  # 止赢
THOST_FTDC_CC_ParkedOrder = b'4'  # 预埋单
THOST_FTDC_CC_LastPriceGreaterThanStopPrice = b'5'  # 最新价大于条件价
THOST_FTDC_CC_LastPriceGreaterEqualStopPrice = b'6'  # 最新价大于等于条件价
THOST_FTDC_CC_LastPriceLesserThanStopPrice = b'7'  # 最新价小于条件价
THOST_FTDC_CC_LastPriceLesserEqualStopPrice = b'8'  # 最新价小于等于条件价
THOST_FTDC_CC_AskPriceGreaterThanStopPrice = b'9'  # 卖一价大于条件价
THOST_FTDC_CC_AskPriceGreaterEqualStopPrice = b'A'  # 卖一价大于等于条件价
THOST_FTDC_CC_AskPriceLesserThanStopPrice = b'B'  # 卖一价小于条件价
THOST_FTDC_CC_AskPriceLesserEqualStopPrice = b'C'  # 卖一价小于等于条件价
THOST_FTDC_CC_BidPriceGreaterThanStopPrice = b'D'  # 买一价大于条件价
THOST_FTDC_CC_BidPriceGreaterEqualStopPrice = b'E'  # 买一价大于等于条件价
THOST_FTDC_CC_BidPriceLesserThanStopPrice = b'F'  # 买一价小于条件价
THOST_FTDC_CC_BidPriceLesserEqualStopPrice = b'H'  # 买一价小于等于条件价

TThostFtdcCurrExDirectionType = ctypes.c_char

THOST_FTDC_CED_Settlement = b'0'  # 结汇
THOST_FTDC_CED_Sale = b'1'  # 售汇

TThostFtdcCurrencySwapStatusType = ctypes.c_char

THOST_FTDC_CSS_Entry = b'1'  # 已录入
THOST_FTDC_CSS_Approve = b'2'  # 已审核
THOST_FTDC_CSS_Refuse = b'3'  # 已拒绝
THOST_FTDC_CSS_Revoke = b'4'  # 已撤销
THOST_FTDC_CSS_Send = b'5'  # 已发送
THOST_FTDC_CSS_Success = b'6'  # 换汇成功
THOST_FTDC_CSS_Failure = b'7'  # 换汇失败

TThostFtdcCusAccountTypeType = ctypes.c_char

THOST_FTDC_CAT_Futures = b'1'  # 期货结算账户
THOST_FTDC_CAT_AssetmgrFuture = b'2'  # 纯期货资管业务下的资管结算账户
THOST_FTDC_CAT_AssetmgrTrustee = b'3'  # 综合类资管业务下的期货资管托管账户
THOST_FTDC_CAT_AssetmgrTransfer = b'4'  # 综合类资管业务下的资金中转账户

TThostFtdcCustTypeType = ctypes.c_char

THOST_FTDC_CUSTT_Person = b'0'  # 自然人
THOST_FTDC_CUSTT_Institution = b'1'  # 机构户

TThostFtdcDAClientTypeType = ctypes.c_char

THOST_FTDC_CACT_Person = b'0'  # 自然人
THOST_FTDC_CACT_Company = b'1'  # 法人
THOST_FTDC_CACT_Other = b'2'  # 其他

TThostFtdcDBOperationType = ctypes.c_char

THOST_FTDC_DBOP_Insert = b'0'  # 插入
THOST_FTDC_DBOP_Update = b'1'  # 更新
THOST_FTDC_DBOP_Delete = b'2'  # 删除

TThostFtdcDCEUploadFileNameType = ctypes.c_char

THOST_FTDC_DUFN_DUFN_O = b'O'  # ^\d{8}_dl_\d{3}
THOST_FTDC_DUFN_DUFN_T = b'T'  # ^\d{8}_成交表
THOST_FTDC_DUFN_DUFN_P = b'P'  # ^\d{8}_持仓表
THOST_FTDC_DUFN_DUFN_F = b'F'  # ^\d{8}_资金结算表
THOST_FTDC_DUFN_DUFN_C = b'C'  # ^\d{8}_优惠组合持仓明细表
THOST_FTDC_DUFN_DUFN_D = b'D'  # ^\d{8}_持仓明细表
THOST_FTDC_DUFN_DUFN_M = b'M'  # ^\d{8}_保证金参数表
THOST_FTDC_DUFN_DUFN_S = b'S'  # ^\d{8}_期权执行表

TThostFtdcDataResourceType = ctypes.c_char

THOST_FTDC_DAR_Settle = b'1'  # 本系统
THOST_FTDC_DAR_Exchange = b'2'  # 交易所
THOST_FTDC_DAR_CSRC = b'3'  # 报送数据

TThostFtdcDataStatusType = ctypes.c_char

THOST_FTDC_AMLDS_Normal = b'0'  # 正常
THOST_FTDC_AMLDS_Deleted = b'1'  # 已删除

TThostFtdcDataSyncStatusType = ctypes.c_char

THOST_FTDC_DS_Asynchronous = b'1'  # 未同步
THOST_FTDC_DS_Synchronizing = b'2'  # 同步中
THOST_FTDC_DS_Synchronized = b'3'  # 已同步

TThostFtdcDceCombinationTypeType = ctypes.c_char

THOST_FTDC_DCECOMBT_SPL = b'0'  # 期货对锁组合
THOST_FTDC_DCECOMBT_OPL = b'1'  # 期权对锁组合
THOST_FTDC_DCECOMBT_SP = b'2'  # 期货跨期组合
THOST_FTDC_DCECOMBT_SPC = b'3'  # 期货跨品种组合
THOST_FTDC_DCECOMBT_BLS = b'4'  # 买入期权垂直价差组合
THOST_FTDC_DCECOMBT_BES = b'5'  # 卖出期权垂直价差组合
THOST_FTDC_DCECOMBT_CAS = b'6'  # 期权日历价差组合
THOST_FTDC_DCECOMBT_STD = b'7'  # 期权跨式组合
THOST_FTDC_DCECOMBT_STG = b'8'  # 期权宽跨式组合
THOST_FTDC_DCECOMBT_BFO = b'9'  # 买入期货期权组合
THOST_FTDC_DCECOMBT_SFO = b'a'  # 卖出期货期权组合

TThostFtdcDeliveryModeType = ctypes.c_char

THOST_FTDC_DM_CashDeliv = b'1'  # 现金交割
THOST_FTDC_DM_CommodityDeliv = b'2'  # 实物交割

TThostFtdcDeliveryTypeType = ctypes.c_char

THOST_FTDC_DT_HandDeliv = b'1'  # 手工交割
THOST_FTDC_DT_PersonDeliv = b'2'  # 到期交割

TThostFtdcDepartmentRangeType = ctypes.c_char

THOST_FTDC_DR_All = b'1'  # 所有
THOST_FTDC_DR_Group = b'2'  # 组织架构
THOST_FTDC_DR_Single = b'3'  # 单一投资者

TThostFtdcDirectionEnType = ctypes.c_char

THOST_FTDC_DEN_Buy = b'0'  # Buy
THOST_FTDC_DEN_Sell = b'1'  # Sell

TThostFtdcDirectionType = ctypes.c_char

THOST_FTDC_D_Buy = b'0'  # 买
THOST_FTDC_D_Sell = b'1'  # 卖

TThostFtdcEnumBoolType = ctypes.c_char

THOST_FTDC_EBL_False = b'0'  # false
THOST_FTDC_EBL_True = b'1'  # true

TThostFtdcEventModeType = ctypes.c_char

THOST_FTDC_EvM_ADD = b'1'  # 增加
THOST_FTDC_EvM_UPDATE = b'2'  # 修改
THOST_FTDC_EvM_DELETE = b'3'  # 删除
THOST_FTDC_EvM_CHECK = b'4'  # 复核
THOST_FTDC_EvM_COPY = b'5'  # 复制
THOST_FTDC_EvM_CANCEL = b'6'  # 注销
THOST_FTDC_EvM_Reverse = b'7'  # 冲销

TThostFtdcExClientIDTypeType = ctypes.c_char

THOST_FTDC_ECIDT_Hedge = b'1'  # 套保
THOST_FTDC_ECIDT_Arbitrage = b'2'  # 套利
THOST_FTDC_ECIDT_Speculation = b'3'  # 投机

TThostFtdcExDirectionType = ctypes.c_char

THOST_FTDC_FBEDIR_Settlement = b'0'  # 结汇
THOST_FTDC_FBEDIR_Sale = b'1'  # 售汇

TThostFtdcExStatusType = ctypes.c_char

THOST_FTDC_EXS_Before = b'0'  # 修改前
THOST_FTDC_EXS_After = b'1'  # 修改后

TThostFtdcExchangeConnectStatusType = ctypes.c_char

THOST_FTDC_ECS_NoConnection = b'1'  # 没有任何连接
THOST_FTDC_ECS_QryInstrumentSent = b'2'  # 已经发出合约查询请求
THOST_FTDC_ECS_GotInformation = b'9'  # 已经获取信息

TThostFtdcExchangeIDTypeType = ctypes.c_char

THOST_FTDC_EIDT_SHFE = b'S'  # 上海期货交易所
THOST_FTDC_EIDT_CZCE = b'Z'  # 郑州商品交易所
THOST_FTDC_EIDT_DCE = b'D'  # 大连商品交易所
THOST_FTDC_EIDT_CFFEX = b'J'  # 中国金融期货交易所
THOST_FTDC_EIDT_INE = b'N'  # 上海国际能源交易中心股份有限公司

TThostFtdcExchangePropertyType = ctypes.c_char

THOST_FTDC_EXP_Normal = b'0'  # 正常
THOST_FTDC_EXP_GenOrderByTrade = b'1'  # 根据成交生成报单

TThostFtdcExchangeSettlementParamIDType = ctypes.c_char

THOST_FTDC_ESPI_MortgageRatio = b'1'  # 质押比例
THOST_FTDC_ESPI_OtherFundItem = b'2'  # 分项资金导入项
THOST_FTDC_ESPI_OtherFundImport = b'3'  # 分项资金入交易所出入金
THOST_FTDC_ESPI_CFFEXMinPrepa = b'6'  # 中金所开户最低可用金额
THOST_FTDC_ESPI_CZCESettlementType = b'7'  # 郑商所结算方式
THOST_FTDC_ESPI_ExchDelivFeeMode = b'9'  # 交易所交割手续费收取方式
THOST_FTDC_ESPI_DelivFeeMode = b'0'  # 投资者交割手续费收取方式
THOST_FTDC_ESPI_CZCEComMarginType = b'A'  # 郑商所组合持仓保证金收取方式
THOST_FTDC_ESPI_DceComMarginType = b'B'  # 大商所套利保证金是否优惠
THOST_FTDC_ESPI_OptOutDisCountRate = b'a'  # 虚值期权保证金优惠比率
THOST_FTDC_ESPI_OptMiniGuarantee = b'b'  # 最低保障系数

TThostFtdcExecOrderCloseFlagType = ctypes.c_char

THOST_FTDC_EOCF_AutoClose = b'0'  # 自动平仓
THOST_FTDC_EOCF_NotToClose = b'1'  # 免于自动平仓

TThostFtdcExecOrderPositionFlagType = ctypes.c_char

THOST_FTDC_EOPF_Reserve = b'0'  # 保留
THOST_FTDC_EOPF_UnReserve = b'1'  # 不保留

TThostFtdcExecResultType = ctypes.c_char

THOST_FTDC_OER_NoExec = b'n'  # 没有执行
THOST_FTDC_OER_Canceled = b'c'  # 已经取消
THOST_FTDC_OER_OK = b'0'  # 执行成功
THOST_FTDC_OER_NoPosition = b'1'  # 期权持仓不够
THOST_FTDC_OER_NoDeposit = b'2'  # 资金不够
THOST_FTDC_OER_NoParticipant = b'3'  # 会员不存在
THOST_FTDC_OER_NoClient = b'4'  # 客户不存在
THOST_FTDC_OER_NoInstrument = b'6'  # 合约不存在
THOST_FTDC_OER_NoRight = b'7'  # 没有执行权限
THOST_FTDC_OER_InvalidVolume = b'8'  # 不合理的数量
THOST_FTDC_OER_NoEnoughHistoryTrade = b'9'  # 没有足够的历史成交
THOST_FTDC_OER_Unknown = b'a'  # 未知

TThostFtdcExportFileTypeType = ctypes.c_char

THOST_FTDC_EFT_CSV = b'0'  # CSV
THOST_FTDC_EFT_EXCEL = b'1'  # Excel
THOST_FTDC_EFT_DBF = b'2'  # DBF

TThostFtdcExprSetModeType = ctypes.c_char

THOST_FTDC_ESM_Relative = b'1'  # 相对已有规则设置
THOST_FTDC_ESM_Typical = b'2'  # 典型设置

TThostFtdcFBEAlreadyTradeType = ctypes.c_char

THOST_FTDC_FBEAT_NotTrade = b'0'  # 未交易
THOST_FTDC_FBEAT_Trade = b'1'  # 已交易

TThostFtdcFBEExchStatusType = ctypes.c_char

THOST_FTDC_FBEES_Normal = b'0'  # 正常
THOST_FTDC_FBEES_ReExchange = b'1'  # 交易重发

TThostFtdcFBEFileFlagType = ctypes.c_char

THOST_FTDC_FBEFG_DataPackage = b'0'  # 数据包
THOST_FTDC_FBEFG_File = b'1'  # 文件

TThostFtdcFBEReqFlagType = ctypes.c_char

THOST_FTDC_FBERF_UnProcessed = b'0'  # 未处理
THOST_FTDC_FBERF_WaitSend = b'1'  # 等待发送
THOST_FTDC_FBERF_SendSuccess = b'2'  # 发送成功
THOST_FTDC_FBERF_SendFailed = b'3'  # 发送失败
THOST_FTDC_FBERF_WaitReSend = b'4'  # 等待重发

TThostFtdcFBEResultFlagType = ctypes.c_char

THOST_FTDC_FBERES_Success = b'0'  # 成功
THOST_FTDC_FBERES_InsufficientBalance = b'1'  # 账户余额不足
THOST_FTDC_FBERES_UnknownTrading = b'8'  # 交易结果未知
THOST_FTDC_FBERES_Fail = b'x'  # 失败

TThostFtdcFBEUserEventTypeType = ctypes.c_char

THOST_FTDC_FBEUET_SignIn = b'0'  # 签到
THOST_FTDC_FBEUET_Exchange = b'1'  # 换汇
THOST_FTDC_FBEUET_ReExchange = b'2'  # 换汇重发
THOST_FTDC_FBEUET_QueryBankAccount = b'3'  # 银行账户查询
THOST_FTDC_FBEUET_QueryExchDetial = b'4'  # 换汇明细查询
THOST_FTDC_FBEUET_QueryExchSummary = b'5'  # 换汇汇总查询
THOST_FTDC_FBEUET_QueryExchRate = b'6'  # 换汇汇率查询
THOST_FTDC_FBEUET_CheckBankAccount = b'7'  # 对账文件通知
THOST_FTDC_FBEUET_SignOut = b'8'  # 签退
THOST_FTDC_FBEUET_Other = b'Z'  # 其他

TThostFtdcFBTEncryModeType = ctypes.c_char

THOST_FTDC_EM_NoEncry = b'0'  # 不加密
THOST_FTDC_EM_DES = b'1'  # DES
THOST_FTDC_EM_3DES = b'2'  # 3DES

TThostFtdcFBTPassWordTypeType = ctypes.c_char

THOST_FTDC_PWT_Query = b'0'  # 查询
THOST_FTDC_PWT_Fetch = b'1'  # 取款
THOST_FTDC_PWT_Transfer = b'2'  # 转帐
THOST_FTDC_PWT_Trade = b'3'  # 交易

TThostFtdcFBTTradeCodeEnumType = str

THOST_FTDC_FTC_BankLaunchBankToBroker = "102001"  # 银行发起银行转期货
THOST_FTDC_FTC_BrokerLaunchBankToBroker = "202001"  # 期货发起银行转期货
THOST_FTDC_FTC_BankLaunchBrokerToBank = "102002"  # 银行发起期货转银行
THOST_FTDC_FTC_BrokerLaunchBrokerToBank = "202002"  # 期货发起期货转银行

TThostFtdcFBTTransferDirectionType = ctypes.c_char

THOST_FTDC_FBTTD_FromBankToFuture = b'1'  # 入金，银行转期货
THOST_FTDC_FBTTD_FromFutureToBank = b'2'  # 出金，期货转银行

TThostFtdcFBTUserEventTypeType = ctypes.c_char

THOST_FTDC_FBTUET_SignIn = b'0'  # 签到
THOST_FTDC_FBTUET_FromBankToFuture = b'1'  # 银行转期货
THOST_FTDC_FBTUET_FromFutureToBank = b'2'  # 期货转银行
THOST_FTDC_FBTUET_OpenAccount = b'3'  # 开户
THOST_FTDC_FBTUET_CancelAccount = b'4'  # 销户
THOST_FTDC_FBTUET_ChangeAccount = b'5'  # 变更银行账户
THOST_FTDC_FBTUET_RepealFromBankToFuture = b'6'  # 冲正银行转期货
THOST_FTDC_FBTUET_RepealFromFutureToBank = b'7'  # 冲正期货转银行
THOST_FTDC_FBTUET_QueryBankAccount = b'8'  # 查询银行账户
THOST_FTDC_FBTUET_QueryFutureAccount = b'9'  # 查询期货账户
THOST_FTDC_FBTUET_SignOut = b'A'  # 签退
THOST_FTDC_FBTUET_SyncKey = b'B'  # 密钥同步
THOST_FTDC_FBTUET_ReserveOpenAccount = b'C'  # 预约开户
THOST_FTDC_FBTUET_CancelReserveOpenAccount = b'D'  # 撤销预约开户
THOST_FTDC_FBTUET_ReserveOpenAccountConfirm = b'E'  # 预约开户确认
THOST_FTDC_FBTUET_Other = b'Z'  # 其他

TThostFtdcFeeAcceptStyleType = ctypes.c_char

THOST_FTDC_FAS_ByTrade = b'1'  # 按交易收取
THOST_FTDC_FAS_ByDeliv = b'2'  # 按交割收取
THOST_FTDC_FAS_None = b'3'  # 不收
THOST_FTDC_FAS_FixFee = b'4'  # 按指定手续费收取

TThostFtdcFeePayFlagType = ctypes.c_char

THOST_FTDC_FPF_BEN = b'0'  # 由受益方支付费用
THOST_FTDC_FPF_OUR = b'1'  # 由发送方支付费用
THOST_FTDC_FPF_SHA = b'2'  # 由发送方支付发起的费用，受益方支付接受的费用

TThostFtdcFileBusinessCodeType = ctypes.c_char

THOST_FTDC_FBC_Others = b'0'  # 其他
THOST_FTDC_FBC_TransferDetails = b'1'  # 转账交易明细对账
THOST_FTDC_FBC_CustAccStatus = b'2'  # 客户账户状态对账
THOST_FTDC_FBC_AccountTradeDetails = b'3'  # 账户类交易明细对账
THOST_FTDC_FBC_FutureAccountChangeInfoDetails = b'4'  # 期货账户信息变更明细对账
THOST_FTDC_FBC_CustMoneyDetail = b'5'  # 客户资金台账余额明细对账
THOST_FTDC_FBC_CustCancelAccountInfo = b'6'  # 客户销户结息明细对账
THOST_FTDC_FBC_CustMoneyResult = b'7'  # 客户资金余额对账结果
THOST_FTDC_FBC_OthersExceptionResult = b'8'  # 其它对账异常结果文件
THOST_FTDC_FBC_CustInterestNetMoneyDetails = b'9'  # 客户结息净额明细
THOST_FTDC_FBC_CustMoneySendAndReceiveDetails = b'a'  # 客户资金交收明细
THOST_FTDC_FBC_CorporationMoneyTotal = b'b'  # 法人存管银行资金交收汇总
THOST_FTDC_FBC_MainbodyMoneyTotal = b'c'  # 主体间资金交收汇总
THOST_FTDC_FBC_MainPartMonitorData = b'd'  # 总分平衡监管数据
THOST_FTDC_FBC_PreparationMoney = b'e'  # 存管银行备付金余额
THOST_FTDC_FBC_BankMoneyMonitorData = b'f'  # 协办存管银行资金监管数据

TThostFtdcFileFormatType = ctypes.c_char

THOST_FTDC_FFT_Txt = b'0'  # 文本文件(.txt)
THOST_FTDC_FFT_Zip = b'1'  # 压缩文件(.zip)
THOST_FTDC_FFT_DBF = b'2'  # DBF文件(.dbf)

TThostFtdcFileGenStyleType = ctypes.c_char

THOST_FTDC_FGS_FileTransmit = b'0'  # 下发
THOST_FTDC_FGS_FileGen = b'1'  # 生成

TThostFtdcFileIDType = ctypes.c_char

THOST_FTDC_FI_SettlementFund = b'F'  # 资金数据
THOST_FTDC_FI_Trade = b'T'  # 成交数据
THOST_FTDC_FI_InvestorPosition = b'P'  # 投资者持仓数据
THOST_FTDC_FI_SubEntryFund = b'O'  # 投资者分项资金数据
THOST_FTDC_FI_CZCECombinationPos = b'C'  # 组合持仓数据
THOST_FTDC_FI_CSRCData = b'R'  # 上报保证金监控中心数据
THOST_FTDC_FI_CZCEClose = b'L'  # 郑商所平仓了结数据
THOST_FTDC_FI_CZCENoClose = b'N'  # 郑商所非平仓了结数据
THOST_FTDC_FI_PositionDtl = b'D'  # 持仓明细数据
THOST_FTDC_FI_OptionStrike = b'S'  # 期权执行文件
THOST_FTDC_FI_SettlementPriceComparison = b'M'  # 结算价比对文件
THOST_FTDC_FI_NonTradePosChange = b'B'  # 上期所非持仓变动明细

TThostFtdcFileStatusType = ctypes.c_char

THOST_FTDC_FIS_NoCreate = b'0'  # 未生成
THOST_FTDC_FIS_Created = b'1'  # 已生成
THOST_FTDC_FIS_Failed = b'2'  # 生成失败

TThostFtdcFileTypeType = ctypes.c_char

THOST_FTDC_FUT_Settlement = b'0'  # 结算
THOST_FTDC_FUT_Check = b'1'  # 核对

TThostFtdcFileUploadStatusType = ctypes.c_char

THOST_FTDC_FUS_SucceedUpload = b'1'  # 上传成功
THOST_FTDC_FUS_FailedUpload = b'2'  # 上传失败
THOST_FTDC_FUS_SucceedLoad = b'3'  # 导入成功
THOST_FTDC_FUS_PartSucceedLoad = b'4'  # 导入部分成功
THOST_FTDC_FUS_FailedLoad = b'5'  # 导入失败

TThostFtdcFindMarginRateAlgoIDType = ctypes.c_char

THOST_FTDC_FMRA_Base = b'1'  # 基本
THOST_FTDC_FMRA_DCE = b'2'  # 大连商品交易所
THOST_FTDC_FMRA_CZCE = b'3'  # 郑州商品交易所

TThostFtdcFlexStatModeType = ctypes.c_char

THOST_FTDC_FSM_Product = b'1'  # 产品统计
THOST_FTDC_FSM_Exchange = b'2'  # 交易所统计
THOST_FTDC_FSM_All = b'3'  # 统计所有

TThostFtdcFlowIDType = ctypes.c_char

THOST_FTDC_EvM_InvestorGroupFlow = b'1'  # 投资者对应投资者组设置
THOST_FTDC_EvM_InvestorRate = b'2'  # 投资者手续费率设置
THOST_FTDC_EvM_InvestorCommRateModel = b'3'  # 投资者手续费率模板关系设置

TThostFtdcForQuoteStatusType = ctypes.c_char

THOST_FTDC_FQST_Submitted = b'a'  # 已经提交
THOST_FTDC_FQST_Accepted = b'b'  # 已经接受
THOST_FTDC_FQST_Rejected = b'c'  # 已经被拒绝

TThostFtdcForceCloseReasonType = ctypes.c_char

THOST_FTDC_FCC_NotForceClose = b'0'  # 非强平
THOST_FTDC_FCC_LackDeposit = b'1'  # 资金不足
THOST_FTDC_FCC_ClientOverPositionLimit = b'2'  # 客户超仓
THOST_FTDC_FCC_MemberOverPositionLimit = b'3'  # 会员超仓
THOST_FTDC_FCC_NotMultiple = b'4'  # 持仓非整数倍
THOST_FTDC_FCC_Violation = b'5'  # 违规
THOST_FTDC_FCC_Other = b'6'  # 其它
THOST_FTDC_FCC_PersonDeliv = b'7'  # 自然人临近交割
THOST_FTDC_FCC_Notverifycapital = b'8'  # 本地强平资金不足忽略敞口
THOST_FTDC_FCC_LocalLackDeposit = b'9'  # 本地强平资金不足
THOST_FTDC_FCC_LocalViolationNocheck = b'a'  # 本地强平违规持仓忽略敞口
THOST_FTDC_FCC_LocalViolation = b'b'  # 本地强平违规持仓

TThostFtdcForceCloseTypeType = ctypes.c_char

THOST_FTDC_FCT_Manual = b'0'  # 手工强平
THOST_FTDC_FCT_Single = b'1'  # 单一投资者辅助强平
THOST_FTDC_FCT_Group = b'2'  # 批量投资者辅助强平

TThostFtdcFreezeStatusType = ctypes.c_char

THOST_FTDC_FRS_Normal = b'1'  # 活跃
THOST_FTDC_FRS_Freeze = b'0'  # 休眠

TThostFtdcFunctionCodeType = ctypes.c_char

THOST_FTDC_FC_DataAsync = b'1'  # 数据异步化
THOST_FTDC_FC_ForceUserLogout = b'2'  # 强制用户登出
THOST_FTDC_FC_UserPasswordUpdate = b'3'  # 变更管理用户口令
THOST_FTDC_FC_BrokerPasswordUpdate = b'4'  # 变更经纪公司口令
THOST_FTDC_FC_InvestorPasswordUpdate = b'5'  # 变更投资者口令
THOST_FTDC_FC_OrderInsert = b'6'  # 报单插入
THOST_FTDC_FC_OrderAction = b'7'  # 报单操作
THOST_FTDC_FC_SyncSystemData = b'8'  # 同步系统数据
THOST_FTDC_FC_SyncBrokerData = b'9'  # 同步经纪公司数据
THOST_FTDC_FC_BachSyncBrokerData = b'A'  # 批量同步经纪公司数据
THOST_FTDC_FC_SuperQuery = b'B'  # 超级查询
THOST_FTDC_FC_ParkedOrderInsert = b'C'  # 预埋报单插入
THOST_FTDC_FC_ParkedOrderAction = b'D'  # 预埋报单操作
THOST_FTDC_FC_SyncOTP = b'E'  # 同步动态令牌
THOST_FTDC_FC_DeleteOrder = b'F'  # 删除未知单
THOST_FTDC_FC_ExitEmergency = b'G'  # 退出紧急状态

TThostFtdcFundDirectionEnType = ctypes.c_char

THOST_FTDC_FDEN_In = b'1'  # Deposit
THOST_FTDC_FDEN_Out = b'2'  # Withdrawal

TThostFtdcFundDirectionType = ctypes.c_char

THOST_FTDC_FD_In = b'1'  # 入金
THOST_FTDC_FD_Out = b'2'  # 出金

TThostFtdcFundEventTypeType = ctypes.c_char

THOST_FTDC_FET_Restriction = b'0'  # 转账限额
THOST_FTDC_FET_TodayRestriction = b'1'  # 当日转账限额
THOST_FTDC_FET_Transfer = b'2'  # 期商流水
THOST_FTDC_FET_Credit = b'3'  # 资金冻结
THOST_FTDC_FET_InvestorWithdrawAlm = b'4'  # 投资者可提资金比例
THOST_FTDC_FET_BankRestriction = b'5'  # 单个银行帐户转账限额
THOST_FTDC_FET_Accountregister = b'6'  # 银期签约账户
THOST_FTDC_FET_ExchangeFundIO = b'7'  # 交易所出入金
THOST_FTDC_FET_InvestorFundIO = b'8'  # 投资者出入金

TThostFtdcFundIOTypeEnType = ctypes.c_char

THOST_FTDC_FIOTEN_FundIO = b'1'  # Deposit/Withdrawal
THOST_FTDC_FIOTEN_Transfer = b'2'  # Bank-Futures Transfer
THOST_FTDC_FIOTEN_SwapCurrency = b'3'  # Bank-Futures FX Exchange

TThostFtdcFundIOTypeType = ctypes.c_char

THOST_FTDC_FIOT_FundIO = b'1'  # 出入金
THOST_FTDC_FIOT_Transfer = b'2'  # 银期转帐
THOST_FTDC_FIOT_SwapCurrency = b'3'  # 银期换汇

TThostFtdcFundMortDirectionEnType = ctypes.c_char

THOST_FTDC_FMDEN_In = b'1'  # Pledge
THOST_FTDC_FMDEN_Out = b'2'  # Redemption

TThostFtdcFundMortDirectionType = ctypes.c_char

THOST_FTDC_FMD_In = b'1'  # 货币质入
THOST_FTDC_FMD_Out = b'2'  # 货币质出

TThostFtdcFundMortgageTypeType = ctypes.c_char

THOST_FTDC_FMT_Mortgage = b'1'  # 质押
THOST_FTDC_FMT_Redemption = b'2'  # 解质

TThostFtdcFundStatusType = ctypes.c_char

THOST_FTDC_FS_Record = b'1'  # 已录入
THOST_FTDC_FS_Check = b'2'  # 已复核
THOST_FTDC_FS_Charge = b'3'  # 已冲销

TThostFtdcFundTypeEnType = ctypes.c_char

THOST_FTDC_FTEN_Deposite = b'1'  # Bank Deposit
THOST_FTDC_FTEN_ItemFund = b'2'  # Payment/Fee
THOST_FTDC_FTEN_Company = b'3'  # Brokerage Adj
THOST_FTDC_FTEN_InnerTransfer = b'4'  # Internal Transfer

TThostFtdcFundTypeType = ctypes.c_char

THOST_FTDC_FT_Deposite = b'1'  # 银行存款
THOST_FTDC_FT_ItemFund = b'2'  # 分项资金
THOST_FTDC_FT_Company = b'3'  # 公司调整
THOST_FTDC_FT_InnerTransfer = b'4'  # 资金内转

TThostFtdcFutureAccTypeType = ctypes.c_char

THOST_FTDC_FAT_BankBook = b'1'  # 银行存折
THOST_FTDC_FAT_SavingCard = b'2'  # 储蓄卡
THOST_FTDC_FAT_CreditCard = b'3'  # 信用卡

TThostFtdcFuturePwdFlagType = ctypes.c_char

THOST_FTDC_FPWD_UnCheck = b'0'  # 不核对
THOST_FTDC_FPWD_Check = b'1'  # 核对

TThostFtdcFutureTypeType = ctypes.c_char

THOST_FTDC_FUTT_Commodity = b'1'  # 商品期货
THOST_FTDC_FUTT_Financial = b'2'  # 金融期货

TThostFtdcGenderType = ctypes.c_char

THOST_FTDC_GD_Unknown = b'0'  # 未知状态
THOST_FTDC_GD_Male = b'1'  # 男
THOST_FTDC_GD_Female = b'2'  # 女

TThostFtdcGiveUpDataSourceType = ctypes.c_char

THOST_FTDC_GUDS_Gen = b'0'  # 系统生成
THOST_FTDC_GUDS_Hand = b'1'  # 手工添加

TThostFtdcHandlePositionAlgoIDType = ctypes.c_char

THOST_FTDC_HPA_Base = b'1'  # 基本
THOST_FTDC_HPA_DCE = b'2'  # 大连商品交易所
THOST_FTDC_HPA_CZCE = b'3'  # 郑州商品交易所

TThostFtdcHandleTradingAccountAlgoIDType = ctypes.c_char

THOST_FTDC_HTAA_Base = b'1'  # 基本
THOST_FTDC_HTAA_DCE = b'2'  # 大连商品交易所
THOST_FTDC_HTAA_CZCE = b'3'  # 郑州商品交易所

TThostFtdcHasBoardType = ctypes.c_char

THOST_FTDC_HB_No = b'0'  # 没有
THOST_FTDC_HB_Yes = b'1'  # 有

TThostFtdcHasTrusteeType = ctypes.c_char

THOST_FTDC_HT_Yes = b'1'  # 有
THOST_FTDC_HT_No = b'0'  # 没有

TThostFtdcHedgeFlagEnType = ctypes.c_char

THOST_FTDC_HFEN_Speculation = b'1'  # Speculation
THOST_FTDC_HFEN_Arbitrage = b'2'  # Arbitrage
THOST_FTDC_HFEN_Hedge = b'3'  # Hedge

TThostFtdcHedgeFlagType = ctypes.c_char

THOST_FTDC_HF_Speculation = b'1'  # 投机
THOST_FTDC_HF_Arbitrage = b'2'  # 套利
THOST_FTDC_HF_Hedge = b'3'  # 套保
THOST_FTDC_HF_MarketMaker = b'5'  # 做市商
THOST_FTDC_HF_SpecHedge = b'6'  # 第一腿投机第二腿套保
THOST_FTDC_HF_HedgeSpec = b'7'  # 第一腿套保第二腿投机

TThostFtdcIdCardTypeType = ctypes.c_char

THOST_FTDC_ICT_EID = b'0'  # 组织机构代码
THOST_FTDC_ICT_IDCard = b'1'  # 中国公民身份证
THOST_FTDC_ICT_OfficerIDCard = b'2'  # 军官证
THOST_FTDC_ICT_PoliceIDCard = b'3'  # 警官证
THOST_FTDC_ICT_SoldierIDCard = b'4'  # 士兵证
THOST_FTDC_ICT_HouseholdRegister = b'5'  # 户口簿
THOST_FTDC_ICT_Passport = b'6'  # 护照
THOST_FTDC_ICT_TaiwanCompatriotIDCard = b'7'  # 台胞证
THOST_FTDC_ICT_HomeComingCard = b'8'  # 回乡证
THOST_FTDC_ICT_LicenseNo = b'9'  # 营业执照号
THOST_FTDC_ICT_TaxNo = b'A'  # 税务登记号/当地纳税ID
THOST_FTDC_ICT_HMMainlandTravelPermit = b'B'  # 港澳居民来往内地通行证
THOST_FTDC_ICT_TwMainlandTravelPermit = b'C'  # 台湾居民来往大陆通行证
THOST_FTDC_ICT_DrivingLicense = b'D'  # 驾照
THOST_FTDC_ICT_SocialID = b'F'  # 当地社保ID
THOST_FTDC_ICT_LocalID = b'G'  # 当地身份证
THOST_FTDC_ICT_BusinessRegistration = b'H'  # 商业登记证
THOST_FTDC_ICT_HKMCIDCard = b'I'  # 港澳永久性居民身份证
THOST_FTDC_ICT_AccountsPermits = b'J'  # 人行开户许可证
THOST_FTDC_ICT_FrgPrmtRdCard = b'K'  # 外国人永久居留证
THOST_FTDC_ICT_CptMngPrdLetter = b'L'  # 资管产品备案函
THOST_FTDC_ICT_HKMCTwResidencePermit = b'M'  # 港澳台居民居住证
THOST_FTDC_ICT_UniformSocialCreditCode = b'N'  # 统一社会信用代码
THOST_FTDC_ICT_CorporationCertNo = b'O'  # 机构成立证明文件
THOST_FTDC_ICT_OtherCard = b'x'  # 其他证件

TThostFtdcIncludeCloseProfitType = ctypes.c_char

THOST_FTDC_ICP_Include = b'0'  # 包含平仓盈利
THOST_FTDC_ICP_NotInclude = b'2'  # 不包含平仓盈利

TThostFtdcInitSettlementType = ctypes.c_char

THOST_FTDC_SIS_UnInitialize = b'0'  # 结算初始化未开始
THOST_FTDC_SIS_Initialize = b'1'  # 结算初始化中
THOST_FTDC_SIS_Initialized = b'2'  # 结算初始化完成

TThostFtdcInstLifePhaseType = ctypes.c_char

THOST_FTDC_IP_NotStart = b'0'  # 未上市
THOST_FTDC_IP_Started = b'1'  # 上市
THOST_FTDC_IP_Pause = b'2'  # 停牌
THOST_FTDC_IP_Expired = b'3'  # 到期

TThostFtdcInstMarginCalIDType = ctypes.c_char

THOST_FTDC_IMID_BothSide = b'1'  # 标准算法收取双边
THOST_FTDC_IMID_MMSA = b'2'  # 单向大边
THOST_FTDC_IMID_SPMM = b'3'  # 新组保SPMM

TThostFtdcInstStatusEnterReasonType = ctypes.c_char

THOST_FTDC_IER_Automatic = b'1'  # 自动切换
THOST_FTDC_IER_Manual = b'2'  # 手动切换
THOST_FTDC_IER_Fuse = b'3'  # 熔断

TThostFtdcInstitutionTypeType = ctypes.c_char

THOST_FTDC_TS_Bank = b'0'  # 银行
THOST_FTDC_TS_Future = b'1'  # 期商
THOST_FTDC_TS_Store = b'2'  # 券商

TThostFtdcInstrumentClassType = ctypes.c_char

THOST_FTDC_EIC_Usual = b'1'  # 一般月份合约
THOST_FTDC_EIC_Delivery = b'2'  # 临近交割合约
THOST_FTDC_EIC_NonComb = b'3'  # 非组合合约

TThostFtdcInstrumentStatusType = ctypes.c_char

THOST_FTDC_IS_BeforeTrading = b'0'  # 开盘前
THOST_FTDC_IS_NoTrading = b'1'  # 非交易
THOST_FTDC_IS_Continous = b'2'  # 连续交易
THOST_FTDC_IS_AuctionOrdering = b'3'  # 集合竞价报单
THOST_FTDC_IS_AuctionBalance = b'4'  # 集合竞价价格平衡
THOST_FTDC_IS_AuctionMatch = b'5'  # 集合竞价撮合
THOST_FTDC_IS_Closed = b'6'  # 收盘
THOST_FTDC_IS_TransactionProcessing = b'7'  # 交易业务处理

TThostFtdcInvestorRangeType = ctypes.c_char

THOST_FTDC_IR_All = b'1'  # 所有
THOST_FTDC_IR_Group = b'2'  # 投资者组
THOST_FTDC_IR_Single = b'3'  # 单一投资者

TThostFtdcInvestorRiskStatusType = ctypes.c_char

THOST_FTDC_IRS_Normal = b'1'  # 正常
THOST_FTDC_IRS_Warn = b'2'  # 警告
THOST_FTDC_IRS_Call = b'3'  # 追保
THOST_FTDC_IRS_Force = b'4'  # 强平
THOST_FTDC_IRS_Exception = b'5'  # 异常

TThostFtdcInvestorSettlementParamIDType = ctypes.c_char

THOST_FTDC_ISPI_MortgageRatio = b'4'  # 质押比例
THOST_FTDC_ISPI_MarginWay = b'5'  # 保证金算法
THOST_FTDC_ISPI_BillDeposit = b'9'  # 结算单结存是否包含质押

TThostFtdcInvestorTypeType = ctypes.c_char

THOST_FTDC_CT_Person = b'0'  # 自然人
THOST_FTDC_CT_Company = b'1'  # 法人
THOST_FTDC_CT_Fund = b'2'  # 投资基金
THOST_FTDC_CT_SpecialOrgan = b'3'  # 特殊法人
THOST_FTDC_CT_Asset = b'4'  # 资管户

TThostFtdcInvstTradingRightType = ctypes.c_char

THOST_FTDC_ITR_CloseOnly = b'1'  # 只能平仓
THOST_FTDC_ITR_Forbidden = b'2'  # 不能交易

TThostFtdcLanguageTypeType = ctypes.c_char

THOST_FTDC_LT_Chinese = b'1'  # 中文
THOST_FTDC_LT_English = b'2'  # 英文

TThostFtdcLastFragmentType = ctypes.c_char

THOST_FTDC_LF_Yes = b'0'  # 是最后分片
THOST_FTDC_LF_No = b'1'  # 不是最后分片

TThostFtdcLimitUseTypeType = ctypes.c_char

THOST_FTDC_LUT_Repeatable = b'1'  # 可重复使用
THOST_FTDC_LUT_Unrepeatable = b'2'  # 不可重复使用

TThostFtdcLinkStatusType = ctypes.c_char

THOST_FTDC_LS_Connected = b'1'  # 已经连接
THOST_FTDC_LS_Disconnected = b'2'  # 没有连接

TThostFtdcLoginModeType = ctypes.c_char

THOST_FTDC_LM_Trade = b'0'  # 交易
THOST_FTDC_LM_Transfer = b'1'  # 转账

TThostFtdcManageStatusType = ctypes.c_char

THOST_FTDC_MSS_Point = b'0'  # 指定存管
THOST_FTDC_MSS_PrePoint = b'1'  # 预指定
THOST_FTDC_MSS_CancelPoint = b'2'  # 撤销指定

TThostFtdcMarginPriceTypeType = ctypes.c_char

THOST_FTDC_MPT_PreSettlementPrice = b'1'  # 昨结算价
THOST_FTDC_MPT_SettlementPrice = b'2'  # 最新价
THOST_FTDC_MPT_AveragePrice = b'3'  # 成交均价
THOST_FTDC_MPT_OpenPrice = b'4'  # 开仓价

TThostFtdcMarginRateTypeType = ctypes.c_char

THOST_FTDC_MRT_Exchange = b'1'  # 交易所保证金率
THOST_FTDC_MRT_Investor = b'2'  # 投资者保证金率
THOST_FTDC_MRT_InvestorTrade = b'3'  # 投资者交易保证金率

TThostFtdcMarginTypeType = ctypes.c_char

THOST_FTDC_MGT_ExchMarginRate = b'0'  # 交易所保证金率
THOST_FTDC_MGT_InstrMarginRate = b'1'  # 投资者保证金率
THOST_FTDC_MGT_InstrMarginRateTrade = b'2'  # 投资者交易保证金率

TThostFtdcMatchTypeType = ctypes.c_char

THOST_FTDC_OTC_MT_DV01 = b'1'  # 基点价值
THOST_FTDC_OTC_MT_ParValue = b'2'  # 面值

TThostFtdcMaxMarginSideAlgorithmType = ctypes.c_char

THOST_FTDC_MMSA_NO = b'0'  # 不使用大额单边保证金算法
THOST_FTDC_MMSA_YES = b'1'  # 使用大额单边保证金算法

TThostFtdcMoneyAccountStatusType = ctypes.c_char

THOST_FTDC_MAS_Normal = b'0'  # 正常
THOST_FTDC_MAS_Cancel = b'1'  # 销户

TThostFtdcMonthBillTradeSumType = ctypes.c_char

THOST_FTDC_MBTS_ByInstrument = b'0'  # 同日同合约
THOST_FTDC_MBTS_ByDayInsPrc = b'1'  # 同日同合约同价格
THOST_FTDC_MBTS_ByDayIns = b'2'  # 同合约

TThostFtdcMortgageFundUseRangeType = ctypes.c_char

THOST_FTDC_MFUR_None = b'0'  # 不能使用
THOST_FTDC_MFUR_Margin = b'1'  # 用于保证金
THOST_FTDC_MFUR_All = b'2'  # 用于手续费、盈亏、保证金
THOST_FTDC_MFUR_CNY3 = b'3'  # 人民币方案3

TThostFtdcMortgageTypeType = ctypes.c_char

THOST_FTDC_MT_Out = b'0'  # 质出
THOST_FTDC_MT_In = b'1'  # 质入

TThostFtdcNoteTypeType = ctypes.c_char

THOST_FTDC_NOTETYPE_TradeSettleBill = b'1'  # 交易结算单
THOST_FTDC_NOTETYPE_TradeSettleMonth = b'2'  # 交易结算月报
THOST_FTDC_NOTETYPE_CallMarginNotes = b'3'  # 追加保证金通知书
THOST_FTDC_NOTETYPE_ForceCloseNotes = b'4'  # 强行平仓通知书
THOST_FTDC_NOTETYPE_TradeNotes = b'5'  # 成交通知书
THOST_FTDC_NOTETYPE_DelivNotes = b'6'  # 交割通知书

TThostFtdcNotifyClassType = ctypes.c_char

THOST_FTDC_NC_NOERROR = b'0'  # 正常
THOST_FTDC_NC_Warn = b'1'  # 警示
THOST_FTDC_NC_Call = b'2'  # 追保
THOST_FTDC_NC_Force = b'3'  # 强平
THOST_FTDC_NC_CHUANCANG = b'4'  # 穿仓
THOST_FTDC_NC_Exception = b'5'  # 异常

TThostFtdcOTCTradeTypeType = ctypes.c_char

THOST_FTDC_OTC_TRDT_Block = b'0'  # 大宗交易
THOST_FTDC_OTC_TRDT_EFP = b'1'  # 期转现

TThostFtdcOTPStatusType = ctypes.c_char

THOST_FTDC_OTPS_Unused = b'0'  # 未使用
THOST_FTDC_OTPS_Used = b'1'  # 已使用
THOST_FTDC_OTPS_Disuse = b'2'  # 注销

TThostFtdcOTPTypeType = ctypes.c_char

THOST_FTDC_OTP_NONE = b'0'  # 无动态令牌
THOST_FTDC_OTP_TOTP = b'1'  # 时间令牌

TThostFtdcOffsetFlagEnType = ctypes.c_char

THOST_FTDC_OFEN_Open = b'0'  # Position Opening
THOST_FTDC_OFEN_Close = b'1'  # Position Close
THOST_FTDC_OFEN_ForceClose = b'2'  # Forced Liquidation
THOST_FTDC_OFEN_CloseToday = b'3'  # Close Today
THOST_FTDC_OFEN_CloseYesterday = b'4'  # Close Prev.
THOST_FTDC_OFEN_ForceOff = b'5'  # Forced Reduction
THOST_FTDC_OFEN_LocalForceClose = b'6'  # Local Forced Liquidation

TThostFtdcOffsetFlagType = ctypes.c_char

THOST_FTDC_OF_Open = b'0'  # 开仓
THOST_FTDC_OF_Close = b'1'  # 平仓
THOST_FTDC_OF_ForceClose = b'2'  # 强平
THOST_FTDC_OF_CloseToday = b'3'  # 平今
THOST_FTDC_OF_CloseYesterday = b'4'  # 平昨
THOST_FTDC_OF_ForceOff = b'5'  # 强减
THOST_FTDC_OF_LocalForceClose = b'6'  # 本地强平

TThostFtdcOpenLimitControlLevelType = ctypes.c_char

THOST_FTDC_PLCL_None = b'0'  # 不控制
THOST_FTDC_PLCL_Product = b'1'  # 产品级别
THOST_FTDC_PLCL_Inst = b'2'  # 合约级别

TThostFtdcOpenOrDestroyType = ctypes.c_char

THOST_FTDC_OOD_Open = b'1'  # 开户
THOST_FTDC_OOD_Destroy = b'0'  # 销户

TThostFtdcOptSelfCloseFlagType = ctypes.c_char

THOST_FTDC_OSCF_CloseSelfOptionPosition = b'1'  # 自对冲期权仓位
THOST_FTDC_OSCF_ReserveOptionPosition = b'2'  # 保留期权仓位
THOST_FTDC_OSCF_SellCloseSelfFuturePosition = b'3'  # 自对冲卖方履约后的期货仓位
THOST_FTDC_OSCF_ReserveFuturePosition = b'4'  # 保留卖方履约后的期货仓位

TThostFtdcOptionRoyaltyPriceTypeType = ctypes.c_char

THOST_FTDC_ORPT_PreSettlementPrice = b'1'  # 昨结算价
THOST_FTDC_ORPT_OpenPrice = b'4'  # 开仓价
THOST_FTDC_ORPT_MaxPreSettlementPrice = b'5'  # 最新价与昨结算价较大值

TThostFtdcOptionsTypeType = ctypes.c_char

THOST_FTDC_CP_CallOptions = b'1'  # 看涨
THOST_FTDC_CP_PutOptions = b'2'  # 看跌

TThostFtdcOrderActionStatusType = ctypes.c_char

THOST_FTDC_OAS_Submitted = b'a'  # 已经提交
THOST_FTDC_OAS_Accepted = b'b'  # 已经接受
THOST_FTDC_OAS_Rejected = b'c'  # 已经被拒绝

TThostFtdcOrderCancelAlgType = ctypes.c_char

THOST_FTDC_OAC_Balance = b'1'  # 轮询席位撤单
THOST_FTDC_OAC_OrigFirst = b'2'  # 优先原报单席位撤单

TThostFtdcOrderFreqControlLevelType = ctypes.c_char

THOST_FTDC_OFCL_None = b'0'  # 不控制
THOST_FTDC_OFCL_Product = b'1'  # 产品级别
THOST_FTDC_OFCL_Inst = b'2'  # 合约级别

TThostFtdcOrderPriceTypeType = ctypes.c_char

THOST_FTDC_OPT_AnyPrice = b'1'  # 任意价
THOST_FTDC_OPT_LimitPrice = b'2'  # 限价
THOST_FTDC_OPT_BestPrice = b'3'  # 最优价
THOST_FTDC_OPT_LastPrice = b'4'  # 最新价
THOST_FTDC_OPT_LastPricePlusOneTicks = b'5'  # 最新价浮动上浮1个ticks
THOST_FTDC_OPT_LastPricePlusTwoTicks = b'6'  # 最新价浮动上浮2个ticks
THOST_FTDC_OPT_LastPricePlusThreeTicks = b'7'  # 最新价浮动上浮3个ticks
THOST_FTDC_OPT_AskPrice1 = b'8'  # 卖一价
THOST_FTDC_OPT_AskPrice1PlusOneTicks = b'9'  # 卖一价浮动上浮1个ticks
THOST_FTDC_OPT_AskPrice1PlusTwoTicks = b'A'  # 卖一价浮动上浮2个ticks
THOST_FTDC_OPT_AskPrice1PlusThreeTicks = b'B'  # 卖一价浮动上浮3个ticks
THOST_FTDC_OPT_BidPrice1 = b'C'  # 买一价
THOST_FTDC_OPT_BidPrice1PlusOneTicks = b'D'  # 买一价浮动上浮1个ticks
THOST_FTDC_OPT_BidPrice1PlusTwoTicks = b'E'  # 买一价浮动上浮2个ticks
THOST_FTDC_OPT_BidPrice1PlusThreeTicks = b'F'  # 买一价浮动上浮3个ticks
THOST_FTDC_OPT_FiveLevelPrice = b'G'  # 五档价

TThostFtdcOrderSourceType = ctypes.c_char

THOST_FTDC_OSRC_Participant = b'0'  # 来自参与者
THOST_FTDC_OSRC_Administrator = b'1'  # 来自管理员

TThostFtdcOrderStatusType = ctypes.c_char

THOST_FTDC_OST_AllTraded = b'0'  # 全部成交
THOST_FTDC_OST_PartTradedQueueing = b'1'  # 部分成交还在队列中
THOST_FTDC_OST_PartTradedNotQueueing = b'2'  # 部分成交不在队列中
THOST_FTDC_OST_NoTradeQueueing = b'3'  # 未成交还在队列中
THOST_FTDC_OST_NoTradeNotQueueing = b'4'  # 未成交不在队列中
THOST_FTDC_OST_Canceled = b'5'  # 撤单
THOST_FTDC_OST_Unknown = b'a'  # 未知
THOST_FTDC_OST_NotTouched = b'b'  # 尚未触发
THOST_FTDC_OST_Touched = b'c'  # 已触发

TThostFtdcOrderSubmitStatusType = ctypes.c_char

THOST_FTDC_OSS_InsertSubmitted = b'0'  # 已经提交
THOST_FTDC_OSS_CancelSubmitted = b'1'  # 撤单已经提交
THOST_FTDC_OSS_ModifySubmitted = b'2'  # 修改已经提交
THOST_FTDC_OSS_Accepted = b'3'  # 已经接受
THOST_FTDC_OSS_InsertRejected = b'4'  # 报单已经被拒绝
THOST_FTDC_OSS_CancelRejected = b'5'  # 撤单已经被拒绝
THOST_FTDC_OSS_ModifyRejected = b'6'  # 改单已经被拒绝

TThostFtdcOrderTypeType = ctypes.c_char

THOST_FTDC_ORDT_Normal = b'0'  # 正常
THOST_FTDC_ORDT_DeriveFromQuote = b'1'  # 报价衍生
THOST_FTDC_ORDT_DeriveFromCombination = b'2'  # 组合衍生
THOST_FTDC_ORDT_Combination = b'3'  # 组合报单
THOST_FTDC_ORDT_ConditionalOrder = b'4'  # 条件单
THOST_FTDC_ORDT_Swap = b'5'  # 互换单
THOST_FTDC_ORDT_DeriveFromBlockTrade = b'6'  # 大宗交易成交衍生
THOST_FTDC_ORDT_DeriveFromEFPTrade = b'7'  # 期转现成交衍生

TThostFtdcOrgSystemIDType = ctypes.c_char

THOST_FTDC_ORGS_Standard = b'0'  # 综合交易平台
THOST_FTDC_ORGS_ESunny = b'1'  # 易盛系统
THOST_FTDC_ORGS_KingStarV6 = b'2'  # 金仕达V6系统

TThostFtdcOrganLevelType = ctypes.c_char

THOST_FTDC_OL_HeadQuarters = b'1'  # 银行总行或期商总部
THOST_FTDC_OL_Branch = b'2'  # 银行分中心或期货公司营业部

TThostFtdcOrganStatusType = ctypes.c_char

THOST_FTDC_OS_Ready = b'0'  # 启用
THOST_FTDC_OS_CheckIn = b'1'  # 签到
THOST_FTDC_OS_CheckOut = b'2'  # 签退
THOST_FTDC_OS_CheckFileArrived = b'3'  # 对帐文件到达
THOST_FTDC_OS_CheckDetail = b'4'  # 对帐
THOST_FTDC_OS_DayEndClean = b'5'  # 日终清理
THOST_FTDC_OS_Invalid = b'9'  # 注销

TThostFtdcOrganTypeType = ctypes.c_char

THOST_FTDC_OT_Bank = b'1'  # 银行代理
THOST_FTDC_OT_Future = b'2'  # 交易前置
THOST_FTDC_OT_PlateForm = b'9'  # 银期转帐平台管理

TThostFtdcParkedOrderStatusType = ctypes.c_char

THOST_FTDC_PAOS_NotSend = b'1'  # 未发送
THOST_FTDC_PAOS_Send = b'2'  # 已发送
THOST_FTDC_PAOS_Deleted = b'3'  # 已删除

TThostFtdcPassWordKeyTypeType = ctypes.c_char

THOST_FTDC_PWKT_ExchangeKey = b'0'  # 交换密钥
THOST_FTDC_PWKT_PassWordKey = b'1'  # 密码密钥
THOST_FTDC_PWKT_MACKey = b'2'  # MAC密钥
THOST_FTDC_PWKT_MessageKey = b'3'  # 报文密钥

TThostFtdcPasswordTypeType = ctypes.c_char

THOST_FTDC_PWDT_Trade = b'1'  # 交易密码
THOST_FTDC_PWDT_Account = b'2'  # 资金密码

TThostFtdcPersonTypeType = ctypes.c_char

THOST_FTDC_PST_Order = b'1'  # 指定下单人
THOST_FTDC_PST_Open = b'2'  # 开户授权人
THOST_FTDC_PST_Fund = b'3'  # 资金调拨人
THOST_FTDC_PST_Settlement = b'4'  # 结算单确认人
THOST_FTDC_PST_Company = b'5'  # 法人
THOST_FTDC_PST_Corporation = b'6'  # 法人代表
THOST_FTDC_PST_LinkMan = b'7'  # 投资者联系人
THOST_FTDC_PST_Ledger = b'8'  # 分户管理资产负责人
THOST_FTDC_PST_Trustee = b'9'  # 托（保）管人
THOST_FTDC_PST_TrusteeCorporation = b'A'  # 托（保）管机构法人代表
THOST_FTDC_PST_TrusteeOpen = b'B'  # 托（保）管机构开户授权人
THOST_FTDC_PST_TrusteeContact = b'C'  # 托（保）管机构联系人
THOST_FTDC_PST_ForeignerRefer = b'D'  # 境外自然人参考证件
THOST_FTDC_PST_CorporationRefer = b'E'  # 法人代表参考证件

TThostFtdcPortfTypeType = ctypes.c_char

THOST_FTDC_EET_None = b'0'  # 使用初版交易所算法
THOST_FTDC_EET_SPBM_AddOnHedge = b'1'  # SPBM算法V1.1.0_附加保证金调整

TThostFtdcPortfolioType = ctypes.c_char

THOST_FTDC_EPF_None = b'0'  # 不使用新型组保算法
THOST_FTDC_EPF_SPBM = b'1'  # SPBM算法
THOST_FTDC_EPF_RULE = b'2'  # RULE算法
THOST_FTDC_EPF_SPMM = b'3'  # SPMM算法
THOST_FTDC_EPF_RCAMS = b'4'  # RCAMS算法

TThostFtdcPosiDirectionType = ctypes.c_char

THOST_FTDC_PD_Net = b'1'  # 净
THOST_FTDC_PD_Long = b'2'  # 多头
THOST_FTDC_PD_Short = b'3'  # 空头

TThostFtdcPositionDateType = ctypes.c_char

THOST_FTDC_PSD_Today = b'1'  # 今日持仓
THOST_FTDC_PSD_History = b'2'  # 历史持仓

TThostFtdcPositionDateTypeType = ctypes.c_char

THOST_FTDC_PDT_UseHistory = b'1'  # 使用历史持仓
THOST_FTDC_PDT_NoUseHistory = b'2'  # 不使用历史持仓

TThostFtdcPositionTypeType = ctypes.c_char

THOST_FTDC_PT_Net = b'1'  # 净持仓
THOST_FTDC_PT_Gross = b'2'  # 综合持仓

TThostFtdcPriceSourceType = ctypes.c_char

THOST_FTDC_PSRC_LastPrice = b'0'  # 前成交价
THOST_FTDC_PSRC_Buy = b'1'  # 买委托价
THOST_FTDC_PSRC_Sell = b'2'  # 卖委托价
THOST_FTDC_PSRC_OTC = b'3'  # 场外成交价

TThostFtdcProcessStatusType = ctypes.c_char

THOST_FTDC_PSS_NotProcess = b'0'  # 未处理
THOST_FTDC_PSS_StartProcess = b'1'  # 开始处理
THOST_FTDC_PSS_Finished = b'2'  # 处理完成

TThostFtdcProdChangeFlagType = ctypes.c_char

THOST_FTDC_PCF_None = b'0'  # 持仓量和冻结量均无变化
THOST_FTDC_PCF_OnlyFrozen = b'1'  # 持仓量无变化，冻结量有变化
THOST_FTDC_PCF_PositionChange = b'2'  # 持仓量有变化

TThostFtdcProductClassType = ctypes.c_char

THOST_FTDC_PC_Futures = b'1'  # 期货
THOST_FTDC_PC_Options = b'2'  # 期货期权
THOST_FTDC_PC_Combination = b'3'  # 组合
THOST_FTDC_PC_Spot = b'4'  # 即期
THOST_FTDC_PC_EFP = b'5'  # 期转现
THOST_FTDC_PC_SpotOption = b'6'  # 现货期权
THOST_FTDC_PC_TAS = b'7'  # TAS合约
THOST_FTDC_PC_MI = b'I'  # 金属指数

TThostFtdcProductLifePhaseType = ctypes.c_char

THOST_FTDC_PLP_Active = b'1'  # 活跃
THOST_FTDC_PLP_NonActive = b'2'  # 不活跃
THOST_FTDC_PLP_Canceled = b'3'  # 注销

TThostFtdcProductStatusType = ctypes.c_char

THOST_FTDC_PS_tradeable = b'1'  # 可交易
THOST_FTDC_PS_untradeable = b'2'  # 不可交易

TThostFtdcProductTypeType = ctypes.c_char

THOST_FTDC_PTE_Futures = b'1'  # 期货
THOST_FTDC_PTE_Options = b'2'  # 期权

TThostFtdcPromptTypeType = ctypes.c_char

THOST_FTDC_CPT_Instrument = b'1'  # 合约上下市
THOST_FTDC_CPT_Margin = b'2'  # 保证金分段生效

TThostFtdcPropertyInvestorRangeType = ctypes.c_char

THOST_FTDC_PIR_All = b'1'  # 所有
THOST_FTDC_PIR_Property = b'2'  # 投资者属性
THOST_FTDC_PIR_Single = b'3'  # 单一投资者

TThostFtdcProtocalIDType = ctypes.c_char

THOST_FTDC_PID_FutureProtocal = b'0'  # 期商协议
THOST_FTDC_PID_ICBCProtocal = b'1'  # 工行协议
THOST_FTDC_PID_ABCProtocal = b'2'  # 农行协议
THOST_FTDC_PID_CBCProtocal = b'3'  # 中国银行协议
THOST_FTDC_PID_CCBProtocal = b'4'  # 建行协议
THOST_FTDC_PID_BOCOMProtocal = b'5'  # 交行协议
THOST_FTDC_PID_FBTPlateFormProtocal = b'X'  # 银期转帐平台协议

TThostFtdcPublishStatusType = ctypes.c_char

THOST_FTDC_PS_None = b'1'  # 未发布
THOST_FTDC_PS_Publishing = b'2'  # 正在发布
THOST_FTDC_PS_Published = b'3'  # 已发布

TThostFtdcPwdFlagType = ctypes.c_char

THOST_FTDC_BPWDF_NoCheck = b'0'  # 不核对
THOST_FTDC_BPWDF_BlankCheck = b'1'  # 明文核对
THOST_FTDC_BPWDF_EncryptCheck = b'2'  # 密文核对

TThostFtdcPwdRcdSrcType = ctypes.c_char

THOST_FTDC_PRS_Init = b'0'  # 来源于Sync初始化数据
THOST_FTDC_PRS_Sync = b'1'  # 来源于实时上场数据
THOST_FTDC_PRS_UserUpd = b'2'  # 来源于用户修改
THOST_FTDC_PRS_SuperUserUpd = b'3'  # 来源于超户修改，很可能来自主席同步数据

TThostFtdcQueryInvestorRangeType = ctypes.c_char

THOST_FTDC_QIR_All = b'1'  # 所有
THOST_FTDC_QIR_Group = b'2'  # 查询分类
THOST_FTDC_QIR_Single = b'3'  # 单一投资者

TThostFtdcQuestionTypeType = ctypes.c_char

THOST_FTDC_QT_Radio = b'1'  # 单选
THOST_FTDC_QT_Option = b'2'  # 多选
THOST_FTDC_QT_Blank = b'3'  # 填空

TThostFtdcRCAMSCombinationTypeType = ctypes.c_char

THOST_FTDC_ERComb_BUC = b'0'  # 牛市看涨价差组合
THOST_FTDC_ERComb_BEC = b'1'  # 熊市看涨价差组合
THOST_FTDC_ERComb_BEP = b'2'  # 熊市看跌价差组合
THOST_FTDC_ERComb_BUP = b'3'  # 牛市看跌价差组合
THOST_FTDC_ERComb_CAS = b'4'  # 日历价差组合

TThostFtdcRateInvestorRangeType = ctypes.c_char

THOST_FTDC_RIR_All = b'1'  # 公司标准
THOST_FTDC_RIR_Model = b'2'  # 模板
THOST_FTDC_RIR_Single = b'3'  # 单一投资者

TThostFtdcRateTypeType = ctypes.c_char

THOST_FTDC_RATETYPE_MarginRate = b'2'  # 保证金率

TThostFtdcRatioAttrType = ctypes.c_char

THOST_FTDC_RA_Trade = b'0'  # 交易费率
THOST_FTDC_RA_Settlement = b'1'  # 结算费率

TThostFtdcReasonType = ctypes.c_char

THOST_FTDC_RN_CD = b'0'  # 错单
THOST_FTDC_RN_ZT = b'1'  # 资金在途
THOST_FTDC_RN_QT = b'2'  # 其它

TThostFtdcReportStatusType = ctypes.c_char

THOST_FTDC_SRS_NoCreate = b'0'  # 未生成报表数据
THOST_FTDC_SRS_Create = b'1'  # 报表数据生成中
THOST_FTDC_SRS_Created = b'2'  # 已生成报表数据
THOST_FTDC_SRS_CreateFail = b'3'  # 生成报表数据失败

TThostFtdcReqFlagType = ctypes.c_char

THOST_FTDC_REQF_NoSend = b'0'  # 未发送
THOST_FTDC_REQF_SendSuccess = b'1'  # 发送成功
THOST_FTDC_REQF_SendFailed = b'2'  # 发送失败
THOST_FTDC_REQF_WaitReSend = b'3'  # 等待重发

TThostFtdcReqRspTypeType = ctypes.c_char

THOST_FTDC_REQRSP_Request = b'0'  # 请求
THOST_FTDC_REQRSP_Response = b'1'  # 响应

TThostFtdcResFlagType = ctypes.c_char

THOST_FTDC_RESF_Success = b'0'  # 成功
THOST_FTDC_RESF_InsuffiCient = b'1'  # 账户余额不足
THOST_FTDC_RESF_UnKnown = b'8'  # 交易结果未知

TThostFtdcReserveOpenAccStasType = ctypes.c_char

THOST_FTDC_ROAST_Processing = b'0'  # 等待处理中
THOST_FTDC_ROAST_Cancelled = b'1'  # 已撤销
THOST_FTDC_ROAST_Opened = b'2'  # 已开户
THOST_FTDC_ROAST_Invalid = b'3'  # 无效请求

TThostFtdcResponseValueType = ctypes.c_char

THOST_FTDC_RV_Right = b'0'  # 检查成功
THOST_FTDC_RV_Refuse = b'1'  # 检查失败

TThostFtdcReturnLevelType = ctypes.c_char

THOST_FTDC_RL_Level1 = b'1'  # 级别1
THOST_FTDC_RL_Level2 = b'2'  # 级别2
THOST_FTDC_RL_Level3 = b'3'  # 级别3
THOST_FTDC_RL_Level4 = b'4'  # 级别4
THOST_FTDC_RL_Level5 = b'5'  # 级别5
THOST_FTDC_RL_Level6 = b'6'  # 级别6
THOST_FTDC_RL_Level7 = b'7'  # 级别7
THOST_FTDC_RL_Level8 = b'8'  # 级别8
THOST_FTDC_RL_Level9 = b'9'  # 级别9

TThostFtdcReturnPatternType = ctypes.c_char

THOST_FTDC_RP_ByVolume = b'1'  # 按成交手数
THOST_FTDC_RP_ByFeeOnHand = b'2'  # 按留存手续费

TThostFtdcReturnStandardType = ctypes.c_char

THOST_FTDC_RSD_ByPeriod = b'1'  # 分阶段返还
THOST_FTDC_RSD_ByStandard = b'2'  # 按某一标准

TThostFtdcReturnStyleType = ctypes.c_char

THOST_FTDC_RS_All = b'1'  # 按所有品种
THOST_FTDC_RS_ByProduct = b'2'  # 按品种

TThostFtdcRightParamTypeType = ctypes.c_char

THOST_FTDC_RPT_Freeze = b'1'  # 休眠户
THOST_FTDC_RPT_FreezeActive = b'2'  # 激活休眠户
THOST_FTDC_RPT_OpenLimit = b'3'  # 开仓权限限制
THOST_FTDC_RPT_RelieveOpenLimit = b'4'  # 解除开仓权限限制

TThostFtdcRiskLevelType = ctypes.c_char

THOST_FTDC_FAS_Low = b'1'  # 低风险客户
THOST_FTDC_FAS_Normal = b'2'  # 普通客户
THOST_FTDC_FAS_Focus = b'3'  # 关注客户
THOST_FTDC_FAS_Risk = b'4'  # 风险客户

TThostFtdcRiskNotifyMethodType = ctypes.c_char

THOST_FTDC_RNM_System = b'0'  # 系统通知
THOST_FTDC_RNM_SMS = b'1'  # 短信通知
THOST_FTDC_RNM_EMail = b'2'  # 邮件通知
THOST_FTDC_RNM_Manual = b'3'  # 人工通知

TThostFtdcRiskNotifyStatusType = ctypes.c_char

THOST_FTDC_RNS_NotGen = b'0'  # 未生成
THOST_FTDC_RNS_Generated = b'1'  # 已生成未发送
THOST_FTDC_RNS_SendError = b'2'  # 发送失败
THOST_FTDC_RNS_SendOk = b'3'  # 已发送未接收
THOST_FTDC_RNS_Received = b'4'  # 已接收未确认
THOST_FTDC_RNS_Confirmed = b'5'  # 已确认

TThostFtdcRiskUserEventType = ctypes.c_char

THOST_FTDC_RUE_ExportData = b'0'  # 导出数据

TThostFtdcSHFEUploadFileNameType = ctypes.c_char

THOST_FTDC_SUFN_SUFN_O = b'O'  # ^\d{4}_\d{8}_\d{8}_DailyFundChg
THOST_FTDC_SUFN_SUFN_T = b'T'  # ^\d{4}_\d{8}_\d{8}_Trade
THOST_FTDC_SUFN_SUFN_P = b'P'  # ^\d{4}_\d{8}_\d{8}_SettlementDetail
THOST_FTDC_SUFN_SUFN_F = b'F'  # ^\d{4}_\d{8}_\d{8}_Capital

TThostFtdcSaveStatusType = ctypes.c_char

THOST_FTDC_SSS_UnSaveData = b'0'  # 归档未完成
THOST_FTDC_SSS_SaveDatad = b'1'  # 归档完成

TThostFtdcSecuAccTypeType = ctypes.c_char

THOST_FTDC_SAT_AccountID = b'1'  # 资金帐号
THOST_FTDC_SAT_CardID = b'2'  # 资金卡号
THOST_FTDC_SAT_SHStockholderID = b'3'  # 上海股东帐号
THOST_FTDC_SAT_SZStockholderID = b'4'  # 深圳股东帐号

TThostFtdcSendMethodType = ctypes.c_char

THOST_FTDC_UOASM_ByAPI = b'1'  # 文件发送
THOST_FTDC_UOASM_ByFile = b'2'  # 电子发送

TThostFtdcSendTypeType = ctypes.c_char

THOST_FTDC_UOAST_NoSend = b'0'  # 未发送
THOST_FTDC_UOAST_Sended = b'1'  # 已发送
THOST_FTDC_UOAST_Generated = b'2'  # 已生成
THOST_FTDC_UOAST_SendFail = b'3'  # 报送失败
THOST_FTDC_UOAST_Success = b'4'  # 接收成功
THOST_FTDC_UOAST_Fail = b'5'  # 接收失败
THOST_FTDC_UOAST_Cancel = b'6'  # 取消报送

TThostFtdcSettArchiveStatusType = ctypes.c_char

THOST_FTDC_SAS_UnArchived = b'0'  # 未归档数据
THOST_FTDC_SAS_Archiving = b'1'  # 数据归档中
THOST_FTDC_SAS_Archived = b'2'  # 已归档数据
THOST_FTDC_SAS_ArchiveFail = b'3'  # 归档数据失败

TThostFtdcSettleManagerGroupType = ctypes.c_char

THOST_FTDC_SMG_Exhcange = b'1'  # 交易所核对
THOST_FTDC_SMG_ASP = b'2'  # 内部核对
THOST_FTDC_SMG_CSRC = b'3'  # 上报数据核对

TThostFtdcSettleManagerLevelType = ctypes.c_char

THOST_FTDC_SML_Must = b'1'  # 必要
THOST_FTDC_SML_Alarm = b'2'  # 警告
THOST_FTDC_SML_Prompt = b'3'  # 提示
THOST_FTDC_SML_Ignore = b'4'  # 不检查

TThostFtdcSettleManagerTypeType = ctypes.c_char

THOST_FTDC_SMT_Before = b'1'  # 结算前准备
THOST_FTDC_SMT_Settlement = b'2'  # 结算
THOST_FTDC_SMT_After = b'3'  # 结算后核对
THOST_FTDC_SMT_Settlemented = b'4'  # 结算后处理

TThostFtdcSettlementBillTypeType = ctypes.c_char

THOST_FTDC_ST_Day = b'0'  # 日报
THOST_FTDC_ST_Month = b'1'  # 月报

TThostFtdcSettlementStatusType = ctypes.c_char

THOST_FTDC_STS_Initialize = b'0'  # 初始
THOST_FTDC_STS_Settlementing = b'1'  # 结算中
THOST_FTDC_STS_Settlemented = b'2'  # 已结算
THOST_FTDC_STS_Finished = b'3'  # 结算完成

TThostFtdcSettlementStyleType = ctypes.c_char

THOST_FTDC_SBS_Day = b'1'  # 逐日盯市
THOST_FTDC_SBS_Volume = b'2'  # 逐笔对冲

TThostFtdcSexType = ctypes.c_char

THOST_FTDC_SEX_None = b'0'  # 未知
THOST_FTDC_SEX_Man = b'1'  # 男
THOST_FTDC_SEX_Woman = b'2'  # 女

TThostFtdcSpecPosiTypeType = ctypes.c_char

THOST_FTDC_SPOST_Common = b'#'  # 普通持仓明细
THOST_FTDC_SPOST_Tas = b'0'  # TAS合约成交产生的标的合约持仓明细

TThostFtdcSpecProductTypeType = ctypes.c_char

THOST_FTDC_SPT_CzceHedge = b'1'  # 郑商所套保产品
THOST_FTDC_SPT_IneForeignCurrency = b'2'  # 货币质押产品
THOST_FTDC_SPT_DceOpenClose = b'3'  # 大连短线开平仓产品

TThostFtdcSpecialCreateRuleType = ctypes.c_char

THOST_FTDC_SC_NoSpecialRule = b'0'  # 没有特殊创建规则
THOST_FTDC_SC_NoSpringFestival = b'1'  # 不包含春节

TThostFtdcSponsorTypeType = ctypes.c_char

THOST_FTDC_SPTYPE_Broker = b'0'  # 期商
THOST_FTDC_SPTYPE_Bank = b'1'  # 银行

TThostFtdcStandardStatusType = ctypes.c_char

THOST_FTDC_STST_Standard = b'0'  # 已规范
THOST_FTDC_STST_NonStandard = b'1'  # 未规范

TThostFtdcStartModeType = ctypes.c_char

THOST_FTDC_SM_Normal = b'1'  # 正常
THOST_FTDC_SM_Emerge = b'2'  # 应急
THOST_FTDC_SM_Restore = b'3'  # 恢复

TThostFtdcStatModeType = ctypes.c_char

THOST_FTDC_SM_Non = b'0'  # ----
THOST_FTDC_SM_Instrument = b'1'  # 按合约统计
THOST_FTDC_SM_Product = b'2'  # 按产品统计
THOST_FTDC_SM_Investor = b'3'  # 按投资者统计

TThostFtdcStrikeModeType = ctypes.c_char

THOST_FTDC_STM_Continental = b'0'  # 欧式
THOST_FTDC_STM_American = b'1'  # 美式
THOST_FTDC_STM_Bermuda = b'2'  # 百慕大

TThostFtdcStrikeOffsetTypeType = ctypes.c_char

THOST_FTDC_STOV_RealValue = b'1'  # 实值额
THOST_FTDC_STOV_ProfitValue = b'2'  # 盈利额
THOST_FTDC_STOV_RealRatio = b'3'  # 实值比例
THOST_FTDC_STOV_ProfitRatio = b'4'  # 盈利比例

TThostFtdcStrikeTypeType = ctypes.c_char

THOST_FTDC_STT_Hedge = b'0'  # 自身对冲
THOST_FTDC_STT_Match = b'1'  # 匹配执行

TThostFtdcSwapSourceTypeType = ctypes.c_char

THOST_FTDC_SST_Manual = b'0'  # 手工
THOST_FTDC_SST_Automatic = b'1'  # 自动生成

TThostFtdcSyncDataStatusType = ctypes.c_char

THOST_FTDC_SDS_Initialize = b'0'  # 未同步
THOST_FTDC_SDS_Settlementing = b'1'  # 同步中
THOST_FTDC_SDS_Settlemented = b'2'  # 已同步

TThostFtdcSyncDeltaStatusType = ctypes.c_char

THOST_FTDC_SDS_Readable = b'1'  # 交易可读
THOST_FTDC_SDS_Reading = b'2'  # 交易在读
THOST_FTDC_SDS_Readend = b'3'  # 交易读取完成
THOST_FTDC_SDS_OptErr = b'e'  # 追平失败 交易本地状态结算不存在

TThostFtdcSyncFlagType = ctypes.c_char

THOST_FTDC_SYNF_Yes = b'0'  # 已同步
THOST_FTDC_SYNF_No = b'1'  # 未同步

TThostFtdcSyncModeType = ctypes.c_char

THOST_FTDC_SRM_ASync = b'0'  # 异步
THOST_FTDC_SRM_Sync = b'1'  # 同步

TThostFtdcSyncTypeType = ctypes.c_char

THOST_FTDC_SYNT_OneOffSync = b'0'  # 一次同步
THOST_FTDC_SYNT_TimerSync = b'1'  # 定时同步
THOST_FTDC_SYNT_TimerFullSync = b'2'  # 定时完全同步

TThostFtdcSysOperModeType = ctypes.c_char

THOST_FTDC_SoM_Add = b'1'  # 增加
THOST_FTDC_SoM_Update = b'2'  # 修改
THOST_FTDC_SoM_Delete = b'3'  # 删除
THOST_FTDC_SoM_Copy = b'4'  # 复制
THOST_FTDC_SoM_AcTive = b'5'  # 激活
THOST_FTDC_SoM_CanCel = b'6'  # 注销
THOST_FTDC_SoM_ReSet = b'7'  # 重置

TThostFtdcSysOperTypeType = ctypes.c_char

THOST_FTDC_SoT_UpdatePassword = b'0'  # 修改操作员密码
THOST_FTDC_SoT_UserDepartment = b'1'  # 操作员组织架构关系
THOST_FTDC_SoT_RoleManager = b'2'  # 角色管理
THOST_FTDC_SoT_RoleFunction = b'3'  # 角色功能设置
THOST_FTDC_SoT_BaseParam = b'4'  # 基础参数设置
THOST_FTDC_SoT_SetUserID = b'5'  # 设置操作员
THOST_FTDC_SoT_SetUserRole = b'6'  # 用户角色设置
THOST_FTDC_SoT_UserIpRestriction = b'7'  # 用户IP限制
THOST_FTDC_SoT_DepartmentManager = b'8'  # 组织架构管理
THOST_FTDC_SoT_DepartmentCopy = b'9'  # 组织架构向查询分类复制
THOST_FTDC_SoT_Tradingcode = b'A'  # 交易编码管理
THOST_FTDC_SoT_InvestorStatus = b'B'  # 投资者状态维护
THOST_FTDC_SoT_InvestorAuthority = b'C'  # 投资者权限管理
THOST_FTDC_SoT_PropertySet = b'D'  # 属性设置
THOST_FTDC_SoT_ReSetInvestorPasswd = b'E'  # 重置投资者密码
THOST_FTDC_SoT_InvestorPersonalityInfo = b'F'  # 投资者个性信息维护

TThostFtdcSysSettlementStatusType = ctypes.c_char

THOST_FTDC_SS_NonActive = b'1'  # 不活跃
THOST_FTDC_SS_Startup = b'2'  # 启动
THOST_FTDC_SS_Operating = b'3'  # 操作
THOST_FTDC_SS_Settlement = b'4'  # 结算
THOST_FTDC_SS_SettlementFinished = b'5'  # 结算完成

TThostFtdcSystemParamIDType = ctypes.c_char

THOST_FTDC_SPI_InvestorIDMinLength = b'1'  # 投资者代码最小长度
THOST_FTDC_SPI_AccountIDMinLength = b'2'  # 投资者帐号代码最小长度
THOST_FTDC_SPI_UserRightLogon = b'3'  # 投资者开户默认登录权限
THOST_FTDC_SPI_SettlementBillTrade = b'4'  # 投资者交易结算单成交汇总方式
THOST_FTDC_SPI_TradingCode = b'5'  # 统一开户更新交易编码方式
THOST_FTDC_SPI_CheckFund = b'6'  # 结算是否判断存在未复核的出入金和分项资金
THOST_FTDC_SPI_CommModelRight = b'7'  # 是否启用手续费模板数据权限
THOST_FTDC_SPI_MarginModelRight = b'9'  # 是否启用保证金率模板数据权限
THOST_FTDC_SPI_IsStandardActive = b'8'  # 是否规范用户才能激活
THOST_FTDC_SPI_UploadSettlementFile = b'U'  # 上传的交易所结算文件路径
THOST_FTDC_SPI_DownloadCSRCFile = b'D'  # 上报保证金监控中心文件路径
THOST_FTDC_SPI_SettlementBillFile = b'S'  # 生成的结算单文件路径
THOST_FTDC_SPI_CSRCOthersFile = b'C'  # 证监会文件标识
THOST_FTDC_SPI_InvestorPhoto = b'P'  # 投资者照片路径
THOST_FTDC_SPI_CSRCData = b'R'  # 全结经纪公司上传文件路径
THOST_FTDC_SPI_InvestorPwdModel = b'I'  # 开户密码录入方式
THOST_FTDC_SPI_CFFEXInvestorSettleFile = b'F'  # 投资者中金所结算文件下载路径
THOST_FTDC_SPI_InvestorIDType = b'a'  # 投资者代码编码方式
THOST_FTDC_SPI_FreezeMaxReMain = b'r'  # 休眠户最高权益
THOST_FTDC_SPI_IsSync = b'A'  # 手续费相关操作实时上场开关
THOST_FTDC_SPI_RelieveOpenLimit = b'O'  # 解除开仓权限限制
THOST_FTDC_SPI_IsStandardFreeze = b'X'  # 是否规范用户才能休眠
THOST_FTDC_SPI_CZCENormalProductHedge = b'B'  # 郑商所是否开放所有品种套保交易

TThostFtdcSystemStatusType = ctypes.c_char

THOST_FTDC_ES_NonActive = b'1'  # 不活跃
THOST_FTDC_ES_Startup = b'2'  # 启动
THOST_FTDC_ES_Initialize = b'3'  # 交易开始初始化
THOST_FTDC_ES_Initialized = b'4'  # 交易完成初始化
THOST_FTDC_ES_Close = b'5'  # 收市开始
THOST_FTDC_ES_Closed = b'6'  # 收市完成
THOST_FTDC_ES_Settlement = b'7'  # 结算

TThostFtdcSystemTypeType = ctypes.c_char

THOST_FTDC_SYT_FutureBankTransfer = b'0'  # 银期转帐
THOST_FTDC_SYT_StockBankTransfer = b'1'  # 银证转帐
THOST_FTDC_SYT_TheThirdPartStore = b'2'  # 第三方存管

TThostFtdcTemplateTypeType = ctypes.c_char

THOST_FTDC_TPT_Full = b'1'  # 全量
THOST_FTDC_TPT_Increment = b'2'  # 增量
THOST_FTDC_TPT_BackUp = b'3'  # 备份

TThostFtdcTimeConditionType = ctypes.c_char

THOST_FTDC_TC_IOC = b'1'  # 立即完成，否则撤销
THOST_FTDC_TC_GFS = b'2'  # 本节有效
THOST_FTDC_TC_GFD = b'3'  # 当日有效
THOST_FTDC_TC_GTD = b'4'  # 指定日期前有效
THOST_FTDC_TC_GTC = b'5'  # 撤销前有效
THOST_FTDC_TC_GFA = b'6'  # 集合竞价有效

TThostFtdcTimeRangeType = ctypes.c_char

THOST_FTDC_ETR_USUAL = b'1'  # 一般月份
THOST_FTDC_ETR_FNSP = b'2'  # 交割月前一个月上半月
THOST_FTDC_ETR_BNSP = b'3'  # 交割月前一个月下半月
THOST_FTDC_ETR_SPOT = b'4'  # 交割月份

TThostFtdcTradeParamIDType = ctypes.c_char

THOST_FTDC_TPID_EncryptionStandard = b'E'  # 系统加密算法
THOST_FTDC_TPID_RiskMode = b'R'  # 系统风险算法
THOST_FTDC_TPID_RiskModeGlobal = b'G'  # 系统风险算法是否全局 0-否 1-是
THOST_FTDC_TPID_modeEncode = b'P'  # 密码加密算法
THOST_FTDC_TPID_tickMode = b'T'  # 价格小数位数参数
THOST_FTDC_TPID_SingleUserSessionMaxNum = b'S'  # 用户最大会话数
THOST_FTDC_TPID_LoginFailMaxNum = b'L'  # 最大连续登录失败数
THOST_FTDC_TPID_IsAuthForce = b'A'  # 是否强制认证
THOST_FTDC_TPID_IsPosiFreeze = b'F'  # 是否冻结证券持仓
THOST_FTDC_TPID_IsPosiLimit = b'M'  # 是否限仓
THOST_FTDC_TPID_ForQuoteTimeInterval = b'Q'  # 郑商所询价时间间隔
THOST_FTDC_TPID_IsFuturePosiLimit = b'B'  # 是否期货限仓
THOST_FTDC_TPID_IsFutureOrderFreq = b'C'  # 是否期货下单频率限制
THOST_FTDC_TPID_IsExecOrderProfit = b'H'  # 行权冻结是否计算盈利
THOST_FTDC_TPID_IsCheckBankAcc = b'I'  # 银期开户是否验证开户银行卡号是否是预留银行账户
THOST_FTDC_TPID_PasswordDeadLine = b'J'  # 弱密码最后修改日期
THOST_FTDC_TPID_IsStrongPassword = b'K'  # 强密码校验
THOST_FTDC_TPID_BalanceMorgage = b'a'  # 自有资金质押比
THOST_FTDC_TPID_MinPwdLen = b'O'  # 最小密码长度
THOST_FTDC_TPID_LoginFailMaxNumForIP = b'U'  # IP当日最大登陆失败次数
THOST_FTDC_TPID_PasswordPeriod = b'V'  # 密码有效期
THOST_FTDC_TPID_PwdHistoryCmp = b'X'  # 历史密码重复限制次数

TThostFtdcTradeSourceType = ctypes.c_char

THOST_FTDC_TSRC_NORMAL = b'0'  # 来自交易所普通回报
THOST_FTDC_TSRC_QUERY = b'1'  # 来自查询

TThostFtdcTradeSumStatModeType = ctypes.c_char

THOST_FTDC_TSSM_Instrument = b'1'  # 按合约统计
THOST_FTDC_TSSM_Product = b'2'  # 按产品统计
THOST_FTDC_TSSM_Exchange = b'3'  # 按交易所统计

TThostFtdcTradeTypeType = ctypes.c_char

THOST_FTDC_TRDT_SplitCombination = b'#'  # 组合持仓拆分为单一持仓,初始化不应包含该类型的持仓
THOST_FTDC_TRDT_Common = b'0'  # 普通成交
THOST_FTDC_TRDT_OptionsExecution = b'1'  # 期权执行
THOST_FTDC_TRDT_OTC = b'2'  # OTC成交
THOST_FTDC_TRDT_EFPDerived = b'3'  # 期转现衍生成交
THOST_FTDC_TRDT_CombinationDerived = b'4'  # 组合衍生成交
THOST_FTDC_TRDT_BlockTrade = b'5'  # 大宗交易成交

TThostFtdcTraderConnectStatusType = ctypes.c_char

THOST_FTDC_TCS_NotConnected = b'1'  # 没有任何连接
THOST_FTDC_TCS_Connected = b'2'  # 已经连接
THOST_FTDC_TCS_QryInstrumentSent = b'3'  # 已经发出合约查询请求
THOST_FTDC_TCS_SubPrivateFlow = b'4'  # 订阅私有流

TThostFtdcTradingRightType = ctypes.c_char

THOST_FTDC_TR_Allow = b'0'  # 可以交易
THOST_FTDC_TR_CloseOnly = b'1'  # 只能平仓
THOST_FTDC_TR_Forbidden = b'2'  # 不能交易

TThostFtdcTradingRoleType = ctypes.c_char

THOST_FTDC_ER_Broker = b'1'  # 代理
THOST_FTDC_ER_Host = b'2'  # 自营
THOST_FTDC_ER_Maker = b'3'  # 做市商

TThostFtdcTradingTypeType = ctypes.c_char

THOST_FTDC_TD_ALL = b'0'  # 所有状态
THOST_FTDC_TD_TRADE = b'1'  # 交易
THOST_FTDC_TD_UNTRADE = b'2'  # 非交易

TThostFtdcTransferDirectionType = ctypes.c_char

THOST_FTDC_TD_Out = b'0'  # 移出
THOST_FTDC_TD_In = b'1'  # 移入

TThostFtdcTransferStatusType = ctypes.c_char

THOST_FTDC_TRFS_Normal = b'0'  # 正常
THOST_FTDC_TRFS_Repealed = b'1'  # 被冲正

TThostFtdcTransferTypeType = ctypes.c_char

THOST_FTDC_TT_BankToFuture = b'0'  # 银行转期货
THOST_FTDC_TT_FutureToBank = b'1'  # 期货转银行

TThostFtdcTransferValidFlagType = ctypes.c_char

THOST_FTDC_TVF_Invalid = b'0'  # 无效或失败
THOST_FTDC_TVF_Valid = b'1'  # 有效
THOST_FTDC_TVF_Reverse = b'2'  # 冲正

TThostFtdcTxnEndFlagType = ctypes.c_char

THOST_FTDC_TEF_NormalProcessing = b'0'  # 正常处理中
THOST_FTDC_TEF_Success = b'1'  # 成功结束
THOST_FTDC_TEF_Failed = b'2'  # 失败结束
THOST_FTDC_TEF_Abnormal = b'3'  # 异常中
THOST_FTDC_TEF_ManualProcessedForException = b'4'  # 已人工异常处理
THOST_FTDC_TEF_CommuFailedNeedManualProcess = b'5'  # 通讯异常 ，请人工处理
THOST_FTDC_TEF_SysErrorNeedManualProcess = b'6'  # 系统出错，请人工处理

TThostFtdcUOAAssetmgrTypeType = ctypes.c_char

THOST_FTDC_UOAAT_Futures = b'1'  # 期货类
THOST_FTDC_UOAAT_SpecialOrgan = b'2'  # 综合类

TThostFtdcUOAAutoSendType = ctypes.c_char

THOST_FTDC_UOAA_ASR = b'1'  # 自动发送并接收
THOST_FTDC_UOAA_ASNR = b'2'  # 自动发送，不自动接收
THOST_FTDC_UOAA_NSAR = b'3'  # 不自动发送，自动接收
THOST_FTDC_UOAA_NSR = b'4'  # 不自动发送，也不自动接收

TThostFtdcUpdateFlagType = ctypes.c_char

THOST_FTDC_UF_NoUpdate = b'0'  # 未更新
THOST_FTDC_UF_Success = b'1'  # 更新全部信息成功
THOST_FTDC_UF_Fail = b'2'  # 更新全部信息失败
THOST_FTDC_UF_TCSuccess = b'3'  # 更新交易编码成功
THOST_FTDC_UF_TCFail = b'4'  # 更新交易编码失败
THOST_FTDC_UF_Cancel = b'5'  # 已丢弃

TThostFtdcUsedStatusType = ctypes.c_char

THOST_FTDC_CHU_Unused = b'0'  # 未生效
THOST_FTDC_CHU_Used = b'1'  # 已生效
THOST_FTDC_CHU_Fail = b'2'  # 生效失败

TThostFtdcUserEventTypeType = ctypes.c_char

THOST_FTDC_UET_Login = b'1'  # 登录
THOST_FTDC_UET_Logout = b'2'  # 登出
THOST_FTDC_UET_Trading = b'3'  # CTP校验通过
THOST_FTDC_UET_TradingError = b'4'  # CTP校验失败
THOST_FTDC_UET_UpdatePassword = b'5'  # 修改密码
THOST_FTDC_UET_Authenticate = b'6'  # 客户端认证
THOST_FTDC_UET_SubmitSysInfo = b'7'  # 终端信息上报
THOST_FTDC_UET_Transfer = b'8'  # 转账
THOST_FTDC_UET_Other = b'9'  # 其他
THOST_FTDC_UET_UpdateTradingAccountPassword = b'a'  # 修改资金密码

TThostFtdcUserRangeType = ctypes.c_char

THOST_FTDC_UR_All = b'0'  # 所有
THOST_FTDC_UR_Single = b'1'  # 单一操作员

TThostFtdcUserRightTypeType = ctypes.c_char

THOST_FTDC_URT_Logon = b'1'  # 登录
THOST_FTDC_URT_Transfer = b'2'  # 银期转帐
THOST_FTDC_URT_EMail = b'3'  # 邮寄结算单
THOST_FTDC_URT_Fax = b'4'  # 传真结算单
THOST_FTDC_URT_ConditionOrder = b'5'  # 条件单

TThostFtdcUserTypeType = ctypes.c_char

THOST_FTDC_UT_Investor = b'0'  # 投资者
THOST_FTDC_UT_Operator = b'1'  # 操作员
THOST_FTDC_UT_SuperUser = b'2'  # 管理员

TThostFtdcValueMethodType = ctypes.c_char

THOST_FTDC_VM_Absolute = b'0'  # 按绝对值
THOST_FTDC_VM_Ratio = b'1'  # 按比率

TThostFtdcVirBankAccTypeType = ctypes.c_char

THOST_FTDC_VBAT_BankBook = b'1'  # 存折
THOST_FTDC_VBAT_BankCard = b'2'  # 储蓄卡
THOST_FTDC_VBAT_CreditCard = b'3'  # 信用卡

TThostFtdcVirDealStatusType = ctypes.c_char

THOST_FTDC_VDS_Dealing = b'1'  # 正在处理
THOST_FTDC_VDS_DeaclSucceed = b'2'  # 处理成功

TThostFtdcVirTradeStatusType = ctypes.c_char

THOST_FTDC_VTS_NaturalDeal = b'0'  # 正常处理中
THOST_FTDC_VTS_SucceedEnd = b'1'  # 成功结束
THOST_FTDC_VTS_FailedEND = b'2'  # 失败结束
THOST_FTDC_VTS_Exception = b'3'  # 异常中
THOST_FTDC_VTS_ManualDeal = b'4'  # 已人工异常处理
THOST_FTDC_VTS_MesException = b'5'  # 通讯异常 ，请人工处理
THOST_FTDC_VTS_SysException = b'6'  # 系统出错，请人工处理

TThostFtdcVirementAvailAbilityType = ctypes.c_char

THOST_FTDC_VAA_NoAvailAbility = b'0'  # 未确认
THOST_FTDC_VAA_AvailAbility = b'1'  # 有效
THOST_FTDC_VAA_Repeal = b'2'  # 冲正

TThostFtdcVirementStatusType = ctypes.c_char

THOST_FTDC_VMS_Natural = b'0'  # 正常
THOST_FTDC_VMS_Canceled = b'9'  # 销户

TThostFtdcVirementTradeCodeType = str

THOST_FTDC_VTC_BankBankToFuture = "102001"  # 银行发起银行资金转期货
THOST_FTDC_VTC_BankFutureToBank = "102002"  # 银行发起期货资金转银行
THOST_FTDC_VTC_FutureBankToFuture = "202001"  # 期货发起银行资金转期货
THOST_FTDC_VTC_FutureFutureToBank = "202002"  # 期货发起期货资金转银行

TThostFtdcVolumeConditionType = ctypes.c_char

THOST_FTDC_VC_AV = b'1'  # 任何数量
THOST_FTDC_VC_MV = b'2'  # 最小数量
THOST_FTDC_VC_CV = b'3'  # 全部数量

TThostFtdcWeakPasswordSourceType = ctypes.c_char

THOST_FTDC_WPSR_Lib = b'1'  # 弱密码库
THOST_FTDC_WPSR_Manual = b'2'  # 手工录入

TThostFtdcWithDrawParamIDType = ctypes.c_char

THOST_FTDC_WDPID_CashIn = b'C'  # 权利金收支是否可提 1 代表可提 0 不可提

TThostFtdcYesNoIndicatorType = ctypes.c_char

THOST_FTDC_YNI_Yes = b'0'  # 是
THOST_FTDC_YNI_No = b'1'  # 否

# ----- 单字符类型 -----

# TThostFtdcNewsUrgencyType 紧急程度类型
TThostFtdcNewsUrgencyType = ctypes.c_char
