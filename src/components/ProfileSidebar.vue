<template>
    <div class="sidebar" style="position: relative;">
        <img src="../assets/edit.png" alt="" style="position: absolute; top:20px; right: 40px;">
        <br>
        <img src="../assets/avatar.png" alt="">
        <h1>USER 12345</h1>
        <hr>
        <router-link to="/profileEditInfo"><button class="button" id="editInfo">Edit Info</button></router-link>
        <button class="button">My Bookings</button>
        <router-link to="/profileRefer"><button class="button" id="refer">Refer a Friend</button></router-link>
        <button class="button">Logout</button>
    </div>
    
</template>
    
  <script>
    export default {
      name: 'ProfileSidebar'
    }
  </script>
    
  <style scoped>
    .sidebar {
        width: 30% ;
        height: 100vh;
        overflow-y: auto;
        background-color: black;
    }
    h1 {
        color: white;
        font-family: 'Poppins';
        margin-top: 15px;
        margin-bottom: 30px;
    }
    hr {
        width: 80% ;
    }
    li {
        color: white;
        list-style: none;
    }
    .button {
        background-color: black;
        color: white;
        padding: 15px 40px;
        margin-top: 14px;
        margin-bottom: 14px;
        font-size: 28px ;
        font-family: 'Poppins';
        border: 0px;
        width: 70%;
        border-radius: 12px;
        cursor: pointer;
    }
    .button:hover {
        background-color: #C7FF9C20;
        color: white;
    }
    /* .button:focus {
        background-color: #C7FF9C20;
        color: white;
    } */
  </style>