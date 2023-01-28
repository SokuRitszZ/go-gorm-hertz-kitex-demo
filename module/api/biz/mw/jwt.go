package mw

import (
	"context"
	"ghkd/kitex_gen/user"
	"ghkd/module/api/biz/model/api"
	"ghkd/module/api/biz/rpc"
	"ghkd/pkg/consts"
	"ghkd/pkg/errno"
	"net/http"
	"time"

	"github.com/cloudwego/hertz/pkg/app"
	"github.com/cloudwego/hertz/pkg/common/utils"
	"github.com/hertz-contrib/jwt"
)

var JwtMiddleware *jwt.HertzJWTMiddleware

func InitJWT() {
	JwtMiddleware, _ = jwt.New(&jwt.HertzJWTMiddleware{
		Key: []byte(consts.SecretKey),
		TokenLookup: "header: Authorization, query: token, cookie: jwt",
		TokenHeadName: "Bearer",
		TimeFunc: time.Now,
		Timeout: time.Hour,
		MaxRefresh: time.Hour,
		IdentityKey: consts.IdentityKey,
		// 将信息存入 JWT
		IdentityHandler: func(ctx context.Context, c *app.RequestContext) interface{} {
			claims := jwt.ExtractClaims(ctx, c)
			return &api.User{
				ID: int64(claims[consts.IdentityKey].(float64)),
			}
		},
		// 荷载（对 JWT 进行验证并获取信息）
		PayloadFunc: func(data interface{}) jwt.MapClaims {
			if v, ok := data.(int64); ok {
				return jwt.MapClaims{
					// IdentityKey 就是 "id"
					consts.IdentityKey: v,
				}
			}
			return jwt.MapClaims{}
		},
		// Authenticator 验证器
		Authenticator: func(ctx context.Context, c *app.RequestContext) (interface{}, error) {
			var err error 
			var req api.CheckUserRequest
			if err = c.BindAndValidate(&req); err != nil {
				return "", jwt.ErrMissingLoginValues
			}
			if len(req.Name) == 0 || len(req.Password) == 0 {
				return "", jwt.ErrMissingLoginValues
			}
			return rpc.CheckUser(context.Background(), &user.CheckUserRequest{
				Name: req.Name,
				Password: req.Password,
			})
		},
		// 登陆响应
		LoginResponse: func(ctx context.Context, c *app.RequestContext, code int, token string, expire time.Time) {
			c.JSON(http.StatusOK, utils.H{
				"code": errno.Success.ErrCode,
				"token": token,
				"expire": expire.Format(time.RFC3339),
			})
		},
		// 认证失败
		Unauthorized: func(ctx context.Context, c *app.RequestContext, code int, message string) {
			c.JSON(http.StatusOK, utils.H{
				"code": errno.AuthorizationErr,
				"message": message,
			})
		},
		// HTTP状态信息显示
		HTTPStatusMessageFunc: func(e error, ctx context.Context, c *app.RequestContext) string {
			switch t := e.(type) {
			case errno.ErrNo:
				return t.ErrMsg
			default:
				return t.Error()
			}
		},
	})
}
