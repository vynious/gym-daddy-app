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
}

.subheader {
  font-size: 34px;
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
                <button type="button" class="btn joinbtn" @click="joinqueue()">Join Queue</button>
            </div>

        </div>
      </div>

</template>

<script>
    export default {
        name: 'JoinQueue',
        methods: {
            joinqueue() {
                this.$router.push({name: 'joinqueue'});
            }
        }
    };
</script>
