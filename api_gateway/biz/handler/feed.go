package handler

import (
	"context"
	"fmt"
	"douyin/kitex_gen/feed"
	"douyin/kitex_gen/feed/feedservice"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"douyin/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

type FeedImpl struct {
	client feedservice.Client
}


type feedResponse struct {
	StatusCode int64 `json:"status_code"` // 状态码，0-成功，其他值-失败
	// StatusMsg  string `json:"status_msg"` // 返回状态描述
	NextTime  int64            `json:"next_time"`  // 本次返回的视频中，发布最早的时间，作为下次请求时的latest_time
	VideoList []*feed.Video `json:"video_list"` // 视频列表
}



func NewFeedImpl() *FeedImpl {
	// c, err := feedservice.NewClient("feed", client.WithHostPorts("127.0.0.1:9991"))
	// if err != nil {
	// 	panic(fmt.Sprintf("create user client error: %v", err))
	// }
	r, err := etcd.NewEtcdResolver([]string{config.ServiceConfigInstance.EtcdAddress})
	if err != nil {
		panic(err)
	}
	ServiceName := config.ServiceConfigInstance.FeedService.Name

	c, err := feedservice.NewClient(
		ServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(fmt.Sprintf("create feed client error: %v", err))
	}

	return &FeedImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
}



// GetFeed: Get请求
func (u *FeedImpl) GetFeed(ctx context.Context, c *app.RequestContext) {
	token, _ := c.Query("token"), c.Query("last_time")
	
	lr, err := u.client.ListVideos(ctx, &feed.DouyinFeedRequest{Token: token, LatestTime: 0})
	if err != nil {
		response := feedResponse{
			StatusCode:  -1,
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}
	if lr.GetStatusCode() == -1 {
		response := feedResponse{
			StatusCode:  -1,
		}
		
		c.JSON(http.StatusInternalServerError,response )
		return
	}

	

	response := feedResponse{
		StatusCode: int64(lr.GetStatusCode()), // 成功状态码
		VideoList: lr.GetVideoList(),
	}

	c.JSON(http.StatusOK, response)
}


