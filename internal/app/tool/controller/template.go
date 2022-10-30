package controller

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/it00021hot/gen-admin/api/v1/tool"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
)

var Template = cTemplate{}

type cTemplate struct{}

func (c *cTemplate) List(ctx context.Context, req *v1.TplListReq) (res *v1.TplListRes, err error) {
	list, err := service.Template().List(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	res = &v1.TplListRes{
		Total: len(list),
		Items: list,
	}
	return
}

func (c *cTemplate) GetTree(ctx context.Context, req *v1.TplTreeReq) (res *v1.TplTreeRes, err error) {
	tree, err := service.Template().Tree(ctx, "", false)
	if err != nil {
		return nil, err
	}
	res = &v1.TplTreeRes{
		Items: tree,
	}
	return
}

func (c *cTemplate) GetTplTree(ctx context.Context, req *v1.GetTplTreeReq) (res *v1.TplTreeRes, err error) {
	tree, err := service.Template().GetTplTree(ctx, req.Name, req.IsDir)
	if err != nil {
		return nil, err
	}
	res = &v1.TplTreeRes{
		Items: tree,
	}
	return
}

func (c *cTemplate) AddTpl(ctx context.Context, req *v1.AddTplReq) (res *v1.AddTplRes, err error) {
	err = service.Template().AddTpl(ctx, req.Name)
	if err != nil {
		return nil, err
	}
	res = &v1.AddTplRes{
		Name: req.Name,
		Path: req.Name,
	}
	return
}

func (c *cTemplate) AddFile(ctx context.Context, req *v1.AddFileReq) (res *v1.AddFileRes, err error) {
	fileInfo, err := service.Template().AddFile(req.AddFileModel)
	if err != nil {
		return nil, err
	}
	res = &v1.AddFileRes{AddFileModel: fileInfo}
	return
}

func (c *cTemplate) GetSettings(ctx context.Context, req *v1.GetSettingsReq) (res *v1.GetSettingsRes, err error) {
	result, err := service.Template().GetSettings(req.Name)
	res = &v1.GetSettingsRes{
		Name:    req.Name,
		Content: result,
	}
	return
}

func (c *cTemplate) SaveSettings(ctx context.Context, req *v1.SaveSettingsReq) (res *v1.EmptyRes, err error) {
	err = service.Template().SaveSettings(req.Name, req.Content)
	return
}

func (c *cTemplate) Rename(ctx context.Context, req *v1.RenameReq) (res *v1.EmptyRes, err error) {
	err = service.Template().Rename(req.Path, req.NewName)
	return
}

func (c *cTemplate) DelTpl(ctx context.Context, req *v1.DelTplReq) (res *v1.EmptyRes, err error) {
	g.Log().Info(ctx, "删除模板: ", req.Name)
	err = service.Template().DelTpl(req.Name)
	return
}

func (c *cTemplate) DelFile(ctx context.Context, req *v1.DelFileReq) (res *v1.EmptyRes, err error) {
	g.Log().Info(ctx, "删除文件(夹): ", req.Path)
	err = service.Template().DelFile(req.Path)
	return
}

func (c *cTemplate) GetContent(ctx context.Context, req *v1.GetContentReq) (res *v1.GetContentRes, err error) {
	result, err := service.Template().GetContent(req.Path)
	if err != nil {
		return nil, err
	}
	res = &v1.GetContentRes{
		GetContentModel: result,
	}
	return
}

func (c *cTemplate) SaveContent(ctx context.Context, req *v1.SaveContentReq) (res *v1.EmptyRes, err error) {
	err = service.Template().SaveContent(req.Path, req.Content)
	return
}
