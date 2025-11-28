package lexer

import (
	"testing"

	"monkey/token"
)

func TestNextToken(t *testing.T) {
	input := `let five = 5;
let ten = 10;
let add = fn(x, y) {
x + y;
};
let result = add(five, ten);
!-/*5;
5 < 10 > 5;
if (5 < 10) {
return true;
} else {
return false;
}
10 == 10;
10 != 9;
"foobar"
"foo bar"
[1, 2];
{"foo": "bar"}
let 🏠 = "house";
macro(x, y) { x + y; };
`

	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.LET, string("let")},
		{token.IDENT, string("five")},
		{token.ASSIGN, string("=")},
		{token.INT, string("5")},
		{token.SEMICOLON, string(";")},
		{token.LET, string("let")},
		{token.IDENT, string("ten")},
		{token.ASSIGN, string("=")},
		{token.INT, string("10")},
		{token.SEMICOLON, string(";")},
		{token.LET, string("let")},
		{token.IDENT, string("add")},
		{token.ASSIGN, string("=")},
		{token.FUNCTION, string("fn")},
		{token.LPAREN, string("(")},
		{token.IDENT, string("x")},
		{token.COMMA, string(",")},
		{token.IDENT, string("y")},
		{token.RPAREN, string(")")},
		{token.LBRACE, string("{")},
		{token.IDENT, string("x")},
		{token.PLUS, string("+")},
		{token.IDENT, string("y")},
		{token.SEMICOLON, string(";")},
		{token.RBRACE, string("}")},
		{token.SEMICOLON, string(";")},
		{token.LET, string("let")},
		{token.IDENT, string("result")},
		{token.ASSIGN, string("=")},
		{token.IDENT, string("add")},
		{token.LPAREN, string("(")},
		{token.IDENT, string("five")},
		{token.COMMA, string(",")},
		{token.IDENT, string("ten")},
		{token.RPAREN, string(")")},
		{token.SEMICOLON, string(";")},
		{token.BANG, string("!")},
		{token.MINUS, string("-")},
		{token.SLASH, string("/")},
		{token.ASTERISK, string("*")},
		{token.INT, string("5")},
		{token.SEMICOLON, string(";")},
		{token.INT, string("5")},
		{token.LT, string("<")},
		{token.INT, string("10")},
		{token.GT, string(">")},
		{token.INT, string("5")},
		{token.SEMICOLON, string(";")},
		{token.IF, string("if")},
		{token.LPAREN, string("(")},
		{token.INT, string("5")},
		{token.LT, string("<")},
		{token.INT, string("10")},
		{token.RPAREN, string(")")},
		{token.LBRACE, string("{")},
		{token.RETURN, string("return")},
		{token.TRUE, string("true")},
		{token.SEMICOLON, string(";")},
		{token.RBRACE, string("}")},
		{token.ELSE, string("else")},
		{token.LBRACE, string("{")},
		{token.RETURN, string("return")},
		{token.FALSE, string("false")},
		{token.SEMICOLON, string(";")},
		{token.RBRACE, string("}")},
		{token.INT, string("10")},
		{token.EQ, string("==")},
		{token.INT, string("10")},
		{token.SEMICOLON, string(";")},
		{token.INT, string("10")},
		{token.NOT_EQ, string("!=")},
		{token.INT, string("9")},
		{token.SEMICOLON, string(";")},
		{token.STRING, string("foobar")},
		{token.STRING, string("foo bar")},
		{token.LBRACKET, string("[")},
		{token.INT, string("1")},
		{token.COMMA, string(",")},
		{token.INT, string("2")},
		{token.RBRACKET, string("]")},
		{token.SEMICOLON, string(";")},
		{token.LBRACE, string("{")},
		{token.STRING, string("foo")},
		{token.COLON, string(":")},
		{token.STRING, string("bar")},
		{token.RBRACE, string("}")},
		{token.LET, string("let")},
		{token.IDENT, string("🏠")},
		{token.ASSIGN, string("=")},
		{token.STRING, string("house")},
		{token.SEMICOLON, string(";")},
		{token.MACRO, "macro"},
		{token.LPAREN, "("},
		{token.IDENT, "x"},
		{token.COMMA, ","},
		{token.IDENT, "y"},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.IDENT, "x"},
		{token.PLUS, "+"},
		{token.IDENT, "y"},
		{token.SEMICOLON, ";"},
		{token.RBRACE, "}"},
		{token.SEMICOLON, ";"},
		{token.EOF, ""},
	}

	l := New(input)

	for i, tt := range tests {
		tok := l.NextToken()

		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q",
				i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - literal wrong. expected=%q, got=%q",
				i, tt.expectedLiteral, tok.Literal)
		}
	}
}
