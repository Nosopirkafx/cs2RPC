:: Start FACEIT Discord RPC in this console window.
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
echo Starting FACEIT Discord RPC. Keep this window open while you play.
echo Press Ctrl+C to stop the application.
echo.
"%EXE%"
echo.
echo FACEIT Discord RPC has stopped.
pause
