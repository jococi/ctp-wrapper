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
MD_SRC := $(CSRC_DIR)/ctpmd_c_api.cpp
TRADER_SRC := $(CSRC_DIR)/ctptrader_c_api.cpp

# 编译器标志
CXXFLAGS := -std=c++17 -fPIC -O3 -Wall -Wextra -Wno-unused-parameter
LDFLAGS :=

# 根据平台设置
ifeq ($(UNAME_S),Darwin)
    # macOS
    PLATFORM := macos
    CXX := clang++
    # 输出到 libs 目录
    MD_LIB := $(LIBS_DIR)/libctpmd_c_api.dylib
    TRADER_LIB := $(LIBS_DIR)/libctptrader_c_api.dylib
    
    # 头文件路径 (macOS上使用软链接,直接指向macos目录即可)
    # 注意: 代码中使用相对路径 ctpapi/macos/xxx.h, 所以需要包含项目根目录
    INCLUDE_MD := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_TRADER := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_DATACOLLECT := -I. -I$(CTPAPI_DIR)/macos
    INCLUDE_COMMON := -I. -I$(CTPAPI_DIR)/macos -I$(CSRC_DIR)
    
    # Framework 路径和链接（编译时使用）
    FRAMEWORK_PATH := -F$(CTPAPI_DIR)/macos
    MD_FRAMEWORKS := -framework thostmduserapi_se
    TRADER_FRAMEWORKS := -framework thosttraderapi_se -framework MacDataCollect \
                         -framework IOKit -framework Foundation -framework CoreFoundation
    
    # 安装名称和rpath（运行时在可执行文件所在目录的 libs/ 子目录查找）
    # install_name 使用 @rpath，由 rpath 控制实际搜索位置
    INSTALL_NAME_MD := -install_name @rpath/libctpmd_c_api.dylib
    INSTALL_NAME_TRADER := -install_name @rpath/libctptrader_c_api.dylib
    # @loader_path 表示加载该库的可执行文件所在目录
    # 设置多个搜索路径：
    # 1. @loader_path/libs - 可执行文件所在目录的 libs/ 子目录（c_api 库和依赖 frameworks 都在这里）
    # 2. @loader_path/../libs - 可执行文件父目录的 libs/（备用）
    # 3. @loader_path - 可执行文件所在目录（备用）
    RPATH := -Wl,-rpath,@loader_path/libs -Wl,-rpath,@loader_path/../libs -Wl,-rpath,@loader_path
    
    CXXFLAGS += $(FRAMEWORK_PATH)
    LDFLAGS += $(RPATH)
    
else ifeq ($(UNAME_S),Linux)
    # Linux
    PLATFORM := linux
    CXX := g++
    # 输出到 libs 目录
    MD_LIB := $(LIBS_DIR)/libctpmd_c_api.so
    TRADER_LIB := $(LIBS_DIR)/libctptrader_c_api.so
    
    # 头文件路径
    # 注意: 代码中使用相对路径 ctpapi/linux/xxx.h, 所以需要包含项目根目录
    INCLUDE_MD := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_TRADER := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_DATACOLLECT := -I. -I$(CTPAPI_DIR)/linux
    INCLUDE_COMMON := -I. -I$(CTPAPI_DIR)/linux -I$(CSRC_DIR)
    
    # 库路径和链接（编译时使用官方库）
    LIB_PATH := -L$(CTPAPI_DIR)/linux
    MD_LIBS := -lthostmduserapi_se
    TRADER_LIBS := -lthosttraderapi_se -lLinuxDataCollect
    
    # SONAME 设置（类似 macOS 的 install_name）
    SONAME_MD := -Wl,-soname,libctpmd_c_api.so
    SONAME_TRADER := -Wl,-soname,libctptrader_c_api.so
    
    # $ORIGIN 表示加载该库的可执行文件所在目录，运行时会在那里查找 .so 文件
    # 设置多个搜索路径：
    # 1. $ORIGIN/libs - 可执行文件所在目录的 libs/ 子目录（c_api 库和依赖 .so 都在这里）
    # 2. $ORIGIN/../libs - 可执行文件父目录的 libs/（备用）
    # 3. $ORIGIN - 可执行文件所在目录（备用）
    LDFLAGS += $(LIB_PATH) -Wl,-rpath,$$ORIGIN/libs -Wl,-rpath,$$ORIGIN/../libs -Wl,-rpath,$$ORIGIN
    
