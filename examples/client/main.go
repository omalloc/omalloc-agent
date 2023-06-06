package main

import (
	"context"

	"github.com/omalloc/agent/examples/client/sdk"
)

func main() {
	exec := sdk.New()

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
