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
		{Name: "ability_score_id", Type: field.TypeInt},
		{Name: "race_ability_bonuses", Type: field.TypeInt, Nullable: true},
		{Name: "subrace_ability_bonuses", Type: field.TypeInt, Nullable: true},
	}
	// AbilityBonusTable holds the schema information for the "ability_bonus" table.
	AbilityBonusTable = &schema.Table{
		Name:       "ability_bonus",
		Columns:    AbilityBonusColumns,
		PrimaryKey: []*schema.Column{AbilityBonusColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ability_bonus_ability_scores_ability_bonuses",
				Columns:    []*schema.Column{AbilityBonusColumns[2]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "ability_bonus_races_ability_bonuses",
				Columns:    []*schema.Column{AbilityBonusColumns[3]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ability_bonus_subraces_ability_bonuses",
				Columns:    []*schema.Column{AbilityBonusColumns[4]},
				RefColumns: []*schema.Column{SubracesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// AbilityScoresColumns holds the columns for the "ability_scores" table.
	AbilityScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "full_name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON},
	}
	// AbilityScoresTable holds the schema information for the "ability_scores" table.
	AbilityScoresTable = &schema.Table{
		Name:       "ability_scores",
		Columns:    AbilityScoresColumns,
		PrimaryKey: []*schema.Column{AbilityScoresColumns[0]},
	}
	// ArmorsColumns holds the columns for the "armors" table.
	ArmorsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "armor_category", Type: field.TypeString},
		{Name: "stealth_disadvantage", Type: field.TypeBool},
		{Name: "min_strength", Type: field.TypeInt},
		{Name: "equipment_id", Type: field.TypeInt, Unique: true},
	}
	// ArmorsTable holds the schema information for the "armors" table.
	ArmorsTable = &schema.Table{
		Name:       "armors",
		Columns:    ArmorsColumns,
		PrimaryKey: []*schema.Column{ArmorsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "armors_equipment_armor",
				Columns:    []*schema.Column{ArmorsColumns[6]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// ArmorClassesColumns holds the columns for the "armor_classes" table.
	ArmorClassesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "base", Type: field.TypeInt},
		{Name: "dex_bonus", Type: field.TypeBool},
		{Name: "max_bonus", Type: field.TypeInt, Nullable: true},
		{Name: "armor_armor_class", Type: field.TypeInt, Nullable: true},
	}
	// ArmorClassesTable holds the schema information for the "armor_classes" table.
	ArmorClassesTable = &schema.Table{
		Name:       "armor_classes",
		Columns:    ArmorClassesColumns,
		PrimaryKey: []*schema.Column{ArmorClassesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "armor_classes_armors_armor_class",
				Columns:    []*schema.Column{ArmorClassesColumns[4]},
				RefColumns: []*schema.Column{ArmorsColumns[0]},
				OnDelete:   schema.SetNull,
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
	// CostsColumns holds the columns for the "costs" table.
	CostsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "quantity", Type: field.TypeInt},
		{Name: "unit", Type: field.TypeString},
	}
	// CostsTable holds the schema information for the "costs" table.
	CostsTable = &schema.Table{
		Name:       "costs",
		Columns:    CostsColumns,
		PrimaryKey: []*schema.Column{CostsColumns[0]},
	}
	// DamageTypesColumns holds the columns for the "damage_types" table.
	DamageTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON},
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
		{Name: "equipment_category", Type: field.TypeEnum, Enums: []string{"weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "other"}, Default: "other"},
		{Name: "equipment_cost", Type: field.TypeInt, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equipment_costs_cost",
				Columns:    []*schema.Column{EquipmentColumns[4]},
				RefColumns: []*schema.Column{CostsColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// GearsColumns holds the columns for the "gears" table.
	GearsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "gear_category", Type: field.TypeEnum, Enums: []string{"ammunition", "standard_gear", "kits", "equipment_packs", "arcane_foci", "druidic_foci", "holy_symbols", "other"}, Default: "other"},
		{Name: "desc", Type: field.TypeJSON},
		{Name: "quantity", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_id", Type: field.TypeInt, Unique: true},
	}
	// GearsTable holds the schema information for the "gears" table.
	GearsTable = &schema.Table{
		Name:       "gears",
		Columns:    GearsColumns,
		PrimaryKey: []*schema.Column{GearsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "gears_equipment_gear",
				Columns:    []*schema.Column{GearsColumns[6]},
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
		{Name: "desc", Type: field.TypeString},
		{Name: "language_type", Type: field.TypeEnum, Enums: []string{"STANDARD", "EXOTIC"}, Default: "STANDARD"},
		{Name: "script", Type: field.TypeEnum, Nullable: true, Enums: []string{"Common", "Dwarvish", "Elvish", "Infernal", "Draconic", "Celestial", "Abyssal", "Giant", "Gnomish", "Goblin", "Halfling", "Orc", "Other"}, Default: "Common"},
	}
	// LanguagesTable holds the schema information for the "languages" table.
	LanguagesTable = &schema.Table{
		Name:       "languages",
		Columns:    LanguagesColumns,
		PrimaryKey: []*schema.Column{LanguagesColumns[0]},
	}
	// MagicSchoolsColumns holds the columns for the "magic_schools" table.
	MagicSchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
	}
	// MagicSchoolsTable holds the schema information for the "magic_schools" table.
	MagicSchoolsTable = &schema.Table{
		Name:       "magic_schools",
		Columns:    MagicSchoolsColumns,
		PrimaryKey: []*schema.Column{MagicSchoolsColumns[0]},
	}
	// ProficienciesColumns holds the columns for the "proficiencies" table.
	ProficienciesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "proficiency_category", Type: field.TypeString},
		{Name: "proficiency_skill", Type: field.TypeInt, Nullable: true},
		{Name: "proficiency_equipment", Type: field.TypeInt, Nullable: true},
		{Name: "proficiency_saving_throw", Type: field.TypeInt, Nullable: true},
	}
	// ProficienciesTable holds the schema information for the "proficiencies" table.
	ProficienciesTable = &schema.Table{
		Name:       "proficiencies",
		Columns:    ProficienciesColumns,
		PrimaryKey: []*schema.Column{ProficienciesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "proficiencies_skills_skill",
				Columns:    []*schema.Column{ProficienciesColumns[4]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "proficiencies_equipment_equipment",
				Columns:    []*schema.Column{ProficienciesColumns[5]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "proficiencies_ability_scores_saving_throw",
				Columns:    []*schema.Column{ProficienciesColumns[6]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ProficiencyChoicesColumns holds the columns for the "proficiency_choices" table.
	ProficiencyChoicesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "desc", Type: field.TypeString, Nullable: true},
		{Name: "choose", Type: field.TypeInt},
	}
	// ProficiencyChoicesTable holds the schema information for the "proficiency_choices" table.
	ProficiencyChoicesTable = &schema.Table{
		Name:       "proficiency_choices",
		Columns:    ProficiencyChoicesColumns,
		PrimaryKey: []*schema.Column{ProficiencyChoicesColumns[0]},
	}
	// RacesColumns holds the columns for the "races" table.
	RacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "alignment", Type: field.TypeString},
		{Name: "age", Type: field.TypeString},
		{Name: "size", Type: field.TypeString},
		{Name: "size_description", Type: field.TypeString},
		{Name: "language_desc", Type: field.TypeString},
		{Name: "speed", Type: field.TypeInt},
		{Name: "race_starting_proficiency_option", Type: field.TypeInt, Nullable: true},
	}
	// RacesTable holds the schema information for the "races" table.
	RacesTable = &schema.Table{
		Name:       "races",
		Columns:    RacesColumns,
		PrimaryKey: []*schema.Column{RacesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "races_proficiency_choices_starting_proficiency_option",
				Columns:    []*schema.Column{RacesColumns[9]},
				RefColumns: []*schema.Column{ProficiencyChoicesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// RulesColumns holds the columns for the "rules" table.
	RulesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
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
		{Name: "desc", Type: field.TypeString},
	}
	// RuleSectionsTable holds the schema information for the "rule_sections" table.
	RuleSectionsTable = &schema.Table{
		Name:       "rule_sections",
		Columns:    RuleSectionsColumns,
		PrimaryKey: []*schema.Column{RuleSectionsColumns[0]},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON},
		{Name: "skill_ability_score", Type: field.TypeInt, Nullable: true},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "skills_ability_scores_ability_score",
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
		{Name: "desc", Type: field.TypeString},
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
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "tool_category", Type: field.TypeString},
		{Name: "equipment_id", Type: field.TypeInt, Unique: true},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "tools_equipment_tool",
				Columns:    []*schema.Column{ToolsColumns[4]},
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
		{Name: "desc", Type: field.TypeJSON},
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
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "vehicle_category", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeString},
		{Name: "equipment_id", Type: field.TypeInt, Unique: true},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "vehicles_equipment_vehicle",
				Columns:    []*schema.Column{VehiclesColumns[5]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WeaponsColumns holds the columns for the "weapons" table.
	WeaponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "weapon_category", Type: field.TypeString},
		{Name: "weapon_range", Type: field.TypeString},
		{Name: "equipment_id", Type: field.TypeInt, Unique: true},
	}
	// WeaponsTable holds the schema information for the "weapons" table.
	WeaponsTable = &schema.Table{
		Name:       "weapons",
		Columns:    WeaponsColumns,
		PrimaryKey: []*schema.Column{WeaponsColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weapons_equipment_weapon",
				Columns:    []*schema.Column{WeaponsColumns[5]},
				RefColumns: []*schema.Column{EquipmentColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WeaponDamagesColumns holds the columns for the "weapon_damages" table.
	WeaponDamagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dice", Type: field.TypeString},
		{Name: "weapon_id", Type: field.TypeInt},
		{Name: "damage_type_id", Type: field.TypeInt},
	}
	// WeaponDamagesTable holds the schema information for the "weapon_damages" table.
	WeaponDamagesTable = &schema.Table{
		Name:       "weapon_damages",
		Columns:    WeaponDamagesColumns,
		PrimaryKey: []*schema.Column{WeaponDamagesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weapon_damages_weapons_weapon_damage",
				Columns:    []*schema.Column{WeaponDamagesColumns[2]},
				RefColumns: []*schema.Column{WeaponsColumns[0]},
				OnDelete:   schema.NoAction,
			},
			{
				Symbol:     "weapon_damages_damage_types_damage_type",
				Columns:    []*schema.Column{WeaponDamagesColumns[3]},
				RefColumns: []*schema.Column{DamageTypesColumns[0]},
				OnDelete:   schema.NoAction,
			},
		},
	}
	// WeaponPropertiesColumns holds the columns for the "weapon_properties" table.
	WeaponPropertiesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeJSON},
	}
	// WeaponPropertiesTable holds the schema information for the "weapon_properties" table.
	WeaponPropertiesTable = &schema.Table{
		Name:       "weapon_properties",
		Columns:    WeaponPropertiesColumns,
		PrimaryKey: []*schema.Column{WeaponPropertiesColumns[0]},
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
	// ClassProficiencyChoicesColumns holds the columns for the "class_proficiency_choices" table.
	ClassProficiencyChoicesColumns = []*schema.Column{
		{Name: "class_id", Type: field.TypeInt},
		{Name: "proficiency_choice_id", Type: field.TypeInt},
	}
	// ClassProficiencyChoicesTable holds the schema information for the "class_proficiency_choices" table.
	ClassProficiencyChoicesTable = &schema.Table{
		Name:       "class_proficiency_choices",
		Columns:    ClassProficiencyChoicesColumns,
		PrimaryKey: []*schema.Column{ClassProficiencyChoicesColumns[0], ClassProficiencyChoicesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_proficiency_choices_class_id",
				Columns:    []*schema.Column{ClassProficiencyChoicesColumns[0]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "class_proficiency_choices_proficiency_choice_id",
				Columns:    []*schema.Column{ClassProficiencyChoicesColumns[1]},
				RefColumns: []*schema.Column{ProficiencyChoicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// ProficiencyChoiceOptionsColumns holds the columns for the "proficiency_choice_options" table.
	ProficiencyChoiceOptionsColumns = []*schema.Column{
		{Name: "proficiency_choice_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// ProficiencyChoiceOptionsTable holds the schema information for the "proficiency_choice_options" table.
	ProficiencyChoiceOptionsTable = &schema.Table{
		Name:       "proficiency_choice_options",
		Columns:    ProficiencyChoiceOptionsColumns,
		PrimaryKey: []*schema.Column{ProficiencyChoiceOptionsColumns[0], ProficiencyChoiceOptionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "proficiency_choice_options_proficiency_choice_id",
				Columns:    []*schema.Column{ProficiencyChoiceOptionsColumns[0]},
				RefColumns: []*schema.Column{ProficiencyChoicesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "proficiency_choice_options_proficiency_id",
				Columns:    []*schema.Column{ProficiencyChoiceOptionsColumns[1]},
				RefColumns: []*schema.Column{ProficienciesColumns[0]},
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
	// RaceProficienciesColumns holds the columns for the "race_proficiencies" table.
	RaceProficienciesColumns = []*schema.Column{
		{Name: "race_id", Type: field.TypeInt},
		{Name: "proficiency_id", Type: field.TypeInt},
	}
	// RaceProficienciesTable holds the schema information for the "race_proficiencies" table.
	RaceProficienciesTable = &schema.Table{
		Name:       "race_proficiencies",
		Columns:    RaceProficienciesColumns,
		PrimaryKey: []*schema.Column{RaceProficienciesColumns[0], RaceProficienciesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "race_proficiencies_race_id",
				Columns:    []*schema.Column{RaceProficienciesColumns[0]},
				RefColumns: []*schema.Column{RacesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "race_proficiencies_proficiency_id",
				Columns:    []*schema.Column{RaceProficienciesColumns[1]},
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
	// RuleRuleSectionsColumns holds the columns for the "rule_rule_sections" table.
	RuleRuleSectionsColumns = []*schema.Column{
		{Name: "rule_id", Type: field.TypeInt},
		{Name: "rule_section_id", Type: field.TypeInt},
	}
	// RuleRuleSectionsTable holds the schema information for the "rule_rule_sections" table.
	RuleRuleSectionsTable = &schema.Table{
		Name:       "rule_rule_sections",
		Columns:    RuleRuleSectionsColumns,
		PrimaryKey: []*schema.Column{RuleRuleSectionsColumns[0], RuleRuleSectionsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "rule_rule_sections_rule_id",
				Columns:    []*schema.Column{RuleRuleSectionsColumns[0]},
				RefColumns: []*schema.Column{RulesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "rule_rule_sections_rule_section_id",
				Columns:    []*schema.Column{RuleRuleSectionsColumns[1]},
				RefColumns: []*schema.Column{RuleSectionsColumns[0]},
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
	// WeaponWeaponPropertiesColumns holds the columns for the "weapon_weapon_properties" table.
	WeaponWeaponPropertiesColumns = []*schema.Column{
		{Name: "weapon_id", Type: field.TypeInt},
		{Name: "weapon_property_id", Type: field.TypeInt},
	}
	// WeaponWeaponPropertiesTable holds the schema information for the "weapon_weapon_properties" table.
	WeaponWeaponPropertiesTable = &schema.Table{
		Name:       "weapon_weapon_properties",
		Columns:    WeaponWeaponPropertiesColumns,
		PrimaryKey: []*schema.Column{WeaponWeaponPropertiesColumns[0], WeaponWeaponPropertiesColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "weapon_weapon_properties_weapon_id",
				Columns:    []*schema.Column{WeaponWeaponPropertiesColumns[0]},
				RefColumns: []*schema.Column{WeaponsColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "weapon_weapon_properties_weapon_property_id",
				Columns:    []*schema.Column{WeaponWeaponPropertiesColumns[1]},
				RefColumns: []*schema.Column{WeaponPropertiesColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityBonusTable,
		AbilityScoresTable,
		ArmorsTable,
		ArmorClassesTable,
		ClassesTable,
		CostsTable,
		DamageTypesTable,
		EquipmentTable,
		GearsTable,
		LanguagesTable,
		MagicSchoolsTable,
		ProficienciesTable,
		ProficiencyChoicesTable,
		RacesTable,
		RulesTable,
		RuleSectionsTable,
		SkillsTable,
		SubracesTable,
		ToolsTable,
		TraitsTable,
		VehiclesTable,
		WeaponsTable,
		WeaponDamagesTable,
		WeaponPropertiesTable,
		ClassProficienciesTable,
		ClassProficiencyChoicesTable,
		ProficiencyChoiceOptionsTable,
		RaceLanguagesTable,
		RaceProficienciesTable,
		RaceTraitsTable,
		RuleRuleSectionsTable,
		SubraceProficienciesTable,
		SubraceTraitsTable,
		WeaponWeaponPropertiesTable,
	}
)

func init() {
	AbilityBonusTable.ForeignKeys[0].RefTable = AbilityScoresTable
	AbilityBonusTable.ForeignKeys[1].RefTable = RacesTable
	AbilityBonusTable.ForeignKeys[2].RefTable = SubracesTable
	ArmorsTable.ForeignKeys[0].RefTable = EquipmentTable
	ArmorClassesTable.ForeignKeys[0].RefTable = ArmorsTable
	EquipmentTable.ForeignKeys[0].RefTable = CostsTable
	GearsTable.ForeignKeys[0].RefTable = EquipmentTable
	ProficienciesTable.ForeignKeys[0].RefTable = SkillsTable
	ProficienciesTable.ForeignKeys[1].RefTable = EquipmentTable
	ProficienciesTable.ForeignKeys[2].RefTable = AbilityScoresTable
	RacesTable.ForeignKeys[0].RefTable = ProficiencyChoicesTable
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
	SubracesTable.ForeignKeys[0].RefTable = RacesTable
	ToolsTable.ForeignKeys[0].RefTable = EquipmentTable
	VehiclesTable.ForeignKeys[0].RefTable = EquipmentTable
	WeaponsTable.ForeignKeys[0].RefTable = EquipmentTable
	WeaponDamagesTable.ForeignKeys[0].RefTable = WeaponsTable
	WeaponDamagesTable.ForeignKeys[1].RefTable = DamageTypesTable
	ClassProficienciesTable.ForeignKeys[0].RefTable = ClassesTable
	ClassProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	ClassProficiencyChoicesTable.ForeignKeys[0].RefTable = ClassesTable
	ClassProficiencyChoicesTable.ForeignKeys[1].RefTable = ProficiencyChoicesTable
	ProficiencyChoiceOptionsTable.ForeignKeys[0].RefTable = ProficiencyChoicesTable
	ProficiencyChoiceOptionsTable.ForeignKeys[1].RefTable = ProficienciesTable
	RaceLanguagesTable.ForeignKeys[0].RefTable = RacesTable
	RaceLanguagesTable.ForeignKeys[1].RefTable = LanguagesTable
	RaceProficienciesTable.ForeignKeys[0].RefTable = RacesTable
	RaceProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	RaceTraitsTable.ForeignKeys[0].RefTable = RacesTable
	RaceTraitsTable.ForeignKeys[1].RefTable = TraitsTable
	RuleRuleSectionsTable.ForeignKeys[0].RefTable = RulesTable
	RuleRuleSectionsTable.ForeignKeys[1].RefTable = RuleSectionsTable
	SubraceProficienciesTable.ForeignKeys[0].RefTable = SubracesTable
	SubraceProficienciesTable.ForeignKeys[1].RefTable = ProficienciesTable
	SubraceTraitsTable.ForeignKeys[0].RefTable = SubracesTable
	SubraceTraitsTable.ForeignKeys[1].RefTable = TraitsTable
	WeaponWeaponPropertiesTable.ForeignKeys[0].RefTable = WeaponsTable
	WeaponWeaponPropertiesTable.ForeignKeys[1].RefTable = WeaponPropertiesTable
}
