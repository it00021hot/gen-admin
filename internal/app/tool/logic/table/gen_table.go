package table

import (
	"context"
	"fmt"
	"github.com/blastrain/vitess-sqlparser/sqlparser"
	"github.com/blastrain/vitess-sqlparser/tidbparser/ast"
	"github.com/blastrain/vitess-sqlparser/tidbparser/parser"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/it00021hot/gen-admin/internal/app/tool/consts"
	"github.com/it00021hot/gen-admin/internal/app/tool/dao"
	"github.com/it00021hot/gen-admin/internal/app/tool/factory"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/do"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
	utils "github.com/it00021hot/gen-admin/internal/app/tool/util"
	"os"
	"reflect"
	"strings"
)

func init() {
	service.RegisterTable(New())
}

type (
	// sDatabase is service struct of module Table.
	sTable struct{}
)

// New returns the interface of Table service.
func New() *sTable {
	return &sTable{}
}

func (s *sTable) PageList(ctx context.Context, in model.TablePageInput) (total int, list []*entity.GenTable, err error) {
	daoModel := dao.GenTable.Ctx(ctx)
	if err != nil {
		return 0, nil, err
	}
	if !reflect.DeepEqual(in, model.TablePageInput{}) {
		if in.TableName != "" {
			daoModel = daoModel.WhereLike(dao.GenTable.Columns().TableName, "%"+in.TableName+"%")
		}
		if in.TableComment != "" {
			daoModel = daoModel.WhereLike(dao.GenTable.Columns().TableComment, "%"+in.TableComment+"%")
		}
		if !in.BeginTime.IsZero() && !in.EndTime.IsZero() {
			daoModel = daoModel.WhereBetween(dao.GenTable.Columns().CreateTime, in.BeginTime, in.EndTime.EndOfDay())
		}
		total, err = daoModel.Count()
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取总行数失败")
			return
		}
		if total == 0 {
			list = make([]*entity.GenTable, 0)
			return total, list, nil
		}
		err = daoModel.Page(in.PageNum, in.PageSize).OrderAsc(dao.GenTable.Columns().TableId).Scan(&list)
		if err != nil {
			g.Log().Error(ctx, err)
			err = gerror.New("获取数据失败")
		}
	}
	return
}

func (s *sTable) DBPageList(ctx context.Context, in model.TablePageInput) (total int, list []*entity.GenTable, err error) {
	db := g.DB()
	var idbType factory.IDBType
	if in.Group != gdb.DefaultGroupName {
		var dbConfig *entity.GenDatabase
		err := dao.GenDatabase.Ctx(ctx).Where(dao.GenDatabase.Columns().Group, in.Group).Scan(&dbConfig)
		if err != nil {
			return 0, nil, err
		}
		gdb.SetConfigGroup(in.Group, gdb.ConfigGroup{gdb.ConfigNode{
			Host:   dbConfig.Host,
			Port:   dbConfig.Port,
			User:   dbConfig.User,
			Pass:   dbConfig.Pass,
			Name:   dbConfig.Name,
			Type:   dbConfig.Type,
			Debug:  true,
			DryRun: true,
		}})
		db, err = gdb.Instance(in.Group)
		if err != nil {
			return 0, nil, err
		}
	}
	idbType = factory.New(db.GetConfig().Type)
	return idbType.DBPageList(ctx, in, db)
}

func (s *sTable) Create(ctx context.Context, sql string) (err error) {
	parse, err := parser.New().Parse(sql, "", "")
	if err != nil {
		return
	}
	for _, stmtNode := range parse {
		switch stmtNode := stmtNode.(type) {
		case *ast.CreateTableStmt:
			err := s.create(ctx, stmtNode.Text())
			if err != nil {
				return err
			}
		}
	}
	return
}

