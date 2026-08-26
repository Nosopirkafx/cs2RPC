:: FACEIT Discord RPC - Firefox / Zen install guide (EN/RU)
@echo off
chcp 65001 >nul
title FACEIT Discord RPC - Firefox / Zen
echo.
echo ===== FACEIT Discord RPC - Firefox / Zen =====
echo.
echo [EN]
echo 1. Run start_daemon.bat and keep its window open.
echo 2. Open:  about:config
echo    Set  xpinstall.signatures.required  =  false   (one time).
echo 3. Open:  about:addons  - gear menu - Install Add-on From File
echo    Choose file:  faceit-rpc.xpi
echo 4. Open a FACEIT match room. In the popup enter your
echo    FACEIT nickname and click Save. Presence shows in Discord.
echo.
echo [RU]
echo 1. Запусти start_daemon.bat и держи окно открытым.
echo 2. Открой:  about:config
echo    Переключи  xpinstall.signatures.required  =  false   (один раз).
echo 3. Открой:  about:addons - меню-шестерёнка - Install Add-on From File
echo    Выбери файл:  faceit-rpc.xpi
echo 4. Открой матч-руму FACEIT. В попапе введи свой
echo    FACEIT-ник и нажми Save. Статус появится в Discord.
echo.
echo Full guide: README.md
echo.
pause
