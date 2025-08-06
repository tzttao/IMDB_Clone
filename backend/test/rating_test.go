package Test

import (
	ratingController "backend/Controllers"
	"github.com/stretchr/testify/assert"

	//"backend/Models"
	"github.com/gin-gonic/gin"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddRating(t *testing.T) {
	router := gin.New()
	ratingGroup := router.Group("/user/rating")
	{
		ratingGroup.POST("/readRating", ratingController.ReadRating)
		ratingGroup.POST("/addRating", ratingController.AddRating)
		ratingGroup.POST("/updateRating", ratingController.UpdateRating)
		ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
		ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	params.Add("userId", "1")
	params.Add("score", "10")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/rating/addRating", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestReadRating(t *testing.T) {
	router := gin.New()
	ratingGroup := router.Group("/user/rating")
	{
		ratingGroup.POST("/readRating", ratingController.ReadRating)
		ratingGroup.POST("/addRating", ratingController.AddRating)
		ratingGroup.POST("/updateRating", ratingController.UpdateRating)
		ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
		ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	params.Add("userId", "1")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/rating/readRating", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestUpdateRating(t *testing.T) {
	router := gin.New()
	ratingGroup := router.Group("/user/rating")
	{
		ratingGroup.POST("/readRating", ratingController.ReadRating)
		ratingGroup.POST("/addRating", ratingController.AddRating)
		ratingGroup.POST("/updateRating", ratingController.UpdateRating)
		ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
		ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	params.Add("userId", "1")
	params.Add("score", "8")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/rating/updateRating", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestComputeAvgGrade(t *testing.T) {
	router := gin.New()
	ratingGroup := router.Group("/user/rating")
	{
		ratingGroup.POST("/readRating", ratingController.ReadRating)
		ratingGroup.POST("/addRating", ratingController.AddRating)
		ratingGroup.POST("/updateRating", ratingController.UpdateRating)
		ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
		ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	//params.Add("userId", "1")
	//params.Add("score", "8")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/rating/computeAvgGrade", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteRating(t *testing.T) {
	router := gin.New()
	ratingGroup := router.Group("/user/rating")
	{
		ratingGroup.POST("/readRating", ratingController.ReadRating)
		ratingGroup.POST("/addRating", ratingController.AddRating)
		ratingGroup.POST("/updateRating", ratingController.UpdateRating)
		ratingGroup.POST("/computeAvgGrade", ratingController.ComputeAvgGrade)
		ratingGroup.POST("/deleteRating", ratingController.DeleteRating)
	}

	params := url.Values{}
	params.Add("movieId", "1")
	params.Add("userId", "1")
	params.Add("score", "8")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/rating/deleteRating", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
