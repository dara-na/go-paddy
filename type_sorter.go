package main

import (
	"go/ast"
)

type structFields struct {
	fields []*ast.Field
	source string
}

func (sf structFields) Len() int {
	return len(sf.fields)
}

func (sf structFields) Swap(i, j int) {
	sf.fields[i], sf.fields[j] = sf.fields[j], sf.fields[i]
}

func (sf structFields) Less(i, j int) bool {
	ti := sf.source[sf.fields[i].Type.Pos()-1 : sf.fields[i].Type.End()-1]
	tj := sf.source[sf.fields[j].Type.Pos()-1 : sf.fields[j].Type.End()-1]
	return GetGoType(ti) < GetGoType(tj)
}
