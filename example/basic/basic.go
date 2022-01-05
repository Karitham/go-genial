package main

import (
	"fmt"

	"github.com/karitham/go-genial"
)

func main() {
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
}
