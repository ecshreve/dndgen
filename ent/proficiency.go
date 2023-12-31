// Code generated by ent, DO NOT EDIT.

package ent

import (
	"encoding/json"
	"fmt"
	"strings"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"github.com/ecshreve/dndgen/ent/proficiency"
)

// Proficiency is the model entity for the Proficiency schema.
type Proficiency struct {
	config `json:"-"`
	// ID of the ent.
	ID int `json:"id,omitempty"`
	// Indx holds the value of the "indx" field.
	Indx string `json:"index"`
	// Name holds the value of the "name" field.
	Name string `json:"name,omitempty"`
	// ProficiencyCategory holds the value of the "proficiency_category" field.
	ProficiencyCategory string `json:"type"`
	// Edges holds the relations/edges for other nodes in the graph.
	// The values are being populated by the ProficiencyQuery when eager-loading is set.
	Edges        ProficiencyEdges `json:"-"`
	selectValues sql.SelectValues
}

// ProficiencyEdges holds the relations/edges for other nodes in the graph.
type ProficiencyEdges struct {
	// Classes holds the value of the classes edge.
	Classes []*Class `json:"classes,omitempty"`
	// Races holds the value of the races edge.
	Races []*Race `json:"races,omitempty"`
	// Subraces holds the value of the subraces edge.
	Subraces []*Subrace `json:"subraces,omitempty"`
	// Choice holds the value of the choice edge.
	Choice []*ProficiencyChoice `json:"choice,omitempty"`
	// Skill holds the value of the skill edge.
	Skill []*Skill `json:"skill,omitempty"`
	// Equipment holds the value of the equipment edge.
	Equipment []*Equipment `json:"equipment,omitempty"`
	// EquipmentCategory holds the value of the equipment_category edge.
	EquipmentCategory []*EquipmentCategory `json:"equipment_category,omitempty"`
	// SavingThrow holds the value of the saving_throw edge.
	SavingThrow []*AbilityScore `json:"saving_throw,omitempty"`
	// loadedTypes holds the information for reporting if a
	// type was loaded (or requested) in eager-loading or not.
	loadedTypes [8]bool
	// totalCount holds the count of the edges above.
	totalCount [8]map[string]int

	namedClasses           map[string][]*Class
	namedRaces             map[string][]*Race
	namedSubraces          map[string][]*Subrace
	namedChoice            map[string][]*ProficiencyChoice
	namedSkill             map[string][]*Skill
	namedEquipment         map[string][]*Equipment
	namedEquipmentCategory map[string][]*EquipmentCategory
	namedSavingThrow       map[string][]*AbilityScore
}

// ClassesOrErr returns the Classes value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) ClassesOrErr() ([]*Class, error) {
	if e.loadedTypes[0] {
		return e.Classes, nil
	}
	return nil, &NotLoadedError{edge: "classes"}
}

// RacesOrErr returns the Races value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) RacesOrErr() ([]*Race, error) {
	if e.loadedTypes[1] {
		return e.Races, nil
	}
	return nil, &NotLoadedError{edge: "races"}
}

// SubracesOrErr returns the Subraces value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) SubracesOrErr() ([]*Subrace, error) {
	if e.loadedTypes[2] {
		return e.Subraces, nil
	}
	return nil, &NotLoadedError{edge: "subraces"}
}

// ChoiceOrErr returns the Choice value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) ChoiceOrErr() ([]*ProficiencyChoice, error) {
	if e.loadedTypes[3] {
		return e.Choice, nil
	}
	return nil, &NotLoadedError{edge: "choice"}
}

// SkillOrErr returns the Skill value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) SkillOrErr() ([]*Skill, error) {
	if e.loadedTypes[4] {
		return e.Skill, nil
	}
	return nil, &NotLoadedError{edge: "skill"}
}

// EquipmentOrErr returns the Equipment value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) EquipmentOrErr() ([]*Equipment, error) {
	if e.loadedTypes[5] {
		return e.Equipment, nil
	}
	return nil, &NotLoadedError{edge: "equipment"}
}

