# go-genial

[![Go Reference](https://pkg.go.dev/badge/github.com/Karitham/go-genial.svg)](https://pkg.go.dev/github.com/Karitham/go-genial)

A prototype code-generator library for golang.

## Example

```go
package main

import (
	"fmt"

	"github.com/karitham/go-genial"
)

func main() {
	f := &genial.FuncBuilder{}
	f.Comment("FooBar is a new example function")
	f.Name("FooBar").Parameters(
			genial.Parameter{
				Name: "foo",
				Type: "int",
			}, genial.Parameter{
				Name: "bar",
				Type: "string",
			},
		)

	f.Write([]byte("\tpanic(\"not implemented\")\n"))

	i := &genial.IfaceBuilder{}
	i.Comment("Iface is an example interface").Functions(f).Name("Iface")

	fmt.Println(f.String())
	fmt.Println(i.String())
}
```

generates

```go
// FooBar is a new example function
func FooBar(foo int, bar string) {
	panic("not implemented")
}

// Iface is an example interface
type Iface interface {
	// FooBar is a new example function
	FooBar(foo int, bar string)
}
```
