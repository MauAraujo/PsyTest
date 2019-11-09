<template>
  <div class="main">
    <a-card>
      <div class="card-body">
        <div style="margin: 10px">
          <a-input placeholder="Nombre de usuario" v-model="userName" ref="userNameInput">
            <a-icon slot="prefix" type="user" />
            <a-tooltip slot="suffix" title="Nombre de usuario">
              <a-icon type="info-circle" style="color: rgba(0,0,0,.45)" />
            </a-tooltip>
          </a-input>
        </div>
        <br />
        <div style="margin: 10px">
          <a-input-password placeholder="Contraseña" v-model="password" ref="passwordInput" />
        </div>
        <a-button v-on:click="validate" style="margin: 10px" type="primary" shape="round">Ingresar</a-button>
      </div>
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: "",
      password: "",
    };
  },
  methods: {
    validate() {
      this.$emit('submitted', [this.userName, this.password])
    }
  }
};
</script>

<style>
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #e8f8f8;
}

.card-body {
  display: flex;
  justify-content: center;
  flex-direction: column;
}
</style>
