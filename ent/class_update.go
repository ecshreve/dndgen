// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/character"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipmententry"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
)

// ClassUpdate is the builder for updating Class entities.
type ClassUpdate struct {
	config
	hooks    []Hook
	mutation *ClassMutation
}

// Where appends a list predicates to the ClassUpdate builder.
func (cu *ClassUpdate) Where(ps ...predicate.Class) *ClassUpdate {
	cu.mutation.Where(ps...)
	return cu
}

// SetIndx sets the "indx" field.
func (cu *ClassUpdate) SetIndx(s string) *ClassUpdate {
	cu.mutation.SetIndx(s)
	return cu
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableIndx(s *string) *ClassUpdate {
	if s != nil {
		cu.SetIndx(*s)
	}
	return cu
}

// SetName sets the "name" field.
func (cu *ClassUpdate) SetName(s string) *ClassUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableName(s *string) *ClassUpdate {
	if s != nil {
		cu.SetName(*s)
	}
	return cu
}

// SetHitDie sets the "hit_die" field.
func (cu *ClassUpdate) SetHitDie(i int) *ClassUpdate {
	cu.mutation.ResetHitDie()
	cu.mutation.SetHitDie(i)
	return cu
}

// SetNillableHitDie sets the "hit_die" field if the given value is not nil.
func (cu *ClassUpdate) SetNillableHitDie(i *int) *ClassUpdate {
	if i != nil {
		cu.SetHitDie(*i)
	}
	return cu
}

// AddHitDie adds i to the "hit_die" field.
func (cu *ClassUpdate) AddHitDie(i int) *ClassUpdate {
	cu.mutation.AddHitDie(i)
	return cu
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (cu *ClassUpdate) AddProficiencyIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddProficiencyIDs(ids...)
	return cu
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (cu *ClassUpdate) AddProficiencies(p ...*Proficiency) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProficiencyIDs(ids...)
}

// AddProficiencyOptionIDs adds the "proficiency_options" edge to the ProficiencyChoice entity by IDs.
func (cu *ClassUpdate) AddProficiencyOptionIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddProficiencyOptionIDs(ids...)
	return cu
}

// AddProficiencyOptions adds the "proficiency_options" edges to the ProficiencyChoice entity.
func (cu *ClassUpdate) AddProficiencyOptions(p ...*ProficiencyChoice) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddProficiencyOptionIDs(ids...)
}

// AddStartingEquipmentIDs adds the "starting_equipment" edge to the EquipmentEntry entity by IDs.
func (cu *ClassUpdate) AddStartingEquipmentIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddStartingEquipmentIDs(ids...)
	return cu
}

// AddStartingEquipment adds the "starting_equipment" edges to the EquipmentEntry entity.
func (cu *ClassUpdate) AddStartingEquipment(e ...*EquipmentEntry) *ClassUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddStartingEquipmentIDs(ids...)
}

// AddSavingThrowIDs adds the "saving_throws" edge to the AbilityScore entity by IDs.
func (cu *ClassUpdate) AddSavingThrowIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddSavingThrowIDs(ids...)
	return cu
}

// AddSavingThrows adds the "saving_throws" edges to the AbilityScore entity.
func (cu *ClassUpdate) AddSavingThrows(a ...*AbilityScore) *ClassUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.AddSavingThrowIDs(ids...)
}

// AddCharacterIDs adds the "characters" edge to the Character entity by IDs.
func (cu *ClassUpdate) AddCharacterIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddCharacterIDs(ids...)
	return cu
}

