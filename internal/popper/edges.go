package popper

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/skill"
)

func (p *Popper) PopulateAbilityScoreEdges(ctx context.Context, raw []ent.AbilityScore) error {
	log.Info("populated ability score edges", "count", 0, "edge", "NONE")
	return nil
}

// PopulateSkillEdges populates the Skill edges.
func (p *Popper) PopulateSkillEdges(ctx context.Context, raw []ent.Skill) error {
	var count = 0
	for _, s := range raw {
		p.Client.Skill.Update().
			Where(skill.Indx(s.Indx)).
			SetAbilityScoreID(p.IndxToId[s.Edges.AbilityScore.Indx]).
			ExecX(ctx)

		count++
	}
	log.Info("populated skill edges", "count", count, "edge", "ability_score")
	return nil
}

// PopulateLanguageEdges populates the Language edges.
func (p *Popper) PopulateLanguageEdges(ctx context.Context, raw []ent.Language) error {
	log.Info("populated language edges", "count", 0, "edge", "NONE")
	return nil
}

// PopulateAlignmentEdges populates the Alignment edges.
func (p *Popper) PopulateAlignmentEdges(ctx context.Context, raw []ent.Alignment) error {
	log.Info("populated alignment edges", "count", 0, "edge", "NONE")
	return nil
}

// PopulateRaceEdges populates the Race edges.
func (p *Popper) PopulateRaceEdges(ctx context.Context, raw []ent.Race) error {
	log.Info("populated race edges", "count", 0, "edge", "TODO")
	return nil
}

// PopulateMagicSchoolEdges populates the MagicSchool edges.
func (p *Popper) PopulateMagicSchoolEdges(ctx context.Context, raw []ent.MagicSchool) error {
	log.Info("populated magic school edges", "count", 0, "edge", "NONE")
	return nil
}

// PopulateDamageTypeEdges populates the DamageType edges.
func (p *Popper) PopulateDamageTypeEdges(ctx context.Context, raw []ent.DamageType) error {
	log.Info("populated damage type edges", "count", 0, "edge", "NONE")
	return nil
}
