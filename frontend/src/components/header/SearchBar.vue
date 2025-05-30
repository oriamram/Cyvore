<script setup>
import { ref } from "vue";
import Button from "../ui/button/Button.vue";
import Input from "../ui/input/Input.vue";
import { Eye } from "lucide-vue-next";
import { toast } from "vue-sonner";
import { ScanService } from "@/services/scanService";

const searchValue = ref("");

const handleScanClick = async () => {
	const domain = searchValue.value;
	if (!domain) {
		toast.warning("Please enter a domain to scan.");
		return;
	}

	try {
		const response = await ScanService.getInstance().startScan(domain);
		toast.success(`Scan initiated for ${domain}!`);
	} catch (error) {
		console.error("Error initiating scan:", error);
		toast.error(`Failed to initiate scan for ${domain}. Error: ${error.message}`);
	}
};
</script>

<template>
	<div class="bg-white rounded-full flex items-center p-1 gap-2 md:gap-3 w-full md:w-1/2 lg:w-1/3">
		<Input class="rounded-full border-none !ring-0 font-medium !text-base md:!text-xl shadow-none" type="text" v-model="searchValue" placeholder="cyvore.com" />
		<Button class="rounded-full w-8 h-8 md:w-10 md:h-10 bg-primary" @click="handleScanClick">
			<Eye class="text-white -scale-x-[100%] w-4 h-4 md:w-5 md:h-5" />
		</Button>
	</div>
</template>
