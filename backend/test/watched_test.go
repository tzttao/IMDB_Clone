package Test

import (
	watchedListController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddWatched(t *testing.T) {
	router := gin.New()
	watchedListGroup := router.Group("/user/watchedList")
	{
		watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
		watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
		watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
	}

	params := url.Values{}
	params.Add("userId", "1")
	params.Add("movieId", "2")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/watchedList/addWatchedList", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteWatched(t *testing.T) {
	router := gin.New()
	watchedListGroup := router.Group("/user/watchedList")
	{
		watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
		watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
		watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
	}

	params := url.Values{}
	params.Add("userId", "1")
	params.Add("movieId", "2")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/watchedList/deleteWatchedList", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestReadWatched(t *testing.T) {
	router := gin.New()
	watchedListGroup := router.Group("/user/watchedList")
	{
		watchedListGroup.POST("/addWatchedList", watchedListController.AddWatched)
		watchedListGroup.POST("/deleteWatchedList", watchedListController.DeleteWatched)
		watchedListGroup.POST("/readWatchedList", watchedListController.ReadWatched)
	}

	params := url.Values{}
	params.Add("userId", "1")
	//params.Add("movieId", "1")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/watchedList/readWatchedList", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
