// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/choice"
	"github.com/ecshreve/dndgen/ent/language"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/race"
	"github.com/ecshreve/dndgen/ent/subrace"
	"github.com/ecshreve/dndgen/ent/trait"
)

// RaceUpdate is the builder for updating Race entities.
type RaceUpdate struct {
	config
	hooks    []Hook
	mutation *RaceMutation
}

// Where appends a list predicates to the RaceUpdate builder.
func (ru *RaceUpdate) Where(ps ...predicate.Race) *RaceUpdate {
	ru.mutation.Where(ps...)
	return ru
}

// SetIndx sets the "indx" field.
func (ru *RaceUpdate) SetIndx(s string) *RaceUpdate {
	ru.mutation.SetIndx(s)
	return ru
}

// SetName sets the "name" field.
func (ru *RaceUpdate) SetName(s string) *RaceUpdate {
	ru.mutation.SetName(s)
	return ru
}

// SetAlignment sets the "alignment" field.
func (ru *RaceUpdate) SetAlignment(s string) *RaceUpdate {
	ru.mutation.SetAlignment(s)
	return ru
}

// SetAge sets the "age" field.
func (ru *RaceUpdate) SetAge(s string) *RaceUpdate {
	ru.mutation.SetAge(s)
	return ru
}

// SetSize sets the "size" field.
func (ru *RaceUpdate) SetSize(s string) *RaceUpdate {
	ru.mutation.SetSize(s)
	return ru
}

// SetSizeDescription sets the "size_description" field.
func (ru *RaceUpdate) SetSizeDescription(s string) *RaceUpdate {
	ru.mutation.SetSizeDescription(s)
	return ru
}

// SetLanguageDesc sets the "language_desc" field.
func (ru *RaceUpdate) SetLanguageDesc(s string) *RaceUpdate {
	ru.mutation.SetLanguageDesc(s)
	return ru
}

// SetSpeed sets the "speed" field.
func (ru *RaceUpdate) SetSpeed(i int) *RaceUpdate {
	ru.mutation.ResetSpeed()
	ru.mutation.SetSpeed(i)
	return ru
}

// AddSpeed adds i to the "speed" field.
func (ru *RaceUpdate) AddSpeed(i int) *RaceUpdate {
	ru.mutation.AddSpeed(i)
	return ru
}

// AddLanguageIDs adds the "languages" edge to the Language entity by IDs.
func (ru *RaceUpdate) AddLanguageIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddLanguageIDs(ids...)
	return ru
}

// AddLanguages adds the "languages" edges to the Language entity.
func (ru *RaceUpdate) AddLanguages(l ...*Language) *RaceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.AddLanguageIDs(ids...)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (ru *RaceUpdate) AddProficiencyIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddProficiencyIDs(ids...)
	return ru
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (ru *RaceUpdate) AddProficiencies(p ...*Proficiency) *RaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.AddProficiencyIDs(ids...)
}

// AddSubraceIDs adds the "subraces" edge to the Subrace entity by IDs.
func (ru *RaceUpdate) AddSubraceIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddSubraceIDs(ids...)
	return ru
}

// AddSubraces adds the "subraces" edges to the Subrace entity.
func (ru *RaceUpdate) AddSubraces(s ...*Subrace) *RaceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.AddSubraceIDs(ids...)
}

// AddTraitIDs adds the "traits" edge to the Trait entity by IDs.
func (ru *RaceUpdate) AddTraitIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddTraitIDs(ids...)
	return ru
}

// AddTraits adds the "traits" edges to the Trait entity.
func (ru *RaceUpdate) AddTraits(t ...*Trait) *RaceUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.AddTraitIDs(ids...)
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (ru *RaceUpdate) AddAbilityBonuseIDs(ids ...int) *RaceUpdate {
	ru.mutation.AddAbilityBonuseIDs(ids...)
	return ru
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (ru *RaceUpdate) AddAbilityBonuses(a ...*AbilityBonus) *RaceUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.AddAbilityBonuseIDs(ids...)
}

