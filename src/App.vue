<script>
import HelloWorld from './components/HelloWorld.vue'
import CardBooker from './components/CardBooker.vue';

export default {
  name: 'App',
  components: {
    HelloWorld,
    CardBooker
  }
}
</script><template>
  <img alt="Vue logo" src="./assets/logo.png">
  <HelloWorld msg="Welcome to Your Vue.js App"/>
  <div>
    <CardBooker/> 
    <!-- just testing it out can remove it! -->
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
