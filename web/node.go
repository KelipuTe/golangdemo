package web

import (
	"regexp"
	"strings"
)

const (
	nodeTypeStatic = iota + 1 //静态路由
	nodeTypeAny               //通配符路由
	nodeTypeParam             //路径参数路由
	nodeTypeRegexp            //正则表达式路由
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

// 路由结点
type s6RoutingNode struct {
	nodeType      int               //结点类型
	part          string            //这个路由结点代表的那段路径
	f8HandlerFunc F8HTTPHandlerFunc //命中路由之后的处理逻辑
	path          string            //从根路由到这个路由结点的全路径

	m3RoutingTree map[string]*s6RoutingNode //路由子树，k=子结点的 path，v=子结点
	p7ParamChild  *s6RoutingNode            //路径参数结点
	paramName     string                    //路径参数路由和正则表达式路由，都会提取路由参数的名字
	p7RegExpChild *s6RoutingNode            //正则表达式结点
	p7RegExp      *regexp.Regexp            //正则表达式
	p7AnyChild    *s6RoutingNode            //通配符结点

	s5f8Middleware      []F8HTTPMiddlewareFunc //结点上注册的中间件
	s5f8MiddlewareCache []F8HTTPMiddlewareFunc //服务启动后，命中结点时，需要用到的所有中间件
}

// 构建路由树时，查询子结点
func (p7this *s6RoutingNode) findChild(part string) *s6RoutingNode {
	//找静态路由
	if nil != p7this.m3RoutingTree {
		t4p7node, ok := p7this.m3RoutingTree[part]
		if ok {
			return t4p7node
		}
	}
	//找路径参数路由和正则表达式路由
	if ':' == part[0] {
		//正则表达式用括号包裹
		t4regIndex1 := strings.Index(part, "(")
		t4regIndex2 := strings.Index(part, ")")
		if -1 != t4regIndex1 && -1 != t4regIndex2 && t4regIndex1 < t4regIndex2 {
			//正则表达式路由
			return p7this.p7RegExpChild
		} else {
			//路径参数路由
			return p7this.p7ParamChild
		}
	}
	//找通配符路由
	if "*" == part {
		return p7this.p7AnyChild
	}
	return nil
}

// 构建路由树时，校验子结点是否可以继续操作
func (p7this *s6RoutingNode) checkChild(part string) {
	//这里需要校验路径参数路由和正则表达式路由是否冲突
	if ':' == part[0] {
		if p7this.part != part {
			panic(StrParamChildExist)
		}
	}
}

// createChild 构建路由树时，构造新的子结点
func (p7this *s6RoutingNode) createChild(part string, path string) *s6RoutingNode {
	if ':' == part[0] {
		t4regIndex1 := strings.Index(part, "(")
		t4regIndex2 := strings.Index(part, ")")
		if -1 != t4regIndex1 && -1 != t4regIndex2 && t4regIndex1 < t4regIndex2 {
			if nil != p7this.p7AnyChild {
				panic(StrRegexpChildClashWithAnyChild)
			}
			if nil != p7this.p7ParamChild {
				panic(StrRegexpChildClashWithParamChild)
			}
			if nil != p7this.p7RegExpChild {
				panic(StrRegexpChildExist)
			}

			p7this.p7RegExpChild = &s6RoutingNode{
				nodeType:  nodeTypeRegexp,
				part:      part,
				path:      path,
				paramName: part[1:t4regIndex1],
				p7RegExp:  regexp.MustCompile(part[t4regIndex1+1 : t4regIndex2]),
			}
			return p7this.p7RegExpChild
		} else {
			if nil != p7this.p7AnyChild {
				panic(StrParamChildClashWithAnyChild)
			}
			if nil != p7this.p7ParamChild {
				panic(StrParamChildExist)
			}
			if nil != p7this.p7RegExpChild {
				panic(StrParamChildClashWithRegexpChild)
			}

			p7this.p7ParamChild = &s6RoutingNode{
				nodeType:  nodeTypeParam,
				part:      part,
				path:      path,
				paramName: part[1:],
			}
			return p7this.p7ParamChild
		}
	}
	if "*" == part {
		if nil != p7this.p7AnyChild {
			panic(StrAnyChildExist)
		}
		if nil != p7this.p7ParamChild {
			panic(StrAnyChildClashWithParamChild)
		}
		if nil != p7this.p7RegExpChild {
			panic(StrAnyChildClashWithRegexpChild)
		}

		p7this.p7AnyChild = &s6RoutingNode{
			nodeType: nodeTypeAny,
			part:     part,
			path:     path,
		}
		return p7this.p7AnyChild
	}

	if nil == p7this.m3RoutingTree {
		p7this.m3RoutingTree = make(map[string]*s6RoutingNode)
	}
	_, ok := p7this.m3RoutingTree[part]
	if ok {
		panic(StrStaticChildExist)
	}

	p7this.m3RoutingTree[part] = &s6RoutingNode{
		nodeType: nodeTypeStatic,
		part:     part,
		path:     path,
	}
	return p7this.m3RoutingTree[part]
}

// 服务启动前，查询并缓存结点需要用到的所有中间件
func (p7this *s6RoutingNode) f8CacheMiddleware(s5f4mw []F8HTTPMiddlewareFunc) {
	t4s5f4mw := make([]F8HTTPMiddlewareFunc, 0, len(s5f4mw))
	//上一层结点的中间件
	if nil != s5f4mw {
		t4s5f4mw = append(t4s5f4mw, s5f4mw...)
	}
	//如果有通配符结点，则其他子结点需要把通配符结点上的中间件也加上
	if nil != p7this.p7AnyChild {
		p7this.p7AnyChild.f8CacheMiddleware(t4s5f4mw)
		if nil != p7this.p7AnyChild.s5f8Middleware {
			t4s5f4mw = append(t4s5f4mw, p7this.p7AnyChild.s5f8Middleware...)
		}
	}
	//添加这个结点上的中间件
	if nil != p7this.s5f8Middleware {
		t4s5f4mw = append(t4s5f4mw, p7this.s5f8Middleware...)
	}

	//如果这个结点有处理方法，那么这个结点就不是中间结点而是有效的路由结点，需要缓存中间件结果
	if nil != p7this.f8HandlerFunc {
		p7this.s5f8MiddlewareCache = make([]F8HTTPMiddlewareFunc, 0, len(t4s5f4mw))
		p7this.s5f8MiddlewareCache = append(p7this.s5f8MiddlewareCache, t4s5f4mw...)
	}
	//处理其余类型的子结点
	if nil != p7this.p7RegExpChild {
		p7this.p7RegExpChild.f8CacheMiddleware(t4s5f4mw)
	}
	if nil != p7this.p7ParamChild {
		p7this.p7ParamChild.f8CacheMiddleware(t4s5f4mw)
	}
	for _, p7childNode := range p7this.m3RoutingTree {
		p7childNode.f8CacheMiddleware(t4s5f4mw)
	}
}

// 查询路由时，匹配子结点
func (p7this *s6RoutingNode) matchChild(part string) *s6RoutingNode {
	//这里的查询优先级可以根据需要进行调整
	//先查询静态路由
	if nil != p7this.m3RoutingTree {
		p7node, ok := p7this.m3RoutingTree[part]
		if ok {
			return p7node
		}
	}
	//然后依次查询，正则表达式路由、路径参数路由、通配符路由
	if nil != p7this.p7RegExpChild {
		return p7this.p7RegExpChild
	} else if nil != p7this.p7ParamChild {
		return p7this.p7ParamChild
	} else if nil != p7this.p7AnyChild {
		return p7this.p7AnyChild
	}
	return nil
}
