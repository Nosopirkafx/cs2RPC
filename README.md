<div align="center">

# FACEIT Discord RPC for CS2

Show your FACEIT match in Discord Rich Presence: map, ELO, score and match phase.

[![License](https://img.shields.io/github/license/Nosopirkafx/cs2RPC)](LICENSE)
[![Latest release](https://img.shields.io/github/v/release/Nosopirkafx/cs2RPC)](https://github.com/Nosopirkafx/cs2RPC/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/Nosopirkafx/cs2RPC?filename=backend/go.mod)](backend/go.mod)

[English installation](#english-installation) · [Установка на русском](#установка-на-русском) · [Project files](#project-files)

</div>

![FACEIT Discord Rich Presence screenshot](docs/rpc-discord.png)

FACEIT RPC consists of a browser extension and a small Windows background service. The extension reads match information visible on the public FACEIT page and sends it to the service on your computer. The service updates your Discord status through Discord's local IPC connection.

> This project does not read game memory or connect to the CS2 process. You need Windows, Discord desktop, and a supported browser.

## English installation

### Download and start

1. Download `faceit-rpc-win.zip` from [Releases](https://github.com/Nosopirkafx/cs2RPC/releases).
2. Extract the ZIP to a regular folder. Do not run the app from inside the ZIP preview.
3. Double-click `start_daemon.bat`. A console window will stay open and show connection messages and match updates. Keep it open while playing; press **Ctrl+C** to stop the app.

### Install the browser extension

Install the extension once. Choose your browser below.

<details>
<summary><b>Chrome, Edge, Brave, Yandex and other Chromium browsers</b></summary>

1. Open `chrome://extensions` (in Edge, open `edge://extensions`).
2. Turn on **Developer mode**.
3. Click **Load unpacked**.
4. Select the `extensions/chromium` folder from the extracted ZIP. Select the folder itself, not a file inside it.
5. Pin **FACEIT Discord RPC** to the browser toolbar.
6. Click its toolbar icon, enter your FACEIT nickname, and click **Save**.

</details>

<details>
<summary><b>Firefox and Zen — permanent installation</b></summary>

The included `faceit-rpc.xpi` is unsigned. Standard Firefox Release blocks unsigned extensions, and changing `xpinstall.signatures.required` does not bypass that restriction there. Mozilla supports disabling the check in Firefox Developer Edition, Nightly, and ESR builds that expose the preference. See [Mozilla's signing guide](https://support.mozilla.org/en-US/kb/add-on-signing-in-firefox).

For a permanent install in a Firefox build that supports the preference:

1. Open `about:config` in the address bar and accept the warning.
2. Search for `xpinstall.signatures.required`.
3. If the preference exists and is editable, set it to **false**. If it is missing, locked, or the add-on is still rejected, that browser build does not allow unsigned permanent add-ons.
4. Open `about:addons`.
5. Click the gear menu and choose **Install Add-on From File…**.
6. Select `faceit-rpc.xpi` from the extracted ZIP and confirm the prompt.
7. Click the extension icon and save your FACEIT nickname.

Zen is based on Firefox, but whether this preference enables unsigned add-ons depends on the Zen build. If Zen rejects the XPI as unverified, there is no reliable `about:config` workaround for that build. A Mozilla-signed XPI is required for permanent installation in builds that enforce signatures. Do not download a pre-signed package from an unofficial source.

</details>

### Play

Make sure Discord desktop is running, keep the console window open, and open a FACEIT match room. The extension popup shows whether the daemon is connected and the latest detected match details. Closing the match tab clears the status; if the browser closes unexpectedly, the popup marks its state inactive after 30 seconds without an update.

## Установка на русском

### Скачать и запустить

1. Откройте [Releases](https://github.com/Nosopirkafx/cs2RPC/releases) и скачайте `faceit-rpc-win.zip`.
2. Распакуйте архив в обычную папку. Не запускайте программу прямо из окна просмотра ZIP.
3. Дважды нажмите `start_daemon.bat`. Откроется консоль с сообщениями о подключении и матче. Оставьте её открытой во время игры; для завершения нажмите **Ctrl+C**.

### Установить расширение

Установите расширение один раз. Выберите свой браузер.

<details>
<summary><b>Chrome, Edge, Brave, Яндекс Браузер и другие Chromium-браузеры</b></summary>

1. Откройте `chrome://extensions` (в Edge — `edge://extensions`).
2. Включите **Режим разработчика**.
3. Нажмите **Загрузить распакованное расширение**.
4. Выберите папку `extensions/chromium` внутри распакованного архива. Выберите именно папку, а не файл в ней.
5. Закрепите **FACEIT Discord RPC** на панели браузера.
6. Нажмите иконку расширения, введите FACEIT-ник и нажмите **Save**.

</details>

<details>
<summary><b>Firefox и Zen — постоянная установка</b></summary>

В архиве находится неподписанный файл `faceit-rpc.xpi`. Обычный Firefox Release блокирует неподписанные расширения. Настройка `xpinstall.signatures.required` не отключает эту защиту в обычной версии Firefox. Mozilla разрешает отключить проверку в Firefox Developer Edition, Nightly и в тех ESR-сборках, где эта настройка доступна. Подробности — в [справке Mozilla о подписи дополнений](https://support.mozilla.org/en-US/kb/add-on-signing-in-firefox).

Для постоянной установки в сборке Firefox, которая поддерживает эту настройку:

1. Введите `about:config` в адресной строке и подтвердите предупреждение.
2. Найдите настройку `xpinstall.signatures.required`.
3. Если настройка есть и доступна для изменения, установите значение **false**. Если настройки нет, она заблокирована или браузер всё равно отклоняет дополнение, эта сборка не разрешает постоянную установку неподписанных расширений.
4. Откройте `about:addons`.
5. Нажмите кнопку с шестерёнкой и выберите **Установить дополнение из файла…**.
6. Выберите `faceit-rpc.xpi` из распакованного архива и подтвердите установку.
7. Нажмите иконку расширения, введите FACEIT-ник и нажмите **Save**.

Zen основан на Firefox, но возможность установки неподписанных расширений зависит от конкретной сборки Zen. Если браузер сообщает, что XPI не проверен, надёжного обхода через `about:config` для этой сборки нет. Для постоянной установки в сборках, которые требуют подпись, нужен XPI, подписанный Mozilla. Не скачивайте подписанные файлы из неизвестных источников.

</details>

### Играть

Запустите Discord, не закрывайте консоль FACEIT RPC и откройте комнату матча FACEIT. В popup расширения отображаются подключение daemon и последние данные матча. Закрытие вкладки матча очищает статус; если браузер завершился аварийно, данные станут неактивными после 30 секунд без обновлений.

## Match data and troubleshooting

The extension detects match data from FACEIT's page markup. If FACEIT changes its site, a field can temporarily be missing or incorrect until the selectors are updated. The match timer starts when the extension first recognizes the match, so it can differ from FACEIT's official start time.

| Problem | What to check |
| --- | --- |
| Popup says **Offline** | Start `start_daemon.bat` and keep its console open. |
| Popup says **Connected**, but no match appears | Open a FACEIT room and refresh that tab. Check that the extension is enabled on FACEIT. |
| ELO is missing | Re-enter your FACEIT nickname in the extension popup, save it, and refresh the match tab. |
| Discord presence is missing | Run Discord desktop and check the console for Discord connection errors. |
| Firefox/Zen rejects the XPI | Check the browser-specific signature notes above. A signature-enforcing build needs an XPI signed by Mozilla. |

The daemon writes `cs2rpc.log` beside `faceit-rpc.exe` and also prints messages to its console.

## Project files

```text
.
├── .gitattributes                  GitHub language labels
├── .gitignore                      Ignores generated builds and local logs
├── LICENSE                         Project license
├── Makefile                        Build shortcuts for developers
├── README.md                       User guide in English and Russian
├── start_daemon.bat                Starts the daemon in a visible console
├── backend/
│   ├── go.mod                      Go module and dependency versions
│   ├── go.sum                      Checksums for Go dependencies
│   ├── main.go                     Local HTTP server, state API and logging
│   ├── console.go                  Centered, colored live match dashboard
│   ├── console_windows.go          Windows console color and width support
│   ├── console_other.go            Console fallback for non-Windows builds
│   ├── main_test.go                Tests for API state and origin validation
│   ├── console_test.go              Tests for logo and match data in the CLI
│   ├── mutex_windows.go            Prevents two daemon instances on Windows
│   ├── mutex_other.go              No-op mutex implementation for other systems
│   └── rpc/
│       ├── client.go               Discord IPC connection and Rich Presence updates
│       └── payload.go              Match state format and input validation
├── docs/
│   └── rpc-discord.png             README screenshot
├── extensions/
│   ├── chromium/
│   │   ├── manifest.json           Chromium extension metadata and permissions
│   │   ├── background.js            Sends match states to the local daemon
│   │   ├── content.js              Reads match details from the FACEIT page
│   │   ├── popup.html               Toolbar popup layout and styles
│   │   └── popup.js                 Nickname, daemon status and match display
│   └── gecko/
│       ├── manifest.json           Firefox/Zen extension metadata and ID
│       ├── background.js            Sends match states to the local daemon
│       ├── content.js              Reads match details from the FACEIT page
│       ├── popup.html               Toolbar popup layout and styles
│       └── popup.js                 Nickname, daemon status and match display
└── install/
    ├── pack_dist.bat                Builds the Windows user ZIP
    └── pack_firefox.bat             Packages the Gecko extension as an XPI
```

`bin/` and `dist/` are generated when building; they are not source folders and are not committed.

## Build from source

Developers need Go **1.25+**. On Windows, run `install\pack_dist.bat` to create `dist\faceit-rpc-win.zip`. With Make installed, the shortcuts are:

```bash
make build   # build bin/faceit-rpc.exe
make xpi     # build dist/faceit-rpc.xpi
make dist    # build the complete Windows bundle
make run     # run the daemon from source
```

The daemon listens on `127.0.0.1:42157`. Set `CS2RPC_PORT` to use a different port; extension configuration must then be updated to match.

## License

See [LICENSE](LICENSE).
