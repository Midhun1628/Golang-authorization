<!-- 

  <script>
  import axios from "/axios";
  
  export default {
    props: {
      isOpen: Boolean,
      userData: Object,
      isEditing: Boolean,
    },
    data() {
      return {
        userForm: {
          user_id: "",
          username: "",
          email: "",
          password: "",
          role: "",
        },
      };
    },
    watch: {
      userData: {
        handler(newUser) {
          if (newUser) {
            this.userForm = { ...newUser }; // Copy existing data if editing
          }
        },
        immediate: true,
      },
    },
    methods: {
      async submitForm() {
        try {
          if (this.isEditing) {
            await axios.put(`/users/${this.userForm.user_id}`, this.userForm);
          } else {
            await axios.post("/register", this.userForm);
          }
          this.$emit("refreshData"); // Emit event to update parent component
          this.closeModal();
        } catch (error) {
          console.error("Error saving user:", error);
        }
      },
      closeModal() {
        this.$emit("close");
      },
    },
  };
  </script>
  
<template>
  <div v-if="isOpen" class="modal-overlay">
    <div class="modal">
      <h3>{{ isEditing ? "Edit User" : "Add User" }}</h3>
      <div class="form-group">
        <label>ID:</label>
        <input v-model="userForm.user_id" type="text" class="form-control" :disabled="isEditing" />
      </div>
      <div class="form-group">
        <label>Username:</label>
        <input v-model="userForm.username" type="text" class="form-control" />
      </div>
      <div class="form-group">
        <label>Email:</label>
        <input v-model="userForm.email" type="email" class="form-control" />
      </div>
      <div class="form-group">
        <label>Password:</label>
        <input v-model="userForm.password" type="password" class="form-control" :disabled="isEditing" />
      </div>
      <div class="form-group">
        <label>Role:</label>
        <select v-model="userForm.role" class="form-control">
          <option>Super Admin</option>
          <option>Admin</option>
          <option>Front Office</option>
        </select>
      </div>
      <div class="modal-actions">
        <button class="btn btn-primary" @click="submitForm">{{ isEditing ? "Update" : "Add" }} User</button>
        <button class="btn btn-secondary" @click="closeModal">Cancel</button>
      </div>
    </div>
  </div>
</template>
  
  <style scoped>
  .modal-overlay {
    position: fixed;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    background: rgba(0, 0, 0, 0.5);
    display: flex;
    justify-content: center;
    align-items: center;
  }
  .modal {
    background: white;
    padding: 20px;
    border-radius: 10px;
    width: 400px;
  }
  .form-group {
    margin-bottom: 10px;
  }
  .modal-actions {
    display: flex;
    justify-content: space-between;
  }
  </style> -->