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

type postUserInput struct {
	uid   uint64
	upass string
}

func newPostUserRequest(in postUserInput) *http.Request {
	f := url.Values{}

	if in.uid != 0 {
		f.Set("userid", fmt.Sprintf("%d", in.uid))
	}
	if in.upass != "" {
		f.Set("password", in.upass)
	}

	req := httptest.NewRequest("POST", "http://localhost/user", nil)
	req.Form = f
	return req
}

var postUserValidRequestInput = []postUserInput{
	{1, "password1"},
	{2121, "$#!_&*()<>"},
}

func validatePostUserResponse(in postUserInput, expectedCode, receivedCode int, t *testing.T) {
	var uidStr, upassStr string

	if in.uid != 0 {
		uidStr = fmt.Sprintf("%d", in.uid)
	} else {
		uidStr = "<empty>"
	}

	if in.upass != "" {
		upassStr = fmt.Sprintf("\"%s\"", in.upass)
	} else {
		upassStr = "<empty>"
	}

	if receivedCode != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d [userid = %s, password = %s]",
			expectedCode, receivedCode, uidStr, upassStr)
	}
}

func Test_postUser_WithValidRequest_ReturnsSuccess(t *testing.T) {
	expectedCode := http.StatusCreated
	model.SetUserRecordDataLayer(data.NewInMemoryUserRecordLayer())

	for _, in := range postUserValidRequestInput {
		// Arrange
		req := newPostUserRequest(in)
		res := httptest.NewRecorder()

		// Act
		postUser(res, req)

		// Assert
		if !model.UIDExists(in.uid) {
			t.Errorf("posted userid (%d) does not exist in data", in.uid)
		}
		validatePostUserResponse(in, expectedCode, res.Code, t)
	}
}

func Test_postUser_MissingUserIDField_ReturnsBadRequest(t *testing.T) {
	// Arrange
	expectedCode := http.StatusBadRequest
	model.SetUserRecordDataLayer(data.NewInMemoryUserRecordLayer())
	in := postUserInput{0, "password1"}
	req := newPostUserRequest(in)
	res := httptest.NewRecorder()

	// Act
	postUser(res, req)

	// Assert
	validatePostUserResponse(in, expectedCode, res.Code, t)
}

func Test_postUser_MissingPasswordField_ReturnsBadRequest(t *testing.T) {
	// Arrange
	expectedCode := http.StatusBadRequest
	model.SetUserRecordDataLayer(data.NewInMemoryUserRecordLayer())
	in := postUserInput{101, ""}
	req := newPostUserRequest(in)
	res := httptest.NewRecorder()

	// Act
	postUser(res, req)

	// Assert
	validatePostUserResponse(in, expectedCode, res.Code, t)
}

func Test_postUser_ExistingUserID_ReturnsConflict(t *testing.T) {
	// Arrange
	expectedCode1 := http.StatusCreated
	expectedCode2 := http.StatusConflict
	model.SetUserRecordDataLayer(data.NewInMemoryUserRecordLayer())
	in1 := postUserInput{174, "password1"}
	in2 := postUserInput{174, "password2"}
	req1 := newPostUserRequest(in1)
	req2 := newPostUserRequest(in2)
	res1 := httptest.NewRecorder()
	res2 := httptest.NewRecorder()

	// Act
	postUser(res1, req1)
	postUser(res2, req2)

	// Assert
	validatePostUserResponse(in1, expectedCode1, res1.Code, t)
	validatePostUserResponse(in2, expectedCode2, res2.Code, t)
}
