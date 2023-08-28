package handler

import (
	"context"
	"fmt"
	"douyin/kitex_gen/feed"
	"douyin/kitex_gen/feed/feedservice"
	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/kitex/client"
	"douyin/pkg/config"
	etcd "github.com/kitex-contrib/registry-etcd"
	"time"
	"github.com/cloudwego/kitex/pkg/retry"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/hertz/pkg/protocol/consts"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
)

type FeedImpl struct {
	client feedservice.Client
}

var FeedImplInst FeedImpl


func init()  {
	// c, err := feedservice.NewClient("feed", client.WithHostPorts("127.0.0.1:9991"))
	// if err != nil {
	// 	panic(fmt.Sprintf("create user client error: %v", err))
	// }
	r, err := etcd.NewEtcdResolver([]string{config.ServiceConfigInstance.EtcdAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.FeedServiceName),
		provider.WithExportEndpoint(config.ExportEndpoint),
		provider.WithInsecure(),
	)
	
	ServiceName := config.ServiceConfigInstance.FeedService.Name

	c, err := feedservice.NewClient(
		ServiceName,
		client.WithMuxConnection(1),                       // mux
		client.WithRPCTimeout(30*time.Second),             // rpc timeout
		client.WithConnectTimeout(30000*time.Millisecond), // conn timeout
		client.WithFailureRetry(retry.NewFailurePolicy()), // retry
		client.WithResolver(r),                            // resolver
		client.WithSuite(tracing.NewClientSuite()),
		// Please keep the same as provider.WithServiceName
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: ServiceName}),
	)
	if err != nil {
		panic(fmt.Sprintf("create feed client error: %v", err))
	}

	FeedImplInst = FeedImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
}



// GetFeed: Get请求
func (u *FeedImpl) GetFeed(ctx context.Context, c *app.RequestContext) {

	var err error
	var req feedRequest
	err = c.BindAndValidate(&req)  //  token, _ := c.Query("token"), c.Query("last_time")
	if err != nil {
		SendResponse(c,consts.StatusInternalServerError , utils.H{
			"StatusMsg": "bind and validate error",
		})
		return
	}

	
	rpc_resp, err := u.client.ListVideos(ctx, &feed.DouyinFeedRequest{Token: req.Token, LatestTime: req.LatestTime})
	if err != nil || rpc_resp.GetStatusCode() == -1 {
		err_response := feedResponse{
			StatusCode:  -1,
		}
		SendResponse(c,consts.StatusInternalServerError, err_response)
		return
	}

	
	succ_response := feedResponse{
		StatusCode: int64(rpc_resp.GetStatusCode()), // 成功状态码
		VideoList: rpc_resp.GetVideoList(),
	}

	SendResponse(c,consts.StatusOK, succ_response)
}





// func NewFeedImpl() *FeedImpl {
// 	c, err := feedservice.NewClient("feed", client.WithHostPorts("127.0.0.1:9991"))
// 	if err != nil {
// 		panic(fmt.Sprintf("create user client error: %v", err))
// 	}
// 	return &FeedImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
// }

