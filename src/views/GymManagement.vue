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
    background-color: #4CAF50;
    /* Green background for increase button */
}

.decrease {
    background-color: #f44336;
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
</style>
<template>
    <div class="background">
        <div class="cardPos">
            <button @click="dequeue" class="entergymbtn">Dequeue</button>
            <button @click="increaseAvailability" class="entergymbtn increase">Increase Gym Availability</button>
            <button @click="decreaseAvailability" class="entergymbtn decrease">Decrease Gym Availability</button>
        </div>
    </div>
</template>

<script>
import axios from 'axios';

export default {
    methods: {
        dequeue() {
            axios.post('/api/dequeue')
                .then(response => {
                    console.log(response);
                    // Handle response
                })
                .catch(error => {
                    console.error(error);
                    // Handle error
                });
        },
        increaseAvailability() {
            axios.post('/api/increase-availability')
                .then(response => {
                    console.log(response);
                    // Handle response
                })
                .catch(error => {
                    console.error(error);
                    // Handle error
                });
        },
        decreaseAvailability() {
            axios.post('/api/decrease-availability')
                .then(response => {
                    console.log(response);
                    // Handle response
                })
                .catch(error => {
                    console.error(error);
                    // Handle error
                });
        }
    }
}
</script>