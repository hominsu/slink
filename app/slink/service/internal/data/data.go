package data

import (
	"context"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"

	"github.com/hominsu/slink/app/slink/service/internal/conf"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent/migrate"

	// driver
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/mattn/go-sqlite3"
)

var ProviderSet = wire.NewSet(
	NewData,
	NewEntClient,
	NewShortLinkRepo,
)

type Data struct {
	db   *ent.Client
	conf *conf.Data
}

// NewData .
func NewData(
	entClient *ent.Client,
	conf *conf.Data,
	logger log.Logger,
) (*Data, func(), error) {
	// NewData
	helper := log.NewHelper(log.With(logger, "module", "data"))

	data := &Data{
		db:   entClient,
		conf: conf,
	}
	return data, func() {
		if err := data.db.Close(); err != nil {
			helper.Error(err)
		}
	}, nil
}

func NewEntClient(conf *conf.Data, logger log.Logger) *ent.Client {
	helper := log.NewHelper(log.With(logger, "module", "data/ent"))

	client, err := ent.Open(
		conf.Database.Driver,
		conf.Database.Source,
	)
	if err != nil {
		helper.Fatalf("failed opening connection to db: %v", err)
	}
	// Run the auto migration tool.
	if err := client.Schema.Create(
		context.Background(),
		migrate.WithForeignKeys(true),
		migrate.WithGlobalUniqueID(true),
	); err != nil {
		helper.Fatalf("failed creating schema resources: %v", err)
	}
	return client
}
