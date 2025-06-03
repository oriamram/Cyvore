<script setup>
import { RouterView } from "vue-router";
import { Toaster } from "vue-sonner";
import "vue-sonner/style.css";
</script>

<template>
	<div class="flex flex-col items-center max-md:justify-center min-h-screen p-6 max-md:p-3 gradient-background font-rubik">
		<RouterView />
	</div>
	<Toaster />
</template>
