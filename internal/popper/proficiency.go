package popper

import (
	"context"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
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

var proficiencyJSON = `
[{
	"index": "saving-throw-int",
	"type": "Saving Throws",
	"name": "Saving Throw: INT",
	"classes": [
		{
			"index": "druid",
			"name": "Druid",
			"url": "/api/classes/druid"
		},
		{
			"index": "rogue",
			"name": "Rogue",
			"url": "/api/classes/rogue"
		},
		{
			"index": "wizard",
			"name": "Wizard",
			"url": "/api/classes/wizard"
		}
	],
	"races": [],
	"url": "/api/proficiencies/saving-throw-int",
	"reference": {
		"index": "int",
		"name": "INT",
		"url": "/api/ability-scores/int"
	}
},
{
	"index": "pan-flute",
	"type": "Musical Instruments",
	"name": "Pan flute",
	"classes": [],
	"races": [],
	"url": "/api/proficiencies/pan-flute",
	"reference": {
		"index": "pan-flute",
		"name": "Pan flute",
		"url": "/api/equipment/pan-flute"
	}
},
{
	"index": "skill-survival",
	"type": "Skills",
	"name": "Skill: Survival",
	"classes": [],
	"races": [],
	"url": "/api/proficiencies/skill-survival",
	"reference": {
		"index": "survival",
		"name": "Survival",
		"url": "/api/skills/survival"
	},
	{
    "index": "light-armor",
    "type": "Armor",
    "name": "Light Armor",
    "classes": [
      {
        "index": "barbarian",
        "name": "Barbarian",
        "url": "/api/classes/barbarian"
      },
      {
        "index": "bard",
        "name": "Bard",
        "url": "/api/classes/bard"
      },
      {
        "index": "cleric",
        "name": "Cleric",
        "url": "/api/classes/cleric"
      },
      {
        "index": "druid",
        "name": "Druid",
        "url": "/api/classes/druid"
      },
      {
        "index": "ranger",
        "name": "Ranger",
        "url": "/api/classes/ranger"
      },
      {
        "index": "rogue",
        "name": "Rogue",
        "url": "/api/classes/rogue"
      },
      {
        "index": "warlock",
        "name": "Warlock",
        "url": "/api/classes/warlock"
      }
    ],
    "races": [],
    "url": "/api/proficiencies/light-armor",
    "reference": {
      "index": "light-armor",
      "name": "Light Armor",
      "url": "/api/equipment-categories/light-armor"
    }
  }
}]`

// PopulateProficiency populates the Proficiency entities from the JSON data files.
func (p *Popper) PopulateProficiency(ctx context.Context) ([]*ent.Proficiency, error) {
	fpath := "data/Proficiency.json"
	var v []ProficiencyWrapper

	// if err := json.Unmarshal([]byte(proficiencyJSON), &v); err != nil {
	// 	return nil, oops.Wrapf(err, "unable to load JSON file %s", proficiencyJSON)
	// }

	if err := LoadJSONFile(fpath, &v); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	allClasses := p.Client.Class.Query().AllX(ctx)
	classIndxToId := map[string]int{}
	for _, c := range allClasses {
		classIndxToId[c.Indx] = c.ID
	}

	allRaces := p.Client.Race.Query().AllX(ctx)
	raceIndxToId := map[string]int{}
	for _, c := range allRaces {
		raceIndxToId[c.Indx] = c.ID
	}

	creates := []*ent.Proficiency{}
	for _, vv := range v {
		created := p.Client.Proficiency.Create().SetProficiency(vv.ToEnt()).SaveX(ctx)
		updated := p.Client.Proficiency.Query().Where(proficiency.ID(created.ID)).OnlyX(ctx).Update()
		if len(vv.Classes) > 0 {
			classIDs := make([]int, len(vv.Classes))
			for ind, c := range vv.Classes {
				classIDs[ind] = classIndxToId[c.Indx]
			}
			updated.AddClassIDs(classIDs...).SaveX(ctx)
		}
		if len(vv.Races) > 0 {
			raceIDs := []int{}
			for _, c := range vv.Races {
				if raceID, ok := raceIndxToId[c.Indx]; ok {
					raceIDs = append(raceIDs, raceID)
				}
			}
			updated.AddRaceIDs(raceIDs...).SaveX(ctx)
		}
		creates = append(creates, created)

		p.PopulateProficiencyEdges(ctx, created, vv.Reference.Indx)
		p.IdToIndx[created.ID] = created.Indx
		p.IndxToId[created.Indx] = created.ID
	}
	log.Infof("created %d entities for type Proficiency", len(creates))

	return creates, nil
	// for _, c := range creates {
	// 	p.IdToIndx[c.ID] = c.Indx
	// 	p.IndxToId[c.Indx] = c.ID
	// }

	// return creates, nil
}

// type ProficiencyWrapper struct {
// 	Indx                string `json:"index"`
// 	Name                string `json:"name"`
// 	ProficiencyCategory string `json:"type"`
// }

// func CleanCategory(s string) string {
// 	ss := strings.ToLower(strings.Replace(s, "-", "_", -1))
// 	ss = strings.Replace(ss, " ", "_", -1)
// 	ss = strings.Replace(ss, "'", "", -1)

// 	return ss
// }

// // PopulateProficiency populates the Proficiency entities from the JSON data files.
// func (p *Popper) PopulateProficiency(ctx context.Context) error {
// 	fpath := "data/Proficiency.json"
// 	type Wrapper struct {
// 		Indx                string   `json:"index"`
// 		Name                string   `json:"name"`
// 		ProficiencyCategory string   `json:"type"`
// 		Classes             []string `json:"classes"`
// 		Races               []string `json:"races"`
// 		Reference           struct {
// 			Indx string `json:"index"`
// 			Url  string `json:"url"`
// 		} `json:"reference"`
// 	}

// 	var v []Wrapper

// 	if err := LoadJSONFile(fpath, &v); err != nil {
// 		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
// 	}

// 	for _, ww := range v {
// 		vv := ent.Proficiency{
// 			Indx:                ww.Indx,
// 			Name:                ww.Name,
// 			ProficiencyCategory: proficiency.ProficiencyCategory(CleanCategory(ww.ProficiencyCategory)),
// 		}

// 		var eq []*ent.Equipment
// 		// var as *ent.Proficiency
// 		var sk *ent.Skill
// 		var parseErr error

// 		switch vv.ProficiencyCategory {
// 		case proficiency.ProficiencyCategoryArmor:
// 			fallthrough
// 		case proficiency.ProficiencyCategoryWeapons:
// 			fallthrough
// 		case proficiency.ProficiencyCategoryArtisansTools:
// 			fallthrough
// 		case proficiency.ProficiencyCategoryVehicles:
// 			fallthrough
// 		case proficiency.ProficiencyCategoryGamingSets:
// 			fallthrough
// 		case proficiency.ProficiencyCategoryMusicalInstruments:
// 			eq, parseErr = p.Reader.Equipment.Query().Where(equipment.Indx(ww.Reference.Indx)).All(ctx)
// 		case proficiency.ProficiencyCategorySavingThrows:
// 			// as, parseErr = p.Reader.Proficiency.Query().Where(Proficiency.Indx(ww.Reference.Indx)).Only(ctx)
// 		case proficiency.ProficiencyCategorySkills:
// 			sk, parseErr = p.Reader.Skill.Query().Where(skill.Indx(ww.Reference.Indx)).Only(ctx)
// 		default:
// 			log.Info("default")
// 		}
// 		if parseErr != nil {
// 			return oops.Wrapf(parseErr, "unable to parse reference %s", ww.Reference.Indx)
// 		}

// 		prof := p.Client.Proficiency.Create().SetProficiency(&vv)
// 		if eq != nil {
// 			prof = prof.AddEquipment(eq...)
// 		}
// 		// if as != nil {
// 		// 	prof = prof.Adds(as)
// 		// }
// 		if sk != nil {
// 			prof = prof.AddSkill(sk)
// 		}

// 		_, err := prof.Save(ctx)
// 		if ent.IsConstraintError(err) {
// 			log.Debugf("constraint failed, skipping %s", vv.Indx)
// 			continue
// 		}
// 		if err != nil {
// 			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
// 		}

// 	}

// 	return nil
// }
