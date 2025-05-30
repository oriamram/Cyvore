<script setup>
import { ref } from "vue";
import Button from "../ui/button/Button.vue";
import Input from "../ui/input/Input.vue";
import { Eye } from "lucide-vue-next";
import { initiateScan } from "@/api/scanService";
import { toast } from "vue-sonner";

const searchValue = ref("");

const handleScanClick = async () => {
	const domain = searchValue.value;
	if (!domain) {
		toast.warning("Please enter a domain to scan.");
		return;
	}

	try {
		const response = await initiateScan(domain);
		toast.success(`Scan initiated for ${domain}!`);
	} catch (error) {
		console.error("Error initiating scan:", error);
		toast.error(`Failed to initiate scan for ${domain}. Error: ${error.message}`);
	}
};
</script>

<template>
	<div class="bg-white rounded-full flex items-center p-1 gap-3 w-1/4">
		<Input class="rounded-full border-none !ring-0 font-medium !text-xl" type="text" v-model="searchValue" placeholder="cyvore.com" />
		<Button class="rounded-full w-10 h-10 bg-primary" @click="handleScanClick">
			<Eye class="text-white -scale-x-[100%]" />
		</Button>
	</div>
</template>
