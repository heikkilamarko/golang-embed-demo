<script>
  import { tick } from "svelte";
  import { fly } from "svelte/transition";
  import store from "../stores/chatStore";
  import Title from "./Title.svelte";

  const { sender, message, messages, canSendMessage, sendMessage } = store;

  let messagesEl;

  async function scrollToBottom() {
    await tick();
    messagesEl.scrollTop = messagesEl.scrollHeight;
  }

  $: $messages && scrollToBottom();
</script>

<Title>Chat</Title>

<form class="row g-3 my-2 mb-5" on:submit|preventDefault={sendMessage}>
  <div class="col-sm-4">
    <div class="input-group">
      <span class="input-group-text bg-white text-primary">@</span>
      <input
        type="text"
        class="form-control"
        placeholder="Type your name"
        bind:value={$sender}
      />
    </div>
  </div>
  <div class="col-sm">
    <div class="input-group">
      <input
        type="text"
        class="form-control"
        placeholder="Type a new message"
        bind:value={$message}
      />
      <button class="btn btn-primary" type="submit" disabled={!$canSendMessage}
        >Send</button
      >
    </div>
  </div>
</form>

<div class="flex-fill overflow-auto" bind:this={messagesEl}>
  {#each $messages as m (m.id)}
    <div
      class="border rounded bg-white py-2 px-3 mb-3 text-start"
      in:fly={{ x: 100, duration: 600 }}
    >
      <div class="sender mb-2">
        <span class="fw-bold">{m.sender}</span>
        <span class="text-muted ms-3">{m.ts.toLocaleString()}</span>
      </div>
      <div>{m.message}</div>
    </div>
  {/each}
</div>

<style>
  .sender {
    font-size: 0.75rem;
  }
</style>
