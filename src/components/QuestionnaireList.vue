<template>
  <a-list
    class="loadmore-list"
    :loading="loading"
    itemLayout="horizontal"
    :dataSource="questionnaires"
  >
    <a-list-item slot="renderItem" slot-scope="item">
      <a-icon class="action-icon" slot="actions" type="eye" />
      <a-icon class="action-icon" slot="actions" type="edit" />
      <a-icon class="action-icon" slot="actions" type="delete" @click="onDelete(item)" />
      <a-list-item-meta :description="item.description">
        <a slot="title" href="https://vue.ant.design/">{{item.testName}}</a>
      </a-list-item-meta>
    </a-list-item>
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
      url: `http://localhost:3000/user/${this.user.uid}/questionnaires/`
    };
  },
  mounted() {
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
        this.questionnaires = response.data;
      })
      .catch(error => {
        console.log(error);
      });
    /*eslint-disable*/
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