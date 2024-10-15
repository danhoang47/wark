package userbiz

import (
	"wark/common"
	"wark/components/hasher"
	"wark/modules/users/usermodels"
)

type RegisterUserRepository interface {
	Add(u *usermodels.CreateUser) (int, error)
}

type RegisterUserBussiness interface {
	Register(u *usermodels.CreateUser) error
}

type registerUserBiz struct {
	userRepo RegisterUserRepository
}

func NewRegisterUserBiz(userRepo RegisterUserRepository) *registerUserBiz {
	return &registerUserBiz{userRepo}
}

func (biz *registerUserBiz) Register(hasher hasher.Hasher, u *usermodels.CreateUser) (int, error) {
	if err := u.Validate(); err != nil {
		return 0, err
	}

	salt := common.GetSalt()
	hashedPassword := hasher.Hash([]byte(u.Password + salt))

	u.Password = hashedPassword
	u.Salt = salt

	return biz.userRepo.Add(u)
}
