<template>
	<div class="h-full rounded-2xl w-full bg-white/20 flex flex-col px-[30%] justify-center py-16">
		<div class="text-center mb-8">
			<h2 class="text-4xl font-bold text-neutral-800 mb-2">Create your account</h2>
			<p class="text-neutral-100">Join us today</p>
		</div>
		<form @submit.prevent="handleSubmit" class="space-y-6">
			<div class="space-y-4">
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
			</div>

			<Button type="submit" :disabled="loading" class="w-full bg-primary">
				<span v-if="loading" class="flex items-center justify-center">Creating account...</span>
				<span v-else>Create account</span>
			</Button>

			<div v-if="error" class="text-red-600 text-sm text-center bg-red-100 p-3 rounded-md border border-red-300">
				{{ error }}
			</div>

			<div class="text-center text-sm">
				<router-link to="/login" class="text-white hover:text-white/70 transition-colors"> Already have an account? Sign in </router-link>
			</div>
		</form>
	</div>
</template>

<script setup lang="ts">
import { reactive, ref } from "vue";
import { useRouter } from "vue-router";
import { AuthService } from "../services/auth";
import type { RegisterRequest } from "../types/auth";
import { Input } from "@/components/ui/input";
import { Button } from "@/components/ui/button";

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
