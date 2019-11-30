package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dgravesa/AuthenticationServer/data"
	"github.com/dgravesa/AuthenticationServer/model"
)

var validLogins = []model.UserLogin{
	{512, "Password1"},
	{512136, "<>LL$#rT^6"},
	{9099, "qwrty$@@!..;"},
	{12371, "5215"},
	{43, "{[password]}:"},
}

func initValidLogins() {
	model.SetUserRecordDataLayer(data.NewInMemoryUserRecordLayer())
	model.SetSessionDataLayer(data.NewInMemorySessionLayer())

	for _, login := range validLogins {
		model.AddUserLogin(login)
	}
}

func newPostLoginRequest(l model.UserLogin) *http.Request {
	f := url.Values{}

	if l.ID != 0 {
		f.Set("userid", fmt.Sprintf("%d", l.ID))
	}
	if l.Password != "" {
		f.Set("password", l.Password)
	}

	req := httptest.NewRequest("POST", "http://localhost/login", nil)
	req.Form = f
	return req
}

func validatePostLoginResponse(res *httptest.ResponseRecorder, expectedCode int, expectSession bool, t *testing.T) {
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}

	_, err := model.DecodeSessionFromJSON(res.Result().Body)
	receivedSession := (err == nil)

	if receivedSession != expectSession {
		t.Errorf("expected session = %t, received session = %t", expectSession, receivedSession)
	}
}

func Test_postLogin_WithValidCredentials_ReturnsSession(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode := http.StatusCreated
	expectSession := true

	for _, login := range validLogins {
		req := newPostLoginRequest(login)
		res := httptest.NewRecorder()

		// Act
		postLogin(res, req)

		// Assert
		validatePostLoginResponse(res, expectedCode, expectSession, t)
	}
}

var fakeLogins = []model.UserLogin{
	{511, "Password1"},
	{512136, "NotPassword"},
	{315, "&(*dggr$E#"},
}

func Test_postLogin_WithInvalidCredentials_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode := http.StatusUnauthorized
	expectSession := false

	for _, fakeLogin := range fakeLogins {
		req := newPostLoginRequest(fakeLogin)
		res := httptest.NewRecorder()

		// Act
		postLogin(res, req)

		// Assert
		validatePostLoginResponse(res, expectedCode, expectSession, t)
	}
}

func Test_postLogin_MissingPassword_ReturnsBadRequest(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode := http.StatusBadRequest
	expectSession := false
	req := newPostLoginRequest(model.UserLogin{512, ""})
	res := httptest.NewRecorder()

	// Act
	postLogin(res, req)

	// Assert
	validatePostLoginResponse(res, expectedCode, expectSession, t)
}

func Test_postLogin_MissingUID_ReturnsBadRequest(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode := http.StatusBadRequest
	expectSession := false
	req := newPostLoginRequest(model.UserLogin{0, "Password1"})
	res := httptest.NewRecorder()

	// Act
	postLogin(res, req)

	// Assert
	validatePostLoginResponse(res, expectedCode, expectSession, t)
}