// AddCharacters adds the "characters" edges to the Character entity.
func (cu *ClassUpdate) AddCharacters(c ...*Character) *ClassUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.AddCharacterIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cu *ClassUpdate) Mutation() *ClassMutation {
	return cu.mutation
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (cu *ClassUpdate) ClearProficiencies() *ClassUpdate {
	cu.mutation.ClearProficiencies()
	return cu
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (cu *ClassUpdate) RemoveProficiencyIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveProficiencyIDs(ids...)
	return cu
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (cu *ClassUpdate) RemoveProficiencies(p ...*Proficiency) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProficiencyIDs(ids...)
}

// ClearProficiencyOptions clears all "proficiency_options" edges to the ProficiencyChoice entity.
func (cu *ClassUpdate) ClearProficiencyOptions() *ClassUpdate {
	cu.mutation.ClearProficiencyOptions()
	return cu
}

// RemoveProficiencyOptionIDs removes the "proficiency_options" edge to ProficiencyChoice entities by IDs.
func (cu *ClassUpdate) RemoveProficiencyOptionIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveProficiencyOptionIDs(ids...)
	return cu
}

// RemoveProficiencyOptions removes "proficiency_options" edges to ProficiencyChoice entities.
func (cu *ClassUpdate) RemoveProficiencyOptions(p ...*ProficiencyChoice) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveProficiencyOptionIDs(ids...)
}

// ClearStartingEquipment clears all "starting_equipment" edges to the EquipmentEntry entity.
func (cu *ClassUpdate) ClearStartingEquipment() *ClassUpdate {
	cu.mutation.ClearStartingEquipment()
	return cu
}

// RemoveStartingEquipmentIDs removes the "starting_equipment" edge to EquipmentEntry entities by IDs.
func (cu *ClassUpdate) RemoveStartingEquipmentIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveStartingEquipmentIDs(ids...)
	return cu
}

// RemoveStartingEquipment removes "starting_equipment" edges to EquipmentEntry entities.
func (cu *ClassUpdate) RemoveStartingEquipment(e ...*EquipmentEntry) *ClassUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveStartingEquipmentIDs(ids...)
}

// ClearSavingThrows clears all "saving_throws" edges to the AbilityScore entity.
func (cu *ClassUpdate) ClearSavingThrows() *ClassUpdate {
	cu.mutation.ClearSavingThrows()
	return cu
}

// RemoveSavingThrowIDs removes the "saving_throws" edge to AbilityScore entities by IDs.
func (cu *ClassUpdate) RemoveSavingThrowIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveSavingThrowIDs(ids...)
	return cu
}

// RemoveSavingThrows removes "saving_throws" edges to AbilityScore entities.
func (cu *ClassUpdate) RemoveSavingThrows(a ...*AbilityScore) *ClassUpdate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cu.RemoveSavingThrowIDs(ids...)
}

// ClearCharacters clears all "characters" edges to the Character entity.
func (cu *ClassUpdate) ClearCharacters() *ClassUpdate {
	cu.mutation.ClearCharacters()
	return cu
}

// RemoveCharacterIDs removes the "characters" edge to Character entities by IDs.
func (cu *ClassUpdate) RemoveCharacterIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveCharacterIDs(ids...)
	return cu
}

// RemoveCharacters removes "characters" edges to Character entities.
func (cu *ClassUpdate) RemoveCharacters(c ...*Character) *ClassUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cu.RemoveCharacterIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, cu.sqlSave, cu.mutation, cu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cu *ClassUpdate) SaveX(ctx context.Context) int {
	affected, err := cu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (cu *ClassUpdate) Exec(ctx context.Context) error {
	_, err := cu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cu *ClassUpdate) ExecX(ctx context.Context) {
	if err := cu.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cu *ClassUpdate) check() error {
	if v, ok := cu.mutation.Indx(); ok {
		if err := class.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Class.indx": %w`, err)}
		}
	}
	if v, ok := cu.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Class.name": %w`, err)}
		}
	}
	if v, ok := cu.mutation.HitDie(); ok {
		if err := class.HitDieValidator(v); err != nil {
			return &ValidationError{Name: "hit_die", err: fmt.Errorf(`ent: validator failed for field "Class.hit_die": %w`, err)}
		}
	}
	return nil
}

