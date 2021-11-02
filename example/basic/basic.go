package main

import (
	"fmt"

	"github.com/karitham/go-genial"
)

func main() {
	t := &genial.StructBuilder{}
	t.Comment("Baz is a implementation of Iface").
		Name("Baz").
		Fields(genial.Field{
			Name: "Foo",
			Type: "*string",
			Tag: []genial.StructTag{
				{Type: "json", Value: "foo", Omitempty: true},
			},
		},
		)

	f := &genial.FuncBuilder{}
	f.Comment("FooBar is a new example function").
		Name("FooBar").
		Receiver(genial.Parameter{
			Name: "b",
			Type: "*Baz",
		}).
		Parameters(
			genial.Parameter{
				Name: "foo",
				Type: "int",
			}, genial.Parameter{
				Name: "bar",
				Type: "string",
			},
		).
		ReturnTypes(
			genial.Parameter{
				Type: "int",
			},
			genial.Parameter{
				Type: "error",
			},
		).
		Write([]byte("\tpanic(\"not implemented\")\n"))

	i := &genial.IfaceBuilder{}
	i.Comment("Iface is an example interface").
		Functions(f).Name("Iface")

	p := &genial.PackageBuilder{}

	p.Comment("example is an example package").
		Blocks(t, i, f).
		Name("example")

	fmt.Println(p.String())
}
