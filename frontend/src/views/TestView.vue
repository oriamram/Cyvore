<script setup lang="ts">
import { ref } from "vue";

const testData = ref([
	{ id: 1, name: "Test Item 1", status: "Active" },
	{ id: 2, name: "Test Item 2", status: "Inactive" },
	{ id: 3, name: "Test Item 3", status: "Active" },
]);
</script>

<template>
	<div class="flex flex-col items-center h-full w-full rounded-lg shadow-lg bg-white/20 p-6 gap-6">
		<h1 class="text-2xl font-bold text-gray-800">Test View</h1>

		<div class="w-full max-w-2xl bg-white rounded-lg shadow p-6">
			<h2 class="text-xl font-semibold mb-4">Test Data Table</h2>
			<div class="overflow-x-auto">
				<table class="min-w-full divide-y divide-gray-200">
					<thead class="bg-gray-50">
						<tr>
							<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">ID</th>
							<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Name</th>
							<th class="px-6 py-3 text-left text-xs font-medium text-gray-500 uppercase tracking-wider">Status</th>
						</tr>
					</thead>
					<tbody class="bg-white divide-y divide-gray-200">
						<tr v-for="item in testData" :key="item.id" class="hover:bg-gray-50">
							<td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ item.id }}</td>
							<td class="px-6 py-4 whitespace-nowrap text-sm text-gray-900">{{ item.name }}</td>
							<td class="px-6 py-4 whitespace-nowrap">
								<span
									class="px-2 inline-flex text-xs leading-5 font-semibold rounded-full"
									:class="item.status === 'Active' ? 'bg-green-100 text-green-800' : 'bg-red-100 text-red-800'"
								>
									{{ item.status }}
								</span>
							</td>
						</tr>
					</tbody>
				</table>
			</div>
		</div>
	</div>
</template>
