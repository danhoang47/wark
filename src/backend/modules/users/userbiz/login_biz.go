package userbiz

import (
	"errors"
	"strings"
	"wark/components/hasher"
	"wark/modules/users/usermodels"
	"wark/modules/users/userrepos"
)

type GetUserRepository interface {
	GetUser(*userrepos.GetOneUserOptions) (*usermodels.User, error)
}

type LoginUserBusiness interface {
	Login(hasher.Hasher, *usermodels.CreateUser) (*usermodels.User, error)
}

type loginUserBiz struct {
	getUserRepo GetUserRepository
}

func NewLoginUserBiz(getUserRepo GetUserRepository) *loginUserBiz { return &loginUserBiz{getUserRepo} }

func (biz *loginUserBiz) Login(hasher hasher.Hasher, loginUser *usermodels.CreateUser) (*usermodels.User, error) {
	user, err := biz.getUserRepo.GetUser(&userrepos.GetOneUserOptions{
		Username: loginUser.Username,
	})

	if err != nil {
		return nil, err
	}

	hashedPassword := hasher.Hash([]byte(loginUser.Password + user.Salt))

	if strings.Compare(hashedPassword, user.Password) != 0 {
		return nil, errors.New("loginUserBiz: invalid password")
	}

	return user, nil
}
