// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/class"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/proficiency"
	"github.com/ecshreve/dndgen/ent/proficiencychoice"
	"github.com/ecshreve/dndgen/ent/race"
)

// ProficiencyChoiceUpdate is the builder for updating ProficiencyChoice entities.
type ProficiencyChoiceUpdate struct {
	config
	hooks    []Hook
	mutation *ProficiencyChoiceMutation
}

// Where appends a list predicates to the ProficiencyChoiceUpdate builder.
func (pcu *ProficiencyChoiceUpdate) Where(ps ...predicate.ProficiencyChoice) *ProficiencyChoiceUpdate {
	pcu.mutation.Where(ps...)
	return pcu
}

// SetChoose sets the "choose" field.
func (pcu *ProficiencyChoiceUpdate) SetChoose(i int) *ProficiencyChoiceUpdate {
	pcu.mutation.ResetChoose()
	pcu.mutation.SetChoose(i)
	return pcu
}

// AddChoose adds i to the "choose" field.
func (pcu *ProficiencyChoiceUpdate) AddChoose(i int) *ProficiencyChoiceUpdate {
	pcu.mutation.AddChoose(i)
	return pcu
}

// SetDesc sets the "desc" field.
func (pcu *ProficiencyChoiceUpdate) SetDesc(s string) *ProficiencyChoiceUpdate {
	pcu.mutation.SetDesc(s)
	return pcu
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (pcu *ProficiencyChoiceUpdate) SetNillableDesc(s *string) *ProficiencyChoiceUpdate {
	if s != nil {
		pcu.SetDesc(*s)
	}
	return pcu
}

// ClearDesc clears the value of the "desc" field.
func (pcu *ProficiencyChoiceUpdate) ClearDesc() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearDesc()
	return pcu
}

// AddProficiencyIDs adds the "proficiency" edge to the Proficiency entity by IDs.
func (pcu *ProficiencyChoiceUpdate) AddProficiencyIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.AddProficiencyIDs(ids...)
	return pcu
}

// AddProficiency adds the "proficiency" edges to the Proficiency entity.
func (pcu *ProficiencyChoiceUpdate) AddProficiency(p ...*Proficiency) *ProficiencyChoiceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.AddProficiencyIDs(ids...)
}

// SetParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID.
func (pcu *ProficiencyChoiceUpdate) SetParentChoiceID(id int) *ProficiencyChoiceUpdate {
	pcu.mutation.SetParentChoiceID(id)
	return pcu
}

// SetNillableParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID if the given value is not nil.
func (pcu *ProficiencyChoiceUpdate) SetNillableParentChoiceID(id *int) *ProficiencyChoiceUpdate {
	if id != nil {
		pcu = pcu.SetParentChoiceID(*id)
	}
	return pcu
}

// SetParentChoice sets the "parent_choice" edge to the ProficiencyChoice entity.
func (pcu *ProficiencyChoiceUpdate) SetParentChoice(p *ProficiencyChoice) *ProficiencyChoiceUpdate {
	return pcu.SetParentChoiceID(p.ID)
}

// AddSubChoiceIDs adds the "sub_choice" edge to the ProficiencyChoice entity by IDs.
func (pcu *ProficiencyChoiceUpdate) AddSubChoiceIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.AddSubChoiceIDs(ids...)
	return pcu
}

// AddSubChoice adds the "sub_choice" edges to the ProficiencyChoice entity.
func (pcu *ProficiencyChoiceUpdate) AddSubChoice(p ...*ProficiencyChoice) *ProficiencyChoiceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.AddSubChoiceIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (pcu *ProficiencyChoiceUpdate) AddClasIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.AddClasIDs(ids...)
	return pcu
}

// AddClass adds the "class" edges to the Class entity.
func (pcu *ProficiencyChoiceUpdate) AddClass(c ...*Class) *ProficiencyChoiceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pcu.AddClasIDs(ids...)
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (pcu *ProficiencyChoiceUpdate) AddRaceIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.AddRaceIDs(ids...)
	return pcu
}

// AddRace adds the "race" edges to the Race entity.
func (pcu *ProficiencyChoiceUpdate) AddRace(r ...*Race) *ProficiencyChoiceUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pcu.AddRaceIDs(ids...)
}

