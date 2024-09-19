package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/internal/utils"
)

type Populator interface {
	ParseData(data []byte) error
	Populate(ctx context.Context) error
}

type EquipmentJSON struct {
	Indx string   `json:"index"`
	Name string   `json:"name"`
	Desc []string `json:"desc"`
	Cost CostJSON `json:"cost"`
}

type CostJSON struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type EquipmentPopulator struct {
	client   *ent.Client
	dataFile string
	data     []EquipmentJSON
}

func NewEquipmentPopulator(client *ent.Client, dataDir string) *EquipmentPopulator {
	return &EquipmentPopulator{
		client:   client,
		dataFile: fmt.Sprintf("%s/Equipment.json", dataDir),
	}
}

func (p *EquipmentPopulator) Populate(ctx context.Context) error {
	if err := utils.LoadJSONFile(p.dataFile, &p.data); err != nil {
		return fmt.Errorf("error loading equipment data: %w", err)
	}

	for _, equipment := range p.data {
		if len(equipment.Desc) == 0 {
			equipment.Desc = []string{}
		}

		eq, err := p.client.Equipment.Create().
			SetIndx(equipment.Indx).
			SetName(equipment.Name).
			SetDesc(equipment.Desc).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating equipment: %w", err)
		}

		coin, err := p.client.Coin.Query().
			Where(coin.IndxHasSuffix(equipment.Cost.Unit)).
			Only(ctx)
		if err != nil || coin == nil {
			return fmt.Errorf("error querying coin: %w", err)
		}

		_, err = p.client.EquipmentCost.Create().
			SetQuantity(equipment.Cost.Quantity).
			SetCoin(coin).
			SetEquipment(eq).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating equipment cost: %w", err)
		}

		// log.Debug("Created equipment", "equipment", eq.Indx)
	}
	log.Info("Created equipment", "count", len(p.data))

	return nil
}
