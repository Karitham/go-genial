package genial

import "strings"

type PackageB struct {
	comment string
	license string
	name    string
	imports []string

	decl []Stringer
}

type Stringer interface {
	String() string
}

type Package interface {
	Comment(string) Package
	License(string) Package
	Name(string) Package
	Imports(...string) Package
	Declarations(...Stringer) Package

	String() string
}

func (p *PackageB) License(s string) Package {
	p.license = s
	return p
}

func (p *PackageB) Declarations(b ...Stringer) Package {
	p.decl = append(p.decl, b...)
	return p
}

func (p *PackageB) Comment(c string) Package {
	p.comment = c
	return p
}

func (p *PackageB) Name(n string) Package {
	p.name = n
	return p
}

func (p *PackageB) Imports(i ...string) Package {
	p.imports = append(p.imports, i...)
	return p
}

// String returns the string representation of the package.
// TODO(@Karitham): Add tests for this.
func (p *PackageB) String() string {
	b := &strings.Builder{}

	if p.license != "" {
		b.WriteString("//")
		b.WriteString(commentSanitizer.Replace(p.license))
		b.WriteString("\n\n")
	}

	if p.comment != "" {
		b.WriteString("// ")
		b.WriteString(commentSanitizer.Replace(p.comment))
		b.WriteString("\n")
	}

	b.WriteString("package ")
	b.WriteString(p.name)
	b.WriteString("\n")

	switch len(p.imports) {
	case 0:
		// do nothing
	case 1:
		b.WriteString(`import "`)
		b.WriteString(p.imports[0])
		b.WriteString(`"\n`)
	default:
		b.WriteString("import (\n")
		for _, i := range p.imports {
			b.WriteString(`\t"`)
			b.WriteString(i)
			b.WriteString(`"\n`)
		}
		b.WriteString(")\n")
	}

	for _, block := range p.decl {
		b.WriteString("\n")
		b.WriteString(block.String())
	}

	return b.String()
}
