<template>
  <div class="login-container">
    <el-form
      :model="loginForm"
      :rules="loginRules"
      ref="loginForm"
      label-position="left"
      label-width="0px"
      class="demo-ruleForm login-page"
    >
      <h3 class="title">Log In</h3>
      <el-form-item prop="username">
        <el-input type="text" v-model="loginForm.username" auto-complete="off" placeholder="用户名">
          <template slot="prepend">
            <i class="el-icon-s-custom"></i>
          </template>
        </el-input>
      </el-form-item>
      <el-form-item prop="password">
        <el-input type="password" v-model="loginForm.password" auto-complete="off" placeholder="密码" show-password>
          <template slot="prepend">
            <i class="el-icon-lock"></i>
          </template>
        </el-input>
      </el-form-item>
      <el-link type="primary" class="register" href="http://localhost:8080/#/UserSignUp">Sign Up</el-link>
      <el-form-item style="width:100%;">
        <el-button type="primary" style="width:100%;" @click="checkLogin" :loading="load">Sign In</el-button>
      </el-form-item>
    </el-form>

  </div>

</template>

<script>
export default {
  name: 'loginForm',
  data() {
    return {
      sId: '',
      load: false,
      token: '',
      loginForm: {
        username: 'qwe',
        password: '123',
      },
      loginRules: {
        username: [{required: true, message: '请输入用户名', trigger: 'blur'}],
        password: [{required: true, message: '请输入密码', trigger: 'blur'}]
      },
      checked: false,
    }
  },
  methods: {
    checkLogin() {
      this.$refs.loginForm.validate(valid => {
          if (valid) {//校验成功
            this.load = true
            setTimeout(() => {//为了看到登录转圈圈的加载效果，这里来一个延迟2秒发送请求
                //封装的公共请求API
                let params = {
                  username: this.loginForm.username,
                  password: this.loginForm.password
                }
                this.$axios.post("/api/user/LogIn", this.$qs.stringify(params))
                  .then(res => {
                    if (res.status == 200) {
                      console.log("token:"+res.data.token)
                      localStorage.setItem('token',res.data.token);
                      this.$router.push('/hello');//否则跳转至首页
                      this.$message({
                        title: '登录提示',
                        message: '登录成功',
                        showClose: true,
                        center: true,
                        type: 'success'
                      });
                    } else {
                      this.$message({
                        title: '登录提示',
                        message: '用户名或密码错误',
                        center: true,
                        type: 'error'
                      });
                      this.load = false;
                    }
                  }).catch(error => {
                  this.load = false;
                  console.log(error);
                  this.$message({//这里采用element ui的一个错误显示效果模板
                    title: '登录提示',
                    message: error.message,
                    center: true,
                    type: 'warning'
                  });
                })
              }
              ,
              500
            )
            ;
          } else {
            return false;
          }
        }
      )
    }
  }
}
;
</script>

<style scoped>
.login-container {
  width: 100%;
  height: 100%;
}

.login-page {
  -webkit-border-radius: 5px;
  border-radius: 5px;
  margin: 180px auto;
  width: 350px;
  padding: 35px 35px 15px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}

.register {
  margin: 0px 0px 15px;
  text-align: left;
}


</style>
