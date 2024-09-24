package popper

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/armor"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/damagetype"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/property"
	"github.com/ecshreve/dndgen/ent/vehicle"
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
		log.Debug("Created equipment", "ceq", eq.Indx, "id", ceq.ID)

		if err := handleEquipmentCost(ctx, p.client, ceq, eq.Cost); err != nil {
			return fmt.Errorf("error handling equipment cost: %w", err)
		}

		if category == "armor" {
			if err := handleArmorWrapper(ctx, p.client, ceq, eq.ArmorJSON); err != nil {
				return fmt.Errorf("error handling armor: %w", err)
			}
		}

		if category == "weapon" {
			if err := handleWeaponWrapper(ctx, p.client, ceq, eq.WeaponJSON); err != nil {
				return fmt.Errorf("error handling weapon: %w", err)
			}
		}

		if category == "vehicles" {
			if err := handleVehicleWrapper(ctx, p.client, ceq, eq.VehicleJSON); err != nil {
				return fmt.Errorf("error handling vehicle: %w", err)
			}
		}

		if category == "gear" {
			if err := handleGearWrapper(ctx, p.client, ceq, eq.GearJSON); err != nil {
				return fmt.Errorf("error handling gear: %w", err)
			}
		}

		if category == "tools" {
			if err := handleToolWrapper(ctx, p.client, ceq, eq.ToolJSON); err != nil {
				return fmt.Errorf("error handling tool: %w", err)
			}
		}
	}
	log.Info("Created equipment", "count", len(p.data))

	return nil
}

// handleEquipmentCost
func handleEquipmentCost(ctx context.Context, client *ent.Client, eq *ent.Equipment, cost QuantityUnitWrapper) error {
	log.Debug("Handling equipment cost", "equipment", eq.Indx, "cost", cost)

	coin, err := client.Coin.Query().
		Where(coin.IndxHasSuffix(cost.Unit)).
		Only(ctx)
	if err != nil || coin == nil {
		return fmt.Errorf("error querying coin: %w", err)
	}

	costEntity, err := client.Cost.Create().
		SetQuantity(int(cost.Quantity)).
		SetCoin(coin).
		SetEquipment(eq).
		Save(ctx)
	if err != nil {
		return fmt.Errorf("error creating equipment cost: %w", err)
	}

	log.Info("Populated equipment cost", "equipment", eq.Indx, "id", costEntity.ID)
	return nil
}

func handleWeaponWrapper(ctx context.Context, client *ent.Client, eq *ent.Equipment, w *WeaponJSON) error {
	log.Debug("Handling weapon", "equipment", eq.Indx)

	wps := make([]int, 0)
	for _, wp := range w.WeaponProperties {
		wp, err := client.Property.Query().Where(property.Indx(wp.Indx)).Only(ctx)
		if err != nil {
			log.Warn("Error querying property", "error", err)
		}
		wps = append(wps, wp.ID)
	}

	wcreate := client.Weapon.Create().
		SetWeaponCategory(weapon.WeaponCategory(strings.ToLower(w.WeaponCategory))).
		SetWeaponSubcategory(weapon.WeaponSubcategory(strings.ToLower(w.WeaponRange))).
		SetEquipment(eq).
		AddPropertyIDs(wps...)

	if w.Damage.DamageType.Indx == "" {
		log.Warn("No damage type", "equipment", eq.Indx)
		w.Damage.DamageType.Indx = "none"
	}
	var dt *ent.DamageType
	var dd *ent.Damage
	var damageError error

	if w.Damage.DamageType.Indx != "none" {
		dt, damageError = client.DamageType.Query().
			Where(damagetype.Indx(w.Damage.DamageType.Indx)).
			Only(ctx)
		if damageError != nil {
			log.Warn("Error querying damage type", "error", damageError)
		}

		dd, damageError = client.Damage.Create().
			SetDamageDice(w.Damage.DamageDice).
			SetDamageType(dt).
			Save(ctx)
		if damageError != nil {
			log.Warn("Error creating damage", "error", damageError)
		}
	}
	if dd != nil {
		wcreate = wcreate.SetDamage(dd)
	}

	rg := ent.WeaponRange{
		RangeNormal:      0,
		RangeLong:        0,
		ThrowRangeNormal: 0,
		ThrowRangeLong:   0,
	}

	if w.Range != nil {
		rg.RangeNormal = w.Range.Normal
		rg.RangeLong = w.Range.Long
	}
	if w.ThrowRange != nil {
		rg.ThrowRangeNormal = w.ThrowRange.Normal
		rg.ThrowRangeLong = w.ThrowRange.Long
	}

	wr, err := client.WeaponRange.Create().
		SetRangeNormal(rg.RangeNormal).
		SetRangeLong(rg.RangeLong).
		SetThrowRangeNormal(rg.ThrowRangeNormal).
		SetThrowRangeLong(rg.ThrowRangeLong).
		Save(ctx)
	if err != nil {
		log.Warn("Error creating weapon range", "error", err)
	}
	wcreate = wcreate.SetWeaponRange(wr)

	ww, err := wcreate.Save(ctx)
	if err != nil {
		log.Warn("Error creating weapon", "error", err)
	}

	log.Info("Populated weapon", "weapon", eq.Indx, "id", ww.ID)
	return nil
}

