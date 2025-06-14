import { ref } from "vue";

interface WebSocketData {
	assets: any[];
	relations: any[];
	total: {
		assets: number;
		relations: number;
	};
}

class WebSocketService {
	private ws: WebSocket | null = null;
	private reconnectAttempts = 0;
	private maxReconnectAttempts = 5;
	private reconnectTimeout = 3000;
	private listeners: Map<string, Set<(data: any) => void>> = new Map();
	private readonly PAGE_SIZE = 10;

	public isConnected = ref(false);
	public assets = ref<any[]>([]);
	public relations = ref<any[]>([]);
	public total = ref({ assets: 0, relations: 0 });
	public currentAssetPage = ref(1);
	public currentRelationPage = ref(1);
	public totalAssetPages = ref(0);
	public totalRelationPages = ref(0);
	public sortColumn = ref("");
	public sortDirection = ref<"asc" | "desc">("asc");
	public assetType = ref("");
	public assetFilter = ref("");
	public relationType = ref("");
	public relationFilter = ref("");

	constructor(private url: string) {}

	connect() {
		try {
			this.ws = new WebSocket(this.url);

			this.ws.onopen = () => {
				console.log("WebSocket connected");
				this.isConnected.value = true;
				this.reconnectAttempts = 0;
				// Send initial state like the test client
				this.sendState();
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
					const data = JSON.parse(event.data) as WebSocketData;
					console.log("📦 WebSocket data received:", data);

					// Update the data directly from the message
					this.assets.value = data.assets || [];
					this.relations.value = data.relations || [];
					this.total.value = data.total || { assets: 0, relations: 0 };

					// Update total pages
					this.totalAssetPages.value = Math.ceil(this.total.value.assets / this.PAGE_SIZE);
					this.totalRelationPages.value = Math.ceil(this.total.value.relations / this.PAGE_SIZE);

					// Notify listeners if needed
					if (this.listeners.has("data")) {
						this.listeners.get("data")?.forEach((callback) => callback(data));
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

	private sendState() {
		if (this.ws?.readyState === WebSocket.OPEN) {
			const state = {
				assetPage: this.currentAssetPage.value,
				assetPageSize: this.PAGE_SIZE,
				assetType: this.assetType.value,
				assetFilter: this.assetFilter.value,
				relationPage: this.currentRelationPage.value,
				relationPageSize: this.PAGE_SIZE,
				relationType: this.relationType.value,
				relationFilter: this.relationFilter.value,
				sortColumn: this.sortColumn.value,
				sortDirection: this.sortDirection.value,
			};
			this.ws.send(JSON.stringify(state));
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

	requestPage(assetPage?: number, relationPage?: number) {
		if (this.ws?.readyState === WebSocket.OPEN) {
			if (assetPage) this.currentAssetPage.value = assetPage;
			if (relationPage) this.currentRelationPage.value = relationPage;

			const state = {
				assetPage: this.currentAssetPage.value,
				assetPageSize: this.PAGE_SIZE,
				assetType: this.assetType.value,
				assetFilter: this.assetFilter.value,
				relationPage: this.currentRelationPage.value,
				relationPageSize: this.PAGE_SIZE,
				relationType: this.relationType.value,
				relationFilter: this.relationFilter.value,
				sortColumn: this.sortColumn.value,
				sortDirection: this.sortDirection.value,
			};
			this.ws.send(JSON.stringify(state));
		} else {
			console.error("WebSocket is not connected");
		}
	}

	requestSort(column: string) {
		if (this.sortColumn.value === column) {
			// Toggle direction if same column
			this.sortDirection.value = this.sortDirection.value === "asc" ? "desc" : "asc";
		} else {
			// New column, default to ascending
			this.sortColumn.value = column;
			this.sortDirection.value = "asc";
		}
		this.sendState();
	}

	disconnect() {
		if (this.ws) {
			this.ws.close();
			this.ws = null;
		}
	}
}

const wsService = new WebSocketService("ws://localhost:8081/ws");

export default wsService;
