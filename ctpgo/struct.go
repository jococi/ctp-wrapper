package ctpgo

// 此文件由代码生成器自动生成，请勿手动修改
// CTP 结构体定义

// ========== CTP 结构体 ==========

// CThostFtdcAccountPropertyField 银行账户属性
type CThostFtdcAccountPropertyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	BankID TThostFtdcBankIDType // 银行统一标识类型
	BankAccount TThostFtdcBankAccountType // 银行账户
	OpenName TThostFtdcInvestorFullNameType // 银行账户的开户人名称
	OpenBank TThostFtdcOpenBankType // 银行账户的开户行
	IsActive TThostFtdcBoolType // 是否活跃
	AccountSourceType TThostFtdcAccountSourceTypeType // 账户来源
	OpenDate TThostFtdcDateType // 开户日期
	CancelDate TThostFtdcDateType // 注销日期
	OperatorID TThostFtdcOperatorIDType // 录入员代码
	OperateDate TThostFtdcDateType // 录入日期
	OperateTime TThostFtdcTimeType // 录入时间
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcAccountregisterField 客户开销户信息表
type CThostFtdcAccountregisterField struct {
	TradeDay TThostFtdcTradeDateType // 交易日期
	BankID TThostFtdcBankIDType // 银行编码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构编码
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BrokerID TThostFtdcBrokerIDType // 期货公司编码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期货公司分支机构编码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	OpenOrDestroy TThostFtdcOpenOrDestroyType // 开销户类别
	RegDate TThostFtdcTradeDateType // 签约日期
	OutDate TThostFtdcTradeDateType // 解约日期
	TID TThostFtdcTIDType // 交易ID
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcAppIDAuthAssignField App客户端权限分配
type CThostFtdcAppIDAuthAssignField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AppID TThostFtdcAppIDType // App代码
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcAuthForbiddenIPField 禁止认证IP
type CThostFtdcAuthForbiddenIPField struct {
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcAuthIPField 用户IP绑定信息
type CThostFtdcAuthIPField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AppID TThostFtdcAppIDType // App代码
	IPAddress TThostFtdcIPAddressType // 用户代码
}

// CThostFtdcAuthUserIDField 终端用户绑定信息
type CThostFtdcAuthUserIDField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AppID TThostFtdcAppIDType // App代码
	UserID TThostFtdcUserIDType // 用户代码
	AuthType TThostFtdcAuthTypeType // 校验类型
}

// CThostFtdcAuthenticationInfoField 客户端认证信息
type CThostFtdcAuthenticationInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	AuthInfo TThostFtdcAuthInfoType // 认证信息
	IsResult TThostFtdcBoolType // 是否为认证结果
	AppID TThostFtdcAppIDType // App代码
	AppType TThostFtdcAppTypeType // App类型
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
}

// CThostFtdcBatchOrderActionField 批量报单操作
type CThostFtdcBatchOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcBrokerDepositField 经纪公司资金
type CThostFtdcBrokerDepositField struct {
	TradingDay TThostFtdcTradeDateType // 交易日期
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	PreBalance TThostFtdcMoneyType // 上次结算准备金
	CurrMargin TThostFtdcMoneyType // 当前保证金总额
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	Balance TThostFtdcMoneyType // 期货结算准备金
	Deposit TThostFtdcMoneyType // 入金金额
	Withdraw TThostFtdcMoneyType // 出金金额
	Available TThostFtdcMoneyType // 可提资金
	Reserve TThostFtdcMoneyType // 基本准备金
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
}

// CThostFtdcBrokerField 经纪公司
type CThostFtdcBrokerField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	BrokerAbbr TThostFtdcBrokerAbbrType // 经纪公司简称
	BrokerName TThostFtdcBrokerNameType // 经纪公司名称
	IsActive TThostFtdcBoolType // 是否活跃
}

// CThostFtdcBrokerSyncField 经纪公司同步
type CThostFtdcBrokerSyncField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcBrokerTradingAlgosField 经纪公司交易算法
type CThostFtdcBrokerTradingAlgosField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HandlePositionAlgoID TThostFtdcHandlePositionAlgoIDType // 持仓处理算法编号
	FindMarginRateAlgoID TThostFtdcFindMarginRateAlgoIDType // 寻找保证金率算法编号
	HandleTradingAccountAlgoID TThostFtdcHandleTradingAccountAlgoIDType // 资金处理算法编号
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcBrokerTradingParamsField 经纪公司交易参数
type CThostFtdcBrokerTradingParamsField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	MarginPriceType TThostFtdcMarginPriceTypeType // 保证金价格类型
	Algorithm TThostFtdcAlgorithmType // 盈亏算法
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType // 可用是否包含平仓盈利
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	OptionRoyaltyPriceType TThostFtdcOptionRoyaltyPriceTypeType // 期权权利金价格类型
	AccountID TThostFtdcAccountIDType // 投资者帐号
}

// CThostFtdcBrokerUserEventField 查询经纪公司用户事件
type CThostFtdcBrokerUserEventField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserEventType TThostFtdcUserEventTypeType // 用户事件类型
	EventSequenceNo TThostFtdcSequenceNoType // 用户事件序号
	EventDate TThostFtdcDateType // 事件发生日期
	EventTime TThostFtdcTimeType // 事件发生时间
	UserEventInfo TThostFtdcUserEventInfoType // 用户事件信息
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
	TradingDay TThostFtdcDateType // 交易日
}

// CThostFtdcBrokerUserField 经纪公司用户
type CThostFtdcBrokerUserField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserName TThostFtdcUserNameType // 用户名称
	UserType TThostFtdcUserTypeType // 用户类型
	IsActive TThostFtdcBoolType // 是否活跃
	IsUsingOTP TThostFtdcBoolType // 是否使用令牌
	IsAuthForce TThostFtdcBoolType // 是否强制终端认证
}

// CThostFtdcBrokerUserFunctionField 经纪公司用户功能权限
type CThostFtdcBrokerUserFunctionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	BrokerFunctionCode TThostFtdcBrokerFunctionCodeType // 经纪公司功能代码
}

// CThostFtdcBrokerUserOTPParamField 用户动态令牌参数
type CThostFtdcBrokerUserOTPParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	OTPVendorsID TThostFtdcOTPVendorsIDType // 动态令牌提供商
	SerialNumber TThostFtdcSerialNumberType // 动态令牌序列号
	AuthKey TThostFtdcAuthKeyType // 令牌密钥
	LastDrift TThostFtdcLastDriftType // 漂移值
	LastSuccess TThostFtdcLastSuccessType // 成功值
	OTPType TThostFtdcOTPTypeType // 动态令牌类型
}

// CThostFtdcBrokerUserPasswordField 经纪公司用户口令
type CThostFtdcBrokerUserPasswordField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	LastUpdateTime TThostFtdcDateTimeType // 上次修改时间
	LastLoginTime TThostFtdcDateTimeType // 上次登陆时间
	ExpireDate TThostFtdcDateType // 密码过期时间
	WeakExpireDate TThostFtdcDateType // 弱密码过期时间
}

// CThostFtdcBrokerUserRightAssignField 经济公司是否有在本标示的交易权限
type CThostFtdcBrokerUserRightAssignField struct {
	BrokerID TThostFtdcBrokerIDType // 应用单元代码
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
	Tradeable TThostFtdcBoolType // 能否交易
}

// CThostFtdcBrokerWithdrawAlgorithmField 经纪公司可提资金算法表
type CThostFtdcBrokerWithdrawAlgorithmField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	WithdrawAlgorithm TThostFtdcAlgorithmType // 可提资金算法
	UsingRatio TThostFtdcRatioType // 资金使用率
	IncludeCloseProfit TThostFtdcIncludeCloseProfitType // 可提是否包含平仓盈利
	AllWithoutTrade TThostFtdcAllWithoutTradeType // 本日无仓且无成交客户是否受可提比例限制
	AvailIncludeCloseProfit TThostFtdcIncludeCloseProfitType // 可用是否包含平仓盈利
	IsBrokerUserEvent TThostFtdcBoolType // 是否启用用户事件
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	FundMortgageRatio TThostFtdcRatioType // 货币质押比率
	BalanceAlgorithm TThostFtdcBalanceAlgorithmType // 权益算法
}

// CThostFtdcBulletinField 交易所公告
type CThostFtdcBulletinField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TradingDay TThostFtdcDateType // 交易日
	BulletinID TThostFtdcBulletinIDType // 公告编号
	SequenceNo TThostFtdcSequenceNoType // 序列号
	NewsType TThostFtdcNewsTypeType // 公告类型
	NewsUrgency TThostFtdcNewsUrgencyType // 紧急程度
	SendTime TThostFtdcTimeType // 发送时间
	Abstract TThostFtdcAbstractType // 消息摘要
	ComeFrom TThostFtdcComeFromType // 消息来源
	Content TThostFtdcContentType // 消息正文
	URLLink TThostFtdcURLLinkType // WEB地址
	MarketID TThostFtdcMarketIDType // 市场代码
}

// CThostFtdcCFMMCBrokerKeyField 保证金监管系统经纪公司密钥
type CThostFtdcCFMMCBrokerKeyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType // 经纪公司统一编码
	CreateDate TThostFtdcDateType // 密钥生成日期
	CreateTime TThostFtdcTimeType // 密钥生成时间
	KeyID TThostFtdcSequenceNoType // 密钥编号
	CurrentKey TThostFtdcCFMMCKeyType // 动态密钥
	KeyKind TThostFtdcCFMMCKeyKindType // 动态密钥类型
}

// CThostFtdcCFMMCTradingAccountKeyField 保证金监管系统经纪公司资金账户密钥
type CThostFtdcCFMMCTradingAccountKeyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType // 经纪公司统一编码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	KeyID TThostFtdcSequenceNoType // 密钥编号
	CurrentKey TThostFtdcCFMMCKeyType // 动态密钥
}

// CThostFtdcCFMMCTradingAccountTokenField 监控中心用户令牌
type CThostFtdcCFMMCTradingAccountTokenField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType // 经纪公司统一编码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	KeyID TThostFtdcSequenceNoType // 密钥编号
	Token TThostFtdcCFMMCTokenType // 动态令牌
}

// CThostFtdcCancelAccountField 银期销户信息
type CThostFtdcCancelAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType // 汇钞标志
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	TID TThostFtdcTIDType // 交易ID
	UserID TThostFtdcUserIDType // 用户标识
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcChangeAccountField 银期变更银行账号信息
type CThostFtdcChangeAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	NewBankAccount TThostFtdcBankAccountType // 新银行帐号
	NewBankPassWord TThostFtdcPasswordType // 新银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	TID TThostFtdcTIDType // 交易ID
	Digest TThostFtdcDigestType // 摘要
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcCombActionField 申请组合
type CThostFtdcCombActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	CombActionRef TThostFtdcOrderRefType // 组合引用
	UserID TThostFtdcUserIDType // 用户代码
	Direction TThostFtdcDirectionType // 买卖方向
	Volume TThostFtdcVolumeType // 数量
	CombDirection TThostFtdcCombDirectionType // 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionLocalID TThostFtdcOrderLocalIDType // 本地申请组合编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ActionStatus TThostFtdcOrderActionStatusType // 组合状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ComTradeID TThostFtdcTradeIDType // 组合编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcCombInstrumentGuardField 组合合约安全系数
type CThostFtdcCombInstrumentGuardField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	GuarantRatio TThostFtdcRatioType
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcCombPromotionParamField 组合优惠比例
type CThostFtdcCombPromotionParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 投机套保标志
	Xparameter TThostFtdcDiscountRatioType // 期权组合保证金比例
}

// CThostFtdcCombinationLegField 组合交易合约的单腿
type CThostFtdcCombinationLegField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	LegID TThostFtdcLegIDType // 单腿编号
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	Direction TThostFtdcDirectionType // 买卖方向
	LegMultiple TThostFtdcLegMultipleType // 单腿乘数
	ImplyLevel TThostFtdcImplyLevelType // 派生层数
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
	LegInstrumentID TThostFtdcInstrumentIDType // 单腿合约代码
}

// CThostFtdcCommPhaseField 通讯阶段
type CThostFtdcCommPhaseField struct {
	TradingDay TThostFtdcDateType // 交易日
	CommPhaseNo TThostFtdcCommPhaseNoType // 通讯时段编号
	SystemID TThostFtdcSystemIDType // 系统编号
}

// CThostFtdcCommRateModelField 投资者手续费率模板
type CThostFtdcCommRateModelField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	CommModelID TThostFtdcInvestorIDType // 手续费率模板代码
	CommModelName TThostFtdcCommModelNameType // 模板名称
}

// CThostFtdcContractBankField 查询签约银行响应
type CThostFtdcContractBankField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	BankID TThostFtdcBankIDType // 银行代码
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码
	BankName TThostFtdcBankNameType // 银行名称
}

