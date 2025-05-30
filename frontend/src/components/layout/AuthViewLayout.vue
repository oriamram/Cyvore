<template>
	<div class="h-full rounded-2xl w-full bg-white/20 flex flex-col md:px-[30%] px-10 justify-center py-16">
		<div class="text-center mb-8">
			<h2 class="text-4xl font-bold text-neutral-800 mb-2">{{ title }}</h2>
			<p class="text-neutral-100">{{ subtitle }}</p>
		</div>
		<form @submit.prevent="$emit('submit')" class="space-y-6">
			<div class="space-y-4">
				<slot name="form-fields"></slot>
			</div>

			<Button type="submit" :disabled="loading" class="w-full bg-primary">
				<span v-if="loading" class="flex items-center justify-center">{{ loadingText }}</span>
				<span v-else>{{ buttonText }}</span>
			</Button>

			<div v-if="error" class="text-red-600 text-sm text-center bg-red-100 p-3 rounded-md border border-red-300">
				{{ error }}
			</div>

			<div class="text-center text-sm">
				<slot name="footer-link"></slot>
			</div>
		</form>
	</div>
</template>

<script setup lang="ts">
import { Button } from "@/components/ui/button";

defineProps<{
	title: string;
	subtitle: string;
	buttonText: string;
	loadingText: string;
	loading: boolean;
	error?: string;
}>();

defineEmits<{
	(e: "submit"): void;
}>();
</script>
