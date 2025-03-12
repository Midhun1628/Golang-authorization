<template>
   <div class="card-body">
      <button class="btn btn-primary" @click="showAddContact">Add Contact</button>
    </div>
  <div class="card">
    <div class="card-body">
      <table class="table table-bordered">
        <thead>
          <tr>
            <th v-for="column in projectColumns" :key="column.field">
              {{ column.title }}
            </th>
            <th>Actions</th> <!-- New column for Edit/Delete -->
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in userData" :key="user.user_id">
            <td>{{ user.user_id }}</td>
            <td>{{ user.username }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.position }}</td>
            <td>
              <!-- Edit Button -->
              <button class="btn btn-sm btn-warning" @click="editUser(user.user_id)">
                <i class="fas fa-edit"></i>
              </button>
              <!-- Delete Button -->
              <button class="btn btn-sm btn-danger" @click="deleteUser(user.user_id)">
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
    </div>
  </div>
</template>

<script>
import axios from "/axios"; // Ensure correct import path

export default {
  name: "UserTable",
  data() {
    return {
      userData: [],
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
        const res = await axios.get("/users");
        this.userData = res.data;
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    },
   async editUser(user) {
      console.log("Editing user:", user);
      // Implement your edit logic here
    },
    async deleteUser(userId) {
      if (confirm("Are you sure you want to delete this user?")) {
        try {
          await axios.delete(`/users/${userId}`);
          this.fetchProjects(); // Refresh user list after deletion
        } catch (error) {
          console.error("Error deleting user:", error);
        }
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
.btn {
  margin-right: 5px;
}
</style>
