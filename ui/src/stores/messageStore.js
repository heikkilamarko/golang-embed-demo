import { writable, get } from "svelte/store";

function createStore() {
  let isLoading = writable(false);
  let message = writable();
  let error = writable();

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
    } catch (error) {
      error.set(error.message);
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
