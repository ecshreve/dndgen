package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
)

type RefWrapper struct {
	Indx string `json:"index"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ProficiencyJSON struct {
	Indx      string        `json:"index"`
	ProfType  string        `json:"type"`
	Name      string        `json:"name"`
	Classes   []IndxWrapper `json:"classes"`
	Races     []IndxWrapper `json:"races"`
	Reference RefWrapper    `json:"reference"`
}

type ProficiencyPopulator struct {
	client   *ent.Client
	dataFile string
	data     []ProficiencyJSON
	indxToId map[string]int
}

func NewProficiencyPopulator(pp *Popper) *ProficiencyPopulator {
	ep := &ProficiencyPopulator{
		client:   pp.Client,
		dataFile: fmt.Sprintf("%s/Proficiency.json", pp.DataDir),
		indxToId: pp.IndxToId,
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (cp *ProficiencyPopulator) Populate(ctx context.Context) error {
	if len(cp.data) == 0 {
		return fmt.Errorf("no Proficiency data to populate")
	}

	for _, prof := range cp.data {
		profCreate := cp.client.Proficiency.Create().
			SetIndx(prof.Indx).
			SetName(prof.Name).
			SetReference(prof.Reference.URL)

		profEntity, err := profCreate.Save(ctx)
		if err != nil {
			log.Warn("error creating proficiency", "error", err)
		}
		log.Info("Populated proficiency", "indx", profEntity.Indx, "id", profEntity.ID)
	}

	return nil
}
