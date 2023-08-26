package popper

import (
	"context"
	"encoding/json"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/equipmentcategory"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/skill"
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

// PopulateEquipmentEdges populates the Equipment edges from the JSON data files.
func (p *Popper) PopulateEquipmentEdges(ctx context.Context, raw []ent.Equipment) error {
	categories := []string{"weapon", "armor", "adventuring-gear", "tools", "mounts-and-vehicles", "other"}
	for _, r := range categories {
		created := p.Client.EquipmentCategory.Create().SetIndx(equipmentcategory.Indx(r)).SaveX(ctx)
		p.IndxToId[r] = created.ID
		p.IdToIndx[created.ID] = r
	}

	for _, r := range raw {
		cost := p.Client.Cost.Create().SetCost(r.Edges.Cost).SaveX(ctx)
		p.Client.Equipment.Query().
			Where(equipment.Indx(r.Indx)).OnlyX(ctx).
			Update().
			SetEquipmentCategoryID(p.IndxToId[string(r.Edges.EquipmentCategory.Indx)]).
			SetCost(cost).SaveX(ctx)
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

	_, err = p.PopulateClass(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}

	_, err = p.PopulateRace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}

	_, err = p.PopulateEquipment(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Equipment entities")
	}

	// createsClass, err := p.PopulateClass(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Class entities")
	// }
	// p.Client.Class.CreateBulk(createsClass...).ExecX(ctx)
	// log.Infof("created %d entities for type Class", len(createsClass))

	// createsRace, err := p.PopulateRace(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Race entities")
	// }
	// p.Client.Race.CreateBulk(createsRace...).ExecX(ctx)
	// log.Infof("created %d entities for type Race", len(createsRace))

	// start := p.Client.Equipment.Query().CountX(ctx)
	// err = p.PopulateEquipment(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Equipment entities")
	// }
	// log.Infof("created %d entities for type Equipment", p.Client.Equipment.Query().CountX(ctx)-start)

	// start = p.Client.Proficiency.Query().CountX(ctx)
	// err = p.PopulateProficiency(ctx)
	// if err != nil {
	// 	return oops.Wrapf(err, "unable to populate Proficiency entities")
	// }
	// log.Infof("created %d entities for type Proficiency", p.Client.Proficiency.Query().CountX(ctx)-start)

	return nil
}