# else ifeq ($(OS),Windows_NT)
#     # Windows (需要 MSVC)
#     PLATFORM := windows
#     CXX := cl
#     # 输出到 libs 目录
#     MD_LIB := $(LIBS_DIR)/ctpmd_c_api.dll
#     TRADER_LIB := $(LIBS_DIR)/ctptrader_c_api.dll
    
#     # 头文件路径
#     # 注意: 代码中使用相对路径 ctpapi/windows/xxx.h, 所以需要包含项目根目录
#     INCLUDE_MD := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
#     INCLUDE_TRADER := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
#     INCLUDE_DATACOLLECT := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
#     INCLUDE_COMMON := /I. /I$(CTPAPI_DIR)\\windows /I$(CSRC_DIR)
    
#     # 库路径和链接（编译时使用官方库）
#     LIB_PATH := /LIBPATH:$(CTPAPI_DIR)\\windows
#     MD_LIBS := thostmduserapi_se.lib
#     TRADER_LIBS := thosttraderapi_se.lib WinDataCollect.lib
    
#     CXXFLAGS := /LD /O2 /EHsc /std:c++17
#     LDFLAGS := $(LIB_PATH)
else
    $(error 不支持的操作系统: $(UNAME_S))
endif

# 默认目标
.PHONY: all clean md trader help check-frameworks install-deps

all: check-frameworks install-deps md trader
	@echo ""
	@echo "✓ 编译完成！"
	@echo "  行情API: $(MD_LIB)"
	@echo "  交易API: $(TRADER_LIB)"
	@echo "  所有库已输出到: $(LIBS_DIR)/"

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

# 安装依赖库到 libs 目录（用于独立分发）
install-deps:
	@echo "准备 libs 目录..."
	@mkdir -p $(LIBS_DIR)
ifeq ($(PLATFORM),linux)
	@echo "复制 Linux 官方库到 libs 目录..."
	@if [ -f "$(CTPAPI_DIR)/linux/thostmduserapi_se.so" ]; then \
		cp -f $(CTPAPI_DIR)/linux/thostmduserapi_se.so $(LIBS_DIR)/; \
		echo "  ✓ thostmduserapi_se.so"; \
	fi
	@if [ -f "$(CTPAPI_DIR)/linux/thosttraderapi_se.so" ]; then \
		cp -f $(CTPAPI_DIR)/linux/thosttraderapi_se.so $(LIBS_DIR)/; \
		echo "  ✓ thosttraderapi_se.so"; \
	fi
	@if [ -f "$(CTPAPI_DIR)/linux/LinuxDataCollect.so" ]; then \
		cp -f $(CTPAPI_DIR)/linux/LinuxDataCollect.so $(LIBS_DIR)/; \
		echo "  ✓ LinuxDataCollect.so"; \
	fi
else ifeq ($(PLATFORM),windows)
	@echo "复制 Windows 官方库到 libs 目录..."
	@if [ -f "$(CTPAPI_DIR)/windows/thostmduserapi_se.dll" ]; then \
		cp -f $(CTPAPI_DIR)/windows/thostmduserapi_se.dll $(LIBS_DIR)/; \
		echo "  ✓ thostmduserapi_se.dll"; \
	fi
	@if [ -f "$(CTPAPI_DIR)/windows/thosttraderapi_se.dll" ]; then \
		cp -f $(CTPAPI_DIR)/windows/thosttraderapi_se.dll $(LIBS_DIR)/; \
		echo "  ✓ thosttraderapi_se.dll"; \
	fi
	@if [ -f "$(CTPAPI_DIR)/windows/WinDataCollect.dll" ]; then \
		cp -f $(CTPAPI_DIR)/windows/WinDataCollect.dll $(LIBS_DIR)/; \
		echo "  ✓ WinDataCollect.dll"; \
	fi
