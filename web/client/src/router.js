import {createRouter, createWebHistory} from 'vue-router';
import auth from "./components/Auth/auth.vue";
import main from "./components/Pages/Main/main.vue";
import reports from "./components/Pages/Reports/reports.vue";
import analytics from "./components/Pages/Analytics/analytics.vue";
import statistic from "./components/Pages/Statistic/statistic.vue";



export default  createRouter({
    history: createWebHistory(),
    routes: [
        { path: '/', component: auth },
        { path: '/dashboard', component: main },
        { path: '/reports', component: reports },
        { path: '/analytics', component: analytics },
        { path: '/statistic', component: statistic },
    ],
})