
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
</style>