import { createApp } from "vue";
import App from "./App.vue";
import router from "./router";
import store from "./store";

var ip = require("ip");
console.dir ( ip.address() );

const app = createApp(App)
//app.config.globalProperties.$serverAddr = 'http://localhost:8080'
app.config.globalProperties.$serverAddr = 'http://192.168.1.119:8080'
app.use(store).use(router).mount("#app");
