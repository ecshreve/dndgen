// Code generated by ent, DO NOT EDIT.

package ent

import (
	"context"
	"database/sql/driver"
	"fmt"
	"math"

	"entgo.io/ent/dialect/sql"
	"entgo.io/ent/dialect/sql/sqlgraph"
	"entgo.io/ent/schema/field"
	"github.com/ecshreve/dndgen/ent/abilityscore"
	"github.com/ecshreve/dndgen/ent/predicate"
	"github.com/ecshreve/dndgen/ent/prerequisite"
)

// PrerequisiteQuery is the builder for querying Prerequisite entities.
type PrerequisiteQuery struct {
	config
	ctx              *QueryContext
	order            []prerequisite.OrderOption
	inters           []Interceptor
	predicates       []predicate.Prerequisite
	withAbilityScore *AbilityScoreQuery
	// intermediate query (i.e. traversal path).
	sql  *sql.Selector
	path func(context.Context) (*sql.Selector, error)
}

// Where adds a new predicate for the PrerequisiteQuery builder.
func (pq *PrerequisiteQuery) Where(ps ...predicate.Prerequisite) *PrerequisiteQuery {
	pq.predicates = append(pq.predicates, ps...)
	return pq
}

// Limit the number of records to be returned by this query.
func (pq *PrerequisiteQuery) Limit(limit int) *PrerequisiteQuery {
	pq.ctx.Limit = &limit
	return pq
}

// Offset to start from.
func (pq *PrerequisiteQuery) Offset(offset int) *PrerequisiteQuery {
	pq.ctx.Offset = &offset
	return pq
}

// Unique configures the query builder to filter duplicate records on query.
// By default, unique is set to true, and can be disabled using this method.
func (pq *PrerequisiteQuery) Unique(unique bool) *PrerequisiteQuery {
	pq.ctx.Unique = &unique
	return pq
}

// Order specifies how the records should be ordered.
func (pq *PrerequisiteQuery) Order(o ...prerequisite.OrderOption) *PrerequisiteQuery {
	pq.order = append(pq.order, o...)
	return pq
}

// QueryAbilityScore chains the current query on the "ability_score" edge.
func (pq *PrerequisiteQuery) QueryAbilityScore() *AbilityScoreQuery {
	query := (&AbilityScoreClient{config: pq.config}).Query()
	query.path = func(ctx context.Context) (fromU *sql.Selector, err error) {
		if err := pq.prepareQuery(ctx); err != nil {
			return nil, err
		}
		selector := pq.sqlQuery(ctx)
		if err := selector.Err(); err != nil {
			return nil, err
		}
		step := sqlgraph.NewStep(
			sqlgraph.From(prerequisite.Table, prerequisite.FieldID, selector),
			sqlgraph.To(abilityscore.Table, abilityscore.FieldID),
			sqlgraph.Edge(sqlgraph.O2M, false, prerequisite.AbilityScoreTable, prerequisite.AbilityScoreColumn),
		)
		fromU = sqlgraph.SetNeighbors(pq.driver.Dialect(), step)
		return fromU, nil
	}
	return query
}

// First returns the first Prerequisite entity from the query.
// Returns a *NotFoundError when no Prerequisite was found.
func (pq *PrerequisiteQuery) First(ctx context.Context) (*Prerequisite, error) {
	nodes, err := pq.Limit(1).All(setContextOp(ctx, pq.ctx, "First"))
	if err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nil, &NotFoundError{prerequisite.Label}
	}
	return nodes[0], nil
}

// FirstX is like First, but panics if an error occurs.
func (pq *PrerequisiteQuery) FirstX(ctx context.Context) *Prerequisite {
	node, err := pq.First(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return node
}

// FirstID returns the first Prerequisite ID from the query.
// Returns a *NotFoundError when no Prerequisite ID was found.
func (pq *PrerequisiteQuery) FirstID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(1).IDs(setContextOp(ctx, pq.ctx, "FirstID")); err != nil {
		return
	}
	if len(ids) == 0 {
		err = &NotFoundError{prerequisite.Label}
		return
	}
	return ids[0], nil
}

// FirstIDX is like FirstID, but panics if an error occurs.
func (pq *PrerequisiteQuery) FirstIDX(ctx context.Context) int {
	id, err := pq.FirstID(ctx)
	if err != nil && !IsNotFound(err) {
		panic(err)
	}
	return id
}

