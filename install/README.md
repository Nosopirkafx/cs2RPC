# FACEIT Discord RPC — Install Guide (EN / RU)

Shows your live FACEIT CS2 match in Discord Rich Presence.
No memory/process reading — 100% safe from FACEIT Anti-Cheat.

---

## EN

### 0. One-time: Discord Developer Portal
1. Go to https://discord.com/developers/applications → **New Application**.
2. Copy the **Application ID** (already set in `backend/main.go`).
3. **Rich Presence → Artwork Assets**: upload two images:
   - key `cs2` — the CS2 game icon (large image)
   - key `faciet` — the FACEIT logo (small image)
4. Keep the Discord desktop client running and logged in.

### 1. Build the daemon (only if `bin/faceit-rpc.exe` is missing)
Requires Go 1.22+. Or just run the installer — it builds automatically.
```bat
install\install_chromium.bat      (Chrome / Edge / Brave / Yandex)
install\install_firefox.bat       (Firefox / Zen)
```
The bat builds (if needed), starts `bin\faceit-rpc.exe`, and opens the
browser with the extension already loaded.

### 2. Chromium (Chrome / Edge / Brave / Yandex)
- Run `install\install_chromium.bat`.
- Chromium launches with the extension auto-loaded.
- Click the **puzzle icon** and **pin "FACEIT Discord RPC"**
  (pinning persists across restarts).
- Open a FACEIT match room → presence appears in Discord within ~1s.

### 3. Gecko (Firefox / Zen)
- Run `install\install_firefox.bat`.
- Firefox opens `about:debugging`.
- Click **Load Temporary Add-on** → select `extensions\gecko\manifest.json`.
- Temporary add-ons reset after restart; reload each session
  (or sign via AMO for a permanent install).

### 4. Use
- Click the toolbar icon → popup shows daemon status + live match (Map / ELO / Score).
- Enter your **FACEIT nickname** and press **Save** (needed for your ELO).
- Closing the match tab clears the presence.

---

## RU

### 0. Один раз: Discord Developer Portal
1. Зайди на https://discord.com/developers/applications → **New Application**.
2. Скопируй **Application ID** (уже прописан в `backend/main.go`).
3. **Rich Presence → Artwork Assets**: загрузи две картинки:
   - ключ `cs2` — иконка CS2 (большая)
   - ключ `faciet` — логотип FACEIT (маленькая)
4. Держи десктоп-клиент Discord запущенным и авторизованным.

### 1. Сборка демона (если нет `bin\faceit-rpc.exe`)
Нужен Go 1.22+. Или просто запусти установщик — он соберёт сам.
```bat
install\install_chromium.bat      (Chrome / Edge / Brave / Yandex)
install\install_firefox.bat       (Firefox / Zen)
```
Bat собирает (при необходимости), запускает `bin\faceit-rpc.exe` и
открывает браузер с уже загруженным расширением.

### 2. Chromium (Chrome / Edge / Brave / Yandex)
- Запусти `install\install_chromium.bat`.
- Chromium откроется с автоматически загруженным расширением.
- Нажми **иконку пазла** и **закрепи "FACEIT Discord RPC"**
  (закрепление сохраняется после перезапусков).
- Открой матч-руму FACEIT → статус появится в Discord через ~1с.

### 3. Gecko (Firefox / Zen)
- Запусти `install\install_firefox.bat`.
- Firefox откроет `about:debugging`.
- Нажми **Load Temporary Add-on** → выбери `extensions\gecko\manifest.json`.
- Временные дополнения сбрасываются после перезапуска — перезагружай
  каждую сессию (или подпиши через AMO для постоянной установки).

### 4. Использование
- Кликни иконку на панели → попап показывает статус демона и матч
  (карта / ELO / счёт) в реальном времени.
- Введи свой **FACEIT-никнейм** и нажми **Save** (нужно для твоего ELO).
- Закрытие вкладки матча очищает статус.
