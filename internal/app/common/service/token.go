package service

import (
	"context"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/core/gftoken"

	"github.com/it00021hot/gen-admin/internal/app/common/consts"
	"github.com/it00021hot/gen-admin/internal/app/common/model"
)

type IGfToken interface {
	GenerateToken(ctx context.Context, key string, data interface{}) (keys string, err error)
	Middleware(group *ghttp.RouterGroup) error
	ParseToken(r *ghttp.Request) (*gftoken.CustomClaims, error)
	IsLogin(r *ghttp.Request) (b bool, failed *gftoken.AuthFailed)
	GetRequestToken(r *ghttp.Request) (token string)
	RemoveToken(ctx context.Context, token string) (err error)
}

type gfTokenImpl struct {
	*gftoken.GfToken
}

var gT = gfTokenImpl{
	GfToken: gftoken.NewGfToken(),
}

func GfToken(options *model.TokenOptions) IGfToken {
	var fun gftoken.OptionFunc
	if options.CacheModel == consts.CacheModelRedis {
		fun = gftoken.WithGRedis()
	} else {
		fun = gftoken.WithGCache()
	}
	gT.GfToken = gftoken.NewGfToken(
		gftoken.WithCacheKey(options.CacheKey),
		gftoken.WithTimeout(options.Timeout),
		gftoken.WithMaxRefresh(options.MaxRefresh),
		gftoken.WithMultiLogin(options.MultiLogin),
		gftoken.WithExcludePaths(options.ExcludePaths),
		fun,
	)
	return &gT
}