// EquipmentCategoryOrErr returns the EquipmentCategory value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) EquipmentCategoryOrErr() ([]*EquipmentCategory, error) {
	if e.loadedTypes[6] {
		return e.EquipmentCategory, nil
	}
	return nil, &NotLoadedError{edge: "equipment_category"}
}

// SavingThrowOrErr returns the SavingThrow value or an error if the edge
// was not loaded in eager-loading.
func (e ProficiencyEdges) SavingThrowOrErr() ([]*AbilityScore, error) {
	if e.loadedTypes[7] {
		return e.SavingThrow, nil
	}
	return nil, &NotLoadedError{edge: "saving_throw"}
}

// scanValues returns the types for scanning values from sql.Rows.
func (*Proficiency) scanValues(columns []string) ([]any, error) {
	values := make([]any, len(columns))
	for i := range columns {
		switch columns[i] {
		case proficiency.FieldID:
			values[i] = new(sql.NullInt64)
		case proficiency.FieldIndx, proficiency.FieldName, proficiency.FieldProficiencyCategory:
			values[i] = new(sql.NullString)
		default:
			values[i] = new(sql.UnknownType)
		}
	}
	return values, nil
}

// assignValues assigns the values that were returned from sql.Rows (after scanning)
// to the Proficiency fields.
func (pr *Proficiency) assignValues(columns []string, values []any) error {
	if m, n := len(values), len(columns); m < n {
		return fmt.Errorf("mismatch number of scan values: %d != %d", m, n)
	}
	for i := range columns {
		switch columns[i] {
		case proficiency.FieldID:
			value, ok := values[i].(*sql.NullInt64)
			if !ok {
				return fmt.Errorf("unexpected type %T for field id", value)
			}
			pr.ID = int(value.Int64)
		case proficiency.FieldIndx:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field indx", values[i])
			} else if value.Valid {
				pr.Indx = value.String
			}
		case proficiency.FieldName:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field name", values[i])
			} else if value.Valid {
				pr.Name = value.String
			}
		case proficiency.FieldProficiencyCategory:
			if value, ok := values[i].(*sql.NullString); !ok {
				return fmt.Errorf("unexpected type %T for field proficiency_category", values[i])
			} else if value.Valid {
				pr.ProficiencyCategory = value.String
			}
		default:
			pr.selectValues.Set(columns[i], values[i])
		}
	}
	return nil
}

// Value returns the ent.Value that was dynamically selected and assigned to the Proficiency.
// This includes values selected through modifiers, order, etc.
func (pr *Proficiency) Value(name string) (ent.Value, error) {
	return pr.selectValues.Get(name)
}

// QueryClasses queries the "classes" edge of the Proficiency entity.
func (pr *Proficiency) QueryClasses() *ClassQuery {
	return NewProficiencyClient(pr.config).QueryClasses(pr)
}

// QueryRaces queries the "races" edge of the Proficiency entity.
func (pr *Proficiency) QueryRaces() *RaceQuery {
	return NewProficiencyClient(pr.config).QueryRaces(pr)
}

// QuerySubraces queries the "subraces" edge of the Proficiency entity.
func (pr *Proficiency) QuerySubraces() *SubraceQuery {
	return NewProficiencyClient(pr.config).QuerySubraces(pr)
}

// QueryChoice queries the "choice" edge of the Proficiency entity.
func (pr *Proficiency) QueryChoice() *ProficiencyChoiceQuery {
	return NewProficiencyClient(pr.config).QueryChoice(pr)
}

// QuerySkill queries the "skill" edge of the Proficiency entity.
func (pr *Proficiency) QuerySkill() *SkillQuery {
	return NewProficiencyClient(pr.config).QuerySkill(pr)
}

// QueryEquipment queries the "equipment" edge of the Proficiency entity.
func (pr *Proficiency) QueryEquipment() *EquipmentQuery {
	return NewProficiencyClient(pr.config).QueryEquipment(pr)
}

