import { HomeStore } from './HomeStore.svelte.js';
import { ChatStore } from './ChatStore.svelte.js';

export const stores = {};

export function createStores() {
	stores.homeStore ??= new HomeStore();
	stores.chatStore ??= new ChatStore();
}
