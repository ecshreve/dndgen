// Code generated by ent, DO NOT EDIT.

package abilityscore

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the abilityscore type in the database.
	Label = "ability_score"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldFullName holds the string denoting the full_name field in the database.
	FieldFullName = "full_name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// EdgeClasses holds the string denoting the classes edge name in mutations.
	EdgeClasses = "classes"
	// EdgeSkills holds the string denoting the skills edge name in mutations.
	EdgeSkills = "skills"
	// Table holds the table name of the abilityscore in the database.
	Table = "ability_scores"
	// ClassesTable is the table that holds the classes relation/edge. The primary key declared below.
	ClassesTable = "class_saving_throws"
	// ClassesInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassesInverseTable = "classes"
	// SkillsTable is the table that holds the skills relation/edge.
	SkillsTable = "skills"
	// SkillsInverseTable is the table name for the Skill entity.
	// It exists in this package in order to avoid circular dependency with the "skill" package.
	SkillsInverseTable = "skills"
	// SkillsColumn is the table column denoting the skills relation/edge.
	SkillsColumn = "skill_ability_score"
)

// Columns holds all SQL columns for abilityscore fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldFullName,
	FieldDesc,
}

var (
	// ClassesPrimaryKey and ClassesColumn2 are the table columns denoting the
	// primary key for the classes relation (M2M).
	ClassesPrimaryKey = []string{"class_id", "ability_score_id"}
)

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
	// IndxValidator is a validator for the "indx" field. It is called by the builders before save.
	IndxValidator func(string) error
)

// OrderOption defines the ordering options for the AbilityScore queries.
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

// ByFullName orders the results by the full_name field.
func ByFullName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFullName, opts...).ToFunc()
}

// ByClassesCount orders the results by classes count.
func ByClassesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newClassesStep(), opts...)
	}
}

// ByClasses orders the results by classes terms.
func ByClasses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySkillsCount orders the results by skills count.
func BySkillsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSkillsStep(), opts...)
	}
}

// BySkills orders the results by skills terms.
func BySkills(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSkillsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newClassesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, ClassesTable, ClassesPrimaryKey...),
	)
}
func newSkillsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SkillsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, SkillsTable, SkillsColumn),
	)
}
