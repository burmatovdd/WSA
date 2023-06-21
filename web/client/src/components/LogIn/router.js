import {createRouter, createWebHistory} from 'vue-router';
import Login from "./login.vue";


export default  createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Login },
    ],
})