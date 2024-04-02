<template>
  <div class="scroll">
    <div id="top">
      <div id="overlap1"></div>
      <div id="overlap2"></div>
      <div class="info">
        <h1 class="gymname">GYM DADDY</h1>
        <p>Your Ultimate Gym Stop</p>
        <p>We Provide You with the Best Experience and Qualified Trainers</p>
        <br />
        <a href="#bottom"><button id="explore">Explore More</button></a>
      </div>
      <a href="#bottom">
        <div id="down">
          <img src="../assets/down.png" alt="" style="width: 50px" />
        </div>
      </a>
    </div>

    <div id="bottom">
      <div class="topinfo">
        <div class="level">
          <h1>Level&nbsp;&nbsp;<span style="font-size: 60px">1</span></h1>
          <p>BEGINNER</p>
          <p>0 to 1 year experience</p>
        </div>

        <div class="separator"></div>

        <div class="level">
          <h1>Level&nbsp;&nbsp;<span style="font-size: 60px">2</span></h1>
          <p>INTERMEDIATE</p>
          <p>1 to 2 years experience</p>
        </div>

        <div class="separator"></div>

        <div class="level">
          <h1>Level&nbsp;&nbsp;<span style="font-size: 60px">3</span></h1>
          <p>ADVANCED</p>
          <p>2 to 3+ years experience</p>
        </div>
      </div>

      <div class="bottominfo">
        <div class="categories">
          <div id="cardio">
            <div class="label">CARDIO</div>
          </div>
          <div id="strength">
            <div class="label">STRENGTH</div>
          </div>
          <div id="yoga">
            <div class="label">YOGA</div>
          </div>
        </div>

        <div class="actions">
          <router-link to="/queue"><button id="joinqueue">JOIN QUEUE</button></router-link>
          <br />
          <router-link to="/book"><button id="bookclass">BOOK CLASSES</button></router-link>
        </div>
      </div>
    </div>

    <div id="footer">
      <div class="copyright">
        <p class="copyrighttext">© 2024 Gym Daddy</p>
      </div>
      <img src="../assets/mail.png" alt="" class="icon" />
      <br />
      <img src="../assets/facebook.png" alt="" class="icon" />
      <br />
      <img src="../assets/instagram.png" alt="" class="icon" />
    </div>
  </div>
  <!-- Modal -->
  <transition name="modal">
    <div class="modal-mask" v-if="showModal">
      <div class="modal-wrapper">
        <div class="modal-container">
          <div class="modal-header">
            <slot name="header">Join Our Telegram Group!</slot>
          </div>

          <div class="modal-body">
            <slot name="body">
              Get the latest updates and notifications by joining our Telegram group.
            </slot>
          </div>

          <div class="modal-footer">
            <slot name="footer">
              <button class="modal-default-button" @click="showModal = false">
                Close
              </button>
              <!-- Add your telegram group link here -->
              <a href="https://t.me/gym_daddy_bot" target="_blank" class="modal-default-button">
                Join Now
              </a>
            </slot>
          </div>
        </div>
      </div>
    </div>
  </transition>
</template>

<script>
export default {
  name: "HomePageDesign",
  data() {
    return {
      showModal: false,
    };
  },
  mounted() {
    // Show the modal after 5 seconds
    setTimeout(() => {
      this.showModal = true;
    }, 5000);
  },

};
</script>

<style scoped>
html {
  scroll-behavior: smooth;
}

.scroll {
  max-height: 100vh;
  overflow-y: auto;
}

#top {
  height: 100vh;
  position: relative;
  background-image: url("../assets/background.png");
  background-size: cover;
  /* Ensure the background image covers the entire div */
  background-position: center;
  /* Center the background image */
}

#overlap1 {
  border: 2px solid #c7ff9c;
  width: 90vw;
  height: 77vh;
  position: absolute;
  background-color: transparent;
  left: 65px;
  top: 70px;
}

#overlap2 {
  border: 2px solid #c7ff9c;
  width: 90vw;
  height: 77vh;
  position: absolute;
  background-color: transparent;
  left: 87px;
  top: 90px;
}

.info {
  width: 40vw;
  height: 30vh;
  position: absolute;
  top: 220px;
  left: 770px;
  color: white;
}

/* .glow {
        font-size: 70px;
        /* font-family: 'Poppins Medium'; */
/* color: white; */
/* Add glow effect using text-shadow */
/* text-shadow: 0 0 10px #C7FF9C, 0 0 20px #C7FF9C, 0 0 30px #C7FF9C, 0 0 40px #C7FF9C, 0 0 50px #C7FF9C, 0 0 60px #C7FF9C, 0 0 70px #C7FF9C; */
/* Add animation to make the glow pulsate */
/* animation: pulse 1.5s infinite alternate; */
/*} */

/* @keyframes pulse {
        from {
            text-shadow: 0 0 10px #C7FF9C, 0 0 20px #C7FF9C, 0 0 30px #C7FF9C, 0 0 40px #C7FF9C, 0 0 50px #C7FF9C, 0 0 60px #C7FF9C, 0 0 70px #C7FF9C;
        }
        to {
            text-shadow: 0 0 20px #C7FF9C, 0 0 30px #C7FF9C, 0 0 40px #C7FF9C, 0 0 50px #C7FF9C, 0 0 60px #C7FF9C, 0 0 70px #C7FF9C, 0 0 80px #C7FF9C;
        }
    } */

