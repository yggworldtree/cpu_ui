<template>
  <div id="digital-flop">
    <div
      class="digital-flop-item"
      v-for="item in digitalFlopData"
      :key="item.title"
    >
      <div class="digital-flop-title">{{ item.title }}</div>
      <div class="digital-flop">
        <dv-digital-flop
          :config="item.number"
          style="width: 100px; height: 50px"
        />
        <div class="unit">{{ item.unit }}</div>
      </div>
    </div>

    <dv-decoration-10 />
  </div>
</template>

<script>
import { queryWarnln, queryWinfoln, queryWinfos } from "@/request/api"; // 导入我们的api接口
export default {
  name: "DigitalFlop",
  data() {
    return {
      digitalFlopData: [],
    };
  },
  methods: {
    getWinfoln() {
      queryWinfoln().then((res) => {
        this.digitalFlopData.push({
          title: "总报告数",
          number: {
            number: [parseInt(res.count)],
            content: "{nt}",
            textAlign: "center",
            style: {
              fill: '#40faee',
              fontWeight: "bold",
            },
          },
        });
      });
      this.getWarnln();
    },
    getWarnln() {
      queryWarnln().then((res) => {
        this.digitalFlopData.push({
          title: "总警告数",
          number: {
            number: [parseInt(res.count)],
            content: "{nt}",
            textAlign: "center",
            style: {
              fill: '#4d99fc',
              fontWeight: "bold",
            },
          },
        });
        this.digitalFlopData.push({
          title: "今日警告数",
          number: {
            number: [parseInt(res.count_today)],
            content: "{nt}",
            textAlign: "center",
            style: {
              fill: '#f46827',
              fontWeight: "bold",
            },
          },
        });
        this.digitalFlopData.push({
          title: "昨日警告数",
          number: {
            number: [parseInt(res.count_lastday)],
            content: "{nt}",
            textAlign: "center",
            style: {
              fill: '#f46827',
              fontWeight: "bold",
            },
          },
        });
      });
    },
    createData() {
      this.digitalFlopData = [];
      this.getWinfoln();
    },
    randomExtend(minNum, maxNum) {
      if (arguments.length === 1) {
        return parseInt(Math.random() * minNum + 1, 10);
      } else {
        return parseInt(Math.random() * (maxNum - minNum + 1) + minNum, 10);
      }
    },
  },
  mounted() {
    const { createData } = this;

    createData();

    setInterval(createData, 300000);
  },
};
</script>

<style lang="less">
#digital-flop {
  position: relative;
  height: 15%;
  flex-shrink: 0;
  display: flex;
  justify-content: space-between;
  align-items: center;
  background-color: rgba(6, 30, 93, 0.5);

  .dv-decoration-10 {
    position: absolute;
    width: 95%;
    left: 2.5%;
    height: 5px;
    bottom: 0px;
  }

  .digital-flop-item {
    width: 11%;
    height: 80%;
    display: flex;
    flex-direction: column;
    justify-content: center;
    align-items: center;
    border-left: 3px solid rgb(6, 30, 93);
    border-right: 3px solid rgb(6, 30, 93);
  }

  .digital-flop-title {
    font-size: 20px;
    margin-bottom: 20px;
  }

  .digital-flop {
    display: flex;
  }

  .unit {
    margin-left: 10px;
    display: flex;
    align-items: flex-end;
    box-sizing: border-box;
    padding-bottom: 13px;
  }
}
</style>
