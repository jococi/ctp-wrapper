@echo off
REM CTP C API Build Script (Windows)
REM Uses MSVC compiler

chcp 65001 >nul 2>&1

REM Check for clean command
if "%1"=="clean" goto clean_target
if "%1"=="help" goto help_target
if not "%1"=="" (
    echo Unknown command: %1
    echo Use: build.bat [clean^|help]
    exit /b 1
)

REM Check if compiler is available
where cl >nul 2>&1
if %errorlevel% equ 0 goto compiler_found

echo MSVC compiler not found, searching for Visual Studio...
echo.

REM Search for vcvars64.bat
set "VCVARS="

if exist "C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files\Microsoft Visual Studio\2022\Enterprise\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files\Microsoft Visual Studio\2022\Professional\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files\Microsoft Visual Studio\2022\Professional\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files\Microsoft Visual Studio\2022\Community\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2019\Enterprise\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2019\Enterprise\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2019\Professional\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2019\Professional\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2019\Community\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2017\Enterprise\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2017\Enterprise\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2017\Professional\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2017\Professional\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)
if exist "C:\Program Files (x86)\Microsoft Visual Studio\2017\Community\VC\Auxiliary\Build\vcvars64.bat" (
    set "VCVARS=C:\Program Files (x86)\Microsoft Visual Studio\2017\Community\VC\Auxiliary\Build\vcvars64.bat"
    goto found_vs
)

echo ERROR: Visual Studio not found!
echo Please install Visual Studio 2017/2019/2022 or run from Developer Command Prompt.
exit /b 1

:found_vs
echo Found: %VCVARS%
echo Setting up MSVC environment...
call "%VCVARS%"
echo.

REM Verify compiler
where cl >nul 2>&1
if %errorlevel% neq 0 (
    echo ERROR: Failed to setup MSVC environment
    exit /b 1
)

:compiler_found
echo MSVC compiler ready.
echo.

REM Directory definitions
set CSRC_DIR=csrc
set CTPAPI_DIR=ctpapi
set LIBS_DIR=libs

REM Create output directory
if not exist "%LIBS_DIR%" mkdir "%LIBS_DIR%"

REM Copy official libraries
echo Copying Windows official libraries to libs...
if exist "%CTPAPI_DIR%\windows\thostmduserapi_se.dll" (
    copy /Y "%CTPAPI_DIR%\windows\thostmduserapi_se.dll" "%LIBS_DIR%\" >nul
    echo   [OK] thostmduserapi_se.dll
)
if exist "%CTPAPI_DIR%\windows\thosttraderapi_se.dll" (
    copy /Y "%CTPAPI_DIR%\windows\thosttraderapi_se.dll" "%LIBS_DIR%\" >nul
    echo   [OK] thosttraderapi_se.dll
)
if exist "%CTPAPI_DIR%\windows\WinDataCollect.dll" (
    copy /Y "%CTPAPI_DIR%\windows\WinDataCollect.dll" "%LIBS_DIR%\" >nul
    echo   [OK] WinDataCollect.dll
)

REM Build Market Data API
echo.
echo ========================================
echo Building Market Data API...
echo ========================================
cl /LD /O2 /EHsc /std:c++17 /W3 /utf-8 /wd4828 /DCTP_EXPORTS ^
   /I. /I"%CTPAPI_DIR%\windows" /I"%CSRC_DIR%" ^
   "%CSRC_DIR%\ctp_md_c_api.cpp" ^
   /link /LIBPATH:"%CTPAPI_DIR%\windows" ^
   thostmduserapi_se.lib ^
   /OUT:"%LIBS_DIR%\ctpmd_c_api.dll"
if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Failed to build Market Data API!
    exit /b 1
)
echo   [SUCCESS] %LIBS_DIR%\ctpmd_c_api.dll
REM Clean temporary files
if exist "ctp_md_c_api.exp" del /Q "ctp_md_c_api.exp" >nul 2>&1
if exist "ctp_md_c_api.lib" del /Q "ctp_md_c_api.lib" >nul 2>&1
if exist "ctp_md_c_api.obj" del /Q "ctp_md_c_api.obj" >nul 2>&1

REM Build Trader API
echo.
echo ========================================
echo Building Trader API...
echo ========================================
cl /LD /O2 /EHsc /std:c++17 /W3 /utf-8 /wd4828 /DCTP_EXPORTS ^
   /I. /I"%CTPAPI_DIR%\windows" /I"%CSRC_DIR%" ^
   "%CSRC_DIR%\ctp_trader_c_api.cpp" ^
   /link /LIBPATH:"%CTPAPI_DIR%\windows" ^
   thosttraderapi_se.lib WinDataCollect.lib ^
   /OUT:"%LIBS_DIR%\ctptrader_c_api.dll"
