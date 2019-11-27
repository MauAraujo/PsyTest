<template>
  <a-list
    class="loadmore-list"
    :loading="loading"
    itemLayout="horizontal"
    :dataSource="users"
  >
    <a-list-item slot="renderItem" slot-scope="user">
      <a-icon class="action-icon" slot="actions" type="edit" @click="onEdit(user)" />
      <a-icon class="action-icon" slot="actions" type="delete" @click="showModal()" />
      <a-list-item-meta :description="user.type">
        <a slot="title">{{user.firstName}}</a>
      </a-list-item-meta>
      <a-modal
        title="Confirmar eliminación de usuario"
        :visible="visible"
        @ok="handleOk(user)"
        :confirmLoading="confirmLoading"
        @cancel="handleCancel"
      >
        <template slot="footer">
          <a-button key="back" @click="handleCancel">Regresar</a-button>
          <a-button key="submit" type="primary" :loading="loading" @click="handleOk(user)">Aceptar</a-button>
        </template>
        <p>{{ModalText}}</p>
      </a-modal>
    </a-list-item>
    <a-empty
      v-if="users.length === 0"
      image="https://gw.alipayobjects.com/mdn/miniapp_social/afts/img/A*pevERLJC9v0AAAAAAAAAAABjAQAAAQ/original"
    >
      <span slot="description">No hay usuarios. Cuando se cree uno aquí aparecerá.</span>
    </a-empty>
  </a-list>
</template>
<script>
export default {
  data() {
    return {
      loading: true,
      loadingMore: false,
      showLoadingMore: true,
      users: [],
      ModalText: "¿Estás seguro que quieres eliminar la información del paciente?",
      visible: false,
      confirmLoading: false
    };
  },
  mounted() {
    this.getData();
  },
  methods: {
    /*eslint-disable*/
    onDelete(user) {
      const axios = require("axios");
      console.log(user);
      const axiosParams = {
        headers: {
          "Content-Type": "application/x-www-form-urlencoded",
          Accept: "application/json"
        }
      };

      axios
        .delete(
          `http://localhost:3000/users/delete/${user._id}`,
          {
            qid: user.id
          },
          axiosParams
        )
        .then(response => {
          console.log(response);
          this.visible = false;
          // this.getData();
        })
        .catch(error => {
          console.log(error);
        });
    },
    onEdit(user) {
      this.$emit("edit", user);
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
        .get(`http://localhost:3000/users/get`, axiosParams)
        .then(response => {
          console.log(response);
          this.loading = false;
          this.users = response.data !== null ? response.data : [];
        })
        .catch(error => {
          console.log(error);
        });
      /*eslint-disable*/
    },
    handleOk(user) {
      console.log(user);
      this.confirmLoading = true;
      this.onDelete(user);
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