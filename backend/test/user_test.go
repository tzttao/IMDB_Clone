package Test

import (
	userController "backend/Controllers"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
	"testing"
)

func TestLogIn(t *testing.T) {

	router := gin.New()
	customerGroup := router.Group("/user")
	{
		customerGroup.POST("/login", userController.LogIn)                 // when create sth
		customerGroup.POST("/signUp", userController.SignUp)               // when query sth
		customerGroup.POST("/checkUsername", userController.CheckUsername) // when update sth

	}

	params := url.Values{}
	params.Add("username", "qwe")
	params.Add("password", "123456")
	para1 := params.Encode()
	w := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/user/login", strings.NewReader(para1))
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w, req)
	assert.Equal(t, 200, w.Code)

	params2 := url.Values{}
	params2.Add("username", "qwe")
	params2.Add("password", "123")
	params2.Add("user_type", "0")
	params2.Add("gender", "female")
	params2.Add("age", "12")
	params2.Add("email", "123@")
	para2 := params2.Encode()
	w2 := httptest.NewRecorder()
	req2, _ := http.NewRequest("POST", "/user/signUp", strings.NewReader(para2))
	req2.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w2, req2)
	assert.Equal(t, 200, w2.Code)

	params3 := url.Values{}
	params3.Add("username", "qwe")
	para3 := params3.Encode()
	w3 := httptest.NewRecorder()
	req3, _ := http.NewRequest("POST", "/user/checkUsername", strings.NewReader(para3))
	req3.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	router.ServeHTTP(w3, req3)
	assert.Equal(t, 201, w3.Code)

}
