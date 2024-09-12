package popper

import (
	"context"
	"encoding/json"
	"fmt"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/charmbracelet/log"

	"github.com/ecshreve/dndgen/ent"

	_ "github.com/mattn/go-sqlite3"
)

//go:generate go run gen.go

// Popper is a populator for the ent database.
type Popper struct {
	Client   *ent.Client
	Reader   *ent.Client
	IdToIndx map[int]string
	IndxToId map[string]int
	Context  *context.Context
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
		Context:  &ctx,
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

// PopulateAll calls entity specific populators in the order they should be populated.
func (p *Popper) PopulateAll(ctx context.Context) error {
	_, err := p.PopulateCoin(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Coin entities: %w", err)
	}
	_, err = p.PopulateAbilityScore(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate AbilityScore entities: %w", err)
	}

	_, err = p.PopulateSkill(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Skill entities: %w", err)
	}

	_, err = p.PopulateLanguage(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Language entities: %w", err)
	}

	_, err = p.PopulateDamageType(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate DamageType entities: %w", err)
	}

	_, err = p.PopulateWeaponProperty(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate WeaponProperty entities: %w", err)
	}

	_, err = p.PopulateEquipment(ctx, "")
	if err != nil {
		return fmt.Errorf("unable to populate Equipment entities: %w", err)
	}

	_, err = p.PopulateClass(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Class entities: %w", err)
	}

	_, err = p.PopulateRace(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Race entities: %w", err)
	}

	_, err = p.PopulateSubrace(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Subrace entities: %w", err)
	}

	_, err = p.PopulateProficiency(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Proficiency entities: %w", err)
	}

	_, err = p.PopulateMagicSchool(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate MagicSchool entities: %w", err)
	}

	_, err = p.PopulateRuleSection(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate RuleSection entities: %w", err)
	}

	_, err = p.PopulateRule(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Rule entities: %w", err)
	}

	_, err = p.PopulateTrait(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate Trait entities: %w", err)
	}

	err = p.PopulateStartingProficiencyOptions(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate StartingProficiencyOptions entities: %w", err)
	}

	_, err = p.PopulateProficiencyChoices(ctx)
	if err != nil {
		return fmt.Errorf("unable to populate ProficiencyProficiencyChoices entities: %w", err)
	}

	return nil
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
