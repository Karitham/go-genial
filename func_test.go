package genial

import "testing"

func TestFuncBuilder_Build(t *testing.T) {
	type fields struct {
		comment    string
		name       string
		receiver   Parameter
		parameters []Parameter
		returnType []Parameter
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "basic",
			fields: fields{
				name: "basic",
				parameters: []Parameter{
					{
						Name: "foo",
						Type: "int",
					},
				},
				comment: "basic is a basic test function",
				returnType: []Parameter{
					{
						Type: "int",
					},
				},
			},
			want: `// basic is a basic test function
func basic(foo int) int {
}
`,
		},
		{
			name: "with receiver",
			fields: fields{
				name: "GetPetByID",
				receiver: Parameter{
					Name: "c",
					Type: "*Client",
				},
				parameters: []Parameter{
					{
						Name: "petID",
						Type: "uint",
					},
				},
				comment: "GetPetByID gets a pet by ID",
				returnType: []Parameter{
					{
						Type: "*Pet",
					},
				},
			},
			want: `// GetPetByID gets a pet by ID
func (c *Client) GetPetByID(petID uint) *Pet {
}
`,
		},
		{
			name: "with receiver and variadic",
			fields: fields{
				name: "GetPetsByTags",
				receiver: Parameter{
					Name: "c",
					Type: "*Client",
				},
				parameters: []Parameter{
					{
						Name: "tags",
						Type: "...string",
					},
				},
				comment: "GetPetsByTags gets multiple pets by tags",
				returnType: []Parameter{
					{
						Type: "[]Pet",
					},
					{
						Type: "error",
					},
				},
			},
			want: `// GetPetsByTags gets multiple pets by tags
func (c *Client) GetPetsByTags(tags ...string) ([]Pet, error) {
}
`,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			f := &FuncBuilder{
				comment:    tt.fields.comment,
				name:       tt.fields.name,
				receiver:   tt.fields.receiver,
				parameters: tt.fields.parameters,
				returnType: tt.fields.returnType,
			}
			if got := f.String(); got != tt.want {
				t.Errorf("FuncBuilder.Build() = %v, want %v", got, tt.want)
			}
		})
	}
}
