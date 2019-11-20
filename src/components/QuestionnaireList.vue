<template>
  <a-list
    class="loadmore-list"
    :loading="loading"
    itemLayout="horizontal"
    :dataSource="questionnaires"
  >
    <a-empty :v-if="questionnaires.length === 0" image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original">
      <span>No hay ningún instrumento. Crea uno y aparecerá aquí.</span>
    </a-empty>
    <a-list-item slot="renderItem" slot-scope="item">
      <a-icon class="action-icon" slot="actions" type="eye" />
      <a-icon class="action-icon" slot="actions" type="edit" />
      <a-icon class="action-icon" slot="actions" type="delete" @click="showModal()" />
      <a-list-item-meta :description="item.description">
        <a slot="title" href="https://vue.ant.design/">{{item.testName}}</a>
      </a-list-item-meta>
    </a-list-item>
    <a-modal
      title="Confirmar eliminación de instrumento"
      :visible="visible"
      @ok="handleOk(item)"
      :confirmLoading="confirmLoading"
      @cancel="handleCancel"
    >
      <p>{{ModalText}}</p>
    </a-modal>
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
      url: `http://localhost:3000/user/${this.user.uid}/questionnaires/`,
      ModalText: "¿Estás seguro que quieres eliiminar el instrumento?",
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
      const axiosParams = {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        }
      };

      axios
        .delete(
          `http://localhost:3000/user/${this.user.uid}/questionnaires/${questionnaire._id}`,
          axiosParams
        )
        .then(response => {
          console.log(response);
        })
        .catch(error => {
          console.log(error);
        });
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
          `http://localhost:3000/user/${this.user.uid}/questionnaires`,
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