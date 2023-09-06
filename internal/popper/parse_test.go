package popper_test

import (
	"context"
	"encoding/json"
	"testing"

	"entgo.io/ent/dialect"
	"entgo.io/ent/dialect/sql/schema"
	"github.com/ecshreve/dndgen/ent"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/weapon"
	"github.com/ecshreve/dndgen/internal/popper"
	"github.com/samsarahq/go/snapshotter"
	log "github.com/sirupsen/logrus"
	"github.com/stretchr/testify/assert"
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

	ctx := context.Background()

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

	p := popper.NewTestPopper(ctx)
	p.PopulateAll(ctx)
	dd := p.Client.Class.Query().Where(class.Indx(v.Indx)).
		WithEquipmentChoices(func(ecq *ent.EquipmentChoiceQuery) {
			ecq.WithEquipment()
		}).
		WithClassEquipment().AllX(ctx)
	snap.Snapshot("class from db", dd)
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

var equipmentJSON = `
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
			"index": "light",
			"name": "Light",
			"url": "/api/weapon-properties/light"
		},
		{
			"index": "monk",
			"name": "Monk",
			"url": "/api/weapon-properties/monk"
		}
	],
	"url": "/api/equipment/club"
}`

func TestParseEquipment(t *testing.T) {
	ctx := context.Background()
	snap := snapshotter.New(t)
	defer snap.Verify()

	var v popper.EquipmentWrapper
	if err := json.Unmarshal([]byte(equipmentJSON), &v); err != nil {
		t.Fatal(err)
	}
	snap.Snapshot("equipment", v)

	client, err := ent.Open(dialect.SQLite, "file:dnd?mode=memory&_fk=1")
	if err != nil {
		log.Fatal(err)
	}

	if err := client.Schema.Create(context.Background(), schema.WithGlobalUniqueID(true)); err != nil {
		log.Fatal(err)
	}

	p := popper.NewPopper(ctx, client)
	if err := p.PopulateAll(ctx); err != nil {
		log.Fatal(err)
	}

	eq := p.Client.Equipment.Query().
		Where(equipment.Indx("club")).
		WithCost().OnlyX(ctx)

	snap.Snapshot("equipment", eq)

	all, err := p.Client.Weapon.Query().
		Where(weapon.Indx("club")).
		WithWeaponDamage().
		WithWeaponProperties().All(ctx)

	assert.NoError(t, err)
	assert.NotNil(t, eq)
	// pretty.Print(eq.Weapon(ctx))
	// p.Client.Equipment.Create().SetEquipment(&v).SaveX(context.Background())

	// all := p.Client.Equipment.Query().AllX(ctx)
	// snap.Snapshot("equipment", v)
	snap.Snapshot("all equipment", all)
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

	ctx := context.Background()
	p := popper.NewTestPopper(ctx)
	p.PopulateAll(ctx)

	// created, _ := p.PopulateProficiency(ctx)
	// snap.Snapshot("proficiency created", created)

	fetched := p.Client.Proficiency.Query().WithEquipment().WithSavingThrow().WithSkill().WithClasses().AllX(ctx)
	snap.Snapshot("proficiency fetched", fetched)

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
	p.PopulateEquipment(ctx)
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
