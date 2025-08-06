package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func AddGenre(c *gin.Context) {
	var genre Models.Genre
	genre.Genre_id, _ = strconv.Atoi(c.PostForm("genreId"))
	genre.Genre_name = c.PostForm("genreName")
	//rating.Score, _ = strconv.Atoi(c.PostForm("score"))
	//movie.Description = c.PostForm("desciption")
	res := Models.AddGenre(genre)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"genre":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}
}

func UpdateGenre(c *gin.Context) {
	var genre Models.Genre
	genre.Genre_id, _ = strconv.Atoi(c.PostForm("genreId"))
	genre.Genre_name = c.PostForm("genreName")
	//rating.Score, _ = strconv.Atoi(c.PostForm("score"))
	//movie.Description = c.PostForm("desciption")
	res := Models.AddGenre(genre)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"genre":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}
}

func DeleteGenre(c *gin.Context) {
	var genreId int
	genreId, _ = strconv.Atoi(c.PostForm("genreId"))
	res := Models.DeleteGenre(genreId)
	if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"genre":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}
}

func SearchGenreName(c *gin.Context) {
	var genreName string
	genreName = c.PostForm("genreName")
	//movieId = c.PostForm("movieId")
	res := Models.SearchGenreName(genreName)
	//rating := Models.Rating{}
	if len(res) == 0 {
		//if res == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
			"genre":   res,
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}

}

func SearchGenreByGenreName(c *gin.Context) {
	var genreName string
	genreName = c.PostForm("genreName")
	res := Models.SearchGenre(genreName)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}

}
func SearchMovieWithGenre(c *gin.Context) {
	var movieId int
	//user.Username = c.PostForm("username")
	movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.SearchMovieWithGenre(movieId)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"genre":   res,
		})
	}

}
