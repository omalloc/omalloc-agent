// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: executor/executor.proto

package executor

import (
	context "context"
	http "github.com/go-kratos/kratos/v2/transport/http"
	binding "github.com/go-kratos/kratos/v2/transport/http/binding"
)

// This is a compile-time assertion to ensure that this generated file
// is compatible with the kratos package it is being compiled against.
var _ = new(context.Context)
var _ = binding.EncodeURL

const _ = http.SupportPackageIsVersion1

const OperationExecutorRegistrarDeregister = "/api.executor.ExecutorRegistrar/Deregister"
const OperationExecutorRegistrarRegister = "/api.executor.ExecutorRegistrar/Register"

type ExecutorRegistrarHTTPServer interface {
	// Deregister Deregister the registration.
	Deregister(context.Context, *ExecutorDeregisterRequest) (*ExecutorDeregisterReply, error)
	// Register Register the registration.
	Register(context.Context, *ExecutorRegisterRequest) (*ExecutorRegisterReply, error)
}

func RegisterExecutorRegistrarHTTPServer(s *http.Server, srv ExecutorRegistrarHTTPServer) {
	r := s.Route("/")
	r.POST("/executor/v1/register", _ExecutorRegistrar_Register0_HTTP_Handler(srv))
	r.POST("/executor/v1/deregister", _ExecutorRegistrar_Deregister0_HTTP_Handler(srv))
}

func _ExecutorRegistrar_Register0_HTTP_Handler(srv ExecutorRegistrarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExecutorRegisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorRegistrarRegister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Register(ctx, req.(*ExecutorRegisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecutorRegisterReply)
		return ctx.Result(200, reply)
	}
}

func _ExecutorRegistrar_Deregister0_HTTP_Handler(srv ExecutorRegistrarHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ExecutorDeregisterRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorRegistrarDeregister)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Deregister(ctx, req.(*ExecutorDeregisterRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ExecutorDeregisterReply)
		return ctx.Result(200, reply)
	}
}

type ExecutorRegistrarHTTPClient interface {
	Deregister(ctx context.Context, req *ExecutorDeregisterRequest, opts ...http.CallOption) (rsp *ExecutorDeregisterReply, err error)
	Register(ctx context.Context, req *ExecutorRegisterRequest, opts ...http.CallOption) (rsp *ExecutorRegisterReply, err error)
}

type ExecutorRegistrarHTTPClientImpl struct {
	cc *http.Client
}

func NewExecutorRegistrarHTTPClient(client *http.Client) ExecutorRegistrarHTTPClient {
	return &ExecutorRegistrarHTTPClientImpl{client}
}

func (c *ExecutorRegistrarHTTPClientImpl) Deregister(ctx context.Context, in *ExecutorDeregisterRequest, opts ...http.CallOption) (*ExecutorDeregisterReply, error) {
	var out ExecutorDeregisterReply
	pattern := "/executor/v1/deregister"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorRegistrarDeregister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExecutorRegistrarHTTPClientImpl) Register(ctx context.Context, in *ExecutorRegisterRequest, opts ...http.CallOption) (*ExecutorRegisterReply, error) {
	var out ExecutorRegisterReply
	pattern := "/executor/v1/register"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorRegistrarRegister))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

const OperationExecutorBeat = "/api.executor.Executor/Beat"
const OperationExecutorGetUnitLog = "/api.executor.Executor/GetUnitLog"
const OperationExecutorIdleBeat = "/api.executor.Executor/IdleBeat"
const OperationExecutorKill = "/api.executor.Executor/Kill"
const OperationExecutorRun = "/api.executor.Executor/Run"

type ExecutorHTTPServer interface {
	// Beat Beat 心跳检测
	Beat(context.Context, *BeatRequest) (*BeatReply, error)
	// GetUnitLog GetUnitLog 获取任务日志
	GetUnitLog(context.Context, *UnitLogRequest) (*UnitLogReply, error)
	// IdleBeat IdleBeat 忙碌检测
	IdleBeat(context.Context, *IdleBeatRequest) (*IdleBeatReply, error)
	// Kill Kill 杀死任务单元
	Kill(context.Context, *KillRequest) (*KillReply, error)
	// Run Run 触发任务单元
	Run(context.Context, *RunRequest) (*RunReply, error)
}

