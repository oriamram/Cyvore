<script setup lang="ts">
import { computed } from "vue";
import wsService from "@/services/websocket";
import { Pagination, PaginationContent, PaginationItem, PaginationPrevious, PaginationNext, PaginationEllipsis } from "@/components/ui/pagination";
import { ChevronLeft, ChevronRight } from "lucide-vue-next";

const props = defineProps<{
	type: "assets" | "relations";
}>();

const currentPage = computed(() => (props.type === "assets" ? wsService.currentAssetPage.value : wsService.currentRelationPage.value));

const totalPages = computed(() => (props.type === "assets" ? wsService.totalAssetPages.value : wsService.totalRelationPages.value));

const changePage = (page: number) => {
	if (page < 1 || page > totalPages.value) return;
	if (props.type === "assets") {
		wsService.requestPage(page, undefined);
	} else {
		wsService.requestPage(undefined, page);
	}
};
</script>

<template>
	<Pagination v-if="totalPages > 1" :items-per-page="108" :total="totalPages * 108" :page="currentPage" class="justify-center md:justify-end">
		<PaginationContent class="flex-wrap gap-1 md:gap-2">
			<PaginationPrevious :disabled="currentPage === 1" @click="changePage(currentPage - 1)" class="size-8 md:size-10">
				<ChevronLeft class="w-3 h-3 md:w-4 md:h-4" />
			</PaginationPrevious>
			<template v-if="totalPages <= 7">
				<PaginationItem
					v-for="page in totalPages"
					:key="page"
					:value="page"
					:is-active="page === currentPage"
					@click="changePage(page)"
					class="size-8 md:size-10 text-sm md:text-base"
				>
					{{ page }}
				</PaginationItem>
			</template>

			<template v-else>
				<PaginationItem v-if="currentPage > 3" :value="1" @click="changePage(1)" class="size-8 md:size-10 text-sm md:text-base">1</PaginationItem>
				<PaginationEllipsis v-if="currentPage > 4" class="size-8 md:size-10" />

				<template v-for="pageOffset in 3" :key="pageOffset">
					<PaginationItem
						v-if="currentPage - 2 + pageOffset > 0 && currentPage - 2 + pageOffset <= totalPages"
						:value="currentPage - 2 + pageOffset"
						:is-active="currentPage - 2 + pageOffset === currentPage"
						@click="changePage(currentPage - 2 + pageOffset)"
						class="size-8 md:size-10 text-sm md:text-base"
					>
						{{ currentPage - 2 + pageOffset }}
					</PaginationItem>
				</template>

				<PaginationEllipsis v-if="currentPage < totalPages - 3" class="size-8 md:size-10" />
				<PaginationItem v-if="currentPage < totalPages - 2" :value="totalPages" @click="changePage(totalPages)" class="size-8 md:size-10 text-sm md:text-base">
					{{ totalPages }}
				</PaginationItem>
			</template>

			<PaginationNext :disabled="currentPage === totalPages" @click="changePage(currentPage + 1)" class="size-8 md:size-10">
				<ChevronRight class="w-3 h-3 md:w-4 md:h-4" />
			</PaginationNext>
		</PaginationContent>
	</Pagination>
</template>
