package fault

import "errors"

var (
	UserConflict = errors.New("user already exists")
	SignUpFail = errors.New("sign up failed")
)