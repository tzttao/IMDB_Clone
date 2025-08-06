<template>
  <div class="block">
    <!-- <span
      class="demonstration"
      style="font-weight: bold; color: #fcd738; font-size: 200%"
    >
      Feature Today</span> -->
    <!-- <el-carousel :interval="4000" indicator-position="outside" height="550px">
      <el-carousel-item v-for="item in imgs1" :key="item.id">
        <el-image :src="item.img"  fit="auto" />
      </el-carousel-item>
    </el-carousel> -->

    <span
      class="demonstration"
      style="font-weight: bold; color: #fcd738; font-size: 200%"
    >
      Top Picks</span>
    <el-carousel :interval="4000" type="card" height="550px">
      <el-carousel-item v-for="item in imgs1" :key="item.id">
        <el-image :src="item.img"  :fit="cover"></el-image>
      </el-carousel-item>
    </el-carousel>
  </div>
</template>

<script>
export default {
  name: 'img_show',
  data() {
    return {
      imgs1: [
        {
          id: 1, img: require('../../assets/doctor-strange-2-mcu-box-office.jpg')
        },
        {
          id: 2, img: require('../../assets/turningRed.jpg')
        },
        {
          id: 3, img: require('../../assets/justiceLeague.jpg')
        },
        {
          id: 4, img: require('../../assets/3-idiots.jpg')
        },
        {
          id: 5, img: require('../../assets/inception1.jpg')
        },
        {
          id: 6 , img: require('../../assets/uncharted.jpg')
        },
        {
          id: 7, img: require('../../assets/theWolf.jpg')
        },
        {
          id: 8, img: require('../../assets/Call-Me-By-Your-Name-Header.jpg')
        },
        {
          id: 9, img: require('../../assets/moonfall.jpeg')
        }
      ],
      // imgs2 : [
      //   {
      //     id: 1, img: require('../../assets/theShawshankRedemption.jpg')
      //   },
      //   {
      //     id: 2, img: require('../../assets/TheGodfather.jpg')
      //   },
      //   {
      //     id: 3, img: require('../../assets/TheDarkKnight.jpg')
      //   },
      //   {
      //     id: 4, img: require('../../assets/12_Angry_Men_(1957_film_poster).jpg')
      //   },
      //   {
      //     id: 5, img: require('../../assets/SchindlerList.jpg')
      //   },
      //   {
      //     id: 6, img: require('../../assets/TheLord.jpg')
      //   },
      //   {
      //     id: 7, img: require('../../assets/truman.jpg')
      //   },
      //   {
      //     id: 8, img: require('../../assets/forrest.jpg')
      //   },
      //   {
      //     id: 9, img: require('../../assets/breakingBad.jpg')
      //   },
      //   {
      //     id: 10, img: require('../../assets/inglouriousBasterds.jpg')
      //   }
      // ]
    }
  },
  methods: {}
}
</script>

<style >
.el-carousel__item h3 {
  color: #475669;
  font-size: 18px;
  opacity: 0.75;
  line-height: 300px;
  margin: 0;
}
.el-carousel__item:nth-child(2n) {
  background-color: #99a9bf;
}
.el-carousel__item:nth-child(2n + 1) {
  background-color: #d3dce6;
}
.time {
  font-size: 13px;
  color: #999;
}
.bottom {
  margin-top: 13px;
  line-height: 12px;
}
.button {
  padding: 0;
  float: right;
}
.image {
  width: 100%;
  display: block;
}
.clearfix:before,
.clearfix:after {
  display: table;
  content: "";
}
.clearfix:after {
  clear: both;
}
</style>