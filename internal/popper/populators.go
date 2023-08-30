package popper

import (
	"context"
	"encoding/json"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/rule"
	"github.com/ecshreve/dndgen/ent/skill"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/samsarahq/go/oops"
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

	return nil
}

func (p *Popper) PopulateRaceEdges(ctx context.Context, raw []ent.Race) error {
	for _, r := range raw {
		langs, _ := json.Marshal(r.Edges.Languages)
		p.Client.Race.Query().
			Where(race.Indx(r.Indx)).OnlyX(ctx).
			Update().
			AddLanguageIDs(p.GetIDsFromIndxs(langs)...).SaveX(ctx)
	}
	return nil
}

// PopulateWeaponPropertyEdges populates the WeaponProperty edges from the JSON data files.
func (p *Popper) PopulateWeaponPropertyEdges(ctx context.Context, raw []ent.WeaponProperty) error {
	return nil
}

// PopulateEquipmentEdges populates the Equipment edges from the JSON data files.
func (p *Popper) PopulateEquipmentEdges(ctx context.Context, raw []EquipmentWrapper) error {
	// categories := []string{"weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "other"}
	// for _, r := range categories {
	// 	created := p.Client.EquipmentCategory.Create().SetIndx(equipmentcategory.Indx(r)).SaveX(ctx)
	// 	p.IndxToId[r] = created.ID
	// 	p.IdToIndx[created.ID] = r
	// }

	for _, r := range raw {
		cost := p.Client.Cost.Create().SetQuantity(r.Cost.Quantity).SetUnit(r.Cost.Unit).SaveX(ctx)
		p.Client.Equipment.Query().
			Where(equipment.Indx(r.Indx)).OnlyX(ctx).
			Update().
			SetCost(cost).Save(ctx)

	}
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
		profs, _ := json.Marshal(r.Edges.Proficiencies)

		p.Client.Subrace.Query().
			Where(subrace.Indx(r.Indx)).OnlyX(ctx).
			Update().
			SetRaceID(p.Client.Race.Query().
				Where(race.Indx(r.Edges.Race.Indx)).OnlyX(ctx).ID).
			AddProficiencyIDs(p.GetIDsFromIndxs(profs)...).
			SaveX(ctx)
	}

	return nil
}

// PopulateAll populates all entities generated from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	// _, err := p.PopulateAbilityScore(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate AbilityScore entities")
	// }
	for _, r := range p.Client.AbilityScore.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateSkill(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Skill entities")
	// }
	for _, r := range p.Client.Skill.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateLanguage(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Language entities")
	// }
	for _, r := range p.Client.Language.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateDamageType(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate DamageType entities")
	// }
	for _, r := range p.Client.DamageType.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateClass(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Class entities")
	// }
	for _, r := range p.Client.Class.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateRace(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Race entities")
	// }
	for _, r := range p.Client.Race.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateWeaponProperty(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate WeaponProperty entities")
	// }
	for _, r := range p.Client.WeaponProperty.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// err = p.PopulateEquipment(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Equipment entities")
	// }
	for _, r := range p.Client.Equipment.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	_, err := p.PopulateProficiency(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Proficiency entities")
	}
	for _, r := range p.Client.Proficiency.Query().AllX(ctx) {
		p.IndxToId[r.Indx] = r.ID
		p.IdToIndx[r.ID] = r.Indx
	}

	// _, err = p.PopulateMagicSchool(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate MagicSchool entities")
	// }

	// _, err = p.PopulateRuleSection(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate RuleSection entities")
	// }

	// _, err = p.PopulateRule(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Rule entities")
	// }

	_, err = p.PopulateSubrace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Subrace entities")
	}

	return nil
}
