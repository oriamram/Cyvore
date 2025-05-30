<script setup>
import HomeView from "./views/HomeView.vue";
import { Toaster } from "vue-sonner";
import "vue-sonner/style.css";
</script>

<template>
	<div class="flex flex-col items-center h-screen p-6 gradient-background font-rubik">
		<HomeView />
	</div>
	<Toaster />
</template>
