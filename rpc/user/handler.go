package main

import (
	"context"
	"douyin/database"
	"douyin/database/models"
	user "douyin/kitex_gen/user"
	"douyin/config"
	"golang.org/x/crypto/bcrypt"
)

// UserServiceImpl implements the last service interface defined in the IDL.
type UserServiceImpl struct{}

// Register implements the UserServiceImpl interface.
func (s *UserServiceImpl) Register(ctx context.Context, request *user.DouyinUserRegisterRequest) (resp *user.DouyinUserRegisterResponse, err error) {
	// TODO: Your code here...
	resp = &user.DouyinUserRegisterResponse{}

	if request.GetUsername() == "" || request.GetPassword() == "" {
		resp = &user.DouyinUserRegisterResponse{StatusCode: -1, StatusMsg: "username or password is empty"}
		return
	}

	hashPassword, err := hashPassword(request.GetPassword())
	if err != nil {
		resp = &user.DouyinUserRegisterResponse{StatusCode: -1, StatusMsg: "密码加密失败"}
		return
	}

	// 创建用户数据模型
	newuser := models.User{
		Name:     request.GetUsername(),
		Password: hashPassword,
		Avatar : config.UserInfoConfigInstance.DefaultAvatarURL,
		BackgroundImage : config.UserInfoConfigInstance.DefaultBackgroundImageURL,
		Signature : config.UserInfoConfigInstance.DefaultSignature,
	}

	// 验证用户名是否已经存在
	err = database.DB.Table("user").Where("name = ?", newuser.Name).First(&newuser).Error
	if err == nil {
		resp = &user.DouyinUserRegisterResponse{StatusCode: -1, StatusMsg: "注册用户名已经存在"}
		return
	}

	//保存用户数据到数据库
	err = database.DB.Table("user").Create(&newuser).Error
	if err != nil {
		resp = &user.DouyinUserRegisterResponse{StatusCode: -1, StatusMsg: "注册信息写入数据库时创建用户失败"}
		return
	}


	resp = &user.DouyinUserRegisterResponse{StatusCode: 0, StatusMsg: "注册用户成功",UserId: newuser.ID ,Token: "DeFault_Token_Wait_For_Inple"}
	return
}





// Login implements the UserServiceImpl interface.
func (s *UserServiceImpl) Login(ctx context.Context, request *user.DouyinUserLoginRequest) (resp *user.DouyinUserLoginResponse, err error) {
	// TODO: Your code here...
	// resp = &user.DouyinUserLoginResponse{}
	
	// 进行用户登录验证，比对用户名和密码是否正确
	login_user, er := getUserByUsername(request.GetUsername())
	if er != nil {
		resp = &user.DouyinUserLoginResponse{StatusCode: -1, StatusMsg: "用户名不存在"}
		return 
	}

	// 验证密码是否正确
	er = bcrypt.CompareHashAndPassword([]byte(login_user.Password), []byte(request.GetPassword()))
	if er != nil {
		resp = &user.DouyinUserLoginResponse{StatusCode: -1, StatusMsg: "用户名密码不正确"}
		return 
	}
	

	resp = &user.DouyinUserLoginResponse{StatusCode: 0, UserId: login_user.ID ,StatusMsg: "登录成功"}

	return
}



// GetUserById implements the UserServiceImpl interface.
func (s *UserServiceImpl) GetUserById(ctx context.Context, request *user.DouyinUserRequest) (resp *user.DouyinUserResponse, err error) {
	// TODO: Your code here...
	return
}


// 加密用户密码
func hashPassword(password string) (string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedBytes), nil
}

// 根据用户名查询用户信息
func getUserByUsername(username string) (models.User, error) {
	var user_re models.User
	if err := database.DB.Table("user").Where("name = ?", username).First(&user_re).Error; err != nil {
		return models.User{}, err
	}
	return user_re, nil
}