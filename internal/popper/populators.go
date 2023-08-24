package popper

import (
	"context"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type WrapperWr interface {
	ToEnt(*Popper) interface{}
	ToCreate(*Popper) interface{}
}

type IndxWrapper struct {
	Indx string `json:"index"`
}

type AbilityScoreWrapper struct {
	ent.AbilityScore
}

// ToEnt converts the AbilityScoreWrapper to an ent.AbilityScore.
func (w *AbilityScoreWrapper) ToEnt(ctx context.Context, p *Popper) *ent.AbilityScore {
	return &w.AbilityScore
}

// ToCreate converts the AbilityScoreWrapper to an ent.AbilityScoreCreate.
func (w *AbilityScoreWrapper) ToCreate(ctx context.Context, p *Popper) *ent.AbilityScoreCreate {
	return p.Client.AbilityScore.Create().SetAbilityScore(w.ToEnt(ctx, p))
}

type SkillWrapper struct {
	Indx         string      `json:"index"`
	Name         string      `json:"name"`
	Desc         []string    `json:"desc"`
	AbilityScore IndxWrapper `json:"ability_score"`
}

// ToEnt converts the SkillWrapper to an ent.Skill.
func (w *SkillWrapper) ToEnt(ctx context.Context, p *Popper) *ent.Skill {
	return &ent.Skill{
		Indx:           w.Indx,
		Name:           w.Name,
		Desc:           w.Desc,
		AbilityScoreID: p.Reader.AbilityScore.Query().Where(abilityscore.Indx(w.AbilityScore.Indx)).OnlyIDX(ctx),
	}
}

func (w *SkillWrapper) ToCreate(ctx context.Context, p *Popper) *ent.SkillCreate {
	return p.Client.Skill.Create().SetSkill(w.ToEnt(ctx, p))
}

type LanguageWrapper struct {
	ent.Language
}

// ToEnt converts the LanguageWrapper to an ent.Language.
func (w *LanguageWrapper) ToEnt(ctx context.Context, p *Popper) *ent.Language {
	return &w.Language
}

func (w *LanguageWrapper) ToCreate(ctx context.Context, p *Popper) *ent.LanguageCreate {
	return p.Client.Language.Create().SetLanguage(w.ToEnt(ctx, p))
}

type RaceWrapper struct {
	ent.Race
	Languages []IndxWrapper `json:"languages"`
}

// ToEnt converts the RWrapper to an ent.R.
func (w *RaceWrapper) ToEnt(ctx context.Context, p *Popper) *ent.Race {
	return &w.Race
}

// ToCreate converts the RWrapper to an ent.RCreate.
func (w *RaceWrapper) ToCreate(ctx context.Context, p *Popper) *ent.RaceCreate {
	return p.Client.Race.Create().SetRace(w.ToEnt(ctx, p)).
		AddLanguageIDs(p.Reader.Language.Query().Where(language.IndxIn(GetIDStrings(w.Languages)...)).IDsX(ctx)...)
}

type ClassWrapper struct {
	Indx   string `json:"index"`
	Name   string `json:"name"`
	HitDie int    `json:"hit_die"`
}

// ToEnt converts the ClassWrapper to an ent.Class.
func (w *ClassWrapper) ToEnt(ctx context.Context, p *Popper) *ent.Class {
	return &ent.Class{
		Indx:   w.Indx,
		Name:   w.Name,
		HitDie: w.HitDie,
	}
}

// ToCreate converts the ClassWrapper to an ent.ClassCreate.
func (w *ClassWrapper) ToCreate(ctx context.Context, p *Popper) *ent.ClassCreate {
	return p.Client.Class.Create().SetClass(w.ToEnt(ctx, p))
}

// PopulateAll populates all entities generated from the JSON data files.
func (p *Popper) PopulateAll(ctx context.Context) error {
	createsAbilityScore, err := p.PopulateAbilityScore(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate AbilityScore entities")
	}
	p.Client.AbilityScore.CreateBulk(createsAbilityScore...).ExecX(ctx)
	log.Infof("created %d entities for type AbilityScore", len(createsAbilityScore))

	createsSkill, err := p.PopulateSkill(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Skill entities")
	}
	p.Client.Skill.CreateBulk(createsSkill...).ExecX(ctx)
	log.Infof("created %d entities for type Skill", len(createsSkill))

	createsLanguage, err := p.PopulateLanguage(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Language entities")
	}
	p.Client.Language.CreateBulk(createsLanguage...).ExecX(ctx)
	log.Infof("created %d entities for type Language", len(createsLanguage))

	createsClass, err := p.PopulateClass(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Class entities")
	}
	p.Client.Class.CreateBulk(createsClass...).ExecX(ctx)
	log.Infof("created %d entities for type Class", len(createsClass))

	createsRace, err := p.PopulateRace(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Race entities")
	}
	p.Client.Race.CreateBulk(createsRace...).ExecX(ctx)
	log.Infof("created %d entities for type Race", len(createsRace))

	start := p.Client.Equipment.Query().CountX(ctx)
	err = p.PopulateEquipment(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Equipment entities")
	}
	log.Infof("created %d entities for type Equipment", p.Client.Equipment.Query().CountX(ctx)-start)

	start = p.Client.Proficiency.Query().CountX(ctx)
	err = p.PopulateProficiency(ctx)
	if err != nil {
		return oops.Wrapf(err, "unable to populate Proficiency entities")
	}
	log.Infof("created %d entities for type Proficiency", p.Client.Proficiency.Query().CountX(ctx)-start)

	return nil
}
