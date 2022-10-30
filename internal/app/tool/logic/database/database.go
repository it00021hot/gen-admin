package database

import (
	"context"
	"github.com/gogf/gf/v2/database/gdb"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/it00021hot/gen-admin/internal/app/tool/dao"
	"github.com/it00021hot/gen-admin/internal/app/tool/factory"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
	"reflect"
)

type (
	// sDatabase is service struct of module Database.
	sDatabase struct{}
)

func init() {
	service.RegisterDatabase(New())
}

func New() *sDatabase {
	return &sDatabase{}
}

// PageList creates user account.
func (s *sDatabase) PageList(ctx context.Context, param model.DatabasePageInput) (total int, databases []*entity.GenDatabase, err error) {
	columns := dao.GenDatabase.Columns()
	daoModel := dao.GenDatabase.Ctx(ctx).FieldsEx(columns.Pass)
	if !reflect.DeepEqual(param, model.DatabasePageInput{}) {
		if param.Group != "" {
			daoModel = daoModel.WhereLike(columns.Group, "%"+param.Group+"%")
		}
		if param.Name != "" {
			daoModel = daoModel.WhereLike(columns.Name, "%"+param.Name+"%")
		}
		if !param.BeginTime.IsZero() && !param.EndTime.IsZero() {
			daoModel = daoModel.WhereBetween(columns.CreateTime, param.BeginTime, param.EndTime.EndOfDay())
		}
	}
	total, err = daoModel.Count()
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取总行数失败")
		return
	}
	if total == 0 {
		databases = make([]*entity.GenDatabase, 0)
		return total, databases, nil
	}
	err = daoModel.Page(param.PageNum, param.PageSize).Scan(&databases)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	return
}

func (s *sDatabase) Add(ctx context.Context, in entity.GenDatabase) (err error) {
	columns := dao.GenDatabase.Columns()
	if in.Group == gdb.DefaultGroupName {
		return gerror.New(in.Group + "已存在")
	}
	count, err := dao.GenDatabase.Ctx(ctx).Count(columns.Group, in.Group)
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(in.Group + "已存在")
	}
	_, err = dao.GenDatabase.Ctx(ctx).Save(in)
	if err != nil {
		return err
	}
	return
}

func (s *sDatabase) Delete(ctx context.Context, ids []int) (err error) {
	for _, id := range ids {
		if id == 0 {
			return gerror.New("默认数据源不可删除")
		}
	}
	_, err = dao.GenDatabase.Ctx(ctx).Delete(dao.GenDatabase.Columns().Id, ids)
	if err != nil {
		return err
	}
	return
}

func (s *sDatabase) List(ctx context.Context, group string, name string) (databases []*entity.GenDatabase, err error) {
	columns := dao.GenDatabase.Columns()
	daoModel := dao.GenDatabase.Ctx(ctx).FieldsEx(columns.Pass)

	if group != "" {
		daoModel = daoModel.WhereLike(columns.Group, "%"+group+"%")
	}
	if name != "" {
		daoModel = daoModel.WhereLike(columns.Name, "%"+name+"%")
	}
	err = daoModel.Scan(&databases)
	if err != nil {
		g.Log().Error(ctx, err)
		err = gerror.New("获取数据失败")
	}
	if err != nil {
		return
	}
	config := g.DB().GetConfig()
	defaultDB := &entity.GenDatabase{
		Id:     0,
		Group:  gdb.DefaultGroupName,
		Host:   config.Host,
		Port:   config.Port,
		User:   config.User,
		Type:   config.Type,
		Name:   config.Name,
		Remark: "默认数据源",
	}
	databases = append(databases, defaultDB)

	return
}

func (s *sDatabase) Test(ctx context.Context, in entity.GenDatabase) (err error) {
	columns := dao.GenDatabase.Columns()
	if in.Group == gdb.DefaultGroupName {
		return gerror.New("不能使用默认分组名")
	}
	count, err := dao.GenDatabase.Ctx(ctx).Count(columns.Group, in.Group)
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(in.Group + "已存在")
	}
	gdb.SetConfigGroup(in.Group, gdb.ConfigGroup{
		gdb.ConfigNode{
			Host:   in.Host,
			Port:   in.Port,
			User:   in.User,
			Pass:   in.Pass,
			Name:   in.Name,
			Type:   in.Type,
			Debug:  true,
			DryRun: true,
		},
	})
	db, err := gdb.Instance(in.Group)
	_, err = db.Query(ctx, factory.New(db.GetConfig().Type).GetTestSql())
	if err != nil {
		return gerror.New("连接失败!")
	}
	return nil
}

func (s *sDatabase) EditTest(ctx context.Context, in entity.GenDatabase) (err error) {
	columns := dao.GenDatabase.Columns()
	if in.Group == gdb.DefaultGroupName {
		return gerror.New("不能使用默认分组名")
	}
	var dbInfo *entity.GenDatabase
	err = dao.GenDatabase.Ctx(ctx).Where(g.Map{
		columns.Id + "!= ?": in.Id,
		columns.Group:       in.Group,
	}).Scan(&dbInfo)
	if in.Pass == "" {
		in.Pass = dbInfo.Pass
	}
	gdb.SetConfigGroup(in.Group, gdb.ConfigGroup{
		gdb.ConfigNode{
			Host:   in.Host,
			Port:   in.Port,
			User:   in.User,
			Pass:   in.Pass,
			Name:   in.Name,
			Type:   in.Type,
			Debug:  true,
			DryRun: true,
		},
	})
	db, err := gdb.Instance(in.Group)
	_, err = db.Query(ctx, factory.New(db.GetConfig().Type).GetTestSql())
	if err != nil {
		return gerror.New("连接失败!")
	}
	return nil
}

func (s *sDatabase) Edit(ctx context.Context, in entity.GenDatabase) (err error) {
	columns := dao.GenDatabase.Columns()
	if in.Group == gdb.DefaultGroupName {
		return gerror.New(in.Group + "已存在")
	}
	count, err := dao.GenDatabase.Ctx(ctx).Where(g.Map{
		columns.Id + "!= ?": in.Id,
		columns.Group:       in.Group,
	}).Count()
	if err != nil {
		return err
	}
	if count > 0 {
		return gerror.New(in.Group + "已存在")
	}
	_, err = dao.GenDatabase.Ctx(ctx).Where(columns.Id, in.Id).Data(in).Update()
	if err != nil {
		return err
	}
	return
}

func (s *sDatabase) Get(ctx context.Context, id int) (result entity.GenDatabase, err error) {
	err = dao.GenDatabase.Ctx(ctx).FieldsEx(dao.GenDatabase.Columns().Pass).Where(id).Scan(&result)
	return
}
