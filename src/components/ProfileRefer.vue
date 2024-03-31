<template>
    <div class="default">
        <div class="refer">
          <div>
            <p>Use this link to invite friends</p>
            <input class="link" value="https://gymdaddy.com/join">
            <br>
            <button class="sharelink">Share</button>
          </div>
        </div>
    </div>
</template>
    
  <script>
    export default {
      name: 'ProfileRefer'
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
    .refer {
        width: 90% ;
        height: 80vh;
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
    .link {
        background-color: white;
        color: black;
        padding: 15px 40px;
        margin: 15px ;
        font-size: 25px ;
        font-family:Arial, Helvetica, sans-serif;
        border: 0px;
        border-radius: 12px;
    }
    .sharelink {
      font-size: 25px;
      /* font-family: 'Poppins'; */
      color: #C7FF9C;
      background-color: black;      
      padding: 15px 40px;
      border: 0px;
      border-radius: 12px;
      cursor: pointer;
    }
    .sharelink:hover {
      background-color: #D9D9D920;
    }
  </style>