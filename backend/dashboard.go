package main

import "net/http"

func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" || r.Method != http.MethodGet || !validHost(r) {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	_, _ = w.Write([]byte(dashboardHTML))
}

const dashboardHTML = `<!doctype html>
<html lang="ru"><head><meta charset="utf-8"><meta name="viewport" content="width=device-width,initial-scale=1">
<title>FACEIT Discord RPC</title><style>
:root{color-scheme:dark;--bg:#0c1018;--card:#151c29;--line:#273247;--muted:#9ba8bf;--orange:#ff5500;--green:#59db91}*{box-sizing:border-box}body{margin:0;font:16px/1.5 Inter,Segoe UI,system-ui,sans-serif;background:radial-gradient(circle at 10% 0,#26304a 0,transparent 36rem),var(--bg);color:#f5f7fb}main{width:min(720px,calc(100% - 32px));margin:0 auto;padding:64px 0}h1{font-size:clamp(32px,7vw,52px);line-height:1;margin:0 0 12px}.orange{color:var(--orange)}.lead{color:var(--muted);margin:0 0 34px}.card{background:color-mix(in srgb,var(--card) 94%,white);border:1px solid var(--line);border-radius:20px;padding:24px;margin:16px 0;box-shadow:0 20px 60px #0004}.status{display:flex;align-items:center;gap:10px;font-weight:700}.dot{width:12px;height:12px;border-radius:50%;background:var(--green);box-shadow:0 0 18px var(--green)}.muted{color:var(--muted)}ol{margin:12px 0 0;padding-left:24px}li+li{margin-top:8px}button{appearance:none;border:0;border-radius:10px;padding:11px 15px;background:#273247;color:#f5f7fb;font:inherit;cursor:pointer}button:hover{background:#34435e}.stop{margin-top:18px;background:#3b2430}.stop:hover{background:#5a2e3c}.match{display:grid;grid-template-columns:repeat(2,1fr);gap:10px;margin-top:16px}.item{background:#0e1420;padding:12px;border-radius:12px}.label{font-size:12px;color:var(--muted);text-transform:uppercase;letter-spacing:.08em}.value{font-weight:700;overflow-wrap:anywhere}@media(max-width:480px){main{padding:32px 0}.match{grid-template-columns:1fr}}
</style></head><body><main>
<h1><span class="orange">FACEIT</span> Discord RPC</h1><p class="lead">Приложение запущено. Оставьте эту страницу или просто сверните её во время игры.</p>
<section class="card"><div class="status"><span class="dot"></span><span>Daemon работает</span></div><p id="state" class="muted">Жду страницу матча FACEIT…</p><div id="match" class="match" hidden></div></section>
<section class="card"><strong>Осталось один раз настроить браузер</strong><ol><li>Установите расширение из папки <code>extensions/chromium</code> через страницу расширений браузера.</li><li>Нажмите его иконку, укажите FACEIT-ник и нажмите Save.</li><li>Откройте комнату FACEIT-матча. Discord обновится сам.</li></ol></section>
<section class="card"><strong>Готово играть?</strong><p class="muted">Не закрывайте daemon, пока нужен статус Discord. Остановка очистит соединение с расширением.</p><button class="stop" id="stop">Остановить приложение</button></section>
</main><script>
const state=document.querySelector('#state'),match=document.querySelector('#match');
function show(s){if(!s.running){state.textContent='Жду страницу матча FACEIT…';match.hidden=true;return}const x=s.state||{},score=x.score?x.score.team_a+' : '+x.score.team_b:'—';state.textContent=x.status==='match'?'Матч обнаружен':'Страница FACEIT подключена';const rows=[['Карта',x.map||'—'],['ELO',x.elo||'—'],['Счёт',score],['Фаза',x.phase||'—']];match.innerHTML=rows.map(([k,v])=>'<div class="item"><div class="label">'+k+'</div><div class="value">'+String(v)+'</div></div>').join('');match.hidden=false}
async function refresh(){try{const r=await fetch('/api/status',{cache:'no-store'});show(await r.json())}catch{state.textContent='Соединение с daemon потеряно. Запустите start_daemon.bat снова.';match.hidden=true}}
document.querySelector('#stop').onclick=async()=>{await fetch('/api/stop',{method:'POST'});state.textContent='Приложение остановлено. Эту страницу можно закрыть.';document.querySelector('#stop').disabled=true};refresh();setInterval(refresh,2000);
</script></body></html>`
