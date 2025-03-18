<template>
  <div class="card-body">
    <!-- Show Add User button only if the user has "Create" permission -->
    <button v-if="hasPermission('Create')" class="btn btn-primary" @click="openAddUser">Add User</button>
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
            <td>{{ user.ID }}</td>
            <td>{{ user.Username }}</td>
            <td>{{ user.Email }}</td>
            <td>{{ user.Position }}</td>
            <td>
              <!-- Show Edit button only if the user has "Update" permission -->
              <button
                v-if="hasPermission('Update')"
                class="btn btn-sm btn-warning"
                @click="openEditUser(user)"
              >
                <i class="fas fa-edit"></i>
              </button>
              <!-- Show Delete button only if the user has "Delete" permission -->
              <button
                v-if="hasPermission('Delete')"
                class="btn btn-sm btn-danger"
                @click="deleteUser(user.ID)"
              >
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
export default {
  data() {
    return {
      userData: [],
      userPermissions: [],
    };
  },
  methods: {
    hasPermission(permission) {
      return this.userPermissions.includes(permission);
    },
    fetchPermissions() {
      this.userPermissions = JSON.parse(localStorage.getItem("permissions")) || [];
    },
  },
  mounted() {
    this.fetchPermissions();
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
