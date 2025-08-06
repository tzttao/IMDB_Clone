package Controllers

import (
	myjwt "backend/Middleware"
	"backend/Models"
	"backend/Services"
	"github.com/golang-jwt/jwt/v4"
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
	
	if user.Username == "" || user.Password == "" {
		c.JSON(http.StatusBadRequest, Models.ValidationErrorResponse("Username and password are required"))
		return
	}
	
	authenticatedUser := Models.LogIn(user)
	if authenticatedUser.User_id == 0 {
		c.JSON(http.StatusUnauthorized, Models.UnauthorizedResponse("Invalid username or password"))
		return
	}
	
	generateToken(c, authenticatedUser)
}

//SignUp for User
func SignUp(c *gin.Context) {
	username := c.PostForm("username")
	password := c.PostForm("password")
	email := c.PostForm("email")
	gender := c.PostForm("gender")
	ageStr := c.PostForm("age")
	
	if username == "" || password == "" || email == "" {
		c.JSON(http.StatusBadRequest, Models.ValidationErrorResponse("Username, password, and email are required"))
		return
	}
	
	age, err := strconv.Atoi(ageStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, Models.ValidationErrorResponse("Invalid age format"))
		return
	}
	
	user := Models.User{
		Username:  username,
		Password:  password,
		User_type: 0,
		Gender:    gender,
		Age:       age,
		Email:     email,
	}
	
	if Models.SignUp(user) == 0 {
		c.JSON(http.StatusInternalServerError, Models.InternalErrorResponse("Failed to create user"))
		return
	}
	
	user.Password = "" // Don't return password in response
	c.JSON(http.StatusOK, Models.SuccessResponse(user, "User created successfully"))
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

// 生成Token
func generateToken(c *gin.Context, user Models.User) {
	j := myjwt.NewJWT()
	expirationTime := time.Now().Add(24 * time.Hour) // 24小时过期
	claims := myjwt.CustomClaims{
		user.User_id,
		user.Username,
		"", // 不在token中存储密码
		jwt.RegisteredClaims{
			NotBefore: jwt.NewNumericDate(time.Now().Add(-1000 * time.Second)),
			ExpiresAt: jwt.NewNumericDate(expirationTime),
			Issuer:    "IMDB-Clone-System",
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
	user.Password = "" // Don't return password in response
	responseData := gin.H{
		"user":  user,
		"token": token,
	}
	c.JSON(http.StatusOK, Models.SuccessResponse(responseData, "Login successful"))
	return
}

//Update user profile

func UpdateUserProfile(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		var user Models.User
		user.User_id, _ = strconv.Atoi(c.PostForm("userId"))
		user.Age, _ = strconv.Atoi(c.PostForm("age"))
		user = Models.User{
			User_id:   user.User_id,
			Username:  c.PostForm("username"),
			Password:  c.PostForm("password"),
			User_type: 0,
			Gender:    c.PostForm("gender"),
			Age:       user.Age,
			Email:     c.PostForm("email"),
		}
		if Models.UpdateProfile(user) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "update error",
				"user":    user,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "update success",
				"user":    user,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 0,
			"msg":    "invalid token",
			"data":   claims,
		})
	}
}

//DeleteUser

func DeleteUser(c *gin.Context) {
	claims := c.MustGet("claims").(*myjwt.CustomClaims)
	if claims != nil {
		var userId, _ = strconv.Atoi(c.PostForm("userId"))
		if Models.DeleteUser(userId) == 0 {
			c.JSON(http.StatusNotFound, gin.H{
				"message": "delete error",
				"userId":  userId,
			})
		} else {
			c.JSON(http.StatusOK, gin.H{
				"message": "delete success",
				"userId":  userId,
			})
		}
	} else {
		c.JSON(http.StatusUnauthorized, gin.H{
			"status": 0,
			"msg":    "invalid token",
			"data":   claims,
		})
	}
}

func QueryUserInfoById(c *gin.Context) {
	var userId, _ = strconv.Atoi(c.PostForm("userId"))
	Models.QueryUserInfoByUserId(userId)
	c.JSON(http.StatusOK, gin.H{
		"userInfo": Models.QueryUserInfoByUserId(userId),
	})
}
