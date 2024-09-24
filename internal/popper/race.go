package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
)

type RaceJSON struct {
	*ent.Race
}

type RacePopulator struct {
	client   *ent.Client
	dataFile string
	data     []RaceJSON
	indxToId map[string]int
}

func NewRacePopulator(pp *Popper) *RacePopulator {
	ep := &RacePopulator{
		client:   pp.Client,
		dataFile: fmt.Sprintf("%s/Race.json", pp.DataDir),
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
			SetIndx(rr.Race.Indx).
			SetName(rr.Race.Name).
			SetSpeed(rr.Race.Speed).
			SetSize(rr.Race.Size).
			SetSizeDesc(rr.Race.SizeDesc).
			SetAlignmentDesc(rr.Race.AlignmentDesc).
			SetAgeDesc(rr.Race.AgeDesc).
			SetLanguageDesc(rr.Race.LanguageDesc).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating race: %w", err)
		}

		log.Info("Created race", "race", raceEntity.Indx)
	}

	return nil
}
