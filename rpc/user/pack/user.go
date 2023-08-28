package pack

import (
	// "douyin/rpc/user/dal/db"
	// "douyin/kitex_gen/user"
	"douyin/pkg/errno"
	"errors"

)

// // User pack user info
// func User(u *db.User) *user.User {
// 	if u == nil {
// 		return nil
// 	}

// 	return &user.User{UserId: int64(u.ID), Username: u.Username}
// }

// // Users pack list of user info
// func Users(us []*db.User) []*user.User {
// 	users := make([]*user.User, 0)
// 	for _, u := range us {
// 		if temp := User(u); temp != nil {
// 			users = append(users, temp)
// 		}
// 	}
// 	return users
// }



func GetErrorMesg(err error) string{
	if err == nil {
		return "Success"
	}
	e := errno.ErrNo{}
	if errors.As(err, &e) {
		return e.ErrMsg
	}

	s := errno.ServiceErr.WithMessage(err.Error())
	return s.ErrMsg

}