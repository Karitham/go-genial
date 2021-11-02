package genial

import "strings"

type PackageBuilder struct {
	comment string
	name    string
	imports []string

	blocks []Stringer
}

type Stringer interface {
	String() string
}

type Package interface {
	Comment(string) Package
	Name(string) Package
	Imports(...string) Package
	Blocks(...Stringer) Package

	String() string
}

func (p *PackageBuilder) Blocks(b ...Stringer) Package {
	p.blocks = append(p.blocks, b...)
	return p
}

func (p *PackageBuilder) Comment(c string) Package {
	p.comment = c
	return p
}

func (p *PackageBuilder) Name(n string) Package {
	p.name = n

	return p
}

func (p *PackageBuilder) Imports(i ...string) Package {
	p.imports = append(p.imports, i...)
	return p
}

// String returns the string representation of the package.
// TODO(@Karitham): Add tests for this.
// TODO(@Karitham): Add support for License header.
func (p *PackageBuilder) String() string {
	b := &strings.Builder{}

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

	for _, block := range p.blocks {
		b.WriteString("\n")
		b.WriteString(block.String())
	}

	return b.String()
}