// CThostFtdcCurrDRIdentityField 当前交易中心
type CThostFtdcCurrDRIdentityField struct {
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcCurrTransferIdentityField 当前银期所属交易中心
type CThostFtdcCurrTransferIdentityField struct {
	IdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcCurrentTimeField 当前时间
type CThostFtdcCurrentTimeField struct {
	CurrDate TThostFtdcDateType // 当前交易日
	CurrTime TThostFtdcTimeType // 当前时间
	CurrMillisec TThostFtdcMillisecType // 当前时间（毫秒）
	ActionDay TThostFtdcDateType // 自然日期
}

// CThostFtdcDRTransferField 灾备交易转换报文
type CThostFtdcDRTransferField struct {
	OrigDRIdentityID TThostFtdcDRIdentityIDType // 原交易中心代码
	DestDRIdentityID TThostFtdcDRIdentityIDType // 目标交易中心代码
	OrigBrokerID TThostFtdcBrokerIDType // 原应用单元代码
	DestBrokerID TThostFtdcBrokerIDType // 目标易用单元代码
}

// CThostFtdcDepartmentUserField 操作员组织架构关系
type CThostFtdcDepartmentUserField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	InvestorRange TThostFtdcDepartmentRangeType // 投资者范围
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcDepositResultInformField 验证期货资金密码和客户信息
type CThostFtdcDepositResultInformField struct {
	DepositSeqNo TThostFtdcDepositSeqNoType // 出入金流水号，该流水号为银期报盘返回的流水号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Deposit TThostFtdcMoneyType // 入金金额
	RequestID TThostFtdcRequestIDType // 请求编号
	ReturnCode TThostFtdcReturnCodeType // 返回代码
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType // 返回码描述
}

// CThostFtdcDepthMarketDataField 深度行情
type CThostFtdcDepthMarketDataField struct {
	TradingDay TThostFtdcDateType // 交易日
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	LastPrice TThostFtdcPriceType // 最新价
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	PreClosePrice TThostFtdcPriceType // 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType // 昨持仓量
	OpenPrice TThostFtdcPriceType // 今开盘
	HighestPrice TThostFtdcPriceType // 最高价
	LowestPrice TThostFtdcPriceType // 最低价
	Volume TThostFtdcVolumeType // 数量
	Turnover TThostFtdcMoneyType // 成交金额
	OpenInterest TThostFtdcLargeVolumeType // 持仓量
	ClosePrice TThostFtdcPriceType // 今收盘
	SettlementPrice TThostFtdcPriceType // 本次结算价
	UpperLimitPrice TThostFtdcPriceType // 涨停板价
	LowerLimitPrice TThostFtdcPriceType // 跌停板价
	PreDelta TThostFtdcRatioType // 昨虚实度
	CurrDelta TThostFtdcRatioType // 今虚实度
	UpdateTime TThostFtdcTimeType // 最后修改时间
	UpdateMillisec TThostFtdcMillisecType // 最后修改毫秒
	BidPrice1 TThostFtdcPriceType // 申买价一
	BidVolume1 TThostFtdcVolumeType // 申买量一
	AskPrice1 TThostFtdcPriceType // 申卖价一
	AskVolume1 TThostFtdcVolumeType // 申卖量一
	BidPrice2 TThostFtdcPriceType // 申买价二
	BidVolume2 TThostFtdcVolumeType // 申买量二
	AskPrice2 TThostFtdcPriceType // 申卖价二
	AskVolume2 TThostFtdcVolumeType // 申卖量二
	BidPrice3 TThostFtdcPriceType // 申买价三
	BidVolume3 TThostFtdcVolumeType // 申买量三
	AskPrice3 TThostFtdcPriceType // 申卖价三
	AskVolume3 TThostFtdcVolumeType // 申卖量三
	BidPrice4 TThostFtdcPriceType // 申买价四
	BidVolume4 TThostFtdcVolumeType // 申买量四
	AskPrice4 TThostFtdcPriceType // 申卖价四
	AskVolume4 TThostFtdcVolumeType // 申卖量四
	BidPrice5 TThostFtdcPriceType // 申买价五
	BidVolume5 TThostFtdcVolumeType // 申买量五
	AskPrice5 TThostFtdcPriceType // 申卖价五
	AskVolume5 TThostFtdcVolumeType // 申卖量五
	AveragePrice TThostFtdcPriceType // 当日均价
	ActionDay TThostFtdcDateType // 业务日期
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	BandingUpperPrice TThostFtdcPriceType // 上带价
	BandingLowerPrice TThostFtdcPriceType // 下带价
}

// CThostFtdcDiscountField 会员资金折扣
type CThostFtdcDiscountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Discount TThostFtdcRatioType // 资金折扣比例
}

// CThostFtdcDisseminationField 信息分发
type CThostFtdcDisseminationField struct {
	SequenceSeries TThostFtdcSequenceSeriesType // 序列系列号
	SequenceNo TThostFtdcSequenceNoType // 序列号
}

// CThostFtdcEWarrantOffsetField 仓单折抵信息
type CThostFtdcEWarrantOffsetField struct {
	TradingDay TThostFtdcTradeDateType // 交易日期
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	Direction TThostFtdcDirectionType // 买卖方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Volume TThostFtdcVolumeType // 数量
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcErrExecOrderActionField 错误执行宣告操作
type CThostFtdcErrExecOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType // 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcErrExecOrderField 错误执行宣告
type CThostFtdcErrExecOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	UserID TThostFtdcUserIDType // 用户代码
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionType TThostFtdcActionTypeType // 执行类型
	PosiDirection TThostFtdcPosiDirectionType // 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType // 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType // 期权行权后生成的头寸是否自动平仓
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcErrOrderActionField 错误报单操作
type CThostFtdcErrOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	OrderRef TThostFtdcOrderRefType // 报单引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeChange TThostFtdcVolumeType // 数量变化
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcErrOrderField 错误报单
type CThostFtdcErrOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	UserForceClose TThostFtdcBoolType // 用户强平标志
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	IsSwapOrder TThostFtdcBoolType // 互换单标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcErrorConditionalOrderField 查询错误报单操作
type CThostFtdcErrorConditionalOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 报单提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	OrderSource TThostFtdcOrderSourceType // 报单来源
	OrderStatus TThostFtdcOrderStatusType // 报单状态
	OrderType TThostFtdcOrderTypeType // 报单类型
	VolumeTraded TThostFtdcVolumeType // 今成交数量
	VolumeTotal TThostFtdcVolumeType // 剩余数量
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 委托时间
	ActiveTime TThostFtdcTimeType // 激活时间
	SuspendTime TThostFtdcTimeType // 挂起时间
	UpdateTime TThostFtdcTimeType // 最后修改时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ActiveTraderID TThostFtdcTraderIDType // 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	UserForceClose TThostFtdcBoolType // 用户强平标志
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerOrderSeq TThostFtdcSequenceNoType // 经纪公司报单编号
	RelativeOrderSysID TThostFtdcOrderSysIDType // 相关报单
	ZCETotalTradedVolume TThostFtdcVolumeType // 郑商所成交数量
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	IsSwapOrder TThostFtdcBoolType // 互换单标志
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeBatchOrderActionField 交易所批量报单操作
type CThostFtdcExchangeBatchOrderActionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeCombActionField 交易所申请组合信息
type CThostFtdcExchangeCombActionField struct {
	Direction TThostFtdcDirectionType // 买卖方向
	Volume TThostFtdcVolumeType // 数量
	CombDirection TThostFtdcCombDirectionType // 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionLocalID TThostFtdcOrderLocalIDType // 本地申请组合编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ActionStatus TThostFtdcOrderActionStatusType // 组合状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ComTradeID TThostFtdcTradeIDType // 组合编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeExecOrderActionField 交易所执行宣告操作
type CThostFtdcExchangeExecOrderActionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType // 本地执行宣告编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	ActionType TThostFtdcActionTypeType // 执行类型
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	Volume TThostFtdcVolumeType // 数量
	IPAddress TThostFtdcIPAddressType // IP地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcExchangeExecOrderField 交易所执行宣告信息
type CThostFtdcExchangeExecOrderField struct {
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionType TThostFtdcActionTypeType // 执行类型
	PosiDirection TThostFtdcPosiDirectionType // 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType // 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType // 期权行权后生成的头寸是否自动平仓
	ExecOrderLocalID TThostFtdcOrderLocalIDType // 本地执行宣告编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 执行宣告提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ExecResult TThostFtdcExecResultType // 执行结果
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeField 交易所
type CThostFtdcExchangeField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExchangeName TThostFtdcExchangeNameType // 交易所名称
	ExchangeProperty TThostFtdcExchangePropertyType // 交易所属性
}

// CThostFtdcExchangeForQuoteField 交易所询价信息
type CThostFtdcExchangeForQuoteField struct {
	ForQuoteLocalID TThostFtdcOrderLocalIDType // 本地询价编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	ForQuoteStatus TThostFtdcForQuoteStatusType // 询价状态
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeMarginRateAdjustField 交易所保证金率调整
type CThostFtdcExchangeMarginRateAdjustField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 跟随交易所投资者多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 跟随交易所投资者多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 跟随交易所投资者空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 跟随交易所投资者空头保证金费
	ExchLongMarginRatioByMoney TThostFtdcRatioType // 交易所多头保证金率
	ExchLongMarginRatioByVolume TThostFtdcMoneyType // 交易所多头保证金费
	ExchShortMarginRatioByMoney TThostFtdcRatioType // 交易所空头保证金率
	ExchShortMarginRatioByVolume TThostFtdcMoneyType // 交易所空头保证金费
	NoLongMarginRatioByMoney TThostFtdcRatioType // 不跟随交易所投资者多头保证金率
	NoLongMarginRatioByVolume TThostFtdcMoneyType // 不跟随交易所投资者多头保证金费
	NoShortMarginRatioByMoney TThostFtdcRatioType // 不跟随交易所投资者空头保证金率
	NoShortMarginRatioByVolume TThostFtdcMoneyType // 不跟随交易所投资者空头保证金费
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcExchangeMarginRateField 交易所保证金率
type CThostFtdcExchangeMarginRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcExchangeOptionSelfCloseActionField 交易所期权自对冲操作
type CThostFtdcExchangeOptionSelfCloseActionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType // 本地期权自对冲编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType // 期权行权的头寸是否自对冲
	IPAddress TThostFtdcIPAddressType // IP地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcExchangeOptionSelfCloseField 交易所期权自对冲信息
type CThostFtdcExchangeOptionSelfCloseField struct {
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType // 期权行权的头寸是否自对冲
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType // 本地期权自对冲编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 期权自对冲提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ExecResult TThostFtdcExecResultType // 自对冲结果
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeOrderActionErrorField 交易所报单操作失败
type CThostFtdcExchangeOrderActionErrorField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcExchangeOrderActionField 交易所报单操作
type CThostFtdcExchangeOrderActionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeChange TThostFtdcVolumeType // 数量变化
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeOrderField 交易所报单
type CThostFtdcExchangeOrderField struct {
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 报单提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	OrderSource TThostFtdcOrderSourceType // 报单来源
	OrderStatus TThostFtdcOrderStatusType // 报单状态
	OrderType TThostFtdcOrderTypeType // 报单类型
	VolumeTraded TThostFtdcVolumeType // 今成交数量
	VolumeTotal TThostFtdcVolumeType // 剩余数量
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 委托时间
	ActiveTime TThostFtdcTimeType // 激活时间
	SuspendTime TThostFtdcTimeType // 挂起时间
	UpdateTime TThostFtdcTimeType // 最后修改时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ActiveTraderID TThostFtdcTraderIDType // 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeOrderInsertErrorField 交易所报单插入失败
type CThostFtdcExchangeOrderInsertErrorField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcExchangeQuoteActionField 交易所报价操作
type CThostFtdcExchangeQuoteActionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType // 报价操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	QuoteLocalID TThostFtdcOrderLocalIDType // 本地报价编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExchangeQuoteField 交易所报价信息
type CThostFtdcExchangeQuoteField struct {
	AskPrice TThostFtdcPriceType // 卖价格
	BidPrice TThostFtdcPriceType // 买价格
	AskVolume TThostFtdcVolumeType // 卖数量
	BidVolume TThostFtdcVolumeType // 买数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType // 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType // 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType // 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType // 买投机套保标志
	QuoteLocalID TThostFtdcOrderLocalIDType // 本地报价编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	NotifySequence TThostFtdcSequenceNoType // 报价提示序号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 报价提交状态
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	QuoteSysID TThostFtdcOrderSysIDType // 报价编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	QuoteStatus TThostFtdcOrderStatusType // 报价状态
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	AskOrderSysID TThostFtdcOrderSysIDType // 卖方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType // 买方报单编号
	ForQuoteSysID TThostFtdcOrderSysIDType // 应价编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
}

// CThostFtdcExchangeRateField 汇率
type CThostFtdcExchangeRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	FromCurrencyID TThostFtdcCurrencyIDType // 源币种
	FromCurrencyUnit TThostFtdcCurrencyUnitType // 源币种单位数量
	ToCurrencyID TThostFtdcCurrencyIDType // 目标币种
	ExchangeRate TThostFtdcExchangeRateType // 汇率
}

// CThostFtdcExchangeSequenceField 交易所状态
type CThostFtdcExchangeSequenceField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SequenceNo TThostFtdcSequenceNoType // 序号
	MarketStatus TThostFtdcInstrumentStatusType // 合约交易状态
}

// CThostFtdcExchangeTradeField 交易所成交
type CThostFtdcExchangeTradeField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TradeID TThostFtdcTradeIDType // 成交编号
	Direction TThostFtdcDirectionType // 买卖方向
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	TradingRole TThostFtdcTradingRoleType // 交易角色
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Price TThostFtdcPriceType // 价格
	Volume TThostFtdcVolumeType // 数量
	TradeDate TThostFtdcDateType // 成交时期
	TradeTime TThostFtdcTimeType // 成交时间
	TradeType TThostFtdcTradeTypeType // 成交类型
	PriceSource TThostFtdcPriceSourceType // 成交价来源
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	SequenceNo TThostFtdcSequenceNoType // 序号
	TradeSource TThostFtdcTradeSourceType // 成交来源
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcExecOrderActionField 执行宣告操作
type CThostFtdcExecOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType // 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	ExecOrderLocalID TThostFtdcOrderLocalIDType // 本地执行宣告编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	ActionType TThostFtdcActionTypeType // 执行类型
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExecOrderField 执行宣告
type CThostFtdcExecOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	UserID TThostFtdcUserIDType // 用户代码
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionType TThostFtdcActionTypeType // 执行类型
	PosiDirection TThostFtdcPosiDirectionType // 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType // 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType // 期权行权后生成的头寸是否自动平仓
	ExecOrderLocalID TThostFtdcOrderLocalIDType // 本地执行宣告编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 执行宣告提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ExecResult TThostFtdcExecResultType // 执行结果
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerExecOrderSeq TThostFtdcSequenceNoType // 经纪公司报单编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcExitEmergencyField 退出紧急状态参数
type CThostFtdcExitEmergencyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcFensUserInfoField Fens用户信息
type CThostFtdcFensUserInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	LoginMode TThostFtdcLoginModeType // 登录模式
}

// CThostFtdcForQuoteField 询价
type CThostFtdcForQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ForQuoteRef TThostFtdcOrderRefType // 询价引用
	UserID TThostFtdcUserIDType // 用户代码
	ForQuoteLocalID TThostFtdcOrderLocalIDType // 本地询价编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	ForQuoteStatus TThostFtdcForQuoteStatusType // 询价状态
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerForQutoSeq TThostFtdcSequenceNoType // 经纪公司询价编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcForQuoteParamField 询价价差参数
type CThostFtdcForQuoteParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	LastPrice TThostFtdcPriceType // 最新价
	PriceInterval TThostFtdcPriceType // 价差
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcForQuoteRspField 发给做市商的询价请求
type CThostFtdcForQuoteRspField struct {
	TradingDay TThostFtdcDateType // 交易日
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ForQuoteSysID TThostFtdcOrderSysIDType // 询价编号
	ForQuoteTime TThostFtdcTimeType // 询价时间
	ActionDay TThostFtdcDateType // 业务日期
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcForceUserLogoutField 强制交易员退出
type CThostFtdcForceUserLogoutField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcFrontInfoField 前置信息
type CThostFtdcFrontInfoField struct {
	FrontAddr TThostFtdcAddressType // 前置地址
	QryFreq TThostFtdcQueryFreqType // 查询频率
	FTDPkgFreq TThostFtdcQueryFreqType // FTD频率
}

// CThostFtdcFrontStatusField 前置状态
type CThostFtdcFrontStatusField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
	LastReportDate TThostFtdcDateType // 上次报告日期
	LastReportTime TThostFtdcTimeType // 上次报告时间
	IsActive TThostFtdcBoolType // 是否活跃
}

// CThostFtdcFutureLimitPosiParamField 期货持仓限制参数
type CThostFtdcFutureLimitPosiParamField struct {
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	SpecOpenVolume TThostFtdcVolumeType // 当日投机开仓数量限制
	ArbiOpenVolume TThostFtdcVolumeType // 当日套利开仓数量限制
	OpenVolume TThostFtdcVolumeType // 当日投机+套利开仓数量限制
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcFutureSignIOField 期商签到签退
type CThostFtdcFutureSignIOField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
}

// CThostFtdcIPListField IP列表
type CThostFtdcIPListField struct {
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	IsWhite TThostFtdcBoolType // 是否白名单
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcIndexPriceField 股指现货指数
type CThostFtdcIndexPriceField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ClosePrice TThostFtdcPriceType // 指数现货收盘价
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInputBatchOrderActionField 输入批量报单操作
type CThostFtdcInputBatchOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	UserID TThostFtdcUserIDType // 用户代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputCombActionField 输入的申请组合
type CThostFtdcInputCombActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	CombActionRef TThostFtdcOrderRefType // 组合引用
	UserID TThostFtdcUserIDType // 用户代码
	Direction TThostFtdcDirectionType // 买卖方向
	Volume TThostFtdcVolumeType // 数量
	CombDirection TThostFtdcCombDirectionType // 组合指令方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputExecOrderActionField 输入执行宣告操作
type CThostFtdcInputExecOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExecOrderActionRef TThostFtdcOrderActionRefType // 执行宣告操作引用
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputExecOrderField 输入的执行宣告
type CThostFtdcInputExecOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExecOrderRef TThostFtdcOrderRefType // 执行宣告引用
	UserID TThostFtdcUserIDType // 用户代码
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ActionType TThostFtdcActionTypeType // 执行类型
	PosiDirection TThostFtdcPosiDirectionType // 保留头寸申请的持仓方向
	ReservePositionFlag TThostFtdcExecOrderPositionFlagType // 期权行权后是否保留期货头寸的标记,该字段已废弃
	CloseFlag TThostFtdcExecOrderCloseFlagType // 期权行权后生成的头寸是否自动平仓
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputForQuoteField 输入的询价
type CThostFtdcInputForQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ForQuoteRef TThostFtdcOrderRefType // 询价引用
	UserID TThostFtdcUserIDType // 用户代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputOptionSelfCloseActionField 输入期权自对冲操作
type CThostFtdcInputOptionSelfCloseActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType // 期权自对冲操作引用
	OptionSelfCloseRef TThostFtdcOrderRefType // 期权自对冲引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputOptionSelfCloseField 输入的期权自对冲
type CThostFtdcInputOptionSelfCloseField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OptionSelfCloseRef TThostFtdcOrderRefType // 期权自对冲引用
	UserID TThostFtdcUserIDType // 用户代码
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType // 期权行权的头寸是否自对冲
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcInputOrderActionField 输入报单操作
type CThostFtdcInputOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	OrderRef TThostFtdcOrderRefType // 报单引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeChange TThostFtdcVolumeType // 数量变化
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcInputOrderField 输入报单
type CThostFtdcInputOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	UserForceClose TThostFtdcBoolType // 用户强平标志
	IsSwapOrder TThostFtdcBoolType // 互换单标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcInputQuoteActionField 输入报价操作
type CThostFtdcInputQuoteActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	QuoteActionRef TThostFtdcOrderActionRefType // 报价操作引用
	QuoteRef TThostFtdcOrderRefType // 报价引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType // 报价操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcInputQuoteField 输入的报价
type CThostFtdcInputQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	QuoteRef TThostFtdcOrderRefType // 报价引用
	UserID TThostFtdcUserIDType // 用户代码
	AskPrice TThostFtdcPriceType // 卖价格
	BidPrice TThostFtdcPriceType // 买价格
	AskVolume TThostFtdcVolumeType // 卖数量
	BidVolume TThostFtdcVolumeType // 买数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType // 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType // 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType // 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType // 买投机套保标志
	AskOrderRef TThostFtdcOrderRefType // 衍生卖报单引用
	BidOrderRef TThostFtdcOrderRefType // 衍生买报单引用
	ForQuoteSysID TThostFtdcOrderSysIDType // 应价编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	ClientID TThostFtdcClientIDType // 交易编码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	ReplaceSysID TThostFtdcOrderSysIDType // 被顶单编号
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcInstrumentCommissionRateField 合约手续费率
type CThostFtdcInstrumentCommissionRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BizType TThostFtdcBizTypeType // 业务类型
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInstrumentField 合约
type CThostFtdcInstrumentField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentName TThostFtdcInstrumentNameType // 合约名称
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	reserve3 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ProductClass TThostFtdcProductClassType // 产品类型
	DeliveryYear TThostFtdcYearType // 交割年份
	DeliveryMonth TThostFtdcMonthType // 交割月
	MaxMarketOrderVolume TThostFtdcVolumeType // 市价单最大下单量
	MinMarketOrderVolume TThostFtdcVolumeType // 市价单最小下单量
	MaxLimitOrderVolume TThostFtdcVolumeType // 限价单最大下单量
	MinLimitOrderVolume TThostFtdcVolumeType // 限价单最小下单量
	VolumeMultiple TThostFtdcVolumeMultipleType // 合约数量乘数
	PriceTick TThostFtdcPriceType // 最小变动价位
	CreateDate TThostFtdcDateType // 创建日
	OpenDate TThostFtdcDateType // 上市日
	ExpireDate TThostFtdcDateType // 到期日
	StartDelivDate TThostFtdcDateType // 开始交割日
	EndDelivDate TThostFtdcDateType // 结束交割日
	InstLifePhase TThostFtdcInstLifePhaseType // 合约生命周期状态
	IsTrading TThostFtdcBoolType // 当前是否交易
	PositionType TThostFtdcPositionTypeType // 持仓类型
	PositionDateType TThostFtdcPositionDateTypeType // 持仓日期类型
	LongMarginRatio TThostFtdcRatioType // 多头保证金率
	ShortMarginRatio TThostFtdcRatioType // 空头保证金率
	MaxMarginSideAlgorithm TThostFtdcMaxMarginSideAlgorithmType // 是否使用大额单边保证金算法
	reserve4 TThostFtdcOldInstrumentIDType // 保留的无效字段
	StrikePrice TThostFtdcPriceType // 执行价
	OptionsType TThostFtdcOptionsTypeType // 期权类型
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType // 合约基础商品乘数
	CombinationType TThostFtdcCombinationTypeType // 组合类型
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
	UnderlyingInstrID TThostFtdcInstrumentIDType // 基础商品代码
}

// CThostFtdcInstrumentMarginRateAdjustField 合约保证金率调整
type CThostFtdcInstrumentMarginRateAdjustField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	IsRelative TThostFtdcBoolType // 是否相对交易所收取
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInstrumentMarginRateField 合约保证金率
type CThostFtdcInstrumentMarginRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	IsRelative TThostFtdcBoolType // 是否相对交易所收取
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInstrumentMarginRateULField 合约保证金率调整
type CThostFtdcInstrumentMarginRateULField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInstrumentOrderCommRateField 当前报单手续费的详细内容
type CThostFtdcInstrumentOrderCommRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	OrderCommByVolume TThostFtdcRatioType // 报单手续费
	OrderActionCommByVolume TThostFtdcRatioType // 撤单手续费
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	OrderCommByTrade TThostFtdcRatioType // 报单手续费
	OrderActionCommByTrade TThostFtdcRatioType // 撤单手续费
}

// CThostFtdcInstrumentStatusField 合约状态
type CThostFtdcInstrumentStatusField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	SettlementGroupID TThostFtdcSettlementGroupIDType // 结算组代码
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentStatus TThostFtdcInstrumentStatusType // 合约交易状态
	TradingSegmentSN TThostFtdcTradingSegmentSNType // 交易阶段编号
	EnterTime TThostFtdcTimeType // 进入本状态时间
	EnterReason TThostFtdcInstStatusEnterReasonType // 进入本状态原因
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInstrumentTradingRightField 投资者合约交易权限
type CThostFtdcInstrumentTradingRightField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	TradingRight TThostFtdcTradingRightType // 交易权限
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInvestUnitField 投资单元
type CThostFtdcInvestUnitField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InvestorUnitName TThostFtdcPartyNameType // 投资者单元名称
	InvestorGroupID TThostFtdcInvestorIDType // 投资者分组代码
	CommModelID TThostFtdcInvestorIDType // 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType // 保证金率模板代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcInvestorAccountField 投资者账户
type CThostFtdcInvestorAccountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcInvestorCommodityGroupSPMMMarginField 投资者商品群SPMM记录
type CThostFtdcInvestorCommodityGroupSPMMMarginField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
	MarginBeforeDiscount TThostFtdcMoneyType // 优惠仓位应收保证金
	MarginNoDiscount TThostFtdcMoneyType // 不优惠仓位应收保证金
	LongRisk TThostFtdcMoneyType // 多头风险
	ShortRisk TThostFtdcMoneyType // 空头风险
	CloseFrozenMargin TThostFtdcMoneyType // 商品群平仓冻结保证金
	InterCommodityRate TThostFtdcSPMMDiscountRatioType // SPMM跨品种优惠系数
	MiniMarginRatio TThostFtdcSPMMDiscountRatioType // 商品群最小保证金比例
	AdjustRatio TThostFtdcRatioType // 投资者保证金和交易所保证金的比例
	IntraCommodityDiscount TThostFtdcMoneyType // SPMM品种内优惠汇总
	InterCommodityDiscount TThostFtdcMoneyType // SPMM跨品种优惠
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	InvestorMargin TThostFtdcMoneyType // 投资者保证金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	Commission TThostFtdcMoneyType // 手续费
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	CashIn TThostFtdcMoneyType // 资金差额
	StrikeFrozenMargin TThostFtdcMoneyType // 行权冻结资金
}

// CThostFtdcInvestorCommoditySPMMMarginField 投资者商品组SPMM记录
type CThostFtdcInvestorCommoditySPMMMarginField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
	MarginBeforeDiscount TThostFtdcMoneyType // 优惠仓位应收保证金
	MarginNoDiscount TThostFtdcMoneyType // 不优惠仓位应收保证金
	LongPosRisk TThostFtdcMoneyType // 多头实仓风险
	LongOpenFrozenRisk TThostFtdcMoneyType // 多头开仓冻结风险
	LongCloseFrozenRisk TThostFtdcMoneyType // 多头被平冻结风险
	ShortPosRisk TThostFtdcMoneyType // 空头实仓风险
	ShortOpenFrozenRisk TThostFtdcMoneyType // 空头开仓冻结风险
	ShortCloseFrozenRisk TThostFtdcMoneyType // 空头被平冻结风险
	IntraCommodityRate TThostFtdcSPMMDiscountRatioType // SPMM品种内跨期优惠系数
	OptionDiscountRate TThostFtdcSPMMDiscountRatioType // SPMM期权优惠系数
	PosDiscount TThostFtdcMoneyType // 实仓对冲优惠金额
	OpenFrozenDiscount TThostFtdcMoneyType // 开仓报单对冲优惠金额
	NetRisk TThostFtdcMoneyType // 品种风险净头
	CloseFrozenMargin TThostFtdcMoneyType // 平仓冻结保证金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	Commission TThostFtdcMoneyType // 手续费
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	CashIn TThostFtdcMoneyType // 资金差额
	StrikeFrozenMargin TThostFtdcMoneyType // 行权冻结资金
}

// CThostFtdcInvestorField 投资者
type CThostFtdcInvestorField struct {
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType // 投资者分组代码
	InvestorName TThostFtdcPartyNameType // 投资者名称
	IdentifiedCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	IsActive TThostFtdcBoolType // 是否活跃
	Telephone TThostFtdcTelephoneType // 联系电话
	Address TThostFtdcAddressType // 通讯地址
	OpenDate TThostFtdcDateType // 开户日期
	Mobile TThostFtdcMobileType // 手机
	CommModelID TThostFtdcInvestorIDType // 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType // 保证金率模板代码
	IsOrderFreq TThostFtdcEnumBoolType // 是否频率控制
	IsOpenVolLimit TThostFtdcEnumBoolType // 是否开仓限制
}

// CThostFtdcInvestorGroupField 投资者组
type CThostFtdcInvestorGroupField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType // 投资者分组代码
	InvestorGroupName TThostFtdcInvestorGroupNameType // 投资者分组名称
}

// CThostFtdcInvestorInfoCntSettingField 投资者申报费阶梯收取设置
type CThostFtdcInvestorInfoCntSettingField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ProductID TThostFtdcProductIDType // 商品代码
	IsCalInfoComm TThostFtdcBoolType // 是否收取申报费
	IsLimitInfoMax TThostFtdcBoolType // 是否限制信息量
	InfoMaxLimit TThostFtdcVolumeType // 信息量限制笔数
}

// CThostFtdcInvestorPortfMarginModelField 新组保保证金系数投资者模板对应关系
type CThostFtdcInvestorPortfMarginModelField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	MarginModelID TThostFtdcInvestorIDType // 保证金系数模板
}

// CThostFtdcInvestorPortfMarginRatioField 投资者新型组合保证金系数
type CThostFtdcInvestorPortfMarginRatioField struct {
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	MarginRatio TThostFtdcRatioType // 会员对投资者收取的保证金和交易所对投资者收取的保证金的比例
	ProductGroupID TThostFtdcProductIDType // 产品群代码
}

// CThostFtdcInvestorPortfSettingField 投资者新组保设置
type CThostFtdcInvestorPortfSettingField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者编号
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	UsePortf TThostFtdcBoolType // 是否开启新组保
}

// CThostFtdcInvestorPositionCombineDetailField 投资者组合持仓明细
type CThostFtdcInvestorPositionCombineDetailField struct {
	TradingDay TThostFtdcDateType // 交易日
	OpenDate TThostFtdcDateType // 开仓日期
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SettlementID TThostFtdcSettlementIDType // 结算编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ComTradeID TThostFtdcTradeIDType // 组合编号
	TradeID TThostFtdcTradeIDType // 撮合编号
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Direction TThostFtdcDirectionType // 买卖
	TotalAmt TThostFtdcVolumeType // 持仓量
	Margin TThostFtdcMoneyType // 投资者保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	LegID TThostFtdcLegIDType // 单腿编号
	LegMultiple TThostFtdcLegMultipleType // 单腿乘数
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	TradeGroupID TThostFtdcTradeGroupIDType // 成交组号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	CombInstrumentID TThostFtdcInstrumentIDType // 组合持仓合约编码
}

// CThostFtdcInvestorPositionDetailField 投资者持仓明细
type CThostFtdcInvestorPositionDetailField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Direction TThostFtdcDirectionType // 买卖
	OpenDate TThostFtdcDateType // 开仓日期
	TradeID TThostFtdcTradeIDType // 成交编号
	Volume TThostFtdcVolumeType // 数量
	OpenPrice TThostFtdcPriceType // 开仓价
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	TradeType TThostFtdcTradeTypeType // 成交类型
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CloseProfitByDate TThostFtdcMoneyType // 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType // 逐笔对冲平仓盈亏
	PositionProfitByDate TThostFtdcMoneyType // 逐日盯市持仓盈亏
	PositionProfitByTrade TThostFtdcMoneyType // 逐笔对冲持仓盈亏
	Margin TThostFtdcMoneyType // 投资者保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	LastSettlementPrice TThostFtdcPriceType // 昨结算价
	SettlementPrice TThostFtdcPriceType // 结算价
	CloseVolume TThostFtdcVolumeType // 平仓量
	CloseAmount TThostFtdcMoneyType // 平仓金额
	TimeFirstVolume TThostFtdcVolumeType // 先开先平剩余数量
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	SpecPosiType TThostFtdcSpecPosiTypeType // 特殊持仓标志
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
}

