<template>
  <div>
    <a-list itemLayout="horizontal" :dataSource="questionnaires">
      <a-list-item slot="renderItem" slot-scope="item">
        <a-list-item-meta
          description="Ant Design, a design language for background applications, is refined by Ant UED Team"
        >
          <a slot="title" href="https://vue.ant.design/">{{item.testName}}</a>
          <a-avatar
            slot="avatar"
            src="https://zos.alipayobjects.com/rmsportal/ODTLcjxAfvqbxHnVXCYX.png"
          />
        </a-list-item-meta>
      </a-list-item>
    </a-list>
  </div>
</template>
<script>
export default {
  props: {
    user: Object
  },
  data() {
    return {
      questionnaires: []
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
        this.questionnaires = response.data;
      })
      .catch(error => {
        console.log(error);
      });
    /*eslint-disable*/
  }
};
</script>
<style>
</style>