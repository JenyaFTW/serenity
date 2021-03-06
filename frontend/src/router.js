import { createRouter, createWebHistory } from 'vue-router';
import Login from './views/Login.vue';
import SignUp from './views/SignUp.vue';
import Home from './views/Home.vue';

const routes = [
    {
        path: '/',
        name: 'Home',
        component: Home
    },
    {
        path: '/login',
        name: 'Login',
        component: Login
    },
    {
        path: '/signup',
        name: 'Signup',
        component: SignUp
    }
];

const Router = createRouter({
    history: createWebHistory(),
    routes
});

export default Router;