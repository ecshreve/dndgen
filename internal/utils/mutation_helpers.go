package utils

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/characterproficiency"
)

// HandleCreateCharacterAbilityScores creates the character's ability scores
func HandleCreateCharacterAbilityScores(ctx context.Context, client *ent.Client, ch *ent.Character) ([]*ent.CharacterAbilityScore, error) {
	// Check if character already has ability scores
	charAbilityScores := ch.QueryCharacterAbilityScores().AllX(ctx)
	if len(charAbilityScores) > 0 {
		return charAbilityScores, fmt.Errorf("character already has ability scores")
	}

	// Create ability scores
	allAS := client.AbilityScore.Query().AllX(ctx)
	for _, asScore := range allAS {
		charAbilityScore := client.CharacterAbilityScore.Create().
			SetCharacter(ch).
			SetAbilityScoreID(asScore.ID).
			SetScore(8).
			SaveX(ctx)
		log.Info("Character ability score created", "character_ability_score", charAbilityScore.QueryAbilityScore().OnlyX(ctx).Indx)
	}

	// Check if all ability scores are set
	refetchedCharAbilityScores := ch.QueryCharacterAbilityScores().AllX(ctx)
	if len(refetchedCharAbilityScores) != 6 {
		return nil, fmt.Errorf("character has wrong number of ability scores, got %d", len(refetchedCharAbilityScores))
	}

	return refetchedCharAbilityScores, nil
}

// HandleCharacterSkills creates the character's skills
func HandleCharacterSkills(ctx context.Context, client *ent.Client, ch *ent.Character, abilityScores []*ent.CharacterAbilityScore) ([]*ent.CharacterSkill, error) {
	casCache := make(map[string]*ent.CharacterAbilityScore)
	for _, cas := range abilityScores {
		casCache[cas.QueryAbilityScore().OnlyX(ctx).Indx] = cas
	}
	// Check if character already has skills
	charSkills := ch.QueryCharacterSkills().AllX(ctx)
	if len(charSkills) > 0 {
		return charSkills, fmt.Errorf("character already has skills")
	}

	skills := client.Skill.Query().WithAbilityScore().AllX(ctx)
	for _, sk := range skills {
		charSkill, err := client.CharacterSkill.Create().
			SetCharacter(ch).
			SetSkill(sk).
			SetCharacterAbilityScore(casCache[sk.Edges.AbilityScore.Indx]).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("error creating character skill: %w", err)
		}
		log.Info("Character skill created", "character_skill", charSkill.QuerySkill().OnlyX(ctx).Indx)
	}

	refetchedCharSkills := ch.QueryCharacterSkills().WithSkill().AllX(ctx)
	if len(refetchedCharSkills) != 18 {
		return nil, fmt.Errorf("character has wrong number of skills, got %d", len(refetchedCharSkills))
	}

	return refetchedCharSkills, nil
}

func HandleCharacterProficiencies(ctx context.Context, client *ent.Client, ch *ent.Character, skills []*ent.CharacterSkill) error {
	log.Info("Handling character proficiencies")

	skillCache := make(map[string]*ent.CharacterSkill)
	for _, sk := range skills {
		skillCache[sk.Edges.Skill.Indx] = sk
	}

	raceProficiencies := ch.QueryRace().QueryStartingProficiencies().AllX(ctx)
	for _, pp := range raceProficiencies {
		log.Info("Race proficiency", "race_proficiency", pp.Indx)
		cp := client.CharacterProficiency.Create().
			SetCharacter(ch).
			SetProficiency(pp).
			SetProficiencyType(GetProficiencyTypeFromReference(pp.Reference)).
			SetProficiencySource("RACE_PROFICIENCY").
			SaveX(ctx)

		if cp.ProficiencyType == characterproficiency.ProficiencyTypeSKILL {
			skillRef := strings.Split(pp.Reference, "/")
			if len(skillRef) != 4 {
				return fmt.Errorf("invalid proficiency reference: %s", pp.Reference)
			}
			skillIndx := skillRef[3]
			skillCache[skillIndx].Update().
				SetCharacterProficiencyID(cp.ID).
				SaveX(ctx)
		}
		log.Info("Character proficiency created", "character_proficiency", cp.ID)
	}

	classProficiencies := ch.QueryClass().QueryProficiencies().AllX(ctx)
	for _, pp := range classProficiencies {
		log.Info("Class proficiency", "class_proficiency", pp.Indx)
		cp := client.CharacterProficiency.Create().
			SetCharacter(ch).
			SetProficiency(pp).
			SetProficiencyType(GetProficiencyTypeFromReference(pp.Reference)).
			SetProficiencySource("CLASS_PROFICIENCY").
			SaveX(ctx)
		log.Info("Character proficiency created", "character_proficiency", cp.ID)
	}

	return nil
}
