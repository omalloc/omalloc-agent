package sdk

import (
	"context"

	pb "github.com/omalloc/agent/api/executor"
)

type ExecutorControlSideServer struct {
	pb.UnimplementedExecutorServer

	executor *executor
}

func NewExecutorControlSideServer() *ExecutorControlSideServer {
	return &ExecutorControlSideServer{}
}

func (s *ExecutorControlSideServer) Beat(context.Context, *pb.BeatRequest) (*pb.BeatReply, error) {
	return &pb.BeatReply{}, nil
}
func (s *ExecutorControlSideServer) IdleBeat(context.Context, *pb.IdleBeatRequest) (*pb.IdleBeatReply, error) {
	return &pb.IdleBeatReply{}, nil
}
func (s *ExecutorControlSideServer) Kill(context.Context, *pb.KillRequest) (*pb.KillReply, error) {
	return &pb.KillReply{}, nil
}
func (s *ExecutorControlSideServer) Run(ctx context.Context, req *pb.RunRequest) (*pb.RunReply, error) {
	// 调用执行器的任务执行方法
	err := s.executor.run(req.UnitName, req.UnitArgs)

	return &pb.RunReply{
		Success: err == nil,
		Key:     s.executor.opts.Key,
	}, err
}
