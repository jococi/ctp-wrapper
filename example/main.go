package main

// 多实例支持说明：
// 本封装库支持创建多个独立的 API 实例，每个实例的回调会正确路由，不会串台。
//
// 创建多实例的要点：
// 1. 每个实例必须使用不同的 flowPath（用于存储订阅信息文件）
//    例如："./flow1", "./flow2", "./flow_strategy1" 等
// 2. 每个实例可以设置不同的 SPI 实现（不同的策略）
// 3. 每个实例可以连接相同或不同的前置地址
// 4. 每个实例有独立的 userData ID，回调通过 userData 正确路由到对应实例
//
// 示例：
//   mdApi1 := ctpgo.NewMdApi("./flow1", false, false)
//   mdApi1.SetSpi(&Strategy1Spi{})
//   mdApi1.RegisterFront("tcp://...")
//
//   mdApi2 := ctpgo.NewMdApi("./flow2", false, false)
//   mdApi2.SetSpi(&Strategy2Spi{})
//   mdApi2.RegisterFront("tcp://...")
//
// 两个实例可以同时运行，回调会正确路由到各自的 SPI。

import (
	"fmt"
	"log"
	"os"
	"time"

	"ctpgo"

	"github.com/joho/godotenv"
)

// MyMdSpi 自定义行情回调实现
// 嵌入 DefaultMdSpi 后，只需实现需要的方法
type MyMdSpi struct {
	ctpgo.DefaultMdSpi
}

// OnFrontConnected 连接成功回调
func (s *MyMdSpi) OnFrontConnected() {
	fmt.Println("行情服务器连接成功")
}

// OnFrontDisconnected 连接断开回调
func (s *MyMdSpi) OnFrontDisconnected(nReason int32) {
	fmt.Printf("行情服务器断开连接，原因: %d\n", nReason)
}

// OnRspUserLogin 登录响应
func (s *MyMdSpi) OnRspUserLogin(pRspUserLogin *ctpgo.CThostFtdcRspUserLoginField, pRspInfo *ctpgo.CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		fmt.Printf("登录失败: %s\n", ctpgo.GB18030(pRspInfo.ErrorMsg[:]))
		return
	}
	if pRspUserLogin != nil {
		fmt.Printf("登录成功，交易日: %s\n", ctpgo.BytesToString(pRspUserLogin.TradingDay[:]))
	}
}

// OnRtnDepthMarketData 行情数据回调
func (s *MyMdSpi) OnRtnDepthMarketData(pDepthMarketData *ctpgo.CThostFtdcDepthMarketDataField) {
	if pDepthMarketData == nil {
		return
	}
	instrumentID := ctpgo.BytesToString(pDepthMarketData.InstrumentID[:])
	lastPrice := pDepthMarketData.LastPrice
	fmt.Printf("行情更新: %s, 最新价: %.2f\n", instrumentID, lastPrice)
}

// MyTraderSpi 自定义交易回调实现
// 嵌入 DefaultTraderSpi 后，只需实现需要的方法
type MyTraderSpi struct {
	ctpgo.DefaultTraderSpi
	authChan chan bool // 认证结果通道
}

// OnFrontConnected 连接成功回调
func (s *MyTraderSpi) OnFrontConnected() {
	fmt.Println("交易服务器连接成功")
}

// OnRspAuthenticate 认证响应
func (s *MyTraderSpi) OnRspAuthenticate(pRspAuthenticateField *ctpgo.CThostFtdcRspAuthenticateField, pRspInfo *ctpgo.CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		fmt.Printf("认证失败: %s\n", ctpgo.GB18030(pRspInfo.ErrorMsg[:]))
		if s.authChan != nil {
			s.authChan <- false
		}
		return
	}
	fmt.Println("认证成功")
	if s.authChan != nil {
		s.authChan <- true
	}
}