func (s *sTable) create(ctx context.Context, sql string) (err error) {
	stmt, err := sqlparser.Parse(sql)
	if err != nil {
		return gerror.New("请输入建表语句")
	}
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return err
	}
	switch stmt := stmt.(type) {
	case *sqlparser.CreateTable:
		tableName := stmt.NewName.Name.String()
		count, err := tx.Model(dao.GenTable.Table()).Where(dao.GenTable.Columns().TableName, tableName).Count()
		if err != nil {
			return err
		}
		if count > 0 {
			return gerror.New(fmt.Sprintf("表%s已存在！", tableName))
		}
		var createTable = &entity.GenTable{
			TableName: tableName,
		}
		for _, option := range stmt.Options {
			if option.Type == sqlparser.TableOptionComment {
				createTable.TableComment = option.StrValue
			}
		}
		s.InitTable(ctx, createTable)
		id, err := tx.Model(dao.GenTable.Table()).FieldsEx(dao.GenTable.Columns().TableId).InsertAndGetId(createTable)
		if err != nil {
			return err
		}
		createTable.TableId = id
		var columns = make([]entity.GenTableColumn, 0)
		for _, column := range stmt.Columns {
			var genTableColumn = &entity.GenTableColumn{
				ColumnName: column.Name,
				ColumnType: column.Type,
				CreateTime: gtime.Now(),
			}
			for _, option := range column.Options {

				switch option.Type {
				case sqlparser.ColumnOptionPrimaryKey:
					genTableColumn.IsPk = true
				case sqlparser.ColumnOptionComment:
					genTableColumn.ColumnComment = strings.Trim(option.Value, "\"")
				}
			}
			service.TableColumn().InitColumnField(genTableColumn, createTable)
			columns = append(columns, *genTableColumn)
		}

		_, err = tx.Model(dao.GenTableColumn.Table()).FieldsEx(dao.GenTableColumn.Columns().ColumnId).Insert(columns)
		if err != nil {
			_ = tx.Rollback()
			return gerror.New("保存列数据失败")
		}
	default:
		return gerror.New("请输入建表语句")
	}
	return tx.Commit()
}

// InitTable 初始化表信息
func (s *sTable) InitTable(ctx context.Context, table *entity.GenTable) {
	table.ClassName = s.ConvertClassName(ctx, table.TableName)
	table.BusinessName = s.GetBusinessName(ctx, table.TableName)
	table.FunctionName = table.TableComment
	table.GenType = "0"
	table.GenPath = "/"
	table.TplCategory = "crud"
	table.CreateTime = gtime.Now()
	table.UpdateTime = table.CreateTime
}

// InitGenInfo 初始化生成信息
func (s *sTable) InitGenInfo(table *model.GenTableDTO, genMap map[string]*g.Var) {
	if table.SystemName == "" {
		table.SystemName = genMap["systemName"].String()

	}
	if table.ModuleName == "" {
		table.ModuleName = genMap["moduleName"].String()

	}
	if table.PackageName == "" {
		table.PackageName = genMap["packageName"].String()

	}
	if table.FunctionAuthor == "" {
		table.FunctionAuthor = genMap["author"].String()
	}
	table.PublicFields = genMap["publicFields"].Strings()

}

// ConvertClassName 表名转换成类名
func (s *sTable) ConvertClassName(ctx context.Context, tableName string) string {
	return gstr.CaseCamel(s.RemoveTablePrefix(ctx, tableName))
}

// GetBusinessName 获取业务名
func (s *sTable) GetBusinessName(ctx context.Context, tableName string) string {
	return s.RemoveTablePrefix(ctx, tableName)
}

// RemoveTablePrefix 删除表前缀
func (s *sTable) RemoveTablePrefix(ctx context.Context, tableName string) string {
	autoRemovePre, _ := g.Cfg().Get(ctx, "gen.autoRemovePre")
	tablePrefix, _ := g.Cfg().Get(ctx, "gen.tablePrefix")
	if autoRemovePre.Bool() && tablePrefix.String() != "" {
		searchList := strings.Split(tablePrefix.String(), ",")
		for _, str := range searchList {
			if strings.HasPrefix(tableName, str) {
				tableName = strings.Replace(tableName, str, "", 1) //注意，只替换一次
			}
		}
	}
	return tableName
}

