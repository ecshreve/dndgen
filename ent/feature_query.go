// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"fmt"
	"math"

	"entgo.io/ent"
	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/feature"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// FeatureQuery is the builder for querying Feature entities.
type FeatureQuery struct {
	config
	ctx        *QueryContext
	order      []feature.OrderOption
	inters     []Interceptor
	predicates []predicate.Feature
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*Feature) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the FeatureQuery builder.
func (fq *FeatureQuery) Where(ps ...predicate.Feature) *FeatureQuery {
	fq.predicates = append(fq.predicates, ps...)
	return fq
}

// Limit the number of records to be returned by this query.
func (fq *FeatureQuery) Limit(limit int) *FeatureQuery {
	fq.ctx.Limit = &limit
	return fq
}

// Offset to start from.
func (fq *FeatureQuery) Offset(offset int) *FeatureQuery {
	fq.ctx.Offset = &offset
	return fq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (fq *FeatureQuery) Unique(unique bool) *FeatureQuery {
	fq.ctx.Unique = &unique
	return fq
}

// Order specifies how the records should be ordered.
func (fq *FeatureQuery) Order(o ...feature.OrderOption) *FeatureQuery {
	fq.order = append(fq.order, o...)
	return fq
}

// First returns the first Feature entity from the query.
// Returns a *NotFoundError when no Feature was found.
func (fq *FeatureQuery) First(ctx context.Context) (*Feature, error) {
	nodes, err := fq.Limit(1).All(setContextOp(ctx, fq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{feature.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (fq *FeatureQuery) FirstX(ctx context.Context) *Feature {
	node, err := fq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Feature ID from the query.
// Returns a *NotFoundError when no Feature ID was found.
func (fq *FeatureQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(1).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{feature.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (fq *FeatureQuery) FirstIDX(ctx context.Context) int {
	id, err := fq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Feature entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Feature entity is found.
// Returns a *NotFoundError when no Feature entities are found.
func (fq *FeatureQuery) Only(ctx context.Context) (*Feature, error) {
	nodes, err := fq.Limit(2).All(setContextOp(ctx, fq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{feature.Label}
	default:
		return nil, &NotSingularError{feature.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (fq *FeatureQuery) OnlyX(ctx context.Context) *Feature {
	node, err := fq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Feature ID in the query.
// Returns a *NotSingularError when more than one Feature ID is found.
// Returns a *NotFoundError when no entities are found.
func (fq *FeatureQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = fq.Limit(2).IDs(setContextOp(ctx, fq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{feature.Label}
	default:
		err = &NotSingularError{feature.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (fq *FeatureQuery) OnlyIDX(ctx context.Context) int {
	id, err := fq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Features.
func (fq *FeatureQuery) All(ctx context.Context) ([]*Feature, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryAll)
	if err := fq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Feature, *FeatureQuery]()
	return withInterceptors[[]*Feature](ctx, fq, qr, fq.inters)
}

// AllX is like All, but panics if an error occurs.
func (fq *FeatureQuery) AllX(ctx context.Context) []*Feature {
	nodes, err := fq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Feature IDs.
func (fq *FeatureQuery) IDs(ctx context.Context) (ids []int, err error) {
	if fq.ctx.Unique == nil && fq.path != nil {
		fq.Unique(true)
	}
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryIDs)
	if err = fq.Select(feature.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (fq *FeatureQuery) IDsX(ctx context.Context) []int {
	ids, err := fq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (fq *FeatureQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryCount)
	if err := fq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, fq, querierCount[*FeatureQuery](), fq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (fq *FeatureQuery) CountX(ctx context.Context) int {
	count, err := fq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (fq *FeatureQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, fq.ctx, ent.OpQueryExist)
	switch _, err := fq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (fq *FeatureQuery) ExistX(ctx context.Context) bool {
	exist, err := fq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the FeatureQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (fq *FeatureQuery) Clone() *FeatureQuery {
	if fq == nil {
		return nil
	}
	return &FeatureQuery{
		config:     fq.config,
		ctx:        fq.ctx.Clone(),
		order:      append([]feature.OrderOption{}, fq.order...),
		inters:     append([]Interceptor{}, fq.inters...),
		predicates: append([]predicate.Feature{}, fq.predicates...),
		// clone intermediate query.
		sql:  fq.sql.Clone(),
		path: fq.path,
	}
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"index"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Feature.Query().
//		GroupBy(feature.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (fq *FeatureQuery) GroupBy(field string, fields ...string) *FeatureGroupBy {
	fq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &FeatureGroupBy{build: fq}
	grbuild.flds = &fq.ctx.Fields
	grbuild.label = feature.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Indx string `json:"index"`
//	}
//
//	client.Feature.Query().
//		Select(feature.FieldIndx).
//		Scan(ctx, &v)
func (fq *FeatureQuery) Select(fields ...string) *FeatureSelect {
	fq.ctx.Fields = append(fq.ctx.Fields, fields...)
	sbuild := &FeatureSelect{FeatureQuery: fq}
	sbuild.label = feature.Label
	sbuild.flds, sbuild.scan = &fq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a FeatureSelect configured with the given aggregations.
func (fq *FeatureQuery) Aggregate(fns ...AggregateFunc) *FeatureSelect {
	return fq.Select().Aggregate(fns...)
}

func (fq *FeatureQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range fq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, fq); err != nil {
				return err
			}
		}
	}
	for _, f := range fq.ctx.Fields {
		if !feature.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if fq.path != nil {
		prev, err := fq.path(ctx)
		if err != nil {
			return err
		}
		fq.sql = prev
	}
	return nil
}

func (fq *FeatureQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Feature, error) {
	var (
		nodes = []*Feature{}
		_spec = fq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Feature).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Feature{config: fq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, fq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range fq.loadTotal {
		if err := fq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (fq *FeatureQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := fq.querySpec()
	if len(fq.modifiers) > 0 {
		_spec.Modifiers = fq.modifiers
	}
	_spec.Node.Columns = fq.ctx.Fields
	if len(fq.ctx.Fields) > 0 {
		_spec.Unique = fq.ctx.Unique != nil && *fq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, fq.driver, _spec)
}

func (fq *FeatureQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(feature.Table, feature.Columns, sqlgraph.NewFieldSpec(feature.FieldID, field.TypeInt))
	_spec.From = fq.sql
	if unique := fq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if fq.path != nil {
		_spec.Unique = true
	}
	if fields := fq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, feature.FieldID)
		for i := range fields {
			if fields[i] != feature.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := fq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := fq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := fq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := fq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (fq *FeatureQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(fq.driver.Dialect())
	t1 := builder.Table(feature.Table)
	columns := fq.ctx.Fields
	if len(columns) == 0 {
		columns = feature.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if fq.sql != nil {
		selector = fq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if fq.ctx.Unique != nil && *fq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range fq.predicates {
		p(selector)
	}
	for _, p := range fq.order {
		p(selector)
	}
	if offset := fq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := fq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// FeatureGroupBy is the group-by builder for Feature entities.
type FeatureGroupBy struct {
	selector
	build *FeatureQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (fgb *FeatureGroupBy) Aggregate(fns ...AggregateFunc) *FeatureGroupBy {
	fgb.fns = append(fgb.fns, fns...)
	return fgb
}

// Scan applies the selector query and scans the result into the given value.
func (fgb *FeatureGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fgb.build.ctx, ent.OpQueryGroupBy)
	if err := fgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeatureQuery, *FeatureGroupBy](ctx, fgb.build, fgb, fgb.build.inters, v)
}

func (fgb *FeatureGroupBy) sqlScan(ctx context.Context, root *FeatureQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(fgb.fns))
	for _, fn := range fgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*fgb.flds)+len(fgb.fns))
		for _, f := range *fgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*fgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// FeatureSelect is the builder for selecting fields of Feature entities.
type FeatureSelect struct {
	*FeatureQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (fs *FeatureSelect) Aggregate(fns ...AggregateFunc) *FeatureSelect {
	fs.fns = append(fs.fns, fns...)
	return fs
}

// Scan applies the selector query and scans the result into the given value.
func (fs *FeatureSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, fs.ctx, ent.OpQuerySelect)
	if err := fs.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*FeatureQuery, *FeatureSelect](ctx, fs.FeatureQuery, fs, fs.inters, v)
}

func (fs *FeatureSelect) sqlScan(ctx context.Context, root *FeatureQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(fs.fns))
	for _, fn := range fs.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*fs.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := fs.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
