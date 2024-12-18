// Code generated by ent, DO NOT EDIT.

package characterabilityscore

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the characterabilityscore type in the database.
	Label = "character_ability_score"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldScore holds the string denoting the score field in the database.
	FieldScore = "score"
	// FieldModifier holds the string denoting the modifier field in the database.
	FieldModifier = "modifier"
	// EdgeCharacter holds the string denoting the character edge name in mutations.
	EdgeCharacter = "character"
	// EdgeAbilityScore holds the string denoting the ability_score edge name in mutations.
	EdgeAbilityScore = "ability_score"
	// EdgeCharacterSkills holds the string denoting the character_skills edge name in mutations.
	EdgeCharacterSkills = "character_skills"
	// Table holds the table name of the characterabilityscore in the database.
	Table = "character_ability_scores"
	// CharacterTable is the table that holds the character relation/edge.
	CharacterTable = "character_ability_scores"
	// CharacterInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharacterInverseTable = "characters"
	// CharacterColumn is the table column denoting the character relation/edge.
	CharacterColumn = "character_character_ability_scores"
	// AbilityScoreTable is the table that holds the ability_score relation/edge.
	AbilityScoreTable = "character_ability_scores"
	// AbilityScoreInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	AbilityScoreInverseTable = "ability_scores"
	// AbilityScoreColumn is the table column denoting the ability_score relation/edge.
	AbilityScoreColumn = "character_ability_score_ability_score"
	// CharacterSkillsTable is the table that holds the character_skills relation/edge.
	CharacterSkillsTable = "character_skills"
	// CharacterSkillsInverseTable is the table name for the CharacterSkill entity.
	// It exists in this package in order to avoid circular dependency with the "characterskill" package.
	CharacterSkillsInverseTable = "character_skills"
	// CharacterSkillsColumn is the table column denoting the character_skills relation/edge.
	CharacterSkillsColumn = "character_skill_character_ability_score"
)

// Columns holds all SQL columns for characterabilityscore fields.
var Columns = []string{
	FieldID,
	FieldScore,
	FieldModifier,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "character_ability_scores"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"character_character_ability_scores",
	"character_ability_score_ability_score",
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
	// ScoreValidator is a validator for the "score" field. It is called by the builders before save.
	ScoreValidator func(int) error
)

// OrderOption defines the ordering options for the CharacterAbilityScore queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByScore orders the results by the score field.
func ByScore(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScore, opts...).ToFunc()
}

// ByModifier orders the results by the modifier field.
func ByModifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifier, opts...).ToFunc()
}

// ByCharacterField orders the results by character field.
func ByCharacterField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterStep(), sql.OrderByField(field, opts...))
	}
}

// ByAbilityScoreField orders the results by ability_score field.
func ByAbilityScoreField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAbilityScoreStep(), sql.OrderByField(field, opts...))
	}
}

// ByCharacterSkillsCount orders the results by character_skills count.
func ByCharacterSkillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCharacterSkillsStep(), opts...)
	}
}

// ByCharacterSkills orders the results by character_skills terms.
func ByCharacterSkills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterSkillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newCharacterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, CharacterTable, CharacterColumn),
	)
}
func newAbilityScoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AbilityScoreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AbilityScoreTable, AbilityScoreColumn),
	)
}
func newCharacterSkillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterSkillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CharacterSkillsTable, CharacterSkillsColumn),
	)
}
