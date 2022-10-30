package util

import (
	"fmt"
	"github.com/it00021hot/gen-admin/internal/app/tool/consts"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gview"
	"github.com/gogf/gf/v2/text/gregex"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/gogf/gf/v2/util/gconv"
	"io/ioutil"
	"strings"
)

// InitView 初始化模板视图解析配置
func InitView() *gview.View {
	view := gview.New()
	err := view.SetConfigWithMap(g.Map{
		"Paths":      consts.BasePath,
		"Delimiters": []string{"{{", "}}"},
	})
	if err != nil {
		return nil
	}
	view.BindFuncMap(g.Map{
		"UcFirst": gstr.UcFirst,
		"Sum": func(a, b int) int {
			return a + b
		},
		"CaseCamelLower":     gstr.CaseCamelLower,     //首字母小写驼峰
		"CaseCamel":          gstr.CaseCamel,          //首字母大写驼峰
		"HasSuffix":          gstr.HasSuffix,          //是否存在后缀
		"ContainsI":          gstr.ContainsI,          //是否包含子字符串
		"SnakeScreamingCase": gstr.CaseSnakeScreaming, //转常量名
		"InArray":            gstr.InArray,            //是否在数组中
		"VueTag": func(t string) string {
			return t
		},
	})
	return view
}

// PrepareContext 设置模板变量信息
func PrepareContext(table *model.GenTableDTO) g.Map {
	//模板变
	tplContext := g.Map{
		"Table":      table,
		"ImportList": GetImportList(table),
	}
	if table.Params != "" {
		tplContext["Params"] = gconv.Map(table.Params)
	}
	return tplContext
}

// GetImportList 根据列类型获取导入包
func GetImportList(extendData *model.GenTableDTO) (importList []string) {
	for _, column := range extendData.Columns {
		if column.JavaType == "Date" {
			importList = append(importList, "java.util.Date", "com.fasterxml.jackson.annotation.JsonFormat")
		} else if column.JavaType == "BigDecimal" {
			importList = append(importList, "java.math.BigDecimal")
		}
	}
	return removeRepByMap(importList)
}

// slice去重
func removeRepByMap(slc []string) []string {
	var result []string          //存放返回的不重复切片
	tempMap := map[string]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0 //当e存在于tempMap中时，再次添加是添加不进去的，，因为key不允许重复
		//如果上一行添加成功，那么长度发生变化且此时元素一定不重复
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e) //当元素不重复时，将元素添加到切片result中
		}
	}
	return result
}

func GetAllTplFile(pathName string, category string, fileList []string) ([]string, error) {

	files, err := ioutil.ReadDir(pathName)
	for _, fi := range files {
		if fi.IsDir() {
			fileList, err = GetAllTplFile(pathName+"/"+fi.Name(), "", fileList)
			if err != nil {
				return nil, err
			}
		} else {
			if strings.HasSuffix(fi.Name(), ".tpl") {
				if category != consts.TplSub {
					if !strings.Contains(fi.Name(), consts.TplSub) {
						fileList = append(fileList, pathName+"/"+fi.Name())
					}
				} else {
					fileList = append(fileList, pathName+"/"+fi.Name())
				}
			}
		}
	}
	return fileList, err
}

// TrimBreak 剔除多余的换行
func TrimBreak(str string) (rStr string, err error) {
	var b []byte
	if b, err = gregex.Replace("(([\\s\t]*)\r?\n){2,}", []byte("$2\n"), []byte(str)); err != nil {
		return
	}
	if b, err = gregex.Replace("(([\\s]*)/{4})", []byte("$2\t"), b); err == nil {
		rStr = gconv.String(b)
	}
	return
}

// GetFileName 获取文件名
func GetFileName(templateName string, tpl string, fileMap map[string]string) (fileName string) {
	trimBase := strings.TrimPrefix(templateName, fmt.Sprintf("%s/%s/", consts.BasePath, tpl))
	if len(fileMap) == 0 {
		if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
			fileName = strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
		}
	} else {
		settingFileName := fileMap[trimBase]
		if settingFileName == "" {
			if lastSeparator := strings.LastIndex(trimBase, "/"); lastSeparator != -1 {
				fileName = strings.TrimSuffix(trimBase[lastSeparator+1:], ".tpl")
			}
		} else {
			fileName = fileMap[trimBase]
		}
	}
	return fileName
}
