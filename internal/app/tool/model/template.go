package model

import (
	"github.com/gogf/gf/v2/os/gtime"
)

type TemplateItem struct {
	Label string `json:"label" dc:"标签"`
	Value string `json:"value" dc:"值"`
}

type FileTree struct {
	Name       string      `json:"name" dc:"名称" v:"required#名称不能为空"`
	Path       string      `json:"path" dc:"路径" v:"required#路径不能为空"`
	ModifyTime *gtime.Time `json:"modifyTime" dc:"修改时间"`
	IsDir      bool        `json:"isDir" dc:"是否是文件夹"`
	Type       int         `json:"type" dc:"类型"`
	Children   []*FileTree `json:"children" dc:"子节点"`
}

type AddTplModel struct {
	Name string `json:"name" dc:"名称" v:"required#模板称不能为空"`
	Path string `json:"path" dc:"路径"`
}
type AddFileModel struct {
	Content    string `json:"content" dc:"内容"`
	Name       string `json:"name" dc:"名称" v:"required#名称不能为空"`
	ParentPath string `json:"parentPath" dc:"路径" v:"required#路径不能为空"`
	IsDir      bool   `json:"isDir" dc:"是否是文件夹"`
}

type GetContentModel struct {
	Content    string      `json:"content" dc:"内容"`
	Name       string      `json:"name" dc:"名称"`
	Path       string      `json:"path" dc:"路径"`
	ParentPath string      `json:"parentPath" dc:"路径"`
	ModifyTime *gtime.Time `json:"modifyTime" dc:"修改时间"`
	Size       string      `json:"size" dc:"文件大小"`
	IsDir      bool        `json:"isDir" dc:"是否是文件夹"`
}
