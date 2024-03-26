
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

body {
  margin: 0;
  padding: 0;
}
</style>

<template>
  <div id="app">
    <router-view></router-view>
  </div>
</template>

<script>

import "./styles.css";

export default {
  name: "App",
};
</script>

<style>
</style>

