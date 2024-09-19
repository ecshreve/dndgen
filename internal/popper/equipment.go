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

type QuantityUnitWrapper struct {
	Quantity float32 `json:"quantity"`
	Unit     string  `json:"unit"`
}

type QuantityItemWrapper struct {
	Quantity int         `json:"quantity"`
	Item     IndxWrapper `json:"item"`
}

type DamageJSON struct {
	DamageDice string      `json:"damage_dice"`
	DamageType IndxWrapper `json:"damage_type"`
}

type RangeJSON struct {
	Normal int `json:"normal"`
	Long   int `json:"long,omitempty"`
}

type WeaponJSON struct {
	WeaponCategory   string        `json:"weapon_category"`
	WeaponRange      string        `json:"weapon_range"`
	CategoryRange    string        `json:"category_range"`
	Damage           DamageJSON    `json:"damage"`
	Range            *RangeJSON    `json:"range,omitempty"`
	ThrowRange       *RangeJSON    `json:"throw_range,omitempty"`
	WeaponProperties []IndxWrapper `json:"properties"`
}

type ArmorClassJSON struct {
	Base     int  `json:"base"`
	DexBonus bool `json:"dex_bonus,omitempty"`
	MaxBonus int  `json:"max_bonus,omitempty"`
}

type ArmorJSON struct {
	ArmorCategory string         `json:"armor_category"`
	ArmorClass    ArmorClassJSON `json:"armor_class"`
	StrMinimum    int            `json:"str_minimum"`
	StealthDis    bool           `json:"stealth_disadvantage"`
}

type GearJSON struct {
	GearCategory IndxWrapper           `json:"gear_category,omitempty"`
	Contents     []QuantityItemWrapper `json:"contents,omitempty"`
}

type ToolJSON struct {
	ToolCategory string `json:"tool_category,omitempty"`
}

type VehicleJSON struct {
	VehicleCategory string              `json:"vehicle_category,omitempty"`
	Speed           QuantityUnitWrapper `json:"speed,omitempty"`
	Capacity        string              `json:"capacity,omitempty"`
}

type EquipmentBaseJSON struct {
	Indx              string              `json:"index"`
	Name              string              `json:"name"`
	Desc              []string            `json:"desc,omitempty"`
	Weight            float32             `json:"weight,omitempty"`
	Cost              QuantityUnitWrapper `json:"cost,omitempty"`
	EquipmentCategory IndxWrapper         `json:"equipment_category"`
}

type EquipmentJSON struct {
	*EquipmentBaseJSON
	*WeaponJSON
	*ArmorJSON
	*GearJSON
	*ToolJSON
	*VehicleJSON
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
			SetQuantity(int(equipment.Cost.Quantity)).
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
