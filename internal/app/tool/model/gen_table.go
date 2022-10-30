package model

import (
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
)

type TablePageInput struct {
	PageReq
	Group        string `q:"eq"`
	TableName    string `q:"like"`
	TableComment string `q:"like"`
}

type TableImportSaveReq struct {
	Group  string
	Tables []string
}

type GenTableDTO struct {
	TableId        int64                    `json:"tableId,omitempty"        dc:"编号" `                                // 编号
	TableName      string                   `json:"tableName,omitempty"      dc:"表名称" `                               // 表名称
	PrefixName     string                   `json:"prefixName,omitempty"     dc:"无前缀表面" `                             // 无前缀表名称
	TableComment   string                   `json:"tableComment,omitempty"   dc:"表描述" `                               // 表描述
	TplCategory    string                   `json:"tplCategory,omitempty"    dc:"使用的模板（crud单表操作 tree树表操作 sub主子表操作）" ` // 使用的模板（crud单表操作 tree树表操作 sub主子表操作）
	SubTableName   string                   `json:"subTableName,omitempty"   dc:"关联子表的表名" `                           // 关联子表的表名
	SubTableFkName string                   `json:"subTableFkName,omitempty" dc:"子表关联的外键名" `                          // 子表关联的外键名
	TreeCode       string                   `json:"treeCode,omitempty"       dc:"树编码字段" `                             // 树编码字段
	TreeName       string                   `json:"treeName,omitempty"       dc:"树编码字段" `                             // 树编码字段
	TreeParentCode string                   `json:"treeParentCode,omitempty" dc:"树编码字段" `                             // 树编码字段
	ClassName      string                   `json:"className,omitempty"      dc:"实体类名称" `                             // 实体类名称
	SystemName     string                   `json:"systemName,omitempty"     dc:"系统名称" `                              // 系统名称
	ModuleName     string                   `json:"moduleName,omitempty"     dc:"生成模块名" `                             // 生成模块名
	PackageName    string                   `json:"packageName,omitempty"    dc:"生成包路径" `                             // 生成包路径
	BusinessName   string                   `json:"businessName,omitempty"   dc:"生成业务名" `                             // 生成业务名
	FunctionName   string                   `json:"functionName,omitempty"   dc:"生成功能名" `                             // 生成功能名
	FunctionAuthor string                   `json:"functionAuthor,omitempty" dc:"生成功能作者" `                            // 生成功能作者
	GenType        string                   `json:"genType,omitempty"        dc:"生成代码方式（0zip压缩包 1自定义路径）" `            // 生成代码方式（0zip压缩包 1自定义路径）
	GenPath        string                   `json:"genPath,omitempty"        dc:"生成路径（不填默认项目路径）" `                    // 生成路径（不填默认项目路径）
	Params         string                   `json:"params,omitempty"         dc:"额外属性" `                              // 额外属性
	CreateBy       string                   `json:"createBy,omitempty"       dc:"创建者" `                               // 创建者
	CreateTime     *gtime.Time              `json:"createTime,omitempty"     dc:"创建时间" `                              // 创建时间
	UpdateBy       string                   `json:"updateBy,omitempty"       dc:"更新者" `                               // 更新者
	UpdateTime     *gtime.Time              `json:"updateTime,omitempty"     dc:"更新时间" `                              // 更新时间
	Remark         string                   `json:"remark,omitempty"         dc:"备注" `                                // 备注
	PkColumn       *entity.GenTableColumn   `json:"pkColumn,omitempty" dc:"主键信息"`                                     // 主键信息
	SubTable       *GenTableDTO             `json:"subTable,omitempty" dc:"子表信息"`                                     // 子表信息
	Columns        []*entity.GenTableColumn `json:"columns,omitempty" dc:"字段集合"`                                      // 字段集合
	PublicFields   []string                 `json:"hideFields,omitempty" dc:"不生成字段"`                                  // 公共字段
}
