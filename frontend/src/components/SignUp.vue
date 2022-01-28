<template>
  <el-container>
    <el-header style="display: flex; align-items: center">
      <el-button style="width: 100px; float: left; margin-left: 20px" v-on:click="getBacktoLogin">back</el-button>
      <span style="margin: 0 auto">Sign Up</span>
    </el-header>
    <el-main style="margin-top: 100px">
      <div class="demo-input-suffix" style="margin-bottom: 15px; padding-right: 40%; height: 50px">
        <div style="float: right">
          <span style="margin-right: 30px">email：</span>
          <el-input clearable v-model="email" style="width: 220px; float: right"></el-input>
        </div>
      </div>
      <div class="demo-input-suffix" style="margin-bottom: 30px; padding-right: 40%; height: 50px">
        <div style="float: right">
          <span style="margin-right: 30px">password：</span>
          <el-input v-model="password" show-password clearable style="width: 220px; float: right"></el-input>
        </div>
      </div>
      <div style="display: inline-block">
        <el-button style="width: 100px; margin-right: 30px" v-on:click="getBacktoLogin">Cancel</el-button>
        <el-button style="width: 100px;" v-on:click="signUp" type="primary">Create</el-button>
      </div>
    </el-main>
  </el-container>
</template>

<script>
import Main from '../App'

export default {
  name: 'SignUp',
  data () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    getBacktoLogin () {
      this.$router.push('/')
    },
    signUp () {
      if (this.email.length === 0 || this.password.length === 0) {
        this.$alert('Please enter your email or password!', 'Alert', {
          confirmButtonText: 'OK',
          callback: action => {
          }
        })
      } else {
        this.$ajax.post('/api/signUp/' + this.email + '/' + this.password).then(res => {
          if (res.data === 'success') {
            this.$ajax.post('/api/getUserInfo/' + this.email).then(res => {
              Main.email = res.data.email
              Main.usertype = res.data.usertype
              this.$message({
                message: 'Welcome to Weather Recording System!',
                type: 'success'
              })
              setTimeout(() => {
                this.$router.push('/welcome')
              }, 500)
            })
          } else if (res.data === 'error') {
            this.$message.error('Sorry, please try again!')
          } else if (res.data === 'Account existed') {
            this.$message.error('This email has already had an account!')
          } else if (res.data === 'Incorrect format') {
            this.$message.error('Please enter the correct format of email!')
          }
        })
      }
    }
  }
}
</script>

<style scoped>

</style>