func (s *sTable) SelectDbTableListByNames(ctx context.Context, in model.TableImportSaveReq) (list []*entity.GenTable, err error) {
	db := g.DB()
	var idbType factory.IDBType
	if in.Group != gdb.DefaultGroupName {
		var dbConfig *entity.GenDatabase
		err = dao.GenDatabase.Ctx(ctx).Where(dao.GenDatabase.Columns().Group, in.Group).Scan(&dbConfig)
		if err != nil {
			return
		}
		gdb.SetConfigGroup(in.Group, gdb.ConfigGroup{gdb.ConfigNode{
			Host:   dbConfig.Host,
			Port:   dbConfig.Port,
			User:   dbConfig.User,
			Pass:   dbConfig.Pass,
			Name:   dbConfig.Name,
			Type:   dbConfig.Type,
			Debug:  true,
			DryRun: true,
		}})
		db, err = gdb.Instance(in.Group)
		if err != nil {
			return
		}
	}
	idbType = factory.New(db.GetConfig().Type)
	return idbType.SelectDbTableListByNames(ctx, in, db)
}

func (s *sTable) ImportGenTable(ctx context.Context, group string, tableList []*entity.GenTable) (err error) {
	if tableList != nil {
		tx, err := g.DB().Begin(ctx)
		if err != nil {
			return err
		}
		for _, importTable := range tableList {
			tableName := importTable.TableName
			count, err := tx.Model(dao.GenTable.Table()).Where(dao.GenTable.Columns().TableName, tableName).Count()
			if err != nil || count > 0 {
				return gerror.Newf("%s已存在", tableName)
			}
			s.InitTable(ctx, importTable)
			tmpId, err := tx.Model(dao.GenTable.Table()).FieldsEx(dao.GenTable.Columns().TableId).InsertAndGetId(importTable)
			if err != nil || tmpId <= 0 {
				err = tx.Rollback()
				return gerror.New("保存数据失败")
			}

			importTable.TableId = tmpId

			// 保存列信息
			genTableColumns, err := service.TableColumn().SelectDbTableColumnsByName(ctx, group, tableName)

			if err != nil || len(genTableColumns) <= 0 {
				_ = tx.Rollback()
				return gerror.New("获取列数据失败")
			}
			for _, column := range genTableColumns {
				service.TableColumn().InitColumnField(column, importTable)
				_, err = tx.Model(dao.GenTableColumn.Table()).FieldsEx(dao.GenTableColumn.Columns().ColumnId).Insert(column)
				if err != nil {
					err = tx.Rollback()
					if err != nil {
						return err
					}
					return gerror.New("保存列数据失败")
				}
			}
		}
		return tx.Commit()
	} else {
		return gerror.New("参数错误")
	}
}

func (s *sTable) SelectGenTableById(ctx context.Context, tableId int) (table *model.GenTableDTO, err error) {
	err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableId, tableId).Scan(&table)
	if err != nil {
		return nil, err
	}
	var columns []*entity.GenTableColumn
	err = dao.GenTableColumn.Ctx(ctx).
		Where(dao.GenTableColumn.Columns().TableId, tableId).Scan(&columns)
	if err != nil {
		return nil, err
	}
	table.Columns = columns
	return
}

func (s *sTable) SelectGenTableAll(ctx context.Context) (table []*model.GenTableDTO, err error) {
	err = dao.GenTable.Ctx(ctx).Scan(&table)
	if err != nil {
		return nil, err
	}
	var columns []*entity.GenTableColumn
	err = dao.GenTableColumn.Ctx(ctx).Scan(&columns)
	if err != nil {
		return nil, err
	}
	for _, tableDTO := range table {
		for _, column := range columns {
			if column.TableId == tableDTO.TableId {
				tableDTO.Columns = append(tableDTO.Columns, column)
			}
		}
	}
	return table, nil
}

