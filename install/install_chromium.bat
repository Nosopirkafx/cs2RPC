@echo off
setlocal
cd /d "%~dp0\.."
title FACEIT Discord RPC - Chromium installer

echo.
echo ===================================================
echo   FACEIT Discord RPC  -  Chromium installer
echo ===================================================
echo Builds/runs the daemon and loads the extension.
echo.

if exist "bin\faceit-rpc.exe" (
  echo [+] Daemon found: bin\faceit-rpc.exe
) else (
  where go >nul 2>nul
  if not errorlevel 1 (
    echo [*] Go found - building daemon...
    pushd backend
    go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
    popd
    if exist "bin\faceit-rpc.exe" (echo [+] Build OK) else (echo [!] Build failed)
  ) else (
    echo [!] Go not installed. Place a ready bin\faceit-rpc.exe or install Go.
  )
)

if exist "bin\faceit-rpc.exe" (
  echo [*] Starting daemon...
  start "" "bin\faceit-rpc.exe"
) else (
  echo [!] Daemon missing - presence will not work until you run it.
)

set "EXT=%cd%\extensions\chromium"
echo [*] Loading extension into Chromium...
start "" chrome.exe --load-extension="%EXT%"

echo.
echo ===================================================
echo The extension is now loaded. To keep it visible, click
echo the puzzle icon and pin "FACEIT Discord RPC".
echo (Pinning persists across restarts.)
echo.
echo See install\README.md for full EN/RU instructions.
echo ===================================================
echo.
pause
