const PORT = 42157;

chrome.runtime.onMessage.addListener((msg) => {
  if (msg && msg.type === "state") {
    fetch(`http://127.0.0.1:${PORT}/api/state`, {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify(msg.state)
    }).catch(() => {
      // Daemon not running yet — keep last state, resume on next message.
    });
  }
  return false;
});
