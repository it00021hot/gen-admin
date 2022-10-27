
package model

import "github.com/it00021hot/gen-admin/internal/app/system/model/entity"

type SysDeptTreeRes struct {
*entity.SysDept
Children []*SysDeptTreeRes `json:"children"`
}
