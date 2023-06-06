package sdk

import (
	"context"
	"errors"
	"os"
	"os/signal"
	"sync"
	"syscall"

	"google.golang.org/grpc"

	pb "github.com/omalloc/agent/api/executor"
)

var (
	ErrExecutorHandleExist = errors.New("executor handle exist")
)

type Callback func(args string) error

type Executor interface {
	Handle(name string, cb Callback) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Signal() <-chan os.Signal
}

type Config struct {
	Key      string
	Name     string
	Protocol string
	Addr     string
	Port     int
}

type executor struct {
	mu sync.RWMutex

	opts        *Config
	signal      chan os.Signal
	handlers    map[string]Callback
	registrar   pb.ExecutorRegistrarClient
	controlSide pb.ExecutorServer
}

func (exec *executor) Handle(name string, cb Callback) error {
	exec.mu.Lock()
	defer exec.mu.Unlock()

	// checker
	if _, exist := exec.handlers[name]; exist {
		return ErrExecutorHandleExist
	}

	exec.handlers[name] = cb
	return nil
}

func (exec *executor) run(name string, args string) error {
	exec.mu.RLock()
	defer exec.mu.RUnlock()

	cb, exist := exec.handlers[name]
	if !exist {
		return errors.New("executor handle not exist")
	}

	return cb(args)
}

func (exec *executor) Start(ctx context.Context) error {
	// run gprc server
	srv := grpc.NewServer()
	pb.RegisterExecutorServer(srv, exec.controlSide)

	// run grpc client
	resp, err := exec.registrar.Register(context.TODO(), &pb.ExecutorRegisterRequest{
		ExecutorProtocol: pb.ExecutorProtocol_GRPC,
		ExecutorName:     exec.opts.Name,
		ExecutorAddr:     exec.opts.Addr,
		ExecutorPort:     int32(exec.opts.Port),
	})
	if err != nil {
		return err
	}

	exec.opts.Key = resp.Key

	// handle os signal
	signal.Notify(exec.signal, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func (exec *executor) Stop(ctx context.Context) error {
	exec.registrar.Deregister(context.TODO(), &pb.ExecutorDeregisterRequest{
		Key: exec.opts.Key,
	})
	return nil
}

func (exec *executor) Signal() <-chan os.Signal {
	println("runnning...")
	return exec.signal
}

func New() Executor {
	return &executor{
		opts:        &Config{},
		signal:      make(chan os.Signal),
		handlers:    make(map[string]Callback),
		registrar:   pb.NewExecutorRegistrarClient(&grpc.ClientConn{}),
		controlSide: NewExecutorControlSideServer(),
	}
}
