// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"errors"
	"fmt"

	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/feature"
	"github.com/ecshreve/dndgen/ent/prerequisite"
)

// FeatureCreate is the builder for creating a Feature entity.
type FeatureCreate struct {
	config
	mutation *FeatureMutation
	hooks    []Hook
}

// SetIndx sets the "indx" field.
func (fc *FeatureCreate) SetIndx(s string) *FeatureCreate {
	fc.mutation.SetIndx(s)
	return fc
}

// SetName sets the "name" field.
func (fc *FeatureCreate) SetName(s string) *FeatureCreate {
	fc.mutation.SetName(s)
	return fc
}

// SetDesc sets the "desc" field.
func (fc *FeatureCreate) SetDesc(s []string) *FeatureCreate {
	fc.mutation.SetDesc(s)
	return fc
}

// SetLevel sets the "level" field.
func (fc *FeatureCreate) SetLevel(i int) *FeatureCreate {
	fc.mutation.SetLevel(i)
	return fc
}

// AddPrerequisiteIDs adds the "prerequisites" edge to the Prerequisite entity by IDs.
func (fc *FeatureCreate) AddPrerequisiteIDs(ids ...int) *FeatureCreate {
	fc.mutation.AddPrerequisiteIDs(ids...)
	return fc
}

// AddPrerequisites adds the "prerequisites" edges to the Prerequisite entity.
func (fc *FeatureCreate) AddPrerequisites(p ...*Prerequisite) *FeatureCreate {
	ids := make([]int, len(p))
	for i := range p {
		ids[i] = p[i].ID
	}
	return fc.AddPrerequisiteIDs(ids...)
}

// Mutation returns the FeatureMutation object of the builder.
func (fc *FeatureCreate) Mutation() *FeatureMutation {
	return fc.mutation
}

// Save creates the Feature in the database.
func (fc *FeatureCreate) Save(ctx context.Context) (*Feature, error) {
	return withHooks(ctx, fc.sqlSave, fc.mutation, fc.hooks)
}

// SaveX calls Save and panics if Save returns an error.
func (fc *FeatureCreate) SaveX(ctx context.Context) *Feature {
	v, err := fc.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fc *FeatureCreate) Exec(ctx context.Context) error {
	_, err := fc.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fc *FeatureCreate) ExecX(ctx context.Context) {
	if err := fc.Exec(ctx); err != nil {
		panic(err)
	}
}

