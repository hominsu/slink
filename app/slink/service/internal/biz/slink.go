package biz

import (
	"context"
	"strconv"

	"github.com/go-kratos/kratos/v2/log"
	cuckoo "github.com/seiflotfy/cuckoofilter"
	"github.com/spaolacci/murmur3"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/internal/conf"
	"github.com/hominsu/slink/pkg/utils"
)

type ShortLink struct {
	Key  string `json:"key,omitempty"`
	Link string `json:"link,omitempty"`
}

type ShortLinkPage struct {
	NextPageToken string
	ShortLinks    []*ShortLink
}

type ShortLinkRepo interface {
	Create(ctx context.Context, sl *ShortLink) (*ShortLink, error)
	GetByKey(ctx context.Context, key string) (*ShortLink, error)
	GetByLink(ctx context.Context, link string) (*ShortLink, error)
	DeleteByKey(ctx context.Context, key string) error
	List(ctx context.Context, pageSize int, pageToken string) (*ShortLinkPage, error)
	Flush(ctx context.Context) error
	Count(ctx context.Context) (int, error)
}

type ShortLinkRepoUsecase struct {
	sr     ShortLinkRepo
	filter *cuckoo.Filter
	log    *log.Helper
	retry  int
}

func NewShortLinkRepoUsecase(sr ShortLinkRepo, conf *conf.Data, logger log.Logger) *ShortLinkRepoUsecase {
	helper := log.NewHelper(log.With(logger, "module", "biz/slink"))
	if conf.Cuckoo.Retry == 0 {
		helper.Fatalf("retry num should be greater than zero")
	}

	return &ShortLinkRepoUsecase{
		sr:     sr,
		filter: cuckoo.NewFilter(uint(conf.Cuckoo.Capacity)),
		log:    log.NewHelper(logger),
		retry:  int(conf.Cuckoo.Retry),
	}
}

func (uc *ShortLinkRepoUsecase) Create(ctx context.Context, link string) (*ShortLink, error) {
	var key string
	fLink := link
	for i := 0; i < uc.retry; i++ {
		key = strconv.FormatUint(uint64(murmur3.Sum32([]byte(fLink))), 36)
		if !uc.filter.Lookup([]byte(key)) {
			s, err := uc.sr.Create(ctx, &ShortLink{
				Key:  key,
				Link: link,
			})
			switch {
			case err == nil:
				uc.filter.InsertUnique([]byte(key))
				return s, nil
			case v1.IsConflict(err):
				return uc.sr.GetByLink(ctx, link)
			default:
				return nil, err
			}
		}
		fLink += utils.RandString(10, utils.AllCharSet)
	}
	return nil, v1.ErrorConflict("can not short link")
}

func (uc *ShortLinkRepoUsecase) Get(ctx context.Context, key string) (*ShortLink, error) {
	if uc.filter.Lookup([]byte(key)) {
		return uc.sr.GetByKey(ctx, key)
	}
	return nil, v1.ErrorNotFound("short link not found")
}

func (uc *ShortLinkRepoUsecase) Del(ctx context.Context, key string) error {
	if uc.filter.Lookup([]byte(key)) {
		uc.filter.Delete([]byte(key))
		return uc.sr.DeleteByKey(ctx, key)
	}
	return v1.ErrorNotFound("short link not found")
}

func (uc *ShortLinkRepoUsecase) List(ctx context.Context, pageSize int, pageToken string) ([]string, string, error) {
	page, err := uc.sr.List(ctx, pageSize, pageToken)
	if err != nil {
		return nil, "", err
	}

	keys := make([]string, len(page.ShortLinks))
	for i, link := range page.ShortLinks {
		keys[i] = link.Key
	}
	return keys, page.NextPageToken, nil
}

func (uc *ShortLinkRepoUsecase) Flush(ctx context.Context) error {
	uc.filter.Reset()
	return uc.sr.Flush(ctx)
}

func (uc *ShortLinkRepoUsecase) Count(ctx context.Context) (int, error) {
	return uc.sr.Count(ctx)
}
