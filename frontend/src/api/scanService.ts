export const initiateScan = async (domain: string) => {
	try {
		const response = await fetch(`http://localhost:4000/scan?domain=${domain}`, {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});

		if (response.ok) {
			return await response.json();
		} else {
			console.error("Failed to send scan request:", response.status, response.statusText);
			throw new Error(`Failed to initiate scan: ${response.statusText}`);
		}
	} catch (error) {
		console.error("Error initiating scan:", error);
		throw error;
	}
};

export const getSystemStatus = async () => {
	try {
		const response = await fetch("http://localhost:4000/scan/status", {
			method: "GET",
			headers: {
				"Content-Type": "application/json",
			},
		});

		if (response.ok) {
			return await response.json();
		} else {
			console.error("Failed to get system status:", response.status, response.statusText);
			throw new Error(`Failed to get system status: ${response.statusText}`);
		}
	} catch (error) {
		console.error("Error getting system status:", error);
		throw error;
	}
};
