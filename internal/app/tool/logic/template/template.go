package template

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/it00021hot/gen-admin/internal/app/tool/consts"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"github.com/it00021hot/gen-admin/internal/app/tool/service"
	utils "github.com/it00021hot/gen-admin/internal/app/tool/util"

	"os"
	"sync"
)

func init() {
	service.RegisterTemplate(New())
}

type (
	sTemplate struct {
		Mu sync.Mutex
	}
)

func New() *sTemplate {
	return &sTemplate{}
}

func (s *sTemplate) List(ctx context.Context, name string) (templates []model.TemplateItem, err error) {
	if name == "" {
		name = consts.Pattern
	}

	fileList, err := gfile.ScanDir(consts.BasePath, name, false)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}

	templates = make([]model.TemplateItem, 0)
	for _, fileName := range fileList {
		baseName := gfile.Basename(fileName)
		templates = append(templates, model.TemplateItem{
			Label: baseName,
			Value: baseName,
		})
	}
	return
}

func (s *sTemplate) GetTplTree(ctx context.Context, name string, isDir bool) (tree []*model.FileTree, err error) {
	tree, err = s.Tree(ctx, name, isDir)
	if err != nil {
		return
	}
	utils.TrimFirstPath(name, tree)
	return
}

func (s *sTemplate) Tree(ctx context.Context, name string, isDir bool) (tree []*model.FileTree, err error) {
	fileList, err := utils.GetPath(name, isDir)
	if err != nil {
		g.Log().Error(ctx, err)
		return nil, err
	}
	tree = make([]*model.FileTree, 0)
	for _, info := range fileList {
		path := info.Name()
		fileType := 0
		if name != "" {
			path = name + "/" + path
			fileType = 1
		}
		fileInfo := &model.FileTree{
			Name:       info.Name(),
			Path:       path,
			ModifyTime: gtime.New(info.ModTime()),
			IsDir:      info.IsDir(),
			Children:   make([]*model.FileTree, 0),
			Type:       fileType,
		}
		tree = append(tree, fileInfo)
		utils.GetTree(fileInfo, isDir)
	}
	return
}

func (s *sTemplate) AddTpl(ctx context.Context, name string) (err error) {

	path := fmt.Sprintf("%s/%s", consts.BasePath, name)
	isFile := gfile.Exists(path)
	if isFile {
		return gerror.New("已存在!")
	}
	err = gfile.Mkdir(path)
	create, err := gfile.Create(fmt.Sprintf("%s/%s", path, consts.Settings))
	defer func(create *os.File) {
		err := create.Close()
		if err != nil {
			g.Log().Error(gctx.New(), err.Error())
		}
	}(create)
	genInfo := g.Cfg().MustGet(ctx, "gen")
	genMap := make(map[string]interface{}, 0)
	genMap["gen"] = genInfo.Map()
	_, err = create.Write(gjson.New(genMap).MustToYaml())
	if err != nil {
		return err
	}
	if err != nil {
		return err
	}
	return
}

func (s *sTemplate) AddFile(tplModel *model.AddFileModel) (fileInfo *model.AddFileModel, err error) {

	//判断父路径，如果不存在或者是文件，不允许添加
	ParentPath := fmt.Sprintf("%s/%s", consts.BasePath, tplModel.ParentPath)
	if !gfile.Exists(ParentPath) || gfile.IsFile(ParentPath) {
		err = gerror.New("当前路径非文件夹，无法添加")
		return
	}
	//如果文件已存在，不允许添加
	path := fmt.Sprintf("%s/%s/%s", consts.BasePath, tplModel.ParentPath, tplModel.Name)
	if gfile.Exists(path) {
		return nil, gerror.New(tplModel.Name + "已存在!")
	}
	//判断新增类型
	if tplModel.IsDir {
		err = gfile.Mkdir(path)
		return tplModel, nil
	}
	//如果不带.tpl后缀，拼接上.tpl
	if gfile.Ext(path) != consts.TplExt {
		path = path + consts.TplExt
	}
	file, err := gfile.Create(path)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {
			g.Log().Error(gctx.New(), err.Error())
		}
	}(file)

	if tplModel.Content != "" {
		_, err = file.WriteString(tplModel.Content)
		if err != nil {
			return nil, err
		}
		if !gfile.Exists(path) {
			return nil, gerror.New("创建失败!")
		}
	}
	return tplModel, nil
}

