package ast

import (
	"bytes"
	"strings"

	"github.com/hmdyt/monkey/token"
)

type CallExpression struct {
	Token     token.Token // '(' token
	Function  Expression  // Identifier or FunctionLiteral
	Arguments []Expression
}

func (c *CallExpression) expressionNode()      {}
func (c *CallExpression) TokenLiteral() string { return c.Token.Literal }
func (c *CallExpression) String() string {
	var out bytes.Buffer

	var args []string
	for _, a := range c.Arguments {
		args = append(args, a.String())
	}

	out.WriteString(c.Function.String())
	out.WriteString("(")
	out.WriteString(strings.Join(args, ", "))
	out.WriteString(")")

	return out.String()
}
