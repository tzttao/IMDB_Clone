<template>
  <div class="bgView">
    <div class="line"></div>
    <div class="inputText">
      <textarea
        v-model="inputText"
        placeholder="Please leave your comment about this movie."
      ></textarea>
    </div>
    <button class="publish" @click="publicComment">Submit</button>
  </div>
</template>

<script>
export default {
  data: function () {
    return {
      texts: ["Terrible", "Bad", "Regular", "Good", "Fantastic"],
      inputText: "",
    };
  },
  methods: {
    // 发布评论
    publicComment() {
      let params = {
        movieId: this.$route.query.movieId,
        userId:  localStorage.getItem('userId'),
        reviewContent: this.inputText
      };
      this.$axios
        .post("/api/user/review/addReview", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.$router.push({path:'/UserContainer/Movies',query:{movieId: params.movieId}});
            this.$message({
              title: "查询提示",
              message: "add comment successfully",
              showClose: true,
              center: true,
              type: "success",
            });
          } else {
            this.$message({
              title: "查询提示",
              message: "Sorry, the movie you search does not exist",
              center: true,
              type: "error",
            });
          }
        })
        .catch((error) => {
          console.log(error);
          this.$message({
            //这里采用element ui的一个错误显示效果模板
            title: "系统提示",
            message: error.message,
            center: true,
            type: "warning",
          });
        });
    },
  },
};
</script>

<style scoped>
.bgView {
  padding: 15px 15px;
}
.headerImg {
  width: 4.8125rem;
  height: 4.8125rem;
}
.score {
  color: #333333;
  font-size: 0.875rem;
  margin-left: 10px;
  display: inline-block;
}
.bgScored {
  position: relative;
  margin-top: -15px;
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 15px;
}

.bgImg img {
  margin-left: 5px;
  width: 1.09375rem;
  height: 1.03125rem;
}
.grade-box {
  position: absolute;
  top: 1.6875rem;
  left: 5.0275rem;
  display: inline-block;
}
.bgImg {
  display: inline-block;
  margin-left: 5px;
}
.line {
  margin-top: -5px;
  margin-left: -23px;
  margin-right: -23px;
  background: #c4c3c3;
  height: 1px;
  opacity: 0.3;
}
.scoreInfo {
  color: #999999;
  font-size: 0.75rem;
}
.inputText {
  width: 100%;
  margin-top: 10px;
  border-bottom: 1px solid #eef1f5;
}
.inputText textarea {
  width: 100%;
  height: 5.5rem;
  /*outline: none;*/
  /*overflow:auto;*/
  border: none;
  font-size: 0.75rem;
  /*background-attachment:fixed;*/
  /*background-repeat:no-repeat;*/
  /*border-style:solid;*/
  /*border-color:#FFFFFF*/
}
.selectPic {
  margin-top: 10px;
  width: 100%;
}
.unName {
  display: inline-block;
  color: #333333;
  margin-left: 5px;
  vertical-align: middle;
  font-size: 0.75rem;
}
.bgName {
  width: 100%;
  height: 50px;
  line-height: 50px;
  display: flex;
  align-items: center;
}
.bgName img {
  width: 17px;
  height: 17px;
}
.line2 {
  margin-top: -5px;
  margin-left: -23px;
  margin-right: -23px;
  background: #c4c3c3;
  height: 8px;
  opacity: 0.3;
}
.attitudeOfService {
  color: #333333;
  font-size: 0.875rem;
  display: inline-block;
}
.severce {
  display: flex;
  height: 60px;
  align-items: center;
}
.publish {
  width: 100%;
  height: 40px;
  background: #3fc9a5;
  color: white;
  font-size: 0.875rem;
  border-radius: 20px;
  line-height: 40px;
}
</style>
