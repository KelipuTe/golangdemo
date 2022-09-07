package knast

import (
	"go/ast"
	"strings"
)

// Comment 每个自定义注释的结构，自定义注释的结构就对应这个结构体的字段
// 这里自定义注释的结构如下
// 单行注释：`// @key value`
// 多行注释，字符串里面可能会有 `\r`、`\n`、`\t`
// `/* @key line1`
// `line2`
// `*/`
type Comment struct {
	Name string
	Text string
}

// ParseCommentGroup 从 ast.CommentGroup 中，解析出每个自定义注释
func ParseCommentGroup(p7cg *ast.CommentGroup) []Comment {
	if nil == p7cg {
		return nil
	}

	p7cgListLen := len(p7cg.List)
	if 0 == p7cgListLen {
		return nil
	}

	s5c := make([]Comment, 0, p7cgListLen)
	for _, t4p7c := range p7cg.List {
		// 解析每个自定义注释
		t4text := t4p7c.Text
		if strings.HasPrefix(t4text, "// ") {
			// 单行注释，把最前面的 `// ` 截掉
			t4text = t4text[3:]
		} else if strings.HasPrefix(t4text, "/* ") {
			// 多行注释，把最前面的 `/* ` 和最后面的 `*/` 截掉
			t4textLen := len(t4text)
			t4text = t4text[3 : t4textLen-2]
		} else {
			// 如果都不是，就跳过
			continue
		}
		// 注释结构差不多长这样，主要注意多行注释会有特殊符号
		// `@key value`、`@key line1\r\n\tline2\r\n\t`
		if strings.HasPrefix(t4text, "@") {
			// 判断一下有没有 @，然后用 ` ` 切成两份
			t4s5text := strings.SplitN(t4text, " ", 2)
			if 2 != len(t4s5text) {
				continue
			}
			// `@key` 要把 @ 去掉
			s5c = append(s5c, Comment{
				Name: t4s5text[0][1:],
				Text: t4s5text[1],
			})
		}
	}

	return s5c
}
