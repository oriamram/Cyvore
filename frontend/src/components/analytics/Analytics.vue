<script setup lang="ts">
import AnalyticsCard from "./AnalyticsCard.vue";
import { ShieldCheck, Cpu, Database, Wifi } from "lucide-vue-next";
import { getSystemStatus } from "@/api/scanService";
import { ref, onMounted, onUnmounted, watch } from "vue";
import wsService from "@/services/websocket";

// Define your cards list
const cards = ref([
	{
		title: "WebSocket",
		content: "Disconnected",
		icon: Wifi,
		contentColor: "text-red-700",
	},
	{
		title: "Security",
		content: "Active Scan",
		icon: ShieldCheck,
		contentColor: "text-green-700",
	},
	{
		title: "CPU Load",
		content: "36%",
		icon: Cpu,
	},
	{
		title: "Database",
		content: "Connected",
		icon: Database,
	},
]);

let statusInterval: number;

const fetchStatus = async () => {
	try {
		const status = await getSystemStatus();
		cards.value[1].content = status.scanning ? "Active" : "Sleeping";
		cards.value[1].contentColor = status.scanning ? "text-green-700" : "text-red-700";
	} catch (error) {
		console.error("Error fetching status:", error);
	}
};

// Watch WebSocket connection status
const updateWebSocketStatus = () => {
	cards.value[0].content = wsService.isConnected.value ? "Connected" : "Disconnected";
	cards.value[0].contentColor = wsService.isConnected.value ? "text-green-700" : "text-red-700";
};

// Watch for changes in WebSocket connection status
watch(() => wsService.isConnected.value, updateWebSocketStatus);

onMounted(() => {
	// Initial fetch
	fetchStatus();
	// Set up interval for polling
	statusInterval = window.setInterval(fetchStatus, 5000);
	// Initial WebSocket status
	updateWebSocketStatus();
});

onUnmounted(() => {
	// Clean up interval when component is unmounted
	if (statusInterval) {
		clearInterval(statusInterval);
	}
});
</script>

<template>
	<div class="flex flex-col gap-4">
		<div class="w-full flex items-center justify-between gap-4 p-4">
			<AnalyticsCard v-for="(card, index) in cards" :key="index" :title="card.title" :content="card.content" :icon="card.icon" :contentColor="card.contentColor" />
		</div>
	</div>
</template>
