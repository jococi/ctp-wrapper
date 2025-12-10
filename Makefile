# CTP C API 编译 Makefile
# 支持 macOS, Linux, Windows

# 检测平台
UNAME_S := $(shell uname -s)
UNAME_M := $(shell uname -m)

# 目录定义
CSRC_DIR := csrc
CTPAPI_DIR := ctpapi
LIBS_DIR := libs
BUILD_DIR := build

# 源文件
MD_SRC := $(CSRC_DIR)/ctp_md_c_api.cpp
TRADER_SRC := $(CSRC_DIR)/ctp_trader_c_api.cpp

# 编译器标志
CXXFLAGS := -std=c++11 -fPIC -O3 -Wall -Wextra
LDFLAGS :=

# 根据平台设置
ifeq ($(UNAME_S),Darwin)
    # macOS
    PLATFORM := macos
    CXX := clang++
    MD_LIB := $(CTPAPI_DIR)/macos/libctpmd_c_api.dylib
    TRADER_LIB := $(CTPAPI_DIR)/macos/libctptrader_c_api.dylib
    
    # 头文件路径 (macOS上使用软链接,直接指向macos目录即可)
    # 注意: 代码中使用相对路径 ctpapi/macos/xxx.h, 所以需要包含项目根目录
    INCLUDE_MD := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_TRADER := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_DATACOLLECT := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_COMMON := -I. -I$(CTPAPI_DIR)/macos -I$(CSRC_DIR)
    
    # Framework 路径和链接
    FRAMEWORK_PATH := -F$(CTPAPI_DIR)/macos
    MD_FRAMEWORKS := -framework thostmduserapi_se
    TRADER_FRAMEWORKS := -framework thosttraderapi_se -framework MacDataCollect \
                         -framework IOKit -framework Foundation -framework CoreFoundation
    
    # 安装名称和rpath
    INSTALL_NAME_MD := -install_name @rpath/libctpmd_c_api.dylib
    INSTALL_NAME_TRADER := -install_name @rpath/libctptrader_c_api.dylib
    RPATH := -Wl,-rpath,@loader_path
    
    CXXFLAGS += $(FRAMEWORK_PATH)
    LDFLAGS += $(RPATH)
    
else ifeq ($(UNAME_S),Linux)
    # Linux
    PLATFORM := linux
    CXX := g++
    MD_LIB := $(CTPAPI_DIR)/linux/libctpmd_c_api.so
    TRADER_LIB := $(CTPAPI_DIR)/linux/libctptrader_c_api.so
    
    # 头文件路径
    # 注意: 代码中使用相对路径 ctpapi/linux/xxx.h, 所以需要包含项目根目录
    INCLUDE_MD := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_TRADER := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_DATACOLLECT := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_COMMON := -I. -I$(CTPAPI_DIR)/linux -I$(CSRC_DIR)
    
    # 库路径和链接
    LIB_PATH := -L$(CTPAPI_DIR)/linux
    MD_LIBS := -lthostmduserapi_se
    TRADER_LIBS := -lthosttraderapi_se -lLinuxDataCollect
    
    LDFLAGS += $(LIB_PATH) -Wl,-rpath,$$ORIGIN
    
else ifeq ($(OS),Windows_NT)
    # Windows (需要 MSVC)
    PLATFORM := windows
    CXX := cl
    MD_LIB := $(CTPAPI_DIR)/windows/ctpmd_c_api.dll
    TRADER_LIB := $(CTPAPI_DIR)/windows/ctptrader_c_api.dll
    
    # 头文件路径
    # 注意: 代码中使用相对路径 ctpapi/windows/xxx.h, 所以需要包含项目根目录
    INCLUDE_MD := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
    INCLUDE_TRADER := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
    INCLUDE_DATACOLLECT := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
    INCLUDE_COMMON := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
    
    # 库路径和链接
    LIB_PATH := /LIBPATH:$(CTPAPI_DIR)\\windows
    MD_LIBS := thostmduserapi_se.lib
    TRADER_LIBS := thosttraderapi_se.lib WinDataCollect.lib
    
    CXXFLAGS := /LD /O2 /EHsc
    LDFLAGS := $(LIB_PATH)
