// Code generated by ent, DO NOT EDIT.

package prerequisite

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the prerequisite type in the database.
	Label = "prerequisite"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldPrerequisiteType holds the string denoting the prerequisite_type field in the database.
	FieldPrerequisiteType = "prerequisite_type"
	// FieldLevelValue holds the string denoting the level_value field in the database.
	FieldLevelValue = "level_value"
	// FieldFeatureValue holds the string denoting the feature_value field in the database.
	FieldFeatureValue = "feature_value"
	// FieldSpellValue holds the string denoting the spell_value field in the database.
	FieldSpellValue = "spell_value"
	// EdgeFeature holds the string denoting the feature edge name in mutations.
	EdgeFeature = "feature"
	// Table holds the table name of the prerequisite in the database.
	Table = "prerequisites"
	// FeatureTable is the table that holds the feature relation/edge.
	FeatureTable = "prerequisites"
	// FeatureInverseTable is the table name for the Feature entity.
	// It exists in this package in order to avoid circular dependency with the "feature" package.
	FeatureInverseTable = "features"
	// FeatureColumn is the table column denoting the feature relation/edge.
	FeatureColumn = "feature_prerequisites"
)

// Columns holds all SQL columns for prerequisite fields.
var Columns = []string{
	FieldID,
	FieldPrerequisiteType,
	FieldLevelValue,
	FieldFeatureValue,
	FieldSpellValue,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "prerequisites"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"feature_prerequisites",
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

// PrerequisiteType defines the type for the "prerequisite_type" enum field.
type PrerequisiteType string

// PrerequisiteType values.
const (
	PrerequisiteTypeLevel   PrerequisiteType = "level"
	PrerequisiteTypeSpell   PrerequisiteType = "spell"
	PrerequisiteTypeFeature PrerequisiteType = "feature"
)

func (pt PrerequisiteType) String() string {
	return string(pt)
}

// PrerequisiteTypeValidator is a validator for the "prerequisite_type" field enum values. It is called by the builders before save.
func PrerequisiteTypeValidator(pt PrerequisiteType) error {
	switch pt {
	case PrerequisiteTypeLevel, PrerequisiteTypeSpell, PrerequisiteTypeFeature:
		return nil
	default:
		return fmt.Errorf("prerequisite: invalid enum value for prerequisite_type field: %q", pt)
	}
}

// OrderOption defines the ordering options for the Prerequisite queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByPrerequisiteType orders the results by the prerequisite_type field.
func ByPrerequisiteType(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldPrerequisiteType, opts...).ToFunc()
}

// ByLevelValue orders the results by the level_value field.
func ByLevelValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevelValue, opts...).ToFunc()
}

// ByFeatureValue orders the results by the feature_value field.
func ByFeatureValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldFeatureValue, opts...).ToFunc()
}

// BySpellValue orders the results by the spell_value field.
func BySpellValue(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpellValue, opts...).ToFunc()
}

// ByFeatureField orders the results by feature field.
func ByFeatureField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newFeatureStep(), sql.OrderByField(field, opts...))
	}
}
func newFeatureStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(FeatureInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, true, FeatureTable, FeatureColumn),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e PrerequisiteType) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *PrerequisiteType) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = PrerequisiteType(str)
	if err := PrerequisiteTypeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid PrerequisiteType", str)
	}
	return nil
}
