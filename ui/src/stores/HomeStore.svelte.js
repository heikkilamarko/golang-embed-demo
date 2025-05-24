export class HomeStore {
	isLoading = $state(false);
	message = $state();
	error = $state();

	async loadMessage() {
		if (this.isLoading) return;
		try {
			this.isLoading = true;
			this.message = null;
			this.error = null;
			const response = await fetch('/api/message');
			if (response.ok) {
				this.message = await response.text();
			} else {
				this.error = await response.text();
			}
		} catch (err) {
			this.error = err.message;
		} finally {
			this.isLoading = false;
		}
	}
}
