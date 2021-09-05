<script>
  import store from "../stores/chatStore";
  import Title from "./Title.svelte";

  const { messages, sendMessage } = store;

  let message = "";

  function handleSubmit() {
    sendMessage(message);
    message = "";
  }
</script>

<Title>Chat</Title>

<form class="mb-5" on:submit|preventDefault={handleSubmit}>
  <div class="input-group">
    <input type="text" class="form-control" bind:value={message} />
    <button class="btn btn-primary" type="submit" disabled={!message}
      >Send</button
    >
  </div>
</form>

{#each $messages as message}
  <div class="border rounded bg-white p-2 my-3">{message}</div>
{/each}
