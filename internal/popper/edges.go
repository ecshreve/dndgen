package popper

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/skill"
)

func (p *Popper) PopulateAbilityScoreEdges(ctx context.Context, raw []ent.AbilityScore) error {
	return nil
}

// PopulateSkillEdges populates the Skill edges.
func (p *Popper) PopulateSkillEdges(ctx context.Context, raw []ent.Skill) error {
	log.Info("populating skill edges")

	for _, s := range raw {
		p.Client.Skill.Update().
			Where(skill.Indx(s.Indx)).
			SetAbilityScoreID(p.IndxToId[s.Edges.AbilityScore.Indx]).
			ExecX(ctx)
	}
	return nil
}

// PopulateLanguageEdges populates the Language edges.
func (p *Popper) PopulateLanguageEdges(ctx context.Context, raw []ent.Language) error {
	return nil
}

// PopulateAlignmentEdges populates the Alignment edges.
func (p *Popper) PopulateAlignmentEdges(ctx context.Context, raw []ent.Alignment) error {
	return nil
}

// PopulateRaceEdges populates the Race edges.
func (p *Popper) PopulateRaceEdges(ctx context.Context, raw []ent.Race) error {
	return nil
}
