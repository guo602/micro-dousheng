package config

import (
	"github.com/spf13/viper"
	"fmt"
)

// AppConfig stores the application configuration
type AppConfig struct {
	DatabaseName             string
	DBHost                   string
	DBPort                   string
	DBUser                   string
	DBPassword               string
	JWTSecretKey             string
	SrcSavedPath             string
	AliyunOSSAccessKeyID     string // 阿里云 OSS Access Key ID
	AliyunOSSAccessKeySecret string // 阿里云 OSS Access Key Secret
	AliyunOSSEndpoint        string // 阿里云 OSS Endpoint
	AliyunOSSBucketName      string // 阿里云 OSS Bucket 名称
}

type UserInfoConfig struct {
	DefaultAvatarURL           string
	DefaultBackgroundImageURL  string
	DefaultSignature           string
}

// AppConfigInstance holds the instance of the application configuration
var AppConfigInstance AppConfig
var UserInfoConfigInstance UserInfoConfig
var ServiceConfigInstance ServiceConfig

func init(){
	LoadAppConfig()
	LoadUserInfoConfig()
	InitServiceConfig()
}

func LoadAppConfig() {
	// 初始化 Viper
	viper.SetConfigFile("config/config.json") // 配置文件的路径

	// 读取配置文件
	if err := viper.ReadInConfig(); err != nil {
		panic(err)
	}

	// 将配置绑定到结构体
	if err := viper.Unmarshal(&AppConfigInstance); err != nil {
		panic(err)
	}
}

func LoadUserInfoConfig(){
	UserInfoConfigInstance.DefaultAvatarURL = fmt.Sprintf("https://%s.%s/%s", AppConfigInstance.AliyunOSSBucketName, AppConfigInstance.AliyunOSSEndpoint, "img/default_avatar.png")
	UserInfoConfigInstance.DefaultBackgroundImageURL= fmt.Sprintf("https://%s.%s/%s", AppConfigInstance.AliyunOSSBucketName, AppConfigInstance.AliyunOSSEndpoint, "img/default_background.jpg")
	UserInfoConfigInstance.DefaultSignature = "这个人很懒，什么都没有留下"
}

func InitServiceConfig(){
	ServiceConfigInstance.EtcdAddress = "127.0.0.1:2379"

	ServiceConfigInstance.GatewayService = Service{
		Name : "Api_gateway",
		Address : "127.0.0.1:8888",
	}

	ServiceConfigInstance.FeedService = Service{
		Name : "FeedService",
		Address : "127.0.0.1:9991",
	}

	ServiceConfigInstance.UserService = Service{
		Name : "UserService",
		Address : "127.0.0.1:9992",
	}

	ServiceConfigInstance.FavoriteService = Service{
		Name : "FavoriteService",
		Address : "127.0.0.1:9993",
	}

	ServiceConfigInstance.CommentService = Service{
		Name : "CommentService",
		Address : "127.0.0.1:9994",
	}

	ServiceConfigInstance.PublishService = Service{
		Name : "PublishService",
		Address : "127.0.0.1:9995",
	}
}


type Service struct {
	Address string 
	Name string    
}


type ServiceConfig struct {
	GatewayService Service
	UserService Service
	FeedService Service
	FavoriteService Service
	PublishService Service
	CommentService Service
	EtcdAddress string
}