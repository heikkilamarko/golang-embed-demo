<script>
	import { tick } from 'svelte';
	import { fly } from 'svelte/transition';
	import { stores } from '$lib/stores/stores.js';
	import PageTitle from '$lib/components/PageTitle.svelte';

	const { chatStore } = stores;

	let messagesEl;

	$effect(() => {
		chatStore.messages && scrollToBottom();
	});

	async function scrollToBottom() {
		await tick();
		messagesEl.scrollTop = messagesEl.scrollHeight;
	}

	function handleSubmit(event) {
		event.preventDefault();
		chatStore.sendMessage();
	}
</script>

<PageTitle>Chat</PageTitle>

<div class="text-center">
	{#if chatStore.isConnected}
		<span class="badge text-bg-success">CONNECTED</span>
	{:else}
		<span class="badge text-bg-primary">CONNECTING...</span>
	{/if}
</div>

<form class="row g-3 my-2 mb-5" autocomplete="off" onsubmit={handleSubmit}>
	<div class="col-sm-4">
		<div class="input-group">
			<span class="input-group-text bg-white text-primary">@</span>
			<input
				type="text"
				name="sender"
				class="form-control"
				placeholder="Type your name"
				bind:value={chatStore.sender}
			/>
		</div>
	</div>
	<div class="col-sm">
		<div class="input-group">
			<input
				type="text"
				name="message"
				class="form-control"
				placeholder="Type a new message"
				bind:value={chatStore.message}
			/>
			<button class="btn btn-primary" type="submit" disabled={!chatStore.canSendMessage}
				>Send</button
			>
		</div>
	</div>
</form>

<div class="flex-fill overflow-auto" bind:this={messagesEl}>
	{#each chatStore.messages as m (m.id)}
		<div
			class="border rounded bg-white py-2 px-3 mb-3 text-start"
			in:fly={{ x: 100, duration: 600 }}
		>
			<div class="sender mb-2">
				<span class="fw-bold" class:text-primary={m.sender === 'INFO'}>{m.sender}</span>
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
