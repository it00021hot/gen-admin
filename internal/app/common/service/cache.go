package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/it00021hot/gen-admin/internal/app/common/consts"
	"github.com/it00021hot/gen-admin/internal/core/cache"
	"sync"
)

type ICache interface {
	cache.IGCache
}

type cacheImpl struct {
	*cache.GfCache
	prefix string
}

var (
	c              = cacheImpl{}
	cacheContainer *cache.GfCache
	lock           = &sync.Mutex{}
)

func Cache() ICache {
	var (
		ch  = c
		ctx = gctx.New()
	)
	prefix := g.Cfg().MustGet(ctx, "system.cache.prefix").String()
	model := g.Cfg().MustGet(ctx, "system.cache.model").String()
	if cacheContainer == nil {
		lock.Lock()
		if cacheContainer == nil {
			if model == consts.CacheModelRedis {
				// redis
				cacheContainer = cache.NewRedis(prefix)
			} else {
				// memory
				cacheContainer = cache.New(prefix)
			}
		}
		lock.Unlock()
	}
	ch.GfCache = cacheContainer
	return &ch
}
