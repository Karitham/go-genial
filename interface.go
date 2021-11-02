package genial

import (
	"strings"
)

type Interface interface {
	Name(string) Interface
	Comment(string) Interface
	Functions(...Signaturer) Interface

	String() string
}

type Signaturer interface {
	Description() string
	Signature() string
}

type IfaceBuilder struct {
	comment     string
	name        string
	signaturers []Signaturer
}

func (i *IfaceBuilder) Functions(s ...Signaturer) Interface {
	i.signaturers = append(i.signaturers, s...)
	return i
}

func (ifb *IfaceBuilder) Comment(comment string) Interface {
	ifb.comment = comment
	return ifb
}

func (ifb *IfaceBuilder) Name(n string) Interface {
	ifb.name = n
	return ifb
}

// String returns a string representation of the iface
func (i *IfaceBuilder) String() string {
	b := &strings.Builder{}

	// comment
	if i.comment != "" {
		b.WriteString("// ")
		b.WriteString(commentSanitizer.Replace(i.comment))
		b.WriteString("\n")
	}

	// top level
	b.WriteString("type ")
	b.WriteString(i.name)
	b.WriteString(" interface {\n")

	// functions

	for _, s := range i.signaturers {
		c := s.Description()
		if c != "" {
			b.WriteString("\t")
			b.WriteString(c)
		}

		b.WriteString("\t")
		b.WriteString(s.Signature())
		b.WriteString("\n")
	}

	b.WriteString("}\n")
	return b.String()
}
