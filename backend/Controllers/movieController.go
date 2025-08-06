package Controllers

import (
	"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func SearchMovieByName(c *gin.Context) {
	var movieName string
	//user.Username = c.PostForm("username")
	movieName = c.PostForm("searchName")
	res := Models.SearchMovieByName(movieName)
	//movie= res
	if len(res) == 0 {
		c.JSON(http.StatusNotFound, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"movie":   res,
		})
	}
}

func SearchMovieByCast(c *gin.Context) {
	var castName string
	//user.Username = c.PostForm("username")
	castName = c.PostForm("castName")
	res := Models.SearchMovieByCast(castName)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"movie":   res,
		})
	}

}

func SearchMovieByYear(c *gin.Context) {
	var year int
	//user.Username = c.PostForm("username")
	year, _ = strconv.Atoi(c.PostForm("year"))
	res := Models.SearchMovieByYear(year)
	if len(res) == 0 {
		c.JSON(http.StatusOK, gin.H{
			"code":    404,
			"message": "error",
		})
	} else {
		c.JSON(http.StatusOK, gin.H{
			"code":    200,
			"message": "success",
			"movie":   res,
		})
	}

}

func SearchMovieByMovieId(c *gin.Context) {
	var movieId int
	movieId, _ = strconv.Atoi(c.PostForm("movieId"))
	//movieId = c.PostForm("movieId")
	res := Models.SearchMovieByMovieId(movieId)
	movie := Models.Movie{}
	if res == movie {
		//if res == 0 {
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

func AddMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	movie.Movie_name = c.PostForm("movieName")
	movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	movie.Description = c.PostForm("desciption")
	res := Models.AddMovie(movie)
	if res == 0 {
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

func TopMovie(c *gin.Context) {
	//var movie Models.Movie
	//movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	//movie.Movie_name = c.PostForm("movieName")
	//movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	//movie.Description = c.PostForm("desciption")
	res := Models.TopMovie()
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

func UpdateMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	movie.Movie_name = c.PostForm("movieName")
	movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	movie.Description = c.PostForm("desciption")
	//res := Models.AddMovie(movie)
	res := Models.UpdateMovie(movie)
	if res == 0 {
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

func DeleteMovie(c *gin.Context) {
	var movie Models.Movie
	movie.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	res := Models.DeleteMovie(movie.Movie_id)
	if res == 0 {
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

func AddMovieGenre(c *gin.Context) {
	var movieGenre Models.Movie_genre
	movieGenre.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	movieGenre.Genre_id, _ = strconv.Atoi(c.PostForm("genreId"))
	//movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	//movie.Description = c.PostForm("desciption")
	res := Models.AddMovieGenre(movieGenre)
	if res == 0 {
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

func AddMovieCast(c *gin.Context) {
	var movieCast Models.Movie_cast
	movieCast.Cast_id, _ = strconv.Atoi(c.PostForm("castId"))
	movieCast.Movie_id, _ = strconv.Atoi(c.PostForm("movieId"))
	//movie.Year, _ = strconv.Atoi(c.PostForm("year"))
	//movie.Description = c.PostForm("desciption")
	res := Models.AddMovieCast(movieCast)
	if res == 0 {
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
