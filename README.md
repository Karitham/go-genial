# go-genial

[![Go Reference](https://pkg.go.dev/badge/github.com/Karitham/go-genial.svg)](https://pkg.go.dev/github.com/Karitham/go-genial)

A prototype code-generator library for golang.

## Install

`go get github.com/karitham/go-genial`

## Example

```go
	t := &genial.StructB{}
	t.Comment("Baz is a implementation of Iface").
		Name("Baz").
		Fields(genial.Field{
			Name: "Foo",
			Type: "*string",
			Tag:  []genial.StructTag{{Type: "json", Value: "foo,omitempty"}},
		},
		)

	f := &genial.FuncB{}
	f.Comment("FooBar is a new example function").
		Name("FooBar").
		Receiver("b", "*Baz").
		Parameters(
			genial.Parameter{Name: "foo", Type: "int"},
			genial.Parameter{Name: "bar", Type: "string"},
		).
		ReturnTypes("int", "error").
		WriteString("\tpanic(\"not implemented\")\n")

	i := &genial.InterfaceB{}
	i.Comment("Iface is an example interface").
		Members(f).
		Name("Iface")

	p := &genial.PackageB{}
	p.Comment("example is an example package").
		Declarations(t, i, f).
		Name("example")

	fmt.Println(p.String())
```

generates

```go
// example is an example package
package example

import "encoding/json"

// Baz is a implementation of Iface
type Baz struct {
	Foo *string `json:"foo,omitempty"`
	rest json.Raw
}

// Iface is an example interface
type Iface interface {
	// FooBar is a new example function
	FooBar(foo int, bar string) (int, error)
}

// FooBar is a new example function
func (b *Baz) FooBar(foo int, bar string) (int, error) {
	panic("not implemented")
}
```
