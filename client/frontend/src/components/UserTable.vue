<template>
  <div class="card">
    <div class="card-body">
      <button class="btn btn-primary" @click="showAddContact">Add Contact</button>
    </div>
    <div class="card-body">
      <bootstrap-table
        :columns="projectColumns"
        :data="projectData"
        data-pagination="true"
      ></bootstrap-table>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      projectData: [],
      projectColumns: [
        { title: "ID", field: "user_id", sortable: true },
        { title: "USERNAME", field: "username", sortable: true },
        { title: "EMAIL", field: "email", sortable: true },
        { title: "Position", field: "position", sortable: true },
      ],
    };
  },
  methods: {
    fetchProjects() {
      axios.get("http://localhost:3000/api/users").then((res) => {
        this.projectData = res.data.USERNAME;
      });
    },
  },
  mounted() {
    this.fetchProjects();
  },
};
</script>
