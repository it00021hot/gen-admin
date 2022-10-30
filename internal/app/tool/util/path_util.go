package util

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/it00021hot/gen-admin/internal/app/tool/consts"
	"github.com/it00021hot/gen-admin/internal/app/tool/model"
	"io/fs"
	"os"
)

func GetPath(path string, isDir bool) (fileList []fs.FileInfo, err error) {
	entries, err := os.ReadDir(consts.BasePath + "/" + path)
	if err != nil {
		return nil, err
	}
	fileList = make([]fs.FileInfo, 0, len(entries))
	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			return nil, err
		}
		fileList = append(fileList, info)
	}
	if err != nil {
		return
	}
	// 忽略以 . 开头的文件
	for i := 0; i < len(fileList); i++ {
		if fileList[i].Name()[0] == '.' {
			fileList = append(fileList[:i], fileList[i+1:]...)
		} else if isDir && !fileList[i].IsDir() {
			fileList = append(fileList[:i], fileList[i+1:]...)
		}
	}
	return

}

func GetTree(fileTree *model.FileTree, isDir bool) {

	files, err := GetPath(fileTree.Path, isDir)
	if err != nil {
		fmt.Println("read file path error", err)
		return
	}
	for _, fi := range files {
		if isDir && fi.IsDir() {
			fileInfo := &model.FileTree{
				Name:       fi.Name(),
				Path:       fmt.Sprintf("%s/%s", fileTree.Path, fi.Name()),
				IsDir:      fi.IsDir(),
				ModifyTime: gtime.New(fi.ModTime()),
			}
			fileTree.Children = append(fileTree.Children, fileInfo)
			if fi.IsDir() {
				fileInfo.Type = 1
				fileInfo.Children = make([]*model.FileTree, 0)
				GetTree(fileInfo, isDir)
			} else {
				fileInfo.Type = 2
			}
		} else if !isDir {
			fileInfo := &model.FileTree{
				Name:       fi.Name(),
				Path:       fmt.Sprintf("%s/%s", fileTree.Path, fi.Name()),
				IsDir:      fi.IsDir(),
				ModifyTime: gtime.New(fi.ModTime()),
			}
			fileTree.Children = append(fileTree.Children, fileInfo)
			if fi.IsDir() {
				fileInfo.Type = 1
				fileInfo.Children = make([]*model.FileTree, 0)
				GetTree(fileInfo, isDir)
			} else {
				fileInfo.Type = 2
			}
		}

	}
}

func TrimFirstPath(name string, tree []*model.FileTree) {
	for _, fileTree := range tree {
		fileTree.Path = gstr.TrimLeft(fileTree.Path, name+"/")
		if len(fileTree.Children) > 0 {
			TrimFirstPath(name, fileTree.Children)
		}
	}
}
