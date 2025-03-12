<template>
  <div id="app">
    <Navbar v-if="isLoggedIn" :currentMenu="selectedMenu" />
    <Sidebar v-if="isLoggedIn" @update-navbar="selectedMenu = $event" />
    
    <main :class="{ 'with-sidebar': isLoggedIn }">
      <router-view></router-view> <!-- Dynamically load the content here -->
    </main>

    <Footer v-if="isLoggedIn"></Footer>
  </div>
</template>

<script>
import Navbar from "./components/Navbar.vue";
import Sidebar from "./components/Sidebar.vue";
import Footer from "./components/Footer.vue";
import "@fortawesome/fontawesome-free/css/all.css";

export default {
  components: {
    Navbar,
    Sidebar,
    Footer,
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

<style>
/* Default style when sidebar is hidden */
main {
  padding: 20px;
  margin-left: 0; /* No margin when sidebar is hidden */
}

/* Apply left margin only when the user is logged in */
main.with-sidebar {
  margin-left: 250px; /* Adjust based on sidebar width */
}
</style>
