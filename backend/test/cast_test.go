package Test

import (
	castController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestSearchCastById(t *testing.T) {
	router := gin.New()
	castGroup := router.Group("/user/cast")
	{
		castGroup.POST("/searchCastById", castController.SearchCastById)
		castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
		castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
		castGroup.POST("/searchCastByName", castController.SearchCastByName)
	}

	params := url.Values{}
	params.Add("castId", "1")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/cast/searchCastById", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestSearchCastByMovieId(t *testing.T) {
	router := gin.New()
	castGroup := router.Group("/user/cast")
	{
		castGroup.POST("/searchCastById", castController.SearchCastById)
		castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
		castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
		castGroup.POST("/searchCastByName", castController.SearchCastByName)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/cast/searchCastByMovieId", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestSearchRelativeMovieByCastId(t *testing.T) {
	router := gin.New()
	castGroup := router.Group("/user/cast")
	{
		castGroup.POST("/searchCastById", castController.SearchCastById)
		castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
		castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
		castGroup.POST("/searchCastByName", castController.SearchCastByName)
	}

	params := url.Values{}
	params.Add("castId", "1")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/cast/searchRelativeMovieByCastId", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestSearchCastByName(t *testing.T) {
	router := gin.New()
	castGroup := router.Group("/user/cast")
	{
		castGroup.POST("/searchCastById", castController.SearchCastById)
		castGroup.POST("/searchCastByMovieId", castController.SearchCastByMovieId)
		castGroup.POST("/searchRelativeMovieByCastId", castController.SearchRelativeMovieByCastId)
		castGroup.POST("/searchCastByName", castController.SearchCastByName)
	}

	params := url.Values{}
	params.Add("castName", "Tao")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/cast/searchCastByName", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
