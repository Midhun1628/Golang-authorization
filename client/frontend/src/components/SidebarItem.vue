<template>
    <li class="nav-item">
      <a href="#" class="nav-link" @click="updateNavbar">
        <i class="nav-icon" :class="icon"></i>
        <p>{{ name }}</p>
      </a>
    </li>
  </template>
  
  <script>
  export default {
    props: {
      name: String,  // Menu name (Dashboard, Manage Users)
      icon: String,  // FontAwesome icon class
    },
    methods: {
      updateNavbar() {
        this.$emit("update-navbar", this.name); // Emit event to parent here passing the data from child to parent (Sidebar.vue)
      },
    },
  };
  </script>
  
  <style scoped>
  /* Custom styles if needed */
  </style>
  