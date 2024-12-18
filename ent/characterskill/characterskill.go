// Code generated by ent, DO NOT EDIT.

package characterskill

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the characterskill type in the database.
	Label = "character_skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldProficient holds the string denoting the proficient field in the database.
	FieldProficient = "proficient"
	// EdgeCharacter holds the string denoting the character edge name in mutations.
	EdgeCharacter = "character"
	// EdgeSkill holds the string denoting the skill edge name in mutations.
	EdgeSkill = "skill"
	// EdgeCharacterAbilityScore holds the string denoting the character_ability_score edge name in mutations.
	EdgeCharacterAbilityScore = "character_ability_score"
	// EdgeCharacterProficiency holds the string denoting the character_proficiency edge name in mutations.
	EdgeCharacterProficiency = "character_proficiency"
	// Table holds the table name of the characterskill in the database.
	Table = "character_skills"
	// CharacterTable is the table that holds the character relation/edge.
	CharacterTable = "character_skills"
	// CharacterInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharacterInverseTable = "characters"
	// CharacterColumn is the table column denoting the character relation/edge.
	CharacterColumn = "character_character_skills"
	// SkillTable is the table that holds the skill relation/edge.
	SkillTable = "character_skills"
	// SkillInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillInverseTable = "skills"
	// SkillColumn is the table column denoting the skill relation/edge.
	SkillColumn = "character_skill_skill"
	// CharacterAbilityScoreTable is the table that holds the character_ability_score relation/edge.
	CharacterAbilityScoreTable = "character_skills"
	// CharacterAbilityScoreInverseTable is the table name for the CharacterAbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "characterabilityscore" package.
	CharacterAbilityScoreInverseTable = "character_ability_scores"
	// CharacterAbilityScoreColumn is the table column denoting the character_ability_score relation/edge.
	CharacterAbilityScoreColumn = "character_skill_character_ability_score"
	// CharacterProficiencyTable is the table that holds the character_proficiency relation/edge.
	CharacterProficiencyTable = "character_proficiencies"
	// CharacterProficiencyInverseTable is the table name for the CharacterProficiency entity.
	// It exists in this package in order to avoid circular dependency with the "characterproficiency" package.
	CharacterProficiencyInverseTable = "character_proficiencies"
	// CharacterProficiencyColumn is the table column denoting the character_proficiency relation/edge.
	CharacterProficiencyColumn = "character_skill_character_proficiency"
)

// Columns holds all SQL columns for characterskill fields.
var Columns = []string{
	FieldID,
	FieldProficient,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "character_skills"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"character_character_skills",
	"character_skill_skill",
	"character_skill_character_ability_score",
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	for i := range ForeignKeys {
		if column == ForeignKeys[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultProficient holds the default value on creation for the "proficient" field.
	DefaultProficient bool
)

// OrderOption defines the ordering options for the CharacterSkill queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByProficient orders the results by the proficient field.
func ByProficient(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldProficient, opts...).ToFunc()
}

// ByCharacterField orders the results by character field.
func ByCharacterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterStep(), sql.OrderByField(field, opts...))
	}
}

// BySkillField orders the results by skill field.
func BySkillField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSkillStep(), sql.OrderByField(field, opts...))
	}
}

// ByCharacterAbilityScoreField orders the results by character_ability_score field.
func ByCharacterAbilityScoreField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterAbilityScoreStep(), sql.OrderByField(field, opts...))
	}
}

// ByCharacterProficiencyField orders the results by character_proficiency field.
func ByCharacterProficiencyField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterProficiencyStep(), sql.OrderByField(field, opts...))
	}
}
func newCharacterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CharacterTable, CharacterColumn),
	)
}
func newSkillStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SkillTable, SkillColumn),
	)
}
func newCharacterAbilityScoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterAbilityScoreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CharacterAbilityScoreTable, CharacterAbilityScoreColumn),
	)
}
func newCharacterProficiencyStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterProficiencyInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2O, false, CharacterProficiencyTable, CharacterProficiencyColumn),
	)
}
