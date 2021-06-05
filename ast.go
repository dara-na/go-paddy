package main

import (
	"bytes"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"sort"
)

func in(s string, names ...string) bool {
	if len(names) == 0 {
		return true
	}
	for _, n := range names {
		if n == s {
			return true
		}
	}
	return false
}

func eachStructs(source string, f func(source string, st *ast.StructType), structs ...string) (string, error) {
	ts := token.NewFileSet()
	node, err := parser.ParseFile(ts, "", source, parser.AllErrors)
	if err != nil {
		return "", err
	}

	for _, decl := range node.Decls {
		if gd, ok := decl.(*ast.GenDecl); ok {
			if len(gd.Specs) > 0 {
				if ts, ok := gd.Specs[0].(*ast.TypeSpec); ok {
					if s, ok := ts.Type.(*ast.StructType); ok {
						if len(structs) == 0 || in(ts.Name.Name, structs...) {
							f(source, s)
						}
					}
				}
			}
		}
	}

	var buff bytes.Buffer

	if err := printer.Fprint(&buff, ts, node); err != nil {
		return "", err
	}

	return buff.String(), nil
}

func structTypeSorter(source string, st *ast.StructType) {
	if len(st.Fields.List) > 0 {
		sort.Sort(structFields{
			fields: st.Fields.List,
			source: source,
		})
	}
}
