package factory

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"reflect"
)

type Mysql struct{}

func (d *Mysql) DBPageList(ctx context.Context, in model.TablePageInput, db gdb.DB) (total int, list []*entity.GenTable, err error) {
	sql := " from information_schema.tables where table_schema = (select database())"
	if !reflect.DeepEqual(in, model.TablePageInput{}) {
		if in.TableName != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(table_name) like lower(?)", []interface{}{"%" + in.TableName + "%"})
		}

		if in.TableComment != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(table_comment) like lower(?)", []interface{}{"%" + in.TableComment + "%"})
		}
	}
	countSql := "select count(1) " + sql
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
	sql = "table_name, table_comment, create_time, update_time " + sql
	page := (in.PageNum - 1) * in.PageSize
	sql += " order by create_time desc,table_name asc limit  " + gconv.String(page) + "," + gconv.String(in.PageSize)
	err = db.GetScan(ctx, &list, "select "+sql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取数据失败")
	}
	return

}

func (d *Mysql) SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq, db gdb.DB) (list []*entity.GenTable, err error) {
	sql := "select table_name, table_comment, create_time, update_time  from information_schema.tables where table_schema = (select database()) "
	if len(in.Tables) > 0 {
		tbs := gstr.TrimRight(gstr.Repeat("?,", len(in.Tables)), ",")
		sql += " and " + gdb.FormatSqlWithArgs("table_name in ("+tbs+")", gconv.SliceAny(in.Tables))
	}
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取表格信息失败")
	}
	return
}

func (d *Mysql) SelectDbTableColumnsByName(ctx context.Context, tableName string, db gdb.DB) (res []*entity.GenTableColumn, err error) {
	sql := " select column_name, (case when (is_nullable = 'no' && column_key != 'PRI') then '1' else null end) as is_required, " +
		"(case when column_key = 'PRI' then '1' else '0' end) as is_pk, ordinal_position as sort, column_comment," +
		" (case when extra = 'auto_increment' then '1' else '0' end) as is_increment, column_type from information_schema.columns" +
		" where table_schema = (select database()) "
	sql += " and " + gdb.FormatSqlWithArgs(" table_name=? ", []interface{}{tableName}) + " order by ordinal_position ASC "
	err = db.GetScan(ctx, &res, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("查询列信息失败")
	}
	return res, nil
}

func (d *Mysql) GetTestSql() (sql string) {
	return "select 1 from dual"
}
