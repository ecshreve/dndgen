package popper

import (
	"context"
	"encoding/json"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func (p *Popper) PopulateAbilityScoreEdges(ctx context.Context, raw []ent.AbilityScore) error {
	return nil
}

func (p *Popper) PopulateLanguageEdges(ctx context.Context, raw []ent.Language) error {
	return nil
}

func (p *Popper) PopulateSkillEdges(ctx context.Context, raw []ent.Skill) {
	for _, r := range raw {
		p.Client.Skill.Query().
			Where(skill.Indx(r.Indx)).OnlyX(ctx).
			Update().
			SetAbilityScoreID(p.IndxToId[r.Edges.AbilityScore.Indx]).SaveX(ctx)
	}
}

func (p *Popper) PopulateDamageTypeEdges(ctx context.Context, raw []ent.DamageType) error {
	return nil
}

func (p *Popper) PopulateClassEdges(ctx context.Context, raw []ent.Class) error {
	filePath := "data/Class.json"

	var v []struct {
		Indx              string `json:"index"`
		StartingEquipment []struct {
			Quantity  int `json:"quantity"`
			Equipment struct {
				Indx string `json:"index"`
			} `json:"equipment"`
		} `json:"starting_equipment,omitempty"`
	}

	if err := LoadJSONFile(filePath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	for _, r := range v {
		cl := p.Client.Class.Query().Where(class.Indx(r.Indx)).OnlyX(ctx)

		seqs := []*ent.StartingEquipmentCreate{}
		for _, seq := range r.StartingEquipment {
			seqEnt := &ent.StartingEquipment{
				ClassID:     cl.ID,
				EquipmentID: p.IndxToId[seq.Equipment.Indx],
				Quantity:    seq.Quantity,
			}
			seqs = append(seqs, p.Client.StartingEquipment.Create().SetStartingEquipment(seqEnt))
		}
		p.Client.StartingEquipment.CreateBulk(seqs...).SaveX(ctx)
		log.Infof("added %d starting equipment to class %s", len(seqs), cl.Name)
	}
	return nil
}

func (p *Popper) PopulateRaceEdges(ctx context.Context, raw []ent.Race) error {
	for _, r := range raw {
		langs, _ := json.Marshal(r.Edges.Languages)
		// abs, _ := json.Marshal(r.Edges.AbilityBonuses)

		race := p.Client.Race.Query().
			Where(race.Indx(r.Indx)).OnlyX(ctx).
			Update().
			AddLanguageIDs(p.GetIDsFromIndxs(langs)...).SaveX(ctx)

		for _, ab := range r.Edges.AbilityBonuses {
			p.Client.AbilityBonus.Create().SetAbilityScoreID(p.IndxToId[ab.Edges.AbilityScore.Indx]).SetBonus(ab.Bonus).SetRaceID(race.ID).SaveX(ctx)
		}
	}
	return nil
}

// PopulateWeaponPropertyEdges populates the WeaponProperty edges from the JSON data files.
func (p *Popper) PopulateWeaponPropertyEdges(ctx context.Context, raw []ent.WeaponProperty) error {
	return nil
}

// PopulateEquipmentEdges populates the Equipment edges from the JSON data files.
func (p *Popper) PopulateEquipmentEdges(ctx context.Context, raw []EquipmentWrapper) error {

	return nil
}

// PopulateMagicSchoolEdges populates the MagicSchool edges from the JSON data files.
func (p *Popper) PopulateMagicSchoolEdges(ctx context.Context, raw []ent.MagicSchool) error {
	return nil
}

// PopulateRuleSectionEdges populates the RuleSection edges from the JSON data files.
func (p *Popper) PopulateRuleSectionEdges(ctx context.Context, raw []ent.RuleSection) error {
	return nil
}

// PopulateRuleEdges populates the Rule edges from the JSON data files.
func (p *Popper) PopulateRuleEdges(ctx context.Context, raw []ent.Rule) error {
	for _, r := range raw {
		rs, _ := json.Marshal(r.Edges.RuleSections)
		p.Client.Rule.Query().
			Where(rule.Indx(r.Indx)).OnlyX(ctx).
			Update().
			AddRuleSectionIDs(p.GetIDsFromIndxs(rs)...).SaveX(ctx)
	}

	return nil
}

// PopulateSubraceEdges populates the Subrace edges from the JSON data files.
func (p *Popper) PopulateSubraceEdges(ctx context.Context, raw []ent.Subrace) error {
	for _, r := range raw {
		sr := p.Client.Subrace.Query().
			Where(subrace.Indx(r.Indx)).OnlyX(ctx).
			Update().
			SetRaceID(p.Client.Race.Query().
				Where(race.Indx(r.Edges.Race.Indx)).OnlyX(ctx).ID).
			SaveX(ctx)

		for _, ab := range r.Edges.AbilityBonuses {
			p.Client.AbilityBonus.Create().SetAbilityScoreID(p.IndxToId[ab.Edges.AbilityScore.Indx]).SetBonus(ab.Bonus).SetSubraceID(sr.ID).SaveX(ctx)
		}
	}

	return nil
}

// PopulateTraitEdges populates the Trait edges from the JSON data files.
func (p *Popper) PopulateTraitEdges(ctx context.Context, raw []ent.Trait) error {
	for _, r := range raw {
		races, _ := json.Marshal(r.Edges.Races)
		subraces, _ := json.Marshal(r.Edges.Subraces)

		p.Client.Trait.Query().
			Where(trait.Indx(r.Indx)).OnlyX(ctx).
			Update().
			AddRaceIDs(p.GetIDsFromIndxs(races)...).
			AddSubraceIDs(p.GetIDsFromIndxs(subraces)...).
			SaveX(ctx)
	}

	return nil
}

// PopulateProficiencyEdges populates the Proficiency edges from the JSON data files.
func (p *Popper) PopulateProficiencyEdges(ctx context.Context, raw []*ent.Proficiency, wrap []ProficiencyWrapper) error {
	for i, vv := range wrap {
		if len(vv.Classes) > 0 {
			classIDs := make([]int, len(vv.Classes))
			for ind, c := range vv.Classes {
				classIDs[ind] = p.IndxToId[c.Indx]
			}
			raw[i].Update().AddClassIDs(classIDs...).SaveX(ctx)
		}
		if len(vv.Races) > 0 {
			raceIDs := []int{}
			for _, c := range vv.Races {
				if c.Indx == "high-elf" {
					continue
				}
				if raceID, ok := p.IndxToId[c.Indx]; ok {
					raceIDs = append(raceIDs, raceID)
				}
			}
			raw[i].Update().AddRaceIDs(raceIDs...).SaveX(ctx)
		}

		ref := vv.Reference.Indx
		switch raw[i].ProficiencyCategory {
		case "skills":
			raw[i].Update().
				SetSkillID(p.IndxToId[ref]).
				SaveX(ctx)
		case "ability_scores":
			raw[i].Update().
				SetSavingThrowID(p.IndxToId[ref]).
				SaveX(ctx)
		case "equipment_categories":
			continue
		case "equipment":
			raw[i].Update().
				SetEquipmentID(p.IndxToId[ref]).
				SaveX(ctx)
		default:
			return oops.Errorf("unknown ProficiencyCategory %s", raw[i].ProficiencyCategory)
		}
	}

	return nil

}

// PopulateAll populates all entities generated from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	_, err := p.PopulateAbilityScore(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}

	_, err = p.PopulateSkill(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}

	_, err = p.PopulateLanguage(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Language entities")
	}

	_, err = p.PopulateDamageType(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate DamageType entities")
	}

	_, err = p.PopulateWeaponProperty(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate WeaponProperty entities")
	}

	err = p.PopulateEquipment(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Equipment entities")
	}

	_, err = p.PopulateClass(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}

	_, err = p.PopulateRace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}

	_, err = p.PopulateSubrace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Subrace entities")
	}

	_, err = p.PopulateProficiency(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Proficiency entities")
	}

	_, err = p.PopulateMagicSchool(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate MagicSchool entities")
	}

	_, err = p.PopulateRuleSection(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate RuleSection entities")
	}

	_, err = p.PopulateRule(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Rule entities")
	}

	_, err = p.PopulateTrait(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Trait entities")
	}

	err = p.PopulateStartingProficiencyOptions(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate StartingProficiencyOptions entities")
	}

	err = p.PopulateProficiencyChoices(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate ProficiencyProficiencyChoices entities")
	}

	return nil
}

