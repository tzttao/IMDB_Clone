package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddWatched(c *gin.Context) {
	var WatchList Models.Watch_list
	WatchList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	WatchList.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.AddWatched(WatchList)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "error",
			"watchList": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "success",
			"watchList": res,
		})
	}
}

func DeleteWatched(c *gin.Context) {
	//var Review_id int
	var WatchList Models.Watch_list
	WatchList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	WatchList.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	//res := Models.DeleteWatched(WatchList.User_id, WatchList.Movie_id)
	res := Models.DeleteWatched(WatchList)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "error",
			"watchList": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "success",
			"watchList": res,
		})
	}
}

func ReadWatched(c *gin.Context) {
	//var Review_id int
	var WatchList Models.Watch_list
	WatchList.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	res := Models.ReadWatched(WatchList.User_id)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message":   "error",
			"watchList": res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message":   "success",
			"watchList": res,
		})
	}
}
