package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonRouter "github.com/it00021hot/gen-admin/internal/app/common/router"
	demoRouter "github.com/it00021hot/gen-admin/internal/app/demo/router"
	systemRouter "github.com/it00021hot/gen-admin/internal/app/system/router"
	toolRouter "github.com/it00021hot/gen-admin/internal/app/tool/router"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/api/v1", func(group *ghttp.RouterGroup) {
		// 绑定后台路由
		systemRouter.BindController(group)
		// 绑定测试路由
		demoRouter.BindController(group)
		// 绑定公共路由
		commonRouter.BindController(group)
		// 绑定工具路由
		toolRouter.BindController(group)
	})

}
