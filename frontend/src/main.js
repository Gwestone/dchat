import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

const app = createApp(App)
//app.config.globalProperties.$serverAddr = 'http://localhost:8080'
app.config.globalProperties.$serverAddr = 'https://1992-185-222-23-142.eu.ngrok.io'
app.use(store).use(router).mount("#app");