func RegisterExecutorHTTPServer(s *http.Server, srv ExecutorHTTPServer) {
	r := s.Route("/")
	r.POST("/executor/v1/run", _Executor_Run0_HTTP_Handler(srv))
	r.POST("/executor/v1/kill", _Executor_Kill0_HTTP_Handler(srv))
	r.POST("/executor/v1/log", _Executor_GetUnitLog0_HTTP_Handler(srv))
	r.POST("/executor/v1/beat", _Executor_Beat0_HTTP_Handler(srv))
	r.POST("/executor/v1/idlebeat", _Executor_IdleBeat0_HTTP_Handler(srv))
}

func _Executor_Run0_HTTP_Handler(srv ExecutorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in RunRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorRun)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Run(ctx, req.(*RunRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*RunReply)
		return ctx.Result(200, reply)
	}
}

func _Executor_Kill0_HTTP_Handler(srv ExecutorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in KillRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorKill)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Kill(ctx, req.(*KillRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*KillReply)
		return ctx.Result(200, reply)
	}
}

func _Executor_GetUnitLog0_HTTP_Handler(srv ExecutorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UnitLogRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorGetUnitLog)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.GetUnitLog(ctx, req.(*UnitLogRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UnitLogReply)
		return ctx.Result(200, reply)
	}
}

func _Executor_Beat0_HTTP_Handler(srv ExecutorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in BeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorBeat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.Beat(ctx, req.(*BeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*BeatReply)
		return ctx.Result(200, reply)
	}
}

func _Executor_IdleBeat0_HTTP_Handler(srv ExecutorHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in IdleBeatRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationExecutorIdleBeat)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.IdleBeat(ctx, req.(*IdleBeatRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*IdleBeatReply)
		return ctx.Result(200, reply)
	}
}

type ExecutorHTTPClient interface {
	Beat(ctx context.Context, req *BeatRequest, opts ...http.CallOption) (rsp *BeatReply, err error)
	GetUnitLog(ctx context.Context, req *UnitLogRequest, opts ...http.CallOption) (rsp *UnitLogReply, err error)
	IdleBeat(ctx context.Context, req *IdleBeatRequest, opts ...http.CallOption) (rsp *IdleBeatReply, err error)
	Kill(ctx context.Context, req *KillRequest, opts ...http.CallOption) (rsp *KillReply, err error)
	Run(ctx context.Context, req *RunRequest, opts ...http.CallOption) (rsp *RunReply, err error)
}

type ExecutorHTTPClientImpl struct {
	cc *http.Client
}

func NewExecutorHTTPClient(client *http.Client) ExecutorHTTPClient {
	return &ExecutorHTTPClientImpl{client}
}

func (c *ExecutorHTTPClientImpl) Beat(ctx context.Context, in *BeatRequest, opts ...http.CallOption) (*BeatReply, error) {
	var out BeatReply
	pattern := "/executor/v1/beat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorBeat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExecutorHTTPClientImpl) GetUnitLog(ctx context.Context, in *UnitLogRequest, opts ...http.CallOption) (*UnitLogReply, error) {
	var out UnitLogReply
	pattern := "/executor/v1/log"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorGetUnitLog))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExecutorHTTPClientImpl) IdleBeat(ctx context.Context, in *IdleBeatRequest, opts ...http.CallOption) (*IdleBeatReply, error) {
	var out IdleBeatReply
	pattern := "/executor/v1/idlebeat"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorIdleBeat))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExecutorHTTPClientImpl) Kill(ctx context.Context, in *KillRequest, opts ...http.CallOption) (*KillReply, error) {
	var out KillReply
	pattern := "/executor/v1/kill"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorKill))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *ExecutorHTTPClientImpl) Run(ctx context.Context, in *RunRequest, opts ...http.CallOption) (*RunReply, error) {
	var out RunReply
	pattern := "/executor/v1/run"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationExecutorRun))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "POST", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
