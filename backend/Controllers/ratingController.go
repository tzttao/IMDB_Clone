package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddRating(c *gin.Context) {
	var rating Models.Rating
	rating.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	rating.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	rating.Score, _ = strconv.ParseFloat(c.PostForm("score"), 64)
	//movie.Description = c.PostForm("desciption")
	res := Models.AddRating(rating)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"rating":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"rating":  res,
		})
	}
}

func UpdateRating(c *gin.Context) {
	var rating Models.Rating
	rating.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	rating.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	rating.Score, _ = strconv.ParseFloat(c.PostForm("score"), 64)
	//movie.Description = c.PostForm("desciption")
	res := Models.UpdateRating(rating)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"rating":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"rating":  res,
		})
	}
}

func DeleteRating(c *gin.Context) {
	var rating Models.Rating
	rating.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	rating.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	rating.Score, _ = strconv.ParseFloat(c.PostForm("score"), 64)
	//movie.Description = c.PostForm("desciption")
	res := Models.DeleteRating(rating)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"rating":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"rating":  res,
		})
	}
}

func ReadRating(c *gin.Context) {
	var movieId, userId int
	movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	userId, _ = strconv.Atoi(c.PostForm("userId"))
	res := Models.ReadRating(userId, movieId)
	rating := Models.Rating{}
	if res == rating {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"rating":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"rating":  res,
		})
	}

}

func ComputeAvgGrade(c *gin.Context) {
	var movieId int
	//rating.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	//rating.Score,_ = strconv.Atoi(c.PostForm("score"))
	//movie.Description = c.PostForm("desciption")
	res := Models.UpdateGrade(movieId)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"rating":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"rating":  res,
		})
	}
}
