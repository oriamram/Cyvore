import { AuthService } from "./auth";

const API_URL = "http://localhost:4000";

export class ScanService {
	private static instance: ScanService;
	private authService: AuthService;

	private constructor() {
		this.authService = AuthService.getInstance();
	}

	public static getInstance(): ScanService {
		if (!ScanService.instance) {
			ScanService.instance = new ScanService();
		}
		return ScanService.instance;
	}

	private getHeaders(): HeadersInit {
		return {
			"Content-Type": "application/json",
			...this.authService.getAuthHeader(),
		};
	}

	public async startScan(target: string): Promise<any> {
		try {
			const response = await fetch(`${API_URL}/scan`, {
				method: "POST",
				headers: this.getHeaders(),
				body: JSON.stringify({ target }),
			});
			return await response.json();
		} catch (error) {
			console.error("Failed to start scan:", error);
			throw error;
		}
	}

	public async stopScan(): Promise<any> {
		try {
			const response = await fetch(`${API_URL}/scan/stop`, {
				method: "POST",
				headers: this.getHeaders(),
			});
			return await response.json();
		} catch (error) {
			console.error("Failed to stop scan:", error);
			throw error;
		}
	}

	public async getStatus(): Promise<any> {
		try {
			console.log(this.getHeaders());

			const response = await fetch(`${API_URL}/scan/status`, {
				method: "GET",
				headers: this.getHeaders(),
			});
			return await response.json();
		} catch (error) {
			console.error("Failed to get scan status:", error);
			throw error;
		}
	}

	public async getAmassData(): Promise<any> {
		try {
			const response = await fetch(`${API_URL}/amass/data`, {
				method: "GET",
				headers: this.getHeaders(),
			});
			return await response.json();
		} catch (error) {
			console.error("Failed to get Amass data:", error);
			throw error;
		}
	}
}
