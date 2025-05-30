<script setup lang="ts">
import Card from "../ui/card/Card.vue";
import CardHeader from "../ui/card/CardHeader.vue";
import CardContent from "../ui/card/CardContent.vue";
import CardTitle from "../ui/card/CardTitle.vue";
import type { Component } from "vue";

const props = defineProps<{
	icon: Component;
	title: string;
	content: string;
	contentColor?: string;
}>();
</script>

<template>
	<Card class="w-[200px] flex flex-col justify-center gap-2 p-4 bg-white/90">
		<CardHeader class="flex items-baseline p-0">
			<div>
				<component :is="props.icon" class="w-6 h-6 text-gray-500" />
			</div>
			<CardTitle class="-translate-y-1 ml-1 font-medium text-sm">{{ props.title }}</CardTitle>
		</CardHeader>

		<CardContent class="p-0">
			<div :class="['text-2xl font-bold', props.contentColor ?? 'text-neutral-900']">
				{{ props.content }}
			</div>
		</CardContent>
	</Card>
</template>
