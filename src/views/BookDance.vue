<script>
import DatePicker from "../components/DatePicker.vue";
//   i suppose use axios(?)

export default {
  components: {
    DatePicker,
  },
  data() {
    return {
      timings: [], // Array to store fetched timings
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
};

// fetch data from backend - time of the class on the selected day
</script>

<template>
  <div class="background">
    <a-card style="width: 600px; margin-top: 10%">
      <h1>Selected Class: Dance</h1>
      <h1>Choose a date</h1>
      <DatePicker @date-selected="fetchTimings" />

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
@import url(https://db.onlinewebfonts.com/c/0c28006f19928dfd146027cfd7024ca0?family=Poppins+Medium);
@import url(https://db.onlinewebfonts.com/c/07ecc0aa9ce268962dea7356eeff50a6?family=Poppins+Bold);

@font-face {
  font-family: "Poppins Medium";
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot");
  src: url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.eot?#iefix")
      format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff2")
      format("woff2"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.woff")
      format("woff"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.ttf")
      format("truetype"),
    url("https://db.onlinewebfonts.com/t/0c28006f19928dfd146027cfd7024ca0.svg#Poppins Medium")
      format("svg");
}

@font-face {
  font-family: "Poppins Bold";
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot");
  src: url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.eot?#iefix")
      format("embedded-opentype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff2")
      format("woff2"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.woff")
      format("woff"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.ttf")
      format("truetype"),
    url("https://db.onlinewebfonts.com/t/07ecc0aa9ce268962dea7356eeff50a6.svg#Poppins Bold")
      format("svg");
}

.background {
  background-image: url("../assets/background.png");
  background-size: cover;
  background-position: center;
  height: 100vh;
  display: flex;
  flex-direction: column; /* Align items in a column */
  justify-content: flex-start; /* Start from the top */
  align-items: center;
}
h1 {
  font-family: "Poppins Bold";
}
</style>
