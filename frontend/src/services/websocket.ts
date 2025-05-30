import { ref } from "vue";

class WebSocketService {
	private ws: WebSocket | null = null;
	private reconnectAttempts = 0;
	private maxReconnectAttempts = 5;
	private reconnectTimeout = 3000;
	private listeners: Map<string, Set<(data: any) => void>> = new Map();

	public isConnected = ref(false);

	constructor(private url: string) {}

	connect() {
		try {
			this.ws = new WebSocket(this.url);

			this.ws.onopen = () => {
				console.log("WebSocket connected");
				this.isConnected.value = true;
				this.reconnectAttempts = 0;
			};

			this.ws.onclose = () => {
				console.log("WebSocket disconnected");
				this.isConnected.value = false;
				this.handleReconnect();
			};

			this.ws.onerror = (error) => {
				console.error("WebSocket error:", error);
			};

			this.ws.onmessage = (event) => {
				try {
					const data = JSON.parse(event.data);
					const { type, payload } = data;

					if (type && this.listeners.has(type)) {
						this.listeners.get(type)?.forEach((callback) => callback(payload));
					}
				} catch (error) {
					console.error("Error parsing WebSocket message:", error);
				}
			};
		} catch (error) {
			console.error("Error creating WebSocket connection:", error);
			this.handleReconnect();
		}
	}

	private handleReconnect() {
		if (this.reconnectAttempts < this.maxReconnectAttempts) {
			this.reconnectAttempts++;
			console.log(`Attempting to reconnect (${this.reconnectAttempts}/${this.maxReconnectAttempts})...`);
			setTimeout(() => this.connect(), this.reconnectTimeout);
		}
	}

	subscribe(type: string, callback: (data: any) => void) {
		if (!this.listeners.has(type)) {
			this.listeners.set(type, new Set());
		}
		this.listeners.get(type)?.add(callback);
	}

	unsubscribe(type: string, callback: (data: any) => void) {
		this.listeners.get(type)?.delete(callback);
	}

	send(type: string, payload: any) {
		if (this.ws?.readyState === WebSocket.OPEN) {
			this.ws.send(JSON.stringify({ type, payload }));
		} else {
			console.error("WebSocket is not connected");
		}
	}

	disconnect() {
		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}
	}
}

// Create a singleton instance
const wsService = new WebSocketService("ws://localhost:8081/ws");

export default wsService;
