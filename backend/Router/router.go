package Router

import (
	castController "backend/Controllers"
	genreController "backend/Controllers"
	movieController "backend/Controllers"
	ratingController "backend/Controllers"
	userController "backend/Controllers"
	watchedListController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"net/http"
	"os"
)

func Main() {
	router := gin.Default()
	router.Use(jwt.CORSMiddleware())

	// Admin Router, put all customer router in this. // there is a Middleware that authentic admin
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", ratingController.AddGenre)
		adminGroup.POST("/updateGenre", ratingController.UpdateGenre)
		adminGroup.POST("/deleteGenre", ratingController.DeleteGenre)
		adminGroup.POST("/addMovie", genreController.AddMovie)
		adminGroup.POST("/updateMovie", genreController.UpdateMovie)
		adminGroup.DELETE("/deleteMovie", genreController.DeleteMovie)
		adminGroup.POST("/addMovieGenre", genreController.AddMovieGenre)
		adminGroup.POST("/addMovieCast", genreController.AddMovieCast)
	}

	// User Router, put all customer router in this.
	userGroup := router.Group("/user")
	{
		userGroup.POST("/logIn", userController.LogIn)
		userGroup.POST("/signUp", userController.SignUp)               // when query sth
		userGroup.POST("/checkUsername", userController.CheckUsername) // when update sth
		movieGroup := router.Group("/user/movie")
		{
			movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
			movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
			movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
			movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
			movieGroup.POST("/searchMovieByYear", genreController.SearchMovieByYear)
			movieGroup.POST("/topMovie", genreController.TopMovie)
		}
		castGroup := router.Group("/user/cast")
		{
			castGroup.POST("/searchCastById", castController.SearchCastById)
			castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
			castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
			castGroup.POST("/searchCastByName", castController.SearchCastByName)

		}
		reviewGroup := router.Group("/user/review")
		{
			reviewGroup.POST("/addReview", movieController.AddReview)
			reviewGroup.POST("/updateReview", movieController.UpdateReview)
			reviewGroup.POST("/deleteReview", movieController.DeleteReview)
			reviewGroup.POST("/readReview", movieController.ReadReview)
			reviewGroup.POST("/readReviewByMovieId", movieController.ReadReviewByMovieId)
		}
		watchedListGroup := router.Group("/user/watchedList")
		{
			watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
			watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
			watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
		}
		ratingGroup := router.Group("/user/rating")
		{
			ratingGroup.POST("/readRating", ratingController.ReadRating)
			ratingGroup.POST("/addRating", ratingController.AddRating)
			ratingGroup.POST("/updateRating", ratingController.UpdateRating)
			ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
			ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
		}
		genreGroup := router.Group("/user/genre")
		{
			genreGroup.POST("/searchGenreName", ratingController.SearchGenreName)
			genreGroup.POST("/searchGenreByGenreName", ratingController.SearchGenreByGenreName)
		}
	}
	userGroup.Use(jwt.JWTAuth())
	{
		userGroup.POST("/UpdateProfile", userController.UpdateUserProfile)
		userGroup.POST("/DeleteUser", userController.DeleteUser)
		userGroup.POST("/queryUserInfoById", userController.QueryUserInfoById)
	}

	// if user input invalid address, it'll send below msg
	router.NoRoute(func(c *gin.Context) {
		c.JSON(http.StatusNotFound, gin.H{"msg": "please use correct address!"})
	})
	// router port
	port := os.Getenv("SERVER_PORT")
	if port == "" {
		port = "8000"
	}
	router.Run(":" + port)
}
