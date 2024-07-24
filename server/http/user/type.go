package user

import (
	"errors"

	utilString "github.com/azizyogo/bank-app/common/util/string"
)

type (
	LoginRequest struct {
		Username string `json:"username"`
		Password string `json:"password"`
	}
)

func (s LoginRequest) Validate() error {
	if s.Username == "" {
		return errors.New("username must not empty")
	}

	if !utilString.IsValidString(s.Username) {
		return errors.New("invalid username, please only use characters")
	}

	if s.Password == "" {
		return errors.New("password must not empty")
	}

	if !utilString.IsValidString(s.Password) {
		return errors.New("invalid password, please only use characters")
	}

	return nil
}
