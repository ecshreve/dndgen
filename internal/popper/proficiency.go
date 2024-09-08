package popper

import (
	"context"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/samsarahq/go/oops"
)

type ProficiencyWrapper struct {
	Indx      string `json:"index"`
	Name      string `json:"name"`
	Reference struct {
		Indx string `json:"index"`
		Url  string `json:"url"`
	} `json:"reference"`
	Classes []*ent.Class `json:"classes"`
	Races   []*ent.Race  `json:"races"`
}

// ToEnt converts the ProficiencyWrapper to an ent.Proficiency.
func (p *ProficiencyWrapper) ToEnt() *ent.Proficiency {
	return &ent.Proficiency{
		Indx:                p.Indx,
		Name:                p.Name,
		ProficiencyCategory: strings.Replace(strings.Split(p.Reference.Url, "/")[2], "-", "_", -1),
	}
}

// PopulateProficiency populates the Proficiency entities from the JSON data files.
func (p *Popper) PopulateProficiency(ctx context.Context) ([]*ent.Proficiency, error) {
	fpath := "data/Proficiency.json"
	var v []ProficiencyWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	creates := make([]*ent.ProficiencyCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Proficiency.Create().SetProficiency(&ent.Proficiency{
			Indx:                vv.Indx,
			Name:                vv.Name,
			ProficiencyCategory: strings.Replace(strings.Split(vv.Reference.Url, "/")[2], "-", "_", -1),
		})
	}

	created, err := p.Client.Proficiency.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to save Proficiency entities")
	}
	log.Infof("created %d entities for type Proficiency", len(created))

	p.PopulateProficiencyEdges(ctx, created, v)

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}
