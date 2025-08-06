<template xmlns="http://www.w3.org/1999/html">
  <el-form :model="tableData" :rules="rules" ref="tableData" inline-message>
    <span class="demonstration" style="float:left; font-weight:bold; color:#fcd738; font-size: 100%">User Profile</span>
    <el-descriptions direction="vertical" :column="3" border :data="tableData"  >
      <template slot="extra">
        <el-button type="primary" @click="edit = !edit">Edit</el-button>
      </template>

      <el-descriptions-item label-style="background: #c6ddf5;" label-class-name="my-label"
                            content-class-name="my-content">

        <template slot="label">
          <i class="el-icon-user"></i>
          Username
        </template>
        <span v-show="!edit">{{ tableData.Username }}</span>
        <el-form-item v-show="edit" prop="Username" style="height: 10px; margin-top: -7.5px">
          <el-input v-model="tableData.Username" style=" width: 200px;"></el-input>
        </el-form-item>
      </el-descriptions-item>
      <el-descriptions-item>
        <template slot="label">
          <i class="el-icon-male"></i>
          Gender
        </template>
        <el-form-item prop="Gender" style="height: 10px; margin-top: -7.5px">
          <span v-show="!edit">{{ tableData.Gender }}</span>
          <el-radio-group v-show="edit" v-model="tableData.Gender">
            <el-radio label="male" sele><i class="el-icon-male">male</i></el-radio>
            <el-radio label="female"><i class="el-icon-female">female</i></el-radio>
          </el-radio-group>
          <!--          <el-input v-show="edit" v-model="tableData.Gender" style=" width: 200px;"></el-input>-->
        </el-form-item>
      </el-descriptions-item>
      <el-descriptions-item>
        <template slot="label">
          <i class="el-icon-lollipop"></i>
          Age
        </template>
        <el-form-item prop="Age" style="height: 10px; margin-top: -7.5px">
          <span v-show="!edit">{{ tableData.Age }}</span>
          <el-input v-show="edit" v-model.number="tableData.Age" style=" width: 200px;"/>
        </el-form-item>
      </el-descriptions-item>
      <el-descriptions-item>
        <template slot="label">
          <i class="el-icon-tickets"></i>
          User Type
        </template>
        <el-tag size="medium" v-if="tableData.User_type===0">User</el-tag>
        <el-tag size="small" v-else>Admin</el-tag>
      </el-descriptions-item>
      <el-descriptions-item :span="2">
        <template slot="label">
          <i class="el-icon-message"></i>
          Email
        </template>
        <el-form-item prop="Email" style="height: 10px; margin-top: -7.5px">
          <span v-show="!edit">{{ tableData.Email }}</span>
          <el-input v-show="edit" v-model="tableData.Email" style=" width: 200px;"></el-input>
        </el-form-item>
      </el-descriptions-item>
      <el-descriptions-item :span="3">
        <template slot="label">
          <i class="el-icon-key"></i>
          Password
        </template>
        <span v-show="!edit">
          <span v-show="!showPsw">show password</span>
        <span v-show="showPsw">{{ tableData.Password }}</span>
          <i
            :class="{'el-icon-edit': !showPsw, 'el-icon-check': showPsw}"
            @click="showPsw = !showPsw"
          ></i>
      </span>
        <el-form-item v-show="edit" label="Password：" prop="Password" style="height: 10px;">
          <el-input v-model="tableData.Password" :show-password="true" style=" width: 200px;"></el-input>
        </el-form-item>
        <br/>
        <el-form-item v-show="edit" label="Confirmed Password：" prop="checkPassword" style="height: 10px;">
          <el-input type="password" v-model="tableData.checkPassword" style="width:200px" show-password/>
        </el-form-item>
        <el-form-item v-show="edit">
          <br/>
          <el-button type="primary" @click="updateUserInfo()" icon="el-icon-mouse">Submit</el-button>
        </el-form-item>
      </el-descriptions-item>
    </el-descriptions>
    <br/>
    <el-button type="danger" @click="deleteAccount()" icon="el-icon-delete" style="float: left">Delete Account
    </el-button>
  </el-form>
</template>

