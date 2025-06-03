import { z } from "zod";

export const loginSchema = z.object({
	username: z.string().min(1, "Username is required"),
	password: z.string().min(8, "Password must be at least 8 characters"),
});

export const registerSchema = z.object({
	username: z.string().min(1, "Username is required"),
	email: z.string().email("Invalid email address"),
	password: z.string().min(8, "Password must be at least 8 characters"),
});

export type LoginFormData = z.infer<typeof loginSchema>;
export type RegisterFormData = z.infer<typeof registerSchema>;
