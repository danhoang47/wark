package usermodels

import (
	"errors"
	"wark/common"
)

type User struct {
	common.SQLModel
	Username string `json:"username" db:"username"`
	Salt     string `json:"-" db:"salt"`
	Password string `json:"-" db:"password"`
}

type CreateUser struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
	Salt     string `json:"-" db:"salt"`
}

var (
	ErrCreateUserPayloadInvalid = errors.New("createUser: payload invalid")
	ErrUsernameInvalid          = errors.New("createUser: username invalid")
	ErrPasswordInvalid          = errors.New("createUser: password invalid")
)

func (user *CreateUser) Validate() error {
	if user.Username == "" || user.Password == "" {
		return ErrCreateUserPayloadInvalid
	}

	if len(user.Username) > 100 {
		return ErrUsernameInvalid
	}

	if len(user.Password) < 6 {
		return ErrPasswordInvalid
	}

	return nil
}
