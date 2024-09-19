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
		SkillsTable,
	}
)

func init() {
	SkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
}