func (s *sTable) SaveEdit(ctx context.Context, req *model.GenTableDTO) (err error) {
	if reflect.DeepEqual(req, model.TablePageInput{}) {
		err = gerror.New("参数错误")
		return
	}

	tableParams := new(do.GenTable)
	err = gconv.Scan(req, tableParams)
	if err != nil {
		return err
	}
	count, err := dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableId, req.TableId).Count()
	if err != nil || count < 1 {
		err = gerror.New("数据不存在")
		return
	}
	tableParams.UpdateTime = gtime.Now()

	// 开启事务
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		return
	}
	if req.TplCategory != consts.TplTree {
		tableParams.TreeCode = ""
		tableParams.TreeName = ""
		tableParams.TreeParentCode = ""

	}
	if req.TplCategory != consts.TplSub {
		tableParams.SubTableName = ""
		tableParams.SubTableFkName = ""
	}
	_, err = tx.Model(dao.GenTable.Table()).FieldsEx(dao.GenTable.Columns().TableId).Data(tableParams).Where(dao.GenTable.Columns().TableId, tableParams.TableId).Update()
	if err != nil {
		err = tx.Rollback()
		return
	}
	//保存列数据
	if req.Columns != nil {
		for _, column := range req.Columns {
			if column.ColumnId > 0 {
				if tc := req.TreeParentCode; tc != "" && tc == column.JavaField {
					column.IsQuery = false
					column.IsList = false
					column.HtmlType = "select"
				}
				_, err = tx.Model(dao.GenTableColumn.Table()).FieldsEx(dao.GenTableColumn.Columns().ColumnId).OmitNil().Data(column).Where(dao.GenTableColumn.Columns().ColumnId, column.ColumnId).Update()
				if err != nil {
					err = tx.Rollback()
					return
				}
			}
		}
	}
	err = tx.Commit()

	return
}

func (s *sTable) Delete(ctx context.Context, ids []int) (err error) {
	tx, err := g.DB().Begin(ctx)
	if err != nil {
		g.Log().Error(ctx, err)
		return gerror.New("开启删除事务出错")
	}
	_, err = tx.Model(dao.GenTable.Table()).WhereIn(dao.GenTable.Columns().TableId, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		err = tx.Rollback()
		return gerror.New("删除表格数据失败")
	}
	_, err = tx.Model(dao.GenTableColumn.Table()).WhereIn(dao.GenTableColumn.Columns().TableId, ids).Delete()
	if err != nil {
		g.Log().Error(ctx, err)
		err = tx.Rollback()
		return gerror.New("删除表格字段数据失败")
	}
	err = tx.Commit()
	return err
}

func (s *sTable) PreviewCode(ctx context.Context, tableId int, tplList []string) (data g.MapStrStr, err error) {
	previewTable, err := s.SelectGenTableById(ctx, tableId)
	if err != nil {
		return
	}
	data, err = s.generatorCode(ctx, previewTable, tplList)
	return

}

