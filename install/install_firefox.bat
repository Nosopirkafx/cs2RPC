:: FACEIT Discord RPC - Firefox Zen guide xpi EN/RU
@echo off
chcp 65001 >nul
title FACEIT Discord RPC - Firefox Zen guide
echo.
echo ==============================================================
echo   FACEIT Discord RPC  -  Firefox / Zen  -  EN / RU
echo   Method: packaged xpi, permanent install
echo ==============================================================
echo.
echo [EN] Steps:
echo   1. Build the daemon if needed:
echo        cd backend
echo        go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
echo   2. Run bin\faceit-rpc.exe and keep its window open.
echo   3. Pack the extension: zip the CONTENTS of  extensions\gecko
echo      into a file  faceit-rpc.xpi  (manifest.json must be at the archive root).
echo   4. Open in browser:  about:config
echo      find  xpinstall.signatures.required  and set it to  false  (one time).
echo   5. Open  about:addons  - gear menu - Install Add-on From File - choose faceit-rpc.xpi.
echo   6. In the popup enter your FACEIT nickname and click Save.
echo   7. Open a FACEIT match room, presence appears in Discord.
echo.
 echo RU: Шаги:
 echo   1. Собери демон при необходимости:
 echo        cd backend
 echo        go build -ldflags="-s -w" -trimpath -o ..\bin\faceit-rpc.exe .
 echo   2. Запусти bin\faceit-rpc.exe, окно держи открытым.
 echo   3. Упакуй расширение: заархивируй содержимое папки extensions\gecko в файл faceit-rpc.xpi. manifest.json должен быть в корне архива.
 echo   4. Открой в браузере:  about:config.  Найди xpinstall.signatures.required и переключи в false, один раз.
 echo   5. Открой about:addons, меню-шестерёнка, Install Add-on From File, выбери faceit-rpc.xpi.
 echo   6. В попапе введи свой FACEIT-ник и нажми Save.
 echo   7. Открой матч-руму FACEIT, статус появится в Discord.
 echo.
 echo Discord Developer Portal: загрузи картинки с ключами  cs2  и  faciet
 echo Полное руководство: install\README.md
 echo.
 pause
