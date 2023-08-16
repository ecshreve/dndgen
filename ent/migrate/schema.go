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
	}
	// AbilityBonusTable holds the schema information for the "ability_bonus" table.
	AbilityBonusTable = &schema.Table{
		Name:       "ability_bonus",
		Columns:    AbilityBonusColumns,
		PrimaryKey: []*schema.Column{AbilityBonusColumns[0]},
	}
	// AbilityScoresColumns holds the columns for the "ability_scores" table.
	AbilityScoresColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
		{Name: "abbr", Type: field.TypeString, Unique: true, Size: 3},
		{Name: "ability_bonus_ability_score", Type: field.TypeInt, Nullable: true},
		{Name: "prerequisite_ability_score", Type: field.TypeInt, Nullable: true},
	}
	// AbilityScoresTable holds the schema information for the "ability_scores" table.
	AbilityScoresTable = &schema.Table{
		Name:       "ability_scores",
		Columns:    AbilityScoresColumns,
		PrimaryKey: []*schema.Column{AbilityScoresColumns[0]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ability_scores_ability_bonus_ability_score",
				Columns:    []*schema.Column{AbilityScoresColumns[5]},
				RefColumns: []*schema.Column{AbilityBonusColumns[0]},
				OnDelete:   schema.SetNull,
			},
			{
				Symbol:     "ability_scores_prerequisites_ability_score",
				Columns:    []*schema.Column{AbilityScoresColumns[6]},
				RefColumns: []*schema.Column{PrerequisitesColumns[0]},
				OnDelete:   schema.SetNull,
			},
		},
	}
	// ConditionsColumns holds the columns for the "conditions" table.
	ConditionsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
	}
	// ConditionsTable holds the schema information for the "conditions" table.
	ConditionsTable = &schema.Table{
		Name:       "conditions",
		Columns:    ConditionsColumns,
		PrimaryKey: []*schema.Column{ConditionsColumns[0]},
	}
	// DamageTypesColumns holds the columns for the "damage_types" table.
	DamageTypesColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
	}
	// DamageTypesTable holds the schema information for the "damage_types" table.
	DamageTypesTable = &schema.Table{
		Name:       "damage_types",
		Columns:    DamageTypesColumns,
		PrimaryKey: []*schema.Column{DamageTypesColumns[0]},
	}
	// MagicSchoolsColumns holds the columns for the "magic_schools" table.
	MagicSchoolsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
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
		{Name: "minimum", Type: field.TypeInt},
	}
	// PrerequisitesTable holds the schema information for the "prerequisites" table.
	PrerequisitesTable = &schema.Table{
		Name:       "prerequisites",
		Columns:    PrerequisitesColumns,
		PrimaryKey: []*schema.Column{PrerequisitesColumns[0]},
	}
	// SkillsColumns holds the columns for the "skills" table.
	SkillsColumns = []*schema.Column{
		{Name: "id", Type: field.TypeInt, Increment: true},
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityBonusTable,
		AbilityScoresTable,
		ConditionsTable,
		DamageTypesTable,
		MagicSchoolsTable,
		PrerequisitesTable,
		SkillsTable,
	}
)

func init() {
	AbilityScoresTable.ForeignKeys[0].RefTable = AbilityBonusTable
	AbilityScoresTable.ForeignKeys[1].RefTable = PrerequisitesTable
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
}
