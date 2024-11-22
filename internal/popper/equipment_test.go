package popper_test

import (
	"testing"

	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/ecshreve/dndgen/internal/utils"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/stretchr/testify/require"
)

func TestEquipmentPopulator(t *testing.T) {
	var data []popper.EquipmentJSON

	err := utils.LoadJSONFile("data/Equipment.json", &data)
	// err := utils.LoadJSONFile("testdata/multiEquipment.json", &data)
	require.NoError(t, err)

	snaps.MatchJSON(t, data)
}
