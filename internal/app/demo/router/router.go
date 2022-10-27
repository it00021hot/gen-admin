package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/app/demo/controller"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/demo", func(group *ghttp.RouterGroup) {
		group.Bind(
			controller.Demo,
		)
	})

}
