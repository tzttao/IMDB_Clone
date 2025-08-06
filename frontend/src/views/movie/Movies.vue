<template>
  <div class="block">
    <el-container>
      <el-header>
        <template>
          {{ movieName }}
        </template>
      </el-header>
      <el-container>
        <el-aside
          >Movie Poster
          <template v-slot="props">
            <el-image
              style="width: 259px; height: 375px"
              :src="require('@/assets/' + $data.image)"
              fit="cover"
            >
            </el-image
          ></template>
        </el-aside>
        <el-container>
          <el-main>
            <h2 style="height:50px;text-align;center">Movie Information</h2>
            <el-table :data="tableData" style="width: 100%">
              <el-table-column label="Genre" width="200">
                <span v-for="(item, i) in genreTable">
                  <el-tag>{{ item.Genre }}</el-tag> &nbsp;
                </span>
              </el-table-column>
              <el-table-column label="Cast">
                <span v-for="(item, i) in castList">
                  {{ item.CastName }},
                </span>
              </el-table-column>
              <el-table-column prop="abstract" label="Abstract">
              </el-table-column>
              <el-table-column label="Rating" width="200">
                <el-rate
                  v-model="score"
                  disabled
                  show-score
                  text-color="#ff9900"
                  score-template="{value}"
                >
                </el-rate>
              </el-table-column>
              <el-table-column label="My Rating" width="200">
                <span v-if="$data.username == null || ''">
                  <el-link
                    type="primary"
                    href="http://localhost:8080/#/UserLogIn"
                    >login or sign up</el-link
                  >
                </span>
                <span v-else>
                  <template v-if="myScore === 0.0 || 0 || null">
                    <el-rate
                      v-model="myScore"
                      show-score
                      allow-half
                      score-template="{value}"
                      @change="addScore()"
                    ></el-rate>
                    <span>you do not rate this movie</span>
                  </template>
                  <template v-else>
                    <el-rate
                      v-model="myScore"
                      show-score
                      text-color="#ff9900"
                      score-template="{value}"
                      allow-half
                      @change="undateScore()"
                    >
                    </el-rate></template
                ></span>
              </el-table-column>
            </el-table>
          </el-main>
          <el-footer> Storyline </el-footer>

          <hgroup
            class="
              ipc-title
              ipc-title--section-title
              ipc-title--base
              ipc-title--on-textPrimary
            "
          >
            <h3 class="ipc-title__text">
              <el-badge class="item">
                <span v-if="$data.username == null || ''">
                  <el-link
                    type="warning"
                    href="http://localhost:8080/#/UserLogIn"
                    >if comment, please login or sign up</el-link
                  >
                </span>
                <span v-else>
                  <el-button size="small" @click="toMovie"
                    >Comments</el-button
                  ></span
                >
              </el-badge>
            </h3>
            <el-table :data="commentTable" style="width: 100%">
              <el-table-column width="90"> </el-table-column>
              <el-table-column label="username" prop="Username" width="180">
              </el-table-column>
              <el-table-column label="content" prop="Review_content">
              </el-table-column>
            </el-table>
            <el-pagination
              @size-change="handleSizeChange"
              @current-change="handleCurrentChange"
              :current-page="currentPage"
              :page-size="pageSize"
              background
              :total="total"
            >
            </el-pagination>
          </hgroup>
        </el-container>
      </el-container>
    </el-container>
  </div>
</template>

