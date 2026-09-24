const PORT = 42157;

const nickInput = document.getElementById("nick");
const saveBtn = document.getElementById("save");
const savedHint = document.getElementById("saved");
const statusEl = document.getElementById("status");
const matchEl = document.getElementById("match");

function renderState(s) {
  if (!s || s.status === "idle") {
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

chrome.storage.local.get("nickname", (r) => {
  if (r.nickname) nickInput.value = r.nickname;
});
renderState(null);

saveBtn.addEventListener("click", () => {
  const v = nickInput.value.trim();
  chrome.storage.local.set({ nickname: v }, () => {
    savedHint.textContent = "Saved: " + (v || "(empty)");
  });
});

function checkDaemon() {
  fetch(`http://127.0.0.1:${PORT}/api/status`, { method: "GET", cache: "no-store" })
    .then(async (res) => {
      if (!res.ok) throw new Error("daemon unavailable");
      const data = await res.json();
      statusEl.textContent = "Connected";
      statusEl.className = "pill ok";
      renderState(data.connected ? data.state : null);
    })
    .catch(() => {
      statusEl.textContent = "Offline";
      statusEl.className = "pill bad";
      renderState(null);
    });
}

checkDaemon();
setInterval(checkDaemon, 5000);
