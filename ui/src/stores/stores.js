import { RouteStore } from './RouteStore.svelte.js';
import { ChatStore } from './ChatStore.svelte.js';
import { HomeStore } from './HomeStore.svelte.js';

export const stores = {};

stores.routeStore = new RouteStore();
stores.chatStore = new ChatStore();
stores.homeStore = new HomeStore();
