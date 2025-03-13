<template>
  <div class="card-body">
    <!-- Show Add User button only if role is Super Admin -->
    <button v-if="userRole === 'SuperAdmin'" class="btn btn-primary" @click="openAddUser">Add User</button>
  </div>
  <div class="card">
    <div class="card-body">
      <table class="table table-bordered">
        <thead>
          <tr>
            <th v-for="column in projectColumns" :key="column.field">
              {{ column.title }}
            </th>
            <th>Actions</th>
          </tr>
        </thead>
        <tbody>
          <tr v-for="user in userData" :key="user.user_id">
            <td>{{ user.user_id }}</td>
            <td>{{ user.username }}</td>
            <td>{{ user.email }}</td>
            <td>{{ user.position }}</td>
            <td>
              <!-- Show Edit & Delete buttons for Super Admin and Admin -->
              <button
                v-if="userRole === 'SuperAdmin' || userRole === 'Admin'"
                class="btn btn-sm btn-warning"
                @click="openEditUser(user)"
              >
                <i class="fas fa-edit"></i>
              </button>
              <button
                v-if="userRole === 'SuperAdmin' || userRole === 'Admin'"
                class="btn btn-sm btn-danger"
                @click="deleteUser(user.user_id)"
              >
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>
      <UserForm
  v-if="isModalOpen" 
  :user="selectedUser" 
  :isEditing="isEditing" 
  @userUpdated="fetchUsers" 
  @close="isModalOpen = false"
/>
    </div>
  </div>
</template>

<script>
import UserForm from "./Userform.vue";
import axios from "/axios";



export default {
  name: "UserTable",
  components: { UserForm },
  data() {
    return {
      userData: [],
      projectColumns: [
        { title: "ID", field: "user_id" },
        { title: "USERNAME", field: "username" },
        { title: "EMAIL", field: "email" },
        { title: "Position", field: "position" },
      ],
      isModalOpen: false,
      isEditing: false,
      selectedUser: null,
      userRole: "", // Store the logged-in user's role
    };
  },
  methods: {



    async fetchUsers() {
      try {
        const res = await axios.get("/users");
        this.userData = res.data;
      } catch (error) {
        console.error("Error fetching users:", error);
      }
    },
    openAddUser() {
      this.isEditing = false;
      this.selectedUser = null;
      this.isModalOpen = true;
    },
    openEditUser(user) {
      this.isEditing = true;
      this.selectedUser = { ...user };
      this.isModalOpen = true;
    },
    async deleteUser(userId) {
      if (confirm("Are you sure you want to delete this user?")) {
        try {
          await axios.delete(`/users/${userId}`);
          this.fetchUsers();
        } catch (error) {
          console.error("Error deleting user:", error);
        }
      }
    },
    getUserRole() {
      // Assuming role is stored in localStorage after login
      this.userRole = localStorage.getItem("job_title") || ""; 
    },
  },
  mounted() {
    this.fetchUsers();
    this.getUserRole();
  },
};
</script>

<style scoped>
.card-body {
  margin: 30px;
}
.card {
  margin: 50px;
}
.table {
  width: 100%;
  text-align: left;
}
.btn {
  margin-right: 5px;
}
</style>
