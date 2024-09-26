// Code generated by ent, DO NOT EDIT.

package character

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the character type in the database.
	Label = "character"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldAge holds the string denoting the age field in the database.
	FieldAge = "age"
	// FieldLevel holds the string denoting the level field in the database.
	FieldLevel = "level"
	// EdgeRace holds the string denoting the race edge name in mutations.
	EdgeRace = "race"
	// EdgeClass holds the string denoting the class edge name in mutations.
	EdgeClass = "class"
	// EdgeAlignment holds the string denoting the alignment edge name in mutations.
	EdgeAlignment = "alignment"
	// EdgeTraits holds the string denoting the traits edge name in mutations.
	EdgeTraits = "traits"
	// EdgeLanguages holds the string denoting the languages edge name in mutations.
	EdgeLanguages = "languages"
	// EdgeProficiencies holds the string denoting the proficiencies edge name in mutations.
	EdgeProficiencies = "proficiencies"
	// EdgeAbilityScores holds the string denoting the ability_scores edge name in mutations.
	EdgeAbilityScores = "ability_scores"
	// EdgeCharacterAbilityScores holds the string denoting the character_ability_scores edge name in mutations.
	EdgeCharacterAbilityScores = "character_ability_scores"
	// Table holds the table name of the character in the database.
	Table = "characters"
	// RaceTable is the table that holds the race relation/edge.
	RaceTable = "characters"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// RaceColumn is the table column denoting the race relation/edge.
	RaceColumn = "character_race"
	// ClassTable is the table that holds the class relation/edge.
	ClassTable = "characters"
	// ClassInverseTable is the table name for the Class entity.
	// It exists in this package in order to avoid circular dependency with the "class" package.
	ClassInverseTable = "classes"
	// ClassColumn is the table column denoting the class relation/edge.
	ClassColumn = "character_class"
	// AlignmentTable is the table that holds the alignment relation/edge.
	AlignmentTable = "characters"
	// AlignmentInverseTable is the table name for the Alignment entity.
	// It exists in this package in order to avoid circular dependency with the "alignment" package.
	AlignmentInverseTable = "alignments"
	// AlignmentColumn is the table column denoting the alignment relation/edge.
	AlignmentColumn = "character_alignment"
	// TraitsTable is the table that holds the traits relation/edge.
	TraitsTable = "traits"
	// TraitsInverseTable is the table name for the Trait entity.
	// It exists in this package in order to avoid circular dependency with the "trait" package.
	TraitsInverseTable = "traits"
	// TraitsColumn is the table column denoting the traits relation/edge.
	TraitsColumn = "character_traits"
	// LanguagesTable is the table that holds the languages relation/edge.
	LanguagesTable = "languages"
	// LanguagesInverseTable is the table name for the Language entity.
	// It exists in this package in order to avoid circular dependency with the "language" package.
	LanguagesInverseTable = "languages"
	// LanguagesColumn is the table column denoting the languages relation/edge.
	LanguagesColumn = "character_languages"
	// ProficienciesTable is the table that holds the proficiencies relation/edge.
	ProficienciesTable = "proficiencies"
	// ProficienciesInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	ProficienciesInverseTable = "proficiencies"
	// ProficienciesColumn is the table column denoting the proficiencies relation/edge.
	ProficienciesColumn = "character_proficiencies"
	// AbilityScoresTable is the table that holds the ability_scores relation/edge. The primary key declared below.
	AbilityScoresTable = "character_ability_scores"
	// AbilityScoresInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	AbilityScoresInverseTable = "ability_scores"
	// CharacterAbilityScoresTable is the table that holds the character_ability_scores relation/edge.
	CharacterAbilityScoresTable = "character_ability_scores"
	// CharacterAbilityScoresInverseTable is the table name for the CharacterAbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "characterabilityscore" package.
	CharacterAbilityScoresInverseTable = "character_ability_scores"
	// CharacterAbilityScoresColumn is the table column denoting the character_ability_scores relation/edge.
	CharacterAbilityScoresColumn = "character_id"
)

// Columns holds all SQL columns for character fields.
var Columns = []string{
	FieldID,
	FieldName,
	FieldAge,
	FieldLevel,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "characters"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"character_race",
	"character_class",
	"character_alignment",
}

var (
	// AbilityScoresPrimaryKey and AbilityScoresColumn2 are the table columns denoting the
	// primary key for the ability_scores relation (M2M).
	AbilityScoresPrimaryKey = []string{"character_id", "ability_score_id"}
)

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
	// NameValidator is a validator for the "name" field. It is called by the builders before save.
	NameValidator func(string) error
	// DefaultAge holds the default value on creation for the "age" field.
	DefaultAge int
	// AgeValidator is a validator for the "age" field. It is called by the builders before save.
	AgeValidator func(int) error
	// DefaultLevel holds the default value on creation for the "level" field.
	DefaultLevel int
	// LevelValidator is a validator for the "level" field. It is called by the builders before save.
	LevelValidator func(int) error
)

// OrderOption defines the ordering options for the Character queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByName orders the results by the name field.
func ByName(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldName, opts...).ToFunc()
}

// ByAge orders the results by the age field.
func ByAge(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldAge, opts...).ToFunc()
}

// ByLevel orders the results by the level field.
func ByLevel(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldLevel, opts...).ToFunc()
}

// ByRaceField orders the results by race field.
func ByRaceField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRaceStep(), sql.OrderByField(field, opts...))
	}
}

// ByClassField orders the results by class field.
func ByClassField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newClassStep(), sql.OrderByField(field, opts...))
	}
}

// ByAlignmentField orders the results by alignment field.
func ByAlignmentField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAlignmentStep(), sql.OrderByField(field, opts...))
	}
}

// ByTraitsCount orders the results by traits count.
func ByTraitsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newTraitsStep(), opts...)
	}
}

// ByTraits orders the results by traits terms.
func ByTraits(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newTraitsStep(), append([]sql.OrderTerm{term}, terms...)...)
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

// ByAbilityScoresCount orders the results by ability_scores count.
func ByAbilityScoresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newAbilityScoresStep(), opts...)
	}
}

// ByAbilityScores orders the results by ability_scores terms.
func ByAbilityScores(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAbilityScoresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCharacterAbilityScoresCount orders the results by character_ability_scores count.
func ByCharacterAbilityScoresCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCharacterAbilityScoresStep(), opts...)
	}
}

// ByCharacterAbilityScores orders the results by character_ability_scores terms.
func ByCharacterAbilityScores(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharacterAbilityScoresStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, RaceTable, RaceColumn),
	)
}
func newClassStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ClassInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, ClassTable, ClassColumn),
	)
}
func newAlignmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AlignmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AlignmentTable, AlignmentColumn),
	)
}
func newTraitsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(TraitsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, TraitsTable, TraitsColumn),
	)
}
func newLanguagesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(LanguagesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, LanguagesTable, LanguagesColumn),
	)
}
func newProficienciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficienciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProficienciesTable, ProficienciesColumn),
	)
}
func newAbilityScoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AbilityScoresInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, AbilityScoresTable, AbilityScoresPrimaryKey...),
	)
}
func newCharacterAbilityScoresStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharacterAbilityScoresInverseTable, CharacterAbilityScoresColumn),
		sqlgraph.Edge(sqlgraph.O2M, true, CharacterAbilityScoresTable, CharacterAbilityScoresColumn),
	)
}
