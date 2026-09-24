<div align="center">

# FACEIT Discord RPC for CS2

**Показывает карту, ELO, счёт и время матча FACEIT в статусе Discord.**

[![License](https://img.shields.io/github/license/Nosopirkafx/cs2RPC)](LICENSE)
[![Latest release](https://img.shields.io/github/v/release/Nosopirkafx/cs2RPC)](https://github.com/Nosopirkafx/cs2RPC/releases)
[![Go](https://img.shields.io/github/go-mod/go-version/Nosopirkafx/cs2RPC?filename=backend/go.mod)](backend/go.mod)

[English](#english) · [Установка за пару минут](#быстрый-старт) · [Сборка из исходников](#для-разработчиков)

</div>

> **Для работы нужны Windows, браузер с расширением FACEIT RPC и запущенный Discord для Windows.** Установка проходит один раз. После неё для запуска достаточно открыть `start_daemon.bat`: приложение запустится в фоне и откроет страницу состояния.

## Быстрый старт

### 1. Скачайте и распакуйте приложение

Откройте [Releases](https://github.com/Nosopirkafx/cs2RPC/releases) и скачайте `faceit-rpc-win.zip`. Распакуйте архив в обычную папку, например на рабочий стол. Запустите **`start_daemon.bat`**. Откроется локальная страница приложения: её можно свернуть, пока играете.

### 2. Установите расширение в браузер

Выберите свой браузер и выполните шаги ниже. Это нужно сделать один раз.

<details>
<summary><b>Chrome, Edge, Brave, Yandex и другие Chromium-браузеры</b></summary>

1. Откройте `chrome://extensions` (в Edge — `edge://extensions`).
2. Включите **Режим разработчика**.
3. Нажмите **Загрузить распакованное расширение** и выберите папку `extensions/chromium` внутри распакованного архива.
4. Закрепите **FACEIT Discord RPC** на панели браузера.

</details>

<details>
<summary><b>Firefox и Zen</b></summary>

В архиве лежит неподписанное дополнение. Обычный Firefox не устанавливает такие дополнения постоянно. Чтобы попробовать его временно:

1. Откройте `about:debugging` → **This Firefox**.
2. Нажмите **Load Temporary Add-on…** и выберите `faceit-rpc.xpi` из распакованного архива.

Firefox удалит временное дополнение после перезапуска. Для постоянной установки разработчику нужно получить подпись Mozilla. Chromium устанавливается обычным способом и подходит для ежедневного использования.

</details>

### 3. Укажите свой FACEIT-ник

Нажмите иконку расширения, введите ник FACEIT и нажмите **Save**. Это нужно для определения вашего ELO.

### 4. Откройте матч

Откройте комнату матча FACEIT. Статус появится в Discord автоматически. Когда закончите, закройте окно daemon.

## Что отображается

- Карта, ELO, счёт команд и фаза матча — если эти данные видны на странице FACEIT.
- Таймер начинается, когда расширение распознаёт активный матч, поэтому он может не совпадать с официальным временем начала.
- Расширение читает публичную страницу FACEIT в браузере. Приложение не читает память игры и не подключается к процессу CS2.

## Если что-то не работает

| Симптом | Что сделать |
| --- | --- |
| Расширение показывает **Offline** | Запустите `start_daemon.bat`; в браузере должна открыться страница приложения с зелёным статусом. |
| Discord не показывает статус | Откройте Discord для Windows и убедитесь, что вы вошли в аккаунт. |
| ELO не отображается | Проверьте ник в настройках расширения и обновите страницу матча. |
| Карта или счёт неверные | FACEIT меняет разметку сайта; сбор данных может временно перестать работать до обновления проекта. |
| Firefox удалил расширение после перезапуска | В этой MVP-сборке Firefox может потребоваться установить XPI снова. |

Лог daemon сохраняется рядом с `faceit-rpc.exe` в файле `cs2rpc.log`.

## Для разработчиков

Требуется Go **1.25+**. В Windows можно собрать пользовательский архив командой `install\pack_dist.bat`; для разработки полезны команды:

```bash
make build   # собрать daemon в bin/faceit-rpc.exe
make xpi     # упаковать Firefox-расширение в dist/faceit-rpc.xpi
make dist    # собрать пользовательский архив dist/faceit-rpc-win.zip
make run     # запустить daemon из исходников
```

Для Chromium загрузите папку `extensions/chromium` через страницу расширений браузера. Для Firefox архив `.xpi` собирается из `extensions/gecko`.

## English

<details>
<summary><b>Show the English quick start</b></summary>

Download `faceit-rpc-win.zip` from [Releases](https://github.com/Nosopirkafx/cs2RPC/releases), extract it, and run `start_daemon.bat`. It starts the app in the background and opens its local status page.

- **Chrome / Edge / Brave / Yandex:** open `chrome://extensions` (or `edge://extensions`), enable Developer mode, choose **Load unpacked**, and select `extensions/chromium` from the extracted folder.
- **Firefox / Zen:** open `about:debugging` → **This Firefox** → **Load Temporary Add-on…**, then select `faceit-rpc.xpi`. Firefox removes temporary add-ons after restart; permanent installation requires a Mozilla-signed release.
- Click the extension icon, enter your FACEIT nickname, and click **Save**. Open a FACEIT match room; Discord Rich Presence updates automatically.

Discord for Windows must be running. The extension reads the public match page DOM; it does not read or modify the game process.

</details>

## License

See [LICENSE](LICENSE).
