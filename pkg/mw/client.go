package mw

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = ClientMiddleware 

// ClientMiddleware 每次显示目的地址、配置 RPCTimeout、ConnectTimeout
func ClientMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)

		// get server information
		klog.Infof(
			"server address: %v, rpc timeout: %v, readwrite timeout: %v\n", 
			ri.To().Address(), 
			ri.Config().RPCTimeout(),
			ri.Config().ConnectTimeout(),
		)
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		return nil
	}
}
