package popper

import (
	"github.com/ecshreve/dndgen/ent"
)

type IndxWrapper struct {
	Indx string `json:"index"`
}

type SkillWrapper struct {
	Indx         string      `json:"index"`
	Name         string      `json:"name"`
	Desc         []string    `json:"desc"`
	AbilityScore IndxWrapper `json:"ability_score"`
}

// ToEnt converts the SkillWrapper to an ent.Skill.
func (w *SkillWrapper) ToEnt(p *Popper) *ent.Skill {
	return &ent.Skill{
		Indx:           w.Indx,
		Name:           w.Name,
		Desc:           w.Desc,
		AbilityScoreID: p.IndxToId[w.AbilityScore.Indx],
	}
}

func (w *SkillWrapper) ToCreate(p *Popper) *ent.SkillCreate {
	return p.Client.Skill.Create().SetSkill(w.ToEnt(p))
}

type LanguageWrapper struct {
	ent.Language
}

//  {
// 	Indx           string `json:"index"`
// 	Name           string `json:"name"`
// 	Desc           string `json:"desc"`
// 	LanguageType   string `json:"type"`
// 	LanguageScript string `json:"script"`
// }

// ToEnt converts the LanguageWrapper to an ent.Language.
func (w *LanguageWrapper) ToEnt(p *Popper) *ent.Language {
	return &w.Language
}

func (w *LanguageWrapper) ToCreate(p *Popper) *ent.LanguageCreate {
	return p.Client.Language.Create().SetLanguage(w.ToEnt(p))
}

type RaceWrapper struct {
	ent.Race
	Languages []IndxWrapper `json:"languages"`
}

// ToEnt converts the RWrapper to an ent.R.
func (w *RaceWrapper) ToEnt(p *Popper) *ent.Race {
	return &w.Race
}

// ToCreate converts the RWrapper to an ent.RCreate.
func (w *RaceWrapper) ToCreate(p *Popper) *ent.RaceCreate {
	return p.Client.Race.Create().SetRace(w.ToEnt(p)).
		AddLanguageIDs(p.GetIDsFromIndxWrappers(w.Languages)...)
}

type ClassWrapper struct {
	Indx   string `json:"index"`
	Name   string `json:"name"`
	HitDie int    `json:"hit_die"`
}

// ToEnt converts the ClassWrapper to an ent.Class.
func (w *ClassWrapper) ToEnt(p *Popper) *ent.Class {
	return &ent.Class{
		Indx:   w.Indx,
		Name:   w.Name,
		HitDie: w.HitDie,
	}
}

// ToCreate converts the ClassWrapper to an ent.ClassCreate.
func (w *ClassWrapper) ToCreate(p *Popper) *ent.ClassCreate {
	return p.Client.Class.Create().SetClass(w.ToEnt(p))
}
