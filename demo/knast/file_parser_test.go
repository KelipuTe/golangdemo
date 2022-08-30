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
	p7af, err := parser.ParseFile(p7fs, "target_file.go", nil, parser.ParseComments)
	pe := &ParsingEntrance{}
	ast.Walk(pe, p7af)
	fmt.Println(pe, err)
}
