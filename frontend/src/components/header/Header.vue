<script setup lang="ts">
import logo from "@/assets/logo.png";
import { ScrollText } from "lucide-vue-next";
import SearchBar from "@/components/header/SearchBar.vue";
import Button from "../ui/button/Button.vue";
</script>

<template>
	<header class="w-full">
		<div class="flex flex-col md:flex-row items-center md:justify-between gap-4 md:gap-0 w-full">
			<!-- Top row: logo and logs button on mobile, only logo on desktop -->
			<div class="flex w-full items-center justify-between md:w-auto md:justify-start md:gap-6">
				<!-- Logo -->
				<a href="https://cyvore.com" target="_blank" rel="noopener noreferrer" class="flex items-center cursor-pointer">
					<img :src="logo" alt="Cyvore Logo" class="h-auto w-24 md:w-32" />
				</a>

				<!-- Logs Button (visible in top row only on mobile) -->
				<Button variant="secondary" size="icon" class="size-10 md:hidden">
					<ScrollText :size="24" class="text-neutral-100" />
				</Button>
			</div>

			<!-- Search bar -->
			<SearchBar class="w-full md:max-w-xl" />

			<!-- Logs Button (visible in main row on desktop) -->
			<div class="w-32 flex justify-end">
				<Button variant="secondary" size="icon" class="hidden md:flex size-10 md:size-12">
					<ScrollText :size="24" class="text-neutral-100 md:text-[30px]" />
				</Button>
			</div>
		</div>
	</header>
</template>
