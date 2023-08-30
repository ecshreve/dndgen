//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/contrib/entproto"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
	"github.com/hedwigz/entviz"
)

func main() {
	// Create an entgql extension.
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./gqlserver/ent.graphql"),
		entgql.WithWhereInputs(true),
		entgql.WithNodeDescriptor(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	epex, err := entproto.NewExtension(
		entproto.WithProtoDir("./ent/proto"),
	)

	opts := []entc.Option{
		entc.Extensions(entviz.Extension{}),
		entc.Extensions(&StructConverter{}),
		entc.Extensions(&EncodeExtension{}),
		entc.Extensions(ex),
		entc.Extensions(epex),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{}, opts...); err != nil {
		log.Fatal("running ent codegen:", err)
	}
}

type StructConverter struct {
	entc.DefaultExtension
}

func (s *StructConverter) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("model/additional/structconverter").
			Parse(`
	{{ $builder := $.CreateName }}
	{{ $receiver := receiver $builder }}

	func ({{ $receiver }} *{{ $builder }}) Set{{ $.Name }}(input *{{ $.Name }}) *{{ $builder }} {
			{{- range $f := $.Fields }}
					{{- $setter := print "Set" $f.StructField }}
					{{ $receiver }}.{{ $setter }}(input.{{ $f.StructField }})
			{{- end }}
			return {{ $receiver }}
	}`)),
	}
}

var _ entc.Extension = (*StructConverter)(nil)

// EncodeExtension is an implementation of entc.Extension that adds a MarshalJSON
// method to each generated type <T> and inlines the Edges field to the top level JSON.
type EncodeExtension struct {
	entc.DefaultExtension
}

// Templates of the extension.
func (e *EncodeExtension) Templates() []*gen.Template {
	return []*gen.Template{
		gen.MustParse(gen.NewTemplate("model/additional/jsonencode").
			Parse(`
{{ if $.Edges }}
	// MarshalJSON implements the json.Marshaler interface.
	func ({{ $.Receiver }} *{{ $.Name }}) MarshalJSON() ([]byte, error) {
			type Alias {{ $.Name }}
			return json.Marshal(&struct {
					*Alias
					{{ $.Name }}Edges
			}{
					Alias: (*Alias)({{ $.Receiver }}),
					{{ $.Name }}Edges: {{ $.Receiver }}.Edges,
			})
	}

	// UnmarshalJSON implements the json.Unmarshaler interface.
	func ({{ $.Receiver }} *{{ $.Name }}) UnmarshalJSON(data []byte) error {
			type Alias {{ $.Name }}
			aux := &struct {
					*Alias
					{{ $.Name }}Edges
			}{
					Alias: (*Alias)({{ $.Receiver }}),
			}

			if err := json.Unmarshal(data, &aux); err != nil {
				  return err
			}

			{{ $.Receiver }}.Edges = aux.{{ $.Name }}Edges
			return nil
	}
{{ end }}
`)),
	}
}

// Hooks of the extension.
func (e *EncodeExtension) Hooks() []gen.Hook {
	return []gen.Hook{}
	// 	func(next gen.Generator) gen.Generator {
	// 		return gen.GenerateFunc(func(g *gen.Graph) error {
	// 			tag := edge.Annotation{StructTag: `json:"-"`}
	// 			for _, n := range g.Nodes {
	// 				n.Annotations.Set(tag.Name(), tag)
	// 			}
	// 			return next.Generate(g)
	// 		})
	// 	},
	// }
}
