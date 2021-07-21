package jesters

import "go/token"

type Tester func(position token.Pos, original string, jested string)