// SetStartingProficiencyOptionID sets the "starting_proficiency_option" edge to the Choice entity by ID.
func (ru *RaceUpdate) SetStartingProficiencyOptionID(id int) *RaceUpdate {
	ru.mutation.SetStartingProficiencyOptionID(id)
	return ru
}

// SetNillableStartingProficiencyOptionID sets the "starting_proficiency_option" edge to the Choice entity by ID if the given value is not nil.
func (ru *RaceUpdate) SetNillableStartingProficiencyOptionID(id *int) *RaceUpdate {
	if id != nil {
		ru = ru.SetStartingProficiencyOptionID(*id)
	}
	return ru
}

// SetStartingProficiencyOption sets the "starting_proficiency_option" edge to the Choice entity.
func (ru *RaceUpdate) SetStartingProficiencyOption(c *Choice) *RaceUpdate {
	return ru.SetStartingProficiencyOptionID(c.ID)
}

// Mutation returns the RaceMutation object of the builder.
func (ru *RaceUpdate) Mutation() *RaceMutation {
	return ru.mutation
}

// ClearLanguages clears all "languages" edges to the Language entity.
func (ru *RaceUpdate) ClearLanguages() *RaceUpdate {
	ru.mutation.ClearLanguages()
	return ru
}

// RemoveLanguageIDs removes the "languages" edge to Language entities by IDs.
func (ru *RaceUpdate) RemoveLanguageIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemoveLanguageIDs(ids...)
	return ru
}

// RemoveLanguages removes "languages" edges to Language entities.
func (ru *RaceUpdate) RemoveLanguages(l ...*Language) *RaceUpdate {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ru.RemoveLanguageIDs(ids...)
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (ru *RaceUpdate) ClearProficiencies() *RaceUpdate {
	ru.mutation.ClearProficiencies()
	return ru
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (ru *RaceUpdate) RemoveProficiencyIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemoveProficiencyIDs(ids...)
	return ru
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (ru *RaceUpdate) RemoveProficiencies(p ...*Proficiency) *RaceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ru.RemoveProficiencyIDs(ids...)
}

// ClearSubraces clears all "subraces" edges to the Subrace entity.
func (ru *RaceUpdate) ClearSubraces() *RaceUpdate {
	ru.mutation.ClearSubraces()
	return ru
}

// RemoveSubraceIDs removes the "subraces" edge to Subrace entities by IDs.
func (ru *RaceUpdate) RemoveSubraceIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemoveSubraceIDs(ids...)
	return ru
}

// RemoveSubraces removes "subraces" edges to Subrace entities.
func (ru *RaceUpdate) RemoveSubraces(s ...*Subrace) *RaceUpdate {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ru.RemoveSubraceIDs(ids...)
}

// ClearTraits clears all "traits" edges to the Trait entity.
func (ru *RaceUpdate) ClearTraits() *RaceUpdate {
	ru.mutation.ClearTraits()
	return ru
}

// RemoveTraitIDs removes the "traits" edge to Trait entities by IDs.
func (ru *RaceUpdate) RemoveTraitIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemoveTraitIDs(ids...)
	return ru
}

// RemoveTraits removes "traits" edges to Trait entities.
func (ru *RaceUpdate) RemoveTraits(t ...*Trait) *RaceUpdate {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ru.RemoveTraitIDs(ids...)
}

// ClearAbilityBonuses clears all "ability_bonuses" edges to the AbilityBonus entity.
func (ru *RaceUpdate) ClearAbilityBonuses() *RaceUpdate {
	ru.mutation.ClearAbilityBonuses()
	return ru
}

// RemoveAbilityBonuseIDs removes the "ability_bonuses" edge to AbilityBonus entities by IDs.
func (ru *RaceUpdate) RemoveAbilityBonuseIDs(ids ...int) *RaceUpdate {
	ru.mutation.RemoveAbilityBonuseIDs(ids...)
	return ru
}

