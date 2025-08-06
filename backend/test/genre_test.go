package Test

import (
	genreController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//func TestSearchGenreName(t *testing.T) {
//	//var castName = ""
//	if Models.SearchGenreName("1") == nil {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func TestSearchGenreName(t *testing.T) {

	router := gin.New()

	genreGroup := router.Group("/user/genre")
	{
		genreGroup.POST("/searchGenreName", genreController.SearchGenreName)
		genreGroup.POST("/searchGenreByGenreName", genreController.SearchGenreByGenreName)
	}

	params := url.Values{}
	params.Add("genreName", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/genre/searchGenreName", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestSearchGenreByGenreName(t *testing.T) {

	router := gin.New()

	genreGroup := router.Group("/user/genre")
	{
		genreGroup.POST("/searchGenreName", genreController.SearchGenreName)
		genreGroup.POST("/searchGenreByGenreName", genreController.SearchGenreByGenreName)
	}

	params := url.Values{}
	params.Add("genreName", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/genre/searchGenreByGenreName", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
