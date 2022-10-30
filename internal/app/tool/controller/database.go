package controller

import (
	"context"
	"github.com/gogf/gf/v2/util/gconv"
	v1 "github.com/it00021hot/gen-admin/api/v1/tool"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
)

var Database = cDatabase{}

type cDatabase struct{}

func (c *cDatabase) PageList(ctx context.Context, req *v1.DBPageListReq) (res *v1.DBPageListRes, err error) {
	total, list, err := service.Database().PageList(ctx, model.DatabasePageInput{
		PageReq: req.PageReq,
		Group:   req.Group,
		Name:    req.Name,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.DBPageListRes{
		Total: total,
		Items: list,
	}
	return
}

func (c *cDatabase) List(ctx context.Context, req *v1.DBListReq) (res *v1.DBListRes, err error) {
	list, err := service.Database().List(ctx, req.Group, req.Name)
	if err != nil {
		return nil, err
	}
	res = &v1.DBListRes{
		Items: list,
	}
	return
}

func (c *cDatabase) Test(ctx context.Context, req *v1.DBTestReq) (res *v1.EmptyRes, err error) {
	in := new(entity.GenDatabase)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Database().Test(ctx, *in)
	return
}

func (c *cDatabase) EditTest(ctx context.Context, req *v1.DBEditTestReq) (res *v1.EmptyRes, err error) {
	in := new(entity.GenDatabase)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Database().EditTest(ctx, *in)
	return
}

func (c *cDatabase) Get(ctx context.Context, req *v1.DBGetReq) (res *v1.DBGetRes, err error) {
	dbInfo, err := service.Database().Get(ctx, req.Id)
	res = &v1.DBGetRes{
		GenDatabase: dbInfo,
	}
	return
}

func (c *cDatabase) Add(ctx context.Context, req *v1.DBAddReq) (res *v1.EmptyRes, err error) {
	in := new(entity.GenDatabase)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Database().Add(ctx, *in)
	return
}

func (c *cDatabase) Edit(ctx context.Context, req *v1.DBEditReq) (res *v1.EmptyRes, err error) {
	in := new(entity.GenDatabase)
	err = gconv.Struct(req, in)
	if err != nil {
		return nil, err
	}
	err = service.Database().Edit(ctx, *in)
	return
}

func (c *cDatabase) Delete(ctx context.Context, req *v1.DBDelReq) (res *v1.EmptyRes, err error) {
	err = service.Database().Delete(ctx, req.Ids)
	return
}
