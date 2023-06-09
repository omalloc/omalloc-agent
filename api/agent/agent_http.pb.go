// Code generated by protoc-gen-go-http. DO NOT EDIT.
// versions:
// - protoc-gen-go-http v2.6.1
// - protoc             v3.21.12
// source: agent/agent.proto

package agent

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

const OperationAgentListTasks = "/api.agent.Agent/ListTasks"
const OperationAgentUpdateTask = "/api.agent.Agent/UpdateTask"

type AgentHTTPServer interface {
	// ListTasks ListTasks list this agent all tasks.
	ListTasks(context.Context, *ListTaskRequest) (*ListTaskReply, error)
	// UpdateTask UpdateTask update task info.
	UpdateTask(context.Context, *UpdateTaskRequest) (*UpdateTaskReply, error)
}

func RegisterAgentHTTPServer(s *http.Server, srv AgentHTTPServer) {
	r := s.Route("/")
	r.GET("/agent/v1/tasks", _Agent_ListTasks0_HTTP_Handler(srv))
	r.PUT("/agent/v1/task/{id}", _Agent_UpdateTask0_HTTP_Handler(srv))
}

func _Agent_ListTasks0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in ListTaskRequest
		if err := ctx.BindQuery(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentListTasks)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.ListTasks(ctx, req.(*ListTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*ListTaskReply)
		return ctx.Result(200, reply.Tasks)
	}
}

func _Agent_UpdateTask0_HTTP_Handler(srv AgentHTTPServer) func(ctx http.Context) error {
	return func(ctx http.Context) error {
		var in UpdateTaskRequest
		if err := ctx.Bind(&in); err != nil {
			return err
		}
		if err := ctx.BindVars(&in); err != nil {
			return err
		}
		http.SetOperation(ctx, OperationAgentUpdateTask)
		h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
			return srv.UpdateTask(ctx, req.(*UpdateTaskRequest))
		})
		out, err := h(ctx, &in)
		if err != nil {
			return err
		}
		reply := out.(*UpdateTaskReply)
		return ctx.Result(200, reply)
	}
}

type AgentHTTPClient interface {
	ListTasks(ctx context.Context, req *ListTaskRequest, opts ...http.CallOption) (rsp *ListTaskReply, err error)
	UpdateTask(ctx context.Context, req *UpdateTaskRequest, opts ...http.CallOption) (rsp *UpdateTaskReply, err error)
}

type AgentHTTPClientImpl struct {
	cc *http.Client
}

func NewAgentHTTPClient(client *http.Client) AgentHTTPClient {
	return &AgentHTTPClientImpl{client}
}

func (c *AgentHTTPClientImpl) ListTasks(ctx context.Context, in *ListTaskRequest, opts ...http.CallOption) (*ListTaskReply, error) {
	var out ListTaskReply
	pattern := "/agent/v1/tasks"
	path := binding.EncodeURL(pattern, in, true)
	opts = append(opts, http.Operation(OperationAgentListTasks))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "GET", path, nil, &out.Tasks, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}

func (c *AgentHTTPClientImpl) UpdateTask(ctx context.Context, in *UpdateTaskRequest, opts ...http.CallOption) (*UpdateTaskReply, error) {
	var out UpdateTaskReply
	pattern := "/agent/v1/task/{id}"
	path := binding.EncodeURL(pattern, in, false)
	opts = append(opts, http.Operation(OperationAgentUpdateTask))
	opts = append(opts, http.PathTemplate(pattern))
	err := c.cc.Invoke(ctx, "PUT", path, in, &out, opts...)
	if err != nil {
		return nil, err
	}
	return &out, err
}
