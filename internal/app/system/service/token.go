package service

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/it00021hot/gen-admin/internal/app/common/model"
	commonService "github.com/it00021hot/gen-admin/internal/app/common/service"
	"github.com/it00021hot/gen-admin/library/liberr"
	"sync"
)

type gft struct {
	options *model.TokenOptions
	gT      commonService.IGfToken
	lock    *sync.Mutex
}

var gftService = &gft{
	options: nil,
	gT:      nil,
	lock:    &sync.Mutex{},
}

func GfToken() commonService.IGfToken {
	if gftService.gT == nil {
		gftService.lock.Lock()
		defer gftService.lock.Unlock()
		if gftService.gT == nil {
			ctx := gctx.New()
			err := g.Cfg().MustGet(ctx, "gfToken").Struct(&gftService.options)
			liberr.ErrIsNil(ctx, err)
			gftService.gT = commonService.GfToken(gftService.options)
		}
	}
	return gftService.gT
}
