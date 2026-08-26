# FACEIT Discord Rich Presence (CS2)

Shows your live FACEIT CS2 match (map, ELO, score, elapsed timer) in your Discord
Rich Presence — fully out-of-process. No memory reading, no process handles, no OCR.
100% isolated from FACEIT Anti-Cheat.

## Architecture

```
FACEIT tab (DOM) ──content.js──▶ background (service worker) ──HTTP POST──▶ Go daemon ──IPC──▶ Discord
                                (localhost only, 127.0.0.1:42157)
```

- **Browser extension** reads the public FACEIT match page DOM and forwards a JSON
  snapshot to `http://127.0.0.1:42157/api/state`.
- **Go daemon** (`faceit-rpc.exe`) validates the payload, keeps a single Discord IPC
  connection with auto-reconnect, and renders the Rich Presence.

## Discord Developer Portal setup

1. Go to https://discord.com/developers/applications and **New Application**.
2. Copy the **Application ID** (already set in `backend/main.go` as `1540354848015388685`).
3. **Rich Presence → Artwork Assets**: upload two images:
   - key `cs2` — the CS2 game icon (Large image)
   - key `faceit` — the FACEIT logo (Small image)
4. No public listing / install required.

## Build

Requires Go 1.22+ and (optionally) UPX.

```bash
make build     # -> bin/faceit-rpc.exe (~6 MB, console window for easy Ctrl+C close)
make pack      # -> UPX-compressed (~3 MB, no console window)
make run       # run from source for development
```

The port can be overridden with `CS2RPC_PORT` (must match `PORT` in the extension
files if changed).

## Install the extension

Quickest path — run the installer (builds the daemon, starts it, and loads the extension):

```bat
install\install_chromium.bat      # Chrome / Edge / Brave / Yandex
install\install_firefox.bat       # Firefox / Zen
```

Full step-by-step (EN/RU) for both browsers: [install/README.md](install/README.md)
and [docs/INSTALLATION.md](docs/INSTALLATION.md).

The extension popup shows the daemon status and a **live match panel** (Map / ELO / Score).
Enter your FACEIT nickname in the popup (Save) so your own ELO is detected.

## Usage

1. Launch `bin/faceit-rpc.exe` (runs in the background, logs to `cs2rpc.log` next to it).
2. Install the extension for your browser (Chromium or Gecko).
3. Open a FACEIT match room. Discord will show the live presence within ~1s.
4. Closing the tab or leaving the match clears the presence.

## Payload (extension → daemon)

```json
{
  "status": "match" | "queue" | "idle",
  "map": "Mirage" | null,
  "elo": 1234 | null,
  "score": { "team_a": 4, "team_b": 2 } | null,
  "phase": "In Queue" | "Captains Pick" | null,
  "match_start": 1700000000 | null,
  "match_id": "abc-123" | null
}
```

## Notes / caveats

- FACEIT DOM selectors are best-effort (`[data-testid]` anchors + heuristics). If the
  score/map stops updating after a site redesign, adjust the selectors in `content.js`.
- The daemon binds **only** `127.0.0.1` and rejects requests with a non-loopback `Host`
  or a non-extension `Origin` (DNS-rebinding / cross-site protection).
- Single-instance: double-launching the exe is ignored via a Windows named mutex.
