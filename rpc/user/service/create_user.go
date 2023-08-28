package service

import (
	"context"
	user "douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/rpc/user/dal/db"
	"golang.org/x/crypto/bcrypt"
	// "fmt"
	"douyin/pkg/config"
)

type CreateUserService struct {
	ctx context.Context
}

// NewCreateUserService new CreateUserService
func NewCreateUserService(ctx context.Context) *CreateUserService {
	return &CreateUserService{ctx: ctx}
}

// CreateUser create user info.
func (s *CreateUserService) CreateUser(req *user.DouyinUserRegisterRequest)( []*db.User ,error) {
	users, err := db.QueryUsers(s.ctx, req.Username)
	if err != nil {
		return nil,err
	}
	if len(users) != 0 {
		return nil,errno.UserAlreadyExistErr
	}

	hashPassword, err := hashPassword(req.GetPassword())
	if err != nil {
		return nil,errno.PasswdHashFailedErr
	}


	return db.CreateUser(s.ctx, []*db.User{{
		Name: req.Username,
		Password: hashPassword,
		Avatar : config.UserInfoConfigInstance.DefaultAvatarURL,
		BackgroundImage : config.UserInfoConfigInstance.DefaultBackgroundImageURL,
		Signature : config.UserInfoConfigInstance.DefaultSignature,
	}})
	// return nil
}


// 加密用户密码
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}