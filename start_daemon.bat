:: Launch the FACEIT Discord RPC daemon (keep this window open)
@echo off
chcp 65001 >nul
cd /d "%~dp0"
if exist "%~dp0faceit-rpc.exe" (
  set "EXE=%~dp0faceit-rpc.exe"
) else if exist "%~dp0bin\faceit-rpc.exe" (
  set "EXE=%~dp0bin\faceit-rpc.exe"
) else (
  echo ERROR: faceit-rpc.exe not found.
  echo Re-download and extract faceit-rpc-win.zip from the release.
  pause
  exit /b 1
)
echo Starting FACEIT Discord RPC daemon...
echo Keep this window open while playing FACEIT.
echo.
"%EXE%"
