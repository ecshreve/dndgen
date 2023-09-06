package popper

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

func (p *Popper) PopulateCoinEdges(ctx context.Context, raw []ent.Coin) error {
	return nil
}

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
		StartingEquipmentOptions []struct {
			Choose int    `json:"choose"`
			Desc   string `json:"desc"`
			From   struct {
				Options []struct {
					Count int `json:"count,omitempty"`
					Of    *struct {
						Indx string `json:"index,omitempty"`
						Url  string `json:"url,omitempty"`
					} `json:"of,omitempty"`
					Choice *struct {
						Choose int    `json:"choose,omitempty"`
						Desc   string `json:"desc,omitempty"`
						From   *struct {
							EquipmentCategory struct {
								Indx string `json:"index,omitempty"`
							} `json:"equipment_category,omitempty"`
						} `json:"from,omitempty"`
					} `json:"choice,omitempty"`
				} `json:"options,omitempty"`
			} `json:"from,omitempty"`
		} `json:"starting_equipment_options,omitempty"`
	}

	if err := LoadJSONFile(filePath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	for _, r := range v {
		cl := p.Client.Class.Query().Where(class.Indx(r.Indx)).OnlyX(ctx)

		sq := []*ent.ClassEquipmentCreate{}
		for _, eq := range r.StartingEquipment {
			cleq := &ent.ClassEquipment{
				ClassID:     cl.ID,
				EquipmentID: p.IndxToId[eq.Equipment.Indx],
				Quantity:    eq.Quantity,
			}
			sq = append(sq, p.Client.ClassEquipment.Create().SetClassEquipment(cleq))
		}
		ceqs := p.Client.ClassEquipment.CreateBulk(sq...).SaveX(ctx)
		log.Infof("added %d starting equipment to class %s", len(ceqs), cl.Name)

		seqs := []*ent.EquipmentChoiceCreate{}
		for _, seq := range r.StartingEquipmentOptions {
			for _, opt := range seq.From.Options {
				if opt.Choice == nil {
					continue
				}
				if opt.Choice.From.EquipmentCategory.Indx == "" {
					continue
				}

				catIndex := opt.Choice.From.EquipmentCategory.Indx
				catIndex = strings.TrimSuffix(catIndex, "-weapons")
				sp := strings.Split(catIndex, "-")

				// get all equipment entities that satisfy all of the categories
				// in named in the sp slice
				allEqInCatQ := p.Client.Equipment.Query()
				for _, spp := range sp {
					allEqInCatQ = allEqInCatQ.Where(equipment.HasEquipmentCategoryWith(equipmentcategory.NameContains(spp)))
				}
				allEqInCat := allEqInCatQ.AllX(ctx)

				seqEnt := &ent.EquipmentChoice{
					Choose: opt.Choice.Choose,
					Desc:   opt.Choice.Desc,
				}
				seqs = append(seqs, p.Client.EquipmentChoice.Create().SetEquipmentChoice(seqEnt).AddClasIDs(cl.ID).AddEquipment(allEqInCat...))
			}
		}
		p.Client.EquipmentChoice.CreateBulk(seqs...).SaveX(ctx)
		log.Infof("added %d starting equipment choices to class %s", len(seqs), cl.Name)
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
				AddSkillIDs(p.IndxToId[ref]).
				SaveX(ctx)
		case "ability_scores":
			raw[i].Update().
				AddSavingThrowIDs(p.IndxToId[ref]).
				SaveX(ctx)
		case "equipment_categories":
			ids := p.Client.EquipmentCategory.Query().
				Where(equipmentcategory.NameContains(strings.Split(ref, "-")[0])).IDsX(ctx)

			raw[i].Update().
				AddEquipmentCategoryIDs(ids...).
				SaveX(ctx)
		case "equipment":
			raw[i].Update().
				AddEquipmentIDs(p.IndxToId[ref]).
				SaveX(ctx)
		default:
			return oops.Errorf("unknown ProficiencyCategory %s", raw[i].ProficiencyCategory)
		}
	}

	return nil

}

// PopulateAll populates all entities generated from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	_, err := p.PopulateCoin(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Coin entities")
	}
	_, err = p.PopulateAbilityScore(ctx)
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

	_, err = p.PopulateProficiencyChoices(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate ProficiencyProficiencyChoices entities")
	}

	return nil
}

type RaceChoiceWrapper struct {
	Indx                       string         `json:"index"`
	StartingProficiencyOptions *ChoiceWrapper `json:"starting_proficiency_options,omitempty"`
}