.gymname {
  font-size: 5.5rem;
  color: #fff;
  text-shadow: 0 0 15px #c7ff9c;
  font-family: "Poppins Bold";
  white-space: nowrap;
}

#explore {
  font-size: 25px;
  /* font-family: 'Poppins'; */
  color: #c7ff9c;
  background-color: black;
  padding: 15px 30px;
  border: 0px;
  /* border-radius: 12px; */
  cursor: pointer;
}

#explore:hover {
  background-color: #ffffff70;
}

#joinqueue:hover,
#bookclass:hover {
  background-color: #c7ff9c;
  color: black;
}

p {
  font-size: 22px;
}

#down {
  cursor: pointer;
  position: absolute;
  top: 770px;
  left: 700px;
}

#bottom {
  height: 100vh;
  width: 100vw;
  background-color: black;
  position: relative;
  color: white;
}

.topinfo,
.bottominfo {
  height: 50vh;
  display: flex;
  justify-content: center;
  align-items: center;
}

.level {
  /* background-color: green ; */
  padding: 20px;
  margin: 70px;
}

.level h1 {
  color: #c7ff9c;
  font-size: 40px;
}

.separator {
  border-left: 2px solid white;
  height: 300px;
  margin: 30px;
}

.categories {
  margin-left: 40px;
  width: 1200px;
  height: 300px;
  position: relative;
  font-size: 35px;
}

#cardio {
  background-image: url("../assets/cardio.png");
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 700px;
  height: 100px;
  border-radius: 12px;
  position: absolute;
  top: 0px;
  left: 30px;
}

#strength {
  background-image: url("../assets/strength.png");
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 700px;
  height: 100px;
  border-radius: 12px;
  position: absolute;
  top: 120px;
  left: 130px;
}

#yoga {
  background-image: url("../assets/yoga.png");
  background-size: cover;
  background-position: center;
  display: flex;
  justify-content: center;
  align-items: center;
  width: 700px;
  height: 100px;
  border-radius: 12px;
  position: absolute;
  top: 240px;
  left: 230px;
}

.label {
  width: 700px;
  height: 100px;
  background-color: #ffffff40;
  display: flex;
  justify-content: center;
  align-items: center;
  border-radius: 12px;
  text-shadow: 2px 2px 2px #c7ff9c;
}

.actions {
  margin-right: 50px;
}

#joinqueue {
  font-size: 27px;
  /* font-family: 'Poppins'; */
  color: #c7ff9c;
  background-color: black;
  width: 280px;
  margin: 20px;
  padding: 20px 30px;
  border: 1px solid #c7ff9c;
  /* border-radius: 12px; */
  cursor: pointer;
}

#bookclass {
  font-size: 27px;
  /* font-family: 'Poppins'; */
  color: #c7ff9c;
  background-color: black;
  width: 280px;
  margin: 20px;
  padding: 20px 30px;
  border: 1px solid #c7ff9c;
  /* border-radius: 12px; */
  cursor: pointer;
}

#footer {
  height: 60px;
  width: 100vw;
  background-color: black;
  display: flex;
  padding: 20px;
}

.icon {
  width: 30px;
  height: 30px;
  display: inline;
  margin-left: 30px;
}

.copyright {
  width: 80vw;
  text-align: left;
}

.copyrighttext {
  font-size: 18px;
  color: white;
}

/* Modal */
.modal-mask {
  position: fixed;
  z-index: 9999;
  top: 0;
  left: 0;
  width: 100%;
  height: 100%;
  background-color: rgba(0, 0, 0, 0.6);
  display: flex;
  align-items: center;
  justify-content: center;
}

.modal-wrapper {
  padding: 10px;
}

.modal-container {
  min-width: 500px;
  background-color: #fff;
  padding: 20px;
  border-radius: 8px;
  box-shadow: 0 0 10px rgba(0, 0, 0, 0.25);
  text-align: center;
}

.modal-header {
  margin-bottom: 20px;
}

.modal-header h3 {
  font-size: 24px;
  color: #333;
}

.modal-body {
  margin-bottom: 20px;
  font-size: 18px;
  color: #555;
}

.modal-footer {
  text-align: right;
}

.modal-default-button {
  padding: 10px 20px;
  margin-right: 10px;
  background-color: #f2f2f2;
  border: 1px solid #ccc;
  border-radius: 4px;
  cursor: pointer;
  font-size: 16px;
}

.modal-default-button:hover {
  background-color: #e6e6e6;
}

.modal-close-button {
  color: #aaa;
  cursor: pointer;
}

.modal-close-button:hover {
  color: #777;
}

/* Optional: Add some transition for the modal */
.modal-enter-active,
.modal-leave-active {
  transition: opacity 0.5s;
}

.modal-enter,
.modal-leave-to

/* .modal-leave-active in <2.1.8 */
  {
  opacity: 0;
}

/* Optional: Animate the modal container to get a little bounce effect */
.modal-enter-active .modal-container {
  transition: all 0.5s;
  transform: translateY(-100px);
}

.modal-enter .modal-container {
  transform: translateY(0px);
}
</style>
