package main

import (
	"context"
	"douyin/database"
	"douyin/database/models"
	feed "douyin/kitex_gen/feed"
	"douyin/middleware"
	"fmt"
	"log"
	"time"
	"gorm.io/gorm"

)

// FeedServiceImpl implements the last service interface defined in the IDL.
type FeedServiceImpl struct{}

// ListVideos implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) ListVideos(ctx context.Context, request *feed.DouyinFeedRequest) (resp *feed.DouyinFeedResponse, err error) {
	// TODO: Your code here...
	resp = &feed.DouyinFeedResponse{}
	
	token := request.GetToken()
	// lastTime := request.GetLatestTime()
	var uid int64
	if token == "" {
		uid = -1
	}else{
		uid , _ = middleware.ParsedJWTToken(token)
	}

	fmt.Println(uid)

	var video_ids []int64
	var videos []models.Video
	result := database.DB.Table("video").Select("id").Find(&videos)
	if result.Error != nil {
		log.Println(result.Error)
	}
	for _, video := range videos {
		video_ids = append(video_ids, video.VideoID)
	}
	

	//新发布的先刷到，将vid倒叙排列
	video_ids = reverseList(video_ids)
	var Videos []*feed.Video
	for _, v_id := range video_ids {
		Videos = append(Videos, Get_Video_for_feed(v_id, uid))
	}

	

	resp = &feed.DouyinFeedResponse{  
		StatusCode: 0,                 // 成功状态码
		NextTime:   time.Now().Unix(), // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time

		// StatusMsg  : "feed get success" ,// 返回状态描述
		VideoList: Videos, // 视频列表
	}
	

	return
}

// QueryVideos implements the FeedServiceImpl interface.
func (s *FeedServiceImpl) QueryVideos(ctx context.Context, video *feed.Video) (resp *feed.VideoIdRequest, err error) {
	// TODO: Your code here...
	return
}

func Get_Video_for_feed(video_id int64, current_userID int64) *feed.Video {
	var video models.Video
	result := database.DB.Table("video").Where("id = ?", video_id).Find(&video)
	if result.Error != nil {
		log.Fatal(result.Error)
	}
	
	autherID := video.AuthorUserID
	author_resp := Get_author_for_feed(autherID, current_userID)
	
	var far models.Favorite
	var isfar bool
	result2 := database.DB.Table("favorite").Where("user_id = ? AND video_id = ? AND is_deleted=-1", current_userID, video_id).First(&far)
	// if result2.Error != nil && result2.Error != gorm.ErrRecordNotFound {
	// 	log.Fatal(result2.Error)
	// }

	if result2.RowsAffected > 0 {
		isfar = true
	} else {
		isfar = false
	}

	var video_resp = feed.Video{
		Id:            video_id,
		Author:        &author_resp,
		PlayUrl:       video.PlayURL,
		CoverUrl:      video.CoverURL,
		FavoriteCount: int64(video.Likes),
		CommentCount:  int64(video.Comments),
		IsFavorite:    isfar,
		Title:         video.Title,
	}

	return &video_resp
}

func Get_author_for_feed(author_id int64, current_userID int64) feed.User {

	var author_resp feed.User
	var author models.User
	var relation models.Relation
	var follow bool

	

	result1 := database.DB.Table("user").Where("id = ?", author_id).First(&author)
	
	if result1.Error != nil {
		log.Println(result1.Error)
	}

	

	result2 := database.DB.Table("relation").Where("follower_id = ? AND followed_id = ?", current_userID, author_id).First(&relation)
	if result2.Error != nil && result2.Error != gorm.ErrRecordNotFound {
		log.Println(result2.Error)
	}

	if result2.RowsAffected > 0 {
		follow = true
	} else {
		follow = false
	}

	author_resp = feed.User{
		Id:              author_id,
		Name:            author.Name,
		BackgroundImage: author.BackgroundImage, // 用户个人页顶部大图
		FavoriteCount:   author.FavoriteCount,   // 喜欢数
		FollowCount:     author.FollowCount,     // 关注总数
		FollowerCount:   author.FollowerCount,   // 粉丝总数
		Signature:       author.Signature,       // 个人简介
		TotalFavorited:  author.TotalFavorited,  // 获赞数量
		WorkCount:       author.WorkCount,       // 作品数
		Avatar:          author.Avatar,
		IsFollow:        follow,
	}

	return author_resp
}

func reverseList(list []int64) []int64 {
	length := len(list)
	reversed := make([]int64, length)
	for i, j := 0, length-1; i < length; i, j = i+1, j-1 {
		reversed[j] = list[i]
	}
	return reversed
}

