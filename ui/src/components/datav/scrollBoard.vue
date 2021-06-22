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
  },
  methods: {
    formatDate(date, fmt) {
      if (/(y+)/.test(fmt)) {
        fmt = fmt.replace(
          RegExp.$1,
          (date.getFullYear() + "").substr(4 - RegExp.$1.length)
        );
      }
      let o = {
        "M+": date.getMonth() + 1,
        "d+": date.getDate(),
        "h+": date.getHours(),
        "m+": date.getMinutes(),
        "s+": date.getSeconds(),
      };
      for (let k in o) {
        if (new RegExp(`(${k})`).test(fmt)) {
          let str = o[k] + "";
          fmt = fmt.replace(
            RegExp.$1,
            RegExp.$1.length === 1 ? str : this.padLeftZero(str)
          );
        }
      }
      return fmt;
    },

    padLeftZero(str) {
      return ("00" + str).substr(str.length);
    },
    getCpuProcs() {
      cpuProcs().then((res) => {
        var datas = [];
        var i = 0;
        res.forEach((element) => {
          var elements = [];
          elements[0] = element.user;
          elements[1] = element.pid;
          elements[2] = parseFloat(Number(element.cpu).toFixed(2)) + "%";
          elements[3] = parseFloat(Number(element.mem).toFixed(2)) + "%";
          elements[4] = this.formatDate(new Date(element.createTime), 'yyyy-MM-dd hh:mm');
          elements[5] =
            element.commandName +
            (element.commandLine != null && element.commandLine.length > 0
              ? " [" + element.commandLine.join(" ") + "]"
              : "");
          datas[i] = elements;
          i++;
        });
        this.config = {
          header: ["User", "Pid", "Cpu占用率", "内存占用率", "Time", "command"],
          data: datas,
          index: true,
          columnWidth: [100, 80, 80, 150, 150, 150, 500],
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
      config: {},
    };
  },
};
</script>

<style lang="less">
#scroll-board {
  width: 100%;
  box-sizing: border-box;
  margin-left: 20px;
  height: 100%;
  overflow: hidden;
}
</style>
