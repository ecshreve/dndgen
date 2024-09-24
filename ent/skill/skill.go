// Code generated by ent, DO NOT EDIT.

package skill

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the skill type in the database.
	Label = "skill"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeAbilityScore holds the string denoting the ability_score edge name in mutations.
	EdgeAbilityScore = "ability_score"
	// EdgeProficiencies holds the string denoting the proficiencies edge name in mutations.
	EdgeProficiencies = "proficiencies"
	// Table holds the table name of the skill in the database.
	Table = "skills"
	// AbilityScoreTable is the table that holds the ability_score relation/edge.
	AbilityScoreTable = "skills"
	// AbilityScoreInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	AbilityScoreInverseTable = "ability_scores"
	// AbilityScoreColumn is the table column denoting the ability_score relation/edge.
	AbilityScoreColumn = "ability_score_skills"
	// ProficienciesTable is the table that holds the proficiencies relation/edge.
	ProficienciesTable = "proficiencies"
	// ProficienciesInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	ProficienciesInverseTable = "proficiencies"
	// ProficienciesColumn is the table column denoting the proficiencies relation/edge.
	ProficienciesColumn = "proficiency_skill"
)

// Columns holds all SQL columns for skill fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "skills"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"ability_score_skills",
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
	// IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	IndxValidator func(string) error
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
)

// OrderOption defines the ordering options for the Skill queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByIndx orders the results by the indx field.
func ByIndx(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldIndx, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAbilityScoreField orders the results by ability_score field.
func ByAbilityScoreField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAbilityScoreStep(), sql.OrderByField(field, opts...))
	}
}

// ByProficienciesCount orders the results by proficiencies count.
func ByProficienciesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProficienciesStep(), opts...)
	}
}

// ByProficiencies orders the results by proficiencies terms.
func ByProficiencies(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProficienciesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAbilityScoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AbilityScoreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, AbilityScoreTable, AbilityScoreColumn),
	)
}
func newProficienciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficienciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, ProficienciesTable, ProficienciesColumn),
	)
}
