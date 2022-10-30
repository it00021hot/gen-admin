package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/app/tool/controller"
)

func InitTpl(group *ghttp.RouterGroup) {
	group.Group("/tpl", func(group *ghttp.RouterGroup) {
		group.Bind(controller.Template)
	})
}
