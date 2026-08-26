const PORT = 42157;

chrome.runtime.onConnect.addListener((port) => {
  if (port.name !== "statefeed") return;

  port.onMessage.addListener(async (msg) => {
    if (msg.type !== "state") return;
    chrome.storage.local.set({ lastState: msg.state });
    try {
      const res = await fetch(`http://127.0.0.1:${PORT}/api/state`, {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify(msg.state)
      });
      if (!res.ok) throw new Error("http " + res.status);
    } catch (e) {
      // Daemon not running yet — keep last state, resume on next change.
    }
  });
});
