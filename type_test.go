package genial

import "testing"

func TestStructBuilder_String(t *testing.T) {
	type fields struct {
		name    string
		fields  []Field
		comment string
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				name:    "TestStruct",
				comment: "This is a basic comment for testing",
				fields: []Field{
					{
						Name:    "TestField",
						Type:    "string",
						Comment: "This is a basic struct for testing",
						Tag: []StructTag{
							{
								Type:      "json",
								Value:     "test_field",
								Omitempty: false,
							},
						},
					},
				},
			},
			want: `// This is a basic comment for testing
type TestStruct struct {
	// This is a basic struct for testing
	TestField string ` + "`json:\"test_field\"`" + `
}
`,
		},
		{
			name: "with omitempty",
			fields: fields{
				name:    "TestStruct",
				comment: "This is a basic comment for testing",
				fields: []Field{
					{
						Name:    "TestField",
						Type:    "string",
						Comment: "This is a basic struct for testing",
						Tag: []StructTag{
							{
								Type:      "json",
								Value:     "test_field",
								Omitempty: true,
							},
						},
					},
				},
			},
			want: `// This is a basic comment for testing
type TestStruct struct {
	// This is a basic struct for testing
	TestField string ` + "`json:\"test_field,omitempty\"`" + `
}
`,
		},
		{
			name: "with omitempty multiple struct tags",
			fields: fields{
				name:    "TestStruct",
				comment: "This is a basic comment for testing",
				fields: []Field{
					{
						Name:    "TestField",
						Type:    "string",
						Comment: "This is a basic struct for testing",
						Tag: []StructTag{
							{
								Type:      "json",
								Value:     "test_field",
								Omitempty: true,
							},
							{
								Type:      "db",
								Value:     "test_field_db",
								Omitempty: false,
							},
						},
					},
				},
			},
			want: `// This is a basic comment for testing
type TestStruct struct {
	// This is a basic struct for testing
	TestField string ` + "`json:\"test_field,omitempty\" db:\"test_field_db\"`" + `
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &StructBuilder{
				name:    tt.fields.name,
				fields:  tt.fields.fields,
				comment: tt.fields.comment,
			}
			if got := s.String(); got != tt.want {
				t.Errorf("StructBuilder.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
