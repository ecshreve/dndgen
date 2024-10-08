package popper_test

import (
	"encoding/json"
	"testing"

	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
)

var skillJSON = `
{
	"index": "acrobatics",
	"name": "Acrobatics",
	"desc": [
		"Your Dexterity (Acrobatics) check covers your attempt to stay on your feet in a tricky situation, such as when you're trying to run across a sheet of ice, balance on a tightrope, or stay upright on a rocking ship's deck. The GM might also call for a Dexterity (Acrobatics) check to see if you can perform acrobatic stunts, including dives, rolls, somersaults, and flips."
	],
	"ability_score": {
		"index": "dex",
		"name": "DEX",
		"url": "/api/ability-scores/dex"
	},
	"url": "/api/skills/acrobatics"
}`

func TestParseSkill(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v = ent.Skill{}
	if err := json.Unmarshal([]byte(skillJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("skill", v)
}

var proficiencyJSON = `
[{
	"index": "saving-throw-int",
	"type": "Saving Throws",
	"name": "Saving Throw: INT",
	"classes": [
		{
			"index": "druid",
			"name": "Druid",
			"url": "/api/classes/druid"
		},
		{
			"index": "rogue",
			"name": "Rogue",
			"url": "/api/classes/rogue"
		},
		{
			"index": "wizard",
			"name": "Wizard",
			"url": "/api/classes/wizard"
		}
	],
	"races": [],
	"url": "/api/proficiencies/saving-throw-int",
	"reference": {
		"index": "int",
		"name": "INT",
		"url": "/api/ability-scores/int"
	}
},
{
	"index": "pan-flute",
	"type": "Musical Instruments",
	"name": "Pan flute",
	"classes": [],
	"races": [],
	"url": "/api/proficiencies/pan-flute",
	"reference": {
		"index": "pan-flute",
		"name": "Pan flute",
		"url": "/api/equipment/pan-flute"
	}
},
{
	"index": "skill-survival",
	"type": "Skills",
	"name": "Skill: Survival",
	"classes": [],
	"races": [],
	"url": "/api/proficiencies/skill-survival",
	"reference": {
		"index": "survival",
		"name": "Survival",
		"url": "/api/skills/survival"
	}
}]`

func TestParseProficiency(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v []popper.ProficiencyWrapper
	if err := json.Unmarshal([]byte(proficiencyJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("proficiency", v)
}

var classJSON = `
{
	"index": "barbarian",
	"name": "Barbarian",
	"hit_die": 12,
	"proficiency_choices": [
		{
			"desc": "Choose two from Animal Handling, Athletics, Intimidation, Nature, Perception, and Survival",
			"choose": 2,
			"type": "proficiencies",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "reference",
						"item": {
							"index": "skill-animal-handling",
							"name": "Skill: Animal Handling",
							"url": "/api/proficiencies/skill-animal-handling"
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
							"index": "skill-intimidation",
							"name": "Skill: Intimidation",
							"url": "/api/proficiencies/skill-intimidation"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-nature",
							"name": "Skill: Nature",
							"url": "/api/proficiencies/skill-nature"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-perception",
							"name": "Skill: Perception",
							"url": "/api/proficiencies/skill-perception"
						}
					},
					{
						"option_type": "reference",
						"item": {
							"index": "skill-survival",
							"name": "Skill: Survival",
							"url": "/api/proficiencies/skill-survival"
						}
					}
				]
			}
		}
	],
	"proficiencies": [
		{
			"index": "light-armor",
			"name": "Light Armor",
			"url": "/api/proficiencies/light-armor"
		},
		{
			"index": "medium-armor",
			"name": "Medium Armor",
			"url": "/api/proficiencies/medium-armor"
		},
		{
			"index": "shields",
			"name": "Shields",
			"url": "/api/proficiencies/shields"
		},
		{
			"index": "simple-weapons",
			"name": "Simple Weapons",
			"url": "/api/proficiencies/simple-weapons"
		},
		{
			"index": "martial-weapons",
			"name": "Martial Weapons",
			"url": "/api/proficiencies/martial-weapons"
		},
		{
			"index": "saving-throw-str",
			"name": "Saving Throw: STR",
			"url": "/api/proficiencies/saving-throw-str"
		},
		{
			"index": "saving-throw-con",
			"name": "Saving Throw: CON",
			"url": "/api/proficiencies/saving-throw-con"
		}
	],
	"saving_throws": [
		{
			"index": "str",
			"name": "STR",
			"url": "/api/ability-scores/str"
		},
		{
			"index": "con",
			"name": "CON",
			"url": "/api/ability-scores/con"
		}
	],
	"starting_equipment": [
		{
			"equipment": {
				"index": "explorers-pack",
				"name": "Explorer's Pack",
				"url": "/api/equipment/explorers-pack"
			},
			"quantity": 1
		},
		{
			"equipment": {
				"index": "javelin",
				"name": "Javelin",
				"url": "/api/equipment/javelin"
			},
			"quantity": 4
		}
	],
	"starting_equipment_options": [
		{
			"desc": "(a) a greataxe or (b) any martial melee weapon",
			"choose": 1,
			"type": "equipment",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "counted_reference",
						"count": 1,
						"of": {
							"index": "greataxe",
							"name": "Greataxe",
							"url": "/api/equipment/greataxe"
						}
					},
					{
						"option_type": "choice",
						"choice": {
							"desc": "any martial melee weapon",
							"choose": 1,
							"type": "equipment",
							"from": {
								"option_set_type": "equipment_category",
								"equipment_category": {
									"index": "martial-melee-weapons",
									"name": "Martial Melee Weapons",
									"url": "/api/equipment-categories/martial-melee-weapons"
								}
							}
						}
					}
				]
			}
		},
		{
			"desc": "(a) two handaxes or (b) any simple weapon",
			"choose": 1,
			"type": "equipment",
			"from": {
				"option_set_type": "options_array",
				"options": [
					{
						"option_type": "counted_reference",
						"count": 2,
						"of": {
							"index": "handaxe",
							"name": "Handaxe",
							"url": "/api/equipment/handaxe"
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
		}
	],
	"class_levels": "/api/classes/barbarian/levels",
	"multi_classing": {
		"prerequisites": [
			{
				"ability_score": {
					"index": "str",
					"name": "STR",
					"url": "/api/ability-scores/str"
				},
				"minimum_score": 13
			}
		],
		"proficiencies": [
			{
				"index": "shields",
				"name": "Shields",
				"url": "/api/proficiencies/shields"
			},
			{
				"index": "simple-weapons",
				"name": "Simple Weapons",
				"url": "/api/proficiencies/simple-weapons"
			},
			{
				"index": "martial-weapons",
				"name": "Martial Weapons",
				"url": "/api/proficiencies/martial-weapons"
			}
		]
	},
	"subclasses": [
		{
			"index": "berserker",
			"name": "Berserker",
			"url": "/api/subclasses/berserker"
		}
	],
	"url": "/api/classes/barbarian"
}`

func TestParseClass(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v struct {
		Indx              string `json:"index"`
		StartingEquipment []struct {
			Quantity  int `json:"quantity"`
			Equipment struct {
				Indx string `json:"index"`
			} `json:"equipment"`
		} `json:"starting_equipment,omitempty"`
		StartingEquipmentOptions []struct {
			Choose int    `json:"choose"`
			Desc   string `json:"desc"`
			From   struct {
				Options []struct {
					Count int `json:"count,omitempty"`
					Of    *struct {
						Indx string `json:"index,omitempty"`
						Url  string `json:"url,omitempty"`
					} `json:"of,omitempty"`
					Choice *struct {
						Choose int    `json:"choose,omitempty"`
						Desc   string `json:"desc,omitempty"`
						From   *struct {
							EquipmentCategory struct {
								Indx string `json:"index,omitempty"`
							} `json:"equipment_category,omitempty"`
						} `json:"from,omitempty"`
					} `json:"choice,omitempty"`
				} `json:"options,omitempty"`
			} `json:"from,omitempty"`
		} `json:"starting_equipment_options,omitempty"`
	}
	if err := json.Unmarshal([]byte(classJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("class", v)
}

var raceJSON = `
{
	"index": "dwarf",
	"name": "Dwarf",
	"speed": 25,
	"ability_bonuses": [
		{
			"ability_score": {
				"index": "con",
				"name": "CON",
				"url": "/api/ability-scores/con"
			},
			"bonus": 2
		}
	],
	"alignment": "Most dwarves are lawful, believing firmly in the benefits of a well-ordered society. They tend toward good as well, with a strong sense of fair play and a belief that everyone deserves to share in the benefits of a just order.",
	"age": "Dwarves mature at the same rate as humans, but they're considered young until they reach the age of 50. On average, they live about 350 years.",
	"size": "Medium",
	"size_description": "Dwarves stand between 4 and 5 feet tall and average about 150 pounds. Your size is Medium.",
	"proficiencies": [
		{
			"index": "battleaxes",
			"name": "Battleaxes",
			"url": "/api/proficiencies/battleaxes"
		},
		{
			"index": "handaxes",
			"name": "Handaxes",
			"url": "/api/proficiencies/handaxes"
		},
		{
			"index": "light-hammers",
			"name": "Light hammers",
			"url": "/api/proficiencies/light-hammers"
		},
		{
			"index": "warhammers",
			"name": "Warhammers",
			"url": "/api/proficiencies/warhammers"
		}
	],
	"starting_proficiency_options": {
		"desc": "You gain proficiency with the artisan’s tools of your choice: smith’s tools, brewer’s supplies, or mason’s tools.",
		"choose": 1,
		"type": "proficiencies",
		"from": {
			"option_set_type": "options_array",
			"options": [
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
						"index": "brewers-supplies",
						"name": "Brewer's Supplies",
						"url": "/api/proficiencies/brewers-supplies"
					}
				},
				{
					"option_type": "reference",
					"item": {
						"index": "masons-tools",
						"name": "Mason's Tools",
						"url": "/api/proficiencies/masons-tools"
					}
				}
			]
		}
	},
	"languages": [
		{
			"index": "common",
			"name": "Common",
			"url": "/api/languages/common"
		},
		{
			"index": "dwarvish",
			"name": "Dwarvish",
			"url": "/api/languages/dwarvish"
		}
	],
	"language_desc": "You can speak, read, and write Common and Dwarvish. Dwarvish is full of hard consonants and guttural sounds, and those characteristics spill over into whatever other language a dwarf might speak.",
	"traits": [
		{
			"index": "darkvision",
			"name": "Darkvision",
			"url": "/api/traits/darkvision"
		},
		{
			"index": "dwarven-resilience",
			"name": "Dwarven Resilience",
			"url": "/api/traits/dwarven-resilience"
		},
		{
			"index": "stonecunning",
			"name": "Stonecunning",
			"url": "/api/traits/stonecunning"
		},
		{
			"index": "dwarven-combat-training",
			"name": "Dwarven Combat Training",
			"url": "/api/traits/dwarven-combat-training"
		},
		{
			"index": "tool-proficiency",
			"name": "Tool Proficiency",
			"url": "/api/traits/tool-proficiency"
		}
	],
	"subraces": [
		{
			"index": "hill-dwarf",
			"name": "Hill Dwarf",
			"url": "/api/subraces/hill-dwarf"
		}
	],
	"url": "/api/races/dwarf"
}`

func TestParseRace(t *testing.T) {
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v ent.Race
	// var v struct {
	// 	Index          string `json:"index"`
	// 	Name           string `json:"name"`
	// 	Speed          int    `json:"speed"`
	// 	AbilityBonuses []struct {
	// 		AbilityScore struct {
	// 			Index string `json:"index"`
	// 			Name  string `json:"name"`
	// 		} `json:"ability_score"`
	// 		Bonus int `json:"bonus"`
	// 	} `json:"ability_bonuses"`
	// 	Alignment string `json:"alignment"`
	// 	Age       string `json:"age"`
	// }
	if err := json.Unmarshal([]byte(raceJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("race", v)

	var vv struct {
		Indx                       string `json:"index"`
		StartingProficiencyOptions struct {
			Choose int `json:"choose"`
			From   struct {
				Options []struct {
					Item struct {
						Indx string `json:"index"`
					} `json:"item"`
				} `json:"options"`
			} `json:"from"`
		} `json:"starting_proficiency_options,omitempty"`
	}

	if err := json.Unmarshal([]byte(raceJSON), &vv); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("vv", vv)
}
