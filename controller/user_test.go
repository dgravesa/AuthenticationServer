package controller

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"strings"
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

	return httptest.NewRequest("POST", "http://localhost/user", strings.NewReader(f.Encode()))
}

var postUserValidRequestInput = []postUserInput{
	{1, "password1"},
	{2121, "$#!_&*()<>"},
}

func Test_postUser_WithValidRequest_ReturnsSuccess(t *testing.T) {
	expectedCode := http.StatusCreated
	model.SetDataLayer(data.NewInMemoryLayer())

	for _, in := range postUserValidRequestInput {
		// Arrange
		req := newPostUserRequest(in)
		res := httptest.NewRecorder()

		// Act
		postUser(res, req)

		// Assert
		if res.Code != expectedCode {
			t.Errorf("expected status code = %d, received status code = %d [userid = %d, password = \"%s\"]",
				expectedCode, res.Code, in.uid, in.upass)
		}
	}
}

func Test_postUser_MissingUserIDField_ReturnsBadRequest(t *testing.T) {
	// Arrange
	expectedCode := http.StatusBadRequest
	model.SetDataLayer(data.NewInMemoryLayer())
	// req := httptest.NewRequest("POST", "http://localhost/user", nil)
	// req.Form.Add("password", "password1")
	req := newPostUserRequest(postUserInput{0, "password1"})
	res := httptest.NewRecorder()

	// Act
	postUser(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d [userid = <empty>, password = \"password1\"]",
			expectedCode, res.Code)
	}
}

func Test_postUser_MissingPasswordField_ReturnsBadRequest(t *testing.T) {
	// Arrange
	expectedCode := http.StatusBadRequest
	model.SetDataLayer(data.NewInMemoryLayer())
	// req := httptest.NewRequest("POST", "http://localhost/user", nil)
	// req.Form.Add("userid", fmt.Sprintf("%d", 100))
	req := newPostUserRequest(postUserInput{101, ""})
	res := httptest.NewRecorder()

	// Act
	postUser(res, req)

	// Assert
	if res.Code != expectedCode {
		t.Errorf("expected status code = %d, received status code = %d [userid = 101, password = <empty>]",
			expectedCode, res.Code)
	}
}
