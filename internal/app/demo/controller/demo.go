package controller

import (
	"context"
	"github.com/it00021hot/gen-admin/api/v1/demo"
)

var Demo = cDemo{}

type cDemo struct {
}

func (c *cDemo) Demo(ctx context.Context, req *demo.DmReq) (res *demo.DmRes, err error) {
	res = &demo.DmRes{Name: "赵四"}
	panic("demo wrong")
	return
}
