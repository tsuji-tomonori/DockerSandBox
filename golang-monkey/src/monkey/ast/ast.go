package ast

import "monkey/token"

type Node interface {
	TokenLiteral() string
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