// QueryEquipmentCategory queries the "equipment_category" edge of the Proficiency entity.
func (pr *Proficiency) QueryEquipmentCategory() *EquipmentCategoryQuery {
	return NewProficiencyClient(pr.config).QueryEquipmentCategory(pr)
}

// QuerySavingThrow queries the "saving_throw" edge of the Proficiency entity.
func (pr *Proficiency) QuerySavingThrow() *AbilityScoreQuery {
	return NewProficiencyClient(pr.config).QuerySavingThrow(pr)
}

// Update returns a builder for updating this Proficiency.
// Note that you need to call Proficiency.Unwrap() before calling this method if this Proficiency
// was returned from a transaction, and the transaction was committed or rolled back.
func (pr *Proficiency) Update() *ProficiencyUpdateOne {
	return NewProficiencyClient(pr.config).UpdateOne(pr)
}

// Unwrap unwraps the Proficiency entity that was returned from a transaction after it was closed,
// so that all future queries will be executed through the driver which created the transaction.
func (pr *Proficiency) Unwrap() *Proficiency {
	_tx, ok := pr.config.driver.(*txDriver)
	if !ok {
		panic("ent: Proficiency is not a transactional entity")
	}
	pr.config.driver = _tx.drv
	return pr
}

// String implements the fmt.Stringer.
func (pr *Proficiency) String() string {
	var builder strings.Builder
	builder.WriteString("Proficiency(")
	builder.WriteString(fmt.Sprintf("id=%v, ", pr.ID))
	builder.WriteString("indx=")
	builder.WriteString(pr.Indx)
	builder.WriteString(", ")
	builder.WriteString("name=")
	builder.WriteString(pr.Name)
	builder.WriteString(", ")
	builder.WriteString("proficiency_category=")
	builder.WriteString(pr.ProficiencyCategory)
	builder.WriteByte(')')
	return builder.String()
}

