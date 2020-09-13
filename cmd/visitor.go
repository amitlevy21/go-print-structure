package cmd

import (
	"fmt"
	"go/ast"
	"go/types"
)

type Visitor struct {
	s map[string][]*FieldData
}

type FieldData struct {
	name   string
	myType string
}

func NewVisitor() *Visitor {
	return &Visitor{make(map[string][]*FieldData)}
}

func NewFieldData(name, myType string) *FieldData {
	return &FieldData{name, myType}
}

func (v Visitor) Visit(n ast.Node) ast.Visitor {
	if n == nil {
		return nil
	}
	d, ok := n.(*ast.GenDecl)
	if !ok {
		return v
	}
	v.structs(d)
	return v
}

func (v Visitor) structs(d *ast.GenDecl) {
	for _, spec := range d.Specs {
		if t, ok := spec.(*ast.TypeSpec); ok {
			if st, ok := t.Type.(*ast.StructType); ok {
				// v.visitNestedStructs(t.Name.Name, st)
				for _, l := range st.Fields.List {
					fd := NewFieldData(l.Names[0].Name, types.ExprString(l.Type))
					v.s[t.Name.Name] = append(v.s[t.Name.Name], fd)
				}
			}
		}
	}
}

// func (v Visitor) visitNestedStructs(name string, st *ast.StructType) {
// 	if v.s[name] == nil {
// 		v.s[name] = NewMyVar(name, name)
// 	}
// 	for _, l := range st.Fields.List {
// 		v.s[name].fields[l.Names[0].Name] = &myVar{l.Names[0].Name, l.Names[0].Name, make(map[string]*myVar)}
// 	}
// }

func (v *Visitor) String() string {
	return fmt.Sprintf("%v", v.s)
}

func (v *FieldData) String() string {
	return fmt.Sprintf("%s, %s", v.name, v.myType)
}
