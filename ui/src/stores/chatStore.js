import { derived, get, readable, writable } from "svelte/store";
import { v4 as uuid } from "uuid";

function createStore() {
  let conn;

  const sender = writable("");
  const message = writable("");
  const messages = writable([]);

  const connectionState = readable(null, (set) => {
    const id = setInterval(() => set(conn?.readyState ?? null), 1000);
    return () => clearInterval(id);
  });

  const isConnected = derived(
    connectionState,
    ($connectionState) => $connectionState === 1
  );

  const canSendMessage = derived(
    [isConnected, sender, message],
    ([$isConnected, $sender, $message]) => $isConnected && $sender && $message
  );

  function addMessage(message) {
    message.ts = new Date(message.ts);
    messages.update((messages) => [...messages, message]);
  }

  function addInfoMessage(message) {
    addMessage({
      sender: "INFO",
      message,
      ts: new Date(),
    });
  }

  function sendMessage() {
    conn?.send(
      JSON.stringify({
        id: uuid(),
        sender: get(sender),
        message: get(message),
        ts: new Date(),
      })
    );
    message.set("");
  }

  function connect() {
    if (conn && conn.readyState !== 3) return;
    conn = new WebSocket(`ws://${document.location.host}/ws`);
    conn.onopen = () => addInfoMessage("Connection opened.");
    conn.onclose = () => addInfoMessage("Connection closed.");
    conn.onerror = () => addInfoMessage("Connection error.");
    conn.onmessage = (e) => addMessage(JSON.parse(e.data));
  }

  connect();
  setInterval(connect, 2000);

  return {
    sender,
    message,
    messages,
    connectionState,
    isConnected,
    canSendMessage,
    sendMessage,
    connect,
  };
}

export default createStore();
