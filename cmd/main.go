package main

import (
	"flag"
	"fmt"
	go_jester "github.com/keanpedersen/go-jester"
)

func main() {
	flag.Parse()
	filename := ""
	if len(flag.Args()) > 0 {
		filename = flag.Arg(0)
	}

	if err := go_jester.Path(filename); err != nil {

		fmt.Printf("Stacktrace: %+v\n", err)

		panic(err)
	}
}
