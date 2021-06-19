<template>
  <div id="scroll-board">
    <dv-scroll-board :config="config" />
  </div>
</template>

<script>
import { cpuProcs } from "@/request/api"; // 导入我们的api接口
export default {
  name: "ScrollBoard",
  mounted: function () {
    const { getCpuProcs } = this;
    getCpuProcs();
    setInterval(this.getCpuProcs, 1000000);
  },
  methods: {
    getCpuProcs() {
      cpuProcs().then((res) => {
        var datas = [];
        var i = 0;
        res.forEach((element) => {
          var elements = [];
          elements[0] = element.CommandName;
          elements[1] = element.Cpu;
          elements[2] = element.Pid;
          elements[3] = element.User;
          datas[i] = elements;
          i++;
        });
        this.config = {
          header: ["CommandName", "Cpu占用率", "Pid", "User"],
          data: datas,
          index: true,
          columnWidth: [50, 170, 170,170,170],
          align: ["center"],
          rowNum: 7,
          headerBGC: "#1981f6",
          headerHeight: 45,
          oddRowBGC: "rgba(0, 44, 81, 0.8)",
          evenRowBGC: "rgba(10, 29, 50, 0.8)",
        };
      });
    },
  },
  data() {
    return {
      config: {
          header: ["CommandName", "Cpu占用率", "Pid", "User"],
          data: [],
          index: true,
          columnWidth: [50, 170, 170,170,170],
          align: ["center"],
          rowNum: 7,
          headerBGC: "#1981f6",
          headerHeight: 45,
          oddRowBGC: "rgba(0, 44, 81, 0.8)",
          evenRowBGC: "rgba(10, 29, 50, 0.8)",
      },
    };
  },
};
</script>

<style lang="less">
#scroll-board {
  width: 50%;
  box-sizing: border-box;
  margin-left: 20px;
  height: 100%;
  overflow: hidden;
}
</style>
