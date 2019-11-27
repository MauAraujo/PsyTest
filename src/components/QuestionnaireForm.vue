<template>
  <div>
    <a-steps :current="current">
      <a-step title="Nombra el instrumento">
        <a-icon type="tool" slot="icon" />
      </a-step>
      <a-step title="Ingresa los campos">
        <a-icon type="form" slot="icon" />
      </a-step>
      <a-step title="Confirma la información">
        <a-icon type="file-done" slot="icon" />
      </a-step>
    </a-steps>

    <div class="steps-content">
      <!-- Step 1 -->
      <div v-if="current === 0">
        <a-form :form="nameForm">
          <a-form-item
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
            label="Nombre del instrumento"
          >
            <a-input
              v-decorator="['testName', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, message: 'El nombre es obligatorio' }] }]"
              allowClear
            />
          </a-form-item>
          <a-form-item
            :label-col="{ span: 5 }"
            :wrapper-col="{ span: 12 }"
            label="Descripción del instrumento"
          >
            <a-textarea
              :autosize="{ minRows: 2, maxRows: 6 }"
              v-decorator="['description', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, message: 'La descripción es obligatoria' }] }]"
              allowClear
            />
          </a-form-item>
        </a-form>
      </div>
      <!-- Step 2  -->
      <div v-if="current === 1">
        <a-divider>Ingresa las respuestas</a-divider>
        <a-form :form="questionsForm">
          <a-form-item :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }" label="Respuesta 1">
            <a-input
              v-decorator="['answer1', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, whitespace: true, message: 'Esta respuesta es obligatoria' }] }]"
              allowClear
            />
          </a-form-item>
          <a-form-item :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }" label="Respuesta 2">
            <a-input
              v-decorator="['answer2', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, whitespace: true,message: 'Esta respuesta es obligatoria' }] }]"
              allowClear
            />
          </a-form-item>
          <a-form-item :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }" label="Respuesta 3">
            <a-input
              v-decorator="['answer3', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, whitespace: true, message: 'Esta respuesta es obligatoria' }] }]"
              allowClear
            />
          </a-form-item>
          <a-form-item :label-col="{ span: 6 }" :wrapper-col="{ span: 12 }" label="Respuesta 4">
            <a-input
              v-decorator="['answer4', { 
                validateTrigger: ['change', 'blur'],
                rules: [{ required: true, whitespace: true, message: 'Esta respuesta es obligatoria' }] }]"
              allowClear
            />
          </a-form-item>
          <a-divider>Ingresa las preguntas</a-divider>

          <div v-if="!edit">
            <a-form-item
              v-for="(k) in questionsForm.getFieldValue('keys')"
              :key="k"
              :wrapper-col="{ span: 18 }"
              :label-col="{ span: 6 }"
              label=" Pregunta"
              :required="false"
            >
              <a-input
                v-decorator="[
          `questions[${k}]`,
          {
            validateTrigger: ['change', 'blur'],
            rules: [
              {
                required: true,
                whitespace: true,
                message: 'Ingresa una pregunta',
              },
            ],
          },
        ]"
                style="width: 60%; margin-right: 8px"
              />
              <a-icon
                v-if="questionsForm.getFieldValue('keys').length > 1"
                class="dynamic-delete-button"
                type="minus-circle-o"
                :disabled="questionsForm.getFieldValue('keys').length === 1"
                @click="() => remove(k)"
              />
            </a-form-item>
            <a-form-item :wrapper-col="{ span: 24 }">
              <a-button type="dashed" style="width: 100%" @click="add">
                <a-icon type="plus" />Agregar pregunta
              </a-button>
            </a-form-item>
          </div>

          <div v-if="edit">
            <a-form-item
              v-for="(question) in questionsForm.getFieldValue('questions')"
              :key="question"
            >
              <a-input :defaultValue="question" />
            </a-form-item>
          </div>
        </a-form>
      </div>

      <!-- Step 3 -->
      <div v-if="current === 2">
        <a-input-group size="large">
          <a-row :gutter="8">
            <a-col :span="8">
              <h3>Nombre del instrumento:</h3>
            </a-col>
            <a-col :span="12">
              <a-input disabled :defaultValue="nameForm.getFieldValue('testName')" />
            </a-col>
          </a-row>
          <a-row :gutter="8">
            <a-col :span="8">
              <h3>Descripción del instrumento:</h3>
            </a-col>
            <a-col :span="12">
              <a-input disabled :defaultValue="nameForm.getFieldValue('description')" />
            </a-col>
          </a-row>
          <a-divider />
          <a-row
            v-for="(question) in questionsForm.getFieldValue('questions')"
            :key="questionsForm.getFieldValue('questions').indexOf(question)"
            :gutter="8"
          >
            <a-col :span="4">
              <h3>Pregunta {{ questionsForm.getFieldValue('questions').indexOf(question) + 1 }}</h3>
            </a-col>
            <a-col :span="10">
              <a-input disabled :defaultValue="question" />
            </a-col>
            <a-col :span="10">
              <a-radio-group>
                <a-radio-button>{{ questionsForm.getFieldValue('answer1') }}</a-radio-button>
                <a-radio-button>{{ questionsForm.getFieldValue('answer2') }}</a-radio-button>
                <a-radio-button>{{ questionsForm.getFieldValue('answer3') }}</a-radio-button>
                <a-radio-button>{{ questionsForm.getFieldValue('answer4') }}</a-radio-button>
              </a-radio-group>
            </a-col>
          </a-row>
        </a-input-group>
      </div>
    </div>
    <div class="steps-action">
      <a-button v-if="current>0" style="margin-left: 8px" @click="prev">Regresar</a-button>
      <a-button v-if="current < 3 - 1" type="primary" @click="next">Siguiente</a-button>
      <a-button v-if="current == 3 - 1" type="primary" @click="next">Terminar</a-button>
    </div>
    <a-modal
      title="Confirmar creación de instrumento"
      :visible="visible"
      @ok="handleOk"
      :confirmLoading="confirmLoading"
      @cancel="handleCancel"
    >
      <p>{{ModalText}}</p>
    </a-modal>
  </div>
