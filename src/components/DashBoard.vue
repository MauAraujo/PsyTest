<template>
  <div>
    <a-layout style="height: 100vh" id="components-layout-custom-trigger">
      <a-layout-sider :trigger="null" collapsible v-model="collapsed">
        <div class="logo" />
        <!-- Admin Menu -->
        <a-menu
          v-if="user.admin"
          @click="handleClick"
          theme="dark"
          mode="inline"
          :defaultSelectedKeys="['0']"
        >
          <a-sub-menu key="sub1">
            <span slot="title">
              <a-icon type="solution" />
              <span>Instrumentos</span>
            </span>
            <a-menu-item key="0">Ver pruebas</a-menu-item>
            <a-menu-item key="1">Crear prueba</a-menu-item>
          </a-sub-menu>
          <a-sub-menu key="sub2">
            <span slot="title">
              <a-icon type="user" />
              <span>Usuarios</span>
            </span>
            <a-menu-item key="2">Ver pacientes</a-menu-item>
            <a-menu-item key="3">Crear paciente</a-menu-item>
          </a-sub-menu>
        </a-menu>
        <!-- Patient Menu -->
        <a-menu
          v-if="user.patient"
          @click="handleClick"
          theme="dark"
          mode="inline"
          :defaultSelectedKeys="['0']"
        >
          <a-sub-menu key="sub1">
            <span slot="title">
              <a-icon type="solution" />
              <span>Instrumentos</span>
            </span>
            <a-menu-item key="0">Contestar prueba</a-menu-item>
            <a-menu-item key="1">Mis pruebas</a-menu-item>
          </a-sub-menu>
        </a-menu>
      </a-layout-sider>
      <a-layout>
        <a-layout-header
          :style="[!viewQuestionnaire ? {background: '#fff', padding: 0} : {background: '#e8f8f8', padding: 0}]"
        >
          <a-icon
            class="trigger"
            :type="collapsed ? 'menu-unfold' : 'menu-fold'"
            @click="()=> collapsed = !collapsed"
          />
        </a-layout-header>
        <!-- Admin Content -->
        <a-layout-content
          v-if="user.admin"
          :style="[!viewQuestionnaire ? { margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' } : {}]"
        >
          <Questionnaire
            v-on:exit="handleExit"
            v-if="viewQuestionnaire"
            v-bind:questionnaire="questionnaire"
            v-bind:preview="true"
          />
          <QuestionnaireForm
            v-if="editQuestionnaire"
            v-bind:user="user"
            v-bind:edit="true"
            v-on:done="handleDone"
            v-bind:questionnaire="questionnaire"
          />
          <QuestionnaireList
            v-if="selectedOption === '0' && !viewQuestionnaire && !editQuestionnaire"
            v-bind:user="user"
            v-bind:own="false"
            v-on:edit="handleEdit"
            v-on:get="handleGet"
          />
          <QuestionnaireForm
            v-bind:user="user"
            v-bind:edit="false"
            v-on:done="handleDone"
            v-if="selectedOption === '1' && !editQuestionnaire"
          />
          <div v-if="selectedOption === '2'">
            <UserList v-on:edit="handleUserEdit" />
          </div>
          <div v-if="selectedOption === '3'">
            <UserForm v-on:done="handleUserDone" />
          </div>
        </a-layout-content>
        <!-- Patient Content -->
        <a-layout-content
          v-if="user.patient"
          :style="[!viewQuestionnaire ? { margin: '24px 16px', padding: '24px', background: '#fff', minHeight: '280px' } : {}]"
        >
          <QuestionnaireList
            v-if="selectedOption === '0' && !viewQuestionnaire"
            v-bind:own="false"
            v-bind:user="user"
            v-on:edit="handleEdit"
            v-on:get="handleGet"
          />
          <QuestionnaireList
            v-if="selectedOption === '1' && !viewQuestionnaire"
            v-bind:own="true"
            v-bind:user="user"
            v-on:edit="handleEdit"
            v-on:get="handleGet"
          />
          <Questionnaire
            v-on:exit="handleExit"
            v-if="viewQuestionnaire"
            v-bind:questionnaire="questionnaire"
            v-bind:user="user"
            v-bind:preview="false"
          />
        </a-layout-content>
      </a-layout>
    </a-layout>
  </div>
</template>
<script>
import QuestionnaireForm from "./QuestionnaireForm";
import QuestionnaireList from "./QuestionnaireList";
import Questionnaire from "./Questionnaire";
import UserList from "./UserList";
import UserForm from "./UserForm";

export default {
  props: {
    user: Object
  },
  components: {
    QuestionnaireForm,
    Questionnaire,
    QuestionnaireList,
    UserList,
    UserForm
  },
  data() {
    return {
      collapsed: false,
      selectedOption: "0",
      editQuestionnaire: false,
      viewQuestionnaire: false,
      questionnaire: Object
    };
  },
  methods: {
    /*eslint-disable*/

    handleClick(e) {
      this.viewQuestionnaire = false;
      this.editQuestionnaire = false;
      this.selectedOption = e.key;
    },
    handleDone() {
      this.editQuestionnaire = false;
      this.viewQuestionnaire = false;
      this.selectedOption = "0";
    },
    handleUserDone() {
      this.selectedOption = "2";
    },
    handleEdit(event) {
      this.editQuestionnaire = true;
      this.questionnaire = event;
    },
    handleUserEdit(event) {
      this.editUser = true;
      this.user = event;
    },
    handleGet(event) {
      this.viewQuestionnaire = true;
      this.questionnaire = event;
    },
    handleExit() {
      this.viewQuestionnaire = false;
    }
    /*eslint-disable*/
  },
  mounted() {
    /*eslint-disable*/
    console.log(this.user);
    /*eslint-disable*/
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
