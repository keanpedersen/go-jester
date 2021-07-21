package go_jester

import (
	"bytes"
	"fmt"
	"github.com/keanpedersen/go-jester/jesters"
	_ "github.com/keanpedersen/go-jester/jesters"
	"github.com/pkg/errors"
	"go/ast"
	"go/parser"
	"go/printer"
	"go/token"
	"io/ioutil"
	"os"
	"os/exec"
	"strings"
)

// Path runs the jester on all non-test files in the given path.
// The path must be reachable as a target for `go test ./[path]` in order
// to use the test cases to determine the outcome of the jester.
func Path(path string) error {

	fileSet := token.NewFileSet()

	astDir, err := parser.ParseDir(fileSet, path, func(info os.FileInfo) bool {
		return !strings.HasSuffix(info.Name(), "_test.go")
	}, parser.ParseComments)
	if err != nil {
		return errors.WithStack(err)
	}

	testFunc := func(position token.Pos, original string, jested string) {
		passed, err := tryTests(path, fileSet, astDir)
		if err != nil {
			panic(err) // todo - bubble/log?
		}
		if passed {
			// TODO: print source code line with pointers
			pos := fileSet.Position(position)
			fmt.Printf("File %v, Line %v, column %v", pos.Filename, pos.Line, pos.Column)
			fmt.Printf(" - changing `%v` to `%v` has no test coverage\n", original, jested)
		}
	}

	// jest away
	for _, packages := range astDir {
		for _, file := range packages.Files {
			ast.Inspect(file, func(n ast.Node) bool {
				for _, jester := range jesters.Jesters {
					jester.Jest(n, testFunc)
				}
				return true
			})

		}
	}

	// Rewrite the original source again at the end
	return writeSource(fileSet, astDir)
}

func render(fset *token.FileSet, x interface{}) string {
	var buf bytes.Buffer
	if err := printer.Fprint(&buf, fset, x); err != nil {
		panic(err)
	}
	return buf.String()
}

func writeSource(fileSet *token.FileSet, astDir map[string]*ast.Package) error {
	for _, packages := range astDir {
		// Print output
		for fileName, file := range packages.Files {

			outFile, err := os.OpenFile(fileName, os.O_TRUNC|os.O_WRONLY, 0)
			if err != nil {
				return errors.WithStack(err)
			}

			if err := printer.Fprint(outFile, fileSet, file); err != nil {
				return errors.WithStack(err)
			}

			if err := outFile.Close(); err != nil {
				return errors.WithStack(err)
			}
		}
	}
	return nil
}

func tryTests(path string, fileSet *token.FileSet, astDir map[string]*ast.Package) (passed bool, err error) {
	if err := writeSource(fileSet, astDir); err != nil {
		return passed, err
	}

	return runTests(path)
}

func runTests(path string) (passed bool, err error) {

	// run the test
	cmd := exec.Command("go", "test", "./"+path)
	cmd.Stderr = ioutil.Discard
	_, err = cmd.Output()
	if err != nil {
		switch err.(type) {
		case *exec.ExitError:
			return false, nil
		default:
			return passed, errors.WithStack(err)
		}
	}

	return true, nil
}
