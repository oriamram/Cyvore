<script setup lang="ts">
import wsService from "@/services/websocket";
import DataViewTable from "./DataViewTable.vue";
import { Tabs, TabsContent, TabsList, TabsTrigger } from "@/components/ui/tabs";
import CustomPagination from "./DataViewPagination.vue";
import DataViewFilters from "./DataViewFilters.vue";

interface Asset {
	id: string;
	created_at: string;
	type: string;
	content: string;
	last_seen: string;
}

interface Relation {
	id: string;
	created_at: string;
	type: string;
	from_asset_id: string;
	to_asset_id: string;
	last_seen: string;
}

const assets = wsService.assets as unknown as Asset[];
const relations = wsService.relations as unknown as Relation[];

const getAssetContent = (content: string) => {
	try {
		const parsed = JSON.parse(content);
		if (parsed.name) return parsed.name;
		if (parsed.address) return `${parsed.address} (${parsed.type})`;
		return content;
	} catch {
		return content;
	}
};

// Format date
const formatDate = (dateStr: string) => {
	return new Date(dateStr).toLocaleString();
};

const assetColumns = [
	{ key: "type", label: "Type", sortable: true },
	{ key: "content", label: "Content", formatter: getAssetContent, sortable: false },
	{ key: "created_at", label: "Created", formatter: formatDate, sortable: true },
	{ key: "last_seen", label: "Last Seen", formatter: formatDate, sortable: true },
];

const relationColumns = [
	{ key: "type", label: "Type", sortable: true },
	{ key: "from_asset_id", label: "From Asset ID", sortable: false },
	{ key: "to_asset_id", label: "To Asset ID", sortable: false },
	{ key: "created_at", label: "Created", formatter: formatDate, sortable: true },
	{ key: "last_seen", label: "Last Seen", formatter: formatDate, sortable: true },
];

import { ref } from "vue";

const selectedTab = ref<"assets" | "relations">("assets");
</script>

<template>
	<Tabs v-model="selectedTab" default-value="assets" class="w-full gap-0">
		<div class="flex flex-col md:flex-row items-start md:items-center justify-between gap-4 mb-4">
			<div class="flex flex-col md:flex-row items-start md:items-center gap-4 w-full md:w-auto">
				<TabsList class="w-full md:w-auto">
					<TabsTrigger value="assets">Assets</TabsTrigger>
					<TabsTrigger value="relations">Relations</TabsTrigger>
				</TabsList>
				<DataViewFilters :type="selectedTab" class="w-full md:w-auto" />
			</div>
			<CustomPagination :type="selectedTab" class="w-full max-md:hidden" />
		</div>

		<TabsContent value="assets" class="overflow-x-auto">
			<DataViewTable :columns="assetColumns" :data="assets" empty-message="No assets available" />
		</TabsContent>
		<TabsContent value="relations" class="overflow-x-auto">
			<DataViewTable :columns="relationColumns" :data="relations" empty-message="No relations available" />
		</TabsContent>
	</Tabs>
	<CustomPagination :type="selectedTab" class="w-full md:hidden" />
</template>