</template>
<script>
let id = 0;
export default {
  props: {
    user: Object,
    questionnaire: Object,
    edit: Boolean
  },
  data() {
    return {
      current: 0,
      validStep: false,
      finalValue: Object,
      ModalText: "¿Estás seguro que quieres guardar el instrumento?",
      visible: false,
      confirmLoading: false
    };
  },
  beforeCreate() {
    this.nameForm = this.$form.createForm(this, { name: "name" });
    this.questionsForm = this.$form.createForm(this, { name: "questionsForm" });
    // this.editQuestions = this.$form.createForm(this, { name: "editQuestions" });

    this.questionsForm.getFieldDecorator("keys", {
      initialValue: [],
      preserve: true
    });
    this.questionsForm.getFieldDecorator("answer1", {
      initialValue: "",
      preserve: true
    });
    this.questionsForm.getFieldDecorator("answer2", {
      initialValue: "",
      preserve: true
    });
    this.questionsForm.getFieldDecorator("answer3", {
      initialValue: "",
      preserve: true
    });
    this.questionsForm.getFieldDecorator("answer4", {
      initialValue: "",
      preserve: true
    });
    this.nameForm.getFieldDecorator("testName", {
      initialValue: "",
      preserve: true
    });
    this.nameForm.getFieldDecorator("description", {
      initialValue: "",
      preserve: true
    });
  },
  mounted() {
    if (this.edit) {
      this.nameForm.setFields({
        testName: { value: this.questionnaire.testName },
        description: { value: this.questionnaire.description }
      });

      let keys = [];
      for (const key in this.questionnaire.questions) {
        /*eslint-disable*/

        console.log(key);
        /*eslint-disable*/

        keys.push(key);
      }
      this.questionsForm.getFieldDecorator("questions", {
        initialValue: [],
        preserve: true
      });
      this.questionsForm.setFields({
        answer1: { value: this.questionnaire.answers[0] },
        answer2: { value: this.questionnaire.answers[1] },
        answer3: { value: this.questionnaire.answers[2] },
        answer4: { value: this.questionnaire.answers[3] },
        questions: { value: this.questionnaire.questions },
        keys: { value: keys }
      });
    }
  },
  methods: {
    next() {
      switch (this.current) {
        case 0:
          this.handleSubmit(this.nameForm);
          break;
        case 1:
          this.handleSubmit(this.questionsForm);
          this.finalValue = {
            testName: this.nameForm.getFieldValue("testName"),
            description: this.nameForm.getFieldValue("description"),
            questions: this.questionsForm.getFieldValue("questions"),
            answers: [
              this.questionsForm.getFieldValue("answer1"),
              this.questionsForm.getFieldValue("answer2"),
              this.questionsForm.getFieldValue("answer3"),
              this.questionsForm.getFieldValue("answer4")
            ]
          };
          break;
        case 2:
          this.showModal();
          break;
        default:
          break;
      }
      if (this.validStep) {
        this.current++;
      }
    },
    prev() {
      this.current--;
    },
    handleSubmit(form) {
      form.validateFields(err => {
        if (err) {
          this.validStep = false;
        } else {
          this.validStep = true;
        }
      });
    },
    remove(k) {
      const { questionsForm } = this;
      const keys = questionsForm.getFieldValue("keys");
      if (keys.length === 1) {
        return;
      }

      questionsForm.setFieldsValue({
        keys: keys.filter(key => key !== k)
      });
    },
    add() {
      const { questionsForm } = this;
      const keys = questionsForm.getFieldValue("keys");
      const nextKeys = keys.concat(id++);
      questionsForm.setFieldsValue({
        keys: nextKeys
      });
    },
    createTest() {
      const axios = require("axios");
      if (this.finalValue) {
        const axiosParams = {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            Accept: "application/json"
          }
        };

        return axios.post(
          `http://localhost:3000/questionnaires/create`,
          {
            testName: this.finalValue.testName,
            description: this.finalValue.description,
            questions: this.finalValue.questions,
            answers: this.finalValue.answers,
            uid: this.user.uid
          },
          axiosParams
        );
      }
    },
    editTest() {
      const axios = require("axios");
      if (this.finalValue) {
        const axiosParams = {
          headers: {
            "Content-Type": "application/x-www-form-urlencoded",
            Accept: "application/json"
          }
        };

        return axios.patch(
          `http://localhost:3000/questionnaires/edit/${this.questionnaire._id}`,
          {
            testName: this.finalValue.testName,
            description: this.finalValue.description,
            questions: this.finalValue.questions,
            answers: this.finalValue.answers,
            uid: this.user.uid
          },
          axiosParams
        );
      }
    },
    showModal() {
      this.visible = true;
    },
    handleOk() {
      this.confirmLoading = true;
      if (this.edit) {
        this.editTest()
          .then(response => {
            if (response.status === 200) {
              this.visible = false;
              this.confirmLoading = false;
              this.$emit("done");
            }
          })
          .catch(() => {
            this.visible = false;
            this.confirmLoading = false;
          });
      } else {
        this.createTest()
          .then(response => {
            if (response.status === 200) {
              this.visible = false;
              this.confirmLoading = false;
              this.$emit("done");
            }
          })
          .catch(() => {
            this.visible = false;
            this.confirmLoading = false;
          });
      }
    },
    handleCancel() {
      this.visible = false;
      this.current = 2;
    }
  }
};
</script>
<style scoped>
.steps-content {
  /* margin-top: 16px; */
  border: 1px dashed #e9e9e9;
  border-radius: 6px;
  background-color: #fafafa;
  min-height: 200px;
  /* text-align: center; */
  padding-top: 80px;
}

.steps-action {
  margin-top: 24px;
}
</style>
