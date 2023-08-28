package service

import (
	"context"
	user "douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"douyin/rpc/user/dal/db"
	"golang.org/x/crypto/bcrypt"

)


type CheckPwService struct {
	ctx context.Context
}

// NewCheckUserService new CheckUserService
func NewCheckUserService(ctx context.Context) *CheckPwService {
	return &CheckPwService{
		ctx: ctx,
	}
}

// CheckUser check user info
func (s *CheckPwService) CheckUser(req *user.DouyinUserLoginRequest) (int64, error) {

	userName := req.Username
	users, err := db.QueryUsers(s.ctx, userName)
	if err != nil {
		return -1,  errno.UserDoNotExistErr
	}

	if len(users) == 0 {
		return -1, errno.UserDoNotExistErr
	}
	
	u := users[0]
	if err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(req.Password)); err != nil {
		return -1, errno.AuthorizationFailedErr
	}

	return int64(u.ID), nil
}