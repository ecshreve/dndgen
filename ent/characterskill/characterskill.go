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
	// FieldModifier holds the string denoting the modifier field in the database.
	FieldModifier = "modifier"
	// FieldCharacterID holds the string denoting the character_id field in the database.
	FieldCharacterID = "character_id"
	// FieldSkillID holds the string denoting the skill_id field in the database.
	FieldSkillID = "skill_id"
	// EdgeCharacter holds the string denoting the character edge name in mutations.
	EdgeCharacter = "character"
	// EdgeSkill holds the string denoting the skill edge name in mutations.
	EdgeSkill = "skill"
	// Table holds the table name of the characterskill in the database.
	Table = "character_skills"
	// CharacterTable is the table that holds the character relation/edge.
	CharacterTable = "character_skills"
	// CharacterInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharacterInverseTable = "characters"
	// CharacterColumn is the table column denoting the character relation/edge.
	CharacterColumn = "character_id"
	// SkillTable is the table that holds the skill relation/edge.
	SkillTable = "character_skills"
	// SkillInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillInverseTable = "skills"
	// SkillColumn is the table column denoting the skill relation/edge.
	SkillColumn = "skill_id"
)

// Columns holds all SQL columns for characterskill fields.
var Columns = []string{
	FieldID,
	FieldProficient,
	FieldModifier,
	FieldCharacterID,
	FieldSkillID,
}

// ValidColumn reports if the column name is valid (part of the table columns).
func ValidColumn(column string) bool {
	for i := range Columns {
		if column == Columns[i] {
			return true
		}
	}
	return false
}

var (
	// DefaultProficient holds the default value on creation for the "proficient" field.
	DefaultProficient bool
	// DefaultModifier holds the default value on creation for the "modifier" field.
	DefaultModifier int
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

// ByModifier orders the results by the modifier field.
func ByModifier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldModifier, opts...).ToFunc()
}

// ByCharacterID orders the results by the character_id field.
func ByCharacterID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldCharacterID, opts...).ToFunc()
}

// BySkillID orders the results by the skill_id field.
func BySkillID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSkillID, opts...).ToFunc()
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
func newCharacterStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, CharacterTable, CharacterColumn),
	)
}
func newSkillStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, SkillTable, SkillColumn),
	)
}