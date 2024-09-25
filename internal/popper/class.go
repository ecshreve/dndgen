package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/internal/utils"
)

type EquipmentWrapper struct {
	Equipment IndxWrapper `json:"equipment"`
	Quantity  int         `json:"quantity"`
}

type ClassJSON struct {
	Indx               string             `json:"index"`
	Name               string             `json:"name"`
	HitDie             int                `json:"hit_die"`
	Proficiencies      []IndxWrapper      `json:"proficiencies"`
	ProficiencyChoices []ChoiceWrapper    `json:"proficiency_choices"`
	StartingEquipment  []EquipmentWrapper `json:"starting_equipment"`
}

type ClassPopulator struct {
	client   *ent.Client
	dataFile string
	data     []ClassJSON
	indxToId map[string]int
}

func NewClassPopulator(pp *Popper) *ClassPopulator {
	ep := &ClassPopulator{
		client:   pp.Client,
		dataFile: fmt.Sprintf("%s/Class.json", pp.DataDir),
		indxToId: pp.IndxToId,
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (cp *ClassPopulator) Populate(ctx context.Context) error {
	log.Info("Populating classes")
	if len(cp.data) == 0 {
		return fmt.Errorf("no class data to populate")
	}

	for _, class := range cp.data {
		cc, err := cp.client.Class.Create().
			SetIndx(class.Indx).
			SetName(class.Name).
			SetHitDie(class.HitDie).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error creating class: %w", err)
		}

		cp.indxToId[class.Indx] = cc.ID
		log.Info("Class created", "index", class.Indx, "id", cc.ID)
	}

	if err := cp.populateProficiencies(ctx); err != nil {
		return fmt.Errorf("error populating proficiencies: %w", err)
	}
	if err := cp.populateStartingEquipment(ctx); err != nil {
		return fmt.Errorf("error populating starting equipment: %w", err)
	}

	return nil
}

func (cp *ClassPopulator) populateProficiencies(ctx context.Context) error {

	for _, cc := range cp.data {
		classUpdate := cp.client.Class.UpdateOneID(cp.indxToId[cc.Indx])

		profIDs := []int{}
		for _, p := range cc.Proficiencies {
			profIDs = append(profIDs, cp.indxToId[p.Indx])
		}
		classUpdate = classUpdate.AddProficiencyIDs(profIDs...)

		if err := classUpdate.Exec(ctx); err != nil {
			return fmt.Errorf("error saving class: %w", err)
		}
		log.Info("Class updated", "index", cc.Indx, "name", cc.Name, "proficiencies", profIDs)
	}

	log.Info("Class proficiencies populated")
	return nil
}

func (cp *ClassPopulator) populateStartingEquipment(ctx context.Context) error {
	for _, cc := range cp.data {
		for _, e := range cc.StartingEquipment {
			equipID := cp.client.Equipment.Query().
				Where(equipment.Indx(e.Equipment.Indx)).
				OnlyIDX(ctx)
			classID := cp.client.Class.Query().
				Where(class.Indx(cc.Indx)).
				OnlyIDX(ctx)

			ee, err := cp.client.EquipmentEntry.Create().
				SetEquipmentID(equipID).
				SetClassID(classID).
				SetQuantity(e.Quantity).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("error creating equipment entry: %w", err)
			}

			log.Info("Equipment entry created", "index", e.Equipment.Indx, "id", ee.ID, "quantity", e.Quantity)
		}
	}

	log.Info("Class starting equipment populated")
	return nil
}
