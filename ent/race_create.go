// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilitybonus"
	"github.com/ecshreve/dndgen/ent/race"
)

// RaceCreate is the builder for creating a Race entity.
type RaceCreate struct {
	config
	mutation *RaceMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (rc *RaceCreate) SetIndx(s string) *RaceCreate {
	rc.mutation.SetIndx(s)
	return rc
}

// SetName sets the "name" field.
func (rc *RaceCreate) SetName(s string) *RaceCreate {
	rc.mutation.SetName(s)
	return rc
}

// SetSpeed sets the "speed" field.
func (rc *RaceCreate) SetSpeed(i int) *RaceCreate {
	rc.mutation.SetSpeed(i)
	return rc
}

// SetSize sets the "size" field.
func (rc *RaceCreate) SetSize(r race.Size) *RaceCreate {
	rc.mutation.SetSize(r)
	return rc
}

// SetNillableSize sets the "size" field if the given value is not nil.
func (rc *RaceCreate) SetNillableSize(r *race.Size) *RaceCreate {
	if r != nil {
		rc.SetSize(*r)
	}
	return rc
}

// SetSizeDesc sets the "size_desc" field.
func (rc *RaceCreate) SetSizeDesc(s string) *RaceCreate {
	rc.mutation.SetSizeDesc(s)
	return rc
}

// SetAlignmentDesc sets the "alignment_desc" field.
func (rc *RaceCreate) SetAlignmentDesc(s string) *RaceCreate {
	rc.mutation.SetAlignmentDesc(s)
	return rc
}

// SetAgeDesc sets the "age_desc" field.
func (rc *RaceCreate) SetAgeDesc(s string) *RaceCreate {
	rc.mutation.SetAgeDesc(s)
	return rc
}

// SetLanguageDesc sets the "language_desc" field.
func (rc *RaceCreate) SetLanguageDesc(s string) *RaceCreate {
	rc.mutation.SetLanguageDesc(s)
	return rc
}

// AddAbilityBonuseIDs adds the "ability_bonuses" edge to the AbilityBonus entity by IDs.
func (rc *RaceCreate) AddAbilityBonuseIDs(ids ...int) *RaceCreate {
	rc.mutation.AddAbilityBonuseIDs(ids...)
	return rc
}

// AddAbilityBonuses adds the "ability_bonuses" edges to the AbilityBonus entity.
func (rc *RaceCreate) AddAbilityBonuses(a ...*AbilityBonus) *RaceCreate {
	ids := make([]int, len(a))
	for i := range a {
		ids[i] = a[i].ID
	}
	return rc.AddAbilityBonuseIDs(ids...)
}

// Mutation returns the RaceMutation object of the builder.
func (rc *RaceCreate) Mutation() *RaceMutation {
	return rc.mutation
}

// Save creates the Race in the database.
func (rc *RaceCreate) Save(ctx context.Context) (*Race, error) {
	rc.defaults()
	return withHooks(ctx, rc.sqlSave, rc.mutation, rc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (rc *RaceCreate) SaveX(ctx context.Context) *Race {
	v, err := rc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rc *RaceCreate) Exec(ctx context.Context) error {
	_, err := rc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rc *RaceCreate) ExecX(ctx context.Context) {
	if err := rc.Exec(ctx); err != nil {
		panic(err)
	}
}

// defaults sets the default values of the builder before save.
func (rc *RaceCreate) defaults() {
	if _, ok := rc.mutation.Size(); !ok {
		v := race.DefaultSize
		rc.mutation.SetSize(v)
	}
}

// check runs all checks and user-defined validators on the builder.
func (rc *RaceCreate) check() error {
	if _, ok := rc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Race.indx"`)}
	}
	if v, ok := rc.mutation.Indx(); ok {
		if err := race.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Race.indx": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Race.name"`)}
	}
	if v, ok := rc.mutation.Name(); ok {
		if err := race.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Race.name": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Speed(); !ok {
		return &ValidationError{Name: "speed", err: errors.New(`ent: missing required field "Race.speed"`)}
	}
	if v, ok := rc.mutation.Speed(); ok {
		if err := race.SpeedValidator(v); err != nil {
			return &ValidationError{Name: "speed", err: fmt.Errorf(`ent: validator failed for field "Race.speed": %w`, err)}
		}
	}
	if _, ok := rc.mutation.Size(); !ok {
		return &ValidationError{Name: "size", err: errors.New(`ent: missing required field "Race.size"`)}
	}
	if v, ok := rc.mutation.Size(); ok {
		if err := race.SizeValidator(v); err != nil {
			return &ValidationError{Name: "size", err: fmt.Errorf(`ent: validator failed for field "Race.size": %w`, err)}
		}
	}
	if _, ok := rc.mutation.SizeDesc(); !ok {
		return &ValidationError{Name: "size_desc", err: errors.New(`ent: missing required field "Race.size_desc"`)}
	}
	if _, ok := rc.mutation.AlignmentDesc(); !ok {
		return &ValidationError{Name: "alignment_desc", err: errors.New(`ent: missing required field "Race.alignment_desc"`)}
	}
	if _, ok := rc.mutation.AgeDesc(); !ok {
		return &ValidationError{Name: "age_desc", err: errors.New(`ent: missing required field "Race.age_desc"`)}
	}
	if _, ok := rc.mutation.LanguageDesc(); !ok {
		return &ValidationError{Name: "language_desc", err: errors.New(`ent: missing required field "Race.language_desc"`)}
	}
	return nil
}

func (rc *RaceCreate) sqlSave(ctx context.Context) (*Race, error) {
	if err := rc.check(); err != nil {
		return nil, err
	}
	_node, _spec := rc.createSpec()
	if err := sqlgraph.CreateNode(ctx, rc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	rc.mutation.id = &_node.ID
	rc.mutation.done = true
	return _node, nil
}

func (rc *RaceCreate) createSpec() (*Race, *sqlgraph.CreateSpec) {
	var (
		_node = &Race{config: rc.config}
		_spec = sqlgraph.NewCreateSpec(race.Table, sqlgraph.NewFieldSpec(race.FieldID, field.TypeInt))
	)
	if value, ok := rc.mutation.Indx(); ok {
		_spec.SetField(race.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := rc.mutation.Name(); ok {
		_spec.SetField(race.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := rc.mutation.Speed(); ok {
		_spec.SetField(race.FieldSpeed, field.TypeInt, value)
		_node.Speed = value
	}
	if value, ok := rc.mutation.Size(); ok {
		_spec.SetField(race.FieldSize, field.TypeEnum, value)
		_node.Size = value
	}
	if value, ok := rc.mutation.SizeDesc(); ok {
		_spec.SetField(race.FieldSizeDesc, field.TypeString, value)
		_node.SizeDesc = value
	}
	if value, ok := rc.mutation.AlignmentDesc(); ok {
		_spec.SetField(race.FieldAlignmentDesc, field.TypeString, value)
		_node.AlignmentDesc = value
	}
	if value, ok := rc.mutation.AgeDesc(); ok {
		_spec.SetField(race.FieldAgeDesc, field.TypeString, value)
		_node.AgeDesc = value
	}
	if value, ok := rc.mutation.LanguageDesc(); ok {
		_spec.SetField(race.FieldLanguageDesc, field.TypeString, value)
		_node.LanguageDesc = value
	}
	if nodes := rc.mutation.AbilityBonusesIDs(); len(nodes) > 0 {
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
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// RaceCreateBulk is the builder for creating many Race entities in bulk.
type RaceCreateBulk struct {
	config
	err      error
	builders []*RaceCreate
}

// Save creates the Race entities in the database.
func (rcb *RaceCreateBulk) Save(ctx context.Context) ([]*Race, error) {
	if rcb.err != nil {
		return nil, rcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(rcb.builders))
	nodes := make([]*Race, len(rcb.builders))
	mutators := make([]Mutator, len(rcb.builders))
	for i := range rcb.builders {
		func(i int, root context.Context) {
			builder := rcb.builders[i]
			builder.defaults()
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*RaceMutation)
				if !ok {
					return nil, fmt.Errorf("unexpected mutation type %T", m)
				}
				if err := builder.check(); err != nil {
					return nil, err
				}
				builder.mutation = mutation
				var err error
				nodes[i], specs[i] = builder.createSpec()
				if i < len(mutators)-1 {
					_, err = mutators[i+1].Mutate(root, rcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, rcb.driver, spec); err != nil {
						if sqlgraph.IsConstraintError(err) {
							err = &ConstraintError{msg: err.Error(), wrap: err}
						}
					}
				}
				if err != nil {
					return nil, err
				}
				mutation.id = &nodes[i].ID
				if specs[i].ID.Value != nil {
					id := specs[i].ID.Value.(int64)
					nodes[i].ID = int(id)
				}
				mutation.done = true
				return nodes[i], nil
			})
			for i := len(builder.hooks) - 1; i >= 0; i-- {
				mut = builder.hooks[i](mut)
			}
			mutators[i] = mut
		}(i, ctx)
	}
	if len(mutators) > 0 {
		if _, err := mutators[0].Mutate(ctx, rcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (rcb *RaceCreateBulk) SaveX(ctx context.Context) []*Race {
	v, err := rcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (rcb *RaceCreateBulk) Exec(ctx context.Context) error {
	_, err := rcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (rcb *RaceCreateBulk) ExecX(ctx context.Context) {
	if err := rcb.Exec(ctx); err != nil {
		panic(err)
	}
}
