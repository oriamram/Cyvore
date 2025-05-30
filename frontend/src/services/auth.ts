import { AuthResponse, LoginRequest, RegisterRequest, ApiResponse } from "../types/auth";

const API_URL = "http://localhost:4000";
const TOKEN_REFRESH_INTERVAL = 55 * 60 * 1000;

export class AuthService {
	private static instance: AuthService;
	private accessToken: string | null = null;
	private refreshInterval: number | null = null;

	private constructor() {
		// Initialize from localStorage if available
		this.accessToken = localStorage.getItem("access_token");
		if (this.accessToken) {
			this.startRefreshInterval();
		}
	}

	public static getInstance(): AuthService {
		if (!AuthService.instance) {
			AuthService.instance = new AuthService();
		}
		return AuthService.instance;
	}

	private startRefreshInterval() {
		// Clear any existing interval
		if (this.refreshInterval) {
			clearInterval(this.refreshInterval);
		}

		// Start new interval
		this.refreshInterval = window.setInterval(() => {
			this.refreshToken();
		}, TOKEN_REFRESH_INTERVAL);
	}

	private async refreshToken() {
		try {
			const response = await fetch(`${API_URL}/auth/refresh`, {
				method: "POST",
				credentials: "include", // Important for cookies
			});

			const result = await response.json();
			if (result.success && result.data?.access_token) {
				this.accessToken = result.data.access_token;
				localStorage.setItem("access_token", result.data.access_token);
			} else {
				// If refresh fails, clear everything and redirect to login
				this.clearAuth();
				window.location.href = "/login";
			}
		} catch (error) {
			console.error("Failed to refresh token:", error);
			this.clearAuth();
			window.location.href = "/login";
		}
	}

	private clearAuth() {
		this.accessToken = null;
		localStorage.removeItem("access_token");
		if (this.refreshInterval) {
			clearInterval(this.refreshInterval);
			this.refreshInterval = null;
		}
	}

	public getAccessToken(): string | null {
		return this.accessToken;
	}

	public isAuthenticated(): boolean {
		return !!this.accessToken;
	}

	public async register(data: RegisterRequest): Promise<ApiResponse<null>> {
		try {
			const response = await fetch(`${API_URL}/auth/register`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(data),
			});

			const result = await response.json();
			return result;
		} catch (error) {
			return {
				success: false,
				error: "Failed to register user",
			};
		}
	}

	public async login(data: LoginRequest): Promise<ApiResponse<AuthResponse>> {
		try {
			const response = await fetch(`${API_URL}/auth/signin`, {
				method: "POST",
				headers: {
					"Content-Type": "application/json",
				},
				body: JSON.stringify(data),
				credentials: "include", // Important for cookies
			});

			const result = await response.json();
			if (result.success && result.data?.access_token) {
				this.accessToken = result.data.access_token;
				localStorage.setItem("access_token", result.data.access_token);
				this.startRefreshInterval();
			}
			return result;
		} catch (error) {
			return {
				success: false,
				error: "Failed to login",
			};
		}
	}

	public async logout(): Promise<ApiResponse<null>> {
		try {
			const response = await fetch(`${API_URL}/auth/signout`, {
				method: "POST",
				credentials: "include",
			});

			const result = await response.json();
			if (result.success) {
				this.clearAuth();
			}
			return result;
		} catch (error) {
			return {
				success: false,
				error: "Failed to logout",
			};
		}
	}

	public getAuthHeader(): { Authorization: string } | {} {
		return this.accessToken ? { Authorization: `Bearer ${this.accessToken}` } : {};
	}
}
