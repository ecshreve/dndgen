package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/internal/utils"
)

type BonusOptionWrapper struct {
	OptionType   string      `json:"option_type"`
	AbilityScore IndxWrapper `json:"ability_score"`
	Bonus        int         `json:"bonus"`
}

type OptionWrapper struct {
	Item   *IndxWrapper   `json:"item"`
	Choice *ChoiceWrapper `json:"choice"`
}

type BonusWrapper struct {
	AbilityScore IndxWrapper `json:"ability_score"`
	Bonus        int         `json:"bonus"`
}

type ChoiceWrapper struct {
	Desc     []string `json:"desc"`
	Choose   int      `json:"choose"`
	FromType string   `json:"type"`
	From     struct {
		OptionSetType string          `json:"option_set_type"`
		Options       []OptionWrapper `json:"options"`
	} `json:"from"`
}

type AbilityScoreChoiceWrapper struct {
	Desc   []string `json:"desc"`
	Choose int      `json:"choose"`
	Type   string   `json:"type"`
	From   struct {
		OptionSetType string               `json:"option_set_type"`
		Options       []BonusOptionWrapper `json:"options"`
	} `json:"from"`
}

type BaseRaceJSON struct {
	Indx                       string                    `json:"index"`
	Name                       string                    `json:"name"`
	Speed                      int                       `json:"speed"`
	Size                       string                    `json:"size"`
	AlignmentDesc              string                    `json:"alignment"`
	AgeDesc                    string                    `json:"age"`
	SizeDesc                   string                    `json:"size_description"`
	LanguageDesc               string                    `json:"language_desc"`
	StartingProficiencies      []IndxWrapper             `json:"proficiencies"`
	StartingProficiencyOptions ChoiceWrapper             `json:"starting_proficiency_options"`
	AbilityBonuses             []BonusWrapper            `json:"ability_bonuses"`
	AbilityBonusOptions        AbilityScoreChoiceWrapper `json:"ability_bonus_options"`
	Traits                     []IndxWrapper             `json:"traits"`
	Languages                  []IndxWrapper             `json:"languages"`
	LanguageOptions            ChoiceWrapper             `json:"language_options"`
}

type RaceJSON struct {
	*BaseRaceJSON
	// AbilityBonuses []AbilityBonus `json:"ability_bonuses"`
	// StartingProficiencies []Proficiency `json:"starting_proficiencies"`
	// StartingProficiencyOptions []ProficiencyOption `json:"starting_proficiency_options"`
}

type SubraceJSON struct {
	Indx            string         `json:"index"`
	Name            string         `json:"name"`
	Desc            []string       `json:"desc"`
	AbilityBonuses  []BonusWrapper `json:"ability_bonuses"`
	Proficiencies   []IndxWrapper  `json:"proficiencies"`
	Traits          []IndxWrapper  `json:"traits"`
	LanguageOptions ChoiceWrapper  `json:"language_options"`
	Race            IndxWrapper    `json:"race"`
}
type RacePopulator struct {
	client   *ent.Client
	dataFile string
	data     []RaceJSON
	indxToId map[string]int
}

