// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/dialect/sql/sqljson"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/equipment"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
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

// SetName sets the "name" field.
func (cu *ClassUpdate) SetName(s string) *ClassUpdate {
	cu.mutation.SetName(s)
	return cu
}

// SetDesc sets the "desc" field.
func (cu *ClassUpdate) SetDesc(s []string) *ClassUpdate {
	cu.mutation.SetDesc(s)
	return cu
}

// AppendDesc appends s to the "desc" field.
func (cu *ClassUpdate) AppendDesc(s []string) *ClassUpdate {
	cu.mutation.AppendDesc(s)
	return cu
}

// ClearDesc clears the value of the "desc" field.
func (cu *ClassUpdate) ClearDesc() *ClassUpdate {
	cu.mutation.ClearDesc()
	return cu
}

// SetHitDie sets the "hit_die" field.
func (cu *ClassUpdate) SetHitDie(i int) *ClassUpdate {
	cu.mutation.ResetHitDie()
	cu.mutation.SetHitDie(i)
	return cu
}

// AddHitDie adds i to the "hit_die" field.
func (cu *ClassUpdate) AddHitDie(i int) *ClassUpdate {
	cu.mutation.AddHitDie(i)
	return cu
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

// AddStartingProficiencyIDs adds the "starting_proficiencies" edge to the Proficiency entity by IDs.
func (cu *ClassUpdate) AddStartingProficiencyIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddStartingProficiencyIDs(ids...)
	return cu
}

// AddStartingProficiencies adds the "starting_proficiencies" edges to the Proficiency entity.
func (cu *ClassUpdate) AddStartingProficiencies(p ...*Proficiency) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.AddStartingProficiencyIDs(ids...)
}

// AddStartingEquipmentIDs adds the "starting_equipment" edge to the Equipment entity by IDs.
func (cu *ClassUpdate) AddStartingEquipmentIDs(ids ...int) *ClassUpdate {
	cu.mutation.AddStartingEquipmentIDs(ids...)
	return cu
}

// AddStartingEquipment adds the "starting_equipment" edges to the Equipment entity.
func (cu *ClassUpdate) AddStartingEquipment(e ...*Equipment) *ClassUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.AddStartingEquipmentIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cu *ClassUpdate) Mutation() *ClassMutation {
	return cu.mutation
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

// ClearStartingProficiencies clears all "starting_proficiencies" edges to the Proficiency entity.
func (cu *ClassUpdate) ClearStartingProficiencies() *ClassUpdate {
	cu.mutation.ClearStartingProficiencies()
	return cu
}

// RemoveStartingProficiencyIDs removes the "starting_proficiencies" edge to Proficiency entities by IDs.
func (cu *ClassUpdate) RemoveStartingProficiencyIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveStartingProficiencyIDs(ids...)
	return cu
}

// RemoveStartingProficiencies removes "starting_proficiencies" edges to Proficiency entities.
func (cu *ClassUpdate) RemoveStartingProficiencies(p ...*Proficiency) *ClassUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cu.RemoveStartingProficiencyIDs(ids...)
}

// ClearStartingEquipment clears all "starting_equipment" edges to the Equipment entity.
func (cu *ClassUpdate) ClearStartingEquipment() *ClassUpdate {
	cu.mutation.ClearStartingEquipment()
	return cu
}

// RemoveStartingEquipmentIDs removes the "starting_equipment" edge to Equipment entities by IDs.
func (cu *ClassUpdate) RemoveStartingEquipmentIDs(ids ...int) *ClassUpdate {
	cu.mutation.RemoveStartingEquipmentIDs(ids...)
	return cu
}

// RemoveStartingEquipment removes "starting_equipment" edges to Equipment entities.
func (cu *ClassUpdate) RemoveStartingEquipment(e ...*Equipment) *ClassUpdate {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cu.RemoveStartingEquipmentIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (cu *ClassUpdate) Save(ctx context.Context) (int, error) {
	return withHooks[int, ClassMutation](ctx, cu.sqlSave, cu.mutation, cu.hooks)
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

func (cu *ClassUpdate) sqlSave(ctx context.Context) (n int, err error) {
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
	if value, ok := cu.mutation.Desc(); ok {
		_spec.SetField(class.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := cu.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, class.FieldDesc, value)
		})
	}
	if cu.mutation.DescCleared() {
		_spec.ClearField(class.FieldDesc, field.TypeJSON)
	}
	if value, ok := cu.mutation.HitDie(); ok {
		_spec.SetField(class.FieldHitDie, field.TypeInt, value)
	}
	if value, ok := cu.mutation.AddedHitDie(); ok {
		_spec.AddField(class.FieldHitDie, field.TypeInt, value)
	}
	if cu.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedSavingThrowsIDs(); len(nodes) > 0 && !cu.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
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
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
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
	if cu.mutation.StartingProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStartingProficienciesIDs(); len(nodes) > 0 && !cu.mutation.StartingProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
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
	if nodes := cu.mutation.StartingProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
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
	if cu.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.RemovedStartingEquipmentIDs(); len(nodes) > 0 && !cu.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cu.mutation.StartingEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
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

// SetName sets the "name" field.
func (cuo *ClassUpdateOne) SetName(s string) *ClassUpdateOne {
	cuo.mutation.SetName(s)
	return cuo
}

// SetDesc sets the "desc" field.
func (cuo *ClassUpdateOne) SetDesc(s []string) *ClassUpdateOne {
	cuo.mutation.SetDesc(s)
	return cuo
}

// AppendDesc appends s to the "desc" field.
func (cuo *ClassUpdateOne) AppendDesc(s []string) *ClassUpdateOne {
	cuo.mutation.AppendDesc(s)
	return cuo
}

// ClearDesc clears the value of the "desc" field.
func (cuo *ClassUpdateOne) ClearDesc() *ClassUpdateOne {
	cuo.mutation.ClearDesc()
	return cuo
}

// SetHitDie sets the "hit_die" field.
func (cuo *ClassUpdateOne) SetHitDie(i int) *ClassUpdateOne {
	cuo.mutation.ResetHitDie()
	cuo.mutation.SetHitDie(i)
	return cuo
}

// AddHitDie adds i to the "hit_die" field.
func (cuo *ClassUpdateOne) AddHitDie(i int) *ClassUpdateOne {
	cuo.mutation.AddHitDie(i)
	return cuo
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

// AddStartingProficiencyIDs adds the "starting_proficiencies" edge to the Proficiency entity by IDs.
func (cuo *ClassUpdateOne) AddStartingProficiencyIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddStartingProficiencyIDs(ids...)
	return cuo
}

// AddStartingProficiencies adds the "starting_proficiencies" edges to the Proficiency entity.
func (cuo *ClassUpdateOne) AddStartingProficiencies(p ...*Proficiency) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.AddStartingProficiencyIDs(ids...)
}

