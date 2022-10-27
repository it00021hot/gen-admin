package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/app/system/controller"
	"github.com/it00021hot/gen-admin/internal/app/system/service"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/system", func(group *ghttp.RouterGroup) {
		// 系统初始化
		group.Bind(
			controller.DbInit,
		)
		group.Bind(
			//登录
			controller.Login,
		)
		//登录验证拦截
		service.GfToken().Middleware(group)
		//context拦截器
		group.Middleware(service.Middleware().Ctx, service.Middleware().Auth)
		group.Bind(
			controller.User,
			controller.Menu,
			controller.Role,
			controller.Dept,
			controller.Post,
			controller.DictType,
			controller.DictData,
			controller.Config,
			controller.Monitor,
			controller.LoginLog,
		)
	})
}
