package popper_test

import (
	"context"
	"testing"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func SetupDB(ctx context.Context) (*popper.Popper, error) {
	client, err := ent.Open("sqlite3", "file:testdb?mode=memory&_fk=1")
	if err != nil {
		return nil, err
	}
	if err := client.Schema.Create(ctx); err != nil {
		return nil, err
	}

	return popper.NewPopper(ctx, client), nil
}

func TestPopulate(t *testing.T) {
	ctx := context.Background()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	assert.NoError(t, err)
}

// TestPopulateAbilityScore tests the Populate methods.
func TestPopulateAbilityScore(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	as, err := p.Client.AbilityScore.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("ability_scores", as)
}

// TestPopulateSkill tests the Populate methods.
func TestPopulateSkill(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	sk, err := p.Client.Skill.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("skills", sk)
}

// TestPopulateLanguage tests the Populate methods.
func TestPopulateLanguage(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	ll, err := p.Client.Language.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("languages", ll)
}

// TestPopulateClass tests the Populate methods.
func TestPopulateClass(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	cl, err := p.Client.Class.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("classes", cl)
}

// TestPopulateRace tests the Populate methods.
func TestPopulateRace(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	cl, err := p.Client.Race.Query().
		WithStartingProficiencyOption().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("races", cl)
}

// TestPopulateEquipment tests the Populate methods.
func TestPopulateEquipment(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	// err := p.PopulateAll(ctx)
	// require.NoError(t, err)

	_, err := p.PopulateWeaponProperty(ctx)
	require.NoError(t, err)

	err = p.PopulateEquipment(ctx)
	require.NoError(t, err)

	cl, err := p.Client.Equipment.Query().WithCost().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("equipment", cl)

	w, err := p.Client.Weapon.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("weapons", w)

	a, err := p.Client.Armor.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("armor", a)

	g, err := p.Client.Gear.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("gear", g)

	v, err := p.Client.Vehicle.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("vehicles", v)

	tl, err := p.Client.Tool.Query().All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("tools", tl)
}

// TestPopulateProficiency tests the Populate methods.
func TestPopulateProficiency(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	p := popper.NewTestPopper(ctx)
	err := p.PopulateAll(ctx)
	require.NoError(t, err)

	prof, err := p.Client.Proficiency.Query().All(ctx)
	// WithClasses(func(q *ent.ClassQuery) {
	// 	q.Select(class.FieldIndx).Order(ent.Asc(class.FieldID))
	// }).
	// WithRaces(func(rq *ent.RaceQuery) {
	// 	rq.Select(race.FieldIndx).Order(ent.Asc(race.FieldID))
	// }).Order(ent.Asc(proficiency.FieldID)).All(ctx)
	assert.NoError(t, err)
	snap.Snapshot("proficiencies", prof)
}
