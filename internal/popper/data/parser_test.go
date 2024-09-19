package popper_test

import (
	"encoding/json"
	"testing"

	"github.com/ecshreve/dndgen/ent"
	"github.com/stretchr/testify/assert"
)

var coinJSON = `
{
	"index": "cp",
	"name": "Copper Piece",
	"desc": "Copper coins are small and roughly the size of a fingernail. A single copper coin is worth 1 cp.",
	"gold_conversion_rate": 0.01
}
`

// TestParseCoin tests the parsing of a coin.
func TestParseCoin(t *testing.T) {
	var v ent.Coin
	err := json.Unmarshal([]byte(coinJSON), &v)
	assert.NoError(t, err)

	assert.Equal(t, "cp", v.Indx)
	assert.Equal(t, "Copper Piece", v.Name)
	assert.Equal(t, "Copper coins are small and roughly the size of a fingernail. A single copper coin is worth 1 cp.", v.Desc)
	assert.Equal(t, 0.01, v.GoldConversionRate)
}
