package server

import (
	"github.com/ouqiang/gocron/internal/modules/logger"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/ouqiang/gocron/internal/modules/rpc/auth"
	pb "github.com/ouqiang/gocron/internal/modules/rpc/proto"
	"github.com/ouqiang/gocron/internal/modules/utils"
	log "github.com/sirupsen/logrus"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/keepalive"
)

type Server struct{}
type ProcessServer struct{}

var keepAlivePolicy = keepalive.EnforcementPolicy{
	MinTime:             10 * time.Second,
	PermitWithoutStream: true,
}

var keepAliveParams = keepalive.ServerParameters{
	MaxConnectionIdle: 30 * time.Second,
	Time:              30 * time.Second,
	Timeout:           3 * time.Second,
}

// Run 实现rpc Task Run方法
func (s Server) Run(ctx context.Context, req *pb.TaskRequest) (*pb.TaskResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			log.Error(err)
		}
	}()
	log.Infof("execute cmd start: [id: %d cmd: %s]", req.Id, req.Command)
	output, err := utils.ExecShell(ctx, req.Command)
	resp := new(pb.TaskResponse)
	resp.Output = output
	if err != nil {
		resp.Error = err.Error()
	} else {
		resp.Error = ""
	}
	log.Infof("execute cmd end: [id: %d cmd: %s err: %s]", req.Id, req.Command, resp.Error)

	return resp, nil
}

// StartWorker 开启进程
func (s ProcessServer) StartWorker(ctx context.Context, req *pb.StartRequest) (*pb.StartResponse, error) {
	pid, err := utils.StartWorker(ctx, req)
	return &pb.StartResponse{Pid: int64(pid)}, err
}

func (s ProcessServer) StopWorker(_ context.Context, req *pb.StopRequest) (*pb.StopResponse, error) {
	err := utils.StopWorker(int(req.Pid))
	if err != nil {
		return &pb.StopResponse{Code: "fail", Message: err.Error()}, err
	}
	return &pb.StopResponse{Code: "success", Message: "Success"}, nil
}

func (s ProcessServer) WorkerStateCheck(_ context.Context, req *pb.StateRequest) (*pb.StateResponse, error) {
	state, _ := utils.WorkerStateCheck(int(req.Pid))
	return &pb.StateResponse{Code: "success", Message: "Success", State: state}, nil
}

func Start(addr string, enableTLS bool, certificate auth.Certificate) {
	defer func() {
		if err := recover(); err != nil {
			logger.Error("服务崩溃启动异常:", err)
		}
	}()
	l, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	opts := []grpc.ServerOption{
		grpc.KeepaliveParams(keepAliveParams),
		grpc.KeepaliveEnforcementPolicy(keepAlivePolicy),
	}
	if enableTLS {
		tlsConfig, err := certificate.GetTLSConfigForServer()
		if err != nil {
			log.Fatal(err)
		}
		opt := grpc.Creds(credentials.NewTLS(tlsConfig))
		opts = append(opts, opt)
	}
	server := grpc.NewServer(opts...)
	pb.RegisterTaskServer(server, Server{})
	pb.RegisterProcessServer(server, ProcessServer{})
	log.Infof("server listen on %s", addr)

	go func() {
		err = server.Serve(l)
		if err != nil {
			log.Fatal(err)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM)
	for {
		s := <-c
		log.Infoln("收到信号 -- ", s)
		switch s {
		case syscall.SIGHUP:
			log.Infoln("收到终端断开信号, 忽略")
		case syscall.SIGINT, syscall.SIGTERM:
			log.Info("应用准备退出")
			server.GracefulStop()
			return
		}
	}

}
