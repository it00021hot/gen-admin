package router

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/it00021hot/gen-admin/internal/app/common/controller"
)

func BindController(group *ghttp.RouterGroup) {
	group.Group("/pub", func(group *ghttp.RouterGroup) {
		group.Group("/captcha", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Captcha,
			)
		})

		// 文件上传
		group.Group("/upload", func(group *ghttp.RouterGroup) {
			group.Bind(
				controller.Upload,
			)
		})
	})
}
