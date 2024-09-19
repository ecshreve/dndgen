package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/rulesection"
	"github.com/ecshreve/dndgen/ent/skill"
)

func (p *Popper) PopulateAbilityScoreEdges(ctx context.Context, raw []ent.AbilityScore) error {
	return nil
}

// PopulateAlignmentEdges populates the Alignment edges.
func (p *Popper) PopulateAlignmentEdges(ctx context.Context, raw []ent.Alignment) error {
	return nil
}

// PopulateCoinEdges populates the Coin edges.
func (p *Popper) PopulateCoinEdges(ctx context.Context, raw []ent.Coin) error {
	return nil
}

// PopulateConditionEdges populates the Condition edges.
func (p *Popper) PopulateConditionEdges(ctx context.Context, raw []ent.Condition) error {
	return nil
}

// PopulateDamageTypeEdges populates the DamageType edges.
func (p *Popper) PopulateDamageTypeEdges(ctx context.Context, raw []ent.DamageType) error {
	return nil
}

// PopulateFeatEdges populates the Feat edges.
func (p *Popper) PopulateFeatEdges(ctx context.Context, raw []ent.Feat) error {
	return nil
}

// PopulateLanguageEdges populates the Language edges.
func (p *Popper) PopulateLanguageEdges(ctx context.Context, raw []ent.Language) error {
	return nil
}

// PopulateMagicSchoolEdges populates the MagicSchool edges.
func (p *Popper) PopulateMagicSchoolEdges(ctx context.Context, raw []ent.MagicSchool) error {
	return nil
}

// PopulateRaceEdges populates the Race edges.
func (p *Popper) PopulateRaceEdges(ctx context.Context, raw []ent.Race) error {
	for _, r := range raw {
		for _, ab := range r.Edges.AbilityBonuses {
			p.Client.AbilityBonus.Create().
				SetAbilityScoreID(p.IndxToId[ab.Edges.AbilityScore.Indx]).
				SetBonus(ab.Bonus).
				SetRaceID(p.IndxToId[r.Indx]).
				ExecX(ctx)
			log.Debug("populated race -> ability_bonus", "race", r.Indx, "ability_score", ab.Edges.AbilityScore.Indx, "bonus", ab.Bonus)
		}

		languageIndxs := make([]string, 0)
		for _, l := range r.Edges.Languages {
			languageIndxs = append(languageIndxs, fmt.Sprintf("lang-%s", l.Indx))
		}

		p.Client.Language.Update().
			Where(language.IndxIn(languageIndxs...)).
			AddRaceIDs(p.IndxToId[r.Indx]).
			ExecX(ctx)
		log.Debug("populated race -> languages", "race", r.Indx, "languages", languageIndxs)

	}
	log.Info("populated race edges")
	return nil
}

// PopulateRuleEdges populates the Rule edges.
func (p *Popper) PopulateRuleEdges(ctx context.Context, raw []ent.Rule) error {
	// iterate over the rules
	for _, r := range raw {
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

// PopulateRuleSectionEdges populates the RuleSection edges.
func (p *Popper) PopulateRuleSectionEdges(ctx context.Context, raw []ent.RuleSection) error {
	return nil
}

// PopulateSkillEdges populates the Skill edges.
func (p *Popper) PopulateSkillEdges(ctx context.Context, raw []ent.Skill) error {
	for _, s := range raw {
		p.Client.Skill.Update().
			Where(skill.Indx(s.Indx)).
			SetAbilityScoreID(p.IndxToId[s.Edges.AbilityScore.Indx]).
			ExecX(ctx)
		log.Debug("populated skill -> ability_score", "skill", s.Indx, "ability_score", s.Edges.AbilityScore.Indx)
	}
	log.Info("populated skill edges")
	return nil
}

// PopulateWeaponPropertyEdges populates the WeaponProperty edges.
func (p *Popper) PopulateWeaponPropertyEdges(ctx context.Context, raw []ent.WeaponProperty) error {
	return nil
}
