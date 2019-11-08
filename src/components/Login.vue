<template>
  <div class="main">
    <a-card>
      <div>
        <a-input placeholder="Nombre de usuario" v-model="userName" ref="userNameInput">
          <a-icon slot="prefix" type="user" />
          <a-tooltip slot="suffix" title="Nombre de usuario">
            <a-icon type="info-circle" style="color: rgba(0,0,0,.45)" />
          </a-tooltip>
        </a-input>
      </div>
      <br />
      <div>
        <a-input-password placeholder="Contraseña" ref="passwordInput" />
      </div>
    </a-card>
  </div>
</template>

<script>
export default {
  data() {
    return {
      userName: ""
    };
  },
  methods: {
    emitEmpty() {
      this.$refs.userNameInput.focus();
      this.userName = "";
      this.passwordInput;
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
}
</style>
