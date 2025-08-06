<template>
  <el-container>
        <!-- <el-image
          style="width: 315px; height: 110px"
          :src="require('@/assets/IMDB-CLone-FinalLogo.png')"
          :fit="fill">
        </el-image> -->
    <el-header>
      <el-menu class="el-menu-demo" mode="horizontal" active-text-color="#409EFF" :router=true>
        <el-menu-item index="/UserContainer/Homepage">
          <i class="el-icon-s-home"></i>
        </el-menu-item>
        <el-menu-item>
          <span>
          <el-button @click="drawer = true" size="medium" icon="el-icon-menu" round>
            Menu
          </el-button>
          <el-drawer
            title="IMDb Clone"
            :visible.sync="drawer"
            :direction="direction"
            :before-close="handleClose">
            <span>Welcome to IMDb!</span>
          </el-drawer>
          </span>
        </el-menu-item>

        <el-menu-item>
          <span>
              <el-input placeholder="Search" prefix-icon="el-icon-search" v-model="searchInput"
                        @keydown.enter.native="onSubmit" style=" width: 975px;" id="search" />
          </span>
        </el-menu-item>

        <el-submenu index="" style="float: right" v-if="username!==null||''">
          <template slot="title">{{ username }}</template>
          <el-menu-item index="Userprofile">My profile</el-menu-item>
          <el-menu-item index="2-2">Watch List</el-menu-item>
          <el-menu-item index="/UserLogIn" @click="removeToken">Sign out</el-menu-item>
        </el-submenu>
        <el-menu-item index="/UserLogIn" style="float: right" v-else>
             <span>
            <el-button type="primary" round>LogIn</el-button></span>
        </el-menu-item>
      </el-menu>
    </el-header>


    <el-main>
      <template>
        <router-view/>
      </template>
    </el-main>


    <el-footer>
      <el-divider></el-divider>
      <el-row :gutter="20" justify="center" align="middle" type="flex">
        <el-col :span="5">
          <el-button type="danger">Instagram</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="primary">Twitter</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="warning">Weibo</el-button>
        </el-col>
        <el-col :span="5">
          <el-button type="success">WeChat</el-button>
        </el-col>
      </el-row>
      <el-divider></el-divider>
    </el-footer>
  </el-container>

</template>
<script>
export default {
  data() {
    return {
      myStyle:{
        backgroundColor:"#16a085"
      },
      drawer: false,
      direction: 'ttb',
      searchInput: '',
      username: localStorage.getItem('username'),
    };
  },
  methods: {
    removeToken() {
      localStorage.clear()
      console.log(localStorage.getItem('username'))
      console.log(localStorage.getItem('token'))
    },
    handleClose(done) {
      done();
    },
    onSubmit() {
      console.log("1111")
      this.$router.push({
        path: 'SearchResult',
        query: {searchInput: this.searchInput}
      });
    },
  }
}
</script>

<style>
.el-container {
  color: #333;
  text-align: center;
  width: 100%;
}
.el-header {
  background-color: #475669;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-menu {
  background-color: #475669;
  color: #333;
  text-align: center;
  width: 100%;
}

.el-footer {

  color: #333;
  text-align: center;
  width: 100%;
}

.el-row {

  color: #333;
  text-align: center;
  width: 100%;
}

.el-aside {
  color: #333;
  text-align: center;
  line-height: 200px;
}

.ui {
  list-style: none;
  width: 700px;
  height: 50px;
  margin: 0;
  padding: 0;
}

.el-main {
  background-color: #7690b9;
  color: #333;
  text-align: center;
  margin-left: 10%;
  margin-right: 10%;
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
  clear: both
}
</style>
