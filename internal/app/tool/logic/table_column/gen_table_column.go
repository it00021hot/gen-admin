package table_column

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/it00021hot/gen-admin/internal/app/tool/dao"
	"github.com/it00021hot/gen-admin/internal/app/tool/factory"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
	"strings"
)

type (
	sTableColumn struct {
		ColumnTypeStr      []string //数据库字符串类型
		ColumnTypeTime     []string //数据库时间类型
		ColumnTypeNumber   []string //数据库数字类型
		ColumnNameNotEdit  []string //页面不需要编辑字段
		ColumnNameNotList  []string //页面不需要显示的列表字段
		ColumnNameNotQuery []string //页面不需要查询字段
	}
)

func init() {
	service.RegisterTableColumn(New())
}

func New() *sTableColumn {
	return &sTableColumn{
		ColumnTypeStr:      []string{"char", "varchar", "narchar", "varchar2", "tinytext", "text", "mediumtext", "longtext"},
		ColumnTypeTime:     []string{"datetime", "time", "date", "timestamp"},
		ColumnTypeNumber:   []string{"tinyint", "smallint", "mediumint", "int", "number", "integer", "bigint", "float", "float", "double", "decimal"},
		ColumnNameNotEdit:  []string{"id", "creator", "creator_id", "create_date", "modifier", "modifier_id", "modify_date", "version", "version_number"},
		ColumnNameNotList:  []string{"id", "version", "version_number"},
		ColumnNameNotQuery: []string{"id", "state", "version", "version_number", "remark"},
	}
}

func (s *sTableColumn) SelectDbTableColumnsByName(ctx context.Context, group string, tableName string) (res []*entity.GenTableColumn, err error) {
	db := g.DB()
	var idbType factory.IDBType
	if group != gdb.DefaultGroupName {
		var dbConfig *entity.GenDatabase
		err = dao.GenDatabase.Ctx(ctx).Where(dao.GenDatabase.Columns().Group, group).Scan(&dbConfig)
		if err != nil {
			return
		}
		gdb.SetConfigGroup(group, gdb.ConfigGroup{gdb.ConfigNode{
			Host:   dbConfig.Host,
			Port:   dbConfig.Port,
			User:   dbConfig.User,
			Pass:   dbConfig.Pass,
			Name:   dbConfig.Name,
			Type:   dbConfig.Type,
			Debug:  true,
			DryRun: true,
		}})
		db, err = gdb.Instance(group)
		if err != nil {
			return
		}
	}
	idbType = factory.New(db.GetConfig().Type)
	return idbType.SelectDbTableColumnsByName(ctx, tableName, db)
}

// InitColumnField 初始化列属性字段
func (s *sTableColumn) InitColumnField(column *entity.GenTableColumn, table *entity.GenTable) {
	dataType := s.GetDbType(column.ColumnType)
	columnName := column.ColumnName
	column.TableId = table.TableId
	//设置字段名
	column.GoField = gstr.CaseCamel(gstr.ToLower(columnName))
	column.JavaField = gstr.CaseCamelLower(gstr.ToLower(columnName))
	column.TsField = column.JavaField
	if s.IsStringObject(dataType) {
		//字段为字符串类型
		column.GoType = "string"
		column.JavaType = "String"
		column.TsType = "string"
		if column.ColumnLength == 0 {
			column.ColumnLength = s.GetColumnLength(column.ColumnType)
		}
		if column.ColumnLength >= 500 {
			column.HtmlType = "textarea"
		} else {
			column.HtmlType = "input"
		}
	} else if s.IsTimeObject(dataType) {
		//字段为时间类型
		column.GoType = "*gtime.Time"
		column.JavaType = "Date"
		column.HtmlType = "datetime"
		column.TsType = "string"
	} else if s.IsNumberObject(dataType) {
		//字段为数字类型
		column.HtmlType = "input"
		column.TsType = "number"
		t, _ := gregex.ReplaceString(`\(.+\)`, "", column.ColumnType)
		t = gstr.Split(gstr.Trim(t), " ")[0]
		t = gstr.ToLower(t)
		// 如果是浮点型
		switch t {
		case "float", "double", "decimal":
			column.GoType = "float64"
			column.JavaType = "BigDecimal"
		case "tinyint", "small_int", "smallint":
			column.GoType = "int8"
			column.JavaType = "Integer"
		case "bit", "int", "medium_int", "mediumint", "int unsigned":
			if gstr.ContainsI(column.ColumnType, "unsigned") {
				column.GoType = "uint"
				column.JavaType = "Integer"
			} else {
				column.GoType = "int"
				column.JavaType = "Integer"
			}
		case "big_int", "bigint":
			if gstr.ContainsI(column.ColumnType, "unsigned") {
				column.GoType = "uint64"
				column.JavaType = "Long"
			} else {
				column.GoType = "int64"
				column.JavaType = "Long"
			}
		}
	}
	//新增字段
	if s.IsNotEdit(columnName) {
		column.IsRequired = false
		column.IsInsert = false
	} else {
		column.IsInsert = true
		if strings.Index(columnName, "name") >= 0 || strings.Index(columnName, "state") >= 0 {
			column.IsRequired = true
		}
	}

	// 编辑字段
	if s.IsNotEdit(columnName) {
		column.IsEdit = false
	} else {
		if column.IsPk {
			column.IsEdit = false
		} else {
			column.IsEdit = true
		}
	}
	// 列表字段
	if s.IsNotList(columnName) {
		column.IsList = false
	} else {
		column.IsList = true
	}
	// 查询字段
	if s.IsNotQuery(columnName) {
		column.IsQuery = false
	} else {
		column.IsQuery = true
	}

	// 查询字段类型
	if s.CheckNameColumn(columnName) {
		column.QueryType = "LIKE"
	} else if s.IsTimeObject(dataType) {
		column.QueryType = "BETWEEN"
	} else {
		column.QueryType = "EQ"
	}

	// 状态字段设置单选框
	if s.CheckStateColumn(columnName) {
		column.HtmlType = "radio"
	} else if s.CheckTypeColumn(columnName) || s.CheckSexColumn(columnName) {
		// 类型&性别字段设置下拉框
		column.HtmlType = "select"
	}
}