<script>
export default {
  name: "UserProfile",
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
        if (this.tableData.checkPassword !== '') {
          this.$refs.tableData.validateField('checkPassword');
        }
        callback();
      }
    };
    var validatePass2 = (rule, value, callback) => {
      if (value === '') {
        callback(new Error('Please Input Password Again'));
      } else if (value !== this.tableData.Password) {
        callback(new Error('Passwords must match!'));
      } else {
        callback();
      }
    };
    let that = this;
    var validateUser = function (rule, value, callback) {
      if (!value) {
        return callback(new Error('username cannot be nothing'));
      } else if(value !== localStorage.getItem('username')){
        let params = {username: value}
        that.$axios.post("/api/user/checkUsername", that.$qs.stringify(params)).then((res) => {
          console.log(res.status)
          if (res.status === 201) {
            return callback(new Error('Username have been registered!!'));
          } else {
            return callback();
          }
        }).catch((error) => {
          console.log(error)       //请求失败返回的数据
        })
      } else {
        callback();
      }
    };
    return {
      tableData: {
        Username: '',
        Password: '',
        Gender: '',
        Email: '',
        Age: 0,
        checkPassword: ''
      },
      edit: false,
      showPsw: false,
      rules: {
        Username: [
          {required: true, validator: validateUser, trigger: "blur"},
          {min: 3, max: 12, message: "Please enter a username in a length of 3-12", trigger: "blur"}
        ],
        Password: [
          {required: true, validator: validatePass, trigger: 'blur'},
          {min: 6, max: 11, message: "Please enter a username in a length of 6-11", trigger: "blur"}
        ],
        checkPassword: [
          {required: true, validator: validatePass2, trigger: 'blur'}
        ],
        Age: [
          {required: true, message: "Please enter your age", trigger: "blur"},
          {validator: checkAge, trigger: 'blur'}
        ],
        Gender: [
          {required: true, message: "Please select your gender", trigger: "change"}
        ],
        Email: [{required: true, message: 'Please enter your e-mail address', trigger: 'blur'},
          {type: 'email', message: 'Please enter your correct e-mail address', trigger: ['blur']}]
      }
    };
  },
  mounted() {
    this.queryUserInfo()
  },
  methods: {
    queryUserInfo() {
      let params = {userId: localStorage.getItem('userId')}
      this.$axios.post("/api/user/queryUserInfoById", this.$qs.stringify(params))
        .then(res => {
          this.tableData = res.data.userInfo
        }).catch(error => {
        console.log(error);
        this.$message({
          title: '系统提示',
          message: error.message,
          center: true,
          type: 'warning'
        });
      });
    },
    updateUserInfo() {
      this.$refs.tableData.validate((valid) => {
        if (valid) {
          let params = {
            userId: localStorage.getItem('userId'),
            username: this.tableData.Username,
            password: this.tableData.Password,
            age: this.tableData.Age,
            gender: this.tableData.Gender,
            email: this.tableData.Email,
          }
          this.$axios.post("/api/user/UpdateProfile", this.$qs.stringify(params))
            .then(res => {
              localStorage.setItem('username', this.tableData.Username)
              if (res.status === 200) {
                this.$router.go(0)
                this.$message({
                  title: '更新提示',
                  message: 'Update successful',
                  showClose: true,
                  center: true,
                  type: 'success'
                });
              }
            }).catch(error => {
            console.log(error);
            this.$message({//这里采用element ui的一个错误显示效果模板
              title: '更新提示',
              message: error.message,
              center: true,
              type: 'warning'
            });
          })
        } else {
          console.log('error submit!!');
          return false;
        }
      });
    },
    deleteAccount() {
      this.$confirm('Are you sure delete your account', 'Alert', {
        confirmButtonText: 'Confirm',
        cancelButtonText: 'Cancel',
        type: 'warning'
      }).then(() => {
        let params = {userId: localStorage.getItem('userId')}
        this.$axios.post("/api/user/DeleteUser", this.$qs.stringify(params))
          .then(res => {
            console.log(res);
            if (res.status === 200) {
              this.$router.push('/UserLogIn')
              this.$message({
                type: 'info',
                message: 'delete successful'
              });
            }
          }).catch(error => {
          console.log(error);
          this.$message({//这里采用element ui的一个错误显示效果模板
            title: '更新提示',
            message: error.message,
            center: true,
            type: 'warning'
          });
        })
      }).catch(() => {
        this.$message({
          type: 'info',
          message: '已取消删除'
        });
      });
    }
  }
}
</script>

<style>
.my-content {
  background: #F6F3FCFF;
  margin-top: 60px;
}

</style>
