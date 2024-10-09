package usermodels

import "wark/common"

type User struct {
	common.SQLModel
	Username string `json:"username" db:"username"`
	Salt     string `json:"-" db:"salt"`
	Password string `json:"-" db:"password"`
}

type CreateUser struct {
	Username string `json:"username" db:"username"`
	Password string `json:"password" db:"password"`
}
