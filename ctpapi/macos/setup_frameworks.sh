#!/bin/bash

# 设置 framework：修复符号链接 + 创建头文件软链接
# 步骤 1: 修复 framework 符号链接
# 步骤 2: 创建头文件软链接到 macos 路径下

set -e

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
cd "$SCRIPT_DIR"

echo "设置 framework..."
echo ""

# 步骤 1: 修复 framework 符号链接
echo "步骤 1: 修复 framework 符号链接..."

# 修复 thosttraderapi_se.framework
if [ -d "thosttraderapi_se.framework" ]; then
    cd thosttraderapi_se.framework
    rm -f thosttraderapi_se
    ln -sf Versions/A/thosttraderapi_se thosttraderapi_se
    # 修复 Headers 和 Resources
    if [ ! -L Headers ] && [ -f Headers ]; then
        rm -f Headers
        ln -sf Versions/A/Headers Headers
    fi
    if [ ! -L Resources ] && [ -f Resources ]; then
        rm -f Resources
        ln -sf Versions/A/Resources Resources
    fi
    # 修复 Versions/Current
    if [ -d "Versions" ]; then
        cd Versions
        if [ ! -L Current ]; then
            rm -f Current
            ln -sf A Current
        fi
        cd ..
    fi
    cd ..
    echo "✓ 修复 thosttraderapi_se.framework"
fi

# 修复 thostmduserapi_se.framework
if [ -d "thostmduserapi_se.framework" ]; then
    cd thostmduserapi_se.framework
    rm -f thostmduserapi_se
    ln -sf Versions/A/thostmduserapi_se thostmduserapi_se
    # 修复 Headers 和 Resources
    if [ ! -L Headers ] && [ -f Headers ]; then
        rm -f Headers
        ln -sf Versions/A/Headers Headers
    fi
    if [ ! -L Resources ] && [ -f Resources ]; then
        rm -f Resources
        ln -sf Versions/A/Resources Resources
    fi
    # 修复 Versions/Current
    if [ -d "Versions" ]; then
        cd Versions
        if [ ! -L Current ]; then
            rm -f Current
            ln -sf A Current
        fi
        cd ..
    fi
    cd ..
    echo "✓ 修复 thostmduserapi_se.framework"
fi

# 修复 MacDataCollect.framework
if [ -d "MacDataCollect.framework" ]; then
    cd MacDataCollect.framework
    rm -f MacDataCollect
    ln -sf Versions/A/MacDataCollect MacDataCollect
    # 修复 Headers 和 Resources
    if [ ! -L Headers ] && [ -f Headers ]; then
        rm -f Headers
        ln -sf Versions/A/Headers Headers
    fi
    if [ ! -L Resources ] && [ -f Resources ]; then
        rm -f Resources
        ln -sf Versions/A/Resources Resources
    fi
    # 修复 Versions/Current
    if [ -d "Versions" ]; then
        cd Versions
        if [ ! -L Current ]; then
            rm -f Current
            ln -sf A Current
        fi
        cd ..
    fi
    cd ..
    echo "✓ 修复 MacDataCollect.framework"
fi

echo ""

# 步骤 2: 创建头文件软链接到 macos 路径下
echo "步骤 2: 创建头文件软链接..."

# thosttraderapi_se.framework 的头文件
ln -sf thosttraderapi_se.framework/Versions/A/Headers/ThostFtdcTraderApi.h ThostFtdcTraderApi.h
ln -sf thosttraderapi_se.framework/Versions/A/Headers/ThostFtdcUserApiDataType.h ThostFtdcUserApiDataType.h
ln -sf thosttraderapi_se.framework/Versions/A/Headers/ThostFtdcUserApiStruct.h ThostFtdcUserApiStruct.h

# thostmduserapi_se.framework 的头文件
ln -sf thostmduserapi_se.framework/Versions/A/Headers/ThostFtdcMdApi.h ThostFtdcMdApi.h
# 注意：UserApiDataType 和 UserApiStruct 可能已经在上面创建了，这里检查一下
if [ ! -L ThostFtdcUserApiDataType.h ]; then
    ln -sf thostmduserapi_se.framework/Versions/A/Headers/ThostFtdcUserApiDataType.h ThostFtdcUserApiDataType.h
fi
if [ ! -L ThostFtdcUserApiStruct.h ]; then
    ln -sf thostmduserapi_se.framework/Versions/A/Headers/ThostFtdcUserApiStruct.h ThostFtdcUserApiStruct.h
fi

# MacDataCollect.framework 的头文件
ln -sf MacDataCollect.framework/Versions/A/Headers/DataCollect.h DataCollect.h

echo "✓ 头文件软链接创建完成"
echo ""

echo "设置完成！"
echo ""
echo "创建的头文件软链接："
ls -lh *.h 2>/dev/null | grep -E "(ThostFtdc|DataCollect)" || echo "未找到头文件"

