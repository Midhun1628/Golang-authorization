import { createRouter, createWebHistory } from "vue-router";
// import Login from "@/components/Login.vue";
// import Dashboard from "@/components/Dashboard.vue";
import Login from "../components/Login.vue";
import Dashboard from "../components/Dashboard.vue";
import UserTable from "../components/UserTable.vue";


const routes = [
  { path: "/", component: Login },
  {
    path: "/dashboard",
    component: Dashboard,
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("access_token");
      if (!token) {
        next("/");
      } else {
        next();
      }
    },
  },
  // {
  //   path: "/manage-user",
  //   name: "UserTable",
  //   component: UserTable,
  // },{
  //   path: "/logout",
  //   name: "Logout",
  //   beforeEnter: (to, from, next) => {
  //     localStorage.removeItem("access_token"); // Clear token on logout
  //     next("/login"); // Redirect to login page
  //   },
  // },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
