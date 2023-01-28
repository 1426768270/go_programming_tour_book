package upload

import (
	"blog-service/global"
	"blog-service/pkg/util"
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

// 类型别名
type FileType int

//计数器，为了扩展 execl/txt
const TypeImage FileType = iota +1

// 获取文件名称，先是通过获取文件后缀并筛出原始文件名进行 MD5 加密，最后返回经过加密处理后的文件名。
func GetFileName(name string) string{
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)
	return fileName + ext
}
//获取文件后缀，主要是通过调用 path.Ext 方法进行循环查找”.“符号，最后通过切片索引返回对应的文化后缀名称。
func GetFileExt(name string) string {
	return path.Ext(name)
}

func GetSavePath() string{
	return global.AppSetting.UploadSavePath
}
// 检查保存目录是否存在
func CheckSavePath(dst string) bool{
	_, err:= os.Stat(dst)
	return os.IsNotExist(err)
}

// 检查文件后缀是否包含在约定的后缀配置项中
func CheckContainExt(t FileType, name string) bool{
	ext :=GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage :
		for _, allowExt := range global.AppSetting.UploadImageAllowExts{
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}
	}
	return false
}

func CheckMaxSize(t FileType, f multipart.File) bool{
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize * 1024 * 1024 {
			return true
		}
	}
	return false
}

// 检查保存目录是否存在
func CheckPermission(dst string) bool{
	_, err := os.Stat(dst)
	return os.IsPermission(err)
}

func CreateSavePath(dst string, perm os.FileMode) error{
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}
	return nil
}

func SaveFile(file *multipart.FileHeader, dst string) error{
	src, err := file.Open()
	if err != nil {
		return err
	}
	defer src.Close()
	out, err := os.Create(dst)
	if err != nil {
		return err
	}
	defer out.Close()
	_, err = io.Copy(out, src)
	return err
}