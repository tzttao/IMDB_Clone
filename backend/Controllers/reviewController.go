package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddReview(c *gin.Context) {
	var review Models.Review
	//review.Review_id, _ = strconv.Atoi(c.PostForm("reviewId"))
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	review.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	review.Review_content = c.PostForm("reviewContent")
	res := Models.AddReview(review)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"review":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"review":  res,
		})
	}
}

func UpdateReview(c *gin.Context) {
	var review Models.Review
	review.Review_id, _ = strconv.Atoi(c.PostForm("reviewId"))
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	review.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	review.Review_content = c.PostForm("reviewContent")
	res := Models.UpdateReview(review)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"review":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"review":  res,
			"code":    200,
		})
	}
}

func DeleteReview(c *gin.Context) {
	var review Models.Review
	review.Review_id, _ = strconv.Atoi(c.PostForm("reviewID"))
	res := Models.DeleteReview(review.Review_id)
	if res == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"review":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"review":  res,
			"code":    200,
		})
	}
}

func ReadReview(c *gin.Context) {
	var review Models.Review
	review.User_id, _ = strconv.Atoi(c.PostForm("userId"))
	res := Models.ReadReview(review.User_id)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
			"review":  res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"message": "success",
			"review":  res,
			"code":    200,
		})
	}
}
func ReadReviewByMovieId(c *gin.Context) {
	var review Models.Review
	review.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	var pageNo, _ = strconv.Atoi(c.PostForm("pageNo"))
	var pageSize, _ = strconv.Atoi(c.PostForm("pageSize"))
	res := Models.ReadReviewByMovieId(review.Movie_id, pageNo, pageSize)
	count := Models.CountNumOfReview(review.Movie_id)
	if res == nil {
		c.JSON(http.StatusNotFound, gin.H{
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"review":  res,
			"count":   count,
		})
	}
}
