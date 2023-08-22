// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/armorclass"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// ArmorClassQuery is the builder for querying ArmorClass entities.
type ArmorClassQuery struct {
	config
	ctx        *QueryContext
	order      []armorclass.OrderOption
	inters     []Interceptor
	predicates []predicate.ArmorClass
	withFKs    bool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*ArmorClass) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the ArmorClassQuery builder.
func (acq *ArmorClassQuery) Where(ps ...predicate.ArmorClass) *ArmorClassQuery {
	acq.predicates = append(acq.predicates, ps...)
	return acq
}

// Limit the number of records to be returned by this query.
func (acq *ArmorClassQuery) Limit(limit int) *ArmorClassQuery {
	acq.ctx.Limit = &limit
	return acq
}

// Offset to start from.
func (acq *ArmorClassQuery) Offset(offset int) *ArmorClassQuery {
	acq.ctx.Offset = &offset
	return acq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (acq *ArmorClassQuery) Unique(unique bool) *ArmorClassQuery {
	acq.ctx.Unique = &unique
	return acq
}

// Order specifies how the records should be ordered.
func (acq *ArmorClassQuery) Order(o ...armorclass.OrderOption) *ArmorClassQuery {
	acq.order = append(acq.order, o...)
	return acq
}

// First returns the first ArmorClass entity from the query.
// Returns a *NotFoundError when no ArmorClass was found.
func (acq *ArmorClassQuery) First(ctx context.Context) (*ArmorClass, error) {
	nodes, err := acq.Limit(1).All(setContextOp(ctx, acq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{armorclass.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (acq *ArmorClassQuery) FirstX(ctx context.Context) *ArmorClass {
	node, err := acq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first ArmorClass ID from the query.
// Returns a *NotFoundError when no ArmorClass ID was found.
func (acq *ArmorClassQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acq.Limit(1).IDs(setContextOp(ctx, acq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{armorclass.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (acq *ArmorClassQuery) FirstIDX(ctx context.Context) int {
	id, err := acq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single ArmorClass entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one ArmorClass entity is found.
// Returns a *NotFoundError when no ArmorClass entities are found.
func (acq *ArmorClassQuery) Only(ctx context.Context) (*ArmorClass, error) {
	nodes, err := acq.Limit(2).All(setContextOp(ctx, acq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{armorclass.Label}
	default:
		return nil, &NotSingularError{armorclass.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (acq *ArmorClassQuery) OnlyX(ctx context.Context) *ArmorClass {
	node, err := acq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only ArmorClass ID in the query.
// Returns a *NotSingularError when more than one ArmorClass ID is found.
// Returns a *NotFoundError when no entities are found.
func (acq *ArmorClassQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = acq.Limit(2).IDs(setContextOp(ctx, acq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{armorclass.Label}
	default:
		err = &NotSingularError{armorclass.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (acq *ArmorClassQuery) OnlyIDX(ctx context.Context) int {
	id, err := acq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of ArmorClasses.
func (acq *ArmorClassQuery) All(ctx context.Context) ([]*ArmorClass, error) {
	ctx = setContextOp(ctx, acq.ctx, "All")
	if err := acq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*ArmorClass, *ArmorClassQuery]()
	return withInterceptors[[]*ArmorClass](ctx, acq, qr, acq.inters)
}

// AllX is like All, but panics if an error occurs.
func (acq *ArmorClassQuery) AllX(ctx context.Context) []*ArmorClass {
	nodes, err := acq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of ArmorClass IDs.
func (acq *ArmorClassQuery) IDs(ctx context.Context) (ids []int, err error) {
	if acq.ctx.Unique == nil && acq.path != nil {
		acq.Unique(true)
	}
	ctx = setContextOp(ctx, acq.ctx, "IDs")
	if err = acq.Select(armorclass.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (acq *ArmorClassQuery) IDsX(ctx context.Context) []int {
	ids, err := acq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (acq *ArmorClassQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, acq.ctx, "Count")
	if err := acq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, acq, querierCount[*ArmorClassQuery](), acq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (acq *ArmorClassQuery) CountX(ctx context.Context) int {
	count, err := acq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (acq *ArmorClassQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, acq.ctx, "Exist")
	switch _, err := acq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (acq *ArmorClassQuery) ExistX(ctx context.Context) bool {
	exist, err := acq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the ArmorClassQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (acq *ArmorClassQuery) Clone() *ArmorClassQuery {
	if acq == nil {
		return nil
	}
	return &ArmorClassQuery{
		config:     acq.config,
		ctx:        acq.ctx.Clone(),
		order:      append([]armorclass.OrderOption{}, acq.order...),
		inters:     append([]Interceptor{}, acq.inters...),
		predicates: append([]predicate.ArmorClass{}, acq.predicates...),
		// clone intermediate query.
		sql:  acq.sql.Clone(),
		path: acq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Base int `json:"base,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.ArmorClass.Query().
//		GroupBy(armorclass.FieldBase).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (acq *ArmorClassQuery) GroupBy(field string, fields ...string) *ArmorClassGroupBy {
	acq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &ArmorClassGroupBy{build: acq}
	grbuild.flds = &acq.ctx.Fields
	grbuild.label = armorclass.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Base int `json:"base,omitempty"`
//	}
//
//	client.ArmorClass.Query().
//		Select(armorclass.FieldBase).
//		Scan(ctx, &v)
func (acq *ArmorClassQuery) Select(fields ...string) *ArmorClassSelect {
	acq.ctx.Fields = append(acq.ctx.Fields, fields...)
	sbuild := &ArmorClassSelect{ArmorClassQuery: acq}
	sbuild.label = armorclass.Label
	sbuild.flds, sbuild.scan = &acq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a ArmorClassSelect configured with the given aggregations.
func (acq *ArmorClassQuery) Aggregate(fns ...AggregateFunc) *ArmorClassSelect {
	return acq.Select().Aggregate(fns...)
}

func (acq *ArmorClassQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range acq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, acq); err != nil {
				return err
			}
		}
	}
	for _, f := range acq.ctx.Fields {
		if !armorclass.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if acq.path != nil {
		prev, err := acq.path(ctx)
		if err != nil {
			return err
		}
		acq.sql = prev
	}
	return nil
}

func (acq *ArmorClassQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*ArmorClass, error) {
	var (
		nodes   = []*ArmorClass{}
		withFKs = acq.withFKs
		_spec   = acq.querySpec()
	)
	if withFKs {
		_spec.Node.Columns = append(_spec.Node.Columns, armorclass.ForeignKeys...)
	}
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*ArmorClass).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &ArmorClass{config: acq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, acq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range acq.loadTotal {
		if err := acq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (acq *ArmorClassQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := acq.querySpec()
	if len(acq.modifiers) > 0 {
		_spec.Modifiers = acq.modifiers
	}
	_spec.Node.Columns = acq.ctx.Fields
	if len(acq.ctx.Fields) > 0 {
		_spec.Unique = acq.ctx.Unique != nil && *acq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, acq.driver, _spec)
}

func (acq *ArmorClassQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(armorclass.Table, armorclass.Columns, sqlgraph.NewFieldSpec(armorclass.FieldID, field.TypeInt))
	_spec.From = acq.sql
	if unique := acq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if acq.path != nil {
		_spec.Unique = true
	}
	if fields := acq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, armorclass.FieldID)
		for i := range fields {
			if fields[i] != armorclass.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := acq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := acq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := acq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := acq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (acq *ArmorClassQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(acq.driver.Dialect())
	t1 := builder.Table(armorclass.Table)
	columns := acq.ctx.Fields
	if len(columns) == 0 {
		columns = armorclass.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if acq.sql != nil {
		selector = acq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if acq.ctx.Unique != nil && *acq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range acq.predicates {
		p(selector)
	}
	for _, p := range acq.order {
		p(selector)
	}
	if offset := acq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := acq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// ArmorClassGroupBy is the group-by builder for ArmorClass entities.
type ArmorClassGroupBy struct {
	selector
	build *ArmorClassQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (acgb *ArmorClassGroupBy) Aggregate(fns ...AggregateFunc) *ArmorClassGroupBy {
	acgb.fns = append(acgb.fns, fns...)
	return acgb
}

// Scan applies the selector query and scans the result into the given value.
func (acgb *ArmorClassGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acgb.build.ctx, "GroupBy")
	if err := acgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmorClassQuery, *ArmorClassGroupBy](ctx, acgb.build, acgb, acgb.build.inters, v)
}

func (acgb *ArmorClassGroupBy) sqlScan(ctx context.Context, root *ArmorClassQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(acgb.fns))
	for _, fn := range acgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*acgb.flds)+len(acgb.fns))
		for _, f := range *acgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*acgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// ArmorClassSelect is the builder for selecting fields of ArmorClass entities.
type ArmorClassSelect struct {
	*ArmorClassQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (acs *ArmorClassSelect) Aggregate(fns ...AggregateFunc) *ArmorClassSelect {
	acs.fns = append(acs.fns, fns...)
	return acs
}

// Scan applies the selector query and scans the result into the given value.
func (acs *ArmorClassSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, acs.ctx, "Select")
	if err := acs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*ArmorClassQuery, *ArmorClassSelect](ctx, acs.ArmorClassQuery, acs, acs.inters, v)
}

func (acs *ArmorClassSelect) sqlScan(ctx context.Context, root *ArmorClassQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(acs.fns))
	for _, fn := range acs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*acs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := acs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}