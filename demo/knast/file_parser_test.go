package knast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParseEnterVisit(t *testing.T) {
	//s5case := []struct {
	//	filename string
	//	src      any
	//	want     FileRes
	//}{
	//	{
	//		filename: "target_file.go",
	//		src:      nil,
	//		want:     FileRes{},
	//	},
	//}

	p7fs := token.NewFileSet()
	p7af, _ := parser.ParseFile(p7fs, "target_file.go", nil, parser.ParseComments)
	p7fpe := &FileParsingEntrance{}
	ast.Walk(p7fpe, p7af)
	fr := p7fpe.GetRes()
	fmt.Println(fr)
}
