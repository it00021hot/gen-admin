package service

import (
	"context"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type IDatabase interface {
	PageList(ctx context.Context, param model.DatabasePageInput) (total int, databases []*entity.GenDatabase, err error)
	List(ctx context.Context, group string, name string) (databases []*entity.GenDatabase, err error)
	Get(ctx context.Context, id int) (result entity.GenDatabase, err error)
	Add(ctx context.Context, in entity.GenDatabase) (err error)
	Edit(ctx context.Context, in entity.GenDatabase) (err error)
	Delete(ctx context.Context, ids []int) (err error)
	Test(ctx context.Context, in entity.GenDatabase) (err error)
	EditTest(ctx context.Context, in entity.GenDatabase) (err error)
}

var localDatabase IDatabase

func Database() IDatabase {
	if localDatabase == nil {
		panic("implement not found for interface IDatabase, forgot register?")
	}
	return localDatabase
}

func RegisterDatabase(i IDatabase) {
	localDatabase = i
}
