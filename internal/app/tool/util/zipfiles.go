package util

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

//@author: [piexlmax](https://github.com/piexlmax)
//@function: ZipFiles
//@description: 压缩文件
//@param: filename string, files []string, oldForm, newForm string
//@return: error

func ZipFiles(filename string, files []string, oldForm, newForm string) error {
	newZipFile, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer func() {
		_ = newZipFile.Close()
	}()

	zipWriter := zip.NewWriter(newZipFile)
	defer func() {
		_ = zipWriter.Close()
	}()

	// 把files添加到zip中
	for _, file := range files {

		err = func(file string) error {
			zipFile, err := os.Open(file)
			if err != nil {
				return err
			}
			defer zipFile.Close()
			// 获取file的基础信息
			info, err := zipFile.Stat()
			if err != nil {
				return err
			}

			header, err := zip.FileInfoHeader(info)
			if err != nil {
				return err
			}

			// 使用上面的FileInfoHeader() 就可以把文件保存的路径替换成我们自己想要的了，如下面
			header.Name = strings.Replace(file, oldForm, newForm, -1)

			// 优化压缩
			// 更多参考see http://golang.org/pkg/archive/zip/#pkg-constants
			header.Method = zip.Deflate

			writer, err := zipWriter.CreateHeader(header)
			if err != nil {
				return err
			}
			if _, err = io.Copy(writer, zipFile); err != nil {
				return err
			}
			return nil
		}(file)
		if err != nil {
			return err
		}
	}
	return nil
}

func ZipByte(filename string, files []string, oldForm, newForm string) (bs []byte, err error) {
	// 创建一个缓冲区用来保存压缩文件内容
	buf := new(bytes.Buffer)
	// 创建一个压缩文档
	w := zip.NewWriter(buf)
	// 将文件加入压缩文档
	for _, file := range files {
		fileByte, err := ioutil.ReadFile(file)
		if err != nil {
			return bs, err
		}
		fileName := strings.Replace(file, oldForm, newForm, -1)
		f, err := w.Create(fileName)
		if err != nil {
			fmt.Println(err)
		}
		_, err = f.Write(fileByte)
		if err != nil {
			fmt.Println(err)
		}
	}
	// 关闭压缩文档
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
	i := buf.Bytes()
	return i, nil
}