// PopulateStartingProficiencyOptions populates the StartingProficiencyOptions edges from the JSON data files.
func (p *Popper) PopulateStartingProficiencyOptions(ctx context.Context) error {
	filePath := "data/Race.json"

	var v []*RaceChoiceWrapper

	if err := LoadJSONFile(filePath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	for _, r := range v {
		if r.StartingProficiencyOptions == nil {
			continue
		}

		ch, err := p.BuildChoice(ctx, r.StartingProficiencyOptions)
		if err != nil {
			return oops.Wrapf(err, "unable to build Choice entity")
		}

		rr := p.Client.Race.Query().Where(race.Indx(r.Indx)).OnlyX(ctx).Update().AddProficiencyChoiceIDs(ch.ID).SaveX(ctx)
		log.Infof("created ProficiencyChoices for race %s", rr.Name)
	}

	return nil

}

type RefChoice struct {
	Options []struct {
		Item struct {
			Indx string `json:"index"`
		} `json:"item"`
	} `json:"options,omitempty"`
}

type NestedChoice struct {
	Options []struct {
		Choice struct {
			Desc   string `json:"desc,omitempty"`
			Choose int    `json:"choose"`
			From   struct {
				Options []struct {
					Item struct {
						Indx string `json:"index"`
					} `json:"item"`
				} `json:"options"`
			} `json:"from"`
		} `json:"choice,omitempty"`
	} `json:"options,omitempty"`
}

type ChoiceWrapper struct {
	Desc   string                 `json:"desc"`
	Choose int                    `json:"choose"`
	From   map[string]interface{} `json:"from"`
}

type ClassChoiceWrapper struct {
	Indx               string           `json:"index"`
	ProficiencyChoices []*ChoiceWrapper `json:"proficiency_choices,omitempty"`
}

// PopulateProficiencyProficiencyChoices populates the PopulateProficiencyProficiencyChoices edges from the JSON data files.
func (p *Popper) PopulateProficiencyChoices(ctx context.Context) ([]*ent.ProficiencyChoice, error) {
	filePath := "data/Class.json"

	var vClasses []*ClassChoiceWrapper

	if err := LoadJSONFile(filePath, &vClasses); err != nil {
		return nil, oops.Wrapf(err, "unable to load JSON file %s", filePath)
	}

	for _, c := range vClasses {
		if len(c.ProficiencyChoices) == 0 {
			continue
		}
		// cl := p.Client.Class.Query().
		// 	Where(class.Indx(c.Indx)).OnlyX(ctx)

		choiceIDs := []int{}
		for _, pc := range c.ProficiencyChoices {
			ch, err := p.BuildChoice(ctx, pc)
			if err != nil {
				return nil, oops.Wrapf(err, "unable to build Choice entity")
			}
			choiceIDs = append(choiceIDs, ch.ID)
		}

		cl := p.Client.Class.Query().
			Where(class.Indx(c.Indx)).OnlyX(ctx).
			Update().AddProficiencyChoiceIDs(choiceIDs...).SaveX(ctx)
		log.Infof("created %d ProficiencyChoices for class %s", len(choiceIDs), cl.Name)
	}

	return nil, nil
}

// buildChoice creates a Choice entity from a JSON struct.
func (p *Popper) BuildChoice(ctx context.Context, c *ChoiceWrapper) (*ent.ProficiencyChoice, error) {
	if c == nil || c.From == nil || c.Choose == 0 {
		return nil, nil
	}

	created := p.Client.ProficiencyChoice.Create().SetChoose(c.Choose).SetDesc(c.Desc).SaveX(ctx)

	fromJson, err := json.Marshal(c.From)
	if err != nil {
		return nil, oops.Wrapf(err, "unable to marshal JSON string")
	}

	var single RefChoice
	if err := json.Unmarshal(fromJson, &single); err != nil {
		return nil, oops.Wrapf(err, "unable to unmarshal JSON string")
	}

	profIDs := []int{}
	for _, prof := range single.Options {
		pid := p.IndxToId[prof.Item.Indx]
		if pid == 0 {
			continue
		}
		profIDs = append(profIDs, pid)
	}
	if len(profIDs) > 0 {
		created.Update().AddProficiencyIDs(profIDs...).SaveX(ctx)
		return created, nil
	}

	var nest NestedChoice
	if err := json.Unmarshal(fromJson, &nest); err != nil {
		return nil, oops.Wrapf(err, "unable to unmarshal JSON string")
	}

	subchoiceIDs := []int{}
	for _, nc := range nest.Options {
		if nc.Choice.Choose == 0 {
			continue
		}

		sf, err := json.Marshal(nc.Choice.From)
		if err != nil {
			return nil, oops.Wrapf(err, "unable to marshal JSON string")
		}
		var subChoiceMap map[string]interface{}
		if err := json.Unmarshal(sf, &subChoiceMap); err != nil {
			return nil, oops.Wrapf(err, "unable to unmarshal JSON string")
		}

		sub, err := p.BuildChoice(ctx, &ChoiceWrapper{
			Desc:   nc.Choice.Desc,
			Choose: nc.Choice.Choose,
			From:   subChoiceMap,
		})
		if err != nil {
			return nil, oops.Wrapf(err, "unable to build Choice entity")
		}
		subchoiceIDs = append(subchoiceIDs, sub.ID)
	}

	if len(subchoiceIDs) > 0 {
		created.Update().AddSubChoiceIDs(subchoiceIDs...).SaveX(ctx)
		return created, nil
	}

	return nil, nil
}
