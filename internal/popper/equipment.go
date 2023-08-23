package popper

import (
	"context"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/ecshreve/dndgen/internal/util"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Range struct {
	Normal int `json:"normal,omitempty"`
	Long   int `json:"long,omitempty"`
}

type WeaponWrapper struct {
	WeaponCategory string `json:"weapon_category"`
	WeaponRange    string `json:"weapon_range"`
	MeleeRange     Range  `json:"range,omitempty"`
	ThrowRange     Range  `json:"throw_range,omitempty"`
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

type EquipmentWrapper struct {
	Name string `json:"name"`
	Cost struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"cost"`
	Weight            float64     `json:"weight"`
	EquipmentCategory IndxWrapper `json:"equipment_category"`
	Desc              []string    `json:"desc"`
}

// PopulateEquipment populates the Equipment entities from the JSON data files.
func (p *Popper) PopulateEquipment(ctx context.Context) error {
	fpath := "internal/popper/data/Equipment.json"
	type Wrapper struct {
		Indx string `json:"index"`
		EquipmentWrapper
		*WeaponWrapper
		*ArmorWrapper
		*GearWrapper
		*ToolWrapper
		*VehicleWrapper
	}
	var v []Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, ww := range v {
		vv := ent.Equipment{
			Indx:              ww.Indx,
			Name:              ww.Name,
			EquipmentCategory: equipment.EquipmentCategory(strings.Replace(ww.EquipmentCategory.Indx, "-", "_", -1)),
		}

		eq, err := p.Client.Equipment.Create().SetEquipment(&vv).Save(ctx)
		if ent.IsConstraintError(err) {
			log.Debugf("constraint failed, skipping %s", vv.Indx)
			log.Debug(err)
			continue
		}
		if err != nil {
			return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
		}

		if ww.WeaponWrapper != nil {
			w := ent.Weapon{
				Indx:        ww.Indx,
				Name:        ww.Name,
				WeaponRange: strings.Replace(ww.WeaponRange, "-", "_", -1),
			}

			_, err := p.Client.Weapon.Create().SetWeapon(&w).AddEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debugf("constraint failed, skipping %s", vv.Indx)
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
			}
		}

		if ww.ArmorWrapper != nil && ww.EquipmentCategory.Indx == "armor" {
			a := ent.Armor{
				Indx:                ww.Indx,
				Name:                ww.Name,
				StealthDisadvantage: ww.ArmorWrapper.StealthDisadvantage,
				MinStrength:         ww.ArmorWrapper.StrMinimum,
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

			_, err = p.Client.Armor.Create().SetArmor(&a).AddArmorClass(ac).AddEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", ac)
			}
		}

		if ww.GearWrapper != nil && ww.EquipmentCategory.Indx == "adventuring-gear" {
			a := ent.Gear{
				Indx:         ww.Indx,
				Name:         ww.Name,
				Quantity:     util.IntOrDef(ww.Quantity, 0),
				GearCategory: gear.GearCategory(strings.Replace(ww.GearCategory.Indx, "-", "_", -1)),
			}

			_, err = p.Client.Gear.Create().SetGear(&a).AddEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
		}

		if ww.ToolWrapper != nil && ww.EquipmentCategory.Indx == "tools" {
			a := ent.Tool{
				Indx: ww.Indx,
				Name: ww.Name,
			}

			_, err = p.Client.Tool.Create().SetTool(&a).AddEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
		}

		if ww.VehicleWrapper != nil && ww.EquipmentCategory.Indx == "mounts-and-vehicles" {
			a := ent.Vehicle{
				Indx:            ww.Indx,
				Name:            ww.Name,
				VehicleCategory: strings.Replace(ww.VehicleCategory, "-", "_", -1),
			}

			_, err = p.Client.Vehicle.Create().SetVehicle(&a).AddEquipment(eq).Save(ctx)
			if ent.IsConstraintError(err) {
				log.Debug("constraint failed, skipping")
				log.Debug(err)
				continue
			}
			if err != nil {
				return oops.Wrapf(err, "unable to create entity %v", a)
			}
		}

		p.IdToIndx[eq.ID] = vv.Indx
		p.IndxToId[vv.Indx] = eq.ID
	}

	return nil
}
