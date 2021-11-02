package main

import (
	"fmt"

	"github.com/karitham/go-genial"
)

func main() {
	f := &genial.FuncBuilder{}
	f.Comment("FooBar is a new example function").
		Name("FooBar").
		Parameters(
			genial.Parameter{
				Name: "foo",
				Type: "int",
			}, genial.Parameter{
				Name: "bar",
				Type: "string",
			},
		).Write([]byte("\tpanic(\"not implemented\")\n"))

	i := &genial.IfaceBuilder{}
	i.Comment("Iface is an example interface").
		Functions(f).Name("Iface")

	fmt.Println(f.String())
	fmt.Println(i.String())
}
