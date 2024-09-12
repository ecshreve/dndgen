package popper_test

import (
	"context"
	"encoding/json"
	"testing"

	"github.com/ecshreve/dndgen/ent/enttest"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
	"github.com/stretchr/testify/assert"

	_ "github.com/mattn/go-sqlite3"
)

// TestGetIDsFromIndxs tests the GetIDsFromIndxs method.
func TestGetIDsFromIndxs(t *testing.T) {

}

// TestPopulate tests the Populate methods.
func TestPopulate(t *testing.T) {
	t.Skip("skipping test")
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	testClient := enttest.Open(t, "sqlite3", "file:ent?mode=memory&_fk=1")
	defer testClient.Close()

	p := popper.NewPopper(ctx, testClient)
	err := p.PopulateAll(ctx)
	assert.NoError(t, err)

	testcases := []struct {
		desc  string
		query []interface{}
	}{
		{
			desc: "ability_scores",
			query: []interface{}{
				testClient.AbilityScore.Query().AllX(ctx),
			},
		},
		{
			desc: "skills",
			query: []interface{}{
				testClient.Skill.Query().WithAbilityScore().AllX(ctx),
			},
		},
		{
			desc: "languages",
			query: []interface{}{
				testClient.Language.Query().AllX(ctx),
			},
		},
		{
			desc: "classes",
			query: []interface{}{
				testClient.Class.Query().
					WithProficiencyChoices().
					WithClassEquipment().AllX(ctx),
			},
		},
		{
			desc: "races",
			query: []interface{}{
				testClient.Race.Query().
					WithSubrace().AllX(ctx),
			},
		},
		{
			desc: "proficiencies",
			query: []interface{}{
				testClient.Proficiency.Query().
					WithClasses().
					WithRaces().AllX(ctx),
			},
		},
		{
			desc: "equipment",
			query: []interface{}{
				testClient.Equipment.Query().WithEquipmentCategory().AllX(ctx),
			},
		},
		{
			desc: "weapons",
			query: []interface{}{
				testClient.Weapon.Query().AllX(ctx),
			},
		},
		{
			desc: "armor",
			query: []interface{}{
				testClient.Armor.Query().AllX(ctx),
			},
		},
	}

	for _, tc := range testcases {
		t.Run(tc.desc, func(t *testing.T) {
			snap.Snapshot(tc.desc, tc.query)
		})
	}

}

