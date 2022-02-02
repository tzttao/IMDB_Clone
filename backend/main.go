package main

import (
	Database "backend/Database"
	Service "backend/Services"
)

func main() {
	//router := gin.Default()
	//router.LoadHTMLGlob("templates/*.html")
	//router.GET("/", func(ctx *gin.Context) {
	//	ctx.HTML(200, "index.html", gin.H{})
	//})
	////定义新的路由，传入data
	//router.GET("/data", func(ctx *gin.Context) {
	//	ctx.HTML(200, "data.html", gin.H{"data": "this is data!"})
	//})
	//router.GET("/form", func(ctx *gin.Context) {
	//	ctx.HTML(200, "form.html", gin.H{})
	//})
	//
	////定义新的路由，返回json类型数据
	//router.GET("/json", func(ctx *gin.Context) {
	//	ctx.JSON(200, gin.H{
	//		"result":    "ok",
	//		"data":      "this is data!",
	//		"developer": "Tao!",
	//	})
	//})
	////接受POST数据，返回json类型数据
	//router.POST("/service", func(ctx *gin.Context) {
	//	username := ctx.PostForm("username")
	//	ctx.JSON(200, gin.H{
	//		"result":    "ok",
	//		"hello":     username,
	//		"developer": "Tao!",
	//	})
	//})
	//router.Run(":8000")

	//Database.StructInsert()
	//Database.StructUpdate()
	//Database.StructQueryField()
	//Database.StructQueryAllField()
	//Database.StructDel()
	//Database.StructTx()
	//Database.RawQueryField()
	//Database.RawQueryAllField()
	//Service.GetOne()
	Service.DeleteOne()
	//Service.UpdateOne()
	defer Database.Db.Close()
}
