package model

import (
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gstructs"
	"github.com/it00021hot/gen-admin/internal/app/tool/model/entity"
	"testing"
)

func TestTableInfo(t *testing.T) {
	table := new(GenTableDTO)
	r, _ := gstructs.Fields(gstructs.FieldsInput{
		Pointer:         table,
		RecursiveOption: 1,
	})
	params := make([]Point, 0)
	for _, field := range r {
		p := Point{
			Label:  field.Name(),
			Detail: field.Tag("dc"),
		}
		if field.Name() == "Columns" || field.Name() == "PkColumn" {
			c := new(entity.GenTableColumn)
			cr, _ := gstructs.Fields(gstructs.FieldsInput{
				Pointer:         c,
				RecursiveOption: 1,
			})
			children := make([]Point, 0)
			for _, f := range cr {
				children = append(children, Point{
					Label:  f.Name(),
					Detail: f.Tag("dc"),
				})
			}
			p.Children = children

		}

		if field.Name() == "SubTable" {
			gt := new(GenTableDTO)
			tr, _ := gstructs.Fields(gstructs.FieldsInput{
				Pointer:         gt,
				RecursiveOption: 1,
			})
			children := make([]Point, 0)
			for _, f := range tr {
				children = append(children, Point{
					Label:  f.Name(),
					Detail: f.Tag("dc"),
				})
			}
			p.Children = children

		}

		params = append(params, p)

	}
	jsonString, _ := gjson.New(params).ToJsonString()
	t.Log(jsonString)
	t.Log("success")
}

type Point struct {
	Label    string  `json:"label,omitempty"`
	Detail   string  `json:"detail,omitempty"`
	Children []Point `json:"children,omitempty"`
}
