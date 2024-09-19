package popper

import (
	"context"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/skill"
)

func (p *Popper) PopulateAbilityScoreEdges(ctx context.Context, raw []*ent.AbilityScore) error {
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
