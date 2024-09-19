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
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityScoresTable,
		AlignmentsTable,
		LanguagesTable,
		SkillsTable,
	}
)

func init() {
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
}
