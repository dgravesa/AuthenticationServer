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

type deleteUserInput struct {
	uid uint64
}

func newDeleteUserRequest(in deleteUserInput) *http.Request {
	f := url.Values{}

	if in.uid != 0 {
		f.Set("userid", fmt.Sprintf("%d", in.uid))
	}

	req := httptest.NewRequest("DELETE", "http://localhost/user", nil)
	req.Form = f
	return req
}

func validateDeleteUserResponse(in deleteUserInput, expectedCode, receivedCode int, t *testing.T) {
	var uidStr string

	if in.uid != 0 {
		uidStr = fmt.Sprintf("%d", in.uid)
	} else {
		uidStr = "<empty>"
	}

	if receivedCode != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d [userid = %s]",
			expectedCode, receivedCode, uidStr)
	}
}

func Test_deleteUser_OnExistingUser_ReturnsSuccess(t *testing.T) {
	// Arrange
	expectedCode := http.StatusOK
	var uidToDelete uint64 = 4765
	var uidToPersist uint64 = 5167
	model.SetDataLayer(data.NewInMemoryLayer())
	model.AddUserLogin(model.UserLogin{uidToDelete, "Password1"})
	model.AddUserLogin(model.UserLogin{uidToPersist, "Password2"})
	if !model.UIDExists(uidToDelete) {
		t.Fatalf("failed to add deletable uid prior to test")
	}
	if !model.UIDExists(uidToPersist) {
		t.Fatalf("failed to add persistent uid prior to test")
	}
	in := deleteUserInput{uidToDelete}
	req := newDeleteUserRequest(in)
	res := httptest.NewRecorder()

	// Act
	deleteUser(res, req)

	// Assert
	if model.UIDExists(uidToDelete) {
		t.Errorf("deleted userid (%d) still exists in data", uidToDelete)
	}
	if !model.UIDExists(uidToPersist) {
		t.Errorf("persistent userid (%d) no longer exists in data", uidToPersist)
	}
	validateDeleteUserResponse(in, expectedCode, res.Code, t)
}

func Test_deleteUser_OnNonexistentUser_ReturnsNotFound(t *testing.T) {
	// Arrange
	expectedCode := http.StatusNotFound
	var uidToPersist uint64 = 151899
	var uidToDelete uint64 = 444125
	model.SetDataLayer(data.NewInMemoryLayer())
	model.AddUserLogin(model.UserLogin{uidToPersist, "PasswordExistingUser"})
	if !model.UIDExists(uidToPersist) {
		t.Fatalf("failed to add persistent uid prior to test")
	}
	in := deleteUserInput{uidToDelete}
	req := newDeleteUserRequest(in)
	res := httptest.NewRecorder()

	// Act
	deleteUser(res, req)

	// Assert
	if !model.UIDExists(uidToPersist) {
		t.Errorf("persistent userid (%d) no longer exists in data", uidToPersist)
	}
	validateDeleteUserResponse(in, expectedCode, res.Code, t)
}

func Test_deleteUser_MissingUserID_ReturnsBadRequest(t *testing.T) {
	// Arrange
	expectedCode := http.StatusBadRequest
	model.SetDataLayer(data.NewInMemoryLayer())
	in := deleteUserInput{0}
	req := newDeleteUserRequest(in)
	res := httptest.NewRecorder()

	// Act
	deleteUser(res, req)

	// Assert
	validateDeleteUserResponse(in, expectedCode, res.Code, t)
}
