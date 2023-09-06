package popper

import (
	"context"
	"encoding/json"
	"fmt"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/coin"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/gear"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
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
	Desc              []string    `json:"desc"`
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
	fpath := "data/Equipment.json"
	var v []EquipmentWrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, ww := range v {

		vv := ent.Equipment{
			Indx:              ww.Indx,
			Name:              ww.Name,
			EquipmentCategory: equipment.EquipmentCategory(strings.Replace(ww.EquipmentCategory.Indx, "-", "_", -1)),
		}

		cn := p.Client.Coin.Query().Where(coin.IndxEQ(ww.Cost.Unit)).FirstX(ctx)
		// val := cn.GoldConversionRate * float64(ww.Cost.Quantity)
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
				WeaponRange:    strings.Replace(ww.WeaponRange, "-", "_", -1),
				WeaponCategory: strings.Replace(ww.WeaponCategory, "-", "_", -1),
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
			subcat := strings.ToLower(fmt.Sprintf("%s_%s", w.WeaponCategory, w.WeaponRange))
			eq.Update().SetEquipmentSubcategory(subcat).SaveX(ctx)

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
			subcat := strings.ToLower(fmt.Sprintf("%s_%s", a.ArmorCategory, "armor"))
			eq.Update().SetEquipmentSubcategory(subcat).SaveX(ctx)
		}

		if ww.GearWrapper != nil && ww.EquipmentCategory.Indx == "adventuring-gear" {
			a := ent.Gear{
				Indx:         ww.Indx,
				Name:         ww.Name,
				Quantity:     intOrDef(ww.Quantity, 0),
				GearCategory: gear.GearCategory(strings.Replace(ww.GearCategory.Indx, "-", "_", -1)),
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
			eq.Update().SetEquipmentSubcategory(string(a.GearCategory)).SaveX(ctx)

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

			eq.Update().SetEquipmentSubcategory(a.ToolCategory).SaveX(ctx)
		}

		if ww.VehicleWrapper != nil && ww.EquipmentCategory.Indx == "mounts-and-vehicles" {
			a := ent.Vehicle{
				Indx:            ww.Indx,
				Name:            ww.Name,
				VehicleCategory: strings.Replace(ww.VehicleCategory, "-", "_", -1),
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

			eq.Update().SetEquipmentSubcategory(a.VehicleCategory).SaveX(ctx)
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
