<script setup lang="ts">
import logo from "@/assets/logo.png";
import { ScrollText } from "lucide-vue-next";
import SearchBar from "@/components/header/SearchBar.vue";
import Button from "../ui/button/Button.vue";
</script>

<template>
	<header class="w-full">
		<div class="flex items-center justify-between">
			<!-- Logo -->
			<a href="https://cyvore.com" target="_blank" rel="noopener noreferrer" class="flex items-center cursor-pointer">
				<img :src="logo" alt="Cyvore Logo" class="h-auto w-32 mr-3" />
			</a>
			<!-- Search Bar -->
			<SearchBar />

			<!-- Logs Button -->
			<Button variant="secondary" size="icon" class="size-12">
				<ScrollText :size="30" class="text-neutral-100" />
			</Button>
		</div>
	</header>
</template>
