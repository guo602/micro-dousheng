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

type UserImpl struct {
	client userservice.Client
}

var UserImplInst UserImpl

func init()  {
	// c, err := userservice.NewClient("user", client.WithHostPorts("127.0.0.1:9990"))
	// if err != nil {
	// 	panic(fmt.Sprintf("create user client error: %v", err))
	// }

	r, err := etcd.NewEtcdResolver([]string{config.ServiceConfigInstance.EtcdAddress})
	if err != nil {
		panic(err)
	}
	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(config.UserServiceName),
		provider.WithExportEndpoint(config.ExportEndpoint),
		provider.WithInsecure(),
	)
	

	ServiceName := config.ServiceConfigInstance.UserService.Name

	c, err := userservice.NewClient(
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
		panic(fmt.Sprintf("create user client error: %v", err))
	}

	UserImplInst = UserImpl{client: c} 

	// return &UserImpl{client: c} // 指定下游的ip，高级用法可以使用resolver去调用服务注册中心
}



// Login: Post请求
func (u *UserImpl) Register(ctx context.Context, c *app.RequestContext) {
	var err error
	var req UserRegisterRequest
	err = c.BindAndValidate(&req)
	if err != nil {
		SendResponse(c,consts.StatusInternalServerError , utils.H{
			"StatusMsg": "bind and validate error",
		})
		return
	}

	rpc_resp, err := u.client.Register(ctx, &user.DouyinUserRegisterRequest{Username: req.Username, Password: req.Password})
	if err != nil {
		err_response := UserRegisterResponse{
			StatusMsg:  "Register fail",
		}
		SendResponse(c,http.StatusInternalServerError, err_response)    // c.JSON(http.StatusInternalServerError, response)
		return
	}

	if rpc_resp.GetStatusCode() == -1 {
		err_response := UserRegisterResponse{
			StatusCode: rpc_resp.GetStatusCode(),
			StatusMsg:  rpc_resp.GetStatusMsg(),
		}
		SendResponse(c,consts.StatusInternalServerError , err_response)
		return
	}

	token := middleware.GenerateJWTToken(rpc_resp.GetUserId())

	success_response := UserRegisterResponse{
		StatusCode: rpc_resp.GetStatusCode(), // 成功状态码
		StatusMsg:  rpc_resp.GetStatusMsg(),
		UserID:     rpc_resp.GetUserId(),
		Token:      token,
	}

	SendResponse(c,consts.StatusOK , success_response)

}
// LogIn: Post请求
func (u *UserImpl) LogIn(ctx context.Context, c *app.RequestContext) {
	username, password := c.PostForm("username"), c.PostForm("password")
	lr, err := u.client.Login(ctx, &user.DouyinUserLoginRequest{Username: username, Password: password})
	if err != nil {

		response := UserLoginResponse{
			StatusMsg:  "fail",
		}
		c.JSON(http.StatusInternalServerError, response)
		return
	}

	if lr.GetStatusCode() == -1 {
		response := UserLoginResponse{
			StatusMsg:  lr.GetStatusMsg(),
		}
		c.JSON(http.StatusInternalServerError,response )
		return
	}

	token := middleware.GenerateJWTToken(lr.GetUserId())
	response := UserLoginResponse{
		StatusCode: lr.GetStatusCode(), // 成功状态码
		StatusMsg:  lr.GetStatusMsg(),
		UserID:     lr.GetUserId(),
		Token:      token,
	}
	c.JSON(consts.StatusOK, response)
	
}

// GetUser: Get请求
func (u *UserImpl) GetUserById(ctx context.Context, c *app.RequestContext) {
	//to be done
}