// RemoveAbilityBonuses removes "ability_bonuses" edges to AbilityBonus entities.
func (ru *RaceUpdate) RemoveAbilityBonuses(a ...*AbilityBonus) *RaceUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ru.RemoveAbilityBonuseIDs(ids...)
}

// ClearStartingProficiencyOption clears the "starting_proficiency_option" edge to the Choice entity.
func (ru *RaceUpdate) ClearStartingProficiencyOption() *RaceUpdate {
	ru.mutation.ClearStartingProficiencyOption()
	return ru
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (ru *RaceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, ru.sqlSave, ru.mutation, ru.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ru *RaceUpdate) SaveX(ctx context.Context) int {
	affected, err := ru.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (ru *RaceUpdate) Exec(ctx context.Context) error {
	_, err := ru.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ru *RaceUpdate) ExecX(ctx context.Context) {
	if err := ru.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ru *RaceUpdate) check() error {
	if v, ok := ru.mutation.Indx(); ok {
		if err := race.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Race.indx": %w`, err)}
		}
	}
	if v, ok := ru.mutation.Name(); ok {
		if err := race.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Race.name": %w`, err)}
		}
	}
	return nil
}

func (ru *RaceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := ru.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(race.Table, race.Columns, sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt))
	if ps := ru.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ru.mutation.Indx(); ok {
		_spec.SetField(race.FieldIndx, field.TypeString, value)
	}
	if value, ok := ru.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
	}
	if value, ok := ru.mutation.Alignment(); ok {
		_spec.SetField(race.FieldAlignment, field.TypeString, value)
	}
	if value, ok := ru.mutation.Age(); ok {
		_spec.SetField(race.FieldAge, field.TypeString, value)
	}
	if value, ok := ru.mutation.Size(); ok {
		_spec.SetField(race.FieldSize, field.TypeString, value)
	}
	if value, ok := ru.mutation.SizeDescription(); ok {
		_spec.SetField(race.FieldSizeDescription, field.TypeString, value)
	}
	if value, ok := ru.mutation.LanguageDesc(); ok {
		_spec.SetField(race.FieldLanguageDesc, field.TypeString, value)
	}
	if value, ok := ru.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := ru.mutation.AddedSpeed(); ok {
		_spec.AddField(race.FieldSpeed, field.TypeInt, value)
	}
	if ru.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedLanguagesIDs(); len(nodes) > 0 && !ru.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.LanguagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !ru.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedSubracesIDs(); len(nodes) > 0 && !ru.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.SubracesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.TraitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedTraitsIDs(); len(nodes) > 0 && !ru.mutation.TraitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.TraitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.RemovedAbilityBonusesIDs(); len(nodes) > 0 && !ru.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ru.mutation.StartingProficiencyOptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   race.StartingProficiencyOptionTable,
			Columns: []string{race.StartingProficiencyOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ru.mutation.StartingProficiencyOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   race.StartingProficiencyOptionTable,
			Columns: []string{race.StartingProficiencyOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, ru.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	ru.mutation.done = true
	return n, nil
}

// RaceUpdateOne is the builder for updating a single Race entity.
type RaceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *RaceMutation
}

// SetIndx sets the "indx" field.
func (ruo *RaceUpdateOne) SetIndx(s string) *RaceUpdateOne {
	ruo.mutation.SetIndx(s)
	return ruo
}

// SetName sets the "name" field.
func (ruo *RaceUpdateOne) SetName(s string) *RaceUpdateOne {
	ruo.mutation.SetName(s)
	return ruo
}

// SetAlignment sets the "alignment" field.
func (ruo *RaceUpdateOne) SetAlignment(s string) *RaceUpdateOne {
	ruo.mutation.SetAlignment(s)
	return ruo
}