func handleArmorWrapper(ctx context.Context, client *ent.Client, eq *ent.Equipment, ar *ArmorJSON) error {
	log.Debug("Handling armor", "equipment", eq.Indx)
	armorClassEntity, err := client.ArmorClass.
		Create().
		SetBase(ar.ArmorClass.Base).
		SetDexBonus(ar.ArmorClass.DexBonus).
		Save(ctx)
	if err != nil {
		return err
	}

	armorEntity, err := client.Armor.Create().
		SetArmorCategory(armor.ArmorCategory(strings.ToLower(ar.ArmorCategory))).
		SetArmorClass(armorClassEntity).
		SetStrMinimum(ar.StrMinimum).
		SetStealthDisadvantage(ar.StealthDis).
		SetEquipment(eq).
		Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated armor", "armor", eq.Indx, "id", armorEntity.ID)
	return nil
}

func handleGearWrapper(ctx context.Context, client *ent.Client, eq *ent.Equipment, g *GearJSON) error {
	log.Debug("Handling gear", "equipment", eq.Indx)
	gearCreate := client.Gear.Create().
		SetGearCategory(g.GearCategory.Indx).
		SetEquipment(eq)

	// for _, item := range g.Contents {
	// 	gearCreate = gearCreate.AddGearContents(item.Item)
	// }

	gearEntity, err := gearCreate.Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated gear", "gear", eq.Indx, "id", gearEntity.ID)
	return nil
}

func handleToolWrapper(ctx context.Context, client *ent.Client, eq *ent.Equipment, t *ToolJSON) error {
	log.Debug("Handling tool", "equipment", eq.Indx)
	toolEntity, err := client.Tool.Create().
		SetToolCategory(t.ToolCategory).
		SetEquipment(eq).
		Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated tool", "tool", eq.Indx, "id", toolEntity.ID)
	return nil
}

func handleVehicleWrapper(ctx context.Context, client *ent.Client, eq *ent.Equipment, v *VehicleJSON) error {
	log.Debug("Handling vehicle", "equipment", eq.Indx)
	vehicleCreate := client.Vehicle.Create().
		SetVehicleCategory(vehicle.VehicleCategory(v.VehicleCategory)).
		SetEquipment(eq)

	if v.Speed.Unit != "" {
		vehicleCreate = vehicleCreate.SetSpeedUnits(vehicle.SpeedUnits(v.Speed.Unit))
		vehicleCreate = vehicleCreate.SetSpeedQuantity(v.Speed.Quantity)
	}

	if v.Capacity != "" {
		vehicleCreate = vehicleCreate.SetCapacity(v.Capacity)
	}

	vehicleEntity, err := vehicleCreate.Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated vehicle", "vehicle", eq.Indx, "id", vehicleEntity.ID)
	return nil
}
