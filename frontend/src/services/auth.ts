import { AuthResponse, LoginRequest, RegisterRequest, ApiResponse } from "../types/auth";

const API_URL = "http://localhost:4000";

export class AuthService {
	private static instance: AuthService;
	private accessToken: string | null = null;

	private constructor() {
		// Initialize from localStorage if available
		this.accessToken = localStorage.getItem("access_token");
	}

	public static getInstance(): AuthService {
		if (!AuthService.instance) {
			AuthService.instance = new AuthService();
		}
		return AuthService.instance;
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
				credentials: "include", // Important for cookies
			});

			const result = await response.json();
			if (result.success) {
				this.accessToken = null;
				localStorage.removeItem("access_token");
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
