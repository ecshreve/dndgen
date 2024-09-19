package popper

import (
	"context"
	"encoding/json"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"

	_ "github.com/mattn/go-sqlite3"
)

// Popper is a populator for the ent database.
type Popper struct {
	Client   *ent.Client
	Reader   *ent.Client
	IdToIndx map[int]string
	IndxToId map[string]int
}

// NewPopper creates a new Popper configured to populate data from JSON files in
// the popper/data directory into a database via the given ent client.
func NewPopper(ctx context.Context, client *ent.Client) *Popper {
	if client == nil {
		client, _ = ent.Open(dialect.SQLite, "file:dev.db?_fk=1")
	}

	return &Popper{
		Client:   client,
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
		Reader:   client,
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

// PopulateAll populates the database with all the data from the JSON files in
// the popper/data directory.
func (p *Popper) PopulateAll(ctx context.Context) error {
	log.Debug("PopulateAll")
	_, err := p.PopulateRuleSection(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateRule(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateAbilityScore(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateSkill(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateLanguage(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateAlignment(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateDamageType(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateMagicSchool(ctx)
	if err != nil {
		return err
	}

	_, err = p.PopulateRace(ctx)
	if err != nil {
		return err
	}

	return nil
}
