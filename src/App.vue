<template>
  <div>
    <Login
      v-if="!authenticated"
      v-bind:is-authenticated="authenticated"
      v-on:submitted="handleLogin"
    />
    <DashBoard v-if="authenticated" v-bind:user="decoded" />
  </div>
</template>

<script>
import Login from "./components/Login";
import DashBoard from "./components/DashBoard";

export default {
  data() {
    return {
      authenticated: false,
      decoded: ""
    };
  },
  components: {
    Login,
    DashBoard
  },
  methods: {
    handleLogin(event) {
      const axios = require("axios");
      const jwtDecode = require("jwt-decode");

      const axiosParams = {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        }
      };
      axios
        .post(
          "http://localhost:3000/login",
          {
            username: event[0],
            password: event[1]
          },
          axiosParams
        )
        .then(response => {
          this.authenticated = true;
          this.decoded = jwtDecode(response.data);
        })
        .catch(error => {
          if (error.response.status === 401) {
            this.authenticated = false;
          }
        });
    }
  }
};
</script>

<style scoped>
</style>