// Only returns a single Prerequisite entity found by the query, ensuring it only returns one.
// Returns a *NotSingularError when more than one Prerequisite entity is found.
// Returns a *NotFoundError when no Prerequisite entities are found.
func (pq *PrerequisiteQuery) Only(ctx context.Context) (*Prerequisite, error) {
	nodes, err := pq.Limit(2).All(setContextOp(ctx, pq.ctx, "Only"))
	if err != nil {
		return nil, err
	}
	switch len(nodes) {
	case 1:
		return nodes[0], nil
	case 0:
		return nil, &NotFoundError{prerequisite.Label}
	default:
		return nil, &NotSingularError{prerequisite.Label}
	}
}

// OnlyX is like Only, but panics if an error occurs.
func (pq *PrerequisiteQuery) OnlyX(ctx context.Context) *Prerequisite {
	node, err := pq.Only(ctx)
	if err != nil {
		panic(err)
	}
	return node
}

// OnlyID is like Only, but returns the only Prerequisite ID in the query.
// Returns a *NotSingularError when more than one Prerequisite ID is found.
// Returns a *NotFoundError when no entities are found.
func (pq *PrerequisiteQuery) OnlyID(ctx context.Context) (id int, err error) {
	var ids []int
	if ids, err = pq.Limit(2).IDs(setContextOp(ctx, pq.ctx, "OnlyID")); err != nil {
		return
	}
	switch len(ids) {
	case 1:
		id = ids[0]
	case 0:
		err = &NotFoundError{prerequisite.Label}
	default:
		err = &NotSingularError{prerequisite.Label}
	}
	return
}

// OnlyIDX is like OnlyID, but panics if an error occurs.
func (pq *PrerequisiteQuery) OnlyIDX(ctx context.Context) int {
	id, err := pq.OnlyID(ctx)
	if err != nil {
		panic(err)
	}
	return id
}

// All executes the query and returns a list of Prerequisites.
func (pq *PrerequisiteQuery) All(ctx context.Context) ([]*Prerequisite, error) {
	ctx = setContextOp(ctx, pq.ctx, "All")
	if err := pq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	qr := querierAll[[]*Prerequisite, *PrerequisiteQuery]()
	return withInterceptors[[]*Prerequisite](ctx, pq, qr, pq.inters)
}

// AllX is like All, but panics if an error occurs.
func (pq *PrerequisiteQuery) AllX(ctx context.Context) []*Prerequisite {
	nodes, err := pq.All(ctx)
	if err != nil {
		panic(err)
	}
	return nodes
}

// IDs executes the query and returns a list of Prerequisite IDs.
func (pq *PrerequisiteQuery) IDs(ctx context.Context) (ids []int, err error) {
	if pq.ctx.Unique == nil && pq.path != nil {
		pq.Unique(true)
	}
	ctx = setContextOp(ctx, pq.ctx, "IDs")
	if err = pq.Select(prerequisite.FieldID).Scan(ctx, &ids); err != nil {
		return nil, err
	}
	return ids, nil
}

// IDsX is like IDs, but panics if an error occurs.
func (pq *PrerequisiteQuery) IDsX(ctx context.Context) []int {
	ids, err := pq.IDs(ctx)
	if err != nil {
		panic(err)
	}
	return ids
}

// Count returns the count of the given query.
func (pq *PrerequisiteQuery) Count(ctx context.Context) (int, error) {
	ctx = setContextOp(ctx, pq.ctx, "Count")
	if err := pq.prepareQuery(ctx); err != nil {
		return 0, err
	}
	return withInterceptors[int](ctx, pq, querierCount[*PrerequisiteQuery](), pq.inters)
}

// CountX is like Count, but panics if an error occurs.
func (pq *PrerequisiteQuery) CountX(ctx context.Context) int {
	count, err := pq.Count(ctx)
	if err != nil {
		panic(err)
	}
	return count
}

// Exist returns true if the query has elements in the graph.
func (pq *PrerequisiteQuery) Exist(ctx context.Context) (bool, error) {
	ctx = setContextOp(ctx, pq.ctx, "Exist")
	switch _, err := pq.FirstID(ctx); {
	case IsNotFound(err):
		return false, nil
	case err != nil:
		return false, fmt.Errorf("ent: check existence: %w", err)
	default:
		return true, nil
	}
}

