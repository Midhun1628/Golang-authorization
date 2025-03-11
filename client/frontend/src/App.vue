<template>
  <div id="app">
    <Navbar v-if="isLoggedIn" :currentMenu="selectedMenu" />
    <Sidebar v-if="isLoggedIn" @update-navbar="selectedMenu = $event" />
    <router-view></router-view>
    <Footer v-if="isLoggedIn">
      </Footer>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Sidebar from "./components/Sidebar.vue";
import Footer from "./components/Footer.vue";
import Login from "./components/Login.vue";

export default {
  components: {
    Navbar,
    Sidebar,
    Footer,
    Login,
  },
  data() {
    return {
      selectedMenu: "Dashboard",  // Default menu
    };
  },
  computed: {
    isLoggedIn() {
      return localStorage.getItem("access_token") !== null;
    },
  },
};
</script>