// CThostFtdcInvestorPositionField 投资者持仓
type CThostFtdcInvestorPositionField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PosiDirection TThostFtdcPosiDirectionType // 持仓多空方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	PositionDate TThostFtdcPositionDateType // 持仓日期
	YdPosition TThostFtdcVolumeType // 上日持仓
	Position TThostFtdcVolumeType // 今日持仓
	LongFrozen TThostFtdcVolumeType // 多头冻结
	ShortFrozen TThostFtdcVolumeType // 空头冻结
	LongFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	OpenVolume TThostFtdcVolumeType // 开仓量
	CloseVolume TThostFtdcVolumeType // 平仓量
	OpenAmount TThostFtdcMoneyType // 开仓金额
	CloseAmount TThostFtdcMoneyType // 平仓金额
	PositionCost TThostFtdcMoneyType // 持仓成本
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	UseMargin TThostFtdcMoneyType // 占用的保证金
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	SettlementPrice TThostFtdcPriceType // 本次结算价
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OpenCost TThostFtdcMoneyType // 开仓成本
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	CombPosition TThostFtdcVolumeType // 组合成交形成的持仓
	CombLongFrozen TThostFtdcVolumeType // 组合多头冻结
	CombShortFrozen TThostFtdcVolumeType // 组合空头冻结
	CloseProfitByDate TThostFtdcMoneyType // 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType // 逐笔对冲平仓盈亏
	TodayPosition TThostFtdcVolumeType // 今日持仓
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	StrikeFrozen TThostFtdcVolumeType // 执行冻结
	StrikeFrozenAmount TThostFtdcMoneyType // 执行冻结金额
	AbandonFrozen TThostFtdcVolumeType // 放弃执行冻结
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	YdStrikeFrozen TThostFtdcVolumeType // 执行冻结的昨仓
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	PositionCostOffset TThostFtdcMoneyType // 持仓成本差值
	TasPosition TThostFtdcVolumeType // tas持仓手数
	TasPositionCost TThostFtdcMoneyType // tas持仓成本
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcInvestorProdRCAMSMarginField 投资者品种RCAMS保证金
type CThostFtdcInvestorProdRCAMSMarginField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	ProductGroupID TThostFtdcProductIDType // 商品群代码
	RiskBeforeDiscount TThostFtdcMoneyType // 品种组合前风险
	IntraInstrRisk TThostFtdcMoneyType // 同合约对冲风险
	BPosRisk TThostFtdcMoneyType // 品种买持仓风险
	SPosRisk TThostFtdcMoneyType // 品种卖持仓风险
	IntraProdRisk TThostFtdcMoneyType // 品种内对冲风险
	NetRisk TThostFtdcMoneyType // 品种净持仓风险
	InterProdRisk TThostFtdcMoneyType // 品种间对冲风险
	ShortOptRiskAdj TThostFtdcMoneyType // 空头期权风险调整
	OptionRoyalty TThostFtdcMoneyType // 空头期权权利金
	MMSACloseFrozenMargin TThostFtdcMoneyType // 大边组合平仓冻结保证金
	CloseCombFrozenMargin TThostFtdcMoneyType // 策略组合平仓/行权冻结保证金
	CloseFrozenMargin TThostFtdcMoneyType // 平仓/行权冻结保证金
	MMSAOpenFrozenMargin TThostFtdcMoneyType // 大边组合开仓冻结保证金
	DeliveryOpenFrozenMargin TThostFtdcMoneyType // 交割月期货开仓冻结保证金
	OpenFrozenMargin TThostFtdcMoneyType // 开仓冻结保证金
	UseFrozenMargin TThostFtdcMoneyType // 投资者冻结保证金
	MMSAExchMargin TThostFtdcMoneyType // 大边组合交易所持仓保证金
	DeliveryExchMargin TThostFtdcMoneyType // 交割月期货交易所持仓保证金
	CombExchMargin TThostFtdcMoneyType // 策略组合交易所保证金
	ExchMargin TThostFtdcMoneyType // 交易所持仓保证金
	UseMargin TThostFtdcMoneyType // 投资者持仓保证金
}

// CThostFtdcInvestorProdRULEMarginField 投资者产品RULE保证金
type CThostFtdcInvestorProdRULEMarginField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	InstrumentClass TThostFtdcInstrumentClassType // 合约类型
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
	BStdPosition TThostFtdcStdPositionType // 买标准持仓
	SStdPosition TThostFtdcStdPositionType // 卖标准持仓
	BStdOpenFrozen TThostFtdcStdPositionType // 买标准开仓冻结
	SStdOpenFrozen TThostFtdcStdPositionType // 卖标准开仓冻结
	BStdCloseFrozen TThostFtdcStdPositionType // 买标准平仓冻结
	SStdCloseFrozen TThostFtdcStdPositionType // 卖标准平仓冻结
	IntraProdStdPosition TThostFtdcStdPositionType // 品种内对冲标准持仓
	NetStdPosition TThostFtdcStdPositionType // 品种内单腿标准持仓
	InterProdStdPosition TThostFtdcStdPositionType // 品种间对冲标准持仓
	SingleStdPosition TThostFtdcStdPositionType // 单腿标准持仓
	IntraProdMargin TThostFtdcMoneyType // 品种内对锁保证金
	InterProdMargin TThostFtdcMoneyType // 品种间对锁保证金
	SingleMargin TThostFtdcMoneyType // 跨品种单腿保证金
	NonCombMargin TThostFtdcMoneyType // 非组合合约保证金
	AddOnMargin TThostFtdcMoneyType // 附加保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	AddOnFrozenMargin TThostFtdcMoneyType // 附加冻结保证金
	OpenFrozenMargin TThostFtdcMoneyType // 开仓冻结保证金
	CloseFrozenMargin TThostFtdcMoneyType // 平仓冻结保证金
	Margin TThostFtdcMoneyType // 品种保证金
	FrozenMargin TThostFtdcMoneyType // 冻结保证金
}

// CThostFtdcInvestorProdSPBMDetailField 投资者产品SPBM明细
type CThostFtdcInvestorProdSPBMDetailField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	IntraInstrMargin TThostFtdcMoneyType // 合约内对锁保证金
	BCollectingMargin TThostFtdcMoneyType // 买归集保证金
	SCollectingMargin TThostFtdcMoneyType // 卖归集保证金
	IntraProdMargin TThostFtdcMoneyType // 品种内合约间对锁保证金
	NetMargin TThostFtdcMoneyType // 净保证金
	InterProdMargin TThostFtdcMoneyType // 产品间对锁保证金
	SingleMargin TThostFtdcMoneyType // 裸保证金
	AddOnMargin TThostFtdcMoneyType // 附加保证金
	DeliveryMargin TThostFtdcMoneyType // 交割月保证金
	CallOptionMinRisk TThostFtdcMoneyType // 看涨期权最低风险
	PutOptionMinRisk TThostFtdcMoneyType // 看跌期权最低风险
	OptionMinRisk TThostFtdcMoneyType // 卖方期权最低风险
	OptionValueOffset TThostFtdcMoneyType // 买方期权冲抵价值
	OptionRoyalty TThostFtdcMoneyType // 卖方期权权利金
	RealOptionValueOffset TThostFtdcMoneyType // 价值冲抵
	Margin TThostFtdcMoneyType // 保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
}

// CThostFtdcInvestorProductGroupMarginField 投资者品种/跨品种保证金
type CThostFtdcInvestorProductGroupMarginField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	LongFrozenMargin TThostFtdcMoneyType // 多头冻结的保证金
	ShortFrozenMargin TThostFtdcMoneyType // 空头冻结的保证金
	UseMargin TThostFtdcMoneyType // 占用的保证金
	LongUseMargin TThostFtdcMoneyType // 多头保证金
	ShortUseMargin TThostFtdcMoneyType // 空头保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	LongExchMargin TThostFtdcMoneyType // 交易所多头保证金
	ShortExchMargin TThostFtdcMoneyType // 交易所空头保证金
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	Commission TThostFtdcMoneyType // 手续费
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	CashIn TThostFtdcMoneyType // 资金差额
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	OffsetAmount TThostFtdcMoneyType // 折抵总金额
	LongOffsetAmount TThostFtdcMoneyType // 多头折抵总金额
	ShortOffsetAmount TThostFtdcMoneyType // 空头折抵总金额
	ExchOffsetAmount TThostFtdcMoneyType // 交易所折抵总金额
	LongExchOffsetAmount TThostFtdcMoneyType // 交易所多头折抵总金额
	ShortExchOffsetAmount TThostFtdcMoneyType // 交易所空头折抵总金额
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	ProductGroupID TThostFtdcInstrumentIDType // 品种/跨品种标示
}

// CThostFtdcInvestorTradingRightField 投资者交易权限设置
type CThostFtdcInvestorTradingRightField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InvstTradingRight TThostFtdcInvstTradingRightType // 交易权限
}

// CThostFtdcInvestorWithdrawAlgorithmField 经纪公司可提资金算法表
type CThostFtdcInvestorWithdrawAlgorithmField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	UsingRatio TThostFtdcRatioType // 可提资金比例
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	FundMortgageRatio TThostFtdcRatioType // 货币质押比率
}

// CThostFtdcLinkManField 联系人
type CThostFtdcLinkManField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PersonType TThostFtdcPersonTypeType // 联系人类型
	IdentifiedCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	PersonName TThostFtdcPartyNameType // 名称
	Telephone TThostFtdcTelephoneType // 联系电话
	Address TThostFtdcAddressType // 通讯地址
	ZipCode TThostFtdcZipCodeType // 邮政编码
	Priority TThostFtdcPriorityType // 优先级
	UOAZipCode TThostFtdcUOAZipCodeType // 开户邮政编码
	PersonFullName TThostFtdcInvestorFullNameType // 全称
}

// CThostFtdcLoadSettlementInfoField 装载结算信息
type CThostFtdcLoadSettlementInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcLoginForbiddenIPField 禁止登录IP
type CThostFtdcLoginForbiddenIPField struct {
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcLoginForbiddenUserField 禁止登录用户
type CThostFtdcLoginForbiddenUserField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcLoginInfoField 登录信息
type CThostFtdcLoginInfoField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	LoginDate TThostFtdcDateType // 登录日期
	LoginTime TThostFtdcTimeType // 登录时间
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	SystemName TThostFtdcSystemNameType // 系统名称
	PasswordDeprecated TThostFtdcPasswordType // 密码,已弃用
	MaxOrderRef TThostFtdcOrderRefType // 最大报单引用
	SHFETime TThostFtdcTimeType // 上期所时间
	DCETime TThostFtdcTimeType // 大商所时间
	CZCETime TThostFtdcTimeType // 郑商所时间
	FFEXTime TThostFtdcTimeType // 中金所时间
	MacAddress TThostFtdcMacAddressType // Mac地址
	OneTimePassword TThostFtdcPasswordType // 动态密码
	INETime TThostFtdcTimeType // 能源中心时间
	IsQryControl TThostFtdcBoolType // 查询时是否需要流控
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	Password TThostFtdcPasswordType // 密码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcLogoutAllField 登录信息
type CThostFtdcLogoutAllField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	SystemName TThostFtdcSystemNameType // 系统名称
}

// CThostFtdcMDTraderOfferField 交易所行情报盘机
type CThostFtdcMDTraderOfferField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	Password TThostFtdcPasswordType // 密码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	TraderConnectStatus TThostFtdcTraderConnectStatusType // 交易所交易员连接状态
	ConnectRequestDate TThostFtdcDateType // 发出连接请求的日期
	ConnectRequestTime TThostFtdcTimeType // 发出连接请求的时间
	LastReportDate TThostFtdcDateType // 上次报告日期
	LastReportTime TThostFtdcTimeType // 上次报告时间
	ConnectDate TThostFtdcDateType // 完成连接日期
	ConnectTime TThostFtdcTimeType // 完成连接时间
	StartDate TThostFtdcDateType // 启动日期
	StartTime TThostFtdcTimeType // 启动时间
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	MaxTradeID TThostFtdcTradeIDType // 本席位最大成交编号
	MaxOrderMessageReference TThostFtdcReturnCodeType // 本席位最大报单备拷
	OrderCancelAlg TThostFtdcOrderCancelAlgType // 撤单时选择席位算法
}

// CThostFtdcMMInstrumentCommissionRateField 做市商合约手续费率
type CThostFtdcMMInstrumentCommissionRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcMMOptionInstrCommRateField 当前做市商期权合约手续费的详细内容
type CThostFtdcMMOptionInstrCommRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	StrikeRatioByMoney TThostFtdcRatioType // 执行手续费率
	StrikeRatioByVolume TThostFtdcRatioType // 执行手续费
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcManualSyncBrokerUserOTPField 手工同步用户动态令牌
type CThostFtdcManualSyncBrokerUserOTPField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	OTPType TThostFtdcOTPTypeType // 动态令牌类型
	FirstOTP TThostFtdcPasswordType // 第一个动态密码
	SecondOTP TThostFtdcPasswordType // 第二个动态密码
}

// CThostFtdcMarginModelField 投资者保证金率模板
type CThostFtdcMarginModelField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	MarginModelID TThostFtdcInvestorIDType // 保证金率模板代码
	MarginModelName TThostFtdcCommModelNameType // 模板名称
}

// CThostFtdcMarketDataAsk23Field 行情申卖二、三属性
type CThostFtdcMarketDataAsk23Field struct {
	AskPrice2 TThostFtdcPriceType // 申卖价二
	AskVolume2 TThostFtdcVolumeType // 申卖量二
	AskPrice3 TThostFtdcPriceType // 申卖价三
	AskVolume3 TThostFtdcVolumeType // 申卖量三
}

// CThostFtdcMarketDataAsk45Field 行情申卖四、五属性
type CThostFtdcMarketDataAsk45Field struct {
	AskPrice4 TThostFtdcPriceType // 申卖价四
	AskVolume4 TThostFtdcVolumeType // 申卖量四
	AskPrice5 TThostFtdcPriceType // 申卖价五
	AskVolume5 TThostFtdcVolumeType // 申卖量五
}

// CThostFtdcMarketDataAveragePriceField 成交均价
type CThostFtdcMarketDataAveragePriceField struct {
	AveragePrice TThostFtdcPriceType // 当日均价
}

// CThostFtdcMarketDataBandingPriceField 行情上下带价
type CThostFtdcMarketDataBandingPriceField struct {
	BandingUpperPrice TThostFtdcPriceType // 上带价
	BandingLowerPrice TThostFtdcPriceType // 下带价
}

// CThostFtdcMarketDataBaseField 行情基础属性
type CThostFtdcMarketDataBaseField struct {
	TradingDay TThostFtdcDateType // 交易日
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	PreClosePrice TThostFtdcPriceType // 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType // 昨持仓量
	PreDelta TThostFtdcRatioType // 昨虚实度
}

// CThostFtdcMarketDataBestPriceField 行情最优价属性
type CThostFtdcMarketDataBestPriceField struct {
	BidPrice1 TThostFtdcPriceType // 申买价一
	BidVolume1 TThostFtdcVolumeType // 申买量一
	AskPrice1 TThostFtdcPriceType // 申卖价一
	AskVolume1 TThostFtdcVolumeType // 申卖量一
}

// CThostFtdcMarketDataBid23Field 行情申买二、三属性
type CThostFtdcMarketDataBid23Field struct {
	BidPrice2 TThostFtdcPriceType // 申买价二
	BidVolume2 TThostFtdcVolumeType // 申买量二
	BidPrice3 TThostFtdcPriceType // 申买价三
	BidVolume3 TThostFtdcVolumeType // 申买量三
}

// CThostFtdcMarketDataBid45Field 行情申买四、五属性
type CThostFtdcMarketDataBid45Field struct {
	BidPrice4 TThostFtdcPriceType // 申买价四
	BidVolume4 TThostFtdcVolumeType // 申买量四
	BidPrice5 TThostFtdcPriceType // 申买价五
	BidVolume5 TThostFtdcVolumeType // 申买量五
}

// CThostFtdcMarketDataExchangeField 行情交易所代码属性
type CThostFtdcMarketDataExchangeField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcMarketDataField 市场行情
type CThostFtdcMarketDataField struct {
	TradingDay TThostFtdcDateType // 交易日
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	LastPrice TThostFtdcPriceType // 最新价
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	PreClosePrice TThostFtdcPriceType // 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType // 昨持仓量
	OpenPrice TThostFtdcPriceType // 今开盘
	HighestPrice TThostFtdcPriceType // 最高价
	LowestPrice TThostFtdcPriceType // 最低价
	Volume TThostFtdcVolumeType // 数量
	Turnover TThostFtdcMoneyType // 成交金额
	OpenInterest TThostFtdcLargeVolumeType // 持仓量
	ClosePrice TThostFtdcPriceType // 今收盘
	SettlementPrice TThostFtdcPriceType // 本次结算价
	UpperLimitPrice TThostFtdcPriceType // 涨停板价
	LowerLimitPrice TThostFtdcPriceType // 跌停板价
	PreDelta TThostFtdcRatioType // 昨虚实度
	CurrDelta TThostFtdcRatioType // 今虚实度
	UpdateTime TThostFtdcTimeType // 最后修改时间
	UpdateMillisec TThostFtdcMillisecType // 最后修改毫秒
	ActionDay TThostFtdcDateType // 业务日期
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcMarketDataLastMatchField 行情最新成交属性
type CThostFtdcMarketDataLastMatchField struct {
	LastPrice TThostFtdcPriceType // 最新价
	Volume TThostFtdcVolumeType // 数量
	Turnover TThostFtdcMoneyType // 成交金额
	OpenInterest TThostFtdcLargeVolumeType // 持仓量
}

// CThostFtdcMarketDataStaticField 行情静态属性
type CThostFtdcMarketDataStaticField struct {
	OpenPrice TThostFtdcPriceType // 今开盘
	HighestPrice TThostFtdcPriceType // 最高价
	LowestPrice TThostFtdcPriceType // 最低价
	ClosePrice TThostFtdcPriceType // 今收盘
	UpperLimitPrice TThostFtdcPriceType // 涨停板价
	LowerLimitPrice TThostFtdcPriceType // 跌停板价
	SettlementPrice TThostFtdcPriceType // 本次结算价
	CurrDelta TThostFtdcRatioType // 今虚实度
}

// CThostFtdcMarketDataUpdateTimeField 行情更新时间属性
type CThostFtdcMarketDataUpdateTimeField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	UpdateTime TThostFtdcTimeType // 最后修改时间
	UpdateMillisec TThostFtdcMillisecType // 最后修改毫秒
	ActionDay TThostFtdcDateType // 业务日期
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcMortgageParamField 质押配比参数
type CThostFtdcMortgageParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	MortgageBalance TThostFtdcRatioType // 质押配比系数
	CheckMortgageRatio TThostFtdcBoolType // 开仓是否验证质押配比
}

// CThostFtdcMulticastInstrumentField MulticastInstrument
type CThostFtdcMulticastInstrumentField struct {
	TopicID TThostFtdcInstallIDType // 主题号
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentNo TThostFtdcInstallIDType // 合约编号
	CodePrice TThostFtdcPriceType // 基准价
	VolumeMultiple TThostFtdcVolumeMultipleType // 合约数量乘数
	PriceTick TThostFtdcPriceType // 最小变动价位
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcNoticeField 客户通知
type CThostFtdcNoticeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	Content TThostFtdcContentType // 消息正文
	SequenceLabel TThostFtdcSequenceLabelType // 经纪公司通知内容序列号
}

// CThostFtdcNotifyFutureSignInField 期商签到通知
type CThostFtdcNotifyFutureSignInField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	PinKey TThostFtdcPasswordKeyType // PIN密钥
	MacKey TThostFtdcPasswordKeyType // MAC密钥
}

// CThostFtdcNotifyFutureSignOutField 期商签退通知
type CThostFtdcNotifyFutureSignOutField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcNotifyQueryAccountField 查询账户信息通知
type CThostFtdcNotifyQueryAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	BankUseAmount TThostFtdcTradeAmountType // 银行可用金额
	BankFetchAmount TThostFtdcTradeAmountType // 银行可取金额
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcNotifySyncKeyField 交易核心向银期报盘发出密钥同步处理结果的通知
type CThostFtdcNotifySyncKeyField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Message TThostFtdcAddInfoType // 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcOpenAccountField 银期开户信息
type CThostFtdcOpenAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType // 汇钞标志
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	TID TThostFtdcTIDType // 交易ID
	UserID TThostFtdcUserIDType // 用户标识
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcOptionInstrCommRateField 当前期权合约手续费的详细内容
type CThostFtdcOptionInstrCommRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	StrikeRatioByMoney TThostFtdcRatioType // 执行手续费率
	StrikeRatioByVolume TThostFtdcRatioType // 执行手续费
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionInstrDeltaField 期权合约delta值
type CThostFtdcOptionInstrDeltaField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Delta TThostFtdcRatioType // Delta值
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionInstrMarginAdjustField 当前期权合约保证金调整系数
type CThostFtdcOptionInstrMarginAdjustField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SShortMarginRatioByMoney TThostFtdcRatioType // 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType // 投机空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType // 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType // 保值空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType // 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType // 套利空头保证金调整系数
	IsRelative TThostFtdcBoolType // 是否跟随交易所收取
	MShortMarginRatioByMoney TThostFtdcRatioType // 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType // 做市商空头保证金调整系数
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionInstrMiniMarginField 当前期权合约最小保证金
type CThostFtdcOptionInstrMiniMarginField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	MinMargin TThostFtdcMoneyType // 单位（手）期权合约最小保证金
	ValueMethod TThostFtdcValueMethodType // 取值方式
	IsRelative TThostFtdcBoolType // 是否跟随交易所收取
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionInstrTradeCostField 期权交易成本
type CThostFtdcOptionInstrTradeCostField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	FixedMargin TThostFtdcMoneyType // 期权合约保证金不变部分
	MiniMargin TThostFtdcMoneyType // 期权合约最小保证金
	Royalty TThostFtdcMoneyType // 期权合约权利金
	ExchFixedMargin TThostFtdcMoneyType // 交易所期权合约保证金不变部分
	ExchMiniMargin TThostFtdcMoneyType // 交易所期权合约最小保证金
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionInstrTradingRightField 投资者期权合约交易权限
type CThostFtdcOptionInstrTradingRightField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Direction TThostFtdcDirectionType // 买卖方向
	TradingRight TThostFtdcTradingRightType // 交易权限
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcOptionSelfCloseActionField 期权自对冲操作
type CThostFtdcOptionSelfCloseActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OptionSelfCloseActionRef TThostFtdcOrderActionRefType // 期权自对冲操作引用
	OptionSelfCloseRef TThostFtdcOrderRefType // 期权自对冲引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType // 本地期权自对冲编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcOptionSelfCloseField 期权自对冲
