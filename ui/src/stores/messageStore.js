import { writable, get } from "svelte/store";

function createStore() {
  const isLoading = writable(false);
  const message = writable();
  const error = writable();

  async function loadMessage() {
    if (get(isLoading)) return;
    try {
      isLoading.set(true);
      message.set(null);
      error.set(null);
      const response = await fetch("/api/message");
      if (response.ok) {
        message.set(await response.text());
      } else {
        error.set(await response.text());
      }
    } catch (err) {
      error.set(err.message);
    } finally {
      isLoading.set(false);
    }
  }

  return {
    isLoading,
    message,
    error,
    loadMessage,
  };
}

export default createStore();
