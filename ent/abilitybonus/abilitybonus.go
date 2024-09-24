// Code generated by ent, DO NOT EDIT.

package abilitybonus

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the abilitybonus type in the database.
	Label = "ability_bonus"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldBonus holds the string denoting the bonus field in the database.
	FieldBonus = "bonus"
	// EdgeAbilityScore holds the string denoting the ability_score edge name in mutations.
	EdgeAbilityScore = "ability_score"
	// EdgeRace holds the string denoting the race edge name in mutations.
	EdgeRace = "race"
	// EdgeOptions holds the string denoting the options edge name in mutations.
	EdgeOptions = "options"
	// Table holds the table name of the abilitybonus in the database.
	Table = "ability_bonus"
	// AbilityScoreTable is the table that holds the ability_score relation/edge.
	AbilityScoreTable = "ability_bonus"
	// AbilityScoreInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	AbilityScoreInverseTable = "ability_scores"
	// AbilityScoreColumn is the table column denoting the ability_score relation/edge.
	AbilityScoreColumn = "ability_bonus_ability_score"
	// RaceTable is the table that holds the race relation/edge. The primary key declared below.
	RaceTable = "race_ability_bonuses"
	// RaceInverseTable is the table name for the Race entity.
	// It exists in this package in order to avoid circular dependency with the "race" package.
	RaceInverseTable = "races"
	// OptionsTable is the table that holds the options relation/edge. The primary key declared below.
	OptionsTable = "ability_bonus_choice_ability_bonuses"
	// OptionsInverseTable is the table name for the AbilityBonusChoice entity.
	// It exists in this package in order to avoid circular dependency with the "abilitybonuschoice" package.
	OptionsInverseTable = "ability_bonus_choices"
)

// Columns holds all SQL columns for abilitybonus fields.
var Columns = []string{
	FieldID,
	FieldBonus,
}

// ForeignKeys holds the SQL foreign-keys that are owned by the "ability_bonus"
// table and are not defined as standalone fields in the schema.
var ForeignKeys = []string{
	"ability_bonus_ability_score",
}

var (
	// RacePrimaryKey and RaceColumn2 are the table columns denoting the
	// primary key for the race relation (M2M).
	RacePrimaryKey = []string{"race_id", "ability_bonus_id"}
	// OptionsPrimaryKey and OptionsColumn2 are the table columns denoting the
	// primary key for the options relation (M2M).
	OptionsPrimaryKey = []string{"ability_bonus_choice_id", "ability_bonus_id"}
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
	// BonusValidator is a validator for the "bonus" field. It is called by the builders before save.
	BonusValidator func(int) error
)

// OrderOption defines the ordering options for the AbilityBonus queries.
type OrderOption func(*sql.Selector)

// ByID orders the results by the id field.
func ByID(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldID, opts...).ToFunc()
}

// ByBonus orders the results by the bonus field.
func ByBonus(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldBonus, opts...).ToFunc()
}

// ByAbilityScoreField orders the results by ability_score field.
func ByAbilityScoreField(field string, opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newAbilityScoreStep(), sql.OrderByField(field, opts...))
	}
}

// ByRaceCount orders the results by race count.
func ByRaceCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newRaceStep(), opts...)
	}
}

// ByRace orders the results by race terms.
func ByRace(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newRaceStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByOptionsCount orders the results by options count.
func ByOptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newOptionsStep(), opts...)
	}
}

// ByOptions orders the results by options terms.
func ByOptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newOptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newAbilityScoreStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(AbilityScoreInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2O, false, AbilityScoreTable, AbilityScoreColumn),
	)
}
func newRaceStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(RaceInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, RaceTable, RacePrimaryKey...),
	)
}
func newOptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(OptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, true, OptionsTable, OptionsPrimaryKey...),
	)
}