func (s *sTemplate) Rename(path string, name string) error {
	oldPath := fmt.Sprintf("%s/%s", consts.BasePath, path)
	if !gfile.Exists(oldPath) {
		return gerror.New("原文件(夹)不存在")
	}
	//如果是文件类型而且新名称不带后缀，拼接上后缀
	if gfile.IsFile(oldPath) && gfile.Ext(name) != consts.TplExt {
		name = name + consts.TplExt
	}
	newPath := fmt.Sprintf("%s\\%s", gfile.Dir(oldPath), name)
	if gfile.Exists(newPath) {
		return gerror.New("已存在同名文件(夹)")

	}
	return gfile.Rename(oldPath, newPath)

}

func (s *sTemplate) DelTpl(name string) (err error) {
	path := fmt.Sprintf("%s/%s", consts.BasePath, name)
	err = gfile.Remove(path)
	return
}

func (s *sTemplate) DelFile(path string) (err error) {
	err = gfile.Remove(fmt.Sprintf("%s/%s", consts.BasePath, path))
	return
}

func (s *sTemplate) GetContent(path string) (result model.GetContentModel, err error) {
	filePath := fmt.Sprintf("%s/%s", consts.BasePath, path)
	if !gfile.IsFile(filePath) {
		err = gerror.New("非模板文件路径")
		return
	}
	contents := gfile.GetContents(filePath)
	result = model.GetContentModel{
		Path:       path,
		Name:       gfile.Basename(path),
		ParentPath: gfile.Dir(path),
		Content:    contents,
		ModifyTime: gtime.NewFromTime(gfile.MTime(filePath)),
		Size:       gfile.ReadableSize(filePath),
		IsDir:      gfile.IsDir(filePath),
	}
	return
}

func (s *sTemplate) SaveContent(path string, content string) error {
	filePath := fmt.Sprintf("%s/%s", consts.BasePath, path)
	if !gfile.Exists(filePath) {
		return gerror.New("文件不存在！")
	}
	if gfile.IsFile(filePath) && gfile.Ext(filePath) != consts.TplExt {
		filePath = filePath + consts.TplExt
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	err := gfile.PutContents(filePath, content)
	return err
}

func (s *sTemplate) GetSettings(name string) (content map[string]interface{}, err error) {
	filePath := fmt.Sprintf("%s/%s/%s", consts.BasePath, name, consts.Settings)
	if !gfile.IsFile(filePath) {
		create, err := gfile.Create(filePath)
		if err != nil {
			return nil, err
		}
		defer func(create *os.File) {
			err := create.Close()
			if err != nil {
				g.Log().Error(gctx.New(), err.Error())
			}
		}(create)
	}
	yaml, err := gjson.LoadYaml(gfile.GetContents(filePath))
	if err != nil {
		return
	}
	content = yaml.Map()
	return
}

func (s *sTemplate) SaveSettings(name string, content map[string]interface{}) (err error) {
	filePath := fmt.Sprintf("%s/%s/%s", consts.BasePath, name, consts.Settings)
	if !gfile.Exists(filePath) {
		return gerror.New("不能对模板文件(夹)设置属性")
	}
	s.Mu.Lock()
	defer s.Mu.Unlock()
	yamlString, err := gjson.New(content).ToYamlString()
	if err != nil {
		return err
	}
	err = gfile.PutContents(filePath, yamlString)
	return
}
