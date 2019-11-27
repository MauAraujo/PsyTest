<template>
  <a-form :form="form" @submit="handleSubmit">
    <a-form-item v-bind="formItemLayout" label="Nombre">
      <a-input
        v-decorator="[
          'name',
          {
            rules: [
              {
                required: true,
                whitespace: true,
                message: 'Ingresa el nombre.',
              },
            ],
          },
        ]"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="Apellido Paterno">
      <a-input
        v-decorator="[
          'middleName',
          {
            rules: [
              {
                required: true,
                message: 'Ingresa el apellido paterno.',
              },
            ],
          },
        ]"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="Apellido Materno">
      <a-input
        v-decorator="[
          'lastName',
          {
            rules: [
              {
                required: true,
                message: 'Ingresa el apellido materno.',
              },
            ],
          },
        ]"
      />
    </a-form-item>
    <a-row>
      <a-col :span="12">
        <a-form-item
          :label-col="{ span: 10 }"
          :wrapper-col="{ span: 12 }"
          v-bind="formItemLayout"
          label="Edad"
        >
          <a-input-number
            :min="0"
            v-decorator="[
          'age',
          {
            rules: [{ required: true, message: 'Ingresa la edad' }],
          },
        ]"
          />
        </a-form-item>
      </a-col>
      <a-col :span="12">
        <a-form-item
          :label-col="{ span: 6 }"
          :wrapper-col="{ span: 12 }"
          v-bind="formItemLayout"
          label="Género"
        >
          <a-select
            v-decorator="['gender',
            {
                rules: [{ required: true, message: 'Ingresa la edad' }]
            },
            { initialValue: '86' }]"
            style="width: 70px"
          >
            <a-select-option value="M">M</a-select-option>
            <a-select-option value="F">F</a-select-option>
          </a-select>
        </a-form-item>
      </a-col>
    </a-row>
    <a-form-item v-bind="formItemLayout" label="Nombre de usuario">
      <a-input
        v-decorator="[
          'username',
          {
            rules: [
              {
                required: true,
                message: 'Ingresa el nombre de usuario',
              },
            ],
          },
        ]"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="Contraseña">
      <a-input
        v-decorator="[
          'password',
          {
            rules: [
              {
                required: true,
                message: 'Ingresa la contraseña.',
              },
              {
                validator: validateToNextPassword,
              },
            ],
          },
        ]"
        type="password"
      />
    </a-form-item>
    <a-form-item v-bind="formItemLayout" label="Confirma la contraseña">
      <a-input
        v-decorator="[
          'confirm',
          {
            rules: [
              {
                required: true,
                message: 'Confirma la contraseña.',
              },
              {
                validator: compareToFirstPassword,
              },
            ],
          },
        ]"
        type="password"
        @blur="handleConfirmBlur"
      />
    </a-form-item>
    <a-form-item v-bind="tailFormItemLayout">
      <a-button type="primary" html-type="submit">Registrar usuario</a-button>
    </a-form-item>
  </a-form>
</template>

<script>
export default {
  data() {
    return {
      confirmDirty: false,
      autoCompleteResult: [],
      formItemLayout: {
        labelCol: {
          xs: { span: 24 },
          sm: { span: 5 }
        },
        wrapperCol: {
          xs: { span: 24 },
          sm: { span: 16 }
        }
      },
      tailFormItemLayout: {
        wrapperCol: {
          xs: {
            span: 24,
            offset: 0
          },
          sm: {
            span: 16,
            offset: 8
          }
        }
      }
    };
  },
  beforeCreate() {
    this.form = this.$form.createForm(this, { name: "register" });
  },
  methods: {
    handleSubmit(e) {
      e.preventDefault();
      this.form.validateFieldsAndScroll((err, values) => {
        if (!err) {
          /*eslint-disable*/

          console.log("Received values of form: ", values);
          /*eslint-disable*/
        }
      });
    },
    handleConfirmBlur(e) {
      const value = e.target.value;
      this.confirmDirty = this.confirmDirty || !!value;
    },
    compareToFirstPassword(rule, value, callback) {
      const form = this.form;
      if (value && value !== form.getFieldValue("password")) {
        callback("Las contraseñas no coinciden.");
      } else {
        callback();
      }
    },
    validateToNextPassword(rule, value, callback) {
      const form = this.form;
      if (value && this.confirmDirty) {
        form.validateFields(["confirm"], { force: true });
      }
      callback();
    }
  }
};
</script>