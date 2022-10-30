package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/app/system/service"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/tool", func(group *ghttp.RouterGroup) {
		//登录验证拦截
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		InitDB(group)    //数据源
		InitTable(group) //表
		InitTpl(group)   //模板
	})
}
