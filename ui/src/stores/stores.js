import { ChatStore } from './ChatStore.svelte.js';
import { HomeStore } from './HomeStore.svelte.js';

export const stores = {};

stores.chatStore = new ChatStore();
stores.homeStore = new HomeStore();
