<style scoped>
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
  font-size: 60px;
  font-family: 'Poppins Bold', sans-serif;
}

.subheader {
  font-size: 34px;
  font-family: 'Poppins Medium', sans-serif;
}

.circle {
  background-color: #c7ff9c;
  border-radius: 50%;
  padding: 5px;
}

.queueimg {
  margin: auto;
  height: auto;
  width: 400px;
}

.joinbtn {
  background-color: #c7ff9c;
  font-size: 35px;
  width: 250px;
  justify-content: center;
  margin: auto;
  align-items: center;
  display: flex;
  transition-duration: 0.2s;
  margin-bottom: 20px;
  font-family: 'Poppins Medium', sans-serif;
}

.joinbtn:hover {
  background-color: #a9e47c;
  color: black;
}

.routerlink {
  text-decoration: none;
}
</style>

<template>
  <div class="background">

    <div class="card cardPos" style="width: 800px; background-color: white; margin-top: 100px;">

      <div class="card-body m-auto">
        <div class="card-title header">Virtual Queue</div>
        <div class="card-text subheader">
          Join the queue now to enter the gym!
        </div>
      </div>

      <img src="../assets/queue.png" class="card-img-top queueimg" alt="...">

      <div class="card-body">
        <button type="button" class="btn joinbtn" @click="joinQueue()">Join Queue</button>
      </div>

    </div>
  </div>

</template>

<script>
import axios from "axios";

export default {
  name: 'JoinQueue',
  methods: {
    joinQueue() {
      const baseURL = "http://127.0.0.1:8000/api/queue/join";
      const authToken = JSON.parse(localStorage.getItem("token"));
      const userId = localStorage.getItem("user_id");

      if (!userId) {
        alert("User ID is missing.");
        return;
      }

      axios.post(baseURL, { user_id: userId }, { headers: { Authorisation: `Bearer ${authToken}` } })
        .then(response => {
          const queueNumber = response.data.data.queue_number;
          console.log(queueNumber)
          localStorage.setItem("userQueue", queueNumber);
          this.$router.push({ name: 'joinqueue' }); // Assuming you have a route named 'QueueStatus' to show the queue status
        })
        .catch(error => {
          console.error("Error joining queue: ", error);
          alert("Failed to join queue. Please try again.");
        });
    },
  },
};
</script>