var classProficiencyJSON = `[{
	"index": "monk",
	"name": "Monk",
	"hit_die": 8,
	"proficiency_choices": [
		{
			"desc": "Choose two from Acrobatics, Athletics, History, Insight, Religion, and Stealth",
			"choose": 2,
			"type": "proficiencies",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "reference",
						"item": {
							"index": "skill-acrobatics",
							"name": "Skill: Acrobatics",
							"url": "/api/proficiencies/skill-acrobatics"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-athletics",
							"name": "Skill: Athletics",
							"url": "/api/proficiencies/skill-athletics"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-history",
							"name": "Skill: History",
							"url": "/api/proficiencies/skill-history"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-insight",
							"name": "Skill: Insight",
							"url": "/api/proficiencies/skill-insight"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-religion",
							"name": "Skill: Religion",
							"url": "/api/proficiencies/skill-religion"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-stealth",
							"name": "Skill: Stealth",
							"url": "/api/proficiencies/skill-stealth"
						}
					}
				]
			}
		},
		{
			"desc": "Choose one type of artisan’s tools or one musical instrument",
			"type": "proficiencies",
			"choose": 1,
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "choice",
						"choice": {
							"desc": "artisan's tools",
							"type": "proficiencies",
							"choose": 1,
							"from": {
								"option_set_type": "options_array",
								"options": [
									{
										"option_type": "reference",
										"item": {
											"index": "alchemists-supplies",
											"name": "Alchemist's Supplies",
											"url": "/api/proficiencies/alchemists-supplies"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "brewers-supplies",
											"name": "Brewer's Supplies",
											"url": "/api/proficiencies/brewers-supplies"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "calligraphers-supplies",
											"name": "Calligrapher's Supplies",
											"url": "/api/proficiencies/calligraphers-supplies"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "carpenters-tools",
											"name": "Carpenter's Tools",
											"url": "/api/proficiencies/carpenters-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "cartographers-tools",
											"name": "Cartographer's Tools",
											"url": "/api/proficiencies/cartographers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "cobblers-tools",
											"name": "Cobbler's Tools",
											"url": "/api/proficiencies/cobblers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "cooks-utensils",
											"name": "Cook's utensils",
											"url": "/api/proficiencies/cooks-utensils"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "glassblowers-tools",
											"name": "Glassblower's Tools",
											"url": "/api/proficiencies/glassblowers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "jewelers-tools",
											"name": "Jeweler's Tools",
											"url": "/api/proficiencies/jewelers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "leatherworkers-tools",
											"name": "Leatherworker's Tools",
											"url": "/api/proficiencies/leatherworkers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "masons-tools",
											"name": "Mason's Tools",
											"url": "/api/proficiencies/masons-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "painters-supplies",
											"name": "Painter's Supplies",
											"url": "/api/proficiencies/painters-supplies"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "potters-tools",
											"name": "Potter's Tools",
											"url": "/api/proficiencies/potters-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "smiths-tools",
											"name": "Smith's Tools",
											"url": "/api/proficiencies/smiths-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "tinkers-tools",
											"name": "Tinker's Tools",
											"url": "/api/proficiencies/tinkers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "weavers-tools",
											"name": "Weaver's Tools",
											"url": "/api/proficiencies/weavers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "woodcarvers-tools",
											"name": "Woodcarver's Tools",
											"url": "/api/proficiencies/woodcarvers-tools"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "disguise-kit",
											"name": "Disguise Kit",
											"url": "/api/proficiencies/disguise-kit"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "forgery-kit",
											"name": "Forgery Kit",
											"url": "/api/proficiencies/forgery-kit"
										}
									}
								]
							}
						}
					},
					{
						"option_type": "choice",
						"choice": {
							"desc": "musical instrument",
							"type": "proficiencies",
							"choose": 1,
							"from": {
								"option_set_type": "options_array",
								"options": [
									{
										"option_type": "reference",
										"item": {
											"index": "bagpipes",
											"name": "Bagpipes",
											"url": "/api/proficiencies/bagpipes"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "drum",
											"name": "Drum",
											"url": "/api/proficiencies/drum"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "dulcimer",
											"name": "Dulcimer",
											"url": "/api/proficiencies/dulcimer"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "flute",
											"name": "Flute",
											"url": "/api/proficiencies/flute"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "lute",
											"name": "Lute",
											"url": "/api/proficiencies/lute"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "lyre",
											"name": "Lyre",
											"url": "/api/proficiencies/lyre"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "horn",
											"name": "Horn",
											"url": "/api/proficiencies/horn"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "pan-flute",
											"name": "Pan flute",
											"url": "/api/proficiencies/pan-flute"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "shawm",
											"name": "Shawm",
											"url": "/api/proficiencies/shawm"
										}
									},
									{
										"option_type": "reference",
										"item": {
											"index": "viol",
											"name": "Viol",
											"url": "/api/proficiencies/viol"
										}
									}
								]
							}
						}
					}
				]
			}
		}
	],
	"proficiencies": [
		{
			"index": "simple-weapons",
			"name": "Simple Weapons",
			"url": "/api/proficiencies/simple-weapons"
		},
		{
			"index": "shortswords",
			"name": "Shortswords",
			"url": "/api/proficiencies/shortswords"
		},
		{
			"index": "saving-throw-dex",
			"name": "Saving Throw: DEX",
			"url": "/api/proficiencies/saving-throw-dex"
		},
		{
			"index": "saving-throw-str",
			"name": "Saving Throw: STR",
			"url": "/api/proficiencies/saving-throw-str"
		}
	],
	"saving_throws": [
		{
			"index": "str",
			"name": "STR",
			"url": "/api/ability-scores/str"
		},
		{
			"index": "dex",
			"name": "DEX",
			"url": "/api/ability-scores/dex"
		}
	],
	"starting_equipment": [
		{
			"equipment": {
				"index": "dart",
				"name": "Dart",
				"url": "/api/equipment/dart"
			},
			"quantity": 10
		}
	],
	"starting_equipment_options": [
		{
			"desc": "(a) a shortsword or (b) any simple weapon",
			"choose": 1,
			"type": "equipment",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "counted_reference",
						"count": 1,
						"of": {
							"index": "shortsword",
							"name": "Shortsword",
							"url": "/api/equipment/shortsword"
						}
					},
					{
						"option_type": "choice",
						"choice": {
							"desc": "any simple weapon",
							"choose": 1,
							"type": "equipment",
							"from": {
								"option_set_type": "equipment_category",
								"equipment_category": {
									"index": "simple-weapons",
									"name": "Simple Weapons",
									"url": "/api/equipment-categories/simple-weapons"
								}
							}
						}
					}
				]
			}
		},
		{
			"desc": "(a) a dungeoneer’s pack or (b) an explorer’s pack",
			"choose": 1,
			"type": "equipment",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "counted_reference",
						"count": 1,
						"of": {
							"index": "dungeoneers-pack",
							"name": "Dungeoneer's Pack",
							"url": "/api/equipment/dungeoneers-pack"
						}
					},
					{
						"option_type": "counted_reference",
						"count": 1,
						"of": {
							"index": "explorers-pack",
							"name": "Explorer's Pack",
							"url": "/api/equipment/explorers-pack"
						}
					}
				]
			}
		}
	],
	"class_levels": "/api/classes/monk/levels",
	"multi_classing": {
		"prerequisites": [
			{
				"ability_score": {
					"index": "dex",
					"name": "DEX",
					"url": "/api/ability-scores/dex"
				},
				"minimum_score": 13
			},
			{
				"ability_score": {
					"index": "wis",
					"name": "WIS",
					"url": "/api/ability-scores/wis"
				},
				"minimum_score": 13
			}
		],
		"proficiencies": [
			{
				"index": "simple-weapons",
				"name": "Simple Weapons",
				"url": "/api/proficiencies/simple-weapons"
			},
			{
				"index": "shortswords",
				"name": "Shortswords",
				"url": "/api/proficiencies/shortswords"
			}
		]
	},
	"subclasses": [
		{
			"index": "open-hand",
			"name": "Open Hand",
			"url": "/api/subclasses/open-hand"
		}
	],
	"url": "/api/classes/monk"
}]`

