import { v4 as uuid } from 'uuid';

export class ChatStore {
	conn;

	sender = $state('');
	message = $state('');
	messages = $state([]);
	connectionState = $state();

	isConnected = $derived(this.connectionState === WebSocket.OPEN);
	canSendMessage = $derived(this.isConnected && !!this.sender && !!this.message);

	constructor() {
		setInterval(() => (this.connectionState = this.conn?.readyState ?? null), 1000);

		this.connect();
		setInterval(() => this.connect(), 2000);
	}

	addMessage(message) {
		message.ts = new Date(message.ts);
		this.messages = [...this.messages, message];
	}

	addInfoMessage(message) {
		this.addMessage({
			sender: 'INFO',
			message,
			ts: new Date()
		});
	}

	sendMessage() {
		this.conn?.send(
			JSON.stringify({
				id: uuid(),
				sender: this.sender,
				message: this.message,
				ts: new Date()
			})
		);
		this.message = '';
	}

	connect() {
		if (this.conn && this.conn.readyState !== WebSocket.CLOSED) return;
		this.conn = new WebSocket(`ws://${window.location.host}/ws`);
		this.conn.onopen = () => this.addInfoMessage('Connection opened.');
		this.conn.onclose = () => this.addInfoMessage('Connection closed.');
		this.conn.onerror = () => this.addInfoMessage('Connection error.');
		this.conn.onmessage = (e) => this.addMessage(JSON.parse(e.data));
	}
}
