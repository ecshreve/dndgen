package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/prerequisite"
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

type prerequisiteJSON struct {
	Type    string `json:"type"`
	Level   int    `json:"level"`
	Feature string `json:"feature"`
	Spell   string `json:"spell"`
}

type featureJSON struct {
	Indx          string             `json:"index"`
	Class         IndxWrapper        `json:"class"`
	Name          string             `json:"name"`
	Level         int                `json:"level"`
	Prerequisites []prerequisiteJSON `json:"prerequisites"`
}

// PopulateFeatureEdges populates the Feature edges.
func (p *Popper) PopulateFeatureEdges(ctx context.Context) error {
	fpath := fmt.Sprintf("%s/Feature.json", p.DataDir)
	var v []featureJSON
	if err := utils.LoadJSONFile(fpath, &v); err != nil {
		return fmt.Errorf("LoadJSONFile: %w", err)
	}

	for _, f := range v {
		// create the prerequisites
		prereqs := make([]*ent.PrerequisiteCreate, 0)
		for _, prereq := range f.Prerequisites {
			switch prereq.Type {
			case "level":
				prereqs = append(prereqs, p.Client.Prerequisite.Create().
					SetPrerequisiteType(prerequisite.PrerequisiteTypeLevel).
					SetFeatureID(p.IndxToId[f.Indx]).
					SetLevelValue(prereq.Level))
			case "feature":
				prereqs = append(prereqs, p.Client.Prerequisite.Create().
					SetPrerequisiteType(prerequisite.PrerequisiteTypeFeature).
					SetFeatureID(p.IndxToId[f.Indx]).
					SetFeatureValue(prereq.Feature))
			case "spell":
				prereqs = append(prereqs, p.Client.Prerequisite.Create().
					SetPrerequisiteType(prerequisite.PrerequisiteTypeSpell).
					SetFeatureID(p.IndxToId[f.Indx]).
					SetSpellValue(prereq.Spell))
			default:
				return fmt.Errorf("unknown prerequisite type: %s", prereq.Type)
			}
		}

		prereqsCreated, err := p.Client.Prerequisite.CreateBulk(prereqs...).Save(ctx)
		if err != nil {
			return fmt.Errorf("CreateBulk: %w", err)
		}
		log.Debug("populated feature -> prerequisites", "feature", f.Indx, "prerequisites", len(prereqsCreated))
	}
	log.Info("populated feature edges")
	return nil
}