<script>
export default {
  data() {
    return {
      username: localStorage.getItem("username"),
      genreTable: [],
      commentTable: [],
      score: 3.7,
      myScore: 0.0,
      texts: ["Terrible", "Bad", "Regular", "Good", "Fantastic"],
      movieName: "",
      year: 0,
      image: "",
      tableData: [
        {
          abstract: "",
        },
      ],
      castList: [],
      total: 0, //数据总数
      pageSize: 5, //每页的数据条数
      currentPage: 1, //默认开始页面
    };
  },
  created() {
      this.fetchData();
      this.showCast();
      this.queryRate();
      this.readReviewByMovieId();
      this.searchMovieWithGenre();
      this.searchMovieByTitle();
  },
  methods: {
    handleSizeChange(val) {
      console.log(`每页 ${val} 条`);
      this.pageSize = val;
      this.readReviewByMovieId();
    },
    handleCurrentChange(val) {
      console.log(`当前页: ${val}`);
      this.currentPage = val;
      this.readReviewByMovieId();
    },
    searchMovieWithGenre() {
      let params = {
        movieId: this.movieId,
      };
      this.$axios
        .post(
          "/api/user/movie/searchMovieWithGenre",
          this.$qs.stringify(params)
        )
        .then((res) => {
          if (res.data.code == 200) {
            this.genreTable = res.data.genre;
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
    readReviewByMovieId() {
      let params = {
        movieId: this.movieId,
        pageNo: this.currentPage,
        pageSize: this.pageSize,
      };
      this.$axios
        .post(
          "/api/user/review/readReviewByMovieId",
          this.$qs.stringify(params)
        )
        .then((res) => {
          if (res.data.code == 200) {
            this.commentTable = res.data.review;
            this.total = res.data.count;
            console.log(this.total);
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
    undateScore() {
      let params = {
        movieId: this.movieId,
        userId: localStorage.getItem("userId"),
        score: this.myScore,
      };
      this.$axios
        .post("/api/user/rating/updateRating", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.$router.go(0);
            this.$message({
              title: "查询提示",
              message: "rating success",
              showClose: true,
              center: true,
              type: "success",
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
    addScore() {
      let params = {
        movieId: this.movieId,
        userId: localStorage.getItem("userId"),
        score: this.myScore,
      };
      this.$axios
        .post("/api/user/rating/addRating", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.$router.go(0);
            this.$message({
              title: "查询提示",
              message: "rating success",
              showClose: true,
              center: true,
              type: "success",
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
    toMovie() {
      this.$router.push({
        path: "Comments",
        query: { movieId: this.fetchData() },
      });
    },
    queryRate() {
      let params = {
        movieId: this.movieId,
        userId: localStorage.getItem("userId"),
      };
      this.$axios
        .post("/api/user/rating/readRating", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.myScore = res.data.rating.Score;
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
    fetchData() {
      this.movieId = this.$route.query.movieId;
      return this.movieId;
    },
    showCast() {
      let params = {
        movieId: this.movieId,
      };
      this.$axios
        .post("/api/user/cast/searchCastByMovieId", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.castList = res.data.cast;
          } else {
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

    //根据movie_id来找movie信息
    searchMovieByTitle() {
      let params = {
        movieId: this.movieId,
      };
      this.$axios
        .post(
          "/api/user/movie/searchMovieByMovieId",
          this.$qs.stringify(params)
        )
        .then((res) => {
          if (res.data.code === 200) {
            this.score = res.data.movie.Grade;
            this.image = res.data.movie.Image;
            this.movieName = res.data.movie.Movie_name;
            this.year = res.data.movie.Year;
            //this.cast = res.data.movie.
            this.tableData[0].abstract = res.data.movie.Description;
            this.$message({
              title: "查询提示",
              message: "This is what we found by movie titles",
              showClose: true,
              center: true,
              type: "success",
            });
          } else {
            this.$message({
              title: "查询提示",
              message: "Sorry, we cannot find any result by movie titles",
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

export default {
data() {
return {
value1: null,
value2: null,
colors: ['#99A9BF', '#F7BA2A', '#FF9900']  // 等同于 { 2: '#99A9BF', 4: { value: '#F7BA2A', excluded: true }, 5: '#FF9900' }
}
}
}
<style scoped>
.el-header,
.el-footer {
  background-color: #b3c0d1;
  color: #333;
  text-align: center;
  line-height: 60px;
}

.el-aside {
  background-color: #d3dce6;
  color: #333;
  text-align: center;
  line-height: 150px;
  width: fit-content;
}

.el-main {
  background-color: #e9eef3;
  color: #333;
  text-align: center;
  line-height: 50px;
  width: 100%;
  margin-left: auto;
  margin-right: auto;
}

body > .el-container {
  margin-bottom: 40px;
}

.el-container:nth-child(5) .el-aside,
.el-container:nth-child(6) .el-aside {
  line-height: 260px;
}

.el-container:nth-child(7) .el-aside {
  line-height: 320px;
}
</style>
