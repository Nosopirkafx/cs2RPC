:: FACEIT Discord RPC - Chromium guide EN/RU
@echo off
chcp 65001 >nul
title FACEIT Discord RPC - Chromium guide
echo.
echo ==============================================================
echo   FACEIT Discord RPC  -  Chromium  -  EN / RU
echo ==============================================================
echo.
echo [EN] Steps:
echo   1. Build the daemon if bin\faceit-rpc.exe is missing:
echo        cd backend
echo        go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
echo   2. Run bin\faceit-rpc.exe and keep its window open.
echo   3. Open in your browser:  chrome://extensions
echo   4. Turn on Developer mode, switch at top-right.
echo   5. Click Load unpacked and select folder:  extensions\chromium
echo   6. Click the puzzle icon and pin FACEIT Discord RPC.
echo   7. Open a FACEIT match room, presence appears in Discord.
echo   8. In the popup, enter your FACEIT nickname and click Save.
echo.
 echo RU: Шаги:
 echo   1. Собери демон, если отсутствует bin\faceit-rpc.exe:
 echo        cd backend
 echo        go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
 echo   2. Запусти bin\faceit-rpc.exe и держи окно открытым.
 echo   3. Открой в браузере:  chrome://extensions
 echo   4. Включи Режим разработчика, переключатель справа сверху.
 echo   5. Нажми Load unpacked и выбери папку:  extensions\chromium
 echo   6. Нажми иконку пазла и закрепи FACEIT Discord RPC.
 echo   7. Открой матч-руму FACEIT, статус появится в Discord.
 echo   8. В попапе введи свой FACEIT-ник и нажми Save.
 echo.
 echo Discord Developer Portal: загрузи картинки с ключами  cs2  и  faciet
 echo Полное руководство: install\README.md
 echo.
 pause
