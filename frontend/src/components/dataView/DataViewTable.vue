<script setup lang="ts">
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow, TableEmpty } from "@/components/ui/table";

interface Column {
	key: string;
	label: string;
	formatter?: (value: any) => string;
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
</script>

<template>
	<Table class="border-separate border-spacing-y-2">
		<TableHeader>
			<TableRow class="hover:bg-transparent">
				<TableHead v-for="column in columns" :key="column.key">
					{{ column.label }}
				</TableHead>
			</TableRow>
		</TableHeader>
		<TableBody>
			<TableRow v-for="row in data" :key="row.id" class="hover:bg-transparent">
				<TableCell :colspan="columns.length" class="p-0">
					<div class="flex w-full bg-neutral-100 rounded-lg px-4 py-2 shadow-md">
						<div v-for="column in columns" :key="column.key" class="flex-1 text-sm font-medium">
							{{ getFormattedValue(row, column) }}
						</div>
					</div>
				</TableCell>
			</TableRow>
			<TableEmpty v-if="data.length === 0" :colspan="columns.length">
				{{ emptyMessage }}
			</TableEmpty>
		</TableBody>
	</Table>
</template>