func (cu *ClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
	if err := cu.check(); err != nil {
		return n, err
	}
	_spec := sqlgraph.NewUpdateSpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt))
	if ps := cu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cu.mutation.Indx(); ok {
		_spec.SetField(class.FieldIndx, field.TypeString, value)
	}
	if value, ok := cu.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
	}
	if value, ok := cu.mutation.HitDie(); ok {
		_spec.SetField(class.FieldHitDie, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedHitDie(); ok {
		_spec.AddField(class.FieldHitDie, field.TypeInt, value)
	}
	if cu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !cu.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
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
	if nodes := cu.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
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
	if cu.mutation.ProficiencyOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedProficiencyOptionsIDs(); len(nodes) > 0 && !cu.mutation.ProficiencyOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.ProficiencyOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStartingEquipmentIDs(); len(nodes) > 0 && !cu.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StartingEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSavingThrowsIDs(); len(nodes) > 0 && !cu.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.SavingThrowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cu.mutation.CharactersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedCharactersIDs(); len(nodes) > 0 && !cu.mutation.CharactersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.CharactersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, cu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	cu.mutation.done = true
	return n, nil
}

// ClassUpdateOne is the builder for updating a single Class entity.
type ClassUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ClassMutation
}

// SetIndx sets the "indx" field.
func (cuo *ClassUpdateOne) SetIndx(s string) *ClassUpdateOne {
	cuo.mutation.SetIndx(s)
	return cuo
}

// SetNillableIndx sets the "indx" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableIndx(s *string) *ClassUpdateOne {
	if s != nil {
		cuo.SetIndx(*s)
	}
	return cuo
}

// SetName sets the "name" field.
func (cuo *ClassUpdateOne) SetName(s string) *ClassUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetNillableName sets the "name" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableName(s *string) *ClassUpdateOne {
	if s != nil {
		cuo.SetName(*s)
	}
	return cuo
}

// SetHitDie sets the "hit_die" field.
func (cuo *ClassUpdateOne) SetHitDie(i int) *ClassUpdateOne {
	cuo.mutation.ResetHitDie()
	cuo.mutation.SetHitDie(i)
	return cuo
}

// SetNillableHitDie sets the "hit_die" field if the given value is not nil.
func (cuo *ClassUpdateOne) SetNillableHitDie(i *int) *ClassUpdateOne {
	if i != nil {
		cuo.SetHitDie(*i)
	}
	return cuo
}

// AddHitDie adds i to the "hit_die" field.
func (cuo *ClassUpdateOne) AddHitDie(i int) *ClassUpdateOne {
	cuo.mutation.AddHitDie(i)
	return cuo
}

// AddProficiencyIDs adds the "proficiencies" edge to the Proficiency entity by IDs.
func (cuo *ClassUpdateOne) AddProficiencyIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddProficiencyIDs(ids...)
	return cuo
}

// AddProficiencies adds the "proficiencies" edges to the Proficiency entity.
func (cuo *ClassUpdateOne) AddProficiencies(p ...*Proficiency) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProficiencyIDs(ids...)
}

// AddProficiencyOptionIDs adds the "proficiency_options" edge to the ProficiencyChoice entity by IDs.
func (cuo *ClassUpdateOne) AddProficiencyOptionIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddProficiencyOptionIDs(ids...)
	return cuo
}

// AddProficiencyOptions adds the "proficiency_options" edges to the ProficiencyChoice entity.
func (cuo *ClassUpdateOne) AddProficiencyOptions(p ...*ProficiencyChoice) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddProficiencyOptionIDs(ids...)
}

// AddStartingEquipmentIDs adds the "starting_equipment" edge to the EquipmentEntry entity by IDs.
func (cuo *ClassUpdateOne) AddStartingEquipmentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddStartingEquipmentIDs(ids...)
	return cuo
}

// AddStartingEquipment adds the "starting_equipment" edges to the EquipmentEntry entity.
func (cuo *ClassUpdateOne) AddStartingEquipment(e ...*EquipmentEntry) *ClassUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddStartingEquipmentIDs(ids...)
}

