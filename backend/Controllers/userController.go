package Controllers

import (
	"backend/Models"
	userService "backend/Services"
	"github.com/gin-gonic/gin"
	"time"
)

//Sign in
func SignIn(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	ctx.JSON(200, gin.H{
		"code":    1,
		"message": "success",
		"data":    userService.Signin(username, password),
		//"username": username,
		//"password": password,
		//"gender": gender,
		//"birthday": birthday,
	})
	//userService.Signin(username, passwsord)
}

//SignUp
func SignUp(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")
	gender := ctx.PostForm("gender")
	birthday := ctx.PostForm("birthday")

	//t := time.Now() //当前时间
	//t.Unix()        //时间戳
	//
	//ts := t.Format(birthday) //time转string
	//fmt.Println(ts)
	st, _ := time.Parse("2006-01-02 15:04:05", birthday) //string转time
	//fmt.Println(st)
	//a := time.Unix(t, 0).Format(birthday)

	ctx.JSON(200, gin.H{
		"username": username,
		"password": password,
		"gender":   gender,
		"birthday": st,
	})

	ggg := uint(1)
	if gender == "female" {
		ggg = uint(1)
	} else {
		ggg = uint(0)
	}

	user := Models.User{
		//User_id:   ,
		Username:  username,
		Password:  password,
		User_type: 0,
		Gender:    ggg,
		Birthday:  st,
	}
	userService.AddUser(user)
}
