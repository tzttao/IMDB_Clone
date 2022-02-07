<template>
    <div class="vue-tempalte">
        <form>
            <h3>Sign Up</h3>

            <div class="form-group">
                <label>Full Name</label>
                <input type="text" class="form-control form-control-lg"/>
            </div>

            <div class="form-group">
                <label>Email address</label>
                <input type="email" class="form-control form-control-lg" />
            </div>

            <div class="form-group">
                <label>Password</label>
                <input type="password" class="form-control form-control-lg" />
            </div>

            <button type="submit" class="btn btn-dark btn-lg btn-block">Sign Up</button>

            <p class="forgot-password text-right">
                Already registered
                <router-link :to="{name: 'login'}">sign in?</router-link>
            </p>
        </form>
    </div>
</template>

<script>
export default {
  data() {
    var checkUsername = (rule, value, callback) => {
      if (!value) {
        return callback(new Error('年龄不能为空'));
      }
      setTimeout(() => {
        if (!Number.isInteger(value)) {
          callback(new Error('请输入数字值'));
        } else {
          if (value < 18) {
            callback(new Error('必须年满18岁'));
          } else if (value < 1 || value > 120) {
            callback(new Error('输入年龄不合法'));
          } else {
            callback();
          }
        }
      }, 500);
    };
    var validatePass = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请输入密码'));
      } else {
        if (this.ruleForm.checkPassword !== '') {
          this.$refs.ruleForm.validateField('checkPassword');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('请再次输入密码'));
      } else if (value !== this.ruleForm.password) {
        callback(new Error('两次输入密码不一致!'));
      } else {
        callback();
      }
    };
    let that = this;
    var user = function (rule, value, callback) {
      if (!value) {
        return callback(new Error('用户名不能为空'));
      } else {
        let params = { username: value }
        that.$axios.post("/api/checkTeacherUsername", that.$qs.stringify(params)).then((res) => {
          if (res.data.code == '500') {
            callback(new Error('用户名已存在'));
          } else {
            callback();
          }    //这里使用了ES6的语法
        }).catch((error) => {
          console.log(error)       //请求失败返回的数据
        })
      }

    };
    return {
      dept:[],
      ruleForm: {
        username: '',
        name: "",
        password: '',
        checkPassword: '',
        dId: '',
        age: '',
        sex: '',
      },
      rules: {
        username: [
          { required: true, validator: user, trigger: "blur" },
          { min: 3, max: 12, message: "长度在 3 到 12 个字符", trigger: "blur" }
        ],
        password: [
          { required: true, validator: validatePass, trigger: 'blur' },
          { min: 6, max: 11, message: "长度在 6 到 11 个字符", trigger: "blur" }
        ],
        checkPassword: [
          { required: true, validator: validatePass2, trigger: 'blur' }
        ],
        name: [
          { required: true, message: "请输入真实姓名", trigger: "blur" },
          { min: 2, max: 8, message: "长度在 2 到 8 个字符", trigger: "blur" }
        ],
        age: [
          { validator: checkAge, trigger: 'blur' }
        ],
        dId: [
          { required: true, message: "请选择学院", trigger: "change" }
        ],
        sex: [
          { required: true, message: "请选择性别", trigger: "change" }
        ]
      }
    };
  },
  mounted:function () {
    this.selectDepartment();
  },
  methods: {
    submitForm() {
      this.$refs.ruleForm.validate(valid => {
        if (valid) {
          let params = {
            username: this.ruleForm.username,
            password: this.ruleForm.password,
            dId: this.ruleForm.dId,
            age: this.ruleForm.age,
            sex: this.ruleForm.sex,
            name: this.ruleForm.name
          }
          this.$axios.post("/api/teacherRegister", this.$qs.stringify(params))
            .then(res => {
              console.log(res);
              if (res.data.code == '200' && res.data.message == '注册成功') {
                this.$router.push('/TeacherLogin');//否则跳转至首页
                this.$message({
                  title: '注册提示',
                  message: '注册成功',
                  showClose: true,
                  center: true,
                  type: 'success'
                });
              } else if (res.data.code == '500') {
                this.$message({
                  title: '注册提示',
                  message: '用户名已存在',
                  center: true,
                  type: 'error'
                });
                this.loading = false;
              }
            }).catch(error => {
            this.loading = false;
            console.log(error);
            this.$message({//这里采用element ui的一个错误显示效果模板
              title: '注册提示',
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
    },
    selectDepartment() {
      let that=this
      let params = {}
      this.$axios.post("/api/selectDepartment", this.$qs.stringify(params))
        .then(function (res) {
          that.dept = res.data;
          console.log(that.dept);
        }).catch(function (error) {
        console.log(error);
      });
    }
  }
};
</script>
