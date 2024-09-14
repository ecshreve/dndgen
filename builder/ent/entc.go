//go:build ignore
// +build ignore

package main

import (
	"log"

	"entgo.io/contrib/entgql"
	"entgo.io/ent/entc"
	"entgo.io/ent/entc/gen"
)

func main() {
	ex, err := entgql.NewExtension(
		entgql.WithConfigPath("./gqlgen.yml"),
		entgql.WithSchemaGenerator(),
		entgql.WithSchemaPath("./graph/ent.graphql"),
		entgql.WithWhereInputs(true),
	)
	if err != nil {
		log.Fatalf("creating entgql extension: %v", err)
	}

	opts := []entc.Option{
		entc.Extensions(&StructConverter{}),
		entc.Extensions(ex),
	}

	if err := entc.Generate("./ent/schema", &gen.Config{
		Features: []gen.Feature{gen.FeatureVersionedMigration},
	}, opts...); err != nil {
		log.Fatalf("running ent codegen: %v", err)
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
