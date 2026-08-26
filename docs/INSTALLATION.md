# Installation Guide

## 1. Discord Application & Assets

1. Open https://discord.com/developers/applications and create a **New Application**.
2. Copy its **Application ID** (numeric snowflake).
   - Already baked into `backend/main.go` (`clientID = "1540354848015388685"`).
3. Go to **Rich Presence → Artwork Assets** and upload:
   - `cs2` — the Counter-Strike 2 icon (used as the large image)
   - `faceit` — the FACEIT logo (used as the small image)
4. Keep Discord desktop client running and logged in.

## 2. Build the daemon

```bash
make build      # produces bin/faceit-rpc.exe
# optional: make pack   (UPX, ~3 MB)
```

Run `bin/faceit-rpc.exe`. A console window will NOT appear (GUI subsystem).
A log file `cs2rpc.log` is written next to the executable.

## 3. Install the browser extension

### Chromium (Chrome, Edge, Brave, Yandex, Opera)

1. Open `chrome://extensions` (or the equivalent page).
2. Enable **Developer mode** (top-right toggle).
3. Click **Load unpacked** and select the folder:
   `extensions/chromium`
4. The extension appears in your toolbar.

### Firefox / Zen (Gecko)

1. Open `about:debugging` → **This Firefox**.
2. Click **Load Temporary Add-on**.
3. Select `extensions/gecko/manifest.json`.
4. (Optional, for a permanent install) package and submit to AMO, or use a signed
   build — the `browser_specific_settings.gecko.id` is already set to
   `faceit-rpc@example.com`.

## 4. Use it

1. Make sure `faceit-rpc.exe` is running.
2. Click the toolbar icon → popup should say **"Daemon connected"** (otherwise launch the exe).
3. Open a FACEIT CS2 match room in a tab.
4. Your Discord profile now shows the live match: map, ELO, score and an elapsed timer.
5. Leave / close the match tab → presence clears automatically.

## Troubleshooting

- **No presence appears:** confirm Discord is running, the app ID is correct, and both
  `cs2` and `faceit` assets are uploaded. Check `cs2rpc.log`.
- **Popup says "Daemon not running":** launch `bin/faceit-rpc.exe`.
- **Score/map not updating:** FACEIT may have changed its DOM. Inspect the page and
  adjust the selectors in `content.js`.
- **Antivirus flags the exe:** this is a common false-positive for UPX-packed Go binaries.
  Build with `make build` (uncompressed) or add an exception.
