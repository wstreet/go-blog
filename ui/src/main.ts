import { createApp } from "vue";
import Antd from 'ant-design-vue'
import App from "./App.vue";
import router from "./router";
import "./index.css";
import 'ant-design-vue/dist/antd.css';

// import "./assets/main.css";

const app = createApp(App);

app.use(router);
app.use(Antd)

app.mount("#app");
