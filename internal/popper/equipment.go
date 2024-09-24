package popper

import (
	"context"
	"fmt"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/armor"
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
	client    *ent.Client
	dataFile  string
	data      []EquipmentJSON
	indxToId  map[string]int
	indxToCat map[string]equipment.EquipmentCategory
}

func NewEquipmentPopulator(client *ent.Client, dataDir string) *EquipmentPopulator {
	ep := &EquipmentPopulator{
		client:    client,
		dataFile:  fmt.Sprintf("%s/Equipment.json", dataDir),
		indxToId:  make(map[string]int),
		indxToCat: make(map[string]equipment.EquipmentCategory),
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

	// Put the coins in the index to id map
	coins := p.client.Coin.Query().AllX(ctx)
	for _, c := range coins {
		p.indxToId[c.Indx] = c.ID
	}

	for _, eq := range p.data {
		log.Debug("Populating equipment", "equipment", eq.Indx)
		catRaw := eq.EquipmentCategory.Indx
		catSplit := strings.Split(catRaw, "-")
		cat := strings.ToUpper(catSplit[len(catSplit)-1])
		if cat[len(cat)-1] == 'S' {
			cat = cat[:len(cat)-1]
		}

		ceq, err := p.client.Equipment.Create().
			SetIndx(eq.Indx).
			SetName(eq.Name).
			SetWeight(eq.Weight).
			SetEquipmentCategory(equipment.EquipmentCategory(cat)).
			SetCost(p.client.Cost.Create().
				SetQuantity(int(eq.Cost.Quantity)).
				SetCoinID(p.indxToId[eq.Cost.Unit]).
				SaveX(ctx)).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating equipment: %w", err)
		}
		p.indxToId[eq.Indx] = ceq.ID
		p.indxToCat[eq.Indx] = equipment.EquipmentCategory(cat)

		log.Debug("Created equipment", "ceq", eq.Indx, "id", ceq.ID)
	}

	if err := p.PopulateEdges(ctx); err != nil {
		return err
	}
	return nil
}

// FIXME
func (p *EquipmentPopulator) PopulateEdges(ctx context.Context) error {
	log.Info("Populating equipment edges")
	if len(p.data) == 0 {
		return fmt.Errorf("no equipment data to populate")
	}

	for _, eq := range p.data {
		cat := p.indxToCat[eq.Indx]
		switch cat {
		case equipment.EquipmentCategoryGear:
			if err := p.handleGearWrapper(ctx, &eq); err != nil {
				return err
			}
		case equipment.EquipmentCategoryTool:
			if err := p.handleToolWrapper(ctx, &eq); err != nil {
				return err
			}
		case equipment.EquipmentCategoryVehicle:
			if err := p.handleVehicleWrapper(ctx, &eq); err != nil {
				return err
			}
		case equipment.EquipmentCategoryWeapon:
			if err := p.handleWeaponWrapper(ctx, &eq); err != nil {
				return err
			}
		case equipment.EquipmentCategoryArmor:
			if err := p.handleArmorWrapper(ctx, &eq); err != nil {
				return err
			}
		default:
			log.Warn("Unknown equipment category", "category", cat)
		}

	}
	log.Info("Created equipment", "count", len(p.data))

	return nil
}

func (p *EquipmentPopulator) handleWeaponWrapper(ctx context.Context, eq *EquipmentJSON) error {
	log.Debug("Handling weapon", "equipment", eq.Indx)

	wps := make([]int, 0)
	for _, wp := range eq.WeaponProperties {
		wp, err := p.client.Property.Query().Where(property.Indx(wp.Indx)).Only(ctx)
		if err != nil {
			log.Warn("Error querying property", "error", err)
		}
		wps = append(wps, wp.ID)
	}

	wcreate := p.client.Weapon.Create().
		SetEquipmentID(p.indxToId[eq.Indx]).
		SetWeaponCategory(weapon.WeaponCategory(strings.ToLower(eq.WeaponCategory))).
		SetWeaponSubcategory(weapon.WeaponSubcategory(strings.ToLower(eq.WeaponRange))).
		AddPropertyIDs(wps...)

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
	if dd != nil {
		wcreate = wcreate.SetDamage(dd)
	}

	rg := ent.WeaponRange{
		RangeNormal:      0,
		RangeLong:        0,
		ThrowRangeNormal: 0,
		ThrowRangeLong:   0,
	}

	if eq.Range != nil {
		rg.RangeNormal = eq.Range.Normal
		rg.RangeLong = eq.Range.Long
	}
	if eq.ThrowRange != nil {
		rg.ThrowRangeNormal = eq.ThrowRange.Normal
		rg.ThrowRangeLong = eq.ThrowRange.Long
	}

	wr, err := p.client.WeaponRange.Create().
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

func (p *EquipmentPopulator) handleArmorWrapper(ctx context.Context, eq *EquipmentJSON) error {
	log.Debug("Handling armor", "equipment", eq.Indx)
	armorClassEntity, err := p.client.ArmorClass.
		Create().
		SetBase(eq.ArmorClass.Base).
		SetDexBonus(eq.ArmorClass.DexBonus).
		Save(ctx)
	if err != nil {
		return err
	}

	armorEntity, err := p.client.Armor.Create().
		SetEquipmentID(p.indxToId[eq.Indx]).
		SetArmorCategory(armor.ArmorCategory(strings.ToLower(eq.ArmorCategory))).
		SetArmorClass(armorClassEntity).
		SetStrMinimum(eq.StrMinimum).
		SetStealthDisadvantage(eq.StealthDis).
		Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated armor", "armor", eq.Indx, "id", armorEntity.ID)
	return nil
}

func (p *EquipmentPopulator) handleGearWrapper(ctx context.Context, eq *EquipmentJSON) error {
	log.Debug("Handling gear", "equipment", eq.Indx)
	gearCreate := p.client.Gear.Create().
		SetEquipmentID(p.indxToId[eq.Indx]).
		SetDesc(eq.Desc).
		SetGearCategory(eq.GearCategory.Indx)

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

func (p *EquipmentPopulator) handleToolWrapper(ctx context.Context, eq *EquipmentJSON) error {
	log.Debug("Handling tool", "equipment", eq.Indx)
	toolEntity, err := p.client.Tool.Create().
		SetEquipmentID(p.indxToId[eq.Indx]).
		SetDesc(eq.Desc).
		SetToolCategory(eq.ToolCategory).
		Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated tool", "tool", eq.Indx, "id", toolEntity.ID)
	return nil
}

func (p *EquipmentPopulator) handleVehicleWrapper(ctx context.Context, eq *EquipmentJSON) error {
	log.Debug("Handling vehicle", "equipment", eq.Indx)
	vehicleCreate := p.client.Vehicle.Create().
		SetEquipmentID(p.indxToId[eq.Indx]).
		SetDesc(eq.Desc).
		SetVehicleCategory(vehicle.VehicleCategory(strings.ToLower(eq.VehicleCategory)))

	if eq.Speed.Unit != "" {
		vehicleCreate = vehicleCreate.SetSpeedUnits(vehicle.SpeedUnits(eq.Speed.Unit))
		vehicleCreate = vehicleCreate.SetSpeedQuantity(eq.Speed.Quantity)
	}

	if eq.Capacity != "" {
		vehicleCreate = vehicleCreate.SetCapacity(eq.Capacity)
	}

	vehicleEntity, err := vehicleCreate.Save(ctx)
	if err != nil {
		return err
	}

	log.Info("Populated vehicle", "vehicle", eq.Indx, "id", vehicleEntity.ID)
	return nil
}
