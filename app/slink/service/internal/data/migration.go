package data

import (
	"context"
	"time"

	"github.com/go-kratos/kratos/v2/log"
	cuckoo "github.com/seiflotfy/cuckoofilter"

	"github.com/hominsu/slink/app/slink/service/internal/data/ent"
)

type MigrationStatus struct{}

func Migration(entClient *ent.Client, filter *cuckoo.Filter, logger log.Logger) *MigrationStatus {
	helper := log.NewHelper(log.With(logger, "module", "data/migration"))

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	setCuckooFilter(ctx, entClient, filter, helper)

	return &MigrationStatus{}
}

func setCuckooFilter(ctx context.Context, client *ent.Client, filter *cuckoo.Filter, helper *log.Helper) {
	targets, err := client.ShortLink.Query().All(ctx)
	if err != nil {
		helper.Fatalf("failed query short link")
	}

	for _, target := range targets {
		filter.InsertUnique([]byte(target.Key))
	}
}
