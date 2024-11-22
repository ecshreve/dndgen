package popper_test

import (
	"context"
	"testing"

	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/ecshreve/dndgen/internal/utils"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"

	_ "github.com/mattn/go-sqlite3"
)

type testRace struct {
	Indx string `json:"index"`
	Name string `json:"name"`
}

func TestRacePopulator(t *testing.T) {
	var data []popper.RaceJSON

	err := utils.LoadJSONFile("testdata/simpleRace.json", &data)
	require.NoError(t, err)

	snaps.MatchJSON(t, data)
}

func TestRacePopulatorFull(t *testing.T) {
	ctx := context.Background()
	// cl, err := ent.Open("sqlite3", "file:ent?cache=shared&mode=memory&_fk=1")
	cl, err := ent.Open("sqlite3", "file:testdb.db?_fk=1")
	require.NoError(t, err)

	err = cl.Schema.Create(ctx, schema.WithGlobalUniqueID(true))
	require.NoError(t, err)

	pp := popper.NewPopper(ctx, cl, "data")
	err = pp.PopulateAll(ctx)
	require.NoError(t, err)

	profPop := popper.NewProficiencyPopulator(pp, "data/Proficiency.json")
	err = profPop.Populate(ctx)
	require.NoError(t, err)

	racePop := popper.NewRacePopulator(pp, "testdata/simpleRace.json")
	err = racePop.Populate(ctx)
	require.NoError(t, err)

	allRaces, err := cl.Race.Query().
		WithStartingProficiencies().
		WithStartingProficiencyOptions(
			func(pc *ent.ProficiencyChoiceQuery) {
				pc.WithProficiencies()
			},
		).
		All(ctx)
	require.NoError(t, err)

	snaps.MatchJSON(t, allRaces)
}
