/**
 * CTP Md API - C 接口实现
 * 
 * 自动生成，请勿手动修改
 */

#include "ctp_md_c_api.h"

// 平台相关的原始 CTP 头文件
// 注意: 编译时需要正确设置头文件搜索路径
#ifdef _WIN32
    #include "ctpapi/windows/ThostFtdcMdApi.h"
    #include "ctpapi/windows/DataCollect.h"
#elif __APPLE__
    #include "ctpapi/macos/ThostFtdcMdApi.h"
    #include "ctpapi/macos/DataCollect.h"
#elif __linux__
    #include "ctpapi/linux/ThostFtdcMdApi.h"
    #include "ctpapi/linux/DataCollect.h"
#endif

#include <cstring>

// ========== SPI 包装类 ==========
class MdSpiWrapper : public CThostFtdcMdSpi {
public:
    virtual ~MdSpiWrapper() = default;

    void* userData = nullptr;

    MdOnFrontConnectedCallback onFrontConnected = nullptr;
    MdOnFrontDisconnectedCallback onFrontDisconnected = nullptr;
    MdOnHeartBeatWarningCallback onHeartBeatWarning = nullptr;
    MdOnRspUserLoginCallback onRspUserLogin = nullptr;
    MdOnRspUserLogoutCallback onRspUserLogout = nullptr;
    MdOnRspQryMulticastInstrumentCallback onRspQryMulticastInstrument = nullptr;
    MdOnRspErrorCallback onRspError = nullptr;
    MdOnRspSubMarketDataCallback onRspSubMarketData = nullptr;
    MdOnRspUnSubMarketDataCallback onRspUnSubMarketData = nullptr;
    MdOnRspSubForQuoteRspCallback onRspSubForQuoteRsp = nullptr;
    MdOnRspUnSubForQuoteRspCallback onRspUnSubForQuoteRsp = nullptr;
    MdOnRtnDepthMarketDataCallback onRtnDepthMarketData = nullptr;
    MdOnRtnForQuoteRspCallback onRtnForQuoteRsp = nullptr;

    void OnFrontConnected() override {
        if (onFrontConnected) {
            onFrontConnected(userData);
        }
    }

    void OnFrontDisconnected(int nReason) override {
        if (onFrontDisconnected) {
            onFrontDisconnected(userData, nReason);
        }
    }

    void OnHeartBeatWarning(int nTimeLapse) override {
        if (onHeartBeatWarning) {
            onHeartBeatWarning(userData, nTimeLapse);
        }
    }

