<template>
  <!-- <NavigationBar></NavigationBar> -->
  <div class="login-container">
    <div class="form-section">
      <div class="form-container">
        <div class="gym-logo">
          <img src="../images/icons-8-dumbbell-32-removebg-preview-10.png" alt="Gym Logo" />
        </div>
        <div class="error-message" v-if="error">{{ error }}</div>
        <div class="form-group">
          <input type="text" v-model="username" placeholder="Username" required />
        </div>
        <div class="form-group">
          <input type="password" v-model="password" placeholder="Password" required />
        </div>
        <div class="form-check">
          <label class="checkbox-container">
            <input type="checkbox" v-model="isChecked" class="checkmark" />
            <span class="checkmark-icon"></span>
            <span class="remember">Remember Me</span>
          </label>
        </div>
        <button class="btn-login" @click="handleLogin">LOGIN</button>
        <p class="sign-up">
          Don't have an account?
          <router-link to="/sign-up">Sign Up</router-link>
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
  
  name: "LoginPage",
  data() {
    return {
      username: "",
      password: "",
      isChecked: false,
      error: "",
    };
  },
  methods: {
    handleLogin() {
      if (!this.username || !this.password) {
        this.error = "Please enter both username and password.";
        return;
      }

      const loginPayload = {
        username: this.username,
        password: this.password,
      };

      axios.post(process.env.VUE_APP_LOGIN_USER_URL, loginPayload)
        .then(response => {
          // Handle success
          console.log("Login successful", response.data);
          this.error = "";

          // You might want to save the received token in local storage
          localStorage.setItem('token', response.data.token);

          // Redirect the user to another page after login
          this.$router.push('/');
        })
        .catch(error => {
          // Handle error
          if (error.response && error.response.data) {
            // Server responded with a status other than 2xx
            console.error("Login error:", error.response.data.message);
            this.error = error.response.data.message;
          } else {
            // Something happened in setting up the request that triggered an Error
            console.error("Login error:", error.message);
            this.error = "An error occurred. Please try again.";
          }
        });
    },
  },
};
</script>

<style scoped>
@import url(https://db.onlinewebfonts.com/c/0c28006f19928dfd146027cfd7024ca0?family=Poppins+Medium);
@import url(https://db.onlinewebfonts.com/c/07ecc0aa9ce268962dea7356eeff50a6?family=Poppins+Bold);

@font-face {
  font-family: "Poppins Medium";
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot");
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot?#iefix") format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff2") format("woff2"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff") format("woff"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.ttf") format("truetype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.svg#Poppins Medium") format("svg");
}

@font-face {
  font-family: "Poppins Bold";
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot");
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot?#iefix") format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff2") format("woff2"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff") format("woff"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.ttf") format("truetype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.svg#Poppins Bold") format("svg");
}

.login-container {
  display: flex;
  height: 100vh;
  background-color: #000;
  color: #fff;
}

.error-message {
  font-family: "Poppins Medium";
  font-size: 15px;
}

.form-section {
  width: 50%;
  display: flex;
  justify-content: center;
  align-items: center;
  background-color: #252525;
  font-family: "Poppins Medium";
}

.form-container {
  /* padding: 2rem; */
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
  /* Adjust size as needed */
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

.form-check {
  font-family: "Poppins Medium";
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 1.5rem;
  padding-left: 0%;
}

/* .checkbox-container {
  display: flex;
  align-items: center;
  color: #fff;
  font-size: 13px;
  white-space: nowrap;
  margin-bottom: 2.5rem;
} */

.remember {
  margin-left: 0.5rem;
}

/* .remember-me-label input[type="checkbox"] {
  margin-right: 0.5rem;
} */

/* .checkbox-container input:checked ~ .checkmark:after {
  content: "\2713";
  font-size: 16px;
  color: #000;
  position: relative;
  z-index: 2;

} */
.checkbox-container {
  display: inline-block;
  position: relative;
  padding-left: 30px;
  cursor: pointer;
  font-size: 15px;
  width: 3px;
  height: 3px;
  white-space: nowrap;
  margin-bottom: 3.5rem;
  left: 0;
}

.checkbox-container input {
  position: absolute;
  opacity: 0;
  cursor: pointer;
}

.checkmark-icon {
  position: absolute;
  top: 0;
  left: 0;
  height: 25px;
  width: 25px;
  background-color: #eee;
  border-radius: 3px;
}

.checkbox-container input:checked~.checkmark-icon {
  background-color: #000000;
}

.checkbox-container input:checked~.checkmark-icon:after {
  content: "\2705";
  font-size: 20px;
  color: white;
  position: absolute;
  top: 50%;
  left: 50%;
  transform: translate(-50%, -50%);
}

/* .forgot-password {
  color: #fff;
  text-decoration: none;
  font-size: 13px;
  margin-bottom: 2.5rem;
} */

.btn-login {
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
}

.btn-login:hover {
  background-color: #000000;
}

.sign-up {
  font-size: 13px;
}

.sign-up a {
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