// Mutation returns the ProficiencyChoiceMutation object of the builder.
func (pcu *ProficiencyChoiceUpdate) Mutation() *ProficiencyChoiceMutation {
	return pcu.mutation
}

// ClearProficiency clears all "proficiency" edges to the Proficiency entity.
func (pcu *ProficiencyChoiceUpdate) ClearProficiency() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearProficiency()
	return pcu
}

// RemoveProficiencyIDs removes the "proficiency" edge to Proficiency entities by IDs.
func (pcu *ProficiencyChoiceUpdate) RemoveProficiencyIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.RemoveProficiencyIDs(ids...)
	return pcu
}

// RemoveProficiency removes "proficiency" edges to Proficiency entities.
func (pcu *ProficiencyChoiceUpdate) RemoveProficiency(p ...*Proficiency) *ProficiencyChoiceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.RemoveProficiencyIDs(ids...)
}

// ClearParentChoice clears the "parent_choice" edge to the ProficiencyChoice entity.
func (pcu *ProficiencyChoiceUpdate) ClearParentChoice() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearParentChoice()
	return pcu
}

// ClearSubChoice clears all "sub_choice" edges to the ProficiencyChoice entity.
func (pcu *ProficiencyChoiceUpdate) ClearSubChoice() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearSubChoice()
	return pcu
}

// RemoveSubChoiceIDs removes the "sub_choice" edge to ProficiencyChoice entities by IDs.
func (pcu *ProficiencyChoiceUpdate) RemoveSubChoiceIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.RemoveSubChoiceIDs(ids...)
	return pcu
}

// RemoveSubChoice removes "sub_choice" edges to ProficiencyChoice entities.
func (pcu *ProficiencyChoiceUpdate) RemoveSubChoice(p ...*ProficiencyChoice) *ProficiencyChoiceUpdate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcu.RemoveSubChoiceIDs(ids...)
}

// ClearClass clears all "class" edges to the Class entity.
func (pcu *ProficiencyChoiceUpdate) ClearClass() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearClass()
	return pcu
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (pcu *ProficiencyChoiceUpdate) RemoveClasIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.RemoveClasIDs(ids...)
	return pcu
}

// RemoveClass removes "class" edges to Class entities.
func (pcu *ProficiencyChoiceUpdate) RemoveClass(c ...*Class) *ProficiencyChoiceUpdate {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pcu.RemoveClasIDs(ids...)
}

// ClearRace clears all "race" edges to the Race entity.
func (pcu *ProficiencyChoiceUpdate) ClearRace() *ProficiencyChoiceUpdate {
	pcu.mutation.ClearRace()
	return pcu
}

// RemoveRaceIDs removes the "race" edge to Race entities by IDs.
func (pcu *ProficiencyChoiceUpdate) RemoveRaceIDs(ids ...int) *ProficiencyChoiceUpdate {
	pcu.mutation.RemoveRaceIDs(ids...)
	return pcu
}

// RemoveRace removes "race" edges to Race entities.
func (pcu *ProficiencyChoiceUpdate) RemoveRace(r ...*Race) *ProficiencyChoiceUpdate {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pcu.RemoveRaceIDs(ids...)
}