    void OnRspUserLogin(CThostFtdcRspUserLoginField* pRspUserLogin, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserLogin) {
            onRspUserLogin(userData, pRspUserLogin, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUserLogout(CThostFtdcUserLogoutField* pUserLogout, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUserLogout) {
            onRspUserLogout(userData, pUserLogout, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspQryMulticastInstrument(CThostFtdcMulticastInstrumentField* pMulticastInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspQryMulticastInstrument) {
            onRspQryMulticastInstrument(userData, pMulticastInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspError(CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspError) {
            onRspError(userData, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspSubMarketData(CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspSubMarketData) {
            onRspSubMarketData(userData, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUnSubMarketData(CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUnSubMarketData) {
            onRspUnSubMarketData(userData, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspSubForQuoteRsp(CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspSubForQuoteRsp) {
            onRspSubForQuoteRsp(userData, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRspUnSubForQuoteRsp(CThostFtdcSpecificInstrumentField* pSpecificInstrument, CThostFtdcRspInfoField* pRspInfo, int nRequestID, bool bIsLast) override {
        if (onRspUnSubForQuoteRsp) {
            onRspUnSubForQuoteRsp(userData, pSpecificInstrument, pRspInfo, nRequestID, bIsLast);
        }
    }

    void OnRtnDepthMarketData(CThostFtdcDepthMarketDataField* pDepthMarketData) override {
        if (onRtnDepthMarketData) {
            onRtnDepthMarketData(userData, pDepthMarketData);
        }
    }

    void OnRtnForQuoteRsp(CThostFtdcForQuoteRspField* pForQuoteRsp) override {
        if (onRtnForQuoteRsp) {
            onRtnForQuoteRsp(userData, pForQuoteRsp);
        }
    }

};

// ========== C 接口实现 ==========

extern "C" {

MdApiHandle MdCreateFtdcMdApi(const char* pszFlowPath, const bool bIsUsingUdp, const bool bIsMulticast) {
    return reinterpret_cast<MdApiHandle>(
        CThostFtdcMdApi::CreateFtdcMdApi(pszFlowPath, bIsUsingUdp, bIsMulticast)
    );
}

const char * MdGetApiVersion(void) {
    return CThostFtdcMdApi::GetApiVersion();
}

void MdRelease(MdApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    api->Release();
}

void MdInit(MdApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    api->Init();
}

int MdJoin(MdApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->Join();
}

const char * MdGetTradingDay(MdApiHandle handle) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->GetTradingDay();
}

void MdRegisterFront(MdApiHandle handle, char* pszFrontAddress) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    api->RegisterFront(pszFrontAddress);
}

void MdRegisterNameServer(MdApiHandle handle, char* pszNsAddress) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    api->RegisterNameServer(pszNsAddress);
}

void MdRegisterFensUserInfo(MdApiHandle handle, struct CThostFtdcFensUserInfoField* pFensUserInfo) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    api->RegisterFensUserInfo(pFensUserInfo);
}

int MdSubscribeMarketData(MdApiHandle handle, char* ppInstrumentID[], int nCount) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->SubscribeMarketData(ppInstrumentID, nCount);
}

int MdUnSubscribeMarketData(MdApiHandle handle, char* ppInstrumentID[], int nCount) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->UnSubscribeMarketData(ppInstrumentID, nCount);
}

int MdSubscribeForQuoteRsp(MdApiHandle handle, char* ppInstrumentID[], int nCount) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->SubscribeForQuoteRsp(ppInstrumentID, nCount);
}

int MdUnSubscribeForQuoteRsp(MdApiHandle handle, char* ppInstrumentID[], int nCount) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->UnSubscribeForQuoteRsp(ppInstrumentID, nCount);
}

int MdReqUserLogin(MdApiHandle handle, struct CThostFtdcReqUserLoginField* pReqUserLoginField, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->ReqUserLogin(pReqUserLoginField, nRequestID);
}

int MdReqUserLogout(MdApiHandle handle, struct CThostFtdcUserLogoutField* pUserLogout, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->ReqUserLogout(pUserLogout, nRequestID);
}

int MdReqQryMulticastInstrument(MdApiHandle handle, struct CThostFtdcQryMulticastInstrumentField* pQryMulticastInstrument, int nRequestID) {
    auto* api = reinterpret_cast<CThostFtdcMdApi*>(handle);
    return api->ReqQryMulticastInstrument(pQryMulticastInstrument, nRequestID);
}

// SPI 创建与销毁
MdSpiHandle MdSpiCreate(void* userData) {
    auto* spi = new MdSpiWrapper();
    spi->userData = userData;
    return reinterpret_cast<MdSpiHandle>(spi);
}

void MdSpiDestroy(MdSpiHandle spi) {
    delete reinterpret_cast<MdSpiWrapper*>(spi);
}

void MdRegisterSpi(MdApiHandle api, MdSpiHandle spi) {
    auto* apiPtr = reinterpret_cast<CThostFtdcMdApi*>(api);
    auto* spiPtr = reinterpret_cast<MdSpiWrapper*>(spi);
    apiPtr->RegisterSpi(spiPtr);
}

void MdSpiSetCallbacks(MdSpiHandle spi, const MdSpiCallbacks* callbacks) {
    auto* spiPtr = reinterpret_cast<MdSpiWrapper*>(spi);
    spiPtr->userData = callbacks->userData;
    spiPtr->onFrontConnected = callbacks->onFrontConnected;
    spiPtr->onFrontDisconnected = callbacks->onFrontDisconnected;
    spiPtr->onHeartBeatWarning = callbacks->onHeartBeatWarning;
    spiPtr->onRspUserLogin = callbacks->onRspUserLogin;
    spiPtr->onRspUserLogout = callbacks->onRspUserLogout;
    spiPtr->onRspQryMulticastInstrument = callbacks->onRspQryMulticastInstrument;
    spiPtr->onRspError = callbacks->onRspError;
    spiPtr->onRspSubMarketData = callbacks->onRspSubMarketData;
    spiPtr->onRspUnSubMarketData = callbacks->onRspUnSubMarketData;
    spiPtr->onRspSubForQuoteRsp = callbacks->onRspSubForQuoteRsp;
    spiPtr->onRspUnSubForQuoteRsp = callbacks->onRspUnSubForQuoteRsp;
    spiPtr->onRtnDepthMarketData = callbacks->onRtnDepthMarketData;
    spiPtr->onRtnForQuoteRsp = callbacks->onRtnForQuoteRsp;
}

void MdSpiSetOnFrontConnected(MdSpiHandle spi, MdOnFrontConnectedCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onFrontConnected = callback;
}

void MdSpiSetOnFrontDisconnected(MdSpiHandle spi, MdOnFrontDisconnectedCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onFrontDisconnected = callback;
}

void MdSpiSetOnHeartBeatWarning(MdSpiHandle spi, MdOnHeartBeatWarningCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onHeartBeatWarning = callback;
}

void MdSpiSetOnRspUserLogin(MdSpiHandle spi, MdOnRspUserLoginCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspUserLogin = callback;
}

void MdSpiSetOnRspUserLogout(MdSpiHandle spi, MdOnRspUserLogoutCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspUserLogout = callback;
}

void MdSpiSetOnRspQryMulticastInstrument(MdSpiHandle spi, MdOnRspQryMulticastInstrumentCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspQryMulticastInstrument = callback;
}

void MdSpiSetOnRspError(MdSpiHandle spi, MdOnRspErrorCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspError = callback;
}

void MdSpiSetOnRspSubMarketData(MdSpiHandle spi, MdOnRspSubMarketDataCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspSubMarketData = callback;
}

void MdSpiSetOnRspUnSubMarketData(MdSpiHandle spi, MdOnRspUnSubMarketDataCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspUnSubMarketData = callback;
}

void MdSpiSetOnRspSubForQuoteRsp(MdSpiHandle spi, MdOnRspSubForQuoteRspCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspSubForQuoteRsp = callback;
}

void MdSpiSetOnRspUnSubForQuoteRsp(MdSpiHandle spi, MdOnRspUnSubForQuoteRspCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRspUnSubForQuoteRsp = callback;
}

void MdSpiSetOnRtnDepthMarketData(MdSpiHandle spi, MdOnRtnDepthMarketDataCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRtnDepthMarketData = callback;
}

void MdSpiSetOnRtnForQuoteRsp(MdSpiHandle spi, MdOnRtnForQuoteRspCallback callback) {
    reinterpret_cast<MdSpiWrapper*>(spi)->onRtnForQuoteRsp = callback;
}

} // extern "C"
