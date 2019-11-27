<template>
  <a-list
    class="loadmore-list"
    :loading="loading"
    itemLayout="horizontal"
    :dataSource="questionnaires"
  >
    <a-list-item slot="renderItem" slot-scope="item">
      <a-icon class="action-icon" slot="actions" type="eye" @click="onSelect(item)" />
      <a-icon class="action-icon" slot="actions" type="edit" @click="onEdit(item)" />
      <a-icon class="action-icon" slot="actions" type="delete" @click="showModal()" />
      <a-list-item-meta :description="item.description">
        <a slot="title">{{item.testName}}</a>
      </a-list-item-meta>
      <a-modal
        title="Confirmar eliminación de instrumento"
        :visible="visible"
        @ok="handleOk(item)"
        :confirmLoading="confirmLoading"
        @cancel="handleCancel"
      >
        <template slot="footer">
          <a-button key="back" @click="handleCancel">Regresar</a-button>
          <a-button key="submit" type="primary" :loading="loading" @click="handleOk(item)">Aceptar</a-button>
        </template>
        <p>{{ModalText}}</p>
      </a-modal>
    </a-list-item>
    <a-empty
      v-if="questionnaires.length === 0"
      image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
    >
      <span slot="description">No hay instrumentos. Cuando se cree uno aquí aparecerá.</span>
    </a-empty>
  </a-list>
</template>
<script>
export default {
  props: {
    user: Object
  },
  data() {
    return {
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
      questionnaires: [],
      ModalText: "¿Estás seguro que quieres eliminar el instrumento?",
      visible: false,
      confirmLoading: false
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    /*eslint-disable*/
    onDelete(questionnaire) {
      const axios = require("axios");
      console.log(questionnaire);
      const axiosParams = {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        }
      };

      axios
        .delete(
          `http://localhost:3000/questionnaires/delete/${questionnaire._id}`,
          axiosParams
        )
        .then(response => {
          console.log(response);
          this.visible = false;
          this.getData();
        })
        .catch(error => {
          console.log(error);
        });
    },
    onEdit(questionnaire) {
      this.$emit("edit", questionnaire);
    },
    onSelect(questionnaire) {
      this.$emit("get", questionnaire);
    },
    getData() {
      const axios = require("axios");
      const axiosParams = {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        }
      };
      /*eslint-disable*/
      axios
        .get(
          `http://localhost:3000/questionnaires/get`,
          axiosParams
        )
        .then(response => {
          console.log(response);
          this.loading = false;
          this.questionnaires = response.data !== null ? response.data : [];
        })
        .catch(error => {
          console.log(error);
        });
      /*eslint-disable*/
    },
    handleOk(questionnaire) {
      console.log(questionnaire);
      this.confirmLoading = true;
      this.onDelete(questionnaire);
    },
    handleCancel() {
      this.visible = false;
      this.current = 2;
    },
    showModal() {
      this.visible = true;
    }
    /*eslint-disable*/
  }
};
</script>
<style>
.loadmore-list {
  min-height: 350px;
}

.action-icon:hover,
.action-icon:focus {
  color: #40a9ff;
  background-color: #fff;
  border-color: #40a9ff;
}
</style>