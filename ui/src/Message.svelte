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

<h1 class="display-6 fw-light text-success">{message}</h1>
