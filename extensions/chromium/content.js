(() => {
  const PORT = 42157;

  let myNick = "";
  chrome.storage.local.get("nickname", (r) => { myNick = (r.nickname || "").trim(); });
  chrome.storage.onChanged.addListener((c) => {
    if (c.nickname) myNick = (c.nickname.newValue || "").trim();
  });

  let port = null;
  let lastState = null;

  function connect() {
    port = chrome.runtime.connect({ name: "statefeed" });
    port.onDisconnect.addListener(() => {
      port = null;
      setTimeout(connect, 1000);
    });
    if (lastState) port.postMessage({ type: "state", state: lastState });
  }
  connect();

  let last = "";
  let matchStart = null;

  function pick(selectors) {
    for (const sel of selectors) {
      const el = document.querySelector(sel);
      if (el) {
        const t = el.textContent.trim();
        if (t) return t;
      }
    }
    return null;
  }

  function findMap() {
    const re = /\b(de_\w+|mirage|inferno|nuke|overpass|ancient|anubis|vertigo|dust2|dust_2)\b/i;
    const els = document.querySelectorAll("*");
    for (const el of els) {
      if (el.children.length !== 0) continue;
      const m = el.textContent.trim().match(re);
      if (m) {
        const v = m[1].toLowerCase();
        if (v === "dust2" || v === "dust_2") return "de_dust2";
        return v;
      }
    }
    return null;
  }

  function findMyElo(nick) {
    if (!nick) return null;
    const lower = nick.toLowerCase();
    const all = document.querySelectorAll("*");

    const eloFrom = (root) => {
      const a = root.querySelector("[class*='fv-ps-mini-elo' i]");
      if (a) {
        const n = parseInt(a.textContent.replace(/[^\d]/g, ""), 10);
        if (!isNaN(n)) return n;
      }
      const b = root.querySelector("div[class*='gap-px' i][class*='font-bold' i]");
      if (b) {
        const n = parseInt(b.textContent.replace(/[^\d]/g, ""), 10);
        if (!isNaN(n) && n >= 300 && n <= 4000) return n;
      }
      const c = root.querySelector("[class*='Subtitle__Holder' i]");
      if (c) {
        const n = parseInt(c.textContent.replace(/[^\d]/g, ""), 10);
        if (!isNaN(n) && n >= 300 && n <= 4000) return n;
      }
      return null;
    };

    const nameNicks = [];
    const otherNicks = [];
    for (const el of all) {
      const t = el.textContent;
      if (!t || !t.toLowerCase().includes(lower) || el.children.length > 5) continue;
      const cls = typeof el.className === "string" ? el.className.toLowerCase() : "";
      if (cls.includes("nickname__name")) nameNicks.push(el);
      else otherNicks.push(el);
    }

    for (const list of [nameNicks, otherNicks]) {
      for (const el of list) {
        let p = el;
        for (let i = 0; i < 12 && p; i++) {
          const v = eloFrom(p);
          if (v !== null) return v;
          p = p.parentElement;
        }
      }
    }
    return null;
  }

  function findScore() {
    const els = document.querySelectorAll("[class*='FactionScore' i]");
    if (els.length >= 2) {
      const a = parseInt(els[0].textContent.trim(), 10);
      const b = parseInt(els[1].textContent.trim(), 10);
      if (!isNaN(a) && !isNaN(b)) return { team_a: a, team_b: b };
    }
    const containers = [
      document.querySelector("[data-testid*='scoreboard' i]"),
      document.querySelector("[class*='scoreboard' i]"),
      document.querySelector("[data-testid*='score' i]"),
      document.querySelector("[class*='scoreHeader' i]")
    ].filter(Boolean);
    const scope = containers[0] || document.body;
    const m = scope.innerText.match(/(\d{1,2})\s*[:\-]\s*(\d{1,2})/);
    if (m) {
      const a = parseInt(m[1], 10);
      const b = parseInt(m[2], 10);
      if (a <= 16 && b <= 16) return { team_a: a, team_b: b };
    }
    return null;
  }

  function extract() {
    const url = location.href;
    const inRoom = /faceit\.com\/.*\/room\//.test(url);

    let map = findMap();

    const elo = findMyElo(myNick);

    const score = findScore();

    let phase = null;
    const txt = document.body.innerText.toLowerCase();
    if (txt.includes("captain") || txt.includes("pick")) phase = "Captains Pick";
    else if (txt.includes("in queue") || txt.includes("searching") || txt.includes("queue")) phase = "In Queue";

    let status = "idle";
    if (inRoom) status = (score || phase) ? "match" : "queue";

    return {
      status,
      map: map || null,
      elo: elo,
      score: score,
      phase: phase || null,
      match_id: null,
      match_start: null
    };
  }

  function tick() {
    const s = extract();
    if (s.status === "match" && matchStart === null) matchStart = Math.floor(Date.now() / 1000);
    if (s.status !== "match") matchStart = null;
    s.match_start = matchStart;

    const key = JSON.stringify(s);
    if (key !== last) {
      last = key;
      lastState = s;
      console.log("[faceit-rpc] state", s);
      if (port) port.postMessage({ type: "state", state: s });
    }
  }

  setInterval(tick, 1000);
  tick();

  setInterval(() => {
    if (port && lastState) port.postMessage({ type: "state", state: lastState });
  }, 20000);

  window.addEventListener("beforeunload", () => {
    if (port) port.postMessage({ type: "state", state: { status: "idle" } });
  });
})();
