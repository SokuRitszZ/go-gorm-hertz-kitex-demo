package mw

import (
	"context"

	"github.com/cloudwego/kitex/pkg/endpoint"
	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
)

var _ endpoint.Middleware = CommonMiddleware

// CommonMiddleware 请求方式 + 请求服务名
func CommonMiddleware(next endpoint.Endpoint) endpoint.Endpoint {
	return func(ctx context.Context, req, resp interface{}) (err error) {
		ri := rpcinfo.GetRPCInfo(ctx)

		// get real request
		klog.Info("real request: %+v\n", req)
		klog.Infof(
			"remote service name: %s, remote method: %s\n",
			ri.To().ServiceName(),
			ri.To().Method(),
		)
		if err = next(ctx, req, resp); err != nil {
			return err
		}
		klog.Infof("real respone: %+v\n", resp)
		return nil
	}
}
