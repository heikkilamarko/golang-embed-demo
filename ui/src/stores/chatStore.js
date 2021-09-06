import { writable } from "svelte/store";

function createStore() {
  const messages = writable([]);

  let conn;

  function addMessage(message) {
    message.ts = new Date(message.ts);
    messages.update((messages) => [...messages, message]);
  }

  function sendMessage(message) {
    message.ts = new Date();
    conn?.send(JSON.stringify(message));
  }

  function connect() {
    conn = new WebSocket(`ws://${document.location.host}/ws`);
    conn.onmessage = (e) => addMessage(JSON.parse(e.data));
    conn.onclose = () =>
      addMessage({
        sender: "Server",
        message: "Connection closed.",
        ts: new Date(),
      });
  }

  return {
    messages,
    addMessage,
    sendMessage,
    connect,
  };
}

export default createStore();
