package fault

import "errors"

var (
	UserConflict = errors.New("user already exists")
	UserNotFound = errors.New("user not found")
	SignUpFail = errors.New("sign up failed")
	UserNotUpdated = errors.New("update user info failed")
) 