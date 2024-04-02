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
  font-size: 35px;
  width: 250px;
  justify-content: center;
  margin: auto;
  align-items: center;
  display: flex;
  transition-duration: 0.2s;
  margin-bottom: 20px;
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
</style>

<template>
    <div class="background">
       
        <div class="card cardPos" style="width: 800px; background-color: white; margin-top: 100px;"> 

            <!-- Its the person's turn and they have not entered gym -->
            <div class="card-body m-auto" v-if="isTurn && !showQR">
                <div class="card-title header">It is your turn!</div>
                <div class="card-body wording">Please head down to the gym and click the button below when you have entered the gym to receive your checkout QR code ☺️</div>

                <img src="../assets/gym.png" class="card-img-top gymimg" alt="...">

                <div class="card-body">
                    <button type="button" class="btn entergymbtn" @click="showQR = true; updateGymAvail()">Enter Gym</button>
                </div>
            </div>

            <!-- Generation of QR code after person enters gym-->
            <div class="card-body m-auto" v-else-if="isTurn && showQR">
                <div class="card-title header">Checkout QR Code</div>
                <div class="qr-modal-body">
                    <img :src="QRcodeURL" height="300px" alt="QR Code">
                </div>
                <div class="card-body" style="font-size: 25px;">Remember to scan this QR code to sign out after exiting the gym. Thank you!</div>
            </div>

            <!-- Waiting in line -->
            <div class="card-body m-auto" v-else>
                <div class="card-title header">You have joined the virtual queue.</div>
                <div class="card-body subheader">
                    Your queue number is
                    <div class="borderstyle">
                        <span class="queueno"> {{ userQueue }} </span>
                    </div>

                    <div class="queue-info">
                        {{ currentQueue }}/{{ userQueue }}
                    </div>

                    <div class="queue-progress">
                        <div class="progress-container">
                            <div class="progress-bar" :style="{width: progressBar}"></div>
                        </div>
                    </div>
                </div>
            </div>
        </div>

    </div>

</template>

<script>
import QRCode from "qrcode";
import axios from "axios";

    export default {
        data() {
            return {
                currentQueue: 36,  // numerator
                userQueue: 36,  // denominator (100%)
                progressBar: "0%",
                isTurn: false,
                showQR: false,
                QRcodeURL: "" 
            };
        },
        created() {
            this.fetchQueueData();
            this.callQueueNo();

            QRCode.toDataURL("http://127.0.0.1:8000/api/gym/update-avail")  
            .then(URL => {
                this.QRcodeURL = URL;
            })
            .catch(err => {
                console.error("Error generating QR code: ", err)
            })
        },
        methods: {
            fetchQueueData() {
                const baseURL = "http://127.0.0.1:8000";
                const authToken = sessionStorage.AuthToken;

                // get user's queue number
                axios.post(`${baseURL}/api/queue/join`, {
                    Headers: {
                        Authorisation: `Bearer ${authToken}`
                    }
                })
                .then (response => {
                    console.log(response.data);
                    this.userQueue = response.data.queue_number;

                    // get current queue number
                    axios.post(`${baseURL}/api/queue/upcoming`, {
                        Headers: {
                            Authorisation: `Bearer ${authToken}`
                        }
                    })
                    .then (response => {
                        console.log(response.data);
                        this.currentQueue = response.data.queue_number;

                        const percentage = (this.currentQueue/this.userQueue) * 100;
                        this.progressBar = `${percentage}%`;
                    })
                    .catch (error => {
                        console.log("Error fetching upcoming queue data: ", error);
                    })
                })
                .catch (error => {
                    console.log("Error fetching queue data: ", error);
                })

                
            },
            callQueueNo() {
                if (this.currentQueue == this.userQueue) {
                    this.isTurn = true;
                }
            },
            deQueue() { // function should be in an admin's page
                const baseURL = "http://127.0.0.1:8000";
                const authToken = sessionStorage.AuthToken;

                axios.get(`${baseURL}/api/gym/avail`, {
                    headers: {
                        Authorisation: `Bearer ${authToken}`
                    }
                })
                .then (response => {
                    console.log(response.data);
                    const gymAvail = response.data;
                    
                    // dequeues next person as long as there's space in gym
                    while (gymAvail > 0) {
                        axios.get(`${baseURL}/api/queue/next`, {
                            headers: {
                            Authorisation: `Bearer ${authToken}`
                        }
                        })
                        .then (response => {
                            console.log(response.data);
                        })
                        .catch (error => {
                            console.log("Error dequeuing next person: ", error);
                        })
                    }
                })
                .catch (error => {
                    console.log("Error obtaining gym availabilities: ", error);
                })

            },
            updateGymAvail() { // call this function from this page but function should be in an admin's page
                const baseURL = "http://127.0.0.1:8000";
                const authToken = sessionStorage.AuthToken;

                axios.post(`${baseURL}/api/gym/update-avail`, {
                    headers: {
                        Authorisation: `Bearer ${authToken}`
                    }
                })

            }
        }
    }
</script>
