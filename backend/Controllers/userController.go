package Controllers

import (
	"backend/Models"
	userService "backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

//Sign in
func SignIn(ctx *gin.Context) {
	ctx.JSON(http.StatusOK, gin.H{
		"code":    1,
		"message": "success",
		"data":    userService.Signin(),
	})
}

//SignUp
func SignUp(ctx *gin.Context) {
	username := ctx.PostForm("username")
	ctx.JSON(200, gin.H{
		"result":    "ok",
		"hello":     username,
		"developer": "Tao!",
	})
	user := Models.User{
		User_id:   3,
		Username:  username,
		Password:  "123456",
		User_type: 0,
		Gender:    0,
		Birthday:  time.Now(),
	}
	userService.AddUser(user)
}
