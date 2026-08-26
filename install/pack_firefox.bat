:: Build dist/faceit-rpc.xpi (Firefox / Zen packaged extension)
@echo off
chcp 65001 >nul
cd /d "%~dp0\.."
setlocal
set OUT=dist\faceit-rpc.xpi
set ZIP=dist\faceit-rpc.zip
if not exist dist mkdir dist
if exist "%OUT%" del "%OUT%"
powershell -NoProfile -ExecutionPolicy Bypass -Command "Compress-Archive -Path 'extensions\gecko\*' -DestinationPath '%ZIP%' -Force"
if exist "%ZIP%" (
  powershell -NoProfile -ExecutionPolicy Bypass -Command "Rename-Item -Path '%ZIP%' -NewName 'faceit-rpc.xpi'"
  if exist "%OUT%" (echo Built: %OUT%) else (echo FAILED to rename %OUT% & exit /b 1)
) else (echo FAILED to build %ZIP% & exit /b 1)
