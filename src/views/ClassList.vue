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
    }
    .heading {
      margin-top: 20px; /* Add some margin to the top */
    }
    h1 {
      color: #c7ff9c;
    }
</style>