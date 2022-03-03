package Controllers

import (
	"backend/Database"
	myjwt "backend/Middleware"
	"backend/Models"
	"backend/Services"
	jwtgo "github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
	"time"
)

// LogIn Sign in
func LogIn(c *gin.Context) {
	var user Models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	isTure := Database.DB.Where("username = ? AND password = ?", user.Username, user.Password).Find(&user).RecordNotFound()
	if isTure == true {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
		})
	} else {
		generateToken(c, user)
	}

}

func SignUp(c *gin.Context) {
	var user Models.User
	user.Username = c.PostForm("username")
	user.Password = c.PostForm("password")
	user.Gender = c.PostForm("gender")
	user.Age, _ = strconv.Atoi(c.PostForm("age"))
	user = Models.User{
		Username:  user.Username,
		Password:  user.Password,
		User_type: 0,
		Gender:    user.Gender,
		Age:       user.Age,
	}
	//Models.SignUp(user)
	if Models.SignUp(user) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"user":    user,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"user":    user,
		})
	}
}

// CheckUsername When sign up, check if the username is unique
func CheckUsername(c *gin.Context) {
	username := c.PostForm("username")
	isTure := Services.CheckUsername(username)
	if isTure == true {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	} else {
		c.JSON(201, gin.H{
			"message": "error",
		})
	}
}

// LoginResult 登录结果结构
type LoginResult struct {
	Token string `json:"token"`
	Models.User
}

func generateToken(c *gin.Context, user Models.User) {
	j := &myjwt.JWT{
		[]byte("Tao"),
	}
	claims := myjwt.CustomClaims{
		user.User_id,
		user.Username,
		user.Password,
		jwtgo.StandardClaims{
			NotBefore: int64(time.Now().Unix() - 1000), // 签名生效时间
			ExpiresAt: int64(time.Now().Unix() + 3600), // 过期时间 一小时
			Issuer:    "Tao",                           //签名的发行者
		},
	}

	token, err := j.CreateToken(claims)

	if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": -1,
			"msg":    err.Error(),
		})
		return
	}

	log.Println(token)

	//data := LoginResult{
	//	User:  user,
	//	Token: token,
	//}
	c.JSON(http.StatusOK, gin.H{
		"status": 0,
		"msg":    "登录成功！",
		"user":   user,
		"token":  token,
	})
	return
}

// GetDataByTime 一个需要token认证的测试接口
func GetDataByTime(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		c.JSON(http.StatusOK, gin.H{
			"status": 0,
			"msg":    "token有效",
			"data":   claims,
		})
	}
}
