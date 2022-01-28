<template>
  <el-container>
    <el-header>Welcome to Weather Recording System</el-header>
    <el-main style="margin-top: 100px">
      <div class="demo-input-suffix" style="margin-bottom: 15px; padding-right: 42%; height: 50px">
        <div style="float: right">
          <span style="margin-right: 30px">email：</span>
          <el-input clearable v-model="email" style="width: 220px; float: right"></el-input>
        </div>
      </div>
      <div class="demo-input-suffix" style="margin-bottom: 30px; padding-right: 42%; height: 50px">
        <div style="float: right">
          <span style="margin-right: 30px">password：</span>
          <el-input v-model="password" show-password clearable style="width: 220px; float: right"></el-input>
        </div>
      </div>
      <div style="display: inline-block">
        <el-button style="width: 100px; margin-right: 30px" v-on:click="signUp">Sign Up</el-button>
        <el-button style="width: 100px;" v-on:click="login">Login</el-button>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import Main from '../App'

export default {
  name: "Login",
  data(){
    return{
      email: '', password: ''
    }
  },
  created() {
  },
  methods: {
    login () {
      if (this.email.length === 0 || this.password.length === 0) {
        this.$alert('Please enter your email or password!', 'Alert', {
          confirmButtonText: 'OK',
          callback: action => {
            // this.$message({
            //   type: 'info',
            //   message: `action: ${action}`
            // })
          }
        })
      } else {
        this.$ajax.post('/api/login/' + this.email + '/' + this.password).then(res1 => {
          if (res1.data === 'success') {
            this.$ajax.post('/api/getUserInfo/' + this.email).then(res2 => {
              Main.email = res2.data.email
              Main.usertype = res2.data.usertype
              console.log(Main.email)
              if (Main.usertype === 'M') {
                this.$router.push('/admin')
              } else if (Main.usertype === 'R') {
                this.$router.push('/welcome')
              }
            })
          } else {
            this.$message.error('Sorry, your email or password is wrong!')
          }
        })
      }
    },
    signUp () {
      this.$router.push('/signUp')
    }
  }
}
</script>


<style scoped>
demo-input-suffix {
  display: inline-block;
}

el-input {
  width: 200px;
}
</style>
