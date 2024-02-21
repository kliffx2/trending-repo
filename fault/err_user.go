package fault

import "errors"

var (
	UserConflict = errors.New("User already exists")
	UserNotFound = errors.New("User not found")
	SignUpFail = errors.New("Sign up failed")
	UserNotUpdated = errors.New("Update user info failed")
) 