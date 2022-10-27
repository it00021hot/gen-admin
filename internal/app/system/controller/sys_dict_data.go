
package controller

import (
"context"
"github.com/it00021hot/gen-admin/api/v1/system"
commonService "github.com/it00021hot/gen-admin/internal/app/common/service"
"github.com/it00021hot/gen-admin/internal/app/system/service"
)

var DictData = dictDataController{}

type dictDataController struct {
}

// GetDictData 获取字典数据
func (c *dictDataController) GetDictData(ctx context.Context, req *system.GetDictReq) (res *system.GetDictRes, err error) {
res, err = commonService.DictData().GetDictWithDataByType(ctx, req)
return
}

// List 获取字典数据列表
func (c *dictDataController) List(ctx context.Context, req *system.DictDataSearchReq) (res *system.DictDataSearchRes, err error) {
res, err = commonService.DictData().List(ctx, req)
return
}

// Add 添加字典数据
func (c *dictDataController) Add(ctx context.Context, req *system.DictDataAddReq) (res *system.DictDataAddRes, err error) {
err = commonService.DictData().Add(ctx, req, service.Context().GetUserId(ctx))
return
}

// Get 获取对应的字典数据
func (c *dictDataController) Get(ctx context.Context, req *system.DictDataGetReq) (res *system.DictDataGetRes, err error) {
res, err = commonService.DictData().Get(ctx, req.DictCode)
return
}

// Edit 修改字典数据
func (c *dictDataController) Edit(ctx context.Context, req *system.DictDataEditReq) (res *system.DictDataEditRes, err error) {
err = commonService.DictData().Edit(ctx, req, service.Context().GetUserId(ctx))
return
}

func (c *dictDataController) Delete(ctx context.Context, req *system.DictDataDeleteReq) (res *system.DictDataDeleteRes, err error) {
err = commonService.DictData().Delete(ctx, req.Ids)
return
}
