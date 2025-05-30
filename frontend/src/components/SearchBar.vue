<script setup>
import Button from "./ui/button/Button.vue";
import Input from "./ui/input/Input.vue";
import { Eye } from "lucide-vue-next";
</script>
<template>
	<div class="bg-white rounded-full flex items-center p-1 gap-3 w-1/4">
		<Input class="rounded-full border-none !ring-0 font-medium !text-xl" type="text" />
		<Button class="rounded-full w-10 h-10 bg-primary">
			<Eye class="text-white -scale-x-[100%]" />
		</Button>
	</div>
</template>
