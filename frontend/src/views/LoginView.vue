<template>
	<AuthViewLayout
		title="Welcome back"
		subtitle="Sign in to your account"
		buttonText="Sign in"
		loadingText="Signing in..."
		:loading="loading"
		:error="error"
		@submit="handleSubmit"
	>
		<template #form-fields>
			<div>
				<label for="username" class="block text-sm font-medium text-neutral-700 mb-1">Username</label>
				<Input
					id="username"
					v-model="form.username"
					placeholder="Enter your username"
					type="text"
					required
					class="bg-white/90 text-neutral-800 placeholder:text-neutral-500"
					:class="{ 'border-red-500': errors.username }"
				/>
				<p v-if="errors.username" class="mt-1 text-sm text-red-500">{{ errors.username }}</p>
			</div>
			<div>
				<label for="password" class="block text-sm font-medium text-neutral-700 mb-1">Password</label>
				<Input
					id="password"
					v-model="form.password"
					placeholder="Enter your password"
					type="password"
					required
					class="bg-white/90 text-neutral-800 placeholder:text-neutral-500"
					:class="{ 'border-red-500': errors.password }"
				/>
				<p v-if="errors.password" class="mt-1 text-sm text-red-500">{{ errors.password }}</p>
			</div>
		</template>

		<template #footer-link>
			<router-link to="/register" class="text-white hover:text-white/70 transition-colors"> Don't have an account? Register </router-link>
		</template>
	</AuthViewLayout>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { AuthService } from "../services/auth";
import type { LoginRequest } from "../types/auth";
import { Input } from "@/components/ui/input";
import AuthViewLayout from "@/components/layout/AuthViewLayout.vue";
import { loginSchema, type LoginFormData } from "@/lib/validation";

const router = useRouter();
const authService = AuthService.getInstance();
const loading = ref(false);
const error = ref("");
const errors = reactive<Record<string, string>>({});

const form = reactive<LoginFormData>({
	username: "",
	password: "",
});

const validateForm = () => {
	try {
		loginSchema.parse(form);
		Object.keys(errors).forEach((key) => delete errors[key]);
		return true;
	} catch (err) {
		if (err instanceof Error) {
			const zodError = JSON.parse(err.message);
			zodError.forEach((error: { path: string[]; message: string }) => {
				errors[error.path[0]] = error.message;
			});
		}
		return false;
	}
};

const handleSubmit = async () => {
	if (!validateForm()) return;

	loading.value = true;
	error.value = "";

	try {
		const response = await authService.login(form);
		if (response.success) {
			router.push("/dashboard");
		} else {
			error.value = response.error || "Failed to login";
		}
	} catch (err) {
		error.value = "An error occurred";
	} finally {
		loading.value = false;
	}
};
</script>
