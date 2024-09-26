// Code generated by ent, DO NOT EDIT.

package class

import (
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
)

const (
	// Label holds the string label denoting the class type in the database.
	Label = "class"
	// FieldID holds the string denoting the id field in the database.
	FieldID = "id"
	// FieldIndx holds the string denoting the indx field in the database.
	FieldIndx = "indx"
	// FieldName holds the string denoting the name field in the database.
	FieldName = "name"
	// FieldHitDie holds the string denoting the hit_die field in the database.
	FieldHitDie = "hit_die"
	// EdgeProficiencies holds the string denoting the proficiencies edge name in mutations.
	EdgeProficiencies = "proficiencies"
	// EdgeProficiencyOptions holds the string denoting the proficiency_options edge name in mutations.
	EdgeProficiencyOptions = "proficiency_options"
	// EdgeStartingEquipment holds the string denoting the starting_equipment edge name in mutations.
	EdgeStartingEquipment = "starting_equipment"
	// EdgeSavingThrows holds the string denoting the saving_throws edge name in mutations.
	EdgeSavingThrows = "saving_throws"
	// EdgeCharacters holds the string denoting the characters edge name in mutations.
	EdgeCharacters = "characters"
	// Table holds the table name of the class in the database.
	Table = "classes"
	// ProficienciesTable is the table that holds the proficiencies relation/edge. The primary key declared below.
	ProficienciesTable = "class_proficiencies"
	// ProficienciesInverseTable is the table name for the Proficiency entity.
	// It exists in this package in order to avoid circular dependency with the "proficiency" package.
	ProficienciesInverseTable = "proficiencies"
	// ProficiencyOptionsTable is the table that holds the proficiency_options relation/edge.
	ProficiencyOptionsTable = "proficiency_choices"
	// ProficiencyOptionsInverseTable is the table name for the ProficiencyChoice entity.
	// It exists in this package in order to avoid circular dependency with the "proficiencychoice" package.
	ProficiencyOptionsInverseTable = "proficiency_choices"
	// ProficiencyOptionsColumn is the table column denoting the proficiency_options relation/edge.
	ProficiencyOptionsColumn = "class_proficiency_options"
	// StartingEquipmentTable is the table that holds the starting_equipment relation/edge. The primary key declared below.
	StartingEquipmentTable = "class_starting_equipment"
	// StartingEquipmentInverseTable is the table name for the EquipmentEntry entity.
	// It exists in this package in order to avoid circular dependency with the "equipmententry" package.
	StartingEquipmentInverseTable = "equipment_entries"
	// SavingThrowsTable is the table that holds the saving_throws relation/edge. The primary key declared below.
	SavingThrowsTable = "class_saving_throws"
	// SavingThrowsInverseTable is the table name for the AbilityScore entity.
	// It exists in this package in order to avoid circular dependency with the "abilityscore" package.
	SavingThrowsInverseTable = "ability_scores"
	// CharactersTable is the table that holds the characters relation/edge.
	CharactersTable = "characters"
	// CharactersInverseTable is the table name for the Character entity.
	// It exists in this package in order to avoid circular dependency with the "character" package.
	CharactersInverseTable = "characters"
	// CharactersColumn is the table column denoting the characters relation/edge.
	CharactersColumn = "character_class"
)

// Columns holds all SQL columns for class fields.
var Columns = []string{
	FieldID,
	FieldIndx,
	FieldName,
	FieldHitDie,
}

var (
	// ProficienciesPrimaryKey and ProficienciesColumn2 are the table columns denoting the
	// primary key for the proficiencies relation (M2M).
	ProficienciesPrimaryKey = []string{"class_id", "proficiency_id"}
	// StartingEquipmentPrimaryKey and StartingEquipmentColumn2 are the table columns denoting the
	// primary key for the starting_equipment relation (M2M).
	StartingEquipmentPrimaryKey = []string{"class_id", "equipment_entry_id"}
	// SavingThrowsPrimaryKey and SavingThrowsColumn2 are the table columns denoting the
	// primary key for the saving_throws relation (M2M).
	SavingThrowsPrimaryKey = []string{"class_id", "ability_score_id"}
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
	// HitDieValidator is a validator for the "hit_die" field. It is called by the builders before save.
	HitDieValidator func(int) error
)

// OrderOption defines the ordering options for the Class queries.
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

// ByHitDie orders the results by the hit_die field.
func ByHitDie(opts ...sql.OrderTermOption) OrderOption {
	return sql.OrderByField(FieldHitDie, opts...).ToFunc()
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

// ByProficiencyOptionsCount orders the results by proficiency_options count.
func ByProficiencyOptionsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newProficiencyOptionsStep(), opts...)
	}
}

// ByProficiencyOptions orders the results by proficiency_options terms.
func ByProficiencyOptions(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newProficiencyOptionsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByStartingEquipmentCount orders the results by starting_equipment count.
func ByStartingEquipmentCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newStartingEquipmentStep(), opts...)
	}
}

// ByStartingEquipment orders the results by starting_equipment terms.
func ByStartingEquipment(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newStartingEquipmentStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// BySavingThrowsCount orders the results by saving_throws count.
func BySavingThrowsCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newSavingThrowsStep(), opts...)
	}
}

// BySavingThrows orders the results by saving_throws terms.
func BySavingThrows(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newSavingThrowsStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}

// ByCharactersCount orders the results by characters count.
func ByCharactersCount(opts ...sql.OrderTermOption) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborsCount(s, newCharactersStep(), opts...)
	}
}

// ByCharacters orders the results by characters terms.
func ByCharacters(term sql.OrderTerm, terms ...sql.OrderTerm) OrderOption {
	return func(s *sql.Selector) {
		sqlgraph.OrderByNeighborTerms(s, newCharactersStep(), append([]sql.OrderTerm{term}, terms...)...)
	}
}
func newProficienciesStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficienciesInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, ProficienciesTable, ProficienciesPrimaryKey...),
	)
}
func newProficiencyOptionsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(ProficiencyOptionsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, false, ProficiencyOptionsTable, ProficiencyOptionsColumn),
	)
}
func newStartingEquipmentStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(StartingEquipmentInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, StartingEquipmentTable, StartingEquipmentPrimaryKey...),
	)
}
func newSavingThrowsStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(SavingThrowsInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.M2M, false, SavingThrowsTable, SavingThrowsPrimaryKey...),
	)
}
func newCharactersStep() *sqlgraph.Step {
	return sqlgraph.NewStep(
		sqlgraph.From(Table, FieldID),
		sqlgraph.To(CharactersInverseTable, FieldID),
		sqlgraph.Edge(sqlgraph.O2M, true, CharactersTable, CharactersColumn),
	)
}
