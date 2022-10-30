package service

import (
	"context"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type ITableColumn interface {
	SelectDbTableColumnsByName(ctx context.Context, group string, tableName string) (res []*entity.GenTableColumn, err error)
	InitColumnField(column *entity.GenTableColumn, table *entity.GenTable)
	GetDbType(columnType string) string
	IsExistInArray(value string, array []string) bool
	GetColumnLength(columnType string) int
	IsStringObject(dataType string) bool
	IsTimeObject(dataType string) bool
	IsNumberObject(dataType string) bool
	IsNotEdit(name string) bool
	IsNotList(name string) bool
	IsNotQuery(name string) bool
	CheckNameColumn(columnName string) bool
	CheckStateColumn(columnName string) bool
	CheckTypeColumn(columnName string) bool
	CheckSexColumn(columnName string) bool
	SelectGenTableColumnListByTableId(ctx context.Context, tableId int) (list []*entity.GenTableColumn, err error)
}

var localTableColumn ITableColumn

func TableColumn() ITableColumn {
	if localTableColumn == nil {
		panic("implement not found for interface IDatabase, forgot register?")
	}
	return localTableColumn
}

func RegisterTableColumn(i ITableColumn) {
	localTableColumn = i
}
