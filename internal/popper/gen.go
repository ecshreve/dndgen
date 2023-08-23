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
func (p *Popper) Populate{{ . }}(ctx context.Context) error {
	fpath := "internal/popper/data/{{ . }}.json"
	var v []ent.{{ . }}

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, vv := range v {
		created, err := p.Client.{{ . }}.Create().Set{{ . }}(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
		p.IdToIndx[created.ID] = vv.Indx
		p.IndxToId[vv.Indx] = created.ID
	}

	return nil
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

// PopulateAllGen populates all entities generated from the JSON data files.
func (p *Popper) PopulateAllGen(ctx context.Context) error {
	var start int

	{{ range . }}
	start = len(p.IdToIndx)
	if err := p.Populate{{ . }}(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate {{ . }} entities")
	}
	log.Infof("created %d entities for type {{ . }}", len(p.IdToIndx)-start)
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
