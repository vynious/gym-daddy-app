import { createApp } from "vue";
import App from "./App.vue";
import router from './router'

// this is for the ant design library
import Antd from "ant-design-vue";
import "ant-design-vue/dist/reset.css";


const app = createApp(App);

app.use(Antd)
app.use(router)

app.mount("#app");