type CThostFtdcOptionSelfCloseField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OptionSelfCloseRef TThostFtdcOrderRefType // 期权自对冲引用
	UserID TThostFtdcUserIDType // 用户代码
	Volume TThostFtdcVolumeType // 数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	OptSelfCloseFlag TThostFtdcOptSelfCloseFlagType // 期权行权的头寸是否自对冲
	OptionSelfCloseLocalID TThostFtdcOrderLocalIDType // 本地期权自对冲编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 期权自对冲提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ExecResult TThostFtdcExecResultType // 自对冲结果
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerOptionSelfCloseSeq TThostFtdcSequenceNoType // 经纪公司报单编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcOrderActionField 报单操作
type CThostFtdcOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	OrderRef TThostFtdcOrderRefType // 报单引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeChange TThostFtdcVolumeType // 数量变化
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcOrderField 报单
type CThostFtdcOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 报单提交状态
	NotifySequence TThostFtdcSequenceNoType // 报单提示序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	OrderSource TThostFtdcOrderSourceType // 报单来源
	OrderStatus TThostFtdcOrderStatusType // 报单状态
	OrderType TThostFtdcOrderTypeType // 报单类型
	VolumeTraded TThostFtdcVolumeType // 今成交数量
	VolumeTotal TThostFtdcVolumeType // 剩余数量
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 委托时间
	ActiveTime TThostFtdcTimeType // 激活时间
	SuspendTime TThostFtdcTimeType // 挂起时间
	UpdateTime TThostFtdcTimeType // 最后修改时间
	CancelTime TThostFtdcTimeType // 撤销时间
	ActiveTraderID TThostFtdcTraderIDType // 最后修改交易所交易员代码
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	UserForceClose TThostFtdcBoolType // 用户强平标志
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerOrderSeq TThostFtdcSequenceNoType // 经纪公司报单编号
	RelativeOrderSysID TThostFtdcOrderSysIDType // 相关报单
	ZCETotalTradedVolume TThostFtdcVolumeType // 郑商所成交数量
	IsSwapOrder TThostFtdcBoolType // 互换单标志
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcParkedOrderActionField 输入预埋单操作
type CThostFtdcParkedOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OrderActionRef TThostFtdcOrderActionRefType // 报单操作引用
	OrderRef TThostFtdcOrderRefType // 报单引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeChange TThostFtdcVolumeType // 数量变化
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType // 预埋撤单单编号
	UserType TThostFtdcUserTypeType // 用户类型
	Status TThostFtdcParkedOrderStatusType // 预埋撤单状态
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcParkedOrderField 预埋单
type CThostFtdcParkedOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	OrderPriceType TThostFtdcOrderPriceTypeType // 报单价格条件
	Direction TThostFtdcDirectionType // 买卖方向
	CombOffsetFlag TThostFtdcCombOffsetFlagType // 组合开平标志
	CombHedgeFlag TThostFtdcCombHedgeFlagType // 组合投机套保标志
	LimitPrice TThostFtdcPriceType // 价格
	VolumeTotalOriginal TThostFtdcVolumeType // 数量
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	GTDDate TThostFtdcDateType // GTD日期
	VolumeCondition TThostFtdcVolumeConditionType // 成交量类型
	MinVolume TThostFtdcVolumeType // 最小成交量
	ContingentCondition TThostFtdcContingentConditionType // 触发条件
	StopPrice TThostFtdcPriceType // 止损价
	ForceCloseReason TThostFtdcForceCloseReasonType // 强平原因
	IsAutoSuspend TThostFtdcBoolType // 自动挂起标志
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	RequestID TThostFtdcRequestIDType // 请求编号
	UserForceClose TThostFtdcBoolType // 用户强平标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParkedOrderID TThostFtdcParkedOrderIDType // 预埋报单编号
	UserType TThostFtdcUserTypeType // 用户类型
	Status TThostFtdcParkedOrderStatusType // 预埋单状态
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	IsSwapOrder TThostFtdcBoolType // 互换单标志
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	ClientID TThostFtdcClientIDType // 交易编码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcPartBrokerField 会员编码和经纪公司编码对照表
type CThostFtdcPartBrokerField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	IsActive TThostFtdcBoolType // 是否活跃
}

// CThostFtdcPortfTradeParamSettingField 组保交易参数设置
type CThostFtdcPortfTradeParamSettingField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Portfolio TThostFtdcPortfolioType // 新型组保算法
	IsActionVerify TThostFtdcBoolType // 撤单是否验资
	IsCloseVerify TThostFtdcBoolType // 平仓是否验资
}

// CThostFtdcPositionProfitAlgorithmField 浮动盈亏算法
type CThostFtdcPositionProfitAlgorithmField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Algorithm TThostFtdcAlgorithmType // 盈亏算法
	Memo TThostFtdcMemoType // 备注
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcProductExchRateField 产品报价汇率
type CThostFtdcProductExchRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	QuoteCurrencyID TThostFtdcCurrencyIDType // 报价币种类型
	ExchangeRate TThostFtdcExchangeRateType // 汇率
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcProductField 产品
type CThostFtdcProductField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ProductName TThostFtdcProductNameType // 产品名称
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductClass TThostFtdcProductClassType // 产品类型
	VolumeMultiple TThostFtdcVolumeMultipleType // 合约数量乘数
	PriceTick TThostFtdcPriceType // 最小变动价位
	MaxMarketOrderVolume TThostFtdcVolumeType // 市价单最大下单量
	MinMarketOrderVolume TThostFtdcVolumeType // 市价单最小下单量
	MaxLimitOrderVolume TThostFtdcVolumeType // 限价单最大下单量
	MinLimitOrderVolume TThostFtdcVolumeType // 限价单最小下单量
	PositionType TThostFtdcPositionTypeType // 持仓类型
	PositionDateType TThostFtdcPositionDateTypeType // 持仓日期类型
	CloseDealType TThostFtdcCloseDealTypeType // 平仓处理类型
	TradeCurrencyID TThostFtdcCurrencyIDType // 交易币种类型
	MortgageFundUseRange TThostFtdcMortgageFundUseRangeType // 质押资金可用范围
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	UnderlyingMultiple TThostFtdcUnderlyingMultipleType // 合约基础商品乘数
	ProductID TThostFtdcInstrumentIDType // 产品代码
	ExchangeProductID TThostFtdcInstrumentIDType // 交易所产品代码
	OpenLimitControlLevel TThostFtdcOpenLimitControlLevelType // 开仓量限制粒度
	OrderFreqControlLevel TThostFtdcOrderFreqControlLevelType // 报单频率控制粒度
}

// CThostFtdcProductGroupField 投资者品种/跨品种保证金产品组
type CThostFtdcProductGroupField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ProductID TThostFtdcInstrumentIDType // 产品代码
	ProductGroupID TThostFtdcInstrumentIDType // 产品组代码
}

// CThostFtdcQryAccountregisterField 请求查询银期签约关系
type CThostFtdcQryAccountregisterField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	BankID TThostFtdcBankIDType // 银行编码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构编码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcQryAuthForbiddenIPField 查询禁止认证IP
type CThostFtdcQryAuthForbiddenIPField struct {
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcQryBatchOrderActionField 查询批量报单操作
type CThostFtdcQryBatchOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryBrokerField 查询经纪公司
type CThostFtdcQryBrokerField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcQryBrokerTradingAlgosField 查询经纪公司交易算法
type CThostFtdcQryBrokerTradingAlgosField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryBrokerTradingParamsField 查询经纪公司交易参数
type CThostFtdcQryBrokerTradingParamsField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
}

// CThostFtdcQryBrokerUserEventField 查询经纪公司用户事件
type CThostFtdcQryBrokerUserEventField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserEventType TThostFtdcUserEventTypeType // 用户事件类型
}

// CThostFtdcQryBrokerUserField 查询经纪公司用户
type CThostFtdcQryBrokerUserField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQryBrokerUserFunctionField 查询经纪公司用户权限
type CThostFtdcQryBrokerUserFunctionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQryBulletinField 查询交易所公告
type CThostFtdcQryBulletinField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BulletinID TThostFtdcBulletinIDType // 公告编号
	SequenceNo TThostFtdcSequenceNoType // 序列号
	NewsType TThostFtdcNewsTypeType // 公告类型
	NewsUrgency TThostFtdcNewsUrgencyType // 紧急程度
}

// CThostFtdcQryCFMMCBrokerKeyField 查询保证金监管系统经纪公司密钥
type CThostFtdcQryCFMMCBrokerKeyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcQryCFMMCTradingAccountKeyField 请求查询保证金监管系统经纪公司资金账户密钥
type CThostFtdcQryCFMMCTradingAccountKeyField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryClassifiedInstrumentField 查询分类合约
type CThostFtdcQryClassifiedInstrumentField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
	TradingType TThostFtdcTradingTypeType // 合约交易状态
	ClassType TThostFtdcClassTypeType // 合约分类类型
}

// CThostFtdcQryCombActionField 申请组合查询
type CThostFtdcQryCombActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryCombInstrumentGuardField 组合合约安全系数查询
type CThostFtdcQryCombInstrumentGuardField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryCombPromotionParamField 查询组合优惠比例
type CThostFtdcQryCombPromotionParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryCombinationLegField 查询组合合约分腿
type CThostFtdcQryCombinationLegField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	LegID TThostFtdcLegIDType // 单腿编号
	reserve2 TThostFtdcOldInstrumentIDType // 保留的无效字段
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
	LegInstrumentID TThostFtdcInstrumentIDType // 单腿合约代码
}

// CThostFtdcQryCommRateModelField 请求查询投资者手续费率模板
type CThostFtdcQryCommRateModelField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	CommModelID TThostFtdcInvestorIDType // 手续费率模板代码
}

// CThostFtdcQryContractBankField 查询签约银行请求
type CThostFtdcQryContractBankField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	BankID TThostFtdcBankIDType // 银行代码
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码
}

// CThostFtdcQryCurrDRIdentityField 查询当前交易中心
type CThostFtdcQryCurrDRIdentityField struct {
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcQryDepthMarketDataField 查询行情
type CThostFtdcQryDepthMarketDataField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProductClass TThostFtdcProductClassType // 产品类型
}

// CThostFtdcQryEWarrantOffsetField 查询仓单折抵信息
type CThostFtdcQryEWarrantOffsetField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryErrExecOrderActionField 查询错误执行宣告操作
type CThostFtdcQryErrExecOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryErrExecOrderField 查询错误执行宣告
type CThostFtdcQryErrExecOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryErrOrderActionField 查询错误报单操作
type CThostFtdcQryErrOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryErrOrderField 查询错误报单
type CThostFtdcQryErrOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryExchangeCombActionField 交易所申请组合查询
type CThostFtdcQryExchangeCombActionField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryExchangeExecOrderActionField 交易所执行宣告操作查询
type CThostFtdcQryExchangeExecOrderActionField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryExchangeExecOrderField 交易所执行宣告查询
type CThostFtdcQryExchangeExecOrderField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryExchangeField 查询交易所
type CThostFtdcQryExchangeField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryExchangeForQuoteField 交易所询价查询
type CThostFtdcQryExchangeForQuoteField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryExchangeMarginRateAdjustField 查询交易所调整保证金率
type CThostFtdcQryExchangeMarginRateAdjustField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryExchangeMarginRateField 查询交易所保证金率
type CThostFtdcQryExchangeMarginRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryExchangeOrderActionField 查询交易所报单操作
type CThostFtdcQryExchangeOrderActionField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryExchangeOrderField 查询交易所报单
type CThostFtdcQryExchangeOrderField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryExchangeQuoteActionField 交易所报价操作查询
type CThostFtdcQryExchangeQuoteActionField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryExchangeQuoteField 交易所报价查询
type CThostFtdcQryExchangeQuoteField struct {
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryExchangeRateField 查询汇率
type CThostFtdcQryExchangeRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	FromCurrencyID TThostFtdcCurrencyIDType // 源币种
	ToCurrencyID TThostFtdcCurrencyIDType // 目标币种
}

// CThostFtdcQryExchangeSequenceField 查询交易所状态
type CThostFtdcQryExchangeSequenceField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryExecOrderActionField 执行宣告操作查询
type CThostFtdcQryExecOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryExecOrderField 执行宣告查询
type CThostFtdcQryExecOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExecOrderSysID TThostFtdcExecOrderSysIDType // 执行宣告编号
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryForQuoteField 询价查询
type CThostFtdcQryForQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryForQuoteParamField 查询询价价差参数
type CThostFtdcQryForQuoteParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryFrontStatusField 查询前置状态
type CThostFtdcQryFrontStatusField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
}

// CThostFtdcQryHisOrderField 查询报单
type CThostFtdcQryHisOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryIPListField 查询IP列表
type CThostFtdcQryIPListField struct {
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcQryInstrumentCommissionRateField 查询手续费率
type CThostFtdcQryInstrumentCommissionRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInstrumentField 查询合约
type CThostFtdcQryInstrumentField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	reserve3 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcQryInstrumentMarginRateField 查询合约保证金率
type CThostFtdcQryInstrumentMarginRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInstrumentOrderCommRateField 报单手续费率查询
type CThostFtdcQryInstrumentOrderCommRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInstrumentStatusField 查询合约状态
type CThostFtdcQryInstrumentStatusField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	reserve1 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcQryInstrumentTradingRightField 查询合约交易权限
type CThostFtdcQryInstrumentTradingRightField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInvestUnitField 查询投资单元
type CThostFtdcQryInvestUnitField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcQryInvestorCommodityGroupSPMMMarginField 投资者商品群SPMM记录查询
type CThostFtdcQryInvestorCommodityGroupSPMMMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
}

// CThostFtdcQryInvestorCommoditySPMMMarginField 投资者商品组SPMM记录查询
type CThostFtdcQryInvestorCommoditySPMMMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
}

// CThostFtdcQryInvestorField 查询投资者
type CThostFtdcQryInvestorField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryInvestorGroupField 查询投资者组
type CThostFtdcQryInvestorGroupField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcQryInvestorPortfMarginRatioField 投资者新型组合保证金系数查询
type CThostFtdcQryInvestorPortfMarginRatioField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductGroupID TThostFtdcProductIDType // 产品群代码
}

// CThostFtdcQryInvestorPortfSettingField 投资者新组保设置查询
type CThostFtdcQryInvestorPortfSettingField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者编号
}

// CThostFtdcQryInvestorPositionCombineDetailField 查询组合持仓明细
type CThostFtdcQryInvestorPositionCombineDetailField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	CombInstrumentID TThostFtdcInstrumentIDType // 组合持仓合约编码
}

// CThostFtdcQryInvestorPositionDetailField 查询投资者持仓明细
type CThostFtdcQryInvestorPositionDetailField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInvestorPositionField 查询投资者持仓
type CThostFtdcQryInvestorPositionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryInvestorProdRCAMSMarginField 投资者品种RCAMS保证金查询
type CThostFtdcQryInvestorProdRCAMSMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
}

// CThostFtdcQryInvestorProdRULEMarginField 投资者产品RULE保证金查询
type CThostFtdcQryInvestorProdRULEMarginField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
}

// CThostFtdcQryInvestorProdSPBMDetailField 投资者产品SPBM明细查询
type CThostFtdcQryInvestorProdSPBMDetailField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQryInvestorProductGroupMarginField 查询投资者品种/跨品种保证金
type CThostFtdcQryInvestorProductGroupMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	ProductGroupID TThostFtdcInstrumentIDType // 品种/跨品种标示
}

// CThostFtdcQryLinkManField 查询联系人
type CThostFtdcQryLinkManField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQryLoginForbiddenIPField 查询禁止登录IP
type CThostFtdcQryLoginForbiddenIPField struct {
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcQryLoginForbiddenUserField 查询禁止登录用户
type CThostFtdcQryLoginForbiddenUserField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQryMDTraderOfferField 查询行情报盘机
type CThostFtdcQryMDTraderOfferField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryMMInstrumentCommissionRateField 查询做市商合约手续费率
type CThostFtdcQryMMInstrumentCommissionRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryMMOptionInstrCommRateField 做市商期权手续费率查询
type CThostFtdcQryMMOptionInstrCommRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryMarginModelField 请求查询投资者保证金率模板
type CThostFtdcQryMarginModelField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	MarginModelID TThostFtdcInvestorIDType // 保证金率模板代码
}

// CThostFtdcQryMaxOrderVolumeField 查询最大报单数量
type CThostFtdcQryMaxOrderVolumeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	Direction TThostFtdcDirectionType // 买卖方向
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	MaxVolume TThostFtdcVolumeType // 最大允许报单数量
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryMaxOrderVolumeWithPriceField 根据价格查询最大报单数量
type CThostFtdcQryMaxOrderVolumeWithPriceField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	Direction TThostFtdcDirectionType // 买卖方向
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	MaxVolume TThostFtdcVolumeType // 最大允许报单数量
	Price TThostFtdcPriceType // 报单价格
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryMulticastInstrumentField QryMulticastInstrument
type CThostFtdcQryMulticastInstrumentField struct {
	TopicID TThostFtdcInstallIDType // 主题号
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryNoticeField 查询客户通知
type CThostFtdcQryNoticeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
}

// CThostFtdcQryOptionInstrCommRateField 期权手续费率查询
type CThostFtdcQryOptionInstrCommRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryOptionInstrTradeCostField 期权交易成本查询
type CThostFtdcQryOptionInstrTradeCostField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	InputPrice TThostFtdcPriceType // 期权合约报价
	UnderlyingPrice TThostFtdcPriceType // 标的价格,填0则用昨结算价
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryOptionInstrTradingRightField 查询期权合约交易权限
type CThostFtdcQryOptionInstrTradingRightField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	Direction TThostFtdcDirectionType // 买卖方向
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryOptionSelfCloseActionField 期权自对冲操作查询
type CThostFtdcQryOptionSelfCloseActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryOptionSelfCloseField 期权自对冲查询
type CThostFtdcQryOptionSelfCloseField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OptionSelfCloseSysID TThostFtdcOrderSysIDType // 期权自对冲编号
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryOrderActionField 查询报单操作
type CThostFtdcQryOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryOrderField 查询报单
type CThostFtdcQryOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryParkedOrderActionField 查询预埋撤单
type CThostFtdcQryParkedOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryParkedOrderField 查询预埋单
type CThostFtdcQryParkedOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryPartBrokerField 查询经纪公司会员代码
type CThostFtdcQryPartBrokerField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
}

// CThostFtdcQryProductExchRateField 产品报价汇率查询
type CThostFtdcQryProductExchRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcQryProductField 查询产品
type CThostFtdcQryProductField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ProductClass TThostFtdcProductClassType // 产品类型
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcQryProductGroupField 查询产品组
type CThostFtdcQryProductGroupField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcQryQuoteActionField 报价操作查询
type CThostFtdcQryQuoteActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQryQuoteField 报价查询
type CThostFtdcQryQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType // 报价编号
	InsertTimeStart TThostFtdcTimeType // 开始时间
	InsertTimeEnd TThostFtdcTimeType // 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryRCAMSCombProductInfoField RCAMS产品组合信息查询
type CThostFtdcQryRCAMSCombProductInfoField struct {
	ProductID TThostFtdcProductIDType // 产品代码
	CombProductID TThostFtdcProductIDType // 商品组代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
}

// CThostFtdcQryRCAMSInstrParameterField RCAMS同合约风险对冲参数查询
type CThostFtdcQryRCAMSInstrParameterField struct {
	ProductID TThostFtdcProductIDType // 产品代码
}

// CThostFtdcQryRCAMSInterParameterField RCAMS跨品种风险折抵参数查询
type CThostFtdcQryRCAMSInterParameterField struct {
	ProductGroupID TThostFtdcProductIDType // 商品群代码
	CombProduct1 TThostFtdcProductIDType // 产品组合代码1
	CombProduct2 TThostFtdcProductIDType // 产品组合代码2
}

// CThostFtdcQryRCAMSIntraParameterField RCAMS品种内风险对冲参数查询
type CThostFtdcQryRCAMSIntraParameterField struct {
	CombProductID TThostFtdcProductIDType // 产品组合代码
}

// CThostFtdcQryRCAMSInvestorCombPositionField RCAMS策略组合持仓查询
type CThostFtdcQryRCAMSInvestorCombPositionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
}

// CThostFtdcQryRCAMSShortOptAdjustParamField RCAMS空头期权风险调整参数查询
type CThostFtdcQryRCAMSShortOptAdjustParamField struct {
	CombProductID TThostFtdcProductIDType // 产品组合代码
}

// CThostFtdcQryRULEInstrParameterField RULE合约保证金参数查询
type CThostFtdcQryRULEInstrParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryRULEInterParameterField RULE跨品种抵扣参数查询
type CThostFtdcQryRULEInterParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
}

// CThostFtdcQryRULEIntraParameterField RULE品种内对锁仓折扣参数查询
type CThostFtdcQryRULEIntraParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQryRiskSettleInvstPositionField 投资者风险结算持仓查询
type CThostFtdcQryRiskSettleInvstPositionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryRiskSettleProductStatusField 风险结算产品查询
type CThostFtdcQryRiskSettleProductStatusField struct {
	ProductID TThostFtdcInstrumentIDType // 产品代码
}

// CThostFtdcQrySPBMAddOnInterParameterField SPBM附加跨品种抵扣参数查询
type CThostFtdcQrySPBMAddOnInterParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
}

// CThostFtdcQrySPBMFutureParameterField SPBM期货合约保证金参数查询
type CThostFtdcQrySPBMFutureParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQrySPBMInterParameterField SPBM跨品种抵扣参数查询
type CThostFtdcQrySPBMInterParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
}

