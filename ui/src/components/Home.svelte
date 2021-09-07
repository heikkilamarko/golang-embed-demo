<script>
  import { fly } from "svelte/transition";
  import store from "../stores/homeStore";
  import Title from "./Title.svelte";

  const { isLoading, message, error, loadMessage } = store;
</script>

<Title>Go Embed Demo</Title>

<section class="text-center">
  <button class="btn btn-primary rounded-pill px-4 my-4" on:click={loadMessage}
    >{$isLoading ? "Waiting..." : "Call API"}</button
  >

  {#if $message}
    <p class="fs-3 text-primary" in:fly={{ x: 100, duration: 600 }}>
      {$message}
    </p>
  {:else if $error}
    <p class="fs-3 text-danger" in:fly={{ x: 100, duration: 600 }}>
      {$error}
    </p>
  {/if}
</section>
