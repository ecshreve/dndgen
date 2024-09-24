package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/rulesection"
	"github.com/ecshreve/dndgen/internal/utils"
)

// PopulateRuleEdges populates the Rule edges.
func (p *Popper) PopulateRuleEdges(ctx context.Context) error {
	fpath := fmt.Sprintf("%s/Rule.json", p.DataDir)
	var v []ent.Rule
	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("LoadJSONFile: %w", err)
	}

	// iterate over the rules
	for _, r := range v {
		sectionIndxs := make([]string, 0)
		for _, rs := range r.Edges.Sections {
			sectionIndxs = append(sectionIndxs, rs.Indx)
		}

		// update the sections
		p.Client.RuleSection.Update().
			Where(rulesection.IndxIn(sectionIndxs...)).
			SetRuleID(p.IndxToId[r.Indx]).
			ExecX(ctx)

		log.Debug("populated rule -> sections", "rule", r.Indx, "sections", len(sectionIndxs))
	}
	log.Info("populated rule edges")
	return nil
}

// PopulateSkillEdges populates the Skill edges.
func (p *Popper) PopulateSkillEdges(ctx context.Context) error {
	fpath := fmt.Sprintf("%s/Skill.json", p.DataDir)
	var v []ent.Skill
	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("LoadJSONFile: %w", err)
	}

	for _, s := range v {
		p.Client.Skill.UpdateOneID(p.IndxToId[s.Indx]).
			SetAbilityScoreID(p.IndxToId[s.Edges.AbilityScore.Indx]).
			ExecX(ctx)
		log.Debug("populated skill -> ability_score", "skill", s.Indx, "ability_score", s.Edges.AbilityScore.Indx)
	}
	log.Info("populated skill edges")
	return nil
}
