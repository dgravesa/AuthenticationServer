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

func Test_getValidate_WithDeletedSession_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	expectedCode1 := http.StatusOK
	expectedCode2 := http.StatusUnauthorized
	initValidLogins()
	for _, login := range validLogins {
		_, _ = model.AuthenticateUser(login)
	}
	validLogin := validLogins[0]
	validSession, loginSucceeded := model.AuthenticateUser(validLogin)
	if !loginSucceeded {
		t.Fatalf("unable to set up test for get validate; failed to log in with valid credentials")
	}
	req := newGetValidateRequest(validSession)
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	// Act
	getValidate(res1, req)
	model.DeleteSession(validSession)
	getValidate(res2, req)

	// Assert
	if res1.Code != expectedCode1 {
		t.Errorf("validate before session deleted: expected status code = %d, received status code = %d",
			expectedCode1, res1.Code)
	}
	if res2.Code != expectedCode2 {
		t.Errorf("validate after session deleted: expected status code = %d, received status code = %d",
			expectedCode2, res2.Code)
	}
}

// TODO: consider moving more stateful, functional tests like this to a component test tier.
func Test_getValidate_OnSessionAfterPasswordChange_ReturnsUnauthorized(t *testing.T) {
	// Arrange
	expectedCodeBefore := http.StatusOK
	expectedCodePutUser := http.StatusOK
	expectedCodeAfter := http.StatusUnauthorized
	initValidLogins()
	oldLogin := validLogins[0]
	oldSession, loginSucceeded := model.AuthenticateUser(oldLogin)
	if !loginSucceeded {
		t.Fatalf("unable to set up test for get validate; failed to log in with valid credentials")
	}
	req := newGetValidateRequest(oldSession)
	resBefore := httptest.NewRecorder()
	resPutUser := httptest.NewRecorder()
	resAfter := httptest.NewRecorder()
	newLogin := model.UserLogin{ID: oldLogin.ID, Password: "newP@ssw0rd"}
	putUserReq := newPutUserRequest(newLogin)

	// Act
	getValidate(resBefore, req)
	putUser(resPutUser, putUserReq)
	getValidate(resAfter, req)

	// Assert
	if resBefore.Code != expectedCodeBefore {
		t.Errorf("validate before login changed: expected status code = %d, received status code = %d",
			expectedCodeBefore, resBefore.Code)
	}
	if resPutUser.Code != expectedCodePutUser {
		t.Errorf("error on login change request: expected status code = %d, received status code = %d",
			expectedCodePutUser, resPutUser.Code)
	}
	if resAfter.Code != expectedCodeAfter {
		t.Errorf("validate before login changed: expected status code = %d, received status code = %d",
			expectedCodeAfter, resAfter.Code)
	}
}
