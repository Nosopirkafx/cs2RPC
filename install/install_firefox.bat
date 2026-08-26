@echo off
title FACEIT Discord RPC - Firefox setup guide
echo.
echo ===================================================
echo   FACEIT Discord RPC  -  Firefox setup guide
echo ===================================================
echo.
echo EN: open about:debugging, click "Load Temporary Add-on",
echo     select: extensions\gecko\manifest.json
echo     Enter your FACEIT nickname in the popup (Save).
echo     Full step-by-step (EN / RU) is in install\README.md.
echo.
echo RU: otkroj about:debugging, nazhmi "Load Temporary Add-on",
echo     vyberi: extensions\gecko\manifest.json
echo     Vvedi svoj FACEIT-nik v popup (Save).
echo     Podrobnyj guid (RU / EN) - v fajle install\README.md.
echo.
echo Opening the add-on debugger and the guide for you...
start "" firefox.exe "about:debugging#/runtime/this-firefox"
start "" "install\README.md"
echo.
pause
