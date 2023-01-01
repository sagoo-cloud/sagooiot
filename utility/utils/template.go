package utils

import (
	"bytes"
	"text/template"
)

func ReplaceTemplate(content string, variable map[string]interface{}) (result string, err error) {
	// 创建模版对象
	tmpl, err := template.New("text").Parse(content)
	if err != nil {
		return
	}
	// 创建缓冲区
	var buf bytes.Buffer
	// 执行模版替换
	if err = tmpl.Execute(&buf, variable); err != nil {
		return
	}
	// 输出结果
	result = buf.String()
	return
}
