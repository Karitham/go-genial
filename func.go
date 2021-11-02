package genial

import (
	"bytes"
	"io"
	"log"
	"strings"
)

type Function interface {
	// to write the body
	io.Writer

	Name(string) Function
	Comment(string) Function
	Receiver(Parameter) Function
	Parameters(...Parameter) Function
	ReturnTypes(...Parameter) Function

	Description() string
	Signature() string

	String() string
}

type Parameter struct {
	Name     string
	Type     string
	Variadic bool
}

type FuncBuilder struct {
	comment    string
	name       string
	receiver   Parameter
	parameters []Parameter
	returnType []Parameter
	body       *bytes.Buffer
}

func (f *FuncBuilder) Description() string {
	if f.comment == "" {
		return ""
	}

	b := &bytes.Buffer{}
	b.WriteString("// ")
	b.WriteString(commentSanitizer.Replace(f.comment))
	b.WriteString("\n")
	return b.String()
}

func (f *FuncBuilder) Signature() string {
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

func (f *FuncBuilder) Comment(c string) Function {
	f.comment = c
	return f
}

func (f *FuncBuilder) Name(n string) Function {
	f.name = n
	return f
}

func (f *FuncBuilder) Receiver(p Parameter) Function {
	f.receiver = p
	return f
}

func (f *FuncBuilder) Parameters(p ...Parameter) Function {
	for _, t := range p {
		if t.Name == "" {
			log.Panicf("Parameter name must not be empty for type: %s", t.Type)
		}
	}
	f.parameters = append(f.parameters, p...)
	return f
}

func (f *FuncBuilder) ReturnTypes(p ...Parameter) Function {
	for _, t := range p {
		if t.Name != "" {
			for i := range f.returnType {
				if f.returnType[i].Name == "" {
					f.returnType[i].Name = f.returnType[i].Type[:1]
				}

				if f.returnType[i].Name == t.Name {
					f.returnType[i].Name = f.returnType[i].Name + "1"
				}
			}
		}
	}
	f.returnType = append(f.returnType, p...)
	return f
}

func (f *FuncBuilder) Write(b []byte) (n int, err error) {
	if f.body == nil {
		f.body = &bytes.Buffer{}
	}
	return f.body.Write(b)
}

var commentSanitizer = strings.NewReplacer("\n", "\n// ")

// String returns a string representation of the function
func (f *FuncBuilder) String() string {
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
