package data

import (
	"context"
	"fmt"
	"strconv"
	"strings"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"github.com/go-kratos/kratos/v2/log"
	"golang.org/x/sync/singleflight"

	v1 "github.com/hominsu/slink/api/slink/service/v1"
	"github.com/hominsu/slink/app/slink/service/internal/biz"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent"
	"github.com/hominsu/slink/app/slink/service/internal/data/ent/shortlink"
	"github.com/hominsu/slink/pkg/pagination"
)

const slinkCacheKeyPrefix = "slink_cache_key_"

var _ biz.ShortLinkRepo = (*shortLinkRepo)(nil)

type shortLinkRepo struct {
	data *Data
	ck   map[string][]string
	sg   *singleflight.Group
	log  *log.Helper
}

// NewShortLinkRepo .
func NewShortLinkRepo(data *Data, logger log.Logger) biz.ShortLinkRepo {
	sr := &shortLinkRepo{
		data: data,
		sg:   &singleflight.Group{},
		log:  log.NewHelper(log.With(logger, "module", "data/slink")),
	}
	sr.ck = make(map[string][]string)
	sr.ck["GetByKey"] = []string{"get", "user", "key"}
	sr.ck["List"] = []string{"list", "user"}
	return sr
}

func (r *shortLinkRepo) Create(ctx context.Context, sl *biz.ShortLink) (*biz.ShortLink, error) {
	m := r.createBuilder(sl)
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		s, tErr := toShortLink(res)
		if tErr != nil {
			return nil, v1.ErrorInternal("internal error: %v", tErr)
		}
		return s, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, v1.ErrorConflict("short link already exists: %v", err)
	case ent.IsConstraintError(err):
		return nil, v1.ErrorConflict("invalid argument: %v", err)
	default:
		return nil, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) Upsert(ctx context.Context, sl *biz.ShortLink) (*biz.ShortLink, error) {
	tx, err := r.data.db.Tx(ctx)
	if err != nil {
		return nil, v1.ErrorInternal("create transactional client error: %v", err)
	}
	defer func() {
		if v := recover(); v != nil {
			if rErr := tx.Rollback(); rErr != nil {
				r.log.Warnf("rollback failed, err: %v", rErr)
			}
			panic(v)
		}
	}()

	err = r.createBuilder(sl).OnConflict().UpdateExpireAt().Exec(ctx)
	switch {
	case err == nil:
		if cErr := tx.Commit(); cErr != nil {
			return nil, v1.ErrorInternal("failed commits the transaction, err: %v", cErr)
		}
		return r.GetByKey(ctx, sl.Key)
	default:
		if rErr := tx.Rollback(); rErr != nil {
			return nil, v1.ErrorInternal("rollback failed, err: %v",
				fmt.Errorf("%w: rolling back transaction: %v", err, rErr))
		}
		return nil, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) GetByKey(ctx context.Context, key string) (*biz.ShortLink, error) {
	cacheKey := r.cacheKey(key, r.ck["GetByKey"]...)
	res, err, _ := r.sg.Do(cacheKey, func() (any, error) {
		// get from db
		get, dErr := r.data.db.ShortLink.Query().
			Where(shortlink.KeyEQ(key)).
			Only(ctx)
		return get, dErr
	})

	switch {
	case err == nil: // db hit
		return toShortLink(res.(*ent.ShortLink))
	case ent.IsNotFound(err): // db miss
		return nil, v1.ErrorNotFound("short link not found: %v", err)
	default: // error
		return nil, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) GetByLink(ctx context.Context, link string) (*biz.ShortLink, error) {
	cacheKey := r.cacheKey(link, r.ck["GetByKey"]...)
	res, err, _ := r.sg.Do(cacheKey, func() (any, error) {
		// get from db
		get, dErr := r.data.db.ShortLink.Query().
			Where(shortlink.LinkEQ(link)).
			Only(ctx)
		return get, dErr
	})

	switch {
	case err == nil: // db hit
		return toShortLink(res.(*ent.ShortLink))
	case ent.IsNotFound(err): // db miss
		return nil, v1.ErrorNotFound("link not found: %v", err)
	default: // error
		return nil, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) DeleteByKey(ctx context.Context, key string) error {
	_, err := r.data.db.ShortLink.Delete().
		Where(shortlink.KeyEQ(key)).
		Exec(ctx)
	switch {
	case err == nil:
		return nil
	case ent.IsNotFound(err):
		return v1.ErrorNotFound("short link not found: %v", err)
	default:
		return v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) List(ctx context.Context, pageSize int, pageToken string) (*biz.ShortLinkPage, error) {
	// list users
	listQuery := r.data.db.ShortLink.Query().
		Order(ent.Asc(shortlink.FieldID)).
		Limit(pageSize + 1)
	if pageToken != "" {
		token, pErr := pagination.DecodePageToken(pageToken)
		if pErr != nil {
			return nil, v1.ErrorInternal("decode page token err: %v", pErr)
		}
		listQuery = listQuery.Where(shortlink.IDGTE(token))
	}

	cacheKey := r.cacheKey(
		strings.Join([]string{strconv.FormatInt(int64(pageSize), 10), pageToken}, "_"),
		r.ck["List"]...,
	)
	res, err, _ := r.sg.Do(cacheKey, func() (any, error) {
		// get from db
		entList, dErr := listQuery.All(ctx)
		return entList, dErr
	})

	switch {
	case err == nil:
		entList := res.([]*ent.ShortLink)
		// generate next page token
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken, err = pagination.EncodePageToken(entList[len(entList)-1].ID)
			if err != nil {
				return nil, v1.ErrorInternal("encode page token error: %v", err)
			}
			entList = entList[:len(entList)-1]
		}

		shortLinkList, tErr := toShortLinkList(entList)
		if tErr != nil {
			return nil, v1.ErrorInternal("internal error: %v", tErr)
		}
		return &biz.ShortLinkPage{
			ShortLinks:    shortLinkList,
			NextPageToken: nextPageToken,
		}, nil
	case ent.IsNotFound(err): // db miss
		return nil, v1.ErrorNotFound("short link not found: %v", err)
	default: // error
		return nil, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) Flush(ctx context.Context) error {
	_, err := r.data.db.ShortLink.Delete().
		Exec(ctx)
	switch {
	case err == nil:
		return nil
	case ent.IsNotFound(err):
		return v1.ErrorNotFound("short link not found: %v", err)
	default:
		return v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) Count(ctx context.Context) (int, error) {
	count, err := r.data.db.ShortLink.Query().
		Count(ctx)
	switch {
	case err == nil:
		return count, nil
	case ent.IsNotFound(err):
		return 0, v1.ErrorNotFound("short link not found: %v", err)
	default:
		return 0, v1.ErrorUnknown("unknown error: %v", err)
	}
}

func (r *shortLinkRepo) createBuilder(s *biz.ShortLink) *ent.ShortLinkCreate {
	m := r.data.db.ShortLink.Create()
	m.SetKey(s.Key)
	m.SetLink(s.Link)
	m.SetExpireAt(s.ExpireAt)
	return m
}

func (r *shortLinkRepo) cacheKey(unique string, a ...string) string {
	s := strings.Join(a, "_")
	return slinkCacheKeyPrefix + s + ":" + unique
}

func toShortLink(e *ent.ShortLink) (*biz.ShortLink, error) {
	s := &biz.ShortLink{
		Key:      e.Key,
		Link:     e.Link,
		ExpireAt: e.ExpireAt,
	}
	return s, nil
}

func toShortLinkList(e []*ent.ShortLink) ([]*biz.ShortLink, error) {
	userList := make([]*biz.ShortLink, len(e))
	for i, entEntity := range e {
		u, err := toShortLink(entEntity)
		if err != nil {
			return nil, err
		}
		userList[i] = u
	}
	return userList, nil
}
