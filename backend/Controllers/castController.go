package Controllers

import (
	"backend/Models"
	"backend/Services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SearchCastById(c *gin.Context) {
	var castId int
	castId, _ = strconv.Atoi(c.PostForm("castId"))
	//movieId = c.PostForm("movieId")
	res := Models.SearchCastById(castId)
	cast := Models.Result{}
	if res == cast {
		//if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"cast":    res,
		})
	}
}

func SearchCastByMovieId(c *gin.Context) {
	var movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Services.SearchCastByMovieId(movieId)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"cast":    res,
		})
	}
}

func SearchRelativeMovieByCastId(c *gin.Context) {
	var castId, _ = strconv.Atoi(c.PostForm("castId"))
	res := Models.SearchRelativeMovieByCastId(castId)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"movie":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"movie":   res,
		})
	}
}
func SearchCastByName(c *gin.Context) {
	var castName = c.PostForm("castName")
	res := Models.SearchCastByName(castName)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"movie":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"movie":   res,
		})
	}
}
