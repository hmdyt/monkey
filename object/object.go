package object

import (
	"bytes"
	"fmt"
	"strings"

	"github.com/hmdyt/monkey/ast"
)

type Type string
type BuiltinFunction func(args ...Object) Object

const (
	IntegerObj     = "INTEGER"
	BooleanObj     = "BOOLEAN"
	StringObj      = "STRING"
	NullObj        = "NULL"
	ReturnValueObj = "RETURN_VALUE"
	ErrorObj       = "ERROR"
	FunctionObj    = "FUNCTION"
	BuiltinObj     = "BUILTIN"
	ArrayObj       = "ARRAY"
)

type Object interface {
	Type() Type
	Inspect() string
}

type Integer struct {
	Value int64
}

func (i *Integer) Type() Type      { return IntegerObj }
func (i *Integer) Inspect() string { return fmt.Sprintf("%d", i.Value) }

type Boolean struct {
	Value bool
}

func (b *Boolean) Type() Type      { return BooleanObj }
func (b *Boolean) Inspect() string { return fmt.Sprintf("%t", b.Value) }

type String struct {
	Value string
}

func (s *String) Type() Type      { return StringObj }
func (s *String) Inspect() string { return s.Value }

type Null struct{}

func (n *Null) Type() Type      { return NullObj }
func (n *Null) Inspect() string { return "null" }

type ReturnValue struct {
	Value Object
}

func (r *ReturnValue) Type() Type      { return ReturnValueObj }
func (r *ReturnValue) Inspect() string { return r.Value.Inspect() }

type Function struct {
	Parameters []*ast.Identifier
	Body       *ast.BlockStatement
	Env        *Environment
}

func (f *Function) Type() Type { return FunctionObj }
func (f *Function) Inspect() string {
	var out bytes.Buffer

	var params []string
	for _, p := range f.Parameters {
		params = append(params, p.String())
	}

	out.WriteString("fn")
	out.WriteString("(")
	out.WriteString(strings.Join(params, ", "))
	out.WriteString(") {\n")
	out.WriteString(f.Body.String())
	out.WriteString("\n}")

	return out.String()
}

type Error struct {
	Message string
}

func (e Error) Type() Type      { return ErrorObj }
func (e Error) Inspect() string { return "ERROR: " + e.Message }

type Builtin struct {
	Function BuiltinFunction
}

func (b *Builtin) Type() Type      { return BuiltinObj }
func (b *Builtin) Inspect() string { return "builtin function" }

type Array struct {
	Elements []Object
}

func (a *Array) Type() Type { return ArrayObj }
func (a *Array) Inspect() string {
	var out bytes.Buffer

	var elements []string
	for _, e := range a.Elements {
		elements = append(elements, e.Inspect())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
