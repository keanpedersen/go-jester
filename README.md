# go-jester

When making automatic unit tests, it is not enough to just look at the test coverage percentage. 
Certain scenarios and code paths can be untested, even with 100% test coverage, giving a false
sense of coverage.

`go-jester` will take your go package and test suite, perform changes to your code, and tell
you what changes did not result in a failed test suite. This will indicate where you have
a lack of test coverage.

# Usage
`go-jester [package]`

You should call `go-jester` from a working directory, so that `[package]` is both pointing at your
source files as a relative directory path, and can be used as a package name for `go test ./[package]`.

Beware that `go-jester` will alter your source code multiple times and run your test suite multiple times.
If you stop `go-jester` (or it crashes), your source code might end up in an altered state, so be sure to only
run `go-jester` on a backup, or so that you have a method for reverting back to a working state. `git` is recommended
for this.