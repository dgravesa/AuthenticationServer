package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dgravesa/AuthenticationServer/model"
)

func newPutUserRequest(in model.UserLogin) *http.Request {
	f := url.Values{}

	if in.ID != 0 {
		f.Set("userid", fmt.Sprintf("%d", in.ID))
	}
	if in.Password != "" {
		f.Set("password", in.Password)
	}

	req := httptest.NewRequest("POST", "http://localhost/user", nil)
	req.Form = f
	return req
}

func Test_putUser_WithSubsequentLoginUsingNewCredentials_ReturnsValidSession(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode1 := http.StatusOK
	expectedCode2 := http.StatusCreated
	expectSession := true
	validLogin := validLogins[0]
	newPassword := "newP@ssw0rd"
	changeLogin := model.UserLogin{ID: validLogin.ID, Password: newPassword}
	changeReq := newPutUserRequest(changeLogin)
	changeRes := httptest.NewRecorder()
	loginReq := newPostLoginRequest(changeLogin)
	loginRes := httptest.NewRecorder()

	// Act
	putUser(changeRes, changeReq)
	postLogin(loginRes, loginReq)

	// Assert
	if changeRes.Code != expectedCode1 {
		t.Fatalf("expected status code = %d, received status code = %d", expectedCode1, changeRes.Code)
	}
	validatePostLoginResponse(loginRes, expectedCode2, expectSession, t)
}

func Test_putUser_WithSubsequentLoginUsingPreviousCredentials_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	initValidLogins()
	expectedCode1 := http.StatusOK
	expectedCode2 := http.StatusUnauthorized
	expectSession := false
	oldLogin := validLogins[0]
	newPassword := "newP@ssw0rd"
	changeLogin := model.UserLogin{ID: oldLogin.ID, Password: newPassword}
	changeReq := newPutUserRequest(changeLogin)
	changeRes := httptest.NewRecorder()
	loginReq := newPostLoginRequest(oldLogin)
	loginRes := httptest.NewRecorder()

	// Act
	putUser(changeRes, changeReq)
	postLogin(loginRes, loginReq)

	// Assert
	if changeRes.Code != expectedCode1 {
		t.Fatalf("expected status code = %d, received status code = %d", expectedCode1, changeRes.Code)
	}
	validatePostLoginResponse(loginRes, expectedCode2, expectSession, t)
}
