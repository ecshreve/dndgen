package popper_test

import (
	"testing"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/utils"
	"github.com/gkampitakis/go-snaps/snaps"
	"github.com/kr/pretty"
	"github.com/stretchr/testify/require"
)

type testRace struct {
	Indx string `json:"index"`
	Name string `json:"name"`
}

func TestRacePopulator(t *testing.T) {
	var data []*ent.Race

	err := utils.LoadJSONFile("testdata/simpleRace.json", &data)
	require.NoError(t, err)

	pretty.Print(data)

	snaps.MatchJSON(t, data)
}
