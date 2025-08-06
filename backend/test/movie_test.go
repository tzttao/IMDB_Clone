package Test

import (
	genreController "backend/Controllers"
	movieController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//func TestSearchMovieWithGenre(t *testing.T) {
//	//var castName = ""
//	if Models.SearchMovieWithGenre(1) == nil {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func TestSearchMovieByName(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params := url.Values{}
	params.Add("searchName", "test")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/movie/searchMovieByName", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestSearchMovieByCast(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params2 := url.Values{}
	params2.Add("castName", "Tao")
	para2 := params2.Encode()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/user/movie/searchMovieByCast", strings.NewReader(para2))
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

}

func TestSearchMovieByMovieId(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params3 := url.Values{}
	params3.Add("movieId", "1")
	para3 := params3.Encode()
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/user/movie/searchMovieByMovieId", strings.NewReader(para3))
	req3.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w3, req3)
	assert.Equal(t, 200, w3.Code)

}

func TestSearchMovieWithGenre(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params4 := url.Values{}
	params4.Add("movieId", "1")
	para4 := params4.Encode()
	w4 := httptest.NewRecorder()
	req4, _ := http.NewRequest("POST", "/user/movie/searchMovieWithGenre", strings.NewReader(para4))
	req4.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w4, req4)
	assert.Equal(t, 200, w4.Code)

}

func TestSearchMovieByYear(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", movieController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params5 := url.Values{}
	params5.Add("year", "2021")
	para5 := params5.Encode()
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("POST", "/user/movie/searchMovieByYear", strings.NewReader(para5))
	req5.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w5, req5)
	assert.Equal(t, 200, w5.Code)

}

func TestTopMovie(t *testing.T) {

	router := gin.New()

	movieGroup := router.Group("/user/movie")
	{
		movieGroup.POST("/searchMovieByName", movieController.SearchMovieByName)
		movieGroup.POST("/searchMovieByCast", movieController.SearchMovieByCast)
		movieGroup.POST("/searchMovieByMovieId", movieController.SearchMovieByMovieId)
		movieGroup.POST("/searchMovieWithGenre", genreController.SearchMovieWithGenre)
		movieGroup.POST("/searchMovieByYear", movieController.SearchMovieByYear)
		movieGroup.POST("/topMovie", movieController.TopMovie)
	}

	params5 := url.Values{}
	//params5.Add("year", "2021")
	para5 := params5.Encode()
	w5 := httptest.NewRecorder()
	req5, _ := http.NewRequest("POST", "/user/movie/topMovie", strings.NewReader(para5))
	req5.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w5, req5)
	assert.Equal(t, 200, w5.Code)

}
