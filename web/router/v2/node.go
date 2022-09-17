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

var (
	StrStaticChildExist = "重复注册静态路由"
	StrParamChildExist  = "重复注册路径参数路由"
	StrRegexpChildExist = "重复注册正则表达式路由"
	StrAnyChildExist    = "重复注册通配符路由"

	StrParamChildClashWithAnyChild    = "路径参数路由和通配符路由冲突"
	StrParamChildClashWithRegexpChild = "路径参数路由和正则表达式路由冲突"
	StrRegexpChildClashWithAnyChild   = "正则表达式路由和通配符路由冲突"
	StrRegexpChildClashWithParamChild = "正则表达式路由和路径参数路由冲突"
	StrAnyChildClashWithParamChild    = "通配符路由和路径参数路由冲突"
	StrAnyChildClashWithRegexpChild   = "通配符路由和正则表达式路由冲突"
)

// routingNode 路由结点
type routingNode struct {
	// nodeType 结点类型
	nodeType int
	// path 路径
	path string

	part string

	// f4handler 命中路由之后的处理逻辑
	f4handler HTTPHandleFunc

	// m3routingTree 路由子树，子结点的 path => 子树根结点
	m3routingTree map[string]*routingNode
	// p7paramChild 路径参数结点
	p7paramChild *routingNode
	// paramName 路径参数路由和正则表达式路由，都会提取路由参数的名字
	paramName string
	// p7regexpChild 正则表达式结点
	p7regexpChild *routingNode
	// p7regexp 正则表达式
	p7regexp *regexp.Regexp
	// p7anyChild 通配符结点
	p7anyChild *routingNode
}

func (p7this *routingNode) findChild(part string) *routingNode {
	// 找静态路由
	if nil != p7this.m3routingTree {
		t4p7node, ok := p7this.m3routingTree[part]
		if ok {
			return t4p7node
		}
	}
	// 找路径参数路由和正则表达式路由
	if ':' == part[0] {
		// 正则表达式用括号包裹
		t4regIndex1 := strings.Index(part, "(")
		t4regIndex2 := strings.Index(part, ")")
		if -1 != t4regIndex1 && -1 != t4regIndex2 && t4regIndex1 < t4regIndex2 {
			// 正则表达式路由
			return p7this.p7regexpChild
		} else {
			// 路径参数路由
			return p7this.p7paramChild
		}
	}
	// 找通配符路由，
	if "*" == part {
		return p7this.p7anyChild
	}
	return nil
}

func (p7this *routingNode) createChild(part string) *routingNode {
	if ':' == part[0] {
		t4regIndex1 := strings.Index(part, "(")
		t4regIndex2 := strings.Index(part, ")")
		if -1 != t4regIndex1 && -1 != t4regIndex2 && t4regIndex1 < t4regIndex2 {
			if nil != p7this.p7anyChild {
				panic(StrRegexpChildClashWithAnyChild)
			}
			if nil != p7this.p7paramChild {
				panic(StrRegexpChildClashWithParamChild)
			}
			if nil != p7this.p7regexpChild {
				panic(StrRegexpChildExist)
			}

			p7this.p7regexpChild = &routingNode{
				nodeType:  nodeTypeRegexp,
				path:      part,
				paramName: part[1:t4regIndex1],
				p7regexp:  regexp.MustCompile(part[t4regIndex1+1 : t4regIndex2]),
			}
			return p7this.p7regexpChild
		} else {
			if nil != p7this.p7anyChild {
				panic(StrParamChildClashWithAnyChild)
			}
			if nil != p7this.p7paramChild {
				panic(StrParamChildExist)
			}
			if nil != p7this.p7regexpChild {
				panic(StrParamChildClashWithRegexpChild)
			}

			p7this.p7paramChild = &routingNode{
				nodeType:  nodeTypeParam,
				path:      part,
				paramName: part[1:],
			}
			return p7this.p7paramChild
		}
	}
	if "*" == part {
		if nil != p7this.p7anyChild {
			panic(StrAnyChildExist)
		}
		if nil != p7this.p7paramChild {
			panic(StrAnyChildClashWithParamChild)
		}
		if nil != p7this.p7regexpChild {
			panic(StrAnyChildClashWithRegexpChild)
		}

		p7this.p7anyChild = &routingNode{
			nodeType: nodeTypeAny,
			path:     part,
		}
		return p7this.p7anyChild
	}

	if nil == p7this.m3routingTree {
		p7this.m3routingTree = make(map[string]*routingNode)
	}
	_, ok := p7this.m3routingTree[part]
	if ok {
		panic(StrStaticChildExist)
	}

	p7this.m3routingTree[part] = &routingNode{
		nodeType: nodeTypeStatic,
		path:     part,
	}
	return p7this.m3routingTree[part]
}

func (p7this *routingNode) matchChild(part string) *routingNode {
	if nil != p7this.m3routingTree {
		p7node, ok := p7this.m3routingTree[part]
		if ok {
			return p7node
		}
	}
	if nil != p7this.p7regexpChild {
		return p7this.p7regexpChild
	} else if nil != p7this.p7paramChild {
		return p7this.p7paramChild
	} else if nil != p7this.p7anyChild {
		return p7this.p7anyChild
	}
	return nil
}
