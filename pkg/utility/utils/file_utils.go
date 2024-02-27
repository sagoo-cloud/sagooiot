package utils

import (
	"archive/zip"
	"encoding/json"
	"fmt"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"io"
	"os"
	"path"
	"path/filepath"
	"strings"
)

// ReadZipFileByFileName 读取压缩包指定文件内容
func ReadZipFileByFileName(file *ghttp.UploadFile, fileName string) (result map[string]interface{}, err error) {
	if fileName == "" {
		err = gerror.New("文件名称不能为空")
		return
	}
	src, err := file.Open()
	if err != nil {
		return
	}
	// 打开zip文件  对于macOS生成的zip文件，需要去掉前面的__MACOSX目录
	var zr *zip.Reader
	zr, err = zip.NewReader(src, file.Size)
	if err != nil {
		return
	}

	var zipFile *zip.File
	for _, f := range zr.File {
		if !f.FileInfo().IsDir() && f.FileInfo().Name() == fileName {
			zipFile = f
			break
		}
	}

	if zipFile == nil {
		err = gerror.Newf("无指定%s文件,请确认后再上传!", fileName)
		return
	}

	var fileJson io.ReadCloser
	fileJson, err = zipFile.Open()
	if err != nil {
		err = gerror.Newf("打开文件失败：%s", err)
		return
	}

	// 读取JSON文件数据
	var data []byte
	data, err = io.ReadAll(fileJson)
	if err != nil {
		err = gerror.Newf("读取文件失败：%s", err)
		return
	}

	err = json.Unmarshal(data, &result)
	if err != nil {
		err = gerror.Newf("解析JSON数据失败：%s", err)
		return
	}
	return
}

// UploadZip 上传ZIP包
func UploadZip(file *ghttp.UploadFile, zipPath string, excludeFiles []string) (err error) {

	src, err := file.Open()
	if err != nil {
		return
	}
	// 打开zip文件  对于macOS生成的zip文件，需要去掉前面的__MACOSX目录
	var zr *zip.Reader
	zr, err = zip.NewReader(src, file.Size)
	if err != nil {
		return
	}

	//上传路径
	if !FileIsExisted(zipPath) {
		// 文件夹不存在，创建文件夹
		if err = os.MkdirAll(zipPath, os.ModePerm); err != nil {
			return
		}
	}

	type closer func()
	var cleanups []closer
	defer func() {
		for _, cleanup := range cleanups {
			cleanup()
		}
	}()

	for _, f := range zr.File {
		// 去掉__MACOSX目录
		if strings.Contains(f.Name, "__MACOSX") {
			continue
		}

		// 判断是否为文件夹
		if f.FileInfo().IsDir() || ShouldExclude(f.Name, excludeFiles) {
			continue
		}

		var rc io.ReadCloser
		rc, err = f.Open()
		if err != nil {
			return
		}
		cleanups = append(cleanups, func() { rc.Close() })

		destFilePath := filepath.Join(zipPath, filepath.Base(f.Name))
		if f.FileInfo().IsDir() {
			if err = os.MkdirAll(destFilePath, os.ModePerm); err != nil {
				return
			}
		} else {
			var destFile *os.File
			destFile, err = os.OpenFile(destFilePath, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, f.Mode())
			if err != nil {
				return
			}
			cleanups = append(cleanups, func() { destFile.Close() })

			_, err = io.Copy(destFile, rc)
			if err != nil {
				return
			}
		}
	}
	return
}

func ShouldExclude(filename string, excludeList []string) bool {
	baseName := filepath.Base(filename)
	for _, excludedFile := range excludeList {
		if baseName == excludedFile {
			return true
		}
	}
	return false
}

// WriteToFile 写入文件
func WriteToFile(fileName string, content string) error {
	f, err := os.OpenFile(fileName, os.O_WRONLY|os.O_TRUNC|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	n, _ := f.Seek(0, io.SeekEnd)
	_, err = f.WriteAt([]byte(content), n)
	defer func(f *os.File) {
		if err := f.Close(); err != nil {
			fmt.Println(err)
		}
	}(f)
	return err
}

// FileIsExisted 文件或文件夹是否存在
func FileIsExisted(filename string) bool {
	existed := true
	if _, err := os.Stat(filename); os.IsNotExist(err) {
		existed = false
	}
	return existed
}

// ParseFilePath 解析路径获取文件名称及后缀
func ParseFilePath(pathStr string) (fileName string, fileType string) {
	fileNameWithSuffix := path.Base(pathStr)
	fileType = path.Ext(fileNameWithSuffix)
	fileName = strings.TrimSuffix(fileNameWithSuffix, fileType)
	return
}
