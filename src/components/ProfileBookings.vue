<template>
    <div class="default">
        <div class="bookingsummary">
            <div>
                <!-- placeholder start -->
                <div class="bookingsect">
                    <button class="booking">
                        <p>Booking Number&nbsp;&nbsp;#39284</p>
                        <p>Date&nbsp;&nbsp;:&nbsp;&nbsp;21 / 01 / 24</p>
                        <p>Time&nbsp;&nbsp;:&nbsp;&nbsp;09 30</p>
                        <p class="details">See Details</p>
                    </button>
                    <button class="delete"><img src="../assets/delete.png" alt=""></button>
                </div>
                <div class="bookingsect">
                    <button class="booking">
                        <p>Booking Number&nbsp;&nbsp;#39299</p>
                        <p>Date&nbsp;&nbsp;:&nbsp;&nbsp;23 / 01 / 24</p>
                        <p>Time&nbsp;&nbsp;:&nbsp;&nbsp;13 00</p>
                        <p class="details">See Details</p>
                    </button>
                    <button class="delete"><img src="../assets/delete.png" alt=""></button>
                </div>
                <!-- placeholder end -->
            </div>
        </div>
    </div>
</template>
    
  <script>
    export default {
      name: 'ProfileBookings'
    }
  </script>
    
  <style scoped>
    .default {
        background: black;
        height: 100vh;
        width: 70%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .bookingsummary {
        width: 90% ;
        height: 90vh;
        overflow-y: auto;
        background-color: #D9D9D920;
        color: white;
        display: flex;
        justify-content: center;
        font-size: 28px;
    }
    .bookingsect{
      display: flex;
      justify-content: center;
      align-items: center;
      margin-top: 50px;
    }
    .booking {
        background-color: #D9D9D920;
        border-radius: 12px;
        width: 700px;
        color: white;
        padding: 15px 40px;
        margin-right: 30px ;
        font-size: 28px ;
        /* font-family: 'Poppins'; */
        border: 0px;
        border-radius: 12px;
    }
    p {
      margin: 3px;
    }
    .details {
        font-size : 20px;
        text-decoration: underline;
        cursor: pointer;
    }
    img {
        width: 70px;
        padding-left: auto;
        padding-right: auto;
    }
    .delete {
        background-color: #D9D9D920;
        border-radius: 12px;
        width: 100px;
        height: 100px;
        color: white;
        padding: 15px 40px;
        margin: 15px ;
        border: 0px;
        border-radius: 12px;
        display: flex;
        justify-content: center;
        align-items: center;
        cursor: pointer;
    }
    .delete:hover {
        /* background-color: #FFFFFF50; */
        background-color: #00000050;

    }
  </style>