package handler

import (
	"context"
	"fmt"
	"douyin/kitex_gen/user"
	"douyin/kitex_gen/user/userservice"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"douyin/middleware"
)

type FeedImpl struct {
	client userservice.Client
}

type feedRequest struct {
	LatestTime int64  `json:"latest_time,omitempty"` // 可选参数，限制返回视频的最新投稿时间戳，精确到秒，不填表示当前时间
	Token      string `json:"token,omitempty"`       // 用户登录状态下设置
}
type feedResponse struct {
	StatusCode int64 `json:"status_code"` // 状态码，0-成功，其他值-失败
	// StatusMsg  string `json:"status_msg"` // 返回状态描述
	NextTime  int64            `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []Video_feedResp `json:"video_list"` // 视频列表
}

// Video
type Video_feedResp struct {
	Author        Author_feedResp `json:"author"`         // 视频作者信息
	CommentCount  int64           `json:"comment_count"`  // 视频的评论总数
	CoverURL      string          `json:"cover_url"`      // 视频封面地址
	FavoriteCount int64           `json:"favorite_count"` // 视频的点赞总数
	ID            int64           `json:"id"`             // 视频唯一标识
	IsFavorite    bool            `json:"is_favorite"`    // true-已点赞，false-未点赞
	PlayURL       string          `json:"play_url"`       // 视频播放地址
	Title         string          `json:"title"`          // 视频标题
}

// 视频作者信息
//
// User
type Author_feedResp struct {
	Avatar          string `json:"avatar"`           // 用户头像
	BackgroundImage string `json:"background_image"` // 用户个人页顶部大图
	FavoriteCount   int64  `json:"favorite_count"`   // 喜欢数
	FollowCount     int64  `json:"follow_count"`     // 关注总数
	FollowerCount   int64  `json:"follower_count"`   // 粉丝总数
	ID              int64  `json:"id"`               // 用户id
	IsFollow        bool   `json:"is_follow"`        // true-已关注，false-未关注
	Name            string `json:"name"`             // 用户名称
	Signature       string `json:"signature"`        // 个人简介
	TotalFavorited  int64  `json:"total_favorited"`  // 获赞数量
	WorkCount       int64  `json:"work_count"`       // 作品数
}

func NewFeedImpl() *FeedImpl {
	c, err := userservice.NewClient("user", client.WithHostPorts("127.0.0.1:9991"))
	if err != nil {
		panic(fmt.Sprintf("create user client error: %v", err))
	}
	return &FeedImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
}



// GetFeed: Get请求
func (u *FeedImpl) GetFeed(ctx context.Context, c *app.RequestContext) {
	username, password := c.PostForm("username"), c.PostForm("password")
	fmt.Println(username)
	fmt.Println(password)
	lr, err := u.client.Register(ctx, &user.DouyinUserRegisterRequest{Username: username, Password: password})
	if err != nil {
		response := UserRegisterResponse{
			StatusMsg:  "fail",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if lr.GetStatusCode() == -1 {
		response := UserRegisterResponse{
			StatusMsg:  lr.GetStatusMsg(),
		}
		
		c.JSON(http.StatusInternalServerError,response )
		return
	}

	token := middleware.GenerateJWTToken(lr.GetUserId())

	response := UserRegisterResponse{
		StatusCode: lr.GetStatusCode(), // 成功状态码
		StatusMsg:  lr.GetStatusMsg(),
		UserID:     lr.GetUserId(),
		Token:      token,
	}

	c.JSON(http.StatusOK, response)
}