// AddStartingEquipmentIDs adds the "starting_equipment" edge to the Equipment entity by IDs.
func (cuo *ClassUpdateOne) AddStartingEquipmentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.AddStartingEquipmentIDs(ids...)
	return cuo
}

// AddStartingEquipment adds the "starting_equipment" edges to the Equipment entity.
func (cuo *ClassUpdateOne) AddStartingEquipment(e ...*Equipment) *ClassUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.AddStartingEquipmentIDs(ids...)
}

// Mutation returns the ClassMutation object of the builder.
func (cuo *ClassUpdateOne) Mutation() *ClassMutation {
	return cuo.mutation
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

// ClearStartingProficiencies clears all "starting_proficiencies" edges to the Proficiency entity.
func (cuo *ClassUpdateOne) ClearStartingProficiencies() *ClassUpdateOne {
	cuo.mutation.ClearStartingProficiencies()
	return cuo
}

// RemoveStartingProficiencyIDs removes the "starting_proficiencies" edge to Proficiency entities by IDs.
func (cuo *ClassUpdateOne) RemoveStartingProficiencyIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveStartingProficiencyIDs(ids...)
	return cuo
}

// RemoveStartingProficiencies removes "starting_proficiencies" edges to Proficiency entities.
func (cuo *ClassUpdateOne) RemoveStartingProficiencies(p ...*Proficiency) *ClassUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return cuo.RemoveStartingProficiencyIDs(ids...)
}

// ClearStartingEquipment clears all "starting_equipment" edges to the Equipment entity.
func (cuo *ClassUpdateOne) ClearStartingEquipment() *ClassUpdateOne {
	cuo.mutation.ClearStartingEquipment()
	return cuo
}

// RemoveStartingEquipmentIDs removes the "starting_equipment" edge to Equipment entities by IDs.
func (cuo *ClassUpdateOne) RemoveStartingEquipmentIDs(ids ...int) *ClassUpdateOne {
	cuo.mutation.RemoveStartingEquipmentIDs(ids...)
	return cuo
}

// RemoveStartingEquipment removes "starting_equipment" edges to Equipment entities.
func (cuo *ClassUpdateOne) RemoveStartingEquipment(e ...*Equipment) *ClassUpdateOne {
	ids := make([]int, len(e))
	for i := range e {
		ids[i] = e[i].ID
	}
	return cuo.RemoveStartingEquipmentIDs(ids...)
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
	return withHooks[*Class, ClassMutation](ctx, cuo.sqlSave, cuo.mutation, cuo.hooks)
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

func (cuo *ClassUpdateOne) sqlSave(ctx context.Context) (_node *Class, err error) {
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
	if value, ok := cuo.mutation.Desc(); ok {
		_spec.SetField(class.FieldDesc, field.TypeJSON, value)
	}
	if value, ok := cuo.mutation.AppendedDesc(); ok {
		_spec.AddModifier(func(u *sql.UpdateBuilder) {
			sqljson.Append(u, class.FieldDesc, value)
		})
	}
	if cuo.mutation.DescCleared() {
		_spec.ClearField(class.FieldDesc, field.TypeJSON)
	}
	if value, ok := cuo.mutation.HitDie(); ok {
		_spec.SetField(class.FieldHitDie, field.TypeInt, value)
	}
	if value, ok := cuo.mutation.AddedHitDie(); ok {
		_spec.AddField(class.FieldHitDie, field.TypeInt, value)
	}
	if cuo.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(abilityscore.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedSavingThrowsIDs(); len(nodes) > 0 && !cuo.mutation.SavingThrowsCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
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
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.SavingThrowsTable,
			Columns: []string{class.SavingThrowsColumn},
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
	if cuo.mutation.StartingProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStartingProficienciesIDs(); len(nodes) > 0 && !cuo.mutation.StartingProficienciesCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
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
	if nodes := cuo.mutation.StartingProficienciesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: false,
			Table:   class.StartingProficienciesTable,
			Columns: class.StartingProficienciesPrimaryKey,
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
	if cuo.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.RemovedStartingEquipmentIDs(); len(nodes) > 0 && !cuo.mutation.StartingEquipmentCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := cuo.mutation.StartingEquipmentIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   class.StartingEquipmentTable,
			Columns: []string{class.StartingEquipmentColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(equipment.FieldID, field.TypeInt),
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
