package service

import (
	pb "github.com/omalloc/agent/api/agent"
	"github.com/omalloc/agent/internal/biz"
)

func (s *AgentService) mapTask(item *biz.Task, _ int) *pb.Task {
	return &pb.Task{
		Id:                     item.ID,
		Name:                   item.Name,
		ScheduleAvailableRatio: item.ScheduleAvailableRatio,
		ScheduleTotalCount:     int32(item.ScheduleTotalCount),
		ScheduleFailedCount:    int32(item.ScheduleFailedCount),
		ScheduleSuccessCount:   int32(item.ScheduleSuccessCount),
	}
}
