package middleware

import (
	"context"
	"douyin/pkg/config"
	"fmt"
	"time"
	"github.com/dgrijalva/jwt-go"
	"net/http"
	"github.com/cloudwego/hertz/pkg/app"
)



func JWTMiddleware() app.HandlerFunc {
	return func(ctx context.Context, c *app.RequestContext) {
		// // 获取请求的路径
		// requestPath := c.

		// // 如果是登录或注册接口，则跳过验证
		// if requestPath == "/douyin/user/register/" || requestPath == "/douyin/user/login/" {
		// 	c.Next()
		// 	return
		// }

		// 需要验证token的接口：视频流接口、用户信息接口、投稿接口、发布列表、赞操作、喜欢列表、评论操作、评论列表

		// 需要验证user_id的接口：用户信息接口、发布列表、喜欢列表，自己在接口中单独额外验证user_id与token中的user_id是否一致

		// 获取请求体中的Token
		tokenString := c.Query("token")
		if tokenString == "" {
			tokenString = c.PostForm("token") // 从 POST form-data 中获取 token
		}

		if tokenString == "" {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "未提供Token"})
			c.Abort()
			return
		}

		// 解析Token
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return []byte(config.AppConfigInstance.JWTSecretKey), nil
		})

		if err != nil {
			c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "无效的Token"})
			fmt.Println("无效的Token")
			c.Abort()
			return
		}

		// 检查Token是否有效
		if token.Valid {
			// 检查Token是否已经过期
			if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
				c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token已过期"})
				c.Abort()
				return
			}

			// 提取用户标识
			if claims, ok := token.Claims.(jwt.MapClaims); ok {
				if userID, ok := claims["user_id"].(float64); ok {
					// 将用户ID存储在请求上下文中
					c.Set("user_id", int64(userID))
					c.Next(ctx)
					return
				}
			}
		}

		c.JSON(http.StatusUnauthorized, map[string]interface{}{"error": "Token验证失败"})
		c.Abort()

	}
}


// 生成 JWT Token
func GenerateJWTToken(userID int64) string {
	// 创建一个新的Token对象
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Token 过期时间为 1 天
	})

	// 使用密钥对 Token 进行签名
	tokenString, err := token.SignedString([]byte(config.AppConfigInstance.JWTSecretKey))
	if err != nil {
		// 处理错误
		return ""
	}

	return tokenString
}

// 解析 JWT Token
func ParsedJWTToken(tokenString string) (int64,error) {
	// 解析Token
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return []byte(config.AppConfigInstance.JWTSecretKey), nil
	})

	if err != nil {
		fmt.Println("无效的Token")
		return -1,err
	}

	// 检查Token是否有效
	if token.Valid {
		// 检查Token是否已经过期
		if _, ok := token.Claims.(jwt.Claims); !ok && !token.Valid {
			fmt.Println("Token已过期")
			return -1,err
		}
	}

	// 提取用户标识
	if claims, ok := token.Claims.(jwt.MapClaims); ok {
		if userID, ok := claims["user_id"].(float64); ok {
		
			return int64(userID),nil
		}
	}

	return 0,nil
}
