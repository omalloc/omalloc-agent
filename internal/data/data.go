package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"agent/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	return &Data{}, cleanup, nil
}

func (d *Data) Check(ctx context.Context) error {
	return nil
}
