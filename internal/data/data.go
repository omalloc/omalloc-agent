package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"github.com/omalloc/contrib/kratos/orm"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"

	"github.com/omalloc/agent/internal/conf"
)

// gorm database transaction context key
type txContextKey struct{}

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData)

// Data .
type Data struct {
	db *gorm.DB
	// TODO wrapped database client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	log := log.NewHelper(logger)

	db, err := orm.New(
		orm.WithDriver(sqlite.Open("file::memory:?cache=shared")),
		orm.WithTracing(),
		orm.WithLogger(
			orm.WithLogHelper(logger),
			orm.WIthSlowThreshold(time.Second*2),
			orm.WithDebug(),
		),
	)

	// panic the database is required.
	if err != nil {
		panic(err)
	}

	cleanup := func() {
		log.Info("closing the data resources")
		if db != nil {
			if sql, err := db.DB(); err == nil {
				sql.Close()
			}
		}
	}
	return &Data{
		db,
	}, cleanup, nil
}

// InTx .
func (d *Data) InTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return d.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, txContextKey{}, tx)
		return fn(ctx)
	})
}

// WithContext .
func (d *Data) WithContext(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value(txContextKey{}).(*gorm.DB); ok {
		return tx
	}

	return d.db.WithContext(ctx)
}

// Check checks whether the Data is valid.
func (d *Data) Check(ctx context.Context) error {
	if sql, err := d.db.DB(); err != nil {
		return err
	} else {
		return sql.PingContext(ctx)
	}
}
