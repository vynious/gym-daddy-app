<template>
    <div class="default">
        <div class="welcome">
            <p>Welcome to Your Profile</p>
        </div>
    </div>
</template>
    
  <script>
    export default {
      name: 'ProfileDefault'
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
    .welcome {
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
  </style>