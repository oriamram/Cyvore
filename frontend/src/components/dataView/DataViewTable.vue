<script setup lang="ts">
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, TableEmpty } from "@/components/ui/table";
import { ArrowUpDown } from "lucide-vue-next";
import wsService from "@/services/websocket";

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

defineProps<Props>();

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
</script>

<template>
	<Table class="border-separate border-spacing-y-2">
		<TableHeader>
			<TableRow class="hover:bg-transparent">
				<TableHead v-for="column in columns" :key="column.key" class="cursor-pointer select-none" @click="handleSort(column)">
					<div class="flex items-center">
						{{ column.label }}
						<component
							:is="getSortIcon(column)"
							v-if="column.sortable !== false"
							class="w-4 h-4 ml-1"
							:class="{ 'rotate-180': wsService.sortColumn.value === column.key && wsService.sortDirection.value === 'asc' }"
						/>
					</div>
				</TableHead>
			</TableRow>
		</TableHeader>
		<TableBody>
			<TableRow v-for="row in data" :key="row.id" class="bg-neutral-100 rounded-xl shadow-sm hover:bg-neutral-200 transition">
				<TableCell v-for="column in columns" :key="column.key" class="text-sm font-medium px-4 first:rounded-l-lg last:rounded-r-lg">
					{{ getFormattedValue(row, column) }}
				</TableCell>
			</TableRow>
			<TableEmpty v-if="data.length === 0" :colspan="columns.length">
				{{ emptyMessage }}
			</TableEmpty>
		</TableBody>
	</Table>
</template>
