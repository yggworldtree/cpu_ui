<template>
  <div id="ranking-board">
    <div class="ranking-board-title">报告</div>
    <dv-scroll-ranking-board :config="config" />
  </div>
</template>

<script>
import { queryWinfos } from "@/request/api"; // 导入我们的api接口
export default {
  name: "RankingBoard",
  mounted: function () {
    const { getWinfos } = this;
    getWinfos();
  },
  methods: {
    getWinfos() {
      queryWinfos().then((res) => {
        var datas = [];
        let i = 0;
        res.forEach((element) => {
          if (i > 9) {
            return;
          }
          var elements = {
            name: this.chooseType(element.type),
            value: element.warnLen,
          };
          datas.push(elements);
          i++;
        });
        this.config = {
          data: datas,
          rowNum: 10,
        };
      });
    },
    chooseType(type) {
      if (type == "day") {
        return "日报";
      }
      return "报告";
    },
  },
  data() {
    return {
      config: {
        data: [],
        rowNum: 10,
      },
    };
  },
};
</script>

<style lang="less">
#ranking-board {
  width: 20%;
  box-shadow: 0 0 3px blue;
  display: flex;
  flex-direction: column;
  background-color: rgba(6, 30, 93, 0.5);
  border-top: 2px solid rgba(1, 153, 209, 0.5);
  box-sizing: border-box;
  padding: 0px 30px;

  .ranking-board-title {
    font-weight: bold;
    height: 50px;
    display: flex;
    align-items: center;
    font-size: 20px;
  }

  .dv-scroll-ranking-board {
    flex: 1;
  }
}
</style>