// CThostFtdcQrySPBMIntraParameterField SPBM品种内对锁仓折扣参数查询
type CThostFtdcQrySPBMIntraParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQrySPBMInvestorPortfDefField 投资者套餐选择查询
type CThostFtdcQrySPBMInvestorPortfDefField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQrySPBMOptionParameterField SPBM期权合约保证金参数查询
type CThostFtdcQrySPBMOptionParameterField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQrySPBMPortfDefinitionField 组合保证金套餐查询
type CThostFtdcQrySPBMPortfDefinitionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	PortfolioDefID TThostFtdcPortfolioDefIDType // 组合保证金套餐代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
}

// CThostFtdcQrySPMMInstParamField SPMM合约参数查询
type CThostFtdcQrySPMMInstParamField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQrySPMMProductParamField SPMM产品参数查询
type CThostFtdcQrySPMMProductParamField struct {
	ProductID TThostFtdcSPMMProductIDType // 产品代码
}

// CThostFtdcQrySecAgentACIDMapField 二级代理操作员银期权限查询
type CThostFtdcQrySecAgentACIDMapField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	AccountID TThostFtdcAccountIDType // 资金账户
	CurrencyID TThostFtdcCurrencyIDType // 币种
}

// CThostFtdcQrySecAgentCheckModeField 查询二级代理商资金校验模式
type CThostFtdcQrySecAgentCheckModeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
}

// CThostFtdcQrySecAgentTradeInfoField 查询二级代理商信息
type CThostFtdcQrySecAgentTradeInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	BrokerSecAgentID TThostFtdcAccountIDType // 境外中介机构资金帐号
}

// CThostFtdcQrySettlementInfoConfirmField 查询结算信息确认域
type CThostFtdcQrySettlementInfoConfirmField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcQrySettlementInfoField 查询投资者结算结果
type CThostFtdcQrySettlementInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	TradingDay TThostFtdcDateType // 交易日
	AccountID TThostFtdcAccountIDType // 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcQryStrikeOffsetField 期权执行偏移值查询
type CThostFtdcQryStrikeOffsetField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQrySuperUserField 查询管理用户
type CThostFtdcQrySuperUserField struct {
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQrySuperUserFunctionField 查询管理用户功能权限
type CThostFtdcQrySuperUserFunctionField struct {
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQrySyncDelaySwapField 查询延时换汇同步
type CThostFtdcQrySyncDelaySwapField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	DelaySwapSeqNo TThostFtdcDepositSeqNoType // 延时换汇流水号
}

// CThostFtdcQrySyncDepositField 查询出入金流水
type CThostFtdcQrySyncDepositField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	DepositSeqNo TThostFtdcDepositSeqNoType // 出入金流水号
}

// CThostFtdcQrySyncFundMortgageField 查询货币质押流水
type CThostFtdcQrySyncFundMortgageField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	MortgageSeqNo TThostFtdcDepositSeqNoType // 货币质押流水号
}

// CThostFtdcQrySyncStatusField 查询组合合约分腿
type CThostFtdcQrySyncStatusField struct {
	TradingDay TThostFtdcDateType // 交易日
}

// CThostFtdcQryThostUserFunctionField Thost终端用户功能权限查询
type CThostFtdcQryThostUserFunctionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQryTradeField 查询成交
type CThostFtdcQryTradeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TradeID TThostFtdcTradeIDType // 成交编号
	TradeTimeStart TThostFtdcTimeType // 开始时间
	TradeTimeEnd TThostFtdcTimeType // 结束时间
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcQryTraderAssignField 席位与交易中心对应关系维护查询
type CThostFtdcQryTraderAssignField struct {
	TraderID TThostFtdcTraderIDType // 交易员代码
}

// CThostFtdcQryTraderField 查询交易员
type CThostFtdcQryTraderField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryTraderOfferField 查询交易员报盘机
type CThostFtdcQryTraderOfferField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
}

// CThostFtdcQryTradingAccountField 查询资金账户
type CThostFtdcQryTradingAccountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	BizType TThostFtdcBizTypeType // 业务类型
	AccountID TThostFtdcAccountIDType // 投资者帐号
}

// CThostFtdcQryTradingCodeField 查询交易编码
type CThostFtdcQryTradingCodeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ClientID TThostFtdcClientIDType // 客户代码
	ClientIDType TThostFtdcClientIDTypeType // 交易编码类型
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcQryTradingNoticeField 查询交易事件通知
type CThostFtdcQryTradingNoticeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcQryTransferBankField 查询转帐银行
type CThostFtdcQryTransferBankField struct {
	BankID TThostFtdcBankIDType // 银行代码
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码
}

// CThostFtdcQryTransferSerialField 请求查询转帐流水
type CThostFtdcQryTransferSerialField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	BankID TThostFtdcBankIDType // 银行编码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcQryUserRightsAssignField 查询用户下单权限分配表
type CThostFtdcQryUserRightsAssignField struct {
	BrokerID TThostFtdcBrokerIDType // 应用单元代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQryUserSessionField 查询用户会话
type CThostFtdcQryUserSessionField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcQueryBrokerDepositField 查询经纪公司资金
type CThostFtdcQueryBrokerDepositField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
}

// CThostFtdcQueryCFMMCTradingAccountTokenField 查询监控中心用户令牌
type CThostFtdcQueryCFMMCTradingAccountTokenField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcQueryFreqField 查询频率，每秒查询比数
type CThostFtdcQueryFreqField struct {
	QueryFreq TThostFtdcQueryFreqType // 查询频率
	FTDPkgFreq TThostFtdcQueryFreqType // FTD频率
}

// CThostFtdcQuoteActionField 报价操作
type CThostFtdcQuoteActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	QuoteActionRef TThostFtdcOrderActionRefType // 报价操作引用
	QuoteRef TThostFtdcOrderRefType // 报价引用
	RequestID TThostFtdcRequestIDType // 请求编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	QuoteSysID TThostFtdcOrderSysIDType // 报价操作编号
	ActionFlag TThostFtdcActionFlagType // 操作标志
	ActionDate TThostFtdcDateType // 操作日期
	ActionTime TThostFtdcTimeType // 操作时间
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	QuoteLocalID TThostFtdcOrderLocalIDType // 本地报价编号
	ActionLocalID TThostFtdcOrderLocalIDType // 操作本地编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	OrderActionStatus TThostFtdcOrderActionStatusType // 报单操作状态
	UserID TThostFtdcUserIDType // 用户代码
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	IPAddress TThostFtdcIPAddressType // IP地址
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcQuoteField 报价
type CThostFtdcQuoteField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	QuoteRef TThostFtdcOrderRefType // 报价引用
	UserID TThostFtdcUserIDType // 用户代码
	AskPrice TThostFtdcPriceType // 卖价格
	BidPrice TThostFtdcPriceType // 买价格
	AskVolume TThostFtdcVolumeType // 卖数量
	BidVolume TThostFtdcVolumeType // 买数量
	RequestID TThostFtdcRequestIDType // 请求编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	AskOffsetFlag TThostFtdcOffsetFlagType // 卖开平标志
	BidOffsetFlag TThostFtdcOffsetFlagType // 买开平标志
	AskHedgeFlag TThostFtdcHedgeFlagType // 卖投机套保标志
	BidHedgeFlag TThostFtdcHedgeFlagType // 买投机套保标志
	QuoteLocalID TThostFtdcOrderLocalIDType // 本地报价编号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	InstallID TThostFtdcInstallIDType // 安装编号
	NotifySequence TThostFtdcSequenceNoType // 报价提示序号
	OrderSubmitStatus TThostFtdcOrderSubmitStatusType // 报价提交状态
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	QuoteSysID TThostFtdcOrderSysIDType // 报价编号
	InsertDate TThostFtdcDateType // 报单日期
	InsertTime TThostFtdcTimeType // 插入时间
	CancelTime TThostFtdcTimeType // 撤销时间
	QuoteStatus TThostFtdcOrderStatusType // 报价状态
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	SequenceNo TThostFtdcSequenceNoType // 序号
	AskOrderSysID TThostFtdcOrderSysIDType // 卖方报单编号
	BidOrderSysID TThostFtdcOrderSysIDType // 买方报单编号
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	StatusMsg TThostFtdcErrorMsgType // 状态信息
	ActiveUserID TThostFtdcUserIDType // 操作用户代码
	BrokerQuoteSeq TThostFtdcSequenceNoType // 经纪公司报价编号
	AskOrderRef TThostFtdcOrderRefType // 衍生卖报单引用
	BidOrderRef TThostFtdcOrderRefType // 衍生买报单引用
	ForQuoteSysID TThostFtdcOrderSysIDType // 应价编号
	BranchID TThostFtdcBranchIDType // 营业部编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	AccountID TThostFtdcAccountIDType // 资金账号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	reserve3 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	IPAddress TThostFtdcIPAddressType // IP地址
	ReplaceSysID TThostFtdcOrderSysIDType // 被顶单编号
	TimeCondition TThostFtdcTimeConditionType // 有效期类型
	OrderMemo TThostFtdcOrderMemoType // 报单回显字段
	SessionReqSeq TThostFtdcSequenceNo12Type // session上请求计数 api自动维护
}

// CThostFtdcRCAMSCombProductInfoField RCAMS产品组合信息
type CThostFtdcRCAMSCombProductInfoField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcProductIDType // 产品代码
	CombProductID TThostFtdcProductIDType // 商品组代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
}

// CThostFtdcRCAMSInstrParameterField RCAMS同合约风险对冲参数
type CThostFtdcRCAMSInstrParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcProductIDType // 产品代码
	HedgeRate TThostFtdcHedgeRateType // 同合约风险对冲比率
}

// CThostFtdcRCAMSInterParameterField RCAMS跨品种风险折抵参数
type CThostFtdcRCAMSInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
	Priority TThostFtdcRCAMSPriorityType // 优先级
	CreditRate TThostFtdcHedgeRateType // 折抵率
	CombProduct1 TThostFtdcProductIDType // 产品组合代码1
	CombProduct2 TThostFtdcProductIDType // 产品组合代码2
}

// CThostFtdcRCAMSIntraParameterField RCAMS品种内风险对冲参数
type CThostFtdcRCAMSIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	HedgeRate TThostFtdcHedgeRateType // 品种内对冲比率
}

// CThostFtdcRCAMSInvestorCombPositionField RCAMS策略组合持仓
type CThostFtdcRCAMSInvestorCombPositionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	PosiDirection TThostFtdcPosiDirectionType // 持仓多空方向
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
	LegID TThostFtdcLegIDType // 单腿编号
	ExchangeInstID TThostFtdcExchangeInstIDType // 交易所组合合约代码
	TotalAmt TThostFtdcVolumeType // 持仓量
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	Margin TThostFtdcMoneyType // 投资者保证金
}

// CThostFtdcRCAMSShortOptAdjustParamField RCAMS空头期权风险调整参数
type CThostFtdcRCAMSShortOptAdjustParamField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	AdjustValue TThostFtdcAdjustValueType // 空头期权风险调整标准
}

// CThostFtdcRULEInstrParameterField RULE合约保证金参数
type CThostFtdcRULEInstrParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InstrumentClass TThostFtdcInstrumentClassType // 合约类型
	StdInstrumentID TThostFtdcInstrumentIDType // 标准合约
	BSpecRatio TThostFtdcRatioType // 投机买折算系数
	SSpecRatio TThostFtdcRatioType // 投机卖折算系数
	BHedgeRatio TThostFtdcRatioType // 套保买折算系数
	SHedgeRatio TThostFtdcRatioType // 套保卖折算系数
	BAddOnMargin TThostFtdcMoneyType // 买附加风险保证金
	SAddOnMargin TThostFtdcMoneyType // 卖附加风险保证金
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
}

// CThostFtdcRULEInterParameterField RULE跨品种抵扣参数
type CThostFtdcRULEInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	InterRate TThostFtdcRatioType // 品种间对锁仓费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
	Leg1PropFactor TThostFtdcCommonIntType // 腿1比例系数
	Leg2PropFactor TThostFtdcCommonIntType // 腿2比例系数
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
	CommodityGroupName TThostFtdcInstrumentNameType // 商品群名称
}

// CThostFtdcRULEIntraParameterField RULE品种内对锁仓折扣参数
type CThostFtdcRULEIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	StdInstrumentID TThostFtdcInstrumentIDType // 标准合约
	StdInstrMargin TThostFtdcMoneyType // 标准合约保证金
	UsualIntraRate TThostFtdcRatioType // 一般月份合约组合保证金系数
	DeliveryIntraRate TThostFtdcRatioType // 临近交割合约组合保证金系数
}

// CThostFtdcRemoveParkedOrderActionField 删除预埋撤单
type CThostFtdcRemoveParkedOrderActionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ParkedOrderActionID TThostFtdcParkedOrderActionIDType // 预埋撤单编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcRemoveParkedOrderField 删除预埋单
type CThostFtdcRemoveParkedOrderField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ParkedOrderID TThostFtdcParkedOrderIDType // 预埋报单编号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcReqApiHandshakeField api握手请求
type CThostFtdcReqApiHandshakeField struct {
	CryptoKeyVersion TThostFtdcCryptoKeyVersionType // api与front通信密钥版本号
}

// CThostFtdcReqAuthenticateField 客户端认证请求
type CThostFtdcReqAuthenticateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	AuthCode TThostFtdcAuthCodeType // 认证码
	AppID TThostFtdcAppIDType // App代码
}

// CThostFtdcReqCancelAccountField 转帐销户请求
type CThostFtdcReqCancelAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType // 汇钞标志
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	TID TThostFtdcTIDType // 交易ID
	UserID TThostFtdcUserIDType // 用户标识
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqChangeAccountField 变更银行账户请求
type CThostFtdcReqChangeAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	NewBankAccount TThostFtdcBankAccountType // 新银行帐号
	NewBankPassWord TThostFtdcPasswordType // 新银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	TID TThostFtdcTIDType // 交易ID
	Digest TThostFtdcDigestType // 摘要
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqDayEndFileReadyField 日终文件就绪请求
type CThostFtdcReqDayEndFileReadyField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	FileBusinessCode TThostFtdcFileBusinessCodeType // 文件业务功能
	Digest TThostFtdcDigestType // 摘要
}

// CThostFtdcReqFutureSignOutField 期商签退请求
type CThostFtdcReqFutureSignOutField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
}

// CThostFtdcReqGenUserCaptchaField 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserCaptchaField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcReqGenUserTextField 用户发出获取安全安全登陆方法请求
type CThostFtdcReqGenUserTextField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcReqOpenAccountField 转帐开户请求
type CThostFtdcReqOpenAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	CashExchangeCode TThostFtdcCashExchangeCodeType // 汇钞标志
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	TID TThostFtdcTIDType // 交易ID
	UserID TThostFtdcUserIDType // 用户标识
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqQueryAccountField 查询账户信息请求
type CThostFtdcReqQueryAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqQueryTradeResultBySerialField 查询指定流水号的交易结果请求
type CThostFtdcReqQueryTradeResultBySerialField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	Reference TThostFtdcSerialType // 流水号
	RefrenceIssureType TThostFtdcInstitutionTypeType // 本流水号发布者的机构类型
	RefrenceIssure TThostFtdcOrganCodeType // 本流水号发布者机构编码
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	Digest TThostFtdcDigestType // 摘要
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqRepealField 冲正请求
type CThostFtdcReqRepealField struct {
	RepealTimeInterval TThostFtdcRepealTimeIntervalType // 冲正时间间隔
	RepealedTimes TThostFtdcRepealedTimesType // 已经冲正次数
	BankRepealFlag TThostFtdcBankRepealFlagType // 银行冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType // 期商冲正标志
	PlateRepealSerial TThostFtdcPlateSerialType // 被冲正平台流水号
	BankRepealSerial TThostFtdcBankSerialType // 被冲正银行流水号
	FutureRepealSerial TThostFtdcFutureSerialType // 被冲正期货流水号
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType // 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType // 费用支付标志
	CustFee TThostFtdcCustFeeType // 应收客户费用
	BrokerFee TThostFtdcFutureFeeType // 应收期货公司费用
	Message TThostFtdcAddInfoType // 发送方给接收方的消息
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	TransferStatus TThostFtdcTransferStatusType // 转账交易状态
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqSyncKeyField 交易核心向银期报盘发出密钥同步请求
type CThostFtdcReqSyncKeyField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Message TThostFtdcAddInfoType // 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
}

// CThostFtdcReqTransferField 转账请求
type CThostFtdcReqTransferField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType // 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType // 费用支付标志
	CustFee TThostFtdcCustFeeType // 应收客户费用
	BrokerFee TThostFtdcFutureFeeType // 应收期货公司费用
	Message TThostFtdcAddInfoType // 发送方给接收方的消息
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	TransferStatus TThostFtdcTransferStatusType // 转账交易状态
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcReqUserAuthMethodField 用户发出获取安全安全登陆方法请求
type CThostFtdcReqUserAuthMethodField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcReqUserLoginField 用户登录请求
type CThostFtdcReqUserLoginField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	OneTimePassword TThostFtdcPasswordType // 动态密码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
}

// CThostFtdcReqUserLoginSMField 国密用户登录请求
type CThostFtdcReqUserLoginSMField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	OneTimePassword TThostFtdcPasswordType // 动态密码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
	BrokerName TThostFtdcBrokerNameType // 经纪公司名称
	AuthCode TThostFtdcAuthCodeType // 认证码
	AppID TThostFtdcAppIDType // App代码
	PIN TThostFtdcPasswordType // PIN码
}

// CThostFtdcReqUserLoginWithCaptchaField 用户发出带图形验证码的登录请求请求
type CThostFtdcReqUserLoginWithCaptchaField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	Captcha TThostFtdcPasswordType // 图形验证码的文字内容
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
}

// CThostFtdcReqUserLoginWithOTPField 用户发出带动态验证码的登录请求请求
type CThostFtdcReqUserLoginWithOTPField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	OTPPassword TThostFtdcPasswordType // OTP密码
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
}

// CThostFtdcReqUserLoginWithTextField 用户发出带短信验证码的登录请求请求
type CThostFtdcReqUserLoginWithTextField struct {
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	Password TThostFtdcPasswordType // 密码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	Text TThostFtdcPasswordType // 短信验证码文字内容
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientIPAddress TThostFtdcIPAddressType // 终端IP地址
}

// CThostFtdcReqVerifyApiKeyField api给front的验证key的请求
type CThostFtdcReqVerifyApiKeyField struct {
	ApiHandshakeDataLen TThostFtdcHandshakeDataLenType // 握手回复数据长度
	ApiHandshakeData TThostFtdcHandshakeDataType // 握手回复数据
}

// CThostFtdcReserveOpenAccountConfirmField 银期预约开户确认请求
type CThostFtdcReserveOpenAccountConfirmField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcLongIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	TID TThostFtdcTIDType // 交易ID
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	BankReserveOpenSeq TThostFtdcBankSerialType // 预约开户银行流水号
	BookDate TThostFtdcTradeDateType // 预约开户日期
	BookPsw TThostFtdcPasswordType // 预约开户验证密码
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcReserveOpenAccountField 银期预约开户
type CThostFtdcReserveOpenAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcLongIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	Gender TThostFtdcGenderType // 性别
	CountryCode TThostFtdcCountryCodeType // 国家代码
	CustType TThostFtdcCustTypeType // 客户类型
	Address TThostFtdcAddressType // 地址
	ZipCode TThostFtdcZipCodeType // 邮编
	Telephone TThostFtdcTelephoneType // 电话号码
	MobilePhone TThostFtdcMobilePhoneType // 手机
	Fax TThostFtdcFaxType // 传真
	EMail TThostFtdcEMailType // 电子邮件
	MoneyAccountStatus TThostFtdcMoneyAccountStatusType // 资金账户状态
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	InstallID TThostFtdcInstallIDType // 安装编号
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	TID TThostFtdcTIDType // 交易ID
	ReserveOpenAccStas TThostFtdcReserveOpenAccStasType // 预约开户状态
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcReturnResultField 返回结果
type CThostFtdcReturnResultField struct {
	ReturnCode TThostFtdcReturnCodeType // 返回代码
	DescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType // 返回码描述
}

// CThostFtdcRiskSettleInvstPositionField 投资者风险结算持仓
type CThostFtdcRiskSettleInvstPositionField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PosiDirection TThostFtdcPosiDirectionType // 持仓多空方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	PositionDate TThostFtdcPositionDateType // 持仓日期
	YdPosition TThostFtdcVolumeType // 上日持仓
	Position TThostFtdcVolumeType // 今日持仓
	LongFrozen TThostFtdcVolumeType // 多头冻结
	ShortFrozen TThostFtdcVolumeType // 空头冻结
	LongFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	OpenVolume TThostFtdcVolumeType // 开仓量
	CloseVolume TThostFtdcVolumeType // 平仓量
	OpenAmount TThostFtdcMoneyType // 开仓金额
	CloseAmount TThostFtdcMoneyType // 平仓金额
	PositionCost TThostFtdcMoneyType // 持仓成本
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	UseMargin TThostFtdcMoneyType // 占用的保证金
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	SettlementPrice TThostFtdcPriceType // 本次结算价
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OpenCost TThostFtdcMoneyType // 开仓成本
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	CombPosition TThostFtdcVolumeType // 组合成交形成的持仓
	CombLongFrozen TThostFtdcVolumeType // 组合多头冻结
	CombShortFrozen TThostFtdcVolumeType // 组合空头冻结
	CloseProfitByDate TThostFtdcMoneyType // 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType // 逐笔对冲平仓盈亏
	TodayPosition TThostFtdcVolumeType // 今日持仓
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	StrikeFrozen TThostFtdcVolumeType // 执行冻结
	StrikeFrozenAmount TThostFtdcMoneyType // 执行冻结金额
	AbandonFrozen TThostFtdcVolumeType // 放弃执行冻结
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	YdStrikeFrozen TThostFtdcVolumeType // 执行冻结的昨仓
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	PositionCostOffset TThostFtdcMoneyType // 持仓成本差值
	TasPosition TThostFtdcVolumeType // tas持仓手数
	TasPositionCost TThostFtdcMoneyType // tas持仓成本
}

// CThostFtdcRiskSettleProductStatusField 风险品种
type CThostFtdcRiskSettleProductStatusField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品编号
	ProductStatus TThostFtdcProductStatusType // 产品结算状态
}