// check runs all checks and user-defined validators on the builder.
func (fc *FeatureCreate) check() error {
	if _, ok := fc.mutation.Indx(); !ok {
		return &ValidationError{Name: "indx", err: errors.New(`ent: missing required field "Feature.indx"`)}
	}
	if v, ok := fc.mutation.Indx(); ok {
		if err := feature.IndxValidator(v); err != nil {
			return &ValidationError{Name: "indx", err: fmt.Errorf(`ent: validator failed for field "Feature.indx": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Name(); !ok {
		return &ValidationError{Name: "name", err: errors.New(`ent: missing required field "Feature.name"`)}
	}
	if v, ok := fc.mutation.Name(); ok {
		if err := feature.NameValidator(v); err != nil {
			return &ValidationError{Name: "name", err: fmt.Errorf(`ent: validator failed for field "Feature.name": %w`, err)}
		}
	}
	if _, ok := fc.mutation.Level(); !ok {
		return &ValidationError{Name: "level", err: errors.New(`ent: missing required field "Feature.level"`)}
	}
	if v, ok := fc.mutation.Level(); ok {
		if err := feature.LevelValidator(v); err != nil {
			return &ValidationError{Name: "level", err: fmt.Errorf(`ent: validator failed for field "Feature.level": %w`, err)}
		}
	}
	return nil
}

func (fc *FeatureCreate) sqlSave(ctx context.Context) (*Feature, error) {
	if err := fc.check(); err != nil {
		return nil, err
	}
	_node, _spec := fc.createSpec()
	if err := sqlgraph.CreateNode(ctx, fc.driver, _spec); err != nil {
		if sqlgraph.IsConstraintError(err) {
			err = &ConstraintError{msg: err.Error(), wrap: err}
		}
		return nil, err
	}
	id := _spec.ID.Value.(int64)
	_node.ID = int(id)
	fc.mutation.id = &_node.ID
	fc.mutation.done = true
	return _node, nil
}

func (fc *FeatureCreate) createSpec() (*Feature, *sqlgraph.CreateSpec) {
	var (
		_node = &Feature{config: fc.config}
		_spec = sqlgraph.NewCreateSpec(feature.Table, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeInt))
	)
	if value, ok := fc.mutation.Indx(); ok {
		_spec.SetField(feature.FieldIndx, field.TypeString, value)
		_node.Indx = value
	}
	if value, ok := fc.mutation.Name(); ok {
		_spec.SetField(feature.FieldName, field.TypeString, value)
		_node.Name = value
	}
	if value, ok := fc.mutation.Desc(); ok {
		_spec.SetField(feature.FieldDesc, field.TypeJSON, value)
		_node.Desc = value
	}
	if value, ok := fc.mutation.Level(); ok {
		_spec.SetField(feature.FieldLevel, field.TypeInt, value)
		_node.Level = value
	}
	if nodes := fc.mutation.PrerequisitesIDs(); len(nodes) > 0 {
		edge := &sqlgraph.EdgeSpec{
			Rel:     sqlgraph.O2M,
			Inverse: false,
			Table:   feature.PrerequisitesTable,
			Columns: []string{feature.PrerequisitesColumn},
			Bidi:    false,
			Target: &sqlgraph.EdgeTarget{
				IDSpec: sqlgraph.NewFieldSpec(prerequisite.FieldID, field.TypeInt),
			},
		}
		for _, k := range nodes {
			edge.Target.Nodes = append(edge.Target.Nodes, k)
		}
		_spec.Edges = append(_spec.Edges, edge)
	}
	return _node, _spec
}

// FeatureCreateBulk is the builder for creating many Feature entities in bulk.
type FeatureCreateBulk struct {
	config
	err      error
	builders []*FeatureCreate
}

// Save creates the Feature entities in the database.
func (fcb *FeatureCreateBulk) Save(ctx context.Context) ([]*Feature, error) {
	if fcb.err != nil {
		return nil, fcb.err
	}
	specs := make([]*sqlgraph.CreateSpec, len(fcb.builders))
	nodes := make([]*Feature, len(fcb.builders))
	mutators := make([]Mutator, len(fcb.builders))
	for i := range fcb.builders {
		func(i int, root context.Context) {
			builder := fcb.builders[i]
			var mut Mutator = MutateFunc(func(ctx context.Context, m Mutation) (Value, error) {
				mutation, ok := m.(*FeatureMutation)
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
					_, err = mutators[i+1].Mutate(root, fcb.builders[i+1].mutation)
				} else {
					spec := &sqlgraph.BatchCreateSpec{Nodes: specs}
					// Invoke the actual operation on the latest mutation in the chain.
					if err = sqlgraph.BatchCreate(ctx, fcb.driver, spec); err != nil {
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
		if _, err := mutators[0].Mutate(ctx, fcb.builders[0].mutation); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

// SaveX is like Save, but panics if an error occurs.
func (fcb *FeatureCreateBulk) SaveX(ctx context.Context) []*Feature {
	v, err := fcb.Save(ctx)
	if err != nil {
		panic(err)
	}
	return v
}

// Exec executes the query.
func (fcb *FeatureCreateBulk) Exec(ctx context.Context) error {
	_, err := fcb.Save(ctx)
	return err
}

// ExecX is like Exec, but panics if an error occurs.
func (fcb *FeatureCreateBulk) ExecX(ctx context.Context) {
	if err := fcb.Exec(ctx); err != nil {
		panic(err)
	}
}
