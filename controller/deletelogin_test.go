package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dgravesa/AuthenticationServer/model"
)

func newDeleteLoginRequest(s model.Session) *http.Request {
	req := httptest.NewRequest("DELETE", "http://localhost/login", nil)
	query := make(url.Values)
	query.Set("userId", fmt.Sprint(s.UID))
	query.Set("key", s.Key)
	req.URL.RawQuery = query.Encode()
	return req
}

func Test_deleteLogin_OnValidRequest_ReturnsSuccess(t *testing.T) {
	// Arrange
	expectedCode := http.StatusOK
	initValidLogins()
	validLogin := validLogins[0]
	validSession, loginSucceeded := model.AuthenticateUser(validLogin)
	if !loginSucceeded {
		t.Fatalf("unable to set up test for delete login; initial login request failed")
	}
	req := newDeleteLoginRequest(validSession)
	res := httptest.NewRecorder()

	// Act
	deleteLogin(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}

func Test_deleteLogin_OnInvalidRequest_ReturnsNotFound(t *testing.T) {
	// Arrange
	expectedCode := http.StatusNotFound
	initValidLogins()
	for _, login := range validLogins {
		_, _ = model.AuthenticateUser(login)
	}
	fakeSession := model.Session{
		UID: validLogins[0].ID,
		Key: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef",
	}
	req := newDeleteLoginRequest(fakeSession)
	res := httptest.NewRecorder()

	// Act
	deleteLogin(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}
