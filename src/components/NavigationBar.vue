<template>
  <nav class="navbar">
    <div class="navbar-container">

      <img src="../assets/logo.png" alt="logo" class="navbar-logo">

      <ul class="navbar-menu">
        <li><router-link to="/">Home</router-link></li>
        <li><router-link to="/queue">Queue</router-link></li>
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
  }

  .navbar-menu li {
    display: inline-block;
  }

  .navbar-menu li:not(:last-child) {
    margin-right: 30px;
  }

  .navbar-menu {
    margin-left: auto;
  }

  .navbar-logo {
    width: 75px;
    height: auto;
  }
</style>
  