// SetAge sets the "age" field.
func (ruo *RaceUpdateOne) SetAge(s string) *RaceUpdateOne {
	ruo.mutation.SetAge(s)
	return ruo
}

// SetSize sets the "size" field.
func (ruo *RaceUpdateOne) SetSize(s string) *RaceUpdateOne {
	ruo.mutation.SetSize(s)
	return ruo
}

// SetSizeDescription sets the "size_description" field.
func (ruo *RaceUpdateOne) SetSizeDescription(s string) *RaceUpdateOne {
	ruo.mutation.SetSizeDescription(s)
	return ruo
}

// SetLanguageDesc sets the "language_desc" field.
func (ruo *RaceUpdateOne) SetLanguageDesc(s string) *RaceUpdateOne {
	ruo.mutation.SetLanguageDesc(s)
	return ruo
}

// SetSpeed sets the "speed" field.
func (ruo *RaceUpdateOne) SetSpeed(i int) *RaceUpdateOne {
	ruo.mutation.ResetSpeed()
	ruo.mutation.SetSpeed(i)
	return ruo
}

// AddSpeed adds i to the "speed" field.
func (ruo *RaceUpdateOne) AddSpeed(i int) *RaceUpdateOne {
	ruo.mutation.AddSpeed(i)
	return ruo
}

// AddLanguageIDs adds the "languages" edge to the Language entity by IDs.
func (ruo *RaceUpdateOne) AddLanguageIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddLanguageIDs(ids...)
	return ruo
}

// AddLanguages adds the "languages" edges to the Language entity.
func (ruo *RaceUpdateOne) AddLanguages(l ...*Language) *RaceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.AddLanguageIDs(ids...)
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (ruo *RaceUpdateOne) AddProficiencyIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddProficiencyIDs(ids...)
	return ruo
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (ruo *RaceUpdateOne) AddProficiencies(p ...*Proficiency) *RaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.AddProficiencyIDs(ids...)
}

// AddSubraceIDs adds the "subraces" edge to the Subrace entity by IDs.
func (ruo *RaceUpdateOne) AddSubraceIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddSubraceIDs(ids...)
	return ruo
}

// AddSubraces adds the "subraces" edges to the Subrace entity.
func (ruo *RaceUpdateOne) AddSubraces(s ...*Subrace) *RaceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.AddSubraceIDs(ids...)
}

// AddTraitIDs adds the "traits" edge to the Trait entity by IDs.
func (ruo *RaceUpdateOne) AddTraitIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddTraitIDs(ids...)
	return ruo
}

// AddTraits adds the "traits" edges to the Trait entity.
func (ruo *RaceUpdateOne) AddTraits(t ...*Trait) *RaceUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.AddTraitIDs(ids...)
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (ruo *RaceUpdateOne) AddAbilityBonuseIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.AddAbilityBonuseIDs(ids...)
	return ruo
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (ruo *RaceUpdateOne) AddAbilityBonuses(a ...*AbilityBonus) *RaceUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.AddAbilityBonuseIDs(ids...)
}

// SetStartingProficiencyOptionID sets the "starting_proficiency_option" edge to the Choice entity by ID.
func (ruo *RaceUpdateOne) SetStartingProficiencyOptionID(id int) *RaceUpdateOne {
	ruo.mutation.SetStartingProficiencyOptionID(id)
	return ruo
}

// SetNillableStartingProficiencyOptionID sets the "starting_proficiency_option" edge to the Choice entity by ID if the given value is not nil.
func (ruo *RaceUpdateOne) SetNillableStartingProficiencyOptionID(id *int) *RaceUpdateOne {
	if id != nil {
		ruo = ruo.SetStartingProficiencyOptionID(*id)
	}
	return ruo
}

// SetStartingProficiencyOption sets the "starting_proficiency_option" edge to the Choice entity.
func (ruo *RaceUpdateOne) SetStartingProficiencyOption(c *Choice) *RaceUpdateOne {
	return ruo.SetStartingProficiencyOptionID(c.ID)
}

