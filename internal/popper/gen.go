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
	"Skill",
	"Language",
	"DamageType",
	"Race",
	"Class",
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
func (p *Popper) Populate{{ . }}(ctx context.Context) ([]*ent.{{ . }}, error) {
	fpath := "data/{{ . }}.json"
	var v []ent.{{ . }}

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.{{ . }}Create, len(v))
	for i, vv := range v {
		creates[i] = p.Client.{{ . }}.Create().Set{{ . }}(&vv)
	}

	created, err := p.Client.{{ . }}.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save {{ . }} entities")
	}
	log.Infof("created %d entities for type {{ . }}", len(created))

	p.Populate{{ . }}Edges(ctx, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}
{{ end }}
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
