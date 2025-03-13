<template>
    <div class="modal fade show" id="modal-lg" style="display: block; padding-right: 15px;" aria-modal="true" role="dialog">
      <div class="modal-dialog modal-lg">
        <div class="modal-content">
          <div class="modal-header">
            <h4 class="modal-title">{{ isEditing ? "Edit User" : "Add User" }}</h4>
            <button type="button" class="close" @click="closeModal" aria-label="Close">
             
            </button>
          </div>
  
          <div class="modal-body">
            <form @submit.prevent="handleUserSubmit">
              <div class="form-group">
                <label for="userId">User ID</label>
                <input type="text" id="userId" class="form-control" v-model="userData.user_id" :disabled="isEditing" required />
              </div>
  
              <div class="form-group">
                <label for="userName">Name</label>
                <input type="text" id="userName" class="form-control" v-model="userData.username" required />
              </div>
  
              <div class="form-group">
                <label for="userEmail">Email</label>
                <input type="email" id="userEmail" class="form-control" v-model="userData.email" required />
              </div>
  
              <div class="form-group">
                <label for="userPosition">Position</label>
                <select id="userPosition" class="form-control" v-model="userData.position" required>
                  <option value="Super Admin">Super Admin</option>
                  <option value="Admin">Admin</option>
                  <option value="Front Office">Front Office</option>
                </select>
              </div>
            </form>
          </div>
  
          <div class="modal-footer justify-content-between">
            <button type="button" class="btn btn-default" @click="closeModal">Close</button>
            <button type="button" class="btn btn-primary" @click="handleUserSubmit">
              {{ isEditing ? "Update User" : "Save User" }}
            </button>
          </div>
        </div>
      </div>
    </div>
  </template>
  
  <script>
  import api from "/axios";
  
  export default {
    props: {
      user: Object,
      isEditing: Boolean,
    },
    data() {
      return {
        userData: {
          user_id: "",
          username: "",
          email: "",
          position: "",
        },
      };
    },
    watch: {
      user: {
        immediate: true,
        handler(newUser) {
          if (newUser) {
            this.userData = { ...newUser };
          }
        },
      },
    },
    methods: {
      async handleUserSubmit() {
        try {
          if (this.isEditing) {
            await api.patch(`/users/${this.userData.user_id}`, this.userData);
          } else {
            await api.post("/register", this.userData);
          }
          this.$emit("userUpdated");
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
  
  <style >
  .modal {
    background: rgba(0, 0, 0, 0.5);
  }
  </style>
  