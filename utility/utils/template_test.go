package utils

import (
	"testing"
)

func TestReplaceTemplate(t *testing.T) {
	// 定义文本模版
	textTmpl := `Hello, {{.Name}}! You are {{.Age}} years old`
	// 定义数据
	data := map[string]interface{}{
		"Name": "John Doe",
		"Age":  30,
	}

	res, err := ReplaceTemplate(textTmpl, data)
	if err != nil {
		t.Error(res)

	}
	t.Log(res)

}
