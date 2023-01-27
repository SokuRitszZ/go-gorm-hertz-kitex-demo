package main

import (
	"ghkd/kitex_gen/note/noteservice"
	"ghkd/module/note/model"
	"ghkd/module/note/rpc"
	"ghkd/pkg/consts"
	"ghkd/pkg/mw"
	"net"

	"github.com/cloudwego/kitex/pkg/klog"
	"github.com/cloudwego/kitex/pkg/limit"
	"github.com/cloudwego/kitex/pkg/rpcinfo"
	"github.com/cloudwego/kitex/server"
	kitexlogrus "github.com/kitex-contrib/obs-opentelemetry/logging/logrus"
	"github.com/kitex-contrib/obs-opentelemetry/provider"
	"github.com/kitex-contrib/obs-opentelemetry/tracing"
	etcd "github.com/kitex-contrib/registry-etcd"
)

func main() {
	r, err := etcd.NewEtcdRegistry([]string{consts.ETCDAddress})
	if err != nil {
		panic(err)
	}

	addr, err := net.ResolveTCPAddr(consts.TCP, consts.NoteServiceAddr)
	if err != nil {
		panic(err)
	}

	Init()

	provider.NewOpenTelemetryProvider(
		provider.WithServiceName(consts.NoteServiceName),
		provider.WithExportEndpoint(consts.ExportEndpoint),
		provider.WithInsecure(),
	)

	svr := noteservice.NewServer(
		new(NoteServiceImpl),
		server.WithServiceAddr(addr),
		server.WithRegistry(r),
		server.WithLimit(&limit.Option{MaxConnections: 1000, MaxQPS: 100}),
		server.WithMuxTransport(),
		server.WithMiddleware(mw.CommonMiddleware), // 输出请求方式 + 请求服务名
		server.WithMiddleware(mw.ServerMiddleware), // 输出请求来源的 IP 地址
		server.WithSuite(tracing.NewServerSuite()),
		server.WithServerBasicInfo(&rpcinfo.EndpointBasicInfo{ServiceName: consts.NoteServiceName}),
	)
	err = svr.Run()
	if err != nil {
		klog.Fatal(err)
	}
}

func Init() {
	rpc.Init()
	model.Init()

	// klog init
	klog.SetLogger(kitexlogrus.NewLogger())
	klog.SetLevel(klog.LevelInfo)	
}
