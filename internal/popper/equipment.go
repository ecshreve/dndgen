package popper

import (
	"context"
	"strings"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/samsarahq/go/oops"
	log "github.com/sirupsen/logrus"
)

type Range struct {
	Normal int `json:"normal,omitempty"`
	Long   int `json:"long,omitempty"`
}

type WeaponWrapper struct {
	WeaponCategoryID string `json:"weapon_category"`
	WeaponRangeID    string `json:"weapon_range"`
	MeleeRange       Range  `json:"range,omitempty"`
	ThrowRange       Range  `json:"throw_range,omitempty"`
}

type ArmorWrapper struct {
	ArmorCategoryID     string `json:"armor_category"`
	StealthDisadvantage bool   `json:"stealth_disadvantage"`
	StrMinimum          int    `json:"str_minimum,omitempty"`
	ArmorClass          struct {
		Base     int  `json:"base"`
		DexBonus bool `json:"dex_bonus"`
		MaxBonus int  `json:"max_bonus,omitempty"`
	} `json:"armor_class"`
}

type CommonWrapper struct {
	Name string `json:"name"`
	Cost struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"cost"`
	Weight   float64     `json:"weight"`
	Category IndxWrapper `json:"equipment_category"`
}

// PopulateEquipment populates the Equipment entities from the JSON data files.
func (p *Popper) PopulateEquipment(ctx context.Context) error {
	fpath := "internal/popper/data/Equipment.json"
	type Wrapper struct {
		Indx string `json:"index"`
		CommonWrapper
		*WeaponWrapper
		*ArmorWrapper
	}
	var v []Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}

	for _, ww := range v {
		vv := ent.Equipment{
			Indx:     ww.Indx,
			Name:     ww.Name,
			Category: equipment.Category(strings.Replace(ww.Category.Indx, "-", "_", -1)),
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
				WeaponRange: strings.Replace(ww.WeaponRangeID, "-", "_", -1),
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

		if ww.ArmorWrapper != nil {
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

	}

	return nil
}
