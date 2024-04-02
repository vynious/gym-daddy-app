<template>
    <div class="default">
        <div class="logout">
          <div>
            <p>Proceed with Logout?</p>
            <br>
            <router-link to="/profile"><button id="choice1">Cancel</button></router-link>
            <!-- <router-link to="/loginpage">  -->
                <button id="choice2">Confirm</button>
            <!-- </router-link> -->
          </div>
        </div>
    </div>
</template>
    
  <script>
    export default {
      name: 'ProfileLogout'
    }
  </script>
    
  <style scoped>
    .default {
        background: black;
        height: 100vh;
        width: 70%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .logout {
        width: 90% ;
        height: 90vh;
        overflow-y: auto;
        background-color: #D9D9D920;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 28px;
    }
    p {
      margin: 0px;
    }
    #choice1 {
      font-size: 25px;
      /* font-family: 'Poppins'; */
      color: #C7FF9C;
      background-color: black;      
      padding: 15px 40px;
      border: 0px;
      border-radius: 12px;
      cursor: pointer;
      margin-right: 50px ;
    }
    #choice1:hover {
        background-color: #FFFFFF70;
    }
    #choice2 {
      font-size: 25px;
      /* font-family: 'Poppins'; */
      color: #C7FF9C;
      background-color: #D9D9D920;      
      padding: 15px 40px;
      border: 0px;
      border-radius: 12px;
      cursor: pointer;
    }
    #choice2:hover {
        background-color: #FFFFFF70;
    }
  </style>