package genial

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"strings"
)

// Function is a function
type Function interface {
	// to write the body
	io.Writer
	Writef(format string, a ...interface{}) (n int, err error)
	WriteString(s string) (n int, err error)

	Name(string) Function
	Namef(string, ...interface{}) Function

	Comment(string) Function
	Commentf(string, ...interface{}) Function

	Receiver(name string, typ string) Function

	Parameters(...Parameter) Function
	Parameter(name string, typ string) Function

	ReturnTypes(...string) Function

	Description() string
	Signature() string

	String() string
}

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
func (f *FuncB) Description() string {
	if f.comment == "" {
		return ""
	}

	b := &bytes.Buffer{}
	b.WriteString("// ")
	b.WriteString(commentSanitizer.Replace(f.comment))
	b.WriteString("\n")
	return b.String()
}

// Signature returns the signature of the function
// Useful for the interface builder
func (f *FuncB) Signature() string {
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

	return b.String()
}

// Comment sets the comment of the function
func (f *FuncB) Comment(c string) Function {
	f.comment = c
	return f
}

// Commentf sets the comment using fmt.Sprintf
func (f *FuncB) Commentf(format string, a ...interface{}) Function {
	return f.Comment(fmt.Sprintf(format, a...))
}

// Name sets the name of the function
func (f *FuncB) Name(n string) Function {
	f.name = n
	return f
}

// Namef sets the name of the function using fmt.Sprintf
func (f *FuncB) Namef(format string, a ...interface{}) Function {
	return f.Name(fmt.Sprintf(format, a...))
}

// Receiver sets the receiver of the function
func (f *FuncB) Receiver(name string, typ string) Function {
	f.receiver = Parameter{Name: name, Type: typ}
	return f
}

// Parameters sets the parameters of the function
func (f *FuncB) Parameters(p ...Parameter) Function {
	for _, t := range p {
		if t.Name == "" {
			log.Panicf("Parameter name must not be empty for type: %s", t.Type)
		}
	}
	f.parameters = append(f.parameters, p...)
	return f
}

// Parameter adds a single parameter to the function
func (f *FuncB) Parameter(name string, typ string) Function {
	return f.Parameters(Parameter{Name: name, Type: typ})
}

// ReturnTypes sets the return types of the function
func (f *FuncB) ReturnTypes(p ...string) Function {
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
	return f.body.WriteString(fmt.Sprintf(format, a...))
}

var commentSanitizer = strings.NewReplacer("\n", "\n// ")

// String returns a string representation of the function
func (f *FuncB) String() string {
	b := &strings.Builder{}

	// top level comment
	b.WriteString(f.Description())

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

	b.WriteString(f.Signature())

	b.WriteString(" {\n")

	// Body
	if f.body != nil {
		b.Write(f.body.Bytes())
	}

	b.WriteString("}\n")
	return b.String()
}
