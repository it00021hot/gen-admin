package factory

import (
	"context"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

type Pgsql struct{}

func (d *Pgsql) DBPageList(ctx context.Context, in model.TablePageInput, db gdb.DB) (total int, list []*entity.GenTable, err error) {
	sql := gdb.FormatSqlWithArgs(`FROM pg_class c LEFT JOIN pg_namespace n ON n.oid = c.relnamespace WHERE 
		((c.relkind = 'r'::char) OR (c.relkind = 'f'::char) OR (c.relkind = 'p'::char)) AND n.nspname = ? `, []interface{}{db.GetConfig().Name})
	if !reflect.DeepEqual(in, model.TablePageInput{}) {
		if in.TableName != "" {
			sql += gdb.FormatSqlWithArgs(" and c.relname like ?", []interface{}{"%" + in.TableName + "%"})
		}
		if in.TableComment != "" {
			sql += gdb.FormatSqlWithArgs(" and obj_description(c.oid) like ?", []interface{}{"%" + in.TableComment + "%"})
		}
	}
	countSql := "SELECT count(1) " + sql
	total, err = db.GetCount(ctx, countSql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取总表数失败")
		return
	}
	if total == 0 {
		list = make([]*entity.GenTable, 0)
		return total, list, nil
	}
	sql = "SELECT c.relname AS table_name, obj_description(c.oid) AS table_comment " + sql
	sql += gdb.FormatSqlWithArgs(" Limit ? offset ?", []interface{}{in.PageSize, (in.PageNum - 1) * in.PageSize})
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取数据失败")
	}
	return

}

func (d *Pgsql) SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq, db gdb.DB) (list []*entity.GenTable, err error) {
	sql := gdb.FormatSqlWithArgs(`SELECT c.relname AS table_name, obj_description(c.oid) AS table_comment
			FROM pg_class c
         	LEFT JOIN pg_namespace n ON n.oid = c.relnamespace
			WHERE ((c.relkind = 'r'::char) OR (c.relkind = 'f'::char) OR (c.relkind = 'p'::char)) AND n.nspname = ?`, []interface{}{db.GetConfig().Name})
	if len(in.Tables) > 0 {
		tbs := gstr.TrimRight(gstr.Repeat("?,", len(in.Tables)), ",")
		sql += " and " + gdb.FormatSqlWithArgs("c.relname in ("+tbs+")", gconv.SliceAny(in.Tables))
	}
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取表格信息失败")
	}
	return
}

func (d *Pgsql) SelectDbTableColumnsByName(ctx context.Context, tableName string, db gdb.DB) (res []*entity.GenTableColumn, err error) {
	sql := `SELECT a.attname                                                            AS column_name
     , a.attnotnull::bool                                                   as is_required
     , case when d.contype = 'p' then true else false end                   as is_pk
     , ic.ordinal_position                                                  as sort
     , b.description                                                        as column_comment
     , case when ic.column_default like 'nextval%' then true else false end as is_increment
     , coalesce(character_maximum_length, numeric_precision, -1)            as length
     , t.typname                                                            AS column_type
	FROM pg_attribute a
         left join pg_class c on a.attrelid = c.oid
         left join pg_constraint d on d.conrelid = c.oid and a.attnum = d.conkey[1]
         left join pg_description b ON a.attrelid = b.objoid AND a.attnum = b.objsubid
         left join pg_type t ON a.atttypid = t.oid
         left join information_schema.columns ic on ic.column_name = a.attname and ic.table_name = c.relname
	WHERE c.relname = ?
  	and a.attisdropped is false
  	and a.attnum > 0
	ORDER BY a.attnum
`
	err = db.GetScan(ctx, &res, gdb.FormatSqlWithArgs(sql, []interface{}{tableName}))
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("查询列信息失败")
	}
	return res, nil
}

func (d *Pgsql) GetTestSql() (sql string) {
	return "select 1"
}
