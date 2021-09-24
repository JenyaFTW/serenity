import { createApp } from "vue";
import App from "./App.vue";
import Router from "./router";
import "./scss/global.scss";

const app = createApp(App);

app.use(Router);

app.mount("#app");
