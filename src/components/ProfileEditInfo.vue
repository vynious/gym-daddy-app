<template>

    <div class="default">

        <div class="info">

            <div>     <!-- div needed, if not items will be displayed in line (because of class="info") -->
              <br>
              <div>
              <!-- <div class="infosect"> -->
                <p>Gender</p> 
                <button id="maleButton">M</button> <button class="buttonG">F</button>
              </div> 

              <div style="margin: 30px;">
              <!-- <div class="infosect"> -->
                <p>Current Goals</p>
                <div style="display: flex; justify-content: center; align-items: center;">
                  <div class="goal">
                    <button class="button">Work out for 3 hours a day</button><img src="../assets/remove.png" alt="" class="badge">
                  </div>
                  <div class="goal">
                    <button class="button">Improve on cardio</button><img src="../assets/remove.png" alt="" class="badge">
                  </div>
                </div>
              </div>

              <div class="infosect">
                <p>Add Workout Goal : </p> &nbsp; &nbsp; &nbsp;
                <div class="goal">
                  <input class="button" style="width: 300px;"> <img src="../assets/add.png" alt="" class="addGoal">
                </div>
              </div>

              <div class="infosect">
                <p>Ideal Height & Weight : </p> &nbsp; &nbsp; &nbsp;
                <input type="number" value="171" class="HnW"> <p>cm</p> &nbsp; &nbsp; <input type="number" value="70" class="HnW"> <p>kg</p>
              </div>

              <div class="infosect">
                <p>Current Height & Weight : </p> &nbsp;
                <input type="number" value="171" class="HnW"> <p>cm</p> &nbsp; &nbsp; <input type="number" value="85" class="HnW"> <p>kg</p>
              </div>   
              
              <!-- <button class="save">Save</button> -->

            </div>
        
        </div>
    
    </div>

</template>
    
  <script>
    export default {
      name: 'ProfileEditInfo',
      // mounted() {
      //   // Set focus to the male button when the component mounts
      //   this.$refs.maleButton.focus();
      // }
    }
  </script>
    
  <style scoped>
    .default {
        background: black;
        height: 100vh;
        overflow-y: auto;
        width: 70%;
        display: flex;
        justify-content: center;
        align-items: center;
    }
    .info {
        width: 90% ;
        height: 90vh;
        background-color: #D9D9D920;
        color: white;
        display: flex;
        justify-content: center;
        align-items: center;
        font-size: 28px;
        padding-left: 20px;
    }
    .infosect{
      display: flex;
      justify-content: center;
      align-items: center;
      margin: 30px;
    }
    p {
      margin: 0px;
    }
    .buttonG {
        background-color: #D9D9D900;
        color: white;
        padding: 15px 40px;
        margin: 15px ;
        font-size: 28px ;
        font-family: 'Poppins';
        border: 0px;
        border-radius: 12px;
    }
    #maleButton {
      background-color: #D9D9D920;
      color: #C7FF9C;
      padding: 15px 40px;
      margin: 15px ;
      font-size: 28px ;
      /* font-family: 'Poppins'; */
      border: 0px;
      border-radius: 12px;
    }
    /* .buttonG::selection {
      background-color: #D9D9D920;
      color: #C7FF9C;
    } */
    .buttonG:hover {
      background-color: #C7FF9C20;
      cursor: pointer;
    }
    /* .buttonG:focus {
      background-color: #C7FF9C20;
    } */
    .button {
        background-color: #D9D9D920;
        color: white;
        padding: 15px 40px;
        margin: 15px ;
        font-size: 28px ;
        /* font-family: 'Poppins'; */
        border: 0px;
        border-radius: 12px;
        width: 300px;
    }
    .badge {
      height: 40px;
      width: 40px;
      position: absolute;
      top: 4px;
      right: 0px;
    }
    .goal {
      position: relative;
      display: flex;
    }
    .addGoal {
      height: 60px;
      width: 60px;
      position: absolute;
      top: 15px;
      right: -50px;
    }
    .HnW {
      background-color: #D9D9D920;
      color: #C7FF9C;
      width: 100px;
      padding: 15px 20px;
      margin: 15px ;
      font-size: 28px ;
      /* font-family: 'Poppins'; */
      border: 0px;
      border-radius: 12px;
    }
    /* .save {
      font-size: 25px;
      font-family: 'Poppins';
      color: #C7FF9C;
      background-color: black;      
      padding: 15px 40px;
      border: 0px;
      border-radius: 12px;
    } */
  </style>