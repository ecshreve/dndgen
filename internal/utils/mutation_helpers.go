package utils

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
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
