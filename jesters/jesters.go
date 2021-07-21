package jesters

import "go/ast"

type jester interface {
	Jest(ast.Node, Tester)
}

type JesterFunc func(ast.Node, Tester)

func (j JesterFunc) Jest(n ast.Node, t Tester) {
	j(n, t)
}

var Jesters []jester

func Register(j jester) {
	Jesters = append(Jesters, j)
}
