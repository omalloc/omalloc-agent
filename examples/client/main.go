package main

import (
	"context"

	"google.golang.org/grpc"

	"github.com/omalloc/agent/examples/client/executor"
)

func main() {
	exec := executor.New(
		executor.WithGRPCServer(grpc.NewServer()),
		executor.WithServeAddr("127.0.0.1"),
		executor.WithServePort(8080),
	)

	if err := exec.Handle("job1", func(args string) error {
		return nil
	}); err != nil {
		panic(err)
	}

	if err := exec.Start(context.TODO()); err != nil {
		panic(err)
	}

	<-exec.Signal()
}
