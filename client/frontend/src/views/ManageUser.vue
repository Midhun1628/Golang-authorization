<template>
    <div class="container-fluid">
      <h5>Manage Project</h5>
      <button class="btn btn-primary" @click="showAddContact">Add Contact</button>
      <DataTable />
    </div>
  </template>
  
  <script>
  import DataTable from "../components/UserTable.vue";
  
  export default {
    components: { DataTable },
    methods: {
      showAddContact() {
        alert("Add Contact Clicked!");
      },
    },
  };
  </script>
  