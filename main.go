package main

import (
	_ "github.com/gogf/gf/contrib/drivers/mysql/v2"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/glog"
	_ "github.com/it00021hot/gen-admin/internal/app/system/packed"
	"github.com/it00021hot/gen-admin/internal/app/system/service"
	"github.com/it00021hot/gen-admin/internal/cmd"
)

func init() {
	ctx := gctx.New()
	if service.SysInitConfig["autoInit"].Bool() && service.SysInit().IsCreateConfigFile() {
		// 加载配置文件
		err := service.SysInit().LoadConfigFile()
		if err != nil {
			g.Log().Panic(ctx, err)
		}
	}
}

func main() {
	g.Log().SetFlags(glog.F_ASYNC | glog.F_TIME_DATE | glog.F_TIME_TIME | glog.F_FILE_LONG)
	cmd.Main.Run(gctx.New())
}
