<script>
  import HomePageDesign from '../components/HomePageDesign.vue'
  export default {
    components: {
      HomePageDesign
    }
  }
</script>

<template>
  <div id="app">
    <HomePageDesign/>
    <router-view></router-view>
  </div>
</template>

<style scoped>
/* #app {
  font-family: 'Poppins';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: black;
} */

/* #app {
    display: flex;
    align-items: center;
    /* background-image: url('../assets/background.png');
    background-size : cover;
    background-position: center; */
/*} */
</style>