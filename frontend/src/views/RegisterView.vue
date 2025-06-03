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
					:class="{ 'border-red-500': errors.username }"
				/>
				<p v-if="errors.username" class="mt-1 text-sm text-red-500">{{ errors.username }}</p>
			</div>
			<div>
				<label for="email" class="block text-sm font-medium text-neutral-700 mb-1">Email</label>
				<Input
					id="email"
					v-model="form.email"
					placeholder="Enter your email"
					type="email"
					required
					class="bg-white/90 text-neutral-800 placeholder:text-neutral-500"
					:class="{ 'border-red-500': errors.email }"
				/>
				<p v-if="errors.email" class="mt-1 text-sm text-red-500">{{ errors.email }}</p>
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
					:class="{ 'border-red-500': errors.password }"
				/>
				<p v-if="errors.password" class="mt-1 text-sm text-red-500">{{ errors.password }}</p>
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
import { Input } from "@/components/ui/input";
import AuthViewLayout from "@/components/layout/AuthViewLayout.vue";
import { registerSchema, type RegisterFormData } from "@/lib/validation";

const router = useRouter();
const authService = AuthService.getInstance();
const loading = ref(false);
const error = ref("");
const errors = reactive<Record<string, string>>({});

const form = reactive<RegisterFormData>({
	username: "",
	email: "",
	password: "",
});

const validateForm = () => {
	try {
		registerSchema.parse(form);
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
