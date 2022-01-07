package genial

import "testing"

func TestIfaceBuilder_String(t *testing.T) {
	type fields struct {
		comment     string
		name        string
		signaturers []Signaturer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "test iface builder",
			fields: fields{
				comment: "IfaceTest tests iface",
				name:    "IfaceTest",
				signaturers: []Signaturer{
					(&FuncB{}).
						Comment("GetPetByID basic function comment").
						Name("GetPetByID").
						Parameter("id", "int").
						ReturnTypes("*Pet", "error"),
					(&FuncB{}).
						Comment("GetPetsByTag function comment").
						Name("GetPetsByTag").
						Parameter("tag", "[]string").
						ReturnTypes("[]Pet", "error"),
				},
			},
			want: `// IfaceTest tests iface
type IfaceTest interface {
	// GetPetByID basic function comment
	GetPetByID(id int) (*Pet, error)
	// GetPetsByTag function comment
	GetPetsByTag(tag []string) ([]Pet, error)
}
`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &InterfaceB{
				comment:     tt.fields.comment,
				name:        tt.fields.name,
				signaturers: tt.fields.signaturers,
			}
			if got := i.String(); got != tt.want {
				t.Errorf("IfaceBuilder.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
