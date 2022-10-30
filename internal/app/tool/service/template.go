package service

import (
	"context"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
)

type ITemplate interface {
	List(ctx context.Context, name string) (templates []model.TemplateItem, err error)
	GetTplTree(ctx context.Context, name string, isDir bool) (tree []*model.FileTree, err error)
	Tree(ctx context.Context, name string, isDir bool) (tree []*model.FileTree, err error)
	AddTpl(ctx context.Context, name string) (err error)
	AddFile(tplModel *model.AddFileModel) (fileInfo *model.AddFileModel, err error)
	Rename(path string, name string) error
	DelTpl(name string) (err error)
	DelFile(path string) (err error)
	GetContent(path string) (result model.GetContentModel, err error)
	SaveContent(path string, content string) error
	GetSettings(name string) (content map[string]interface{}, err error)
	SaveSettings(name string, content map[string]interface{}) (err error)
}

var localTemplate ITemplate

func Template() ITemplate {
	if localTemplate == nil {
		panic("implement not found for interface ITemplate, forgot register?")
	}
	return localTemplate
}

func RegisterTemplate(i ITemplate) {
	localTemplate = i
}
