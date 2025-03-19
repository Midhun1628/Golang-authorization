<template>
  <div class="card-body">
    <!-- Show Add User button only if role has "Create" permission -->
    <button v-if="userPermissions.includes('Create')" class="btn btn-primary" @click="openAddUser">
      Add User
    </button>
  </div>

  <div class="card">
    <div class="card-body">
      <table class="table table-bordered">
        <thead>
          <tr>
            <th v-for="column in projectColumns" :key="column.field">
              {{ column.title }}
            </th>
            <th v-if="userPermissions.includes('Update')">Edit</th>
            <th v-if="userPermissions.includes('Delete')">Delete</th>

          </tr>
        </thead>
        <tbody>
          <tr v-for="user in userData" :key="user.ID">
            <td>{{ user.ID }}</td>
            <td>{{ user.Username }}</td>
            <td>{{ user.Email }}</td>
            <td>{{ user.Position }}</td>
            <td v-if="userPermissions.includes('Update') ">
              <button v-if="userPermissions.includes('Update')" class="btn btn-sm btn-warning" @click="openEditUser(user)">
                <i class="fas fa-edit"></i>
              </button>
              </td>
            <td v-if="userPermissions.includes('Delete') ">
             <button v-if="userPermissions.includes('Delete')" class="btn btn-sm btn-danger" @click="deleteUser(user.ID)">
                <i class="fas fa-trash-alt"></i>
              </button>
            </td>
          </tr>
        </tbody>
      </table>

      <UserForm v-if="isModalOpen" :user="selectedUser" :isEditing="isEditing" @userUpdated="fetchUsers" @close="isModalOpen = false" />
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
        { title: "ID", field: "ID" },
        { title: "USERNAME", field: "Username" },
        { title: "EMAIL", field: "Email" },
        { title: "Position", field: "EmployeePosition" },
      ],
      isModalOpen: false,
      isEditing: false,
      selectedUser: null,
      userRole: "",
      userPermissions: [], // Store permissions of logged-in user
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
    getUserPermissions() {
      // Fetch permissions from localStorage after login
      const permissions = JSON.parse(localStorage.getItem("permissions")) || [];
      this.userPermissions = permissions;
    },
  },
  mounted() {
    this.fetchUsers();
    this.getUserPermissions();
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