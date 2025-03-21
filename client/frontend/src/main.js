import { createApp } from "vue";
import App from "./App.vue";
import router from "./routers/router"; // If using Vue Router
import "bootstrap/dist/css/bootstrap.min.css";
import "@fortawesome/fontawesome-free/css/all.css";
// import store from "./store";
import "bootstrap";
import axios from "axios";
import VueCookies from "vue-cookies";

const app = createApp(App);
app.use(router); // If using Vue Router
// app.use(store);
app.use(VueCookies);
app.config.globalProperties.$axios = axios;
app.mount("#app");
