// Code generated by ent, DO NOT EDIT.

package language

import (
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the language type in the database.
	Label = "language"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldDesc holds the string denoting the desc field in the database.
	FieldDesc = "desc"
	// FieldTier holds the string denoting the tier field in the database.
	FieldTier = "tier"
	// FieldScript holds the string denoting the script field in the database.
	FieldScript = "script"
	// EdgeSpeakers holds the string denoting the speakers edge name in mutations.
	EdgeSpeakers = "speakers"
	// Table holds the table name of the language in the database.
	Table = "languages"
	// SpeakersTable is the table that holds the speakers relation/edge. The primary key declared below.
	SpeakersTable = "race_languages"
	// SpeakersInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	SpeakersInverseTable = "races"
)

// Columns holds all SQL columns for language fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldDesc,
	FieldTier,
	FieldScript,
}

var (
	// SpeakersPrimaryKey and SpeakersColumn2 are the table columns denoting the
	// primary key for the speakers relation (M2M).
	SpeakersPrimaryKey = []string{"race_id", "language_id"}
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

// Tier defines the type for the "tier" enum field.
type Tier string

// TierStandard is the default value of the Tier enum.
const DefaultTier = TierStandard

// Tier values.
const (
	TierStandard Tier = "standard"
	TierExotic   Tier = "exotic"
)

func (t Tier) String() string {
	return string(t)
}

// TierValidator is a validator for the "tier" field enum values. It is called by the builders before save.
func TierValidator(t Tier) error {
	switch t {
	case TierStandard, TierExotic:
		return nil
	default:
		return fmt.Errorf("language: invalid enum value for tier field: %q", t)
	}
}

// OrderOption defines the ordering options for the Language queries.
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

// ByDesc orders the results by the desc field.
func ByDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldDesc, opts...).ToFunc()
}

// ByTier orders the results by the tier field.
func ByTier(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldTier, opts...).ToFunc()
}

// ByScript orders the results by the script field.
func ByScript(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldScript, opts...).ToFunc()
}

// BySpeakersCount orders the results by speakers count.
func BySpeakersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSpeakersStep(), opts...)
	}
}

// BySpeakers orders the results by speakers terms.
func BySpeakers(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSpeakersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newSpeakersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SpeakersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, SpeakersTable, SpeakersPrimaryKey...),
	)
}