// Save executes the query and returns the number of nodes affected by the update operation.
func (pcu *ProficiencyChoiceUpdate) Save(ctx context.Context) (int, error) {
	return withHooks(ctx, pcu.sqlSave, pcu.mutation, pcu.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcu *ProficiencyChoiceUpdate) SaveX(ctx context.Context) int {
	affected, err := pcu.Save(ctx)
	if err != nil {
		panic(err)
	}
	return affected
}

// Exec executes the query.
func (pcu *ProficiencyChoiceUpdate) Exec(ctx context.Context) error {
	_, err := pcu.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcu *ProficiencyChoiceUpdate) ExecX(ctx context.Context) {
	if err := pcu.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcu *ProficiencyChoiceUpdate) sqlSave(ctx context.Context) (n int, err error) {
	_spec := sqlgraph.NewUpdateSpec(proficiencychoice.Table, proficiencychoice.Columns, sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt))
	if ps := pcu.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcu.mutation.Choose(); ok {
		_spec.SetField(proficiencychoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.AddedChoose(); ok {
		_spec.AddField(proficiencychoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := pcu.mutation.Desc(); ok {
		_spec.SetField(proficiencychoice.FieldDesc, field.TypeString, value)
	}
	if pcu.mutation.DescCleared() {
		_spec.ClearField(proficiencychoice.FieldDesc, field.TypeString)
	}
	if pcu.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedProficiencyIDs(); len(nodes) > 0 && !pcu.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
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
	if nodes := pcu.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
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
	if pcu.mutation.ParentChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proficiencychoice.ParentChoiceTable,
			Columns: []string{proficiencychoice.ParentChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ParentChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proficiencychoice.ParentChoiceTable,
			Columns: []string{proficiencychoice.ParentChoiceColumn},
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
	if pcu.mutation.SubChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedSubChoiceIDs(); len(nodes) > 0 && !pcu.mutation.SubChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
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
	if nodes := pcu.mutation.SubChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
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
	if pcu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedClassIDs(); len(nodes) > 0 && !pcu.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RemovedRaceIDs(); len(nodes) > 0 && !pcu.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcu.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if n, err = sqlgraph.UpdateNodes(ctx, pcu.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{proficiencychoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return 0, err
	}
	pcu.mutation.done = true
	return n, nil
}

// ProficiencyChoiceUpdateOne is the builder for updating a single ProficiencyChoice entity.
type ProficiencyChoiceUpdateOne struct {
	config
	fields   []string
	hooks    []Hook
	mutation *ProficiencyChoiceMutation
}

// SetChoose sets the "choose" field.
func (pcuo *ProficiencyChoiceUpdateOne) SetChoose(i int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ResetChoose()
	pcuo.mutation.SetChoose(i)
	return pcuo
}

// AddChoose adds i to the "choose" field.
func (pcuo *ProficiencyChoiceUpdateOne) AddChoose(i int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.AddChoose(i)
	return pcuo
}

// SetDesc sets the "desc" field.
func (pcuo *ProficiencyChoiceUpdateOne) SetDesc(s string) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.SetDesc(s)
	return pcuo
}

// SetNillableDesc sets the "desc" field if the given value is not nil.
func (pcuo *ProficiencyChoiceUpdateOne) SetNillableDesc(s *string) *ProficiencyChoiceUpdateOne {
	if s != nil {
		pcuo.SetDesc(*s)
	}
	return pcuo
}

// ClearDesc clears the value of the "desc" field.
func (pcuo *ProficiencyChoiceUpdateOne) ClearDesc() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearDesc()
	return pcuo
}

// AddProficiencyIDs adds the "proficiency" edge to the Proficiency entity by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) AddProficiencyIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.AddProficiencyIDs(ids...)
	return pcuo
}

// AddProficiency adds the "proficiency" edges to the Proficiency entity.
func (pcuo *ProficiencyChoiceUpdateOne) AddProficiency(p ...*Proficiency) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.AddProficiencyIDs(ids...)
}

// SetParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID.
func (pcuo *ProficiencyChoiceUpdateOne) SetParentChoiceID(id int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.SetParentChoiceID(id)
	return pcuo
}

// SetNillableParentChoiceID sets the "parent_choice" edge to the ProficiencyChoice entity by ID if the given value is not nil.
func (pcuo *ProficiencyChoiceUpdateOne) SetNillableParentChoiceID(id *int) *ProficiencyChoiceUpdateOne {
	if id != nil {
		pcuo = pcuo.SetParentChoiceID(*id)
	}
	return pcuo
}

// SetParentChoice sets the "parent_choice" edge to the ProficiencyChoice entity.
func (pcuo *ProficiencyChoiceUpdateOne) SetParentChoice(p *ProficiencyChoice) *ProficiencyChoiceUpdateOne {
	return pcuo.SetParentChoiceID(p.ID)
}

// AddSubChoiceIDs adds the "sub_choice" edge to the ProficiencyChoice entity by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) AddSubChoiceIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.AddSubChoiceIDs(ids...)
	return pcuo
}

// AddSubChoice adds the "sub_choice" edges to the ProficiencyChoice entity.
func (pcuo *ProficiencyChoiceUpdateOne) AddSubChoice(p ...*ProficiencyChoice) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.AddSubChoiceIDs(ids...)
}

