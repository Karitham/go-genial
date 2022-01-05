package genial

import (
	"fmt"
	"strings"
)

// Interface is a go interface
type Interface interface {
	Name(string) Interface
	Namef(string, ...interface{}) Interface
	Comment(string) Interface
	Commentf(string, ...interface{}) Interface
	Members(...Signaturer) Interface

	String() string
}

// Signaturer is implemented by functions,
// which enables us to pass them to the interface builder
type Signaturer interface {
	Description() string
	Signature() string
}

// InterfaceB is an interface builder
type InterfaceB struct {
	comment     string
	name        string
	signaturers []Signaturer
}

// Members adds members to the interface
func (i *InterfaceB) Members(s ...Signaturer) Interface {
	i.signaturers = append(i.signaturers, s...)
	return i
}

// Comment sets the comment on the interface
func (i *InterfaceB) Comment(comment string) Interface {
	i.comment = comment
	return i
}

// Commentf sets the comment on the interface using fmt.Sprintf
func (i *InterfaceB) Commentf(format string, args ...interface{}) Interface {
	return i.Comment(fmt.Sprintf(format, args...))
}

// Name sets the name of the interface
func (i *InterfaceB) Name(n string) Interface {
	i.name = n
	return i
}

// Namef sets the name of the interface using fmt.Sprintf
func (i *InterfaceB) Namef(format string, args ...interface{}) Interface {
	return i.Name(fmt.Sprintf(format, args...))
}

// String returns a string representation of the iface
func (i *InterfaceB) String() string {
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
