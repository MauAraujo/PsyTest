<template>
  <div>
    <a-carousel arrows dotsClass="slick-dots slick-thumb">
      <div slot="prevArrow" class="custom-slick-arrow" style="left: 10px;zIndex: 1">
        <a-icon type="left-circle" />
      </div>
      <div slot="nextArrow" class="custom-slick-arrow" style="right: 10px">
        <a-icon type="right-circle" />
      </div>

      <div>
        <a-card title="Default size card" style="width: 300px">
          <a href="#" slot="extra">more</a>
          <p>card content</p>
          <p>card content</p>
          <p>card content</p>
        </a-card>
      </div>
      <div>
        <a-card title="Default size card" style="width: 300px">
          <a href="#" slot="extra">more</a>
          <p>card content</p>
          <p>card content</p>
          <p>card content</p>
        </a-card>
      </div>
      <div>
        <a-card title="Default size card" style="width: 300px">
          <a href="#" slot="extra">more</a>
          <p>card content</p>
          <p>card content</p>
          <p>card content</p>
        </a-card>
      </div>
      <div>
        <a-card title="Default size card" style="width: 300px">
          <a href="#" slot="extra">more</a>
          <p>card content</p>
          <p>card content</p>
          <p>card content</p>
        </a-card>
      </div>
    </a-carousel>
  </div>
</template>

<script>
export default {};
</script>

<style scoped>
.ant-carousel >>> .slick-slide {
  text-align: center;
  height: 100vh;
  line-height: 160px;
  background: #e8f8f8;
  overflow: hidden;
}

.ant-carousel >>> li {
  background: rgb(0, 0, 0);
}

.ant-carousel >>> .custom-slick-arrow {
  width: 25px;
  height: 25px;
  font-size: 25px;
  color: rgb(0, 186, 186);
  background-color: rgba(255, 255, 255, 0.11);
  opacity: 0.6;
}
.ant-carousel >>> .custom-slick-arrow:before {
  display: none;
}
.ant-carousel >>> .custom-slick-arrow:hover {
  color: rgba(0, 186, 186);
  opacity: 1;
}

.ant-carousel >>> .slick-slide .ant-card {
  color: rgb(0, 0, 0);
  position: relative;
  top: 20em;
  left: 40%;
}
</style>
