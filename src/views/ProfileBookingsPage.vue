<script>
  import ProfileSidebar from '@/components/ProfileSidebar.vue'
  import ProfileBookings from '@/components/ProfileBookings.vue'
  export default {
    components: {
      ProfileSidebar,
      ProfileBookings
    }
  }
</script>

<template>
  <div id="app">
    <ProfileSidebar/>
    <ProfileBookings/>
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