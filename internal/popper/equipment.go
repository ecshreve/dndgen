package popper

import (
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/internal/utils"
)

type IndxWrapper struct {
	Indx string `json:"index"`
}

type Range struct {
	Normal int `json:"normal,omitempty"`
	Long   int `json:"long,omitempty"`
}

type WeaponWrapper struct {
	WeaponCategory string `json:"weapon_category"`
	WeaponRange    string `json:"weapon_range"`
	MeleeRange     *Range `json:"range,omitempty"`
	ThrowRange     *Range `json:"throw_range,omitempty"`
	Damage         struct {
		Dice  string      `json:"damage_dice,omitempty"`
		DType IndxWrapper `json:"damage_type,omitempty"`
	} `json:"damage,omitempty"`
	Properties []IndxWrapper `json:"properties,omitempty"`
}

type ArmorWrapper struct {
	ArmorCategory       string `json:"armor_category"`
	StealthDisadvantage bool   `json:"stealth_disadvantage"`
	StrMinimum          int    `json:"str_minimum,omitempty"`
	ArmorClass          struct {
		Base     int  `json:"base"`
		DexBonus bool `json:"dex_bonus"`
		MaxBonus int  `json:"max_bonus,omitempty"`
	} `json:"armor_class"`
}

type GearWrapper struct {
	GearCategory IndxWrapper `json:"gear_category"`
	Quantity     *int        `json:"quantity,omitempty"`
}

type ToolWrapper struct {
	ToolCategory string `json:"tool_category"`
}

type VehicleWrapper struct {
	VehicleCategory string `json:"vehicle_category"`
}

type CostWrapper struct {
	Quantity int    `json:"quantity"`
	Unit     string `json:"unit"`
}

type EquipmentSubWrapper struct {
	Name              string      `json:"name"`
	Cost              CostWrapper `json:"cost"`
	Weight            float64     `json:"weight"`
	EquipmentCategory IndxWrapper `json:"equipment_category"`
}

type EquipmentWrapper struct {
	Indx string `json:"index"`
	EquipmentSubWrapper
	*WeaponWrapper
	*ArmorWrapper
	*GearWrapper
	*ToolWrapper
	*VehicleWrapper
}

func (p *Popper) PopulateEquipmentFromFile(ctx context.Context, jsonFile string) ([]int, error) {
	os.Setenv("EQUIPMENT_FILE", jsonFile)
	return p.PopulateEquipment(ctx)
}

