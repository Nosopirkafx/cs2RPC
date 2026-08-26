@echo off
chcp 65001 >nul
cd /d "%~dp0\.."
setlocal
echo [1/3] Building daemon (bin/faceit-rpc.exe)...
cd backend
go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
if errorlevel 1 (echo Build FAILED & exit /b 1)
cd ..
echo [2/3] Building .xpi...
call install\pack_firefox.bat
if errorlevel 1 (echo .xpi build FAILED & exit /b 1)
echo [3/3] Packing user bundle (dist/faceit-rpc-win.zip)...
powershell -NoProfile -ExecutionPolicy Bypass -Command "$s='dist/faceit-rpc-win'; if(Test-Path $s){Remove-Item $s -Recurse}; New-Item -ItemType Directory $s | Out-Null; Copy-Item 'bin/faceit-rpc.exe' $s; Copy-Item 'dist/faceit-rpc.xpi' $s; New-Item -ItemType Directory (Join-Path $s 'extensions') | Out-Null; Copy-Item 'extensions/chromium' (Join-Path $s 'extensions') -Recurse; Copy-Item 'install/install_chromium.bat' $s; Copy-Item 'install/install_firefox.bat' $s; Copy-Item 'start_daemon.bat' $s; Copy-Item 'README.md' $s; $z='dist/faceit-rpc-win.zip'; if(Test-Path $z){Remove-Item $z}; Compress-Archive -Path $s -DestinationPath $z -Force; Remove-Item $s -Recurse; Write-Host ('Bundle ready: ' + $z)"
echo.
echo All done. Distribute dist/faceit-rpc-win.zip to your users.
