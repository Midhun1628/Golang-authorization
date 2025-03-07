import { createRouter, createWebHistory } from "vue-router";
import Dashboard from "@/views/Dashboard.vue";
import ManageUser from "@/views/ManageUser.vue";

const routes = [
  { path: "/", component: Dashboard },
  { path: "/manage-user", component: ManageUser },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
