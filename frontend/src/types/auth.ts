export interface User {
	id: string;
	username: string;
	email: string;
}

export interface AuthResponse {
	access_token: string;
	user: User;
}

export interface RegisterRequest {
	username: string;
	password: string;
	email: string;
}

export interface LoginRequest {
	username: string;
	password: string;
}

export interface ApiResponse<T> {
	success: boolean;
	data?: T;
	error?: string;
}
