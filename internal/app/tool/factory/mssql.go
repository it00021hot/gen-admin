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

type Mssql struct{}

func (d *Mssql) DBPageList(ctx context.Context, in model.TablePageInput, db gdb.DB) (total int, list []*entity.GenTable, err error) {
	sql := `FROM SYSOBJECTS D
         	inner JOIN SYS.EXTENDED_PROPERTIES F ON D.ID = F.MAJOR_ID
    		AND F.MINOR_ID = 0 AND D.XTYPE = 'U' AND D.NAME != 'DTPROPERTIES'`
	if !reflect.DeepEqual(in, model.TablePageInput{}) {
		if in.TableName != "" {
			sql += gdb.FormatSqlWithArgs(" and cast(D.NAME as nvarchar) like ?", []interface{}{"%" + in.TableName + "%"})
		}
		if in.TableComment != "" {
			sql += gdb.FormatSqlWithArgs(" and cast(F.VALUE as nvarchar) like ?", []interface{}{"%" + in.TableComment + "%"})
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
	sql = "SELECT cast(D.NAME as nvarchar)  as table_name, cast(F.VALUE as nvarchar) as table_comment, crdate as create_time, refdate as update_time " + sql + " order by create_time desc"
	sql += gdb.FormatSqlWithArgs(" OFFSET 0 ROWS FETCH NEXT 10 ROWS ONLY ", []interface{}{(in.PageNum - 1) * in.PageSize, in.PageSize})
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("读取数据失败")
	}
	return

}

func (d *Mssql) SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq, db gdb.DB) (list []*entity.GenTable, err error) {
	sql := `SELECT cast(D.NAME as nvarchar)  as table_name,
       		cast(F.VALUE as nvarchar) as table_comment,
       		crdate                    as create_time,
       		refdate                   as update_time
			FROM SYSOBJECTS D
         	inner JOIN SYS.EXTENDED_PROPERTIES F ON D.ID = F.MAJOR_ID
    		AND F.MINOR_ID = 0 AND D.XTYPE = 'U' AND D.NAME != 'DTPROPERTIES'`
	if len(in.Tables) > 0 {
		tbs := gstr.TrimRight(gstr.Repeat("?,", len(in.Tables)), ",")
		sql += gdb.FormatSqlWithArgs(" and cast(D.NAME as nvarchar) in ("+tbs+")", gconv.SliceAny(in.Tables))
	}
	err = db.GetScan(ctx, &list, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("获取表格信息失败")
	}
	return
}

func (d *Mssql) SelectDbTableColumnsByName(ctx context.Context, tableName string, db gdb.DB) (res []*entity.GenTableColumn, err error) {
	sql := `SELECT cast(A.NAME as nvarchar)                                                     as column_name,
       		cast(B.NAME as nvarchar) + (case
                                       when B.NAME = 'numeric'
                                           then '(' + cast(A.prec as nvarchar) + ',' + cast(A.scale as nvarchar) + ')'
                                       else '' end)                                 as column_type,
       		cast(G.[VALUE] as nvarchar)                                                  as column_comment,
       		(SELECT 1
        	FROM INFORMATION_SCHEMA.KEY_COLUMN_USAGE Z
        	WHERE TABLE_NAME = D.NAME
          	and A.NAME = Z.column_name)                                               as is_pk,
       		(case when COLUMNPROPERTY(a.id, a.name, 'IsIdentity') = 1 then 1 else 0 end) as is_increment,
       		(case when A.isnullable = 0 then 1 else 0 end)                               as is_required,
       		a.colorder                                                                   as sort
			FROM SYS.SYSCOLUMNS A
         	LEFT JOIN SYS.SYSTYPES B ON A.XTYPE = B.XUSERTYPE
         	INNER JOIN SYS.SYSOBJECTS D ON A.ID = D.ID AND D.XTYPE = 'U' AND D.NAME != 'DTPROPERTIES'
         	LEFT JOIN SYS.EXTENDED_PROPERTIES G ON A.ID = G.MAJOR_ID AND A.COLID = G.MINOR_ID
         	LEFT JOIN SYS.EXTENDED_PROPERTIES F ON D.ID = F.MAJOR_ID AND F.MINOR_ID = 0
			where D.NAME = ?
			ORDER BY A.COLORDER`
	sql = gdb.FormatSqlWithArgs(sql, []interface{}{tableName})
	err = db.GetScan(ctx, &res, sql)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, gerror.New("查询列信息失败")
	}
	return res, nil
}

func (d *Mssql) GetTestSql() (sql string) {
	return "select 1"
}
