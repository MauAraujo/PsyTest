<template>
  <Login v-on:submitted="handleLogin" />
</template>

<script>
import Login from "./components/Login";
// import Questionnaire from "./components/Questionnaire";
// import DashBoard from './components/DashBoard'

export default {
  components: {
    Login
    // Questionnaire
    // DashBoard,
  },
  methods: {
    handleLogin(event) {
      const axios = require("axios");
      /*eslint-disable */

      console.log(event);
      axios
        .post("http://localhost:3000/login", {
          username: event[0],
          password: event[1]
        })
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.error(error);
        });
      /*eslint-disable */
    }
  }
};
</script>

<style scoped>
</style>
