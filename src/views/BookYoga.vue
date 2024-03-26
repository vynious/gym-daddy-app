
<script>
  import DatePicker from '../components/DatePicker.vue'
  export default {
    components: {
      DatePicker
    }
  }

// fetch data from backend - time of the class on the selected day 
</script>

<template>
  <div class="background">
    <a-card style="width: 600px">
      <h1>Selected Class: Yoga</h1>
      <h1>Choose a date</h1>
      <DatePicker/>
    </a-card>
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
    }
    h1 {
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