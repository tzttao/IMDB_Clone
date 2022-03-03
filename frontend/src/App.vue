<template>
    <!--App -->
    <div class="vue-tempalte">
      <!-- Navigation -->
      <nav class="navbar shadow bg-white rounded justify-content-between flex-nowrap flex-row fixed-top">
        <div class="container">
          <!--        <a class="navbar-brand float-left" href="https://www.positronx.io" target="_blank">-->
          <!--           positronX.io-->
          <!--        </a>-->
          <ul class="nav navbar-nav flex-row float-right">
            <li class="nav-item">
              <router-link class="nav-link pr-3" to="/login">Sign in</router-link>
            </li>
            <li class="nav-item">
              <router-link class="btn btn-outline-primary" to="/">Sign up</router-link>
            </li>
          </ul>
          <button @click="testAxios">TestAxios</button>
        </div>
      </nav>

      <!-- Main -->
      <div class="App">
        <div class="vertical-center">
          <div class="inner-block">
            <router-view />
          </div>
        </div>
      </div>
    </div>
</template>


<script>
    export default {
        methods: {
            testAxios() {
                // 由于 main.js 里全局定义的 axios,此处直接使用 $axios 即可。
                // 由于 main.js 里定义了每个请求前缀，此处的 / 即为 /api/，
                // 经过 vue.config.js 配置文件的代理设置，会自动转为 https://www.baidu.com/，从而解决跨域问题
                this.$axios.get("v1/signin").then(response => {
                    if (response.data) {
                        console.log(response.data)
                    }
                }).catch(err => {
                    alert('请求失败')
                })
            }
        }
    }
</script>

<style>

</style>


