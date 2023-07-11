import {createRouter, createWebHistory} from 'vue-router';
import Login from "./login.vue";
import Dashboard from "../Dashboard/dashboard.vue";


export default  createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: Login },
        { path: '/dashboard', component: Dashboard },
    ],
})