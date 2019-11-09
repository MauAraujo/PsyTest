<template>
  <div class="main">
    <a-row>
      <a-col :span="24">
        <a-card style="width: 600px;">
          <p slot="extra">Me da miedo dormir solo</p>
          <a-radio-group @change="onChange" v-model="value">
            <a-radio :style="radioStyle" :value="1">No me siento triste.</a-radio>
            <a-radio :style="radioStyle" :value="2">Me siento triste.</a-radio>
            <a-radio
              :style="radioStyle"
              :value="3"
            >Me siento triste continuamente y no puedo dejar de estarlo.</a-radio>
            <a-radio
              :style="radioStyle"
              :value="4"
            >Me siento tan triste o tan desgraciado que no puedo soportarlo.</a-radio>
          </a-radio-group>
        </a-card>
      </a-col>
    </a-row>
  </div>
</template>

<script>
export default {
  data() {
    return {
      value: 1,
      radioStyle: {
        display: "flex",
        height: "30px",
        lineHeight: "30px"
      }
    };
  },
  methods: {
    onChange() {}
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
</style>
