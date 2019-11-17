<template>
  <a-layout style="height: 100vh" id="components-layout-custom-trigger">
    <a-layout-sider :trigger="null" collapsible v-model="collapsed">
      <div class="logo" />
      <a-menu @click="handleClick" theme="dark" mode="inline" :defaultSelectedKeys="['1']">
        <a-sub-menu key="sub1">
          <span slot="title">
            <a-icon type="solution" />
            <span>Instrumentos</span>
          </span>
          <a-menu-item key="1">Crear</a-menu-item>
          <a-menu-item key="2">Modificar</a-menu-item>
          <a-menu-item key="3">Eliminar</a-menu-item>
        </a-sub-menu>
        <a-menu-item key="sub2">
          <a-icon type="user" />
          <span>Usuarios</span>
        </a-menu-item>
      </a-menu>
    </a-layout-sider>
    <a-layout>
      <a-layout-header style="background: #fff; padding: 0">
        <a-icon
          class="trigger"
          :type="collapsed ? 'menu-unfold' : 'menu-fold'"
          @click="()=> collapsed = !collapsed"
        />
      </a-layout-header>
      <a-layout-content
        :style="{ margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' }"
      >
        <QuestionnaireList v-if="selectedOption === '0'" v-bind:user="user"/>
        <NewQuestionnaire v-bind:user="user" v-on:done="handleDone" v-if="selectedOption === '1'" />
        <h1 v-if="selectedOption === '2'">Modificar</h1>
        <h1 v-if="selectedOption === '3'">Eliminar</h1>
        <h1 v-if="selectedOption === 'sub2'">Usuarios</h1>
      </a-layout-content>
    </a-layout>
  </a-layout>
</template>
<script>
import NewQuestionnaire from "./NewQuestionnaire";
import QuestionnaireList from "./QuestionnaireList";

export default {
  props: {
    user: Object
  },
  components: {
    NewQuestionnaire,
    QuestionnaireList
  },
  data() {
    return {
      collapsed: false,
      selectedOption: "0"
    };
  },
  methods: {
    handleClick(e) {
      this.selectedOption = e.key;
      /*eslint-disable*/

      console.log(this.user);
      /*eslint-disable*/
    },
    handleDone() {
      this.selectedOption = "0";
    }
  }
};
</script>
<style>
#components-layout-custom-trigger .trigger {
  font-size: 18px;
  line-height: 64px;
  padding: 0 24px;
  cursor: pointer;
  transition: color 0.3s;
}

#components-layout-custom-trigger .trigger:hover {
  color: #1890ff;
}

#components-layout-custom-trigger .logo {
  height: 32px;
  background: rgba(255, 255, 255, 0.2);
  margin: 16px;
}
</style>
