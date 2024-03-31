<style scoped>
    .background {
        background-image: url("../assets/background.png");
        background-size: cover;
        background-position: center;
        height: 100vh;
        display: flex;
        justify-content: center;
        align-items: center;
        font-family: 'Poppins medium';
    }

    .cardPos {
        justify-content: center;
        display: flex;
        margin: auto;
        align-items: center;
        margin-top: 50px;
    }

    .header {
        font-family: 'Poppins bold';
        font-size: 58px;
    }

    .subheader {
        font-size: 30px;
    }

    .circle {
        background-color: #C7FF9C;
        border-radius: 50%;
        padding: 5px;
    }

    .queueimg {
        margin: auto;
        height: auto;
        width: 400px;
    }

    .joinbtn {
        background-color: #C7FF9C;
        font-size: 28px;
        width: 250px;
        justify-content: center;
        margin: auto;
        align-items: center;
        display: flex;
        transition-duration: 0.2s;
        margin-bottom: 20px;
    }

    .joinbtn:hover {
        background-color: #A9E47C; 
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
                    There are currently
                    <span class="circle"> {{ queueNo }} </span>
                    people in line.
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
        data() {
            return {
                queueNo: 0
            }
        },
        methods: {
            joinqueue() {
                this.$router.push({name: 'joinqueue'});
                // const baseURL = "";

                // this.$axios.get(`${baseURL}/`,{  // to check path
                //     // headers: {
                //     //     Authorization: ``
                //     // }
                // }) 
                // .then(response => {
                //     console.log(response.data);
                //     // this.queueNo = 
                // })
                
            }
        }
    };
</script>