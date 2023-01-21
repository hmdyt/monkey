package ast

import (
	"bytes"

	"github.com/hmdyt/monkey/token"
)

type IfExpression struct {
	Token       token.Token
	Condition   Expression
	Consequence *BlockStatement
	Alternative *BlockStatement
}

func (i *IfExpression) expressionNode()      {}
func (i *IfExpression) TokenLiteral() string { return i.Token.Literal }
func (i *IfExpression) String() string {
	var out bytes.Buffer

	out.WriteString("if")
	out.WriteString(i.Condition.String())
	out.WriteString(" ")
	out.WriteString(i.Consequence.String())

	if i.Alternative != nil {
		out.WriteString("else ")
		out.WriteString(i.Alternative.String())
	}

	return out.String()
}
