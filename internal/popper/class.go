package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
)

type ClassJSON struct {
	Indx         string        `json:"index"`
	Name         string        `json:"name"`
	HitDie       int           `json:"hit_die"`
	SavingThrows []IndxWrapper `json:"saving_throws"`
}

type ClassPopulator struct {
	client   *ent.Client
	dataFile string
	data     []ClassJSON
	indxToId map[string]int
}

func NewClassPopulator(pp *Popper) *ClassPopulator {
	ep := &ClassPopulator{
		client:   pp.Client,
		dataFile: fmt.Sprintf("%s/Class.json", pp.DataDir),
		indxToId: pp.IndxToId,
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (cp *ClassPopulator) Populate(ctx context.Context) error {
	log.Info("Populating classes")
	if len(cp.data) == 0 {
		return fmt.Errorf("no class data to populate")
	}

	for _, class := range cp.data {
		svIDs := make([]int, 0)
		for _, savingThrow := range class.SavingThrows {
			svIDs = append(svIDs, cp.indxToId[savingThrow.Indx])
		}

		cc, err := cp.client.Class.Create().
			SetIndx(class.Indx).
			SetName(class.Name).
			SetHitDie(class.HitDie).
			AddSavingThrowIDs(svIDs...).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating class: %w", err)
		}

		cp.indxToId[class.Indx] = cc.ID
		log.Info("Class created", "index", class.Indx, "id", cc.ID)
	}

	return nil
}
