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
				/>
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
				/>
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

const router = useRouter();
const authService = AuthService.getInstance();
const loading = ref(false);
const error = ref("");

const form = reactive<LoginRequest>({
	username: "",
	password: "",
});

const handleSubmit = async () => {
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
