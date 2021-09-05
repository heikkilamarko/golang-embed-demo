import { writable } from "svelte/store";

function createStore() {
  const messages = writable([]);

  let conn;

  function addMessage(message) {
    messages.update((messages) => [...messages, message]);
  }

  function sendMessage(message) {
    conn?.send(message);
  }

  function connect() {
    conn = new WebSocket("ws://" + document.location.host + "/ws");
    conn.onclose = () => addMessage("Connection closed.");
    conn.onmessage = (e) => e.data.split("\n").forEach(addMessage);
  }

  return {
    messages,
    addMessage,
    sendMessage,
    connect,
  };
}

export default createStore();
