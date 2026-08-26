const PORT = 42157;

const nickInput = document.getElementById("nick");
const saveBtn = document.getElementById("save");
const savedHint = document.getElementById("saved");
const statusEl = document.getElementById("status");
const matchEl = document.getElementById("match");

function renderState(s) {
  if (!s || s.status === "idle" || (!s.map && !s.elo && !s.score && !s.phase)) {
    matchEl.innerHTML = '<div class="empty">No active match detected</div>';
    return;
  }
  const row = (k, v) => `<div class="stat"><span class="k">${k}</span><span class="v">${v}</span></div>`;
  let html = "";
  if (s.map) html += row("Map", s.map);
  if (s.elo) html += row("ELO", s.elo);
  if (s.score) html += row("Score", s.score.team_a + " : " + s.score.team_b);
  if (s.phase) html += row("Phase", s.phase);
  if (!html) html = '<div class="empty">Match found, waiting for data…</div>';
  matchEl.innerHTML = html;
}

chrome.storage.local.get(["nickname", "lastState"], (r) => {
  if (r.nickname) nickInput.value = r.nickname;
  renderState(r.lastState);
});

saveBtn.addEventListener("click", () => {
  const v = nickInput.value.trim();
  chrome.storage.local.set({ nickname: v }, () => {
    savedHint.textContent = "Saved: " + (v || "(empty)");
  });
});

chrome.storage.onChanged.addListener((c) => {
  if (c.lastState) renderState(c.lastState.newValue);
});

function checkDaemon() {
  fetch(`http://127.0.0.1:${PORT}/api/state`, { method: "GET" })
    .then((res) => {
      if (res.ok) {
        statusEl.textContent = "Connected";
        statusEl.className = "pill ok";
      } else {
        statusEl.textContent = "Error";
        statusEl.className = "pill bad";
      }
    })
    .catch(() => {
      statusEl.textContent = "Offline";
      statusEl.className = "pill bad";
    });
}

checkDaemon();
setInterval(checkDaemon, 5000);
