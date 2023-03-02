package ast

import (
	"bytes"
	"strings"

	"github.com/hmdyt/monkey/token"
)

type ArrayLiteral struct {
	Token    token.Token
	Elements []Expression
}

func (a ArrayLiteral) expressionNode()      {}
func (a ArrayLiteral) TokenLiteral() string { return a.Token.Literal }
func (a ArrayLiteral) String() string {
	var out bytes.Buffer

	var elements []string
	for _, element := range a.Elements {
		elements = append(elements, element.String())
	}

	out.WriteString("[")
	out.WriteString(strings.Join(elements, ", "))
	out.WriteString("]")

	return out.String()
}
