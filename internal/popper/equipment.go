package popper

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/internal/utils"
)

type Populator interface {
	PopulateFields(ctx context.Context) error
	PopulateEdges(ctx context.Context) error
}

type QuantityUnitWrapper struct {
	Quantity float64 `json:"quantity"`
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
	Weight            float64             `json:"weight,omitempty"`
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
	indxToId map[string]int
}

func NewEquipmentPopulator(client *ent.Client, dataDir string) *EquipmentPopulator {
	ep := &EquipmentPopulator{
		client:   client,
		dataFile: fmt.Sprintf("%s/Equipment.json", dataDir),
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (p *EquipmentPopulator) PopulateFields(ctx context.Context) error {
	log.Info("Populating equipment fields")
	if len(p.data) == 0 {
		return fmt.Errorf("no equipment data to populate")
	}

	// Create a map to store the index to ID mappings
	p.indxToId = make(map[string]int)

	// Add the WeaponProperties to the map
	allWeaponProperties := p.client.Property.Query().AllX(ctx)
	for _, wp := range allWeaponProperties {
		p.indxToId[wp.Indx] = wp.ID
	}

	for _, eq := range p.data {
		if len(eq.Desc) == 0 {
			eq.Desc = []string{}
		}

		rawCat := strings.Split(eq.EquipmentCategory.Indx, "-")
		category := rawCat[len(rawCat)-1]

		ceq, err := p.client.Equipment.Create().
			SetIndx(eq.Indx).
			SetName(eq.Name).
			SetDesc(eq.Desc).
			SetWeight(eq.Weight).
			SetEquipmentCategory(equipment.EquipmentCategory(category)).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating equipment: %w", err)
		}
		p.indxToId[eq.Indx] = ceq.ID
		log.Debug("Populated equipment", "equipment", eq.Indx, "id", ceq.ID)

		coin, err := p.client.Coin.Query().
			Where(coin.IndxHasSuffix(eq.Cost.Unit)).
			Only(ctx)
		if err != nil || coin == nil {
			return fmt.Errorf("error querying coin: %w", err)
		}
		p.indxToId[coin.Indx] = coin.ID

		_, err = p.client.EquipmentCost.Create().
			SetQuantity(int(eq.Cost.Quantity)).
			SetCoin(coin).
			SetEquipment(ceq).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating equipment cost: %w", err)
		}

		if eq.EquipmentCategory.Indx == "weapon" {
			if eq.Damage.DamageType.Indx == "" {
				log.Warn("No damage type", "equipment", eq.Indx)
				eq.Damage.DamageType.Indx = "none"
			}
			var dt *ent.DamageType
			var dd *ent.Damage
			var damageError error

			if eq.Damage.DamageType.Indx != "none" {
				dt, damageError = p.client.DamageType.Query().
					Where(damagetype.Indx(eq.Damage.DamageType.Indx)).
					Only(ctx)
				if damageError != nil {
					log.Warn("Error querying damage type", "error", damageError)
				}

				dd, damageError = p.client.Damage.Create().
					SetDamageDice(eq.Damage.DamageDice).
					SetDamageType(dt).
					Save(ctx)
				if damageError != nil {
					log.Warn("Error creating damage", "error", damageError)
				}
			}

			wps := make([]int, 0)
			for _, wp := range eq.WeaponProperties {
				wps = append(wps, p.indxToId[wp.Indx])
			}

			wcreate := p.client.Weapon.Create().
				SetWeaponCategory(weapon.WeaponCategory(strings.ToLower(eq.WeaponCategory))).
				SetWeaponSubcategory(weapon.WeaponSubcategory(strings.ToLower(eq.WeaponRange))).
				SetEquipment(ceq).
				AddPropertyIDs(wps...)
			if err != nil {
				log.Warn("Error creating weapon", "error", err)
			}

			if dd != nil {
				wcreate = wcreate.SetDamage(dd)
			}
			ww, err := wcreate.Save(ctx)
			if err != nil {
				log.Warn("Error creating weapon", "error", err)
			}
			log.Debug("Populated weapon", "weapon", eq.Indx, "id", ww.ID)
		}
	}
	log.Info("Created equipment", "count", len(p.data))

	return nil
}
