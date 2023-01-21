package parser

import (
	"fmt"
	"testing"

	"github.com/hmdyt/monkey/ast"
)

func checkParserErrors(t *testing.T, p *Parser) {
	errors := p.Errors()
	if len(errors) == 0 {
		return
	}

	t.Errorf("parser has %d errors", len(errors))
	for _, msg := range errors {
		t.Errorf("parser error: %q", msg)
	}
	t.FailNow()
}

func testIntegerLiteral(t *testing.T, intLiteral ast.Expression, value int64) bool {
	integer, ok := intLiteral.(*ast.IntegerLiteral)
	if !ok {
		t.Errorf("il not *ast.IntegerLiteral. got=%T", intLiteral)
		return false
	}

	if integer.Value != value {
		t.Errorf("integer.Value not %d. got=%d", value, integer.Value)
		return false
	}

	if integer.TokenLiteral() != fmt.Sprintf("%d", value) {
		t.Errorf("integ.TokenLiteral not %d. got=%s", value,
			integer.TokenLiteral())
		return false
	}

	return true
}

func testIdentifier(t *testing.T, expression ast.Expression, value string) bool {
	identifier, ok := expression.(*ast.Identifier)
	if !ok {
		t.Errorf("expression not *ast.Identifier. got=%T", expression)
		return false
	}

	if identifier.Value != value {
		t.Errorf("identifier.Value not %s. got=%s", value, identifier.Value)
		return false
	}

	if identifier.TokenLiteral() != value {
		t.Errorf("identifier.Value not %s. got=%s", value, identifier.TokenLiteral())
		return false
	}

	return true
}

func testLiteralExpression(
	t *testing.T,
	expression ast.Expression,
	expected interface{},
) bool {
	switch v := expected.(type) {
	case int:
		return testIntegerLiteral(t, expression, int64(v))
	case int64:
		return testIntegerLiteral(t, expression, v)
	case string:
		return testIdentifier(t, expression, v)
	}
	t.Errorf("type of expression not handled. got=%T", expression)
	return false
}

func testInfixExpression(
	t *testing.T,
	expression ast.Expression,
	left interface{},
	operator string,
	right interface{},
) bool {
	infixExpression, ok := expression.(*ast.InfixExpression)
	if !ok {
		t.Errorf(
			"expression is not ast.InfixExpression. got=%T(%s)",
			expression,
			expression,
		)
	}

	if !testLiteralExpression(t, infixExpression.Left, left) {
		return false
	}

	if infixExpression.Operator != operator {
		t.Errorf(
			"expression.Operator is not '%s'. got=%q",
			operator,
			infixExpression.Operator,
		)
	}

	if !testLiteralExpression(t, infixExpression.Right, right) {
		return false
	}

	return true
}