// ExistX is like Exist, but panics if an error occurs.
func (pq *PrerequisiteQuery) ExistX(ctx context.Context) bool {
	exist, err := pq.Exist(ctx)
	if err != nil {
		panic(err)
	}
	return exist
}

// Clone returns a duplicate of the PrerequisiteQuery builder, including all associated steps. It can be
// used to prepare common query builders and use them differently after the clone is made.
func (pq *PrerequisiteQuery) Clone() *PrerequisiteQuery {
	if pq == nil {
		return nil
	}
	return &PrerequisiteQuery{
		config:           pq.config,
		ctx:              pq.ctx.Clone(),
		order:            append([]prerequisite.OrderOption{}, pq.order...),
		inters:           append([]Interceptor{}, pq.inters...),
		predicates:       append([]predicate.Prerequisite{}, pq.predicates...),
		withAbilityScore: pq.withAbilityScore.Clone(),
		// clone intermediate query.
		sql:  pq.sql.Clone(),
		path: pq.path,
	}
}

// WithAbilityScore tells the query-builder to eager-load the nodes that are connected to
// the "ability_score" edge. The optional arguments are used to configure the query builder of the edge.
func (pq *PrerequisiteQuery) WithAbilityScore(opts ...func(*AbilityScoreQuery)) *PrerequisiteQuery {
	query := (&AbilityScoreClient{config: pq.config}).Query()
	for _, opt := range opts {
		opt(query)
	}
	pq.withAbilityScore = query
	return pq
}

// GroupBy is used to group vertices by one or more fields/columns.
// It is often used with aggregate functions, like: count, max, mean, min, sum.
//
// Example:
//
//	var v []struct {
//		Minimum int `json:"minimum,omitempty"`
//		Count int `json:"count,omitempty"`
//	}
//
//	client.Prerequisite.Query().
//		GroupBy(prerequisite.FieldMinimum).
//		Aggregate(ent.Count()).
//		Scan(ctx, &v)
func (pq *PrerequisiteQuery) GroupBy(field string, fields ...string) *PrerequisiteGroupBy {
	pq.ctx.Fields = append([]string{field}, fields...)
	grbuild := &PrerequisiteGroupBy{build: pq}
	grbuild.flds = &pq.ctx.Fields
	grbuild.label = prerequisite.Label
	grbuild.scan = grbuild.Scan
	return grbuild
}

// Select allows the selection one or more fields/columns for the given query,
// instead of selecting all fields in the entity.
//
// Example:
//
//	var v []struct {
//		Minimum int `json:"minimum,omitempty"`
//	}
//
//	client.Prerequisite.Query().
//		Select(prerequisite.FieldMinimum).
//		Scan(ctx, &v)
func (pq *PrerequisiteQuery) Select(fields ...string) *PrerequisiteSelect {
	pq.ctx.Fields = append(pq.ctx.Fields, fields...)
	sbuild := &PrerequisiteSelect{PrerequisiteQuery: pq}
	sbuild.label = prerequisite.Label
	sbuild.flds, sbuild.scan = &pq.ctx.Fields, sbuild.Scan
	return sbuild
}

// Aggregate returns a PrerequisiteSelect configured with the given aggregations.
func (pq *PrerequisiteQuery) Aggregate(fns ...AggregateFunc) *PrerequisiteSelect {
	return pq.Select().Aggregate(fns...)
}

func (pq *PrerequisiteQuery) prepareQuery(ctx context.Context) error {
	for _, inter := range pq.inters {
		if inter == nil {
			return fmt.Errorf("ent: uninitialized interceptor (forgotten import ent/runtime?)")
		}
		if trv, ok := inter.(Traverser); ok {
			if err := trv.Traverse(ctx, pq); err != nil {
				return err
			}
		}
	}
	for _, f := range pq.ctx.Fields {
		if !prerequisite.ValidColumn(f) {
			return &ValidationError{Name: f, err: fmt.Errorf("ent: invalid field %q for query", f)}
		}
	}
	if pq.path != nil {
		prev, err := pq.path(ctx)
		if err != nil {
			return err
		}
		pq.sql = prev
	}
	return nil
}