// Mutation returns the RaceMutation object of the builder.
func (ruo *RaceUpdateOne) Mutation() *RaceMutation {
	return ruo.mutation
}

// ClearLanguages clears all "languages" edges to the Language entity.
func (ruo *RaceUpdateOne) ClearLanguages() *RaceUpdateOne {
	ruo.mutation.ClearLanguages()
	return ruo
}

// RemoveLanguageIDs removes the "languages" edge to Language entities by IDs.
func (ruo *RaceUpdateOne) RemoveLanguageIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemoveLanguageIDs(ids...)
	return ruo
}

// RemoveLanguages removes "languages" edges to Language entities.
func (ruo *RaceUpdateOne) RemoveLanguages(l ...*Language) *RaceUpdateOne {
	ids := make([]int, len(l))
	for i := range l {
		ids[i] = l[i].ID
	}
	return ruo.RemoveLanguageIDs(ids...)
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (ruo *RaceUpdateOne) ClearProficiencies() *RaceUpdateOne {
	ruo.mutation.ClearProficiencies()
	return ruo
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (ruo *RaceUpdateOne) RemoveProficiencyIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemoveProficiencyIDs(ids...)
	return ruo
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (ruo *RaceUpdateOne) RemoveProficiencies(p ...*Proficiency) *RaceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return ruo.RemoveProficiencyIDs(ids...)
}

// ClearSubraces clears all "subraces" edges to the Subrace entity.
func (ruo *RaceUpdateOne) ClearSubraces() *RaceUpdateOne {
	ruo.mutation.ClearSubraces()
	return ruo
}

// RemoveSubraceIDs removes the "subraces" edge to Subrace entities by IDs.
func (ruo *RaceUpdateOne) RemoveSubraceIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemoveSubraceIDs(ids...)
	return ruo
}

// RemoveSubraces removes "subraces" edges to Subrace entities.
func (ruo *RaceUpdateOne) RemoveSubraces(s ...*Subrace) *RaceUpdateOne {
	ids := make([]int, len(s))
	for i := range s {
		ids[i] = s[i].ID
	}
	return ruo.RemoveSubraceIDs(ids...)
}

// ClearTraits clears all "traits" edges to the Trait entity.
func (ruo *RaceUpdateOne) ClearTraits() *RaceUpdateOne {
	ruo.mutation.ClearTraits()
	return ruo
}

// RemoveTraitIDs removes the "traits" edge to Trait entities by IDs.
func (ruo *RaceUpdateOne) RemoveTraitIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemoveTraitIDs(ids...)
	return ruo
}

// RemoveTraits removes "traits" edges to Trait entities.
func (ruo *RaceUpdateOne) RemoveTraits(t ...*Trait) *RaceUpdateOne {
	ids := make([]int, len(t))
	for i := range t {
		ids[i] = t[i].ID
	}
	return ruo.RemoveTraitIDs(ids...)
}

// ClearAbilityBonuses clears all "ability_bonuses" edges to the AbilityBonus entity.
func (ruo *RaceUpdateOne) ClearAbilityBonuses() *RaceUpdateOne {
	ruo.mutation.ClearAbilityBonuses()
	return ruo
}

// RemoveAbilityBonuseIDs removes the "ability_bonuses" edge to AbilityBonus entities by IDs.
func (ruo *RaceUpdateOne) RemoveAbilityBonuseIDs(ids ...int) *RaceUpdateOne {
	ruo.mutation.RemoveAbilityBonuseIDs(ids...)
	return ruo
}

// RemoveAbilityBonuses removes "ability_bonuses" edges to AbilityBonus entities.
func (ruo *RaceUpdateOne) RemoveAbilityBonuses(a ...*AbilityBonus) *RaceUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return ruo.RemoveAbilityBonuseIDs(ids...)
}

