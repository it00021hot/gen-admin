package controller

import (
	"github.com/gogf/gf/v2/net/ghttp"
	commonController "github.com/it00021hot/gen-admin/internal/app/common/controller"
)

type BaseController struct {
	commonController.BaseController
}

// Init 自动执行的初始化方法
func (c *BaseController) Init(r *ghttp.Request) {
	c.BaseController.Init(r)
}
