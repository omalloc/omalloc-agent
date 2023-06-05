package server

import (
	"context"

	"github.com/google/wire"
	"github.com/omalloc/contrib/kratos/health"
	"github.com/omalloc/contrib/kratos/registry"
	"github.com/omalloc/contrib/protobuf"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/omalloc/agent/internal/conf"
	"github.com/omalloc/agent/internal/data"
)

// ProviderSet is server providers.
var ProviderSet = wire.NewSet(
	NewGRPCServer,
	NewHTTPServer,
	NewRegistryConfig,
	NewChecker,

	registry.NewEtcd,
	registry.NewRegistrar,
	registry.NewDiscovery,

	health.NewServer,
)

func NewRegistryConfig(bc *conf.Bootstrap) *protobuf.Registry {
	return bc.Registry
}

func NewTracingConfig(bc *conf.Bootstrap) *protobuf.Tracing {
	return bc.Tracing
}

func NewChecker(c1 *data.Data, etcd *clientv3.Client) []health.Checker {
	return []health.Checker{c1, &ConnectorETCD{etcd}}
}

type ConnectorETCD struct {
	*clientv3.Client
}

// Checke is a health-checker.
func (c *ConnectorETCD) Check(ctx context.Context) error {
	for _, e := range c.Endpoints() {
		if _, err := c.Status(ctx, e); err != nil {
			return err
		}
	}
	// TODO check the database
	return nil
}
