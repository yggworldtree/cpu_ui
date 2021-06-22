<template>
  <div id="rose-chart">
    <div class="rose-chart-title">累计计量资金分布</div>
    <dv-capsule-chart :config="config" style="width: 100%; height: 100%" />
  </div>
</template>

<script>
import { queryWarnln, queryWinfoln, queryWinfos } from "@/request/api"; // 导入我们的api接口
export default {
  name: "RoseChart",
  data() {
    return {
      datas: [],
      config: {},
    };
  },
  methods: {
    createData() {
      console.log(this.datas);
      let dataa = this.datas;
      this.config = {
        data:dataa,
        colors: ["#e062ae", "#fb7293", "#e690d1", "#32c5e9", "#96bfff"],
        unit: "单位",
        showValue: true,
      };
      this.datas = [];
    },

  },
  mounted() {
    const { getWinfos } = this;

    getWinfos();

    setInterval(getWinfos, 30000);
  },
};
</script>

<style lang="less">
#rose-chart {
  width: 30%;
  height: 100%;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, 0.5);
  box-sizing: border-box;

  .rose-chart-title {
    height: 50px;
    font-weight: bold;
    text-indent: 20px;
    font-size: 20px;
    display: flex;
    align-items: center;
  }

  .dv-charts-container {
    height: calc(~"100% - 50px");
  }
}
</style>
