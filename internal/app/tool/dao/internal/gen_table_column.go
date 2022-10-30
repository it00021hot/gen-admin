package internal

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/frame/g"
)

type GenTableColumnDao struct {
	table     string                // 表名称
	group     string                // 数据源分组，默认default
	columns   GenTableColumnColumns // 表字段
	columnMap map[string]string     //表字段map
}

// GenTableColumnColumns defines and stores column names for table gen_database.
type GenTableColumnColumns struct {
	ColumnId      string //编号
	TableId       string //归属表编号
	ColumnName    string //列名称
	ColumnComment string //列描述
	ColumnType    string //列类型
	ColumnLength  string //列长度
	JavaType      string //JAVA类型
	JavaField     string //JAVA字段名
	GoType        string //GO类型
	GoField       string //GO字段名
	TsType        string //ts类型
	TsField       string //ts字段
	IsPk          string //是否主键（1是）
	IsIncrement   string //是否自增（1是）
	IsRequired    string //是否必填（1是）
	IsInsert      string //是否为插入字段（1是）
	IsEdit        string //是否编辑字段（1是）
	IsList        string //是否列表字段（1是）
	IsQuery       string //是否查询字段（1是）
	QueryType     string //查询方式（等于、不等于、大于、小于、范围）
	HtmlType      string //显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType      string //字典类型
	Sort          string //排序
	CreateBy      string //创建者
	CreateTime    string //创建时间
	UpdateBy      string //更新者
	UpdateTime    string //更新时间
}

// genTableColumnColumns holds the columns for table mes_gf_mst_yield_line.
var genTableColumnColumns = GenTableColumnColumns{
	ColumnId:      "column_id",      //编号
	TableId:       "table_id",       //归属表编号
	ColumnName:    "column_name",    //列名称
	ColumnComment: "column_comment", //列描述
	ColumnType:    "column_type",    //列类型
	ColumnLength:  "column_length",  //列长度
	JavaType:      "java_type",      //JAVA类型
	JavaField:     "java_field",     //JAVA字段名
	GoType:        "go_type",        //GO类型
	GoField:       "go_field",       //GO字段名
	TsType:        "ts_type",        //ts类型
	TsField:       "ts_field",       //ts字段
	IsPk:          "is_pk",          //是否主键（1是）
	IsIncrement:   "is_increment",   //是否自增（1是）
	IsRequired:    "is_required",    //是否必填（1是）
	IsInsert:      "is_insert",      //是否为插入字段（1是）
	IsEdit:        "is_edit",        //是否编辑字段（1是）
	IsList:        "is_list",        //是否列表字段（1是）
	IsQuery:       "is_query",       //是否查询字段（1是）
	QueryType:     "query_type",     //查询方式（等于、不等于、大于、小于、范围）
	HtmlType:      "html_type",      //显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	DictType:      "dict_type",      //字典类型
	Sort:          "sort",           //排序
	CreateBy:      "create_by",      //创建者
	CreateTime:    "create_time",    //创建时间
	UpdateBy:      "update_by",      //更新者
	UpdateTime:    "update_time",    //更新时间
}
var genTableColumnColumnMap = map[string]string{
	"ColumnId":      "column_id",      //编号
	"TableId":       "table_id",       //归属表编号
	"ColumnName":    "column_name",    //列名称
	"ColumnComment": "column_comment", //列描述
	"ColumnType":    "column_type",    //列类型
	"ColumnLength":  "column_length",  //列长度
	"JavaType":      "java_type",      //JAVA类型
	"JavaField":     "java_field",     //JAVA字段名
	"GoType":        "go_type",        //GO类型
	"GoField":       "go_field",       //GO字段名
	"TsType":        "ts_type",        //ts类型
	"TsField":       "ts_field",       //ts字段
	"IsPk":          "is_pk",          //是否主键（1是）
	"IsIncrement":   "is_increment",   //是否自增（1是）
	"IsRequired":    "is_required",    //是否必填（1是）
	"IsInsert":      "is_insert",      //是否为插入字段（1是）
	"IsEdit":        "is_edit",        //是否编辑字段（1是）
	"IsList":        "is_list",        //是否列表字段（1是）
	"IsQuery":       "is_query",       //是否查询字段（1是）
	"QueryType":     "query_type",     //查询方式（等于、不等于、大于、小于、范围）
	"HtmlType":      "html_type",      //显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）
	"DictType":      "dict_type",      //字典类型
	"Sort":          "sort",           //排序
	"CreateBy":      "create_by",      //创建者
	"CreateTime":    "create_time",    //创建时间
	"UpdateBy":      "update_by",      //更新者
	"UpdateTime":    "update_time",    //更新时间
}

// NewGenTableColumnDao creates and returns a new DAO object for table data access.
func NewGenTableColumnDao() *GenTableColumnDao {
	return &GenTableColumnDao{
		group:     "default",
		table:     "gen_table_column",
		columns:   genTableColumnColumns,
		columnMap: genTableColumnColumnMap,
	}
}

// DB retrieves and returns the underlying raw database management object of current DAO.
func (dao *GenTableColumnDao) DB() gdb.DB {
	return g.DB(dao.group)
}

// Table returns the table name of current dao.
func (dao *GenTableColumnDao) Table() string {
	return dao.table
}

// Columns returns all column names of current dao.
func (dao *GenTableColumnDao) Columns() GenTableColumnColumns {
	return dao.columns
}

// ColumnMap returns all column map of current dao.
func (dao *GenTableColumnDao) ColumnMap() map[string]string {
	return dao.columnMap
}

// Group returns the configuration group name of database of current dao.
func (dao *GenTableColumnDao) Group() string {
	return dao.group
}

// Ctx creates and returns the Model for current DAO, It automatically sets the context for current operation.
func (dao *GenTableColumnDao) Ctx(ctx context.Context) *gdb.Model {
	return dao.DB().Model(dao.table).Safe().Ctx(ctx)
}

// Transaction wraps the transaction logic using function f.
// It rollbacks the transaction and returns the error from function f if it returns non-nil error.
// It commits the transaction and returns nil if function f returns nil.
//
// Note that, you should not Commit or Rollback the transaction in function f
// as it is automatically handled by this function.
func (dao *GenTableColumnDao) Transaction(ctx context.Context, f func(ctx context.Context, tx *gdb.TX) error) (err error) {
	return dao.Ctx(ctx).Transaction(ctx, f)
}
