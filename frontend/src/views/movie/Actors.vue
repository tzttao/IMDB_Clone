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
          >Actor Image
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
          <el-main
            >Actors
            <el-table :data="tableData" style="width: 100%">
              <el-table-column
                prop="description"
                label="Description"
                width="700"
              >
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
            <h3 class="ipc-title__text">Relative movie you might like</h3>
            <el-table :data="relativeMovie" style="width: 100%">
              <el-table-column width="90"> </el-table-column>
              <el-table-column label="Movie Name" width="180">
                <template v-slot="props">
                  <span @click="handleClick(props.row)">{{
                    props.row.Movie_name
                  }}</span></template
                >
              </el-table-column>
              <el-table-column label="Movie">
                <template v-slot="prop">
                  <img
                    style="width: 100px; height: 150px"
                    :src="require('@/assets/' + prop.row.Image)"
                    fit="cover"
                  />
                </template>
              </el-table-column>
            </el-table>
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
      movieName: "",
      movieImage: [],
      image: "",
      grade: 0,
      tableData: [
        {
          genre: "",
          abstract:
            "A mentally troubled stand-up comedian embarks on a downward spiral that leads to the creation of an iconic villain",
          cast: "Joaquin Phoenix",
        },
      ],
      relativeMovie: [],
    };
  },
  mounted() {
    this.fetchData();
    this.searchMovieByTitle();
    this.SearchRelativeMovieByCastId();
  },
  methods: {
    handleClick(row) {
      this.$router.push({
        path: 'Movies',
        query: {movieId: row.Movie_id}
      });
    },
    fetchData() {
      this.CastId = this.$route.query.SrCastId;
    },
    //根据cast_id来搜索演员名字
    searchMovieByTitle() {
      console.log("this cast id:" + this.CastId);
      let params = {
        castId: this.CastId,
      };
      this.$axios
        .post("/api/user/cast/searchCastById", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.grade = res.data.cast.Grade;
            this.image = res.data.cast.CastImage;
            this.movieName = res.data.cast.CastName;
            this.bornyear = res.data.cast.Year;
            this.tableData[0].description = res.data.cast.Description;
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
    SearchRelativeMovieByCastId() {
      console.log("this cast id:" + this.CastId);
      let params = {
        castId: this.CastId,
      };
      this.$axios
        .post(
          "/api/user/cast/searchRelativeMovieByCastId",
          this.$qs.stringify(params)
        )
        .then((res) => {
          if (res.data.code === 200) {
            this.relativeMovie = res.data.movie;
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
  line-height: 150px;
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
