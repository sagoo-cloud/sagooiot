package utility

import (
	"log"
	"os"
	"path/filepath"
)

func InitFatalLog(loggerPath string) *os.File {
	loggerPath = loggerPath + "-error.log"
	// 确保目录存在
	dir := filepath.Dir(loggerPath)
	if err := os.MkdirAll(dir, 0755); err != nil {
		log.Fatalf("创建目录失败: %v", err)
		return nil
	}

	// 打开（如果不存在则创建）日志文件
	file, err := os.OpenFile(loggerPath, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("打开文件失败: %v", err)
		return nil
	}

	return file
}
