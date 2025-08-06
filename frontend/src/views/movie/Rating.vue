<template>
<!--  <el-table-->
<!--    :data="tableData"-->
<!--    style="width: 100%">-->
<!--    <el-table-column-->
<!--      type="index"-->
<!--      :index="indexMethod">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="date"-->
<!--      label="日期"-->
<!--      width="180">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="name"-->
<!--      label="姓名"-->
<!--      width="180">-->
<!--    </el-table-column>-->
<!--    <el-table-column-->
<!--      prop="address"-->
<!--      label="地址">-->
<!--    </el-table-column>-->
<!--  </el-table>-->


  <div>
    <div class="pie" style="height: 400px" ref="chartEducation"></div>
    <div class="review-container" style="max-height: none;">
      <div class="lister-item-content">
        <div class="ipl-ratings-bar">
              <span class="rating-other-user-rating">
              <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">
                  <path d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>
                  <path d="M0 0h24v24H0z" fill="none"></path>
              </svg>
                  <span>9</span><span class="point-scale">/10</span>
              </span>
        </div>
        <a href="/review/rw1488015/?ref_=tt_urv" class="title"> An excellent show about really nasty people and the heinous things they do for money. Big money.
        </a>            <div class="display-name-date">
        <span class="display-name-link"><a href="/user/ur12122603/?ref_=tt_urv">swetepete</a></span><span class="review-date">28 September 2006</span>
      </div>
        <div class="content">
          <div class="text show-more__control">This is one of the best crime dramas on TV, maybe almost as good as early 'Sopranos' or 'Shield.' It has a wicked soundtrack, it's beautifully shot and choreographed, and sports some of the best acting I've seen on a network show. The characters are, basically, incorrigible scum; but, they pull off their heists with such style, I find myself very much hoping they stick around for a couple dozen episodes. No doubt most or all of them will meet nasty ends, but in the meantime, 'Smith' is a good place to indulge in some cathartic mayhem and villainy. The soundtrack is phenomenal--Rancid, Dirty Pretty Things, Imogen Heap, Thom Yorke, and the Ramones have all popped up early on in the series. I expect the music to stay high-end as the first season progresses. 'Smith' is not for an audience who needs to cheer for the forces of law and order. It's all about the bad guys, and not saying sorry.</div>
          <div class="actions text-muted">
            46 out of 51 found this helpful.
            <span>
                              Was this review helpful? <a href="/registration/signin?ref_=urv"> Sign in</a> to vote.
                          </span>
            <br>
            <a href="/review/rw1488015/?ref_=tt_urv">Permalink</a>
          </div>
        </div>
      </div>
      <div class="clear"></div>
      <div class="lister-item-content">
        <div class="ipl-ratings-bar">
              <span class="rating-other-user-rating">
              <svg class="ipl-icon ipl-star-icon  " xmlns="http://www.w3.org/2000/svg" fill="#000000" height="24" viewBox="0 0 24 24" width="24">
                  <path d="M0 0h24v24H0z" fill="none"></path>
                  <path d="M12 17.27L18.18 21l-1.64-7.03L22 9.24l-7.19-.61L12 2 9.19 8.63 2 9.24l5.46 4.73L5.82 21z"></path>
                  <path d="M0 0h24v24H0z" fill="none"></path>
              </svg>
                  <span>9</span><span class="point-scale">/10</span>
              </span>
        </div>
        <a href="/review/rw2138436/?ref_=tt_urv" class="title"> strange that CBS canceled this after only one season
        </a>            <div class="display-name-date">
        <span class="display-name-link"><a href="/user/ur9298705/?ref_=tt_urv">tom-rusch</a></span><span class="review-date">9 October 2009</span>
      </div>
        <div class="content">
          <div class="text show-more__control clickable">after watching many of the balance of the episodes on DirecTV's 101, i'm puzzled why this didn't get more that a showing of the pilot.<br><br>interesting characters. definite possibilities for the direction the show might have gone. <br><br>all the characters might have great back stories that would give the writers an option to move beyond the mere caper plot of other similar series.<br><br>the second to the last episode on 101 suggested the background for simon baker's character and suggested a little redemptive option.<br><br>and the same episode hints at the history of ray liotta's and virginia madsen's troubled relationship.<br><br>ah well, another example of 'if i like it, they'll cancel it.'</div>

        </div>
      </div>
      <div class="clear"></div>
    </div>
  </div>





