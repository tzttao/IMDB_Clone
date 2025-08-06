package Test

import (
	reviewController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestAddReview(t *testing.T) {
	router := gin.New()
	reviewGroup := router.Group("/user/review")
	{
		reviewGroup.POST("/addReview", reviewController.AddReview)
		reviewGroup.POST("/updateReview", reviewController.UpdateReview)
		reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
		reviewGroup.POST("/readReview", reviewController.ReadReview)
		reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
	}

	params := url.Values{}
	params.Add("userId", "")
	params.Add("movieId", "3")
	params.Add("reviewContent", "orz")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/review/addReview", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestUpdateReview(t *testing.T) {
	router := gin.New()
	reviewGroup := router.Group("/user/review")
	{
		reviewGroup.POST("/addReview", reviewController.AddReview)
		reviewGroup.POST("/updateReview", reviewController.UpdateReview)
		reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
		reviewGroup.POST("/readReview", reviewController.ReadReview)
		reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
	}

	params := url.Values{}
	params.Add("userId", "2")
	params.Add("movieId", "1")
	params.Add("reviewContent", "QAQ")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/review/updateReview", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestDeleteReview(t *testing.T) {
	router := gin.New()
	reviewGroup := router.Group("/user/review")
	{
		reviewGroup.POST("/addReview", reviewController.AddReview)
		reviewGroup.POST("/updateReview", reviewController.UpdateReview)
		reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
		reviewGroup.POST("/readReview", reviewController.ReadReview)
		reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
	}

	params := url.Values{}
	params.Add("reviewID", "3")
	//params.Add("movieId", "1")
	//params.Add("reviewContent", "QAQ")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/review/deleteReview", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestReadReview(t *testing.T) {
	router := gin.New()
	reviewGroup := router.Group("/user/review")
	{
		reviewGroup.POST("/addReview", reviewController.AddReview)
		reviewGroup.POST("/updateReview", reviewController.UpdateReview)
		reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
		reviewGroup.POST("/readReview", reviewController.ReadReview)
		reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
	}

	params := url.Values{}
	params.Add("userId", "1")
	//params.Add("movieId", "1")
	//params.Add("reviewContent", "QAQ")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/review/readReview", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}

func TestReadReviewByMovieId(t *testing.T) {
	router := gin.New()
	reviewGroup := router.Group("/user/review")
	{
		reviewGroup.POST("/addReview", reviewController.AddReview)
		reviewGroup.POST("/updateReview", reviewController.UpdateReview)
		reviewGroup.POST("/deleteReview", reviewController.DeleteReview)
		reviewGroup.POST("/readReview", reviewController.ReadReview)
		reviewGroup.POST("/readReviewByMovieId", reviewController.ReadReviewByMovieId)
	}

	params := url.Values{}
	params.Add("movieId", "3")
	params.Add("pageNo", "1")
	params.Add("pageSize", "5")

	//params.Add("movieId", "1")
	//params.Add("reviewContent", "QAQ")
	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/review/readReviewByMovieId", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