// MarshalJSON implements the json.Marshaler interface.
func (pr *Proficiency) MarshalJSON() ([]byte, error) {
	type Alias Proficiency
	return json.Marshal(&struct {
		*Alias
		ProficiencyEdges
	}{
		Alias:            (*Alias)(pr),
		ProficiencyEdges: pr.Edges,
	})
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (pr *Proficiency) UnmarshalJSON(data []byte) error {
	type Alias Proficiency
	aux := &struct {
		*Alias
		ProficiencyEdges
	}{
		Alias: (*Alias)(pr),
	}

	if err := json.Unmarshal(data, &aux); err != nil {
		return err
	}

	pr.Edges = aux.ProficiencyEdges
	return nil
}

func (pc *ProficiencyCreate) SetProficiency(input *Proficiency) *ProficiencyCreate {
	pc.SetIndx(input.Indx)
	pc.SetName(input.Name)
	pc.SetProficiencyCategory(input.ProficiencyCategory)
	return pc
}

// NamedClasses returns the Classes named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedClasses(name string) ([]*Class, error) {
	if pr.Edges.namedClasses == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedClasses[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedClasses(name string, edges ...*Class) {
	if pr.Edges.namedClasses == nil {
		pr.Edges.namedClasses = make(map[string][]*Class)
	}
	if len(edges) == 0 {
		pr.Edges.namedClasses[name] = []*Class{}
	} else {
		pr.Edges.namedClasses[name] = append(pr.Edges.namedClasses[name], edges...)
	}
}

// NamedRaces returns the Races named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedRaces(name string) ([]*Race, error) {
	if pr.Edges.namedRaces == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedRaces[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedRaces(name string, edges ...*Race) {
	if pr.Edges.namedRaces == nil {
		pr.Edges.namedRaces = make(map[string][]*Race)
	}
	if len(edges) == 0 {
		pr.Edges.namedRaces[name] = []*Race{}
	} else {
		pr.Edges.namedRaces[name] = append(pr.Edges.namedRaces[name], edges...)
	}
}

// NamedSubraces returns the Subraces named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedSubraces(name string) ([]*Subrace, error) {
	if pr.Edges.namedSubraces == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedSubraces[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedSubraces(name string, edges ...*Subrace) {
	if pr.Edges.namedSubraces == nil {
		pr.Edges.namedSubraces = make(map[string][]*Subrace)
	}
	if len(edges) == 0 {
		pr.Edges.namedSubraces[name] = []*Subrace{}
	} else {
		pr.Edges.namedSubraces[name] = append(pr.Edges.namedSubraces[name], edges...)
	}
}

// NamedChoice returns the Choice named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedChoice(name string) ([]*ProficiencyChoice, error) {
	if pr.Edges.namedChoice == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedChoice[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedChoice(name string, edges ...*ProficiencyChoice) {
	if pr.Edges.namedChoice == nil {
		pr.Edges.namedChoice = make(map[string][]*ProficiencyChoice)
	}
	if len(edges) == 0 {
		pr.Edges.namedChoice[name] = []*ProficiencyChoice{}
	} else {
		pr.Edges.namedChoice[name] = append(pr.Edges.namedChoice[name], edges...)
	}
}

// NamedSkill returns the Skill named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedSkill(name string) ([]*Skill, error) {
	if pr.Edges.namedSkill == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedSkill[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedSkill(name string, edges ...*Skill) {
	if pr.Edges.namedSkill == nil {
		pr.Edges.namedSkill = make(map[string][]*Skill)
	}
	if len(edges) == 0 {
		pr.Edges.namedSkill[name] = []*Skill{}
	} else {
		pr.Edges.namedSkill[name] = append(pr.Edges.namedSkill[name], edges...)
	}
}

// NamedEquipment returns the Equipment named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedEquipment(name string) ([]*Equipment, error) {
	if pr.Edges.namedEquipment == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedEquipment[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedEquipment(name string, edges ...*Equipment) {
	if pr.Edges.namedEquipment == nil {
		pr.Edges.namedEquipment = make(map[string][]*Equipment)
	}
	if len(edges) == 0 {
		pr.Edges.namedEquipment[name] = []*Equipment{}
	} else {
		pr.Edges.namedEquipment[name] = append(pr.Edges.namedEquipment[name], edges...)
	}
}

// NamedEquipmentCategory returns the EquipmentCategory named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedEquipmentCategory(name string) ([]*EquipmentCategory, error) {
	if pr.Edges.namedEquipmentCategory == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedEquipmentCategory[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedEquipmentCategory(name string, edges ...*EquipmentCategory) {
	if pr.Edges.namedEquipmentCategory == nil {
		pr.Edges.namedEquipmentCategory = make(map[string][]*EquipmentCategory)
	}
	if len(edges) == 0 {
		pr.Edges.namedEquipmentCategory[name] = []*EquipmentCategory{}
	} else {
		pr.Edges.namedEquipmentCategory[name] = append(pr.Edges.namedEquipmentCategory[name], edges...)
	}
}

// NamedSavingThrow returns the SavingThrow named value or an error if the edge was not
// loaded in eager-loading with this name.
func (pr *Proficiency) NamedSavingThrow(name string) ([]*AbilityScore, error) {
	if pr.Edges.namedSavingThrow == nil {
		return nil, &NotLoadedError{edge: name}
	}
	nodes, ok := pr.Edges.namedSavingThrow[name]
	if !ok {
		return nil, &NotLoadedError{edge: name}
	}
	return nodes, nil
}

func (pr *Proficiency) appendNamedSavingThrow(name string, edges ...*AbilityScore) {
	if pr.Edges.namedSavingThrow == nil {
		pr.Edges.namedSavingThrow = make(map[string][]*AbilityScore)
	}
	if len(edges) == 0 {
		pr.Edges.namedSavingThrow[name] = []*AbilityScore{}
	} else {
		pr.Edges.namedSavingThrow[name] = append(pr.Edges.namedSavingThrow[name], edges...)
	}
}

// Proficiencies is a parsable slice of Proficiency.
type Proficiencies []*Proficiency
