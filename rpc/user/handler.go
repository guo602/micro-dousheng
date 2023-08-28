package main

import (
	"context"
	"fmt"
	user "douyin/kitex_gen/user"
	"douyin/rpc/user/pack"
	"douyin/rpc/user/service"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserRegisterResponse) //resp = &user.DouyinUserRegisterResponse{}

	if request.GetUsername() == "" || request.GetPassword() == "" {
		resp = &user.DouyinUserRegisterResponse{StatusCode: -1, StatusMsg: "username or password is empty"}
		return
	}

	users,er := service.NewCreateUserService(ctx).CreateUser(request)

	if er != nil{
		fmt.Println(pack.GetErrorMesg(err))
		resp = &user.DouyinUserRegisterResponse{
					StatusCode: -1, 
					StatusMsg: pack.GetErrorMesg(er),
				}
		return
	}

	resp = &user.DouyinUserRegisterResponse{
					StatusCode: 0, 
					StatusMsg: "注册用户成功",
					UserId: users[0].ID ,
					Token: "DeFault_Token_Wait_For_Imple"}
	return
}





// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	resp = new(user.DouyinUserLoginResponse)  //resp = &user.DouyinUserLoginResponse{}
	
	ids,er := service.NewCheckUserService(ctx).CheckUser(request)
	
	if er != nil {
		resp = &user.DouyinUserLoginResponse{
						StatusCode: -1, 
						StatusMsg: pack.GetErrorMesg(er)}
		return 
	}
	

	resp = &user.DouyinUserLoginResponse{
						StatusCode: 0, 
						UserId: ids ,
						StatusMsg: pack.GetErrorMesg(er),
					}

	return
}



// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, request *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}




// // 根据用户名查询用户信息
// func getUserByUsername(username string) (models.User, error) {
// 	var user_re models.User
// 	if err := database.DB.Table("user").Where("name = ?", username).First(&user_re).Error; err != nil {
// 		return models.User{}, err
// 	}
// 	return user_re, nil
// }