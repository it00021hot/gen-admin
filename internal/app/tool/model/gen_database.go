package model

type DatabasePageInput struct {
	PageReq
	Group string // 分组名称
	Name  string // 数据库名称
}