func (s *sTable) generatorCode(ctx context.Context, table *model.GenTableDTO, tplList []string) (data g.MapStrStr, err error) {
	data = map[string]string{}
	table.PrefixName = s.RemoveTablePrefix(ctx, table.TableName)
	//设置主子表信息
	s.setSubTable(ctx, table)
	//设置主键列信息
	s.setPkColumn(table)
	view := utils.InitView()
	//设置模板变量
	for _, tpl := range tplList {
		// 获取 BasePath 文件夹下所有tpl文件
		templates, err := utils.GetAllTplFile(fmt.Sprintf("%s/%s", consts.BasePath, tpl), table.TplCategory, nil)
		if err != nil {
			return nil, err
		}
		//渲染生成文件配置
		settingsStr := gfile.GetContents(fmt.Sprintf("%s/%s/%s", consts.BasePath, tpl, consts.Settings))

		settingsJson, err := gjson.LoadYaml(settingsStr)
		if err != nil {
			return nil, err
		}
		//设置生成信息
		s.InitGenInfo(table, settingsJson.Get("gen").MapStrVar())
		tplContext := utils.PrepareContext(table)
		settings, err := view.ParseContent(ctx, settingsStr, tplContext)
		if err != nil {
			return nil, err
		}
		fileYaml, err := gjson.LoadYaml(settings)
		if err != nil {
			return nil, err
		}
		fileMap := fileYaml.Get("file").MapStrStr()
		for _, templateName := range templates {
			fileName := utils.GetFileName(templateName, tpl, fileMap)
			if tmpDao, err := view.Parse(ctx, templateName, tplContext); err == nil {
				daoValue, _ := utils.TrimBreak(tmpDao)
				data[fileName] = daoValue
			} else {
				g.Log().Error(ctx, err)
			}
		}
	}
	if consts.TplSub == table.TplCategory {
		if table.SubTable.TableName != table.TableName {
			subData, err := s.generatorCode(ctx, table.SubTable, tplList)
			if err != nil {
				return nil, err
			}
			for template, code := range subData {
				data[template] = code
			}
		}

	}
	return
}

func (s *sTable) setSubTable(ctx context.Context, table *model.GenTableDTO) {
	subTableName := table.SubTableName
	if subTableName != "" && consts.TplSub == table.TplCategory {
		subTable, err := s.SelectGenTableByName(ctx, subTableName)
		if err != nil {
			return
		}
		table.SubTable = subTable
	}
}

func (s *sTable) setPkColumn(table *model.GenTableDTO) {
	for _, column := range table.Columns {
		if column.IsPk {
			table.PkColumn = column
			break
		}
	}
	if table.PkColumn == nil {
		table.PkColumn = table.Columns[0]
	}
	if consts.TplSub == table.TplCategory {
		for _, column := range table.SubTable.Columns {
			if column.IsPk {
				table.SubTable.PkColumn = column
				break
			}
		}
		if table.SubTable.PkColumn == nil {
			table.SubTable.PkColumn = table.Columns[0]
		}

	}
}

func (s *sTable) SelectGenTableByName(ctx context.Context, tableName string) (result *model.GenTableDTO, err error) {
	err = dao.GenTable.Ctx(ctx).Where(dao.GenTable.Columns().TableName, tableName).Scan(&result)
	if err != nil {
		return nil, err
	}
	var columns []*entity.GenTableColumn
	err = dao.GenTableColumn.Ctx(ctx).Where(dao.GenTableColumn.Columns().TableId, result.TableId).Scan(&columns)
	if err != nil {
		return nil, err
	}
	result.Columns = columns
	return result, nil

}

func (s *sTable) BatchGenCode(ctx context.Context, ids []int, tplList []string) (fileList []string, err error) {
	for _, id := range ids {
		files, err := s.GenCode(ctx, id, tplList)
		if err != nil {
			return nil, err
		}
		fileList = append(fileList, files...)
	}
	return
}

func (s *sTable) GenCode(ctx context.Context, id int, tplList []string) (fileList []string, err error) {
	table, err := s.SelectGenTableById(ctx, id)
	if err != nil {
		return nil, err
	}
	genData, err := s.generatorCode(ctx, table, tplList)
	if err != nil {
		return nil, err
	}
	for template, code := range genData {
		path := strings.Join([]string{consts.AutoPath, "/", template}, "")
		err = s.createFile(path, code, false)
		if err != nil {
			return nil, err
		}
		fileList = append(fileList, path)
	}
	return
}

// createFile 创建文件
func (s *sTable) createFile(fileName, data string, cover bool) (err error) {
	if !gfile.Exists(fileName) || cover {
		var f *os.File
		f, err = gfile.Create(fileName)
		if err != nil {
			return err
		}
		_, err = f.WriteString(data)
		if err != nil {
			return err
		}
		err = f.Close()
		if err != nil {
			return err
		}

	}
	return
}
