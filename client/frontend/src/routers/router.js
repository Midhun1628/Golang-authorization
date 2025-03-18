import { createRouter, createWebHistory } from "vue-router";
// import Login from "@/components/Login.vue";
// import Dashboard from "@/components/Dashboard.vue";
import Login from "../components/Login.vue";
import Dashboard from "../components/Dashboard.vue";
import UserTable from "../components/UserTable.vue";


const routes = [
  { path: "/login", component: Login ,
  beforeEnter:(to, from, next) => {
    const token = localStorage.getItem("access_token");
    if (token) {
      next("/"); // Redirect to dashboard if logged in
    } else {
      next(); // Otherwise, show login page
    }
  },
},
  {
    path: "/",
    component: Dashboard,
    
    beforeEnter: (to, from, next) => {
      const token = localStorage.getItem("access_token");
      if (!token) {
        next("/login");
      } else {
        next();  // stays in the same route which is currently in, (next() when you just want to stay on the current page without forcing a reload.)
      }
    },
  },
  {
    path: "/user-table",
    name: "UserTable",
    component: UserTable,
    beforeEnter: (to, from, next) => {
      let token = localStorage.getItem("access_token");

      if (!token) {
        next("/login");// Redirect to login page
      } else {
        next()
      }},
  },{
    path: "/logout",
    name: "Logout",
    beforeEnter: (to, from, next) => {
      localStorage.clear()    // Clear token on logout
      next("/login"); // Redirect to login page
    },
  },
];

const router = createRouter({
  history: createWebHistory(),
  routes,
});

export default router;
