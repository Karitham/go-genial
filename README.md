# go-genial

[![Go Reference](https://pkg.go.dev/badge/github.com/Karitham/go-genial.svg)](https://pkg.go.dev/github.com/Karitham/go-genial)

A golang code-generation library

## Install

`go get github.com/karitham/go-genial`

## Example

```go
	p := &genial.PackageB{}
	p.Comment("example is an example package").
		Name("example").
		Imports("encoding/json")

	t := &genial.StructB{}
	t.Comment("Baz is a implementation of Iface").
		Name("Baz").
		Field("Foo", "*string", genial.StructTag{Type: "json", Value: "foo,omitempty"}).
		Field("rest", "json.Raw")

	f := &genial.FuncB{}
	f.Comment("FooBar is a new example function").
		Name("FooBar").
		Receiver("b", "*Baz").
		Parameter("foo", "int").
		Parameter("bar", "string").
		ReturnTypes("int", "error").
		WriteString(`panic("not implemented")`)

	i := &genial.InterfaceB{}
	i.Comment("Iface is an example interface").
		Members(f).
		Name("Iface")

	p.Declarations(t, i, f).WriteTo(os.Stdout)
```

generates

```go
// example is an example package
package example

import "encoding/json"

// Baz is a implementation of Iface
type Baz struct {
	Foo  *string `json:"foo,omitempty"`
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