func (pq *PrerequisiteQuery) sqlAll(ctx context.Context, hooks ...queryHook) ([]*Prerequisite, error) {
	var (
		nodes       = []*Prerequisite{}
		_spec       = pq.querySpec()
		loadedTypes = [1]bool{
			pq.withAbilityScore != nil,
		}
	)
	_spec.ScanValues = func(columns []string) ([]any, error) {
		return (*Prerequisite).scanValues(nil, columns)
	}
	_spec.Assign = func(columns []string, values []any) error {
		node := &Prerequisite{config: pq.config}
		nodes = append(nodes, node)
		node.Edges.loadedTypes = loadedTypes
		return node.assignValues(columns, values)
	}
	for i := range hooks {
		hooks[i](ctx, _spec)
	}
	if err := sqlgraph.QueryNodes(ctx, pq.driver, _spec); err != nil {
		return nil, err
	}
	if len(nodes) == 0 {
		return nodes, nil
	}
	if query := pq.withAbilityScore; query != nil {
		if err := pq.loadAbilityScore(ctx, query, nodes,
			func(n *Prerequisite) { n.Edges.AbilityScore = []*AbilityScore{} },
			func(n *Prerequisite, e *AbilityScore) { n.Edges.AbilityScore = append(n.Edges.AbilityScore, e) }); err != nil {
			return nil, err
		}
	}
	return nodes, nil
}

func (pq *PrerequisiteQuery) loadAbilityScore(ctx context.Context, query *AbilityScoreQuery, nodes []*Prerequisite, init func(*Prerequisite), assign func(*Prerequisite, *AbilityScore)) error {
	fks := make([]driver.Value, 0, len(nodes))
	nodeids := make(map[int]*Prerequisite)
	for i := range nodes {
		fks = append(fks, nodes[i].ID)
		nodeids[nodes[i].ID] = nodes[i]
		if init != nil {
			init(nodes[i])
		}
	}
	query.withFKs = true
	query.Where(predicate.AbilityScore(func(s *sql.Selector) {
		s.Where(sql.InValues(s.C(prerequisite.AbilityScoreColumn), fks...))
	}))
	neighbors, err := query.All(ctx)
	if err != nil {
		return err
	}
	for _, n := range neighbors {
		fk := n.prerequisite_ability_score
		if fk == nil {
			return fmt.Errorf(`foreign-key "prerequisite_ability_score" is nil for node %v`, n.ID)
		}
		node, ok := nodeids[*fk]
		if !ok {
			return fmt.Errorf(`unexpected referenced foreign-key "prerequisite_ability_score" returned %v for node %v`, *fk, n.ID)
		}
		assign(node, n)
	}
	return nil
}

func (pq *PrerequisiteQuery) sqlCount(ctx context.Context) (int, error) {
	_spec := pq.querySpec()
	_spec.Node.Columns = pq.ctx.Fields
	if len(pq.ctx.Fields) > 0 {
		_spec.Unique = pq.ctx.Unique != nil && *pq.ctx.Unique
	}
	return sqlgraph.CountNodes(ctx, pq.driver, _spec)
}

