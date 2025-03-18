<template>
  <div id="app">
    <!-- Navbar and Sidebar appear only when logged in -->
    <Navbar v-if="isLoggedIn" :currentMenu="selectedSidebar" />
    <Sidebar v-if="isLoggedIn" @update-navbar="selectedSidebar = $event" />

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
      selectedSidebar: "Dashboard",  // Default menu name
      isLoggedIn: localStorage.getItem("access_token") !== null, // Initial check
    };
  },

 
  watch: {
    $route() {
      // Automatically check login state on route change
      this.isLoggedIn = localStorage.getItem("access_token") !== null;
    },
  },
  methods: {
    syncLoginState() {
      this.isLoggedIn = localStorage.getItem("access_token") !== null;
    },
   
  },
};
</script>

<style>
/* Default style when sidebar is hidden */
main {
  padding: px;
  margin-left: 0; /* No margin when sidebar is hidden */
}

/* Apply left margin only when the user is logged in */
main.with-sidebar {
  margin-left: 230px; /* Adjust based on sidebar width */
}
</style>
