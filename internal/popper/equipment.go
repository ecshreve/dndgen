package popper

import (
	"context"

	"github.com/kr/pretty"
	"github.com/samsarahq/go/oops"
)

type Range struct {
	Normal int `json:"normal,omitempty"`
	Long   int `json:"long,omitempty"`
}

type Weap struct {
	WeaponCategoryID string `json:"weapon_category"`
	WeaponRangeID    string `json:"weapon_range"`
	MeleeRange       Range  `json:"range,omitempty"`
	ThrowRange       Range  `json:"throw_range,omitempty"`
}

type Arm struct {
	ArmorCategoryID     string `json:"armor_category"`
	StealthDisadvantage bool   `json:"stealth_disadvantage"`
}

type Wrapcom struct {
	Name string `json:"name"`
	Cost struct {
		Quantity int    `json:"quantity"`
		Unit     string `json:"unit"`
	} `json:"cost"`
	Weight   float64   `json:"weight"`
	Category IDWrapper `json:"equipment_category"`
}

// PopulateEquipment populates the Equipment entities from the JSON data files.
func (p *Popper) PopulateEquipment(ctx context.Context) error {
	fpath := "internal/popper/data/Equipment.json"
	type Wrapper struct {
		Indx string `json:"index"`
		Wrapcom
		Weap
		Arm
	}
	var v []Wrapper

	if err := LoadJSONFile(fpath, &v); err != nil {
		return oops.Wrapf(err, "unable to load JSON file %s", fpath)
	}
	pretty.Print(v[0])
	pretty.Print(v[40])

	// for _, ww := range v {
	// 	vv := ent.Equipment{ww}
	// 	_, err := p.Client.Equipment.Create().SetEquipment(&vv).Save(ctx)
	// 	if ent.IsConstraintError(err) {
	// 		log.Debugf("constraint failed, skipping %s", vv.Indx)
	// 		continue
	// 	}
	// 	if err != nil {
	// 		return oops.Wrapf(err, "unable to create entity %s", vv.Indx)
	// 	}
	// }

	return nil
}
