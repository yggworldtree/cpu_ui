<template>
  <div id="water-level-chart">
    <div class="water-level-chart-title">内存使用状况</div>
    <div class="chart-container">
      <dv-water-level-pond :config="config" />
    </div>
    <div class="water-level-chart-title">总内存:{{ this.memoryTotal }}</div>
  </div>
</template>

<script>
import { cpuMem } from "@/request/api"; // 导入我们的api接口
export default {
  name: "WaterLevelChart",
  mounted() {
    const { createData } = this;
    createData();
    setInterval(this.createData, 10000);
  },
  methods: {
    change(limit) {
      var size = "";
      if (limit < 0.1 * 1024) {
        //小于0.1KB，则转化成B
        size = limit.toFixed(2) + "B";
      } else if (limit < 0.1 * 1024 * 1024) {
        //小于0.1MB，则转化成KB
        size = (limit / 1024).toFixed(2) + "KB";
      } else if (limit < 0.1 * 1024 * 1024 * 1024) {
        //小于0.1GB，则转化成MB
        size = (limit / (1024 * 1024)).toFixed(2) + "MB";
      } else {
        //其他转化成GB
        size = (limit / (1024 * 1024 * 1024)).toFixed(2) + "GB";
      }

      var sizeStr = size + ""; //转成字符串
      var index = sizeStr.indexOf("."); //获取小数点处的索引
      var dou = sizeStr.substr(index + 1, 2); //获取小数点后两位的值
      if (dou == "00") {
        //判断后两位是否为00，如果是则删除00
        return sizeStr.substring(0, index) + sizeStr.substr(index + 3, 2);
      }
      return size;
    },
    createData() {
      cpuMem().then((res) => {
        this.config = {
          data: [
            Math.floor((res.virtualMem.used / res.virtualMem.total) * 100),
          ],
          formatter:"{value}%",
          shape: "round",
          waveHeight: 25,
          waveNum: 2,
        };
        this.memoryTotal = this.change(17179869184);
      });
    },
  },
  data() {
    return {
      memoryTotal: "",
      config: {
        data: [45],
        shape: "round",
        waveHeight: 25,
        waveNum: 2,
      },
    };
  },
};
</script>

<style lang="less">
#water-level-chart {
  width: 20%;
  box-sizing: border-box;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, 0.5);
  display: flex;
  flex-direction: column;

  .water-level-chart-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
    justify-content: center;
  }

  .water-level-chart-details {
    height: 15%;
    display: flex;
    justify-content: center;
    font-size: 17px;
    align-items: flex-end;

    span {
      font-size: 35px;
      font-weight: bold;
      color: #58a1ff;
      margin: 0 5px;
      margin-bottom: -5px;
    }
  }

  .chart-container {
    flex: 1;
    display: flex;
    justify-content: center;
    align-items: center;
  }

  .dv-water-pond-level {
    max-width: 90%;
    width: 200px;
    height: 200px;
    border: 10px solid #19c3eb;
    border-radius: 50%;

    ellipse {
      stroke: transparent !important;
    }

    text {
      font-size: 40px;
    }
  }
}
</style>
