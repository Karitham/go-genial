package genial

import (
	"bytes"
	"fmt"
	"go/format"
	"io"
)

type Byter interface {
	Bytes() []byte
}

// PackageB is a package builder
//
// Do not edit once `Bytes`, `String`, `WriteTo`, `Write` or `Read` are called.
type PackageB struct {
	comment string
	license string
	name    string
	imports []string

	decl []Byter

	// nil until Bytes has been called.
	b *bytes.Buffer
}

// License sets the license header for the generated code
func (p *PackageB) License(s string) *PackageB {
	p.license = s
	return p
}

// Declarations sets the declarations for the package
func (p *PackageB) Declarations(b ...Byter) *PackageB {
	p.decl = append(p.decl, b...)
	return p
}

// Comment sets the comment for the package
func (p *PackageB) Comment(c string) *PackageB {
	p.comment = c
	return p
}

// Name sets the name of the package
func (p *PackageB) Name(n string) *PackageB {
	p.name = n
	return p
}

// Namef sets the name of the package using fmt.Sprintf
func (p *PackageB) Namef(format string, args ...interface{}) *PackageB {
	return p.Name(fmt.Sprintf(format, args...))
}

// Imports appends to imports
func (p *PackageB) Imports(i ...string) *PackageB {
	p.imports = append(p.imports, i...)
	return p
}

// String returns the string representation of the package.
// TODO(@Karitham): Add tests for this.
func (p *PackageB) String() string {
	if p.b == nil {
		p.fillBuf()
	}
	return p.b.String()
}

// WriteTo writes the package to the given writer.
func (p *PackageB) WriteTo(w io.Writer) (int64, error) {
	return p.b.WriteTo(w)
}

// Bytes returns the bytes representation of the package.
func (p *PackageB) Bytes() []byte {
	if p.b == nil {
		p.fillBuf()
	}
	return p.b.Bytes()
}

// fillBuf fills the package buffer
func (p *PackageB) fillBuf() {
	b := &bytes.Buffer{}
	b.Write(p.licenseB())
	b.Write(p.commentB())
	b.Write(p.packageB())
	b.Write(p.importsB())

	for _, block := range p.decl {
		b.WriteString("\n")
		b.Write(block.Bytes())
	}

	bf, _ := format.Source(b.Bytes())
	p.b = bytes.NewBuffer(bf)
}

// licenseB returns the license header for the generated code
func (p *PackageB) licenseB() []byte {
	if p.license != "" {
		b := bytes.Buffer{}
		b.WriteString("// ")
		b.WriteString(commentSanitizer.Replace(p.license))
		b.WriteString("\n\n")
		return b.Bytes()
	}
	return nil
}

// commentB returns the comment for the package
func (p *PackageB) commentB() []byte {
	if p.comment != "" {
		b := bytes.Buffer{}
		b.WriteString("// ")
		b.WriteString(commentSanitizer.Replace(p.comment))
		b.WriteString("\n")
		return b.Bytes()
	}
	return nil
}

// packageB returns the package declaration
func (p *PackageB) packageB() []byte {
	return []byte("package " + p.name + "\n")
}

// importsB returns the imports for the package
func (p *PackageB) importsB() []byte {
	if len(p.imports) == 0 {
		return nil
	}

	b := bytes.Buffer{}
	switch len(p.imports) {
	case 1:
		b.WriteString(`import "`)
		b.WriteString(p.imports[0])
		b.WriteRune('"')
		b.WriteRune('\n')
	default:
		b.WriteString("import (\n")
		for _, i := range p.imports {
			b.WriteString("\t\"")
			b.WriteString(i)
			b.WriteString("\"\n")
		}
		b.WriteString(")\n")
	}

	return b.Bytes()
}

// Read the package source
func (p *PackageB) Read(b []byte) (int, error) {
	if p.b == nil {
		p.fillBuf()
	}

	return p.b.Read(b)
}
