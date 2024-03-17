
<script>
  import DatePicker from '../components/DatePicker.vue'
//   i suppose use axios(?)

  export default {
    components: {
      DatePicker
    },
    data() {
      return {
        timings: [] // Array to store fetched timings
      };
    },
    // methods: {
    //   // Function to fetch timings based on the selected date
    //   async fetchTimings(selectedDate) {
    //     try {
    //       // Make an HTTP request to backend API to fetch timings
    //       const response = await axios.get(`/api/timings?date=${selectedDate}`);
    //       // Assuming your backend returns an array of timings, update the timings array
    //       this.timings = response.data;
    //     } catch (error) {
    //       console.error('Error fetching timings:', error);
    //     }
    //   }
    // }
  }

// fetch data from backend - time of the class on the selected day 
</script>

<template>
  <div class="background">
    <a-card style="width: 600px">
      <h1>Selected Class: Dance</h1>
      <h1>Choose a date</h1>
      <DatePicker @date-selected = "fetchTimings"/>

      <!-- Display the fetched timings here -->
      <div v-if="timings.length">
        <h2>Available Timings:</h2>
        <ul>
          <li v-for="timing in timings" :key="timing.id">{{ timing.time }}</li> 
        </ul>
      </div>
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