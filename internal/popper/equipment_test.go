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
)

func TestEquipmentPopulator(t *testing.T) {
	var data []popper.EquipmentJSON

	err := utils.LoadJSONFile("testdata/singleEquipment.json", &data)
	require.NoError(t, err)

	snaps.MatchJSON(t, data)

	var multiData []popper.EquipmentJSON
	err = utils.LoadJSONFile("testdata/multiEquipment.json", &multiData)
	require.NoError(t, err)

	snaps.MatchJSON(t, multiData)
}

func TestEquipmentPopulatorFull(t *testing.T) {
	ctx := context.Background()
	cl, err := ent.Open("sqlite3", "file:ent?mode=memory&_fk=1")
	require.NoError(t, err)
	defer cl.Close()

	err = cl.Schema.Create(ctx, schema.WithGlobalUniqueID(true))
	require.NoError(t, err)

	pp := popper.NewPopper(ctx, cl, "../../data")
	err = pp.PopulateAll(ctx)
	require.NoError(t, err)

	equipPop := popper.NewEquipmentPopulator(pp)
	err = equipPop.Populate(ctx)
	require.NoError(t, err)

	allEquipment, err := cl.Equipment.Query().
		WithCost().
		All(ctx)
	require.NoError(t, err)

	snaps.MatchJSON(t, allEquipment)
}