// GetDbType 获取数据库类型字段
func (s *sTableColumn) GetDbType(columnType string) string {
	if strings.Index(columnType, "(") > 0 {
		return columnType[0:strings.Index(columnType, "(")]
	} else {
		return columnType
	}
}

// IsExistInArray 判断 value 是否存在在切片array中
func (s *sTableColumn) IsExistInArray(value string, array []string) bool {
	for _, v := range array {
		if v == value {
			return true
		}
	}
	return false
}

// GetColumnLength 获取字段长度
func (s *sTableColumn) GetColumnLength(columnType string) int {
	start := strings.Index(columnType, "(")
	end := strings.Index(columnType, ")")
	result := ""
	if start >= 0 && end >= 0 {
		result = columnType[start+1 : end]
	}
	return gconv.Int(result)
}

// IsStringObject 判断是否是数据库字符串类型
func (s *sTableColumn) IsStringObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeStr)
}

// IsTimeObject 判断是否是数据库时间类型
func (s *sTableColumn) IsTimeObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeTime)
}

// IsNumberObject 是否数字类型
func (s *sTableColumn) IsNumberObject(dataType string) bool {
	return s.IsExistInArray(dataType, s.ColumnTypeNumber)
}

// IsNotEdit 是否不可编辑
func (s *sTableColumn) IsNotEdit(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotEdit)
}

// IsNotList 不在列表显示
func (s *sTableColumn) IsNotList(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotList)
}

// IsNotQuery 不可用于查询
func (s *sTableColumn) IsNotQuery(name string) bool {
	return s.IsExistInArray(name, s.ColumnNameNotQuery)
}

// CheckNameColumn 检查字段名后4位是否是name
func (s *sTableColumn) CheckNameColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4
		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]
		if tmp == "name" {
			return true
		}
	}
	return false
}

// CheckStateColumn 检查字段名后5位是否是state
func (s *sTableColumn) CheckStateColumn(columnName string) bool {
	if len(columnName) >= 5 {
		end := len(columnName)
		start := end - 5

		if start <= 0 {
			start = 0
		}
		tmp := columnName[start:end]

		if tmp == "state" {
			return true
		}
	}

	return false
}

// CheckTypeColumn 检查字段名后4位是否是type
func (s *sTableColumn) CheckTypeColumn(columnName string) bool {
	if len(columnName) >= 4 {
		end := len(columnName)
		start := end - 4

		if start <= 0 {
			start = 0
		}

		if columnName[start:end] == "type" {
			return true
		}
	}
	return false
}

// CheckSexColumn 检查字段名后3位是否是sex
func (s *sTableColumn) CheckSexColumn(columnName string) bool {
	if len(columnName) >= 3 {
		end := len(columnName)
		start := end - 3
		if start <= 0 {
			start = 0
		}
		if columnName[start:end] == "sex" {
			return true
		}
	}
	return false
}

func (s *sTableColumn) SelectGenTableColumnListByTableId(ctx context.Context, tableId int) (list []*entity.GenTableColumn, err error) {
	err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, tableId).
		Order(dao.GenTableColumn.Columns().Sort + " asc, " + dao.GenTableColumn.Columns().ColumnId + " asc").Scan(&list)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取字段信息出错")
	}
	return list, nil
}
