<template>
  <!-- <NavigationBar></NavigationBar> -->
  <div class="signup-container">
    <div class="form-section">
      <div class="form-container">
        <div class="gym-logo">
          <img src="../images/icons-8-dumbbell-32-removebg-preview-10.png" alt="Gym Logo" />
        </div>
        <p class="sign-up-status" v-if="signupStatus">{{ signupStatus }}</p>
        <div class="form-group">
          <input type="text" v-model="username" placeholder="Username" required />
        </div>
        <div class="form-group">
          <input type="text" v-model="firstName" placeholder="First Name" required />
        </div>
        <div class="form-group">
          <input type="text" v-model="lastName" placeholder="Last Name" required />
        </div>
        <div class="form-group">
          <input type="text" v-model="telegramHandle" placeholder="Telegram Handle" required />
        </div>
        <!-- <div class="form-group"> -->
        <!-- <input type="text" v-model="roleId" placeholder="Role ID" required /> -->
        <!-- </div> -->
        <div class="form-group">
          <input type="email" v-model="email" placeholder="Email" required />
        </div>
        <div class="form-group">
          <input type="password" v-model="password" placeholder="Password" required />
        </div>
        <div class="form-group">
          <input type="password" v-model="confirmPassword" placeholder="Confirm password" required />
        </div>
        <button class="btn-signup" @click="handleSignup">SIGN UP</button>
        <p class="login-link">
          Already have an account? <router-link to="/login">Login</router-link>
        </p>
      </div>
    </div>
    <div class="image-section">
      <h1 class="gym_name">GYM DADDY</h1>
      <img src="../images/gym-daddy-character.png" alt="Gym Daddy Character" />
    </div>
  </div>
</template>

<script>

import axios from 'axios';

export default {
  
  name: "SignupPage",
  data() {
    return {
      email: "",
      password: "",
      confirmPassword: "",
      signupStatus: "",
      username: "",
      firstName: "",
      lastName: "",
      telegramHandle: "",
      roleId: 1, // Ensure you have a way to set this, e.g., based on user selection or a default value
    };
  },
  methods: {
    handleSignup() {
      // Implement your sign-up logic here
      console.log("Email:", this.email);
      console.log("Password:", this.password);
      console.log("Confirm Password:", this.confirmPassword);

      // Perform validation and submit the form data to the server


      // Example: Check if email and password already exist
      if (!this.email || !this.password || !this.confirmPassword) {
        this.signupStatus = "Please fill in all fields.";
        return;
      }
      if (this.password !== this.confirmPassword) {
        this.signupStatus = "Passwords do not match.";
        return;
      }
      const payload = {
        username: this.username,
        email: this.email,
        password: this.password,
        first_name: this.firstName,
        last_name: this.lastName,
        telegram_handle: this.telegramHandle,
        role_id: this.roleId,
      };
      axios.post(process.env.VUE_APP_REGISTER_USER_URL, payload)
        .then((response) => {
          console.log(response.data);
          this.signupStatus = "Sign up successful!";
          // Redirect to login page or another page
          this.$router.push("/login");
        })
        .catch((error) => {
          if (error.response && error.response.data) {
            // Now safely access error.response.data
            console.error("Signup error:", error.response.data.message);
            this.signupStatus = error.response.data.message || "An error occurred. Please try again.";
          } else {
            console.error("Signup error:", error.message);
            this.signupStatus = "An error occurred. Please try again.";
          }
        });

    },
  },
};
</script>

<style scoped>
/* Import the same font styles from the login page */

.signup-container {
  display: flex;
  height: 100vh;
  background-color: #000;
  color: #fff;
}

.form-section {
  width: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #252525;
  font-family: "Poppins Medium";
}

.sign-up-status {
  font-family: "Poppins Medium";
  font-size: 15px;
}

.form-container {
  border-radius: 10px;
  text-align: center;
  max-width: 500px;
  width: 100%;
}

.gym-logo {
  position: relative;
  margin-bottom: 2rem;
}

.gym-logo img {
  max-width: 200px;
  height: auto;
}

.form-group {
  margin-bottom: 1.5rem;
}

input {
  width: 100%;
  padding: 1rem;
  border-radius: 6px;
  border: none;
  background-color: #ffffff;
  color: #000;
}

.btn-signup {
  width: 100%;
  padding: 0.25rem;
  border-radius: 6px;
  border: none;
  background-color: #000000;
  color: #c7ff9c;
  font-weight: bold;
  cursor: pointer;
  transition: background-color 0.3s;
  font-family: "Poppins Bold";
  font-size: 25px;
  margin-bottom: 0.75rem;
  margin-top: 2rem;
}

.btn-signup:hover {
  background-color: #000000;
}

.login-link {
  font-size: 13px;
}

.login-link a {
  color: #c7ff9c;
  text-decoration: underline;
}

.image-section {
  width: 50%;
  display: flex;
  position: relative;
  justify-content: center;
  align-items: flex-end;
  background-image: url(../images/image-70.png);
}

.gym_name {
  position: absolute;
  top: 20%;
  left: 50%;
  transform: translateX(-50%);
  font-size: 5.5rem;
  color: #fff;
  text-shadow: 0 0 15px #c7ff9c;
  font-family: "Poppins Bold";
  white-space: nowrap;
}

.image-section img {
  max-width: 100%;
  height: 55vh;
}
</style>
