<script>
	import { fly } from 'svelte/transition';
	import { stores } from '$lib/stores/stores.js';
	import PageTitle from '$lib/components/PageTitle.svelte';

	const { homeStore } = stores;
</script>

<PageTitle>Go Embed Demo</PageTitle>

<section class="text-center">
	<button class="btn btn-primary rounded-pill px-4 my-4" onclick={() => homeStore.loadMessage()}
		>{homeStore.isLoading ? 'Waiting...' : 'Call API'}</button
	>
	{#if homeStore.message}
		<p class="fs-3 text-primary" in:fly={{ x: 100, duration: 600 }}>
			{homeStore.message}
		</p>
	{:else if homeStore.error}
		<p class="fs-3 text-danger" in:fly={{ x: 100, duration: 600 }}>
			{homeStore.error}
		</p>
	{/if}
</section>
