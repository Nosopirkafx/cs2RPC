# FACEIT Discord RPC — Install Guide (EN)

Shows your live FACEIT CS2 match in Discord Rich Presence.
No memory/process reading — 100% safe from FACEIT Anti-Cheat.

The Russian version of these steps is in the main `README.md`
("Installation → Russian" section).

---

## EN

### 0. One-time: Discord Developer Portal
1. Go to https://discord.com/developers/applications → **New Application**.
2. The **Application ID** is already set in `backend/main.go`
   (`1540354848015388685`).
3. **Rich Presence → Artwork Assets**: upload two images:
   - key `cs2` — the CS2 game icon (large image)
   - key `faciet` — the FACEIT logo (small image). Note: `faciet`, not `faceit`.
4. Keep the Discord desktop client running and logged in.

### 1. Get the daemon running
Either download the user bundle `faceit-rpc-win.zip` from GitHub Releases
(it contains `start_daemon.bat` + `faceit-rpc.xpi`), or build from source:

```bat
make build     # bin/faceit-rpc.exe
make xpi       # dist/faceit-rpc.xpi   (or install\pack_firefox.bat)
make dist      # full Windows bundle   (or install\pack_dist.bat)
```

Then run `start_daemon.bat` (or `bin\faceit-rpc.exe`) and keep its window open.

### 2. Chromium (Chrome / Edge / Brave / Yandex)
- Open `chrome://extensions` and enable **Developer mode**.
- Click **Load unpacked** → select the `extensions/chromium` folder
  (from the bundle, or from this repo).
- Click the **puzzle icon** and pin "FACEIT Discord RPC"
  (pinning persists across restarts).
- Open a FACEIT match room → presence appears in Discord within ~1s.

### 3. Gecko (Firefox / Zen) — packaged, permanent
- One time: open `about:config` and set `xpinstall.signatures.required` = `false`.
- Open `about:addons` → gear menu → **Install Add-on From File** →
  choose `faceit-rpc.xpi` (built via `make xpi`, or in the bundle).
- One restart may be required; after that it stays installed.

### 4. Use
- Click the toolbar icon → enter your **FACEIT nickname** → **Save**
  (needed so your own ELO is detected).
- Open a FACEIT match room → live presence (Map / ELO / Score) in Discord.
- Closing the match tab clears the presence.
