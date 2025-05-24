import UniversalRouter from 'universal-router';

export class RouteStore {
	routes = $state([]);
	route = $state();
	router;

	start() {
		this.router = new UniversalRouter(this.routes);

		document.addEventListener('click', (e) => {
			const link = e.target.closest('a[data-link]');
			if (!link) return;
			e.preventDefault();
			const href = link.getAttribute('href');
			history.pushState(null, '', href);
			this.render(href);
		});

		window.addEventListener('popstate', () => this.render(window.location.href));

		this.render(window.location.href);
	}

	async render(href) {
		console.debug(href);
		try {
			const url = new URL(href, window.location.origin);
			const route = await this.router.resolve({ pathname: url.pathname });
			if (route.redirect) {
				window.location = route.redirect;
				return;
			}
			this.route = route;
		} catch (err) {
			console.error(err);
		}
	}
}
