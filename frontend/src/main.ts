import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import wsService from "./services/websocket";

// Initialize WebSocket connection
wsService.connect();

const app = createApp(App);
app.mount("#app");