else ifeq ($(PLATFORM),macos)
	@echo "复制 macOS frameworks 到 libs 目录..."
	@if [ -d "$(CTPAPI_DIR)/macos/thostmduserapi_se.framework" ]; then \
		if [ ! -d "$(LIBS_DIR)/thostmduserapi_se.framework" ]; then \
			cp -R $(CTPAPI_DIR)/macos/thostmduserapi_se.framework $(LIBS_DIR)/; \
			echo "  ✓ thostmduserapi_se.framework"; \
		else \
			echo "  ⊙ thostmduserapi_se.framework (已存在，跳过)"; \
		fi \
	fi
	@if [ -d "$(CTPAPI_DIR)/macos/thosttraderapi_se.framework" ]; then \
		if [ ! -d "$(LIBS_DIR)/thosttraderapi_se.framework" ]; then \
			cp -R $(CTPAPI_DIR)/macos/thosttraderapi_se.framework $(LIBS_DIR)/; \
			echo "  ✓ thosttraderapi_se.framework"; \
		else \
			echo "  ⊙ thosttraderapi_se.framework (已存在，跳过)"; \
		fi \
	fi
	@if [ -d "$(CTPAPI_DIR)/macos/MacDataCollect.framework" ]; then \
		if [ ! -d "$(LIBS_DIR)/MacDataCollect.framework" ]; then \
			cp -R $(CTPAPI_DIR)/macos/MacDataCollect.framework $(LIBS_DIR)/; \
			echo "  ✓ MacDataCollect.framework"; \
		else \
			echo "  ⊙ MacDataCollect.framework (已存在，跳过)"; \
		fi \
	fi
endif
	@echo "✓ 依赖库准备完成"

# 编译行情API
md: $(MD_LIB)

$(MD_LIB): $(MD_SRC) $(CSRC_DIR)/ctpmd_c_api.h
	@echo "编译行情API: $@"
	@mkdir -p $(dir $(MD_LIB))
ifeq ($(PLATFORM),windows)
	$(CXX) $(CXXFLAGS) $(INCLUDE_MD) $(INCLUDE_COMMON) $(MD_SRC) \
		$(LDFLAGS) $(MD_LIBS) /OUT:$@
else ifeq ($(PLATFORM),macos)
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_MD) $(INCLUDE_COMMON) \
		$(MD_SRC) -o $@ \
		$(INSTALL_NAME_MD) $(LDFLAGS) $(MD_FRAMEWORKS) $(MD_LIBS)
else
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_MD) $(INCLUDE_COMMON) \
		$(MD_SRC) -o $@ \
		$(SONAME_MD) $(LDFLAGS) $(MD_LIBS)
endif
	@echo "✓ $@ 编译完成"

# 编译交易API
trader: $(TRADER_LIB)

$(TRADER_LIB): $(TRADER_SRC) $(CSRC_DIR)/ctptrader_c_api.h
	@echo "编译交易API: $@"
	@mkdir -p $(dir $(TRADER_LIB))
ifeq ($(PLATFORM),windows)
	$(CXX) $(CXXFLAGS) $(INCLUDE_TRADER) $(INCLUDE_DATACOLLECT) $(INCLUDE_COMMON) \
		$(TRADER_SRC) $(LDFLAGS) $(TRADER_LIBS) /OUT:$@
else ifeq ($(PLATFORM),macos)
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_TRADER) $(INCLUDE_DATACOLLECT) $(INCLUDE_COMMON) \
		$(TRADER_SRC) -o $@ \
		$(INSTALL_NAME_TRADER) $(LDFLAGS) $(TRADER_FRAMEWORKS) $(TRADER_LIBS)
else
	$(CXX) -shared $(CXXFLAGS) $(INCLUDE_TRADER) $(INCLUDE_DATACOLLECT) $(INCLUDE_COMMON) \
		$(TRADER_SRC) -o $@ \
		$(SONAME_TRADER) $(LDFLAGS) $(TRADER_LIBS)
endif
	@echo "✓ $@ 编译完成"

# 清理
clean:
	@echo "清理编译产物..."
	@rm -f $(LIBS_DIR)/libctpmd_c_api.* $(LIBS_DIR)/libctptrader_c_api.*
	@rm -f $(LIBS_DIR)/ctpmd_c_api.* $(LIBS_DIR)/ctptrader_c_api.*
	@rm -f $(LIBS_DIR)/thostmduserapi_se.so $(LIBS_DIR)/thosttraderapi_se.so
	@rm -f $(LIBS_DIR)/LinuxDataCollect.so
	@rm -f $(LIBS_DIR)/thostmduserapi_se.dll $(LIBS_DIR)/thosttraderapi_se.dll
	@rm -f $(LIBS_DIR)/WinDataCollect.dll
	@rm -rf $(LIBS_DIR)/*.framework
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
	@echo "输出目录: $(LIBS_DIR)/"
	@echo "  行情API: $(MD_LIB)"
	@echo "  交易API: $(TRADER_LIB)"
	@echo ""
	@echo "说明:"
	@echo "  - 所有库（包括官方库）都会输出到 $(LIBS_DIR)/ 目录"
	@echo "  - 运行时使用相对路径（@loader_path/$ORIGIN）查找依赖"
	@echo "  - 可执行文件和 libs/ 目录放在一起即可独立分发"
