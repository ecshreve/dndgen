//go:build ignore
// +build ignore

package main

import (
	"os"
	"text/template"

	"github.com/charmbracelet/log"
)

var TypesToPopulate = []string{
	"AbilityScore",
	"Alignment",
	"Class",
	"Race",
	// "Skill",
	// "Language",
	// "DamageType",
	// "Race",
	// "Class",
	// "WeaponProperty",
	// "MagicSchool",
	// "RuleSection",
	// "Rule",
	// "Subrace",
	// "Trait",
	// "Coin",
}

var POP_TEMPLATE = `// GENERATED BY seeder
package seeder

import (
	"context"
	"fmt"

	"builder/ent"
	"builder/utils"
	
	"github.com/charmbracelet/log"
)
{{ range . }}
// Seed{{ . }} seeds the {{ . }} entities from the JSON data files.
func Seed{{ . }}(ctx context.Context, client *ent.Client) ([]*ent.{{ . }}, error) {
	fpath := "data/{{ . }}.json"
	var v []ent.{{ . }}

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.{{ . }}Create, len(v))
	for i, vv := range v {
		creates[i] = client.{{ . }}.Create().Set{{ . }}(&vv)
	}

	created, err := client.{{ . }}.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("{{ . }} bulk creation success", "created", len(created))

	return created, nil
}
{{ end }}
`

func main() {
	t := template.Must(template.New("pop").Parse(POP_TEMPLATE))
	// Create the output file.
	ofile, err := os.Create("seeders_gen.go")
	if err != nil {
		log.Fatal(err)
	}
	defer ofile.Close()

	if err := t.Execute(ofile, TypesToPopulate); err != nil {
		log.Fatal(err)
	}
}