// AddSavingThrowIDs adds the "saving_throws" edge to the AbilityScore entity by IDs.
func (cuo *ClassUpdateOne) AddSavingThrowIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddSavingThrowIDs(ids...)
	return cuo
}

// AddSavingThrows adds the "saving_throws" edges to the AbilityScore entity.
func (cuo *ClassUpdateOne) AddSavingThrows(a ...*AbilityScore) *ClassUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.AddSavingThrowIDs(ids...)
}

// AddCharacterIDs adds the "characters" edge to the Character entity by IDs.
func (cuo *ClassUpdateOne) AddCharacterIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddCharacterIDs(ids...)
	return cuo
}

// AddCharacters adds the "characters" edges to the Character entity.
func (cuo *ClassUpdateOne) AddCharacters(c ...*Character) *ClassUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.AddCharacterIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cuo *ClassUpdateOne) Mutation() *ClassMutation {
	return cuo.mutation
}

// ClearProficiencies clears all "proficiencies" edges to the Proficiency entity.
func (cuo *ClassUpdateOne) ClearProficiencies() *ClassUpdateOne {
	cuo.mutation.ClearProficiencies()
	return cuo
}

// RemoveProficiencyIDs removes the "proficiencies" edge to Proficiency entities by IDs.
func (cuo *ClassUpdateOne) RemoveProficiencyIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveProficiencyIDs(ids...)
	return cuo
}

// RemoveProficiencies removes "proficiencies" edges to Proficiency entities.
func (cuo *ClassUpdateOne) RemoveProficiencies(p ...*Proficiency) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProficiencyIDs(ids...)
}

// ClearProficiencyOptions clears all "proficiency_options" edges to the ProficiencyChoice entity.
func (cuo *ClassUpdateOne) ClearProficiencyOptions() *ClassUpdateOne {
	cuo.mutation.ClearProficiencyOptions()
	return cuo
}

// RemoveProficiencyOptionIDs removes the "proficiency_options" edge to ProficiencyChoice entities by IDs.
func (cuo *ClassUpdateOne) RemoveProficiencyOptionIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveProficiencyOptionIDs(ids...)
	return cuo
}

// RemoveProficiencyOptions removes "proficiency_options" edges to ProficiencyChoice entities.
func (cuo *ClassUpdateOne) RemoveProficiencyOptions(p ...*ProficiencyChoice) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveProficiencyOptionIDs(ids...)
}

// ClearStartingEquipment clears all "starting_equipment" edges to the EquipmentEntry entity.
func (cuo *ClassUpdateOne) ClearStartingEquipment() *ClassUpdateOne {
	cuo.mutation.ClearStartingEquipment()
	return cuo
}

// RemoveStartingEquipmentIDs removes the "starting_equipment" edge to EquipmentEntry entities by IDs.
func (cuo *ClassUpdateOne) RemoveStartingEquipmentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveStartingEquipmentIDs(ids...)
	return cuo
}

// RemoveStartingEquipment removes "starting_equipment" edges to EquipmentEntry entities.
func (cuo *ClassUpdateOne) RemoveStartingEquipment(e ...*EquipmentEntry) *ClassUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveStartingEquipmentIDs(ids...)
}

// ClearSavingThrows clears all "saving_throws" edges to the AbilityScore entity.
func (cuo *ClassUpdateOne) ClearSavingThrows() *ClassUpdateOne {
	cuo.mutation.ClearSavingThrows()
	return cuo
}

// RemoveSavingThrowIDs removes the "saving_throws" edge to AbilityScore entities by IDs.
func (cuo *ClassUpdateOne) RemoveSavingThrowIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveSavingThrowIDs(ids...)
	return cuo
}

// RemoveSavingThrows removes "saving_throws" edges to AbilityScore entities.
func (cuo *ClassUpdateOne) RemoveSavingThrows(a ...*AbilityScore) *ClassUpdateOne {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return cuo.RemoveSavingThrowIDs(ids...)
}

