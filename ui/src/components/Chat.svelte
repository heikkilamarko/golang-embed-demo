<script>
  import store from "../stores/chatStore";
  import Title from "./Title.svelte";

  const { messages, sendMessage } = store;

  let sender = "";
  let message = "";

  function handleSubmit() {
    sendMessage({ sender, message });
    message = "";
  }

  $: canSend = sender && message;
</script>

<Title>Chat</Title>

<form class="row g-3 my-2 mb-5" on:submit|preventDefault={handleSubmit}>
  <div class="col-sm-4">
    <div class="input-group">
      <span class="input-group-text bg-white text-primary">@</span>
      <input
        type="text"
        class="form-control"
        placeholder="username..."
        bind:value={sender}
      />
    </div>
  </div>
  <div class="col-sm">
    <div class="input-group">
      <input
        type="text"
        class="form-control"
        placeholder="message..."
        bind:value={message}
      />
      <button class="btn btn-primary" type="submit" disabled={!canSend}
        >Send</button
      >
    </div>
  </div>
</form>

{#each $messages as m}
  <div class="border rounded bg-white py-2 px-3 mb-3 text-start">
    <div class="sender mb-2">
      <span class="fw-bold">{m.sender}</span>
      <span class="text-muted ms-3">{m.ts.toLocaleString()}</span>
    </div>
    <div>{m.message}</div>
  </div>
{/each}

<style>
  .sender {
    font-size: 0.75rem;
  }
</style>
