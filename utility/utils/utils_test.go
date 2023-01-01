package utils

import (
	"testing"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/test/gtest"
)

func TestMarkdownToHtml(t *testing.T) {
	gtest.C(t, func(t *gtest.T) {
		content := `
## 参与贡献

1. 框架代码：参与框架功能开发、单元测试、ISSUE提交、反馈建议等等，https://github.com/gogf/gf
2. 开发文档：参与开发文档的撰写，便于更多的人了解、热爱并加入团队，https://github.com/gogf/gf-doc

`
		g.Dump(MarkdownToHtml(content))
	})
}
