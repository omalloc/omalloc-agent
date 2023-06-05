package biz

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/go-kratos/kratos/v2/log"
)

var (
	ErrCurrentAgentNoTask = errors.NotFound("CURRENT_AGENT_NO_TASK", "当前 Agent 没有任务")
)

type TaskUnit struct {
	Key           string        `json:"key"`
	Status        int           `json:"status"`
	RouteStrategy int           `json:"route_strategy"`
	BlockStrategy int           `json:"block_strategy"`
	MisfirePolicy int           `json:"misfire_policy"`
	Type          string        `json:"type"`      // 模式(CRON=定时触发, FIX_RATE=固定频率)
	TypeConf      string        `json:"type_conf"` // 模式配置单元时间(CRON~='0 * * * *' |  RATE=30s)
	Args          string        `json:"args"`      // 允许传递JSON对象
	Timeout       time.Duration `json:"timeout"`   // 超时
}

type Task struct {
	ID                     int64       `json:"id"`
	Name                   string      `json:"name"`
	Status                 int         `json:"status"`
	Type                   int         `json:"type"`
	ScheduleAvailableRatio float32     `json:"schedule_available_ratio"`
	ScheduleTotalCount     int         `json:"schedule_total_count"`
	ScheduleSuccessCount   int         `json:"schedule_success_count"`
	ScheduleFailedCount    int         `json:"schedule_failed_count"`
	TaskUnit               []*TaskUnit `json:"task_unit"`
}

type TaskRepo interface{}

type TaskUsecase struct {
	log *log.Helper
}

func NewTask(logger log.Logger) *TaskUsecase {
	return &TaskUsecase{
		log: log.NewHelper(logger),
	}
}

func (u *TaskUsecase) ListTasks(ctx context.Context, names []string) ([]*Task, error) {
	// TODO: 从内存获取任务信息, 通过 names 过滤
	// return nil, ErrCurrentAgentNoTask

	return []*Task{
		{
			ID:                     1,
			Name:                   "mock:taskname",
			Status:                 1,
			Type:                   1,
			ScheduleAvailableRatio: 50.00,
			ScheduleTotalCount:     100,
			ScheduleSuccessCount:   50,
			ScheduleFailedCount:    50,
			TaskUnit:               []*TaskUnit{},
		},
		{
			ID:                     2,
			Name:                   "mock:cleanlog",
			Status:                 1,
			Type:                   1,
			ScheduleAvailableRatio: 90.09,
			ScheduleTotalCount:     11,
			ScheduleSuccessCount:   10,
			ScheduleFailedCount:    1,
		},
	}, nil
}
