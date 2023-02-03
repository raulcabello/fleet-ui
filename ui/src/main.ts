import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import {createPinia} from "pinia";
import axios from "axios";
import {configStore} from "@/store/config";

import "bootstrap/dist/css/bootstrap.min.css"
import "bootstrap"
import './assets/main.css'

const app = createApp(App)
const pinia = createPinia()

app.use(router)
app.use(pinia)
app.mount('#app')

const config = configStore()

axios.interceptors.request.use(function (request) {
    config.loading = true
    return request;
}, function (error) {
    config.loading = true
    return Promise.reject(error);
});

// Add a response interceptor
axios.interceptors.response.use(function (response) {
    config.loading = false
    return response;
}, function (error) {
    alert(error.message + '\n' + error.response.data)
    config.loading = false
    return Promise.reject(error);
});

