package main

import (
	"bufio"
	"bytes"
	"demo-golang/homework/shi2zhan4/week03/gen/annotation"
	http2 "demo-golang/homework/shi2zhan4/week03/gen/http"
	"errors"
	"fmt"
	"go/ast"
	"go/format"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"strings"
	"unicode"
)

// 实际上 main 函数这里要考虑接收参数
// src 源目标
// dst 目标目录
// type src 里面可能有很多类型，那么用户可能需要指定具体的类型
// 这里我们简化操作，只读取当前目录下的数据，并且扫描下面的所有源文件，然后生成代码
// 在当前目录下运行 go install 就将 main 安装成功了，
// 可以在命令行中运行 gen
// 在 testdata 里面运行 gen，则会生成能够通过所有测试的代码
func main() {
	err := gen(".")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	fmt.Println("success")
}

func gen(src string) error {
	// 第一步找出符合条件的文件
	srcFiles, err := scanFiles(src)
	if err != nil {
		return err
	}
	// 第二步，AST 解析源代码文件，拿到 service definition 定义
	defs, err := parseFiles(srcFiles)
	if err != nil {
		return err
	}
	// 生成代码
	return genFiles(src, defs)
}

// 根据 defs 来生成代码
// src 是源代码所在目录，在测试里面它是 ./testdata
func genFiles(src string, defs []http2.ServiceDefinition) error {
	for _, d := range defs {
		filename := underscoreName(d.GenName())
		filename = filename + ".go"
		filePath, _ := filepath.Abs(src + "/" + filename)
		bs := &bytes.Buffer{}
		http2.Gen(bs, d)
		b := bs.Bytes()
		bs2, _ := format.Source(b)
		txt := string(bs2)
		txt2 := strings.Replace(txt, "\n", "\r\n", -1)

		var p1file *os.File
		p1file, _ = os.OpenFile(filePath, os.O_APPEND|os.O_CREATE, 0666)
		defer p1file.Close()                //延迟关闭，写在前面防止忘了
		p1writer := bufio.NewWriter(p1file) //使用带缓存的*Writer
		p1writer.WriteString(txt2)          //在调用WriterString方法时，内容先写入缓存
		p1writer.Flush()                    //调用flush方法，将缓存的数据真正写入到文件中
	}
	return nil
}

func parseFiles(srcFiles []string) ([]http2.ServiceDefinition, error) {
	defs := make([]http2.ServiceDefinition, 0, 20)
	for _, src := range srcFiles {
		//fmt.Println(src)
		// 你需要利用 annotation 里面的东西来扫描 src，然后生成 file
		var file annotation.File

		fset := token.NewFileSet()
		f, err := parser.ParseFile(fset, src, nil, parser.ParseComments)
		if err != nil {
		}
		tv := &annotation.SingleFileEntryVisitor{}
		ast.Walk(tv, f)
		file = tv.Get()

		for _, typ := range file.Types {
			_, ok := typ.Annotations.Get("HttpClient")
			if !ok {
				continue
			}
			def, err := parseServiceDefinition(file.Node.Name.Name, typ)
			if err != nil {
				return nil, err
			}
			defs = append(defs, def)
		}
	}
	return defs, nil
}

// 你需要利用 typ 来构造一个 http.ServiceDefinition
// 注意你可能需要检测用户的定义是否符合你的预期
func parseServiceDefinition(pkg string, typ annotation.Type) (http2.ServiceDefinition, error) {
	sd := http2.ServiceDefinition{}
	sd.Package = pkg
	sd.Name = typ.Annotations.Node.Name.Name
	for _, t4ans := range typ.Annotations.Ans {
		if "ServiceName" == t4ans.Key {
			sd.Name = t4ans.Value
		}
	}

	sd.Methods = make([]http2.ServiceMethod, 0)
	for _, t4f := range typ.Fields {
		t4sm := http2.ServiceMethod{}
		t4sm.Name = t4f.Annotations.Node.Names[0].Name
		t4sm.Path = "/" + t4sm.Name

		for _, t4ans := range t4f.Annotations.Ans {
			if "Path" == t4ans.Key {
				t4sm.Path = t4ans.Value
			}
		}

		t4t := t4f.Annotations.Node.Type
		t4ft := t4t.(*ast.FuncType)
		if 2 != len(t4ft.Params.List) {
			return sd, errors.New("gen: 方法必须接收两个参数，其中第一个参数是 context.Context，第二个参数请求")
		}
		for _, t4p := range t4ft.Params.List {
			t4param, ok := (t4p.Type).(*ast.StarExpr)
			if !ok {
				continue
			}
			x := (t4param.X).(*ast.Ident)
			t4sm.ReqTypeName = x.Name
		}

		if 2 != len(t4ft.Results.List) {
			return sd, errors.New("gen: 方法必须返回两个参数，其中第一个返回值是响应，第二个返回值是error")
		}
		for _, t4p := range t4ft.Results.List {
			t4param, ok := (t4p.Type).(*ast.StarExpr)
			if !ok {
				continue
			}
			x := (t4param.X).(*ast.Ident)
			t4sm.RespTypeName = x.Name
		}

		sd.Methods = append(sd.Methods, t4sm)
	}

	return sd, nil
}

// 返回符合条件的 Go 源代码文件，也就是你要用 AST 来分析这些文件的代码
func scanFiles(src string) ([]string, error) {
	s5file := make([]string, 0)

	s5fileName, err := os.ReadDir(src)
	if err != nil {
		log.Fatal(err)
	}
	for i := range s5fileName {
		if strings.Contains(s5fileName[i].Name(), "gen") {
			continue
		}
		if strings.Contains(s5fileName[i].Name(), ".go") {
			t4fp, _ := filepath.Abs(src + "/" + s5fileName[i].Name())
			s5file = append(s5file, t4fp)
		}
	}

	return s5file, nil
}

// underscoreName 驼峰转字符串命名，在决定生成的文件名的时候需要这个方法
// 可以用正则表达式，然而我写不出来，我是正则渣
func underscoreName(name string) string {
	var buf []byte
	for i, v := range name {
		if unicode.IsUpper(v) {
			if i != 0 {
				buf = append(buf, '_')
			}
			buf = append(buf, byte(unicode.ToLower(v)))
		} else {
			buf = append(buf, byte(v))
		}

	}
	return string(buf)
}
