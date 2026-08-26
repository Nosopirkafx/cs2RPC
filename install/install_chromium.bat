@echo off
title FACEIT Discord RPC - Chromium setup guide
echo.
echo ===================================================
echo   FACEIT Discord RPC  -  Chromium setup guide
echo ===================================================
echo.
echo EN: open the extensions page, enable "Developer mode",
echo     click "Load unpacked" and select: extensions\chromium
echo     Enter your FACEIT nickname in the popup (Save).
echo     Full step-by-step (EN / RU) is in install\README.md.
echo.
echo RU: otkroj stranicu rasshirenij, vklyuchi "Rezhim razrabotchika",
echo     nazhmi "Load unpacked" i vyberi papku: extensions\chromium
echo     Vvedi svoj FACEIT-nik v popup (Save).
echo     Podrobnyj guid (RU / EN) - v fajle install\README.md.
echo.
echo Opening the extensions page and the guide for you...
start "" chrome.exe "chrome://extensions"
start "" "install\README.md"
echo.
pause
