package genial

import (
	"bytes"
	"strings"
)

type Field struct {
	Name    string
	Type    string
	Comment string
	Tag     []StructTag
}

type StructTag struct {
	Type  string // json:
	Value string // "value"
}

type Struct interface {
	Comment(string) Struct
	Name(string) Struct
	Fields(...Field) Struct

	String() string
}

type StructB struct {
	name    string
	comment string
	fields  []Field
}

func (s *StructB) Comment(comment string) Struct {
	s.comment = comment
	return s
}

func (s *StructB) Name(name string) Struct {
	s.name = name
	return s
}

func (s *StructB) Fields(fields ...Field) Struct {
	s.fields = append(s.fields, fields...)
	return s
}

var tabbedCommentSanitizer = strings.NewReplacer("\n", "\n\t// ")

func (s *StructB) String() string {
	buf := &bytes.Buffer{}

	if s.comment != "" {
		buf.WriteString("// " + commentSanitizer.Replace(s.comment) + "\n")
	}

	buf.WriteString("type ")
	buf.WriteString(s.name)
	buf.WriteString(" struct {\n")

	for _, field := range s.fields {
		if field.Comment != "" {
			buf.WriteString("\t// " + tabbedCommentSanitizer.Replace(field.Comment) + "\n")
		}

		buf.WriteString("\t")
		buf.WriteString(field.Name)
		buf.WriteString(" ")
		buf.WriteString(field.Type)

		// struct tag
		if len(field.Tag) > 0 {
			buf.WriteString(" `")
			for i, tag := range field.Tag {
				if i > 0 {
					buf.WriteString(" ")
				}
				buf.WriteString(tag.Type)
				buf.WriteString(`:"`)
				buf.WriteString(tag.Value)
				buf.WriteString(`"`)
			}
			buf.WriteString("`")
		}

		buf.WriteString("\n")
	}

	buf.WriteString("}\n")

	return buf.String()
}