// CThostFtdcRspApiHandshakeField front发给api的握手回复
type CThostFtdcRspApiHandshakeField struct {
	FrontHandshakeDataLen TThostFtdcHandshakeDataLenType // 握手回复数据长度
	FrontHandshakeData TThostFtdcHandshakeDataType // 握手回复数据
	IsApiAuthEnabled TThostFtdcBoolType // API认证是否开启
}

// CThostFtdcRspAuthenticateField 客户端认证响应
type CThostFtdcRspAuthenticateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	AppID TThostFtdcAppIDType // App代码
	AppType TThostFtdcAppTypeType // App类型
}

// CThostFtdcRspFutureSignInField 期商签到响应
type CThostFtdcRspFutureSignInField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	PinKey TThostFtdcPasswordKeyType // PIN密钥
	MacKey TThostFtdcPasswordKeyType // MAC密钥
}

// CThostFtdcRspFutureSignOutField 期商签退响应
type CThostFtdcRspFutureSignOutField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Digest TThostFtdcDigestType // 摘要
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcRspGenUserCaptchaField 生成的图片验证码信息
type CThostFtdcRspGenUserCaptchaField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	CaptchaInfoLen TThostFtdcCaptchaInfoLenType // 图片信息长度
	CaptchaInfo TThostFtdcCaptchaInfoType // 图片信息
}

// CThostFtdcRspGenUserTextField 短信验证码生成的回复
type CThostFtdcRspGenUserTextField struct {
	UserTextSeq TThostFtdcUserTextSeqType // 短信验证码序号
}

// CThostFtdcRspInfoField 响应信息
type CThostFtdcRspInfoField struct {
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcRspQueryAccountField 查询账户信息响应
type CThostFtdcRspQueryAccountField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	BankUseAmount TThostFtdcTradeAmountType // 银行可用金额
	BankFetchAmount TThostFtdcTradeAmountType // 银行可取金额
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcRspQueryTradeResultBySerialField 查询指定流水号的交易结果响应
type CThostFtdcRspQueryTradeResultBySerialField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	Reference TThostFtdcSerialType // 流水号
	RefrenceIssureType TThostFtdcInstitutionTypeType // 本流水号发布者的机构类型
	RefrenceIssure TThostFtdcOrganCodeType // 本流水号发布者机构编码
	OriginReturnCode TThostFtdcReturnCodeType // 原始返回代码
	OriginDescrInfoForReturnCode TThostFtdcDescrInfoForReturnCodeType // 原始返回码描述
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	Digest TThostFtdcDigestType // 摘要
}

// CThostFtdcRspRepealField 冲正响应
type CThostFtdcRspRepealField struct {
	RepealTimeInterval TThostFtdcRepealTimeIntervalType // 冲正时间间隔
	RepealedTimes TThostFtdcRepealedTimesType // 已经冲正次数
	BankRepealFlag TThostFtdcBankRepealFlagType // 银行冲正标志
	BrokerRepealFlag TThostFtdcBrokerRepealFlagType // 期商冲正标志
	PlateRepealSerial TThostFtdcPlateSerialType // 被冲正平台流水号
	BankRepealSerial TThostFtdcBankSerialType // 被冲正银行流水号
	FutureRepealSerial TThostFtdcFutureSerialType // 被冲正期货流水号
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType // 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType // 费用支付标志
	CustFee TThostFtdcCustFeeType // 应收客户费用
	BrokerFee TThostFtdcFutureFeeType // 应收期货公司费用
	Message TThostFtdcAddInfoType // 发送方给接收方的消息
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	TransferStatus TThostFtdcTransferStatusType // 转账交易状态
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcRspSyncKeyField 交易核心向银期报盘发出密钥同步响应
type CThostFtdcRspSyncKeyField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	InstallID TThostFtdcInstallIDType // 安装编号
	UserID TThostFtdcUserIDType // 用户标识
	Message TThostFtdcAddInfoType // 交易核心给银期报盘的消息
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcRspTransferField 银行发起银行资金转期货响应
type CThostFtdcRspTransferField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	InstallID TThostFtdcInstallIDType // 安装编号
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	UserID TThostFtdcUserIDType // 用户标识
	VerifyCertNoFlag TThostFtdcYesNoIndicatorType // 验证客户证件号码标志
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 转帐金额
	FutureFetchAmount TThostFtdcTradeAmountType // 期货可取金额
	FeePayFlag TThostFtdcFeePayFlagType // 费用支付标志
	CustFee TThostFtdcCustFeeType // 应收客户费用
	BrokerFee TThostFtdcFutureFeeType // 应收期货公司费用
	Message TThostFtdcAddInfoType // 发送方给接收方的消息
	Digest TThostFtdcDigestType // 摘要
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	DeviceID TThostFtdcDeviceIDType // 渠道标志
	BankSecuAccType TThostFtdcBankAccTypeType // 期货单位帐号类型
	BrokerIDByBank TThostFtdcBankCodingForFutureType // 期货公司银行编码
	BankSecuAcc TThostFtdcBankAccountType // 期货单位帐号
	BankPwdFlag TThostFtdcPwdFlagType // 银行密码标志
	SecuPwdFlag TThostFtdcPwdFlagType // 期货资金密码核对标志
	OperNo TThostFtdcOperNoType // 交易柜员
	RequestID TThostFtdcRequestIDType // 请求编号
	TID TThostFtdcTIDType // 交易ID
	TransferStatus TThostFtdcTransferStatusType // 转账交易状态
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcRspUserAuthMethodField 用户发出获取安全安全登陆方法回复
type CThostFtdcRspUserAuthMethodField struct {
	UsableAuthMethod TThostFtdcCurrentAuthMethodType // 当前可以用的认证模式
}

// CThostFtdcRspUserLogin2Field 用户登录应答2
type CThostFtdcRspUserLogin2Field struct {
	TradingDay TThostFtdcDateType // 交易日
	LoginTime TThostFtdcTimeType // 登录成功时间
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	SystemName TThostFtdcSystemNameType // 交易系统名称
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	MaxOrderRef TThostFtdcOrderRefType // 最大报单引用
	SHFETime TThostFtdcTimeType // 上期所时间
	DCETime TThostFtdcTimeType // 大商所时间
	CZCETime TThostFtdcTimeType // 郑商所时间
	FFEXTime TThostFtdcTimeType // 中金所时间
	INETime TThostFtdcTimeType // 能源中心时间
	RandomString TThostFtdcRandomStringType // 随机串
}

// CThostFtdcRspUserLoginField 用户登录应答
type CThostFtdcRspUserLoginField struct {
	TradingDay TThostFtdcDateType // 交易日
	LoginTime TThostFtdcTimeType // 登录成功时间
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	SystemName TThostFtdcSystemNameType // 交易系统名称
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	MaxOrderRef TThostFtdcOrderRefType // 最大报单引用
	SHFETime TThostFtdcTimeType // 上期所时间
	DCETime TThostFtdcTimeType // 大商所时间
	CZCETime TThostFtdcTimeType // 郑商所时间
	FFEXTime TThostFtdcTimeType // 中金所时间
	INETime TThostFtdcTimeType // 能源中心时间
	SysVersion TThostFtdcSysVersionType // 后台版本信息
	GFEXTime TThostFtdcTimeType // 广期所时间
}

// CThostFtdcSPBMAddOnInterParameterField SPBM附加跨品种抵扣参数
type CThostFtdcSPBMAddOnInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	AddOnInterRateZ2 TThostFtdcRatioType // 品种间对锁仓附加费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
}

// CThostFtdcSPBMFutureParameterField SPBM期货合约保证金参数
type CThostFtdcSPBMFutureParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	Cvf TThostFtdcVolumeMultipleType // 期货合约因子
	TimeRange TThostFtdcTimeRangeType // 阶段标识
	MarginRate TThostFtdcRatioType // 品种保证金标准
	LockRateX TThostFtdcRatioType // 期货合约内部对锁仓费率折扣比例
	AddOnRate TThostFtdcRatioType // 提高保证金标准
	PreSettlementPrice TThostFtdcPriceType // 昨结算价
	AddOnLockRateX2 TThostFtdcRatioType // 期货合约内部对锁仓附加费率折扣比例
}

// CThostFtdcSPBMInterParameterField SPBM跨品种抵扣参数
type CThostFtdcSPBMInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	InterRateZ TThostFtdcRatioType // 品种间对锁仓费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
}

// CThostFtdcSPBMIntraParameterField SPBM品种内对锁仓折扣参数
type CThostFtdcSPBMIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	IntraRateY TThostFtdcRatioType // 品种内合约间对锁仓费率折扣比例
	AddOnIntraRateY2 TThostFtdcRatioType // 品种内合约间对锁仓附加费率折扣比例
}

// CThostFtdcSPBMInvestorPortfDefField 投资者套餐选择
type CThostFtdcSPBMInvestorPortfDefField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PortfolioDefID TThostFtdcPortfolioDefIDType // 组合保证金套餐代码
}

// CThostFtdcSPBMOptionParameterField SPBM期权合约保证金参数
type CThostFtdcSPBMOptionParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	Cvf TThostFtdcVolumeMultipleType // 期权合约因子
	DownPrice TThostFtdcPriceType // 期权冲抵价格
	Delta TThostFtdcDeltaType // Delta值
	SlimiDelta TThostFtdcDeltaType // 卖方期权风险转换最低值
	PreSettlementPrice TThostFtdcPriceType // 昨结算价
}

// CThostFtdcSPBMPortfDefinitionField 组合保证金套餐
type CThostFtdcSPBMPortfDefinitionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	PortfolioDefID TThostFtdcPortfolioDefIDType // 组合保证金套餐代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	IsSPBM TThostFtdcBoolType // 是否启用SPBM
}

// CThostFtdcSPMMInstParamField SPMM合约参数
type CThostFtdcSPMMInstParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InstMarginCalID TThostFtdcInstMarginCalIDType // SPMM合约保证金算法
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
}

// CThostFtdcSPMMProductParamField SPMM产品参数
type CThostFtdcSPMMProductParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcSPMMProductIDType // 产品代码
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
}

// CThostFtdcSecAgentACIDMapField 二级代理操作员银期权限
type CThostFtdcSecAgentACIDMapField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	AccountID TThostFtdcAccountIDType // 资金账户
	CurrencyID TThostFtdcCurrencyIDType // 币种
	BrokerSecAgentID TThostFtdcAccountIDType // 境外中介机构资金帐号
}

// CThostFtdcSecAgentCheckModeField 二级代理商资金校验模式
type CThostFtdcSecAgentCheckModeField struct {
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	CurrencyID TThostFtdcCurrencyIDType // 币种
	BrokerSecAgentID TThostFtdcAccountIDType // 境外中介机构资金帐号
	CheckSelfAccount TThostFtdcBoolType // 是否需要校验自己的资金账户
}

// CThostFtdcSecAgentTradeInfoField 二级代理商信息
type CThostFtdcSecAgentTradeInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	BrokerSecAgentID TThostFtdcAccountIDType // 境外中介机构资金帐号
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	LongCustomerName TThostFtdcLongIndividualNameType // 二级代理商姓名
}

// CThostFtdcSettlementInfoConfirmField 投资者结算结果确认信息
type CThostFtdcSettlementInfoConfirmField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ConfirmDate TThostFtdcDateType // 确认日期
	ConfirmTime TThostFtdcTimeType // 确认时间
	SettlementID TThostFtdcSettlementIDType // 结算编号
	AccountID TThostFtdcAccountIDType // 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcSettlementInfoField 投资者结算结果
