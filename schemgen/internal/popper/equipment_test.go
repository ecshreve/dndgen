package popper_test

import (
	"context"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/ecshreve/dndgen/internal/utils"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestParseEquipment(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v []popper.EquipmentWrapper
	err := utils.LoadJSONFile("testdata/singleEquipment.json", &v)
	require.NoError(t, err)

	snap.Snapshot("equipment", v)
}

func TestPopulateEquipment(t *testing.T) {
	ctx := context.Background()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	snap := snapshotter.New(t)
	defer snap.Verify()

	p := dbHelperEquipment(t, ctx)

	allEquipmentIds, err := p.PopulateEquipmentFromFile(ctx, "testdata/multiEquipment.json")
	require.NoError(t, err)

	equipment := p.Client.Equipment.Query().
		Order(ent.Asc(equipment.FieldIndx)).
		AllX(ctx)
	assert.Equal(t, len(allEquipmentIds), len(equipment))

	snap.Snapshot("equipment", sanitizeIds(equipment))
}

// TestPopulateEquipmentDuplicate is a test that populates the equipment and checks for duplicate constraints.
func TestPopulateEquipmentDuplicate(t *testing.T) {
	ctx := context.Background()
	p := dbHelperEquipment(t, ctx)

	_, err := p.PopulateEquipmentFromFile(ctx, "testdata/multiEquipment.json")
	require.NoError(t, err)

	dupe, err := p.PopulateEquipmentFromFile(ctx, "testdata/singleEquipment.json")
	assert.Error(t, err)
	assert.Equal(t, len(dupe), 0)
}

func dbHelperEquipment(t *testing.T, ctx context.Context) *popper.Popper {
	p := popper.NewTestPopper(ctx)

	_, err := p.PopulateCoin(ctx)
	require.NoError(t, err)

	_, err = p.PopulateAbilityScore(ctx)
	require.NoError(t, err)

	_, err = p.PopulateSkill(ctx)
	require.NoError(t, err)

	_, err = p.PopulateLanguage(ctx)
	require.NoError(t, err)

	_, err = p.PopulateDamageType(ctx)
	require.NoError(t, err)

	_, err = p.PopulateWeaponProperty(ctx)
	require.NoError(t, err)

	return p
}

// sanitizeIds sanitizes the IDs in the equipment so that the snapshot tests pass.
func sanitizeIds(equipment []*ent.Equipment) []*ent.Equipment {
	for _, e := range equipment {
		e.ID = 0
		e.EquipmentCategoryID = 0
	}
	return equipment
}
