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
	SavingThrows       []IndxWrapper      `json:"saving_throws"`
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
	if err := cp.populateSavingThrows(ctx); err != nil {
		return fmt.Errorf("error populating saving throws: %w", err)
	}

	if err := cp.populateProficiencyOptions(ctx); err != nil {
		return fmt.Errorf("error populating proficiency options: %w", err)
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

			se, err := cp.client.EquipmentEntry.Create().
				SetEquipmentID(equipID).
				AddClasIDs(classID).
				SetQuantity(e.Quantity).
				Save(ctx)
			if err != nil {
				return fmt.Errorf("error creating starting equipment entry: %w", err)
			}

			log.Info("Starting equipment entry created", "index", e.Equipment.Indx, "id", se.ID, "quantity", e.Quantity)
		}
	}

	log.Info("Class starting equipment populated")
	return nil
}

func (cp *ClassPopulator) populateSavingThrows(ctx context.Context) error {
	for _, cc := range cp.data {
		abilityScoreIDs := []int{}
		for _, s := range cc.SavingThrows {
			abilityScoreIDs = append(abilityScoreIDs, cp.indxToId[s.Indx])
		}
		classUpdated, err := cp.client.Class.UpdateOneID(cp.indxToId[cc.Indx]).
			AddSavingThrowIDs(abilityScoreIDs...).
			Save(ctx)
		if err != nil {
			return fmt.Errorf("error saving class: %w", err)
		}
		log.Info("Class updated", "index", cc.Indx, "name", classUpdated.Name, "saving throws", abilityScoreIDs)
	}

	log.Info("Class saving throws populated")
	return nil
}

func (cp *ClassPopulator) populateProficiencyOptions(ctx context.Context) error {
	for _, cc := range cp.data {
		classUpdate := cp.client.Class.UpdateOneID(cp.indxToId[cc.Indx])
		for _, pc := range cc.ProficiencyChoices {
			if pc.From.OptionSetType != "options_array" {
				log.Warn("Skipping proficiency choice", "class", cc.Indx, "choice", pc.From.OptionSetType)
				continue
			}
			poIDs := []int{}
			for _, po := range pc.From.Options {
				poIDs = append(poIDs, cp.indxToId[po.Item.Indx])
			}
			classUpdate = classUpdate.AddProficiencyOptions(
				cp.client.ProficiencyChoice.Create().
					SetDesc(pc.Desc).
					SetChoose(pc.Choose).
					AddProficiencyIDs(poIDs...).
					SaveX(ctx),
			)
			log.Info("Added starting proficiency options", "class", cc.Indx, "choose", pc.Choose, "from", len(pc.From.Options))
		}
		if err := classUpdate.Exec(ctx); err != nil {
			return fmt.Errorf("error saving class: %w", err)
		}
		log.Info("Class updated", "index", cc.Indx, "name", cc.Name, "proficiency choices", len(cc.ProficiencyChoices))
	}

	log.Info("Class proficiency options populated")
	return nil
}
