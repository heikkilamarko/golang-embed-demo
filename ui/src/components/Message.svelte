<script>
  import store from "../stores/messageStore";

  const { isLoading, message, error, loadMessage } = store;
</script>

<button class="btn btn-primary rounded-pill px-4 my-4" on:click={loadMessage}
  >{$isLoading ? "Loading..." : "Load message"}</button
>

{#if $message}
  <p class="fs-3 text-primary">{$message}</p>
{:else if $error}
  <p class="fs-3 text-danger">{$error}</p>
{/if}
