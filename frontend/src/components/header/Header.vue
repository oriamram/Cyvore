<script setup lang="ts">
import logo from "@/assets/logo.png";
import { LogOut, Square, Trash2 } from "lucide-vue-next";
import SearchBar from "@/components/header/SearchBar.vue";
import Button from "../ui/button/Button.vue";
import { useRouter } from "vue-router";
import { AuthService } from "@/services/auth";
import { toast } from "vue-sonner";
import { ScanService } from "@/services/scanService";

const router = useRouter();
const authService = AuthService.getInstance();

const handleSignOut = async () => {
	try {
		const response = await authService.logout();
		if (response.success) {
			router.push("/login");
		} else {
			toast.error("Failed to sign out");
		}
	} catch (error) {
		console.error("Error signing out:", error);
		toast.error("Failed to sign out");
	}
};

const handleStopScan = async () => {
	try {
		const response = await ScanService.getInstance().stopScan();
		toast.success(response.message);
	} catch (error) {
		console.error("Error stopping scan:", error);
		toast.error("Failed to stop scan");
	}
};

const handleClean = async () => {
	try {
		const response = await ScanService.getInstance().clean();
		if (response.success) {
			toast.success("Database and logs cleaned successfully");
		} else {
			toast.error("Failed to clean database and logs");
		}
	} catch (error) {
		console.error("Error cleaning database and logs:", error);
		toast.error("Failed to clean database and logs");
	}
};
</script>

<template>
	<header class="w-full">
		<div class="flex flex-col md:flex-row items-center md:justify-between gap-4 md:gap-0 w-full">
			<!-- Top row: logo and signout button on mobile, only logo on desktop -->
			<div class="flex w-full items-center justify-between md:w-auto md:justify-start md:gap-6">
				<!-- Logo -->
				<a href="https://cyvore.com" target="_blank" rel="noopener noreferrer" class="flex items-center cursor-pointer">
					<img :src="logo" alt="Cyvore Logo" class="h-auto w-24 md:w-32" />
				</a>

				<!-- Signout Button (visible in top row only on mobile) -->
				<div class="flex gap-2 md:hidden">
					<Button variant="secondary" size="icon" class="size-10 bg-blue-400 hover:bg-blue-300" @click="handleClean">
						<Trash2 :size="24" class="text-neutral-100" />
					</Button>
					<Button variant="secondary" size="icon" class="size-10 bg-red-400 hover:bg-red-300" @click="handleStopScan">
						<Square :size="24" class="text-neutral-100" />
					</Button>
					<Button variant="secondary" size="icon" class="size-10" @click="handleSignOut">
						<LogOut :size="24" class="text-neutral-100" />
					</Button>
				</div>
			</div>

			<!-- Search bar -->
			<SearchBar class="w-full md:max-w-xl" />

			<!-- Signout Button (visible in main row on desktop) -->
			<div class="w-32 flex justify-end gap-2">
				<Button variant="secondary" size="icon" class="hidden md:flex size-10 md:size-12 bg-blue-400 hover:bg-blue-300" @click="handleClean">
					<Trash2 :size="24" class="text-neutral-100 md:text-[30px]" />
				</Button>
				<Button variant="secondary" size="icon" class="hidden md:flex size-10 md:size-12 bg-red-400 hover:bg-red-300" @click="handleStopScan">
					<Square :size="24" class="text-neutral-100 md:text-[30px]" />
				</Button>
				<Button variant="secondary" size="icon" class="hidden md:flex size-10 md:size-12" @click="handleSignOut">
					<LogOut :size="24" class="text-neutral-100 md:text-[30px]" />
				</Button>
			</div>
		</div>
	</header>
</template>
