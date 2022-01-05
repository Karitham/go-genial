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

// Struct represents a struct
type Struct interface {
	Comment(string) Struct
	Commentf(string, ...interface{}) Struct
	Name(string) Struct
	Namef(string, ...interface{}) Struct
	Fields(...Field) Struct
	Field(string, string, ...StructTag) Struct

	String() string
}

// StructB is a struct builder
type StructB struct {
	name    string
	comment string
	fields  []Field
}

// Comment sets the comment
func (s *StructB) Comment(comment string) Struct {
	s.comment = comment
	return s
}

// Commentf sets the comment using fmt.Sprintf
func (s *StructB) Commentf(format string, args ...interface{}) Struct {
	return s.Comment(fmt.Sprintf(format, args...))
}

// Name sets the name
func (s *StructB) Name(name string) Struct {
	s.name = name
	return s
}

// Namef sets the name using fmt.Sprintf
func (s *StructB) Namef(format string, args ...interface{}) Struct {
	return s.Name(fmt.Sprintf(format, args...))
}

// Fields sets the fields
func (s *StructB) Fields(fields ...Field) Struct {
	s.fields = append(s.fields, fields...)
	return s
}

// Field appends a basic field
func (s *StructB) Field(name string, typeName string, tags ...StructTag) Struct {
	return s.Fields(Field{
		Name: name,
		Type: typeName,
		Tag:  tags,
	})
}

var tabbedCommentSanitizer = strings.NewReplacer("\n", "\n\t// ")

// String returns the string representation of the struct
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
