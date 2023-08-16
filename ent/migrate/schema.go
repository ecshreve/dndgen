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
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
		{Name: "abbr", Type: field.TypeString, Unique: true, Size: 3},
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
		{Name: "indx", Type: field.TypeString},
		{Name: "name", Type: field.TypeString},
		{Name: "desc", Type: field.TypeString},
	}
	// SkillsTable holds the schema information for the "skills" table.
	SkillsTable = &schema.Table{
		Name:       "skills",
		Columns:    SkillsColumns,
		PrimaryKey: []*schema.Column{SkillsColumns[0]},
	}
	// AbilityScoreSkillsColumns holds the columns for the "ability_score_skills" table.
	AbilityScoreSkillsColumns = []*schema.Column{
		{Name: "ability_score_id", Type: field.TypeInt},
		{Name: "skill_id", Type: field.TypeInt},
	}
	// AbilityScoreSkillsTable holds the schema information for the "ability_score_skills" table.
	AbilityScoreSkillsTable = &schema.Table{
		Name:       "ability_score_skills",
		Columns:    AbilityScoreSkillsColumns,
		PrimaryKey: []*schema.Column{AbilityScoreSkillsColumns[0], AbilityScoreSkillsColumns[1]},
		ForeignKeys: []*schema.ForeignKey{
			{
				Symbol:     "ability_score_skills_ability_score_id",
				Columns:    []*schema.Column{AbilityScoreSkillsColumns[0]},
				RefColumns: []*schema.Column{AbilityScoresColumns[0]},
				OnDelete:   schema.Cascade,
			},
			{
				Symbol:     "ability_score_skills_skill_id",
				Columns:    []*schema.Column{AbilityScoreSkillsColumns[1]},
				RefColumns: []*schema.Column{SkillsColumns[0]},
				OnDelete:   schema.Cascade,
			},
		},
	}
	// Tables holds all the tables in the schema.
	Tables = []*schema.Table{
		AbilityScoresTable,
		SkillsTable,
		AbilityScoreSkillsTable,
	}
)

func init() {
	AbilityScoreSkillsTable.ForeignKeys[0].RefTable = AbilityScoresTable
	AbilityScoreSkillsTable.ForeignKeys[1].RefTable = SkillsTable
}
