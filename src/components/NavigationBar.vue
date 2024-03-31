<style scoped>
@import url(https://db.onlinewebfonts.com/c/07ecc0aa9ce268962dea7356eeff50a6?family=Poppins+Bold);

@font-face {
  font-family: "Poppins Bold";
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot");
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot?#iefix")
      format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff2")
      format("woff2"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff")
      format("woff"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.ttf")
      format("truetype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.svg#Poppins Bold")
      format("svg");
}

.navbar {
  background-color: transparent;
  color: #fff;
  padding: 10px 20px;
  font-family: "Poppins Bold";
  position: fixed;
  width: 100%;
  z-index: 2;
}

.navbar-container {
  display: flex;
  justify-content: flex-end;
}

.navbar-menu {
  list-style-type: none;
  padding: 0;
  margin: 0;
  font-weight: bold;
  font-size: 20px;
}

.navbar-menu li {
  display: inline-block;
  margin-left: 10px;
}

.navbar-menu li:not(:last-child) {
  margin-right: 30px;
}

.navbar-menu a {
  text-decoration: none;
  color: white;
}

.navbar-logo {
  width: 75px;
  height: auto;
}
</style>

<template>
  <nav class="navbar">
    <img
      src="../images/icons-8-dumbbell-32-removebg-preview-10.png"
      alt="logo"
      class="navbar-logo"
    />
    <div class="navbar-container">
      <ul class="navbar-menu">
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/queue">Queue</router-link></li>
        <li><router-link to="/book">Classes</router-link></li>
        <li><router-link to="/profile">Profile</router-link></li>
        <li><router-link to="/profile">{{ loggedIn ? 'Profile' : 'Login' }}</router-link></li>
      </ul>
    </div>
  </nav>
</template>

<script>
export default {
  name: "NavigationBar",
  data() {
    return {
      loggedIn: false, // Assuming the user is not logged in initially
    };
  },
  created() {
    // Check if user is already logged in (you would replace this with your actual login check logic)
    const userLoggedIn = localStorage.getItem("loggedIn");
    this.loggedIn = userLoggedIn === "true";
  },
};
</script>

<!-- <script>
export default {
  name: "NavigationBar",
};
</script>
<template>
  <nav class="navbar">
    <div class="navbar-container">

      <img src="../assets/logo.png" alt="logo" class="navbar-logo">

      <ul class="navbar-menu">
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/queue">Queue</router-link></li>
        <li><router-link to='/book'>Book</router-link></li>
<<<<<<< HEAD
        <li><router-link to="/login">Login</router-link></li>
=======
        <li><router-link to='/profile'>Profile</router-link></li>
>>>>>>> 91dc67c5fbc2a4881e4ca381627ad5c20e8208aa
      </ul>
    </div>
  </nav>
</template>
  
<script>
  export default {
    name: 'NavigationBar'
  }
</script>

<style scoped>
  .navbar {
    background-color: black;
    color: #fff;
    padding: 10px 20px;
  }

  .navbar-container {
    display: flex;
    justify-content: space-between;
    align-items: center;
    
  }

  .navbar-menu {
    list-style-type: none;
    padding: 0;
    margin: 0;
    font-weight: bold;
    font-size: 28px;
    margin-left: auto;
  }

  .navbar-menu li {
    display: inline-block;
  }

  .navbar-menu li:not(:last-child) {
    margin-right: 30px;
  }

  .navbar-menu a {
    text-decoration: none;
    color: white;
  }

  .navbar-logo {
    width: 75px;
    height: auto;
  }
</style> -->
