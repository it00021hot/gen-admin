// =================================================================================
// Code generated by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// GenDatabase is the golang structure for table gen_database.
type GenDatabase struct {
	Id         int64       `json:"id"         ` // 编号
	Group      string      `json:"group"      ` // 分组名称
	Host       string      `json:"host"       ` // 地址
	Port       string      `json:"port"       ` // 端口
	User       string      `json:"user"       ` // 账号
	Pass       string      `json:"pass"       ` // 密码
	Name       string      `json:"name"       ` // 数据库名称
	Type       string      `json:"type"       ` // 数据库类型
	CreateBy   string      `json:"createBy"   ` // 创建者
	CreateTime *gtime.Time `json:"createTime" ` // 创建时间
	UpdateBy   string      `json:"updateBy"   ` // 更新者
	UpdateTime *gtime.Time `json:"updateTime" ` // 更新时间
	Remark     string      `json:"remark"     ` // 备注
}