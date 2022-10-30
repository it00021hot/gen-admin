package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type DBPageListReq struct {
	g.Meta `path:"/pageList" tags:"Database" method:"post" summary:"分页获取数据源列表"`
	model.PageReq
	Group string `json:"group" dc:"分组名称"`
	Name  string `json:"name" dc:"数据库名称"`
}

type DBListReq struct {
	g.Meta `path:"/list" tags:"Database" method:"post" summary:"获取数据源列表"`
	Group  string `json:"group" dc:"分组名称"`
	Name   string `json:"name" dc:"数据库名称"`
}

type DBPageListRes struct {
	Total int                   `json:"total" dc:"总数"`
	Items []*entity.GenDatabase `json:"items" dc:"数据"`
}

type DBListRes struct {
	Items []*entity.GenDatabase `json:"items" dc:"数据"`
}

type DBTestReq struct {
	g.Meta `path:"/test" tags:"Database" method:"post" summary:"测试连接"`
	Group  string `json:"group" v:"required#分组名称不能为空" dc:"分组名称"`
	Host   string `json:"host" v:"required#地址不能为空" dc:"地址"`
	Port   string `json:"port" v:"required#端口不能为空" dc:"端口"`
	User   string `json:"user" v:"required#账号不能为空" dc:"账号"`
	Pass   string `json:"pass" v:"required#密码不能为空" dc:"密码"`
	Name   string `json:"name" v:"required#请填写数据库名称" dc:"数据库名称"`
	Type   string `json:"type" d:"mysql" dc:"数据库类型"`
}

type DBEditTestReq struct {
	g.Meta `path:"/edit/test" tags:"Database" method:"post" summary:"测试连接"`
	Group  string `json:"group" v:"required#分组名称不能为空" dc:"分组名称"`
	Host   string `json:"host" v:"required#地址不能为空" dc:"地址"`
	Port   string `json:"port" v:"required#端口不能为空" dc:"端口"`
	User   string `json:"user" v:"required#账号不能为空" dc:"账号"`
	Pass   string `json:"pass" dc:"密码"`
	Name   string `json:"name" v:"required#请填写数据库名称" dc:"数据库名称"`
	Type   string `json:"type" d:"mysql" dc:"数据库类型"`
}

type DBGetReq struct {
	g.Meta `path:"/:id" tags:"Database" method:"get" summary:"获取详情"`
	Id     int `json:"id" dc:"id"`
}

type DBGetRes struct {
	entity.GenDatabase
}

type DBAddReq struct {
	g.Meta `path:"/" tags:"Database" method:"post" summary:"添加数据源"`
	Group  string `json:"group" v:"required#分组名称不能为空" dc:"分组名称"`
	Host   string `json:"host" v:"required#地址不能为空" dc:"地址"`
	Port   string `json:"port" v:"required#端口不能为空" dc:"端口"`
	User   string `json:"user" v:"required#账号不能为空" dc:"账号"`
	Pass   string `json:"pass" v:"required#密码不能为空" dc:"密码"`
	Name   string `json:"name" v:"required#请填写数据库名称" dc:"数据库名称"`
	Type   string `json:"type" d:"mysql" dc:"数据库类型"`
	Remark string `json:"remark" dc:"备注"`
}

type DBEditReq struct {
	g.Meta `path:"/" tags:"Database" method:"put" summary:"修改数据源"`
	Id     int64  `json:"id"         v:"required#id不能为空" dc:"id"`
	Group  string `json:"group" v:"required#分组名称不能为空" dc:"分组名称"`
	Host   string `json:"host" v:"required#地址不能为空" dc:"地址"`
	Port   string `json:"port" v:"required#端口不能为空" dc:"端口"`
	User   string `json:"user" v:"required#账号不能为空" dc:"账号"`
	Pass   string `json:"pass" dc:"密码"`
	Name   string `json:"name" v:"required#请填写数据库名称" dc:"数据库名称"`
	Type   string `json:"type" d:"mysql" dc:"数据库类型"`
	Remark string `json:"remark" dc:"备注"`
}

type DBDelReq struct {
	g.Meta `path:"/" tags:"Database" method:"delete" summary:"删除数据源"`
	Ids    []int `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}
