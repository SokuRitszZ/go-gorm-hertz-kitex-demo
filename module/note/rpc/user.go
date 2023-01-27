package rpc

import (
	"context"
	"ghkd/kitex_gen/user"
	"ghkd/kitex_gen/user/userservice"
	"ghkd/pkg/consts"
	"ghkd/pkg/mw"

	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"

	// "github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

var userClient userservice.Client

func initUser() {
	r, err := etcd.NewEtcdResolver([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.NoteServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)
	c, err := userservice.NewClient(
		consts.UserServiceName,
		client.WithResolver(r),
		client.WithMuxConnection(1),
		client.WithMiddleware(mw.CommonMiddleware), // 输出请求方法 + 服务
		client.WithMiddleware(mw.ClientMiddleware), // 输出rpc 调用信息
		client.WithSuite(tracing.NewClientSuite()),
		client.WithClientBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}),
	)
	if err != nil {
		panic(err)
	}
	userClient = c
}

func MGetUser(ctx context.Context, req *user.MGetUserRequest) (map[int64]*user.User, error) {
	resp, err := userClient.MGetUser(ctx, req)
	if err != nil {
		return nil, err
	}
	if resp.BaseResp.StatusCode != 0 {
		return nil, err
	}
	res := make(map[int64]*user.User)
	for _, u := range resp.Users {
		res[u.ID] = u
	}
	return res, nil
}
