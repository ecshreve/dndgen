package popper

import (
	"context"
	"encoding/json"
	"strings"

	"github.com/charmbracelet/log"
	"github.com/samsarahq/go/oops"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/coin"
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
	MeleeRange     Range  `json:"range,omitempty"`
	ThrowRange     Range  `json:"throw_range,omitempty"`
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

// // PopulateEquipment populates the Equipment entities from the JSON data files.
func (p *Popper) PopulateEquipment(ctx context.Context) error {
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
		parent := p.Client.EquipmentCategory.Create().SetName(cat).SaveX(ctx)
		log.Infof("created parent entity for type EquipmentCategory %s", cat)

		creates := []*ent.EquipmentCategoryCreate{}
		for _, sub := range subs {
			creates = append(creates, p.Client.EquipmentCategory.Create().SetName(sub).SetParentCategoryID(parent.ID))
		}
		cats := p.Client.EquipmentCategory.CreateBulk(creates...).SaveX(ctx)
		log.Infof("created %d entities for type EquipmentCategory %s", len(cats), cat)

		p.IdToIndx[parent.ID] = string(cat)
		p.IndxToId[cat] = parent.ID
		for _, c := range cats {
			p.IdToIndx[c.ID] = string(c.Name)
			p.IndxToId[string(c.Name)] = c.ID
		}
	}

	fpath := "data/Equipment.json"
	var v []EquipmentWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, ww := range v {
		vv := ent.Equipment{
			Indx:   ww.Indx,
			Name:   ww.Name,
			Weight: int(ww.Weight), // TODO: handle float conversion
		}

		cn := p.Client.Coin.Query().Where(coin.IndxEQ(ww.Cost.Unit)).FirstX(ctx)

		eq, err := p.Client.Equipment.Create().
			SetEquipment(&vv).
			Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			log.Debug(err)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}
		eq.Update().AddEquipmentCategoryIDs(p.IndxToId[strings.ToLower(strings.Replace(ww.EquipmentCategory.Indx, "-", "_", -1))]).SaveX(ctx)

		eqc := &ent.EquipmentCost{
			EquipmentID: eq.ID,
			Quantity:    ww.Cost.Quantity,
			CoinID:      cn.ID,
			GpValue:     cn.GoldConversionRate * float64(ww.Cost.Quantity),
		}
		_, err = p.Client.EquipmentCost.Create().SetEquipmentCost(eqc).Save(ctx)
		if err != nil {
			return oops.Wrapf(err, "unable to create eq cost %s", vv.Indx)
		}

		if ww.WeaponWrapper != nil {
			w := ent.Weapon{
				Indx:           ww.Indx,
				Name:           ww.Name,
				WeaponRange:    strings.ToLower(strings.Replace(ww.WeaponRange, "-", "_", -1)),
				WeaponCategory: strings.ToLower(strings.Replace(ww.WeaponCategory, "-", "_", -1)),
			}

			created, err := p.Client.Weapon.Create().SetWeapon(&w).SetEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debugf("constraint failed, skipping %s", vv.Indx)
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
			}
			eq.Update().AddEquipmentCategoryIDs(p.IndxToId[w.WeaponCategory], p.IndxToId[created.WeaponRange]).SaveX(ctx)

			jj, _ := json.Marshal(ww.Properties)
			propIDs := p.GetIDsFromIndxs(jj)
			created.Update().AddWeaponPropertyIDs(propIDs...).SaveX(ctx)

			if ww.WeaponWrapper.Damage.Dice != "" {
				did := p.IndxToId[ww.WeaponWrapper.Damage.DType.Indx]
				_, err = p.Client.WeaponDamage.Create().SetWeaponID(created.ID).SetDamageTypeID(did).SetDice(ww.WeaponWrapper.Damage.Dice).Save(ctx)
				if ent.IsConstraintError(err) {
					log.Debugf("constraint failed, skipping %s", vv.Indx)
					log.Debug(err)
					continue
				}
				if err != nil {
					return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
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
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", ac)
			}

			_, err = p.Client.Armor.Create().SetArmor(&a).AddArmorClass(ac).SetEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", ac)
			}
			eq.Update().AddEquipmentCategoryIDs(p.IndxToId[a.ArmorCategory]).SaveX(ctx)
		}

		if ww.GearWrapper != nil && ww.EquipmentCategory.Indx == "adventuring-gear" {
			a := ent.Gear{
				Indx:         ww.Indx,
				Name:         ww.Name,
				Quantity:     intOrDef(ww.Quantity, 0),
				GearCategory: strings.Replace(ww.GearCategory.Indx, "-", "_", -1),
			}

			_, err = p.Client.Gear.Create().SetGear(&a).SetEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
			eq.Update().AddEquipmentCategoryIDs(p.IndxToId[string(a.GearCategory)]).SaveX(ctx)

		}

		if ww.ToolWrapper != nil && ww.EquipmentCategory.Indx == "tools" {
			a := ent.Tool{
				Indx:         ww.Indx,
				Name:         ww.Name,
				ToolCategory: cleanString(ww.ToolCategory),
			}

			_, err = p.Client.Tool.Create().SetTool(&a).SetEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
			eq.Update().AddEquipmentCategoryIDs(p.IndxToId[a.ToolCategory]).SaveX(ctx)
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
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
			eq.Update().AddEquipmentCategoryIDs(p.IndxToId[a.VehicleCategory]).SaveX(ctx)
		}

		p.IdToIndx[eq.ID] = vv.Indx
		p.IndxToId[vv.Indx] = eq.ID
	}
	log.Infof("created %d entities for type Equipment", p.Client.Equipment.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Weapon", p.Client.Weapon.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Armor", p.Client.Armor.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Gear", p.Client.Gear.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Tool", p.Client.Tool.Query().CountX(ctx))
	log.Infof("-> created %d entities for type Vehicle", p.Client.Vehicle.Query().CountX(ctx))

	return nil
}
