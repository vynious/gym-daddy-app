<script>
  import ProfileSidebar from '@/components/ProfileSidebar.vue'
  import ProfileEditInfo from '@/components/ProfileEditInfo.vue'
  export default {
    components: {
      ProfileSidebar,
      ProfileEditInfo
    }
  }
</script>

<template>
  <div id="app">
    <ProfileSidebar/>
    <ProfileEditInfo/>
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