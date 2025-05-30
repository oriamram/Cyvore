<script setup lang="ts">
import AnalyticsCard from "./AnalyticsCard.vue";
import { Radar, Database, Wifi } from "lucide-vue-next";
import { ref, onMounted, onUnmounted, watch } from "vue";
import wsService from "@/services/websocket";
import { ScanService } from "@/services/scanService";

// Define your cards list
const cards = ref([
	{
		title: "WebSocket",
		content: "Disconnected",
		icon: Wifi,
		contentColor: "text-red-700",
	},
	{
		title: "Scanner",
		content: "Active Scan",
		icon: Radar,
		contentColor: "text-green-700",
	},
	{
		title: "Total Assets",
		content: "0",
		icon: Database,
		contentColor: "text-blue-700",
	},
	{
		title: "Total Relations",
		content: "0",
		icon: Database,
		contentColor: "text-purple-700",
	},
]);

let statusInterval: number;

const fetchStatus = async () => {
	try {
		const status = await ScanService.getInstance().getStatus();
		cards.value[1].content = status.scanning ? "Active" : "Sleeping";
		cards.value[1].contentColor = status.scanning ? "text-green-700" : "text-yellow-600";
	} catch (error) {
		console.error("Error fetching status:", error);
	}
};

// Update WebSocket status
const updateWebSocketStatus = () => {
	cards.value[0].content = wsService.isConnected.value ? "Connected" : "Disconnected";
	cards.value[0].contentColor = wsService.isConnected.value ? "text-green-700" : "text-red-700";
};

// Update total counts
const updateTotalCounts = () => {
	cards.value[2].content = wsService.total.value.assets.toString();
	cards.value[3].content = wsService.total.value.relations.toString();
};

// Watch for changes in WebSocket connection status
watch(() => wsService.isConnected.value, updateWebSocketStatus);

// Watch for changes in total counts
watch(() => wsService.total.value.assets, updateTotalCounts);
watch(() => wsService.total.value.relations, updateTotalCounts);

onMounted(() => {
	fetchStatus();
	statusInterval = window.setInterval(fetchStatus, 5000);
	updateWebSocketStatus();
	updateTotalCounts();
});

onUnmounted(() => {
	if (statusInterval) {
		clearInterval(statusInterval);
	}
});
</script>

<template>
	<div class="max-md:w-full grid grid-cols-2 sm:grid-cols-2 md:grid-cols-2 lg:grid-cols-4 gap-4 md:gap-8">
		<AnalyticsCard v-for="(card, index) in cards" :key="index" :title="card.title" :content="card.content" :icon="card.icon" :contentColor="card.contentColor" />
	</div>
</template>
