<template>
  <div>
    <div>
      <span
        class="demonstration"
        style="float: left; font-weight: bold; color: #fcd738; font-size: 100%"
        >Titles</span
      >
      <el-table :data="titleTableData" style="width: 100%">
        <el-table-column type="expand">
          <template v-slot="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <el-form-item label="Movie Name">
                <span @click="handleClick(props.row)">{{
                  props.row.Movie_name
                }}</span>
              </el-form-item>
              <el-form-item label="Grade">
                <span>{{ props.row.Grade }}</span>
              </el-form-item>
              <el-form-item label="Year">
                <span>{{ props.row.Year }}</span>
              </el-form-item>
              <el-form-item label="Desciption">
                <span>{{ props.row.Description }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column label="Movie">
          <template v-slot="props">
            <div class="demo-image">
              <el-image
                style="width: 100px; height: 150px"
                :src="require('@/assets/' + props.row.Image)"
                fit="cover"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Movie Name">
          <template v-slot="props">
            <span @click="handleClick(props.row)">{{
              props.row.Movie_name
            }}</span></template
          >
        </el-table-column>
        <el-table-column label="Year" prop="Year"> </el-table-column>
        <el-table-column label="Grade" prop="Grade">
          <template v-slot="props">
          <el-rate
            v-model="props.row.Grade "
            disabled
            show-score
            text-color="#ff9900"
            score-template="{value}"
          >
          </el-rate></template>
        </el-table-column>
      </el-table>
    </div>
    <br />
    <div>
      <span
        class="demonstration"
        style="float: left; font-weight: bold; color: #fcd738; font-size: 100%"
        >Names</span
      >
      <el-table :data="nameTableData" style="width: 100%">
        <el-table-column type="expand">
          <template v-slot="props">
            <el-form label-position="left" inline class="demo-table-expand">
              <!-- <el-form-item label="Movie Name">
                <span>{{ props.row.MovieName }}</span>
              </el-form-item> -->
              <el-form-item label="Actor Name">
                <span @click="Click(props.row)">{{ props.row.Cast_name }}</span>
              </el-form-item>
              <!-- <el-form-item label="Grade">
                <span>{{ props.row.Grade }}</span>
              </el-form-item> -->
              <!-- <el-form-item label="Year">
                <span>{{ props.row.Year }}</span>
              </el-form-item> -->
              <el-form-item label="Description">
                <span>{{ props.row.Cast_description }}</span>
              </el-form-item>
            </el-form>
          </template>
        </el-table-column>
        <el-table-column label="Actor" width="200px">
          <template v-slot="props">
            <div class="demo-image">
              <el-image
                style="width: 100px; height: 150px"
                :src="require('@/assets/' + props.row.Cast_image)"
                fit="cover"
              />
            </div>
          </template>
        </el-table-column>
        <el-table-column label="Actor Name" prop="CastName" width="200px">
          <template v-slot="props">
            <span @click="Click(props.row)">{{
              props.row.Cast_name
            }}</span></template
          >
        </el-table-column>
        <el-table-column label="Description" prop="Cast_description">
        </el-table-column>
      </el-table>
    </div>
  </div>
</template>

<script>
export default {
  name: "SearchResult",
  data() {
    return {
      titleTableData: [],
      nameTableData: [],
      searchInput: "",
    };
  },
  watch: {
    $route(to, from) {
      this.$router.go(0);
      this.fetchData();
      this.searchMovieByTitle();
      this.searchMovieByName();
    },
  },
  mounted() {
    this.fetchData();
    this.searchMovieByTitle();
    this.searchCastByName();
  },
  methods: {
    handleClick(row) {
      console.log("!!:" + row.Movie_id);
      this.$router.push({
        path: "Movies",
        query: { movieId: row.Movie_id },
      });
    },
    Click(row) {
      this.$router.push({
        path: "Actors",
        query: { SrCastId: row.Cast_id },
      });
    },
    fetchData() {
      this.searchInput = this.$route.query.searchInput;
      console.log("searchInput:" + this.searchInput);
    },
    searchMovieByTitle() {
      let params = {
        searchName: this.searchInput,
      };
      this.$axios
        .post("/api/user/movie/searchMovieByName", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.titleTableData = res.data.movie;
          } else {
          }
        })
        .catch((error) => {});
    },
    searchCastByName() {
      let params = {
        castName: this.searchInput,
      };
      this.$axios
        .post("/api/user/cast/searchCastByName", this.$qs.stringify(params))
        .then((res) => {
          if (res.data.code === 200) {
            this.nameTableData = res.data.movie;
          } else {
            this.nameTableData.movie = null;
          }
        })
        .catch((error) => {
          console.log(error);
        });
    },
  },
};
</script>

<style scoped>
.demo-table-expand {
  font-size: 0;
}

.demo-table-expand label {
  width: 90px;
  color: #99a9bf;
}

.demo-table-expand .el-form-item {
  margin-right: 0;
  margin-bottom: 0;
  width: 50%;
}
</style>