// // PopulateEquipment populates the Equipment entities from the JSON data files.
func (p *Popper) PopulateEquipment(ctx context.Context) ([]int, error) {
	fname := os.Getenv("EQUIPMENT_FILE")
	if fname == "" {
		fname = "data/Equipment.json"
	}

	// TODO: move to a const or a config file or parse from the JSON file
	equipmentCategories := map[string][]string{
		"weapon": {
			"simple",
			"martial",
			"melee",
			"ranged",
		},
		"armor": {
			"light",
			"medium",
			"heavy",
			"shield",
		},
		"adventuring_gear": {
			"equipment_packs",
			"standard_gear",
			"ammunition",
			"arcane_foci",
			"druidic_foci",
			"holy_symbols",
			"kits",
		},
		"tools": {
			"artisans_tools",
			"gaming_sets",
			"musical_instruments",
			"other_tools",
		},
		"mounts_and_vehicles": {
			"land",
			"waterborne",
		},
		"other": {},
	}

	for cat, subs := range equipmentCategories {
		parent, err := p.Client.EquipmentCategory.Create().SetIndx(cat).SetName(cat).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create parent entity for type EquipmentCategory %s: %w", cat, err)
		}
		log.Infof("created parent entity for type EquipmentCategory %s", cat)

		// Create sub-categories
		creates := []*ent.EquipmentCategoryCreate{}
		for _, sub := range subs {
			creates = append(creates, p.Client.EquipmentCategory.Create().SetIndx(sub).SetName(sub).SetParentCategoryID(parent.ID))
		}
		subcats, err := p.Client.EquipmentCategory.CreateBulk(creates...).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create sub-categories for type EquipmentCategory %s: %w", cat, err)
		}
		log.Infof("created %d sub-categories for type EquipmentCategory %s", len(subcats), cat)

		p.IdToIndx[parent.ID] = cat
		p.IndxToId[cat] = parent.ID

		// Add sub-categories to the map
		for _, c := range subcats {
			p.IdToIndx[c.ID] = c.Indx
			p.IndxToId[c.Indx] = c.ID
		}
	}

	var v []EquipmentWrapper
	if err := utils.LoadJSONFile(fname, &v); err != nil {
		return nil, fmt.Errorf("LoadJSONFile: %w", err)
	}

	allEquipmentIds := []int{}
	for _, ww := range v {
		vv := ent.Equipment{
			Indx:                ww.Indx,
			Name:                ww.Name,
			Weight:              int(ww.Weight), // TODO: handle float conversion
			EquipmentCategoryID: p.IndxToId[strings.ToLower(strings.Replace(ww.EquipmentCategory.Indx, "-", "_", -1))],
		}

		eq, err := p.Client.Equipment.Create().
			SetEquipment(&vv).
			Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create entity %s: %w", vv.Indx, err)
		}

		cn := p.Client.Coin.Query().Where(coin.IndxContains(ww.Cost.Unit)).FirstX(ctx)
		eqc := &ent.EquipmentCost{
			EquipmentID: eq.ID,
			Quantity:    ww.Cost.Quantity,
			CoinID:      cn.ID,
			GpValue:     cn.GoldConversionRate * float64(ww.Cost.Quantity),
		}
		_, err = p.Client.EquipmentCost.Create().SetEquipmentCost(eqc).Save(ctx)
		if err != nil {
			return nil, fmt.Errorf("unable to create eq cost %s: %w", vv.Indx, err)
		}

		if ww.WeaponWrapper != nil {
			w := ent.Weapon{
				Indx:           ww.Indx,
				Name:           ww.Name,
				WeaponRange:    strings.ToLower(strings.Replace(ww.WeaponRange, "-", "_", -1)),
				WeaponCategory: strings.ToLower(strings.Replace(ww.WeaponCategory, "-", "_", -1)),
			}

			propIDs := p.GetIDsFromIndxWrappers(ww.Properties)
			created, err := p.Client.Weapon.Create().SetWeapon(&w).SetEquipment(eq).AddWeaponPropertyIDs(propIDs...).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %s: %w", vv.Indx, err)
			}

			if ww.WeaponWrapper.Damage.Dice != "" {
				did := p.IndxToId[ww.WeaponWrapper.Damage.DType.Indx]
				_, err = p.Client.WeaponDamage.Create().SetWeaponID(created.ID).SetDamageTypeID(did).SetDice(ww.WeaponWrapper.Damage.Dice).Save(ctx)
				if err != nil {
					return nil, fmt.Errorf("unable to create entity %s: %w", vv.Indx, err)
				}
			}
		}

		if ww.ArmorWrapper != nil && ww.EquipmentCategory.Indx == "armor" {
			a := ent.Armor{
				Indx:                ww.Indx,
				Name:                ww.Name,
				StealthDisadvantage: ww.ArmorWrapper.StealthDisadvantage,
				MinStrength:         ww.ArmorWrapper.StrMinimum,
				ArmorCategory:       strings.ToLower(strings.Split(ww.ArmorWrapper.ArmorCategory, "-")[0]),
			}

			b := ent.ArmorClass{
				Base:     ww.ArmorWrapper.ArmorClass.Base,
				DexBonus: ww.ArmorWrapper.ArmorClass.DexBonus,
				MaxBonus: ww.ArmorWrapper.ArmorClass.MaxBonus,
			}

			ac, err := p.Client.ArmorClass.Create().SetArmorClass(&b).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %v: %w", ac, err)
			}

			_, err = p.Client.Armor.Create().SetArmor(&a).AddArmorClass(ac).SetEquipment(eq).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %v: %w", ac, err)
			}
		}

		if ww.GearWrapper != nil && ww.EquipmentCategory.Indx == "adventuring-gear" {
			a := ent.Gear{
				Indx:         ww.Indx,
				Name:         ww.Name,
				Quantity:     utils.IntOrDefault(ww.Quantity, 0),
				GearCategory: strings.Replace(ww.GearCategory.Indx, "-", "_", -1),
			}

			_, err = p.Client.Gear.Create().SetGear(&a).SetEquipment(eq).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %v: %w", a, err)
			}

		}

		if ww.ToolWrapper != nil && ww.EquipmentCategory.Indx == "tools" {
			a := ent.Tool{
				Indx:         ww.Indx,
				Name:         ww.Name,
				ToolCategory: utils.CleanString(ww.ToolCategory),
			}

			_, err = p.Client.Tool.Create().SetTool(&a).SetEquipment(eq).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %v: %w", a, err)
			}
		}

		if ww.VehicleWrapper != nil && ww.EquipmentCategory.Indx == "mounts-and-vehicles" {
			a := ent.Vehicle{
				Indx:            ww.Indx,
				Name:            ww.Name,
				VehicleCategory: "land",
			}

			if catraw := strings.Contains(ww.VehicleCategory, "Waterborne"); catraw {
				a.VehicleCategory = "waterborne"
			}
			_, err = p.Client.Vehicle.Create().SetVehicle(&a).SetEquipment(eq).Save(ctx)
			if err != nil {
				return nil, fmt.Errorf("unable to create entity %v: %w", a, err)
			}
		}

		p.IdToIndx[eq.ID] = vv.Indx
		p.IndxToId[vv.Indx] = eq.ID
		allEquipmentIds = append(allEquipmentIds, eq.ID)
	}
	log.Infof("created %d entities for type Equipment", p.Client.Equipment.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Weapon", p.Client.Weapon.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Armor", p.Client.Armor.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Gear", p.Client.Gear.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Tool", p.Client.Tool.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Vehicle", p.Client.Vehicle.Query().CountX(ctx))

	return allEquipmentIds, nil
}
