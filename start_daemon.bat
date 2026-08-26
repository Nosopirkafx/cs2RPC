:: Launch the FACEIT Discord RPC daemon (keep this window open)
@echo off
chcp 65001 >nul
cd /d "%~dp0"
if not exist bin\faceit-rpc.exe (
  echo ERROR: bin\faceit-rpc.exe not found.
  echo Run install\pack_dist.bat first to build the application.
  pause
  exit /b 1
)
echo Starting FACEIT Discord RPC daemon...
echo Keep this window open while playing FACEIT.
echo.
bin\faceit-rpc.exe
