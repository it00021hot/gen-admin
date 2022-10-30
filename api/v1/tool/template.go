package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
)

type TplListReq struct {
	g.Meta `path:"/list" tags:"Template" method:"post" summary:"获取模板列表"`
	Name   string `json:"name" dc:"模板名称"`
}

type TplListRes struct {
	Total int                  `json:"total" dc:"总数"`
	Items []model.TemplateItem `json:"items" dc:"模板列表名称"`
}

type TplTreeReq struct {
	g.Meta `path:"/tree" tags:"Template" method:"get" summary:"获取模板树"`
}

type TplTreeRes struct {
	Items []*model.FileTree `json:"items" dc:"模板列表名称"`
}

type GetTplTreeReq struct {
	g.Meta `path:"/get/tpl/tree" tags:"Template" method:"get" summary:"获取指定模板树"`
	Name   string `json:"name" dc:"模板名称" v:"required#模板名称不能为空"`
	IsDir  bool   `json:"isDir" dc:"是否只查询文件夹"`
}

type AddTplReq struct {
	g.Meta `path:"/add/tpl" tags:"Template" method:"post" summary:"新增模板"`
	Name   string `json:"name" dc:"模板名称" v:"required#模板名称不能为空"`
}

type AddTplRes struct {
	Name string `json:"fileName" dc:"名称" v:"required#模板称不能为空"`
	Path string `json:"filePath" dc:"路径"`
}

type AddFileReq struct {
	g.Meta `path:"/add/file" tags:"Template" method:"post" summary:"新增文件夹/模板文件"`
	*model.AddFileModel
}

type AddFileRes struct {
	*model.AddFileModel
}

type GetSettingsReq struct {
	g.Meta `path:"/settings" tags:"Template" method:"get" summary:"获取模板设置"`
	Name   string `json:"name" dc:"模板名称" v:"required#模板名称不能为空"`
}

type GetSettingsRes struct {
	Name    string                 `json:"name" dc:"模板名称"`
	Content map[string]interface{} `json:"content" dc:"内容"`
}

type SaveSettingsReq struct {
	g.Meta  `path:"/settings" tags:"Template" method:"put" summary:"保存模板设置"`
	Name    string                 `json:"name" dc:"模板名称" v:"required#模板名称不能为空"`
	Content map[string]interface{} `json:"content" dc:"内容"`
}

type GetContentReq struct {
	g.Meta `path:"/getContent" tags:"Template" method:"get" summary:"获取模板内容"`
	Path   string `json:"path" dc:"路径" v:"required#路径不能为空"`
}

type GetContentRes struct {
	model.GetContentModel
}

type SaveContentReq struct {
	g.Meta  `path:"/saveContent" tags:"Template" method:"put" summary:"保存模板内容"`
	Path    string `json:"path" dc:"路径" v:"required#路径不能为空"`
	Content string `json:"content" dc:"内容"`
}

type RenameReq struct {
	g.Meta  `path:"/rename" tags:"Template" method:"post" summary:"重命名"`
	Path    string `json:"path" dc:"路径" v:"required#路径不能为空"`
	NewName string `json:"newName" dc:"模板名称" v:"required#名称不能为空"`
}

type DelTplReq struct {
	g.Meta `path:"/del/tpl" tags:"Template" method:"delete" summary:"删除模板"`
	Name   string `json:"name" dc:"模板名称" v:"required#模板名称不能为空"`
}

type DelFileReq struct {
	g.Meta `path:"/del/file" tags:"Template" method:"delete" summary:"删除文件(夹)"`
	Path   string `json:"path" dc:"路径" v:"required#路径不能为空"`
	Name   string `json:"name" dc:"模板名称" v:"required#名称不能为空"`
}
