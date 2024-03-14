<style scoped>
    .background {
        background-image: url("../assets/background.png");
        background-size: cover;
        background-position: center;
        height: 90.9vh;
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

    .borderstyle {
        border-style: solid;
        border-color: #C7FF9C;
        border-width: 9px;
        border-radius: 10px;
        width: 300px;
        display: flex;
        margin: auto;
        justify-content: center;
        margin-top: 25px;
    }

    .queueno {
        font-size: 150px;
        font-weight: bolder;
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
        border-color: #D9D9D9;
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
        background-color: #D9D9D9;
        border-radius: 50px;
    }

</style>

<template>
    <div class="background">
       
        <div class="card cardPos" style="width: 800px; background-color: white;"> 

            <div class="card-body m-auto">
                <div class="card-title header">You have joined the virtual queue.</div>
                <div class="card-body subheader">
                    Your queue number is
                    <div class="borderstyle">
                        <span class="queueno">36</span>
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
    export default {
        data() {
            return {
                currentQueue: 25,  // numerator
                userQueue: 36,  // denominator (100%)
                progressBar: "0%"
            };
        },
        created() {
            // call backend here
            this.fetchQueueData();
        },
        methods: {
            fetchQueueData() {
                // i need to know what the backend looks like ahh
                const progress = this.currentQueue / this.userQueue * 100;
                this.progressBar = `${progress}%`;
            }
        }
    }
</script>