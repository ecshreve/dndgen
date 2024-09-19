// GENERATED BY popgen
// DO NOT EDIT
package popper

import (
	"fmt"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
	
	"github.com/charmbracelet/log"
)

// PopulateAbilityScore populates the AbilityScore entities from the JSON data files.
func (p *Popper) PopulateAbilityScore() ([]*ent.AbilityScore, error) {
	ctx := *p.Context
	fpath := "internal/popper/data/AbilityScore.json"
	var v []ent.AbilityScore

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.AbilityScoreCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.AbilityScore.Create().SetAbilityScore(&vv)
	}

	created, err := p.Client.AbilityScore.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "AbilityScore")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	return created, nil
}

// PopulateSkill populates the Skill entities from the JSON data files.
func (p *Popper) PopulateSkill() ([]*ent.Skill, error) {
	ctx := *p.Context
	fpath := "internal/popper/data/Skill.json"
	var v []ent.Skill

	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	creates := make([]*ent.SkillCreate, len(v))
	for i, vv := range v {
		creates[i] = p.Client.Skill.Create().SetSkill(&vv)
	}

	created, err := p.Client.Skill.CreateBulk(creates...).Save(ctx)
	if err != nil {
		return nil, fmt.Errorf("CreateBulk: %w", err)
	}
	log.Info("bulk creation success", "created", len(created), "entity", "Skill")

	for _, c := range created {
		p.IdToIndx[c.ID] = c.Indx
		p.IndxToId[c.Indx] = c.ID
	}

	p.PopulateSkillEdges(ctx, v)

	return created, nil
}