// ClearCharacters clears all "characters" edges to the Character entity.
func (cuo *ClassUpdateOne) ClearCharacters() *ClassUpdateOne {
	cuo.mutation.ClearCharacters()
	return cuo
}

// RemoveCharacterIDs removes the "characters" edge to Character entities by IDs.
func (cuo *ClassUpdateOne) RemoveCharacterIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveCharacterIDs(ids...)
	return cuo
}

// RemoveCharacters removes "characters" edges to Character entities.
func (cuo *ClassUpdateOne) RemoveCharacters(c ...*Character) *ClassUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return cuo.RemoveCharacterIDs(ids...)
}

// Where appends a list predicates to the ClassUpdate builder.
func (cuo *ClassUpdateOne) Where(ps ...predicate.Class) *ClassUpdateOne {
	cuo.mutation.Where(ps...)
	return cuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (cuo *ClassUpdateOne) Select(field string, fields ...string) *ClassUpdateOne {
	cuo.fields = append([]string{field}, fields...)
	return cuo
}

// Save executes the query and returns the updated Class entity.
func (cuo *ClassUpdateOne) Save(ctx context.Context) (*Class, error) {
	return withHooks(ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (cuo *ClassUpdateOne) SaveX(ctx context.Context) *Class {
	node, err := cuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (cuo *ClassUpdateOne) Exec(ctx context.Context) error {
	_, err := cuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (cuo *ClassUpdateOne) ExecX(ctx context.Context) {
	if err := cuo.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (cuo *ClassUpdateOne) check() error {
	if v, ok := cuo.mutation.Indx(); ok {
		if err := class.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Class.indx": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.Name(); ok {
		if err := class.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Class.name": %w`, err)}
		}
	}
	if v, ok := cuo.mutation.HitDie(); ok {
		if err := class.HitDieValidator(v); err != nil {
			return &ValidationError{Name: "hit_die", err: fmt.Errorf(`ent: validator failed for field "Class.hit_die": %w`, err)}
		}
	}
	return nil
}

func (cuo *ClassUpdateOne) sqlSave(ctx context.Context) (_node *Class, err error) {
	if err := cuo.check(); err != nil {
		return _node, err
	}
	_spec := sqlgraph.NewUpdateSpec(class.Table, class.Columns, sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt))
	id, ok := cuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "Class.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := cuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, class.FieldID)
		for _, f := range fields {
			if !class.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != class.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := cuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := cuo.mutation.Indx(); ok {
		_spec.SetField(class.FieldIndx, field.TypeString, value)
	}
	if value, ok := cuo.mutation.Name(); ok {
		_spec.SetField(class.FieldName, field.TypeString, value)
	}
	if value, ok := cuo.mutation.HitDie(); ok {
		_spec.SetField(class.FieldHitDie, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedHitDie(); ok {
		_spec.AddField(class.FieldHitDie, field.TypeInt, value)
	}
	if cuo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProficienciesIDs(); len(nodes) > 0 && !cuo.mutation.ProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
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
	if nodes := cuo.mutation.ProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.ProficienciesTable,
			Columns: class.ProficienciesPrimaryKey,
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
	if cuo.mutation.ProficiencyOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedProficiencyOptionsIDs(); len(nodes) > 0 && !cuo.mutation.ProficiencyOptionsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.ProficiencyOptionsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.ProficiencyOptionsTable,
			Columns: []string{class.ProficiencyOptionsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStartingEquipmentIDs(); len(nodes) > 0 && !cuo.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StartingEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: class.StartingEquipmentPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipmententry.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSavingThrowsIDs(); len(nodes) > 0 && !cuo.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.SavingThrowsIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: class.SavingThrowsPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if cuo.mutation.CharactersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedCharactersIDs(); len(nodes) > 0 && !cuo.mutation.CharactersCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.CharactersIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: true,
			Table:   class.CharactersTable,
			Columns: []string{class.CharactersColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(character.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &Class{config: cuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, cuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{class.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	cuo.mutation.done = true
	return _node, nil
}
