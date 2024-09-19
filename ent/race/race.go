// Code generated by ent, DO NOT EDIT.

package race

import (
	"fmt"
	"io"
	"strconv"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the race type in the database.
	Label = "race"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldSpeed holds the string denoting the speed field in the database.
	FieldSpeed = "speed"
	// FieldSize holds the string denoting the size field in the database.
	FieldSize = "size"
	// FieldSizeDesc holds the string denoting the size_desc field in the database.
	FieldSizeDesc = "size_desc"
	// FieldAlignmentDesc holds the string denoting the alignment_desc field in the database.
	FieldAlignmentDesc = "alignment_desc"
	// FieldAgeDesc holds the string denoting the age_desc field in the database.
	FieldAgeDesc = "age_desc"
	// FieldLanguageDesc holds the string denoting the language_desc field in the database.
	FieldLanguageDesc = "language_desc"
	// EdgeAbilityBonuses holds the string denoting the ability_bonuses edge name in mutations.
	EdgeAbilityBonuses = "ability_bonuses"
	// EdgeLanguages holds the string denoting the languages edge name in mutations.
	EdgeLanguages = "languages"
	// Table holds the table name of the race in the database.
	Table = "races"
	// AbilityBonusesTable is the table that holds the ability_bonuses relation/edge.
	AbilityBonusesTable = "ability_bonus"
	// AbilityBonusesInverseTable is the table name for the AbilityBonus entity.
	// It exists in this package in order to avoid circular dependency with the "abilitybonus" package.
	AbilityBonusesInverseTable = "ability_bonus"
	// AbilityBonusesColumn is the table column denoting the ability_bonuses relation/edge.
	AbilityBonusesColumn = "race_id"
	// LanguagesTable is the table that holds the languages relation/edge. The primary key declared below.
	LanguagesTable = "race_languages"
	// LanguagesInverseTable is the table name for the Language entity.
	// It exists in this package in order to avoid circular dependency with the "language" package.
	LanguagesInverseTable = "languages"
)

// Columns holds all SQL columns for race fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldSpeed,
	FieldSize,
	FieldSizeDesc,
	FieldAlignmentDesc,
	FieldAgeDesc,
	FieldLanguageDesc,
}

var (
	// LanguagesPrimaryKey and LanguagesColumn2 are the table columns denoting the
	// primary key for the languages relation (M2M).
	LanguagesPrimaryKey = []string{"race_id", "language_id"}
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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// SpeedValidator is a validator for the "speed" field. It is called by the builders before save.
	SpeedValidator func(int) error
)

// Size defines the type for the "size" enum field.
type Size string

// SizeMedium is the default value of the Size enum.
const DefaultSize = SizeMedium

// Size values.
const (
	SizeSmall  Size = "Small"
	SizeMedium Size = "Medium"
	SizeLarge  Size = "Large"
)

func (s Size) String() string {
	return string(s)
}

// SizeValidator is a validator for the "size" field enum values. It is called by the builders before save.
func SizeValidator(s Size) error {
	switch s {
	case SizeSmall, SizeMedium, SizeLarge:
		return nil
	default:
		return fmt.Errorf("race: invalid enum value for size field: %q", s)
	}
}

// OrderOption defines the ordering options for the Race queries.
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

// BySpeed orders the results by the speed field.
func BySpeed(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSpeed, opts...).ToFunc()
}

// BySize orders the results by the size field.
func BySize(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSize, opts...).ToFunc()
}

// BySizeDesc orders the results by the size_desc field.
func BySizeDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldSizeDesc, opts...).ToFunc()
}

// ByAlignmentDesc orders the results by the alignment_desc field.
func ByAlignmentDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAlignmentDesc, opts...).ToFunc()
}

// ByAgeDesc orders the results by the age_desc field.
func ByAgeDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAgeDesc, opts...).ToFunc()
}

// ByLanguageDesc orders the results by the language_desc field.
func ByLanguageDesc(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLanguageDesc, opts...).ToFunc()
}

// ByAbilityBonusesCount orders the results by ability_bonuses count.
func ByAbilityBonusesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAbilityBonusesStep(), opts...)
	}
}

// ByAbilityBonuses orders the results by ability_bonuses terms.
func ByAbilityBonuses(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAbilityBonusesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByLanguagesCount orders the results by languages count.
func ByLanguagesCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newLanguagesStep(), opts...)
	}
}

// ByLanguages orders the results by languages terms.
func ByLanguages(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newLanguagesStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAbilityBonusesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AbilityBonusesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, AbilityBonusesTable, AbilityBonusesColumn),
	)
}
func newLanguagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LanguagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, LanguagesTable, LanguagesPrimaryKey...),
	)
}

// MarshalGQL implements graphql.Marshaler interface.
func (e Size) MarshalGQL(w io.Writer) {
	io.WriteString(w, strconv.Quote(e.String()))
}

// UnmarshalGQL implements graphql.Unmarshaler interface.
func (e *Size) UnmarshalGQL(val interface{}) error {
	str, ok := val.(string)
	if !ok {
		return fmt.Errorf("enum %T must be a string", val)
	}
	*e = Size(str)
	if err := SizeValidator(*e); err != nil {
		return fmt.Errorf("%s is not a valid Size", str)
	}
	return nil
}
