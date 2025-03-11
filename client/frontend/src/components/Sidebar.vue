<template>
  <aside class="main-sidebar sidebar-dark-primary elevation-4">
    <a href="#" class="brand-link">
      <img src="/dist/img/AdminLTELogo.png" alt="Logo" class="brand-image img-circle elevation-3">
      <span class="brand-text font-weight-light">User Authorization</span>
    </a>

    <div class="sidebar">
      <nav class="mt-2">
        <ul class="nav nav-pills nav-sidebar flex-column">
          <!-- Dashboard Menu -->
          <SidebarItem name="Dashboard" icon="fas fa-bars" @update-navbar="updateNavbarText" />

          <!-- Manage Users Menu -->
          <SidebarItem name="Manage Users" icon="fas fa-user" @update-navbar="updateNavbarText" />
          <SidebarItem name="Logout"  icon="fas fa-sign-out-alt" @update-navbar="updateNavbarText"  />
          <SidebarItem name="klogjn"  icon="fas fa-sign-out-alt" @update-navbar="updateNavbarText"  />


        </ul>
      </nav>
    </div>
  </aside>
</template>

<script>
import SidebarItem from "./SidebarItem.vue";

export default {
  components: {
    SidebarItem,
  },
  methods: {
    updateNavbarText(newText) {
      this.$emit("update-navbar", newText); // Emit event to App.vue
    },
  },
};
</script>

<style scoped>
/* Custom styles if needed */
</style>
