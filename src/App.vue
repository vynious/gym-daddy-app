
<script>
  import NavigationBar from './components/NavigationBar.vue'
  export default {
    components: {
      NavigationBar
    }
  }
</script>

<template>
  <div id="app">
    <NavigationBar />
    <router-view></router-view>
  </div>

</template>

<style>
#app {
  font-family: 'Poppins';
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: black;
}
</style>
