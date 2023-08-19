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
	"Alignment",
	"Skill",
	"Proficiency",
	"MagicSchool",
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
		_, err := p.Client.{{ . }}.Create().Set{{ . }}(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
	}

	return nil
}
{{ end }}

// PopulateAll populates all entities from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	var start int
	{{ range . }}
	start = p.Client.{{ . }}.Query().CountX(ctx)
	if err := p.Populate{{ . }}(ctx); err != nil {
		return oops.Wrapf(err, "unable to populate {{ . }} entities")
	}
	log.Infof("created %d entities for type {{ . }}", p.Client.{{ . }}.Query().CountX(ctx)-start)
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
