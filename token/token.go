package token

//这个是最小的单元，会将用户输入的文字，按对应规则，依次创建对应的tok对象，例如：x := 123

type TokenType string
type Token struct {
	Type    TokenType //当前字符的类型
	Literal string    // 当前字符
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
	ASSIGN = "="
	PLUS   = "+"
	//分隔符￼
	COMMA     = ","
	SEMICOLON = ";"
	LPAREN    = "("
	RPAREN    = ")"
	LBRACE    = "{"
	RBRACE    = "}"
	// 关键字￼
	FUNCTION = "FUNCTION"
	LET      = "LET"
)

// 定义一个map，根据关键字返回字符串对象
var keywords = map[string]TokenType{
	"fn":  FUNCTION,
	"let": LET,
}

func LookupIdent(ident string) TokenType {
	if tok, ok := keywords[ident]; ok {
		return tok
	}
	return IDENT
}
