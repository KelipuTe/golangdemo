package knast

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"testing"
)

func TestParseEnterVisit(t *testing.T) {
	p7fs := token.NewFileSet()
	p7af, _ := parser.ParseFile(p7fs, "target_file.go", nil, parser.ParseComments)
	p7fpe := &FileParsingEntrance{}
	ast.Walk(p7fpe, p7af)
	fr := p7fpe.GetRes()
	fmt.Printf("%+v", fr)
}
