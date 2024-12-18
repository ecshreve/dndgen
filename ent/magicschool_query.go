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
	"github.com/ecshreve/dndgen/ent/magicschool"
	"github.com/ecshreve/dndgen/ent/predicate"
)

// MagicSchoolQuery is the builder for querying MagicSchool entities.
type MagicSchoolQuery struct {
	config
	ctx        *QueryContext
	order      []magicschool.OrderOption
	inters     []Interceptor
	predicates []predicate.MagicSchool
	modifiers  []func(*sql.Selector)
	loadTotal  []func(context.Context, []*MagicSchool) error
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the MagicSchoolQuery builder.
func (msq *MagicSchoolQuery) Where(ps ...predicate.MagicSchool) *MagicSchoolQuery {
	msq.predicates = append(msq.predicates, ps...)
	return msq
}

// Limit the number of records to be returned by this query.
func (msq *MagicSchoolQuery) Limit(limit int) *MagicSchoolQuery {
	msq.ctx.Limit = &limit
	return msq
}

// Offset to start from.
func (msq *MagicSchoolQuery) Offset(offset int) *MagicSchoolQuery {
	msq.ctx.Offset = &offset
	return msq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (msq *MagicSchoolQuery) Unique(unique bool) *MagicSchoolQuery {
	msq.ctx.Unique = &unique
	return msq
}

// Order specifies how the records should be ordered.
func (msq *MagicSchoolQuery) Order(o ...magicschool.OrderOption) *MagicSchoolQuery {
	msq.order = append(msq.order, o...)
	return msq
}

// First returns the first MagicSchool entity from the query.
// Returns a *NotFoundError when no MagicSchool was found.
func (msq *MagicSchoolQuery) First(ctx context.Context) (*MagicSchool, error) {
	nodes, err := msq.Limit(1).All(setContextOp(ctx, msq.ctx, ent.OpQueryFirst))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{magicschool.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (msq *MagicSchoolQuery) FirstX(ctx context.Context) *MagicSchool {
	node, err := msq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first MagicSchool ID from the query.
// Returns a *NotFoundError when no MagicSchool ID was found.
func (msq *MagicSchoolQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(1).IDs(setContextOp(ctx, msq.ctx, ent.OpQueryFirstID)); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{magicschool.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (msq *MagicSchoolQuery) FirstIDX(ctx context.Context) int {
	id, err := msq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single MagicSchool entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one MagicSchool entity is found.
// Returns a *NotFoundError when no MagicSchool entities are found.
func (msq *MagicSchoolQuery) Only(ctx context.Context) (*MagicSchool, error) {
	nodes, err := msq.Limit(2).All(setContextOp(ctx, msq.ctx, ent.OpQueryOnly))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{magicschool.Label}
	default:
		return nil, &NotSingularError{magicschool.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (msq *MagicSchoolQuery) OnlyX(ctx context.Context) *MagicSchool {
	node, err := msq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only MagicSchool ID in the query.
// Returns a *NotSingularError when more than one MagicSchool ID is found.
// Returns a *NotFoundError when no entities are found.
func (msq *MagicSchoolQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = msq.Limit(2).IDs(setContextOp(ctx, msq.ctx, ent.OpQueryOnlyID)); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{magicschool.Label}
	default:
		err = &NotSingularError{magicschool.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (msq *MagicSchoolQuery) OnlyIDX(ctx context.Context) int {
	id, err := msq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of MagicSchools.
func (msq *MagicSchoolQuery) All(ctx context.Context) ([]*MagicSchool, error) {
	ctx = setContextOp(ctx, msq.ctx, ent.OpQueryAll)
	if err := msq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*MagicSchool, *MagicSchoolQuery]()
	return withInterceptors[[]*MagicSchool](ctx, msq, qr, msq.inters)
}

// AllX is like All, but panics if an error occurs.
func (msq *MagicSchoolQuery) AllX(ctx context.Context) []*MagicSchool {
	nodes, err := msq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of MagicSchool IDs.
func (msq *MagicSchoolQuery) IDs(ctx context.Context) (ids []int, err error) {
	if msq.ctx.Unique == nil && msq.path != nil {
		msq.Unique(true)
	}
	ctx = setContextOp(ctx, msq.ctx, ent.OpQueryIDs)
	if err = msq.Select(magicschool.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (msq *MagicSchoolQuery) IDsX(ctx context.Context) []int {
	ids, err := msq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (msq *MagicSchoolQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, msq.ctx, ent.OpQueryCount)
	if err := msq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, msq, querierCount[*MagicSchoolQuery](), msq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (msq *MagicSchoolQuery) CountX(ctx context.Context) int {
	count, err := msq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (msq *MagicSchoolQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, msq.ctx, ent.OpQueryExist)
	switch _, err := msq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (msq *MagicSchoolQuery) ExistX(ctx context.Context) bool {
	exist, err := msq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the MagicSchoolQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (msq *MagicSchoolQuery) Clone() *MagicSchoolQuery {
	if msq == nil {
		return nil
	}
	return &MagicSchoolQuery{
		config:     msq.config,
		ctx:        msq.ctx.Clone(),
		order:      append([]magicschool.OrderOption{}, msq.order...),
		inters:     append([]Interceptor{}, msq.inters...),
		predicates: append([]predicate.MagicSchool{}, msq.predicates...),
		// clone intermediate query.
		sql:  msq.sql.Clone(),
		path: msq.path,
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
//	client.MagicSchool.Query().
//		GroupBy(magicschool.FieldIndx).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (msq *MagicSchoolQuery) GroupBy(field string, fields ...string) *MagicSchoolGroupBy {
	msq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &MagicSchoolGroupBy{build: msq}
	grbuild.flds = &msq.ctx.Fields
	grbuild.label = magicschool.Label
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
//	client.MagicSchool.Query().
//		Select(magicschool.FieldIndx).
//		Scan(ctx, &v)
func (msq *MagicSchoolQuery) Select(fields ...string) *MagicSchoolSelect {
	msq.ctx.Fields = append(msq.ctx.Fields, fields...)
	sbuild := &MagicSchoolSelect{MagicSchoolQuery: msq}
	sbuild.label = magicschool.Label
	sbuild.flds, sbuild.scan = &msq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a MagicSchoolSelect configured with the given aggregations.
func (msq *MagicSchoolQuery) Aggregate(fns ...AggregateFunc) *MagicSchoolSelect {
	return msq.Select().Aggregate(fns...)
}

func (msq *MagicSchoolQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range msq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, msq); err != nil {
				return err
			}
		}
	}
	for _, f := range msq.ctx.Fields {
		if !magicschool.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if msq.path != nil {
		prev, err := msq.path(ctx)
		if err != nil {
			return err
		}
		msq.sql = prev
	}
	return nil
}

func (msq *MagicSchoolQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*MagicSchool, error) {
	var (
		nodes = []*MagicSchool{}
		_spec = msq.querySpec()
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*MagicSchool).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &MagicSchool{config: msq.config}
		nodes = append(nodes, node)
		return node.assignValues(columns, values)
	}
	if len(msq.modifiers) > 0 {
		_spec.Modifiers = msq.modifiers
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, msq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	for i := range msq.loadTotal {
		if err := msq.loadTotal[i](ctx, nodes); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (msq *MagicSchoolQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := msq.querySpec()
	if len(msq.modifiers) > 0 {
		_spec.Modifiers = msq.modifiers
	}
	_spec.Node.Columns = msq.ctx.Fields
	if len(msq.ctx.Fields) > 0 {
		_spec.Unique = msq.ctx.Unique != nil && *msq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, msq.driver, _spec)
}

func (msq *MagicSchoolQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(magicschool.Table, magicschool.Columns, sqlgraph.NewFieldSpec(magicschool.FieldID, field.TypeInt))
	_spec.From = msq.sql
	if unique := msq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if msq.path != nil {
		_spec.Unique = true
	}
	if fields := msq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, magicschool.FieldID)
		for i := range fields {
			if fields[i] != magicschool.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := msq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := msq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := msq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := msq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (msq *MagicSchoolQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(msq.driver.Dialect())
	t1 := builder.Table(magicschool.Table)
	columns := msq.ctx.Fields
	if len(columns) == 0 {
		columns = magicschool.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if msq.sql != nil {
		selector = msq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if msq.ctx.Unique != nil && *msq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range msq.predicates {
		p(selector)
	}
	for _, p := range msq.order {
		p(selector)
	}
	if offset := msq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := msq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// MagicSchoolGroupBy is the group-by builder for MagicSchool entities.
type MagicSchoolGroupBy struct {
	selector
	build *MagicSchoolQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (msgb *MagicSchoolGroupBy) Aggregate(fns ...AggregateFunc) *MagicSchoolGroupBy {
	msgb.fns = append(msgb.fns, fns...)
	return msgb
}

// Scan applies the selector query and scans the result into the given value.
func (msgb *MagicSchoolGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, msgb.build.ctx, ent.OpQueryGroupBy)
	if err := msgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MagicSchoolQuery, *MagicSchoolGroupBy](ctx, msgb.build, msgb, msgb.build.inters, v)
}

func (msgb *MagicSchoolGroupBy) sqlScan(ctx context.Context, root *MagicSchoolQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(msgb.fns))
	for _, fn := range msgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*msgb.flds)+len(msgb.fns))
		for _, f := range *msgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*msgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := msgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// MagicSchoolSelect is the builder for selecting fields of MagicSchool entities.
type MagicSchoolSelect struct {
	*MagicSchoolQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (mss *MagicSchoolSelect) Aggregate(fns ...AggregateFunc) *MagicSchoolSelect {
	mss.fns = append(mss.fns, fns...)
	return mss
}

// Scan applies the selector query and scans the result into the given value.
func (mss *MagicSchoolSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, mss.ctx, ent.OpQuerySelect)
	if err := mss.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*MagicSchoolQuery, *MagicSchoolSelect](ctx, mss.MagicSchoolQuery, mss, mss.inters, v)
}

func (mss *MagicSchoolSelect) sqlScan(ctx context.Context, root *MagicSchoolQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(mss.fns))
	for _, fn := range mss.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*mss.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := mss.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}
