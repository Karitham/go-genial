package genial

import (
	"bytes"
	"fmt"
	"strings"
)

// Field is a struct field
type Field struct {
	Name    string
	Type    string
	Comment string
	Tag     []StructTag
}

// StructTag is a struct tag
type StructTag struct {
	Type  string // json:
	Value string // "value"
}

// StructB is a struct builder
type StructB struct {
	name    string
	comment string
	fields  []Field
}

// Comment sets the comment
func (s *StructB) Comment(comment string) *StructB {
	s.comment = comment
	return s
}

// Commentf sets the comment using fmt.Sprintf
func (s *StructB) Commentf(format string, args ...interface{}) *StructB {
	return s.Comment(fmt.Sprintf(format, args...))
}

// Name sets the name
func (s *StructB) Name(name string) *StructB {
	s.name = name
	return s
}

// Namef sets the name using fmt.Sprintf
func (s *StructB) Namef(format string, args ...interface{}) *StructB {
	return s.Name(fmt.Sprintf(format, args...))
}

// Fields sets the fields
func (s *StructB) Fields(fields ...Field) *StructB {
	s.fields = append(s.fields, fields...)
	return s
}

// Field appends a basic field
func (s *StructB) Field(name string, typeName string, tags ...StructTag) *StructB {
	return s.Fields(Field{
		Name: name,
		Type: typeName,
		Tag:  tags,
	})
}

var tabbedCommentSanitizer = strings.NewReplacer("\n", "\n\t// ")

// String returns the string representation of the struct
func (s *StructB) String() string {
	return string(s.Bytes())
}

// Bytes returns the byte representation of the struct
func (s *StructB) Bytes() []byte {
	buf := &bytes.Buffer{}

	if s.comment != "" {
		buf.WriteString("// ")
		buf.WriteString(commentSanitizer.Replace(s.comment))
		buf.WriteString("\n")
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

	return buf.Bytes()
}
