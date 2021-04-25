package tools

import (
	"archive/zip"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"mime/multipart"
	"os"
	"path"
	"path/filepath"
	"strings"
	// "github.com/spf13/afero"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func PathCreate(dir string) error {
	return os.MkdirAll(dir, os.ModePerm)
}

func FileCreate(content bytes.Buffer, name string) (err error) {
	file, err := os.Create(name)
	if err != nil {
		return
	}
	defer file.Close()
	_, err = file.WriteString(content.String())
	if err != nil {
		return
	}
	// for i := n; i < len(content.Len()); i++ {
	// 	//写入byte的slice数据
	// 	file.Write(content)
	// 	//写入字符串
	// 	//
	// }
	return
}

func PathRemove(name string) (err error) {
	if err = os.RemoveAll(name); err != nil {
		return
	}
	return
}

func FileRemove(name string) {
	err := os.Remove(name)
	if err != nil {
		log.Println(err)
	}
}

func FileZip(dst, src string, notContPath string) (err error) {
	//创建准备写入的文件
	fw, err := os.Create(dst)
	defer fw.Close()
	if err != nil {
		return err
	}

	// 通过 fw 来创建 zip.Write
	zw := zip.NewWriter(fw)
	defer func() {
		// 检测一下是否成功关闭
		if err := zw.Close(); err != nil {
			log.Fatalln(err)
		}
	}()
	return filepath.Walk(src, func(path string, fi os.FileInfo, errBack error) (err error) {
		if errBack != nil {
			return errBack
		}

		fh, err := zip.FileInfoHeader(fi)
		if err != nil {
			return
		}

		fh.Name = strings.TrimPrefix(path, string(filepath.Separator))

		if fi.IsDir() {
			fh.Name += "/"
		}
		fh.Name = strings.Replace(fh.Name, notContPath, "", -1)

		w, err := zw.CreateHeader(fh)
		if err != nil {
			return
		}

		if !fh.Mode().IsRegular() {
			return nil
		}

		fr, err := os.Open(path)
		defer fr.Close()
		if err != nil {
			return
		}

		n, err := io.Copy(w, fr)
		if err != nil {
			return
		}

		fmt.Printf("成功压缩文件： %s, 共写入了 %d 个字符的数据\n", path, n)

		return nil
	})
}

type ReplaceHelper struct {
	Root    string //路径
	OldText string //需要替换的文本
	NewText string //新的文本
}

func (h *ReplaceHelper) DoWrok() error {
	return filepath.Walk(h.Root, h.walkCallback)
}

func (h ReplaceHelper) walkCallback(path string, f os.FileInfo, err error) error {
	if err != nil {
		return err
	}
	if f == nil {
		return nil
	}
	if f.IsDir() {
		log.Println("DIR:", path)
		return nil
	}

	//文件类型需要进行过滤

	buf, err := ioutil.ReadFile(path)
	if err != nil {
		//err
		return err
	}
	content := string(buf)
	log.Printf("h.OldText: %s \n", h.OldText)
	log.Printf("h.NewText: %s \n", h.NewText)

	//替换
	newContent := strings.Replace(content, h.OldText, h.NewText, -1)

	//重新写入
	ioutil.WriteFile(path, []byte(newContent), 0)

	return err
}

// 获取文件大小
func GetSize(f multipart.File) (int, error) {
	content, err := ioutil.ReadAll(f)

	return len(content), err
}

// 获取文件后缀
func GetExt(fileName string) string {
	return path.Ext(fileName)
}

// CheckExist 检查文件是否存在
func CheckExist(src string) (bool, error) {
	_, err := os.Stat(src)
	if err != nil {
		if os.IsNotExist(err) {
			return false, nil
		} else if os.IsExist(err) {
			return true, nil
		}
		return false, err
	}
	return true, nil
}

// CheckPermission 检查文件权限
func CheckPermission(src string) (bool, error) {
	_, err := os.Stat(src)
	if err != nil {
		if os.IsPermission(err) {
			return true, nil
		}
		return false, err
	}
	return false, nil
}

//IsNotExistMkDir 如果不存在则新建文件夹
func IsNotExistMkDir(src string) error {
	exist, err := CheckExist(src)
	if err != nil {
		return err
	}
	if !exist {
		if err := MkDir(src); err != nil {
			return err
		}
	}
	return nil
}

//MkDir 新建文件夹
func MkDir(src string) error {
	err := os.Mkdir(src, os.ModePerm)
	if err != nil {
		return err
	}
	return nil
}

// 打开文件
func Open(name string, flag int, perm os.FileMode) (*os.File, error) {
	f, err := os.OpenFile(name, flag, perm)
	if err != nil {
		return nil, err
	}
	return f, nil
}
