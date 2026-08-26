@echo off
setlocal
cd /d "%~dp0\.."
title FACEIT Discord RPC - Firefox installer

echo.
echo ===================================================
echo   FACEIT Discord RPC  -  Firefox installer
echo ===================================================
echo Builds/runs the daemon and opens the add-on debugger.
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

set "EXT=%cd%\extensions\gecko"
echo [*] Opening Firefox add-on debugger...
start "" firefox.exe "about:debugging#/runtime/this-firefox"

echo.
echo ===================================================
echo In Firefox: click "Load Temporary Add-on" and choose:
echo   extensions\gecko\manifest.json
echo (Temporary add-ons reset after restart; re-load each
echo  session, or sign via AMO for a permanent install.)
echo.
echo See install\README.md for full EN/RU instructions.
echo ===================================================
echo.
pause
