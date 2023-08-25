// Code generated by hertz generator.

package main

import (
	handler "douyin/api_gateway/biz/handler"
	"github.com/cloudwego/hertz/pkg/app/server"
	"douyin/middleware"
)

// customizeRegister registers customize routers.
func customizedRegister(r *server.Hertz) {
	r.GET("/ping", middleware.JWTMiddleware(),handler.Ping)

	// your code ...

	
	rg := r.Group("/douyin")
	{
		userGroup := rg.Group("/user")
		{
			ui := handler.NewUserImpl()
			userGroup.POST("/register/",  ui.Register)
		}

	}

	
}