// AddClasIDs adds the "class" edge to the Class entity by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) AddClasIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.AddClasIDs(ids...)
	return pcuo
}

// AddClass adds the "class" edges to the Class entity.
func (pcuo *ProficiencyChoiceUpdateOne) AddClass(c ...*Class) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pcuo.AddClasIDs(ids...)
}

// AddRaceIDs adds the "race" edge to the Race entity by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) AddRaceIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.AddRaceIDs(ids...)
	return pcuo
}

// AddRace adds the "race" edges to the Race entity.
func (pcuo *ProficiencyChoiceUpdateOne) AddRace(r ...*Race) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pcuo.AddRaceIDs(ids...)
}

// Mutation returns the ProficiencyChoiceMutation object of the builder.
func (pcuo *ProficiencyChoiceUpdateOne) Mutation() *ProficiencyChoiceMutation {
	return pcuo.mutation
}

// ClearProficiency clears all "proficiency" edges to the Proficiency entity.
func (pcuo *ProficiencyChoiceUpdateOne) ClearProficiency() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearProficiency()
	return pcuo
}

// RemoveProficiencyIDs removes the "proficiency" edge to Proficiency entities by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveProficiencyIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.RemoveProficiencyIDs(ids...)
	return pcuo
}

// RemoveProficiency removes "proficiency" edges to Proficiency entities.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveProficiency(p ...*Proficiency) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.RemoveProficiencyIDs(ids...)
}

// ClearParentChoice clears the "parent_choice" edge to the ProficiencyChoice entity.
func (pcuo *ProficiencyChoiceUpdateOne) ClearParentChoice() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearParentChoice()
	return pcuo
}

// ClearSubChoice clears all "sub_choice" edges to the ProficiencyChoice entity.
func (pcuo *ProficiencyChoiceUpdateOne) ClearSubChoice() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearSubChoice()
	return pcuo
}

// RemoveSubChoiceIDs removes the "sub_choice" edge to ProficiencyChoice entities by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveSubChoiceIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.RemoveSubChoiceIDs(ids...)
	return pcuo
}

// RemoveSubChoice removes "sub_choice" edges to ProficiencyChoice entities.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveSubChoice(p ...*ProficiencyChoice) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return pcuo.RemoveSubChoiceIDs(ids...)
}

// ClearClass clears all "class" edges to the Class entity.
func (pcuo *ProficiencyChoiceUpdateOne) ClearClass() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearClass()
	return pcuo
}

// RemoveClasIDs removes the "class" edge to Class entities by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveClasIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.RemoveClasIDs(ids...)
	return pcuo
}

// RemoveClass removes "class" edges to Class entities.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveClass(c ...*Class) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(c))
	for i := range c {
		ids[i] = c[i].ID
	}
	return pcuo.RemoveClasIDs(ids...)
}

// ClearRace clears all "race" edges to the Race entity.
func (pcuo *ProficiencyChoiceUpdateOne) ClearRace() *ProficiencyChoiceUpdateOne {
	pcuo.mutation.ClearRace()
	return pcuo
}

// RemoveRaceIDs removes the "race" edge to Race entities by IDs.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveRaceIDs(ids ...int) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.RemoveRaceIDs(ids...)
	return pcuo
}

// RemoveRace removes "race" edges to Race entities.
func (pcuo *ProficiencyChoiceUpdateOne) RemoveRace(r ...*Race) *ProficiencyChoiceUpdateOne {
	ids := make([]int, len(r))
	for i := range r {
		ids[i] = r[i].ID
	}
	return pcuo.RemoveRaceIDs(ids...)
}

// Where appends a list predicates to the ProficiencyChoiceUpdate builder.
func (pcuo *ProficiencyChoiceUpdateOne) Where(ps ...predicate.ProficiencyChoice) *ProficiencyChoiceUpdateOne {
	pcuo.mutation.Where(ps...)
	return pcuo
}

// Select allows selecting one or more fields (columns) of the returned entity.
// The default is selecting all fields defined in the entity schema.
func (pcuo *ProficiencyChoiceUpdateOne) Select(field string, fields ...string) *ProficiencyChoiceUpdateOne {
	pcuo.fields = append([]string{field}, fields...)
	return pcuo
}

