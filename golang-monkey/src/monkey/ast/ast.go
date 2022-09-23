// 抽象構文木(Abstract syntax tree)
// 構文解析器の出力
package ast

import (
	"bytes"
	"monkey/token"
)

type Node interface {
	TokenLiteral() string
	String() string
}

// 文: 値を生成しない
type Statement interface {
	Node
	statementNode()
}

// 式: 値を生成する
type Expression interface {
	Node
	expressionNode()
}

type Program struct {
	// root ノード
	Statement []Statement
}

func (p *Program) TokenLiteral() string {
	if len(p.Statement) > 0 {
		return p.Statement[0].TokenLiteral()
	} else {
		return ""
	}
}

func (p *Program) String() string {
	var out bytes.Buffer

	for _, s := range p.Statement {
		out.WriteString(s.String())
	}

	return out.String()
}

type LetStatement struct {
	Token token.Token // token.LET トークン
	Name  *Identifier // 識別のため
	Value Expression  // let文の中で値を生成する式のため
}

func (ls *LetStatement) statementNode() {

}

func (ls *LetStatement) TokenLiteral() string {
	return ls.Token.Literal
}

func (ls *LetStatement) String() string {
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

type Identifier struct {
	// let x = 5;におけるxのような束縛の識別子を保持する
	Token token.Token // token.IDENT トークン
	Value string
}

func (i *Identifier) expressionNode() {
	// ノードの種類を少なく保つため変数束縛の名前を表現するために式として扱う
}

func (i *Identifier) TokenLiteral() string {
	return i.Token.Literal
}

func (i *Identifier) String() string {
	return i.Value
}

type ReturnStatement struct {
	Token       token.Token // token.RETURN トークン
	ReturnValue Expression
}

func (rs *ReturnStatement) statementNode() {}

func (rs *ReturnStatement) TokenLiteral() string {
	return rs.Token.Literal
}

func (rs *ReturnStatement) String() string {
	var out bytes.Buffer

	out.WriteString(rs.TokenLiteral() + " ")

	if rs.ReturnValue != nil {
		out.WriteString(rs.ReturnValue.String())
	}

	out.WriteString(";")

	return out.String()
}

// 式文
type ExpressionStatement struct {
	Token      token.Token // 式の最初のトークン
	Expression Expression
}

func (es *ExpressionStatement) statementNode() {

}

func (es *ExpressionStatement) TokenLiteral() string {
	return es.Token.Literal
}

func (es *ExpressionStatement) String() string {
	if es.Expression != nil {
		return es.Expression.String()
	}
	return ""
}

// 整数リテラル
type IntegerLiteral struct {
	Token token.Token
	Value int64
}

func (il *IntegerLiteral) expressionNode() {

}

func (il *IntegerLiteral) TokenLiteral() string {
	return il.Token.Literal
}

func (il *IntegerLiteral) String() string {
	return il.Token.Literal
}

// 前置詞
type PrefixExpression struct {
	Token    token.Token // 前置トークン 例えば「!」
	Operator string      // ! or - を格納する
	Right    Expression  // 演算子の右側の式を格納する
}

func (pe *PrefixExpression) expressionNode() {

}

func (pe *PrefixExpression) TokenLiteral() string {
	return pe.Token.Literal
}

func (pe *PrefixExpression) String() string {
	var out bytes.Buffer

	out.WriteString("(")
	out.WriteString(pe.Operator)
	out.WriteString(pe.Right.String())
	out.WriteString(")")

	return out.String()
}
