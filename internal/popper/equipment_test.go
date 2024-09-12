package popper_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/charmbracelet/log"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

var equipmentSingleJSON = `
[
	{
		"index": "club",
		"name": "Club",
	"equipment_category": {
		"index": "weapon",
		"name": "Weapon",
		"url": "/api/equipment-categories/weapon"
	},
	"weapon_category": "Simple",
	"weapon_range": "Melee",
	"category_range": "Simple Melee",
	"cost": {
		"quantity": 1,
		"unit": "sp"
	},
	"damage": {
		"damage_dice": "1d4",
		"damage_type": {
			"index": "bludgeoning",
			"name": "Bludgeoning",
			"url": "/api/damage-types/bludgeoning"
		}
	},
	"range": {
		"normal": 5
	},
	"weight": 2,
	"properties": [
		{
			"index": "weap-light",
			"name": "Light",
			"url": "/api/weapon-properties/light"
		},
		{
			"index": "weap-monk",
			"name": "Monk",
			"url": "/api/weapon-properties/monk"
		}
	],
	"url": "/api/equipment/club"
}]
`

var equipmentManyJSON = `
[
	{
		"index": "club",
		"name": "Club",
    "equipment_category": {
      "index": "weapon",
      "name": "Weapon",
      "url": "/api/equipment-categories/weapon"
    },
    "weapon_category": "Simple",
    "weapon_range": "Melee",
    "category_range": "Simple Melee",
    "cost": {
      "quantity": 1,
      "unit": "sp"
    },
    "damage": {
      "damage_dice": "1d4",
      "damage_type": {
        "index": "bludgeoning",
        "name": "Bludgeoning",
        "url": "/api/damage-types/bludgeoning"
      }
    },
    "range": {
      "normal": 5
    },
    "weight": 2,
    "properties": [
      {
        "index": "weap-light",
        "name": "Light",
        "url": "/api/weapon-properties/light"
      },
      {
        "index": "weap-monk",
        "name": "Monk",
        "url": "/api/weapon-properties/monk"
      }
    ],
    "url": "/api/equipment/club"
  },
  {
    "index": "padded-armor",
    "name": "Padded Armor",
    "equipment_category": {
      "index": "armor",
      "name": "Armor",
      "url": "/api/equipment-categories/armor"
    },
    "armor_category": "Light",
    "armor_class": {
      "base": 11,
      "dex_bonus": true
    },
    "str_minimum": 0,
    "stealth_disadvantage": true,
    "weight": 8,
    "cost": {
      "quantity": 5,
      "unit": "gp"
    },
    "url": "/api/equipment/padded-armor"
  },
  {
    "index": "scholars-pack",
    "name": "Scholar's Pack",
    "equipment_category": {
      "index": "adventuring-gear",
      "name": "Adventuring Gear",
      "url": "/api/equipment-categories/adventuring-gear"
    },
    "gear_category": {
      "index": "equipment-packs",
      "name": "Equipment Packs",
      "url": "/api/equipment-categories/equipment-packs"
    },
    "cost": {
      "quantity": 40,
      "unit": "gp"
    },
    "contents": [
      {
        "item": {
          "index": "backpack",
          "name": "Backpack",
          "url": "/api/equipment/backpack"
        },
        "quantity": 1
      },
      {
        "item": {
          "index": "book",
          "name": "Book",
          "url": "/api/equipment/book"
        },
        "quantity": 1
      },
      {
        "item": {
          "index": "ink-1-ounce-bottle",
          "name": "Ink (1 ounce bottle)",
          "url": "/api/equipment/ink-1-ounce-bottle"
        },
        "quantity": 1
      },
      {
        "item": {
          "index": "ink-pen",
          "name": "Ink pen",
          "url": "/api/equipment/ink-pen"
        },
        "quantity": 1
      },
      {
        "item": {
          "index": "parchment-one-sheet",
          "name": "Parchment (one sheet)",
          "url": "/api/equipment/parchment-one-sheet"
        },
        "quantity": 10
      },
      {
        "item": {
          "index": "little-bag-of-sand",
          "name": "Little bag of sand",
          "url": "/api/equipment/little-bag-of-sand"
        },
        "quantity": 1
      },
      {
        "item": {
          "index": "small-knife",
          "name": "Small knife",
          "url": "/api/equipment/small-knife"
        },
        "quantity": 1
      }
    ],
    "url": "/api/equipment/scholars-pack"
  },
  {
    "index": "alchemists-supplies",
    "name": "Alchemist's Supplies",
    "equipment_category": {
      "index": "tools",
      "name": "Tools",
      "url": "/api/equipment-categories/tools"
    },
    "tool_category": "Artisan's Tools",
    "cost": {
      "quantity": 50,
      "unit": "gp"
    },
    "weight": 8,
    "desc": [
      "These special tools include the items needed to pursue a craft or trade. The table shows examples of the most common types of tools, each providing items related to a single craft. Proficiency with a set of artisan's tools lets you add your proficiency bonus to any ability checks you make using the tools in your craft. Each type of artisan's tools requires a separate proficiency."
    ],
    "url": "/api/equipment/alchemists-supplies"
  },
  {
    "index": "elephant",
    "name": "Elephant",
    "equipment_category": {
      "index": "mounts-and-vehicles",
      "name": "Mounts and Vehicles",
      "url": "/api/equipment-categories/mounts-and-vehicles"
    },
    "vehicle_category": "Mounts and Other Animals",
    "cost": {
      "quantity": 200,
      "unit": "gp"
    },
    "speed": {
      "quantity": 40,
      "unit": "ft/round"
    },
    "capacity": "1,320 lb.",
    "url": "/api/equipment/elephant"
  }]
`

func TestParseEquipment(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v []popper.EquipmentWrapper
	if err := json.Unmarshal([]byte(equipmentSingleJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("equipment", v)
}

func TestPopulateEquipment(t *testing.T) {
	ctx := context.Background()
	log.SetLevel(log.DebugLevel)
	log.SetReportCaller(true)

	snap := snapshotter.New(t)
	defer snap.Verify()

	p := dbHelperEquipment(t, ctx)

	allEquipmentIds, err := p.PopulateEquipment(ctx, equipmentManyJSON)
	assert.NoError(t, err)

	equipment := p.Client.Equipment.Query().
		Order(ent.Asc(equipment.FieldIndx)).
		AllX(ctx)
	assert.Equal(t, len(allEquipmentIds), len(equipment))

	snap.Snapshot("equipment", sanitizeIds(equipment))
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
