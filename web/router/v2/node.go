package v2

import (
	"regexp"
	"strings"
)

const (
	// 静态路由
	nodeTypeStatic = iota
	// 通配符路由
	nodeTypeAny
	// 路径参数路由
	nodeTypeParam
	// 正则表达式路由
	nodeTypeRegexp
)

// routingNode 路由结点
type routingNode struct {
	// nodeType 结点类型
	nodeType int
	// path 路径
	path string

	// m3routingTree 路由子树，子结点的 path => 子树根结点
	m3routingTree map[string]*routingNode
	// p7anyChild 通配符结点
	p7anyChild *routingNode
	// p7paramChild 路径参数结点
	p7paramChild *routingNode
	// paramName 参数路由和正则路由会提取路由参数的名字
	paramName string
	// p7regexpChild 正则表达式结点
	p7regexpChild *routingNode
	// p7regexp 正则表达式
	p7regexp *regexp.Regexp

	// f4handler 命中路由之后的处理逻辑
	f4handler HTTPHandleFunc
}

func newRoutingNode(part string) *routingNode {
	return &routingNode{}
}

func (p7this *routingNode) findChild(part string) *routingNode {
	// 通配符路由
	if "*" == part {
		return p7this.p7anyChild
	}
	// 路径参数路由和正则表达式路由以 `:` 开头
	if ':' == part[0] {
		// 正则表达式用括号包裹
		t4regIndex := strings.Index(part, "(")
		if -1 != t4regIndex {
			// 正则表达式路由
			return p7this.p7regexpChild
		} else {
			// 路径参数路由
			return p7this.p7paramChild
		}
	}
	return nil
}
