package popper

import (
	"context"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type ProficiencyWrapper struct {
	Indx                string `json:"index"`
	Name                string `json:"name"`
	ProficiencyCategory string `json:"type"`
}

func CleanCategory(s string) string {
	ss := strings.ToLower(strings.Replace(s, "-", "_", -1))
	ss = strings.Replace(ss, " ", "_", -1)
	ss = strings.Replace(ss, "'", "", -1)

	return ss
}

// PopulateProficiency populates the Proficiency entities from the JSON data files.
func (p *Popper) PopulateProficiency(ctx context.Context) error {
	fpath := "data/Proficiency.json"
	type Wrapper struct {
		Indx                string   `json:"index"`
		Name                string   `json:"name"`
		ProficiencyCategory string   `json:"type"`
		Classes             []string `json:"classes"`
		Races               []string `json:"races"`
		Reference           struct {
			Indx string `json:"index"`
			Url  string `json:"url"`
		} `json:"reference"`
	}

	var v []Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, ww := range v {
		vv := ent.Proficiency{
			Indx:                ww.Indx,
			Name:                ww.Name,
			ProficiencyCategory: proficiency.ProficiencyCategory(CleanCategory(ww.ProficiencyCategory)),
		}

		var eq []*ent.Equipment
		// var as *ent.AbilityScore
		var sk *ent.Skill
		var parseErr error

		switch vv.ProficiencyCategory {
		case proficiency.ProficiencyCategoryArmor:
			fallthrough
		case proficiency.ProficiencyCategoryWeapons:
			fallthrough
		case proficiency.ProficiencyCategoryArtisansTools:
			fallthrough
		case proficiency.ProficiencyCategoryVehicles:
			fallthrough
		case proficiency.ProficiencyCategoryGamingSets:
			fallthrough
		case proficiency.ProficiencyCategoryMusicalInstruments:
			eq, parseErr = p.Reader.Equipment.Query().Where(equipment.Indx(ww.Reference.Indx)).All(ctx)
		case proficiency.ProficiencyCategorySavingThrows:
			// as, parseErr = p.Reader.AbilityScore.Query().Where(abilityscore.Indx(ww.Reference.Indx)).Only(ctx)
		case proficiency.ProficiencyCategorySkills:
			sk, parseErr = p.Reader.Skill.Query().Where(skill.Indx(ww.Reference.Indx)).Only(ctx)
		default:
			log.Info("default")
		}
		if parseErr != nil {
			return oops.Wrapf(parseErr, "unable to parse reference %s", ww.Reference.Indx)
		}

		prof := p.Client.Proficiency.Create().SetProficiency(&vv)
		if eq != nil {
			prof = prof.AddEquipment(eq...)
		}
		// if as != nil {
		// 	prof = prof.Adds(as)
		// }
		if sk != nil {
			prof = prof.AddSkill(sk)
		}

		_, err := prof.Save(ctx)
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