type CThostFtdcSettlementInfoField struct {
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SequenceNo TThostFtdcSequenceNoType // 序号
	Content TThostFtdcContentType // 消息正文
	AccountID TThostFtdcAccountIDType // 投资者帐号
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcSettlementRefField 结算引用
type CThostFtdcSettlementRefField struct {
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
}

// CThostFtdcSpecificInstrumentField 指定的合约
type CThostFtdcSpecificInstrumentField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcStrikeOffsetField 当前期权合约执行偏移值的详细内容
type CThostFtdcStrikeOffsetField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Offset TThostFtdcMoneyType // 执行偏移值
	OffsetType TThostFtdcStrikeOffsetTypeType // 执行偏移类型
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcSuperUserField 管理用户
type CThostFtdcSuperUserField struct {
	UserID TThostFtdcUserIDType // 用户代码
	UserName TThostFtdcUserNameType // 用户名称
	Password TThostFtdcPasswordType // 密码
	IsActive TThostFtdcBoolType // 是否活跃
}

// CThostFtdcSuperUserFunctionField 管理用户功能权限
type CThostFtdcSuperUserFunctionField struct {
	UserID TThostFtdcUserIDType // 用户代码
	FunctionCode TThostFtdcFunctionCodeType // 功能代码
}

// CThostFtdcSyncDelaySwapField 延时换汇同步
type CThostFtdcSyncDelaySwapField struct {
	DelaySwapSeqNo TThostFtdcDepositSeqNoType // 换汇流水号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	FromCurrencyID TThostFtdcCurrencyIDType // 源币种
	FromAmount TThostFtdcMoneyType // 源金额
	FromFrozenSwap TThostFtdcMoneyType // 源换汇冻结金额(可用冻结)
	FromRemainSwap TThostFtdcMoneyType // 源剩余换汇额度(可提冻结)
	ToCurrencyID TThostFtdcCurrencyIDType // 目标币种
	ToAmount TThostFtdcMoneyType // 目标金额
	IsManualSwap TThostFtdcBoolType // 是否手工换汇
	IsAllRemainSetZero TThostFtdcBoolType // 是否将所有外币的剩余换汇额度设置为0
}

// CThostFtdcSyncDelaySwapFrozenField 换汇可提冻结
type CThostFtdcSyncDelaySwapFrozenField struct {
	DelaySwapSeqNo TThostFtdcDepositSeqNoType // 换汇流水号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	FromCurrencyID TThostFtdcCurrencyIDType // 源币种
	FromRemainSwap TThostFtdcMoneyType // 源剩余换汇额度(可提冻结)
	IsManualSwap TThostFtdcBoolType // 是否手工换汇
}

// CThostFtdcSyncDeltaDceCombInstrumentField 风险结算追平组合优先级
type CThostFtdcSyncDeltaDceCombInstrumentField struct {
	CombInstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	TradeGroupID TThostFtdcTradeGroupIDType // 成交组号
	CombHedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	CombinationType TThostFtdcDceCombinationTypeType // 组合类型
	Direction TThostFtdcDirectionType // 买卖
	ProductID TThostFtdcInstrumentIDType // 产品代码
	Xparameter TThostFtdcDiscountRatioType // 期权组合保证金比例
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaDepthMarketDataField 风险结算追平行情
type CThostFtdcSyncDeltaDepthMarketDataField struct {
	TradingDay TThostFtdcDateType // 交易日
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
	LastPrice TThostFtdcPriceType // 最新价
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	PreClosePrice TThostFtdcPriceType // 昨收盘
	PreOpenInterest TThostFtdcLargeVolumeType // 昨持仓量
	OpenPrice TThostFtdcPriceType // 今开盘
	HighestPrice TThostFtdcPriceType // 最高价
	LowestPrice TThostFtdcPriceType // 最低价
	Volume TThostFtdcVolumeType // 数量
	Turnover TThostFtdcMoneyType // 成交金额
	OpenInterest TThostFtdcLargeVolumeType // 持仓量
	ClosePrice TThostFtdcPriceType // 今收盘
	SettlementPrice TThostFtdcPriceType // 本次结算价
	UpperLimitPrice TThostFtdcPriceType // 涨停板价
	LowerLimitPrice TThostFtdcPriceType // 跌停板价
	PreDelta TThostFtdcRatioType // 昨虚实度
	CurrDelta TThostFtdcRatioType // 今虚实度
	UpdateTime TThostFtdcTimeType // 最后修改时间
	UpdateMillisec TThostFtdcMillisecType // 最后修改毫秒
	BidPrice1 TThostFtdcPriceType // 申买价一
	BidVolume1 TThostFtdcVolumeType // 申买量一
	AskPrice1 TThostFtdcPriceType // 申卖价一
	AskVolume1 TThostFtdcVolumeType // 申卖量一
	BidPrice2 TThostFtdcPriceType // 申买价二
	BidVolume2 TThostFtdcVolumeType // 申买量二
	AskPrice2 TThostFtdcPriceType // 申卖价二
	AskVolume2 TThostFtdcVolumeType // 申卖量二
	BidPrice3 TThostFtdcPriceType // 申买价三
	BidVolume3 TThostFtdcVolumeType // 申买量三
	AskPrice3 TThostFtdcPriceType // 申卖价三
	AskVolume3 TThostFtdcVolumeType // 申卖量三
	BidPrice4 TThostFtdcPriceType // 申买价四
	BidVolume4 TThostFtdcVolumeType // 申买量四
	AskPrice4 TThostFtdcPriceType // 申卖价四
	AskVolume4 TThostFtdcVolumeType // 申卖量四
	BidPrice5 TThostFtdcPriceType // 申买价五
	BidVolume5 TThostFtdcVolumeType // 申买量五
	AskPrice5 TThostFtdcPriceType // 申卖价五
	AskVolume5 TThostFtdcVolumeType // 申卖量五
	AveragePrice TThostFtdcPriceType // 当日均价
	ActionDay TThostFtdcDateType // 业务日期
	BandingUpperPrice TThostFtdcPriceType // 上带价
	BandingLowerPrice TThostFtdcPriceType // 下带价
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaEWarrantOffsetField 风险结算追平仓单折抵
type CThostFtdcSyncDeltaEWarrantOffsetField struct {
	TradingDay TThostFtdcTradeDateType // 交易日期
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	Direction TThostFtdcDirectionType // 买卖方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Volume TThostFtdcVolumeType // 数量
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaExchMarginRateField 风险结算追平交易所期货保证金率
type CThostFtdcSyncDeltaExchMarginRateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaIndexPriceField 风险结算追平现货指数
type CThostFtdcSyncDeltaIndexPriceField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ClosePrice TThostFtdcPriceType // 指数现货收盘价
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInfoField 风险结算追平信息
type CThostFtdcSyncDeltaInfoField struct {
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
	SyncDeltaStatus TThostFtdcSyncDeltaStatusType // 追平状态
	SyncDescription TThostFtdcSyncDescriptionType // 追平描述
	IsOnlyTrdDelta TThostFtdcBoolType // 是否只有资金追平
}

// CThostFtdcSyncDeltaInitInvstMarginField 投资者风险结算总保证金
type CThostFtdcSyncDeltaInitInvstMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	LastRiskTotalInvstMargin TThostFtdcMoneyType // 追平前总风险保证金
	LastRiskTotalExchMargin TThostFtdcMoneyType // 追平前交易所总风险保证金
	ThisSyncInvstMargin TThostFtdcMoneyType // 本次追平品种总保证金
	ThisSyncExchMargin TThostFtdcMoneyType // 本次追平品种交易所总保证金
	RemainRiskInvstMargin TThostFtdcMoneyType // 本次未追平品种总保证金
	RemainRiskExchMargin TThostFtdcMoneyType // 本次未追平品种交易所总保证金
	LastRiskSpecTotalInvstMargin TThostFtdcMoneyType // 追平前总特殊产品风险保证金
	LastRiskSpecTotalExchMargin TThostFtdcMoneyType // 追平前总特殊产品交易所风险保证金
	ThisSyncSpecInvstMargin TThostFtdcMoneyType // 本次追平品种特殊产品总保证金
	ThisSyncSpecExchMargin TThostFtdcMoneyType // 本次追平品种特殊产品交易所总保证金
	RemainRiskSpecInvstMargin TThostFtdcMoneyType // 本次未追平品种特殊产品总保证金
	RemainRiskSpecExchMargin TThostFtdcMoneyType // 本次未追平品种特殊产品交易所总保证金
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvestorSPMMModelField 风险结算追平投资者SPMM模板选择
type CThostFtdcSyncDeltaInvestorSPMMModelField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SPMMModelID TThostFtdcSPMMModelIDType // SPMM模板ID
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvstCommRateField 风险结算追平期货手续费率
type CThostFtdcSyncDeltaInvstCommRateField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvstMarginRateField 风险结算追平投资者期货保证金率
type CThostFtdcSyncDeltaInvstMarginRateField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	IsRelative TThostFtdcBoolType // 是否相对交易所收取
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvstMarginRateULField 风险结算追平期权标的调整保证金率
type CThostFtdcSyncDeltaInvstMarginRateULField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvstPosCombDtlField 风险结算追平组合持仓明细
type CThostFtdcSyncDeltaInvstPosCombDtlField struct {
	TradingDay TThostFtdcDateType // 交易日
	OpenDate TThostFtdcDateType // 开仓日期
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SettlementID TThostFtdcSettlementIDType // 结算编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	ComTradeID TThostFtdcTradeIDType // 组合编号
	TradeID TThostFtdcTradeIDType // 撮合编号
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Direction TThostFtdcDirectionType // 买卖
	TotalAmt TThostFtdcVolumeType // 持仓量
	Margin TThostFtdcMoneyType // 投资者保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	LegID TThostFtdcLegIDType // 单腿编号
	LegMultiple TThostFtdcLegMultipleType // 单腿乘数
	TradeGroupID TThostFtdcTradeGroupIDType // 成交组号
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaInvstPosDtlField 风险结算追平持仓明细
type CThostFtdcSyncDeltaInvstPosDtlField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Direction TThostFtdcDirectionType // 买卖
	OpenDate TThostFtdcDateType // 开仓日期
	TradeID TThostFtdcTradeIDType // 成交编号
	Volume TThostFtdcVolumeType // 数量
	OpenPrice TThostFtdcPriceType // 开仓价
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	TradeType TThostFtdcTradeTypeType // 成交类型
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CloseProfitByDate TThostFtdcMoneyType // 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType // 逐笔对冲平仓盈亏
	PositionProfitByDate TThostFtdcMoneyType // 逐日盯市持仓盈亏
	PositionProfitByTrade TThostFtdcMoneyType // 逐笔对冲持仓盈亏
	Margin TThostFtdcMoneyType // 投资者保证金
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	LastSettlementPrice TThostFtdcPriceType // 昨结算价
	SettlementPrice TThostFtdcPriceType // 结算价
	CloseVolume TThostFtdcVolumeType // 平仓量
	CloseAmount TThostFtdcMoneyType // 平仓金额
	TimeFirstVolume TThostFtdcVolumeType // 先开先平剩余数量
	SpecPosiType TThostFtdcSpecPosiTypeType // 特殊持仓标志
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaOptExchMarginField 风险结算追平中金现货期权交易所保证金率
type CThostFtdcSyncDeltaOptExchMarginField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	SShortMarginRatioByMoney TThostFtdcRatioType // 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType // 投机空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType // 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType // 保值空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType // 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType // 套利空头保证金调整系数
	MShortMarginRatioByMoney TThostFtdcRatioType // 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType // 做市商空头保证金调整系数
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaOptInvstCommRateField 风险结算追平期权手续费率
type CThostFtdcSyncDeltaOptInvstCommRateField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	StrikeRatioByMoney TThostFtdcRatioType // 执行手续费率
	StrikeRatioByVolume TThostFtdcRatioType // 执行手续费
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaOptInvstMarginField 风险结算追平中金现货期权投资者保证金率
type CThostFtdcSyncDeltaOptInvstMarginField struct {
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SShortMarginRatioByMoney TThostFtdcRatioType // 投机空头保证金调整系数
	SShortMarginRatioByVolume TThostFtdcMoneyType // 投机空头保证金调整系数
	HShortMarginRatioByMoney TThostFtdcRatioType // 保值空头保证金调整系数
	HShortMarginRatioByVolume TThostFtdcMoneyType // 保值空头保证金调整系数
	AShortMarginRatioByMoney TThostFtdcRatioType // 套利空头保证金调整系数
	AShortMarginRatioByVolume TThostFtdcMoneyType // 套利空头保证金调整系数
	IsRelative TThostFtdcBoolType // 是否跟随交易所收取
	MShortMarginRatioByMoney TThostFtdcRatioType // 做市商空头保证金调整系数
	MShortMarginRatioByVolume TThostFtdcMoneyType // 做市商空头保证金调整系数
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaProductExchRateField 风险结算追平交叉汇率
type CThostFtdcSyncDeltaProductExchRateField struct {
	ProductID TThostFtdcInstrumentIDType // 产品代码
	QuoteCurrencyID TThostFtdcCurrencyIDType // 报价币种类型
	ExchangeRate TThostFtdcExchangeRateType // 汇率
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaProductStatusField 风险结算追平产品信息
type CThostFtdcSyncDeltaProductStatusField struct {
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcInstrumentIDType // 产品代码
	ProductStatus TThostFtdcProductStatusType // 是否允许交易
}

// CThostFtdcSyncDeltaRCAMSCombProdInfoField 风险结算追平RCAMS产品组合信息
type CThostFtdcSyncDeltaRCAMSCombProdInfoField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcProductIDType // 产品代码
	CombProductID TThostFtdcProductIDType // 商品组代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSCombRuleDtlField 风险结算追平RCAMS策略组合规则明细
type CThostFtdcSyncDeltaRCAMSCombRuleDtlField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdGroup TThostFtdcProductIDType // 策略产品
	RuleId TThostFtdcRuleIdType // 策略id
	Priority TThostFtdcRCAMSPriorityType // 优先级
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	CombMargin TThostFtdcMoneyType // 组合保证金标准
	ExchangeInstID TThostFtdcExchangeInstIDType // 交易所组合合约代码
	LegID TThostFtdcLegIDType // 单腿编号
	LegInstrumentID TThostFtdcInstrumentIDType // 单腿合约代码
	Direction TThostFtdcDirectionType // 买卖方向
	LegMultiple TThostFtdcLegMultipleType // 单腿乘数
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSInstrParameterField 风险结算追平RCAMS同合约风险对冲参数
type CThostFtdcSyncDeltaRCAMSInstrParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcProductIDType // 产品代码
	HedgeRate TThostFtdcHedgeRateType // 同合约风险对冲比率
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSInterParameterField 风险结算追平RCAMS跨品种风险折抵参数
type CThostFtdcSyncDeltaRCAMSInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductGroupID TThostFtdcProductIDType // 商品群代码
	Priority TThostFtdcRCAMSPriorityType // 优先级
	CreditRate TThostFtdcHedgeRateType // 折抵率
	CombProduct1 TThostFtdcProductIDType // 产品组合代码1
	CombProduct2 TThostFtdcProductIDType // 产品组合代码2
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSIntraParameterField 风险结算追平RCAMS品种内风险对冲参数
type CThostFtdcSyncDeltaRCAMSIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	HedgeRate TThostFtdcHedgeRateType // 品种内对冲比率
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSInvstCombPosField 风险结算追平RCAMS策略组合持仓
type CThostFtdcSyncDeltaRCAMSInvstCombPosField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	PosiDirection TThostFtdcPosiDirectionType // 持仓多空方向
	CombInstrumentID TThostFtdcInstrumentIDType // 组合合约代码
	LegID TThostFtdcLegIDType // 单腿编号
	ExchangeInstID TThostFtdcExchangeInstIDType // 交易所组合合约代码
	TotalAmt TThostFtdcVolumeType // 持仓量
	ExchMargin TThostFtdcMoneyType // 交易所保证金
	Margin TThostFtdcMoneyType // 投资者保证金
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRCAMSSOptAdjParamField 风险结算追平RCAMS空头期权风险调整参数
type CThostFtdcSyncDeltaRCAMSSOptAdjParamField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	CombProductID TThostFtdcProductIDType // 产品组合代码
	HedgeFlag TThostFtdcHedgeFlagType // 投套标志
	AdjustValue TThostFtdcAdjustValueType // 空头期权风险调整标准
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRULEInstrParameterField 风险结算追平RULE合约保证金参数
type CThostFtdcSyncDeltaRULEInstrParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InstrumentClass TThostFtdcInstrumentClassType // 合约类型
	StdInstrumentID TThostFtdcInstrumentIDType // 标准合约
	BSpecRatio TThostFtdcRatioType // 投机买折算系数
	SSpecRatio TThostFtdcRatioType // 投机卖折算系数
	BHedgeRatio TThostFtdcRatioType // 套保买折算系数
	SHedgeRatio TThostFtdcRatioType // 套保卖折算系数
	BAddOnMargin TThostFtdcMoneyType // 买附加风险保证金
	SAddOnMargin TThostFtdcMoneyType // 卖附加风险保证金
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRULEInterParameterField 风险结算追平RULE跨品种抵扣参数
type CThostFtdcSyncDeltaRULEInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	InterRate TThostFtdcRatioType // 品种间对锁仓费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
	Leg1PropFactor TThostFtdcCommonIntType // 腿1比例系数
	Leg2PropFactor TThostFtdcCommonIntType // 腿2比例系数
	CommodityGroupID TThostFtdcCommodityGroupIDType // 商品群号
	CommodityGroupName TThostFtdcInstrumentNameType // 商品群名称
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaRULEIntraParameterField 风险结算追平RULE品种内对锁仓折扣参数
type CThostFtdcSyncDeltaRULEIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	StdInstrumentID TThostFtdcInstrumentIDType // 标准合约
	StdInstrMargin TThostFtdcMoneyType // 标准合约保证金
	UsualIntraRate TThostFtdcRatioType // 一般月份合约组合保证金系数
	DeliveryIntraRate TThostFtdcRatioType // 临近交割合约组合保证金系数
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMAddOnInterParamField 风险结算追平SPBM附加跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMAddOnInterParamField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	AddOnInterRateZ2 TThostFtdcRatioType // 品种间对锁仓附加费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMFutureParameterField 风险结算追平SPBM期货合约保证金参数
type CThostFtdcSyncDeltaSPBMFutureParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	Cvf TThostFtdcVolumeMultipleType // 期货合约因子
	TimeRange TThostFtdcTimeRangeType // 阶段标识
	MarginRate TThostFtdcRatioType // 品种保证金标准
	LockRateX TThostFtdcRatioType // 期货合约内部对锁仓费率折扣比例
	AddOnRate TThostFtdcRatioType // 提高保证金标准
	PreSettlementPrice TThostFtdcPriceType // 昨结算价
	AddOnLockRateX2 TThostFtdcRatioType // 期货合约内部对锁仓附加费率折扣比例
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMInterParameterField 风险结算追平SPBM跨品种抵扣参数
type CThostFtdcSyncDeltaSPBMInterParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SpreadId TThostFtdcSpreadIdType // 优先级
	InterRateZ TThostFtdcRatioType // 品种间对锁仓费率折扣比例
	Leg1ProdFamilyCode TThostFtdcInstrumentIDType // 第一腿构成品种
	Leg2ProdFamilyCode TThostFtdcInstrumentIDType // 第二腿构成品种
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMIntraParameterField 风险结算追平SPBM品种内对锁仓折扣参数
type CThostFtdcSyncDeltaSPBMIntraParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	IntraRateY TThostFtdcRatioType // 品种内合约间对锁仓费率折扣比例
	AddOnIntraRateY2 TThostFtdcRatioType // 品种内合约间对锁仓附加费率折扣比例
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMInvstPortfDefField 风险结算追平投资者SPBM套餐选择
type CThostFtdcSyncDeltaSPBMInvstPortfDefField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PortfolioDefID TThostFtdcPortfolioDefIDType // 组合保证金套餐代码
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMOptionParameterField 风险结算追平SPBM期权合约保证金参数
type CThostFtdcSyncDeltaSPBMOptionParameterField struct {
	TradingDay TThostFtdcDateType // 交易日
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	Cvf TThostFtdcVolumeMultipleType // 期权合约因子
	DownPrice TThostFtdcPriceType // 期权冲抵价格
	Delta TThostFtdcDeltaType // Delta值
	SlimiDelta TThostFtdcDeltaType // 卖方期权风险转换最低值
	PreSettlementPrice TThostFtdcPriceType // 昨结算价
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPBMPortfDefinitionField 风险结算追平SPBM组合保证金套餐
type CThostFtdcSyncDeltaSPBMPortfDefinitionField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	PortfolioDefID TThostFtdcPortfolioDefIDType // 组合保证金套餐代码
	ProdFamilyCode TThostFtdcInstrumentIDType // 品种代码
	IsSPBM TThostFtdcBoolType // 是否启用SPBM
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPMMInstParamField 风险结算追平SPMM合约参数
type CThostFtdcSyncDeltaSPMMInstParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	InstMarginCalID TThostFtdcInstMarginCalIDType // SPMM合约保证金算法
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPMMModelParamField 风险结算追平SPMM模板参数设置
type CThostFtdcSyncDeltaSPMMModelParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	SPMMModelID TThostFtdcSPMMModelIDType // SPMM模板ID
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
	IntraCommodityRate TThostFtdcSPMMDiscountRatioType // SPMM品种内跨期优惠系数
	InterCommodityRate TThostFtdcSPMMDiscountRatioType // SPMM品种间优惠系数
	OptionDiscountRate TThostFtdcSPMMDiscountRatioType // SPMM期权优惠系数
	MiniMarginRatio TThostFtdcSPMMDiscountRatioType // 商品群最小保证金比例
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaSPMMProductParamField 风险结算追平SPMM产品相关参数
type CThostFtdcSyncDeltaSPMMProductParamField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ProductID TThostFtdcSPMMProductIDType // 产品代码
	CommodityID TThostFtdcSPMMProductIDType // 商品组代码
	CommodityGroupID TThostFtdcSPMMProductIDType // 商品群代码
	ActionDirection TThostFtdcActionDirectionType // 操作标志
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDeltaTradingAccountField 风险结算追平资金
type CThostFtdcSyncDeltaTradingAccountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	PreMortgage TThostFtdcMoneyType // 上次质押金额
	PreCredit TThostFtdcMoneyType // 上次信用额度
	PreDeposit TThostFtdcMoneyType // 上次存款额
	PreBalance TThostFtdcMoneyType // 上次结算准备金
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	InterestBase TThostFtdcMoneyType // 利息基数
	Interest TThostFtdcMoneyType // 利息收入
	Deposit TThostFtdcMoneyType // 入金金额
	Withdraw TThostFtdcMoneyType // 出金金额
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CurrMargin TThostFtdcMoneyType // 当前保证金总额
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	Balance TThostFtdcMoneyType // 期货结算准备金
	Available TThostFtdcMoneyType // 可用资金
	WithdrawQuota TThostFtdcMoneyType // 可取资金
	Reserve TThostFtdcMoneyType // 基本准备金
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	Credit TThostFtdcMoneyType // 信用额度
	Mortgage TThostFtdcMoneyType // 质押金额
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	DeliveryMargin TThostFtdcMoneyType // 投资者交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType // 交易所交割保证金
	ReserveBalance TThostFtdcMoneyType // 保底期货结算准备金
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	PreFundMortgageIn TThostFtdcMoneyType // 上次货币质入金额
	PreFundMortgageOut TThostFtdcMoneyType // 上次货币质出金额
	FundMortgageIn TThostFtdcMoneyType // 货币质入金额
	FundMortgageOut TThostFtdcMoneyType // 货币质出金额
	FundMortgageAvailable TThostFtdcMoneyType // 货币质押余额
	MortgageableFund TThostFtdcMoneyType // 可质押货币金额
	SpecProductMargin TThostFtdcMoneyType // 特殊产品占用保证金
	SpecProductFrozenMargin TThostFtdcMoneyType // 特殊产品冻结保证金
	SpecProductCommission TThostFtdcMoneyType // 特殊产品手续费
	SpecProductFrozenCommission TThostFtdcMoneyType // 特殊产品冻结手续费
	SpecProductPositionProfit TThostFtdcMoneyType // 特殊产品持仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType // 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType // 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin TThostFtdcMoneyType // 特殊产品交易所保证金
	FrozenSwap TThostFtdcMoneyType // 延时换汇冻结金额
	RemainSwap TThostFtdcMoneyType // 剩余换汇额度
	SyncDeltaSequenceNo TThostFtdcSequenceNoType // 追平序号
}

// CThostFtdcSyncDepositField 出入金同步
type CThostFtdcSyncDepositField struct {
	DepositSeqNo TThostFtdcDepositSeqNoType // 出入金流水号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Deposit TThostFtdcMoneyType // 入金金额
	IsForce TThostFtdcBoolType // 是否强制进行
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	IsFromSopt TThostFtdcBoolType // 是否是个股期权内转
	TradingPassword TThostFtdcPasswordType // 资金密码
	IsSecAgentTranfer TThostFtdcBoolType // 是否二级代理商的内转
}

// CThostFtdcSyncFundMortgageField 货币质押同步
type CThostFtdcSyncFundMortgageField struct {
	MortgageSeqNo TThostFtdcDepositSeqNoType // 货币质押流水号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	FromCurrencyID TThostFtdcCurrencyIDType // 源币种
	MortgageAmount TThostFtdcMoneyType // 质押金额
	ToCurrencyID TThostFtdcCurrencyIDType // 目标币种
}

// CThostFtdcSyncSPBMParameterEndField 同步SPBM参数结束
type CThostFtdcSyncSPBMParameterEndField struct {
	TradingDay TThostFtdcDateType // 交易日
}

// CThostFtdcSyncStatusField 数据同步状态
type CThostFtdcSyncStatusField struct {
	TradingDay TThostFtdcDateType // 交易日
	DataSyncStatus TThostFtdcDataSyncStatusType // 数据同步状态
}

// CThostFtdcSyncingInstrumentCommissionRateField 正在同步中的合约手续费率
type CThostFtdcSyncingInstrumentCommissionRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OpenRatioByMoney TThostFtdcRatioType // 开仓手续费率
	OpenRatioByVolume TThostFtdcRatioType // 开仓手续费
	CloseRatioByMoney TThostFtdcRatioType // 平仓手续费率
	CloseRatioByVolume TThostFtdcRatioType // 平仓手续费
	CloseTodayRatioByMoney TThostFtdcRatioType // 平今手续费率
	CloseTodayRatioByVolume TThostFtdcRatioType // 平今手续费
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcSyncingInstrumentMarginRateField 正在同步中的合约保证金率
type CThostFtdcSyncingInstrumentMarginRateField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	LongMarginRatioByMoney TThostFtdcRatioType // 多头保证金率
	LongMarginRatioByVolume TThostFtdcMoneyType // 多头保证金费
	ShortMarginRatioByMoney TThostFtdcRatioType // 空头保证金率
	ShortMarginRatioByVolume TThostFtdcMoneyType // 空头保证金费
	IsRelative TThostFtdcBoolType // 是否相对交易所收取
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcSyncingInstrumentTradingRightField 正在同步中的合约交易权限
type CThostFtdcSyncingInstrumentTradingRightField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	TradingRight TThostFtdcTradingRightType // 交易权限
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcSyncingInvestorField 正在同步中的投资者
type CThostFtdcSyncingInvestorField struct {
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType // 投资者分组代码
	InvestorName TThostFtdcPartyNameType // 投资者名称
	IdentifiedCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	IsActive TThostFtdcBoolType // 是否活跃
	Telephone TThostFtdcTelephoneType // 联系电话
	Address TThostFtdcAddressType // 通讯地址
	OpenDate TThostFtdcDateType // 开户日期
	Mobile TThostFtdcMobileType // 手机
	CommModelID TThostFtdcInvestorIDType // 手续费率模板代码
	MarginModelID TThostFtdcInvestorIDType // 保证金率模板代码
	IsOrderFreq TThostFtdcEnumBoolType // 是否频率控制
	IsOpenVolLimit TThostFtdcEnumBoolType // 是否开仓限制
}

// CThostFtdcSyncingInvestorGroupField 正在同步中的投资者分组
type CThostFtdcSyncingInvestorGroupField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorGroupID TThostFtdcInvestorIDType // 投资者分组代码
	InvestorGroupName TThostFtdcInvestorGroupNameType // 投资者分组名称
}

// CThostFtdcSyncingInvestorPositionField 正在同步中的投资者持仓
type CThostFtdcSyncingInvestorPositionField struct {
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	PosiDirection TThostFtdcPosiDirectionType // 持仓多空方向
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	PositionDate TThostFtdcPositionDateType // 持仓日期
	YdPosition TThostFtdcVolumeType // 上日持仓
	Position TThostFtdcVolumeType // 今日持仓
	LongFrozen TThostFtdcVolumeType // 多头冻结
	ShortFrozen TThostFtdcVolumeType // 空头冻结
	LongFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	ShortFrozenAmount TThostFtdcMoneyType // 开仓冻结金额
	OpenVolume TThostFtdcVolumeType // 开仓量
	CloseVolume TThostFtdcVolumeType // 平仓量
	OpenAmount TThostFtdcMoneyType // 开仓金额
	CloseAmount TThostFtdcMoneyType // 平仓金额
	PositionCost TThostFtdcMoneyType // 持仓成本
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	UseMargin TThostFtdcMoneyType // 占用的保证金
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	PreSettlementPrice TThostFtdcPriceType // 上次结算价
	SettlementPrice TThostFtdcPriceType // 本次结算价
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	OpenCost TThostFtdcMoneyType // 开仓成本
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	CombPosition TThostFtdcVolumeType // 组合成交形成的持仓
	CombLongFrozen TThostFtdcVolumeType // 组合多头冻结
	CombShortFrozen TThostFtdcVolumeType // 组合空头冻结
	CloseProfitByDate TThostFtdcMoneyType // 逐日盯市平仓盈亏
	CloseProfitByTrade TThostFtdcMoneyType // 逐笔对冲平仓盈亏
	TodayPosition TThostFtdcVolumeType // 今日持仓
	MarginRateByMoney TThostFtdcRatioType // 保证金率
	MarginRateByVolume TThostFtdcRatioType // 保证金率(按手数)
	StrikeFrozen TThostFtdcVolumeType // 执行冻结
	StrikeFrozenAmount TThostFtdcMoneyType // 执行冻结金额
	AbandonFrozen TThostFtdcVolumeType // 放弃执行冻结
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	YdStrikeFrozen TThostFtdcVolumeType // 执行冻结的昨仓
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	PositionCostOffset TThostFtdcMoneyType // 持仓成本差值
	TasPosition TThostFtdcVolumeType // tas持仓手数
	TasPositionCost TThostFtdcMoneyType // tas持仓成本
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
}

// CThostFtdcSyncingTradingAccountField 正在同步中的交易账号
type CThostFtdcSyncingTradingAccountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	PreMortgage TThostFtdcMoneyType // 上次质押金额
	PreCredit TThostFtdcMoneyType // 上次信用额度
	PreDeposit TThostFtdcMoneyType // 上次存款额
	PreBalance TThostFtdcMoneyType // 上次结算准备金
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	InterestBase TThostFtdcMoneyType // 利息基数
	Interest TThostFtdcMoneyType // 利息收入
	Deposit TThostFtdcMoneyType // 入金金额
	Withdraw TThostFtdcMoneyType // 出金金额
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CurrMargin TThostFtdcMoneyType // 当前保证金总额
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	Balance TThostFtdcMoneyType // 期货结算准备金
	Available TThostFtdcMoneyType // 可用资金
	WithdrawQuota TThostFtdcMoneyType // 可取资金
	Reserve TThostFtdcMoneyType // 基本准备金
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	Credit TThostFtdcMoneyType // 信用额度
	Mortgage TThostFtdcMoneyType // 质押金额
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	DeliveryMargin TThostFtdcMoneyType // 投资者交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType // 交易所交割保证金
	ReserveBalance TThostFtdcMoneyType // 保底期货结算准备金
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	PreFundMortgageIn TThostFtdcMoneyType // 上次货币质入金额
	PreFundMortgageOut TThostFtdcMoneyType // 上次货币质出金额
	FundMortgageIn TThostFtdcMoneyType // 货币质入金额
	FundMortgageOut TThostFtdcMoneyType // 货币质出金额
	FundMortgageAvailable TThostFtdcMoneyType // 货币质押余额
	MortgageableFund TThostFtdcMoneyType // 可质押货币金额
	SpecProductMargin TThostFtdcMoneyType // 特殊产品占用保证金
	SpecProductFrozenMargin TThostFtdcMoneyType // 特殊产品冻结保证金
	SpecProductCommission TThostFtdcMoneyType // 特殊产品手续费
	SpecProductFrozenCommission TThostFtdcMoneyType // 特殊产品冻结手续费
	SpecProductPositionProfit TThostFtdcMoneyType // 特殊产品持仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType // 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType // 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin TThostFtdcMoneyType // 特殊产品交易所保证金
	FrozenSwap TThostFtdcMoneyType // 延时换汇冻结金额
	RemainSwap TThostFtdcMoneyType // 剩余换汇额度
}

// CThostFtdcSyncingTradingCodeField 正在同步中的交易代码
type CThostFtdcSyncingTradingCodeField struct {
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ClientID TThostFtdcClientIDType // 客户代码
	IsActive TThostFtdcBoolType // 是否活跃
	ClientIDType TThostFtdcClientIDTypeType // 交易编码类型
}

// CThostFtdcThostUserFunctionField Thost终端用户功能权限
type CThostFtdcThostUserFunctionField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	ThostFunctionCode TThostFtdcThostFunctionCodeType // Thost终端功能代码
}

// CThostFtdcTradeField 成交
type CThostFtdcTradeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	reserve1 TThostFtdcOldInstrumentIDType // 保留的无效字段
	OrderRef TThostFtdcOrderRefType // 报单引用
	UserID TThostFtdcUserIDType // 用户代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TradeID TThostFtdcTradeIDType // 成交编号
	Direction TThostFtdcDirectionType // 买卖方向
	OrderSysID TThostFtdcOrderSysIDType // 报单编号
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	ClientID TThostFtdcClientIDType // 客户代码
	TradingRole TThostFtdcTradingRoleType // 交易角色
	reserve2 TThostFtdcOldExchangeInstIDType // 保留的无效字段
	OffsetFlag TThostFtdcOffsetFlagType // 开平标志
	HedgeFlag TThostFtdcHedgeFlagType // 投机套保标志
	Price TThostFtdcPriceType // 价格
	Volume TThostFtdcVolumeType // 数量
	TradeDate TThostFtdcDateType // 成交时期
	TradeTime TThostFtdcTimeType // 成交时间
	TradeType TThostFtdcTradeTypeType // 成交类型
	PriceSource TThostFtdcPriceSourceType // 成交价来源
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	ClearingPartID TThostFtdcParticipantIDType // 结算会员编号
	BusinessUnit TThostFtdcBusinessUnitType // 业务单元
	SequenceNo TThostFtdcSequenceNoType // 序号
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	BrokerOrderSeq TThostFtdcSequenceNoType // 经纪公司报单编号
	TradeSource TThostFtdcTradeSourceType // 成交来源
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
	InstrumentID TThostFtdcInstrumentIDType // 合约代码
	ExchangeInstID TThostFtdcExchangeInstIDType // 合约在交易所的代码
}

// CThostFtdcTradeParamField 交易参数
type CThostFtdcTradeParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	TradeParamID TThostFtdcTradeParamIDType // 参数代码
	TradeParamValue TThostFtdcSettlementParamValueType // 参数代码值
	Memo TThostFtdcMemoType // 备注
}

// CThostFtdcTraderAssignField 席位与交易中心对应关系
type CThostFtdcTraderAssignField struct {
	BrokerID TThostFtdcBrokerIDType // 应用单元代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcTraderField 交易所交易员
type CThostFtdcTraderField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	Password TThostFtdcPasswordType // 密码
	InstallCount TThostFtdcInstallCountType // 安装数量
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	OrderCancelAlg TThostFtdcOrderCancelAlgType // 撤单时选择席位算法
	TradeInstallCount TThostFtdcInstallCountType // 交易报盘安装数量
	MDInstallCount TThostFtdcInstallCountType // 行情报盘安装数量
}

// CThostFtdcTraderOfferField 交易所交易员报盘机
type CThostFtdcTraderOfferField struct {
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	TraderID TThostFtdcTraderIDType // 交易所交易员代码
	ParticipantID TThostFtdcParticipantIDType // 会员代码
	Password TThostFtdcPasswordType // 密码
	InstallID TThostFtdcInstallIDType // 安装编号
	OrderLocalID TThostFtdcOrderLocalIDType // 本地报单编号
	TraderConnectStatus TThostFtdcTraderConnectStatusType // 交易所交易员连接状态
	ConnectRequestDate TThostFtdcDateType // 发出连接请求的日期
	ConnectRequestTime TThostFtdcTimeType // 发出连接请求的时间
	LastReportDate TThostFtdcDateType // 上次报告日期
	LastReportTime TThostFtdcTimeType // 上次报告时间
	ConnectDate TThostFtdcDateType // 完成连接日期
	ConnectTime TThostFtdcTimeType // 完成连接时间
	StartDate TThostFtdcDateType // 启动日期
	StartTime TThostFtdcTimeType // 启动时间
	TradingDay TThostFtdcDateType // 交易日
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	MaxTradeID TThostFtdcTradeIDType // 本席位最大成交编号
	MaxOrderMessageReference TThostFtdcReturnCodeType // 本席位最大报单备拷
	OrderCancelAlg TThostFtdcOrderCancelAlgType // 撤单时选择席位算法
}

// CThostFtdcTradingAccountField 资金账户
type CThostFtdcTradingAccountField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	PreMortgage TThostFtdcMoneyType // 上次质押金额
	PreCredit TThostFtdcMoneyType // 上次信用额度
	PreDeposit TThostFtdcMoneyType // 上次存款额
	PreBalance TThostFtdcMoneyType // 上次结算准备金
	PreMargin TThostFtdcMoneyType // 上次占用的保证金
	InterestBase TThostFtdcMoneyType // 利息基数
	Interest TThostFtdcMoneyType // 利息收入
	Deposit TThostFtdcMoneyType // 入金金额
	Withdraw TThostFtdcMoneyType // 出金金额
	FrozenMargin TThostFtdcMoneyType // 冻结的保证金
	FrozenCash TThostFtdcMoneyType // 冻结的资金
	FrozenCommission TThostFtdcMoneyType // 冻结的手续费
	CurrMargin TThostFtdcMoneyType // 当前保证金总额
	CashIn TThostFtdcMoneyType // 资金差额
	Commission TThostFtdcMoneyType // 手续费
	CloseProfit TThostFtdcMoneyType // 平仓盈亏
	PositionProfit TThostFtdcMoneyType // 持仓盈亏
	Balance TThostFtdcMoneyType // 期货结算准备金
	Available TThostFtdcMoneyType // 可用资金
	WithdrawQuota TThostFtdcMoneyType // 可取资金
	Reserve TThostFtdcMoneyType // 基本准备金
	TradingDay TThostFtdcDateType // 交易日
	SettlementID TThostFtdcSettlementIDType // 结算编号
	Credit TThostFtdcMoneyType // 信用额度
	Mortgage TThostFtdcMoneyType // 质押金额
	ExchangeMargin TThostFtdcMoneyType // 交易所保证金
	DeliveryMargin TThostFtdcMoneyType // 投资者交割保证金
	ExchangeDeliveryMargin TThostFtdcMoneyType // 交易所交割保证金
	ReserveBalance TThostFtdcMoneyType // 保底期货结算准备金
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	PreFundMortgageIn TThostFtdcMoneyType // 上次货币质入金额
	PreFundMortgageOut TThostFtdcMoneyType // 上次货币质出金额
	FundMortgageIn TThostFtdcMoneyType // 货币质入金额
	FundMortgageOut TThostFtdcMoneyType // 货币质出金额
	FundMortgageAvailable TThostFtdcMoneyType // 货币质押余额
	MortgageableFund TThostFtdcMoneyType // 可质押货币金额
	SpecProductMargin TThostFtdcMoneyType // 特殊产品占用保证金
	SpecProductFrozenMargin TThostFtdcMoneyType // 特殊产品冻结保证金
	SpecProductCommission TThostFtdcMoneyType // 特殊产品手续费
	SpecProductFrozenCommission TThostFtdcMoneyType // 特殊产品冻结手续费
	SpecProductPositionProfit TThostFtdcMoneyType // 特殊产品持仓盈亏
	SpecProductCloseProfit TThostFtdcMoneyType // 特殊产品平仓盈亏
	SpecProductPositionProfitByAlg TThostFtdcMoneyType // 根据持仓盈亏算法计算的特殊产品持仓盈亏
	SpecProductExchangeMargin TThostFtdcMoneyType // 特殊产品交易所保证金
	BizType TThostFtdcBizTypeType // 业务类型
	FrozenSwap TThostFtdcMoneyType // 延时换汇冻结金额
	RemainSwap TThostFtdcMoneyType // 剩余换汇额度
}

// CThostFtdcTradingAccountPasswordField 资金账户口令域
type CThostFtdcTradingAccountPasswordField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 密码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcTradingAccountPasswordUpdateField 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	OldPassword TThostFtdcPasswordType // 原来的口令
	NewPassword TThostFtdcPasswordType // 新的口令
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcTradingAccountPasswordUpdateV1Field 资金账户口令变更域
type CThostFtdcTradingAccountPasswordUpdateV1Field struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	OldPassword TThostFtdcPasswordType // 原来的口令
	NewPassword TThostFtdcPasswordType // 新的口令
}

