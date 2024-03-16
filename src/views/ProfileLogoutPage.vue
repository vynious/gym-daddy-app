<script>
  import ProfileSidebar from '@/components/ProfileSidebar.vue'
  import ProfileLogout from '@/components/ProfileLogout.vue'
  export default {
    components: {
      ProfileSidebar,
      ProfileLogout
    }
  }
</script>

<template>
  <div id="app">
    <ProfileSidebar/>
    <ProfileLogout/>
    <!-- <router-view></router-view> -->
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

#app {
    display: flex;
    align-items: center;
}

body {
    background-color: black;
}

</style>