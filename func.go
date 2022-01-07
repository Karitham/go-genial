package genial

import (
	"bytes"
	"fmt"
	"log"
	"strings"
)

// Parameter is a go function parameter
type Parameter struct {
	Name     string
	Type     string
	Variadic bool
}

// FuncB is a function builder
type FuncB struct {
	comment    string
	name       string
	receiver   Parameter
	parameters []Parameter
	returnType []Parameter
	body       *bytes.Buffer
}

// Description returns the description of the function, (as in, the top level comment)
func (f *FuncB) Description() []byte {
	if f.comment == "" {
		return nil
	}

	b := &bytes.Buffer{}
	b.WriteString("// ")
	b.WriteString(commentSanitizer.Replace(f.comment))
	b.WriteString("\n")
	return b.Bytes()
}

// Signature returns the signature of the function
// Useful for the interface builder
func (f *FuncB) Signature() []byte {
	b := &bytes.Buffer{}

	b.WriteString(f.name)
	b.WriteRune('(')
	for i, p := range f.parameters {
		if i > 0 {
			b.WriteString(", ")
		}

		if p.Name != "" {
			b.WriteString(p.Name)
			b.WriteString(" ")
		}

		if p.Variadic {
			b.WriteString("...")
		}
		b.WriteString(p.Type)
	}
	b.WriteRune(')')
	b.WriteString(" ")

	// Return type
	if len(f.returnType) > 1 {
		b.WriteString("(")
	}
	for i, p := range f.returnType {
		if i > 0 {
			b.WriteString(", ")
		}

		if p.Name != "" {
			b.WriteString(p.Name)
			b.WriteString(" ")
		}
		b.WriteString(p.Type)
	}
	if len(f.returnType) > 1 {
		b.WriteRune(')')
	}

	return b.Bytes()
}

// Comment sets the comment of the function
func (f *FuncB) Comment(c string) *FuncB {
	f.comment = c
	return f
}

// Commentf sets the comment using fmt.Sprintf
func (f *FuncB) Commentf(format string, a ...interface{}) *FuncB {
	return f.Comment(fmt.Sprintf(format, a...))
}

// Name sets the name of the function
func (f *FuncB) Name(n string) *FuncB {
	f.name = n
	return f
}

// Namef sets the name of the function using fmt.Sprintf
func (f *FuncB) Namef(format string, a ...interface{}) *FuncB {
	return f.Name(fmt.Sprintf(format, a...))
}

// Receiver sets the receiver of the function
func (f *FuncB) Receiver(name string, typ string) *FuncB {
	f.receiver = Parameter{Name: name, Type: typ}
	return f
}

// Parameters sets the parameters of the function
func (f *FuncB) Parameters(p ...Parameter) *FuncB {
	for _, t := range p {
		if t.Name == "" {
			log.Panicf("Parameter name must not be empty for type: %s", t.Type)
		}
	}
	f.parameters = append(f.parameters, p...)
	return f
}

// Parameter adds a single parameter to the function
func (f *FuncB) Parameter(name string, typ string) *FuncB {
	return f.Parameters(Parameter{Name: name, Type: typ})
}

// ReturnTypes sets the return types of the function
func (f *FuncB) ReturnTypes(p ...string) *FuncB {
	for _, t := range p {
		f.returnType = append(f.returnType, Parameter{Type: t})
	}
	return f
}

// Write directly into the body of the function
func (f *FuncB) Write(b []byte) (n int, err error) {
	if f.body == nil {
		f.body = &bytes.Buffer{}
	}
	return f.body.Write(b)
}

// WriteString writes a string into the body of the function
func (f *FuncB) WriteString(s string) (n int, err error) {
	if f.body == nil {
		f.body = &bytes.Buffer{}
	}
	return f.body.WriteString(s)
}

// Writef writes directly to the body of the function using fmt.Sprintf
func (f *FuncB) Writef(format string, a ...interface{}) (n int, err error) {
	if f.body == nil {
		f.body = &bytes.Buffer{}
	}
	return fmt.Fprintf(f.body, format, a...)
}

var commentSanitizer = strings.NewReplacer("\n", "\n// ")

// String returns a string representation of the function
func (f *FuncB) String() string {
	return string(f.Bytes())
}

// Bytes returns the bytes of the function
func (f *FuncB) Bytes() []byte {
	b := &bytes.Buffer{}

	// top level comment
	b.Write(f.Description())

	b.WriteString("func ")

	// Method receiver
	if f.receiver.Name != "" {
		b.WriteRune('(')
		b.WriteString(f.receiver.Name)
		b.WriteString(" ")
		b.WriteString(f.receiver.Type)
		b.WriteRune(')')
		b.WriteString(" ")
	}

	b.Write(f.Signature())
	b.WriteString(" {\n")

	// Body
	if f.body != nil {
		b.Write(f.body.Bytes())
	}

	b.WriteString("}\n")
	return b.Bytes()
}
