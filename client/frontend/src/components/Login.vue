<script setup>
import { ref } from "vue";
import api from "/axios";
import { useRouter } from "vue-router";

const email = ref("midhun@gmail.com");
const password = ref("web");
const errorMessage = ref("");
const router = useRouter();

const login = async () => {
  errorMessage.value = "";

  try {
    const response = await api.post("/login", {
      email: email.value,
      password: password.value,
    });
    if (response.data.access_token) {

      // Store access token & user details
      localStorage.setItem("access_token", response.data.access_token);
      localStorage.setItem("username", response.data.username);
      localStorage.setItem("job_title", response.data.job_title);
      router.push("/dashboard"); // Redirect after successful login
    }
  } catch (error) {
    if (error.response) {
      // Capture API error messages
      errorMessage.value = error.response.data.error || "Login failed";
    } else {
      // Network or unexpected errors
      errorMessage.value = "An error occurred. Please try again.";
    }
  }
};
</script>


<template>
  <div class="login-container">
    <div class="login-box">
      <h2>Login</h2>
      <form @submit.prevent="login">
        <div class="input-group">
          <label>Email:</label>
          <input v-model="email" type="email" required placeholder="Enter your email"  />
        </div>
        <div class="input-group">
          <label>Password:</label>
          <input v-model="password" type="password" required placeholder="Enter your password" />
        </div>
        <button type="submit">Login</button>
      </form>
      <p v-if="errorMessage" class="error">{{ errorMessage }}</p>
    </div>
  </div>
</template>

<style scoped>
.login-container {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: linear-gradient(to right, #4facfe, #00f2fe);
}

.login-box {
  background: rgb(176, 185, 199);
  padding: 30px;
  border-radius: 20px;
  box-shadow: 0 4px 10px rgba(0, 0, 0, 0.1);
  text-align: center;
  width: 350px;
}

h2 {
  margin-bottom: 20px;
  color: #333;
}

.input-group {
  margin-bottom: 25px;
  text-align: left;
}

label {
  display: block;
  margin-bottom: 5px;
  font-weight: bold;
  color: #555;
}

input {
  width: 100%;
  padding: 10px;
  border: 1px solid #ccc;
  border-radius: 5px;
  font-size: 16px;
}

button {
  width: 100%;
  padding: 10px;
  background: #007bff;
  color: white;
  border: none;
  border-radius: 5px;
  font-size: 18px;
  cursor: pointer;

}

button:hover {
  background: #0056b3;
}

.error {
  color: red;
  margin-top: 10px;
}
</style>
