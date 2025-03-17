import axios from "axios";

const api = axios.create({
  baseURL: "http://localhost:3000",
  withCredentials: true, // Ensures cookies (refresh token) are sent
});

// Request Interceptor: Attach access token to requests
api.interceptors.request.use((config) => {
  const token = localStorage.getItem("access_token");
  
  if (token) {
    config.headers["Authorization"] = `Bearer ${token}`;
  }
  return config;
});

// Flag to prevent multiple redirects
let isRefreshing = false;

// Response Interceptor: Handle expired tokens
api.interceptors.response.use(
  (response) => response,
  async (error) => {
    const originalRequest = error.config;

    if (error.response?.status === 401 && !originalRequest._retry) {
      if (isRefreshing) {
        return Promise.reject(error); // Prevent multiple refresh attempts
      }
      originalRequest._retry = true;
      isRefreshing = true;

      try {
        // Attempt to refresh the token
        const res = await axios.post("http://localhost:3000/refresh", {}, { withCredentials: true });

        localStorage.setItem("access_token", res.data.access_token);

        // Retry the failed request with new access token
        originalRequest.headers["Authorization"] = `Bearer ${res.data.access_token}`;
        isRefreshing = false; // Reset flag
        return api.request(originalRequest);
      } catch (refreshError) {
        console.error("Session expired. Please log in again.");
        localStorage.removeItem("access_token");
        localStorage.removeItem("username");
        localStorage.removeItem("job_title");
        isRefreshing = false; // Reset flag

        // Redirect only if already on a protected route
        if (window.location.pathname !== "/") {
          window.location.href = "/"; // Redirect to login page
        }
      }
    }
    return Promise.reject(error);
  }
);

export default api;