// ClearStartingProficiencyOption clears the "starting_proficiency_option" edge to the Choice entity.
func (ruo *RaceUpdateOne) ClearStartingProficiencyOption() *RaceUpdateOne {
	ruo.mutation.ClearStartingProficiencyOption()
	return ruo
}

// Where appends a list predicates to the RaceUpdate builder.
func (ruo *RaceUpdateOne) Where(ps ...predicate.Race) *RaceUpdateOne {
	ruo.mutation.Where(ps...)
	return ruo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (ruo *RaceUpdateOne) Select(field string, fields ...string) *RaceUpdateOne {
	ruo.fields = append([]string{field}, fields...)
	return ruo
}

// Save executes the query and returns the updated Race entity.
func (ruo *RaceUpdateOne) Save(ctx context.Context) (*Race, error) {
	return withHooks(ctx, ruo.sqlSave, ruo.mutation, ruo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (ruo *RaceUpdateOne) SaveX(ctx context.Context) *Race {
	node, err := ruo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (ruo *RaceUpdateOne) Exec(ctx context.Context) error {
	_, err := ruo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (ruo *RaceUpdateOne) ExecX(ctx context.Context) {
	if err := ruo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (ruo *RaceUpdateOne) check() error {
	if v, ok := ruo.mutation.Indx(); ok {
		if err := race.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Race.indx": %w`, err)}
		}
	}
	if v, ok := ruo.mutation.Name(); ok {
		if err := race.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Race.name": %w`, err)}
		}
	}
	return nil
}

func (ruo *RaceUpdateOne) sqlSave(ctx context.Context) (_node *Race, err error) {
	if err := ruo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(race.Table, race.Columns, sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt))
	id, ok := ruo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Race.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := ruo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, race.FieldID)
		for _, f := range fields {
			if !race.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != race.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := ruo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := ruo.mutation.Indx(); ok {
		_spec.SetField(race.FieldIndx, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Alignment(); ok {
		_spec.SetField(race.FieldAlignment, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Age(); ok {
		_spec.SetField(race.FieldAge, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Size(); ok {
		_spec.SetField(race.FieldSize, field.TypeString, value)
	}
	if value, ok := ruo.mutation.SizeDescription(); ok {
		_spec.SetField(race.FieldSizeDescription, field.TypeString, value)
	}
	if value, ok := ruo.mutation.LanguageDesc(); ok {
		_spec.SetField(race.FieldLanguageDesc, field.TypeString, value)
	}
	if value, ok := ruo.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
	}
	if value, ok := ruo.mutation.AddedSpeed(); ok {
		_spec.AddField(race.FieldSpeed, field.TypeInt, value)
	}
	if ruo.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedLanguagesIDs(); len(nodes) > 0 && !ruo.mutation.LanguagesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.LanguagesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.LanguagesTable,
			Columns: race.LanguagesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(language.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !ruo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.ProficienciesTable,
			Columns: race.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedSubracesIDs(); len(nodes) > 0 && !ruo.mutation.SubracesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.SubracesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.SubracesTable,
			Columns: []string{race.SubracesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(subrace.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.TraitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedTraitsIDs(); len(nodes) > 0 && !ruo.mutation.TraitsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.TraitsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   race.TraitsTable,
			Columns: race.TraitsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(trait.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.RemovedAbilityBonusesIDs(); len(nodes) > 0 && !ruo.mutation.AbilityBonusesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   race.AbilityBonusesTable,
			Columns: []string{race.AbilityBonusesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilitybonus.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if ruo.mutation.StartingProficiencyOptionCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   race.StartingProficiencyOptionTable,
			Columns: []string{race.StartingProficiencyOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := ruo.mutation.StartingProficiencyOptionIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2O,
			Inverse: false,
			Table:   race.StartingProficiencyOptionTable,
			Columns: []string{race.StartingProficiencyOptionColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(choice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Race{config: ruo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, ruo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{race.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	ruo.mutation.done = true
	return _node, nil
}
