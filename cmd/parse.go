package cmd

import (
	"go/ast"
	"go/parser"
	"go/token"
	"log"
)

func parse(dirname string) *Visitor {
	v := NewVisitor()
	pkgs := packages(dirname)
	for _, p := range pkgs {
		for _, f := range p.Files {
			ast.Walk(v, f)
		}
	}
	return v
}

func parseFile(filePath string) *Visitor {
	v := NewVisitor()
	fs := token.NewFileSet()
	file, err := parser.ParseFile(fs, filePath, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("could not parse file %s: %v", filePath, err)
	}
	ast.Walk(v, file)
	return v
}

func packages(dirname string) map[string]*ast.Package {
	fs := token.NewFileSet()
	pkgs, err := parser.ParseDir(fs, dirname, nil, parser.AllErrors)
	if err != nil {
		log.Fatalf("could not parse dir %s: %v", dirname, err)
	}
	return pkgs
}
