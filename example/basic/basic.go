package main

import (
	"os"

	"github.com/karitham/go-genial"
)

func main() {
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
}
