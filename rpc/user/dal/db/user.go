package db

import (
	"context"
	"douyin/pkg/config"
	// "fmt"
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID              int64      `gorm:"column:id;primary_key"`
	Name            string     `gorm:"column:name"`
	FollowCount     int64      `gorm:"column:follow_count"`
	FollowerCount   int64      `gorm:"column:follower_count"`
	IsFollow        bool       `gorm:"column:is_follow"`
	Avatar          string     `gorm:"column:avatar"`
	BackgroundImage string     `gorm:"column:background_image"`
	Signature       string     `gorm:"column:signature"`
	TotalFavorited  int64      `gorm:"column:total_favorited"`
	WorkCount       int64      `gorm:"column:work_count"`
	FavoriteCount   int64      `gorm:"column:favorite_count"`
	Password        string     `gorm:"column:password"`
	CreatedAt       time.Time  `gorm:"column:created_at"`
	UpdatedAt       time.Time  `gorm:"column:updated_at"`
	DeletedAt       *time.Time `gorm:"column:deleted_at;index"`
}

type UserDTO struct {
	ID              int64  `json:"id"`
	Name            string `json:"name"`
	FollowCount     int64  `json:"follow_count"`
	FollowerCount   int64  `json:"follower_count"`
	IsFollow        bool   `json:"is_follow"`
	Avatar          string `json:"avatar"`
	BackgroundImage string `json:"background_image"`
	Signature       string `json:"signature"`
	TotalFavorited  int64  `json:"total_favorited"`
	WorkCount       int64  `json:"work_count"`
	FavoriteCount   int64  `json:"favorite_count"`
}

func (u *User) TableName() string {
	return config.UserTableName
}

// MGetUsers multiple get list of user info
func MGetUsers(ctx context.Context, userIDs []int64) ([]*User, error) {
	res := make([]*User, 0)
	if len(userIDs) == 0 {
		return res, nil
	}

	if err := DB.Table(config.UserTableName).Where("id in ?", userIDs).Find(&res).Error; err != nil {
		return nil, err
	}
	return res, nil
}

// CreateUser create user info
func CreateUser(ctx context.Context, users []*User) ( []*User ,error) {
	return users,DB.Table(config.UserTableName).Create(users).Error
}

// QueryUser query list of user info
func QueryUsers(ctx context.Context, userName string) ([]*User, error) {
	res := make([]*User, 0)

	if err := DB.Table(config.UserTableName).Where("name = ?", userName).Find(&res).Error; err != nil {
		return nil, err
	}
	
	return res, nil
}

// QueryUser query list of user info
func QueryUser(ctx context.Context, userName string) (*User, error) {
	var res *User

	if err := DB.Table(config.UserTableName).Where("name = ?", userName).First(&res).Error; err != nil {
		return nil, err
	}
	
	return res, nil
}