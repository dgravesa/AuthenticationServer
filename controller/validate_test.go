package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/dgravesa/AuthenticationServer/model"
)

func newGetValidateRequest(s model.Session) *http.Request {
	req := httptest.NewRequest("GET", "http://localhost/validate", nil)
	query := make(url.Values)
	query.Set("userId", fmt.Sprint(s.UID))
	query.Set("key", s.Key)
	req.URL.RawQuery = query.Encode()
	return req
}

func Test_getValidate_WithValidSession_ReturnsSuccess(t *testing.T) {
	// Arrange
	expectedCode := http.StatusOK
	initValidLogins()                             // from postlogin_test.go
	validLogin := validLogins[len(validLogins)-1] // from postlogin_test.go
	validSession, loginSucceeded := model.AuthenticateUser(validLogin)
	if !loginSucceeded {
		t.Fatalf("failed to create session prior to test")
	}
	req := newGetValidateRequest(validSession)
	res := httptest.NewRecorder()

	// Act
	getValidate(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}

func Test_getValidate_WithInvalidSessionKey_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	expectedCode := http.StatusUnauthorized
	initValidLogins()                             // from postlogin_test.go
	validLogin := validLogins[len(validLogins)-1] // from postlogin_test.go
	validSession, loginSucceeded := model.AuthenticateUser(validLogin)
	if !loginSucceeded {
		t.Fatalf("failed to create session prior to test")
	}
	invalidSession := model.Session{
		UID: validSession.UID,
		Key: "0123456789abcdef0123456789abcdef0123456789abcdef0123456789abcdef", // fake key
	}
	req := newGetValidateRequest(invalidSession)
	res := httptest.NewRecorder()

	// Act
	getValidate(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}

func Test_getValidate_WithInvalidSessionUID_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	expectedCode := http.StatusUnauthorized
	initValidLogins()                             // from postlogin_test.go
	validLogin := validLogins[len(validLogins)-1] // from postlogin_test.go
	validSession, loginSucceeded := model.AuthenticateUser(validLogin)
	if !loginSucceeded {
		t.Fatalf("failed to create session prior to test")
	}
	invalidSession := model.Session{
		UID: 9099, // other user
		Key: validSession.Key,
	}
	req := newGetValidateRequest(invalidSession)
	res := httptest.NewRecorder()

	// Act
	getValidate(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d", expectedCode, res.Code)
	}
}
