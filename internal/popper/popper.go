package popper

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"

	_ "github.com/mattn/go-sqlite3"
)

// Popper is a populator for the ent database.
type Popper struct {
	Client   *ent.Client
	DataDir  string
	IdToIndx map[int]string
	IndxToId map[string]int
}

// NewPopper creates a new Popper configured to populate data from JSON files in
// the popper/data directory into a database via the given ent client.
func NewPopper(ctx context.Context, client *ent.Client, dataDir string) *Popper {
	if client == nil {
		log.Fatal("client is nil")
	}

	if dataDir == "" {
		dataDir = "data"
	}

	return &Popper{
		Client:   client,
		DataDir:  dataDir,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
	}
}

// NewTestPopper creates a new Popper for testing.
func NewTestPopper(ctx context.Context) *Popper {
	// client, err := ent.Open(dialect.SQLite, "file:devtest.db?_fk=1")
	client, err := ent.Open("sqlite3", "file:testdb?mode=memory&_fk=1")
	if err != nil {
		log.Error(err)
		return nil
	}
	if err := client.Schema.Create(ctx, schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	return &Popper{
		Client:   client,
		IdToIndx: map[int]string{},
		IndxToId: map[string]int{},
	}
}

type IndxWrapper struct {
	Indx string `json:"index"`
}

// GetIDsFromIndxWrapperString gets the IDs from the given JSON string.
func (p *Popper) GetIDsFromIndxWrapperString(v []byte) []int {
	var indxs []IndxWrapper
	if err := json.Unmarshal(v, &indxs); err != nil {
		log.Fatal(err)
	}

	return p.GetIDsFromIndxWrappers(indxs)
}

// GetIDsFromIndxWrappers gets the IDs from the given indx wrappers.
func (p *Popper) GetIDsFromIndxWrappers(indxs []IndxWrapper) []int {
	ids := make([]int, len(indxs))
	for i, indx := range indxs {
		ids[i] = p.IndxToId[indx.Indx]
	}

	return ids
}

// PopulateCustom populates custom entities.
func (p *Popper) PopulateCustom(ctx context.Context) error {
	if err := p.PopulateRuleEdges(ctx); err != nil {
		return fmt.Errorf("error populating rule edges: %w", err)
	}
	if err := p.PopulateSkillEdges(ctx); err != nil {
		return fmt.Errorf("error populating skill edges: %w", err)
	}

	equipmentPopulator := NewEquipmentPopulator(p)
	if err := equipmentPopulator.Populate(ctx); err != nil {
		return fmt.Errorf("error populating equipment: %w", err)
	}

	proficiencyPopulator := NewProficiencyPopulator(p, "data/Proficiency.json")
	if err := proficiencyPopulator.Populate(ctx); err != nil {
		return fmt.Errorf("error populating proficiencies: %w", err)
	}

	racePopulator := NewRacePopulator(p, "data/Race.json")
	if err := racePopulator.Populate(ctx); err != nil {
		return fmt.Errorf("error populating races: %w", err)
	}

	subracePopulator := NewRacePopulator(p, "data/Subrace.json")
	if err := subracePopulator.PopulateSubraces(ctx); err != nil {
		return fmt.Errorf("error populating subraces: %w", err)
	}

	classPopulator := NewClassPopulator(p)
	if err := classPopulator.Populate(ctx); err != nil {
		return fmt.Errorf("error populating classes: %w", err)
	}

	if err := p.PopulateFeatureEdges(ctx); err != nil {
		return fmt.Errorf("error populating feature edges: %w", err)
	}

	return nil
}
