package knast

import (
	"go/ast"
)

// FileParsingEntrance 文件的 ast 解析入口
type FileParsingEntrance struct {
	// p7node *ast.File，整个文件
	p7node ast.Node
	p7fv   *fileVisitor
}

// Visit 实现 ast.Visitor
func (p7this *FileParsingEntrance) Visit(node ast.Node) ast.Visitor {
	p7af, ok := node.(*ast.File)
	if ok {
		p7this.p7node = node
		// 解析 package 的注释
		t4s5c := ParseCommentGroup(p7af.Doc)
		p7this.p7fv = &fileVisitor{
			s5comment: t4s5c,
			s5p7node:  make([]ast.Node, 0),
			s5p7tv:    make([]*typeVisitor, 0),
		}
		return p7this.p7fv
	}
	return p7this
}

func (p7this *FileParsingEntrance) GetRes() FileRes {
	if nil != p7this.p7fv {
		return p7this.p7fv.GetRes()
	}
	return FileRes{
		S5comment: nil,
		S5type:    nil,
	}
}

// fileVisitor 文件层面的访问
type fileVisitor struct {
	// s5comment package 的注释
	s5comment []Comment
	// s5p7node []*ast.TypeSpec，文件中的每个 type
	s5p7node []ast.Node
	s5p7tv   []*typeVisitor
}

func (p7this *fileVisitor) Visit(node ast.Node) ast.Visitor {
	p7ats, ok := node.(*ast.TypeSpec)
	if ok {
		p7this.s5p7node = append(p7this.s5p7node, node)
		// 解析 type 的注释
		t4s5c := ParseCommentGroup(p7ats.Doc)
		t4p7tv := &typeVisitor{
			s5comment: t4s5c,
			s5p7node:  make([]ast.Node, 0),
			s5p7tf:    make([]*fieldVisitor, 0),
		}
		p7this.s5p7tv = append(p7this.s5p7tv, t4p7tv)
		return t4p7tv
	}
	return p7this
}

func (p7this *fileVisitor) GetRes() FileRes {
	t4s5t := make([]TypeRes, 0, len(p7this.s5p7tv))
	for _, t4p7tv := range p7this.s5p7tv {
		t4s5t = append(t4s5t, t4p7tv.GetRes())
	}
	return FileRes{
		S5comment: p7this.s5comment,
		S5type:    t4s5t,
	}
}

type FileRes struct {
	S5comment []Comment
	S5type    []TypeRes
}

// typeVisitor type 层面的访问
type typeVisitor struct {
	// type 的注释
	s5comment []Comment
	// s5p7node []*ast.Field，type 的字段
	s5p7node []ast.Node
	s5p7tf   []*fieldVisitor
}

func (p7this *typeVisitor) Visit(node ast.Node) (w ast.Visitor) {
	p7af, ok := node.(*ast.Field)
	if ok {
		p7this.s5p7node = append(p7this.s5p7node, p7af)
		// 解析 type 的字段的注释
		t4s5c := ParseCommentGroup(p7af.Doc)
		// 解析方法的参数
		var t4s5param []string = nil
		var t4s5result []string = nil
		t4p7ft, ok := (p7af.Type).(*ast.FuncType)
		if ok {
			if nil != t4p7ft.Params {
				t4paramNum := len(t4p7ft.Params.List)
				if t4paramNum > 0 {
					t4s5param = make([]string, 0, t4paramNum)
					for _, t4p7f := range t4p7ft.Params.List {
						t4name := "_"
						if nil != t4p7f.Names {
							t4name = t4p7f.Names[0].Name
						}
						t4type := t4p7f.Type.(*ast.Ident)
						t4s5param = append(t4s5param, t4name+":"+t4type.Name)
					}
				}
			}
			if nil != t4p7ft.Results {
				t4resultNum := len(t4p7ft.Results.List)
				if t4resultNum > 0 {
					t4s5result = make([]string, 0, t4resultNum)
					for _, t4p7f := range t4p7ft.Results.List {
						t4name := "_"
						if nil != t4p7f.Names {
							t4name = t4p7f.Names[0].Name
						}
						t4type := t4p7f.Type.(*ast.Ident)
						t4s5result = append(t4s5result, t4name+":"+t4type.Name)
					}
				}
			}
		}
		t4p7tf := &fieldVisitor{
			s5comment: t4s5c,
			s5param:   t4s5param,
			s5result:  t4s5result,
		}
		p7this.s5p7tf = append(p7this.s5p7tf, t4p7tf)
		return nil
	}
	return p7this
}

func (p7this *typeVisitor) GetRes() TypeRes {
	t4s5f := make([]FieldRes, 0, len(p7this.s5p7tf))
	for _, t4p7tv := range p7this.s5p7tf {
		t4s5f = append(t4s5f, t4p7tv.GetRes())
	}
	return TypeRes{
		S5comment: p7this.s5comment,
		S5field:   t4s5f,
	}
}

type TypeRes struct {
	S5comment []Comment
	S5field   []FieldRes
}

type fieldVisitor struct {
	s5comment []Comment
	s5param   []string
	s5result  []string
}

func (p7this *fieldVisitor) GetRes() FieldRes {
	return FieldRes{
		S5comment: p7this.s5comment,
		S5param:   p7this.s5param,
		S5result:  p7this.s5result,
	}
}

type FieldRes struct {
	S5comment []Comment
	S5param   []string
	S5result  []string
}