if %errorlevel% neq 0 (
    echo.
    echo [ERROR] Failed to build Trader API!
    exit /b 1
)
echo   [SUCCESS] %LIBS_DIR%\ctptrader_c_api.dll
REM Clean temporary files
if exist "ctp_trader_c_api.exp" del /Q "ctp_trader_c_api.exp" >nul 2>&1
if exist "ctp_trader_c_api.lib" del /Q "ctp_trader_c_api.lib" >nul 2>&1
if exist "ctp_trader_c_api.obj" del /Q "ctp_trader_c_api.obj" >nul 2>&1

echo.
echo ========================================
echo Build completed successfully!
echo All libraries output to: %LIBS_DIR%\
echo ========================================
exit /b 0

:clean_target
echo ========================================
echo Cleaning build artifacts...
echo ========================================
set LIBS_DIR=libs
if exist "%LIBS_DIR%\ctpmd_c_api.dll" (
    del /Q "%LIBS_DIR%\ctpmd_c_api.dll" >nul 2>&1
    echo   [DEL] ctpmd_c_api.dll
)
if exist "%LIBS_DIR%\ctpmd_c_api.lib" (
    del /Q "%LIBS_DIR%\ctpmd_c_api.lib" >nul 2>&1
    echo   [DEL] ctpmd_c_api.lib
)
if exist "%LIBS_DIR%\ctptrader_c_api.dll" (
    del /Q "%LIBS_DIR%\ctptrader_c_api.dll" >nul 2>&1
    echo   [DEL] ctptrader_c_api.dll
)
if exist "%LIBS_DIR%\ctptrader_c_api.lib" (
    del /Q "%LIBS_DIR%\ctptrader_c_api.lib" >nul 2>&1
    echo   [DEL] ctptrader_c_api.lib
)
if exist "%LIBS_DIR%\thostmduserapi_se.dll" (
    del /Q "%LIBS_DIR%\thostmduserapi_se.dll" >nul 2>&1
    echo   [DEL] thostmduserapi_se.dll
)
if exist "%LIBS_DIR%\thosttraderapi_se.dll" (
    del /Q "%LIBS_DIR%\thosttraderapi_se.dll" >nul 2>&1
    echo   [DEL] thosttraderapi_se.dll
)
if exist "%LIBS_DIR%\WinDataCollect.dll" (
    del /Q "%LIBS_DIR%\WinDataCollect.dll" >nul 2>&1
    echo   [DEL] WinDataCollect.dll
)
REM Clean temporary files (.exp, .lib, .obj)
if exist "ctp_md_c_api.exp" (
    del /Q "ctp_md_c_api.exp" >nul 2>&1
    echo   [DEL] ctp_md_c_api.exp
)
if exist "ctp_md_c_api.lib" (
    del /Q "ctp_md_c_api.lib" >nul 2>&1
    echo   [DEL] ctp_md_c_api.lib
)
if exist "ctp_md_c_api.obj" (
    del /Q "ctp_md_c_api.obj" >nul 2>&1
    echo   [DEL] ctp_md_c_api.obj
)
if exist "ctp_trader_c_api.exp" (
    del /Q "ctp_trader_c_api.exp" >nul 2>&1
    echo   [DEL] ctp_trader_c_api.exp
)
if exist "ctp_trader_c_api.lib" (
    del /Q "ctp_trader_c_api.lib" >nul 2>&1
    echo   [DEL] ctp_trader_c_api.lib
)
if exist "ctp_trader_c_api.obj" (
    del /Q "ctp_trader_c_api.obj" >nul 2>&1
    echo   [DEL] ctp_trader_c_api.obj
)
if exist "*.obj" (
    del /Q *.obj >nul 2>&1
    echo   [DEL] *.obj
)
echo.
echo [SUCCESS] Clean completed!
echo ========================================
exit /b 0

:help_target
echo ========================================
echo CTP C API Build Script (Windows)
echo ========================================
echo.
echo Usage:
echo   build.bat          - Build all libraries (Market Data API + Trader API)
echo   build.bat clean    - Clean build artifacts
echo   build.bat help     - Show this help message
echo.
echo Output directory: libs\
echo   Market Data API: libs\ctpmd_c_api.dll
echo   Trader API:      libs\ctptrader_c_api.dll
echo.
echo Note:
echo   - All libraries (including official CTP libraries) are output to libs\
echo   - Use relative paths for runtime dependencies
echo   - Executable and libs\ directory can be distributed together
echo.
exit /b 0
