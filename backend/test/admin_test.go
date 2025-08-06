package Test

import (
	genreController "backend/Controllers"
	jwt "backend/Middleware"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

//func TestAddGenre(t *testing.T) {
//	var genre = Models.Genre{Genre_id: 3, Genre_name: "3"}
//	if Models.AddGenre(genre) == 0 {
//		t.Error("result is wrong!")
//	} else {
//		t.Log("result is right")
//	}
//}

func TestAddGenre(t *testing.T) {
	router := gin.New()
	adminGroup := router.Group("/admin")
	adminGroup.Use(jwt.JWTAuth())
	{
		adminGroup.POST("/addGenre", genreController.AddGenre)

	}

	params := url.Values{}
	params.Add("genreId", "3")
	params.Add("genreName", "5555")

	//params.Add("password", "123")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/admin/addGenre", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

}
