package Router

import (
	userController "backend/Controllers"
	"github.com/gin-gonic/gin"
)

func Main() {
	router := gin.Default()

	// Simple group: v1
	v1 := router.Group("/v1")
	{
		v1.GET("/signin", userController.SignIn)
		v1.POST("/signup", userController.SignUp)
		//v1.POST("/read", readEndpoint)
	}
	router.LoadHTMLGlob("templates/*.html")
	router.GET("/form", func(ctx *gin.Context) {
		ctx.HTML(200, "form.html", gin.H{})
	})
	// Simple group: v2
	//v2 := router.Group("/v2")
	//{
	//	v2.POST("/login", loginEndpoint)
	//	v2.POST("/submit", submitEndpoint)
	//	v2.POST("/read", readEndpoint)
	//}

	router.Run(":8000")
}