else
    $(error 不支持的操作系统: $(UNAME_S))
endif

# 默认目标
.PHONY: all clean md trader help check-frameworks

all: check-frameworks md trader
	@echo ""
	@echo "✓ 编译完成！"
	@echo "  行情API: $(MD_LIB)"
	@echo "  交易API: $(TRADER_LIB)"

# 检查 macOS frameworks
check-frameworks:
ifeq ($(PLATFORM),macos)
	@if [ ! -d "$(CTPAPI_DIR)/macos/thostmduserapi_se.framework" ]; then \
		echo "错误: 未找到 thostmduserapi_se.framework"; \
		exit 1; \
	fi
	@if [ ! -d "$(CTPAPI_DIR)/macos/thosttraderapi_se.framework" ]; then \
		echo "错误: 未找到 thosttraderapi_se.framework"; \
		exit 1; \
	fi
	@if [ ! -d "$(CTPAPI_DIR)/macos/MacDataCollect.framework" ]; then \
		echo "错误: 未找到 MacDataCollect.framework"; \
		exit 1; \
	fi
	@echo "✓ Framework 检查通过"
endif

# 编译行情API
md: $(MD_LIB)

$(MD_LIB): $(MD_SRC) $(CSRC_DIR)/ctp_md_c_api.h
	@echo "编译行情API: $@"
	@mkdir -p $(dir $(MD_LIB))
ifeq ($(PLATFORM),windows)
	$(CXX) $(CXXFLAGS) $(INCLUDE_MD) $(INCLUDE_COMMON) $(MD_SRC) \
		$(LDFLAGS) $(MD_LIBS) /OUT:$@
else
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_MD) $(INCLUDE_COMMON) \
		$(MD_SRC) -o $@ \
		$(INSTALL_NAME_MD) $(LDFLAGS) $(MD_FRAMEWORKS) $(MD_LIBS)
endif
	@echo "✓ $@ 编译完成"

# 编译交易API
trader: $(TRADER_LIB)

$(TRADER_LIB): $(TRADER_SRC) $(CSRC_DIR)/ctp_trader_c_api.h
	@echo "编译交易API: $@"
	@mkdir -p $(dir $(TRADER_LIB))
ifeq ($(PLATFORM),windows)
	$(CXX) $(CXXFLAGS) $(INCLUDE_TRADER) $(INCLUDE_DATACOLLECT) $(INCLUDE_COMMON) \
		$(TRADER_SRC) $(LDFLAGS) $(TRADER_LIBS) /OUT:$@
else
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_TRADER) $(INCLUDE_DATACOLLECT) $(INCLUDE_COMMON) \
		$(TRADER_SRC) -o $@ \
		$(INSTALL_NAME_TRADER) $(LDFLAGS) $(TRADER_FRAMEWORKS) $(TRADER_LIBS)
endif
	@echo "✓ $@ 编译完成"

# 清理
clean:
	@echo "清理编译产物..."
	@rm -f $(CTPAPI_DIR)/macos/libctpmd_c_api.* $(CTPAPI_DIR)/macos/libctptrader_c_api.*
	@rm -f $(CTPAPI_DIR)/linux/libctpmd_c_api.* $(CTPAPI_DIR)/linux/libctptrader_c_api.*
	@rm -f $(CTPAPI_DIR)/windows/ctpmd_c_api.* $(CTPAPI_DIR)/windows/ctptrader_c_api.*
	@rm -rf $(BUILD_DIR)
	@echo "✓ 清理完成"

# 帮助信息
help:
	@echo "CTP C API 编译系统"
	@echo ""
	@echo "用法:"
	@echo "  make          - 编译所有库 (行情API + 交易API)"
	@echo "  make md       - 仅编译行情API"
	@echo "  make trader   - 仅编译交易API"
	@echo "  make clean    - 清理编译产物"
	@echo "  make help     - 显示此帮助信息"
	@echo ""
	@echo "当前平台: $(PLATFORM)"
	@echo "编译器: $(CXX)"
	@echo ""
	@echo "输出目录: $(CTPAPI_DIR)/$(PLATFORM)"
	@echo "  行情API: $(MD_LIB)"
	@echo "  交易API: $(TRADER_LIB)"