</template>

<script>
import * as echarts from "echarts";

export default {
  mounted() {
    this.$nextTick(() => {
      this.initCharts4sEducation();
    });
  },
  methods: {
    initCharts4sEducation() {
      let myChart = echarts.init(this.$refs.chartEducation);
      myChart.setOption({
        title: "",
        tooltip: {},
        legend: {
          orient: "horizontal",
          x: "right",
          y: "bottom",
          itemWidth: 18, // 设置图例图形的宽
          itemHeight: 30, // 设置图例图形的高
          itemGap: 30,
          backgroundColor: "#eee", // 设置整个图例区域背景颜色
          data: ["5 Stars", "4 Stars", "3 Stars", "2 Stars", "1 Star"],
        },
        series: [
          {
            name: "User Ratings",
            type: "pie", // 设置图表类型为饼图
            radius: ["55%", "65%"], // 饼图的半径，外半径为可视区尺寸（容器高宽中较小一项）的 55% 长度。
            data: [
              // 数据数组，name 为数据项名称，value 为数据项值
              { value: 30, name: "5 Stars" },
              { value: 40, name: "4 Stars" },
              { value: 20, name: "3 Stars" },
              { value: 10, name: "2 Stars" },
              { value: 10, name: "1 Star" },
            ],
            color: ["green", "blue", "Cyan", "red", "yellow"], //各个区域颜色
            itemStyle: {
              normal: {
                label: {
                  show: true,
                  formatter: "{b} {d}%",
                }, //饼图图形上的文本标签，可用于说明图形的一些数据信息，比如值，名称等。可以与itemStyle属性同级，具体看文档
                labelLine: {
                  show: true,
                },
              },
            },
          },
        ],
      });
    },
  },
  data() {
    return {
      // initCharts4sEducation() {
      //   let myChart = echarts.init(this.$refs.chartEducation);
      //   myChart.setOption({
      //     title: "",
      //     tooltip: {},
      //     legend: {
      //       orient: "horizontal",
      //       x: "right",
      //       y: "top",
      //       itemWidth: 18, // 设置图例图形的宽
      //       itemHeight: 30, // 设置图例图形的高
      //       itemGap: 30,
      //       backgroundColor: "#eee", // 设置整个图例区域背景颜色
      //       data: ["大学及以上", "高中及以下", "初中及以下", "小学及以下"],
      //     },
      //     series: [
      //       {
      //         name: "教育程度",
      //         type: "pie", // 设置图表类型为饼图
      //         radius: ["55%", "65%"], // 饼图的半径，外半径为可视区尺寸（容器高宽中较小一项）的 55% 长度。
      //         data: [
      //           // 数据数组，name 为数据项名称，value 为数据项值
      //           { value: 30, name: "大学及以上" },
      //           { value: 40, name: "高中及以下" },
      //           { value: 20, name: "初中及以下" },
      //           { value: 10, name: "小学及以下" },
      //         ],
      //         color: ["green", "blue", "Cyan", "red"], //各个区域颜色
      //         itemStyle: {
      //           normal: {
      //             label: {
      //               show: true,
      //               formatter: "{b} {d}%",
      //             }, //饼图图形上的文本标签，可用于说明图形的一些数据信息，比如值，名称等。可以与itemStyle属性同级，具体看文档
      //             labelLine: {
      //               show: true,
      //             },
      //           },
      //         },
      //       },
      //     ],
      //   });
      // },
      tableData: [{
        date: '2016-05-02',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1518 弄',
        zip: 200333,
        tag: '家'
      }, {
        date: '2016-05-04',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1517 弄',
        zip: 200333,
        tag: '公司'
      }, {
        date: '2016-05-01',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1519 弄',
        zip: 200333,
        tag: '家'
      }, {
        date: '2016-05-03',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1516 弄',
        zip: 200333,
        tag: '公司'
      },{
        date: '2016-05-03',
        name: '王小虎',
        province: '上海',
        city: '普陀区',
        address: '上海市普陀区金沙江路 1516 弄',
        zip: 200333,
        tag: '公司'
      }],
    }
  },
  // methods: {
  //   indexMethod(index) {
  //     return index +1 ;
  //   }
  // }
};
</script>

<style scoped>
.chartEducation h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}


</style>
