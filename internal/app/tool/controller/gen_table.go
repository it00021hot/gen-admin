package controller

import (
	"context"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	v1 "github.com/it00021hot/gen-admin/api/v1/tool"
	"github.com/it00021hot/gen-admin/internal/app/tool/consts"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
	utils "github.com/it00021hot/gen-admin/internal/app/tool/util"

	"os"
)

var Table = cTable{}

type cTable struct{}

func (c *cTable) PageList(ctx context.Context, req *v1.TablePageListReq) (res *v1.TableDBListRes, err error) {
	total, list, err := service.Table().PageList(ctx, model.TablePageInput{
		PageReq:      req.PageReq,
		TableName:    req.TableName,
		TableComment: req.TableComment,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.TableDBListRes{
		Total: total,
		Items: list,
	}
	return
}

func (c *cTable) DBPageList(ctx context.Context, req *v1.TableDBListReq) (res *v1.TableDBListRes, err error) {
	total, list, err := service.Table().DBPageList(ctx, model.TablePageInput{
		PageReq:      req.PageReq,
		Group:        req.Group,
		TableName:    req.TableName,
		TableComment: req.TableComment,
	})
	if err != nil {
		return nil, err
	}
	res = &v1.TableDBListRes{
		Total: total,
		Items: list,
	}
	return
}

func (c *cTable) Create(ctx context.Context, req *v1.TableCreateReq) (res *v1.EmptyRes, err error) {
	err = service.Table().Create(ctx, req.Sql)
	return
}

func (c *cTable) ImportSave(ctx context.Context, req *v1.TableImportSaveReq) (res *v1.EmptyRes, err error) {
	tableList, err := service.Table().SelectDbTableListByNames(ctx, model.TableImportSaveReq{
		Group:  req.Group,
		Tables: req.Tables,
	})
	if err != nil {
		return
	}
	if tableList == nil {
		return nil, gerror.New("表信息不存在")
	}
	err = service.Table().ImportGenTable(ctx, req.Group, tableList)
	return
}

func (c *cTable) EditTable(ctx context.Context, req *v1.TableGetReq) (res *v1.TableGetRes, err error) {
	if req.TableId == 0 {
		return nil, gerror.New("参数错误")
	}
	table, err := service.Table().SelectGenTableById(ctx, req.TableId)
	if err != nil {
		return
	}
	genTableList, err := service.Table().SelectGenTableAll(ctx)
	res = &v1.TableGetRes{
		Info:   table,
		Tables: genTableList,
		Items:  table.Columns,
		Total:  len(table.Columns),
	}
	return
}

func (c *cTable) EditSave(ctx context.Context, req *v1.TableEditReq) (res *v1.EmptyRes, err error) {
	err = service.Table().SaveEdit(ctx, req.GenTableDTO)
	return
}

func (c *cTable) Delete(ctx context.Context, req *v1.TableDelReq) (res *v1.EmptyRes, err error) {
	err = service.Table().Delete(ctx, req.Ids)
	return
}

func (c *cTable) Preview(ctx context.Context, req *v1.TablePreviewReq) (res *v1.TablePreviewRes, err error) {
	tableId := req.TableId
	if tableId == 0 {
		return nil, gerror.New("参数错误")
	}
	data, err := service.Table().PreviewCode(ctx, tableId, req.TplList)
	if err != nil {
		return nil, err
	}
	res = &v1.TablePreviewRes{
		Data: data,
	}
	return
}

func (c *cTable) genCode(fileList []string) (bs []byte, err error) {
	defer func() { // 移除中间文件
		if err := os.RemoveAll(consts.AutoPath); err != nil {
			return
		}
	}()
	if bs, err = utils.ZipByte("genCode.zip", fileList, consts.AutoPath, ""); err != nil {
		return nil, err
	}
	return
}

func (c *cTable) BatchGenCode(ctx context.Context, req *v1.TableGenCodeReq) (res *v1.TableGenCodeRes, err error) {
	ids := req.Ids
	if len(ids) == 0 {
		return nil, gerror.New("参数错误")
	}
	fileList, err := service.Table().BatchGenCode(ctx, ids, req.TplList)
	bs, err := c.genCode(fileList)
	if err != nil {
		return nil, err
	}
	g.RequestFromCtx(ctx).Response.Write(bs)
	//g.RequestFromCtx(ctx).Response.ServeFileDownload("genCode.zip")
	_ = os.Remove("./genCode.zip")
	return
}