func NewRacePopulator(pp *Popper, dataFile string) *RacePopulator {
	ep := &RacePopulator{
		client:   pp.Client,
		dataFile: dataFile,
		indxToId: pp.IndxToId,
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (cp *RacePopulator) Populate(ctx context.Context) error {
	log.Info("Populating races")
	if len(cp.data) == 0 {
		return fmt.Errorf("no race data to populate")
	}

	for _, rr := range cp.data {
		raceEntity, err := cp.client.Race.Create().
			SetIndx(rr.Indx).
			SetName(rr.Name).
			SetSpeed(rr.Speed).
			SetSize(race.Size(rr.Size)).
			SetSizeDesc(rr.SizeDesc).
			SetAlignmentDesc(rr.AlignmentDesc).
			SetAgeDesc(rr.AgeDesc).
			SetLanguageDesc(rr.LanguageDesc).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating race: %w", err)
		}
		cp.indxToId[rr.Indx] = raceEntity.ID

		log.Info("Created race", "race", raceEntity.Indx)
	}

	if err := cp.populateStartingProficiencies(ctx); err != nil {
		return fmt.Errorf("error populating starting proficiencies: %w", err)
	}

	if err := cp.populateAbilityBonuses(ctx); err != nil {
		return fmt.Errorf("error populating ability bonuses: %w", err)
	}

	if err := cp.populateTraits(ctx); err != nil {
		return fmt.Errorf("error populating traits: %w", err)
	}

	if err := cp.populateLanguages(ctx); err != nil {
		return fmt.Errorf("error populating languages: %w", err)
	}

	return nil
}

func (cp *RacePopulator) PopulateSubraces(ctx context.Context) error {
	log.Info("Populating subraces")
	var subraceData []SubraceJSON
	if err := utils.LoadJSONFile(cp.dataFile, &subraceData); err != nil {
		return fmt.Errorf("error loading subrace data: %w", err)
	}

	for _, rr := range subraceData {
		subraceCreate := cp.client.Subrace.Create().
			SetIndx(rr.Indx).
			SetName(rr.Name).
			SetDesc(rr.Desc).
			SetRaceID(cp.indxToId[rr.Race.Indx])

		for _, ab := range rr.AbilityBonuses {
			subraceCreate = subraceCreate.AddAbilityBonuses(
				cp.client.AbilityBonus.Create().
					SetBonus(ab.Bonus).
					SetAbilityScoreID(cp.indxToId[ab.AbilityScore.Indx]).
					SaveX(ctx),
			)
		}

		traitIDs := []int{}
		for _, t := range rr.Traits {
			traitIDs = append(traitIDs, cp.indxToId[t.Indx])
		}
		subraceCreate = subraceCreate.AddTraitIDs(traitIDs...)
		log.Info("Added traits", "subrace", rr.Indx, "traits", rr.Traits)

		if rr.LanguageOptions.Choose > 0 {
			lIDs := []int{}
			for _, l := range rr.LanguageOptions.From.Options {
				lIDs = append(lIDs, cp.indxToId[fmt.Sprintf("lang-%s", l.Item.Indx)])
			}
			subraceCreate = subraceCreate.AddLanguageOptions(
				cp.client.LanguageChoice.Create().
					SetChoose(rr.LanguageOptions.Choose).
					AddLanguageIDs(lIDs...).
					SaveX(ctx),
			)
			log.Info("Added language options", "race", rr.Indx, "choose", rr.LanguageOptions.Choose, "from", len(rr.LanguageOptions.From.Options))
		}

		if rr.Proficiencies != nil {
			proficiencyIDs := []int{}
			for _, p := range rr.Proficiencies {
				proficiencyIDs = append(proficiencyIDs, cp.indxToId[p.Indx])
			}
			subraceCreate = subraceCreate.AddProficiencyIDs(proficiencyIDs...)
			log.Info("Added proficiencies", "subrace", rr.Indx, "proficiencies", rr.Proficiencies)
		}

		subraceSaved, err := subraceCreate.Save(ctx)
		if err != nil {
			return fmt.Errorf("error adding subrace to race: %w", err)
		}
		log.Info("Saved subrace", "subrace", subraceSaved.Indx)

	}

	return nil
}

// populateAbilityBonuses populates the ability bonuses for a race
func (cp *RacePopulator) populateAbilityBonuses(ctx context.Context) error {
	for _, rr := range cp.data {
		raceUpdate := cp.client.Race.UpdateOneID(cp.indxToId[rr.Indx])

		for _, ab := range rr.AbilityBonuses {
			raceUpdate = raceUpdate.AddAbilityBonuses(
				cp.client.AbilityBonus.Create().
					SetBonus(ab.Bonus).
					SetAbilityScoreID(cp.indxToId[ab.AbilityScore.Indx]).
					AddRaceIDs(cp.indxToId[rr.Indx]).
					SaveX(ctx),
			)
			log.Info("Added ability bonus", "race", rr.Indx, "ability_score", ab.AbilityScore.Indx, "bonus", ab.Bonus)
		}

		if rr.AbilityBonusOptions.Choose > 0 {
			bonusOptions := []*ent.AbilityBonus{}
			for _, bonusOption := range rr.AbilityBonusOptions.From.Options {
				bonusOptions = append(bonusOptions, cp.client.AbilityBonus.Create().
					SetBonus(bonusOption.Bonus).
					SetAbilityScoreID(cp.indxToId[bonusOption.AbilityScore.Indx]).
					AddRaceIDs(cp.indxToId[rr.Indx]).
					SaveX(ctx),
				)
			}

			_, err := raceUpdate.SetAbilityBonusOptions(
				cp.client.AbilityBonusChoice.Create().
					SetChoose(rr.AbilityBonusOptions.Choose).
					AddAbilityBonuses(bonusOptions...).
					SaveX(ctx),
			).Save(ctx)
			if err != nil {
				return fmt.Errorf("error adding ability bonus options to race: %w", err)
			}
			log.Info("Added ability bonus options", "race", rr.Indx, "choose", rr.AbilityBonusOptions.Choose, "from", rr.AbilityBonusOptions.From.Options)
		}

		raceSaved, err := raceUpdate.Save(ctx)
		if err != nil {
			return fmt.Errorf("error adding ability bonuses to race: %w", err)
		}
		log.Info("Saved race", "race", raceSaved.Indx)
	}

	return nil
}

func (cp *RacePopulator) populateStartingProficiencies(ctx context.Context) error {
	for _, rr := range cp.data {
		raceUpdate := cp.client.Race.UpdateOneID(cp.indxToId[rr.Indx])

		spIDs := []int{}
		for _, sp := range rr.StartingProficiencies {
			spIDs = append(spIDs, cp.indxToId[sp.Indx])
		}
		raceUpdate = raceUpdate.AddStartingProficiencyIDs(spIDs...)
		log.Info("Added starting proficiencies", "race", rr.Indx, "proficiencies", spIDs)

		if rr.StartingProficiencyOptions.Choose > 0 {
			poIDs := []int{}
			for _, po := range rr.StartingProficiencyOptions.From.Options {
				poIDs = append(poIDs, cp.indxToId[po.Item.Indx])
			}
			raceUpdate = raceUpdate.SetStartingProficiencyOptions(
				cp.client.ProficiencyChoice.Create().
					SetDesc(rr.StartingProficiencyOptions.Desc).
					SetChoose(rr.StartingProficiencyOptions.Choose).
					AddProficiencyIDs(poIDs...).
					SaveX(ctx),
			)
			log.Info("Added starting proficiency options", "race", rr.Indx, "choose", rr.StartingProficiencyOptions.Choose, "from", len(rr.StartingProficiencyOptions.From.Options))
		}

		raceSaved, err := raceUpdate.Save(ctx)
		if err != nil {
			return fmt.Errorf("error adding starting proficiencies to race: %w", err)
		}
		log.Info("Saved race", "race", raceSaved.Indx)
	}

	return nil
}

func (cp *RacePopulator) populateTraits(ctx context.Context) error {
	for _, rr := range cp.data {
		raceUpdate := cp.client.Race.UpdateOneID(cp.indxToId[rr.Indx])

		traitIDs := []int{}
		for _, t := range rr.Traits {
			traitIDs = append(traitIDs, cp.indxToId[t.Indx])
		}
		raceUpdate = raceUpdate.AddTraitIDs(traitIDs...)
		log.Info("Added traits", "race", rr.Indx, "traits", rr.Traits)

		raceSaved, err := raceUpdate.Save(ctx)
		if err != nil {
			return fmt.Errorf("error adding traits to race: %w", err)
		}
		log.Info("Saved race", "race", raceSaved.Indx)
	}

	return nil
}

func (cp *RacePopulator) populateLanguages(ctx context.Context) error {
	for _, rr := range cp.data {
		raceUpdate := cp.client.Race.UpdateOneID(cp.indxToId[rr.Indx])

		languageIDs := []int{}
		for _, l := range rr.Languages {
			languageIDs = append(languageIDs, cp.indxToId[fmt.Sprintf("lang-%s", l.Indx)])
		}
		raceUpdate = raceUpdate.AddLanguageIDs(languageIDs...)
		log.Info("Added languages", "race", rr.Indx, "languages", rr.Languages)

		if rr.LanguageOptions.Choose > 0 {
			lIDs := []int{}
			for _, l := range rr.LanguageOptions.From.Options {
				lIDs = append(lIDs, cp.indxToId[fmt.Sprintf("lang-%s", l.Item.Indx)])
			}
			raceUpdate = raceUpdate.SetLanguageOptions(
				cp.client.LanguageChoice.Create().
					SetChoose(rr.LanguageOptions.Choose).
					AddLanguageIDs(lIDs...).
					SaveX(ctx),
			)
			log.Info("Added language options", "race", rr.Indx, "choose", rr.LanguageOptions.Choose, "from", len(rr.LanguageOptions.From.Options))
		}

		raceSaved, err := raceUpdate.Save(ctx)
		if err != nil {
			return fmt.Errorf("error adding languages to race: %w", err)
		}
		log.Info("Saved race", "race", raceSaved.Indx)
	}

	return nil
}
