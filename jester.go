package go_jester

import (
	"fmt"
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

	// jest something
	for _, packages := range astDir {
		for _, file := range packages.Files {

			ast.Inspect(file, func(n ast.Node) bool {
				ifStmt, ok := n.(*ast.IfStmt)
				if !ok {
					return true
				}

				binary, ok := ifStmt.Cond.(*ast.BinaryExpr)
				if !ok {
					return true
				}

				binary.Op = token.NEQ
				// TODO: For each binary.Op, try to build a list of things that can be jestered.
				// For each thing that can be jestered: Remember the old version, change it, save, run tests, and revert to the old version

				return true
			})

		}
	}

	// Write out all files, ready for testing
	for pkg, packages := range astDir {
		// Print output
		fmt.Printf("Package: %v\n", pkg)
		for fileName, file := range packages.Files {
			fmt.Printf("Filename: %v\n", fileName)

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

	passed, err := runTests(path)
	if err != nil {
		return err
	}

	fmt.Printf("Passed: %v\n", passed)

	return nil
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
