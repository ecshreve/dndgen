// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
	// AbilityBonusColumns holds the columns for the "ability_bonus" table.
	AbilityBonusColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "bonus", Type: field.TypeInt},
		{Name: "ability_bonus_ability_score", Type: field.TypeInt},
	}
	// AbilityBonusTable holds the schema information for the "ability_bonus" table.
	AbilityBonusTable = &schema.Table{
		Name:       "ability_bonus",
		Columns:    AbilityBonusColumns,
		PrimaryKey: []*schema.Column{AbilityBonusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ability_bonus_ability_scores_ability_score",
				Columns:    []*schema.Column{AbilityBonusColumns[2]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// AbilityBonusChoicesColumns holds the columns for the "ability_bonus_choices" table.
	AbilityBonusChoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "choose", Type: field.TypeInt},
	}
	// AbilityBonusChoicesTable holds the schema information for the "ability_bonus_choices" table.
	AbilityBonusChoicesTable = &schema.Table{
		Name:       "ability_bonus_choices",
		Columns:    AbilityBonusChoicesColumns,
		PrimaryKey: []*schema.Column{AbilityBonusChoicesColumns[0]},
	}
	// AbilityScoresColumns holds the columns for the "ability_scores" table.
	AbilityScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "full_name", Type: field.TypeString},
	}
	// AbilityScoresTable holds the schema information for the "ability_scores" table.
	AbilityScoresTable = &schema.Table{
		Name:       "ability_scores",
		Columns:    AbilityScoresColumns,
		PrimaryKey: []*schema.Column{AbilityScoresColumns[0]},
	}
	// AlignmentsColumns holds the columns for the "alignments" table.
	AlignmentsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "abbr", Type: field.TypeString},
	}
	// AlignmentsTable holds the schema information for the "alignments" table.
	AlignmentsTable = &schema.Table{
		Name:       "alignments",
		Columns:    AlignmentsColumns,
		PrimaryKey: []*schema.Column{AlignmentsColumns[0]},
	}
	// ArmorsColumns holds the columns for the "armors" table.
	ArmorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "armor_category", Type: field.TypeEnum, Enums: []string{"light", "medium", "heavy", "shield"}},
		{Name: "str_minimum", Type: field.TypeInt},
		{Name: "stealth_disadvantage", Type: field.TypeBool},
		{Name: "ac_base", Type: field.TypeInt},
		{Name: "ac_dex_bonus", Type: field.TypeBool, Default: false},
		{Name: "ac_max_bonus", Type: field.TypeInt, Default: 0},
		{Name: "equipment_armor", Type: field.TypeInt, Unique: true},
	}
	// ArmorsTable holds the schema information for the "armors" table.
	ArmorsTable = &schema.Table{
		Name:       "armors",
		Columns:    ArmorsColumns,
		PrimaryKey: []*schema.Column{ArmorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "armors_equipment_armor",
				Columns:    []*schema.Column{ArmorsColumns[7]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ClassesColumns holds the columns for the "classes" table.
	ClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "hit_die", Type: field.TypeInt},
	}
	// ClassesTable holds the schema information for the "classes" table.
	ClassesTable = &schema.Table{
		Name:       "classes",
		Columns:    ClassesColumns,
		PrimaryKey: []*schema.Column{ClassesColumns[0]},
	}
	// CoinsColumns holds the columns for the "coins" table.
	CoinsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "gold_conversion_rate", Type: field.TypeFloat64},
	}
	// CoinsTable holds the schema information for the "coins" table.
	CoinsTable = &schema.Table{
		Name:       "coins",
		Columns:    CoinsColumns,
		PrimaryKey: []*schema.Column{CoinsColumns[0]},
	}
	// ConditionsColumns holds the columns for the "conditions" table.
	ConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// ConditionsTable holds the schema information for the "conditions" table.
	ConditionsTable = &schema.Table{
		Name:       "conditions",
		Columns:    ConditionsColumns,
		PrimaryKey: []*schema.Column{ConditionsColumns[0]},
	}
	// CostsColumns holds the columns for the "costs" table.
	CostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt, Default: 1},
		{Name: "cost_coin", Type: field.TypeInt},
		{Name: "equipment_cost", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// CostsTable holds the schema information for the "costs" table.
	CostsTable = &schema.Table{
		Name:       "costs",
		Columns:    CostsColumns,
		PrimaryKey: []*schema.Column{CostsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "costs_coins_coin",
				Columns:    []*schema.Column{CostsColumns[2]},
				RefColumns: []*schema.Column{CoinsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "costs_equipment_cost",
				Columns:    []*schema.Column{CostsColumns[3]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// DamageTypesColumns holds the columns for the "damage_types" table.
	DamageTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// DamageTypesTable holds the schema information for the "damage_types" table.
	DamageTypesTable = &schema.Table{
		Name:       "damage_types",
		Columns:    DamageTypesColumns,
		PrimaryKey: []*schema.Column{DamageTypesColumns[0]},
	}
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "equipment_category", Type: field.TypeEnum, Enums: []string{"GEAR", "TOOL", "WEAPON", "VEHICLE", "ARMOR"}},
		{Name: "weight", Type: field.TypeFloat64, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
	}
	// FeatsColumns holds the columns for the "feats" table.
	FeatsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// FeatsTable holds the schema information for the "feats" table.
	FeatsTable = &schema.Table{
		Name:       "feats",
		Columns:    FeatsColumns,
		PrimaryKey: []*schema.Column{FeatsColumns[0]},
	}
	// FeaturesColumns holds the columns for the "features" table.
	FeaturesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "level", Type: field.TypeInt},
	}
	// FeaturesTable holds the schema information for the "features" table.
	FeaturesTable = &schema.Table{
		Name:       "features",
		Columns:    FeaturesColumns,
		PrimaryKey: []*schema.Column{FeaturesColumns[0]},
	}
	// GearsColumns holds the columns for the "gears" table.
	GearsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "gear_category", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "equipment_gear", Type: field.TypeInt, Unique: true},
	}
	// GearsTable holds the schema information for the "gears" table.
	GearsTable = &schema.Table{
		Name:       "gears",
		Columns:    GearsColumns,
		PrimaryKey: []*schema.Column{GearsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "gears_equipment_gear",
				Columns:    []*schema.Column{GearsColumns[3]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// LanguagesColumns holds the columns for the "languages" table.
	LanguagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "language_type", Type: field.TypeEnum, Enums: []string{"STANDARD", "EXOTIC"}, Default: "STANDARD"},
		{Name: "script", Type: field.TypeEnum, Enums: []string{"Common", "Dwarvish", "Elvish", "Infernal", "Draconic", "Celestial", "Abyssal", "Giant", "Gnomish", "Goblin", "Halfling", "Orc", "Other"}, Default: "Common"},
	}
	// LanguagesTable holds the schema information for the "languages" table.
	LanguagesTable = &schema.Table{
		Name:       "languages",
		Columns:    LanguagesColumns,
		PrimaryKey: []*schema.Column{LanguagesColumns[0]},
	}
	// LanguageChoicesColumns holds the columns for the "language_choices" table.
	LanguageChoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "choose", Type: field.TypeInt},
		{Name: "race_language_options", Type: field.TypeInt, Unique: true, Nullable: true},
		{Name: "subrace_language_options", Type: field.TypeInt, Nullable: true},
	}
	// LanguageChoicesTable holds the schema information for the "language_choices" table.
	LanguageChoicesTable = &schema.Table{
		Name:       "language_choices",
		Columns:    LanguageChoicesColumns,
		PrimaryKey: []*schema.Column{LanguageChoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "language_choices_races_language_options",
				Columns:    []*schema.Column{LanguageChoicesColumns[2]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "language_choices_subraces_language_options",
				Columns:    []*schema.Column{LanguageChoicesColumns[3]},
				RefColumns: []*schema.Column{SubracesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// MagicSchoolsColumns holds the columns for the "magic_schools" table.
	MagicSchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// MagicSchoolsTable holds the schema information for the "magic_schools" table.
	MagicSchoolsTable = &schema.Table{
		Name:       "magic_schools",
		Columns:    MagicSchoolsColumns,
		PrimaryKey: []*schema.Column{MagicSchoolsColumns[0]},
	}
	// PrerequisitesColumns holds the columns for the "prerequisites" table.
	PrerequisitesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "prerequisite_type", Type: field.TypeEnum, Enums: []string{"level", "spell", "feature"}},
		{Name: "level_value", Type: field.TypeInt, Nullable: true},
		{Name: "feature_value", Type: field.TypeString, Nullable: true},
		{Name: "spell_value", Type: field.TypeString, Nullable: true},
		{Name: "feature_prerequisites", Type: field.TypeInt, Nullable: true},
	}
	// PrerequisitesTable holds the schema information for the "prerequisites" table.
	PrerequisitesTable = &schema.Table{
		Name:       "prerequisites",
		Columns:    PrerequisitesColumns,
		PrimaryKey: []*schema.Column{PrerequisitesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "prerequisites_features_prerequisites",
				Columns:    []*schema.Column{PrerequisitesColumns[5]},
				RefColumns: []*schema.Column{FeaturesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProficienciesColumns holds the columns for the "proficiencies" table.
	ProficienciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "reference", Type: field.TypeString},
	}
	// ProficienciesTable holds the schema information for the "proficiencies" table.
	ProficienciesTable = &schema.Table{
		Name:       "proficiencies",
		Columns:    ProficienciesColumns,
		PrimaryKey: []*schema.Column{ProficienciesColumns[0]},
	}
	// ProficiencyChoicesColumns holds the columns for the "proficiency_choices" table.
	ProficiencyChoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "choose", Type: field.TypeInt},
		{Name: "desc", Type: field.TypeJSON},
		{Name: "class_proficiency_choices", Type: field.TypeInt, Nullable: true},
		{Name: "proficiency_choice_subchoices", Type: field.TypeInt, Nullable: true},
		{Name: "race_starting_proficiency_options", Type: field.TypeInt, Unique: true, Nullable: true},
	}
	// ProficiencyChoicesTable holds the schema information for the "proficiency_choices" table.
	ProficiencyChoicesTable = &schema.Table{
		Name:       "proficiency_choices",
		Columns:    ProficiencyChoicesColumns,
		PrimaryKey: []*schema.Column{ProficiencyChoicesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "proficiency_choices_classes_proficiency_choices",
				Columns:    []*schema.Column{ProficiencyChoicesColumns[3]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "proficiency_choices_proficiency_choices_subchoices",
				Columns:    []*schema.Column{ProficiencyChoicesColumns[4]},
				RefColumns: []*schema.Column{ProficiencyChoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "proficiency_choices_races_starting_proficiency_options",
				Columns:    []*schema.Column{ProficiencyChoicesColumns[5]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// PropertiesColumns holds the columns for the "properties" table.
	PropertiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// PropertiesTable holds the schema information for the "properties" table.
	PropertiesTable = &schema.Table{
		Name:       "properties",
		Columns:    PropertiesColumns,
		PrimaryKey: []*schema.Column{PropertiesColumns[0]},
	}
	// RacesColumns holds the columns for the "races" table.
	RacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "speed", Type: field.TypeInt},
		{Name: "size", Type: field.TypeEnum, Enums: []string{"Small", "Medium", "Large"}, Default: "Medium"},
		{Name: "size_desc", Type: field.TypeString},
		{Name: "alignment_desc", Type: field.TypeString},
		{Name: "age_desc", Type: field.TypeString},
		{Name: "language_desc", Type: field.TypeString},
		{Name: "race_ability_bonus_options", Type: field.TypeInt, Nullable: true},
	}
	// RacesTable holds the schema information for the "races" table.
	RacesTable = &schema.Table{
		Name:       "races",
		Columns:    RacesColumns,
		PrimaryKey: []*schema.Column{RacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "races_ability_bonus_choices_ability_bonus_options",
				Columns:    []*schema.Column{RacesColumns[9]},
				RefColumns: []*schema.Column{AbilityBonusChoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RulesColumns holds the columns for the "rules" table.
	RulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// RulesTable holds the schema information for the "rules" table.
	RulesTable = &schema.Table{
		Name:       "rules",
		Columns:    RulesColumns,
		PrimaryKey: []*schema.Column{RulesColumns[0]},
	}
	// RuleSectionsColumns holds the columns for the "rule_sections" table.
	RuleSectionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "rule_id", Type: field.TypeInt, Nullable: true},
	}
	// RuleSectionsTable holds the schema information for the "rule_sections" table.
	RuleSectionsTable = &schema.Table{
		Name:       "rule_sections",
		Columns:    RuleSectionsColumns,
		PrimaryKey: []*schema.Column{RuleSectionsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rule_sections_rules_sections",
				Columns:    []*schema.Column{RuleSectionsColumns[4]},
				RefColumns: []*schema.Column{RulesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "ability_score_skills", Type: field.TypeInt, Nullable: true},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "skills_ability_scores_skills",
				Columns:    []*schema.Column{SkillsColumns[4]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// SubracesColumns holds the columns for the "subraces" table.
	SubracesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON},
		{Name: "race_subraces", Type: field.TypeInt, Nullable: true},
	}
	// SubracesTable holds the schema information for the "subraces" table.
	SubracesTable = &schema.Table{
		Name:       "subraces",
		Columns:    SubracesColumns,
		PrimaryKey: []*schema.Column{SubracesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subraces_races_subraces",
				Columns:    []*schema.Column{SubracesColumns[4]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "tool_category", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "equipment_tool", Type: field.TypeInt, Unique: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_equipment_tool",
				Columns:    []*schema.Column{ToolsColumns[3]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// TraitsColumns holds the columns for the "traits" table.
	TraitsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
	}
	// TraitsTable holds the schema information for the "traits" table.
	TraitsTable = &schema.Table{
		Name:       "traits",
		Columns:    TraitsColumns,
		PrimaryKey: []*schema.Column{TraitsColumns[0]},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "vehicle_category", Type: field.TypeEnum, Enums: []string{"mounts_and_other_animals", "tack_harness_and_drawn_vehicles", "waterborne"}},
		{Name: "capacity", Type: field.TypeString, Nullable: true},
		{Name: "desc", Type: field.TypeJSON, Nullable: true},
		{Name: "speed_quantity", Type: field.TypeFloat64, Nullable: true},
		{Name: "speed_units", Type: field.TypeEnum, Nullable: true, Enums: []string{"miles_per_hour", "feet_per_round"}},
		{Name: "equipment_vehicle", Type: field.TypeInt, Unique: true},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicles_equipment_vehicle",
				Columns:    []*schema.Column{VehiclesColumns[6]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WeaponsColumns holds the columns for the "weapons" table.
	WeaponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "weapon_category", Type: field.TypeEnum, Enums: []string{"simple", "martial", "exotic", "other"}},
		{Name: "weapon_subcategory", Type: field.TypeEnum, Enums: []string{"melee", "ranged", "other"}},
		{Name: "range_normal", Type: field.TypeInt, Nullable: true},
		{Name: "range_long", Type: field.TypeInt, Nullable: true},
		{Name: "throw_range_normal", Type: field.TypeInt, Nullable: true},
		{Name: "throw_range_long", Type: field.TypeInt, Nullable: true},
		{Name: "damage_dice", Type: field.TypeString, Nullable: true},
		{Name: "equipment_weapon", Type: field.TypeInt, Unique: true},
		{Name: "weapon_damage_type", Type: field.TypeInt, Nullable: true},
	}
	// WeaponsTable holds the schema information for the "weapons" table.
	WeaponsTable = &schema.Table{
		Name:       "weapons",
		Columns:    WeaponsColumns,
		PrimaryKey: []*schema.Column{WeaponsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weapons_equipment_weapon",
				Columns:    []*schema.Column{WeaponsColumns[8]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "weapons_damage_types_damage_type",
				Columns:    []*schema.Column{WeaponsColumns[9]},
				RefColumns: []*schema.Column{DamageTypesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AbilityBonusChoiceAbilityBonusesColumns holds the columns for the "ability_bonus_choice_ability_bonuses" table.
	AbilityBonusChoiceAbilityBonusesColumns = []*schema.Column{
		{Name: "ability_bonus_choice_id", Type: field.TypeInt},
		{Name: "ability_bonus_id", Type: field.TypeInt},
	}
	// AbilityBonusChoiceAbilityBonusesTable holds the schema information for the "ability_bonus_choice_ability_bonuses" table.
	AbilityBonusChoiceAbilityBonusesTable = &schema.Table{
		Name:       "ability_bonus_choice_ability_bonuses",
		Columns:    AbilityBonusChoiceAbilityBonusesColumns,
		PrimaryKey: []*schema.Column{AbilityBonusChoiceAbilityBonusesColumns[0], AbilityBonusChoiceAbilityBonusesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ability_bonus_choice_ability_bonuses_ability_bonus_choice_id",
				Columns:    []*schema.Column{AbilityBonusChoiceAbilityBonusesColumns[0]},
				RefColumns: []*schema.Column{AbilityBonusChoicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ability_bonus_choice_ability_bonuses_ability_bonus_id",
				Columns:    []*schema.Column{AbilityBonusChoiceAbilityBonusesColumns[1]},
				RefColumns: []*schema.Column{AbilityBonusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ClassProficienciesColumns holds the columns for the "class_proficiencies" table.
	ClassProficienciesColumns = []*schema.Column{
		{Name: "class_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// ClassProficienciesTable holds the schema information for the "class_proficiencies" table.
	ClassProficienciesTable = &schema.Table{
		Name:       "class_proficiencies",
		Columns:    ClassProficienciesColumns,
		PrimaryKey: []*schema.Column{ClassProficienciesColumns[0], ClassProficienciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_proficiencies_class_id",
				Columns:    []*schema.Column{ClassProficienciesColumns[0]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "class_proficiencies_proficiency_id",
				Columns:    []*schema.Column{ClassProficienciesColumns[1]},
				RefColumns: []*schema.Column{ProficienciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// LanguageChoiceLanguagesColumns holds the columns for the "language_choice_languages" table.
	LanguageChoiceLanguagesColumns = []*schema.Column{
		{Name: "language_choice_id", Type: field.TypeInt},
		{Name: "language_id", Type: field.TypeInt},
	}
	// LanguageChoiceLanguagesTable holds the schema information for the "language_choice_languages" table.
	LanguageChoiceLanguagesTable = &schema.Table{
		Name:       "language_choice_languages",
		Columns:    LanguageChoiceLanguagesColumns,
		PrimaryKey: []*schema.Column{LanguageChoiceLanguagesColumns[0], LanguageChoiceLanguagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "language_choice_languages_language_choice_id",
				Columns:    []*schema.Column{LanguageChoiceLanguagesColumns[0]},
				RefColumns: []*schema.Column{LanguageChoicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "language_choice_languages_language_id",
				Columns:    []*schema.Column{LanguageChoiceLanguagesColumns[1]},
				RefColumns: []*schema.Column{LanguagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProficiencyChoiceProficienciesColumns holds the columns for the "proficiency_choice_proficiencies" table.
	ProficiencyChoiceProficienciesColumns = []*schema.Column{
		{Name: "proficiency_choice_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// ProficiencyChoiceProficienciesTable holds the schema information for the "proficiency_choice_proficiencies" table.
	ProficiencyChoiceProficienciesTable = &schema.Table{
		Name:       "proficiency_choice_proficiencies",
		Columns:    ProficiencyChoiceProficienciesColumns,
		PrimaryKey: []*schema.Column{ProficiencyChoiceProficienciesColumns[0], ProficiencyChoiceProficienciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "proficiency_choice_proficiencies_proficiency_choice_id",
				Columns:    []*schema.Column{ProficiencyChoiceProficienciesColumns[0]},
				RefColumns: []*schema.Column{ProficiencyChoicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "proficiency_choice_proficiencies_proficiency_id",
				Columns:    []*schema.Column{ProficiencyChoiceProficienciesColumns[1]},
				RefColumns: []*schema.Column{ProficienciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RaceTraitsColumns holds the columns for the "race_traits" table.
	RaceTraitsColumns = []*schema.Column{
		{Name: "race_id", Type: field.TypeInt},
		{Name: "trait_id", Type: field.TypeInt},
	}
	// RaceTraitsTable holds the schema information for the "race_traits" table.
	RaceTraitsTable = &schema.Table{
		Name:       "race_traits",
		Columns:    RaceTraitsColumns,
		PrimaryKey: []*schema.Column{RaceTraitsColumns[0], RaceTraitsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "race_traits_race_id",
				Columns:    []*schema.Column{RaceTraitsColumns[0]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "race_traits_trait_id",
				Columns:    []*schema.Column{RaceTraitsColumns[1]},
				RefColumns: []*schema.Column{TraitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RaceStartingProficienciesColumns holds the columns for the "race_starting_proficiencies" table.
	RaceStartingProficienciesColumns = []*schema.Column{
		{Name: "race_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// RaceStartingProficienciesTable holds the schema information for the "race_starting_proficiencies" table.
	RaceStartingProficienciesTable = &schema.Table{
		Name:       "race_starting_proficiencies",
		Columns:    RaceStartingProficienciesColumns,
		PrimaryKey: []*schema.Column{RaceStartingProficienciesColumns[0], RaceStartingProficienciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "race_starting_proficiencies_race_id",
				Columns:    []*schema.Column{RaceStartingProficienciesColumns[0]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "race_starting_proficiencies_proficiency_id",
				Columns:    []*schema.Column{RaceStartingProficienciesColumns[1]},
				RefColumns: []*schema.Column{ProficienciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RaceAbilityBonusesColumns holds the columns for the "race_ability_bonuses" table.
	RaceAbilityBonusesColumns = []*schema.Column{
		{Name: "race_id", Type: field.TypeInt},
		{Name: "ability_bonus_id", Type: field.TypeInt},
	}
	// RaceAbilityBonusesTable holds the schema information for the "race_ability_bonuses" table.
	RaceAbilityBonusesTable = &schema.Table{
		Name:       "race_ability_bonuses",
		Columns:    RaceAbilityBonusesColumns,
		PrimaryKey: []*schema.Column{RaceAbilityBonusesColumns[0], RaceAbilityBonusesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "race_ability_bonuses_race_id",
				Columns:    []*schema.Column{RaceAbilityBonusesColumns[0]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "race_ability_bonuses_ability_bonus_id",
				Columns:    []*schema.Column{RaceAbilityBonusesColumns[1]},
				RefColumns: []*schema.Column{AbilityBonusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// RaceLanguagesColumns holds the columns for the "race_languages" table.
	RaceLanguagesColumns = []*schema.Column{
		{Name: "race_id", Type: field.TypeInt},
		{Name: "language_id", Type: field.TypeInt},
	}
	// RaceLanguagesTable holds the schema information for the "race_languages" table.
	RaceLanguagesTable = &schema.Table{
		Name:       "race_languages",
		Columns:    RaceLanguagesColumns,
		PrimaryKey: []*schema.Column{RaceLanguagesColumns[0], RaceLanguagesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "race_languages_race_id",
				Columns:    []*schema.Column{RaceLanguagesColumns[0]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "race_languages_language_id",
				Columns:    []*schema.Column{RaceLanguagesColumns[1]},
				RefColumns: []*schema.Column{LanguagesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubraceAbilityBonusesColumns holds the columns for the "subrace_ability_bonuses" table.
	SubraceAbilityBonusesColumns = []*schema.Column{
		{Name: "subrace_id", Type: field.TypeInt},
		{Name: "ability_bonus_id", Type: field.TypeInt},
	}
	// SubraceAbilityBonusesTable holds the schema information for the "subrace_ability_bonuses" table.
	SubraceAbilityBonusesTable = &schema.Table{
		Name:       "subrace_ability_bonuses",
		Columns:    SubraceAbilityBonusesColumns,
		PrimaryKey: []*schema.Column{SubraceAbilityBonusesColumns[0], SubraceAbilityBonusesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subrace_ability_bonuses_subrace_id",
				Columns:    []*schema.Column{SubraceAbilityBonusesColumns[0]},
				RefColumns: []*schema.Column{SubracesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subrace_ability_bonuses_ability_bonus_id",
				Columns:    []*schema.Column{SubraceAbilityBonusesColumns[1]},
				RefColumns: []*schema.Column{AbilityBonusColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubraceProficienciesColumns holds the columns for the "subrace_proficiencies" table.
	SubraceProficienciesColumns = []*schema.Column{
		{Name: "subrace_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// SubraceProficienciesTable holds the schema information for the "subrace_proficiencies" table.
	SubraceProficienciesTable = &schema.Table{
		Name:       "subrace_proficiencies",
		Columns:    SubraceProficienciesColumns,
		PrimaryKey: []*schema.Column{SubraceProficienciesColumns[0], SubraceProficienciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subrace_proficiencies_subrace_id",
				Columns:    []*schema.Column{SubraceProficienciesColumns[0]},
				RefColumns: []*schema.Column{SubracesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subrace_proficiencies_proficiency_id",
				Columns:    []*schema.Column{SubraceProficienciesColumns[1]},
				RefColumns: []*schema.Column{ProficienciesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// SubraceTraitsColumns holds the columns for the "subrace_traits" table.
	SubraceTraitsColumns = []*schema.Column{
		{Name: "subrace_id", Type: field.TypeInt},
		{Name: "trait_id", Type: field.TypeInt},
	}
	// SubraceTraitsTable holds the schema information for the "subrace_traits" table.
	SubraceTraitsTable = &schema.Table{
		Name:       "subrace_traits",
		Columns:    SubraceTraitsColumns,
		PrimaryKey: []*schema.Column{SubraceTraitsColumns[0], SubraceTraitsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "subrace_traits_subrace_id",
				Columns:    []*schema.Column{SubraceTraitsColumns[0]},
				RefColumns: []*schema.Column{SubracesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "subrace_traits_trait_id",
				Columns:    []*schema.Column{SubraceTraitsColumns[1]},
				RefColumns: []*schema.Column{TraitsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// WeaponPropertiesColumns holds the columns for the "weapon_properties" table.
	WeaponPropertiesColumns = []*schema.Column{
		{Name: "weapon_id", Type: field.TypeInt},
		{Name: "property_id", Type: field.TypeInt},
	}
	// WeaponPropertiesTable holds the schema information for the "weapon_properties" table.
	WeaponPropertiesTable = &schema.Table{
		Name:       "weapon_properties",
		Columns:    WeaponPropertiesColumns,
		PrimaryKey: []*schema.Column{WeaponPropertiesColumns[0], WeaponPropertiesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weapon_properties_weapon_id",
				Columns:    []*schema.Column{WeaponPropertiesColumns[0]},
				RefColumns: []*schema.Column{WeaponsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "weapon_properties_property_id",
				Columns:    []*schema.Column{WeaponPropertiesColumns[1]},
				RefColumns: []*schema.Column{PropertiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityBonusTable,
		AbilityBonusChoicesTable,
		AbilityScoresTable,
		AlignmentsTable,
		ArmorsTable,
		ClassesTable,
		CoinsTable,
		ConditionsTable,
		CostsTable,
		DamageTypesTable,
		EquipmentTable,
		FeatsTable,
		FeaturesTable,
		GearsTable,
		LanguagesTable,
		LanguageChoicesTable,
		MagicSchoolsTable,
		PrerequisitesTable,
		ProficienciesTable,
		ProficiencyChoicesTable,
		PropertiesTable,
		RacesTable,
		RulesTable,
		RuleSectionsTable,
		SkillsTable,
		SubracesTable,
		ToolsTable,
		TraitsTable,
		VehiclesTable,
		WeaponsTable,
		AbilityBonusChoiceAbilityBonusesTable,
		ClassProficienciesTable,
		LanguageChoiceLanguagesTable,
		ProficiencyChoiceProficienciesTable,
		RaceTraitsTable,
		RaceStartingProficienciesTable,
		RaceAbilityBonusesTable,
		RaceLanguagesTable,
		SubraceAbilityBonusesTable,
		SubraceProficienciesTable,
		SubraceTraitsTable,
		WeaponPropertiesTable,
	}
)

func init() {
	AbilityBonusTable.ForeignKeys[0].RefTable = AbilityScoresTable
	ArmorsTable.ForeignKeys[0].RefTable = EquipmentTable
	CostsTable.ForeignKeys[0].RefTable = CoinsTable
	CostsTable.ForeignKeys[1].RefTable = EquipmentTable
	GearsTable.ForeignKeys[0].RefTable = EquipmentTable
	LanguageChoicesTable.ForeignKeys[0].RefTable = RacesTable
	LanguageChoicesTable.ForeignKeys[1].RefTable = SubracesTable
	PrerequisitesTable.ForeignKeys[0].RefTable = FeaturesTable
	ProficiencyChoicesTable.ForeignKeys[0].RefTable = ClassesTable
	ProficiencyChoicesTable.ForeignKeys[1].RefTable = ProficiencyChoicesTable
	ProficiencyChoicesTable.ForeignKeys[2].RefTable = RacesTable
	RacesTable.ForeignKeys[0].RefTable = AbilityBonusChoicesTable
	RuleSectionsTable.ForeignKeys[0].RefTable = RulesTable
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
	SubracesTable.ForeignKeys[0].RefTable = RacesTable
	ToolsTable.ForeignKeys[0].RefTable = EquipmentTable
	VehiclesTable.ForeignKeys[0].RefTable = EquipmentTable
	WeaponsTable.ForeignKeys[0].RefTable = EquipmentTable
	WeaponsTable.ForeignKeys[1].RefTable = DamageTypesTable
	AbilityBonusChoiceAbilityBonusesTable.ForeignKeys[0].RefTable = AbilityBonusChoicesTable
	AbilityBonusChoiceAbilityBonusesTable.ForeignKeys[1].RefTable = AbilityBonusTable
	ClassProficienciesTable.ForeignKeys[0].RefTable = ClassesTable
	ClassProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	LanguageChoiceLanguagesTable.ForeignKeys[0].RefTable = LanguageChoicesTable
	LanguageChoiceLanguagesTable.ForeignKeys[1].RefTable = LanguagesTable
	ProficiencyChoiceProficienciesTable.ForeignKeys[0].RefTable = ProficiencyChoicesTable
	ProficiencyChoiceProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	RaceTraitsTable.ForeignKeys[0].RefTable = RacesTable
	RaceTraitsTable.ForeignKeys[1].RefTable = TraitsTable
	RaceStartingProficienciesTable.ForeignKeys[0].RefTable = RacesTable
	RaceStartingProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	RaceAbilityBonusesTable.ForeignKeys[0].RefTable = RacesTable
	RaceAbilityBonusesTable.ForeignKeys[1].RefTable = AbilityBonusTable
	RaceLanguagesTable.ForeignKeys[0].RefTable = RacesTable
	RaceLanguagesTable.ForeignKeys[1].RefTable = LanguagesTable
	SubraceAbilityBonusesTable.ForeignKeys[0].RefTable = SubracesTable
	SubraceAbilityBonusesTable.ForeignKeys[1].RefTable = AbilityBonusTable
	SubraceProficienciesTable.ForeignKeys[0].RefTable = SubracesTable
	SubraceProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	SubraceTraitsTable.ForeignKeys[0].RefTable = SubracesTable
	SubraceTraitsTable.ForeignKeys[1].RefTable = TraitsTable
	WeaponPropertiesTable.ForeignKeys[0].RefTable = WeaponsTable
	WeaponPropertiesTable.ForeignKeys[1].RefTable = PropertiesTable
}
