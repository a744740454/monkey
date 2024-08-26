package ast

import (
	"bytes"
	"monky/token"
)

type Node interface {
	// TokenLiteral 该方法必须实现，用于返回词法单元的值
	TokenLiteral() string
	String() string
}

// Statement 声明
type Statement interface {
	Node
	statementNode()
}

// Expression 表达式
type Expression interface {
	Node
	expressionNode()
}

// Program 当前程序的结构体，存放关键式声明
type Program struct {
	Statements []Statement
}

// LetStatement let的声明
type LetStatement struct {
	Token token.Token //词法单元，包含类型以及值
	Name  *Identifier // 绑定的标识符
	Value Expression  //值的表达式
}

// ReturnStatement let的声明
type ReturnStatement struct {
	Token       token.Token //词法单元，包含类型以及值
	ReturnValue Expression  //值的表达式
}

// ExpressionStatement 表达式的声明
type ExpressionStatement struct {
	Token      token.Token //词法单元，包含类型以及值
	Expression Expression  //值的表达式
}

// Identifier 标识符
type Identifier struct {
	Token token.Token
	Value string
}

func (p *Program) statementNode() {
	return
}

func (p *Program) TokenLiteral() string {
	if len(p.Statements) > 0 {
		return p.Statements[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer
	for _, s := range p.Statements {
		out.WriteString(s.String())
	}
	return out.String()

}

func (i *Identifier) expressionNode() {

}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

func (ls *LetStatement) statementNode() {
	return
}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
	//拼接字符串，好处是省内存
	var out bytes.Buffer
	out.WriteString(ls.TokenLiteral() + " ")
	out.WriteString(ls.Name.String())
	out.WriteString(" = ")
	if ls.Value != nil {
		out.WriteString(ls.Value.String())
	}
	out.WriteString(";")
	return out.String()
}

func (rs *ReturnStatement) statementNode() {
	return
}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	//拼接字符串，好处是省内存
	var out bytes.Buffer
	out.WriteString(rs.TokenLiteral() + " ")
	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}
	out.WriteString(";")
	return out.String()
}

func (es *ExpressionStatement) statementNode() {
	return
}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	//拼接字符串，好处是省内存
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}
