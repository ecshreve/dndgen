//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
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
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(entviz.Extension{}),
		entc.Extensions(&StructConverter{}),
		entc.Extensions(ex),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
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
{{ end }}
`)),
	}
}

var _ entc.Extension = (*EncodeExtension)(nil)
