<script>
	import '../app.scss';
	import { page } from '$app/state';

	let { children } = $props();
</script>

<svelte:head>
	<title>Go Embed Demo</title>
	<meta name="description" content="A demo of Go Embed with SvelteKit" />
</svelte:head>

<main class="container vh-100 d-flex flex-column">
	<ul class="nav nav-pills my-2">
		<li class="nav-item">
			<a class="nav-link" class:active={page.url.pathname === '/'} href="/">HOME</a>
		</li>
		<li class="nav-item">
			<a class="nav-link" class:active={page.url.pathname === '/chat'} href="/chat">CHAT</a>
		</li>
	</ul>
	{@render children()}
</main>