func TestParseClassProfs(t *testing.T) {
	t.Skip("skipping test")
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v []*popper.ClassChoiceWrapper
	if err := json.Unmarshal([]byte(classProficiencyJSON), &v); err != nil {
		t.Fatal(err)
	}

	ctx := context.Background()
	p := popper.NewTestPopper(ctx)
	p.PopulateCoin(ctx)
	p.PopulateAbilityScore(ctx)
	p.PopulateSkill(ctx)
	p.PopulateLanguage(ctx)
	p.PopulateDamageType(ctx)
	p.PopulateWeaponProperty(ctx)
	p.PopulateEquipment(ctx, "")
	p.PopulateRace(ctx)
	p.PopulateSubrace(ctx)
	p.PopulateClass(ctx)
	p.PopulateProficiency(ctx)
	p.PopulateProficiencyChoices(ctx)

	cts := p.Client.Class.Query().
		WithProficiencyChoices().AllX(ctx)
	snap.Snapshot("choices", cts)

	// profCh := v[0].ProficiencyChoices

	// for _, p := range profCh {
	// 	fromJson, err := json.Marshal(p.From)
	// 	require.NoError(t, err)

	// 	var nest popper.NestedChoice
	// 	err = json.Unmarshal(fromJson, &nest)
	// 	require.NoError(t, err)
	// 	pretty.Print(nest)

	// 	var single popper.RefChoice
	// 	err = json.Unmarshal(fromJson, &single)
	// 	require.NoError(t, err)
	// 	pretty.Print(single)
	// }

	// created, _ := p.PopulateClass(ctx)
	// snap.Snapshot("class created", created)

	// fetched := p.Client.Class.Query().WithProficiencyChoices().WithProficiencies().WithSavingThrows().WithStartingEquipment().WithStartingEquipmentOptions().WithSubclasses().AllX(ctx)
	// snap.Snapshot("class fetched", fetched)
}
