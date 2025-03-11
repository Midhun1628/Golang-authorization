<template>
  <div class="card">
    <div class="card-body">
      <button class="btn btn-primary" @click="showAddContact">Add Contact</button>
    </div>
    <div class="card-body">
      <table class="table table-bordered">
        <thead>
          <tr>
            <th v-for="column in projectColumns" :key="column.field">
              {{ column.title }}
            </th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in projectData" :key="user.user_id">
            <td>{{ user.user_id }}</td>
            <td>{{ user.username }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.position }}</td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "UserTable",
  data() {
    return {
      projectData: [],
      projectColumns: [
        { title: "ID", field: "user_id" },
        { title: "USERNAME", field: "username" },
        { title: "EMAIL", field: "email" },
        { title: "Position", field: "position" },
      ],
    };
  },
  methods: {
    async fetchProjects() {
      try {
        const res = await axios.get("http://localhost:3000/api/users");
        this.projectData = res.data; // Ensure the API response matches this structure
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    },
  },
  mounted() {
    this.fetchProjects();
  },
};
</script>

<style scoped>
.card {
  margin: 20px;
}
.table {
  width: 100%;
  text-align: left;
}
</style>
