<template>
  <div id="app">
    <Navbar />
    <Sidebar />
    <router-view></router-view> <!-- This will change dynamically -->
    <Footer>
    </Footer>

  </div>
</template>

<script>
import Navbar from "@/components/Navbar.vue";
import Sidebar from "@/components/Sidebar.vue";
import Footer from "@/components/Footer.vue";

export default {
  components: {
    Navbar,
    Sidebar,
    Footer,
  },
};
</script>

<style>
/* Global styles */
</style>