func (pq *PrerequisiteQuery) querySpec() *sqlgraph.QuerySpec {
	_spec := sqlgraph.NewQuerySpec(prerequisite.Table, prerequisite.Columns, sqlgraph.NewFieldSpec(prerequisite.FieldID, field.TypeInt))
	_spec.From = pq.sql
	if unique := pq.ctx.Unique; unique != nil {
		_spec.Unique = *unique
	} else if pq.path != nil {
		_spec.Unique = true
	}
	if fields := pq.ctx.Fields; len(fields) > 0 {
		_spec.Node.Columns = make([]string, 0, len(fields))
		_spec.Node.Columns = append(_spec.Node.Columns, prerequisite.FieldID)
		for i := range fields {
			if fields[i] != prerequisite.FieldID {
				_spec.Node.Columns = append(_spec.Node.Columns, fields[i])
			}
		}
	}
	if ps := pq.predicates; len(ps) > 0 {
		_spec.Predicate = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	if limit := pq.ctx.Limit; limit != nil {
		_spec.Limit = *limit
	}
	if offset := pq.ctx.Offset; offset != nil {
		_spec.Offset = *offset
	}
	if ps := pq.order; len(ps) > 0 {
		_spec.Order = func(selector *sql.Selector) {
			for i := range ps {
				ps[i](selector)
			}
		}
	}
	return _spec
}

func (pq *PrerequisiteQuery) sqlQuery(ctx context.Context) *sql.Selector {
	builder := sql.Dialect(pq.driver.Dialect())
	t1 := builder.Table(prerequisite.Table)
	columns := pq.ctx.Fields
	if len(columns) == 0 {
		columns = prerequisite.Columns
	}
	selector := builder.Select(t1.Columns(columns...)...).From(t1)
	if pq.sql != nil {
		selector = pq.sql
		selector.Select(selector.Columns(columns...)...)
	}
	if pq.ctx.Unique != nil && *pq.ctx.Unique {
		selector.Distinct()
	}
	for _, p := range pq.predicates {
		p(selector)
	}
	for _, p := range pq.order {
		p(selector)
	}
	if offset := pq.ctx.Offset; offset != nil {
		// limit is mandatory for offset clause. We start
		// with default value, and override it below if needed.
		selector.Offset(*offset).Limit(math.MaxInt32)
	}
	if limit := pq.ctx.Limit; limit != nil {
		selector.Limit(*limit)
	}
	return selector
}

// PrerequisiteGroupBy is the group-by builder for Prerequisite entities.
type PrerequisiteGroupBy struct {
	selector
	build *PrerequisiteQuery
}

// Aggregate adds the given aggregation functions to the group-by query.
func (pgb *PrerequisiteGroupBy) Aggregate(fns ...AggregateFunc) *PrerequisiteGroupBy {
	pgb.fns = append(pgb.fns, fns...)
	return pgb
}

// Scan applies the selector query and scans the result into the given value.
func (pgb *PrerequisiteGroupBy) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, pgb.build.ctx, "GroupBy")
	if err := pgb.build.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrerequisiteQuery, *PrerequisiteGroupBy](ctx, pgb.build, pgb, pgb.build.inters, v)
}

func (pgb *PrerequisiteGroupBy) sqlScan(ctx context.Context, root *PrerequisiteQuery, v any) error {
	selector := root.sqlQuery(ctx).Select()
	aggregation := make([]string, 0, len(pgb.fns))
	for _, fn := range pgb.fns {
		aggregation = append(aggregation, fn(selector))
	}
	if len(selector.SelectedColumns()) == 0 {
		columns := make([]string, 0, len(*pgb.flds)+len(pgb.fns))
		for _, f := range *pgb.flds {
			columns = append(columns, selector.C(f))
		}
		columns = append(columns, aggregation...)
		selector.Select(columns...)
	}
	selector.GroupBy(selector.Columns(*pgb.flds...)...)
	if err := selector.Err(); err != nil {
		return err
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := pgb.build.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}

// PrerequisiteSelect is the builder for selecting fields of Prerequisite entities.
type PrerequisiteSelect struct {
	*PrerequisiteQuery
	selector
}

// Aggregate adds the given aggregation functions to the selector query.
func (ps *PrerequisiteSelect) Aggregate(fns ...AggregateFunc) *PrerequisiteSelect {
	ps.fns = append(ps.fns, fns...)
	return ps
}

// Scan applies the selector query and scans the result into the given value.
func (ps *PrerequisiteSelect) Scan(ctx context.Context, v any) error {
	ctx = setContextOp(ctx, ps.ctx, "Select")
	if err := ps.prepareQuery(ctx); err != nil {
		return err
	}
	return scanWithInterceptors[*PrerequisiteQuery, *PrerequisiteSelect](ctx, ps.PrerequisiteQuery, ps, ps.inters, v)
}

func (ps *PrerequisiteSelect) sqlScan(ctx context.Context, root *PrerequisiteQuery, v any) error {
	selector := root.sqlQuery(ctx)
	aggregation := make([]string, 0, len(ps.fns))
	for _, fn := range ps.fns {
		aggregation = append(aggregation, fn(selector))
	}
	switch n := len(*ps.selector.flds); {
	case n == 0 && len(aggregation) > 0:
		selector.Select(aggregation...)
	case n != 0 && len(aggregation) > 0:
		selector.AppendSelect(aggregation...)
	}
	rows := &sql.Rows{}
	query, args := selector.Query()
	if err := ps.driver.Query(ctx, query, args, rows); err != nil {
		return err
	}
	defer rows.Close()
	return sql.ScanSlice(rows, v)
}