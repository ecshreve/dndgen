// Code generated by ent, DO NOT EDIT.

package migrate

import (
	"entgo.io/ent/dialect/sql/schema"
	"entgo.io/ent/schema/field"
)

var (
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
		{Name: "stealth_disadvantage", Type: field.TypeBool},
		{Name: "min_strength", Type: field.TypeInt},
	}
	// ArmorsTable holds the schema information for the "armors" table.
	ArmorsTable = &schema.Table{
		Name:       "armors",
		Columns:    ArmorsColumns,
		PrimaryKey: []*schema.Column{ArmorsColumns[0]},
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
		{Name: "weapon_damage_damage_type", Type: field.TypeInt, Nullable: true},
	}
	// DamageTypesTable holds the schema information for the "damage_types" table.
	DamageTypesTable = &schema.Table{
		Name:       "damage_types",
		Columns:    DamageTypesColumns,
		PrimaryKey: []*schema.Column{DamageTypesColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "damage_types_weapon_damages_damage_type",
				Columns:    []*schema.Column{DamageTypesColumns[4]},
				RefColumns: []*schema.Column{WeaponDamagesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// EquipmentColumns holds the columns for the "equipment" table.
	EquipmentColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "equipment_category", Type: field.TypeEnum, Enums: []string{"weapon", "armor", "adventuring_gear", "tools", "mounts_and_vehicles", "other"}, Default: "other"},
		{Name: "equipment_weapon", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_armor", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_gear", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_tool", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_vehicle", Type: field.TypeInt, Nullable: true},
		{Name: "equipment_cost", Type: field.TypeInt, Nullable: true},
	}
	// EquipmentTable holds the schema information for the "equipment" table.
	EquipmentTable = &schema.Table{
		Name:       "equipment",
		Columns:    EquipmentColumns,
		PrimaryKey: []*schema.Column{EquipmentColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "equipment_weapons_weapon",
				Columns:    []*schema.Column{EquipmentColumns[4]},
				RefColumns: []*schema.Column{WeaponsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_armors_armor",
				Columns:    []*schema.Column{EquipmentColumns[5]},
				RefColumns: []*schema.Column{ArmorsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_gears_gear",
				Columns:    []*schema.Column{EquipmentColumns[6]},
				RefColumns: []*schema.Column{GearsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_tools_tool",
				Columns:    []*schema.Column{EquipmentColumns[7]},
				RefColumns: []*schema.Column{ToolsColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_vehicles_vehicle",
				Columns:    []*schema.Column{EquipmentColumns[8]},
				RefColumns: []*schema.Column{VehiclesColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "equipment_costs_cost",
				Columns:    []*schema.Column{EquipmentColumns[9]},
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
	}
	// GearsTable holds the schema information for the "gears" table.
	GearsTable = &schema.Table{
		Name:       "gears",
		Columns:    GearsColumns,
		PrimaryKey: []*schema.Column{GearsColumns[0]},
	}
	// RacesColumns holds the columns for the "races" table.
	RacesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "speed", Type: field.TypeInt},
	}
	// RacesTable holds the schema information for the "races" table.
	RacesTable = &schema.Table{
		Name:       "races",
		Columns:    RacesColumns,
		PrimaryKey: []*schema.Column{RacesColumns[0]},
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
	// ToolsColumns holds the columns for the "tools" table.
	ToolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "tool_category", Type: field.TypeString},
	}
	// ToolsTable holds the schema information for the "tools" table.
	ToolsTable = &schema.Table{
		Name:       "tools",
		Columns:    ToolsColumns,
		PrimaryKey: []*schema.Column{ToolsColumns[0]},
	}
	// VehiclesColumns holds the columns for the "vehicles" table.
	VehiclesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "vehicle_category", Type: field.TypeString},
		{Name: "capacity", Type: field.TypeString},
	}
	// VehiclesTable holds the schema information for the "vehicles" table.
	VehiclesTable = &schema.Table{
		Name:       "vehicles",
		Columns:    VehiclesColumns,
		PrimaryKey: []*schema.Column{VehiclesColumns[0]},
	}
	// WeaponsColumns holds the columns for the "weapons" table.
	WeaponsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString, Unique: true},
		{Name: "name", Type: field.TypeString},
		{Name: "weapon_range", Type: field.TypeString},
	}
	// WeaponsTable holds the schema information for the "weapons" table.
	WeaponsTable = &schema.Table{
		Name:       "weapons",
		Columns:    WeaponsColumns,
		PrimaryKey: []*schema.Column{WeaponsColumns[0]},
	}
	// WeaponDamagesColumns holds the columns for the "weapon_damages" table.
	WeaponDamagesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "dice", Type: field.TypeString},
	}
	// WeaponDamagesTable holds the schema information for the "weapon_damages" table.
	WeaponDamagesTable = &schema.Table{
		Name:       "weapon_damages",
		Columns:    WeaponDamagesColumns,
		PrimaryKey: []*schema.Column{WeaponDamagesColumns[0]},
	}
	// ClassSavingThrowsColumns holds the columns for the "class_saving_throws" table.
	ClassSavingThrowsColumns = []*schema.Column{
		{Name: "class_id", Type: field.TypeInt},
		{Name: "ability_score_id", Type: field.TypeInt},
	}
	// ClassSavingThrowsTable holds the schema information for the "class_saving_throws" table.
	ClassSavingThrowsTable = &schema.Table{
		Name:       "class_saving_throws",
		Columns:    ClassSavingThrowsColumns,
		PrimaryKey: []*schema.Column{ClassSavingThrowsColumns[0], ClassSavingThrowsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "class_saving_throws_class_id",
				Columns:    []*schema.Column{ClassSavingThrowsColumns[0]},
				RefColumns: []*schema.Column{ClassesColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "class_saving_throws_ability_score_id",
				Columns:    []*schema.Column{ClassSavingThrowsColumns[1]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityScoresTable,
		ArmorsTable,
		ArmorClassesTable,
		ClassesTable,
		CostsTable,
		DamageTypesTable,
		EquipmentTable,
		GearsTable,
		RacesTable,
		SkillsTable,
		ToolsTable,
		VehiclesTable,
		WeaponsTable,
		WeaponDamagesTable,
		ClassSavingThrowsTable,
	}
)

func init() {
	ArmorClassesTable.ForeignKeys[0].RefTable = ArmorsTable
	DamageTypesTable.ForeignKeys[0].RefTable = WeaponDamagesTable
	EquipmentTable.ForeignKeys[0].RefTable = WeaponsTable
	EquipmentTable.ForeignKeys[1].RefTable = ArmorsTable
	EquipmentTable.ForeignKeys[2].RefTable = GearsTable
	EquipmentTable.ForeignKeys[3].RefTable = ToolsTable
	EquipmentTable.ForeignKeys[4].RefTable = VehiclesTable
	EquipmentTable.ForeignKeys[5].RefTable = CostsTable
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
	ClassSavingThrowsTable.ForeignKeys[0].RefTable = ClassesTable
	ClassSavingThrowsTable.ForeignKeys[1].RefTable = AbilityScoresTable
}
