<script>
import CardBooker from '../components/CardBooker.vue'
export default {
  components: {
    CardBooker,
},
}

</script>

<template>
    <h1>Classes</h1>
    <CardBooker/>
</template>

<style scoped>
body {
  /* background-image: url('/images/image.png');
  background-repeat: no-repeat;
  background-size: cover; */ 
  /* somehow the background image seems to be very blur so i just put it as black first */
  background-color: black;
}
h1 {
  color: #c7ff9c;
}
</style>