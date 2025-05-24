<script>
	import page from 'page';
	import Home from './Home.svelte';
	import Chat from './Chat.svelte';

	let currentPage = $state();

	page('/', () => (currentPage = Home));
	page('/chat', () => (currentPage = Chat));
	page('*', '/');
	page();

	let homeActive = $derived(currentPage === Home);
	let chatActive = $derived(currentPage === Chat);

	const Page = $derived(currentPage);
</script>

<main class="container vh-100 d-flex flex-column">
	<ul class="nav nav-pills my-2">
		<li class="nav-item">
			<a class="nav-link" class:active={homeActive} href="/">HOME</a>
		</li>
		<li class="nav-item">
			<a class="nav-link" class:active={chatActive} href="/chat">CHAT</a>
		</li>
	</ul>
	<Page />
</main>
