<script>
  import { onMount } from "svelte";

  let message = "Loading...";

  async function fetchMessage() {
    try {
      const response = await fetch("/api/message");
      return await response.text();
    } catch (error) {
      return error.message;
    }
  }

  onMount(async () => {
    message = await fetchMessage();
  });
</script>

<div class="message">{message}</div>

<style>
  .message {
    padding: 2rem;
    color: green;
    font-size: 1.5rem;
  }
</style>
