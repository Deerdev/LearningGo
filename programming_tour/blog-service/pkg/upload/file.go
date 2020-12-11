package upload

import (
	"io"
	"io/ioutil"
	"mime/multipart"
	"os"
	"path"
	"strings"

	"github.com/go-programming-tour-book/blog-service/global"
	"github.com/go-programming-tour-book/blog-service/pkg/util"
)

type FileType int

/*
（1）定义FileType为int的类型别名，并把FileType作为类别标识的基础类型。（2）把iota作为它的初始值。
iota又是什么呢？实际上，在Go语言中，iota相当于一个const的常量计数器，也可以理解为枚举值。
第一个声明的iota的值为0，在新的一行被使用时，它的值会自动递增.

为什么要在FileType类型中使用iota呢？其实是为了在后续有其他需求时，能标准化地进行处理(自增)
const (
	TypeImage FileType = iota + 1
	TypeExcel
	TypeTxt
)
 */
const TypeImage FileType = iota + 1

// 获取文件名称。通过获取文件后缀筛选出原始文件名称，对其进行MD5加密，最后返回经过加密处理后的文件名称
func GetFileName(name string) string {
	ext := GetFileExt(name)
	fileName := strings.TrimSuffix(name, ext)
	fileName = util.EncodeMD5(fileName)

	return fileName + ext
}

// 获取文件后缀。通过调用 path.Ext 方法循环查找″.″符号，最后通过切片索引返回对应的文件后缀
func GetFileExt(name string) string {
	return path.Ext(name)
}

// 获取文件保存地址。这里直接返回配置中的文件保存目录即可，以便后续的调整
func GetSavePath() string {
	return global.AppSetting.UploadSavePath
}

func GetServerUrl() string {
	return global.AppSetting.UploadServerUrl
}

// 检查保存目录是否存在，通过调用 os.Stat 方法获取文件的描述信息FileInfo，并调用os.IsNotExist方法进行判断。
// 其原理是，对os.Stat方法返回的error值与系统中定义的oserror.ErrNotExist值进行比较，以达到校验效果
func CheckSavePath(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsNotExist(err)
}

// 检查文件后缀是否包含在约定的后缀配置项中。需要注意的是，上传的文件的后缀有可能是大写、小写、大小写混合等，
// 因此我们需要调用strings.ToUpper方法把后缀统一转为大写（固定的格式）后再进行匹配。
func CheckContainExt(t FileType, name string) bool {
	ext := GetFileExt(name)
	ext = strings.ToUpper(ext)
	switch t {
	case TypeImage:
		for _, allowExt := range global.AppSetting.UploadImageAllowExts {
			if strings.ToUpper(allowExt) == ext {
				return true
			}
		}

	}

	return false
}

// 检查文件大小是否超出限制。
func CheckMaxSize(t FileType, f multipart.File) bool {
	content, _ := ioutil.ReadAll(f)
	size := len(content)
	switch t {
	case TypeImage:
		if size >= global.AppSetting.UploadImageMaxSize*1024*1024 {
			return true
		}
	}

	return false
}

// 检查文件权限是否足够。与 CheckSavePath 方法原理一致，即与oserror.ErrPermission值进行比较，进而做出判断。
func CheckPermission(dst string) bool {
	_, err := os.Stat(dst)

	return os.IsPermission(err)
}

// 创建保存上传文件的目录。在方法内部调用os.MkdirAll方法，该方法会以传入的 os.FileMode 权限位递归创建所需的所有目录结构。
// 若涉及的目录均已存在，则不进行任何操作，直接返回nil
func CreateSavePath(dst string, perm os.FileMode) error {
	err := os.MkdirAll(dst, perm)
	if err != nil {
		return err
	}

	return nil
}

// 保存上传的文件。该方法通过调用os.Create方法创建目标地址文件，再通过file.Open方法打开源地址的文件，结合io.Copy方法实现两者之间的文件内容拷贝。
func SaveFile(file *multipart.FileHeader, dst string) error {
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
