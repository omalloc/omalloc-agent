package service

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/samber/lo"

	pb "github.com/omalloc/agent/api/agent"
	"github.com/omalloc/agent/internal/biz"
)

type AgentService struct {
	pb.UnimplementedAgentServer

	log         *log.Helper
	taskUsecase *biz.TaskUsecase
}

func NewAgentService(logger log.Logger, taskUsecase *biz.TaskUsecase) *AgentService {
	return &AgentService{
		log:         log.NewHelper(logger),
		taskUsecase: taskUsecase,
	}
}

func (s *AgentService) ListTasks(ctx context.Context, req *pb.ListTaskRequest) (*pb.ListTaskReply, error) {
	results, err := s.taskUsecase.ListTasks(ctx, req.Name)
	if err != nil {
		return nil, err
	}

	return &pb.ListTaskReply{
		Tasks: lo.Map(results, s.mapTask),
	}, nil
}

func (s *AgentService) UpdateTask(ctx context.Context, req *pb.UpdateTaskRequest) (*pb.UpdateTaskReply, error) {
	return &pb.UpdateTaskReply{}, nil
}
