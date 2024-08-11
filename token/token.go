package token

//这个是最小的单元，会将用户输入的文字，按对应规则，依次创建对应的tok对象，例如：x := 123

type TokenType string
type Token struct {
	Type    TokenType //当前字符的类型
	Literal string    // "literal" 是一个术语，用于指代在源代码中直接写出的固定值
}

const (
	// 未知语法
	ILLEGAL = "ILLEGAL"
	// 代码结尾
	EOF = "EOF"
	//标识符+字面量
	IDENT = "IDENT" // add, foobar, x, y, ...￼
	INT   = "INT"   //
	//运算符￼
	ASSIGN   = "="
	PLUS     = "+"
	MINUS    = "-"
	BANG     = "!"
	ASTERISK = "*"
	SLASH    = "/"
	LT       = "<"
	GT       = ">"
	EQ       = "=="
	NOT_EQ   = "!="

	//分隔符￼
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"

	// 关键字
	FUNCTION = "FUNCTION"
	LET      = "LET"
	TRUE     = "TRUE"
	FALSE    = "FALSE"
	IF       = "IF"
	ELSE     = "ELSE"
	RETURN   = "RETURN"
)

// 定义一个map，根据关键字返回字符串对象
var keywords = map[string]TokenType{
	"fn":     FUNCTION,
	"let":    LET,
	"true":   TRUE,
	"false":  FALSE,
	"if":     IF,
	"else":   ELSE,
	"return": RETURN,
}

func LookupIdent(ident string) TokenType {
	// 判断是不是关键字
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
