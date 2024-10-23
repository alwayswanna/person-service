import { createRouter, createWebHistory } from 'vue-router';

const routes = [
  {
    path: '/',
    name: 'HomeView',
    component: () => import('@/components/HomeView.vue'),
  },
  {
    path: '/callback',
    name: 'Callback',
    component: () => import('@/components/auth/Callback.vue'),
  },
  {
    path: '/users',
    name: 'UsersView',
    component: () => import('@/components/Users.vue'),
  },
  {
    path: '/information',
    name: 'InformationView',
    component: () => import('@/components/Information.vue'),
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
