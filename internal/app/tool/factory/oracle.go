package factory

import (
	"context"
	"fmt"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"

	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"reflect"
)

type Oracle struct{}

func (d *Oracle) DBPageList(ctx context.Context, in model.TablePageInput, db gdb.DB) (total int, list []*entity.GenTable, err error) {
	sql := " from user_tables dt,user_tab_comments dtc,user_objects uo where dt.table_name = dtc.table_name and dt.table_name = uo.object_name and uo.object_type = 'TABLE' "
	if !reflect.DeepEqual(in, model.TablePageInput{}) {
		if in.TableName != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(dt.table_name) like lower(?)", []interface{}{"%" + in.TableName + "%"})
		}

		if in.TableComment != "" {
			sql += gdb.FormatSqlWithArgs(" and lower(dtc.comments) like lower(?)", []interface{}{"%" + in.TableComment + "%"})
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
	sql = "SELECT lower(dt.table_name) as table_name,dtc.comments as table_comment,uo.created as create_time,uo.last_ddl_time as update_time " + sql
	sql += " order by uo.created desc"
	baseSql := fmt.Sprintf(`SELECT table_name,table_comment,create_time,update_time FROM (SELECT ROWNUM RN, tb.* FROM (%s) tb WHERE ROWNUM <= %d * %d)
	WHERE RN > (%d - 1) * %d`, sql, in.PageSize, in.PageNum, in.PageNum, in.PageSize)
	err = db.GetScan(ctx, &list, baseSql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取数据失败")
	}
	return

}

func (d *Oracle) SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq, db gdb.DB) (list []*entity.GenTable, err error) {
	sql := `select lower(dt.table_name) as table_name,
	dtc.comments         as table_comment,
	uo.created           as create_time,
	uo.last_ddl_time     as update_time
	from user_tables dt,
	user_tab_comments dtc,
	user_objects uo
	where dt.table_name = dtc.table_name
  	and dt.table_name = uo.object_name
  	and uo.object_type = 'TABLE'`
	if len(in.Tables) > 0 {
		tbs := gstr.TrimRight(gstr.Repeat("?,", len(in.Tables)), ",")
		sql += " and " + gdb.FormatSqlWithArgs("lower(dt.table_name) in ("+tbs+")", gconv.SliceAny(in.Tables))
	}
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取表格信息失败")
	}
	return
}

func (d *Oracle) SelectDbTableColumnsByName(ctx context.Context, tableName string, db gdb.DB) (res []*entity.GenTableColumn, err error) {
	sql := `select lower(temp.column_name)                                                      as column_name,
       lower(temp.data_type)                                                                    as column_type,
       (case when (temp.nullable = 'N' and temp.constraint_type != 'P') then '1' else null end) as is_required,
       (case when temp.constraint_type = 'P' then '1' else '0' end)                             as is_pk,
       temp.column_id                                                                           as sort,
       temp.comments                                                                            as column_comment,
       (case when temp.constraint_type = 'P' then '1' else '0' end)                             as is_increment
		from (
         select col.column_id
              , col.column_name
              , col.nullable
              , col.data_type
              , colc.comments
              , uc.constraint_type
              , row_number() over (partition by col.column_name order by uc.constraint_type desc) as row_flg
         from user_tab_columns col
                  left join user_col_comments colc
                            on colc.table_name = col.table_name and colc.column_name = col.column_name
                  left join user_cons_columns ucc
                            on ucc.table_name = col.table_name and ucc.column_name = col.column_name
                  left join user_constraints uc on uc.constraint_name = ucc.constraint_name
         where col.table_name = upper(?)
     ) temp
	WHERE temp.row_flg = 1
	ORDER BY temp.column_id`
	err = db.GetScan(ctx, &res, gdb.FormatSqlWithArgs(sql, []interface{}{tableName}))
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("查询列信息失败")
	}
	return res, nil
}

func (d *Oracle) GetTestSql() (sql string) {
	return "select 1 from DUAL"
}
