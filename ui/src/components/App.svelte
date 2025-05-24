<script>
	import { stores } from '../stores/stores.js';
	import Home from './Home.svelte';
	import Chat from './Chat.svelte';

	const { routeStore } = stores;

	routeStore.routes = [
		{
			path: '/',
			action: () => ({
				component: Home
			})
		},
		{
			path: '/chat',
			action: () => ({
				component: Chat
			})
		},
		{
			path: '*all',
			action: () => ({
				redirect: '/'
			})
		}
	];

	routeStore.start();

	const Page = $derived(routeStore.route?.component);
</script>

<main class="container vh-100 d-flex flex-column">
	<ul class="nav nav-pills my-2">
		<li class="nav-item">
			<a class="nav-link" class:active={Page === Home} href="/" data-link>HOME</a>
		</li>
		<li class="nav-item">
			<a class="nav-link" class:active={Page === Chat} href="/chat" data-link>CHAT</a>
		</li>
	</ul>
	<Page />
</main>
