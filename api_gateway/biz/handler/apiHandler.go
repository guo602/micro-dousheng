package handler

import (
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"douyin/kitex_gen/feed"
)

type Response struct {
	Code    int64       `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type UserRegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// UserLoginRequest 是用户登录请求的结构体
type UserLoginRequest struct {
	Username string `form:"username" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// UserProfileRequest 是获取用户信息请求的结构体
type UserProfileRequest struct {
	UserID int64  `form:"user_id" binding:"required"`
	Token  string `form:"token" binding:"required"`
}

type UserRegisterResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}


// UserLoginResponse 是用户登录响应的结构体
type UserLoginResponse struct {
	StatusCode int32  `json:"status_code"`
	StatusMsg  string `json:"status_msg"`
	UserID     int64  `json:"user_id"`
	Token      string `json:"token"`
}

type feedRequest struct {
	LatestTime int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token,omitempty"`       // 用户登录状态下设置
}

type feedResponse struct {
	StatusCode int64 `json:"status_code"` // 状态码，0-成功，其他值-失败
	// StatusMsg  string `json:"status_msg"` // 返回状态描述
	NextTime  int64            `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*feed.Video `json:"video_list"` // 视频列表
}

// SendErrorResponse pack response
func SendErrorResponse(c *app.RequestContext, err error, data interface{}) {
	c.JSON(consts.StatusInternalServerError, Response{
		Data:    data,
	})
}


// SendErrorResponse pack response
func SendResponse(c *app.RequestContext, statusCode int, data interface{}) {
	c.JSON(statusCode,
		   data,
	)
}