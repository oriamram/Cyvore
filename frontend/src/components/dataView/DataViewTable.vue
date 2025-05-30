<script setup lang="ts">
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, TableEmpty } from "@/components/ui/table";
import { ArrowUpDown } from "lucide-vue-next";
import wsService from "@/services/websocket";
import { ref, watch } from "vue";

interface Column {
	key: string;
	label: string;
	formatter?: (value: any) => string;
	sortable?: boolean;
}

interface Props {
	columns: Column[];
	data: any[];
	emptyMessage?: string;
}

const props = defineProps<Props>();
const highlightRow = ref<Set<string>>(new Set());

const getFormattedValue = (row: any, column: Column) => {
	const value = row[column.key];
	if (column.formatter) {
		return column.formatter(value);
	}
	return value;
};

const handleSort = (column: Column) => {
	if (column.sortable !== false) {
		wsService.requestSort(column.key);
	}
};

const getSortIcon = (column: Column) => {
	if (column.sortable === false) return null;
	if (wsService.sortColumn.value !== column.key) {
		return ArrowUpDown;
	}
	return wsService.sortDirection.value === "asc" ? ArrowUpDown : ArrowUpDown;
};

// Watch for changes in data and highlight new rows
watch(
	() => props.data,
	(newData, oldData) => {
		if (!oldData) return;

		// Find new rows by comparing IDs
		const oldIds = new Set(oldData.map((row) => row.id));
		const newIds = new Set(newData.map((row) => row.id));

		// Add new IDs to highlight set
		newIds.forEach((id) => {
			if (!oldIds.has(id)) {
				highlightRow.value.add(id);
				// Remove highlight after animation
				setTimeout(() => {
					highlightRow.value.delete(id);
				}, 2000);
			}
		});
	},
	{ deep: true }
);
</script>

<template>
	<div class="w-full overflow-x-auto">
		<Table class="border-separate border-spacing-y-2 min-w-[640px]">
			<TableHeader>
				<TableRow class="hover:bg-transparent">
					<TableHead v-for="column in columns" :key="column.key" class="cursor-pointer select-none whitespace-nowrap" @click="handleSort(column)">
						<div class="flex items-center">
							<span class="text-sm md:text-base">{{ column.label }}</span>
							<component
								:is="getSortIcon(column)"
								v-if="column.sortable !== false"
								class="w-3 h-3 md:w-4 md:h-4 ml-1"
								:class="{ 'rotate-180': wsService.sortColumn.value === column.key && wsService.sortDirection.value === 'asc' }"
							/>
						</div>
					</TableHead>
				</TableRow>
			</TableHeader>
			<TableBody>
				<TableRow
					v-for="row in data"
					:key="row.id"
					class="bg-neutral-100 rounded-xl shadow-sm hover:bg-neutral-200 transition"
					:class="{ 'animate-highlight': highlightRow.has(row.id) }"
				>
					<TableCell
						v-for="column in columns"
						:key="column.key"
						class="text-xs md:text-sm font-medium px-2 md:px-4 first:rounded-l-lg last:rounded-r-lg whitespace-nowrap"
					>
						{{ getFormattedValue(row, column) }}
					</TableCell>
				</TableRow>
				<TableEmpty v-if="data.length === 0" :colspan="columns.length">
					{{ emptyMessage }}
				</TableEmpty>
			</TableBody>
		</Table>
	</div>
</template>

<style scoped>
@keyframes highlight {
	0% {
		background-color: rgb(34 197 94 / 0.2);
	}
	100% {
		background-color: rgb(243 244 246);
	}
}

.animate-highlight {
	animation: highlight 2s ease-out;
}
</style>
