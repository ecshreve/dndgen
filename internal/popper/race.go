package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/internal/utils"
)

type OptionWrapper struct {
	OptionType string      `json:"option_type"`
	Item       IndxWrapper `json:"item"`
}

type BonusWrapper struct {
	AbilityScore IndxWrapper `json:"ability_score"`
	Bonus        int         `json:"bonus"`
}

type ChoiceWrapper struct {
	Desc   []string `json:"desc"`
	Choose int      `json:"choose"`
	Type   string   `json:"type"`
	From   struct {
		OptionSetType string          `json:"option_set_type"`
		Options       []OptionWrapper `json:"options"`
	} `json:"from"`
}

type BaseRaceJSON struct {
	Indx                       string         `json:"index"`
	Name                       string         `json:"name"`
	Speed                      int            `json:"speed"`
	Size                       string         `json:"size"`
	AlignmentDesc              string         `json:"alignment"`
	AgeDesc                    string         `json:"age"`
	SizeDesc                   string         `json:"size_description"`
	LanguageDesc               string         `json:"language_desc"`
	StartingProficiencies      []IndxWrapper  `json:"proficiencies"`
	StartingProficiencyOptions ChoiceWrapper  `json:"starting_proficiency_options"`
	AbilityBonuses             []BonusWrapper `json:"ability_bonuses"`
}

type RaceJSON struct {
	*BaseRaceJSON
	// AbilityBonuses []AbilityBonus `json:"ability_bonuses"`
	// StartingProficiencies []Proficiency `json:"starting_proficiencies"`
	// StartingProficiencyOptions []ProficiencyOption `json:"starting_proficiency_options"`
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
					SetAbilityScoreID(cp.indxToId[ab.AbilityScore.Indx]).SaveX(ctx),
			)
			log.Info("Added ability bonus", "race", rr.Indx, "ability_score", ab.AbilityScore.Indx, "bonus", ab.Bonus)
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