// CThostFtdcTradingAccountReserveField 资金账户基本准备金
type CThostFtdcTradingAccountReserveField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Reserve TThostFtdcMoneyType // 基本准备金
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcTradingCodeField 交易编码
type CThostFtdcTradingCodeField struct {
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	ExchangeID TThostFtdcExchangeIDType // 交易所代码
	ClientID TThostFtdcClientIDType // 客户代码
	IsActive TThostFtdcBoolType // 是否活跃
	ClientIDType TThostFtdcClientIDTypeType // 交易编码类型
	BranchID TThostFtdcBranchIDType // 营业部编号
	BizType TThostFtdcBizTypeType // 业务类型
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcTradingNoticeField 用户事件通知
type CThostFtdcTradingNoticeField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorRange TThostFtdcInvestorRangeType // 投资者范围
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SequenceSeries TThostFtdcSequenceSeriesType // 序列系列号
	UserID TThostFtdcUserIDType // 用户代码
	SendTime TThostFtdcTimeType // 发送时间
	SequenceNo TThostFtdcSequenceNoType // 序列号
	FieldContent TThostFtdcContentType // 消息正文
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcTradingNoticeInfoField 用户事件通知信息
type CThostFtdcTradingNoticeInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	SendTime TThostFtdcTimeType // 发送时间
	FieldContent TThostFtdcContentType // 消息正文
	SequenceSeries TThostFtdcSequenceSeriesType // 序列系列号
	SequenceNo TThostFtdcSequenceNoType // 序列号
	InvestUnitID TThostFtdcInvestUnitIDType // 投资单元代码
}

// CThostFtdcTransferBankField 转帐银行
type CThostFtdcTransferBankField struct {
	BankID TThostFtdcBankIDType // 银行代码
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码
	BankName TThostFtdcBankNameType // 银行名称
	IsActive TThostFtdcBoolType // 是否活跃
}

// CThostFtdcTransferBankToFutureReqField 银行资金转期货请求，TradeCode=202001
type CThostFtdcTransferBankToFutureReqField struct {
	FutureAccount TThostFtdcAccountIDType // 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType // 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType // 密码
	TradeAmt TThostFtdcMoneyType // 转账金额
	CustFee TThostFtdcMoneyType // 客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType // 币种：RMB-人民币 USD-美圆 HKD-港元
}

// CThostFtdcTransferBankToFutureRspField 银行资金转期货请求响应
type CThostFtdcTransferBankToFutureRspField struct {
	RetCode TThostFtdcRetCodeType // 响应代码
	RetInfo TThostFtdcRetInfoType // 响应信息
	FutureAccount TThostFtdcAccountIDType // 资金账户
	TradeAmt TThostFtdcMoneyType // 转帐金额
	CustFee TThostFtdcMoneyType // 应收客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType // 币种
}

// CThostFtdcTransferFutureToBankReqField 期货资金转银行请求，TradeCode=202002
type CThostFtdcTransferFutureToBankReqField struct {
	FutureAccount TThostFtdcAccountIDType // 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType // 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType // 密码
	TradeAmt TThostFtdcMoneyType // 转账金额
	CustFee TThostFtdcMoneyType // 客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType // 币种：RMB-人民币 USD-美圆 HKD-港元
}

// CThostFtdcTransferFutureToBankRspField 期货资金转银行请求响应
type CThostFtdcTransferFutureToBankRspField struct {
	RetCode TThostFtdcRetCodeType // 响应代码
	RetInfo TThostFtdcRetInfoType // 响应信息
	FutureAccount TThostFtdcAccountIDType // 资金账户
	TradeAmt TThostFtdcMoneyType // 转帐金额
	CustFee TThostFtdcMoneyType // 应收客户手续费
	CurrencyCode TThostFtdcCurrencyCodeType // 币种
}

// CThostFtdcTransferHeaderField 银期转帐报文头
type CThostFtdcTransferHeaderField struct {
	Version TThostFtdcVersionType // 版本号，常量，1.0
	TradeCode TThostFtdcTradeCodeType // 交易代码，必填
	TradeDate TThostFtdcTradeDateType // 交易日期，必填，格式：yyyymmdd
	TradeTime TThostFtdcTradeTimeType // 交易时间，必填，格式：hhmmss
	TradeSerial TThostFtdcTradeSerialType // 发起方流水号，N/A
	FutureID TThostFtdcFutureIDType // 期货公司代码，必填
	BankID TThostFtdcBankIDType // 银行代码，根据查询银行得到，必填
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码，根据查询银行得到，必填
	OperNo TThostFtdcOperNoType // 操作员，N/A
	DeviceID TThostFtdcDeviceIDType // 交易设备类型，N/A
	RecordNum TThostFtdcRecordNumType // 记录数，N/A
	SessionID TThostFtdcSessionIDType // 会话编号，N/A
	RequestID TThostFtdcRequestIDType // 请求编号，N/A
}

// CThostFtdcTransferQryBankReqField 查询银行资金请求，TradeCode=204002
type CThostFtdcTransferQryBankReqField struct {
	FutureAccount TThostFtdcAccountIDType // 期货资金账户
	FuturePwdFlag TThostFtdcFuturePwdFlagType // 密码标志
	FutureAccPwd TThostFtdcFutureAccPwdType // 密码
	CurrencyCode TThostFtdcCurrencyCodeType // 币种：RMB-人民币 USD-美圆 HKD-港元
}

// CThostFtdcTransferQryBankRspField 查询银行资金请求响应
type CThostFtdcTransferQryBankRspField struct {
	RetCode TThostFtdcRetCodeType // 响应代码
	RetInfo TThostFtdcRetInfoType // 响应信息
	FutureAccount TThostFtdcAccountIDType // 资金账户
	TradeAmt TThostFtdcMoneyType // 银行余额
	UseAmt TThostFtdcMoneyType // 银行可用余额
	FetchAmt TThostFtdcMoneyType // 银行可取余额
	CurrencyCode TThostFtdcCurrencyCodeType // 币种
}

// CThostFtdcTransferQryDetailReqField 查询银行交易明细请求，TradeCode=204999
type CThostFtdcTransferQryDetailReqField struct {
	FutureAccount TThostFtdcAccountIDType // 期货资金账户
}

// CThostFtdcTransferQryDetailRspField 查询银行交易明细请求响应
type CThostFtdcTransferQryDetailRspField struct {
	TradeDate TThostFtdcDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	TradeCode TThostFtdcTradeCodeType // 交易代码
	FutureSerial TThostFtdcTradeSerialNoType // 期货流水号
	FutureID TThostFtdcFutureIDType // 期货公司代码
	FutureAccount TThostFtdcFutureAccountType // 资金帐号
	BankSerial TThostFtdcTradeSerialNoType // 银行流水号
	BankID TThostFtdcBankIDType // 银行代码
	BankBrchID TThostFtdcBankBrchIDType // 银行分中心代码
	BankAccount TThostFtdcBankAccountType // 银行账号
	CertCode TThostFtdcCertCodeType // 证件号码
	CurrencyCode TThostFtdcCurrencyCodeType // 货币代码
	TxAmount TThostFtdcMoneyType // 发生金额
	Flag TThostFtdcTransferValidFlagType // 有效标志
}

// CThostFtdcTransferSerialField 银期转账交易流水表
type CThostFtdcTransferSerialField struct {
	PlateSerial TThostFtdcPlateSerialType // 平台流水号
	TradeDate TThostFtdcTradeDateType // 交易发起方日期
	TradingDay TThostFtdcDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	TradeCode TThostFtdcTradeCodeType // 交易代码
	SessionID TThostFtdcSessionIDType // 会话编号
	BankID TThostFtdcBankIDType // 银行编码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构编码
	BankAccType TThostFtdcBankAccTypeType // 银行帐号类型
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankSerial TThostFtdcBankSerialType // 银行流水号
	BrokerID TThostFtdcBrokerIDType // 期货公司编码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	FutureAccType TThostFtdcFutureAccTypeType // 期货公司帐号类型
	AccountID TThostFtdcAccountIDType // 投资者帐号
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	FutureSerial TThostFtdcFutureSerialType // 期货公司流水号
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	TradeAmount TThostFtdcTradeAmountType // 交易金额
	CustFee TThostFtdcCustFeeType // 应收客户费用
	BrokerFee TThostFtdcFutureFeeType // 应收期货公司费用
	AvailabilityFlag TThostFtdcAvailabilityFlagType // 有效标志
	OperatorCode TThostFtdcOperatorCodeType // 操作员
	BankNewAccount TThostFtdcBankAccountType // 新银行帐号
	ErrorID TThostFtdcErrorIDType // 错误代码
	ErrorMsg TThostFtdcErrorMsgType // 错误信息
}

// CThostFtdcUserIPField 用户IP
type CThostFtdcUserIPField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	reserve2 TThostFtdcOldIPAddressType // 保留的无效字段
	MacAddress TThostFtdcMacAddressType // Mac地址
	IPAddress TThostFtdcIPAddressType // IP地址
	IPMask TThostFtdcIPAddressType // IP地址掩码
}

// CThostFtdcUserLogoutField 用户登出请求
type CThostFtdcUserLogoutField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
}

// CThostFtdcUserPasswordUpdateField 用户口令变更
type CThostFtdcUserPasswordUpdateField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	OldPassword TThostFtdcPasswordType // 原来的口令
	NewPassword TThostFtdcPasswordType // 新的口令
}

// CThostFtdcUserRightField 用户权限
type CThostFtdcUserRightField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	UserRightType TThostFtdcUserRightTypeType // 客户权限类型
	IsForbidden TThostFtdcBoolType // 是否禁止
}

// CThostFtdcUserRightsAssignField 灾备中心交易权限
type CThostFtdcUserRightsAssignField struct {
	BrokerID TThostFtdcBrokerIDType // 应用单元代码
	UserID TThostFtdcUserIDType // 用户代码
	DRIdentityID TThostFtdcDRIdentityIDType // 交易中心代码
}

// CThostFtdcUserSessionField 用户会话
type CThostFtdcUserSessionField struct {
	FrontID TThostFtdcFrontIDType // 前置编号
	SessionID TThostFtdcSessionIDType // 会话编号
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	LoginDate TThostFtdcDateType // 登录日期
	LoginTime TThostFtdcTimeType // 登录时间
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	UserProductInfo TThostFtdcProductInfoType // 用户端产品信息
	InterfaceProductInfo TThostFtdcProductInfoType // 接口端产品信息
	ProtocolInfo TThostFtdcProtocolInfoType // 协议信息
	MacAddress TThostFtdcMacAddressType // Mac地址
	LoginRemark TThostFtdcLoginRemarkType // 登录备注
	IPAddress TThostFtdcIPAddressType // IP地址
}

// CThostFtdcUserSystemInfoField 用户系统信息
type CThostFtdcUserSystemInfoField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	UserID TThostFtdcUserIDType // 用户代码
	ClientSystemInfoLen TThostFtdcSystemInfoLenType // 用户端系统内部信息长度
	ClientSystemInfo TThostFtdcClientSystemInfoType // 用户端系统内部信息
	reserve1 TThostFtdcOldIPAddressType // 保留的无效字段
	ClientIPPort TThostFtdcIPPortType // 终端IP端口
	ClientLoginTime TThostFtdcTimeType // 登录成功时间
	ClientAppID TThostFtdcAppIDType // App代码
	ClientPublicIP TThostFtdcIPAddressType // 用户公网IP
	ClientLoginRemark TThostFtdcClientLoginRemarkType // 客户登录备注2
}

// CThostFtdcVerifyCustInfoField 验证客户信息
type CThostFtdcVerifyCustInfoField struct {
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcVerifyFuturePasswordAndCustInfoField 验证期货资金密码和客户信息
type CThostFtdcVerifyFuturePasswordAndCustInfoField struct {
	CustomerName TThostFtdcIndividualNameType // 客户姓名
	IdCardType TThostFtdcIdCardTypeType // 证件类型
	IdentifiedCardNo TThostFtdcIdentifiedCardNoType // 证件号码
	CustType TThostFtdcCustTypeType // 客户类型
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
	LongCustomerName TThostFtdcLongIndividualNameType // 长客户姓名
}

// CThostFtdcVerifyFuturePasswordField 验证期货资金密码
type CThostFtdcVerifyFuturePasswordField struct {
	TradeCode TThostFtdcTradeCodeType // 业务功能码
	BankID TThostFtdcBankIDType // 银行代码
	BankBranchID TThostFtdcBankBrchIDType // 银行分支机构代码
	BrokerID TThostFtdcBrokerIDType // 期商代码
	BrokerBranchID TThostFtdcFutureBranchIDType // 期商分支机构代码
	TradeDate TThostFtdcTradeDateType // 交易日期
	TradeTime TThostFtdcTradeTimeType // 交易时间
	BankSerial TThostFtdcBankSerialType // 银行流水号
	TradingDay TThostFtdcTradeDateType // 交易系统日期
	PlateSerial TThostFtdcSerialType // 银期平台消息流水号
	LastFragment TThostFtdcLastFragmentType // 最后分片标志
	SessionID TThostFtdcSessionIDType // 会话号
	AccountID TThostFtdcAccountIDType // 投资者帐号
	Password TThostFtdcPasswordType // 期货密码
	BankAccount TThostFtdcBankAccountType // 银行帐号
	BankPassWord TThostFtdcPasswordType // 银行密码
	InstallID TThostFtdcInstallIDType // 安装编号
	TID TThostFtdcTIDType // 交易ID
	CurrencyID TThostFtdcCurrencyIDType // 币种代码
}

// CThostFtdcVerifyInvestorPasswordField 校验投资者密码
type CThostFtdcVerifyInvestorPasswordField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	InvestorID TThostFtdcInvestorIDType // 投资者代码
	Password TThostFtdcPasswordType // 密码
}

// CThostFtdcWithDrawParamField 可提控制参数
type CThostFtdcWithDrawParamField struct {
	BrokerID TThostFtdcBrokerIDType // 经纪公司代码
	AccountID TThostFtdcAccountIDType // 投资者帐号
	WithDrawParamID TThostFtdcWithDrawParamIDType // 参数代码
	WithDrawParamValue TThostFtdcWithDrawParamValueType // 参数代码值
}
