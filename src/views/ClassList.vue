<script>
import CardBooker from '../components/CardBooker.vue'
export default {
  components: {
    CardBooker,
},
}

</script>

<template>
    <div class="background">
      <div class="heading">
        <h1>Classes <router-link to="/createClass">Admin</router-link></h1>
        
      </div>
      <CardBooker/>
    </div>
</template>

<style scoped>
    .background {
        background-image: url("../assets/background.png");
        background-size: cover;
        background-position: center;
        height: 90.9vh;
        display: flex;
        flex-direction: column; /* Align items in a column */
        justify-content: flex-start; /* Start from the top */
        align-items: center;
        font-family: Poppins Medium;
    }
    .heading {
      margin-top: 20px; /* Add some margin to the top */
    }
    h1 {
      color: #c7ff9c;
      font-family: Poppins Bold;
    }
    @import url(https://db.onlinewebfonts.com/c/0c28006f19928dfd146027cfd7024ca0?family=Poppins+Medium);
    @import url(https://db.onlinewebfonts.com/c/07ecc0aa9ce268962dea7356eeff50a6?family=Poppins+Bold);

    @font-face {
      font-family: "Poppins Medium";
      src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot");
      src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot?#iefix")
          format("embedded-opentype"),
        url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff2")
          format("woff2"),
        url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff")
          format("woff"),
        url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.ttf")
          format("truetype"),
        url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.svg#Poppins Medium")
          format("svg");
    }

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

</style>