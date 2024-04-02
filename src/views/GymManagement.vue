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
  justify-content: center;
  align-items: center;
}

.cardPos {
  justify-content: center;
  display: flex;
  margin: auto;
  align-items: center;
  margin-top: 50px;
}

.header {
  font-weight: bold;
  font-size: 48px;
  margin-bottom: 0px;
}

.subheader {
  font-weight: medium;
  font-size: 34px;
}

.wording {
  font-weight: medium;
  font-size: 28px;
  font-family: "Poppins Bold";
}

.borderstyle {
  border-style: solid;
  border-color: #c7ff9c;
  border-width: 9px;
  border-radius: 10px;
  width: 300px;
  display: flex;
  margin: auto;
  justify-content: center;
  margin-top: 25px;
}

.gymimg {
  margin: auto;
  height: auto;
  width: 400px;
}

.queueno {
  font-size: 150px;
  font-weight: bolder;
}

.entergymbtn {
  background-color: #c7ff9c;
  font-size: 25px;
  width: 50%;
  justify-content: center;
  margin: auto;
  align-items: center;
  display: flex;
  transition-duration: 0.2s;
  margin-bottom: 20px;
  border-radius: 5px;
  padding: 10px;
}

.queue-info {
  margin-top: 40px;
  font-size: 24px;
  font-weight: medium;
  text-align: right;
}

.queue-progress {
  width: 100%;
  text-align: center;
  margin-top: 10px;
  border-style: solid;
  border-color: #d9d9d9;
  border-radius: 50px;
}

.progress-container {
  width: 100%;
  height: 30px;
  background-color: white;
  border-radius: 50px;
}

.progress-bar {
  height: 100%;
  background-color: #d9d9d9;
  border-radius: 50px;
}

.entergymbtn {
  color: #fff;
  cursor: pointer;
  transition: all 0.3s ease;
  box-shadow: 0 4px 14px 0 rgba(0, 118, 255, 0.39);
  border: none;
  outline: none;
}

.entergymbtn:hover {
  transform: translateY(-2px);
}

.increase {
  background-color: #47b14a;
  /* Green background for increase button */
}

.decrease {
  background-color: #e23529;
  /* Red background for decrease button */
}

.cardPos {
  justify-content: center;
  display: flex;
  margin: auto;
  align-items: center;
  flex-direction: column;
  /* Align buttons vertically */
}
.queue-info {
  text-align: center;
  margin: 20px 0;
  background-color: #f7f7f7;
  border-radius: 8px;
  padding: 20px;
  box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
  font-family: "Poppins", sans-serif;
}

.queue-info p {
  font-size: 20px;
  color: #333;
  margin: 10px 0;
  letter-spacing: 0.5px;
}

.queue-info p span {
  display: block;
  font-size: 24px;
  color: #2c3e50;
  font-weight: 600;
}
</style>
<template>
  <div class="background" style="font-family: 'Poppins Bold';">
    <div class="cardPos">
      <div class="queue-info" style="font-family: 'Poppins Bold'">
        <p>Current Gym Availability: {{ gymAvailability }}</p>
        <p>Upcoming Queue Number: {{ upcomingQueueNumber }}</p>
      </div>
      <button @click="dequeue" class="entergymbtn" style="color: black;">Dequeue</button>
      <button @click="increaseAvailability" class="entergymbtn increase" style="color: black;">
        Increase Gym Availability
      </button>
      <button @click="decreaseAvailability" class="entergymbtn decrease" style="color: black;">
        Decrease Gym Availability
      </button>
    </div>
  </div>
</template>

<script>
import axios from "axios";

export default {
  data() {
    return {
      gymAvailability: 0,
      upcomingQueueNumber: "No one in Queue",
    };
  },
  created() {
    this.fetchGymAvailability();
    this.fetchUpcomingQueueNumber();
  },
  methods: {
    fetchGymAvailability() {
      const authToken = JSON.parse(localStorage.getItem("token"));
      axios
        .get("http://127.0.0.1:8000/api/gym/avail", {
          headers: { Authorisation: `Bearer ${authToken}` },
        })
        .then((response) => {
          this.gymAvailability = response.data.count;
        })
        .catch((error) => {
          console.error("Error fetching gym availability: ", error);
        });
    },
    fetchUpcomingQueueNumber() {
      const authToken = JSON.parse(localStorage.getItem("token"));
      axios
        .get("http://127.0.0.1:8000/api/queue/upcoming", {
          headers: { Authorisation: `Bearer ${authToken}` },
        })
        .then((response) => {
          this.upcomingQueueNumber = response.data.data.queue_number;
        })
        .catch((error) => {
          console.error("Error fetching upcoming queue number: ", error);
        });
    },
    dequeue() {
      const authToken = JSON.parse(localStorage.getItem("token"));
      axios
        .get("http://127.0.0.1:8000/api/queue/next", {
          headers: { Authorisation: `Bearer ${authToken}` },
        })
        .then((response) => {
          console.log(response);
          alert("Removed first person from the queue.");
          this.fetchGymAvailability();
          this.fetchUpcomingQueueNumber();
        })
        .catch((error) => {
          console.error(error);
          alert(error.response.data.message);
          console.log(error.response.data);

          if (error.response.data == "failed to get next ticket") {
            this.upcomingQueueNumber = "No one in Queue";
          }
        });
    },
    increaseAvailability() {
      const authToken = JSON.parse(localStorage.getItem("token"));
      axios
        .post(
          "http://127.0.0.1:8000/api/gym/update-avail",
          { update_type: "increment", quantity: 1 },
          { headers: { Authorisation: `Bearer ${authToken}` } }
        )
        .then((response) => {
          console.log(response);
          alert("Successfully updated gym availability.");
          this.fetchGymAvailability();
        })
        .catch((error) => {
          console.error(error);
          alert(error.response.data.error);
        });
    },
    decreaseAvailability() {
      const authToken = JSON.parse(localStorage.getItem("token"));
      axios
        .post(
          "http://127.0.0.1:8000/api/gym/update-avail",
          { update_type: "decrement", quantity: 1 },
          { headers: { Authorisation: `Bearer ${authToken}` } }
        )
        .then((response) => {
          console.log(response);
          alert("Successfully updated gym availability.");
          this.fetchGymAvailability();
        })
        .catch((error) => {
          console.error(error);
          alert(error.response.data.error);
        });
    },
  },
};
</script>
