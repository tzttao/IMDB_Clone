package Router

import (
	userController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Main() {
	router := gin.Default()

	// Customer Router, put all customer router in this.
	customerGroup := router.Group("/customer")
	{
		customerGroup.POST("/add", userController.LogIn)      // when create sth
		customerGroup.GET("/query", userController.LogIn)     // when query sth
		customerGroup.PUT("/update", userController.LogIn)    // when update sth
		customerGroup.DELETE("/delete", userController.LogIn) // when delete sth
	}

	// User Router, put all customer router in this.
	userGroup := router.Group("/user")
	{
		userGroup.POST("/LogIn", userController.LogIn)
		userGroup.POST("/SignUp", userController.SignUp)               // when query sth
		userGroup.POST("/CheckUsername", userController.CheckUsername) // when update sth
		userGroup.DELETE("/delete", userController.LogIn)              // when delete sth
	}
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.GET("/jwt", userController.GetDataByTime)
	}
	// Admin Router, put all customer router in this. // there is a Middleware that authentic admin
	adminGroup := router.Group("/admin", gin.BasicAuth(gin.Accounts{"admin": "123"}))
	{
		adminGroup.GET("/", userController.LogIn)
		adminGroup.GET("/query", userController.LogIn)     // when query sth
		adminGroup.PUT("/update", userController.LogIn)    // when update sth
		adminGroup.DELETE("/delete", userController.LogIn) // when delete sth
	}

	// if user input invalid address, it'll send below msg
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "please use correct address!"})
	})
	// router port
	router.Run(":8000")
}
