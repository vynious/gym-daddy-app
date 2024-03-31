<script>
  import ProfileSidebar from '../components/ProfileSidebar.vue'
  import ProfileDefault from '../components/ProfileDefault.vue'
  export default {
    components: {
      ProfileSidebar,
      ProfileDefault
    }
  }
</script>

<template>
  <div id="app">
    <ProfileSidebar/>
    <ProfileDefault/>
    <router-view></router-view>
  </div>
</template>

<style scoped>
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
/* #app {
  font-family: 'Poppins';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: black;
} */

#app {
    display: flex;
    align-items: center;
    font-family: 'Poppins Medium';
}

body {
    background-color: black;
}



</style>