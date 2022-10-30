package v1

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type TablePageListReq struct {
	g.Meta `path:"/pageList" tags:"GenTable" method:"post" summary:"分页查询表信息"`
	model.PageReq
	TableName    string `  json:"tableName" dc:"表名称"`
	TableComment string `  json:"tableComment" dc:"表描述"`
}

type TableDBListReq struct {
	g.Meta `path:"/db/list" tags:"GenTable" method:"post" summary:"分页查询数据库表"`
	model.PageReq
	Group        string `json:"group" d:"default" dc:"分组"`
	TableName    string `json:"tableName" dc:"表名称"`
	TableComment string `json:"tableComment" dc:"表描述"`
}

type TableDBListRes struct {
	Total int                `json:"total" dc:"总数"`
	Items []*entity.GenTable `json:"items" dc:"数据"`
}

type TableCreateReq struct {
	g.Meta `path:"/createTable" tags:"GenTable" method:"post" summary:"创建"`
	Sql    string `json:"sql" v:"required#分组名称不能为空"`
}

type TableImportSaveReq struct {
	g.Meta `path:"/importTable" tags:"GenTable" method:"post" summary:"导入"`
	Group  string   `json:"group" d:"default" dc:"分组"`
	Tables []string ` json:"tables"  v:"required#分组名称不能为空" dc:"表名称集合"`
}

type TableGetReq struct {
	g.Meta  `path:"/:tableId" tags:"GenTable" method:"get" summary:"获取详情"`
	TableId int `json:"tableId" dc:"表id"`
}

type TableGetRes struct {
	Info   *model.GenTableDTO       `json:"info" dc:"表信息"`
	Tables []*model.GenTableDTO     `json:"tables" dc:"所有表信息"`
	Items  []*entity.GenTableColumn `json:"items" dc:"字段信息"`
	Total  int                      `json:"total" dc:"总数"`
}

type TableEditReq struct {
	g.Meta `path:"/" tags:"GenTable" method:"put" summary:"保存表"`
	*model.GenTableDTO
}

type TableDelReq struct {
	g.Meta `path:"/" tags:"GenTable" method:"delete" summary:"删除表"`
	Ids    []int `json:"ids" v:"required#请选择需要删除的数据" dc:"id集合"`
}

type TablePreviewReq struct {
	g.Meta  `path:"/preview" tags:"GenTable" method:"post" summary:"预览"`
	TableId int      `json:"tableId" dc:"表id"`
	TplList []string `json:"tplList" dc:"模板名称" d:"go"`
}

type TablePreviewRes struct {
	Data g.MapStrStr `json:"data"`
}

type TableGenCodeReq struct {
	g.Meta  `path:"/batchGenCode" tags:"GenTable" method:"post" summary:"下载"`
	Ids     []int    `json:"ids" dc:"表名集合"`
	TplList []string `json:"tplList" dc:"模板名称" d:"go"`
}

type TableGenCodeRes struct {
	Data []byte `json:"data" dc:"文件流"`
}
