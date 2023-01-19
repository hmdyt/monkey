package token

type TokenType string

const (
	ILLEGAL = "ILLEGAL"
	EOF     = "EOF"

	// 識別子, リテラル
	INDENT = "INDENT"
	INT    = "INT"

	// operator
	ASSIGN = "="
	PLUS   = "+"

	// delimitor
	COMMA     = ","
	SEMICOLON = ";"

	// brackets
	LPAREN = "("
	RPAREN = ")"
	LBRACE = "{"
	RBRACE = "}"

	// keywords
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

type Token struct {
	Type    TokenType
	Literal string
}

var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdentifier(indent string) TokenType {
	if tok, ok := keywords[indent]; ok {
		return tok
	} else {
		return INDENT
	}
}
