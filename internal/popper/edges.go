package popper

import (
	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/skill"
)

func (p *Popper) PopulateAbilityScoreEdges(raw []ent.AbilityScore) error {
	return nil
}

// PopulateSkillEdges populates the Skill edges.
func (p *Popper) PopulateSkillEdges(raw []ent.Skill) error {
	log.Info("populating skill edges")
	ctx := *p.Context

	for _, s := range raw {
		p.Client.Skill.Update().
			Where(skill.Indx(s.Indx)).
			SetAbilityScoreID(p.IndxToId[s.Edges.AbilityScore.Indx]).
			ExecX(ctx)
	}
	return nil
}

// PopulateLanguageEdges populates the Language edges.
func (p *Popper) PopulateLanguageEdges(raw []ent.Language) error {
	return nil
}

// PopulateAlignmentEdges populates the Alignment edges.
func (p *Popper) PopulateAlignmentEdges(raw []ent.Alignment) error {
	return nil
}

// PopulateRaceEdges populates the Race edges.
func (p *Popper) PopulateRaceEdges(raw []ent.Race) error {
	return nil
}
