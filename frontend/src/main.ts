import { createApp } from "vue";
import "./style.css";
import App from "./App.vue";
import wsService from "./services/websocket";
import router from "./router";

// Initialize WebSocket connection
wsService.connect();

const app = createApp(App);
app.use(router);
app.mount("#app");
