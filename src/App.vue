<script>
import HelloWorld from './components/HelloWorld.vue'
import CardBooker from './components/CardBooker.vue';
import BookingConfirmation from './components/BookingConfirmation.vue';
import ClassConfirmation from './components/ClassConfirmation.vue'
import ErrorMessage from './components/ErrorMessage.vue'

export default {
  name: 'App',
  components: {
    HelloWorld,
    CardBooker,
    BookingConfirmation,
    ClassConfirmation,
    ErrorMessage,
  }
}
</script><template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <div>
    <!-- just testing it out can remove it! -->
    <!-- it shows an error when I don't call the components so i'll place it here first  -->
    <CardBooker/> 
    <BookingConfirmation/>
    <ClassConfirmation/>
    <ErrorMessage/>
  </div>
</template>



<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>
