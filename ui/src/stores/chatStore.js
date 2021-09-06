import { derived, get, writable } from "svelte/store";

function createStore() {
  const sender = writable("");
  const message = writable("");
  const messages = writable([]);

  const canSendMessage = derived(
    [sender, message],
    ([$sender, $message]) => $sender && $message
  );

  let conn;

  function addMessage(message) {
    message.ts = new Date(message.ts);
    messages.update((messages) => [...messages, message]);
  }

  function sendMessage() {
    conn?.send(
      JSON.stringify({
        sender: get(sender),
        message: get(message),
        ts: new Date(),
      })
    );
    message.set("");
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
    sender,
    message,
    messages,
    canSendMessage,
    sendMessage,
    connect,
  };
}

export default createStore();
