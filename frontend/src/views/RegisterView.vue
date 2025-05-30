<template>
	<AuthViewLayout
		title="Create your account"
		subtitle="Join us today"
		buttonText="Create account"
		loadingText="Creating account..."
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
					placeholder="Choose a username"
					type="text"
					required
					class="bg-white/90 text-neutral-800 placeholder:text-neutral-500"
				/>
			</div>
			<div>
				<label for="email" class="block text-sm font-medium text-neutral-700 mb-1">Email</label>
				<Input id="email" v-model="form.email" placeholder="Enter your email" type="email" required class="bg-white/90 text-neutral-800 placeholder:text-neutral-500" />
			</div>
			<div>
				<label for="password" class="block text-sm font-medium text-neutral-700 mb-1">Password</label>
				<Input
					id="password"
					v-model="form.password"
					placeholder="Create a password"
					type="password"
					required
					class="bg-white/90 text-neutral-800 placeholder:text-neutral-500"
				/>
			</div>
		</template>

		<template #footer-link>
			<router-link to="/login" class="text-white hover:text-white/70 transition-colors"> Already have an account? Sign in </router-link>
		</template>
	</AuthViewLayout>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { AuthService } from "../services/auth";
import type { RegisterRequest } from "../types/auth";
import { Input } from "@/components/ui/input";
import AuthViewLayout from "@/components/layout/AuthViewLayout.vue";

const router = useRouter();
const authService = AuthService.getInstance();
const loading = ref(false);
const error = ref("");

const form = reactive<RegisterRequest>({
	username: "",
	email: "",
	password: "",
});

const handleSubmit = async () => {
	loading.value = true;
	error.value = "";

	try {
		const response = await authService.register(form);
		if (response.success) {
			router.push("/login");
		} else {
			error.value = response.error || "Failed to register";
		}
	} catch (err) {
		error.value = "An error occurred";
	} finally {
		loading.value = false;
	}
};
</script>