// PopulateStartingProficiencyOptions populates the StartingProficiencyOptions edges from the JSON data files.
func (p *Popper) PopulateStartingProficiencyOptions(ctx context.Context) error {
	filePath := "data/Race.json"

	var v []struct {
		Indx                       string `json:"index"`
		StartingProficiencyOptions struct {
			Choose int `json:"choose"`
			From   struct {
				Options []struct {
					Item struct {
						Indx string `json:"index"`
					} `json:"item"`
				} `json:"options"`
			} `json:"from"`
		} `json:"starting_proficiency_options,omitempty"`
	}

	if err := LoadJSONFile(filePath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	startCount := p.Client.ProficiencyChoice.Query().CountX(ctx)
	for _, r := range v {
		if r.StartingProficiencyOptions.Choose == 0 {
			continue
		}

		spoCreate := p.Client.ProficiencyChoice.Create().SetChoose(r.StartingProficiencyOptions.Choose).SaveX(ctx)

		profIDs := []int{}
		for _, prof := range r.StartingProficiencyOptions.From.Options {
			profIDs = append(profIDs, p.IndxToId[prof.Item.Indx])
		}

		spoUpdate := spoCreate.Update().AddOptionIDs(profIDs...).SaveX(ctx)
		p.Client.Race.Query().Where(race.Indx(r.Indx)).OnlyX(ctx).Update().SetStartingProficiencyOptionID(spoUpdate.ID).SaveX(ctx)
		// spoCreated.Update().AddProficiencyIDs(profIDs...).SaveX(ctx)
	}
	endCount := p.Client.ProficiencyChoice.Query().CountX(ctx)
	log.Infof("created %d ProficiencyChoices -- starting_proficiency_options", endCount-startCount)

	return nil

}

// PopulateProficiencyProficiencyChoices populates the PopulateProficiencyProficiencyChoices edges from the JSON data files.
func (p *Popper) PopulateProficiencyChoices(ctx context.Context) error {
	filePath := "data/Class.json"

	type refArray struct {
		Options []struct {
			Item struct {
				Indx string `json:"index"`
			} `json:"item"`
		} `json:"options"`
	}

	type refChoice struct {
		From struct {
			Options []struct {
				Item struct {
					Indx string `json:"index"`
				} `json:"item"`
			} `json:"options"`
		} `json:"from"`
	}

	type nestedChoice struct {
		From []struct {
			Options []struct {
				Choice struct {
					Desc   string `json:"desc"`
					Choose int    `json:"choose"`
					From   struct {
						Options []struct {
							Item struct {
								Indx string `json:"index"`
							} `json:"item"`
						} `json:"options"`
					} `json:"from"`
				} `json:"choice"`
			} `json:"options"`
		} `json:"from"`
	}

	type vChoice struct {
		Desc   string `json:"desc"`
		Choose int    `json:"choose"`
		From   struct {
			Options []struct {
				Item struct {
					Indx string `json:"index"`
				} `json:"item"`
			} `json:"options"`
		} `json:"from"`
	}

	var vClasses []struct {
		Indx    string    `json:"index"`
		Choices []vChoice `json:"proficiency_choices,omitempty"`
	}

	if err := LoadJSONFile(filePath, &vClasses); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	for _, c := range vClasses {
		if len(c.Choices) == 0 {
			continue
		}

		choiceIDs := []int{}
		for _, pc := range c.Choices {
			if c.Indx == "monk" {
				continue
			}
			created := p.Client.ProficiencyChoice.Create().SetChoose(pc.Choose).SetDesc(pc.Desc).SaveX(ctx)
			profIDs := []int{}
			for _, prof := range pc.From.Options {
				profIDs = append(profIDs, p.IndxToId[prof.Item.Indx])
			}

			created.Update().AddOptionIDs(profIDs...).SaveX(ctx)
			choiceIDs = append(choiceIDs, created.ID)
		}
		cl := p.Client.Class.Query().
			Where(class.Indx(c.Indx)).OnlyX(ctx).
			Update().AddProficiencyChoiceIDs(choiceIDs...).SaveX(ctx)
		log.Infof("created %d ProficiencyChoices for class %s", len(choiceIDs), cl.Name)
	}

	return nil
}
