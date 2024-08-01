package lexer

import (
	"fmt"
	"monky/token"
	"testing"
)

func TestName(t *testing.T) {
	input := `let five = 5;
	let ten = 10;
	let add = fn(x,y){
		x+y;
	};
	let result = add(five,ten);
`
	tests := []struct {
		expectedType    token.TokenType
		expectedLiteral string
	}{
		{token.ASSIGN, "="},
		{token.PLUS, "+"},
		{token.LPAREN, "("},
		{token.RPAREN, ")"},
		{token.LBRACE, "{"},
		{token.RBRACE, "}"},
		{token.EOF, ""},
	}

	// 返回一个lexer对象
	l := New(input)
	for i, tt := range tests {
		// 获取一个token对象
		tok := l.NextToken()
		fmt.Println(tt.expectedLiteral)
		fmt.Println(tok.Literal)
		if tok.Type != tt.expectedType {
			t.Fatalf("tests[%d] - tokentype wrong. expected=%q, got=%q", i, tt.expectedType, tok.Type)
		}
		if tok.Literal != tt.expectedLiteral {
			t.Fatalf("tests[%d] - Literal wrong. expected=%q, got=%q", i, tt.expectedLiteral, tok.Literal)
		}
	}
}