// OnRspUserLogin 登录响应
func (s *MyTraderSpi) OnRspUserLogin(pRspUserLogin *ctpgo.CThostFtdcRspUserLoginField, pRspInfo *ctpgo.CThostFtdcRspInfoField, nRequestID int32, bIsLast bool) {
	if pRspInfo != nil && pRspInfo.ErrorID != 0 {
		fmt.Printf("登录失败: %s\n", ctpgo.GB18030(pRspInfo.ErrorMsg[:]))
		return
	}
	if pRspUserLogin != nil {
		fmt.Printf("登录成功，交易日: %s\n", ctpgo.BytesToString(pRspUserLogin.TradingDay[:]))
	}
}

// OnRtnOrder 报单通知
func (s *MyTraderSpi) OnRtnOrder(pOrder *ctpgo.CThostFtdcOrderField) {
	if pOrder == nil {
		return
	}
	instrumentID := ctpgo.BytesToString(pOrder.InstrumentID[:])
	orderStatus := pOrder.OrderStatus
	fmt.Printf("报单状态更新: %s, 状态: %c\n", instrumentID, orderStatus)
}

func main() {
	// 加载 .env 文件
	if err := godotenv.Load(); err != nil {
		log.Println("未找到 .env 文件，将使用系统环境变量")
	}

	// 从环境变量读取配置
	userID := os.Getenv("CTP_USER_ID")
	password := os.Getenv("CTP_PASSWORD")
	brokerID := os.Getenv("CTP_BROKER_ID")
	mdFront := os.Getenv("CTP_MD_FRONT")
	traderFront := os.Getenv("CTP_TRADER_FRONT")
	appID := os.Getenv("CTP_APP_ID")
	authCode := os.Getenv("CTP_AUTH_CODE")

	// 检查必要的环境变量
	if userID == "" || password == "" || brokerID == "" {
		log.Fatal("请设置环境变量: CTP_USER_ID, CTP_PASSWORD, CTP_BROKER_ID")
	}
	if mdFront == "" {
		mdFront = "tcp://182.254.243.31:30011" // 默认 simnow 行情地址
	}
	if traderFront == "" {
		traderFront = "tcp://182.254.243.31:30001" // 默认 simnow 交易地址
	}

	// 示例：使用行情 API
	fmt.Println("=== 行情 API 示例 ===")

	// 注意：库会在首次创建 API 时自动加载
	// 可以通过设置 CTP_LIB_PATH 环境变量指定库路径，否则使用默认路径 "../libs" 或 "./libs"

	// 1. 创建行情 API 实例（会自动加载库）
	// flowPath 应该是目录路径，用于存储订阅信息文件
	// 注意：NewMdApi 会自动将相对路径转换为绝对路径，确保 .con 文件保存到正确位置
	flowPath := "./flow"
	// 确保目录存在
	if err := os.MkdirAll(flowPath, 0755); err != nil {
		log.Fatalf("创建 flow 目录失败: %v\n", err)
	}
	mdApi := ctpgo.NewMdApi(flowPath, false, false)
	if mdApi == nil {
		log.Fatal("创建行情 API 失败，可能是库加载失败。请确保库文件在 ../libs 或 ./libs 目录下，或设置 CTP_LIB_PATH 环境变量")
	}
	defer mdApi.Release()

	// 2. 创建并设置回调
	// 注意：SetSpi 方法内部会自动管理 userData
	// userData 是 API 实例的 ID，用于在回调中定位对应的 API 实例
	// 回调函数通过 userData 找到 API 实例，然后调用 api.spi 的方法
	mdSpi := &MyMdSpi{}
	mdApi.SetSpi(mdSpi)
	fmt.Println("mdApi.GetApiVersion()", mdApi.GetApiVersion())

	// 3. 注册前置地址
	mdApi.RegisterFront(mdFront)

	// 4. 初始化
	mdApi.Init()

	// 5. 等待连接
	time.Sleep(1 * time.Second)

	// 6. 登录
	loginReq := &ctpgo.CThostFtdcReqUserLoginField{}
	copy(loginReq.UserID[:], userID)
	copy(loginReq.Password[:], password)
	copy(loginReq.BrokerID[:], brokerID)

	requestID := int32(1)
	ret := mdApi.ReqUserLogin(loginReq, requestID)
	if ret != 0 {
		log.Printf("登录请求失败，错误码: %d\n", ret)
	}

	// 7. 等待登录响应
	time.Sleep(1 * time.Second)

	// 8. 订阅行情（需要替换为真实的合约代码）
	instruments := []string{"rb2605", "lc2605"}
	nCount := int32(len(instruments))
	ret = mdApi.SubscribeMarketData(instruments, nCount)
	if ret != 0 {
		log.Printf("订阅行情失败，错误码: %d\n", ret)
	}

	fmt.Println("等待行情数据...")
	time.Sleep(1 * time.Second)

	// 示例：使用交易 API
	fmt.Println("\n=== 交易 API 示例 ===")

	// 1. 创建交易 API 实例
	// 交易 API 也需要 flow 目录，可以使用不同的子目录或同一个目录
	traderFlowPath := "./flow_trader"
	// 确保目录存在
	if err := os.MkdirAll(traderFlowPath, 0755); err != nil {
		log.Fatalf("创建 trader flow 目录失败: %v\n", err)
	}
	// 注意：NewTraderApi 会自动将相对路径转换为绝对路径，确保 .con 文件保存到正确位置
	traderApi := ctpgo.NewTraderApi(traderFlowPath)
	if traderApi == nil {
		log.Fatal("创建交易 API 失败，可能是库加载失败。请确保库文件在 ../libs 或 ./libs 目录下，或设置 CTP_LIB_PATH 环境变量")
	}
	defer traderApi.Release()

	// 2. 创建并设置回调
	// 注意：SetSpi 方法内部会自动管理 userData
	// userData 是 API 实例的 ID，用于在回调中定位对应的 API 实例
	// 回调函数通过 userData 找到 API 实例，然后调用 api.spi 的方法
	authChan := make(chan bool, 1)
	traderSpi := &MyTraderSpi{
		authChan: authChan,
	}
	traderApi.SetSpi(traderSpi)
	fmt.Println("traderApi.GetApiVersion()", traderApi.GetApiVersion())
	fmt.Println("ctpgo.GetDataCollectApiVersion()", ctpgo.GetDataCollectApiVersion())

	// 3. 注册前置地址
	traderApi.RegisterFront(traderFront)

	// 4. 初始化
	traderApi.Init()

	// 5. 等待连接
	time.Sleep(2 * time.Second)

	// 6. 穿透式认证（如果配置了 AppID 和 AuthCode）
	if appID != "" && authCode != "" {
		fmt.Println("开始穿透式认证...")
		authReq := &ctpgo.CThostFtdcReqAuthenticateField{}
		copy(authReq.BrokerID[:], brokerID)
		copy(authReq.UserID[:], userID)
		copy(authReq.AppID[:], appID)
		copy(authReq.AuthCode[:], authCode)
		// UserProductInfo 可以留空或设置产品信息
		// copy(authReq.UserProductInfo[:], "your_product_info")

		authRequestID := int32(0)
		ret := traderApi.ReqAuthenticate(authReq, authRequestID)
		if ret != 0 {
			log.Printf("认证请求失败，错误码: %d\n", ret)
		} else {
			// 等待认证响应（设置超时）
			select {
			case authSuccess := <-authChan:
				if !authSuccess {
					log.Fatal("认证失败，无法继续登录")
				}
				fmt.Println("认证成功，准备登录...")
			case <-time.After(10 * time.Second):
				log.Fatal("认证超时")
			}
		}
	} else {
		fmt.Println("未配置 AppID 和 AuthCode，跳过穿透式认证（仅适用于仿真环境）")
	}

	// 7. 登录
	traderLoginReq := &ctpgo.CThostFtdcReqUserLoginField{}
	copy(traderLoginReq.UserID[:], userID)
	copy(traderLoginReq.Password[:], password)
	copy(traderLoginReq.BrokerID[:], brokerID)

	ret = traderApi.ReqUserLogin(traderLoginReq, requestID)
	if ret != 0 {
		log.Printf("登录请求失败，错误码: %d\n", ret)
	}

	fmt.Println("等待交易响应...")
	time.Sleep(10 * time.Second)

	fmt.Println("\n示例程序结束")
}
