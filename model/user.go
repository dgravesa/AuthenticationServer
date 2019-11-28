package model

import (
	"fmt"
	"net/url"
	"strconv"
)

// UserLogin is a struct containing a user ID and password pair.
type UserLogin struct {
	ID       uint64
	Password string
}

// ParseUserLogin extracts a UserLogin from http request form values.
func ParseUserLogin(v *url.Values) (UserLogin, error) {
	var u UserLogin
	var err error

	// also errors in case of nil url values
	if u.ID, err = ParseUID(v); err != nil {
		return UserLogin{}, err
	}

	if u.Password = v.Get("password"); u.Password == "" {
		return UserLogin{}, fmt.Errorf("No password given")
	}

	return u, nil
}

// ParseUID extracts a user ID from http request form values.
func ParseUID(v *url.Values) (uint64, error) {
	var uid uint64
	var err error

	if v == nil {
		return 0, fmt.Errorf("No form data given")
	}

	if uid, err = strconv.ParseUint(v.Get("userid"), 10, 64); err != nil {
		return 0, err
	}

	return uid, nil
}