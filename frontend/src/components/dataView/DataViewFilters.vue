<script setup lang="ts">
import { ref } from "vue";
import wsService from "@/services/websocket";
import { Select, SelectContent, SelectGroup, SelectItem, SelectLabel, SelectTrigger, SelectValue } from "@/components/ui/select";

const props = defineProps<{
	type: "assets" | "relations";
}>();

const selectedType = ref("");

// Types from the actual data
const assetTypes = ["FQDN", "IPAddress", "Netblock", "ASN", "RIROrganization"];

const relationTypes = ["ns_record", "a_record", "aaaa_record", "node", "mx_record", "contains", "managed_by", "announces"];

// Get types based on current view
const getTypes = () => {
	return props.type === "assets" ? assetTypes : relationTypes;
};

// Update filters when type changes
const handleTypeChange = (value: any) => {
	selectedType.value = value || "";
	if (props.type === "assets") {
		wsService.assetType.value = value || "";
		wsService.requestPage(1, undefined);
	} else {
		wsService.relationType.value = value || "";
		wsService.requestPage(undefined, 1);
	}
};

// Reset filters
const resetFilters = () => {
	selectedType.value = "";
	if (props.type === "assets") {
		wsService.assetType.value = "";
	} else {
		wsService.relationType.value = "";
	}
	wsService.requestPage(props.type === "assets" ? 1 : undefined, props.type === "relations" ? 1 : undefined);
};
</script>

<template>
	<div class="flex-1">
		<Select v-model="selectedType" @update:modelValue="handleTypeChange">
			<SelectTrigger class="w-full bg-white">
				<SelectValue :placeholder="`Filter by ${type.slice(0, -1)} type`" />
			</SelectTrigger>
			<SelectContent>
				<SelectGroup>
					<SelectLabel class="text-xs font-medium text-gray-500 uppercase tracking-wider"> {{ type.slice(0, -1) }} Types </SelectLabel>
					<SelectItem v-for="type in getTypes()" :key="type" :value="type" class="text-sm">
						{{ type }}
					</SelectItem>
				</SelectGroup>
			</SelectContent>
		</Select>
	</div>
	<Button
		@click="resetFilters"
		class="px-3 py-2 text-sm font-medium text-gray-600 hover:text-gray-900 bg-white border border-gray-200 rounded-md hover:bg-gray-50 transition-colors"
	>
		Reset
	</Button>
</template>
