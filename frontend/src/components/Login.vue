<template>
  <div class="vue-tempalte">
    <form>
      <h3>Sign In</h3>

      <div class="form-group" name="">
        <label>Email address</label>
        <input type="text" class="form-control form-control-lg" name="username"/>
      </div>

      <div class="form-group">
        <label>Password</label>
        <input type="password" class="form-control form-control-lg"name="password"/>
      </div>

      <button type="submit" class="btn btn-dark btn-lg btn-block">Sign In</button>

      <p class="forgot-password text-right mt-2 mb-4">
        <router-link to="/forgot-password">Forgot password ?</router-link>
      </p>

      <div class="social-icons">
        <ul>
          <li><a href="#"><i class="fa fa-google"></i></a></li>
          <li><a href="#"><i class="fa fa-facebook"></i></a></li>
          <li><a href="#"><i class="fa fa-twitter"></i></a></li>
        </ul>
      </div>

    </form>
  </div>
</template>

<script>
export default {
  data() {
    return {
      // loginForm: {
      //   username: ,
      //   password: this.password,
      // },
    }
  },
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
    },
    // Login(){
    //   this.$axios.post("v1/signin",this.$qs.stringify(params))
    //     .then(res => {
    //       console.log(res);
    //       if (res.data != 0) {
    //         sessionStorage.setItem('user', this.loginForm.username);
    //         this.sId=res.data;
    //         sessionStorage.setItem('sId',this.sId);
    //         if(this.$route.query.rediret){//如果存在参数
    //           let redirect=this.$route.query.redirect
    //           this.$router.push(redirect)//则跳转至进入登录页前的路由
    //         }else{
    //           this.$router.push('/StudentContainer');//否则跳转至首页
    //         }
    //         this.$message({
    //           title: '登录提示',
    //           message: '登录成功',
    //           showClose: true,
    //           center: true,
    //           type: 'success'
    //         });
    //       } else {
    //         this.$message({
    //           title: '登录提示',
    //           message: '用户名或密码错误',
    //           center: true,
    //           type: 'error'
    //         });
    //         this.loading = false;
    //       }
    //     }).catch(error => {
    //     this.loading = false;
    //     console.log(error);
    //     this.$message({//这里采用element ui的一个错误显示效果模板
    //       title: '登录提示',
    //       message: error.message,
    //       center: true,
    //       type: 'warning'
    //     });
    //   })
    // }, 1000);
    // }
  }
}
</script>
