<template>
  <div class="main">
    <a-row>
      <a-col :span="24">
        <a-card :title="questionnaire.questions[index]" style="width: 800;">
          <a-radio-group v-model="value">
            <a-radio :style="radioStyle" :value="1">{{ questionnaire.answers[0] }}</a-radio>
            <a-radio :style="radioStyle" :value="2">{{ questionnaire.answers[1] }}</a-radio>
            <a-radio :style="radioStyle" :value="3">{{ questionnaire.answers[2] }}</a-radio>
            <a-radio :style="radioStyle" :value="4">{{ questionnaire.answers[3] }}</a-radio>
          </a-radio-group>
        </a-card>
        <div class="buttons">
          <a-button
            key="back"
            :disabled="questionnaire.questions.length <= 1 || index === 0"
            @click="handlePrevious"
          >Anterior</a-button>
          <a-button key="exit" type="danger" @click="handleExit">Salir</a-button>
          <a-button
            key="next"
            v-if="index !== questionnaire.questions.length - 1"
            :disabled="questionnaire.questions.length <= 1"
            @click="handleNext"
          >Siguiente</a-button>
          <a-button
            key="submit"
            v-if="index === questionnaire.questions.length - 1"
            @click="handleDone"
            type="primary"
          >Terminar</a-button>
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  props: {
    questionnaire: Object,
    preview: Boolean
  },
  data() {
    return {
      value: 1,
      index: 0,
      results: [],
      radioStyle: {
        display: "flex",
        height: "35px",
        lineHeight: "15px"
      }
    };
  },
  methods: {
    /*eslint-disable*/
    handlePrevious() {
      this.index--;
    },
    handleNext() {
      this.results.push({
        question: this.questionnaire.questions[this.index],
        answer: this.questionnaire.answers[this.value - 1]
      });
      this.index === this.questionnaire.questions.length - 1
        ? this.index
        : this.index++;
    },
    handleExit() {
      this.$emit("exit");
    },
    handleDone() {
      this.handleNext();
      if (this.preview) {
        this.handleExit();
      } else {
      }
    }
  }
};
</script>

<style scoped>
.main {
  display: flex;
  justify-content: center;
  align-items: center;
  height: 100vh;
  background: #e8f8f8;
}
.buttons {
  display: flex;
  justify-content: space-around;
  padding: 30px;
}
</style>
