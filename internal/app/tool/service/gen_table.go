package service

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type ITable interface {
	PageList(ctx context.Context, in model.TablePageInput) (total int, list []*entity.GenTable, err error)
	DBPageList(ctx context.Context, in model.TablePageInput) (total int, list []*entity.GenTable, err error)
	Create(ctx context.Context, sql string) (err error)
	InitTable(ctx context.Context, table *entity.GenTable)
	InitGenInfo(table *model.GenTableDTO, genMap map[string]*g.Var)
	ConvertClassName(ctx context.Context, tableName string) string
	GetBusinessName(ctx context.Context, tableName string) string
	RemoveTablePrefix(ctx context.Context, tableName string) string
	SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq) (list []*entity.GenTable, err error)
	ImportGenTable(ctx context.Context, group string, tableList []*entity.GenTable) (err error)
	SelectGenTableById(ctx context.Context, tableId int) (table *model.GenTableDTO, err error)
	SelectGenTableAll(ctx context.Context) (table []*model.GenTableDTO, err error)
	SaveEdit(ctx context.Context, req *model.GenTableDTO) (err error)
	Delete(ctx context.Context, ids []int) (err error)
	PreviewCode(ctx context.Context, tableId int, tplList []string) (data g.MapStrStr, err error)
	SelectGenTableByName(ctx context.Context, tableName string) (result *model.GenTableDTO, err error)
	BatchGenCode(ctx context.Context, ids []int, tplList []string) (fileList []string, err error)
	GenCode(ctx context.Context, id int, tplList []string) (fileList []string, err error)
}

var localTable ITable

func Table() ITable {
	if localTable == nil {
		panic("implement not found for interface ITable, forgot register?")
	}
	return localTable
}

func RegisterTable(i ITable) {
	localTable = i
}
