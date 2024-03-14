import { createApp } from "vue";
import App from "./App.vue";
import router from './router'
import 'bootstrap/dist/css/bootstrap.min.css'

// import Antd from "ant-design-vue";
// import "ant-design-vue/dist/reset.css";



const app = createApp(App);
// app.use(Antd).mount("#app");

app.use(router)

app.mount("#app");

