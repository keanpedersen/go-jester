package jesters

import (
	"go/ast"
	"go/token"
)

func init() {
	Register(JesterFunc(equalityJester))
}

func equalityJester(n ast.Node, t Tester) {
	binaryExpr, ok := n.(*ast.BinaryExpr)
	if !ok {
		return
	}

	var newOp token.Token
	switch binaryExpr.Op {
	case token.NEQ:
		newOp = token.EQL
	case token.EQL:
		newOp = token.NEQ
	default:
		return
	}

	oldOp := binaryExpr.Op
	binaryExpr.Op = newOp
	t(binaryExpr.Pos(), oldOp.String(), newOp.String())
	binaryExpr.Op = oldOp
}