// Save executes the query and returns the updated ProficiencyChoice entity.
func (pcuo *ProficiencyChoiceUpdateOne) Save(ctx context.Context) (*ProficiencyChoice, error) {
	return withHooks(ctx, pcuo.sqlSave, pcuo.mutation, pcuo.hooks)
}

// SaveX is like Save, but panics if an error occurs.
func (pcuo *ProficiencyChoiceUpdateOne) SaveX(ctx context.Context) *ProficiencyChoice {
	node, err := pcuo.Save(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// Exec executes the query on the entity.
func (pcuo *ProficiencyChoiceUpdateOne) Exec(ctx context.Context) error {
	_, err := pcuo.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (pcuo *ProficiencyChoiceUpdateOne) ExecX(ctx context.Context) {
	if err := pcuo.Exec(ctx); err != nil {
		panic(err)
	}
}

func (pcuo *ProficiencyChoiceUpdateOne) sqlSave(ctx context.Context) (_node *ProficiencyChoice, err error) {
	_spec := sqlgraph.NewUpdateSpec(proficiencychoice.Table, proficiencychoice.Columns, sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt))
	id, ok := pcuo.mutation.ID()
	if !ok {
		return nil, &ValidationError{Name: "id", err: errors.New(`ent: missing "ProficiencyChoice.id" for update`)}
	}
	_spec.Node.ID.Value = id
	if fields := pcuo.fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, proficiencychoice.FieldID)
		for _, f := range fields {
			if !proficiencychoice.ValidColumn(f) {
				return nil, &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
			}
			if f != proficiencychoice.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, f)
			}
		}
	}
	if ps := pcuo.mutation.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if value, ok := pcuo.mutation.Choose(); ok {
		_spec.SetField(proficiencychoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.AddedChoose(); ok {
		_spec.AddField(proficiencychoice.FieldChoose, field.TypeInt, value)
	}
	if value, ok := pcuo.mutation.Desc(); ok {
		_spec.SetField(proficiencychoice.FieldDesc, field.TypeString, value)
	}
	if pcuo.mutation.DescCleared() {
		_spec.ClearField(proficiencychoice.FieldDesc, field.TypeString)
	}
	if pcuo.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiency.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedProficiencyIDs(); len(nodes) > 0 && !pcuo.mutation.ProficiencyCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
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
	if nodes := pcuo.mutation.ProficiencyIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ProficiencyTable,
			Columns: proficiencychoice.ProficiencyPrimaryKey,
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
	if pcuo.mutation.ParentChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proficiencychoice.ParentChoiceTable,
			Columns: []string{proficiencychoice.ParentChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ParentChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2O,
			Inverse: true,
			Table:   proficiencychoice.ParentChoiceTable,
			Columns: []string{proficiencychoice.ParentChoiceColumn},
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
	if pcuo.mutation.SubChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(proficiencychoice.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedSubChoiceIDs(); len(nodes) > 0 && !pcuo.mutation.SubChoiceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
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
	if nodes := pcuo.mutation.SubChoiceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   proficiencychoice.SubChoiceTable,
			Columns: []string{proficiencychoice.SubChoiceColumn},
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
	if pcuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedClassIDs(); len(nodes) > 0 && !pcuo.mutation.ClassCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.ClassIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.ClassTable,
			Columns: proficiencychoice.ClassPrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(class.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	if pcuo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RemovedRaceIDs(); len(nodes) > 0 && !pcuo.mutation.RaceCleared() {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Clear = append(_spec.Edges.Clear, edge)
	}
	if nodes := pcuo.mutation.RaceIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.M2M,
			Inverse: true,
			Table:   proficiencychoice.RaceTable,
			Columns: proficiencychoice.RacePrimaryKey,
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges.Add = append(_spec.Edges.Add, edge)
	}
	_node = &ProficiencyChoice{config: pcuo.config}
	_spec.Assign = _node.assignValues
	_spec.ScanValues = _node.scanValues
	if err = sqlgraph.UpdateNode(ctx, pcuo.driver, _spec); err != nil {
		if _, ok := err.(*sqlgraph.NotFoundError); ok {
			err = &NotFoundError{proficiencychoice.Label}
		} else if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	pcuo.mutation.done = true
	return _node, nil
}