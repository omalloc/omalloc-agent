package executor

import (
	"context"
	"errors"
	"fmt"
	"net/http"
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

type ServeMode string

const (
	ServeModeGRPC ServeMode = "grpc"
	ServeModeHTTP ServeMode = "http"
)

type Callback func(args string) error

type Option func(*config)

type Executor interface {
	Handle(name string, cb Callback) error
	Start(ctx context.Context) error
	Stop(ctx context.Context) error
	Signal() <-chan os.Signal
}

type config struct {
	serveMode  ServeMode // `grpc` or `http`
	grpcServer grpc.ServiceRegistrar
	Key        string
	Name       string
	Protocol   string
	Addr       string
	Port       int
}

type executor struct {
	mu sync.RWMutex

	cfg         *config
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
	// pb.RegisterExecutorServer(exec.cfg.grpcServer, exec.controlSide)

	// // run grpc client
	// resp, err := exec.registrar.Register(context.TODO(), &pb.ExecutorRegisterRequest{
	// 	ExecutorProtocol: pb.ExecutorProtocol_GRPC,
	// 	ExecutorName:     exec.cfg.Name,
	// 	ExecutorAddr:     exec.cfg.Addr,
	// 	ExecutorPort:     int32(exec.cfg.Port),
	// })
	// if err != nil {
	// 	return err
	// }

	// exec.cfg.Key = resp.Key

	// handle os signal
	signal.Notify(exec.signal, syscall.SIGINT, syscall.SIGTERM)
	return nil
}

func (exec *executor) Stop(ctx context.Context) error {
	exec.registrar.Deregister(context.TODO(), &pb.ExecutorDeregisterRequest{
		Key: exec.cfg.Key,
	})
	return nil
}

func (exec *executor) Signal() <-chan os.Signal {
	println("runnning...")
	return exec.signal
}

func New(opts ...Option) Executor {
	cfg := &config{
		Addr: "0.0.0.0",
		Port: 9000,
	}
	for _, apply := range opts {
		apply(cfg)
	}

	fmt.Printf("%#+v", cfg)

	r := &executor{
		cfg:         cfg,
		signal:      make(chan os.Signal),
		handlers:    make(map[string]Callback),
		registrar:   pb.NewExecutorRegistrarClient(&grpc.ClientConn{}),
		controlSide: NewExecutorControlSideServer(),
	}
	return r
}

// WithGRPCServer 用于设置 gRPC 服务端(与原生 grpc 端口复用)
func WithGRPCServer(server *grpc.Server) Option {
	return func(cfg *config) {
		cfg.grpcServer = server
		cfg.serveMode = ServeModeGRPC
	}
}

// WithHTTPServer 用于设置 http 服务端(与原生 http 端口复用)
func WithHTTPServer(server *http.Server) Option {
	return func(cfg *config) {
		cfg.serveMode = ServeModeHTTP
	}
}

func WithServeAddr(addr string) Option {
	return func(cfg *config) {
		cfg.Addr = addr
	}
}

func WithServePort(port int) Option {
	return func(cfg *config) {
		cfg.Port = port
	}
}
