package factory

import (
	"context"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
)

type IDBType interface {
	// DBPageList 获取数据库表信息
	DBPageList(ctx context.Context, in model.TablePageInput, db gdb.DB) (total int, list []*entity.GenTable, err error)
	// SelectDbTableListByNames 根据表名、注释查询表信息
	SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq, db gdb.DB) (list []*entity.GenTable, err error)
	// SelectDbTableColumnsByName 查询表字段信息
	SelectDbTableColumnsByName(ctx context.Context, tableName string, db gdb.DB) (res []*entity.GenTableColumn, err error)
	// GetTestSql 测试连接sql
	GetTestSql() (sql string)
}

func New(dbType string) IDBType {
	var IdbType IDBType
	switch dbType {
	case "mysql":
		IdbType = &Mysql{}
	case "oracle":
		IdbType = &Oracle{}
	case "pgsql":
		IdbType = &Pgsql{}
	case "mssql":
		IdbType = &Mssql{}
	}

	return IdbType
}
