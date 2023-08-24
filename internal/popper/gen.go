//go:build ignore
// +build ignore

package main

import (
	"os"
	"text/template"

	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

var TypesToPopulate = []string{
	"AbilityScore",
	"Class",
	"Race",
	"Skill",
	"Language",
}

var POP_TEMPLATE = `package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)
{{ range . }}
// Populate{{ . }} populates the {{ . }} entities from the JSON data files.
func (p *Popper) Populate{{ . }}(ctx context.Context) ([]*ent.{{ . }}Create, error) {
	fpath := "data/{{ . }}.json"
	var v []{{ . }}Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.{{ . }}Create, len(v))
	for i, vv := range v {
		creates[i] = vv.ToCreate(ctx, p)
	}

	return creates, nil
}
{{ end }}
// CleanUp clears all entities from the database.
func (p *Popper) CleanUp(ctx context.Context) error {
	p.Client.Equipment.Delete().ExecX(ctx)
	p.Client.Weapon.Delete().ExecX(ctx)
	p.Client.Armor.Delete().ExecX(ctx)
	p.Client.Gear.Delete().ExecX(ctx)
	p.Client.Vehicle.Delete().ExecX(ctx)
	p.Client.Tool.Delete().ExecX(ctx)

	log.Infof("deleted all entities for type Equipment and subtypes")

	p.Client.Proficiency.Delete().ExecX(ctx)
	log.Infof("deleted all entities for type Proficiency")
	{{ range . }}
	if _, err := p.Client.{{ . }}.Delete().Exec(ctx); err != nil {
		return oops.Wrapf(err, "unable to delete all {{ . }} entities")
	}
	log.Infof("deleted all entities for type {{ . }}")
	{{ end }}
	return nil
}
`

func main() {
	t := template.Must(template.New("pop").Parse(POP_TEMPLATE))
	// Create the output file.
	ofile, err := os.Create("populators_gen.go")
	if err != nil {
		log.Fatalln(oops.Wrapf(err, "unable to create output file for populators"))
	}
	defer ofile.Close()

	if err := t.Execute(ofile, TypesToPopulate); err != nil {
		log.Fatalln(oops.Wrapf(err, "unable to execute template"))
	}
}
