<template>
  <div class="register-container">
    <el-form
      :model="ruleForm"
      :rules="rules"
      ref="ruleForm"
      label-width="200px"
      class="demo-ruleForm register-page"
      status-icon
    >
      <div align="center" style="padding-right:30px">
          <span>
            <h2>User Sign Up</h2>
          </span>
      </div>
      <br/>
      <el-form-item label="Username：" prop="username">
        <el-input v-model="ruleForm.username" style="width:250px"></el-input>
      </el-form-item>
      <el-form-item label="Password：" prop="password">
        <el-input v-model="ruleForm.password" style="width:250px" show-password></el-input>
      </el-form-item>
      <el-form-item label="Confirmed Password：" prop="checkPassword">
        <el-input
          type="password"
          v-model="ruleForm.checkPassword"
          style="width:250px"
          show-password
        ></el-input>
      </el-form-item>
      <el-form-item label="Age：" prop="age">
        <el-input v-model.number="ruleForm.age" style="width:100px"></el-input>
      </el-form-item>
      <el-form-item label="Gender：" prop="gender">
        <el-radio-group v-model="ruleForm.gender">
          <el-radio label="male"><i class="el-icon-male">male</i></el-radio>
          <el-radio label="female"><i class="el-icon-female">female</i></el-radio>
        </el-radio-group>
      </el-form-item>
      <el-form-item>
        <el-button @click="resetForm" icon="el-icon-refresh">Reset</el-button>
        <el-button type="primary" @click="submitForm('ruleForm')" icon="el-icon-mouse">Sign Up</el-button>
      </el-form-item>
    </el-form>

  </div>

</template>
<script>
export default {
  data() {
    var checkAge = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('The age is required!'));
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('Please Input Integer!'));
        } else {
          if (value < 12) {
            callback(new Error('Must be over 12 years old!'));
          } else if (value < 1 || value > 120) {
            callback(new Error('Invalid age input!'));
          } else {
            callback();
          }
        }
      }, 500);
    };
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please Input Password'));
      } else {
        if (this.ruleForm.checkPassword !== '') {
          this.$refs.ruleForm.validateField('checkPassword');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please Input Password Again'));
      } else if (value !== this.ruleForm.password) {
        callback(new Error('Passwords must match!'));
      } else {
        callback();
      }
    };
    let that = this;
    var validateUser = function (rule, value, callback) {
      if (!value) {
        return callback(new Error('Username cannot be empty'));
      } else {
        let params = {username: value}
        that.$axios.post("/api/user/checkUsername", that.$qs.stringify(params)).then((res) => {
          console.log(res.status)
          if (res.status == 201) {
            return callback(new Error('Username have been registered!!'));
          } else {
            return callback();
          }
        }).catch((error) => {
          console.log(error)       //请求失败返回的数据
        })
      }
    };
    return {
      ruleForm: {
        username: '',
        name: "",
        password: '',
        checkPassword: '',
        age: '',
        gender: '',
      },
      rules: {
        username: [
          {required: true, validator: validateUser, trigger: "blur"},
          {min: 3, max: 12, message: "Length has to be 3 to 12 characters", trigger: "blur"}
        ],
        password: [
          {required: true, validator: validatePass, trigger: 'blur'},
          {min: 6, max: 11, message: "Length has to be 6 to 11 characters", trigger: "blur"}
        ],
        checkPassword: [
          {required: true, validator: validatePass2, trigger: 'blur'}
        ],
        age: [
          {required: true, message: "Please enter your age", trigger: "blur"},
          {validator: checkAge, trigger: 'blur'}
        ],
        gender: [
          {required: true, message: "Please enter your gender", trigger: "change"}
        ]
      }
    };
  },
  methods: {
    submitForm() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          let params = {
            username: this.ruleForm.username,
            password: this.ruleForm.password,
            age: this.ruleForm.age,
            gender: this.ruleForm.gender,
          }
          this.$axios.post("/api/user/signUp", this.$qs.stringify(params))
            .then(res => {
              console.log(res);
              if (res.status == 200) {
                this.$router.push('/UserLogIn');//否则跳转至首页
                this.$message({
                  title: 'Registration',
                  message: 'registered successfully',
                  showClose: true,
                  center: true,
                  type: 'success'
                });
              }
            }).catch(error => {
            this.loading = false;
            console.log(error);
            this.$message({//这里采用element ui的一个错误显示效果模板
              title: 'Registration',
              message: error.message,
              center: true,
              type: 'warning'
            });
          })
        } else {
          console.log("error submit!!");
          return false;
        }
      });
    },
    resetForm() {
      this.$refs.ruleForm.resetFields();
    }
  }
};
</script>


<style scoped>
.register-container {
  width: 100%;
  height: 100%;
  padding-top: 10px;
}

.register-page {
  -webkit-border-radius: 5px;
  border-radius: 5px;
  margin: 0px auto;
  width: 600px;
  padding: 20px;
  padding-left: 50px;
  background: #fff;
  border: 1px solid #eaeaea;
  box-shadow: 0 0 25px #cac6c6;
}
</style>
