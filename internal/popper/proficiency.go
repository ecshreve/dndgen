package popper

import (
	"context"
	"fmt"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
)

type RefWrapper struct {
	Indx string `json:"index"`
	Name string `json:"name"`
	URL  string `json:"url"`
}

type ProficiencyJSON struct {
	Indx      string        `json:"index"`
	ProfType  string        `json:"type"`
	Name      string        `json:"name"`
	Classes   []IndxWrapper `json:"classes"`
	Races     []IndxWrapper `json:"races"`
	Reference RefWrapper    `json:"reference"`
}

type ProficiencyPopulator struct {
	client   *ent.Client
	dataFile string
	data     []ProficiencyJSON
	indxToId map[string]int
}

func NewProficiencyPopulator(pp *Popper) *ProficiencyPopulator {
	ep := &ProficiencyPopulator{
		client:   pp.Client,
		dataFile: fmt.Sprintf("%s/Proficiency.json", pp.DataDir),
		indxToId: pp.IndxToId,
	}

	if err := utils.LoadJSONFile(ep.dataFile, &ep.data); err != nil {
		log.Fatal("Error loading equipment data", "error", err)
	}

	return ep
}

func (cp *ProficiencyPopulator) Populate(ctx context.Context) error {
	if len(cp.data) == 0 {
		return fmt.Errorf("no Proficiency data to populate")
	}

	// for _, prof := range cp.data {
	// 	pt := strings.Split(strings.TrimPrefix(prof.Reference.URL, "/api/"), "/")[0]
	// 	pt = strings.Replace(pt, "ability-scores", "SAVING_THROW", 1)
	// 	pt = strings.Replace(pt, "skills", "SKILL", 1)
	// 	pt = strings.Replace(pt, "equipment", "EQUIPMENT", 1)

	// 	profCreate := cp.client.Proficiency.Create().
	// 		SetIndx(prof.Indx).
	// 		SetName(prof.Name).
	// 		SetCategory(pt)

	// 	switch pt {
	// 	case "SAVING_THROW":
	// 		profCreate.SetSavingThrowID(cp.indxToId[prof.Reference.Indx])
	// 	case "SKILL":
	// 		profCreate.SetSkillID(cp.indxToId[prof.Reference.Indx])
	// 	case "EQUIPMENT":
	// 		profCreate.SetEquipmentID(cp.indxToId[prof.Reference.Indx])
	// 	default:
	// 		return fmt.Errorf("unknown proficiency type: %s", pt)
	// 	}

	// 	clIDs := make([]int, 0)
	// 	for _, class := range prof.Classes {
	// 		clIDs = append(clIDs, cp.indxToId[class.Indx])
	// 	}
	// 	profCreate.AddClassIDs(clIDs...)

	// 	raIDs := make([]int, 0)
	// 	skipRace := false
	// 	for _, race := range prof.Races {
	// 		if _, ok := cp.indxToId[race.Indx]; !ok {
	// 			log.Warn("unknown race", "race", race.Indx)
	// 			skipRace = true
	// 			break
	// 		}
	// 		raIDs = append(raIDs, cp.indxToId[race.Indx])
	// 	}
	// 	if !skipRace {
	// 		profCreate.AddRaceIDs(raIDs...)
	// 	}

	// 	profEntity, err := profCreate.Save(ctx)
	// 	if err != nil {
	// 		log.Warn("error creating proficiency", "error", err)
	// 	}

	// 	log.Info("Populated proficiency", "indx", profEntity.Indx, "category", pt, "id", profEntity.ID)
	// }
	return nil
}
