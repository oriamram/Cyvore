<script setup>
import Analytics from "@/components/analytics/Analytics.vue";
import Header from "@/components/header/Header.vue";
import DataView from "@/components/dataView/DataView.vue";
import Button from "@/components/ui/button/Button.vue";
import wsService from "@/services/websocket";
import Separator from "@/components/ui/separator/Separator.vue";
</script>

<template>
	<div class="flex flex-col grow items-center h-full w-full rounded-lg shadow-lg bg-white/20 p-6 px-24 max-md:p-3 md:gap-16 gap-10">
		<Header />
		<Separator class="bg-neutral-300 md:hidden" />
		<Analytics />
		<Separator class="bg-neutral-300 md:hidden" />
		<DataView />
	</div>
</template>
